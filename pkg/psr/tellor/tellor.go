// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package tellor

import (
	"encoding/binary"
	"encoding/json"
	"math"
	"strconv"
	"time"

	ethereum_t "github.com/cryptoriums/packages/ethereum"
	"github.com/cryptoriums/telliot/pkg/aggregator"
	"github.com/cryptoriums/telliot/pkg/tracker/blocks"
	"github.com/cryptoriums/telliot/pkg/tracker/index"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type PsrID struct {
	Pair            string
	Aggr            string
	ConfidenceQuery string
	Inactive        bool
	MinTipAmmount   float64
	Granularity     float64
}

const (
	ComponentName      = "psrTellor"
	defaultGranularity = 1e6

	Median               = "Median"
	MedianEOD            = "Median EOD"
	Mean                 = "Mean"
	TimeWeightedAvg1h    = "TWAP 1h"
	TimeWeightedAvg24h   = "TWAP 24h"
	Week                 = 168 * time.Hour
	TimeWeightedAvg7Days = "TWAP 1w"

	MetricErrCount = "errors_total"
)

type Query struct {
	Type string
	ID   int64 `json:"legacy_id,omitempty"`
}

func (self *Query) Bytes() []byte {
	b, err := json.Marshal(self)
	if err != nil {
		panic(err) // This should never happen.
	}
	return b
}

func (self *Query) String() string {
	return string(self.Bytes())
}

func PsrByID(id int64) (PsrID, error) {
	psr, ok := Psrs[IntToQueryID(id)]
	if !ok {
		return PsrID{}, errors.Errorf("invalid ID:%v", id)
	}
	return psr, nil
}

func NewQueryData(id int64) []byte {
	if id > 100 {
		panic("invalid query id and can't create query data from it")
	}
	return (&Query{Type: "LegacyRequest", ID: id}).Bytes()
}

var Psrs = map[[32]byte]PsrID{
	IntToQueryID(1): {Pair: "ETH/USD", Aggr: Median, Granularity: defaultGranularity},
	IntToQueryID(2): {Pair: "BTC/USD", Aggr: Median, Granularity: defaultGranularity},
	IntToQueryID(3): {Inactive: true, Pair: "BNB/USD", Aggr: Median},
	IntToQueryID(4): {Inactive: true, Pair: "BTC/USD", Aggr: TimeWeightedAvg24h},
	IntToQueryID(5): {Inactive: true, Pair: "ETH/BTC", Aggr: Median},
	IntToQueryID(6): {Inactive: true, Pair: "BNB/BTC", Aggr: Median},
	IntToQueryID(7): {Inactive: true, Pair: "BNB/ETH", Aggr: Median},
	IntToQueryID(8): {Inactive: true, Pair: "ETH/USD", Aggr: TimeWeightedAvg24h},
	IntToQueryID(9): {Inactive: true, Pair: "ETH/USD", Aggr: MedianEOD},
	// // For more details see https://docs.google.com/document/d/1RFCApk1PznMhSRVhiyFl_vBDPA4mP2n1dTmfqjvuTNw/edit
	// // For now this uses third party APIs and don't do local aggregation.
	IntToQueryID(10): {Pair: "AMPL/USD/VWAP", Aggr: Median, MinTipAmmount: 1, Granularity: 1e18},
	IntToQueryID(11): {Inactive: true, Pair: "ZEC/ETH", Aggr: Median},
	IntToQueryID(12): {Inactive: true, Pair: "TRX/ETH", Aggr: Median},
	IntToQueryID(13): {Inactive: true, Pair: "XRP/USD", Aggr: Median},
	IntToQueryID(14): {Inactive: true, Pair: "XMR/ETH", Aggr: Median},
	IntToQueryID(15): {Inactive: true, Pair: "ATOM/USD", Aggr: Median},
	IntToQueryID(16): {Inactive: true, Pair: "LTC/USD", Aggr: Median},
	IntToQueryID(17): {Inactive: true, Pair: "WAVES/BTC", Aggr: Median},
	IntToQueryID(18): {Inactive: true, Pair: "REP/BTC", Aggr: Median},
	IntToQueryID(19): {Inactive: true, Pair: "TUSD/ETH", Aggr: Median},
	IntToQueryID(20): {Inactive: true, Pair: "EOS/USD", Aggr: Median},
	IntToQueryID(21): {Inactive: true, Pair: "IOTA/USD", Aggr: Median},
	IntToQueryID(22): {Inactive: true, Pair: "ETC/USD", Aggr: Median},
	IntToQueryID(23): {Inactive: true, Pair: "ETH/PAX", Aggr: Median},
	IntToQueryID(24): {Inactive: true, Pair: "ETH/BTC", Aggr: TimeWeightedAvg1h},
	IntToQueryID(25): {Inactive: true, Pair: "USDC/USDT", Aggr: Median},
	IntToQueryID(26): {Inactive: true, Pair: "XTZ/USD", Aggr: Median},
	IntToQueryID(27): {Inactive: true, Pair: "LINK/USD", Aggr: Median},
	IntToQueryID(28): {Inactive: true, Pair: "ZRX/BNB", Aggr: Median},
	IntToQueryID(29): {Inactive: true, Pair: "ZEC/USD", Aggr: Median},
	IntToQueryID(30): {Inactive: true, Pair: "XAU/USD", Aggr: Median},
	IntToQueryID(31): {Pair: "MATIC/USD", Aggr: Median, Granularity: defaultGranularity},
	IntToQueryID(32): {Inactive: true, Pair: "BAT/USD", Aggr: Median},
	IntToQueryID(33): {Inactive: true, Pair: "ALGO/USD", Aggr: Median},
	IntToQueryID(34): {Inactive: true, Pair: "ZRX/USD", Aggr: Median},
	IntToQueryID(35): {Inactive: true, Pair: "COS/USD", Aggr: Median},
	IntToQueryID(36): {Inactive: true, Pair: "BCH/USD", Aggr: Median},
	IntToQueryID(37): {Inactive: true, Pair: "REP/USD", Aggr: Median},
	IntToQueryID(38): {Inactive: true, Pair: "GNO/USD", Aggr: Median},
	IntToQueryID(39): {Inactive: true, Pair: "DAI/USD", Aggr: Median},
	IntToQueryID(40): {Inactive: true, Pair: "STEEM/BTC", Aggr: Median},
	// It is three month average for US PCE (monthly levels): https://www.bea.gov/data/personal-consumption-expenditures-price-index-excluding-food-and-energy
	IntToQueryID(41): {Pair: "USPCE", Aggr: Median, MinTipAmmount: 1, Granularity: 1e18},
	IntToQueryID(42): {Inactive: true, Pair: "BTC/USD", Aggr: MedianEOD},
	IntToQueryID(43): {Inactive: true, Pair: "TRB/ETH", Aggr: Median},
	IntToQueryID(44): {Inactive: true, Pair: "BTC/USD", Aggr: TimeWeightedAvg1h},
	IntToQueryID(45): {Inactive: true, Pair: "TRB/USD", Aggr: MedianEOD},
	IntToQueryID(46): {Inactive: true, Pair: "ETH/USD", Aggr: TimeWeightedAvg1h},
	IntToQueryID(47): {Inactive: true, Pair: "BSV/USD", Aggr: Median},
	IntToQueryID(48): {Inactive: true, Pair: "MAKER/USD", Aggr: Median},
	IntToQueryID(49): {Inactive: true, Pair: "BCH/USD", Aggr: TimeWeightedAvg24h},
	IntToQueryID(50): {Pair: "TRB/USD", Aggr: Median, Granularity: defaultGranularity},
	IntToQueryID(51): {Inactive: true, Pair: "XMR/USD", Aggr: Median},
	IntToQueryID(52): {Inactive: true, Pair: "XFT/USD", Aggr: Median},
	IntToQueryID(53): {Inactive: true, Pair: "BTCDOMINANCE", Aggr: Median},
	IntToQueryID(54): {Inactive: true, Pair: "WAVES/USD", Aggr: Median},
	IntToQueryID(55): {Inactive: true, Pair: "OGN/USD", Aggr: Median},
	IntToQueryID(56): {Inactive: true, Pair: "VIXEOD", Aggr: Median},
	IntToQueryID(57): {Inactive: true, Pair: "DEFITVL", Aggr: Median},
	IntToQueryID(58): {Inactive: true, Pair: "DEFIMCAP", Aggr: Mean},
	IntToQueryID(59): {Pair: "ETH/JPY", Aggr: Median, Granularity: defaultGranularity},
	IntToQueryID(60): {Inactive: true, Pair: blocks.MetricSymbolBlockGasPriceAvg, Granularity: defaultGranularity, Aggr: TimeWeightedAvg7Days, ConfidenceQuery: `sum(count_over_time(` + index.MetricIndexValue + `{symbol="` + blocks.MetricSymbolBlockGasPriceAvg + `"}[` + Week.String() + `]))/ sum(` + index.MetricIndexValue + `{symbol="` + blocks.MetricSymbolBlockNum + `"} - ` + index.MetricIndexValue + `{symbol="` + blocks.MetricSymbolBlockNum + `"} offset ` + Week.String() + `)`},
}

func New(logger log.Logger, cfg Config, aggregator *aggregator.Aggregator) *Psr {
	return &Psr{
		logger:     log.With(logger, "component", ComponentName),
		aggregator: aggregator,
		cfg:        cfg,
		errCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      MetricErrCount,
			Help:      "Critical errors in the component",
		}),
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
	errCount   prometheus.Counter
}

func (self *Psr) ConfidenceThreshold(pair string) float64 {
	return self.cfg.MinConfidencePerSymbol[pair]
}

func (self *Psr) GetValue(queryID [32]byte, ts time.Time) (val float64, err error) {
	defer func() {
		if err != nil {
			self.errCount.Inc()
		}
	}()
	psr, ok := Psrs[queryID]
	if !ok {
		return 0, errors.Errorf("invalid queryID:%v", queryID)
	}

	if psr.Inactive {
		level.Info(self.logger).Log("msg", "adding 0 for inactive PSR", "queryID", common.Bytes2Hex(queryID[:]))
		return 0, nil
	}
	val, err = self.getValue(psr, ts)

	return math.Round(val * psr.Granularity), err
}

func (self *Psr) getValue(psr PsrID, ts time.Time) (float64, error) {
	var val float64
	var err error

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

	if confS, ok := self.cfg.MinConfidencePerSymbol[psr.Pair]; ok {
		if conf < confS {
			return 0, errors.Errorf("not enough confidence based on the aggregator calculations - query:%v, value:%v, conf:%v,symbol confidence threshold:%v", psr.Pair, val, conf, self.cfg.MinConfidencePerSymbol[psr.Pair])
		}
	} else if conf < self.cfg.MinConfidenceDefault {
		return 0, errors.Errorf("not enough confidence based on the aggregator calculations - query:%v, value:%v, conf:%v,default confidence threshold:%v", psr.Pair, val, conf, self.cfg.MinConfidenceDefault)
	}

	return val, err
}

func IntToQueryID(i int64) [32]byte {
	// Ids under 100 are encoded differently so that the legacy values can be red from the contract.
	if i <= 100 {
		bs := make([]byte, 4) // uint32 uses only 4 bytes
		binary.BigEndian.PutUint32(bs, uint32(i))

		// Add the bytes at the end of the array.
		var a = [32]byte{}
		a[28] = bs[0]
		a[29] = bs[1]
		a[30] = bs[2]
		a[31] = bs[3]
		return a
	}

	return ethereum_t.Keccak256(string(NewQueryData(i)))
}

func QueryIDToInt(i [32]byte) (int64, error) {
	return strconv.ParseInt(common.BytesToHash(i[:]).String()[2:], 16, 64)
}
