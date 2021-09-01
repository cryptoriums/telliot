// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package gas_price

import "context"

type GasPriceQuerier interface {
	// Query returns the suggested base and tip fee in GWEI.
	Query(ctx context.Context, confidence int) (baseFeeGwei, tipFeeGwei float64, err error)
}
