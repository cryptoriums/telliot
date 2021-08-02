// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package gas_estimator

import (
	"context"

	"github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/pkg/errors"
)

type GasEstimator interface {
	EstimateGas(ctx context.Context, account *ethereum.Account, slot uint64) (uint64, error)
}

// Default returns hard coded gas usage for all slots.
// In theory can use the eth client to esimtate it, but this requires a valid nonce and a staked addr so not that easy.
// Also the client estimation doesn't account for refunds due to freeing up storage slots so
// this needs to be tracked separately and further complicates the logic.
type Default struct {
}

func NewDefault() *Default {
	return &Default{}
}

func (self *Default) EstimateGas(ctx context.Context, _ *ethereum.Account, slot uint64) (uint64, error) {
	switch slot {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		return 400_000, nil
	case 4:
		return 2_100_000, nil
	}
	return 0, errors.Errorf("invalid slot:%v", slot)
}
