// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package transactor

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/gas_price"
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "transactorTellorMesosphere"

type Config struct {
	LogLevel      string
	GasMax        uint
	GasMultiplier int
}

// TellorMesosphere implements the Transactor interface.
type TellorMesosphere struct {
	netID           *big.Int
	cfg             Config
	logger          log.Logger
	gasPriceQuerier gas_price.GasPriceQuerier
	client          *ethclient.Client
	account         *ethereum.Account
	contract        *contracts.ITellorMesosphere
}

func New(
	logger log.Logger,
	cfg Config,
	gasPriceQuerier gas_price.GasPriceQuerier,
	client *ethclient.Client,
	account *ethereum.Account,
	contract *contracts.ITellorMesosphere,
) (*TellorMesosphere, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}

	ctx, cncl := context.WithTimeout(context.Background(), 2*time.Second)
	defer cncl()
	netID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting network id")
	}

	return &TellorMesosphere{
		netID:           netID,
		cfg:             cfg,
		logger:          log.With(logger, "component", ComponentName),
		gasPriceQuerier: gasPriceQuerier,
		client:          client,
		account:         account,
		contract:        contract,
	}, nil
}

func (self *TellorMesosphere) Transact(ctx context.Context, reqId *big.Int, val *big.Int) (*types.Transaction, error) {
	gasEstimate := 3_000_000
	nonce, err := self.client.NonceAt(ctx, self.account.Address, nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting nonce for miner address")
	}

	// Use the same nonce in case there is a stuck transaction so that it resubmits the same TX with higher gas price.
	IntNonce := int64(nonce)

	gasPrice, err := self.gasPriceQuerier.Query(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting data from the db")
	}

	mul := self.cfg.GasMultiplier
	if mul > 0 {
		level.Info(self.logger).Log("msg", "settings gas price multiplier", "value", mul)
		gasPrice = gasPrice.Mul(gasPrice, big.NewInt(int64(mul)))
	}

	var finalError error
	for i := 0; i <= 5; i++ {
		balance, err := self.client.BalanceAt(ctx, self.account.Address, nil)
		if err != nil {
			finalError = err
			continue
		}

		cost := big.NewInt(1)
		cost = cost.Mul(gasPrice, big.NewInt(200000))
		if balance.Cmp(cost) < 0 {
			finalError = errors.Errorf("insufficient funds to send transaction: %v < %v", balance, cost)
			continue
		}

		opts, err := bind.NewKeyedTransactorWithChainID(self.account.PrivateKey, self.netID)
		if err != nil {
			return nil, errors.Wrap(err, "creating transactor")
		}
		opts.Context = ctx
		opts.Nonce = big.NewInt(IntNonce)
		opts.Value = big.NewInt(0) // in weiF
		opts.GasLimit = uint64(gasEstimate + 100_000)

		if gasPrice.Cmp(big.NewInt(0)) == 0 {
			gasPrice = big.NewInt(100)
		}
		if i > 1 {
			gasPrice1 := new(big.Int).Set(gasPrice)
			gasPrice1.Mul(gasPrice1, big.NewInt(int64(i*11))).Div(gasPrice1, big.NewInt(int64(100)))
			opts.GasPrice = gasPrice1.Add(gasPrice, gasPrice1)
		} else {
			// First time, try base gas price.
			opts.GasPrice = gasPrice
		}
		max := self.cfg.GasMax
		var maxGasPrice *big.Int
		gasPrice1 := big.NewInt(params.GWei)
		if max > 0 {
			maxGasPrice = gasPrice1.Mul(gasPrice1, big.NewInt(int64(max)))
		} else {
			maxGasPrice = gasPrice1.Mul(gasPrice1, big.NewInt(int64(100)))
		}

		if opts.GasPrice.Cmp(maxGasPrice) > 0 {
			level.Info(self.logger).Log("msg", "gas price too high, will default to the max price", "current", opts.GasPrice, "defaultMax", maxGasPrice)
			opts.GasPrice = maxGasPrice
		}

		tx, err := self.contract.SubmitValue(opts, reqId, val)
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "nonce too low") { // Can't use error type matching because of the way the eth client is implemented.
				IntNonce = IntNonce + 1
				level.Warn(self.logger).Log("msg", "last transaction has been confirmed so will increase the nonce and resend the transaction.")

			} else if strings.Contains(strings.ToLower(err.Error()), "replacement transaction underpriced") { // Can't use error type matching because of the way the eth client is implemented.
				level.Warn(self.logger).Log("msg", "last transaction is stuck so will increase the gas price and try to resend")
				finalError = err
			} else {
				finalError = errors.Wrap(err, "contract call")
			}

			delay := 15 * time.Second
			level.Info(self.logger).Log("msg", "will retry a send", "retryDelay", delay)
			select {
			case <-ctx.Done():
				return nil, errors.New("the submit context was canceled")
			case <-time.After(delay):
				continue
			}
		}

		return tx, nil
	}
	return nil, errors.Wrapf(finalError, "submit tx after 5 attempts")
}
