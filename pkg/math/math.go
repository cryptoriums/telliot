// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package math

import (
	"encoding/binary"
	"math"
	"math/big"
)

func PercentageDiff(old, new float64) float64 {
	diff := new - old

	if old == 0 {
		if new > 0 {
			return math.MaxFloat64
		}
		return -math.MaxFloat64
	}
	if new == 0 {
		if old > 0 {
			return math.MaxFloat64
		}
		return -math.MaxFloat64
	}

	if old > new {
		return -math.Abs((diff / new) * 100)
	}
	return math.Abs((diff / old) * 100)

}

func BigIntToFloat(input *big.Int) float64 {
	fl, _ := big.NewFloat(0).SetInt(input).Float64()
	return fl
}

func BigIntToFloatDiv(input *big.Int, devider float64) float64 {
	if devider == 1 {
		fl, _ := big.NewFloat(0).SetInt(input).Float64()
		return fl
	}
	f := 0.0
	if input != nil {
		divisor := big.NewInt(int64(devider / float64(100)))
		divisor = big.NewInt(0).Div(input, divisor)
		f = float64(divisor.Int64()) / 100
	}
	return f
}

func FloatToBigIntMul(input float64, multiplier float64) *big.Int {
	bigE18 := big.NewFloat(0).Mul(big.NewFloat(input), big.NewFloat(multiplier))
	result, _ := bigE18.Int(nil)
	return result
}

func FloatToBigInt(input float64) *big.Int {
	result, _ := big.NewFloat(input).Int(nil)
	return result
}

func IntTo32Byte(i uint32) [32]byte {
	bs := make([]byte, 4) // uint32 uses only 4 bytes
	binary.BigEndian.PutUint32(bs, i)

	// Add the bytes at the end of the array.
	var a = [32]byte{}
	a[28] = bs[0]
	a[29] = bs[1]
	a[30] = bs[2]
	a[31] = bs[3]
	return a
}
