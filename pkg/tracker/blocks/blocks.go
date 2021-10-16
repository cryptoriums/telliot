// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package blocks

import (
	"context"
	"fmt"
	"time"

	"github.com/cryptoriums/telliot/pkg/db"
	ethereum_t "github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/ethereum/go-ethereum"

	"github.com/cryptoriums/telliot/pkg/logging"
	"github.com/cryptoriums/telliot/pkg/math"
	"github.com/cryptoriums/telliot/pkg/tracker/head"
	"github.com/cryptoriums/telliot/pkg/tracker/index"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/tsdb"
)

const (
	ComponentName                = "trackerBlocks"
	MetricSymbolBlockNum         = "BLOCK_NUM"
	MetricSymbolBlockGasPriceAvg = "BLOCK_GAS_PRICE_AVG"
	MetricSymbolBlockBaseFee     = "BLOCK_BASE_FEE"
	MetricSymbolBlockTip         = "BLOCK_TIP_AVG"
)

type Config struct {
	LogLevel string
}

type TrackerBlocks struct {
	logger log.Logger
	ctx    context.Context
	stop   context.CancelFunc
	cfg    Config
	tsDB   *tsdb.DB

	headTracker *head.TrackerHead
	headBlocks  chan *types.Block
}

func New(
	ctx context.Context,
	logger log.Logger,
	cfg Config,
	tsDB *tsdb.DB,
	client ethereum.ChainReader,
) (*TrackerBlocks, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)

	headTracker, headBlocks, err := head.New(
		ctx,
		logger,
		client,
		ethereum_t.ReorgEventWait,
	)
	if err != nil {
		return nil, errors.Wrap(err, "creating head tracker")
	}

	ctx, stop := context.WithCancel(ctx)

	return &TrackerBlocks{
		cfg:         cfg,
		ctx:         ctx,
		stop:        stop,
		logger:      logger,
		tsDB:        tsDB,
		headTracker: headTracker,
		headBlocks:  headBlocks,
	}, nil
}

func (self *TrackerBlocks) Start() {
	level.Info(self.logger).Log("msg", "starting", "logLevel", self.cfg.LogLevel)

	go func() {
		if err := self.headTracker.Start(); err != nil {
			level.Error(self.logger).Log("msg", "starting the head tracker", "err", err)
			self.Stop()
		}
	}()

	self.monitorHead()
}

func (self *TrackerBlocks) monitorHead() {
	for {
		select {
		case <-self.ctx.Done():
			return
		case block := <-self.headBlocks:
			time.Sleep(time.Second)
			self.processBlock(block)
		}
	}
}

func (self *TrackerBlocks) processBlock(block *types.Block) {
	logger := log.With(self.logger, "block", block.Number().Int64())
	ctx, cncl := context.WithTimeout(self.ctx, time.Minute)
	defer cncl()

	var blockTxsPriceTotal float64
	var blokTxsTipTotal float64
	for _, tx := range block.Transactions() {
		fee := math.BigIntToFloat(block.BaseFee()) + math.BigIntToFloat(tx.GasTipCap())
		tip := math.BigIntToFloat(tx.GasTipCap())

		maxAllowed := math.BigIntToFloat(tx.GasFeeCap())
		maxAllowedTip := maxAllowed - math.BigIntToFloat(block.BaseFee())

		if fee > maxAllowed {
			fee = maxAllowed
		}
		if tip > maxAllowedTip {
			tip = maxAllowedTip
		}

		blockTxsPriceTotal += fee
		blokTxsTipTotal += tip

	}

	if len(block.Transactions()) > 0 { // Some testnets have blocks without TXs.
		baseFee := math.BigIntToFloatDiv(block.BaseFee(), params.GWei)
		blockGasAvg := blockTxsPriceTotal / float64(len(block.Transactions())) / params.GWei // Record the value in GWEI.
		blockTipAvg := blokTxsTipTotal / float64(len(block.Transactions())) / params.GWei    // Record the value in GWEI.

		self.record(ctx, logger, MetricSymbolBlockNum, float64(block.Number().Int64()))
		self.record(ctx, logger, MetricSymbolBlockGasPriceAvg, blockGasAvg)
		self.record(ctx, logger, MetricSymbolBlockBaseFee, baseFee)
		self.record(ctx, logger, MetricSymbolBlockTip, blockTipAvg)

		level.Debug(logger).Log("msg", "added block details",
			"blockTs", block.Time(),
			"txCount", len(block.Transactions()),
			"blockGasAvg", blockGasAvg,
			"baseFee", fmt.Sprintf("%.2f", baseFee),
		)
	}
}

func (self *TrackerBlocks) record(ctx context.Context, logger log.Logger, symbol string, v float64) {
	lbls := labels.Labels{
		labels.Label{Name: "__name__", Value: index.MetricIndexValue},
		labels.Label{Name: "symbol", Value: symbol},
	}
	err := db.Add(ctx, self.tsDB, lbls, v)
	if err != nil {
		level.Error(logger).Log("msg", "add values to the DB", "err", err)
		return
	}
}

func (self *TrackerBlocks) Stop() {
	self.stop()
}
