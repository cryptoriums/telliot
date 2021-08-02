// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package db

import (
	"context"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/cryptoriums/telliot/pkg/format"
	"github.com/pkg/errors"
	promConfig "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/pkg/timestamp"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/storage/remote"
	"github.com/prometheus/prometheus/tsdb"
)

const ComponentName = "db"

type Config struct {
	LogLevel string
	Path     string
	// Connect to this remote DB.
	RemoteHost    string
	RemotePort    uint
	RemoteTimeout format.Duration
}

func NewRemoteDB(cfg Config) (storage.SampleAndChunkQueryable, error) {

	url, err := url.Parse("http://" + cfg.RemoteHost + ":" + strconv.Itoa(int(cfg.RemotePort)) + "/api/v1/read")
	if err != nil {
		return nil, err
	}
	client, err := remote.NewReadClient("", &remote.ClientConfig{
		URL:     &promConfig.URL{URL: url},
		Timeout: model.Duration(cfg.RemoteTimeout.Duration),
		HTTPClientConfig: promConfig.HTTPClientConfig{
			FollowRedirects: true,
		},
	})
	if err != nil {
		return nil, err
	}
	return remote.NewSampleAndChunkQueryableClient(
		client,
		labels.Labels{},
		[]*labels.Matcher{},
		true,
		func() (i int64, err error) { return 0, nil },
	), nil
}

func Add(ctx context.Context, tsdb *tsdb.DB, lbls labels.Labels, value float64) error {
	var err error
	appender := tsdb.Appender(ctx)

	defer func() { // An appender always needs to be committed or rolled back.
		if err != nil {
			if errR := appender.Rollback(); errR != nil {
				err = errors.Wrap(err, "db rollback failed")
			}
			return
		}
		if errC := appender.Commit(); errC != nil {
			err = errors.Wrap(err, "db append commit failed")
		}
	}()
	sort.Sort(lbls) // This is important! The labels need to be sorted to avoid creating the same series with duplicate reference.

	ts := timestamp.FromTime(time.Now())
	_, err = appender.Append(0, lbls, ts, float64(value))
	if err != nil {
		return errors.Wrapf(err, "append values to the DB ts:%v val:%v", ts, float64(value))
	}
	return nil
}
