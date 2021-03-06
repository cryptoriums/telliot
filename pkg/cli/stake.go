// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"math/big"
	"time"

	"github.com/cryptoriums/packages/ethereum"
	"github.com/cryptoriums/packages/math"
	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

type DepositCmd struct {
	GasAccount
	NoChks
}

func (self *DepositCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, client, master, _, _, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}
	defer client.Close()

	account, err := ethereum.GetAccountByPubAddress(logger, self.Account, cfg.EnvVars)
	if err != nil {
		return err
	}

	addrToCheck := account.PublicKey
	// When using proxy contract need to get its status.
	if cli.Contract != "" {
		addrToCheck = common.HexToAddress(cli.Contract)
	}

	if !self.NoChecks {
		balance, err := master.BalanceOf(&bind.CallOpts{Context: ctx}, addrToCheck)
		if err != nil {
			return errors.Wrap(err, "get TRB balance")
		}

		status, startTime, err := master.GetStakerInfo(&bind.CallOpts{Context: ctx}, addrToCheck)
		if err != nil {
			return errors.Wrap(err, "get stake status")
		}

		if status.Uint64() != 0 && status.Uint64() != 2 {
			PrintReporterStatus(logger, status, startTime)
			return nil
		}

		stakeAmt, err := master.GetUintVar(nil, crypto.Keccak256Hash([]byte("_STAKE_AMOUNT")))
		if err != nil {
			return errors.Wrap(err, "fetching stake amount")
		}

		if balance.Cmp(stakeAmt) < 0 {
			return errors.Errorf("insufficient reporting stake TRB balance actual: %v, required:%v",
				math.BigIntToFloatDiv(balance, params.Ether),
				math.BigIntToFloatDiv(stakeAmt, params.Ether))
		}
	}

	opts, err := ethereum.NewTxOpts(ctx, client, account, self.GasPrice, contracts.StakeDepositGasLimit)
	if err != nil {
		return errors.Wrap(err, "prepare ethereum transaction")
	}

	tx, err := master.DepositStake(opts)
	if err != nil {
		return errors.Wrap(err, "contract failed")
	}
	level.Info(logger).Log("msg", "stake depositied tx created", "hash", tx.Hash())
	return nil
}

type WithdrawCmd struct {
	GasAccount
}

func (self *WithdrawCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, client, master, _, _, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}
	defer client.Close()

	account, err := ethereum.GetAccountByPubAddress(logger, self.Account, cfg.EnvVars)
	if err != nil {
		return err
	}

	status, startTime, err := master.GetStakerInfo(nil, account.PublicKey)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}
	if status.Uint64() != 2 {
		level.Info(logger).Log("msg", "can't withdraw")
		PrintReporterStatus(logger, status, startTime)
		return nil
	}

	opts, err := ethereum.NewTxOpts(ctx, client, account, self.GasPrice, contracts.StakeWithdrawGasLimit)
	if err != nil {
		return errors.Wrap(err, "prepare ethereum transaction")
	}

	tx, err := master.WithdrawStake(opts)
	if err != nil {
		return errors.Wrap(err, "contract")
	}
	level.Info(logger).Log("msg", "withdrew stake", "txHash", tx.Hash().Hex())

	return nil
}

type RequestCmd struct {
	GasAccount
}

func (self *RequestCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, client, master, _, _, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}
	defer client.Close()

	account, err := ethereum.GetAccountByPubAddress(logger, self.Account, cfg.EnvVars)
	if err != nil {
		return err
	}

	status, startTime, err := master.GetStakerInfo(nil, account.PublicKey)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}
	if status.Uint64() != 1 {
		PrintReporterStatus(logger, status, startTime)
		return nil
	}

	opts, err := ethereum.NewTxOpts(ctx, client, account, self.GasPrice, contracts.StakeRequestWithdrawGasLimit)
	if err != nil {
		return errors.Wrap(err, "prepare ethereum transaction")
	}

	tx, err := master.RequestStakingWithdraw(opts)
	if err != nil {
		return errors.Wrap(err, "contract")
	}

	level.Info(logger).Log("msg", "withdrawal request sent", "txHash", tx.Hash().Hex())

	return nil
}

type StatusCmd struct {
	AccountArg
}

func (self *StatusCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	_, client, master, _, _, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}
	defer client.Close()

	addr := common.HexToAddress(self.Account)

	status, startTime, err := master.GetStakerInfo(&bind.CallOpts{Context: ctx}, addr)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}

	PrintReporterStatus(logger, status, startTime)
	return nil
}

func PrintReporterStatus(logger log.Logger, statusID *big.Int, started *big.Int) {

	level.Info(logger).Log("msg", "reporter status", "status", contracts.ReporterStatusName(statusID.Int64()))

	stakeTime := time.Unix(started.Int64(), 0)
	switch status := statusID.Int64(); status {
	case 1:
		level.Info(logger).Log("msg", "staked since", "UTC", stakeTime.UTC())
	case 2:
		startedRound := started.Int64()
		startedRound = ((startedRound + 86399) / 86400) * 86400
		target := time.Unix(startedRound, 0)
		timePassed := time.Since(target)
		delta := timePassed - (time.Hour * 24 * 7)
		if delta > 0 {
			level.Info(logger).Log("msg", "stake has been eligbile to withdraw for", "delta", delta)
		} else {
			level.Info(logger).Log("msg", "stake will be eligible to withdraw in", "delta", -delta)
		}
	}
}
