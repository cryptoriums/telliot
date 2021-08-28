// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package web

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"strings"

	"github.com/cryptoriums/telliot/pkg/db"
	"github.com/cryptoriums/telliot/pkg/format"
	"github.com/cryptoriums/telliot/pkg/logging"
	"github.com/cryptoriums/telliot/pkg/web/api"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/route"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/storage"
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

func New(ctx context.Context, logger log.Logger, tsDB storage.SampleAndChunkQueryable, cfg Config) (*Web, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	router := route.New()

	router.Get("/debug/*subpath", serveDebug)
	router.Post("/debug/*subpath", serveDebug)

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
