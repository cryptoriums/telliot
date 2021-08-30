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

	"github.com/cryptoriums/telliot/pkg/aggregator"
	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/db"
	ethereumT "github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/logging"
	"github.com/cryptoriums/telliot/pkg/math"
	"github.com/cryptoriums/telliot/pkg/psr/tellor"
	psrTellor "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
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
	Slot      int64 `required:""  help:"the reporter index to dispute"`
}

func (self *NewDisputeCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	_, client, contract, err := ConfigClientContract(ctx, logger, cli.Config, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	if self.Slot < 0 || self.Slot > 4 {
		return errors.Errorf("reporter index should be between 0 and 4 (got %v)", self.Slot)
	}

	if self.Timestamp == 0 {
		return errors.New("timestamp can't be zero")
	}

	tsTime := time.Unix(self.Timestamp, 0)
	if time.Since(tsTime) < 0 {
		return errors.New("timestamp can't be in the future")
	}

	account, err := ethereumT.GetAccountByPubAddress(logger, self.Account)
	if err != nil {
		return err
	}

	if !self.NoChecks {
		balance, err := contract.BalanceOf(&bind.CallOpts{Context: ctx}, account.Address)
		if err != nil {
			return errors.Wrap(err, "fetch balance")
		}

		disputeCost, err := contract.GetUintVar(&bind.CallOpts{Context: ctx}, ethereumT.Keccak256([]byte("_DISPUTE_FEE")))
		if err != nil {
			return errors.Wrap(err, "get dispute cost")
		}

		if balance.Cmp(disputeCost) < 0 {
			return errors.Errorf("insufficient balance TRB actual: %v, TRB required:%v)",
				math.BigIntToFloatDiv(balance, params.Ether),
				math.BigIntToFloatDiv(disputeCost, params.Ether))
		}
	}

	opts, err := ethereumT.PrepareEthTransaction(ctx, client, account, self.GasMaxFee, self.GasMaxTip, contracts.BeginDisputeGasUsage)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}

	tx, err := contract.BeginDispute(opts, big.NewInt(self.DataID), big.NewInt(self.Timestamp), big.NewInt(self.Slot))
	if err != nil {
		return errors.Wrap(err, "send dispute txn")
	}
	level.Info(logger).Log("msg", "dispute started",
		"dataID", self.DataID,
		"ts", self.Timestamp,
		"slot", self.Slot,
		"tx", tx.Hash())
	return nil
}

type VoteCmd struct {
	GasAccount
	DispID
	NoChks
	Support bool `required:"" help:"true or false"`
}

func (self *VoteCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	_, client, contract, err := ConfigClientContract(ctx, logger, cli.Config, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	account, err := ethereumT.GetAccountByPubAddress(logger, self.Account)
	if err != nil {
		return err
	}

	dispute, err := contracts.GetDisputeInfo(ctx, big.NewInt(self.DisputeID), contract)
	if err != nil {
		return errors.Wrap(err, "getting dispute details")
	}
	if dispute.Executed {
		return errors.New("dispute already executed")
	}

	if !self.NoChecks {
		voted, err := contract.DidVote(&bind.CallOpts{Context: ctx}, big.NewInt(self.DisputeID), account.Address)
		if err != nil {
			return errors.Wrapf(err, "check if you've already voted")
		}
		if voted {
			level.Info(logger).Log("msg", "you have already voted on this dispute")
			return nil
		}
	}

	opts, err := ethereumT.PrepareEthTransaction(ctx, client, account, self.GasMaxFee, self.GasMaxTip, contracts.VoteGasUSage)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}
	tx, err := contract.Vote(opts, big.NewInt(self.DisputeID), self.Support)
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
	_, client, contract, err := ConfigClientContract(ctx, logger, cli.Config, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	var disputes []*contracts.DisputeLog

	if self.All {
		disputes, err = contracts.GetDisputeLogs(ctx, client, contract, 300*time.Hour)
		if err != nil {
			return errors.Wrap(err, "get dispute logs")
		}

	} else {
		if self.DisputeID == 0 {
			return errors.New("dispute ID is empty")
		}
		dispute, err := contracts.GetDisputeInfo(ctx, big.NewInt(self.DisputeID), contract)
		if err != nil {
			return errors.Wrap(err, "get dispute info")
		}
		disputes = append(disputes, dispute)
	}

	accounts, err := ethereumT.GetAccounts(logger)
	if err != nil {
		return err
	}

	for _, dispute := range disputes {
		if dispute.Executed {
			level.Info(logger).Log("msg", "dipsute already executed", "id", dispute.ID)
			continue
		}

		opts, err := ethereumT.PrepareEthTransaction(ctx, client, accounts[0], self.GasMaxFee, self.GasMaxTip, contracts.TallyGasUsage)
		if err != nil {
			return errors.Wrapf(err, "prepare ethereum transaction")
		}

		tx, err := contract.TallyVotes(opts, big.NewInt(dispute.ID))
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
	_, client, contract, err := ConfigClientContract(ctx, logger, cli.Config, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	logs, err := contracts.GetTallyLogs(ctx, client, contract, self.LookBack)
	if err != nil {
		return errors.Wrap(err, "getting latest disputes")
	}

	for _, log := range logs {
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Println()
		fmt.Fprintln(w, "TxHash:\t", log.Raw.TxHash, "\t")
		fmt.Fprintln(w, "ID:\t", log.DisputeID, "\t")
		fmt.Fprintln(w, "Passed:\t", log.Passed, "\t")
		fmt.Fprintln(w, "Reported:\t", log.ReportedMiner.Hex(), "\t")
		fmt.Fprintln(w, "Reporter:\t", log.ReportingParty.Hex(), "\t")
		w.Flush()
	}

	return nil
}

type UnlockFeeCmd struct {
	NoChks
	Gas
	DispIDOptional
	All bool
}

func (self *UnlockFeeCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	_, client, contract, err := ConfigClientContract(ctx, logger, cli.Config, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	accounts, err := ethereumT.GetAccounts(logger)
	if err != nil {
		return err
	}

	var disputes []*contracts.DisputeLog

	if self.All {
		disputes, err = contracts.GetDisputeLogs(ctx, client, contract, 300*time.Hour)
		if err != nil {
			return errors.Wrap(err, "get dispute logs")
		}

	} else {
		if self.DisputeID == 0 {
			return errors.New("dispute ID is empty")
		}
		dispute, err := contracts.GetDisputeInfo(ctx, big.NewInt(self.DisputeID), contract)
		if err != nil {
			return errors.Wrap(err, "get dispute info")
		}
		disputes = append(disputes, dispute)
	}

	for _, dispute := range disputes {
		if !self.NoChecks {
			tallyTs, err := contract.GetDisputeUintVars(
				&bind.CallOpts{Context: ctx},
				big.NewInt(dispute.ID),
				ethereumT.Keccak256([]byte("_TALLY_DATE")),
			)
			if err != nil {
				return errors.Wrap(err, "get _TALLY_DATE")
			}

			if tallyTs.Int64() == 0 {
				return errors.Errorf("dispute not executed yet")
			}

			timePassed := time.Since(time.Unix(tallyTs.Int64(), 0))

			if timePassed < 24*time.Hour {
				return errors.Errorf("not enough time has passed after tallying - required:24h, current:%v", timePassed)
			}

			paid, err := contract.GetDisputeUintVars(
				&bind.CallOpts{Context: ctx},
				big.NewInt(dispute.ID),
				ethereumT.Keccak256([]byte("_PAID")),
			)
			if err != nil {
				return errors.Wrap(err, "get _PAID")
			}

			if paid.Int64() == 1 {
				return errors.Errorf("dispute already paid out")
			}
		}

		opts, err := ethereumT.PrepareEthTransaction(ctx, client, accounts[0], self.GasMaxFee, self.GasMaxTip, contracts.UnlockFeeGasUsage)
		if err != nil {
			return errors.Wrapf(err, "prepare ethereum transaction")
		}

		tx, err := contract.UnlockDisputeFee(opts, big.NewInt(dispute.ID))
		if err != nil {
			return errors.Wrapf(err, "run unlock fee")
		}

		level.Info(logger).Log("msg", "unlock fee submitted", "tx", tx.Hash().Hex())
	}

	return nil
}

type ListCmd struct {
	ShowClosed bool `help:"also show executed disputes"`
	LookBck
}

func (self *ListCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, client, contract, err := ConfigClientContract(ctx, logger, cli.Config, cli.Contract, contracts.DefaultParams)
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

	psr := psrTellor.New(logger, cfg.PsrTellor, aggregator)

	logs, err := contracts.GetDisputeLogs(ctx, client, contract, self.LookBack)
	if err != nil {
		return errors.Wrap(err, "getting latest disputes")
	}

	level.Info(logger).Log("msg", "disputes count", "lookBackPeriod", self.LookBack, "count", len(logs))

	for _, log := range logs {
		if !self.ShowClosed && log.Executed {
			continue
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Println()
		fmt.Fprintln(w, "ID: \t", log.ID, "\t")
		fmt.Fprintln(w, "Executed: \t", log.Executed, "\t")
		fmt.Fprintln(w, "Passed: \t", log.Passed, "\t")
		fmt.Fprintln(w, "Fee: \t", log.Fee, "\t")
		fmt.Fprintln(w, "Created :\t", log.Created.Format(logging.DefaultTimeFormat), "\t")
		fmt.Fprintln(w, "Ends: \t", -time.Since(log.Ends), "\t")
		fmt.Fprintln(w, "Tally: \t", log.Tally, "\t")
		fmt.Fprintln(w, "Votes: \t", log.Votes, "\t")
		fmt.Fprintln(w, "Pairs: \t", tellor.Psrs[log.DataID].Pair, "\t")
		fmt.Fprintln(w, "DataId: \t", log.DataID, "\t")
		fmt.Fprintln(w, "Ts: \t", strconv.Itoa(int(log.DataTime.Unix()))+" "+log.DataTime.Format(logging.DefaultTimeFormat), "\t")
		fmt.Fprintln(w, "Slot:\t", log.DisputedSlot, "\t")
		fmt.Fprintln(w, "Disputer: \t", log.Disputer.Hex(), "\t")
		fmt.Fprintln(w, "Disputed: \t", log.Disputed.Hex(), "\t")
		fmt.Fprintln(w, "TxHash: \t", log.TxHash.Hex(), "\t")
		w.Flush()

		var suggested float64
		if cfg.Db.RemoteHost != "" {
			suggested, err = psr.GetValue(log.DataID, log.DataTime)
			if err != nil {
				level.Error(logger).Log("msg", "look up recommended value", "id", log.ID, "err", err)
			}
		}

		_submits, err := contract.GetSubmissionsByTimestamp(
			&bind.CallOpts{Context: ctx},
			big.NewInt(log.DataID),
			big.NewInt(log.DataTime.Unix()),
		)
		if err != nil {
			level.Error(logger).Log("msg", "getting all submits", "err", err)
			continue
		}

		for i, submit := range _submits {
			var disputed string
			if i == int(log.DisputedSlot) {
				disputed = "disputed"
			}

			//lint:ignore faillint looks cleaner with print instead of logs
			fmt.Printf("value:%.6f, suggestedValue:%.6f, minerIndex:%v, %v \n",
				float64(submit.Int64())/psrTellor.DefaultGranularity,
				suggested/psrTellor.DefaultGranularity,
				i,
				disputed,
			)
		}
		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Printf("\n\n")
	}

	return nil
}
