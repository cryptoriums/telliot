// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package tellor

import (
	"math"
	"strconv"
	"time"

	"github.com/cryptoriums/telliot/pkg/aggregator"
	"github.com/cryptoriums/telliot/pkg/tracker/blocks"
	"github.com/cryptoriums/telliot/pkg/tracker/index"
	"github.com/go-kit/log"
	"github.com/pkg/errors"
)

type PsrID struct {
	Pair            string
	Aggr            string
	ConfidenceQuery string
}

const (
	ComponentName      = "psrTellor"
	DefaultGranularity = 1000000

	Median               = "Median"
	MedianEOD            = "Median EOD"
	Mean                 = "Mean"
	TimeWeightedAvg1h    = "TWAP 1h"
	TimeWeightedAvg24h   = "TWAP 24h"
	Week                 = 168 * time.Hour
	TimeWeightedAvg7Days = "TWAP 1w"
)

var Psrs = map[int64]PsrID{
	1: {Pair: "ETH/USD", Aggr: Median},
	2: {Pair: "BTC/USD", Aggr: Median},
	3: {Pair: "BNB/USD", Aggr: Median},
	4: {Pair: "BTC/USD", Aggr: TimeWeightedAvg24h},
	5: {Pair: "ETH/BTC", Aggr: Median},
	6: {Pair: "BNB/BTC", Aggr: Median},
	7: {Pair: "BNB/ETH", Aggr: Median},
	8: {Pair: "ETH/USD", Aggr: TimeWeightedAvg24h},
	9: {Pair: "ETH/USD", Aggr: MedianEOD},
	// For more details see https://docs.google.com/document/d/1RFCApk1PznMhSRVhiyFl_vBDPA4mP2n1dTmfqjvuTNw/edit
	// For now this uses third party APIs and don't do local aggregation.
	10: {Pair: "AMPL/USD/VWAP", Aggr: Median},
	11: {Pair: "ZEC/ETH", Aggr: Median},
	12: {Pair: "TRX/ETH", Aggr: Median},
	13: {Pair: "XRP/USD", Aggr: Median},
	14: {Pair: "XMR/ETH", Aggr: Median},
	15: {Pair: "ATOM/USD", Aggr: Median},
	16: {Pair: "LTC/USD", Aggr: Median},
	17: {Pair: "WAVES/BTC", Aggr: Median},
	18: {Pair: "REP/BTC", Aggr: Median},
	19: {Pair: "TUSD/ETH", Aggr: Median},
	20: {Pair: "EOS/USD", Aggr: Median},
	21: {Pair: "IOTA/USD", Aggr: Median},
	22: {Pair: "ETC/USD", Aggr: Median},
	23: {Pair: "ETH/PAX", Aggr: Median},
	24: {Pair: "ETH/BTC", Aggr: TimeWeightedAvg1h},
	25: {Pair: "USDC/USDT", Aggr: Median},
	26: {Pair: "XTZ/USD", Aggr: Median},
	27: {Pair: "LINK/USD", Aggr: Median},
	28: {Pair: "ZRX/BNB", Aggr: Median},
	29: {Pair: "ZEC/USD", Aggr: Median},
	30: {Pair: "XAU/USD", Aggr: Median},
	31: {Pair: "MATIC/USD", Aggr: Median},
	32: {Pair: "BAT/USD", Aggr: Median},
	33: {Pair: "ALGO/USD", Aggr: Median},
	34: {Pair: "ZRX/USD", Aggr: Median},
	35: {Pair: "COS/USD", Aggr: Median},
	36: {Pair: "BCH/USD", Aggr: Median},
	37: {Pair: "REP/USD", Aggr: Median},
	38: {Pair: "GNO/USD", Aggr: Median},
	39: {Pair: "DAI/USD", Aggr: Median},
	40: {Pair: "STEEM/BTC", Aggr: Median},
	// It is three month average for US PCE (monthly levels): https://www.bea.gov/data/personal-consumption-expenditures-price-index-excluding-food-and-energy
	41: {Pair: "USPCE", Aggr: Median},
	42: {Pair: "BTC/USD", Aggr: MedianEOD},
	43: {Pair: "TRB/ETH", Aggr: Median},
	44: {Pair: "BTC/USD", Aggr: TimeWeightedAvg1h},
	45: {Pair: "TRB/USD", Aggr: MedianEOD},
	46: {Pair: "ETH/USD", Aggr: TimeWeightedAvg1h},
	47: {Pair: "BSV/USD", Aggr: Median},
	48: {Pair: "MAKER/USD", Aggr: Median},
	49: {Pair: "BCH/USD", Aggr: TimeWeightedAvg24h},
	50: {Pair: "TRB/USD", Aggr: Median},
	51: {Pair: "XMR/USD", Aggr: Median},
	52: {Pair: "XFT/USD", Aggr: Median},
	53: {Pair: "BTCDOMINANCE", Aggr: Median},
	54: {Pair: "WAVES/USD", Aggr: Median},
	55: {Pair: "OGN/USD", Aggr: Median},
	56: {Pair: "VIXEOD", Aggr: Median},
	57: {Pair: "DEFITVL", Aggr: Median},
	58: {Pair: "DEFIMCAP", Aggr: Mean},
	59: {Pair: "ETH/JPY", Aggr: Median},
	60: {Pair: blocks.MetricSymbolBlockGasPriceAvg, Aggr: TimeWeightedAvg7Days, ConfidenceQuery: `sum(count_over_time(` + index.MetricIndexValue + `{symbol="` + blocks.MetricSymbolBlockGasPriceAvg + `"}[` + Week.String() + `]))/ sum(` + index.MetricIndexValue + `{symbol="` + blocks.MetricSymbolBlockNum + `"} - ` + index.MetricIndexValue + `{symbol="` + blocks.MetricSymbolBlockNum + `"} offset ` + Week.String() + `)`},
}

func New(logger log.Logger, cfg Config, aggregator *aggregator.Aggregator) *Psr {
	return &Psr{
		logger:     log.With(logger, "component", ComponentName),
		aggregator: aggregator,
		cfg:        cfg,
	}
}

type Config struct {
	MinConfidenceDefault   float64
	MinConfidencePerSymbol map[string]float64
}

type Psr struct {
	logger     log.Logger
	aggregator *aggregator.Aggregator
	cfg        Config
}

func (self *Psr) ConfidenceThreshold(id int64) float64 {
	return self.cfg.MinConfidencePerSymbol[strconv.Itoa(int(id))]
}

func (self *Psr) GetValue(reqID int64, ts time.Time) (float64, error) {
	val, err := self.getValue(reqID, ts)
	return math.Round(val * DefaultGranularity), err
}

func (self *Psr) getValue(reqID int64, ts time.Time) (float64, error) {
	var val float64
	var err error

	if _, ok := Psrs[reqID]; !ok {
		return 0, errors.Errorf("invalid reqID id:%v", reqID)
	}

	var conf float64
	switch Psrs[reqID].Aggr {
	case Median:
		val, conf, err = self.aggregator.MedianAt(Psrs[reqID].Pair, ts)
	case TimeWeightedAvg24h:
		val, conf, err = self.aggregator.TimeWeightedAvg(Psrs[reqID].Pair, Psrs[reqID].ConfidenceQuery, ts, 24*time.Hour)
	case TimeWeightedAvg1h:
		val, conf, err = self.aggregator.TimeWeightedAvg(Psrs[reqID].Pair, Psrs[reqID].ConfidenceQuery, ts, time.Hour)
	case TimeWeightedAvg7Days:
		val, conf, err = self.aggregator.TimeWeightedAvg(Psrs[reqID].Pair, Psrs[reqID].ConfidenceQuery, ts, Week)
	case MedianEOD:
		val, conf, err = self.aggregator.MedianAtEOD(Psrs[reqID].Pair, ts)
	case Mean:
		val, conf, err = self.aggregator.MeanAt(Psrs[reqID].Pair, ts)
	}

	if err != nil {
		return 0, err
	}

	if confS, ok := self.cfg.MinConfidencePerSymbol[strconv.Itoa(int(reqID))]; ok {
		if conf < confS {
			return 0, errors.Errorf("not enough confidence based on the aggregator calculations - reqID:%v, value:%v, conf:%v,symbol confidence threshold:%v", reqID, val, conf, self.cfg.MinConfidencePerSymbol[strconv.Itoa(int(reqID))])
		}
	} else if conf < self.cfg.MinConfidenceDefault {
		return 0, errors.Errorf("not enough confidence based on the aggregator calculations - reqID:%v, value:%v, conf:%v,default confidence threshold:%v", reqID, val, conf, self.cfg.MinConfidenceDefault)
	}

	return val, err
}

func IsInactive(id int64) bool {
	switch id {
	case 3:
		return true
	case 4:
		return true
	case 5:
		return true
	case 6:
		return true
	case 7:
		return true
	case 8:
		return true
	case 9:
		return true
	case 11:
		return true
	case 12:
		return true
	case 13:
		return true
	case 14:
		return true
	case 15:
		return true
	case 16:
		return true
	case 17:
		return true
	case 18:
		return true
	case 19:
		return true
	case 20:
		return true
	case 21:
		return true
	case 22:
		return true
	case 23:
		return true
	case 24:
		return true
	case 25:
		return true
	case 26:
		return true
	case 27:
		return true
	case 28:
		return true
	case 29:
		return true
	case 30:
		return true
	case 32:
		return true
	case 33:
		return true
	case 34:
		return true
	case 35:
		return true
	case 36:
		return true
	case 37:
		return true
	case 38:
		return true
	case 39:
		return true
	case 40:
		return true
	case 42:
		return true
	case 43:
		return true
	case 44:
		return true
	case 45:
		return true
	case 46:
		return true
	case 47:
		return true
	case 48:
		return true
	case 49:
		return true
	case 51:
		return true
	case 52:
		return true
	case 53:
		return true
	case 54:
		return true
	case 55:
		return true
	case 56:
		return true
	case 57:
		return true
	case 58:
		return true
	default:
		return false
	}
}
