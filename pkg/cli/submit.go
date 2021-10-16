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

	"github.com/cryptoriums/telliot/pkg/aggregator"
	"github.com/cryptoriums/telliot/pkg/config"
	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/db"
	"github.com/cryptoriums/telliot/pkg/ethereum"
	ethereumT "github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/prompt"
	psrTellor "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
)

type SubmitCmd struct {
	Gas
	AccountArgOptional
	CustomSubmit bool
	SkipConfirm  bool `help:"submit without confirming, useful for testing"`
}

func (self *SubmitCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, client, contract, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	resp, err := contract.GetNewCurrentVariables(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "get current DATA ids")
	}

	_, aggr, err := self.CreateAggr(ctx, logger, cfg)
	if err != nil {
		return errors.Wrap(err, "couldn't get create an aggregator")
	}

	ids := resp.RequestIds
	vals, err := self.GetValuesFromDB(ctx, logger, cfg, aggr, ids)
	if err != nil {
		level.Error(logger).Log("msg", "couldn't get values from the DB so will proceed with manual submit", "err", err)
		vals = GetValuesFromInput(logger, resp.RequestIds)
	}

	shouldContinue := FinalPrompt(ctx, logger, contract, self.SkipConfirm, self.GasPrice, resp.RequestIds, vals)
	if !shouldContinue {
		return errors.New("canceled")
	}

	account, err := self.SelectAccount(logger)
	if err != nil {
		return err
	}

	return self.Submit(cli, ctx, logger, account, client, contract, ids, vals)
}

func (self *SubmitCmd) Submit(
	cli *CLI,
	ctx context.Context,
	logger log.Logger,
	account *ethereum.Account,
	client ethereum.EthClient,
	contract contracts.TellorCaller,
	ids [5]*big.Int,
	vals [5]*big.Int,
) error {

	opts, err := ethereumT.PrepareTx(ctx, client, account, self.GasPrice, contracts.SubmitMiningSolutionGasLimit)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}

	level.Info(logger).Log(
		"msg", "submitting",
		"account", account.Address.Hex()[:8],
		"ids", fmt.Sprintf("%+v", ids),
		"vals", fmt.Sprintf("%+v", vals),
	)

	var tx *types.Transaction
	if cli.Contract != "" && self.CustomSubmit {
		contractP, err := contracts.NewITellorProxy(ctx, common.HexToAddress(cli.Contract), client)
		if err != nil {
			return err
		}
		tx, err = contractP.SubmitMiningSolution0(opts, "krasi!", ids, vals, big.NewInt(0))
		if err != nil {
			return errors.Wrap(err, "creting TX")
		}
	} else {
		tx, err = contract.SubmitMiningSolution(opts, "krasi!", ids, vals)
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

func (self *SubmitCmd) SelectAccount(logger log.Logger) (*ethereum.Account, error) {
	var accounts []*ethereum.Account
	var err error
	if self.Account != "" {
		acc, err := ethereum.GetAccountByPubAddress(logger, self.Account)
		if err != nil {
			return nil, errors.Wrap(err, "getting accounts")
		}
		accounts = append(accounts, acc)
	} else {
		accounts, err = ethereum.GetAccounts(logger)
		if err != nil {
			return nil, errors.Wrap(err, "getting accounts")
		}
	}

	if len(accounts) > 1 {
		level.Warn(logger).Log("msg", "multiple accounts loaded, but will use the first one on the list", "address", accounts[0].Address)
	}
	return accounts[0], nil
}

func GetValuesFromInput(logger log.Logger, dataIDs [5]*big.Int) [5]*big.Int {
	var vals [5]*big.Int
	for i, dataID := range dataIDs {
		if psrTellor.IsInactive(dataID.Int64()) {
			level.Warn(logger).Log("msg", "dataID is inactive so adding 0 value", "datID", dataID.Int64())
			vals[i] = big.NewInt(0)
			continue
		}
		//lint:ignore faillint for prompts can't use logs.
		fmt.Println("Enter values in the format (1.123456)")
		for {
			_val, err := prompt.Prompt(dataID.String()+":"+DataIdFullName(dataID.Int64())+" Val:", false)
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
			vals[i] = big.NewInt(int64(val * float64(psrTellor.DefaultGranularity)))
			break
		}
	}
	return vals
}

func (self *SubmitCmd) GetValuesFromDB(ctx context.Context, logger log.Logger, cfg *config.Config, aggregator *aggregator.Aggregator, dataIDs [5]*big.Int) ([5]*big.Int, error) {
	psr := psrTellor.New(logger, cfg.PsrTellor, aggregator)
	vals, err := requestVals(logger, psr, dataIDs)
	if err != nil {
		return vals, errors.Wrap(err, "getting variable values")
	}
	return vals, nil
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

func requestVals(logger log.Logger, psr *psrTellor.Psr, dataIDs [5]*big.Int) ([5]*big.Int, error) {
	var vals [5]*big.Int
	for i, dataID := range dataIDs {
		if psrTellor.IsInactive(dataID.Int64()) {
			level.Warn(logger).Log("msg", "dataID is inactive so adding 0 value", "datID", dataID.Int64())
			vals[i] = big.NewInt(0)
			continue
		}
		val, err := psr.GetValue(dataID.Int64(), time.Now())
		if err != nil {
			return vals, errors.Wrapf(err, "getting value for request ID:%v", dataID)
		}
		vals[i] = big.NewInt(int64(val))
	}
	return vals, nil
}

func FinalPrompt(
	ctx context.Context,
	logger log.Logger,
	contract contracts.TellorCaller,
	skipConfirm bool,
	gasMaxFee float64,
	dataIDs,
	vals [5]*big.Int,
) bool {

	slot, err := contract.GetUintVar(&bind.CallOpts{Context: ctx}, ethereum.Keccak256([]byte("_SLOT_PROGRESS")))
	if err != nil {
		level.Error(logger).Log("msg", "getting current slot", "err", err)
	}

	timeOfLastNewValue, err := contract.GetUintVar(nil, ethereum.Keccak256([]byte("_TIME_OF_LAST_NEW_VALUE")))
	if err != nil {
		level.Error(logger).Log("msg", "getting last submit time", "err", err)
	}

	//lint:ignore faillint for prompts can't use logs.
	fmt.Printf(">>>>>>>> Current Slot:%v GasPrice:%v LastSubmit:%v \n", slot.String(), gasMaxFee, time.Since(time.Unix(timeOfLastNewValue.Int64(), 0)))
	//lint:ignore faillint for prompts can't use logs.
	fmt.Println("Here are the final values before applying the default granularity of :" + strconv.Itoa(psrTellor.DefaultGranularity))

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)

	var dataIDstr, dataValstr string
	for i, id := range dataIDs {
		dataIDstr += id.String() + ":" + DataIdFullName(id.Int64()) + "\t"
		dataValstr += fmt.Sprintf("%g", float64(vals[i].Int64())/float64(psrTellor.DefaultGranularity)) + "\t"
	}

	fmt.Fprintln(w, "IDs:\t"+dataIDstr)
	fmt.Fprintln(w, "Vals:\t"+dataValstr)
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
		vals = GetValuesFromInput(logger, dataIDs)
		return FinalPrompt(ctx, logger, contract, skipConfirm, gasMaxFee, dataIDs, vals)
	}
	return false
}

func DataIdFullName(id int64) string {
	return psrTellor.Psrs[id].Pair + "(" + psrTellor.Psrs[id].Aggr + ")"
}
