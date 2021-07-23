// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package reward

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"

	"github.com/tellor-io/telliot/pkg/aggregator"
	"github.com/tellor-io/telliot/pkg/contracts"
	eth "github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "rewardTracker"
const DefaultRetry = 30 * time.Second

type Config struct {
	LogLevel string
}

type Reward struct {
	client           *ethclient.Client
	logger           log.Logger
	contractInstance *contracts.ITellor
	ctx              context.Context
	stop             context.CancelFunc
	aggr             aggregator.IAggregator
	gasUsage         map[uint64]int64
	mtx              sync.Mutex
}

func NewReward(
	logger log.Logger,
	ctx context.Context,
	cfg Config,
	client *ethclient.Client,
	contractInstance *contracts.ITellor,
	aggr aggregator.IAggregator,
) (*Reward, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)

	ctx, cncl := context.WithCancel(ctx)
	return &Reward{
		client:           client,
		logger:           logger,
		contractInstance: contractInstance,
		ctx:              ctx,
		stop:             cncl,
		aggr:             aggr,
		gasUsage:         make(map[uint64]int64),
	}, nil
}

// recordGasUsageRefund gets the gas estimation before the TX was mined and
// subtracts the actual gas usage to record the refund.
func (self *Reward) RecordGasUsageRefund(value int64, slot uint64) {
	self.mtx.Lock()
	defer self.mtx.Unlock()
	self.gasUsage[slot] = value
}

// Current returns the profit in percents based on the current TRB price and gas usage.
func (self *Reward) Current(ctx context.Context, slot uint64, gasPriceEth1e18 *big.Int, gasUsed *big.Int) (int64, error) {

	rewardEth1e18, err := self.rewardInEth1e18()
	if err != nil {
		return 0, errors.New("getting trb current TRB price")
	}

	txCostEth1e18 := big.NewInt(0).Mul(gasPriceEth1e18, gasUsed)
	profit := big.NewInt(0).Sub(rewardEth1e18, txCostEth1e18)
	profitPercentFloat := float64(profit.Int64()) / float64(txCostEth1e18.Int64()) * 100
	profitPercent := int64(profitPercentFloat)

	level.Debug(self.logger).Log(
		"msg", "profit checking",
		"reward", fmt.Sprintf("%.2e", float64(rewardEth1e18.Int64())),
		"txCost", fmt.Sprintf("%.2e", float64(txCostEth1e18.Int64())),
		"slot", slot,
		"gasUsed", gasUsed,
		"gasPrice", gasPriceEth1e18,
		"profit", fmt.Sprintf("%.2e", float64(profit.Int64())),
		"profitMargin", profitPercent,
	)

	return profitPercent, nil
}

func (self *Reward) Refund(slot *big.Int) (int64, error) {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	if refund, ok := self.gasUsage[slot.Uint64()]; ok {
		return refund, nil
	}
	return 0, errors.New("no refund value")
}

func (self *Reward) rewardInEth1e18() (*big.Int, error) {
	trbAmount1e18, err := self.contractInstance.CurrentReward(nil)
	if err != nil {
		return nil, errors.New("getting currentReward from the chain")
	}

	trbPrice, confidence, err := self.aggr.TimeWeightedAvg("TRB/ETH", time.Now(), time.Hour)
	if err != nil {
		return nil, errors.New("getting the trb price from the aggregator")
	}

	if confidence < 0.5 {
		return nil, errors.New("trb price confidence too low")

	}

	rewardEth1e18 := big.NewFloat(0).Mul(big.NewFloat(0).SetInt(trbAmount1e18), big.NewFloat(trbPrice))

	rewardEth1e18Int, accuracy := rewardEth1e18.Int64()
	if accuracy != big.Exact {
		return nil, errors.New("conversion precision loss")
	}

	return big.NewInt(rewardEth1e18Int), nil
}

func (self *Reward) Slot() (*big.Int, error) {
	slot, err := self.contractInstance.GetUintVar(nil, eth.Keccak256([]byte("_SLOT_PROGRESS")))
	if err != nil {
		return nil, errors.Wrap(err, "getting _SLOT_PROGRESS")
	}
	return slot, nil
}
