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
	"github.com/cryptoriums/telliot/pkg/logging"
	"github.com/cryptoriums/telliot/pkg/prompt"
	psrTellor "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
)

type SubmitCmd struct {
	config *config.Config
	Gas
	ContractFlag
	AccountArgOptional
	SkipConfirm bool `help:"submit without confirming, useful for testing"`
}

func (self *SubmitCmd) SetConfig(config *config.Config) {
	self.config = config
}

func (self *SubmitCmd) Run() error {
	logger := logging.NewLogger()
	ctx := context.Background()

	client, netID, err := ethereum.NewClient(logger, ctx)
	if err != nil {
		return errors.Wrap(err, "creating ethereum client")
	}

	var accounts []*ethereum.Account
	if self.Account != "" {
		acc, err := ethereum.GetAccountByPubAddess(self.Account)
		if err != nil {
			return errors.Wrap(err, "getting accounts")
		}
		accounts = append(accounts, acc)
	} else {
		accounts, err = ethereum.GetAccounts()
		if err != nil {
			return errors.Wrap(err, "getting accounts")
		}
	}

	address := accounts[0].Address
	if len(accounts) > 1 {
		level.Warn(logger).Log("msg", "multiple accounts loaded, but will use the first one on the list", "address", address)
	}

	var contractCaller contracts.ContractCaller
	if self.Contract != "" {
		contractCaller, err = contracts.NewITellorWithAddr(logger, ctx, common.HexToAddress(self.Contract), client, netID, contracts.DefaultParams)
		if err != nil {
			return errors.Wrap(err, "creating contract interface")
		}
	} else {
		contractCaller, err = contracts.NewITellor(logger, ctx, client, netID, contracts.DefaultParams)
		if err != nil {
			return errors.Wrap(err, "creating contract interface")
		}
	}

	resp, err := contractCaller.GetNewCurrentVariables(&bind.CallOpts{Context: ctx})
	if err != nil {
		return errors.Wrap(err, "get current DATA ids")
	}

	vals, err := self.getValuesFromDB(logger, ctx, resp.RequestIds)
	if err != nil {
		level.Error(logger).Log("msg", "couldn't get values from the DB so will proceed with manual submit", "err", err)
		vals = getValuesFromInput(logger, resp.RequestIds)
	}

	shouldContinue := finalPrompt(logger, self.SkipConfirm, resp.RequestIds, vals)
	if !shouldContinue {
		return nil
	}

	level.Info(logger).Log(
		"msg", "submitting",
		"ids", fmt.Sprintf("%+v", resp.RequestIds),
		"vals", fmt.Sprintf("%+v", vals),
	)

	opts, err := ethereumT.PrepareEthTransaction(ctx, client, accounts[0], self.GasBaseFee, self.GasTip, contracts.SubmitMiningSolutionGasUsage)
	if err != nil {
		return errors.Wrapf(err, "prepare ethereum transaction")
	}

	tx, err := contractCaller.SubmitMiningSolution(opts, "xxx", resp.RequestIds, vals)
	if err != nil {
		return errors.Wrap(err, "creting TX")
	}
	level.Info(logger).Log(
		"msg", "complete",
		"tx", tx.Hash(),
	)

	return nil
}

func getValuesFromInput(logger log.Logger, dataIDs [5]*big.Int) [5]*big.Int {
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
			_val, err := prompt.Prompt("Data ID:"+dataID.String()+" Val:", false)
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

func (self *SubmitCmd) getValuesFromDB(logger log.Logger, ctx context.Context, dataIDs [5]*big.Int) ([5]*big.Int, error) {
	var vals [5]*big.Int
	var querable storage.SampleAndChunkQueryable
	var err error
	if self.config.Db.RemoteHost != "" {
		querable, err = db.NewRemoteDB(self.config.Db)
		if err != nil {
			return vals, errors.Wrap(err, "coudn't open remote DB")
		}
		level.Info(logger).Log("msg", "connected to remote DB",
			"host", self.config.Db.RemoteHost,
			"port", self.config.Db.RemotePort,
		)
	} else {
		tsdbOptions := tsdb.DefaultOptions()
		tsdbOptions.NoLockfile = true
		querable, err = tsdb.Open(self.config.Db.Path, nil, nil, tsdbOptions)
		if err != nil {
			return vals, errors.Wrap(err, "coudn't open local DB")
		}
	}
	aggregator, err := aggregator.New(logger, ctx, self.config.Aggregator, querable)
	if err != nil {
		return vals, errors.Wrap(err, "couldn't create aggregator")

	}

	psr := psrTellor.New(logger, self.config.PsrTellor, aggregator)
	vals, err = requestVals(logger, psr, dataIDs)
	if err != nil {
		return vals, errors.Wrap(err, "getting variable values")
	}
	return vals, nil
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

func finalPrompt(logger log.Logger, skipConfirm bool, dataIDs, vals [5]*big.Int) bool {
	//lint:ignore faillint for prompts can't use logs.
	fmt.Println("Here are the final values before applying the default granularity of :" + strconv.Itoa(psrTellor.DefaultGranularity))

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)

	var dataIDstr, dataValstr string
	for i, id := range dataIDs {
		dataIDstr += id.String() + "\t"
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
		vals = getValuesFromInput(logger, dataIDs)
		return finalPrompt(logger, skipConfirm, dataIDs, vals)
	}
	return false
}
