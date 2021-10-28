// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"net/http"
	"os"
	"syscall"

	"github.com/cryptoriums/telliot/pkg/aggregator"
	"github.com/cryptoriums/telliot/pkg/config"
	"github.com/cryptoriums/telliot/pkg/contracts"
	psr_tellor "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/cryptoriums/telliot/pkg/tracker/blocks"
	"github.com/cryptoriums/telliot/pkg/tracker/index"
	"github.com/cryptoriums/telliot/pkg/web"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/oklog/run"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/tsdb"
)

type DataserverCmd struct {
}

func (self *DataserverCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, err := config.LoadConfig(ctx, logger, cli.Config, true)
	if err != nil {
		return err
	}
	// We define our run groups here.
	var g run.Group
	// Run groups.
	{
		// Handle interupts.
		g.Add(run.SignalHandler(ctx, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM))

		// Open the TSDB database.
		tsdbOptions := tsdb.DefaultOptions()

		tsdbOptions.RetentionDuration = cfg.Db.Retention.Milliseconds()
		if err := os.MkdirAll(cfg.Db.Path, 0777); err != nil {
			return errors.Wrap(err, "creating tsdb DB folder")
		}
		tsDB, err := tsdb.Open(cfg.Db.Path, nil, nil, tsdbOptions, tsdb.NewDBStats())
		if err != nil {
			return errors.Wrap(err, "creating tsdb DB")
		}
		level.Info(logger).Log("msg", "opened local db", "path", cfg.Db.Path)

		defer func() {
			if err := tsDB.Close(); err != nil {
				level.Error(logger).Log("msg", "closing the tsdb", "err", err)
			}
		}()

		trackerIndex, err := index.New(ctx, logger, cfg.TrackerIndex, tsDB)
		if err != nil {
			return errors.Wrap(err, "creating component:"+index.ComponentName)
		}

		g.Add(func() error {
			err := trackerIndex.Start()
			level.Info(logger).Log("msg", "shutdown complete", "component", index.ComponentName)
			return err
		}, func(error) {
			trackerIndex.Stop()
		})

		cfg, client, _, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
		if err != nil {
			return err
		}

		trackerBlocks, err := blocks.New(ctx, logger, cfg.TrackerBlocks, tsDB, client)
		if err != nil {
			return errors.Wrap(err, "creating com:"+blocks.ComponentName)
		}

		g.Add(func() error {
			trackerBlocks.Start()
			level.Info(logger).Log("msg", "shutdown complete", "component", blocks.ComponentName)
			return err
		}, func(error) {
			trackerBlocks.Stop()
		})

		// Web/Api server.
		{
			// Aggregator.
			aggregator, err := aggregator.New(ctx, logger, cfg.Aggregator, tsDB)
			if err != nil {
				return errors.Wrap(err, "creating aggregator")
			}
			handlers := make(map[string]http.HandlerFunc)
			handlers["/psrs"] = web.PSRs(ctx, logger, psr_tellor.New(logger, cfg.PsrTellor, aggregator))
			srv, err := web.New(ctx, logger, nil, tsDB, cfg.Web)
			if err != nil {
				return errors.Wrap(err, "creating component:"+web.ComponentName)
			}
			g.Add(func() error {
				err := srv.Start()
				level.Info(logger).Log("msg", "shutdown complete", "component", web.ComponentName)
				return err
			}, func(error) {
				srv.Stop()
			})
		}
	}

	if err := g.Run(); err != nil {
		level.Error(logger).Log("msg", "main exited with error", "err", err)
		return err
	}

	level.Info(logger).Log("msg", "main shutdown complete")
	return nil
}
