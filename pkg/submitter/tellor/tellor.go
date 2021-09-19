// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package tellor

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/format"
	"github.com/cryptoriums/telliot/pkg/logging"
	"github.com/cryptoriums/telliot/pkg/mining"
	psr "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

const ComponentName = "submitterTellor"

type Transactor interface {
	Transact(ctx context.Context, _nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error)
}

type Config struct {
	LogLevel        string
	MinSubmitPeriod format.Duration `help:"The time limit between each submit for a staked reporter."`
}

/**
* The submitter has one purpose: to either submit the solution on-chain
* or to reject it if the reporter has already submitted a solution for the challenge
* or the the solution'self challenge does not match current challenge
 */

type Submitter struct {
	ctx      context.Context
	stop     context.CancelFunc
	logger   log.Logger
	cfg      Config
	account  *ethereum.Account
	client   ethereum.EthClient
	contract contracts.TellorCaller
	resultCh chan *mining.Result

	transactor Transactor
	psr        *psr.Psr
}

func New(
	ctx context.Context,
	logger log.Logger,
	cfg Config,
	client ethereum.EthClient,
	contract contracts.TellorCaller,
	account *ethereum.Account,
	transactor Transactor,
	psr *psr.Psr,
) (*Submitter, chan *mining.Result, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)
	ctx, stop := context.WithCancel(ctx)
	submitter := &Submitter{
		ctx:        ctx,
		stop:       stop,
		client:     client,
		cfg:        cfg,
		resultCh:   make(chan *mining.Result),
		account:    account,
		logger:     logger,
		contract:   contract,
		transactor: transactor,
		psr:        psr,
	}

	return submitter, submitter.resultCh, nil
}

func (self *Submitter) Start() error {
	level.Info(self.logger).Log("msg", "starting", "logLevel", self.cfg.LogLevel)
	for {
		select {
		case <-self.ctx.Done():
			return self.ctx.Err()
		case result := <-self.resultCh:
			level.Info(self.logger).Log("msg", "received a solution",
				"challenge", fmt.Sprintf("%x", result.Work.Challenge),
				"solution", result.Nonce,
				"difficulty", result.Work.Challenge.Difficulty,
				"requestIDs", fmt.Sprintf("%+v", result.Work.Challenge.RequestIDs),
			)
			self.Submit(result.Work.Context, result)
		}
	}
}

func (self *Submitter) Stop() {
	self.stop()
}

func (self *Submitter) blockUntilTimeToSubmit(newChallengeReplace context.Context) {
	var (
		timeSincelastSubmit time.Duration
		timeOfLastSubmit    *time.Time
		err                 error
	)
	for {
		select {
		case <-newChallengeReplace.Done():
			level.Info(self.logger).Log("msg", "canceled pending submit while gettting last submit time")
		default:
		}
		timeSincelastSubmit, timeOfLastSubmit, err = contracts.LastSubmit(self.contract, self.account.Address)
		if err != nil {
			level.Debug(self.logger).Log("msg", "checking last submit time", "err", err)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	if timeSincelastSubmit < self.cfg.MinSubmitPeriod.Duration {
		nextSubmit := timeOfLastSubmit.Add(self.cfg.MinSubmitPeriod.Duration)
		level.Info(self.logger).Log("msg", "min transaction submit threshold hasn't passed",
			"nextSubmitTs", nextSubmit.Unix(),
			"nextSubmit", nextSubmit,
			"minSubmitPeriod", self.cfg.MinSubmitPeriod,
		)
		timeToSubmit, cncl := context.WithDeadline(newChallengeReplace, nextSubmit)
		defer cncl()
		select {
		case <-newChallengeReplace.Done():
			level.Info(self.logger).Log("msg", "canceled pending submit while waiting for the time to submit")
		case <-timeToSubmit.Done(): // 15min since last submit has passed so can unblock.
		}
	}
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

				statusID, _, err := self.contract.GetStakerInfo(&bind.CallOpts{Context: newChallengeReplace}, self.account.Address)
				if err != nil {
					level.Info(self.logger).Log("msg", "getting reporter status", "err", err)
					<-ticker.C
					continue
				}
				if statusID.Int64() != 1 {
					level.Info(self.logger).Log("msg", "reporter is not in a status that can submit", "status", contracts.ReporterStatusName(statusID.Int64()))
					<-ticker.C
					continue
				}

				reqVals, err := self.addRrequestVals(result.Work.Challenge.RequestIDs)
				if err != nil {
					level.Error(self.logger).Log("msg", "adding the request ids, retrying", "err", err)
					<-ticker.C
					continue
				}

				level.Info(self.logger).Log(
					"msg", "sending solution to the chain",
					"solutionNonce", result.Nonce,
					"IDs", fmt.Sprintf("%+v", result.Work.Challenge.RequestIDs),
					"vals", fmt.Sprintf("%+v", reqVals),
				)

				tx, err := self.transactor.Transact(newChallengeReplace, result.Nonce, result.Work.Challenge.RequestIDs, reqVals)
				if err != nil {
					level.Error(self.logger).Log("msg", "transacting", "err", err)
					return
				}
				level.Info(self.logger).Log("msg", "successfully submited solution",
					"txHash", tx.Hash().Hex(),
					"nonce", tx.Nonce(),
					"gasPrice", tx.GasPrice(),
					"gasLimit", tx.Gas(),
				)
				return
			}
		}
	}(newChallengeReplace, result)
}

func (self *Submitter) addRrequestVals(requestIDs [5]*big.Int) ([5]*big.Int, error) {
	var currentValues [5]*big.Int
	for i, reqID := range requestIDs {

		// For inactive values just zero values.
		// Otherwise the psr will return an error as it doesn't hold any value for those.
		if psr.IsInactive(reqID.Int64()) {
			currentValues[i] = big.NewInt(0)
			continue
		}

		val, err := self.psr.GetValue(reqID.Int64(), time.Now())
		if err != nil {
			return currentValues, errors.Wrapf(err, "getting value for request ID:%v", reqID)
		}
		currentValues[i] = big.NewInt(int64(val))
	}
	return currentValues, nil
}
