// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/cryptoriums/telliot/pkg/aggregator"
	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/db"
	ethereum_t "github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/logging"
	"github.com/cryptoriums/telliot/pkg/math"
	math_t "github.com/cryptoriums/telliot/pkg/math"
	"github.com/cryptoriums/telliot/pkg/prompt"
	psr_tellor "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	_, client, master, oracle, govern, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	if self.Timestamp == 0 {
		return errors.New("timestamp can't be zero")
	}

	tsTime := time.Unix(self.Timestamp, 0)
	if time.Since(tsTime) < 0 {
		return errors.New("timestamp can't be in the future")
	}

	account, err := ethereum_t.GetAccountByPubAddress(logger, self.Account)
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

		balance, err := master.BalanceOf(&bind.CallOpts{Context: ctx}, account.Address)
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
			promptResp, err := prompt.Prompt("Disputed reporter is not in staked status. Press Y to continue despite its status:", false)
			if err == nil && strings.ToLower(promptResp) != "y" {
				return errors.New("canceled")
			}
		}
	}

	promptResp, err := prompt.Prompt("Dispute fee is TRB: "+fmt.Sprintf("%.2f", math_t.BigIntToFloatDiv(disputeCost, params.Ether))+" Press Y to continue:", false)
	if err == nil && strings.ToLower(promptResp) != "y" {
		return errors.New("canceled")
	}

	opts, err := ethereum_t.PrepareTx(ctx, client, account, self.GasPrice, contracts.DisputeNewGasLimit)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}

	tx, err := govern.BeginDispute(opts, queryID, big.NewInt(self.Timestamp))
	if err != nil {
		return errors.Wrap(err, "send dispute txn")
	}
	level.Info(logger).Log("msg", "dispute tx created",
		"queryID", queryID,
		"ts", self.Timestamp,
		"tx", tx.Hash())
	return nil
}

type VoteCmd struct {
	GasAccount
	DispID
	NoChks
	Support bool `help:"true or false"`
	Invalid bool `help:"true or false"`
}

func (self *VoteCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	if !self.Support && !self.Invalid {
		return errors.New("need to either support or invalid for a vote")
	}
	_, client, _, _, govern, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	account, err := ethereum_t.GetAccountByPubAddress(logger, self.Account)
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
		voted, err := govern.DidVote(&bind.CallOpts{Context: ctx}, big.NewInt(self.DisputeID), account.Address)
		if err != nil {
			return errors.Wrapf(err, "check if you've already voted")
		}
		if voted {
			level.Info(logger).Log("msg", "you have already voted on this dispute")
			return nil
		}
	}

	opts, err := ethereum_t.PrepareTx(ctx, client, account, self.GasPrice, contracts.VoteGasUSage)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}
	tx, err := govern.Vote(opts, big.NewInt(self.DisputeID), self.Support, self.Invalid)
	if err != nil {
		return errors.Wrapf(err, "submit vote transaction")
	}

	level.Info(logger).Log("msg", "vote submitted", "disputeID", self.DisputeID, "supports", self.Support, "tx", tx.Hash())
	return nil
}

type TallyCmd struct {
	Gas
	DispIDOptional
	All bool
}

func (self *TallyCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	_, client, _, _, govern, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
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

	accounts, err := ethereum_t.GetAccounts(logger)
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

		opts, err := ethereum_t.PrepareTx(ctx, client, accounts[0], self.GasPrice, contracts.VoteTallyGasLimit)
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

type VoteExecuteCmd struct {
	NoChks
	Gas
	DispIDOptional
	All bool
}

func (self *VoteExecuteCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	_, client, _, _, govern, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	accounts, err := ethereum_t.GetAccounts(logger)
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

		opts, err := ethereum_t.PrepareTx(ctx, client, accounts[0], self.GasPrice, contracts.VoteExecuteGasLimit)
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

	logs, err := contracts.GetDisputeLogs(ctx, logger, client, govern, self.LookBack)
	if err != nil {
		return errors.Wrap(err, "GetDisputeLogs")
	}

	level.Info(logger).Log("msg", "disputes count", "lookBackPeriod", self.LookBack, "count", len(logs))

	for _, log := range logs {
		if !self.ShowClosed && log.Executed {
			continue
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

		psr, ok := psr_tellor.Psrs[log.QueryID]
		if !ok {
			level.Error(logger).Log("msg", "getting psr", "queryID", log.QueryID, "err", err)
			continue
		}
		var suggested float64
		if cfg.Db.RemoteHost != "" {
			suggested, err = psrT.GetValue(log.QueryID, log.DataTime)
			if err != nil {
				level.Error(logger).Log("msg", "look up recommended value", "id", log.ID, "err", err)
			}
		}

		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Println()
		fmt.Fprintln(w, "ID: \t", log.ID, "\t")
		fmt.Fprintln(w, "Executed: \t", log.Executed, "\t")
		fmt.Fprintln(w, "Result: \t", log.ResultName, "\t")
		fmt.Fprintln(w, "Fee: \t", log.Fee, "\t")
		fmt.Fprintln(w, "Vote Ends: \t", -time.Since(log.VoteEnds), "\t")
		fmt.Fprintln(w, "Tally Ts: \t", log.TallyTs, "\t")
		fmt.Fprintln(w, "Votes Support: \t", log.VotesSupport, "\t")
		fmt.Fprintln(w, "Votes Against: \t", log.VotesAgainst, "\t")
		fmt.Fprintln(w, "Votes Invalid: \t", log.VotesInvalid, "\t")
		fmt.Fprintln(w, "Pairs: \t", psr.Pair, "\t")
		fmt.Fprintln(w, "Ts: \t", strconv.Itoa(int(log.DataTime.Unix()))+" "+log.DataTime.Format(logging.DefaultTimeFormat), "\t")
		fmt.Fprintln(w, "Disputer: \t", log.Initiator.Hex(), "\t")
		fmt.Fprintln(w, "Reporter: \t", log.Reporter.Hex(), "\t")
		fmt.Fprintln(w, "Disputed    Id: \t", log.QueryID, "\t")
		fmt.Fprintln(w, "Disputed  Time: \t", log.DataTime, "\t")
		fmt.Fprintln(w, "Disputed   Val: \t", fmt.Sprintf("%.6f", log.DataVal), "\t")
		fmt.Fprintln(w, "Suggested  Val: \t", fmt.Sprintf("%.6f", suggested/psr_tellor.DefaultGranularity), "\t")
		fmt.Fprintln(w, "TxHash: \t", log.TxHash.Hex(), "\t")
		w.Flush()

		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Printf("\n")
	}

	return nil
}
