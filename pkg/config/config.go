// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package config

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cryptoriums/telliot/pkg/aggregator"
	"github.com/cryptoriums/telliot/pkg/db"
	"github.com/cryptoriums/telliot/pkg/format"
	"github.com/cryptoriums/telliot/pkg/private_file"
	psr_tellor "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/cryptoriums/telliot/pkg/tracker/blocks"
	"github.com/cryptoriums/telliot/pkg/tracker/index"
	"github.com/cryptoriums/telliot/pkg/web"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"golang.org/x/tools/godoc/util"
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

	err = LoadEnvFile(ctx, logger, cfg)
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

	cfgI, err := DeеpCopy(logger, path, cfg, DefaultConfig, strictParsing)
	if err != nil {
		return nil, errors.Wrap(err, "making a config deep copy")
	}
	cfg = cfgI.(*Config)

	return cfg, err
}

func LoadEnvFile(ctx context.Context, logger log.Logger, cfg *Config) error {
	_env, err := os.Open(cfg.EnvFile)
	if err != nil {
		// When running inside k8s the .env file is converted into secrets
		// so the .env file in this case doesn't exist.
		if os.IsNotExist(err) {
			level.Warn(logger).Log("msg", ".env file doesn't exist so using the existing env variables", "path", cfg.EnvFile)
			return nil
		}
		return errors.Wrapf(err, "opening the env file:%v", cfg.EnvFile)
	}

	env, err := ioutil.ReadAll(_env)
	if err != nil {
		return errors.Wrapf(err, "reading the env file:%v", cfg.EnvFile)
	}

	if !util.IsText(env) {
		level.Info(logger).Log("msg", "env file is encrypted", "path", cfg.EnvFile)
		if os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
			transacting := `<span style="color:grey">disabled</span>`
			if cfg.Transactor.Transact {
				transacting = `<span style="color:red">enabled</span>`
			}

			level.Info(logger).Log("msg", "running inside k8s so will wait for web password decrypt input")
			env = private_file.DecryptWithWebPassword(ctx, logger, "<h2>Transacting is:"+transacting+"</h2>", env, cfg.Web.ListenHost, cfg.Web.ListenPort)
		} else {
			env, err = private_file.DecryptWithPasswordLoop(env)
			if err != nil {
				return errors.Wrap(err, "decrypt input file")
			}
		}
	}

	rr := bytes.NewReader(env)
	envMap, err := godotenv.Parse(rr)
	if err != nil {
		return errors.Wrapf(err, "parsing the env file:%v", cfg.EnvFile)
	}

	// Copied from the godotenv source code.
	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}

	for key, value := range envMap {
		if !currentEnv[key] {
			os.Setenv(key, value)
		}
	}

	return nil
}

func DeеpCopy(logger log.Logger, path string, cfg, cfgDefault interface{}, strictParsing bool) (interface{}, error) {
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
