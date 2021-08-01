// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/tellor-io/telliot/pkg/aggregator"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	tEthereum "github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/math"
	psrTellor "github.com/tellor-io/telliot/pkg/psr/tellor"
)

type disputeID struct {
	DisputeID int64 `arg:"" required:"" help:"the dispute id"`
}

type newDisputeCmd struct {
	cfgGasAddr
	RequestID  int64 `arg:""  help:"the request id to dispute it"`
	Timestamp  int64 `arg:""  help:"the submitted timestamp to dispute"`
	MinerIndex int64 `arg:""  help:"the miner index to dispute"`
}

func (self newDisputeCmd) Run() error {
	logger := logging.NewLogger()
	ctx := context.Background()

	_, err := config.ParseConfig(logger, string(self.Config)) // Load the env file.
	if err != nil {
		return errors.Wrap(err, "creating config")
	}

	client, err := tEthereum.NewClient(ctx, logger)
	if err != nil {
		return errors.Wrap(err, "creating ethereum client")
	}

	account, err := tEthereum.GetAccountByPubAddess(self.Addr)
	if err != nil {
		return err
	}
	contract, err := contracts.NewITellor(client)
	if err != nil {
		return errors.Wrap(err, "create tellor contract instance")
	}

	if self.MinerIndex < 0 || self.MinerIndex > 4 {
		return errors.Errorf("miner index should be between 0 and 4 (got %v)", self.MinerIndex)
	}

	balance, err := contract.BalanceOf(&bind.CallOpts{Context: ctx}, account.Address)
	if err != nil {
		return errors.Wrap(err, "fetch balance")
	}

	disputeCost, err := contract.GetUintVar(&bind.CallOpts{Context: ctx}, tEthereum.Keccak256([]byte("_DISPUTE_FEE")))
	if err != nil {
		return errors.Wrap(err, "get dispute cost")
	}

	if balance.Cmp(disputeCost) < 0 {
		return errors.Errorf("insufficient balance TRB actual: %v, TRB required:%v)",
			math.BigInt18eToFloat(balance),
			math.BigInt18eToFloat(disputeCost))
	}

	var gasPrice *big.Int
	if self.GasPrice > 0 {
		gasPrice = big.NewInt(int64(self.GasPrice) * params.GWei)
	}

	auth, err := tEthereum.PrepareEthTransaction(ctx, client, account, gasPrice)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}

	tx, err := contract.BeginDispute(auth, big.NewInt(self.RequestID), big.NewInt(self.Timestamp), big.NewInt(self.MinerIndex))
	if err != nil {
		return errors.Wrap(err, "send dispute txn")
	}
	level.Info(logger).Log("msg", "dispute started", "tx", tx.Hash())
	return nil
}

type voteCmd struct {
	cfgGasAddr
	disputeID
	Support bool `arg:"" required:"" help:"true or false"`
}

func (self voteCmd) Run() error {
	logger := logging.NewLogger()
	ctx := context.Background()

	_, err := config.ParseConfig(logger, string(self.Config)) // Load the env file.
	if err != nil {
		return errors.Wrap(err, "creating config")
	}

	client, err := tEthereum.NewClient(ctx, logger)
	if err != nil {
		return errors.Wrap(err, "creating ethereum client")
	}

	account, err := tEthereum.GetAccountByPubAddess(self.Addr)
	if err != nil {
		return err
	}
	contract, err := contracts.NewITellor(client)
	if err != nil {
		return errors.Wrap(err, "create tellor contract instance")
	}

	voted, err := contract.DidVote(&bind.CallOpts{Context: ctx}, big.NewInt(self.DisputeID), account.Address)
	if err != nil {
		return errors.Wrapf(err, "check if you've already voted")
	}
	if voted {
		level.Info(logger).Log("msg", "you have already voted on this dispute")
		return nil
	}

	var gasPrice *big.Int
	if self.GasPrice > 0 {
		gasPrice = big.NewInt(int64(self.GasPrice) * params.GWei)
	}

	auth, err := tEthereum.PrepareEthTransaction(ctx, client, account, gasPrice)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}
	tx, err := contract.Vote(auth, big.NewInt(self.DisputeID), self.Support)
	if err != nil {
		return errors.Wrapf(err, "submit vote transaction")
	}

	level.Info(logger).Log("msg", "vote submitted with transaction", "tx", tx.Hash())
	return nil
}

type tallyCmd struct {
	cfgGas
	disputeID
}

func (self tallyCmd) Run() error {
	logger := logging.NewLogger()
	ctx := context.Background()

	_, err := config.ParseConfig(logger, string(self.Config)) // Load the env file.
	if err != nil {
		return errors.Wrap(err, "creating config")
	}

	client, err := tEthereum.NewClient(ctx, logger)
	if err != nil {
		return errors.Wrap(err, "creating ethereum client")
	}

	accounts, err := tEthereum.GetAccounts()
	if err != nil {
		return err
	}
	contract, err := contracts.NewITellor(client)
	if err != nil {
		return errors.Wrap(err, "create tellor contract instance")
	}

	var gasPrice *big.Int
	if self.GasPrice > 0 {
		gasPrice = big.NewInt(int64(self.GasPrice) * params.GWei)
	}

	auth, err := tEthereum.PrepareEthTransaction(ctx, client, accounts[0], gasPrice)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}

	tx, err := contract.TallyVotes(auth, big.NewInt(self.DisputeID))
	if err != nil {
		return errors.Wrapf(err, "run tally votes if you've already voted")
	}

	level.Info(logger).Log("msg", "tally votes submitted", "tx", tx.Hash().Hex())
	return nil
}

type listCmd struct {
	cfg
}

func (self listCmd) Run() error {
	logger := logging.NewLogger()
	ctx := context.Background()

	cfg, err := config.ParseConfig(logger, string(self.Config)) // Load the env file.
	if err != nil {
		return errors.Wrap(err, "creating config")
	}

	client, err := tEthereum.NewClient(ctx, logger)
	if err != nil {
		return errors.Wrap(err, "creating ethereum client")
	}

	contract, err := contracts.NewITellor(client)
	if err != nil {
		return errors.Wrap(err, "create tellor contract instance")
	}

	// Open the TSDB database.
	var querable storage.SampleAndChunkQueryable
	if cfg.Db.RemoteHost != "" {
		querable, err = db.NewRemoteDB(cfg.Db)
		if err != nil {
			return errors.Wrap(err, "opening remote tsdb DB")
		}
	} else {
		tsdbOptions := tsdb.DefaultOptions()
		tsdbOptions.NoLockfile = true
		querable, err = tsdb.Open(cfg.Db.Path, nil, nil, tsdbOptions)
		if err != nil {
			return errors.Wrap(err, "opening tsdb DB")
		}
	}

	aggregator, err := aggregator.New(logger, ctx, cfg.Aggregator, querable)
	if err != nil {
		return errors.Wrap(err, "creating aggregator")
	}

	psr := psrTellor.New(logger, cfg.PsrTellor, aggregator)

	abi, err := abi.JSON(strings.NewReader(contracts.ITellorABI))
	if err != nil {
		return errors.Wrap(err, "parse abi")
	}

	// Just use nil for most of the variables, only using this object to call UnpackLog which only uses the abi.
	bar := bind.NewBoundContract(contract.Address, abi, nil, nil, nil)

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "get latest eth block header")
	}

	queryDays := int64(10)
	blocksPerDay := (24 * 60 * 60) / tEthereum.BlockTime
	lookBackDelta := big.NewInt(queryDays * blocksPerDay) // Interested in only 5 days worth of blocks in the past since disputes can be voted only for 2 days.
	startBlock := big.NewInt(0).Sub(header.Number, lookBackDelta)

	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   nil,
		Addresses: []common.Address{contract.Address},
		Topics:    [][]common.Hash{{abi.Events["NewDispute"].ID}},
	}

	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return errors.Wrap(err, "filter eth logs")
	}

	level.Info(logger).Log("msg", "disputes count", "daysLookBack", queryDays, "count", len(logs))
	for _, rawDispute := range logs {
		disputeI := contracts.ITellorNewDispute{}
		err := bar.UnpackLog(&disputeI, "NewDispute", rawDispute)
		if err != nil {
			return errors.Wrap(err, "unpack dispute event from logs")
		}

		fmt.Println("disputeI.DisputeId", disputeI.DisputeId)
		_, executed, votePassed, _, reportedAddr, reportingMiner, _, disputeVars, currTally, err := contract.GetAllDisputeVars(&bind.CallOpts{Context: ctx}, disputeI.DisputeId)
		if err != nil {
			return errors.Wrap(err, "get dispute details")
		}

		rounds, err := contract.GetDisputeUintVars(
			&bind.CallOpts{Context: ctx},
			disputeI.DisputeId,
			tEthereum.Keccak256([]byte("_DISPUTE_ROUNDS")),
		)
		if err != nil {
			return errors.Wrap(err, "get dispute rounds")
		}

		votingEnds := time.Unix(disputeVars[3].Int64(), 0)
		votingWindow := 2 * 24 * time.Hour // 2 days.
		createdTime := votingEnds.Add(-votingWindow * time.Duration(rounds.Int64()))

		level.Info(logger).Log(
			"msg", "dispute details",
			"created", createdTime.Format("3:04:05 PM January 02, 2006 MST"),
			"executed", executed,
			"passed", votePassed,
			"disputeId", disputeI.DisputeId.String(),
			"requestId", disputeI.RequestId.Uint64(),
			"disputedTimestamp", time.Unix(disputeI.Timestamp.Int64(), 0).Format("3:04:05 PM January 02, 2006 MST"),
			"reported", reportedAddr.Hex(),
			"miner", reportingMiner.Hex(),
			"fee", math.BigInt18eToFloat(disputeVars[8]),
		)

		fmt.Println("currTally", currTally.String())

		tmp := new(big.Float)
		tmp.SetInt(currTally)
		currTallyFloat, _ := tmp.Float64()
		tmp.SetInt(disputeVars[7])
		currQuorum, _ := tmp.Float64()
		currTallyFloat += currQuorum
		currTallyRatio := currTallyFloat / 2 * currQuorum

		level.Info(logger).Log(
			"msg", "current dispute results",
			"currTallyRatio", fmt.Sprintf("%0.f%%", currTallyRatio*100),
			"TRB", math.BigInt18eToFloat(disputeVars[7]),
			"votes", disputeVars[4],
		)

		submits, err := contract.GetSubmissionsByTimestamp(
			&bind.CallOpts{Context: ctx},
			disputeI.RequestId,
			disputeI.Timestamp,
		)
		if err != nil {
			level.Error(logger).Log("msg", "getting all submits", "err", err)
			continue
		}

		suggested, err := psr.GetValue(disputeI.RequestId.Int64(), time.Unix(disputeI.Timestamp.Int64(), 0))
		if err != nil {
			// level.Error(logger).Log("msg", "look up recomended value", "id", disputeI.RequestId.Int64(), "err", err)
		}

		for i, submit := range submits {
			var disputed bool
			if i == int(disputeVars[6].Int64()) {
				disputed = true
			}
			level.Info(logger).Log(
				"msg", "submit details",
				"submitValue", submit,
				"suggestedValue", suggested,
				"minerIndex", i,
				"disputed", disputed,
			)
		}
	}

	return nil
}
