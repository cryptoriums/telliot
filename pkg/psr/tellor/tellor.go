// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package tellor

import (
	"encoding/json"
	"math"
	"math/big"
	"time"

	"github.com/cryptoriums/telliot/pkg/aggregator"
	ethereum_t "github.com/cryptoriums/telliot/pkg/ethereum"
	math_t "github.com/cryptoriums/telliot/pkg/math"
	"github.com/cryptoriums/telliot/pkg/tracker/blocks"
	"github.com/cryptoriums/telliot/pkg/tracker/index"
	"github.com/go-kit/log"
	"github.com/pkg/errors"
)

type PsrID struct {
	Pair            string
	Aggr            string
	ConfidenceQuery string
	QueryID         [32]byte
	Inactive        bool
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

type Query struct {
	Type string
	ID   int64
}

func (self *Query) Bytes() ([]byte, error) {
	return json.Marshal(self)
}

func PsrByID(id int64) (PsrID, error) {
	return PsrByQuery(*NewQuery(id))
}

func PsrByQueryBytes(query []byte) (PsrID, error) {
	qj := &Query{}
	err := json.Unmarshal(query, qj)
	if err != nil {
		return PsrID{}, err
	}
	return PsrByQuery(*qj)
}

func PsrByQuery(query Query) (PsrID, error) {
	psr, ok := Psrs[query]
	if !ok {
		return PsrID{}, errors.Errorf("invalid query:%+v", query)
	}
	return addQueryHash(psr, query)
}

func addQueryHash(psr PsrID, q Query) (PsrID, error) {
	if psr.QueryID == [32]byte{} {
		qj, err := json.Marshal(q)
		if err != nil {
			return PsrID{}, err
		}
		psr.QueryID = ethereum_t.Keccak256(string(qj))
	}
	return psr, nil
}

func NewQuery(id int64) *Query {
	return &Query{Type: "LegacyRequest", ID: id}
}

func BytesToQuery(val []byte) *Query {
	return NewQuery(big.NewInt(0).SetBytes(val).Int64())
}

var Psrs = map[Query]PsrID{
	*NewQuery(1): {Pair: "ETH/USD", Aggr: Median, QueryID: math_t.IntTo32Byte(1)},
	*NewQuery(2): {Pair: "BTC/USD", Aggr: Median, QueryID: math_t.IntTo32Byte(2)},
	*NewQuery(3): {Inactive: true, Pair: "BNB/USD", Aggr: Median},
	*NewQuery(4): {Inactive: true, Pair: "BTC/USD", Aggr: TimeWeightedAvg24h},
	*NewQuery(5): {Inactive: true, Pair: "ETH/BTC", Aggr: Median},
	*NewQuery(6): {Inactive: true, Pair: "BNB/BTC", Aggr: Median},
	*NewQuery(7): {Inactive: true, Pair: "BNB/ETH", Aggr: Median},
	*NewQuery(8): {Inactive: true, Pair: "ETH/USD", Aggr: TimeWeightedAvg24h},
	*NewQuery(9): {Inactive: true, Pair: "ETH/USD", Aggr: MedianEOD},
	// // For more details see https://docs.google.com/document/d/1RFCApk1PznMhSRVhiyFl_vBDPA4mP2n1dTmfqjvuTNw/edit
	// // For now this uses third party APIs and don't do local aggregation.
	*NewQuery(10): {Pair: "AMPL/USD/VWAP", Aggr: Median, QueryID: math_t.IntTo32Byte(10)},
	*NewQuery(11): {Inactive: true, Pair: "ZEC/ETH", Aggr: Median},
	*NewQuery(12): {Inactive: true, Pair: "TRX/ETH", Aggr: Median},
	*NewQuery(13): {Inactive: true, Pair: "XRP/USD", Aggr: Median},
	*NewQuery(14): {Inactive: true, Pair: "XMR/ETH", Aggr: Median},
	*NewQuery(15): {Inactive: true, Pair: "ATOM/USD", Aggr: Median},
	*NewQuery(16): {Inactive: true, Pair: "LTC/USD", Aggr: Median},
	*NewQuery(17): {Inactive: true, Pair: "WAVES/BTC", Aggr: Median},
	*NewQuery(18): {Inactive: true, Pair: "REP/BTC", Aggr: Median},
	*NewQuery(19): {Inactive: true, Pair: "TUSD/ETH", Aggr: Median},
	*NewQuery(20): {Inactive: true, Pair: "EOS/USD", Aggr: Median},
	*NewQuery(21): {Inactive: true, Pair: "IOTA/USD", Aggr: Median},
	*NewQuery(22): {Inactive: true, Pair: "ETC/USD", Aggr: Median},
	*NewQuery(23): {Inactive: true, Pair: "ETH/PAX", Aggr: Median},
	*NewQuery(24): {Inactive: true, Pair: "ETH/BTC", Aggr: TimeWeightedAvg1h},
	*NewQuery(25): {Inactive: true, Pair: "USDC/USDT", Aggr: Median},
	*NewQuery(26): {Inactive: true, Pair: "XTZ/USD", Aggr: Median},
	*NewQuery(27): {Inactive: true, Pair: "LINK/USD", Aggr: Median},
	*NewQuery(28): {Inactive: true, Pair: "ZRX/BNB", Aggr: Median},
	*NewQuery(29): {Inactive: true, Pair: "ZEC/USD", Aggr: Median},
	*NewQuery(30): {Inactive: true, Pair: "XAU/USD", Aggr: Median},
	*NewQuery(31): {Pair: "MATIC/USD", Aggr: Median, QueryID: math_t.IntTo32Byte(31)},
	*NewQuery(32): {Inactive: true, Pair: "BAT/USD", Aggr: Median},
	*NewQuery(33): {Inactive: true, Pair: "ALGO/USD", Aggr: Median},
	*NewQuery(34): {Inactive: true, Pair: "ZRX/USD", Aggr: Median},
	*NewQuery(35): {Inactive: true, Pair: "COS/USD", Aggr: Median},
	*NewQuery(36): {Inactive: true, Pair: "BCH/USD", Aggr: Median},
	*NewQuery(37): {Inactive: true, Pair: "REP/USD", Aggr: Median},
	*NewQuery(38): {Inactive: true, Pair: "GNO/USD", Aggr: Median},
	*NewQuery(39): {Inactive: true, Pair: "DAI/USD", Aggr: Median},
	*NewQuery(40): {Inactive: true, Pair: "STEEM/BTC", Aggr: Median},
	// It is three month average for US PCE (monthly levels): https://www.bea.gov/data/personal-consumption-expenditures-price-index-excluding-food-and-energy
	*NewQuery(41): {Pair: "USPCE", Aggr: Median, QueryID: math_t.IntTo32Byte(41)},
	*NewQuery(42): {Inactive: true, Pair: "BTC/USD", Aggr: MedianEOD},
	*NewQuery(43): {Inactive: true, Pair: "TRB/ETH", Aggr: Median},
	*NewQuery(44): {Inactive: true, Pair: "BTC/USD", Aggr: TimeWeightedAvg1h},
	*NewQuery(45): {Inactive: true, Pair: "TRB/USD", Aggr: MedianEOD},
	*NewQuery(46): {Inactive: true, Pair: "ETH/USD", Aggr: TimeWeightedAvg1h},
	*NewQuery(47): {Inactive: true, Pair: "BSV/USD", Aggr: Median},
	*NewQuery(48): {Inactive: true, Pair: "MAKER/USD", Aggr: Median},
	*NewQuery(49): {Inactive: true, Pair: "BCH/USD", Aggr: TimeWeightedAvg24h},
	*NewQuery(50): {Pair: "TRB/USD", Aggr: Median, QueryID: math_t.IntTo32Byte(50)},
	*NewQuery(51): {Inactive: true, Pair: "XMR/USD", Aggr: Median},
	*NewQuery(52): {Inactive: true, Pair: "XFT/USD", Aggr: Median},
	*NewQuery(53): {Inactive: true, Pair: "BTCDOMINANCE", Aggr: Median},
	*NewQuery(54): {Inactive: true, Pair: "WAVES/USD", Aggr: Median},
	*NewQuery(55): {Inactive: true, Pair: "OGN/USD", Aggr: Median},
	*NewQuery(56): {Inactive: true, Pair: "VIXEOD", Aggr: Median},
	*NewQuery(57): {Inactive: true, Pair: "DEFITVL", Aggr: Median},
	*NewQuery(58): {Inactive: true, Pair: "DEFIMCAP", Aggr: Mean},
	*NewQuery(59): {Pair: "ETH/JPY", Aggr: Median, QueryID: math_t.IntTo32Byte(59)},
	*NewQuery(60): {Inactive: true, Pair: blocks.MetricSymbolBlockGasPriceAvg, Aggr: TimeWeightedAvg7Days, QueryID: math_t.IntTo32Byte(60), ConfidenceQuery: `sum(count_over_time(` + index.MetricIndexValue + `{symbol="` + blocks.MetricSymbolBlockGasPriceAvg + `"}[` + Week.String() + `]))/ sum(` + index.MetricIndexValue + `{symbol="` + blocks.MetricSymbolBlockNum + `"} - ` + index.MetricIndexValue + `{symbol="` + blocks.MetricSymbolBlockNum + `"} offset ` + Week.String() + `)`},
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
	MinConfidencePerSymbol map[int64]float64
}

type Psr struct {
	logger     log.Logger
	aggregator *aggregator.Aggregator
	cfg        Config
}

func (self *Psr) ConfidenceThreshold(id int64) float64 {
	return self.cfg.MinConfidencePerSymbol[id]
}

func (self *Psr) GetValue(q Query, ts time.Time) (float64, error) {
	psr, ok := Psrs[q]
	if !ok {
		return 0, errors.Errorf("invalid query:%v", q)
	}

	if psr.Inactive {
		return 0, nil
	}
	val, err := self.getValue(q, ts)
	return math.Round(val * DefaultGranularity), err
}

func (self *Psr) getValue(q Query, ts time.Time) (float64, error) {
	var val float64
	var err error

	psr, ok := Psrs[q]
	if !ok {
		return 0, errors.Errorf("invalid query:%v", q)
	}

	var conf float64
	switch psr.Aggr {
	case Median:
		val, conf, err = self.aggregator.MedianAt(psr.Pair, ts)
	case TimeWeightedAvg24h:
		val, conf, err = self.aggregator.TimeWeightedAvg(psr.Pair, psr.ConfidenceQuery, ts, 24*time.Hour)
	case TimeWeightedAvg1h:
		val, conf, err = self.aggregator.TimeWeightedAvg(psr.Pair, psr.ConfidenceQuery, ts, time.Hour)
	case TimeWeightedAvg7Days:
		val, conf, err = self.aggregator.TimeWeightedAvg(psr.Pair, psr.ConfidenceQuery, ts, Week)
	case MedianEOD:
		val, conf, err = self.aggregator.MedianAtEOD(psr.Pair, ts)
	case Mean:
		val, conf, err = self.aggregator.MeanAt(psr.Pair, ts)
	}

	if err != nil {
		return 0, err
	}

	if confS, ok := self.cfg.MinConfidencePerSymbol[q.ID]; ok {
		if conf < confS {
			return 0, errors.Errorf("not enough confidence based on the aggregator calculations - query:%v, value:%v, conf:%v,symbol confidence threshold:%v", q, val, conf, self.cfg.MinConfidencePerSymbol[q.ID])
		}
	} else if conf < self.cfg.MinConfidenceDefault {
		return 0, errors.Errorf("not enough confidence based on the aggregator calculations - query:%v, value:%v, conf:%v,default confidence threshold:%v", q, val, conf, self.cfg.MinConfidenceDefault)
	}

	return val, err
}
