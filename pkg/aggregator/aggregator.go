// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package aggregator

import (
	"context"
	"math"
	"sort"
	"strconv"
	"time"

	"github.com/cryptoriums/packages/format"
	"github.com/cryptoriums/packages/logging"
	"github.com/cryptoriums/telliot/pkg/db"
	"github.com/cryptoriums/telliot/pkg/tracker/index"
	"github.com/go-kit/log"
	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/storage"
)

const ComponentName = "aggregator"

type IAggregator interface {
	Min(symbol string, domain string, start time.Time, lookBack time.Duration) (float64, float64, error)
}

type Config struct {
	LogLevel string
}

type Aggregator struct {
	logger       log.Logger
	ctx          context.Context
	tsDB         storage.Queryable
	promqlEngine *promql.Engine
	cfg          Config
}

func New(
	ctx context.Context,
	logger log.Logger,
	cfg Config,
	tsDB storage.Queryable,
) (*Aggregator, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}

	engine := promql.NewEngine(db.NewPromqlEngineOpts(logger))

	return &Aggregator{
		logger:       log.With(logger, "component", ComponentName),
		ctx:          ctx,
		tsDB:         tsDB,
		promqlEngine: engine,
		cfg:          cfg,
	}, nil
}

func (self *Aggregator) MedianAt(symbol string, at time.Time) (float64, float64, error) {
	vals, confidence, err := self.valsAtWithConfidence(symbol, at)
	if err != nil {
		return 0, 0, err
	}
	if len(vals) == 0 {
		return 0, 0, errors.Errorf("no vals at:%v", at)
	}
	median, confidenceM := self.median(vals)
	if confidenceM < confidence {
		confidence = confidenceM
	}

	return median, confidence, nil
}

func (self *Aggregator) MedianAtEOD(symbol string, at time.Time) (float64, float64, error) {
	return self.MedianAt(symbol, time.Unix(format.EOD(), 0))
}

func (self *Aggregator) MeanAt(symbol string, at time.Time) (float64, float64, error) {
	vals, confidence, err := self.valsAtWithConfidence(symbol, at)
	if err != nil {
		return 0, 0, err
	}
	price, confidenceM := self.mean(vals)
	if confidenceM < confidence {
		confidence = confidenceM
	}
	return price, confidence * 100, nil
}

func (self *Aggregator) mean(vals []float64) (float64, float64) {
	if len(vals) == 1 {
		return vals[0], 100
	}

	priceSum := 0.0
	min, max := vals[0], vals[0]
	for _, val := range vals {
		priceSum += val
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return priceSum / float64(len(vals)), ConfidenceInDifference(min, max)
}

func (self *Aggregator) TimeWeightedAvg(
	symbol string,
	confidenceQuery string,
	start time.Time,
	lookBack time.Duration,
) (float64, float64, error) {
	return self.overTime("avg_over_time", symbol, "", confidenceQuery, start, lookBack)
}

func (self *Aggregator) Min(
	symbol string,
	domain string,
	start time.Time,
	lookBack time.Duration,
) (float64, float64, error) {
	return self.overTime("min_over_time", symbol, domain, "", start, lookBack)
}

// overTime runs an  aggregation function over a period of time and returns price and confidence level for a given symbol.
// Confidence is calculated based on maximum possible samples over the actual samples for a given period.
// avg(maxPossibleSamplesCount/actualSamplesCount)
// For example with 1h look back and source interval of 60sec maxPossibleSamplesCount = 36
// with actualSamplesCount = 18 this is 50% confidence.
// The same calculation is done for all source and the final value is the average of all.
//
// Example for 1h.
// maxDataPointCount is calculated by deviding the seconds in 1h by how often the tracker queries the APIs.
// avg(count_over_time(trackerIndex_value{symbol="AMPL_USD"}[1h]) / (3.6e+12/trackerIndex_interval)).
func (self *Aggregator) overTime(
	funcName string,
	symbol string,
	domain string,
	confidenceQuery string,
	start time.Time,
	lookBack time.Duration,
) (float64, float64, error) {
	// Confidence level.
	if confidenceQuery == "" {
		resolution, err := self.resolution(symbol, start)
		if err != nil {
			return 0, 0, err
		}

		confidenceQuery = `avg(
			count_over_time(
				` + index.MetricIndexValue + `{ symbol="` + symbol + `" }
			[` + lookBack.String() + `:` + resolution.String() + `])
			/
			(` + strconv.Itoa(int(lookBack.Nanoseconds())) + ` / ` + strconv.Itoa(int(resolution.Nanoseconds())) + `)
		)`
	}
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		confidenceQuery,
		start,
	)
	if err != nil {
		return 0, 0, errors.Wrap(err, "creating query")
	}
	defer query.Close()
	_confidence := query.Exec(self.ctx)
	if _confidence.Err != nil {
		return 0, 0, errors.Wrapf(_confidence.Err, "evaluating query:%v", query.Statement())
	}
	if len(_confidence.Value.(promql.Vector)) == 0 {
		return 0, 0, errors.Errorf("no result for TWAP confidence query:%v", query.Statement())
	}
	confidence := _confidence.Value.(promql.Vector)[0].V * 100

	domainQ := ""
	if domain != "" {
		domainQ = `,domain="` + domain + `"`
	}
	// Avg value over the look back period.
	query, err = self.promqlEngine.NewInstantQuery(
		self.tsDB,
		funcName+`(
				`+index.MetricIndexValue+`{symbol="`+symbol+`"`+domainQ+`}
		[`+lookBack.String()+`:])
		`,
		start,
	)
	if err != nil {
		return 0, 0, errors.Wrap(err, "creating query")
	}
	defer query.Close()
	vals := query.Exec(self.ctx)
	if vals.Err != nil {
		return 0, 0, errors.Wrapf(vals.Err, "evaluating query:%v", query.Statement())
	}
	if len(vals.Value.(promql.Vector)) == 0 {
		return 0, 0, errors.Errorf("no result for TWAP vals query:%v", query.Statement())
	}

	var prices []float64
	for _, price := range vals.Value.(promql.Vector) {
		prices = append(prices, price.V)
	}

	price, confidenceM := self.mean(prices)
	if confidenceM < confidence {
		confidence = confidenceM
	}

	return price, confidence, err
}

// VolumWeightedAvg returns price and confidence level for a given symbol.
// Confidence is calculated based on maximum possible samples over the actual samples for a given period.
// avg(maxPossibleSamplesCount/actualSamplesCount)
// For example with 1h look back and source interval of 60sec maxPossibleSamplesCount = 36
// with actualSamplesCount = 18 this is 50% confidence.
// The same calculation is done for all source and the final value is the average of all.
// maxDataPointCount is calculated by deviding the seconds in 1h by how often the tracker queries the APIs.
//
// Example for 1h.
// avg(count_over_time(trackerIndex_value{symbol="AMPL_USD"}[1h]) / (3.6e+12/30s)).
//
// vals are calculated using the official VWAP formula from
// https://tradingtuitions.com/vwap-trading-strategy-excel-sheet/
// UNUSED - Need to add logic to to compare the VWAP results from all APIS and reduce the confidence if the difference is too much like in the other aggregators.
//func (self *Aggregator) VolumWeightedAvg(
// 	symbol string,
// 	start time.Time,
// 	end time.Time,
// 	aggrWindow time.Duration,
// ) (float64, float64, error) {
// 	_timeWindow := end.Sub(start).Round(time.Minute).Seconds()
// 	timeWindow := strconv.Itoa(int(_timeWindow)) + "s"

// 	resolution, err := self.resolution(symbol, end)
// 	if err != nil {
// 		return 0, 0, err
// 	}

// 	query, err := self.promqlEngine.NewInstantQuery(
// 		self.tsDB,
// 		`avg(
// 			sum_over_time(
// 				(
// 					sum_over_time( `+index.MetricIndexValue+`{symbol="`+symbol+`_VOLUME"}[`+aggrWindow.String()+`]
// 					) * on(domain)
// 					avg_over_time(`+index.MetricIndexValue+`{symbol="`+symbol+`"}[`+aggrWindow.String()+`])
// 				)
// 			[`+timeWindow+`:`+aggrWindow.String()+`])
// 			/ on(domain)
// 			sum_over_time(
// 					sum_over_time(`+index.MetricIndexValue+`{symbol="`+symbol+`_VOLUME"}[`+aggrWindow.String()+`])
// 			[`+timeWindow+`:`+aggrWindow.String()+`])
// 		)`,
// 		end,
// 	)

// 	if err != nil {
// 		return 0, 0, errors.Wrap(err, "creating query")
// 	}
// 	defer query.Close()

// 	_result := query.Exec(self.ctx)
// 	if _result.Err != nil {
// 		return 0, 0, errors.Wrapf(_result.Err, "evaluating query:%v", query.Statement())
// 	}
// 	result := _result.Value.(promql.Vector)
// 	if len(result) == 0 {
// 		return 0, 0, errors.Errorf("no result for VWAP vals query:%v", query.Statement())
// 	}

// 	// Confidence level for prices.
// 	query, err = self.promqlEngine.NewInstantQuery(
// 		self.tsDB,
// 		`avg(
// 			count_over_time(`+index.MetricIndexValue+`{symbol="`+symbol+`"}[`+timeWindow+`])
// 			/
// 			(`+strconv.Itoa(int(end.Sub(start).Nanoseconds()))+` / `+strconv.Itoa(int(resolution.Nanoseconds()))+`)
// 		)`,
// 		end,
// 	)
// 	if err != nil {
// 		return 0, 0, errors.Wrap(err, "creating query")
// 	}
// 	defer query.Close()

// 	confidenceP := query.Exec(self.ctx)
// 	if confidenceP.Err != nil {
// 		return 0, 0, errors.Wrapf(confidenceP.Err, "evaluating query:%v", query.Statement())
// 	}

// 	// Confidence level for volumes.
// 	resolution, err = self.resolution(symbol+"/VOLUME", end)
// 	if err != nil {
// 		return 0, 0, err
// 	}
// 	query, err = self.promqlEngine.NewInstantQuery(
// 		self.tsDB,
// 		`avg(
// 			count_over_time(`+index.MetricIndexValue+`{symbol="`+symbol+`_VOLUME"}[`+timeWindow+`])
// 			/
// 			(`+strconv.Itoa(int(end.Sub(start).Nanoseconds()))+` / `+strconv.Itoa(int(resolution.Nanoseconds()))+`)
// 		)`,
// 		end,
// 	)
// 	if err != nil {
// 		return 0, 0, errors.Wrap(err, "creating query")
// 	}
// 	defer query.Close()
// 	confidenceV := query.Exec(self.ctx)
// 	if confidenceV.Err != nil {
// 		return 0, 0, errors.Wrapf(confidenceV.Err, "evaluating query:%v", query.Statement())
// 	}

// 	if len(confidenceP.Value.(promql.Vector)) == 0 || len(confidenceV.Value.(promql.Vector)) == 0 {
// 		return 0, 0, errors.Errorf("no result for VWAP confidence query:%v", query.Statement())
// 	}

// 	// Use the smaller confidence of volume or value.
// 	confidence := confidenceP.Value.(promql.Vector)[0].V
// 	if confidence > confidenceV.Value.(promql.Vector)[0].V {
// 		confidence = confidenceV.Value.(promql.Vector)[0].V
// 	}

// 	// Return the last VWAP price.
// 	return result[len(result)-1].V, confidence * 100, nil
// }

// median returns the median and the confidence calculated with
// the difference between the min and max values.
func (self *Aggregator) median(vals []float64) (float64, float64) {
	if len(vals) == 1 {
		return vals[0], 100
	}
	sort.Slice(vals, func(i, j int) bool {
		return vals[i] < vals[j]
	})

	position := len(vals) / 2
	price := vals[position]

	// When number of vals is even need to use the mean
	// of the 2 middle vals.
	if len(vals)%2 == 0 {
		price = (vals[position-1] + vals[position]) / 2
	}

	return price, ConfidenceInDifference(vals[0], vals[len(vals)-1])
}

// ConfidenceInDifference calculates the percentage difference between the max and min and subtract this from 100%.
// Example:
// min 1, max 2
// Difference is 1 which is 100% so the final confidence is 100-100 equals 0%.
func ConfidenceInDifference(min, max float64) float64 {
	return 100 - (math.Abs(min-max)/min)*100
}

// valsAtWithConfidence returns the value from all sources for a given symbol with the confidence level.
// Confidence drops to 50% when there is only one api source.
func (self *Aggregator) valsAtWithConfidence(symbol string, at time.Time) ([]float64, float64, error) {
	resolution, err := self.resolution(symbol, at)
	if err != nil {
		return nil, 0, err
	}
	lookBack := time.Duration(resolution + 1e+9) // 1 sec more then the pull interval to make sure the tracker has added a value. Interval is in nanosecond granularity.
	var prices []float64
	pricesVector, err := self.valsAt(symbol, at, lookBack)
	if err != nil {
		return nil, 0, err
	}

	for _, price := range pricesVector {
		prices = append(prices, price.V)
	}

	// Also add an avg price as if there is too much deviation this means that something is wrong
	//  and the values difference will trigger lower confidence.
	avg, _, err := self.TimeWeightedAvg(symbol, "", at, time.Hour)
	if err != nil {
		return nil, 0, errors.Wrap(err, "getting avg for calculating the conf")
	}
	prices = append(prices, avg)

	// Confidence level.
	// Max available endpoint/ currently available endpoints.
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`100 * (count(`+index.MetricIndexValue+`{symbol="`+symbol+`"} )/ max_over_time(count(`+index.MetricInterval+`{symbol="`+symbol+`"})[1h:]))`,
		at,
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "creating query")
	}
	defer query.Close()
	confidence := query.Exec(self.ctx)
	if confidence.Err != nil {
		return nil, 0, errors.Wrapf(confidence.Err, "evaluating query:%v", query.Statement())
	}
	if len(confidence.Value.(promql.Vector)) == 0 {
		return nil, 0, errors.Errorf("no vals for confidence at:%v, query:%v", at, query.Statement())
	}

	return prices, confidence.Value.(promql.Vector)[0].V, nil
}

// valsAt returns all vals from all indexes at a given time.
func (self *Aggregator) valsAt(symbol string, at time.Time, lookBack time.Duration) (promql.Vector, error) {
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`last_over_time( `+index.MetricIndexValue+`{symbol="`+symbol+`"} [`+lookBack.String()+`])`,
		at,
	)
	if err != nil {
		return nil, errors.Wrap(err, "creating query")
	}
	defer query.Close()
	result := query.Exec(self.ctx)
	if result.Err != nil {
		return nil, errors.Wrapf(result.Err, "evaluating query:%v", query.Statement())
	}

	return result.Value.(promql.Vector), nil
}

func (self *Aggregator) resolution(symbol string, at time.Time) (time.Duration, error) {
	query, err := self.promqlEngine.NewInstantQuery(
		self.tsDB,
		`last_over_time(`+index.MetricInterval+`{symbol="`+symbol+`"}[3h])`, // The interval is recorded on every index tracker cycle so this lookback should be sufficient.
		at,
	)
	if err != nil {
		return 0, errors.Wrap(err, "creating query")
	}
	defer query.Close()
	_trackerInterval := query.Exec(self.ctx)
	if _trackerInterval.Err != nil {
		return 0, errors.Wrapf(_trackerInterval.Err, "evaluating query:%v", query.Statement())
	}
	if len(_trackerInterval.Value.(promql.Vector)) == 0 {
		return 0, errors.Errorf("no vals for tracker interval at:%v, query:%v", at, query.Statement())
	}

	return time.Duration(_trackerInterval.Value.(promql.Vector)[0].V), nil
}
