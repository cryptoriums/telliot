// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package web

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/http/pprof"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/cryptoriums/packages/ethereum"
	"github.com/cryptoriums/packages/logging"
	math_t "github.com/cryptoriums/packages/math"
	"github.com/cryptoriums/packages/private_file"
	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/db"
	"github.com/cryptoriums/telliot/pkg/format"
	psr_tellor "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/cryptoriums/telliot/pkg/web/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/route"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/storage"
	"golang.org/x/tools/godoc/util"
)

const ComponentName = "web"

type Componentor interface {
	Status() bool
	ChangeStatus()
	Name() string
}

type Config struct {
	LogLevel    string
	ListenHost  string
	ListenPort  uint
	ReadTimeout format.Duration
}

type Web struct {
	logger log.Logger
	cfg    Config
	ctx    context.Context
	stop   context.CancelFunc
	srv    *http.Server
}

func New(
	ctx context.Context,
	logger log.Logger,
	handlers map[string]http.HandlerFunc,
	tsDB storage.SampleAndChunkQueryable,
	cfg Config,
) (*Web, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	router := route.New()

	router.Get("/debug/*subpath", serveDebug)
	router.Post("/debug/*subpath", serveDebug)

	for url, handler := range handlers {
		router.Get(url, handler)
		router.Post(url, handler)
	}

	router.Get("/metrics", promhttp.Handler().ServeHTTP)

	engine := promql.NewEngine(db.NewPromqlEngineOpts(logger))

	api := api.New(logger, ctx, engine, tsDB)
	api.Register(router.WithPrefix("/api/v1"))

	mux := http.NewServeMux()
	mux.Handle("/", router)

	srv := &http.Server{
		Handler:     mux,
		ReadTimeout: cfg.ReadTimeout.Duration,
		Addr:        fmt.Sprintf("%s:%d", cfg.ListenHost, cfg.ListenPort),
	}

	ctx, stop := context.WithCancel(ctx)

	return &Web{
		logger: log.With(logger, "component", ComponentName),
		cfg:    cfg,
		ctx:    ctx,
		stop:   stop,
		srv:    srv,
	}, nil

}

func Data(
	ctx context.Context,
	logger log.Logger,
	componentor Componentor,
	accounts []*ethereum.Account,
	client ethereum.EthClient,
	master contracts.TellorMasterCaller,
	oracle contracts.TellorOracleCaller,
	govern contracts.TellorGovernCaller,
	envFilePath string,
) http.HandlerFunc {
	var err error
	accountsMap := make(map[common.Address]struct{})

	for _, account := range accounts {
		accountsMap[account.Address] = struct{}{}
	}

	t := template.New("template").Funcs(template.FuncMap{
		"timeSince": func(ts int64) int {
			return int(time.Since(time.Unix(ts, 0)).Minutes())
		},
		"reporterFormat": func(addr common.Address) string {
			addrS := addr.Hex()[:8]
			if _, ok := accountsMap[addr]; ok {
				addrS = "<b>" + addrS + "</b>"
			}
			return addrS
		},
		"tsToTime": func(ts int64) string {
			return time.Unix(ts, 0).UTC().Format("15:04")
		},
		"bytesToHash": func(val [32]byte) string {
			return common.BytesToHash(val[:]).String()
		},
		"psrDetails": func(queryID [32]byte) string {
			psr, ok := psr_tellor.Psrs[queryID]
			if !ok {
				return fmt.Sprintf("unable to get psr for queryID:%v ", queryID)
			}

			inactive := ""
			if psr.Inactive {
				inactive = "(in)"
			}
			return strings.Replace(psr.Pair+"_"+psr.Aggr, " ", "_", -1) + inactive
		},

		"applyGranularity": func(val []byte, queryID [32]byte) string {
			psr, ok := psr_tellor.Psrs[queryID]
			if !ok {
				return fmt.Sprintf("unable to get psr for queryID:%v ", queryID)
			}

			return fmt.Sprintf("%.6f", math_t.BigIntToFloat(big.NewInt(0).SetBytes(val))/psr.Granularity)
		},
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.RequestURI)
		values := u.Query()

		var envVars map[string]string
		postResult := ""
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, fmt.Sprintf("parsing form data:%v", err), http.StatusInternalServerError)
				return
			}
			envVars, err = loadEnvVars(client, envFilePath, r.PostForm.Get("pass"))
			if err != nil {
				http.Error(w, fmt.Sprintf("checking  password:%v", err), http.StatusInternalServerError)
				return
			}
			switch r.PostForm.Get("action") {
			case "dispute":
				tx, err := createDispute(ctx, logger, r, client, govern, envVars)
				if err != nil {
					http.Error(w, fmt.Sprintf("creating a dispute:%v", err), http.StatusInternalServerError)
					return
				}
				postResult = `created new dispute: <a href="` + ethereum.GetEtherscanURL(client.NetworkID()) + `/tx/` + tx.Hash().String() + `">` + tx.Hash().String() + `</a><br/><br/>`
			case "changeStatus":
				componentor.ChangeStatus()
			}
		}

		lookBack := time.Hour
		if d, ok := values["look-back"]; ok {
			lookBack, err = time.ParseDuration(d[0])
			if err != nil {
				http.Error(w, fmt.Sprintf("parsing look back param:%v", err), http.StatusInternalServerError)
				return
			}
		}
		ctx, cncl := context.WithTimeout(ctx, 5*time.Minute)
		defer cncl()
		submits, err := contracts.GetSubmitLogs(
			ctx,
			client,
			oracle,
			0,
			lookBack,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("getting submit logs:%v", err), http.StatusInternalServerError)
			return
		}

		// Reverse the order to shows the more recent first.
		sort.Slice(submits[:], func(i, j int) bool {
			return submits[i].Time.Int64() > submits[j].Time.Int64()
		})

		postForm := ``
		if values.Get("ts") != "" && postResult == "" {
			postForm = prepareDisputeForm(ctx, logger, client, master, values, accounts)
		}

		netWarning := ""
		if client.NetworkID() == 1 {
			netWarning = `<b style="color:red">MAINNET</b>`
		}

		lookBackOpts := ""
		for i := 1; i < 48; i++ {
			selected := ""
			if int(lookBack.Hours()) == i {
				selected = `selected="selected"`
			}
			lookBackOpts += `<option ` + selected + ` value="` + strconv.Itoa(i) + `h">` + strconv.Itoa(i) + `h</option>`
		}

		t, err := t.Parse(`
		<!DOCTYPE html>
		<html lang='en'>
		<head>
			<meta name='viewport' content='width=device-width, initial-scale=1, maximum-scale=1'>
			<title>Telliot GUI</title>
			<style>
				body {
					font-family: arial;
					position:absolute;
				}
				label {
					min-width: 9em;
					float: left;
				}
				table {
					border-spacing: 0;
					border-collapse: collapse;
				}
				td, tr {
					padding: 0.2em;
					margin: 0;
				}
				fieldset {
					width: 100%;
					margin: 0;
					min-width: 30em;
				}
				input[type=text] {
					min-width: 70%;
				}
				input:read-only {
					background-color: grey;
				}
			</style>
		</head>
		<body>
		<h2>` + componentor.Name() + ` status: ` + formatStatus(componentor.Status()) + `
			<form style="display:inline" id="changeStatus" method="post">
				<input type="hidden" id="action" name="action" value="changeStatus" >
				<input placeholder="EnvFile Pass" type="password" name="pass" id="pass"/>
				<input type="submit" value="Change Status">
			</form>
		</h2>
		` + netWarning + `
		` + postResult + `
		` + postForm + `
		<form id="data">
			<label for="look-back">Look Back:</label>
			<select name="look-back" id="look-back">
				` + lookBackOpts + `
			</select>
			<input type="submit" value="GO">
		</form>
		<table>
		{{range $index, $submit := .}}
			<tr {{if $submit.Disputed}}style="color:red"{{end}}>
				<td>{{if $submit.Disputed}}DISPUTED - {{end}}Mins:{{  timeSince $submit.Time.Int64 }}</td>
				<td>Time:{{ tsToTime $submit.Time.Int64 }}</td>
				<td>Ts:{{ $submit.Time.String}}</td>
				<td>{{ psrDetails $submit.QueryId}}</td>
				<td style="text-align:right">{{ applyGranularity $submit.Value $submit.QueryId}}</td>
				<td>
					Reporter:<a  href="` + ethereum.GetEtherscanURL(client.NetworkID()) + `/address/{{ $submit.Reporter }}">{{ reporterFormat $submit.Reporter }}</td>
				<td>
					Tx:<a  href="` + ethereum.GetEtherscanURL(client.NetworkID()) + `/tx/{{ $submit.Raw.TxHash.Hex }}">{{ slice ($submit.Raw.TxHash).Hex 0 8 }}</a>
				</td>
				<td>
					{{if  not $submit.Disputed}}
						<form id="data">
							<input type="hidden" id="queryID" name="queryID" value="{{ bytesToHash $submit.QueryId }}" >
							<input type="hidden" id="ts" name="ts"  value="{{$submit.Time.String}}" >
							<input type="hidden" id="look-back" name="look-back"  value="` + strconv.Itoa(int(lookBack.Hours())) + `h" >
							<input type="submit" value="Dispute">
						</form>
					{{end}}
				</td>
			</tr>
		{{end}}
		</table>
		</body>
		</html>
		`)

		if err != nil {
			http.Error(w, fmt.Sprintf("parsing html template submit logs:%v", err), http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, submits)
		if err != nil {
			http.Error(w, fmt.Sprintf("executing the html template:%v", err), http.StatusInternalServerError)
			return
		}
	})
}

func formatStatus(status bool) string {
	switch status {
	case true:
		return `<span style="color:red">enabled</span>`
	default:
		return `<span style="color:grey">disabled</span>`
	}
}

func prepareDisputeForm(
	ctx context.Context,
	logger log.Logger,
	client ethereum.EthClient,
	master contracts.TellorMasterCaller,
	values url.Values,
	accounts []*ethereum.Account,
) string {
	accOpts := ""
	for _, account := range accounts {
		metaData := ""
		status, _, err := master.GetStakerInfo(&bind.CallOpts{Context: ctx}, account.Address)
		if err != nil {
			level.Error(logger).Log("msg", "get stake status", "addr", account.Address.Hex()[:8], "err", err)
		} else {
			metaData += contracts.ReporterStatusName(status.Int64()) + " "
		}
		trbBalance, err := master.BalanceOf(&bind.CallOpts{Context: ctx}, account.Address)
		if err != nil {
			level.Error(logger).Log("msg", "get trb balance", "addr", account.Address.Hex()[:8], "err", err)
		} else {
			metaData += fmt.Sprintf("TRB:%g ", math_t.BigIntToFloatDiv(trbBalance, params.Ether))
		}
		ethBalance, err := client.BalanceAt(ctx, account.Address, nil)
		if err != nil {
			level.Error(logger).Log("msg", "get eth balance", "addr", account.Address.Hex()[:8], "err", err)
		} else {
			metaData += fmt.Sprintf("ETH:%g", math_t.BigIntToFloatDiv(ethBalance, params.Ether))
		}

		accOpts += `<option value="` + string(account.Address.Hex()) + `">` + string(account.Address.Hex()[:8]) + ` ` + metaData + `</option>`

	}
	postForm := `
			<fieldset>
    		<legend>Confirm dispute details:</legend>
				<form id="dispute" method="post">
					<input type="hidden" id="action" name="action" value="dispute" >
					<label for="pass">EnvFile Pass:</label><input type="password" name="pass" id="pass"/><br/>
					<label for="account">Account:</label>
						<select name="account" id="account">` + accOpts + `</select><br/>
					<label for="queryID">Query Hash:</label>
						<input type="text" readonly="readonly" id="queryID" name="queryID" value="` + values.Get("queryID") + `" ><br/>
					<label for="ts">Timestamp:</label>
						<input type="text" readonly="readonly" id="ts" name="ts" value="` + values.Get("ts") + `" ><br/>
					<input type="submit" value="BEGIN DISPUTE">
				</form>
			</fieldset>
			`

	if values.Get("queryID") == "" {
		postForm = `missing Query Hash`
	}
	if values.Get("ts") == "" {
		postForm = `missing timestamp`
	}

	return postForm
}

func createDispute(
	ctx context.Context,
	logger log.Logger,
	r *http.Request,
	client ethereum.EthClient,
	contract contracts.TellorGovernCaller,
	envVars map[string]string,
) (*types.Transaction, error) {

	ctx, cncl := context.WithTimeout(ctx, 20*time.Second)
	defer cncl()
	account, err := ethereum.GetAccountByPubAddress(logger, r.PostForm.Get("account"), envVars)
	if err != nil {
		return nil, errors.Wrap(err, "getting account from selection")
	}

	_queryID := r.PostForm.Get("queryID")
	var queryID [32]byte
	copy(queryID[:], common.HexToHash(_queryID).Bytes())
	ts, err := strconv.ParseInt(r.PostForm.Get("ts"), 0, 64)
	if err != nil {
		return nil, errors.Wrap(err, "parsing the ts value")
	}

	opts, err := ethereum.PrepareTxOpts(ctx, client, account, 10, contracts.DisputeNewGasLimit)
	if err != nil {
		return nil, errors.Wrap(err, "preparing dispute TX")
	}
	tx, err := contract.BeginDispute(opts, queryID, big.NewInt(ts))
	if err != nil {
		return nil, errors.Wrap(err, "creating dispute TX")
	}
	return tx, nil
}

func loadEnvVars(client ethereum.EthClient, envFilePath string, pass string) (map[string]string, error) {
	envFileData, err := ioutil.ReadFile(envFilePath)
	if err != nil {
		return nil, errors.Wrapf(err, "opening the env file:%v", envFilePath)
	}

	if util.IsText(envFileData) {
		if client.NetworkID() != 4 && client.NetworkID() != 5 {
			return nil, errors.Errorf("env file is not encrypted. this is not allowed on the current network:%v", client.NetworkID())
		}
	} else {
		envFileData, err = private_file.Decrypt(envFileData, pass)
		if err != nil {
			return nil, errors.New("incorrect password")
		}
	}

	return private_file.SetEnvVars(envFileData)
}

func PSRs(
	ctx context.Context,
	logger log.Logger,
	psr *psr_tellor.Psr,
) http.HandlerFunc {
	t := template.New("template")

	type val struct {
		ID    string
		Name  string
		Data  string
		Value float64
		Error string
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var vals []val
		for queryID, psrI := range psr_tellor.Psrs {
			if psrI.Inactive {
				continue
			}
			v, err := psr.GetValue(queryID, time.Now())
			val := val{
				ID:    common.Bytes2Hex(queryID[:]),
				Name:  psrI.Pair + "-" + psrI.Aggr,
				Value: v / psrI.Granularity,
			}
			if err != nil {
				val.Error = err.Error()
			}
			id, err := psr_tellor.QueryIDToInt(queryID)
			if err != nil {
				val.Error = err.Error()
			}
			val.Data = common.Bytes2Hex(psr_tellor.NewQueryData(id))

			vals = append(vals, val)
		}

		sort.Slice(vals, func(i, j int) bool {
			return vals[i].ID < vals[j].ID
		})

		t, err := t.Parse(`
		<!DOCTYPE html>
		<html lang='en'>
		<head>
			<meta name='viewport' content='width=device-width, initial-scale=1, maximum-scale=1'>
			<title>Telliot PSRS</title>
			<style>
				body {
					font-family: arial;
					position:absolute;
				}
				td, tr {
					padding: 0.2em;
					margin: 0;
				}
			</style>
		</head>
		<body>

		<table>
			<tr>
				<td>QueryID</td><td>QueryData</td><td>Pair Name</td><td>Value</td><td></td>
			</tr>
		{{range $index, $val := .}}
			<tr>
			<td>{{ $val.ID }}</td>
			<td>{{ $val.Data }}</td>
			<td>{{ $val.Name }}</td>
			<td>{{ $val.Value }}</td>
				<td>{{ $val.Error }}</td>
			</tr>
		{{end}}
		</table>
		</body>
		</html>
		`)

		if err != nil {
			http.Error(w, fmt.Sprintf("parsing html template submit logs:%v", err), http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, vals)
		if err != nil {
			http.Error(w, fmt.Sprintf("executing the html template:%v", err), http.StatusInternalServerError)
			return
		}
	})
}

func (self *Web) Start() error {
	level.Info(self.logger).Log("msg", "starting", "logLevel", self.cfg.LogLevel, "addr", self.srv.Addr)
	if err := self.srv.ListenAndServe(); err != http.ErrServerClosed {
		return errors.Wrapf(err, "ListenAndServe")
	}
	return nil
}

func (self *Web) Stop() {
	self.stop()
	if err := self.srv.Close(); err != nil {
		level.Error(self.logger).Log("msg", "closing srv", "err", err)
	}
}

func serveDebug(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	subpath := route.Param(ctx, "subpath")

	if subpath == "/pprof" {
		http.Redirect(w, req, req.URL.Path+"/", http.StatusMovedPermanently)
		return
	}

	if !strings.HasPrefix(subpath, "/pprof/") {
		http.NotFound(w, req)
		return
	}
	subpath = strings.TrimPrefix(subpath, "/pprof/")

	switch subpath {
	case "cmdline":
		pprof.Cmdline(w, req)
	case "profile":
		pprof.Profile(w, req)
	case "symbol":
		pprof.Symbol(w, req)
	case "trace":
		pprof.Trace(w, req)
	default:
		req.URL.Path = "/debug/pprof/" + subpath
		pprof.Index(w, req)
	}
}
