package contracts

import (
	"context"
	"encoding/hex"
	"math/big"
	"strings"
	"time"

	"github.com/cryptoriums/telliot/pkg/contracts/balancer"
	"github.com/cryptoriums/telliot/pkg/contracts/tellor"
	"github.com/cryptoriums/telliot/pkg/contracts/tellor_testing"
	"github.com/cryptoriums/telliot/pkg/contracts/uniswap"
	ethereumT "github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/math"
	psr "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

const (
	TellorAddress        = "0x88dF592F8eb5D7Bd38bFeF7dEb0fBc02cf3778a0"
	TellorAddressRinkeby = TellorAddress
	TellorAddressGoerli  = "0xf6c0da9600526C4D8411Ce5BD18059F1088a9580"
	TellorAddressHardhat = "0x8920050E1126125a27A4EaC5122AD3586c056E51"

	WithdrawStakeGasUsage          = 50_000
	RequestStakingWithdrawGasUsage = 100_000
	TallyGasUsage                  = 150_000
	DepositStakeGasUsage           = 160_000
	UnlockFeeGasUsage              = 200_000
	VoteGasUSage                   = 200_000
	BeginDisputeGasUsage           = 700_000
	SubmitMiningSolutionGasUsage   = 3_000_000

	EventNameNewTask    = "NewChallenge"
	EventNameNewSubmit  = "NonceSubmitted"
	EventNameTally      = "DisputeVoteTallied"
	EventNameNewDispute = "NewDispute"
	// When creating the bindings the Transfer event is aliased as Trasnfered
	// because it collides with the same function name.
	EventNameTransfer        = "Transferred"
	EventNameNewVote         = "Voted"
	EventNameContractUpgrade = "NewTellorAddress"

	DefaultDisputeVotingWindow = 24 * 2 * time.Hour // 2 days to vote for a dispute.
	DefaultDisputeUnlockWindow = 24 * time.Hour     // 1 day to wait after a dispute before can get the dispute fee.
)

type ContractCaller interface {
	GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error)
	SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error)
	GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error)
	GetNewCurrentVariables(opts *bind.CallOpts) (struct {
		Challenge  [32]byte
		RequestIds [5]*big.Int
		Difficulty *big.Int
		Tip        *big.Int
	}, error)
	Addr() common.Address
	NetID() int64
	Abi() abi.ABI
	BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error)
	GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error)
	GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error)
	TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error)
	UnlockDisputeFee(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error)
	DepositStake(opts *bind.TransactOpts) (*types.Transaction, error)
	UnpackLog(out interface{}, event string, log types.Log) error
	WatchLogs(opts *bind.WatchOpts, name string, query ...[]interface{}) (chan types.Log, event.Subscription, error)
	DisputeVotingWindow() time.Duration
	DisputeUnlockWindow() time.Duration
	GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error)
	GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error)
	GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error)
	GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error)
	GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error)
}

type RewardQuerier interface {
	// Returns the current reward in percent relative to the provided tx cost.
	Query(ctx context.Context, txCostGwei float64) (int64, error)
}

var DefaultParams = Params{
	DisputeVotingWindow: DefaultDisputeVotingWindow,
	DisputeUnlockWindow: DefaultDisputeUnlockWindow,
}

type Params struct {
	DisputeVotingWindow time.Duration
	DisputeUnlockWindow time.Duration
}

type (
	TellorNewDispute         = tellor.TellorNewDispute
	TellorNonceSubmitted     = tellor.TellorNonceSubmitted
	TellorTransferred        = tellor.TellorTransferred
	TellorNewValue           = tellor.TellorNewValue
	TellorDisputeVoteTallied = tellor.ExtensionDisputeVoteTallied
	NewTellorAddress         = tellor.ExtensionNewTellorAddress
	TellorNewChallenge       = tellor.TellorNewChallenge
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
	netID         int64
	params        Params
	*tellor.ITellor
	address common.Address
}

func (self *ITellor) NetID() int64 {
	return self.netID
}

func (self *ITellor) Addr() common.Address {
	return self.address
}

func (self *ITellor) Abi() abi.ABI {
	return self.abi
}

func (self *ITellor) DisputeVotingWindow() time.Duration {
	return self.params.DisputeVotingWindow
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

func NewITellor(logger log.Logger, ctx context.Context, client *ethclient.Client, netID int64, params Params) (*ITellor, error) {
	contractAddr, err := GetTellorAddress(netID)
	if err != nil {
		return nil, errors.Wrap(err, "getting contract address")
	}

	return NewITellorWithAddr(logger, ctx, contractAddr, client, netID, params)
}

func NewITellorWithAddr(
	logger log.Logger,
	ctx context.Context,
	contractAddr common.Address,
	client *ethclient.Client,
	netID int64,
	params Params,
) (*ITellor, error) {
	if params.DisputeVotingWindow == 0 {
		return nil, errors.New("DisputeVotingWindow should not be zero")
	}
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
		netID:         netID,
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

func NewITellorTest(ctx context.Context, client *ethclient.Client, netID int64) (*ITellorTest, error) {
	conractAddr, err := GetTellorAddress(netID)
	if err != nil {
		return nil, errors.Wrap(err, "getting contract address")
	}

	tellorInstance, err := tellor_testing.NewITellor(conractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "creating contract interface")
	}

	return &ITellorTest{
		Address: conractAddr,
		ITellor: tellorInstance,
	}, nil
}

type ITellorFilterer struct {
	*tellor.TellorFilterer
	*tellor.ExtensionFilterer
}

func NewITellorFilterer(addr common.Address, client *ethclient.Client) (*ITellorFilterer, error) {

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
	DataTs       time.Time
	Executed     bool
	Passed       bool
	Disputer     common.Address
	Disputed     common.Address
	DisputedSlot uint64
	Tally        float64
	Votes        uint64
	Created      time.Time
	Ends         time.Time
	Fee          float64
	TxHash       common.Hash
}

func GetDisputeLogs(
	ctx context.Context,
	client *ethclient.Client,
	contract ContractCaller,
	lookBackDuration time.Duration,
) ([]*DisputeLog, error) {
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "get latest eth block header")
	}

	lookBackDelta := big.NewInt(int64(lookBackDuration.Minutes() * ethereumT.BlocksPerMinute))
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

		logUnpacked, err := GetDisputeInfo(ctx, log.DisputeId, contract)
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
	client *ethclient.Client,
	contract ContractCaller,
	lookBackDuration time.Duration,
) ([]*TellorDisputeVoteTallied, error) {

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "get latest eth block header")
	}

	lookBackDelta := big.NewInt(int64(lookBackDuration.Minutes() * ethereumT.BlocksPerMinute))
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

func GetDisputeInfo(ctx context.Context, disputeID *big.Int, contract ContractCaller) (*DisputeLog, error) {
	_, executed, passed, _, disputed, disputer, _, disputeVars, tally, err := contract.GetAllDisputeVars(&bind.CallOpts{Context: ctx}, disputeID)
	if err != nil {
		return nil, errors.Wrap(err, "get dispute details")
	}

	rounds, err := contract.GetDisputeUintVars(
		&bind.CallOpts{Context: ctx},
		disputeID,
		ethereumT.Keccak256([]byte("_DISPUTE_ROUNDS")),
	)
	if err != nil {
		return nil, errors.Wrap(err, "get dispute rounds")
	}

	votingEnds := time.Unix(disputeVars[3].Int64(), 0)
	votingWindow := contract.DisputeVotingWindow()
	created := votingEnds.Add(-votingWindow * time.Duration(rounds.Int64()))

	return &DisputeLog{
		ID:           disputeID.Int64(),
		DataID:       disputeVars[0].Int64(),
		DataTs:       time.Unix(disputeVars[1].Int64(), 0),
		DataVal:      math.BigIntToFloatDiv(disputeVars[2], psr.DefaultGranularity),
		Executed:     executed,
		Passed:       passed,
		Disputer:     disputer,
		Disputed:     disputed,
		DisputedSlot: disputeVars[6].Uint64(),
		Tally:        math.BigIntToFloatDiv(tally, params.Ether),
		Votes:        disputeVars[4].Uint64(),
		Created:      created,
		Ends:         votingEnds,
		Fee:          math.BigIntToFloatDiv(disputeVars[8], params.Ether),
	}, nil
}

type SubmitBlock struct {
	Timestamp   *big.Int
	DataIDs     [5]*big.Int
	FinalValues map[int64]*big.Int
	Reporters   map[int64][5]common.Address
	Values      map[int64][5]*big.Int
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
	client *ethclient.Client,
	contract ContractCaller,
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
		blockNums := int64(ethereumT.BlocksPerMinute * time.Since(time.Unix(from, 0)).Minutes())
		// Subtract form the current header block number to use as the upper limit.
		endBlock = endBlock - blockNums
	}

	lookBackDelta := big.NewInt(int64(lookBackDuration.Minutes() * ethereumT.BlocksPerMinute))
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
		}
		submitBlocks = append(submitBlocks, submitBlock)
	}
	return submitBlocks, err
}

func LastSubmit(contract ContractCaller, reporter common.Address) (time.Duration, *time.Time, error) {
	address := "000000000000000000000000" + reporter.Hex()[2:]
	decoded, err := hex.DecodeString(address)
	if err != nil {
		return 0, nil, errors.Wrapf(err, "decoding address")
	}
	last, err := contract.GetUintVar(nil, ethereumT.Keccak256(decoded))

	if err != nil {
		return 0, nil, errors.Wrapf(err, "getting last submit time for:%v", reporter.Hex())
	}
	// The reporter has never submitted so put a timestamp at the beginning of unix time.
	if last.Int64() == 0 {
		last.Set(big.NewInt(1))
	}

	lastInt := last.Int64()
	var lastSubmit time.Duration
	var tm time.Time
	if lastInt > 0 {
		tm = time.Unix(lastInt, 0)
		lastSubmit = time.Since(tm)
	}

	return lastSubmit, &tm, nil
}

func Slot(caller ContractCaller) (*big.Int, error) {
	slot, err := caller.GetUintVar(nil, ethereumT.Keccak256([]byte("_SLOT_PROGRESS")))
	if err != nil {
		return nil, errors.Wrap(err, "getting _SLOT_PROGRESS")
	}
	return slot, nil
}
