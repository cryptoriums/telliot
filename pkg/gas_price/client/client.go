// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package client

import (
	"context"

	"github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/gas_price"
	"github.com/cryptoriums/telliot/pkg/math"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
)

const ComponentName = "gasPriceClient"

type Client struct {
	client ethereum.EthClient
	logger log.Logger
}

func New(ctx context.Context, logger log.Logger, client ethereum.EthClient) (gas_price.GasPriceQuerier, error) {
	return &Client{
		client: client,
		logger: log.With(logger, "component", ComponentName),
	}, nil

}

func (self *Client) Name() string {
	return ComponentName
}

func (self *Client) Query(ctx context.Context, confidence int) (baseFeeGwei, tipFeeGwei float64, err error) {
	header, err := self.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, 0, err
	}
	baseFeeGwei = math.BigIntToFloatDiv(header.BaseFee, params.GWei)

	tipFee, err := self.client.SuggestGasTipCap(ctx)
	if err != nil {
		return 0, 0, err
	}
	tipFeeGwei = math.BigIntToFloatDiv(tipFee, params.GWei)
	return
}
