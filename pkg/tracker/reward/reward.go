// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package reward

import (
	"context"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/tsdb"

	"github.com/tellor-io/telliot/pkg/aggregator"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
	"github.com/tellor-io/telliot/pkg/db"
	eth "github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "rewardTracker"
const DefaultRetry = 30 * time.Second

type Config struct {
	LogLevel string
}

type RewardTracker struct {
	client           *ethclient.Client
	logger           log.Logger
	contractInstance *contracts.ITellor
	ctx              context.Context
	stop             context.CancelFunc
	addr             common.Address

	tsDB   *tsdb.DB
	aggr   aggregator.IAggregator
	engine *promql.Engine

	lastGasEstimation map[uint64]uint64
}

func NewRewardTracker(
	logger log.Logger,
	ctx context.Context,
	cfg Config,
	tsDB *tsdb.DB,
	client *ethclient.Client,
	contractInstance *contracts.ITellor,
	addr common.Address,
	aggr aggregator.IAggregator,
) (*RewardTracker, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)

	opts := promql.EngineOpts{
		Logger:               logger,
		Reg:                  nil,
		MaxSamples:           30000,
		Timeout:              10 * time.Second,
		LookbackDelta:        5 * time.Minute,
		EnableAtModifier:     true,
		EnableNegativeOffset: true,
	}
	engine := promql.NewEngine(opts)

	ctx, cncl := context.WithCancel(ctx)
	return &RewardTracker{
		client:           client,
		logger:           logger,
		contractInstance: contractInstance,
		addr:             addr,
		ctx:              ctx,
		stop:             cncl,
		tsDB:             tsDB,
		engine:           engine,
		aggr:             aggr,
	}, nil
}

func (self *RewardTracker) Start() error {
	level.Info(self.logger).Log("msg", "starting")

	var err error
	ticker := time.NewTicker(DefaultRetry)
	defer ticker.Stop()

	logger := log.With(self.logger, "event", "NonceSubmitted")

	var sub event.Subscription
	events := make(chan *tellor.TellorNonceSubmitted)
	for {
		select {
		case <-self.ctx.Done():
			return errors.New("context canceled")
		default:
		}
		sub, err = self.nonceSubmittedSub(events)
		if err != nil {
			level.Error(logger).Log("msg", "initial subscribing to events failed")
			<-ticker.C
			continue
		}
		break
	}

	for {
		select {
		case <-self.ctx.Done():
			return errors.New("context canceled")
		case err := <-sub.Err():
			if err != nil {
				level.Error(logger).Log(
					"msg",
					"subscription error",
					"err", err)
			}

			// Trying to resubscribe until it succeeds.
			for {
				select {
				case <-self.ctx.Done():
					return errors.New("context canceled")
				default:
				}
				sub, err = self.nonceSubmittedSub(events)
				if err != nil {
					level.Error(logger).Log("msg", "re-subscribing to events failed", "err", err)
					<-ticker.C
					continue
				}
				break
			}
			level.Info(logger).Log("msg", "re-subscribed to events")
		case event := <-events:
			err := self.recordGasEstimation()
			if err != nil {
				level.Error(self.logger).Log("msg", "record gas estimation", "err", err)
			}
			gasUsed, err := self.recordGasUsage(event)
			if err != nil {
				level.Error(self.logger).Log("msg", "record gas usage", "err", err)
			}
			err = self.recordGasUsageRefund(gasUsed, event.Slot.Uint64())
			if err != nil {
				level.Error(self.logger).Log("msg", "record gas usage estimation", "err", err)
			}
		}
	}
}

func (self *RewardTracker) Stop() {
	self.stop()
}

func (self *RewardTracker) recordGasUsage(event *tellor.TellorNonceSubmitted) (uint64, error) {
	receipt, err := self.client.TransactionReceipt(self.ctx, event.Raw.TxHash)
	if err != nil {
		return 0, errors.Wrap(err, "receipt retrieval")
	} else if receipt != nil && receipt.Status == types.ReceiptStatusSuccessful {
		level.Debug(self.logger).Log("msg", "adding gas used", "gasUsed", receipt.GasUsed)

		lbls := labels.Labels{
			labels.Label{Name: "__name__", Value: "gas_usage_actual"},
			labels.Label{Name: "slot", Value: event.Slot.String()},
		}
		return receipt.GasUsed, db.Add(self.ctx, self.tsDB, lbls, float64(receipt.GasUsed))
	}
	return 0, errors.New("transaction not yet mined")
}

func (self *RewardTracker) recordGasEstimation() error {
	slotCurrent, err := self.slot()
	if err != nil {
		return errors.Wrap(err, "getting current slot")
	}

	gasEstimation, err := contracts.EstimateGasUsageSubmitMiningSolution(
		self.ctx,
		self.contractInstance,
		self.client,
		self.addr,
	)
	if err != nil {
		return errors.Wrap(err, "estimate gas usage from client")
	}

	self.lastGasEstimation[slotCurrent.Uint64()] = gasEstimation
	return nil
}

// recordGasUsageRefund gets the gas estimation before the TX was mined and
// subtracts the actual gas usage to record the refund.
func (self *RewardTracker) recordGasUsageRefund(gasUsedActual uint64, slot uint64) error {
	gasEstimation, ok := self.lastGasEstimation[slot]
	if !ok {
		level.Warn(self.logger).Log("msg", "no gas estimation for slot, this is normal for the first event iteration", "num", slot)
	}
	lbls := labels.Labels{
		labels.Label{Name: "__name__", Value: "gas_usage_refund"},
		labels.Label{Name: "slot", Value: strconv.Itoa(int(slot))},
	}
	return db.Add(self.ctx, self.tsDB, lbls, float64(gasEstimation-gasUsedActual))
}

func (self *RewardTracker) nonceSubmittedSub(output chan *tellor.TellorNonceSubmitted) (event.Subscription, error) {
	tellorFilterer, err := tellor.NewTellorFilterer(self.contractInstance.Address, self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting instance")
	}
	sub, err := tellorFilterer.WatchNonceSubmitted(&bind.WatchOpts{Context: self.ctx}, output, []common.Address{}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting channel")
	}
	return sub, nil
}

func (self *RewardTracker) slot() (*big.Int, error) {
	slot, err := self.contractInstance.GetUintVar(nil, eth.Keccak256([]byte("_SLOT_PROGRESS")))
	if err != nil {
		return nil, errors.Wrap(err, "getting _SLOT_PROGRESS")
	}
	return slot, nil
}
