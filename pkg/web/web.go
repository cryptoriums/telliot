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

	"github.com/cryptoriums/telliot/pkg/componentor"
	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/db"
	"github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/format"
	"github.com/cryptoriums/telliot/pkg/logging"
	"github.com/cryptoriums/telliot/pkg/math"
	"github.com/cryptoriums/telliot/pkg/private_file"
	"github.com/cryptoriums/telliot/pkg/psr/tellor"
	psr_tellor "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/cryptoriums/telliot/pkg/web/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/route"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/storage"
	"golang.org/x/tools/godoc/util"
)

const ComponentName = "web"

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
	componentor componentor.Componentor,
	client ethereum.EthClient,
	contract contracts.TellorCaller,
	envFilePath string,
) http.HandlerFunc {
	var err error
	t := template.New("template").Funcs(template.FuncMap{
		"timeSince": func(ts int64) int {
			return int(time.Since(time.Unix(ts, 0)).Minutes())
		},
		"tsToTime": func(ts int64) string {
			return time.Unix(ts, 0).UTC().Format("15:04")
		},
		"psrDetails": func(id *big.Int) string {
			return strings.Replace(tellor.Psrs[id.Int64()].Pair+"_"+tellor.Psrs[id.Int64()].Aggr, " ", "_", -1)
		},

		"isInactive": func(id *big.Int) string {
			inactive := ""
			if psr_tellor.IsInactive(id.Int64()) {
				inactive = "(in)"
			}
			return id.String() + inactive
		},

		"applyGranularity": func(val *big.Int) string {
			return fmt.Sprintf("%.6f", float64(val.Int64())/tellor.DefaultGranularity)
		},
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.RequestURI)
		values := u.Query()

		postResult := ""
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, fmt.Sprintf("parsing form data:%v", err), http.StatusInternalServerError)
				return
			}
			if err := checkPass(client, envFilePath, r.PostForm.Get("pass")); err != nil {
				http.Error(w, fmt.Sprintf("checking  password:%v", err), http.StatusInternalServerError)
				return
			}
			switch r.PostForm.Get("action") {
			case "dispute":
				tx, err := createDispute(ctx, logger, r, client, contract)
				if err != nil {
					http.Error(w, fmt.Sprintf("creating a dispute:%v", err), http.StatusInternalServerError)
					return
				}
				postResult = `created new dispute<br/><a href="` + ethereum.GetEtherscanURL(client.NetworkID()) + `/tx/` + tx.Hash().String() + `">` + tx.Hash().String() + `</a><br/><br/>`
			case "toggleStatus":
				componentor.ToggleStatus()
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
		ctx, cncl := context.WithTimeout(ctx, time.Minute)
		defer cncl()
		submits, err := contracts.GetSubmitLogs(
			ctx,
			client,
			contract,
			0,
			lookBack,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("getting submit logs:%v", err), http.StatusInternalServerError)
			return
		}

		sort.Slice(submits[:], func(i, j int) bool {
			return submits[i].Timestamp.Int64() > submits[j].Timestamp.Int64()
		})

		postForm := ``
		if values.Get("reporterIndex") != "" && postResult == "" {
			postForm, err = prepareDisputeForm(ctx, logger, client, contract, values)
			if err != nil {
				http.Error(w, fmt.Sprintf("prepare post form:%v", err), http.StatusInternalServerError)
				return
			}
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
				input:read-only {
					background-color: grey;
				  }
			</style>
		</head>
		<body>
		<h2>` + componentor.ID() + ` status: ` + formatStatus(componentor.Status()) + `
			<form style="display:inline" id="toggleStatus" method="post">
				<input type="hidden" id="action" name="action" value="toggleStatus" >
				<input placeholder="EnvFile Pass" type="password" name="pass" id="pass"/>
				<input type="submit" value="Toggle Status">
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
			{{range $di,$dv := .DataIDs}}
				{{ if eq $di 0 }}
					<tr>
						<td colspan="100%"><b>Mins:{{  timeSince $submit.Timestamp.Int64 }} Time:{{ tsToTime $submit.Timestamp.Int64 }} Timestamp:{{ $submit.Timestamp.String}}</b> Tx:<a  href="` + ethereum.GetEtherscanURL(client.NetworkID()) + `/tx/{{ $submit.TxHash.Hex }}">{{ $submit.TxHash.Hex }}</a> </td>
					</tr>
				{{end}}
				<tr>
					<td>
						{{ psrDetails $dv}}
					</td>
					<td>
						id:{{isInactive $dv}}
					</td>
					{{ range $vi,$vv := (index $submit.Values .Int64)}}
						<td style="text-align:right">{{ applyGranularity . }}</td>
						<td style="text-align:right">{{ slice (index (index $submit.Reporters $dv.Int64) $vi).Hex 0 8 }}</td>
						<td>
							<form id="data">
								<input type="hidden" id="dataID" name="dataID" value="{{$dv}}" >
								<input type="hidden" id="ts" name="ts"  value="{{$submit.Timestamp.String}}" >
								<input type="hidden" id="reporterIndex" name="reporterIndex"  value="{{$vi}}" >
								<input type="hidden" id="look-back" name="look-back"  value="` + strconv.Itoa(int(lookBack.Hours())) + `h" >
								<input type="submit" value="GO">
							</form>
						</td>
					{{end}}
				</tr>
			{{end}}
			<tr>
				<td colspan="100%"><br/></td>
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
		return `<span style="color:grey">disabled</span>`
	default:
		return `<span style="color:red">enabled</span>`
	}
}

func prepareDisputeForm(ctx context.Context, logger log.Logger, client ethereum.EthClient, contract contracts.TellorCaller, values url.Values) (string, error) {
	accOpts := ""
	accounts, err := ethereum.GetAccounts(logger)
	if err != nil {
		return "", errors.Wrap(err, "getting accounts")
	}
	for _, account := range accounts {
		metaData := ""

		status, _, err := contract.GetStakerInfo(&bind.CallOpts{Context: ctx}, account.Address)
		if err != nil {
			level.Error(logger).Log("msg", "get stake status", "addr", account.Address.Hex()[:8], "err", err)
		} else {
			metaData += contracts.ReporterStatusName(status.Int64()) + " "
		}
		trbBalance, err := contract.BalanceOf(&bind.CallOpts{Context: ctx}, account.Address)
		if err != nil {
			level.Error(logger).Log("msg", "get trb balance", "addr", account.Address.Hex()[:8], "err", err)
		} else {
			metaData += fmt.Sprintf("TRB:%g ", math.BigIntToFloatDiv(trbBalance, params.Ether))
		}
		ethBalance, err := client.BalanceAt(ctx, account.Address, nil)
		if err != nil {
			level.Error(logger).Log("msg", "get eth balance", "addr", account.Address.Hex()[:8], "err", err)
		} else {
			metaData += fmt.Sprintf("ETH:%g", math.BigIntToFloatDiv(ethBalance, params.Ether))
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
					<label for="dataID">Data ID:</label>
						<input type="text" readonly="readonly" id="dataID" name="dataID" value="` + values.Get("dataID") + `" ><br/>
					<label for="ts">Timestamp:</label>
						<input type="text" readonly="readonly" id="ts" name="ts" value="` + values.Get("ts") + `" ><br/>
					<label for="reporterIndex">Reporter Index:</label>
						<input type="text" readonly="readonly" id="reporterIndex" name="reporterIndex"  value="` + values.Get("reporterIndex") + `" ><br/>
					<input type="submit" value="BEGIN DISPUTE">
				</form>
			</fieldset>
			`

	if values.Get("dataID") == "" {
		postForm = `missing DATA ID`
	}
	if values.Get("ts") == "" {
		postForm = `missing timestamp`
	}
	if values.Get("reporterIndex") == "" {
		postForm = `missing reporterIndex`
	}
	return postForm, nil
}

func createDispute(
	ctx context.Context,
	logger log.Logger,
	r *http.Request,
	client ethereum.EthClient,
	contract contracts.TellorCaller,
) (*types.Transaction, error) {

	ctx, cncl := context.WithTimeout(ctx, 20*time.Second)
	defer cncl()
	account, err := ethereum.GetAccountByPubAddress(logger, r.PostForm.Get("account"))
	if err != nil {
		return nil, errors.Wrap(err, "getting account from selection")
	}
	dataID, err := strconv.ParseInt(r.PostForm.Get("dataID"), 0, 64)
	if err != nil {
		return nil, errors.Wrap(err, "parsing the dataID value")
	}
	ts, err := strconv.ParseInt(r.PostForm.Get("ts"), 0, 64)
	if err != nil {
		return nil, errors.Wrap(err, "parsing the ts value")
	}
	reporterIndex, err := strconv.ParseInt(r.PostForm.Get("reporterIndex"), 0, 64)
	if err != nil {
		return nil, errors.Wrap(err, "parsing the reporterIndex value")
	}
	opts, err := ethereum.PrepareTx(ctx, client, account, 10, contracts.NewDisputeGasLimit)
	if err != nil {
		return nil, errors.Wrap(err, "preparing dispute TX")
	}
	tx, err := contract.BeginDispute(opts, big.NewInt(dataID), big.NewInt(ts), big.NewInt(reporterIndex))
	if err != nil {
		return nil, errors.Wrap(err, "creating dispute TX")
	}
	return tx, nil
}

func checkPass(client ethereum.EthClient, envFilePath string, pass string) error {
	env, err := ioutil.ReadFile(envFilePath)
	if err != nil {
		return errors.Wrap(err, "reading env file")
	}

	if !util.IsText(env) {
		_, err = private_file.Decrypt(env, pass)
		if err != nil {
			return errors.New("incorrect password")
		}
	} else {
		switch client.NetworkID() {
		case 4:
		case 5:
		default:
			return errors.New("file is not encrypted - action is forbidden via the web with unencrypted env file for the current network")
		}
	}
	return nil
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
