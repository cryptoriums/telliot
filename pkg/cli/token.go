// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"

	"github.com/cryptoriums/packages/math"
	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

type tokenCmd struct {
	Gas
	From   string  `required:""`
	To     string  `required:""`
	Amount float64 `arg:""`
}

func (self *tokenCmd) Validate() error {
	if self.To == "" {
		return nil
	}
	if !common.IsHexAddress(self.To) {
		return errors.Errorf("the address is not a hex string:%v", self.To)
	}
	if self.From == "" {
		return nil
	}
	if !common.IsHexAddress(self.From) {
		return errors.Errorf("the address is not a hex string:%v", self.From)
	}
	return nil
}

type TransferCmd tokenCmd

func (self *TransferCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, client, master, _, _, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	from := common.HexToAddress(self.From)

	balance, err := master.BalanceOf(&bind.CallOpts{Context: ctx}, from)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}
	level.Info(logger).Log("msg", "current balance", "amount", math.BigIntToFloatDiv(balance, params.Ether))

	amount := math.FloatToBigIntMul(self.Amount, params.Ether)
	if err != nil {
		return errors.Wrap(err, "invalid input amount")
	}
	if balance.Cmp(amount) < 0 {
		return errors.Errorf("insufficient balance TRB actual: %v, requested: %v",
			math.BigIntToFloatDiv(balance, params.Ether),
			math.BigIntToFloatDiv(amount, params.Ether))
	}

	acc, err := ethereum.GetAccountByPubAddress(logger, self.From, cfg.EnvVars)
	if err != nil {
		return errors.Wrap(err, "getting auth account")
	}
	opts, err := ethereum.PrepareTx(ctx, client, acc, self.GasPrice, 200_000)
	if err != nil {
		return errors.Wrap(err, "preparing ethereum transaction")
	}

	to := common.HexToAddress(self.To)

	tx, err := master.Transfer(opts, to, amount)
	if err != nil {
		return errors.Wrap(err, "calling transfer")
	}
	level.Info(logger).Log(
		"msg", "transferred",
		"amount", math.BigIntToFloatDiv(amount, params.Ether),
		"to", to.Hex()[:12],
		"tx", tx.Hash(),
	)
	return nil
}

type ApproveCmd tokenCmd

func (self *ApproveCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, client, master, _, _, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	from := common.HexToAddress(self.From)

	balance, err := master.BalanceOf(&bind.CallOpts{Context: ctx}, from)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	amount := math.FloatToBigIntMul(self.Amount, params.Ether)
	if err != nil {
		return errors.Wrap(err, "invalid input amount")
	}
	if balance.Cmp(amount) < 0 {
		return errors.Errorf("insufficient balance TRB actual: %v, requested: %v",
			math.BigIntToFloatDiv(balance, params.Ether),
			math.BigIntToFloatDiv(amount, params.Ether),
		)
	}

	acc, err := ethereum.GetAccountByPubAddress(logger, self.From, cfg.EnvVars)
	if err != nil {
		return errors.Wrap(err, "getting auth account")
	}

	opts, err := ethereum.PrepareTx(ctx, client, acc, self.GasPrice, 100_000)
	if err != nil {
		return errors.Wrap(err, "preparing ethereum transaction")
	}

	spender := common.HexToAddress(self.To)

	tx, err := master.Approve(opts, spender, amount)
	if err != nil {
		return errors.Wrap(err, "calling approve")
	}
	level.Info(logger).Log("msg", "approved", "amount", math.BigIntToFloatDiv(amount, params.Ether), "spender", spender.Hex()[:12], "tx", tx.Hash())
	return nil

}

type BalanceCmd struct {
	AccountArg
}

func (self *BalanceCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	_, client, master, _, _, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	addr := common.HexToAddress(self.Account)

	ethBalance, err := client.BalanceAt(ctx, addr, nil)
	if err != nil {
		return errors.Wrap(err, "get eth balance")
	}
	trbBalance, err := master.BalanceOf(&bind.CallOpts{Context: ctx}, addr)
	if err != nil {
		return errors.Wrapf(err, "getting trb balance")
	}

	level.Info(logger).Log(
		"msg", "balance check",
		"address", addr.Hex(),
		"ETH", math.BigIntToFloatDiv(ethBalance, params.Ether),
		"TRB", math.BigIntToFloatDiv(trbBalance, params.Ether),
	)
	return nil
}
