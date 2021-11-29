// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package gas_usage

import (
	"context"
	"math/big"

	"github.com/cryptoriums/packages/ethereum"
)

type GasUsageQuerier interface {
	Query(ctx context.Context, account *ethereum.Account, queryID [32]byte, val []byte, nonce *big.Int, queryData []byte) (float64, error)
	Name() string
}

// Default returns hard coded gas usage.
// In theory can use the eth client to esimtate it, but this requires a valid nonce and a staked addr so not that easy.
// Also the client estimation doesn't account for refunds due to freeing up storage slots so
// this needs to be tracked separately and further complicates the logic.
type Default struct {
}

func NewDefault() *Default {
	return &Default{}
}

func (self *Default) Name() string {
	return "default"
}

func (self *Default) Query(ctx context.Context, _ *ethereum.Account, _ [32]byte, _ []byte, _ *big.Int, _ []byte) (float64, error) {
	return 300_000, nil
}
