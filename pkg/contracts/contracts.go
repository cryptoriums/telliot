package contracts

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math"
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/cryptoriums/telliot/pkg/contracts/balancer"
	"github.com/cryptoriums/telliot/pkg/contracts/tellor"
	"github.com/cryptoriums/telliot/pkg/contracts/tellorX_oracle"
	"github.com/cryptoriums/telliot/pkg/contracts/tellor_proxy"
	"github.com/cryptoriums/telliot/pkg/contracts/tellor_testing"
	"github.com/cryptoriums/telliot/pkg/contracts/uniswap"
	ethereum_t "github.com/cryptoriums/telliot/pkg/ethereum"
	mathT "github.com/cryptoriums/telliot/pkg/math"
	psr "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

const (
	TellorAddress            = "0x88dF592F8eb5D7Bd38bFeF7dEb0fBc02cf3778a0"
	TellorAddressRinkeby     = "0x6807d197dAc4c131f2390B7a7F0c9199df6f70f4"
	TellorAddressGoerli      = "0xe5e09e1C64Eab3cA8bCAD722b0966B69931879ae"
	TellorAddressGoerliProxy = "0x84Ec18B070D84e347eE6B7D5fA2d9fcFfbf759bA" // Proxy contract for testing.
	TellorAddressHardhat     = "0x8920050E1126125a27A4EaC5122AD3586c056E51"

	WithdrawStakeGasLimit          = 50_000
	RequestStakingWithdrawGasLimit = 100_000
	TallyGasLimit                  = 150_000
	DepositStakeGasLimit           = 160_000
	UnlockFeeGasLimit              = 300_000
	VoteGasUSage                   = 200_000
	NewDisputeGasLimit             = 700_000
	SubmitGasLimit                 = 3_000_000

	MethodNameSubmit     = "submitValue"
	MethodNameNewDispute = "beginDispute"

	EventNameNewReport       = "NewReport"
	EventNameNewTip          = "TipAdded"
	EventNameTally           = "DisputeVoteTallied"
	EventNameNewDispute      = "NewDispute"
	EventNameTransfer        = "Transfer"
	EventNameNewVote         = "Voted"
	EventNameContractUpgrade = "NewTellorAddress"

	DefaultDisputeUnlockWindow = 24 * time.Hour // 1 day to wait after a dispute before can get the dispute fee.
)

type TellorCaller interface {
	GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error)
	SubmitValue(opts *bind.TransactOpts, _nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error)
	GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error)
	// GetTimestampCountById returns the current nonce that needs to be used when making a submit.
	GetTimestampCountById(opts *bind.CallOpts, _queryId [32]byte) (*big.Int, error)
	DepositStake(opts *bind.TransactOpts) (*types.Transaction, error)
	WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error)
	RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error)

	// GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error)
	// GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error)

	Addr() common.Address
	Abi() abi.ABI
	AbiRaw() string
	UnpackLog(out interface{}, event string, log types.Log) error
	WatchLogs(opts *bind.WatchOpts, name string, query ...[]interface{}) (chan types.Log, event.Subscription, error)

	BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error)
	Transfer(*bind.TransactOpts, common.Address, *big.Int) (*types.Transaction, error)
	Approve(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error)

	TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error)
	BeginDispute(opts *bind.TransactOpts, _queryId [32]byte, _timestamp *big.Int) (*types.Transaction, error)
	DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error)
	Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supports bool, _invalidQuery bool) (*types.Transaction, error)

	// GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error)
	// GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error)
	// UnlockDisputeFee(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error)
	// DisputeUnlockWindow() time.Duration

}

type RewardQuerier interface {
	// Returns the current reward in percent relative to the provided tx cost.
	Query(ctx context.Context, txCostGwei float64) (int64, error)
}

var DefaultParams = Params{
	DisputeUnlockWindow: DefaultDisputeUnlockWindow,
}

type Params struct {
	DisputeUnlockWindow time.Duration
}

type (
	TellorNewDispute         = tellor.TellorNewDispute
	TellorNonceSubmitted     = tellor.TellorNonceSubmitted
	TellorTransferred        = tellor.TellorTransferred
	TellorNewValue           = tellor.TellorNewValue
	TellorDisputeVoteTallied = tellor.ExtensionDisputeVoteTallied
	NewTellorAddress         = tellor.ExtensionNewTellorAddress
	TellorNewReport          = tellorX_oracle.OracleNewReport
	TellorNewTip             = tellorX_oracle.OracleTipAdded
	TellorVoted              = tellor.TellorVoted
)

var (
	BPoolABI          = balancer.BPoolABI
	BTokenABI         = balancer.BTokenABI
	IERC20ABI         = uniswap.IERC20ABI
	IUniswapV2PairABI = uniswap.IUniswapV2PairABI
	TellorABI         = tellor.ITellorABI
)

type ITellor struct {
	boundContract *bind.BoundContract
	abi           abi.ABI
	abiRaw        string
	params        Params
	*tellor.ITellor
	address common.Address
}

func (self *ITellor) Addr() common.Address {
	return self.address
}

func (self *ITellor) Abi() abi.ABI {
	return self.abi
}

func (self *ITellor) AbiRaw() string {
	return self.abiRaw
}

func (self *ITellor) DisputeUnlockWindow() time.Duration {
	return self.params.DisputeUnlockWindow
}

func (self *ITellor) UnpackLog(out interface{}, event string, log types.Log) error {
	return self.boundContract.UnpackLog(out, event, log)
}

func (self *ITellor) WatchLogs(opts *bind.WatchOpts, name string, query ...[]interface{}) (chan types.Log, event.Subscription, error) {
	return self.boundContract.WatchLogs(opts, name, query...)
}

func NewITellor(
	ctx context.Context,
	logger log.Logger,
	client ethereum_t.EthClient,
	address common.Address,
	params Params,
) (*ITellor, error) {

	var contractAddr common.Address
	if address != (common.Address{}) {
		contractAddr = address
	} else {
		var err error
		contractAddr, err = GetTellorAddress(client.NetworkID())
		if err != nil {
			return nil, errors.Wrap(err, "getting contract address")
		}
	}

	contract, err := newITellorWithAddr(logger, contractAddr, client, params)
	if err != nil {
		return nil, errors.Wrap(err, "creating contract instance")
	}
	return contract, nil
}

func newITellorWithAddr(
	logger log.Logger,
	contractAddr common.Address,
	client ethereum_t.EthClient,
	params Params,
) (*ITellor, error) {
	if params.DisputeUnlockWindow == 0 {
		return nil, errors.New("DisputeUnlockWindow should not be zero")
	}
	tellorInstance, err := tellor.NewITellor(contractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "creating contract interface")
	}
	level.Info(logger).Log("msg", "tellor contract address", "addr", contractAddr)

	abi, err := abi.JSON(strings.NewReader(TellorABI))
	if err != nil {
		return nil, errors.New("parsing contract ABI")
	}

	boundContract := bind.NewBoundContract(contractAddr, abi, client, client, client)

	return &ITellor{
		abi:           abi,
		abiRaw:        tellor.ITellorABI,
		address:       contractAddr,
		boundContract: boundContract,
		params:        params,
		ITellor:       tellorInstance,
	}, nil
}

type ITellorTest struct {
	*tellor_testing.ITellor
	Address common.Address
}

func NewITellorTest(ctx context.Context, contractAddr common.Address, client ethereum_t.EthClient) (*ITellorTest, error) {
	var err error
	if contractAddr == (common.Address{}) {
		contractAddr, err = GetTellorAddress(client.NetworkID())
		if err != nil {
			return nil, errors.Wrap(err, "getting contract address")
		}
	}

	tellorInstance, err := tellor_testing.NewITellor(contractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "creating contract interface")
	}

	return &ITellorTest{
		Address: contractAddr,
		ITellor: tellorInstance,
	}, nil
}

type ITellorProxy struct {
	*tellor_proxy.Reporter
	Address common.Address
}

func NewITellorProxy(ctx context.Context, contractAddr common.Address, client ethereum_t.EthClient) (*ITellorProxy, error) {
	var err error
	if contractAddr == (common.Address{}) {
		contractAddr = common.HexToAddress(TellorAddressGoerliProxy)
	}

	tellorInstance, err := tellor_proxy.NewReporter(contractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "creating contract interface")
	}

	return &ITellorProxy{
		Address:  contractAddr,
		Reporter: tellorInstance,
	}, nil
}

type ITellorFilterer struct {
	*tellor.TellorFilterer
	*tellor.ExtensionFilterer
}

func NewITellorFilterer(addr common.Address, client ethereum_t.EthClient) (*ITellorFilterer, error) {

	ft, err := tellor.NewTellorFilterer(addr, client)
	if err != nil {
		return nil, errors.Wrap(err, "create log filtered")
	}

	fe, err := tellor.NewExtensionFilterer(addr, client)
	if err != nil {
		return nil, errors.Wrap(err, "create log filtered")
	}
	return &ITellorFilterer{
		TellorFilterer:    ft,
		ExtensionFilterer: fe,
	}, nil
}

func GetTellorAddress(netID int64) (common.Address, error) {
	switch netID {
	case 1:
		return common.HexToAddress(TellorAddress), nil
	case 4:
		return common.HexToAddress(TellorAddressRinkeby), nil
	case 5:
		return common.HexToAddress(TellorAddressGoerli), nil
	case 31337:
		return common.HexToAddress(TellorAddressHardhat), nil
	default:
		return common.Address{}, errors.Errorf("tellor contract address for current network id not found:%v", netID)
	}
}

type DisputeLog struct {
	ID           int64
	DataID       int64
	DataVal      float64
	DataTime     time.Time
	Executed     bool
	Passed       bool
	Disputer     common.Address
	Disputed     common.Address
	DisputedSlot uint64
	Tally        float64
	Votes        uint64
	Ends         time.Time
	Fee          float64
	TxHash       common.Hash
}

func GetDisputeLogs(
	ctx context.Context,
	logger log.Logger,
	client ethereum_t.EthClient,
	contract TellorCaller,
	lookBackDuration time.Duration,
) ([]*DisputeLog, error) {
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "get latest eth block header")
	}

	lookBackDelta := big.NewInt(int64(lookBackDuration.Minutes() * ethereum_t.BlocksPerMinute))
	startBlock := big.NewInt(0).Sub(header.Number, lookBackDelta)

	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   nil,
		Addresses: []common.Address{contract.Addr()},
		Topics:    [][]common.Hash{{contract.Abi().Events[EventNameNewDispute].ID}},
	}

	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "get logs")
	}

	var unpackedLogs []*DisputeLog
	for _, logRaw := range logs {
		log := TellorNewDispute{}
		err := contract.UnpackLog(&log, "NewDispute", logRaw)
		if err != nil {
			return nil, errors.Wrap(err, "unpack event from logs")
		}

		logUnpacked, err := GetDisputeInfo(ctx, logger, log.DisputeId, contract)
		if err != nil {
			return nil, errors.Wrap(err, "getting dispute details")
		}
		logUnpacked.TxHash = logRaw.TxHash
		unpackedLogs = append(unpackedLogs, logUnpacked)

	}
	return unpackedLogs, err
}

func GetTallyLogs(
	ctx context.Context,
	client ethereum_t.EthClient,
	contract TellorCaller,
	lookBackDuration time.Duration,
) ([]*TellorDisputeVoteTallied, error) {

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "get latest eth block header")
	}

	lookBackDelta := big.NewInt(int64(lookBackDuration.Minutes() * ethereum_t.BlocksPerMinute))
	startBlock := big.NewInt(0).Sub(header.Number, lookBackDelta)

	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   nil,
		Addresses: []common.Address{contract.Addr()},
		Topics:    [][]common.Hash{{contract.Abi().Events[EventNameTally].ID}},
	}

	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "get logs")
	}

	var unpackedLogs []*TellorDisputeVoteTallied
	for _, logRaw := range logs {
		log := &TellorDisputeVoteTallied{}
		err := contract.UnpackLog(log, "DisputeVoteTallied", logRaw)
		if err != nil {
			return nil, errors.Wrap(err, "unpack event from logs")
		}
		log.Raw = logRaw
		unpackedLogs = append(unpackedLogs, log)

	}
	return unpackedLogs, err
}

func GetDisputeInfo(ctx context.Context, logger log.Logger, disputeID *big.Int, contract TellorCaller) (*DisputeLog, error) {
	ticker := time.NewTicker(1)
	defer ticker.Stop()
	var resetTicker sync.Once
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			resetTicker.Do(func() { ticker.Reset(5 * time.Second) })
		}
		_, executed, passed, _, disputed, disputer, _, disputeVars, tally, err := contract.GetAllDisputeVars(&bind.CallOpts{Context: ctx}, disputeID)
		if err != nil {
			return nil, errors.Wrap(err, "get dispute details")
		}
		if disputer == (common.Address{}) {
			level.Error(logger).Log("msg", "dispute doesn't exist", "id", disputeID)
			continue
		}

		votingEnds := time.Unix(disputeVars[3].Int64(), 0)

		return &DisputeLog{
			ID:           disputeID.Int64(),
			DataID:       disputeVars[0].Int64(),
			DataTime:     time.Unix(disputeVars[1].Int64(), 0),
			DataVal:      mathT.BigIntToFloatDiv(disputeVars[2], psr.DefaultGranularity),
			Executed:     executed,
			Passed:       passed,
			Disputer:     disputer,
			Disputed:     disputed,
			DisputedSlot: disputeVars[6].Uint64(),
			Tally:        mathT.BigIntToFloatDiv(tally, params.Ether),
			Votes:        disputeVars[4].Uint64(),
			Ends:         votingEnds,
			Fee:          mathT.BigIntToFloatDiv(disputeVars[8], params.Ether),
		}, nil
	}
}

type SubmitBlock struct {
	Timestamp   *big.Int
	DataIDs     [5]*big.Int
	FinalValues map[int64]*big.Int
	Reporters   map[int64][5]common.Address
	Values      map[int64][5]*big.Int
	TxHash      common.Hash
}

func NewSubmitBlock() SubmitBlock {
	return SubmitBlock{
		FinalValues: make(map[int64]*big.Int),
		Reporters:   make(map[int64][5]common.Address),
		Values:      make(map[int64][5]*big.Int),
	}
}

func GetSubmitLogs(
	ctx context.Context,
	client ethereum_t.EthClient,
	contract TellorCaller,
	from int64,
	lookBackDuration time.Duration,
) ([]SubmitBlock, error) {
	headerNow, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "get latest eth block header")
	}

	endBlock := headerNow.Number.Int64()

	if from != 0 {
		// Total block numbers that correspond to this TS calculated from the current time.
		blockNums := int64(ethereum_t.BlocksPerMinute * time.Since(time.Unix(from, 0)).Minutes())
		// Subtract form the current header block number to use as the upper limit.
		endBlock = endBlock - blockNums
	}

	lookBackDelta := big.NewInt(int64(lookBackDuration.Minutes() * ethereum_t.BlocksPerMinute))
	startBlock := big.NewInt(0).Sub(big.NewInt(endBlock), lookBackDelta)

	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   big.NewInt(endBlock),
		Addresses: []common.Address{contract.Addr()},
		Topics:    [][]common.Hash{{contract.Abi().Events["NewValue"].ID}},
	}

	events, err := client.FilterLogs(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "get events")
	}

	var submitBlocks []SubmitBlock
	for _, eventRaw := range events {
		submitBlock := NewSubmitBlock()

		eventNewBlock := &TellorNewValue{}
		err := contract.UnpackLog(eventNewBlock, "NewValue", eventRaw)
		if err != nil {
			return nil, errors.Wrap(err, "unpack event from logs")
		}

		submitBlock.Timestamp = eventNewBlock.Time
		submitBlock.DataIDs = eventNewBlock.RequestId
		for index, dataID := range eventNewBlock.RequestId {
			reporters, err := contract.GetMinersByRequestIdAndTimestamp(&bind.CallOpts{Context: ctx}, dataID, eventNewBlock.Time)
			if err != nil {
				return nil, errors.Wrap(err, "getting submit block reporters")
			}
			vals, err := contract.GetSubmissionsByTimestamp(&bind.CallOpts{Context: ctx}, dataID, eventNewBlock.Time)
			if err != nil {
				return nil, errors.Wrap(err, "getting submit block vals")
			}
			submitBlock.FinalValues[dataID.Int64()] = eventNewBlock.Value[index]
			submitBlock.Reporters[dataID.Int64()] = reporters
			submitBlock.Values[dataID.Int64()] = vals
			submitBlock.TxHash = eventRaw.TxHash
		}
		sort.Slice(submitBlock.DataIDs[:], func(i, j int) bool {
			return submitBlock.DataIDs[i].Int64() < submitBlock.DataIDs[j].Int64()
		})
		submitBlocks = append(submitBlocks, submitBlock)
	}

	return submitBlocks, err
}

func LastSubmit(ctx context.Context, contract TellorCaller, reporter common.Address) (time.Duration, *time.Time, error) {
	address := "000000000000000000000000" + reporter.Hex()[2:]
	decoded, err := hex.DecodeString(address)
	if err != nil {
		return 0, nil, errors.Wrapf(err, "decoding address")
	}
	last, err := contract.GetUintVar(&bind.CallOpts{Context: ctx}, ethereum_t.Keccak256(decoded))

	if err != nil {
		return 0, nil, errors.Wrapf(err, "getting last submit time for:%v", reporter.Hex())
	}
	// The reporter has never submitted so put a timestamp at the beginning of unix time.
	if last.Int64() == 0 {
		last.Set(big.NewInt(1))
	}

	lastInt := last.Int64()
	sinceLastSubmit := time.Duration(math.MaxInt64)
	var tm time.Time
	if lastInt > 0 {
		tm = time.Unix(lastInt, 0)
		sinceLastSubmit = time.Since(tm)
	}

	return sinceLastSubmit, &tm, nil
}

func Slot(caller TellorCaller) (*big.Int, error) {
	slot, err := caller.GetUintVar(nil, ethereum_t.Keccak256([]byte("_SLOT_PROGRESS")))
	if err != nil {
		return nil, errors.Wrap(err, "getting _SLOT_PROGRESS")
	}
	return slot, nil
}

func CreateTellorTx(
	ctx context.Context,
	prvKey *ecdsa.PrivateKey,
	to common.Address,
	client ethereum_t.EthClient,
	overwritePending bool,
	gasLimit uint64,
	gasMaxFee float64,
	methodName string,
	args []interface{},
) (*types.Transaction, string, error) {
	var nonce uint64
	var err error
	if overwritePending {
		nonce, err = client.NonceAt(ctx, crypto.PubkeyToAddress(prvKey.PublicKey), nil)
		if err != nil {
			return nil, "", errors.Wrap(err, "getting last nonce")
		}
	} else {
		nonce, err = client.PendingNonceAt(ctx, crypto.PubkeyToAddress(prvKey.PublicKey))
		if err != nil {
			return nil, "", errors.Wrap(err, "getting pending nonce")
		}
	}

	abiP, err := abi.JSON(strings.NewReader(tellor.TellorABI))
	if err != nil {
		return nil, "", errors.Wrap(err, "read contract ABI")
	}

	data, err := abiP.Pack(methodName, args...)
	if err != nil {
		return nil, "", errors.Wrap(err, "packing ABI")
	}

	return ethereum_t.NewSignedTX(
		to,
		data,
		nonce,
		prvKey,
		client.NetworkID(),
		gasLimit,
		gasMaxFee,
	)
}

func ReporterStatusName(statusID int64) string {
	// From https://github.com/tellor-io/tellor3/blob/7c2f38a0e3f96631fb0f96e0d0a9f73e7b355766/contracts/TellorStorage.sol#L41
	switch statusID {
	case 0:
		return "Not staked"
	case 1:
		return "Staked"
	case 2:
		return "LockedForWithdraw"
	case 3:
		return "OnDispute"
	case 4:
		return "ReadyForUnlocking"
	case 5:
		return "Unlocked"
	default:
		return "Unknown"
	}
}
