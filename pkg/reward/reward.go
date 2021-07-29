// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package reward

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"

	"github.com/tellor-io/telliot/pkg/aggregator"
	"github.com/tellor-io/telliot/pkg/contracts"
	eth "github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "reward"
const DefaultRetry = 30 * time.Second

type Config struct {
	LogLevel string
}

type Reward struct {
	logger         log.Logger
	contractCaller contracts.ContractCaller
	ctx            context.Context
	stop           context.CancelFunc
	aggr           aggregator.IAggregator
}

func New(
	logger log.Logger,
	ctx context.Context,
	cfg Config,
	contractCaller contracts.ContractCaller,
	aggr aggregator.IAggregator,
) (*Reward, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)

	ctx, cncl := context.WithCancel(ctx)
	return &Reward{
		logger:         logger,
		contractCaller: contractCaller,
		ctx:            ctx,
		stop:           cncl,
		aggr:           aggr,
	}, nil
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
		"msg", "profit checking current",
		"reward", float64(rewardEth1e18.Int64())/params.Ether,
		"txCost", float64(txCostEth1e18.Int64())/params.Ether,
		"slot", slot,
		"gasUsed", gasUsed,
		"gasPrice", float64(gasPriceEth1e18.Int64())/params.GWei,
		"profit", fmt.Sprintf("%.2e", float64(profit.Int64())),
		"profitMargin", profitPercent,
	)

	return profitPercent, nil
}

// GasPriceForProfitThreshold returns the gas price in WEI for a given profit percent based on the current TRB price and gas usage.
func (self *Reward) GasPriceForProfitThreshold(ctx context.Context, slot uint64, profitTargetPercent float64, gasUsed *big.Int) (int64, error) {
	rewardEth1e18, err := self.rewardInEth1e18()
	if err != nil {
		return 0, errors.New("getting trb current TRB price")
	}

	reward := float64(rewardEth1e18.Int64()) / params.Ether

	profitTarget := reward * (profitTargetPercent / 100)

	cost := reward - profitTarget
	gasPrice := cost / float64(gasUsed.Int64())

	level.Debug(self.logger).Log(
		"msg", "profit checking for threshold",
		"reward", float64(rewardEth1e18.Int64())*params.GWei,
		"txCost", cost/params.GWei,
		"slot", slot,
		"gasUsed", gasUsed,
		"gasPrice", gasPrice*params.GWei,
		"profitTarget", profitTarget,
	)

	return int64(gasPrice * params.Ether), nil
}

func (self *Reward) rewardInEth1e18() (*big.Int, error) {
	trbAmount1e18, err := self.contractCaller.CurrentReward(&bind.CallOpts{Context: self.ctx})
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
	slot, err := self.contractCaller.GetUintVar(nil, eth.Keccak256([]byte("_SLOT_PROGRESS")))
	if err != nil {
		return nil, errors.Wrap(err, "getting _SLOT_PROGRESS")
	}
	return slot, nil
}
