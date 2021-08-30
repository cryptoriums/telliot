// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"syscall"
	"time"

	"github.com/cryptoriums/telliot/pkg/aggregator"
	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/db"
	"github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/gas_price/gas_station"
	"github.com/cryptoriums/telliot/pkg/mining"
	psrTellor "github.com/cryptoriums/telliot/pkg/psr/tellor"
	tellorSubmit "github.com/cryptoriums/telliot/pkg/submitter/tellor"
	"github.com/cryptoriums/telliot/pkg/tasker"
	"github.com/cryptoriums/telliot/pkg/tracker/index"
	transactorTellor "github.com/cryptoriums/telliot/pkg/transactor/tellor"
	"github.com/cryptoriums/telliot/pkg/web"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/oklog/run"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
)

type ReportCmd struct {
}

func (self *ReportCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, client, contract, err := ConfigClientContract(ctx, logger, cli.Config, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	accounts, err := ethereum.GetAccounts(logger)
	if err != nil {
		return errors.Wrap(err, "getting accounts")
	}

	// We define our run groups here.
	var g run.Group
	// Run groups.
	{
		// Handle interupts.
		g.Add(run.SignalHandler(ctx, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM))

		// Open a local or remote instance of the TSDB database.
		var tsDB storage.SampleAndChunkQueryable
		if cfg.Db.RemoteHost != "" {
			tsDB, err = db.NewRemoteDB(cfg.Db)
			if err != nil {
				return errors.Wrap(err, "opening remote tsdb DB")
			}
			level.Info(logger).Log("msg", "connected to remote db", "host", cfg.Db.RemoteHost, "port", cfg.Db.RemotePort)
		} else {
			// Open the TSDB database.
			tsdbOptions := tsdb.DefaultOptions()
			// 30 days are enough as the maximum data we need it for disputes which don't run for more than 2-3 days.
			tsdbOptions.RetentionDuration = int64(30 * 24 * time.Hour)
			_tsDB, err := tsdb.Open(cfg.Db.Path, nil, nil, tsdbOptions, tsdb.NewDBStats())
			if err != nil {
				return errors.Wrap(err, "opening local tsdb DB")
			}
			defer func() {
				if err := _tsDB.Close(); err != nil {
					level.Error(logger).Log("msg", "closing the tsdb", "err", err)
				}
			}()
			tsDB = _tsDB
			level.Info(logger).Log("msg", "opened local db", "path", cfg.Db.Path)
			level.Warn(logger).Log("msg", "FOR NEW DB INSTANCES IT IS NORMAL TO SEE SOME QUERY ERRORS AS THE DATABASE IS NOT YET POPULATED WITH VALUES")
		}

		// Web/Api server.
		{
			srv, err := web.New(ctx, logger, tsDB, cfg.Web)
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

		// Aggregator.
		aggregator, err := aggregator.New(ctx, logger, cfg.Aggregator, tsDB)
		if err != nil {
			return errors.Wrap(err, "creating aggregator")
		}

		// Run these only when not using remote DB as it needs to write to the local db.
		if cfg.Db.RemoteHost == "" {
			_tsDB, ok := tsDB.(*tsdb.DB)
			if !ok {
				return errors.New("tsdb is not a writable DB instance")
			}

			// Index Tracker.
			trackerIndex, err := index.New(ctx, logger, cfg.TrackerIndex, _tsDB)
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
		}

		gasPriceQuerier, err := gas_station.New(ctx, logger, client, contract.NetID())
		if err != nil {
			return errors.Wrap(err, "creating gas price tracker")
		}

		if cfg.SubmitterTellor.Enabled {
			// Event tasker.
			taskerE, taskerChs, err := tasker.New(ctx, logger, cfg.Tasker, client, contract, accounts)
			if err != nil {
				return errors.Wrap(err, "creating component:"+tasker.ComponentName)
			}
			g.Add(func() error {
				err := taskerE.Start()
				level.Info(logger).Log("msg", "shutdown complete", "component", tasker.ComponentName)
				return err
			}, func(error) {
				taskerE.Stop()
			})

			psr := psrTellor.New(logger, cfg.PsrTellor, aggregator)

			// Create a submitter for each account.
			for _, account := range accounts {
				loggerWithAddr := log.With(logger, "addr", account.Address.Hex()[:8])

				transactor, err := transactorTellor.New(
					loggerWithAddr,
					cfg.TransactorTellor,
					gasPriceQuerier,
					client,
					account,
					contract,
				)
				if err != nil {
					return errors.Wrap(err, "creating transactor")
				}

				// Get a channel on which it listens for new data to submit.
				submitter, submitterCh, err := tellorSubmit.New(
					ctx,
					loggerWithAddr,
					cfg.SubmitterTellor,
					client,
					contract,
					account,
					transactor,
					psr,
				)
				if err != nil {
					return errors.Wrap(err, "creating component:"+tellorSubmit.ComponentName)
				}
				g.Add(func() error {
					err := submitter.Start()
					level.Info(loggerWithAddr).Log("msg", "shutdown complete", "component", tellorSubmit.ComponentName)
					return err
				}, func(error) {
					submitter.Stop()
				})

				// The Miner component.
				miner, err := mining.NewManager(ctx, loggerWithAddr, cfg.Mining, contract, taskerChs[account.Address.Hex()], submitterCh, client)
				if err != nil {
					return errors.Wrap(err, "creating component:"+mining.ComponentName)
				}
				g.Add(func() error {
					err := miner.Start()
					level.Info(loggerWithAddr).Log("msg", "shutdown complete", "component", mining.ComponentName)
					return err
				}, func(error) {
					miner.Stop()
				})
			}
		}
	}

	if err := g.Run(); err != nil {
		level.Error(logger).Log("msg", "main exited with error", "err", err)
		return err
	}

	level.Info(logger).Log("msg", "main shutdown complete")
	return nil
}
