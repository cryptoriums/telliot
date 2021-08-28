// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"math/big"
	"time"

	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/math"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

type DepositCmd struct {
	GasAccount
	ProxyFlag
}

func (self *DepositCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	_, client, contract, err := ConfigClientContract(ctx, logger, cli.Config, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	account, err := ethereum.GetAccountByPubAddess(logger, self.Account)
	if err != nil {
		return err
	}

	addrToCheck := account.Address
	// When using proxy contract need to get its status.
	if self.Proxy != "" {
		addrToCheck = common.HexToAddress(self.Proxy)
	}

	balance, err := contract.BalanceOf(&bind.CallOpts{Context: ctx}, addrToCheck)
	if err != nil {
		return errors.Wrap(err, "get TRB balance")
	}

	status, startTime, err := contract.GetStakerInfo(&bind.CallOpts{Context: ctx}, addrToCheck)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}

	if status.Uint64() != 0 && status.Uint64() != 2 {
		PrintReporterStatus(logger, status, startTime)
		return nil
	}

	stakeAmt, err := contract.GetUintVar(nil, ethereum.Keccak256([]byte("_STAKE_AMOUNT")))
	if err != nil {
		return errors.Wrap(err, "fetching stake amount")
	}

	if balance.Cmp(stakeAmt) < 0 {
		return errors.Errorf("insufficient reporting stake TRB balance actual: %v, required:%v",
			math.BigIntToFloatDiv(balance, params.Ether),
			math.BigIntToFloatDiv(stakeAmt, params.Ether))
	}

	opts, err := ethereum.PrepareEthTransaction(ctx, client, account, self.GasBaseFee, self.GasTip, contracts.DepositStakeGasUsage)
	if err != nil {
		return errors.Wrap(err, "prepare ethereum transaction")
	}

	tx, err := contract.DepositStake(opts)
	if err != nil {
		return errors.Wrap(err, "contract failed")
	}
	level.Info(logger).Log("msg", "stake depositied", "tx", tx.Hash())
	return nil
}

type WithdrawCmd struct {
	GasAccount
}

func (self *WithdrawCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	_, client, contract, err := ConfigClientContract(ctx, logger, cli.Config, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	account, err := ethereum.GetAccountByPubAddess(logger, self.Account)
	if err != nil {
		return err
	}

	status, startTime, err := contract.GetStakerInfo(nil, account.Address)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}
	if status.Uint64() != 2 {
		level.Info(logger).Log("msg", "can't withdraw")
		PrintReporterStatus(logger, status, startTime)
		return nil
	}

	opts, err := ethereum.PrepareEthTransaction(ctx, client, account, self.GasBaseFee, self.GasTip, contracts.WithdrawStakeGasUsage)
	if err != nil {
		return errors.Wrap(err, "prepare ethereum transaction")
	}

	tx, err := contract.WithdrawStake(opts)
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
	_, client, contract, err := ConfigClientContract(ctx, logger, cli.Config, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	account, err := ethereum.GetAccountByPubAddess(logger, self.Account)
	if err != nil {
		return err
	}

	status, startTime, err := contract.GetStakerInfo(nil, account.Address)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}
	if status.Uint64() != 1 {
		PrintReporterStatus(logger, status, startTime)
		return nil
	}

	opts, err := ethereum.PrepareEthTransaction(ctx, client, account, self.GasBaseFee, self.GasTip, contracts.RequestStakingWithdrawGasUsage)
	if err != nil {
		return errors.Wrap(err, "prepare ethereum transaction")
	}

	tx, err := contract.RequestStakingWithdraw(opts)
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
	_, _, contract, err := ConfigClientContract(ctx, logger, cli.Config, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	addr := common.HexToAddress(self.Account)

	status, startTime, err := contract.GetStakerInfo(&bind.CallOpts{Context: ctx}, addr)
	if err != nil {
		return errors.Wrap(err, "get stake status")
	}

	PrintReporterStatus(logger, status, startTime)
	return nil
}

func PrintReporterStatus(logger log.Logger, statusID *big.Int, started *big.Int) {

	level.Info(logger).Log("msg", "reporter status", "name", contracts.ReporterStatusName(statusID.Int64()))

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