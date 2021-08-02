// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package web

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func Get(ctx context.Context, url string, headers map[string]string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives:  true,
		DisableCompression: true,
	}
	client := http.Client{Transport: tr}

	ctx, cncl := context.WithTimeout(ctx, 10*time.Second)
	defer cncl()

	req, err := http.NewRequestWithContext(ctx, "GET", ExpandTimeVars(url), nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	r, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "fetching data")
	}
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}

	if r.StatusCode/100 != 2 {
		return nil, errors.Errorf("response status code not OK code:%v, payload:%v", r.StatusCode, string(data))
	}
	return data, nil
}

func ExpandTimeVars(url string) string {
	nowMillisecons := strconv.Itoa(int(time.Now().Unix() * 1000))
	url = strings.Replace(url, "$NOW", nowMillisecons, -1)

	millsIn1day := 86400000
	eodMillisecons := strconv.Itoa(int(time.Now().Unix()*1000) - millsIn1day)
	url = strings.Replace(url, "$EOD", eodMillisecons, -1)

	return url
}
