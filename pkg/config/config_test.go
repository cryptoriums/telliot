// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package config

import (
	"os"
	"testing"

	"github.com/cryptoriums/packages/testutil"
)

func createEnvFile(t *testing.T) func() {
	f, err := os.Create(".env")
	testutil.Ok(t, err)

	_, err = f.WriteString("ETH_PRIVATE_KEY=\"0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef\"")
	testutil.Ok(t, err)
	_, err = f.WriteString("NODE_URLS=\"wss://mainnet.infura.io/v3/ws/xxxxxxxxxxxxx\"")
	testutil.Ok(t, err)
	testutil.Ok(t, f.Close())

	return func() {
		os.Remove(".env")
	}
}

func TestConfig(t *testing.T) {
	//Creating a mock .ENV file to go around this issue with godotenv:
	//https://github.com/joho/godotenv/issues/43
	cleanup := createEnvFile(t)
	defer t.Cleanup(cleanup)

	_, err := OpenTestConfig("../..")
	testutil.Ok(t, err)

}
