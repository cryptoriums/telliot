// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package helpers

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
		MaxIdleConns:       50,
		IdleConnTimeout:    30 * time.Second,
	}
	client := http.Client{Transport: tr}
	defer client.CloseIdleConnections()

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

var (
	now = time.Now // Set as a var to allowe overriging in the tests.
)

func ExpandTimeVars(url string) string {
	year, month, day := now().Date()
	eod := time.Date(year, month, day, 0, 0, 0, 0, now().Location())

	secsInAday := int64(86400)

	bod := eod.Unix() - secsInAday

	// Need to be first so that the longer string substitution happens first.
	url = strings.Replace(url, "$EODM", strconv.Itoa(int(eod.Unix()*1000)), -1)
	url = strings.Replace(url, "$BODM", strconv.Itoa(int(bod*1000)), -1)

	url = strings.Replace(url, "$EOD", strconv.Itoa(int(eod.Unix())), -1)
	url = strings.Replace(url, "$BOD", strconv.Itoa(int(bod)), -1)

	return url
}
