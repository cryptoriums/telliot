// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package tasker

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/logging"
	"github.com/cryptoriums/telliot/pkg/mining"
	"github.com/cryptoriums/telliot/pkg/tracker/events"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

const (
	ComponentName = "taskerNewChallenge"
)

const defaultDelay = 10 * time.Second

type Config struct {
	LogLevel string
}

type Tasker struct {
	ctx           context.Context
	stop          context.CancelFunc
	logger        log.Logger
	cfg           Config
	accounts      []*ethereum.Account
	contract      contracts.ContractCaller
	client        *ethclient.Client
	workSinks     map[string]chan *mining.Work
	eventsTracker *events.TrackerEvents
	logsOutput    chan types.Log
}

func New(
	ctx context.Context,
	logger log.Logger,
	cfg Config,
	client *ethclient.Client,
	contract contracts.ContractCaller,
	accounts []*ethereum.Account,
) (*Tasker, map[string]chan *mining.Work, error) {
	workSinks := make(map[string]chan *mining.Work)
	for _, acc := range accounts {
		workSinks[acc.Address.Hex()] = make(chan *mining.Work)
	}
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "apply filter logger")
	}

	ctx, stop := context.WithCancel(ctx)
	defer func() {
		if err != nil {
			stop()
		}
	}()
	eventsTracker, logsOutput, err := events.New(
		ctx,
		logger,
		client,
		contract,
		0,
		contracts.EventNameNewTask,
		5*time.Second,
	)
	if err != nil {
		return nil, nil, errors.Wrap(err, "creating logs tracker")
	}

	tasker := &Tasker{
		ctx:           ctx,
		cfg:           cfg,
		stop:          stop,
		accounts:      accounts,
		contract:      contract,
		eventsTracker: eventsTracker,
		logsOutput:    logsOutput,
		workSinks:     workSinks,
		logger:        log.With(logger, "component", ComponentName),
		client:        client,
	}
	return tasker, tasker.workSinks, nil
}

func (self *Tasker) sendWork(ctx context.Context, challenge *contracts.TellorNewChallenge) {
	newChallenge := &mining.MiningChallenge{
		Challenge:  challenge.CurrentChallenge[:],
		Difficulty: challenge.Difficulty,
		RequestIDs: challenge.CurrentRequestId,
	}
	for _, acc := range self.accounts {
		level.Info(self.logger).Log("msg", "new event",
			"addr", acc.Address.Hex()[:9],
			"challenge", fmt.Sprintf("%x", newChallenge.Challenge),
			"difficulty", newChallenge.Difficulty,
			"requestIDs", fmt.Sprintf("%+v", newChallenge.RequestIDs),
		)

		select {
		case self.workSinks[acc.Address.Hex()] <- &mining.Work{
			Context:    ctx,
			Challenge:  newChallenge,
			PublicAddr: acc.Address.Hex(),
			Start:      uint64(rand.Int63()),
			N:          math.MaxInt64,
		}:
		case <-self.ctx.Done():
			return
		}
	}

}

func (self *Tasker) Start() error {
	level.Info(self.logger).Log("msg", "starting", "logLevel", self.cfg.LogLevel)

	go func() {
		if err := self.eventsTracker.Start(); err != nil {
			level.Error(self.logger).Log("msg", "starting the logs tracker", "err", err)
			self.Stop()
		}
	}()
	var err error
	ticker := time.NewTicker(defaultDelay)
	defer ticker.Stop()

	// Getting current challenge from the contract.
	currentVariables, err := self.contract.GetNewCurrentVariables(&bind.CallOpts{Context: self.ctx})
	if err != nil {
		return errors.Wrap(err, "getting GetNewCurrentVariables")
	}

	currentChallenge := &contracts.TellorNewChallenge{
		CurrentChallenge: currentVariables.Challenge,
		Difficulty:       currentVariables.Difficulty,
		CurrentRequestId: currentVariables.RequestIds,
		TotalTips:        currentVariables.Tip,
	}

	level.Info(self.logger).Log("msg", "sending the initial event")
	self.sendWork(context.Background(), currentChallenge)

	for {
		select {
		case <-self.ctx.Done():
			return nil
		case logRaw := <-self.logsOutput:
			level.Debug(self.logger).Log("msg", "new event", "hash", logRaw.TxHash)

			log := &contracts.TellorNewChallenge{}
			err := self.contract.UnpackLog(log, contracts.EventNameNewTask, logRaw)
			if err != nil {
				level.Error(self.logger).Log("msg", "unpack event from logs", "err", err)
				continue
			}
			self.sendWork(context.Background(), log)
		}
	}
}

func (self *Tasker) Stop() {
	self.stop()
}
