// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"text/tabwriter"
	"time"

	"github.com/cryptoriums/telliot/pkg/contracts"
	math_t "github.com/cryptoriums/telliot/pkg/math"
	"github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

type DataCmd struct {
	From     uint64        `help:"the unix timestamp to use as a starting point for the data retrieval"`
	LookBack time.Duration `default:"2h" help:"how far to lookback"`
}

func (self *DataCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	_, client, _, oracle, _, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	level.Info(logger).Log("msg", "params", "from", self.From, "lookBack", self.LookBack)

	submits, err := contracts.GetSubmitLogs(ctx, client, oracle, int64(self.From), self.LookBack)
	if err != nil {
		return errors.Wrap(err, "getting submit logs")
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	for _, submit := range submits {
		psr, err := tellor.PsrByQueryBytes(submit.QueryData)
		if err != nil {
			return errors.Wrapf(err, "getting PSR by query")
		}
		inactive := ""
		if psr.Inactive {
			inactive = "(in)"
		}

		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Fprintf(w, "ts: %v \tmins: %v \tReporter: %v \tID%v:%v \tName %v:%v \tVal: %.6f \t \n",
			submit.Time,
			int(time.Since(time.Unix(submit.Time.Int64(), 0)).Minutes()),
			submit.Reporter.Hex()[:8],
			psr.QueryHash, inactive,
			psr.Pair, psr.Aggr,
			math_t.BigIntToFloat(big.NewInt(0).SetBytes(submit.Value))/tellor.DefaultGranularity,
		)
	}
	w.Flush()

	return nil
}
