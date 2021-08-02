// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/oklog/run"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/tellor-io/telliot/pkg/aggregator"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/db"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/gas_price/gas_station"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/mining"
	psrTellor "github.com/tellor-io/telliot/pkg/psr/tellor"
	psrTellorMesosphere "github.com/tellor-io/telliot/pkg/psr/tellorMesosphere"
	"github.com/tellor-io/telliot/pkg/reward"
	"github.com/tellor-io/telliot/pkg/submitter/tellor"
	"github.com/tellor-io/telliot/pkg/submitter/tellorMesosphere"
	"github.com/tellor-io/telliot/pkg/tasker"
	"github.com/tellor-io/telliot/pkg/tracker/dispute/count"
	"github.com/tellor-io/telliot/pkg/tracker/index"
	"github.com/tellor-io/telliot/pkg/tracker/profit"
	"github.com/tellor-io/telliot/pkg/tracker/submitted_values"
	transactorTellor "github.com/tellor-io/telliot/pkg/transactor/tellor"
	transactorTellorMesosphere "github.com/tellor-io/telliot/pkg/transactor/tellorMesosphere"
	"github.com/tellor-io/telliot/pkg/web"
)

type mineCmd struct {
	Config configPath `type:"existingfile" help:"path to config file"`
}

func (self mineCmd) Run() error {
	logger := logging.NewLogger()

	cfg, err := config.ParseConfig(logger, string(self.Config))
	if err != nil {
		return errors.Wrap(err, "creating config")
	}

	// Defining a global context for starting and stopping of components.
	ctx := context.Background()

	client, err := ethereum.NewClient(ctx, logger)
	if err != nil {
		return errors.Wrap(err, "creating ethereum client")
	}

	accounts, err := ethereum.GetAccounts()
	if err != nil {
		return errors.Wrap(err, "getting accounts")
	}

	// We define our run groups here.
	var g run.Group
	// Run groups.
	{
		// Handle interupts.
		g.Add(run.SignalHandler(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM))

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
			// 5 days are enough as the aggregator needs data only 24 hours in the past.
			tsdbOptions.RetentionDuration = int64(5 * 24 * time.Hour)
			_tsDB, err := tsdb.Open(cfg.Db.Path, nil, nil, tsdbOptions)
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
			srv, err := web.New(logger, ctx, tsDB, cfg.Web)
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
		aggregator, err := aggregator.New(logger, ctx, cfg.Aggregator, tsDB)
		if err != nil {
			return errors.Wrap(err, "creating aggregator")
		}

		// Index tracker.
		// Run only when not using remote DB as it needs to write to the local db.
		if cfg.Db.RemoteHost == "" {
			_tsDB, ok := tsDB.(*tsdb.DB)
			if !ok {
				return errors.New("tsdb is not a writable DB instance")
			}

			// Index Tracker.
			indexTracker, err := index.New(logger, ctx, cfg.IndexTracker, _tsDB, client)
			if err != nil {
				return errors.Wrap(err, "creating component:"+index.ComponentName)
			}

			g.Add(func() error {
				err := indexTracker.Run()
				level.Info(logger).Log("msg", "shutdown complete", "component", index.ComponentName)
				return err
			}, func(error) {
				indexTracker.Stop()
			})

			_netID, err := client.NetworkID(ctx)
			if err != nil {
				return errors.Wrap(err, "getting network ID")
			}
			netID := _netID.Int64()

			// Run some trackers it only when not connected to a remote DB.
			// A remote DB already runs a dispute tracker so no need to run another one.
			// Also run and only for mainnet or rinkeby as the tellor oracle exists only on those networks.
			if netID == 1 || netID == 4 {
				contractTellor, err := contracts.NewITellor(client)
				if err != nil {
					return errors.Wrap(err, "create tellor contract instance")
				}

				submittedValuesTracker, err := submitted_values.New(
					logger,
					ctx,
					cfg.SubmittedValuesTracker,
					_tsDB,
					client,
					contractTellor,
					psrTellor.New(logger, cfg.PsrTellor, aggregator),
				)
				if err != nil {
					return errors.Wrap(err, "creating component:"+submitted_values.ComponentName)
				}
				g.Add(func() error {
					submittedValuesTracker.Start()
					level.Info(logger).Log("msg", "shutdown complete", "component", submitted_values.ComponentName)
					return nil
				}, func(error) {
					submittedValuesTracker.Stop()
				})

				disputeCount, err := count.New(
					logger,
					ctx,
					cfg.DisputeCountTracker,
					client,
					contractTellor,
				)
				if err != nil {
					return errors.Wrap(err, "creating component:"+count.ComponentName)
				}
				g.Add(func() error {
					disputeCount.Start()
					level.Info(logger).Log("msg", "shutdown complete", "component", count.ComponentName)
					return nil
				}, func(error) {
					disputeCount.Stop()
				})
			}

		}

		gasPriceQuerier, err := gas_station.New(logger, cfg.GasStation, client)
		if err != nil {
			return errors.Wrap(err, "creating gas price tracker")
		}

		if cfg.SubmitterTellor.Enabled {
			contractTellor, err := contracts.NewITellor(client)
			if err != nil {
				return errors.Wrap(err, "create tellor contract instance")
			}

			// Profit tracker.
			profitTracker, err := profit.NewProfitTracker(logger, ctx, cfg.ProfitTracker, client, nil, contractTellor, accounts)
			if err != nil {
				return errors.Wrap(err, "creating component:"+profit.ComponentName)
			}
			g.Add(func() error {
				err := profitTracker.Start()
				level.Info(logger).Log("msg", "shutdown complete", "component", profit.ComponentName)
				return err
			}, func(error) {
				profitTracker.Stop()
			})

			// Event tasker.
			taskerTracker, taskerChs, err := tasker.New(ctx, logger, cfg.Tasker, client, contractTellor, accounts)
			if err != nil {
				return errors.Wrap(err, "creating component:"+tasker.ComponentName)
			}
			g.Add(func() error {
				err := taskerTracker.Start()
				level.Info(logger).Log("msg", "shutdown complete", "component", tasker.ComponentName)
				return err
			}, func(error) {
				taskerTracker.Stop()
			})

			reward, err := reward.New(logger, ctx, cfg.Reward, contractTellor, aggregator)
			if err != nil {
				return errors.Wrap(err, "creating reward tracker")
			}

			// Create a submitter for each account.
			for _, account := range accounts {
				loggerWithAddr := log.With(logger, "addr", account.Address.String()[:6])

				transactor, err := transactorTellor.New(loggerWithAddr, cfg.TransactorTellor, gasPriceQuerier, client, account, reward, contractTellor)
				if err != nil {
					return errors.Wrap(err, "creating transactor")
				}

				psr := psrTellor.New(loggerWithAddr, cfg.PsrTellor, aggregator)

				// Get a channel on which it listens for new data to submit.
				submitter, submitterCh, err := tellor.New(
					ctx,
					loggerWithAddr,
					cfg.SubmitterTellor,
					client,
					contractTellor,
					account,
					transactor,
					psr,
				)
				if err != nil {
					return errors.Wrap(err, "creating component:"+tellor.ComponentName)
				}
				g.Add(func() error {
					err := submitter.Start()
					level.Info(loggerWithAddr).Log("msg", "shutdown complete", "component", tellor.ComponentName)
					return err
				}, func(error) {
					submitter.Stop()
				})

				// The Miner component.
				miner, err := mining.NewMiningManager(loggerWithAddr, ctx, cfg.Mining, contractTellor, taskerChs[account.Address.String()], submitterCh, client)
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

		if cfg.SubmitterTellorMesosphere.Enabled {
			contract, err := contracts.NewITellorMesosphere(client)
			if err != nil {
				return errors.Wrap(err, "create contract instance")
			}

			// Create a submitter for each account.
			for _, account := range accounts {
				loggerWithAddr := log.With(logger, "addr", account.Address.String()[:6])
				psr := psrTellorMesosphere.New(loggerWithAddr, cfg.PsrTellorMesosphere, aggregator)
				transactor, err := transactorTellorMesosphere.New(loggerWithAddr, cfg.TransactorTellorMesosphere, gasPriceQuerier, client, account, contract)
				if err != nil {
					return errors.Wrap(err, "creating transactor")
				}

				submitter, err := tellorMesosphere.New(
					ctx,
					loggerWithAddr,
					cfg.SubmitterTellorMesosphere,
					client,
					contract,
					account,
					transactor,
					psr,
				)
				if err != nil {
					return errors.Wrap(err, "creating component:"+tellorMesosphere.ComponentName)
				}
				g.Add(func() error {
					err := submitter.Start()
					level.Info(loggerWithAddr).Log("msg", "shutdown complete", "component", tellorMesosphere.ComponentName)
					return err
				}, func(error) {
					submitter.Stop()
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
