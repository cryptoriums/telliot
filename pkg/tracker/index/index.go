// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package index

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cryptoriums/packages/ethereum"
	"github.com/cryptoriums/packages/format"
	http_p "github.com/cryptoriums/packages/http"
	"github.com/cryptoriums/packages/logging"
	"github.com/cryptoriums/telliot/pkg/db"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/itchyny/gojq"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/yalp/jsonpath"
)

const (
	ComponentName = "trackerIndex"

	MetricIndexValue       = ComponentName + "_" + MetricIndexValueSuffix
	MetricIndexValueSuffix = "value"
	MetricInterval         = ComponentName + "_interval"
	MetricRecCount         = ComponentName + "_requests_total"

	paramNameValue = "value"
	paramNameTs    = "timestamp"
)

type DataSource interface {
	// Source returns the data source.
	Source() string
	// Get returns current api value.
	Get(context.Context) (float64, float64, error)
	// The recommended interval for calling the Get method.
	// Some APIs will return an error if called more often
	// Due to API rate limiting of the provider.
	Interval() time.Duration
}

type Parser interface {
	Parse(data []byte, paramName string) (value float64, err error)
}

type Config struct {
	LogLevel  string
	Interval  format.Duration
	IndexFile string
}

type TrackerIndex struct {
	logger      log.Logger
	ctx         context.Context
	stop        context.CancelFunc
	db          *tsdb.DB
	cfg         Config
	dataSources map[string][]DataSource

	mtx              sync.Mutex
	requestCount     map[string]float64
	requestCountErrs map[string]float64

	values *prometheus.GaugeVec
}

func New(
	ctx context.Context,
	logger log.Logger,
	cfg Config,
	db *tsdb.DB,
	envVars map[string]string,
) (*TrackerIndex, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}

	dataSources, err := createDataSources(logger, ctx, cfg, envVars)
	if err != nil {
		return nil, errors.Wrap(err, "create data sources")
	}

	ctx, stop := context.WithCancel(ctx)

	return &TrackerIndex{
		logger:           log.With(logger, "component", ComponentName),
		ctx:              ctx,
		stop:             stop,
		dataSources:      dataSources,
		db:               db,
		cfg:              cfg,
		requestCount:     make(map[string]float64),
		requestCountErrs: make(map[string]float64),
		values: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      MetricIndexValueSuffix,
			Help:      "Values by source and symbol",
		},
			[]string{"source", "domain", "symbol"},
		),
	}, nil
}

func createDataSources(logger log.Logger, ctx context.Context, cfg Config, envVars map[string]string) (map[string][]DataSource, error) {
	// Load index file.
	indexes := make(map[string]Apis)
	{
		file, err := os.Open(cfg.IndexFile)
		if err != nil {
			return nil, errors.Wrap(err, "open index file")
		}
		dec := json.NewDecoder(file)
		dec.DisallowUnknownFields()
		for {
			if err := dec.Decode(&indexes); err == io.EOF {
				break
			} else if err != nil {
				return nil, errors.Wrap(err, "parse config")
			}

		}
	}

	dataSources := make(map[string][]DataSource)
	var err error
	for symbol, api := range indexes {
		for _, endpoint := range api.Endpoints {
			var source DataSource

			// Default value for the api type.
			if endpoint.Type == "" {
				endpoint.Type = httpSource
			}

			// Default value for the parser.
			if endpoint.Parser == "" {
				endpoint.Parser = jsonPathParser
			}

			if api.Interval.Duration == 0 {
				api.Interval = cfg.Interval
			}

			if api.Interval.Duration > db.DbLookback {
				return nil, errors.Errorf("index interval:%v for symbol:%v can't be lower than the default DB look back of:%v, this will cause look up errors", api.Interval.Duration, symbol, db.DbLookback)
			}

			// Fail early when the url has env that is not set.
			url := http_p.ExpandTimeVars(endpoint.URL)
			os.Expand(url, func(key string) string {
				v, ok := envVars[key]
				if !ok {
					err = errors.Errorf("missing required env variable in index url:%v", key)
				}
				return v
			})
			if err != nil {
				return nil, err
			}

			switch endpoint.Type {
			case manual:
				{
					source = NewManual(endpoint.URL, api.Interval.Duration, NewParser(endpoint))
				}
			case httpSource:
				{
					source = NewJSONapi(api.Interval.Duration, endpoint.URL, NewParser(endpoint), envVars)
					if strings.Contains(strings.ToLower(symbol), "volume") {
						source = NewJSONapiVolume(api.Interval.Duration, endpoint.URL, NewParser(endpoint), envVars)
					}
				}
			case bravenewcoin:
				{
					source, err = NewBravenewcoin(api.Interval.Duration, endpoint.URL, NewParser(endpoint), envVars)
					if err != nil {
						return nil, errors.Wrap(err, "creating Bravenewcoin source")
					}
				}
			case ethereumSource:
				{
					client, err := ethereum.NewClient(ctx, logger, envVars)
					if err != nil {
						return nil, errors.Wrap(err, "creating ethereum client")
					}

					// Validate and pick an ethereum address for current network id.
					address, err := ethereum.GetAddressForNetwork(endpoint.URL, client.NetworkID())
					if err != nil {
						return nil, errors.Wrap(err, "getting address for network id")
					}
					if endpoint.Parser == uniswapParser {
						source = NewUniswap(symbol, address, api.Interval.Duration, client)

					} else if endpoint.Parser == balancerParser {
						source = NewBalancer(symbol, address, api.Interval.Duration, client)
					} else {
						return nil, errors.Wrapf(err, "unknown source for on-chain index tracker")
					}
				}
			default:
				return nil, errors.Errorf("unknown index type for index object:%v", endpoint.Type)
			}

			dataSources[symbol] = append(dataSources[symbol], source)
		}

	}
	return dataSources, nil

}

func (self *TrackerIndex) Start() error {
	level.Info(self.logger).Log("msg", "starting", "logLevel", self.cfg.LogLevel)
	delay := time.Second
	for symbol, dataSources := range self.dataSources {
		for _, dataSource := range dataSources {
			go self.record(delay, symbol, dataSource)
			delay += time.Second
		}
	}
	<-self.ctx.Done()
	return nil
}

// record from all API calls.
// The request delay is used to avoid rate limiting at startup
// for when all API calls try to happen at the same time.
func (self *TrackerIndex) record(delay time.Duration, symbol string, dataSource DataSource) {
	delayTicker := time.NewTicker(delay)
	select {
	case <-delayTicker.C:
		break
	case <-self.ctx.Done():
		level.Debug(self.logger).Log("msg", "values record loop exited")
		return
	}
	delayTicker.Stop()

	logger := log.With(self.logger, "source", dataSource.Source())

	url, err := url.Parse(dataSource.Source())
	if err != nil {
		level.Error(logger).Log("msg", "parsing url from data source", "err", err)
		return
	}

	ticker := time.NewTicker(1)
	defer ticker.Stop()
	var resetTicker sync.Once
	for {
		select {
		case <-self.ctx.Done():
			level.Debug(logger).Log("msg", "values record loop exited")
			return
		case <-ticker.C:
			resetTicker.Do(func() { ticker.Reset(dataSource.Interval()) })
		}

		func() {
			self.mtx.Lock()
			defer self.mtx.Unlock()
			self.requestCount[url.String()]++

			// Record the source interval to use it for the confidence calculation.
			// Confidence = avg(actualSamplesCount/expectedMaxSamplesCount) for a given period.
			if err := self.recordInterval(dataSource.Interval(), symbol, url); err != nil {
				level.Error(logger).Log("msg", "record interval", "err", err)
			}

			if err := self.recordValue(symbol, dataSource); err != nil {
				self.requestCountErrs[url.String()]++
				level.Error(logger).Log("msg", "record value", "err", err)
			}

			// Record the count metric.
			{
				lbls := labels.Labels{
					labels.Label{Name: "__name__", Value: MetricRecCount},
					labels.Label{Name: "source", Value: url.String()},
					labels.Label{Name: "domain", Value: url.Host},
					labels.Label{Name: "status", Value: "success"},
				}
				if err := db.Add(self.ctx, self.db, lbls, self.requestCount[dataSource.Source()]); err != nil {
					level.Error(logger).Log("msg", "add values to the DB", "err", err)
				}
			}

			// Record the count err metric.
			{
				lbls := labels.Labels{
					labels.Label{Name: "__name__", Value: MetricRecCount},
					labels.Label{Name: "source", Value: url.String()},
					labels.Label{Name: "domain", Value: url.Host},
					labels.Label{Name: "status", Value: "fail"},
				}
				if err := db.Add(self.ctx, self.db, lbls, self.requestCountErrs[dataSource.Source()]); err != nil {
					level.Error(logger).Log("msg", "add values to the DB", "err", err)
				}
			}
		}()
	}
}

func (self *TrackerIndex) recordInterval(interval time.Duration, symbol string, url *url.URL) (err error) {
	lbls := labels.Labels{
		labels.Label{Name: "__name__", Value: MetricInterval},
		labels.Label{Name: "source", Value: url.String()},
		labels.Label{Name: "domain", Value: url.Host},
		labels.Label{Name: "symbol", Value: symbol},
	}

	if err := db.Add(self.ctx, self.db, lbls, float64(interval)); err != nil {
		return errors.Wrap(err, "add values to the DB")
	}
	return nil
}

func (self *TrackerIndex) recordValue(symbol string, dataSource DataSource) error {
	source, err := url.Parse(dataSource.Source())
	if err != nil {
		return errors.Wrap(err, "parsing url from data source")
	}

	ctx, cncl := context.WithTimeout(self.ctx, time.Minute)
	defer cncl()

	value, timestamp, err := dataSource.Get(ctx)
	if err != nil {
		return errors.Wrap(err, "getting values from data source")
	}

	// When value is older then the interval can skip it
	// as it makes sense to add only recent data.
	valueAge := time.Since(time.Unix(int64(timestamp), 0))
	// Few times the data source interval
	// as some exchanges don't process many orders
	// and don't update the price so often.
	// Still a good failsafe to not use data that is hours or days old.
	if float64(valueAge) > math.Max(float64(10*time.Minute), float64(10*dataSource.Interval())) {
		return errors.Errorf("data source returned old data valueAge:%v, dataSource interval:%v", valueAge, dataSource.Interval())
	}

	self.values.WithLabelValues(dataSource.Source(), source.Host, symbol).Set(value)

	lbls := labels.Labels{
		labels.Label{Name: "__name__", Value: MetricIndexValue},
		labels.Label{Name: "source", Value: dataSource.Source()},
		labels.Label{Name: "domain", Value: source.Host},
		labels.Label{Name: "symbol", Value: symbol},
	}
	if err = db.Add(self.ctx, self.db, lbls, value); err != nil {
		return errors.Wrap(err, "append values to the DB")
	}

	return nil
}

func (self *TrackerIndex) Stop() {
	self.stop()
}

// IndexType -> index type for Api.
type IndexType string

const (
	httpSource     IndexType = "http"
	bravenewcoin   IndexType = "bravenewcoin"
	ethereumSource IndexType = "ethereum"
	manual         IndexType = "manual"
)

// ParserType -> index parser for Api.
type ParserType string

const (
	jsonPathParser ParserType = "jsonPath"
	jqParser       ParserType = "jq"
	uniswapParser  ParserType = "Uniswap"
	balancerParser ParserType = "Balancer"
)

type Endpoint struct {
	URL        string
	Type       IndexType
	Parser     ParserType
	Value      string
	Time       string
	TimeFormat string
}

// Apis will be used in parsing index file.
type Apis struct {
	// The recommended interval for calling the Get method.
	// Some APIs will return an error if called more often
	// Due to API rate limiting of the provider.
	Interval  format.Duration
	Endpoints []Endpoint
}

type Manual struct {
	data []byte
	Parser
	interval time.Duration
}

func NewManual(data string, interval time.Duration, parser Parser) *Manual {
	return &Manual{
		[]byte(data), parser, interval}
}

func (self *Manual) Get(ctx context.Context) (float64, float64, error) {
	value, err := self.Parse(self.data, paramNameValue)
	if err != nil {
		return 0, 0, errors.Wrapf(err, "parsing value")
	}
	timestamp, err := self.Parse(self.data, paramNameTs)
	if err != nil {
		return 0, 0, errors.Wrapf(err, "parsing timestamp")
	}
	return value, timestamp, err
}

func (self *Manual) Interval() time.Duration {
	return self.interval
}

func (self *Manual) Source() string {
	return "http://manual/" + string(self.data)
}

// NewJSONapiVolume are treated differently and return 0 values when the api returns the same timestamp.
// This is to avoid double counting volumes for the same time period.
// Another way is to skip adding the data, but this messes up the confidence calculations
// which counts total added data points.
func NewJSONapiVolume(interval time.Duration, url string, parser Parser, envVars map[string]string) *JSONapiVolume {
	return &JSONapiVolume{
		JSONapi: NewJSONapi(interval, url, parser, envVars),
	}
}

type URLEnvExpander struct {
	url     string
	envVars map[string]string
}

func (self *URLEnvExpander) UrlExpanded() (string, error) {
	var err error
	url := http_p.ExpandTimeVars(self.url)
	url = os.Expand(url, func(key string) string {
		v, ok := self.envVars[key]
		if !ok {
			err = errors.Errorf("missing required env variable in index url:%v", key)
		}
		return v
	})
	if err != nil {
		return "", err
	}
	return url, err
}

func (self *URLEnvExpander) UrlRaw() string {
	return self.url
}

type JSONapiVolume struct {
	*JSONapi
	lastTS float64
}

func (self *JSONapiVolume) Get(ctx context.Context) (float64, float64, error) {
	url, err := self.UrlExpanded()
	if err != nil {
		return 0, 0, errors.Wrap(err, "getting api url")
	}
	vals, err := http_p.Get(ctx, url, nil)
	if err != nil {
		return 0, 0, errors.Wrapf(err, "fetching data from API url:%v", url)
	}
	value, err := self.Parse(vals, paramNameValue)
	if err != nil {
		return 0, 0, errors.Wrapf(err, "parsing data from API url:%v", url)
	}
	timestamp, err := self.Parse(vals, paramNameTs)
	if err != nil {
		return 0, 0, errors.Wrapf(err, "parsing data from API url:%v", url)
	}

	// Use 0 value for the volume as this has already been requested.
	if self.lastTS == timestamp {
		value = 0
	}
	self.lastTS = timestamp

	return value, timestamp, nil

}

func NewJSONapi(interval time.Duration, url string, parser Parser, envVars map[string]string) *JSONapi {
	return &JSONapi{
		interval:    interval,
		Parser:      parser,
		URLExpander: &URLEnvExpander{url: url, envVars: envVars},
	}
}

func NewBravenewcoin(interval time.Duration, urlString string, parser Parser, envVars map[string]string) (*Bravenewcoin, error) {
	return &Bravenewcoin{
		JSONapi: NewJSONapi(interval, urlString, parser, envVars),
	}, nil
}

type Bravenewcoin struct {
	apiKey      string
	bearerToken string
	*JSONapi
}

func (self *Bravenewcoin) Get(ctx context.Context) (float64, float64, error) {
	withRefresh := false
	if self.apiKey == "" || self.bearerToken == "" {
		withRefresh = true
	}
	return self.get(ctx, withRefresh, true)
}

func (self *Bravenewcoin) get(ctx context.Context, withRefresh, withRetry bool) (float64, float64, error) {
	if withRefresh {
		if err := self.refreshApiKey(); err != nil {
			return 0, 0, errors.Wrap(err, "refresh api key")
		}
		if err := self.refreshBearer(); err != nil {
			return 0, 0, errors.Wrap(err, "refresh bearer token")
		}
	}
	headers := make(map[string]string)
	headers["x-rapidapi-key"] = self.apiKey
	headers["authorization"] = "Bearer " + self.bearerToken
	headers["x-rapidapi-host"] = "bravenewcoin.p.rapidapi.com"

	urlExpanded, err := self.UrlExpanded()
	if err != nil {
		return 0, 0, errors.Wrap(err, "getting api url")
	}

	result, err := http_p.Get(ctx, urlExpanded, headers)
	if err != nil {
		if withRetry {
			return self.get(ctx, true, false)
		}
		return 0, 0, errors.Wrap(err, "fetching data")
	}
	value, err := self.Parse(result, paramNameValue)
	if err != nil {
		return 0, 0, errors.Wrapf(err, "parsing value")
	}
	timestamp, err := self.Parse(result, paramNameTs)
	if err != nil {
		return 0, 0, errors.Wrapf(err, "parsing timestamp")
	}
	return value, timestamp, err
}
func (self *Bravenewcoin) refreshApiKey() error {
	urlExpanded, err := self.UrlExpanded()
	if err != nil {
		return errors.Wrap(err, "expanding url")
	}
	u, err := url.Parse(urlExpanded)
	if err != nil {
		return errors.Wrap(err, "parse bravecoin url")
	}
	apiKey := u.Query().Get("rapidapi-key")
	if apiKey == "" {
		return errors.New("rapid api key is empty")
	}
	self.apiKey = apiKey
	return nil
}

func (self *Bravenewcoin) refreshBearer() error {
	url := "https://bravenewcoin.p.rapidapi.com/oauth/token?rapidapi-key=" + self.apiKey

	payload := strings.NewReader("{\n    \"audience\": \"https://api.bravenewcoin.com\",\n    \"client_id\": \"oCdQoZoI96ERE9HY3sQ7JmbACfBf55RY\",\n    \"grant_type\": \"client_credentials\"\n}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return errors.Wrap(err, "create client request")
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-rapidapi-host", "bravenewcoin.p.rapidapi.com")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr}

	res, err := client.Do(req)

	if err != nil {
		return errors.Wrap(err, "client request")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrapf(err, "read body")
	}

	if res.StatusCode/100 != 2 {
		return errors.Wrapf(err, "request status not ok:%v", string(body))
	}

	output := struct {
		Access_token string `json:"access_token,omitempty"`
	}{}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return errors.Wrapf(err, "json marshal:%v", string(body))
	}

	self.bearerToken = output.Access_token
	return nil
}

type JSONapi struct {
	interval time.Duration
	Parser
	URLExpander
}

type URLExpander interface {
	UrlExpanded() (string, error)
	UrlRaw() string
}

func (self *JSONapi) Get(ctx context.Context) (float64, float64, error) {
	url, err := self.UrlExpanded()
	if err != nil {
		return 0, 0, errors.Wrapf(err, "getting api url")
	}
	result, err := http_p.Get(ctx, url, nil)
	if err != nil {
		return 0, 0, errors.Wrapf(err, "fetching data from API url:%v", url)
	}
	value, err := self.Parse(result, paramNameValue)
	if err != nil {
		return 0, 0, errors.Wrapf(err, "parsing value")
	}
	timestamp, err := self.Parse(result, paramNameTs)
	if err != nil {
		return 0, 0, errors.Wrapf(err, "parsing timestamp")
	}
	return value, timestamp, err
}

func (self *JSONapi) Interval() time.Duration {
	return self.interval
}

func (self *JSONapi) Source() string {
	return self.UrlRaw()
}

type params struct {
	valueParam      string
	timeParam       string
	timeFormatParam string
}

type JsonPathParser struct {
	params
}

func (self *JsonPathParser) Parse(input []byte, paramName string) (float64, error) {
	switch paramName {
	case paramNameValue:
		result, err := self.parse(input, self.valueParam, false)
		if err != nil {
			return 0, errors.Wrap(err, "parsing value")
		}
		return result, nil
	case paramNameTs:
		if self.timeParam != "" {
			result, err := self.parse(input, self.timeParam, true)
			if err != nil {
				return 0, errors.Wrap(err, "parsing timestamp")
			}
			if result > 9999999999 { // The TS is with Millisecond granularity.
				result = result / 1000
			}
			if result <= 0 {
				return 0, errors.Errorf("timestamp parsing returns invalid number:%v", result)
			}
			return result, nil
		}

		return float64(time.Now().Unix()), nil
	default:
		return 0, errors.New("invalid param name:" + paramName)
	}
}

func (self *JsonPathParser) parse(input []byte, query string, isTime bool) (float64, error) {
	maxErrL := len(string(input)) - 1
	if maxErrL > 200 {
		maxErrL = 200
	}

	var inputToParse interface{}
	err := json.Unmarshal(input, &inputToParse)
	if err != nil {
		return 0, errors.Wrapf(err, "json marshal:%v", string(input)[:maxErrL])
	}

	out, err := jsonpath.Read(inputToParse, query)
	if err != nil {
		return 0, errors.Wrapf(err, "json path read:%v", string(input)[:maxErrL])
	}
	str := fmt.Sprintf("%v", out)

	return parseStrToFloat(str, isTime, self.timeFormatParam)
}

type JqParser struct {
	params
}

func (self *JqParser) Parse(input []byte, paramName string) (float64, error) {

	switch paramName {
	case paramNameValue:
		result, err := self.parse(input, self.valueParam, false)
		if err != nil {
			return 0, errors.Wrap(err, "parsing value")
		}
		return result, nil
	case paramNameTs:
		if self.timeParam != "" {
			result, err := self.parse(input, self.timeParam, true)
			if err != nil {
				return 0, errors.Wrap(err, "parsing timestamp")
			}
			if result > 9999999999 { // The TS is with Millisecond granularity.
				result = result / 1000
			}
			if result <= 0 {
				return 0, errors.Errorf("timestamp parsing returns invalid number:%v", result)
			}
			return result, nil
		}

		return float64(time.Now().Unix()), nil
	default:
		return 0, errors.New("invalid param name:" + paramName)
	}
}

func (self *JqParser) parse(input []byte, _query string, isTime bool) (float64, error) {
	maxErrL := len(string(input)) - 1
	if maxErrL > 200 {
		maxErrL = 200
	}

	var inputToParse interface{}
	err := json.Unmarshal(input, &inputToParse)
	if err != nil {
		return 0, errors.Wrapf(err, "json marshal:%v", string(input)[:maxErrL])
	}

	query, err := gojq.Parse(_query)
	if err != nil {
		return 0, errors.Wrapf(err, "jq qyery:%v", string(input)[:maxErrL])
	}

	iter := query.Run(inputToParse)

	output, ok := iter.Next()
	if !ok {
		return 0, errors.Wrapf(err, "jq iterate:%v", string(input)[:maxErrL])
	}

	_, ok = iter.Next()
	if ok {
		return 0, errors.Errorf("jq parsing contains multiple values:%v", string(input)[:maxErrL])
	}

	if err, ok := output.(error); ok {
		return 0, errors.Wrapf(err, "jq parse:%v", string(input)[:maxErrL])
	}

	if v, ok := output.(float64); ok {
		return v, nil
	}
	if v, ok := output.(string); ok {
		return parseStrToFloat(v, isTime, self.timeFormatParam)
	}

	return 0, errors.Errorf("jq results not a supported type:%+v", output)
}

func NewParser(t Endpoint) Parser {
	switch t.Parser {
	case jsonPathParser:
		return &JsonPathParser{
			params{
				valueParam:      t.Value,
				timeParam:       t.Time,
				timeFormatParam: t.TimeFormat,
			},
		}
	case jqParser:
		return &JqParser{
			params{
				valueParam:      t.Value,
				timeParam:       t.Time,
				timeFormatParam: t.TimeFormat,
			},
		}
	default:
		return nil
	}
}

func parseStrToFloat(input string, isTime bool, timeFormat string) (float64, error) {
	var value float64
	if !isTime {
		return strconv.ParseFloat(input, 64)
	}

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		if timeFormat == "" {
			timeFormat = time.RFC3339Nano
		}
		t, err := time.Parse(timeFormat, input)
		if err != nil {
			return 0, errors.Wrapf(err, "value not a supported time format or a timestamp:%v", input)
		}
		return float64(t.Unix()), nil

	}
	return value, nil
}
