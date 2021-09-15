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
	"github.com/cryptoriums/telliot/pkg/mining"
	"github.com/cryptoriums/telliot/pkg/private_file"
	psrTellor "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/cryptoriums/telliot/pkg/submitter/tellor"
	"github.com/cryptoriums/telliot/pkg/tasker"
	"github.com/cryptoriums/telliot/pkg/tracker/index"
	transactorTellor "github.com/cryptoriums/telliot/pkg/transactor/tellor"
	"github.com/cryptoriums/telliot/pkg/web"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"golang.org/x/tools/godoc/util"
)

// Config is the top-level configuration that holds configs for all components.
type Config struct {
	Web              web.Config
	Mining           mining.Config
	SubmitterTellor  tellor.Config
	Tasker           tasker.Config
	TransactorTellor transactorTellor.Config
	TrackerIndex     index.Config
	Aggregator       aggregator.Config
	PsrTellor        psrTellor.Config
	Db               db.Config
	// EnvFile location that include all private details like private key etc.
	EnvFile string `json:"envFile"`
}

var DefaultConfig = Config{
	Mining: mining.Config{
		Heartbeat: time.Minute,
	},
	Web: web.Config{
		ListenHost: "", // Listen on all addresses.
		ListenPort: 9090,
	},
	Db: db.Config{
		RemotePort:    9090,
		Path:          "db",
		RemoteTimeout: format.Duration{Duration: 5 * time.Second},
	},
	TransactorTellor: transactorTellor.Config{
		GasMaxTipGwei: 10,
	},
	SubmitterTellor: tellor.Config{
		Enabled: true,
		// With a 30 second delay here as a workaround to prevent a race condition in the oracle contract check.
		MinSubmitPeriod: format.Duration{Duration: 15*time.Minute + 30*time.Second},
	},
	PsrTellor: psrTellor.Config{
		MinConfidence: 80,
	},
	Aggregator: aggregator.Config{
		ManualDataFile: "configs/manualData.json",
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

	err = LoadEnvFile(ctx, logger, cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "loading the enf file:%v", cfg.EnvFile)
	}
	return cfg, nil
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
			level.Info(logger).Log("msg", "running inside k8s so will wait for web password decrypt input")
			env = private_file.DecryptWithWebPassword(ctx, logger, env, cfg.Web.ListenHost, cfg.Web.ListenPort)
		} else {
			env = private_file.DecryptWithPasswordLoop(env)
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
