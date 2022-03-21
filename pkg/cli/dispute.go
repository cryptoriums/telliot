// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	ethereum_p "github.com/cryptoriums/packages/ethereum"
	"github.com/cryptoriums/packages/logging"
	"github.com/cryptoriums/packages/math"
	math_t "github.com/cryptoriums/packages/math"
	"github.com/cryptoriums/telliot/pkg/aggregator"
	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/db"
	psr_tellor "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/console/prompt"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/storage"
)

type DispID struct {
	DisputeID int64 `required:"" help:"the dispute id"`
}
type DispIDOptional struct {
	DisputeID int64 `help:"the dispute id"`
}

type LookBck struct {
	LookBack time.Duration `default:"120h" help:"how far to lookback, the default only few days since disputes can be voted only for 2 days."`
}

type NewDisputeCmd struct {
	GasAccount
	NoChks
	DataID    int64 `required:""  help:"the request id to dispute"`
	Timestamp int64 `required:""  help:"the submitted timestamp to dispute"`
}

func (self *NewDisputeCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, client, master, oracle, govern, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}
	defer client.Close()

	if self.Timestamp == 0 {
		return errors.New("timestamp can't be zero")
	}

	tsTime := time.Unix(self.Timestamp, 0)
	if time.Since(tsTime) < 0 {
		return errors.New("timestamp can't be in the future")
	}

	account, err := ethereum_p.GetAccountByPubAddress(logger, self.Account, cfg.EnvVars)
	if err != nil {
		return err
	}

	psr, err := psr_tellor.PsrByID(self.DataID)
	if err != nil {
		return errors.Wrap(err, "get psr")
	}

	queryID := psr_tellor.IntToQueryID(self.DataID)

	if psr.Inactive {
		return errors.Wrap(err, "psr is inactive so opening a dispute is pointless")
	}

	disputeCost, err := govern.DisputeFee(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "get dispute cost")
	}

	if !self.NoChecks {
		val, err := oracle.GetValueByTimestamp(&bind.CallOpts{Context: ctx}, queryID, big.NewInt(self.Timestamp))
		if err != nil {
			return errors.Wrap(err, "getting the val to dispute")
		}
		if len(val) == 0 {
			return errors.Errorf("value with timestamp:%v doesn't exist. it is possible that it has been removed by another dispute", self.Timestamp)
		}

		balance, err := master.BalanceOf(&bind.CallOpts{Context: ctx}, account.PublicKey)
		if err != nil {
			return errors.Wrap(err, "fetch balance")
		}

		if balance.Cmp(disputeCost) < 0 {
			return errors.Errorf("insufficient balance TRB actual: %v, TRB required:%v)",
				math.BigIntToFloatDiv(balance, params.Ether),
				math.BigIntToFloatDiv(disputeCost, params.Ether))
		}

		reporter, err := oracle.GetReporterByTimestamp(&bind.CallOpts{Context: ctx}, queryID, big.NewInt(self.Timestamp))
		if err != nil {
			return errors.Wrap(err, "getting submit reporter")
		}

		status, _, err := master.GetStakerInfo(&bind.CallOpts{Context: ctx}, reporter)
		if err != nil {
			return errors.Wrap(err, "getting reporter status")
		}
		level.Info(logger).Log("msg", "disputed reporter status", "addr", reporter.Hex()[:8], "status", contracts.ReporterStatusName(status.Int64()))
		if status.Int64() != 1 {
			confirmed, err := prompt.Stdin.PromptConfirm("Disputed reporter is not in staked status: " + contracts.ReporterStatusName(status.Int64()) + ". confirm you want to continue despite its status:")
			if err != nil || !confirmed {
				return errors.New("canceled")
			}
		}
	}

	confirmed, err := prompt.Stdin.PromptConfirm("Dispute fee is TRB: " + fmt.Sprintf("%.2f", math_t.BigIntToFloatDiv(disputeCost, params.Ether)) + " confirm to continue:")
	if err != nil || !confirmed {
		return errors.New("canceled")
	}

	opts, err := ethereum_p.NewTxOpts(ctx, client, account, self.GasPrice, contracts.DisputeNewGasLimit)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}

	tx, err := govern.BeginDispute(opts, queryID, big.NewInt(self.Timestamp))
	if err != nil {
		return errors.Wrap(err, "send dispute txn")
	}
	level.Info(logger).Log("msg", "dispute tx created",
		"queryID", common.Bytes2Hex(queryID[:]),
		"ts", self.Timestamp,
		"tx", tx.Hash())
	return nil
}

type VoteCmd struct {
	GasAccount
	DispID
	NoChks
	Support bool `help:"true or false"`
	Against bool `help:"true or false"`
	Invalid bool `help:"true or false"`
}

func (self *VoteCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	if !self.Support && !self.Against && !self.Invalid {
		return errors.New("need to either support, against or invalid for a vote")
	}
	cfg, client, master, _, govern, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}
	defer client.Close()

	account, err := ethereum_p.GetAccountByPubAddress(logger, self.Account, cfg.EnvVars)
	if err != nil {
		return err
	}

	dispute, err := contracts.GetDisputeInfo(ctx, big.NewInt(self.DisputeID), govern)
	if err != nil {
		return errors.Wrap(err, "getting dispute details")
	}
	if dispute.Executed {
		return errors.New("dispute already executed")
	}

	if !self.NoChecks {
		voted, err := govern.DidVote(&bind.CallOpts{Context: ctx}, big.NewInt(self.DisputeID), account.PublicKey)
		if err != nil {
			return errors.Wrapf(err, "check if you've already voted")
		}
		if voted {
			return errors.New("you have already voted on this dispute")
		}

		status, _, err := master.GetStakerInfo(&bind.CallOpts{Context: ctx}, account.PublicKey)
		if err != nil {
			return errors.Wrap(err, "getting reporter status")
		}
		if status.Int64() == contracts.StatusOnDispute {
			return errors.New("address is in dispute so can't vote")
		}
	}

	opts, err := ethereum_p.NewTxOpts(ctx, client, account, self.GasPrice, contracts.VoteGasUSage)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}
	tx, err := govern.Vote(opts, big.NewInt(self.DisputeID), self.Support, self.Invalid)
	if err != nil {
		return errors.Wrapf(err, "submit vote transaction")
	}

	level.Info(logger).Log("msg", "vote submitted", "disputeID", self.DisputeID, "invalid", self.Invalid, "against", self.Against, "supports", self.Support, "tx", tx.Hash())
	return nil
}

type TallyCmd struct {
	Gas
	DispIDOptional
	All bool
}

func (self *TallyCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, client, _, _, govern, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}
	defer client.Close()

	var disputes []*contracts.DisputeLog

	if self.All {
		disputes, err = contracts.GetDisputeLogs(ctx, logger, client, govern, 30*24*time.Hour)
		if err != nil {
			return errors.Wrap(err, "get dispute logs")
		}

	} else {
		if self.DisputeID == 0 {
			return errors.New("dispute ID is empty")
		}
		dispute, err := contracts.GetDisputeInfo(ctx, big.NewInt(self.DisputeID), govern)
		if err != nil {
			return errors.Wrap(err, "get dispute info")
		}
		disputes = append(disputes, dispute)
	}

	accounts, err := ethereum_p.GetAccounts(logger, cfg.EnvVars)
	if err != nil {
		return err
	}

	for _, dispute := range disputes {
		if dispute.TallyTs > 0 {
			level.Info(logger).Log("msg", "dispute already tallied", "id", dispute.ID)
			continue
		}

		if time.Now().After(dispute.VoteEnds) {
			level.Info(logger).Log("msg", "dispute voting period hasn't passed", "id", dispute.ID, "timeLeft", time.Since(dispute.VoteEnds))
			continue
		}

		opts, err := ethereum_p.NewTxOpts(ctx, client, accounts[0], self.GasPrice, contracts.VoteTallyGasLimit)
		if err != nil {
			return errors.Wrapf(err, "prepare ethereum transaction")
		}

		tx, err := govern.TallyVotes(opts, big.NewInt(dispute.ID))
		if err != nil {
			return errors.Wrapf(err, "run tally votes")
		}

		level.Info(logger).Log("msg", "tally votes submitted", "tx", tx.Hash().Hex())
	}
	return nil
}

type TallyListCmd struct {
	LookBck
}

func (self *TallyListCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	_, client, _, _, govern, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}
	defer client.Close()

	logs, err := contracts.GetTallyLogs(ctx, client, govern, self.LookBack)
	if err != nil {
		return errors.Wrap(err, "GetTallyLogs")
	}

	for _, log := range logs {
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Println()
		fmt.Fprintln(w, "TxHash:\t", log.Raw.TxHash, "\t")
		fmt.Fprintln(w, "ID:\t", log.DisputeId, "\t")
		fmt.Fprintln(w, "Result:\t", log.Result, "\t")
		w.Flush()
	}

	return nil
}

type ExecuteCmd struct {
	NoChks
	Gas
	DispIDOptional
	All bool
}

func (self *ExecuteCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, client, _, _, govern, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}
	defer client.Close()

	accounts, err := ethereum_p.GetAccounts(logger, cfg.EnvVars)
	if err != nil {
		return err
	}

	var disputes []*contracts.DisputeLog

	if self.All {
		disputes, err = contracts.GetDisputeLogs(ctx, logger, client, govern, 30*24*time.Hour)
		if err != nil {
			return errors.Wrap(err, "get dispute logs")
		}

	} else {
		if self.DisputeID == 0 {
			return errors.New("dispute ID is empty")
		}
		dispute, err := contracts.GetDisputeInfo(ctx, big.NewInt(self.DisputeID), govern)
		if err != nil {
			return errors.Wrap(err, "get dispute info")
		}
		disputes = append(disputes, dispute)
	}

	for _, dispute := range disputes {
		if dispute.TallyTs == 0 {
			level.Info(logger).Log("msg", "dispute needs to be tallied first", "id", dispute.ID)
			continue
		}
		if time.Now().After(dispute.VoteEnds) {
			level.Info(logger).Log("msg", "voting period haven't passed", "id", dispute.ID, "ends", time.Since(dispute.VoteEnds))
		}
		if dispute.Executed {
			level.Info(logger).Log("msg", "dispute already paid out", "id", dispute.ID)
			continue
		}

		if time.Now().After(dispute.ExecuteTimeLock) {
			level.Info(logger).Log("msg", "not enough time has passed after tallying", "id", dispute.ID, "time", time.Since(dispute.ExecuteTimeLock))
			continue
		}

		opts, err := ethereum_p.NewTxOpts(ctx, client, accounts[0], self.GasPrice, contracts.VoteExecuteGasLimit)
		if err != nil {
			return errors.Wrapf(err, "prepare ethereum transaction")
		}

		tx, err := govern.ExecuteVote(opts, big.NewInt(dispute.ID))
		if err != nil {
			return errors.Wrapf(err, "run vote execute fee")
		}

		level.Info(logger).Log("msg", "vote execute submitted", "id", dispute.ID, "tx", tx.Hash().Hex())
	}

	return nil
}

type ListCmd struct {
	ShowClosed bool `help:"also show executed disputes"`
	LookBck
}

func (self *ListCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, client, _, _, govern, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}
	defer client.Close()

	// Open the TSDB database.
	var querable storage.SampleAndChunkQueryable
	if cfg.Db.RemoteHost != "" {
		querable, err = db.NewRemoteDB(cfg.Db)
		if err != nil {
			return errors.Wrap(err, "opening remote tsdb DB")
		}
	} else {
		level.Warn(logger).Log("msg", "no remote DB set so will not suggest values")
	}

	aggregator, err := aggregator.New(ctx, logger, cfg.Aggregator, querable)
	if err != nil {
		return errors.Wrap(err, "creating aggregator")
	}

	psrT := psr_tellor.New(logger, cfg.PsrTellor, aggregator)

	disputes, err := contracts.GetDisputeLogs(ctx, logger, client, govern, self.LookBack)
	if err != nil {
		return errors.Wrap(err, "GetDisputeLogs")
	}

	level.Info(logger).Log("msg", "disputes count", "lookBackPeriod", self.LookBack, "count", len(disputes))

	for _, dispute := range disputes {
		if !self.ShowClosed && dispute.StatusID != contracts.VoteStatusOpen && dispute.StatusID != contracts.VoteStatusTallied {
			continue
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

		psr, ok := psr_tellor.Psrs[dispute.QueryID]
		if !ok {
			level.Error(logger).Log("msg", "getting psr", "queryID", common.Bytes2Hex(dispute.QueryID[:]), "err", err)
			continue
		}
		var suggested float64
		if cfg.Db.RemoteHost != "" {
			suggested, err = psrT.GetValue(dispute.QueryID, dispute.DataTime)
			if err != nil {
				level.Error(logger).Log("msg", "look up recommended value", "id", dispute.ID, "err", err)
			}
		}

		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Println()
		fmt.Fprintln(w, "ID: \t", dispute.ID, "\t")
		fmt.Fprintln(w, "Executed: \t", dispute.Executed, "\t")
		fmt.Fprintln(w, "Status: \t", dispute.StatusName, "\t")
		fmt.Fprintln(w, "Fee: \t", dispute.Fee, "\t")
		fmt.Fprintln(w, "Vote Ends: \t", -time.Since(dispute.VoteEnds), "\t")
		fmt.Fprintln(w, "Tally Ts: \t", dispute.TallyTs, "\t")
		fmt.Fprintln(w, "Votes Support: \t", dispute.VotesSupport, "\t")
		fmt.Fprintln(w, "Votes Against: \t", dispute.VotesAgainst, "\t")
		fmt.Fprintln(w, "Votes Invalid: \t", dispute.VotesInvalid, "\t")
		fmt.Fprintln(w, "Pairs: \t", psr.Pair, "\t")
		fmt.Fprintln(w, "Ts: \t", strconv.Itoa(int(dispute.DataTime.Unix()))+" "+dispute.DataTime.Format(logging.DefaultTimeFormat), "\t")
		fmt.Fprintln(w, "Disputer: \t", dispute.Initiator.Hex(), "\t")
		fmt.Fprintln(w, "Reporter: \t", dispute.Reporter.Hex(), "\t")
		fmt.Fprintln(w, "Disputed    Id: \t", dispute.QueryID, "\t")
		fmt.Fprintln(w, "Disputed  Time: \t", dispute.DataTime, "\t")
		fmt.Fprintln(w, "Disputed   Val: \t", fmt.Sprintf("%.6f", dispute.DataVal), "\t")
		fmt.Fprintln(w, "Suggested  Val: \t", fmt.Sprintf("%.6f", suggested/psr.Granularity), "\t")
		fmt.Fprintln(w, "TxHash: \t", dispute.TxHash.Hex(), "\t")
		w.Flush()

		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Printf("\n")
	}

	return nil
}
