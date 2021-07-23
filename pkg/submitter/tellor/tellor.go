// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package tellor

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/format"
	"github.com/tellor-io/telliot/pkg/gasPrice"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/mining"
	psr "github.com/tellor-io/telliot/pkg/psr/tellor"
	"github.com/tellor-io/telliot/pkg/reward"
	"github.com/tellor-io/telliot/pkg/transactor"
)

const ComponentName = "submitterTellor"

type Config struct {
	Enabled         bool
	LogLevel        string
	ProfitThreshold uint64          `help:"Minimum percent of profit when submitting a solution. For example if the tx cost is 0.01 ETH and current reward is 0.02 ETH a ProfitThreshold of 200% or more will wait until the reward is increased or the gas cost is lowered a ProfitThreshold of 199% or less will submit."`
	MinSubmitPeriod format.Duration `help:"The time limit between each submit for a staked miner."`
}

/**
* The submitter has one purpose: to either submit the solution on-chain
* or to reject it if the miner has already submitted a solution for the challenge
* or the the solution'self challenge does not match current challenge
 */

type Submitter struct {
	ctx                     context.Context
	close                   context.CancelFunc
	logger                  log.Logger
	cfg                     Config
	account                 *ethereum.Account
	client                  *ethclient.Client
	contract                contracts.ContractCaller
	resultCh                chan *mining.Result
	submitCount             prometheus.Counter
	submitFailCount         prometheus.Counter
	submitValue             *prometheus.GaugeVec
	submitGasUsageEstimated *prometheus.GaugeVec
	submitGasUsageActual    *prometheus.GaugeVec
	lastSubmitCncl          context.CancelFunc
	transactor              transactor.Transactor
	reward                  *reward.Reward
	gasPriceQuerier         gasPrice.GasPriceQuerier
	psr                     *psr.Psr
}

func New(
	ctx context.Context,
	logger log.Logger,
	cfg Config,
	client *ethclient.Client,
	contract contracts.ContractCaller,
	account *ethereum.Account,
	reward *reward.Reward,
	transactor transactor.Transactor,
	gasPriceQuerier gasPrice.GasPriceQuerier,
	psr *psr.Psr,
) (*Submitter, chan *mining.Result, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)
	ctx, close := context.WithCancel(ctx)
	submitter := &Submitter{
		ctx:             ctx,
		close:           close,
		client:          client,
		cfg:             cfg,
		resultCh:        make(chan *mining.Result),
		account:         account,
		reward:          reward,
		logger:          logger,
		contract:        contract,
		transactor:      transactor,
		gasPriceQuerier: gasPriceQuerier,
		psr:             psr,
		submitCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace:   "telliot",
			Subsystem:   ComponentName,
			Name:        "submit_total",
			Help:        "The total number of submitted solutions",
			ConstLabels: prometheus.Labels{"account": account.Address.String()},
		}),
		submitFailCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace:   "telliot",
			Subsystem:   ComponentName,
			Name:        "submit_fails_total",
			Help:        "The total number of failed submission",
			ConstLabels: prometheus.Labels{"account": account.Address.String()},
		}),
		submitValue: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace:   "telliot",
			Subsystem:   ComponentName,
			Name:        "submit_value",
			Help:        "The submitted value",
			ConstLabels: prometheus.Labels{"account": account.Address.String()},
		},
			[]string{"id"},
		),
		submitGasUsageEstimated: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace:   "telliot",
			Subsystem:   ComponentName,
			Name:        "submit_gas_usage_estimated",
			Help:        "The submitted estimation for the gas usage",
			ConstLabels: prometheus.Labels{"account": account.Address.String()},
		},
			[]string{"slot"},
		),
		submitGasUsageActual: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace:   "telliot",
			Subsystem:   ComponentName,
			Name:        "submit_gas_usage_actual",
			Help:        "The submitted actual result for the gas usage",
			ConstLabels: prometheus.Labels{"account": account.Address.String()},
		},
			[]string{"slot"},
		),
	}

	return submitter, submitter.resultCh, nil
}

func (self *Submitter) Start() error {
	for {
		select {
		case <-self.ctx.Done():
			self.CancelPendingSubmit()
			return self.ctx.Err()
		case result := <-self.resultCh:
			self.CancelPendingSubmit()
			var ctx context.Context
			ctx, self.lastSubmitCncl = context.WithCancel(self.ctx)

			level.Info(self.logger).Log("msg", "received a solution",
				"challenge", fmt.Sprintf("%x", result.Work.Challenge),
				"solution", result.Nonce,
				"difficulty", result.Work.Challenge.Difficulty,
				"requestIDs", fmt.Sprintf("%+v", result.Work.Challenge.RequestIDs),
			)
			self.Submit(ctx, result)
		}
	}
}

func (self *Submitter) CancelPendingSubmit() {
	if self.lastSubmitCncl != nil {
		self.lastSubmitCncl()
	}
}

func (self *Submitter) Stop() {
	self.close()
}

func (self *Submitter) blockUntilTimeToSubmit(newChallengeReplace context.Context) {
	var (
		lastSubmit time.Duration
		timestamp  *time.Time
		err        error
	)
	for {
		select {
		case <-newChallengeReplace.Done():
			level.Info(self.logger).Log("msg", "canceled pending submit while gettting last submit time")
		default:
		}
		lastSubmit, timestamp, err = self.lastSubmit()
		if err != nil {
			level.Debug(self.logger).Log("msg", "checking last submit time", "err", err)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	if lastSubmit < self.cfg.MinSubmitPeriod.Duration {
		level.Info(self.logger).Log("msg", "min transaction submit threshold hasn't passed",
			"nextSubmit", time.Duration(self.cfg.MinSubmitPeriod.Nanoseconds())-lastSubmit,
			"lastSubmit", lastSubmit,
			"lastSubmitTimestamp", timestamp.Format("2006-01-02 15:04:05.000000"),
			"minSubmitPeriod", self.cfg.MinSubmitPeriod,
		)
		timeToSubmit, cncl := context.WithDeadline(newChallengeReplace, timestamp.Add(self.cfg.MinSubmitPeriod.Duration))
		defer cncl()
		select {
		case <-newChallengeReplace.Done():
			level.Info(self.logger).Log("msg", "canceled pending submit while waiting for the time to submit")
		case <-timeToSubmit.Done(): // 15min since last submit has passed so can unblock.
		}
	}
}

func (self *Submitter) profitPercent(slot uint64, gasUsage int64) (int64, error) {
	gasPrice, err := self.gasPriceQuerier.Query(self.ctx)
	if err != nil {
		return 0, errors.Wrapf(err, "getting current Gas price")
	}
	ctx, cncl := context.WithTimeout(self.ctx, 3*time.Second)
	defer cncl()

	return self.reward.Current(ctx, slot, gasPrice, big.NewInt(int64(gasUsage)))
}

func (self *Submitter) Submit(newChallengeReplace context.Context, result *mining.Result) {
	go func(newChallengeReplace context.Context, result *mining.Result) {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-newChallengeReplace.Done():
				level.Info(self.logger).Log("msg", "pending submit canceled")
				return
			default:
			}

			self.blockUntilTimeToSubmit(newChallengeReplace)

			for {
				select {
				case <-newChallengeReplace.Done():
					level.Info(self.logger).Log("msg", "pending submit canceled")
					return
				default:
				}

				statusID, err := self.minerStatus()
				if err != nil {
					level.Info(self.logger).Log("msg", "getting miner status", "err", err)
					<-ticker.C
					continue
				}
				if statusID != 1 {
					level.Info(self.logger).Log("msg", "miner is not in a status that can submit", "status", minerStatusName(statusID))
					<-ticker.C
					continue
				}

				reqVals, err := self.addRrequestVals(result.Work.Challenge.RequestIDs)
				if err != nil {
					level.Error(self.logger).Log("msg", "adding the request ids, retrying", "err", err)
					<-ticker.C
					continue
				}

				slot, err := self.reward.Slot()
				if err != nil {
					level.Error(self.logger).Log("msg", "getting slot number", "err", err)
				}

				gasEstimate, err := contracts.EstimateGasUsageSubmitMiningSolution(
					self.ctx,
					self.contract,
					self.client,
					self.account.Address,
					result.Nonce,
				)
				if err != nil {
					level.Error(self.logger).Log("msg", "estimating the gas usage", "err", err)
					<-ticker.C
					continue
				}

				refund, err := self.reward.Refund(slot)
				if err != nil {
					level.Error(self.logger).Log("msg", "getting gas refund value", "slot", slot, "err", err)
				}

				gasEstimateAfterRefund := gasEstimate - refund

				if self.cfg.ProfitThreshold > 0 { // Profit check is enabled.
					profitPercent, err := self.profitPercent(slot.Uint64(), gasEstimateAfterRefund)
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
				}

				level.Info(self.logger).Log(
					"msg", "sending solution to the chain",
					"solutionNonce", result.Nonce,
					"IDs", fmt.Sprintf("%+v", result.Work.Challenge.RequestIDs),
					"vals", fmt.Sprintf("%+v", reqVals),
				)
				f := func(auth *bind.TransactOpts) (*types.Transaction, error) {
					return self.contract.SubmitMiningSolution(auth, result.Nonce, result.Work.Challenge.RequestIDs, reqVals)
				}
				tx, recieipt, err := self.transactor.Transact(newChallengeReplace, uint64(gasEstimate+100_000), f)
				select {
				case <-newChallengeReplace.Done():
					level.Info(self.logger).Log("msg", "pending submit canceled")
					return
				default:
				}
				if err != nil {
					self.submitFailCount.Inc()
					level.Error(self.logger).Log("msg", "submiting a solution", "err", err)
					return
				}

				if recieipt.Status != types.ReceiptStatusSuccessful {
					self.submitFailCount.Inc()
					level.Error(self.logger).Log("msg", "submiting solution status not success", "status", recieipt.Status, "hash", tx.Hash())
					return
				}

				// Cast to ints to avoid underflow when estimate is lower than actual gas usage.
				self.reward.RecordGasUsageRefund(gasEstimate-int64(recieipt.GasUsed), slot.Uint64())

				self.submitGasUsageEstimated.With(
					prometheus.Labels{
						"slot": slot.String(),
					},
				).(prometheus.Gauge).Set(float64(gasEstimateAfterRefund))

				self.submitGasUsageActual.With(
					prometheus.Labels{
						"slot": slot.String(),
					},
				).(prometheus.Gauge).Set(float64(recieipt.GasUsed))

				level.Info(self.logger).Log("msg", "successfully submited solution",
					"txHash", tx.Hash().String(),
					"nonce", tx.Nonce(),
					"slot", slot.String(),
					"gasPrice", tx.GasPrice(),
					"gasUsed", recieipt.GasUsed,
					"gasUsedEstimated", gasEstimateAfterRefund,
					"gasUsedRefund", refund,
					"gasDifference", int64(recieipt.GasUsed)-int64(gasEstimateAfterRefund),
					"gasLimit", tx.Gas(),
				)
				self.submitCount.Inc()

				for i, id := range result.Work.Challenge.RequestIDs {
					self.submitValue.With(
						prometheus.Labels{
							"id": id.String(),
						},
					).(prometheus.Gauge).Set(float64(reqVals[i].Int64()))
				}

				return
			}
		}
	}(newChallengeReplace, result)
}

func (self *Submitter) addRrequestVals(requestIDs [5]*big.Int) ([5]*big.Int, error) {
	var currentValues [5]*big.Int
	for i, reqID := range requestIDs {
		val, err := self.psr.GetValue(reqID.Int64(), time.Now())
		if err != nil {
			return currentValues, errors.Wrapf(err, "getting value for request ID:%v", reqID)
		}
		currentValues[i] = big.NewInt(int64(val))
	}
	return currentValues, nil
}

func (self *Submitter) minerStatus() (int64, error) {
	// Check if the staked account is in dispute before sending a transaction.
	statusID, _, err := self.contract.GetStakerInfo(&bind.CallOpts{}, self.account.Address)
	if err != nil {
		return 0, errors.Wrapf(err, "getting staker info from contract addr:%v", self.account.Address)
	}
	return statusID.Int64(), nil
}

func (self *Submitter) lastSubmit() (time.Duration, *time.Time, error) {
	address := "000000000000000000000000" + self.account.Address.Hex()[2:]
	decoded, err := hex.DecodeString(address)
	if err != nil {
		return 0, nil, errors.Wrapf(err, "decoding address")
	}
	last, err := self.contract.GetUintVar(nil, ethereum.Keccak256(decoded))

	if err != nil {
		return 0, nil, errors.Wrapf(err, "getting last submit time for:%v", self.account.Address.String())
	}
	// The Miner has never submitted so put a timestamp at the beginning of unix time.
	if last.Int64() == 0 {
		last.Set(big.NewInt(1))
	}

	lastInt := last.Int64()
	now := time.Now()
	var lastSubmit time.Duration
	var tm time.Time
	if lastInt > 0 {
		tm = time.Unix(lastInt, 0)
		lastSubmit = now.Sub(tm)
	}

	return lastSubmit, &tm, nil
}

func minerStatusName(statusID int64) string {
	// From https://github.com/tellor-io/tellor3/blob/7c2f38a0e3f96631fb0f96e0d0a9f73e7b355766/contracts/TellorStorage.sol#L41
	switch statusID {
	case 0:
		return "Not staked"
	case 1:
		return "Staked"
	case 2:
		return "LockedForWithdraw"
	case 3:
		return "OnDispute"
	case 4:
		return "ReadyForUnlocking"
	case 5:
		return "Unlocked"
	default:
		return "Unknown"
	}
}
