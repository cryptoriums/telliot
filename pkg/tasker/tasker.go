// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tasker

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/mining"
)

const ComponentName = "taskerNewChallenge"

const defaultDelay = 10 * time.Second

type Config struct {
	LogLevel string
}

type Tasker struct {
	ctx       context.Context
	close     context.CancelFunc
	logger    log.Logger
	accounts  []*ethereum.Account
	contract  *contracts.ITellor
	client    *ethclient.Client
	workSinks map[string]chan *mining.Work

	mtx            sync.Mutex
	pendingSubmits []context.CancelFunc
}

func New(
	ctx context.Context,
	logger log.Logger,
	cfg Config,
	client *ethclient.Client,
	contract *contracts.ITellor,
	accounts []*ethereum.Account,
) (*Tasker, map[string]chan *mining.Work, error) {
	ctx, close := context.WithCancel(ctx)
	workSinks := make(map[string]chan *mining.Work)
	for _, acc := range accounts {
		workSinks[acc.Address.String()] = make(chan *mining.Work)
	}
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		close()
		return nil, nil, errors.Wrap(err, "apply filter logger")
	}
	tasker := &Tasker{
		ctx:       ctx,
		close:     close,
		accounts:  accounts,
		contract:  contract,
		workSinks: workSinks,
		logger:    log.With(logger, "component", ComponentName),
		client:    client,
	}
	return tasker, tasker.workSinks, nil
}

func (self *Tasker) newSub(output chan *tellor.ITellorNewChallenge) (event.Subscription, error) {
	tellorFilterer, err := tellor.NewITellorFilterer(self.contract.Address, self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting filter instance")
	}
	sub, err := tellorFilterer.WatchNewChallenge(&bind.WatchOpts{Context: self.ctx}, output, nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting subscription channel")
	}
	return sub, nil
}

func (self *Tasker) sendWork(ctx context.Context, challenge *tellor.ITellorNewChallenge) {
	newChallenge := &mining.MiningChallenge{
		Challenge:  challenge.CurrentChallenge[:],
		Difficulty: challenge.Difficulty,
		RequestIDs: challenge.CurrentRequestId,
	}
	for _, acc := range self.accounts {
		level.Info(self.logger).Log("msg", "new event",
			"addr", acc.Address.String(),
			"challenge", fmt.Sprintf("%x", newChallenge.Challenge),
			"difficulty", newChallenge.Difficulty,
			"requestIDs", fmt.Sprintf("%+v", newChallenge.RequestIDs),
		)

		select {
		case self.workSinks[acc.Address.String()] <- &mining.Work{
			Context:    ctx,
			Challenge:  newChallenge,
			PublicAddr: acc.Address.String(),
			Start:      uint64(rand.Int63()),
			N:          math.MaxInt64,
		}:
		case <-self.ctx.Done():
			return
		}
	}

}

func (self *Tasker) Start() error {
	var err error
	ticker := time.NewTicker(defaultDelay)
	defer ticker.Stop()
	level.Info(self.logger).Log("msg", "starting")

	// Getting current challenge from the contract.
	currentVariables, err := self.contract.GetNewCurrentVariables(&bind.CallOpts{Context: self.ctx})
	if err != nil {
		level.Warn(self.logger).Log("msg", "getting new current variables", "err", err)
		return errors.Wrap(err, "getting GetNewCurrentVariables")
	}

	currentChallenge := &tellor.ITellorNewChallenge{
		CurrentChallenge: currentVariables.Challenge,
		Difficulty:       currentVariables.Difficutly,
		CurrentRequestId: currentVariables.RequestIds,
		TotalTips:        currentVariables.Tip,
	}

	level.Info(self.logger).Log("msg", "sending the initial event")

	self.mtx.Unlock()
	ctx, cncl := context.WithCancel(self.ctx)
	self.pendingSubmits = append(self.pendingSubmits, cncl)
	self.mtx.Unlock()
	self.sendWork(ctx, currentChallenge)

	// Subscribe and wait until the context cancellation event.
	events := make(chan *tellor.ITellorNewChallenge)
	var sub event.Subscription

	// Initial subscription.
	for {
		select {
		case <-self.ctx.Done():
			return nil
		default:
		}
		sub, err = self.newSub(events)
		if err != nil {
			level.Error(self.logger).Log("msg", "initial subscription to events failed")
			<-ticker.C
			continue
		}
		break
	}

	for {
		select {
		case <-self.ctx.Done():
			return nil
		case err := <-sub.Err():
			if err != nil {
				level.Error(self.logger).Log(
					"msg",
					"subscription error",
					"err", err)
			}

			// Trying to resubscribe until it succeeds.
			for {
				select {
				case <-self.ctx.Done():
					return nil
				default:
				}
				sub, err = self.newSub(events)
				if err != nil {
					level.Error(self.logger).Log("msg", "re-subscribing to events failed")
					<-ticker.C
					continue
				}
				break
			}
			level.Info(self.logger).Log("msg", "re-subscribed to events")
		case event := <-events:
			level.Debug(self.logger).Log("msg", "new event", "hash", event.Raw.TxHash, "reorg", event.Raw.Removed)
			if !event.Raw.Removed { // For reorg events just cancel the old TXs without sending this one.
				self.mtx.Unlock()
				ctx, cncl := context.WithCancel(self.ctx)
				self.pendingSubmits = append(self.pendingSubmits, cncl)
				self.mtx.Unlock()

				go self.sendWhenConfirmed(ctx, event)
			}
		}
	}
}

func (self *Tasker) sendWhenConfirmed(ctx context.Context, vLog *tellor.ITellorNewChallenge) {
	ticker := time.NewTicker(defaultDelay)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
		// Send the event only when the tx that emitted the event has been confirmed.
		receipt, err := self.client.TransactionReceipt(ctx, vLog.Raw.TxHash)
		if receipt != nil {
			// Send it only when the TX ReceiptStatusSuccessful or otherwise ignore.
			if receipt.Status == types.ReceiptStatusSuccessful {
				self.mtx.Unlock()
				for _, cncl := range self.pendingSubmits {
					cncl()
				}
				self.pendingSubmits = nil
				ctx, cncl := context.WithCancel(self.ctx)
				self.pendingSubmits = append(self.pendingSubmits, cncl)
				self.mtx.Unlock()
				self.sendWork(ctx, vLog)
			}
			return
		}
		if err != nil {
			level.Error(self.logger).Log("msg", "getting TX receipt", "err", err)
		} else {
			level.Debug(self.logger).Log("msg", "transaction not yet mined", "tx", vLog.Raw.TxHash)
		}
	}
}

func (self *Tasker) Stop() {
	self.close()
}
