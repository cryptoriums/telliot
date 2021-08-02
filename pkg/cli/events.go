// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"fmt"
	"syscall"
	"time"

	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/logging"
	"github.com/cryptoriums/telliot/pkg/tracker/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/kit/log/level"
	"github.com/oklog/run"
	"github.com/pkg/errors"
)

type eventsCmd struct {
	LookBack time.Duration `help:"how far to look for the initiali qyery"`
	ContractFlag
	EventName string        `required:"" enum:"NonceSubmitted,NewTellorAddress" help:"the name of the log to watch"`
	ReorgWait time.Duration `default:"3s" help:"how long to wait for removed logs from reorg events"`
}

func (self *eventsCmd) Run() error {
	logger := logging.NewLogger()
	ctx, cncl := context.WithCancel(context.Background())
	defer cncl()

	client, netID, err := ethereum.NewClient(logger, ctx)
	if err != nil {
		return errors.Wrap(err, "creating ethereum client")
	}

	if self.Contract == "" {
		contractAddr, err := contracts.GetTellorAddress(netID)
		if err != nil {
			return errors.Wrap(err, "getting tellor address")
		}
		self.Contract = contractAddr.Hex()
	}

	contract, err := contracts.NewITellorWithAddr(logger, ctx, common.HexToAddress(self.Contract), client, netID, contracts.DefaultParams)
	if err != nil {
		return errors.Wrap(err, "creating contract instance")
	}

	trackerEvents, output, err := events.New(
		ctx,
		logger,
		client,
		contract,
		self.LookBack,
		self.EventName,
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
		g.Add(run.SignalHandler(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM))

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

func (self *eventsCmd) unpack(contract contracts.ContractCaller, logRaw types.Log) error {
	switch eventName := self.EventName; eventName {
	case contracts.EventNameNewSubmit:
		log := &contracts.TellorNonceSubmitted{}
		err := contract.UnpackLog(log, eventName, logRaw)
		if err != nil {
			return errors.Wrap(err, "unpack event from logs")
		}
		log.Raw = logRaw
		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Printf("%+v \n", log)
	case contracts.EventNameContractUpgrade:
		log := &contracts.NewTellorAddress{}
		err := contract.UnpackLog(log, eventName, logRaw)
		if err != nil {
			return errors.Wrap(err, "unpack event from logs")
		}
		log.Raw = logRaw
		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Printf("%+v \n", log)
	default:
		return errors.Errorf("unsuporter log name:%v", eventName)
	}

	return nil
}
