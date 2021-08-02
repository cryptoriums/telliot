// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package gas_price

import (
	"context"
	"math/big"
)

type GasPriceQuerier interface {
	Query(ctx context.Context) (*big.Int, error)
}
