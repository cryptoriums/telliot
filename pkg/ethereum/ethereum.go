// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package ethereum

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

const (
	PrivateKeysEnvName = "ETH_PRIVATE_KEYS"
	NodeURLEnvName     = "NODE_URL"
	ComponentName      = "ethereum"
	BlockTime          = float64(15)
	BlocksPerSecond    = float64(1 / BlockTime)
	BlocksPerMinute    = float64(60 / BlockTime)
	ReorgEventWait     = time.Minute
)

var ethAddressRE *regexp.Regexp = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

// ValidateAddress checks if an ethereum URL is valid?
func ValidateAddress(address string) error {
	if match := ethAddressRE.MatchString(address); !match {
		return errors.New("invalid ethereum address")
	}
	return nil
}

// GetAddressForNetwork returns an ethereum address based on ethereum node network id.
func GetAddressForNetwork(addresses string, networkID int64) (string, error) {
	// Parse addresses to a map.
	networkToAddress := make(map[string]string)
	_addresses := strings.Split(addresses, ",")
	for _, address := range _addresses {
		parts := strings.Split(strings.TrimSpace(address), ":")
		if len(parts) != 2 {
			return "", errors.New("malformed ethereum <network:address> string")
		}
		if err := ValidateAddress(parts[1]); err != nil {
			return "", err
		}
		networkToAddress[parts[0]] = parts[1]
	}

	switch networkID {
	case 1:
		if val, ok := networkToAddress["Mainnet"]; ok {
			return val, nil
		}
		return "", errors.New("address for the Mainnet network not found in the address list")
	case 4:
		if val, ok := networkToAddress["Rinkeby"]; ok {
			return val, nil
		}
		return "", errors.New("address for the Rinkeby network not found in the address list")
	default:
		return "", errors.New("unhandled network id")
	}
}

func DecodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}

func PrepareEthTransaction(
	ctx context.Context,
	client *ethclient.Client,
	account *Account,
	baseFeeGwei float64,
	tipFeeGwei float64,
	gasLimit uint64,
) (*bind.TransactOpts, error) {

	var baseFee *big.Int
	if baseFeeGwei > 0 {
		baseFee = big.NewInt(int64(baseFeeGwei) * params.GWei)
	}
	var tipFee *big.Int
	if tipFeeGwei > 0 {
		tipFee = big.NewInt(int64(tipFeeGwei) * params.GWei)
	}

	nonce, err := client.PendingNonceAt(ctx, account.GetAddress())
	if err != nil {
		return nil, errors.Wrap(err, "getting pending nonce")
	}

	if tipFee == nil {
		tipFee, err = client.SuggestGasTipCap(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "getting suggested gas tip")
		}
	}

	if baseFee == nil {
		header, err := client.HeaderByNumber(ctx, nil)
		if err != nil {
			return nil, errors.Wrap(err, "getting chain header")
		}
		// Add 25% more for the base fee as a safe margin in case of a network load surge.
		// At high network load the base fee increases 12.5% per block
		// so 25% will allow including the TX in the next 2 blocks if the network load surges.
		safeMargin := big.NewInt(0).Div(header.BaseFee, big.NewInt(4))
		baseFee = big.NewInt(0).Add(header.BaseFee, safeMargin)
	}

	gasPriceCap := big.NewInt(0).Add(baseFee, tipFee)

	ethBalance, err := client.BalanceAt(ctx, account.GetAddress(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting balance")
	}

	cost := new(big.Int)
	cost.Mul(gasPriceCap, big.NewInt(700000))
	if ethBalance.Cmp(cost) < 0 {
		return nil, errors.Errorf("insufficient ethereum to send a transaction: %v < %v", ethBalance, cost)
	}

	netID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting network id")
	}

	opts, err := bind.NewKeyedTransactorWithChainID(account.GetPrivateKey(), netID)
	if err != nil {
		return nil, errors.Wrap(err, "creating transactor")
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(0) // in wei

	if gasLimit == 0 {
		gasLimit = 3_000_000
	}
	opts.GasLimit = gasLimit // in units
	opts.GasTipCap = tipFee
	opts.GasFeeCap = gasPriceCap
	opts.Context = ctx
	return opts, nil
}

func Keccak256(input []byte) [32]byte {
	hash := crypto.Keccak256(input)
	var hashed [32]byte
	copy(hashed[:], hash)

	return hashed
}

type Account struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

func (a *Account) GetAddress() common.Address {
	return a.Address
}

func (a *Account) GetPrivateKey() *ecdsa.PrivateKey {
	return a.PrivateKey
}

func GetAccountByPubAddress(logger log.Logger, pubAddr string) (*Account, error) {
	accounts, err := GetAccounts(logger)
	if err != nil {
		return nil, errors.Wrap(err, "getting accounts")
	}

	for i, addr := range accounts {
		if addr.Address.Hex() == pubAddr {
			return accounts[i], nil
		}
	}
	return nil, errors.Errorf("account not found:%v", pubAddr)
}

// GetAccounts returns a slice of Account from private keys in
// PrivateKeysEnvName environment variable.
func GetAccounts(logger log.Logger) ([]*Account, error) {
	_privateKeys := os.Getenv(PrivateKeysEnvName)
	privateKeys := strings.Split(_privateKeys, ",")

	// Create an Account instance per private keys.
	accounts := make([]*Account, len(privateKeys))
	for i, pkey := range privateKeys {
		privateKey, err := crypto.HexToECDSA(strings.TrimSpace(pkey))
		if err != nil {
			return nil, errors.Wrap(err, "getting private key to ECDSA")
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, errors.New("casting public key to ECDSA")
		}

		publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		accounts[i] = &Account{Address: publicAddress, PrivateKey: privateKey}
		level.Info(logger).Log("msg", "registered account", "addr", publicAddress.Hex())
	}
	return accounts, nil
}

func NewClient(ctx context.Context, logger log.Logger) (*ethclient.Client, int64, error) {
	nodeURL := os.Getenv(NodeURLEnvName)

	client, err := ethclient.DialContext(ctx, nodeURL)
	if err != nil {
		return nil, 0, errors.Wrap(err, "create rpc client instance")
	}

	if !strings.Contains(strings.ToLower(nodeURL), "arbitrum") { // Arbitrum nodes doesn't support sync checking.
		// Issue #55, halt if client is still syncing with Ethereum network
		s, err := client.SyncProgress(ctx)
		if err != nil {
			return nil, 0, errors.Wrap(err, "determining if Ethereum client is syncing")
		}
		if s != nil {
			return nil, 0, errors.New("ethereum node is still syncing with the network")
		}
	}

	id, err := client.NetworkID(ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "get nerwork ID")
	}

	level.Info(logger).Log("msg", "created ethereum client", "netID", id.Int64())

	return client, id.Int64(), nil
}
