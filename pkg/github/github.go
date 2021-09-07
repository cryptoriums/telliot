// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package github

import (
	"context"
	"io/ioutil"
	"strings"
	"time"

	"github.com/google/go-github/v35/github"
	"github.com/pkg/errors"
)

func CheckNewVersion(owner, repo string, current string) (string, error) {
	if current == "" { // Can be empty when using `go run`.
		return "", nil
	}
	ctx, cncl := context.WithTimeout(context.Background(), 3*time.Second)
	defer cncl()
	client := github.NewClient(nil)
	release, resp, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return "", errors.Wrap(err, "checking for a new release")
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		rbody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", errors.Errorf("bad response status %v", resp.Status)
		}
		return "", errors.Errorf("bad response status %v from %q", resp.Status, string(rbody))
	}

	parts := strings.Split(current, "-")

	if len(parts) == 0 {
		return "", errors.New("failed to process current tag name")
	}

	if parts[0] != release.GetTagName() {
		return release.GetHTMLURL(), nil
	}

	return "", nil
}
