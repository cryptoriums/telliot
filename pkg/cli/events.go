// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"fmt"
	"syscall"
	"time"

	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/contracts/tellorX_master"
	"github.com/cryptoriums/telliot/pkg/contracts/tellorX_oracle"
	"github.com/cryptoriums/telliot/pkg/tracker/events"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/oklog/run"
	"github.com/pkg/errors"
)

type EventsCmd struct {
	LookBack     time.Duration `help:"how far to look for the initiali qyery"`
	Name         string        `required:"" help:"the name of the log to watch"`
	ContractName string        `required:"" enum:"master,oracle,govern" help:"the name of the contract to watch"`
	ReorgWait    time.Duration `default:"3s" help:"how long to wait for removed logs from reorg events"`
}

func (self *EventsCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	_, client, master, oracle, govern, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}

	ctx, cncl := context.WithCancel(ctx)
	defer cncl()

	var contract contracts.ContractCaller
	switch self.ContractName {
	case "master":
		contract = master
	case "oracle":
		contract = oracle
	case "govern":
		contract = govern

	}

	trackerEvents, output, err := events.New(
		ctx,
		logger,
		events.Config{LogLevel: "debug"},
		client,
		contract,
		self.LookBack,
		self.Name,
		self.ReorgWait,
	)
	if err != nil {
		return errors.Wrap(err, "creating events tracker")
	}

	// We define our run groups here.
	var g run.Group
	// Run groups.
	{
		// Handle interupts.
		g.Add(run.SignalHandler(ctx, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM))

		g.Add(func() error {
			err := trackerEvents.Start()
			level.Info(logger).Log("msg", "shutdown complete", "component", events.ComponentName)
			return err
		}, func(error) {
			trackerEvents.Stop()
		})

		g.Add(func() error {
			for {
				select {
				case log := <-output:
					err := self.unpack(contract, log)
					if err != nil {
						return err
					}

				case <-ctx.Done():
					return nil
				}
			}

		}, func(error) {
			cncl()
		})

	}

	if err := g.Run(); err != nil {
		level.Error(logger).Log("msg", "main exited with error", "err", err)
		return err
	}

	level.Info(logger).Log("msg", "main shutdown complete")

	return nil
}

func (self *EventsCmd) unpack(contract contracts.ContractCaller, logRaw types.Log) error {
	switch eventName := self.Name; eventName {
	case contracts.EventNameNewSubmit:
		log := &tellorX_oracle.OracleNewReport{}
		err := contract.UnpackLog(log, eventName, logRaw)
		if err != nil {
			return errors.Wrap(err, "unpack event from logs")
		}
		log.Raw = logRaw
		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Printf("Hash:%v, %+v \n", logRaw.TxHash, log)
	case contracts.EventNameNewTip:
		log := &tellorX_oracle.OracleTipAdded{}
		err := contract.UnpackLog(log, eventName, logRaw)
		if err != nil {
			return errors.Wrap(err, "unpack event from logs")
		}
		log.Raw = logRaw
		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Printf("Hash:%v, %+v \n", logRaw.TxHash, log)
	case contracts.EventNameTransfer:
		log := &tellorX_master.ControllerTransfer{}
		err := contract.UnpackLog(log, eventName, logRaw)
		if err != nil {
			return errors.Wrap(err, "unpack event from logs")
		}
		log.Raw = logRaw
		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Printf("Hash:%v, %+v \n", logRaw.TxHash, log)
	default:
		return errors.Errorf("unsuporter log name:%v", eventName)
	}

	return nil
}
