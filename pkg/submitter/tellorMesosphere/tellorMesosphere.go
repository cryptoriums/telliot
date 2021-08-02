// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tellorMesosphere

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/format"
	"github.com/tellor-io/telliot/pkg/logging"
	mathU "github.com/tellor-io/telliot/pkg/math"
	psr "github.com/tellor-io/telliot/pkg/psr/tellorMesosphere"
)

const (
	ComponentName = "submitterTellorMesosphere"
)

type Transactor interface {
	Transact(ctx context.Context, reqId *big.Int, val *big.Int) (*types.Transaction, error)
}

type Config struct {
	Enabled              bool
	LogLevel             string
	MinSubmitPeriod      format.Duration `help:"The time limit between each submit for a staked miner."`
	MinSubmitPriceChange float64         `help:" Submit only if that price changed at least that much percent."`
}

/**
* The submitter has one purpose: to either submit the solution on-chain
* or to reject it if the miner has already submitted a solution for the challenge
* or the the solution'self challenge does not match current challenge
 */

type Submitter struct {
	ctx             context.Context
	close           context.CancelFunc
	logger          log.Logger
	cfg             Config
	account         *ethereum.Account
	client          *ethclient.Client
	contract        *contracts.ITellorMesosphere
	transactor      Transactor
	psr             *psr.Psr
	lastSubmitValue map[int64]float64
	lastSubmitTime  map[int64]time.Time
	reqIDs          []int64
}

func New(
	ctx context.Context,
	logger log.Logger,
	cfg Config,
	client *ethclient.Client,
	contract *contracts.ITellorMesosphere,
	account *ethereum.Account,
	transactor Transactor,
	psr *psr.Psr,
) (*Submitter, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)
	ctx, close := context.WithCancel(ctx)
	submitter := &Submitter{
		ctx:             ctx,
		close:           close,
		client:          client,
		cfg:             cfg,
		account:         account,
		logger:          logger,
		contract:        contract,
		transactor:      transactor,
		psr:             psr,
		reqIDs:          []int64{1, 2},
		lastSubmitValue: make(map[int64]float64),
		lastSubmitTime:  make(map[int64]time.Time),
	}

	// Set the initial values
	for _, reqID := range submitter.reqIDs {
		submitter.lastSubmitValue[reqID] = 0
		submitter.lastSubmitTime[reqID] = time.Unix(0, 0)
	}

	return submitter, nil
}

func (self *Submitter) Start() error {
	for _, reqID := range self.reqIDs {
		exists, val, ts, err := self.contract.GetCurrentValue(&bind.CallOpts{Context: self.ctx}, big.NewInt(1))
		if err != nil {
			level.Error(self.logger).Log("msg", "retrieve current value while checking for last submit", "reqID", reqID, "err", err)
			break
		}
		if !exists {
			level.Error(self.logger).Log("msg", "current value doesn't exist for checking for last submit", "reqID", reqID)
			break
		}
		self.lastSubmitValue[reqID] = float64(val.Int64())
		self.lastSubmitTime[reqID] = time.Unix(ts.Int64(), 0)
		level.Debug(self.logger).Log(
			"msg", "recorded initial values",
			"reqID", reqID,
			"lastSubmitValue", self.lastSubmitValue[reqID],
			"lastSubmitTime", time.Since(self.lastSubmitTime[reqID]),
		)
	}

	for _, reqID := range self.reqIDs {
		if err := self.Submit(reqID); err != nil {
			level.Error(self.logger).Log("msg", "submit", "err", err)
		}
	}

	ticker := time.NewTicker(self.cfg.MinSubmitPeriod.Duration)
	defer ticker.Stop()
	for {
		select {
		case <-self.ctx.Done():
			return self.ctx.Err()
		case <-ticker.C:
			for _, reqID := range self.reqIDs {
				if err := self.Submit(reqID); err != nil {
					level.Error(self.logger).Log("msg", "submit", "reqID", reqID, "err", err)
				}
			}
		}
	}
}

func (self *Submitter) Stop() {
	self.close()
}

func (self *Submitter) Submit(reqID int64) error {
	ctx, cncl := context.WithTimeout(self.ctx, time.Minute)
	defer cncl()
	isReporter, err := self.contract.IsReporter(&bind.CallOpts{Context: ctx}, self.account.Address)
	if err != nil {
		return errors.Wrap(err, "checking reporter status")
	}
	if !isReporter {
		return errors.New("addr not a reporter")
	}

	val, err := self.psr.GetValue(reqID, time.Now())
	if err != nil {
		return errors.Wrap(err, "getting the value from the aggregator")
	}

	if !self.shouldSubmit(reqID, float64(val)) {
		return nil
	}
	level.Info(self.logger).Log(
		"msg", "sending values to the chain",
		"ID", reqID,
		"val", val,
	)

	tx, err := self.transactor.Transact(ctx, big.NewInt(reqID), big.NewInt(val))
	if err != nil {
		return errors.Wrap(err, "transacting")
	}

	level.Info(self.logger).Log("msg", "successfully submited solution",
		"txHash", tx.Hash().String(),
		"nonce", tx.Nonce(),
		"gasPrice", tx.GasPrice(),
		"gasLimit", tx.Gas(),
	)

	self.lastSubmitValue[reqID] = float64(val)
	self.lastSubmitTime[reqID] = time.Now()
	level.Debug(self.logger).Log(
		"msg", "recorded new values after a submit",
		"reqID", reqID,
		"lastSubmitValue", self.lastSubmitValue[reqID],
		"lastSubmitTime", time.Since(self.lastSubmitTime[reqID]),
	)
	return nil
}

func (self *Submitter) shouldSubmit(reqID int64, newVal float64) bool {
	logger := log.With(self.logger, "msg", "should submit check passed", "reqID", reqID)

	if self.lastSubmitTime[reqID].IsZero() {
		level.Info(logger).Log(
			"reason", "first submit",
		)
		return true
	}

	if lastSubmitTime, ok := self.lastSubmitTime[reqID]; ok && time.Since(lastSubmitTime) > (5*time.Minute) {
		level.Info(logger).Log(
			"reason", "more then 5 minutes since last submit",
			"timePassed", time.Since(lastSubmitTime),
		)
		return true
	}

	lastSubmitValue, ok := self.lastSubmitValue[reqID]
	if !ok {
		level.Error(self.logger).Log("msg", "last value check - no record for last value")
	}

	PercentageDiff := math.Abs(mathU.PercentageDiff(lastSubmitValue, newVal))
	if PercentageDiff > self.cfg.MinSubmitPriceChange {
		level.Info(logger).Log(
			"reason", "value change more then threshold",
			"PercentageDiff", PercentageDiff,
			"percentageThresohld", self.cfg.MinSubmitPriceChange,
			"lastSubmitValue", lastSubmitValue,
			"newValue", newVal,
		)
		return true
	}
	return false
}
