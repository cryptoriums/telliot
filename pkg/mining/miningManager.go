// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package mining

import (
	"context"
	"fmt"
	"time"

	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/logging"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

type Config struct {
	LogLevel  string
	Heartbeat time.Duration
}

type SolutionSink interface {
	Submit(context.Context, *Result) (*types.Transaction, error)
}

const NumProcessors = 1

func SetupMiningGroup(logger log.Logger, ctx context.Context, cfg Config, contractInstance contracts.TellorCaller) (*MiningGroup, error) {
	var hashers []Hasher
	level.Info(logger).Log("msg", "starting CPU mining", "threads", NumProcessors)
	for i := 0; i < NumProcessors; i++ {
		hashers = append(hashers, NewCpuMiner(int64(i)))
	}
	miningGrp, err := NewMiningGroup(logger, ctx, cfg, hashers, contractInstance)
	if err != nil {
		return nil, errors.Wrap(err, "creating new mining group")
	}
	return miningGrp, nil
}

// MiningMgr manages mining, submiting a solution and requesting data.
// In the tellor contract a solution is saved in slots where a value is valid only when it has 5 confirmed slots.
// The manager tracks tx costs and profitThreshold is set it skips any transactions below the profit threshold.
// The profit is calculated the same way as in the Tellor contract.
// Transaction cost for submitting in each slot might be different so because of this
// the manager needs to complete few transaction to gather the tx cost for each slot.
type MiningMgr struct {
	ctx              context.Context
	stop             context.CancelFunc
	logger           log.Logger
	ethClient        ethereum.EthClient
	group            *MiningGroup
	taskerCh         chan *Work
	submitterCh      chan *Result
	contractInstance contracts.TellorCaller
	toMineInput      chan *Work
	solutionOutput   chan *Result
}

// NewManager is the MiningMgr constructor.
func NewManager(
	ctx context.Context,
	logger log.Logger,
	cfg Config,
	contractInstance contracts.TellorCaller,
	taskerCh chan *Work,
	submitterCh chan *Result,
	client ethereum.EthClient,
) (*MiningMgr, error) {

	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)

	group, err := SetupMiningGroup(logger, ctx, cfg, contractInstance)
	if err != nil {
		return nil, errors.Wrap(err, "setup MiningGroup")
	}

	ctx, stop := context.WithCancel(ctx)
	mng := &MiningMgr{
		ctx:              ctx,
		stop:             stop,
		logger:           logger,
		group:            group,
		taskerCh:         taskerCh,
		submitterCh:      submitterCh,
		contractInstance: contractInstance,
		ethClient:        client,
		toMineInput:      make(chan *Work),
		solutionOutput:   make(chan *Result),
	}
	return mng, nil
}

// Start will start the mining run loop.
func (mgr *MiningMgr) Start() error {
	level.Info(mgr.logger).Log("msg", "starting")

	// Start the mining group.
	go mgr.group.Mine(mgr.ctx, mgr.toMineInput, mgr.solutionOutput)

	for {
		select {

		// Boss wants us to quit for the day.
		case <-mgr.ctx.Done():
			return mgr.ctx.Err()

		// Found a solution.
		case solution := <-mgr.solutionOutput:
			level.Info(mgr.logger).Log("msg", "sending the solution to the submitter")
			mgr.submitterCh <- solution

		// Listen for new work from the tasker and send for mining.
		case work := <-mgr.taskerCh:
			mgr.toMineInput <- work
			level.Info(mgr.logger).Log("msg", "sent new challenge to the mining group",
				"challenge", fmt.Sprintf("%x", work.Challenge.Challenge),
				"difficulty", work.Challenge.Difficulty,
				"requestIDs", fmt.Sprintf("%+v", work.Challenge.RequestIDs),
			)
		}
	}

}

// Stop will take care of stopping the miner component.
func (mgr *MiningMgr) Stop() {
	mgr.stop()
}
