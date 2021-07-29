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
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/format"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/mining"
	psr "github.com/tellor-io/telliot/pkg/psr/tellor"
)

const ComponentName = "submitterTellor"

type Transactor interface {
	Transact(ctx context.Context, _nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error)
}

type Config struct {
	Enabled         bool
	LogLevel        string
	MinSubmitPeriod format.Duration `help:"The time limit between each submit for a staked miner."`
}

/**
* The submitter has one purpose: to either submit the solution on-chain
* or to reject it if the miner has already submitted a solution for the challenge
* or the the solution'self challenge does not match current challenge
 */

type Submitter struct {
	ctx      context.Context
	close    context.CancelFunc
	logger   log.Logger
	cfg      Config
	account  *ethereum.Account
	client   *ethclient.Client
	contract contracts.ContractCaller
	resultCh chan *mining.Result

	transactor Transactor
	psr        *psr.Psr
}

func New(
	ctx context.Context,
	logger log.Logger,
	cfg Config,
	client *ethclient.Client,
	contract contracts.ContractCaller,
	account *ethereum.Account,
	transactor Transactor,
	psr *psr.Psr,
) (*Submitter, chan *mining.Result, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)
	ctx, close := context.WithCancel(ctx)
	submitter := &Submitter{
		ctx:        ctx,
		close:      close,
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

				level.Info(self.logger).Log(
					"msg", "sending solution to the chain",
					"solutionNonce", result.Nonce,
					"IDs", fmt.Sprintf("%+v", result.Work.Challenge.RequestIDs),
					"vals", fmt.Sprintf("%+v", reqVals),
				)

				tx, err := self.transactor.Transact(newChallengeReplace, result.Nonce, result.Work.Challenge.RequestIDs, reqVals)
				if err != nil {
					level.Error(self.logger).Log("msg", "transacting", "err", err)
				}
				level.Info(self.logger).Log("msg", "successfully submited solution",
					"txHash", tx.Hash().String(),
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
