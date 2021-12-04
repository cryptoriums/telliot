package contracts

import (
	"context"
	"math"
	"math/big"
	"strings"
	"time"

	ethereum_t "github.com/cryptoriums/packages/ethereum"
	math_t "github.com/cryptoriums/packages/math"
	"github.com/cryptoriums/telliot/pkg/contracts/tellorX_governance"
	"github.com/cryptoriums/telliot/pkg/contracts/tellorX_master"
	"github.com/cryptoriums/telliot/pkg/contracts/tellorX_oracle"
	psr_tellor "github.com/cryptoriums/telliot/pkg/psr/tellor"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
)

const (
	MasterAddress            = "0x88dF592F8eb5D7Bd38bFeF7dEb0fBc02cf3778a0"
	MasterAddressRinkeby     = "0x5373Fc8Cf8E1dfa91796f9c308A0756EE5b9ABC0"
	MasterAddressGoerli      = "0xd9991C48443eFdaEdE8453C9B8fB017f0642be66"
	MasterAddressGoerliProxy = "0x84Ec18B070D84e347eE6B7D5fA2d9fcFfbf759bA" // Proxy contract for testing.
	MasterAddressHardhat     = "0x8920050E1126125a27A4EaC5122AD3586c056E51"

	// All of these are 50% more than the actual usage.
	SubmitGasLimit = 400_000

	StakeRequestWithdrawGasLimit = 100_000
	StakeWithdrawGasLimit        = 50_000
	StakeDepositGasLimit         = 150_000
	DisputeNewGasLimit           = 1_000_000
	VoteGasUSage                 = 200_000
	VoteTallyGasLimit            = 100_000
	VoteExecuteGasLimit          = 220_000

	MethodNameSubmit      = "submitValue"
	MethodNameNewDispute  = "beginDispute"
	MethodNameTallyVote   = "tallyVotes"
	MethodNameVote        = "vote"
	MethodNameNewVote     = "proposeVote"
	MethodNameExecuteVote = "executeVote"
	MethodNameAddTip      = "tipQuery"

	EventNameNewSubmit       = "NewReport"
	EventNameNewTip          = "TipAdded"
	EventNameNewDispute      = "NewDispute"
	EventNameTransfer        = "Transfer"
	EventNameNewVote         = "NewVote"
	EventNameVoted           = "Voted"
	EventNameTallyVote       = "VoteTallied"
	EventNameExecuteVote     = "VoteExecuted"
	EventNameContractUpgrade = "NewContractAddress"

	DefaultVotingDuration          = 7 * 24 * time.Hour // The default voting period.
	DefaultVoteExecuteWaitDuration = 24 * time.Hour     // The default wait period for executing the dispute and collecing the dispute fee.
)

type ContractCaller interface {
	Addr() common.Address
	Abi() abi.ABI
	AbiRaw() string
	UnpackLog(out interface{}, event string, log types.Log) error
	WatchLogs(opts *bind.WatchOpts, name string, query ...[]interface{}) (chan types.Log, event.Subscription, error)
}

type TellorMasterCaller interface {
	ContractCaller

	BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error)
	Transfer(*bind.TransactOpts, common.Address, *big.Int) (*types.Transaction, error)
	Approve(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error)

	GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error)

	RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error)
	DepositStake(opts *bind.TransactOpts) (*types.Transaction, error)
	GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error)
	WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error)

	ChangeOracleContract(opts *bind.TransactOpts, _newOracle common.Address) (*types.Transaction, error)
}

type TellorOracleCaller interface {
	ContractCaller

	SubmitValue(opts *bind.TransactOpts, _queryId [32]byte, _value []byte, _nonce *big.Int, _queryData []byte) (*types.Transaction, error)
	// GetTimestampCountById returns the current nonce that needs to be used when making a submit.
	GetTimestampCountById(opts *bind.CallOpts, _queryId [32]byte) (*big.Int, error)
	TimeOfLastNewValue(opts *bind.CallOpts) (*big.Int, error)
	GetCurrentReward(opts *bind.CallOpts, _queryId [32]byte) (*big.Int, *big.Int, error)
	GetReporterLastTimestamp(opts *bind.CallOpts, _reporter common.Address) (*big.Int, error)
	GetReporterByTimestamp(opts *bind.CallOpts, _queryId [32]byte, _timestamp *big.Int) (common.Address, error)
	GetValueByTimestamp(opts *bind.CallOpts, _queryId [32]byte, _timestamp *big.Int) ([]byte, error)
	ReportingLock(opts *bind.CallOpts) (*big.Int, error)
	TipsInContract(opts *bind.CallOpts) (*big.Int, error)
	Tips(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error)
}

const (
	VoteStatusOpen = iota - 1
	VoteStatusRejected
	VoteStatusPassed
	VoteStatusInvalid
	VoteStatusTallied

	VoteIDLabel = "id"
)

type TellorGovernCaller interface {
	ContractCaller

	GetDisputeInfo(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, *big.Int, []byte, common.Address, error)
	GetVoteInfo(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, [8]*big.Int, [2]bool, uint8, []byte, [4]byte, [2]common.Address, error)
	TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error)
	ExecuteVote(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error)
	BeginDispute(opts *bind.TransactOpts, _queryId [32]byte, _timestamp *big.Int) (*types.Transaction, error)
	DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error)
	Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supports bool, _invalidQuery bool) (*types.Transaction, error)
	DisputeFee(opts *bind.CallOpts) (*big.Int, error)
	VoteCount(opts *bind.CallOpts) (*big.Int, error) // TODO remove once the NewVote event includes the vote ID.

	VotingDuration() time.Duration
	VoteExecuteWaitDuration() time.Duration
}

type RewardQuerier interface {
	// Returns the current reward in percent relative to the provided tx cost.
	Query(ctx context.Context, txCostGwei float64) (int64, error)
}

var DefaultParams = Params{
	VotingDuration:          DefaultVotingDuration,
	VoteExecuteWaitDuration: DefaultVoteExecuteWaitDuration,
}

type Params struct {
	VotingDuration          time.Duration
	VoteExecuteWaitDuration time.Duration
}

type TellorMaster struct {
	abi    abi.ABI
	abiRaw string
	*tellorX_master.Controller
	*bind.BoundContract
	address common.Address
}

func (self *TellorMaster) Transfer(opts *bind.TransactOpts, to common.Address, amnt *big.Int) (*types.Transaction, error) {
	return self.Controller.Transfer(opts, to, amnt)
}

func (self *TellorMaster) Addr() common.Address {
	return self.address
}

func (self *TellorMaster) Abi() abi.ABI {
	return self.abi
}

func (self *TellorMaster) AbiRaw() string {
	return self.abiRaw
}

type TellorOracle struct {
	*bind.BoundContract
	abi    abi.ABI
	abiRaw string
	*tellorX_oracle.Oracle
	address common.Address
}

func (self *TellorOracle) Addr() common.Address {
	return self.address
}

func (self *TellorOracle) Abi() abi.ABI {
	return self.abi
}

func (self *TellorOracle) AbiRaw() string {
	return self.abiRaw
}

type TellorGovern struct {
	*bind.BoundContract
	abi    abi.ABI
	abiRaw string
	params Params
	*tellorX_governance.Governance
	address common.Address
}

func (self *TellorGovern) Addr() common.Address {
	return self.address
}

func (self *TellorGovern) Abi() abi.ABI {
	return self.abi
}

func (self *TellorGovern) AbiRaw() string {
	return self.abiRaw
}

func (self *TellorGovern) VotingDuration() time.Duration {
	return self.params.VotingDuration
}

func (self *TellorGovern) VoteExecuteWaitDuration() time.Duration {
	return self.params.VoteExecuteWaitDuration
}

func NewContracts(
	ctx context.Context,
	logger log.Logger,
	client ethereum_t.EthClient,
	address common.Address,
	params Params,
) (TellorMasterCaller, TellorOracleCaller, TellorGovernCaller, error) {

	var masterAddr common.Address
	if address != (common.Address{}) {
		masterAddr = address
	} else {
		var err error
		masterAddr, err = GetMasterAddress(client.NetworkID())
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "getting contract address")
		}
	}

	master, oracle, govern, err := newContractsWithAddr(ctx, logger, masterAddr, client, params)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "creating contract instance")
	}
	return master, oracle, govern, nil
}

func newContractsWithAddr(
	ctx context.Context,
	logger log.Logger,
	addrMaster common.Address,
	client ethereum_t.EthClient,
	params Params,
) (TellorMasterCaller, TellorOracleCaller, TellorGovernCaller, error) {
	if params.VotingDuration == 0 {
		return nil, nil, nil, errors.New("VotingDuration should not be zero")
	}
	if params.VoteExecuteWaitDuration == 0 {
		return nil, nil, nil, errors.New("VoteExecuteWaitDuration should not be zero")
	}
	master, err := tellorX_master.NewController(addrMaster, client)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "creating contract interface")
	}

	addrOracle, err := master.Addresses(&bind.CallOpts{Context: ctx}, ethereum_t.Keccak256("_ORACLE_CONTRACT"))
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "getting contract address")
	}
	if addrOracle == common.BytesToAddress([]byte("")) {
		return nil, nil, nil, errors.New("no address for the oracle contract")
	}

	oracle, err := tellorX_oracle.NewOracle(addrOracle, client)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "creating contract interface")
	}

	addrGovern, err := master.Addresses(&bind.CallOpts{Context: ctx}, ethereum_t.Keccak256("_GOVERNANCE_CONTRACT"))
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "getting contract address")
	}
	if addrGovern == common.BytesToAddress([]byte("")) {
		return nil, nil, nil, errors.Wrap(err, "no address for the govern contract")
	}
	govern, err := tellorX_governance.NewGovernance(addrGovern, client)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "creating contract interface")
	}

	level.Info(logger).Log("msg", "contract addresses",
		"master", addrMaster,
		"oracle", addrOracle,
		"govern", addrGovern,
	)

	abiMaster, err := abi.JSON(strings.NewReader(tellorX_master.ControllerABI))
	if err != nil {
		return nil, nil, nil, errors.New("parsing contract ABI")
	}

	abiOracle, err := abi.JSON(strings.NewReader(tellorX_oracle.OracleABI))
	if err != nil {
		return nil, nil, nil, errors.New("parsing contract ABI")
	}

	abiGovern, err := abi.JSON(strings.NewReader(tellorX_governance.GovernanceABI))
	if err != nil {
		return nil, nil, nil, errors.New("parsing contract ABI")
	}

	boundMaster := bind.NewBoundContract(addrMaster, abiMaster, client, client, client)
	boundtOracle := bind.NewBoundContract(addrOracle, abiOracle, client, client, client)
	boundGovern := bind.NewBoundContract(addrGovern, abiGovern, client, client, client)

	return &TellorMaster{
			abi:           abiMaster,
			abiRaw:        tellorX_master.ControllerABI,
			address:       addrMaster,
			Controller:    master,
			BoundContract: boundMaster,
		}, &TellorOracle{
			BoundContract: boundtOracle,
			abi:           abiOracle,
			abiRaw:        tellorX_oracle.OracleABI,
			address:       addrOracle,
			Oracle:        oracle,
		}, &TellorGovern{
			BoundContract: boundGovern,
			abi:           abiGovern,
			abiRaw:        tellorX_governance.GovernanceABI,
			address:       addrGovern,
			Governance:    govern,
			params:        params,
		}, nil
}

func GetMasterAddress(netID int64) (common.Address, error) {
	switch netID {
	case 1:
		return common.HexToAddress(MasterAddress), nil
	case 4:
		return common.HexToAddress(MasterAddressRinkeby), nil
	case 5:
		return common.HexToAddress(MasterAddressGoerli), nil
	case 31337:
		return common.HexToAddress(MasterAddressHardhat), nil
	default:
		return common.Address{}, errors.Errorf("tellor contract address for current network id not found:%v", netID)
	}
}

type VoteLog struct {
	ID              int64
	Executed        bool
	IsDispute       bool
	ExecuteTimeLock time.Time
	TallyTs         int64
	StatusID        int64
	StatusName      string
	Initiator       common.Address
	VoteRound       int64
	VotesSupport    float64
	VotesAgainst    float64
	VotesInvalid    float64
	VoteEnds        time.Time
	VoteFunc        [4]byte
	Fee             float64
}

type DisputeLog struct {
	QueryID  [32]byte
	DataVal  float64
	DataTime time.Time
	Reporter common.Address
	TxHash   common.Hash
	VoteLog
}

func GetDisputeLogs(
	ctx context.Context,
	logger log.Logger,
	client ethereum_t.EthClient,
	contract TellorGovernCaller,
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
		log := tellorX_governance.GovernanceNewDispute{}
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
	client ethereum_t.EthClient,
	contract TellorGovernCaller,
	lookBackDuration time.Duration,
) ([]*tellorX_governance.GovernanceVoteTallied, error) {

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
		Topics:    [][]common.Hash{{contract.Abi().Events[EventNameTallyVote].ID}},
	}

	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "get logs")
	}

	var unpackedLogs []*tellorX_governance.GovernanceVoteTallied
	for _, logRaw := range logs {
		log := &tellorX_governance.GovernanceVoteTallied{}
		err := contract.UnpackLog(log, "VoteTallied", logRaw)
		if err != nil {
			return nil, errors.Wrap(err, "unpack event from logs")
		}
		log.Raw = logRaw
		unpackedLogs = append(unpackedLogs, log)

	}
	return unpackedLogs, err
}

func GetDisputeInfo(ctx context.Context, disputeID *big.Int, contract TellorGovernCaller) (*DisputeLog, error) {

	queryID, timestamp, val, reporter, err := contract.GetDisputeInfo(&bind.CallOpts{Context: ctx}, disputeID)
	if err != nil {
		return nil, errors.Wrap(err, "get dispute details")
	}

	voteInfo, err := GetVoteInfo(ctx, disputeID, contract)
	if err != nil {
		return nil, errors.Wrap(err, "get vote details")
	}

	psr, ok := psr_tellor.Psrs[queryID]
	if !ok {
		return nil, errors.Errorf("unable to get psr details for query ID:%v", queryID)
	}

	return &DisputeLog{
		QueryID:  queryID,
		DataTime: time.Unix(timestamp.Int64(), 0),
		DataVal:  math_t.BigIntToFloatDiv(big.NewInt(0).SetBytes(val), psr.Granularity),
		Reporter: reporter,
		VoteLog:  *voteInfo,
	}, nil
}

func GetVoteInfo(ctx context.Context, voteID *big.Int, contract TellorGovernCaller) (*VoteLog, error) {

	identifierHash, voteVars, statusVars, result, _, voteFunc, addrVars, err := contract.GetVoteInfo(&bind.CallOpts{Context: ctx}, voteID)
	if err != nil {
		return nil, errors.Wrap(err, "get vote details")
	}

	executed := statusVars[0]

	var statusName string
	var statusID int64
	if voteVars[4].Int64() != 0 {
		statusName = "tallied"
		statusID = VoteStatusTallied
	} else if !executed {
		statusName = "open"
		statusID = VoteStatusOpen
	}

	if executed {
		switch int(result) {
		case VoteStatusPassed:
			statusName = "passed"
			statusID = VoteStatusPassed
		case VoteStatusRejected:
			statusName = "rejected"
			statusID = VoteStatusRejected
		case VoteStatusInvalid:
			statusName = "invalid"
			statusID = VoteStatusInvalid
		}
	}

	if identifierHash == [32]byte{} {
		return nil, errors.Errorf("dispute doesn't exist id:%v", voteID)
	}

	return &VoteLog{
		ID:              voteID.Int64(),
		TallyTs:         voteVars[4].Int64(),
		Executed:        executed,
		IsDispute:       statusVars[1],
		ExecuteTimeLock: time.Unix(voteVars[4].Int64(), 0).Add(contract.VoteExecuteWaitDuration() * time.Duration(math_t.BigIntToFloat(voteVars[0]))),
		StatusID:        statusID,
		StatusName:      statusName,
		Initiator:       addrVars[1],
		VoteFunc:        voteFunc,
		VoteRound:       int64(math_t.BigIntToFloat(voteVars[0])),
		VotesSupport:    math_t.BigIntToFloatDiv(voteVars[5], params.Ether),
		VotesAgainst:    math_t.BigIntToFloatDiv(voteVars[6], params.Ether),
		VotesInvalid:    math_t.BigIntToFloatDiv(voteVars[7], params.Ether),
		VoteEnds:        time.Unix(voteVars[1].Int64(), 0).Add(contract.VotingDuration()),
		Fee:             math_t.BigIntToFloatDiv(voteVars[3], params.Ether),
	}, nil
}

type Report struct {
	tellorX_oracle.OracleNewReport
	Disputed bool
}

func GetSubmitLogs(
	ctx context.Context,
	client ethereum_t.EthClient,
	oracle TellorOracleCaller,
	from int64,
	lookBackDuration time.Duration,
) ([]Report, error) {
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
		Addresses: []common.Address{oracle.Addr()},
		Topics:    [][]common.Hash{{oracle.Abi().Events[EventNameNewSubmit].ID}},
	}

	events, err := client.FilterLogs(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "get events")
	}

	var reports []Report
	for _, eventRaw := range events {
		eventNewReport := &tellorX_oracle.OracleNewReport{}
		err := oracle.UnpackLog(eventNewReport, EventNameNewSubmit, eventRaw)
		if err != nil {
			return nil, errors.Wrap(err, "unpack event from logs")
		}

		eventNewReport.Raw = eventRaw

		val, err := oracle.GetValueByTimestamp(&bind.CallOpts{Context: ctx}, eventNewReport.QueryId, eventNewReport.Time)
		if err != nil {
			return nil, errors.Wrap(err, "getting the val to dispute")
		}
		disputed := false
		if len(val) == 0 {
			disputed = true
		}

		reports = append(reports, Report{Disputed: disputed, OracleNewReport: *eventNewReport})
	}

	return reports, err
}

func LastSubmit(ctx context.Context, contract TellorOracleCaller, reporter common.Address) (time.Duration, *time.Time, error) {
	last, err := contract.GetReporterLastTimestamp(&bind.CallOpts{Context: ctx}, reporter)
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

var StatusOnDispute = int64(3)

func ReporterStatusName(statusID int64) string {
	// From https://github.com/tellor-io/tellorX/blob/f63b260375cf119b2c7c0fd920f0ce9441ca06e8/contracts/tellor3/TellorStorage.sol#L28
	switch statusID {
	case 0:
		return "Not staked"
	case 1:
		return "Staked"
	case 2:
		return "LockedForWithdraw"
	case StatusOnDispute:
		return "OnDispute"
	case 4:
		return "ReadyForUnlocking"
	case 5:
		return "Unlocked"
	default:
		return "Unknown"
	}
}
