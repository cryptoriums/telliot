// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package gas_station

import (
	"context"

	"github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/math"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/log"
)

const ComponentName = "gasPriceGasStation"

type GasStation struct {
	netID  int64
	ctx    context.Context
	client ethereum.EthClient
	logger log.Logger
}

// GasStation is what ETHGasStation returns from queries. Not all fields are filled in.
type GasStationModel struct {
	Fast    float32 `json:"fast"`
	Fastest float32 `json:"fastest"`
	Average float32 `json:"average"`
}

func New(ctx context.Context, logger log.Logger, client ethereum.EthClient) (*GasStation, error) {
	return &GasStation{
		netID:  client.NetworkID(),
		ctx:    ctx,
		client: client,
		logger: log.With(logger, "component", ComponentName),
	}, nil

}

func (self *GasStation) Name() string {
	return ComponentName
}

// Query implements the gas query interface.
// For now gas station doesn't have 1559 api so just returning the client suggested values.
func (self *GasStation) Query(ctx context.Context, confidence int) (baseFeeGwei, tipFeeGwei float64, err error) {
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

	// ctx, cncl := context.WithTimeout(self.ctx, 5*time.Second)
	// defer cncl()
	// if self.netID != 1 {
	// 	gasPrice, err := self.client.SuggestGasPrice(ctx)
	// 	if err != nil {
	// 		return nil, errors.Wrap(err, "getting suggested gas price")
	// 	}
	// 	return gasPrice, nil
	// }

	// defer func() {
	// 	if errFinal != nil {
	// 		level.Error(self.logger).Log("msg", "fetching eth gas price falling back to client suggested price", "err", errFinal)
	// 		gasPrice, err := self.client.SuggestGasPrice(ctx)
	// 		if err != nil {
	// 			errFinal = errors.Wrapf(errFinal, "failed to get price from chain client:%v", err)
	// 			return
	// 		}
	// 		gasPriceFinal = gasPrice
	// 	}
	// }()

	// resp, err := web.Get(ctx, "https://ethgasstation.info/json/ethgasAPI.json", nil)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "fetch price from provider")
	// }

	// gpModel := GasStationModel{}
	// err = json.Unmarshal(resp, &gpModel)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "provider response json unmarshal")
	// }

	// var gasPrice float32
	// switch conf := confidence; {
	// case conf < 60:
	// 	gasPrice = gpModel.Average
	// case conf < 80:
	// 	gasPrice = gpModel.Fast
	// case conf < 100:
	// 	gasPrice = gpModel.Fastest
	// }

	// gasPriceB := big.NewInt(int64(gasPrice / 10))
	// return big.NewInt(0).Mul(gasPriceB, big.NewInt(params.GWei)), nil
}
