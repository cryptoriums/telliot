// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package tellor

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/gas_estimator"
	"github.com/cryptoriums/telliot/pkg/gas_price"
	"github.com/cryptoriums/telliot/pkg/logging"
	"github.com/cryptoriums/telliot/pkg/math"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

const ComponentName = "transactorTellor"

type Config struct {
	LogLevel        string
	GasMaxTipGwei   uint   `help:"Hard limit of the gas tip in Gwei."`
	ProfitThreshold uint64 `help:"Minimum percent of profit when submitting a solution. For example if the tx cost is 0.01 ETH and current reward is 0.02 ETH a ProfitThreshold of 200% or more will wait until the reward is increased or the gas cost is lowered a ProfitThreshold of 199% or less will submit."`
}

// Tellor implements the Transactor interface.
type Tellor struct {
	contract        contracts.ContractCaller
	cfg             Config
	logger          log.Logger
	gasPriceQuerier gas_price.GasPriceQuerier
	client          *ethclient.Client
	account         *ethereum.Account
	rewardQ         contracts.RewardQuerier
	gasEstimator    gas_estimator.GasEstimator
}

func New(
	logger log.Logger,
	cfg Config,
	gasPriceQuerier gas_price.GasPriceQuerier,
	client *ethclient.Client,
	account *ethereum.Account,
	rewardQ contracts.RewardQuerier,
	contract contracts.ContractCaller,
) (*Tellor, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}

	return &Tellor{
		cfg:             cfg,
		logger:          log.With(logger, "component", ComponentName),
		gasPriceQuerier: gasPriceQuerier,
		client:          client,
		account:         account,
		rewardQ:         rewardQ,
		gasEstimator:    gas_estimator.NewDefault(),
		contract:        contract,
	}, nil
}

func (self *Tellor) Transact(ctx context.Context, solution string, ids [5]*big.Int, vals [5]*big.Int) (*types.Transaction, error) {
	slot, err := contracts.Slot(self.contract)
	if err != nil {
		level.Error(self.logger).Log("msg", "getting slot number", "err", err)
	}

	gasEstimate, err := self.gasEstimator.EstimateGas(ctx, self.account, slot.Uint64())
	if err != nil {
		level.Error(self.logger).Log("msg", "getting gas estimate", "err", err)
	}

	if self.cfg.ProfitThreshold > 0 && self.rewardQ != nil { // Profit check is enabled.
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		gasBaseFeeGwei, gasTipGwei, err := self.gasPriceQuerier.Query(ctx, 80)
		if err != nil {
			return nil, errors.Wrap(err, "getting gas price from guerier")
		}

		txCostGwei := (gasBaseFeeGwei + gasTipGwei) * float64(gasEstimate)

		for {
			select {
			case <-ctx.Done():
				return nil, errors.New("pending submit canceled")
			default:
			}
			profitPercent, err := self.rewardQ.Query(ctx, txCostGwei)
			if err != nil {
				level.Info(self.logger).Log("msg", "submit solution profit check", "err", err)
				<-ticker.C
				continue
			}
			if profitPercent < int64(self.cfg.ProfitThreshold) {
				level.Info(self.logger).Log("msg", "profit lower then the profit threshold", "profit", profitPercent, "threshold", self.cfg.ProfitThreshold)
				<-ticker.C
				continue
			}
			break
		}
	}

	// Use the pending nonce in case there is a stuck transaction so that it resubmits the same TX with higher gas price.
	nonce, err := self.client.PendingNonceAt(ctx, self.account.Address)
	if err != nil {
		return nil, errors.Wrap(err, "getting nonce for reporter address")
	}

	delay := 15 * time.Second
	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	var finalError error
	for i := 0; i <= 5; i++ {
		gasBaseFeeGwei, gasTipGwei, err := self.gasPriceQuerier.Query(ctx, 80)
		if err != nil {
			return nil, errors.Wrap(err, "getting gas price from guerier")
		}

		balance, err := self.client.BalanceAt(ctx, self.account.Address, nil)
		if err != nil {
			finalError = err
			continue
		}

		balanceGwei := big.NewInt(0).Div(balance, big.NewInt(params.GWei))

		txCostGwei := (gasBaseFeeGwei + gasTipGwei) * float64(gasEstimate)

		if float64(balanceGwei.Int64()) < txCostGwei {
			return nil, errors.Errorf("insufficient funds to send transaction: %v gwei < %v gwei", balanceGwei, txCostGwei)
		}

		opts, err := bind.NewKeyedTransactorWithChainID(self.account.PrivateKey, big.NewInt(self.contract.NetID()))
		if err != nil {
			return nil, errors.Wrap(err, "creating transactor")
		}

		maxGasTipGwei := float64(self.cfg.GasMaxTipGwei)
		if gasTipGwei > maxGasTipGwei {
			gasTipGwei = maxGasTipGwei
		}

		// Tip + BaseFee and also adding 25% to the base fee to accommodate for any network load surged.
		// When there is a surge the base fee increases 12.5% on every next block.
		gasFeeCap := math.FloatToBigIntMul(gasBaseFeeGwei+(gasBaseFeeGwei*0.50)+gasTipGwei, params.GWei)

		opts.GasLimit = uint64(gasEstimate + 300_000)
		opts.Context = ctx
		opts.Nonce = big.NewInt(int64(nonce))
		opts.Value = big.NewInt(0)
		opts.GasTipCap = big.NewInt(0).Mul(big.NewInt(int64(gasTipGwei)), big.NewInt(int64(params.GWei)))
		opts.GasFeeCap = gasFeeCap

		tx, err := self.contract.SubmitMiningSolution(opts, solution, ids, vals)
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "nonce too low") { // Can't use error type matching because of the way the eth client is implemented.
				nonce = nonce + 1
				level.Warn(self.logger).Log("msg", "last transaction has been confirmed so will increase the nonce and resend the transaction.")

			} else {
				finalError = errors.Wrap(err, "contract call")
			}

			level.Info(self.logger).Log("msg", "will retry a send", "retryDelay", delay, "err", err)
			select {
			case <-ctx.Done():
				return nil, errors.New("the submit context was canceled")
			case <-ticker.C:
				continue
			}
		}
		return tx, nil
	}
	return nil, errors.Wrapf(finalError, "submit tx after 5 attempts")
}
