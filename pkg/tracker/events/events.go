// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package events

import (
	"context"
	"math"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/bluele/gcache"
	ethereum_t "github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

const (
	ComponentName = "trackerEvent"
	defaultDelay  = 10 * time.Second
)

type Contract interface {
	Abi() abi.ABI
	WatchLogs(opts *bind.WatchOpts, name string, query ...[]interface{}) (chan types.Log, event.Subscription, error)
	Addr() common.Address
}

type TrackerEvents struct {
	logger      log.Logger
	ctx         context.Context
	stop        context.CancelFunc
	client      ethereum_t.EthClient
	mtx         sync.Mutex
	cacheHeadTX gcache.Cache
	contract    Contract
	lookBack    time.Duration
	eventName   string
	dstChan     chan types.Log

	reorgWaitPeriod  time.Duration
	reorgWaitPending map[string]context.CancelFunc
}

func New(
	ctx context.Context,
	logger log.Logger,
	client ethereum_t.EthClient,
	contract Contract,
	lookBack time.Duration,
	eventName string,
	reorgWaitPeriod time.Duration,
) (*TrackerEvents, chan types.Log, error) {
	logger, err := logging.ApplyFilter("info", logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)
	ctx, stop := context.WithCancel(ctx)

	dstChan := make(chan types.Log)
	return &TrackerEvents{
		client:           client,
		ctx:              ctx,
		stop:             stop,
		logger:           logger,
		contract:         contract,
		lookBack:         lookBack,
		eventName:        eventName,
		dstChan:          dstChan,
		reorgWaitPeriod:  reorgWaitPeriod,
		reorgWaitPending: make(map[string]context.CancelFunc),
		// To be on the safe side create the cache 2 times bigger then the expected block count during the reorg wait.
		cacheHeadTX: gcache.New(int(math.Max(50, 2*ethereum_t.BlocksPerSecond*reorgWaitPeriod.Seconds()))).LRU().Build(),
	}, dstChan, nil
}
func (self *TrackerEvents) Start() error {
	level.Info(self.logger).Log("msg", "starting",
		"monitorAddr", self.contract.Addr(),
		"lookBack", self.lookBack,
		"eventName", self.eventName,
		"reorgWaitPeriod", self.reorgWaitPeriod,
	)
	if self.lookBack != 0 {
		ctx, cncl := context.WithTimeout(self.ctx, time.Minute)
		defer cncl()
		blockNums := ethereum_t.BlocksPerMinute * self.lookBack.Minutes()

		headerNow, err := self.client.HeaderByNumber(ctx, nil)
		if err != nil {
			return errors.Wrap(err, "get latest eth block header")
		}
		fromBlock := headerNow.Number.Int64() - int64(blockNums)
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(fromBlock),
			ToBlock:   nil,
			Addresses: []common.Address{self.contract.Addr()},
			Topics:    [][]common.Hash{{self.contract.Abi().Events[self.eventName].ID}},
		}

		logs, err := self.client.FilterLogs(self.ctx, query)
		if err != nil {
			return errors.Wrap(err, "getting historical logs")
		}
		for _, log := range logs {
			select {
			case self.dstChan <- log:
			case <-self.ctx.Done():
				return nil
			}
		}

	}

	// Initial subscription.
	src, subs := self.waitSubscribe()
	defer func() {
		if subs != nil {
			subs.Unsubscribe()
		}
	}()

	for {
		select {
		case <-self.ctx.Done():
			return nil
		case event := <-src:
			if event.Removed {
				self.cancelPending(hashFromLog(event))
				continue
			}
			ctx, cncl := context.WithCancel(self.ctx)
			self.addPending(hashFromLog(event), cncl)
			go func(ctxReorg context.Context, event types.Log) {
				waitReorg := time.NewTicker(self.reorgWaitPeriod)
				defer waitReorg.Stop()

				select {
				case <-waitReorg.C:
					self.cancelPending(hashFromLog(event))
					select {
					case self.dstChan <- event:
						return
					case <-self.ctx.Done():
						return
					}
				case <-ctxReorg.Done():
					return
				}

			}(ctx, event)
		case err := <-subs.Err():
			level.Error(self.logger).Log("msg", "subscription failed will try to resubscribe", "err", err)
			src, subs = self.waitSubscribe()

		}
	}
}

func (self *TrackerEvents) Stop() {
	self.stop()
}

func (self *TrackerEvents) waitSubscribe() (chan types.Log, event.Subscription) {

	ticker := time.NewTicker(defaultDelay)
	defer ticker.Stop()
MainLoop:
	for {
		select {
		case <-self.ctx.Done():
			return nil, &NoopSubs{} // To avoid panics in the caller.
		default:
		}

		opts := &bind.WatchOpts{
			Context: self.ctx,
		}
		src, subs, err := self.contract.WatchLogs(opts, self.eventName)
		if err != nil {
			level.Error(self.logger).Log("msg", "subscription to events failed", "err", err)
			select {
			case <-self.ctx.Done():
				return nil, &NoopSubs{} // To avoid panics in the caller.
			case <-ticker.C:
				continue MainLoop
			}
		}
		level.Info(self.logger).Log("msg", "subscription created", "eventName", self.eventName)
		return src, subs
	}
}

func hashFromLog(log types.Log) string {
	return log.TxHash.Hex() + "-indexInTheBlock:" + strconv.Itoa(int(log.Index))
}

func (self *TrackerEvents) addPending(hash string, ctx context.CancelFunc) {
	self.mtx.Lock()
	defer self.mtx.Unlock()
	self.reorgWaitPending[hash] = ctx
}

func (self *TrackerEvents) cancelPending(hash string) {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	if cncl, ok := self.reorgWaitPending[hash]; ok {
		cncl()
		delete(self.reorgWaitPending, hash)
	}
}

type NoopSubs struct{}

func (self *NoopSubs) Unsubscribe()      {}
func (self *NoopSubs) Err() <-chan error { return nil }
