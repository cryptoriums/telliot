// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package format

import (
	"testing"
	"time"

	"github.com/cryptoriums/packages/testutil"
)

func Test_ExpandTimeVars(t *testing.T) {

	type testcase struct {
		expectedEOD string
		expectedBOD string
	}

	cases := []testcase{
		{
			"2000-01-02 00:00:00 +0000 UTC",
			"2000-01-01 00:00:00 +0000 UTC",
		},
		{
			"2000-01-02 00:00:00 +0000 UTC",
			"2000-01-01 00:00:00 +0000 UTC",
		},
		{
			"2000-01-02 00:00:00 +0000 UTC",
			"2000-01-01 00:00:00 +0000 UTC",
		},
	}

	for _, tc := range cases {

		ts := EOD()
		testutil.Equals(t, tc.expectedEOD, time.Unix(int64(ts), 0).UTC().String())

		ts = BOD()
		testutil.Equals(t, tc.expectedBOD, time.Unix(int64(ts), 0).UTC().String())

	}
}
