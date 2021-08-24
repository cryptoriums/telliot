// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

type DataCmd struct {
	From     uint64        `help:"the unix timestamp to use as a starting point for the data retrieval"`
	LookBack time.Duration `default:"2h" help:"how far to lookback"`
}

func (self *DataCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	_, client, contract, err := ConfigClientContract(ctx, logger, cli.Config, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	level.Info(logger).Log("msg", "params", "from", self.From, "lookBack", self.LookBack)

	submits, err := contracts.GetSubmitLogs(ctx, client, contract, int64(self.From), self.LookBack)
	if err != nil {
		return errors.Wrap(err, "getting submit logs")
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)

	for _, submit := range submits {
		for i, id := range submit.DataIDs {
			if i == 0 {
				fmt.Fprintf(w, "ts: %v\t0\t1\tmedian:2\t3\t4\t \n",
					submit.Timestamp,
				)
			}
			fmt.Fprintf(w, "%v:%v id:%v\t",
				tellor.Psrs[id.Int64()].Pair,
				tellor.Psrs[id.Int64()].Aggr,
				id.Int64(),
			)
			for i, val := range submit.Values[id.Int64()] {
				//lint:ignore faillint looks cleaner with print instead of logs
				fmt.Fprintf(w, "%.6f:%v\t", float64(val.Int64())/tellor.DefaultGranularity, submit.Reporters[id.Int64()][i].Hex()[:8])
			}
			fmt.Fprint(w, "\n")

		}
		fmt.Fprint(w, "\n")
		w.Flush()
	}

	return nil
}
