// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package reward

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/tellor-io/telliot/pkg/testutil"
)

func TestGasPriceForProfitThreshold(t *testing.T) {
	ctx := context.Background()
	mock := NewMockContractCaller()
	reward, err := New(log.NewNopLogger(), ctx, Config{LogLevel: "debug"}, mock, mock)
	testutil.Ok(t, err)

	for _, tcase := range []struct {
		Name             string
		CurrentReward    *big.Int
		ProfitTarget     float64
		GasPriceExpected int64
		GasUsage         *big.Int
	}{
		{
			"50% profit and high gas usage",
			big.NewInt(2 * params.Ether),
			50,
			1,
			big.NewInt(params.Ether),
		},
		{
			"50% profit and usual gas usage",
			big.NewInt(2 * params.Ether),
			50,
			params.GWei,
			big.NewInt(params.GWei),
		},
		{
			"50% profit, usual gas usage and very small reward",
			big.NewInt(params.Ether),
			50,
			params.GWei / 2,
			big.NewInt(params.GWei),
		},
		{
			"50% profit, usual gas usage and smaller reward",
			big.NewInt(params.Ether / 2),
			50,
			params.GWei / 4,
			big.NewInt(params.GWei),
		},
	} {
		t.Run(tcase.Name, func(t *testing.T) {
			mock.SetReward(tcase.CurrentReward)
			gasPrice, err := reward.GasPriceForProfitThreshold(ctx, 4, tcase.ProfitTarget, tcase.GasUsage)
			testutil.Ok(t, err)
			testutil.Equals(t, tcase.GasPriceExpected, gasPrice)
		})
	}
}

type MockContractCaller struct {
	currentReward *big.Int
}

func (self *MockContractCaller) SetReward(r *big.Int) {
	self.currentReward = r
}
func (self *MockContractCaller) TimeWeightedAvg(symbol string, start time.Time, lookBack time.Duration) (float64, float64, error) {
	// 1 trb equals 1 eth with 100% confidence.
	return 1, 100, nil
}

func (self *MockContractCaller) CurrentReward(opts *bind.CallOpts) (*big.Int, error) {
	return self.currentReward, nil
}

func NewMockContractCaller() *MockContractCaller {
	return &MockContractCaller{}
}

func (self *MockContractCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	return nil, nil
}
func (self *MockContractCaller) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return nil, nil
}
func (self *MockContractCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	return nil, nil, nil

}
func (self *MockContractCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return struct {
		Challenge  [32]byte
		RequestIds [5]*big.Int
		Difficutly *big.Int
		Tip        *big.Int
	}{[32]byte{}, [5]*big.Int{}, nil, nil}, nil
}
func (self *MockContractCaller) GetAddr() *common.Address {
	return nil
}
