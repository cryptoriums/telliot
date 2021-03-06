// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package config

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/cryptoriums/packages/client"
	"github.com/cryptoriums/packages/format"
	"github.com/cryptoriums/packages/private_file"
	"github.com/cryptoriums/telliot/pkg/aggregator"
	"github.com/cryptoriums/telliot/pkg/db"
	psr_tellor "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/cryptoriums/telliot/pkg/tracker/blocks"
	"github.com/cryptoriums/telliot/pkg/tracker/index"
	"github.com/cryptoriums/telliot/pkg/web"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

type Transactor struct {
	Transact bool
}

// Config is the top-level configuration that holds configs for all components.
type Config struct {
	Transactor    Transactor
	Web           web.Config
	TrackerIndex  index.Config
	TrackerBlocks blocks.Config
	Aggregator    aggregator.Config
	PsrTellor     psr_tellor.Config
	Db            db.Config
	// EnvFile location that include all private details like private key etc.
	EnvFile string `json:"envFile"`
	// Is is set by expanding the os env vars.
	// Not using os.SetEnv so that it doesn't expose those to someone that has access to the machine.
	// If using os.SetEnv for the private key anyone can see it that has access to the machine.
	EnvVars   map[string]string
	EthClient client.Config
}

var DefaultConfig = Config{
	Web: web.Config{
		ListenHost: "", // Listen on all addresses.
		ListenPort: 9090,
	},
	Db: db.Config{
		RemotePort:    9090,
		Path:          "db",
		RemoteTimeout: format.Duration{Duration: 5 * time.Second},
		Retention:     format.Duration{Duration: 60 * 24 * time.Hour}, // 60 days.
	},
	PsrTellor: psr_tellor.Config{
		MinConfidenceDefault: 80,
		MinConfidencePerSymbol: map[string]float64{
			"USPCE": 100,
		},
	},
	TrackerIndex: index.Config{
		Interval:  format.Duration{Duration: time.Minute},
		IndexFile: "configs/index.json",
	},
	EnvFile: "configs/.env",
}

func LoadConfig(ctx context.Context, logger log.Logger, path string, strictParsing bool) (*Config, error) {
	cfg, err := Parse(logger, string(path), strictParsing)
	if err != nil {
		return nil, errors.Wrap(err, "creating config")
	}

	if err := Validate(cfg); err != nil {
		return nil, errors.Wrap(err, "validate config")
	}

	cfg.EnvVars, err = private_file.LoadEnvFile(ctx, logger, cfg.EnvFile, cfg.Transactor.Transact, cfg.Web.ListenHost, cfg.Web.ListenPort)
	if err != nil {
		return nil, errors.Wrapf(err, "loading the enf file:%v", cfg.EnvFile)
	}

	return cfg, nil
}

func Validate(cfg *Config) error {
MainLoop:
	for pair := range cfg.PsrTellor.MinConfidencePerSymbol {
		for _, psr := range psr_tellor.Psrs {
			if pair == psr.Pair {
				continue MainLoop
			}
		}
		return errors.Errorf("confidence level for invalid psr pair:%v", pair)
	}
	return nil
}

func Parse(logger log.Logger, path string, strictParsing bool) (*Config, error) {

	cfg := &Config{}

	cfgI, err := De??pCopy(logger, path, cfg, DefaultConfig, strictParsing)
	if err != nil {
		return nil, errors.Wrap(err, "making a config deep copy")
	}
	cfg = cfgI.(*Config)

	return cfg, err
}

func De??pCopy(logger log.Logger, path string, cfg, cfgDefault interface{}, strictParsing bool) (interface{}, error) {
	var file *os.File
	var fileLoaded bool
	var err error
	if path == "" { // Config flag not provided so try to load the default location.
		path = filepath.Join("configs", "config.json")
		file, err = os.Open(path)
		if err != nil && !os.IsNotExist(err) {
			return nil, errors.Wrap(err, "open config file")
		}
		if err == nil {
			fileLoaded = true
		}
	} else { // When config file location is provided should error when the file doesn't exist.
		// In this case obviously the user want to use some custom file.
		file, err = os.Open(path)
		if err != nil {
			return nil, errors.Wrap(err, "open config file")
		}
		fileLoaded = true
	}

	// Populate the defaults.
	{
		b, err := json.Marshal(cfgDefault)
		if err != nil {
			return nil, errors.Wrap(err, "marshal default config")
		}

		if err := json.Unmarshal(b, cfg); err != nil {
			return nil, errors.Wrap(err, "copy default config")
		}
	}

	//  When a config file is loaded also overwrite the defaults with the values from the config file.
	if fileLoaded {
		dec := json.NewDecoder(file)
		if strictParsing {
			dec.DisallowUnknownFields()
		}
		for {
			// Override defaults with the custom configs.
			if err := dec.Decode(cfg); err == io.EOF {
				break
			} else if err != nil {
				return nil, errors.Wrap(err, "parse config")
			}

		}
	} else {
		level.Warn(logger).Log("msg", "no config file loaded and using the defaults")
	}

	return cfg, nil
}
