// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/cryptoriums/packages/ethereum"
	ethereum_t "github.com/cryptoriums/packages/ethereum"
	math_t "github.com/cryptoriums/packages/math"
	"github.com/cryptoriums/packages/prompt"
	"github.com/cryptoriums/telliot/pkg/aggregator"
	"github.com/cryptoriums/telliot/pkg/config"
	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/db"
	psr_tellor "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
)

type SubmitCmd struct {
	DataID int64 `required:""  help:"the request id to submit"`
	Gas
	AccountArgOptional
	CustomSubmit bool
	SkipConfirm  bool `help:"submit without confirming, useful for testing"`
}

func (self *SubmitCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, client, master, oracle, _, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	account, err := self.SelectAccount(ctx, logger, client, master, oracle, cfg.EnvVars)
	if err != nil {
		return err
	}

	_, aggr, err := self.CreateAggr(ctx, logger, cfg)
	if err != nil {
		return errors.Wrap(err, "creating an aggregator")
	}

	psr, err := psr_tellor.PsrByID(self.DataID)
	if err != nil {
		return errors.Wrap(err, "getting psr")
	}

	val, err := self.GetValueFromDB(ctx, logger, cfg, aggr, psr_tellor.IntToQueryID(self.DataID))
	if err != nil {
		level.Error(logger).Log("msg", "couldn't get values from the DB so will proceed with manual submit", "err", err)
		val = GetValueFromInput(logger, psr)
	}

	shouldContinue := FinalPrompt(ctx, logger, oracle, self.SkipConfirm, self.GasPrice, psr, val)
	if !shouldContinue {
		return errors.New("canceled")
	}

	return self.Submit(cli, ctx, logger, account, client, oracle, self.DataID, val)
}

func (self *SubmitCmd) Submit(
	cli *CLI,
	ctx context.Context,
	logger log.Logger,
	account *ethereum.Account,
	client ethereum.EthClient,
	contract contracts.TellorOracleCaller,
	id int64,
	val float64,
) error {
	queryID := psr_tellor.IntToQueryID(id)

	nonce, err := contract.GetTimestampCountById(&bind.CallOpts{Context: ctx}, queryID)
	if err != nil {
		return errors.Wrap(err, "get current DATA ids")
	}

	opts, err := ethereum_t.PrepareTxOpts(ctx, client, account, self.GasPrice, contracts.SubmitGasLimit)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}

	level.Info(logger).Log(
		"msg", "submitting",
		"account", account.Address.Hex()[:8],
		"id", fmt.Sprintf("%+v", id),
		"val", fmt.Sprintf("%.0f", val),
	)

	var tx *types.Transaction
	if cli.Contract != "" && self.CustomSubmit {
		// TODO re-implement if needed.
		// contractP, err := contracts.NewITellorProxy(ctx, common.HexToAddress(cli.Contract), client)
		// if err != nil {
		// 	return err
		// }
	} else {
		tx, err = contract.SubmitValue(opts, queryID, math_t.FloatToBigInt(val).Bytes(), nonce, psr_tellor.NewQueryData(self.DataID))
		if err != nil {
			return errors.Wrap(err, "creting TX")
		}
	}
	level.Info(logger).Log(
		"msg", "complete",
		"tx", tx.Hash(),
	)
	return nil
}

func (self *SubmitCmd) SelectAccount(
	ctx context.Context,
	logger log.Logger,
	client ethereum.EthClient,
	master contracts.TellorMasterCaller,
	oracle contracts.TellorOracleCaller,
	envVars map[string]string,
) (*ethereum.Account, error) {
	var accounts []*ethereum.Account
	var err error
	if self.Account != "" {
		acc, err := ethereum.GetAccountByPubAddress(logger, self.Account, envVars)
		if err != nil {
			return nil, errors.Wrap(err, "getting accounts")
		}
		accounts = append(accounts, acc)
	} else {
		accounts, err = ethereum.GetAccounts(logger, envVars)
		if err != nil {
			return nil, errors.Wrap(err, "getting accounts")
		}
	}

	if len(accounts) > 1 {
		// Print the accounts details and prompt for a selection.
		PrintAccounts(ctx, logger, accounts, client, master, oracle)
		_accIndex, err := prompt.Prompt("select an account from 1 to "+strconv.Itoa(len(accounts))+":", false)
		if err != nil {
			return nil, errors.Wrap(err, "selecting an accounts")
		}
		accIndex, err := strconv.Atoi(_accIndex)
		if err != nil {
			return nil, errors.Wrap(err, "parsing account index selection")
		}

		accIndex-- // So the the selection starts from one
		if accIndex < 0 || accIndex > len(accounts) {
			return nil, errors.New("account selection not in range")
		}

		acc := accounts[accIndex]
		level.Info(logger).Log("msg", "selected account", "addr", acc.Address.Hex()[:8])
		return acc, nil

	}
	return accounts[0], nil
}

func GetValueFromInput(logger log.Logger, psr psr_tellor.PsrID) float64 {
	//lint:ignore faillint for prompts can't use logs.
	fmt.Println("Enter values in the format (1.123456)")
	for {
		_val, err := prompt.Prompt(PSRDetails(psr)+" Val:", false)
		if err != nil {
			//lint:ignore faillint for prompts can't use logs.
			fmt.Println(err)
			continue
		}
		val, err := strconv.ParseFloat(_val, 64)
		if err != nil {
			//lint:ignore faillint for prompts can't use logs.
			fmt.Println(err)
			continue
		}
		return val * float64(psr_tellor.DefaultGranularity)
	}
}

func (self *SubmitCmd) GetValueFromDB(
	ctx context.Context,
	logger log.Logger,
	cfg *config.Config,
	aggregator *aggregator.Aggregator,
	queryID [32]byte,
) (float64, error) {
	psr := psr_tellor.New(logger, cfg.PsrTellor, aggregator)

	val, err := psr.GetValue(queryID, time.Now())
	if err != nil {
		return 0, errors.Wrapf(err, "getting value for queryID:%v", queryID)
	}

	return val, nil
}

func (self *SubmitCmd) CreateAggr(ctx context.Context, logger log.Logger, cfg *config.Config) (storage.Queryable, *aggregator.Aggregator, error) {
	var querable storage.Queryable
	var err error
	if cfg.Db.RemoteHost != "" {
		querable, err = db.NewRemoteDB(cfg.Db)
		if err != nil {
			return nil, nil, errors.Wrap(err, "coudn't open remote DB")
		}
		level.Info(logger).Log("msg", "connected to remote DB",
			"host", cfg.Db.RemoteHost,
			"port", cfg.Db.RemotePort,
		)
	} else {
		tsdbOptions := tsdb.DefaultOptions()
		querable, err = tsdb.Open(cfg.Db.Path, nil, nil, tsdbOptions, tsdb.NewDBStats())
		if err != nil {
			return nil, nil, errors.Wrap(err, "coudn't open local DB")
		}
	}
	aggregator, err := aggregator.New(ctx, logger, cfg.Aggregator, querable)
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't create aggregator")

	}

	return querable, aggregator, nil
}

func FinalPrompt(
	ctx context.Context,
	logger log.Logger,
	contract contracts.TellorOracleCaller,
	skipConfirm bool,
	gasMaxFee float64,
	psr psr_tellor.PsrID,
	val float64,
) bool {
	timeOfLastNewValue, err := contract.TimeOfLastNewValue(&bind.CallOpts{Context: ctx})
	if err != nil {
		timeOfLastNewValue = big.NewInt(0)
		level.Error(logger).Log("msg", "getting last submit time", "err", err)
	}

	//lint:ignore faillint for prompts can't use logs.
	fmt.Printf(">>>>>>>> GasPrice:%v LastSubmit:%v \n", gasMaxFee, time.Since(time.Unix(timeOfLastNewValue.Int64(), 0)))
	//lint:ignore faillint for prompts can't use logs.
	fmt.Println("Here are the final values before applying the default granularity of :" + strconv.Itoa(psr_tellor.DefaultGranularity))

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "Psr:\t"+PSRDetails(psr)+"\t")
	fmt.Fprintln(w, "Val:\t"+fmt.Sprintf("%g", float64(val)/float64(psr_tellor.DefaultGranularity))+"\t")
	w.Flush()

	if skipConfirm {
		return true
	}

	promptResp, err := prompt.Prompt("Press Y to continue with the submit:", false)
	if err == nil && strings.ToLower(promptResp) == "y" {
		return true
	}
	promptResp, err = prompt.Prompt("Press Y if you want to enter values manually?:", false)
	if err == nil && strings.ToLower(promptResp) == "y" {
		val = GetValueFromInput(logger, psr)
		return FinalPrompt(ctx, logger, contract, skipConfirm, gasMaxFee, psr, val)
	}
	return false
}

func PSRDetails(psr psr_tellor.PsrID) string {
	return psr.Pair + ":" + psr.Aggr + ", inactive:" + strconv.FormatBool(psr.Inactive)
}
