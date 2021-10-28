// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tellorX_oracle

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IControllerMetaData contains all meta data concerning the IController contract.
var IControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approveAndTransferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newController\",\"type\":\"address\"}],\"name\":\"changeControllerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newDeity\",\"type\":\"address\"}],\"name\":\"changeDeity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGovernance\",\"type\":\"address\"}],\"name\":\"changeGovernanceContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOracle\",\"type\":\"address\"}],\"name\":\"changeOracleContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"changeStakingStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tContract\",\"type\":\"address\"}],\"name\":\"changeTellorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newTreasury\",\"type\":\"address\"}],\"name\":\"changeTreasuryContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_target\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"changeUint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_c\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_r\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_d\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_t\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"idsOnDeck\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"tipsOnDeck\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reciever\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_disputer\",\"type\":\"address\"}],\"name\":\"slashReporter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tipQuery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"unlockDisputeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"699f200f": "addresses(bytes32)",
		"dd62ed3e": "allowance(address,address)",
		"999cf26c": "allowedToTrade(address,uint256)",
		"095ea7b3": "approve(address,uint256)",
		"288c9c9d": "approveAndTransferFrom(address,address,uint256)",
		"70a08231": "balanceOf(address)",
		"4ee2cd7e": "balanceOfAt(address,uint256)",
		"8581af19": "beginDispute(uint256,uint256,uint256)",
		"42966c68": "burn(uint256)",
		"3c46a185": "changeControllerContract(address)",
		"47abd7f1": "changeDeity(address)",
		"e8ce51d7": "changeGovernanceContract(address)",
		"1cbd3151": "changeOracleContract(address)",
		"a6f9dae1": "changeOwner(address)",
		"a1332c5c": "changeStakingStatus(address,uint256)",
		"ae0a8279": "changeTellorContract(address)",
		"bd87e0c9": "changeTreasuryContract(address)",
		"740358e6": "changeUint(bytes32,uint256)",
		"313ce567": "decimals()",
		"0d2d76a2": "depositStake()",
		"133bee5e": "getAddressVars(bytes32)",
		"da379941": "getDisputeIdByDisputeHash(bytes32)",
		"3180f8df": "getLastNewValueById(uint256)",
		"4049f198": "getNewCurrentVariables()",
		"46eee1c4": "getNewValueCountbyRequestId(uint256)",
		"9a7077ab": "getNewVariablesOnDeck()",
		"733bdef0": "getStakerInfo(address)",
		"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
		"612c8f7f": "getUintVar(bytes32)",
		"e1c7392a": "init()",
		"8fd3ab80": "migrate()",
		"40c10f19": "mint(address,uint256)",
		"06fdde03": "name()",
		"28449c3a": "requestStakingWithdraw()",
		"93fa4915": "retrieveData(uint256,uint256)",
		"4dfc2a34": "slashReporter(address,address)",
		"95d89b41": "symbol()",
		"4d318b0e": "tallyVotes(uint256)",
		"2227dff8": "tipQuery(uint256,uint256,bytes)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"b59e14d4": "uints(bytes32)",
		"9a01ca13": "unlockDisputeFee(uint256)",
		"c9d27afe": "vote(uint256,bool)",
		"bed9d861": "withdrawStake()",
	},
}

// IControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use IControllerMetaData.ABI instead.
var IControllerABI = IControllerMetaData.ABI

// Deprecated: Use IControllerMetaData.Sigs instead.
// IControllerFuncSigs maps the 4-byte function signature to its string representation.
var IControllerFuncSigs = IControllerMetaData.Sigs

// IController is an auto generated Go binding around an Ethereum contract.
type IController struct {
	IControllerCaller     // Read-only binding to the contract
	IControllerTransactor // Write-only binding to the contract
	IControllerFilterer   // Log filterer for contract events
}

// IControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IControllerSession struct {
	Contract     *IController      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IControllerCallerSession struct {
	Contract *IControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IControllerTransactorSession struct {
	Contract     *IControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IControllerRaw struct {
	Contract *IController // Generic contract binding to access the raw methods on
}

// IControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IControllerCallerRaw struct {
	Contract *IControllerCaller // Generic read-only contract binding to access the raw methods on
}

// IControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IControllerTransactorRaw struct {
	Contract *IControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIController creates a new instance of IController, bound to a specific deployed contract.
func NewIController(address common.Address, backend bind.ContractBackend) (*IController, error) {
	contract, err := bindIController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IController{IControllerCaller: IControllerCaller{contract: contract}, IControllerTransactor: IControllerTransactor{contract: contract}, IControllerFilterer: IControllerFilterer{contract: contract}}, nil
}

// NewIControllerCaller creates a new read-only instance of IController, bound to a specific deployed contract.
func NewIControllerCaller(address common.Address, caller bind.ContractCaller) (*IControllerCaller, error) {
	contract, err := bindIController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IControllerCaller{contract: contract}, nil
}

// NewIControllerTransactor creates a new write-only instance of IController, bound to a specific deployed contract.
func NewIControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*IControllerTransactor, error) {
	contract, err := bindIController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IControllerTransactor{contract: contract}, nil
}

// NewIControllerFilterer creates a new log filterer instance of IController, bound to a specific deployed contract.
func NewIControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*IControllerFilterer, error) {
	contract, err := bindIController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IControllerFilterer{contract: contract}, nil
}

// bindIController binds a generic wrapper to an already deployed contract.
func bindIController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IController *IControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IController.Contract.IControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IController *IControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IController.Contract.IControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IController *IControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IController.Contract.IControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IController *IControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IController *IControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IController *IControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IController.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_IController *IControllerCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "allowance", _user, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_IController *IControllerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _IController.Contract.Allowance(&_IController.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_IController *IControllerCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _IController.Contract.Allowance(&_IController.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_IController *IControllerCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "allowedToTrade", _user, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_IController *IControllerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _IController.Contract.AllowedToTrade(&_IController.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_IController *IControllerCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _IController.Contract.AllowedToTrade(&_IController.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_IController *IControllerCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_IController *IControllerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _IController.Contract.BalanceOf(&_IController.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_IController *IControllerCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _IController.Contract.BalanceOf(&_IController.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_IController *IControllerCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "balanceOfAt", _user, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_IController *IControllerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _IController.Contract.BalanceOfAt(&_IController.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_IController *IControllerCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _IController.Contract.BalanceOfAt(&_IController.CallOpts, _user, _blockNumber)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IController *IControllerCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IController *IControllerSession) Decimals() (uint8, error) {
	return _IController.Contract.Decimals(&_IController.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IController *IControllerCallerSession) Decimals() (uint8, error) {
	return _IController.Contract.Decimals(&_IController.CallOpts)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_IController *IControllerCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "getAddressVars", _data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_IController *IControllerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _IController.Contract.GetAddressVars(&_IController.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_IController *IControllerCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _IController.Contract.GetAddressVars(&_IController.CallOpts, _data)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_IController *IControllerCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "getDisputeIdByDisputeHash", _hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_IController *IControllerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _IController.Contract.GetDisputeIdByDisputeHash(&_IController.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_IController *IControllerCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _IController.Contract.GetDisputeIdByDisputeHash(&_IController.CallOpts, _hash)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_IController *IControllerCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "getLastNewValueById", _requestId)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_IController *IControllerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _IController.Contract.GetLastNewValueById(&_IController.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_IController *IControllerCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _IController.Contract.GetLastNewValueById(&_IController.CallOpts, _requestId)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _c, uint256[5] _r, uint256 _d, uint256 _t)
func (_IController *IControllerCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	C [32]byte
	R [5]*big.Int
	D *big.Int
	T *big.Int
}, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "getNewCurrentVariables")

	outstruct := new(struct {
		C [32]byte
		R [5]*big.Int
		D *big.Int
		T *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.C = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.R = *abi.ConvertType(out[1], new([5]*big.Int)).(*[5]*big.Int)
	outstruct.D = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.T = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _c, uint256[5] _r, uint256 _d, uint256 _t)
func (_IController *IControllerSession) GetNewCurrentVariables() (struct {
	C [32]byte
	R [5]*big.Int
	D *big.Int
	T *big.Int
}, error) {
	return _IController.Contract.GetNewCurrentVariables(&_IController.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _c, uint256[5] _r, uint256 _d, uint256 _t)
func (_IController *IControllerCallerSession) GetNewCurrentVariables() (struct {
	C [32]byte
	R [5]*big.Int
	D *big.Int
	T *big.Int
}, error) {
	return _IController.Contract.GetNewCurrentVariables(&_IController.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_IController *IControllerCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_IController *IControllerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _IController.Contract.GetNewValueCountbyRequestId(&_IController.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_IController *IControllerCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _IController.Contract.GetNewValueCountbyRequestId(&_IController.CallOpts, _requestId)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_IController *IControllerCaller) GetNewVariablesOnDeck(opts *bind.CallOpts) (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "getNewVariablesOnDeck")

	outstruct := new(struct {
		IdsOnDeck  [5]*big.Int
		TipsOnDeck [5]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IdsOnDeck = *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)
	outstruct.TipsOnDeck = *abi.ConvertType(out[1], new([5]*big.Int)).(*[5]*big.Int)

	return *outstruct, err

}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_IController *IControllerSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _IController.Contract.GetNewVariablesOnDeck(&_IController.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_IController *IControllerCallerSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _IController.Contract.GetNewVariablesOnDeck(&_IController.CallOpts)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_IController *IControllerCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "getStakerInfo", _staker)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_IController *IControllerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _IController.Contract.GetStakerInfo(&_IController.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_IController *IControllerCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _IController.Contract.GetStakerInfo(&_IController.CallOpts, _staker)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_IController *IControllerCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestID, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_IController *IControllerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _IController.Contract.GetTimestampbyRequestIDandIndex(&_IController.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_IController *IControllerCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _IController.Contract.GetTimestampbyRequestIDandIndex(&_IController.CallOpts, _requestID, _index)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_IController *IControllerCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "getUintVar", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_IController *IControllerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _IController.Contract.GetUintVar(&_IController.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_IController *IControllerCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _IController.Contract.GetUintVar(&_IController.CallOpts, _data)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IController *IControllerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IController *IControllerSession) Name() (string, error) {
	return _IController.Contract.Name(&_IController.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IController *IControllerCallerSession) Name() (string, error) {
	return _IController.Contract.Name(&_IController.CallOpts)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_IController *IControllerCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_IController *IControllerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _IController.Contract.RetrieveData(&_IController.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_IController *IControllerCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _IController.Contract.RetrieveData(&_IController.CallOpts, _requestId, _timestamp)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IController *IControllerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IController *IControllerSession) Symbol() (string, error) {
	return _IController.Contract.Symbol(&_IController.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IController *IControllerCallerSession) Symbol() (string, error) {
	return _IController.Contract.Symbol(&_IController.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IController *IControllerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IController.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IController *IControllerSession) TotalSupply() (*big.Int, error) {
	return _IController.Contract.TotalSupply(&_IController.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IController *IControllerCallerSession) TotalSupply() (*big.Int, error) {
	return _IController.Contract.TotalSupply(&_IController.CallOpts)
}

// Addresses is a paid mutator transaction binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) returns(address)
func (_IController *IControllerTransactor) Addresses(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "addresses", arg0)
}

// Addresses is a paid mutator transaction binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) returns(address)
func (_IController *IControllerSession) Addresses(arg0 [32]byte) (*types.Transaction, error) {
	return _IController.Contract.Addresses(&_IController.TransactOpts, arg0)
}

// Addresses is a paid mutator transaction binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) returns(address)
func (_IController *IControllerTransactorSession) Addresses(arg0 [32]byte) (*types.Transaction, error) {
	return _IController.Contract.Addresses(&_IController.TransactOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_IController *IControllerTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_IController *IControllerSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.Contract.Approve(&_IController.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_IController *IControllerTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.Contract.Approve(&_IController.TransactOpts, _spender, _amount)
}

// ApproveAndTransferFrom is a paid mutator transaction binding the contract method 0x288c9c9d.
//
// Solidity: function approveAndTransferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_IController *IControllerTransactor) ApproveAndTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "approveAndTransferFrom", _from, _to, _amount)
}

// ApproveAndTransferFrom is a paid mutator transaction binding the contract method 0x288c9c9d.
//
// Solidity: function approveAndTransferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_IController *IControllerSession) ApproveAndTransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.Contract.ApproveAndTransferFrom(&_IController.TransactOpts, _from, _to, _amount)
}

// ApproveAndTransferFrom is a paid mutator transaction binding the contract method 0x288c9c9d.
//
// Solidity: function approveAndTransferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_IController *IControllerTransactorSession) ApproveAndTransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.Contract.ApproveAndTransferFrom(&_IController.TransactOpts, _from, _to, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_IController *IControllerTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_IController *IControllerSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _IController.Contract.BeginDispute(&_IController.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_IController *IControllerTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _IController.Contract.BeginDispute(&_IController.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_IController *IControllerTransactor) Burn(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "burn", _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_IController *IControllerSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _IController.Contract.Burn(&_IController.TransactOpts, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_IController *IControllerTransactorSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _IController.Contract.Burn(&_IController.TransactOpts, _amount)
}

// ChangeControllerContract is a paid mutator transaction binding the contract method 0x3c46a185.
//
// Solidity: function changeControllerContract(address _newController) returns()
func (_IController *IControllerTransactor) ChangeControllerContract(opts *bind.TransactOpts, _newController common.Address) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "changeControllerContract", _newController)
}

// ChangeControllerContract is a paid mutator transaction binding the contract method 0x3c46a185.
//
// Solidity: function changeControllerContract(address _newController) returns()
func (_IController *IControllerSession) ChangeControllerContract(_newController common.Address) (*types.Transaction, error) {
	return _IController.Contract.ChangeControllerContract(&_IController.TransactOpts, _newController)
}

// ChangeControllerContract is a paid mutator transaction binding the contract method 0x3c46a185.
//
// Solidity: function changeControllerContract(address _newController) returns()
func (_IController *IControllerTransactorSession) ChangeControllerContract(_newController common.Address) (*types.Transaction, error) {
	return _IController.Contract.ChangeControllerContract(&_IController.TransactOpts, _newController)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_IController *IControllerTransactor) ChangeDeity(opts *bind.TransactOpts, _newDeity common.Address) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "changeDeity", _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_IController *IControllerSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _IController.Contract.ChangeDeity(&_IController.TransactOpts, _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_IController *IControllerTransactorSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _IController.Contract.ChangeDeity(&_IController.TransactOpts, _newDeity)
}

// ChangeGovernanceContract is a paid mutator transaction binding the contract method 0xe8ce51d7.
//
// Solidity: function changeGovernanceContract(address _newGovernance) returns()
func (_IController *IControllerTransactor) ChangeGovernanceContract(opts *bind.TransactOpts, _newGovernance common.Address) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "changeGovernanceContract", _newGovernance)
}

// ChangeGovernanceContract is a paid mutator transaction binding the contract method 0xe8ce51d7.
//
// Solidity: function changeGovernanceContract(address _newGovernance) returns()
func (_IController *IControllerSession) ChangeGovernanceContract(_newGovernance common.Address) (*types.Transaction, error) {
	return _IController.Contract.ChangeGovernanceContract(&_IController.TransactOpts, _newGovernance)
}

// ChangeGovernanceContract is a paid mutator transaction binding the contract method 0xe8ce51d7.
//
// Solidity: function changeGovernanceContract(address _newGovernance) returns()
func (_IController *IControllerTransactorSession) ChangeGovernanceContract(_newGovernance common.Address) (*types.Transaction, error) {
	return _IController.Contract.ChangeGovernanceContract(&_IController.TransactOpts, _newGovernance)
}

// ChangeOracleContract is a paid mutator transaction binding the contract method 0x1cbd3151.
//
// Solidity: function changeOracleContract(address _newOracle) returns()
func (_IController *IControllerTransactor) ChangeOracleContract(opts *bind.TransactOpts, _newOracle common.Address) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "changeOracleContract", _newOracle)
}

// ChangeOracleContract is a paid mutator transaction binding the contract method 0x1cbd3151.
//
// Solidity: function changeOracleContract(address _newOracle) returns()
func (_IController *IControllerSession) ChangeOracleContract(_newOracle common.Address) (*types.Transaction, error) {
	return _IController.Contract.ChangeOracleContract(&_IController.TransactOpts, _newOracle)
}

// ChangeOracleContract is a paid mutator transaction binding the contract method 0x1cbd3151.
//
// Solidity: function changeOracleContract(address _newOracle) returns()
func (_IController *IControllerTransactorSession) ChangeOracleContract(_newOracle common.Address) (*types.Transaction, error) {
	return _IController.Contract.ChangeOracleContract(&_IController.TransactOpts, _newOracle)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_IController *IControllerTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_IController *IControllerSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _IController.Contract.ChangeOwner(&_IController.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_IController *IControllerTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _IController.Contract.ChangeOwner(&_IController.TransactOpts, _newOwner)
}

// ChangeStakingStatus is a paid mutator transaction binding the contract method 0xa1332c5c.
//
// Solidity: function changeStakingStatus(address _reporter, uint256 _status) returns()
func (_IController *IControllerTransactor) ChangeStakingStatus(opts *bind.TransactOpts, _reporter common.Address, _status *big.Int) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "changeStakingStatus", _reporter, _status)
}

// ChangeStakingStatus is a paid mutator transaction binding the contract method 0xa1332c5c.
//
// Solidity: function changeStakingStatus(address _reporter, uint256 _status) returns()
func (_IController *IControllerSession) ChangeStakingStatus(_reporter common.Address, _status *big.Int) (*types.Transaction, error) {
	return _IController.Contract.ChangeStakingStatus(&_IController.TransactOpts, _reporter, _status)
}

// ChangeStakingStatus is a paid mutator transaction binding the contract method 0xa1332c5c.
//
// Solidity: function changeStakingStatus(address _reporter, uint256 _status) returns()
func (_IController *IControllerTransactorSession) ChangeStakingStatus(_reporter common.Address, _status *big.Int) (*types.Transaction, error) {
	return _IController.Contract.ChangeStakingStatus(&_IController.TransactOpts, _reporter, _status)
}

// ChangeTellorContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeTellorContract(address _tContract) returns()
func (_IController *IControllerTransactor) ChangeTellorContract(opts *bind.TransactOpts, _tContract common.Address) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "changeTellorContract", _tContract)
}

// ChangeTellorContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeTellorContract(address _tContract) returns()
func (_IController *IControllerSession) ChangeTellorContract(_tContract common.Address) (*types.Transaction, error) {
	return _IController.Contract.ChangeTellorContract(&_IController.TransactOpts, _tContract)
}

// ChangeTellorContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeTellorContract(address _tContract) returns()
func (_IController *IControllerTransactorSession) ChangeTellorContract(_tContract common.Address) (*types.Transaction, error) {
	return _IController.Contract.ChangeTellorContract(&_IController.TransactOpts, _tContract)
}

// ChangeTreasuryContract is a paid mutator transaction binding the contract method 0xbd87e0c9.
//
// Solidity: function changeTreasuryContract(address _newTreasury) returns()
func (_IController *IControllerTransactor) ChangeTreasuryContract(opts *bind.TransactOpts, _newTreasury common.Address) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "changeTreasuryContract", _newTreasury)
}

// ChangeTreasuryContract is a paid mutator transaction binding the contract method 0xbd87e0c9.
//
// Solidity: function changeTreasuryContract(address _newTreasury) returns()
func (_IController *IControllerSession) ChangeTreasuryContract(_newTreasury common.Address) (*types.Transaction, error) {
	return _IController.Contract.ChangeTreasuryContract(&_IController.TransactOpts, _newTreasury)
}

// ChangeTreasuryContract is a paid mutator transaction binding the contract method 0xbd87e0c9.
//
// Solidity: function changeTreasuryContract(address _newTreasury) returns()
func (_IController *IControllerTransactorSession) ChangeTreasuryContract(_newTreasury common.Address) (*types.Transaction, error) {
	return _IController.Contract.ChangeTreasuryContract(&_IController.TransactOpts, _newTreasury)
}

// ChangeUint is a paid mutator transaction binding the contract method 0x740358e6.
//
// Solidity: function changeUint(bytes32 _target, uint256 _amount) returns()
func (_IController *IControllerTransactor) ChangeUint(opts *bind.TransactOpts, _target [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "changeUint", _target, _amount)
}

// ChangeUint is a paid mutator transaction binding the contract method 0x740358e6.
//
// Solidity: function changeUint(bytes32 _target, uint256 _amount) returns()
func (_IController *IControllerSession) ChangeUint(_target [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _IController.Contract.ChangeUint(&_IController.TransactOpts, _target, _amount)
}

// ChangeUint is a paid mutator transaction binding the contract method 0x740358e6.
//
// Solidity: function changeUint(bytes32 _target, uint256 _amount) returns()
func (_IController *IControllerTransactorSession) ChangeUint(_target [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _IController.Contract.ChangeUint(&_IController.TransactOpts, _target, _amount)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_IController *IControllerTransactor) DepositStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "depositStake")
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_IController *IControllerSession) DepositStake() (*types.Transaction, error) {
	return _IController.Contract.DepositStake(&_IController.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_IController *IControllerTransactorSession) DepositStake() (*types.Transaction, error) {
	return _IController.Contract.DepositStake(&_IController.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_IController *IControllerTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_IController *IControllerSession) Init() (*types.Transaction, error) {
	return _IController.Contract.Init(&_IController.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_IController *IControllerTransactorSession) Init() (*types.Transaction, error) {
	return _IController.Contract.Init(&_IController.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_IController *IControllerTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_IController *IControllerSession) Migrate() (*types.Transaction, error) {
	return _IController.Contract.Migrate(&_IController.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_IController *IControllerTransactorSession) Migrate() (*types.Transaction, error) {
	return _IController.Contract.Migrate(&_IController.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _reciever, uint256 _amount) returns()
func (_IController *IControllerTransactor) Mint(opts *bind.TransactOpts, _reciever common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "mint", _reciever, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _reciever, uint256 _amount) returns()
func (_IController *IControllerSession) Mint(_reciever common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.Contract.Mint(&_IController.TransactOpts, _reciever, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _reciever, uint256 _amount) returns()
func (_IController *IControllerTransactorSession) Mint(_reciever common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.Contract.Mint(&_IController.TransactOpts, _reciever, _amount)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_IController *IControllerTransactor) RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "requestStakingWithdraw")
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_IController *IControllerSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _IController.Contract.RequestStakingWithdraw(&_IController.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_IController *IControllerTransactorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _IController.Contract.RequestStakingWithdraw(&_IController.TransactOpts)
}

// SlashReporter is a paid mutator transaction binding the contract method 0x4dfc2a34.
//
// Solidity: function slashReporter(address _reporter, address _disputer) returns()
func (_IController *IControllerTransactor) SlashReporter(opts *bind.TransactOpts, _reporter common.Address, _disputer common.Address) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "slashReporter", _reporter, _disputer)
}

// SlashReporter is a paid mutator transaction binding the contract method 0x4dfc2a34.
//
// Solidity: function slashReporter(address _reporter, address _disputer) returns()
func (_IController *IControllerSession) SlashReporter(_reporter common.Address, _disputer common.Address) (*types.Transaction, error) {
	return _IController.Contract.SlashReporter(&_IController.TransactOpts, _reporter, _disputer)
}

// SlashReporter is a paid mutator transaction binding the contract method 0x4dfc2a34.
//
// Solidity: function slashReporter(address _reporter, address _disputer) returns()
func (_IController *IControllerTransactorSession) SlashReporter(_reporter common.Address, _disputer common.Address) (*types.Transaction, error) {
	return _IController.Contract.SlashReporter(&_IController.TransactOpts, _reporter, _disputer)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_IController *IControllerTransactor) TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "tallyVotes", _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_IController *IControllerSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _IController.Contract.TallyVotes(&_IController.TransactOpts, _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_IController *IControllerTransactorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _IController.Contract.TallyVotes(&_IController.TransactOpts, _disputeId)
}

// TipQuery is a paid mutator transaction binding the contract method 0x2227dff8.
//
// Solidity: function tipQuery(uint256 , uint256 , bytes ) returns()
func (_IController *IControllerTransactor) TipQuery(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "tipQuery", arg0, arg1, arg2)
}

// TipQuery is a paid mutator transaction binding the contract method 0x2227dff8.
//
// Solidity: function tipQuery(uint256 , uint256 , bytes ) returns()
func (_IController *IControllerSession) TipQuery(arg0 *big.Int, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _IController.Contract.TipQuery(&_IController.TransactOpts, arg0, arg1, arg2)
}

// TipQuery is a paid mutator transaction binding the contract method 0x2227dff8.
//
// Solidity: function tipQuery(uint256 , uint256 , bytes ) returns()
func (_IController *IControllerTransactorSession) TipQuery(arg0 *big.Int, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _IController.Contract.TipQuery(&_IController.TransactOpts, arg0, arg1, arg2)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_IController *IControllerTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_IController *IControllerSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.Contract.Transfer(&_IController.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_IController *IControllerTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.Contract.Transfer(&_IController.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_IController *IControllerTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_IController *IControllerSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.Contract.TransferFrom(&_IController.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_IController *IControllerTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IController.Contract.TransferFrom(&_IController.TransactOpts, _from, _to, _amount)
}

// Uints is a paid mutator transaction binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) returns(uint256)
func (_IController *IControllerTransactor) Uints(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "uints", arg0)
}

// Uints is a paid mutator transaction binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) returns(uint256)
func (_IController *IControllerSession) Uints(arg0 [32]byte) (*types.Transaction, error) {
	return _IController.Contract.Uints(&_IController.TransactOpts, arg0)
}

// Uints is a paid mutator transaction binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) returns(uint256)
func (_IController *IControllerTransactorSession) Uints(arg0 [32]byte) (*types.Transaction, error) {
	return _IController.Contract.Uints(&_IController.TransactOpts, arg0)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_IController *IControllerTransactor) UnlockDisputeFee(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "unlockDisputeFee", _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_IController *IControllerSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _IController.Contract.UnlockDisputeFee(&_IController.TransactOpts, _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_IController *IControllerTransactorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _IController.Contract.UnlockDisputeFee(&_IController.TransactOpts, _disputeId)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_IController *IControllerTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_IController *IControllerSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _IController.Contract.Vote(&_IController.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_IController *IControllerTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _IController.Contract.Vote(&_IController.TransactOpts, _disputeId, _supportsDispute)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_IController *IControllerTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IController.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_IController *IControllerSession) WithdrawStake() (*types.Transaction, error) {
	return _IController.Contract.WithdrawStake(&_IController.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_IController *IControllerTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _IController.Contract.WithdrawStake(&_IController.TransactOpts)
}

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newMiningLock\",\"type\":\"uint256\"}],\"name\":\"MiningLockChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_queryData\",\"type\":\"bytes\"}],\"name\":\"NewReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newTimeBasedReward\",\"type\":\"uint256\"}],\"name\":\"TimeBasedRewardsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_queryData\",\"type\":\"bytes\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMiningLock\",\"type\":\"uint256\"}],\"name\":\"changeMiningLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newTimeBasedReward\",\"type\":\"uint256\"}],\"name\":\"changeTimeBasedReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getBlockNumberByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"}],\"name\":\"getCurrentReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"}],\"name\":\"getCurrentValue\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMiningLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReportTimestampByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getReporterByTimestamp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter\",\"type\":\"address\"}],\"name\":\"getReportsSubmittedByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeBasedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeOfLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"}],\"name\":\"getTimestampCountById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getTimestampIndexByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"}],\"name\":\"getTipsById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getTipsByUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getValueByTimestamp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miningLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"removeValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_queryData\",\"type\":\"bytes\"}],\"name\":\"submitValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeBasedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeOfLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_queryData\",\"type\":\"bytes\"}],\"name\":\"tipQuery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tipsInContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e280e8e8": "changeMiningLock(uint256)",
		"6d53585f": "changeTimeBasedReward(uint256)",
		"935408d0": "getBlockNumberByTimestamp(bytes32,uint256)",
		"a1e588a5": "getCurrentReward(bytes32)",
		"adf1639d": "getCurrentValue(bytes32)",
		"bf1b032e": "getMiningLock()",
		"7c37b8b4": "getReportTimestampByIndex(bytes32,uint256)",
		"e07c5486": "getReporterByTimestamp(bytes32,uint256)",
		"3878293e": "getReportsSubmittedByAddress(address)",
		"14d66b9a": "getTimeBasedReward()",
		"c0f95d52": "getTimeOfLastNewValue()",
		"35e72432": "getTimestampCountById(bytes32)",
		"9d9b16ed": "getTimestampIndexByTimestamp(bytes32,uint256)",
		"ef4c262d": "getTipsById(bytes32)",
		"b736ec36": "getTipsByUser(address)",
		"0b2d2b0d": "getValueByTimestamp(bytes32,uint256)",
		"d804b663": "miningLock()",
		"5b5edcfc": "removeValue(bytes32,uint256)",
		"5eaa9ced": "submitValue(bytes32,bytes,uint256,bytes)",
		"96426d97": "timeBasedReward()",
		"6fd4f229": "timeOfLastNewValue()",
		"ef0234ad": "tipQuery(bytes32,uint256,bytes)",
		"602bf227": "tips(bytes32)",
		"69d43bd3": "tipsInContract()",
		"fc735e99": "verify()",
	},
	Bin: "0x60806040524260015561a8c06002556706f05b59d3b2000060035534801561002657600080fd5b50611d46806100366000396000f3fe608060405234801561001057600080fd5b50600436106101735760003560e01c806396426d97116100de578063c0f95d5211610097578063e280e8e811610071578063e280e8e8146103b0578063ef0234ad146103c3578063ef4c262d146103d6578063fc735e99146103f657610173565b8063c0f95d521461034f578063d804b66314610357578063e07c54861461036057610173565b806396426d97146102ab5780639d9b16ed146102b4578063a1e588a5146102e3578063adf1639d1461030b578063b736ec361461031e578063bf1b032e1461034757610173565b8063602bf22711610130578063602bf2271461022457806369d43bd3146102445780636d53585f1461024d5780636fd4f229146102605780637c37b8b414610269578063935408d01461027c57610173565b80630b2d2b0d1461017857806314d66b9a146101a157806335e72432146101b35780633878293e146101d35780635b5edcfc146101fc5780635eaa9ced14610211575b600080fd5b61018b610186366004611a55565b6103fe565b6040516101989190611bad565b60405180910390f35b6003545b604051908152602001610198565b6101a56101c1366004611998565b60009081526005602052604090205490565b6101a56101e1366004611939565b6001600160a01b031660009081526007602052604090205490565b61020f61020a366004611a55565b6104af565b005b61020f61021f3660046119b0565b610841565b6101a5610232366004611998565b60046020526000908152604090205481565b6101a560005481565b61020f61025b366004611998565b610ecb565b6101a560015481565b6101a5610277366004611a55565b611018565b6101a561028a366004611a55565b60009182526005602090815260408084209284526002909201905290205490565b6101a560035481565b6101a56102c2366004611a55565b60009182526005602090815260408084209284526001909201905290205490565b6102f66102f1366004611998565b611059565b60408051928352602083019190915201610198565b61018b610319366004611998565b6111d0565b6101a561032c366004611939565b6001600160a01b031660009081526008602052604090205490565b6002546101a5565b6001546101a5565b6101a560025481565b61039861036e366004611a55565b6000918252600560209081526040808420928452600490920190529020546001600160a01b031690565b6040516001600160a01b039091168152602001610198565b61020f6103be366004611998565b6112bd565b61020f6103d1366004611a76565b61144f565b6101a56103e4366004611998565b60009081526004602052604090205490565b61270f6101a5565b6000828152600560209081526040808320848452600301909152902080546060919061042990611c56565b80601f016020809104026020016040519081016040528092919081815260200182805461045590611c56565b80156104a25780601f10610477576101008083540402835291602001916104a2565b820191906000526020600020905b81548152906001019060200180831161048557829003601f168201915b5050505050905092915050565b60405163699f200f60e01b8152600080516020611cf183398151915260048201527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a09063699f200f90602401602060405180830381600087803b15801561050957600080fd5b505af115801561051d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610541919061195c565b6001600160a01b0316336001600160a01b03161480610614575060405163699f200f60e01b81527ffa522e460446113e8fd353d7fa015625a68bc0369712213a42e006346440891e60048201527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a09063699f200f90602401602060405180830381600087803b1580156105c757600080fd5b505af11580156105db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ff919061195c565b6001600160a01b0316336001600160a01b0316145b61068b5760405162461bcd60e51b815260206004820152603d60248201527f63616c6c6572206d7573742062652074686520676f7665726e616e636520636f60448201527f6e7472616374206f7220746865206f7261636c6520636f6e747261637400000060648201526084015b60405180910390fd5b60008281526005602090815260408083208484526001810190925290912054805b82546106ba90600190611c3f565b81101561079257826106cd826001611be8565b815481106106eb57634e487b7160e01b600052603260045260246000fd5b906000526020600020015483600001828154811061071957634e487b7160e01b600052603260045260246000fd5b9060005260206000200181905550600183600101600085600001848154811061075257634e487b7160e01b600052603260045260246000fd5b90600052602060002001548152602001908152602001600020600082825461077a9190611c3f565b9091555081905061078a81611c91565b9150506106ac565b50815482906107a390600190611c3f565b815481106107c157634e487b7160e01b600052603260045260246000fd5b600091825260208220015581548290806107eb57634e487b7160e01b600052603160045260246000fd5b60008281526020808220830160001990810183905590920190925560408051808301808352848252878552600387019093529220915161082c9291906117a5565b50506000918252600101602052604081205550565b60008581526005602052604090208054831461089f5760405162461bcd60e51b815260206004820181905260248201527f6e6f6e6365206d757374206d617463682074696d657374616d7020696e6465786044820152606401610682565b600254336000908152600660205260409020546108bc9042611c3f565b1161091b5760405162461bcd60e51b815260206004820152602960248201527f7374696c6c20696e207265706f727465722074696d65206c6f636b2c20706c6560448201526861736520776169742160b81b6064820152608401610682565b60405163699f200f60e01b81527ffa522e460446113e8fd353d7fa015625a68bc0369712213a42e006346440891e60048201527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a09063699f200f90602401602060405180830381600087803b15801561098757600080fd5b505af115801561099b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109bf919061195c565b6001600160a01b0316306001600160a01b0316146109dc57600080fd5b815160208301208614806109f1575060648611155b610a3d5760405162461bcd60e51b815260206004820152601d60248201527f6964206d7573742062652068617368206f6620627974657320646174610000006044820152606401610682565b3360008181526006602052604080822042905551630733bdef60e41b815260048101929092527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a091829063733bdef090602401604080518083038186803b158015610a9b57600080fd5b505afa158015610aaf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad39190611adc565b50905080600114610b265760405162461bcd60e51b815260206004820152601d60248201527f5265706f7274657220737461747573206973206e6f74207374616b65720000006044820152606401610682565b604051632d67853560e21b81527f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc9760048201526001600160a01b0383169063b59e14d490602401602060405180830381600087803b158015610b8757600080fd5b505af1158015610b9b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bbf9190611ac4565b6040516370a0823160e01b81523360048201526001600160a01b038416906370a082319060240160206040518083038186803b158015610bfe57600080fd5b505afa158015610c12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c369190611ac4565b1015610c965760405162461bcd60e51b815260206004820152602960248201527f62616c616e6365206d7573742062652067726561746572207468616e207374616044820152681ad948185b5bdd5b9d60ba1b6064820152608401610682565b4260009081526004840160205260409020546001600160a01b031615610cfe5760405162461bcd60e51b815260206004820152601e60248201527f74696d657374616d7020616c7265616479207265706f7274656420666f7200006044820152606401610682565b82544260008181526001808701602090815260408084208690559185018855878352808320909401839055918152600286018352818120439055600386019092529020610d4c908888611829565b50426000908152600484016020526040812080546001600160a01b0319163317905580610d788a611059565b9150915081600080828254610d8d9190611c3f565b9091555060009050610d9f8383611be8565b1115610e3f576001600160a01b03841663a9059cbb33610dbf8585611be8565b6040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604401602060405180830381600087803b158015610e0557600080fd5b505af1158015610e19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e3d9190611978565b505b60008a81526004602090815260408083208390554260015533835260079091528120805491610e6d83611c91565b909155507fb5e84abaa3916191eacb4f925fb0d61505d1485414b50288e5bb52dd63c75cb790508a428b8b610ea28688611be8565b8c8c604051610eb79796959493929190611b4a565b60405180910390a150505050505050505050565b60405163699f200f60e01b8152600080516020611cf183398151915260048201527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a09063699f200f90602401602060405180830381600087803b158015610f2557600080fd5b505af1158015610f39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f5d919061195c565b6001600160a01b0316336001600160a01b031614610fdc5760405162461bcd60e51b815260206004820152603660248201527f4f6e6c7920676f7665726e616e636520636f6e74726163742063616e206368616044820152753733b2903a34b6b2903130b9b2b2103932bbb0b9321760511b6064820152608401610682565b60038190556040518181527f01432c5d54689da3dbfec924c89d1e453ecaec349928ceabd50207e909e8f7de906020015b60405180910390a150565b600082815260056020526040812080548390811061104657634e487b7160e01b600052603260045260246000fd5b9060005260206000200154905092915050565b60008060007388df592f8eb5d7bd38bfef7deb0fbc02cf3778a090506000600154426110859190611c3f565b9050600061012c6003548361109a9190611c20565b6110a49190611c00565b9050600054816110b49190611be8565b6040516370a0823160e01b81523060048201526001600160a01b038516906370a082319060240160206040518083038186803b1580156110f357600080fd5b505afa158015611107573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061112b9190611ac4565b10156111b8576000546040516370a0823160e01b81523060048201526001600160a01b038516906370a082319060240160206040518083038186803b15801561117357600080fd5b505afa158015611187573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111ab9190611ac4565b6111b59190611c3f565b90505b60009586526004602052604090952054959350505050565b60008181526005602052604081208054606092600383019290916111f690600190611c3f565b8154811061121457634e487b7160e01b600052603260045260246000fd5b90600052602060002001548152602001908152602001600020805461123890611c56565b80601f016020809104026020016040519081016040528092919081815260200182805461126490611c56565b80156112b15780601f10611286576101008083540402835291602001916112b1565b820191906000526020600020905b81548152906001019060200180831161129457829003601f168201915b50505050509050919050565b60405163699f200f60e01b8152600080516020611cf183398151915260048201527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a09063699f200f90602401602060405180830381600087803b15801561131757600080fd5b505af115801561132b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061134f919061195c565b6001600160a01b0316336001600160a01b0316146113c85760405162461bcd60e51b815260206004820152603060248201527f4f6e6c7920676f7665726e616e636520636f6e74726163742063616e2063686160448201526f3733b29036b4b734b733903637b1b59760811b6064820152608401610682565b6283d600811061141a5760405162461bcd60e51b815260206004820152601c60248201527f496e76616c6964205f6e65774d696e696e674c6f636b2076616c7565000000006044820152606401610682565b60028190556040518181527f095309b0b6e827e9df0e249aa7da4ad3362ad0617052642f3a237eb0da21ce6f9060200161100d565b6001821161149f5760405162461bcd60e51b815260206004820152601c60248201527f5469702073686f756c642062652067726561746572207468616e2031000000006044820152606401610682565b60405163288c9c9d60e01b8152336004820152306024820152604481018390527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a09063288c9c9d90606401602060405180830381600087803b1580156114f857600080fd5b505af115801561150c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115309190611978565b61156f5760405162461bcd60e51b815260206004820152601060248201526f1d1a5c081b5d5cdd081899481c185a5960821b6044820152606401610682565b80516020820120831480611584575060648311155b80611631575060405163699f200f60e01b8152600080516020611cf183398151915260048201527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a09063699f200f90602401602060405180830381600087803b1580156115e457600080fd5b505af11580156115f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061161c919061195c565b6001600160a01b0316336001600160a01b0316145b61167d5760405162461bcd60e51b815260206004820152601d60248201527f6964206d7573742062652068617368206f6620627974657320646174610000006044820152606401610682565b611688600283611c00565b604051630852cd8d60e31b8152600481018290529092507388df592f8eb5d7bd38bfef7deb0fbc02cf3778a0906342966c6890602401600060405180830381600087803b1580156116d857600080fd5b505af11580156116ec573d6000803e3d6000fd5b50505060008481526004602052604081208054859350909190611710908490611be8565b90915550503360009081526008602052604081208054849290611734908490611be8565b925050819055508160008082825461174c9190611be8565b909155505060008381526004602052604090819020549051849133917fd951d408a0f5057da5c25b826fb5ce403d56542962b6ac6994dbc6d5432fbff591611798918791908790611bc0565b60405180910390a3505050565b8280546117b190611c56565b90600052602060002090601f0160209004810192826117d35760008555611819565b82601f106117ec57805160ff1916838001178555611819565b82800160010185558215611819579182015b828111156118195782518255916020019190600101906117fe565b5061182592915061189d565b5090565b82805461183590611c56565b90600052602060002090601f0160209004810192826118575760008555611819565b82601f106118705782800160ff19823516178555611819565b82800160010185558215611819579182015b82811115611819578235825591602001919060010190611882565b5b80821115611825576000815560010161189e565b600082601f8301126118c2578081fd5b813567ffffffffffffffff808211156118dd576118dd611cc2565b604051601f8301601f19908116603f0116810190828211818310171561190557611905611cc2565b8160405283815286602085880101111561191d578485fd5b8360208701602083013792830160200193909352509392505050565b60006020828403121561194a578081fd5b813561195581611cd8565b9392505050565b60006020828403121561196d578081fd5b815161195581611cd8565b600060208284031215611989578081fd5b81518015158114611955578182fd5b6000602082840312156119a9578081fd5b5035919050565b6000806000806000608086880312156119c7578081fd5b85359450602086013567ffffffffffffffff808211156119e5578283fd5b818801915088601f8301126119f8578283fd5b813581811115611a06578384fd5b896020828501011115611a17578384fd5b60208301965080955050604088013593506060880135915080821115611a3b578283fd5b50611a48888289016118b2565b9150509295509295909350565b60008060408385031215611a67578182fd5b50508035926020909101359150565b600080600060608486031215611a8a578283fd5b8335925060208401359150604084013567ffffffffffffffff811115611aae578182fd5b611aba868287016118b2565b9150509250925092565b600060208284031215611ad5578081fd5b5051919050565b60008060408385031215611aee578182fd5b505080516020909101519092909150565b60008151808452815b81811015611b2457602081850181015186830182015201611b08565b81811115611b355782602083870101525b50601f01601f19169290920160200192915050565b600088825287602083015260c060408301528560c0830152858760e08401378060e08784010152601f19601f870116820185606084015284608084015260e08382030160a0840152611b9f60e0820185611aff565b9a9950505050505050505050565b6000602082526119556020830184611aff565b600084825283602083015260606040830152611bdf6060830184611aff565b95945050505050565b60008219821115611bfb57611bfb611cac565b500190565b600082611c1b57634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615611c3a57611c3a611cac565b500290565b600082821015611c5157611c51611cac565b500390565b600181811c90821680611c6a57607f821691505b60208210811415611c8b57634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415611ca557611ca5611cac565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114611ced57600080fd5b5056feefa19baa864049f50491093580c5433e97e8d5e41f8db1a61108b4fa44cacd93a26469706673582212209a8fa940aae289a2817201ec1522e7000e981f209c89b66e50734c3b29a7429164736f6c63430008030033",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// Deprecated: Use OracleMetaData.Sigs instead.
// OracleFuncSigs maps the 4-byte function signature to its string representation.
var OracleFuncSigs = OracleMetaData.Sigs

// OracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OracleMetaData.Bin instead.
var OracleBin = OracleMetaData.Bin

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// GetBlockNumberByTimestamp is a free data retrieval call binding the contract method 0x935408d0.
//
// Solidity: function getBlockNumberByTimestamp(bytes32 _queryId, uint256 _timestamp) view returns(uint256)
func (_Oracle *OracleCaller) GetBlockNumberByTimestamp(opts *bind.CallOpts, _queryId [32]byte, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getBlockNumberByTimestamp", _queryId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumberByTimestamp is a free data retrieval call binding the contract method 0x935408d0.
//
// Solidity: function getBlockNumberByTimestamp(bytes32 _queryId, uint256 _timestamp) view returns(uint256)
func (_Oracle *OracleSession) GetBlockNumberByTimestamp(_queryId [32]byte, _timestamp *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetBlockNumberByTimestamp(&_Oracle.CallOpts, _queryId, _timestamp)
}

// GetBlockNumberByTimestamp is a free data retrieval call binding the contract method 0x935408d0.
//
// Solidity: function getBlockNumberByTimestamp(bytes32 _queryId, uint256 _timestamp) view returns(uint256)
func (_Oracle *OracleCallerSession) GetBlockNumberByTimestamp(_queryId [32]byte, _timestamp *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetBlockNumberByTimestamp(&_Oracle.CallOpts, _queryId, _timestamp)
}

// GetCurrentReward is a free data retrieval call binding the contract method 0xa1e588a5.
//
// Solidity: function getCurrentReward(bytes32 _queryId) view returns(uint256, uint256)
func (_Oracle *OracleCaller) GetCurrentReward(opts *bind.CallOpts, _queryId [32]byte) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getCurrentReward", _queryId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCurrentReward is a free data retrieval call binding the contract method 0xa1e588a5.
//
// Solidity: function getCurrentReward(bytes32 _queryId) view returns(uint256, uint256)
func (_Oracle *OracleSession) GetCurrentReward(_queryId [32]byte) (*big.Int, *big.Int, error) {
	return _Oracle.Contract.GetCurrentReward(&_Oracle.CallOpts, _queryId)
}

// GetCurrentReward is a free data retrieval call binding the contract method 0xa1e588a5.
//
// Solidity: function getCurrentReward(bytes32 _queryId) view returns(uint256, uint256)
func (_Oracle *OracleCallerSession) GetCurrentReward(_queryId [32]byte) (*big.Int, *big.Int, error) {
	return _Oracle.Contract.GetCurrentReward(&_Oracle.CallOpts, _queryId)
}

// GetCurrentValue is a free data retrieval call binding the contract method 0xadf1639d.
//
// Solidity: function getCurrentValue(bytes32 _queryId) view returns(bytes)
func (_Oracle *OracleCaller) GetCurrentValue(opts *bind.CallOpts, _queryId [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getCurrentValue", _queryId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCurrentValue is a free data retrieval call binding the contract method 0xadf1639d.
//
// Solidity: function getCurrentValue(bytes32 _queryId) view returns(bytes)
func (_Oracle *OracleSession) GetCurrentValue(_queryId [32]byte) ([]byte, error) {
	return _Oracle.Contract.GetCurrentValue(&_Oracle.CallOpts, _queryId)
}

// GetCurrentValue is a free data retrieval call binding the contract method 0xadf1639d.
//
// Solidity: function getCurrentValue(bytes32 _queryId) view returns(bytes)
func (_Oracle *OracleCallerSession) GetCurrentValue(_queryId [32]byte) ([]byte, error) {
	return _Oracle.Contract.GetCurrentValue(&_Oracle.CallOpts, _queryId)
}

// GetMiningLock is a free data retrieval call binding the contract method 0xbf1b032e.
//
// Solidity: function getMiningLock() view returns(uint256)
func (_Oracle *OracleCaller) GetMiningLock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getMiningLock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMiningLock is a free data retrieval call binding the contract method 0xbf1b032e.
//
// Solidity: function getMiningLock() view returns(uint256)
func (_Oracle *OracleSession) GetMiningLock() (*big.Int, error) {
	return _Oracle.Contract.GetMiningLock(&_Oracle.CallOpts)
}

// GetMiningLock is a free data retrieval call binding the contract method 0xbf1b032e.
//
// Solidity: function getMiningLock() view returns(uint256)
func (_Oracle *OracleCallerSession) GetMiningLock() (*big.Int, error) {
	return _Oracle.Contract.GetMiningLock(&_Oracle.CallOpts)
}

// GetReportTimestampByIndex is a free data retrieval call binding the contract method 0x7c37b8b4.
//
// Solidity: function getReportTimestampByIndex(bytes32 _queryId, uint256 _index) view returns(uint256)
func (_Oracle *OracleCaller) GetReportTimestampByIndex(opts *bind.CallOpts, _queryId [32]byte, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getReportTimestampByIndex", _queryId, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportTimestampByIndex is a free data retrieval call binding the contract method 0x7c37b8b4.
//
// Solidity: function getReportTimestampByIndex(bytes32 _queryId, uint256 _index) view returns(uint256)
func (_Oracle *OracleSession) GetReportTimestampByIndex(_queryId [32]byte, _index *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetReportTimestampByIndex(&_Oracle.CallOpts, _queryId, _index)
}

// GetReportTimestampByIndex is a free data retrieval call binding the contract method 0x7c37b8b4.
//
// Solidity: function getReportTimestampByIndex(bytes32 _queryId, uint256 _index) view returns(uint256)
func (_Oracle *OracleCallerSession) GetReportTimestampByIndex(_queryId [32]byte, _index *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetReportTimestampByIndex(&_Oracle.CallOpts, _queryId, _index)
}

// GetReporterByTimestamp is a free data retrieval call binding the contract method 0xe07c5486.
//
// Solidity: function getReporterByTimestamp(bytes32 _queryId, uint256 _timestamp) view returns(address)
func (_Oracle *OracleCaller) GetReporterByTimestamp(opts *bind.CallOpts, _queryId [32]byte, _timestamp *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getReporterByTimestamp", _queryId, _timestamp)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReporterByTimestamp is a free data retrieval call binding the contract method 0xe07c5486.
//
// Solidity: function getReporterByTimestamp(bytes32 _queryId, uint256 _timestamp) view returns(address)
func (_Oracle *OracleSession) GetReporterByTimestamp(_queryId [32]byte, _timestamp *big.Int) (common.Address, error) {
	return _Oracle.Contract.GetReporterByTimestamp(&_Oracle.CallOpts, _queryId, _timestamp)
}

// GetReporterByTimestamp is a free data retrieval call binding the contract method 0xe07c5486.
//
// Solidity: function getReporterByTimestamp(bytes32 _queryId, uint256 _timestamp) view returns(address)
func (_Oracle *OracleCallerSession) GetReporterByTimestamp(_queryId [32]byte, _timestamp *big.Int) (common.Address, error) {
	return _Oracle.Contract.GetReporterByTimestamp(&_Oracle.CallOpts, _queryId, _timestamp)
}

// GetReportsSubmittedByAddress is a free data retrieval call binding the contract method 0x3878293e.
//
// Solidity: function getReportsSubmittedByAddress(address _reporter) view returns(uint256)
func (_Oracle *OracleCaller) GetReportsSubmittedByAddress(opts *bind.CallOpts, _reporter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getReportsSubmittedByAddress", _reporter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportsSubmittedByAddress is a free data retrieval call binding the contract method 0x3878293e.
//
// Solidity: function getReportsSubmittedByAddress(address _reporter) view returns(uint256)
func (_Oracle *OracleSession) GetReportsSubmittedByAddress(_reporter common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetReportsSubmittedByAddress(&_Oracle.CallOpts, _reporter)
}

// GetReportsSubmittedByAddress is a free data retrieval call binding the contract method 0x3878293e.
//
// Solidity: function getReportsSubmittedByAddress(address _reporter) view returns(uint256)
func (_Oracle *OracleCallerSession) GetReportsSubmittedByAddress(_reporter common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetReportsSubmittedByAddress(&_Oracle.CallOpts, _reporter)
}

// GetTimeBasedReward is a free data retrieval call binding the contract method 0x14d66b9a.
//
// Solidity: function getTimeBasedReward() view returns(uint256)
func (_Oracle *OracleCaller) GetTimeBasedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getTimeBasedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeBasedReward is a free data retrieval call binding the contract method 0x14d66b9a.
//
// Solidity: function getTimeBasedReward() view returns(uint256)
func (_Oracle *OracleSession) GetTimeBasedReward() (*big.Int, error) {
	return _Oracle.Contract.GetTimeBasedReward(&_Oracle.CallOpts)
}

// GetTimeBasedReward is a free data retrieval call binding the contract method 0x14d66b9a.
//
// Solidity: function getTimeBasedReward() view returns(uint256)
func (_Oracle *OracleCallerSession) GetTimeBasedReward() (*big.Int, error) {
	return _Oracle.Contract.GetTimeBasedReward(&_Oracle.CallOpts)
}

// GetTimeOfLastNewValue is a free data retrieval call binding the contract method 0xc0f95d52.
//
// Solidity: function getTimeOfLastNewValue() view returns(uint256)
func (_Oracle *OracleCaller) GetTimeOfLastNewValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getTimeOfLastNewValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeOfLastNewValue is a free data retrieval call binding the contract method 0xc0f95d52.
//
// Solidity: function getTimeOfLastNewValue() view returns(uint256)
func (_Oracle *OracleSession) GetTimeOfLastNewValue() (*big.Int, error) {
	return _Oracle.Contract.GetTimeOfLastNewValue(&_Oracle.CallOpts)
}

// GetTimeOfLastNewValue is a free data retrieval call binding the contract method 0xc0f95d52.
//
// Solidity: function getTimeOfLastNewValue() view returns(uint256)
func (_Oracle *OracleCallerSession) GetTimeOfLastNewValue() (*big.Int, error) {
	return _Oracle.Contract.GetTimeOfLastNewValue(&_Oracle.CallOpts)
}

// GetTimestampCountById is a free data retrieval call binding the contract method 0x35e72432.
//
// Solidity: function getTimestampCountById(bytes32 _queryId) view returns(uint256)
func (_Oracle *OracleCaller) GetTimestampCountById(opts *bind.CallOpts, _queryId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getTimestampCountById", _queryId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampCountById is a free data retrieval call binding the contract method 0x35e72432.
//
// Solidity: function getTimestampCountById(bytes32 _queryId) view returns(uint256)
func (_Oracle *OracleSession) GetTimestampCountById(_queryId [32]byte) (*big.Int, error) {
	return _Oracle.Contract.GetTimestampCountById(&_Oracle.CallOpts, _queryId)
}

// GetTimestampCountById is a free data retrieval call binding the contract method 0x35e72432.
//
// Solidity: function getTimestampCountById(bytes32 _queryId) view returns(uint256)
func (_Oracle *OracleCallerSession) GetTimestampCountById(_queryId [32]byte) (*big.Int, error) {
	return _Oracle.Contract.GetTimestampCountById(&_Oracle.CallOpts, _queryId)
}

// GetTimestampIndexByTimestamp is a free data retrieval call binding the contract method 0x9d9b16ed.
//
// Solidity: function getTimestampIndexByTimestamp(bytes32 _queryId, uint256 _timestamp) view returns(uint256)
func (_Oracle *OracleCaller) GetTimestampIndexByTimestamp(opts *bind.CallOpts, _queryId [32]byte, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getTimestampIndexByTimestamp", _queryId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampIndexByTimestamp is a free data retrieval call binding the contract method 0x9d9b16ed.
//
// Solidity: function getTimestampIndexByTimestamp(bytes32 _queryId, uint256 _timestamp) view returns(uint256)
func (_Oracle *OracleSession) GetTimestampIndexByTimestamp(_queryId [32]byte, _timestamp *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetTimestampIndexByTimestamp(&_Oracle.CallOpts, _queryId, _timestamp)
}

// GetTimestampIndexByTimestamp is a free data retrieval call binding the contract method 0x9d9b16ed.
//
// Solidity: function getTimestampIndexByTimestamp(bytes32 _queryId, uint256 _timestamp) view returns(uint256)
func (_Oracle *OracleCallerSession) GetTimestampIndexByTimestamp(_queryId [32]byte, _timestamp *big.Int) (*big.Int, error) {
	return _Oracle.Contract.GetTimestampIndexByTimestamp(&_Oracle.CallOpts, _queryId, _timestamp)
}

// GetTipsById is a free data retrieval call binding the contract method 0xef4c262d.
//
// Solidity: function getTipsById(bytes32 _queryId) view returns(uint256)
func (_Oracle *OracleCaller) GetTipsById(opts *bind.CallOpts, _queryId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getTipsById", _queryId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTipsById is a free data retrieval call binding the contract method 0xef4c262d.
//
// Solidity: function getTipsById(bytes32 _queryId) view returns(uint256)
func (_Oracle *OracleSession) GetTipsById(_queryId [32]byte) (*big.Int, error) {
	return _Oracle.Contract.GetTipsById(&_Oracle.CallOpts, _queryId)
}

// GetTipsById is a free data retrieval call binding the contract method 0xef4c262d.
//
// Solidity: function getTipsById(bytes32 _queryId) view returns(uint256)
func (_Oracle *OracleCallerSession) GetTipsById(_queryId [32]byte) (*big.Int, error) {
	return _Oracle.Contract.GetTipsById(&_Oracle.CallOpts, _queryId)
}

// GetTipsByUser is a free data retrieval call binding the contract method 0xb736ec36.
//
// Solidity: function getTipsByUser(address _user) view returns(uint256)
func (_Oracle *OracleCaller) GetTipsByUser(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getTipsByUser", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTipsByUser is a free data retrieval call binding the contract method 0xb736ec36.
//
// Solidity: function getTipsByUser(address _user) view returns(uint256)
func (_Oracle *OracleSession) GetTipsByUser(_user common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetTipsByUser(&_Oracle.CallOpts, _user)
}

// GetTipsByUser is a free data retrieval call binding the contract method 0xb736ec36.
//
// Solidity: function getTipsByUser(address _user) view returns(uint256)
func (_Oracle *OracleCallerSession) GetTipsByUser(_user common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetTipsByUser(&_Oracle.CallOpts, _user)
}

// GetValueByTimestamp is a free data retrieval call binding the contract method 0x0b2d2b0d.
//
// Solidity: function getValueByTimestamp(bytes32 _queryId, uint256 _timestamp) view returns(bytes)
func (_Oracle *OracleCaller) GetValueByTimestamp(opts *bind.CallOpts, _queryId [32]byte, _timestamp *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getValueByTimestamp", _queryId, _timestamp)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetValueByTimestamp is a free data retrieval call binding the contract method 0x0b2d2b0d.
//
// Solidity: function getValueByTimestamp(bytes32 _queryId, uint256 _timestamp) view returns(bytes)
func (_Oracle *OracleSession) GetValueByTimestamp(_queryId [32]byte, _timestamp *big.Int) ([]byte, error) {
	return _Oracle.Contract.GetValueByTimestamp(&_Oracle.CallOpts, _queryId, _timestamp)
}

// GetValueByTimestamp is a free data retrieval call binding the contract method 0x0b2d2b0d.
//
// Solidity: function getValueByTimestamp(bytes32 _queryId, uint256 _timestamp) view returns(bytes)
func (_Oracle *OracleCallerSession) GetValueByTimestamp(_queryId [32]byte, _timestamp *big.Int) ([]byte, error) {
	return _Oracle.Contract.GetValueByTimestamp(&_Oracle.CallOpts, _queryId, _timestamp)
}

// MiningLock is a free data retrieval call binding the contract method 0xd804b663.
//
// Solidity: function miningLock() view returns(uint256)
func (_Oracle *OracleCaller) MiningLock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "miningLock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningLock is a free data retrieval call binding the contract method 0xd804b663.
//
// Solidity: function miningLock() view returns(uint256)
func (_Oracle *OracleSession) MiningLock() (*big.Int, error) {
	return _Oracle.Contract.MiningLock(&_Oracle.CallOpts)
}

// MiningLock is a free data retrieval call binding the contract method 0xd804b663.
//
// Solidity: function miningLock() view returns(uint256)
func (_Oracle *OracleCallerSession) MiningLock() (*big.Int, error) {
	return _Oracle.Contract.MiningLock(&_Oracle.CallOpts)
}

// TimeBasedReward is a free data retrieval call binding the contract method 0x96426d97.
//
// Solidity: function timeBasedReward() view returns(uint256)
func (_Oracle *OracleCaller) TimeBasedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "timeBasedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeBasedReward is a free data retrieval call binding the contract method 0x96426d97.
//
// Solidity: function timeBasedReward() view returns(uint256)
func (_Oracle *OracleSession) TimeBasedReward() (*big.Int, error) {
	return _Oracle.Contract.TimeBasedReward(&_Oracle.CallOpts)
}

// TimeBasedReward is a free data retrieval call binding the contract method 0x96426d97.
//
// Solidity: function timeBasedReward() view returns(uint256)
func (_Oracle *OracleCallerSession) TimeBasedReward() (*big.Int, error) {
	return _Oracle.Contract.TimeBasedReward(&_Oracle.CallOpts)
}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(uint256)
func (_Oracle *OracleCaller) TimeOfLastNewValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "timeOfLastNewValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(uint256)
func (_Oracle *OracleSession) TimeOfLastNewValue() (*big.Int, error) {
	return _Oracle.Contract.TimeOfLastNewValue(&_Oracle.CallOpts)
}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(uint256)
func (_Oracle *OracleCallerSession) TimeOfLastNewValue() (*big.Int, error) {
	return _Oracle.Contract.TimeOfLastNewValue(&_Oracle.CallOpts)
}

// Tips is a free data retrieval call binding the contract method 0x602bf227.
//
// Solidity: function tips(bytes32 ) view returns(uint256)
func (_Oracle *OracleCaller) Tips(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "tips", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tips is a free data retrieval call binding the contract method 0x602bf227.
//
// Solidity: function tips(bytes32 ) view returns(uint256)
func (_Oracle *OracleSession) Tips(arg0 [32]byte) (*big.Int, error) {
	return _Oracle.Contract.Tips(&_Oracle.CallOpts, arg0)
}

// Tips is a free data retrieval call binding the contract method 0x602bf227.
//
// Solidity: function tips(bytes32 ) view returns(uint256)
func (_Oracle *OracleCallerSession) Tips(arg0 [32]byte) (*big.Int, error) {
	return _Oracle.Contract.Tips(&_Oracle.CallOpts, arg0)
}

// TipsInContract is a free data retrieval call binding the contract method 0x69d43bd3.
//
// Solidity: function tipsInContract() view returns(uint256)
func (_Oracle *OracleCaller) TipsInContract(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "tipsInContract")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TipsInContract is a free data retrieval call binding the contract method 0x69d43bd3.
//
// Solidity: function tipsInContract() view returns(uint256)
func (_Oracle *OracleSession) TipsInContract() (*big.Int, error) {
	return _Oracle.Contract.TipsInContract(&_Oracle.CallOpts)
}

// TipsInContract is a free data retrieval call binding the contract method 0x69d43bd3.
//
// Solidity: function tipsInContract() view returns(uint256)
func (_Oracle *OracleCallerSession) TipsInContract() (*big.Int, error) {
	return _Oracle.Contract.TipsInContract(&_Oracle.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0xfc735e99.
//
// Solidity: function verify() pure returns(uint256)
func (_Oracle *OracleCaller) Verify(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "verify")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xfc735e99.
//
// Solidity: function verify() pure returns(uint256)
func (_Oracle *OracleSession) Verify() (*big.Int, error) {
	return _Oracle.Contract.Verify(&_Oracle.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0xfc735e99.
//
// Solidity: function verify() pure returns(uint256)
func (_Oracle *OracleCallerSession) Verify() (*big.Int, error) {
	return _Oracle.Contract.Verify(&_Oracle.CallOpts)
}

// ChangeMiningLock is a paid mutator transaction binding the contract method 0xe280e8e8.
//
// Solidity: function changeMiningLock(uint256 _newMiningLock) returns()
func (_Oracle *OracleTransactor) ChangeMiningLock(opts *bind.TransactOpts, _newMiningLock *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "changeMiningLock", _newMiningLock)
}

// ChangeMiningLock is a paid mutator transaction binding the contract method 0xe280e8e8.
//
// Solidity: function changeMiningLock(uint256 _newMiningLock) returns()
func (_Oracle *OracleSession) ChangeMiningLock(_newMiningLock *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.ChangeMiningLock(&_Oracle.TransactOpts, _newMiningLock)
}

// ChangeMiningLock is a paid mutator transaction binding the contract method 0xe280e8e8.
//
// Solidity: function changeMiningLock(uint256 _newMiningLock) returns()
func (_Oracle *OracleTransactorSession) ChangeMiningLock(_newMiningLock *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.ChangeMiningLock(&_Oracle.TransactOpts, _newMiningLock)
}

// ChangeTimeBasedReward is a paid mutator transaction binding the contract method 0x6d53585f.
//
// Solidity: function changeTimeBasedReward(uint256 _newTimeBasedReward) returns()
func (_Oracle *OracleTransactor) ChangeTimeBasedReward(opts *bind.TransactOpts, _newTimeBasedReward *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "changeTimeBasedReward", _newTimeBasedReward)
}

// ChangeTimeBasedReward is a paid mutator transaction binding the contract method 0x6d53585f.
//
// Solidity: function changeTimeBasedReward(uint256 _newTimeBasedReward) returns()
func (_Oracle *OracleSession) ChangeTimeBasedReward(_newTimeBasedReward *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.ChangeTimeBasedReward(&_Oracle.TransactOpts, _newTimeBasedReward)
}

// ChangeTimeBasedReward is a paid mutator transaction binding the contract method 0x6d53585f.
//
// Solidity: function changeTimeBasedReward(uint256 _newTimeBasedReward) returns()
func (_Oracle *OracleTransactorSession) ChangeTimeBasedReward(_newTimeBasedReward *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.ChangeTimeBasedReward(&_Oracle.TransactOpts, _newTimeBasedReward)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x5b5edcfc.
//
// Solidity: function removeValue(bytes32 _queryId, uint256 _timestamp) returns()
func (_Oracle *OracleTransactor) RemoveValue(opts *bind.TransactOpts, _queryId [32]byte, _timestamp *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "removeValue", _queryId, _timestamp)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x5b5edcfc.
//
// Solidity: function removeValue(bytes32 _queryId, uint256 _timestamp) returns()
func (_Oracle *OracleSession) RemoveValue(_queryId [32]byte, _timestamp *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.RemoveValue(&_Oracle.TransactOpts, _queryId, _timestamp)
}

// RemoveValue is a paid mutator transaction binding the contract method 0x5b5edcfc.
//
// Solidity: function removeValue(bytes32 _queryId, uint256 _timestamp) returns()
func (_Oracle *OracleTransactorSession) RemoveValue(_queryId [32]byte, _timestamp *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.RemoveValue(&_Oracle.TransactOpts, _queryId, _timestamp)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x5eaa9ced.
//
// Solidity: function submitValue(bytes32 _queryId, bytes _value, uint256 _nonce, bytes _queryData) returns()
func (_Oracle *OracleTransactor) SubmitValue(opts *bind.TransactOpts, _queryId [32]byte, _value []byte, _nonce *big.Int, _queryData []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "submitValue", _queryId, _value, _nonce, _queryData)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x5eaa9ced.
//
// Solidity: function submitValue(bytes32 _queryId, bytes _value, uint256 _nonce, bytes _queryData) returns()
func (_Oracle *OracleSession) SubmitValue(_queryId [32]byte, _value []byte, _nonce *big.Int, _queryData []byte) (*types.Transaction, error) {
	return _Oracle.Contract.SubmitValue(&_Oracle.TransactOpts, _queryId, _value, _nonce, _queryData)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x5eaa9ced.
//
// Solidity: function submitValue(bytes32 _queryId, bytes _value, uint256 _nonce, bytes _queryData) returns()
func (_Oracle *OracleTransactorSession) SubmitValue(_queryId [32]byte, _value []byte, _nonce *big.Int, _queryData []byte) (*types.Transaction, error) {
	return _Oracle.Contract.SubmitValue(&_Oracle.TransactOpts, _queryId, _value, _nonce, _queryData)
}

// TipQuery is a paid mutator transaction binding the contract method 0xef0234ad.
//
// Solidity: function tipQuery(bytes32 _queryId, uint256 _tip, bytes _queryData) returns()
func (_Oracle *OracleTransactor) TipQuery(opts *bind.TransactOpts, _queryId [32]byte, _tip *big.Int, _queryData []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "tipQuery", _queryId, _tip, _queryData)
}

// TipQuery is a paid mutator transaction binding the contract method 0xef0234ad.
//
// Solidity: function tipQuery(bytes32 _queryId, uint256 _tip, bytes _queryData) returns()
func (_Oracle *OracleSession) TipQuery(_queryId [32]byte, _tip *big.Int, _queryData []byte) (*types.Transaction, error) {
	return _Oracle.Contract.TipQuery(&_Oracle.TransactOpts, _queryId, _tip, _queryData)
}

// TipQuery is a paid mutator transaction binding the contract method 0xef0234ad.
//
// Solidity: function tipQuery(bytes32 _queryId, uint256 _tip, bytes _queryData) returns()
func (_Oracle *OracleTransactorSession) TipQuery(_queryId [32]byte, _tip *big.Int, _queryData []byte) (*types.Transaction, error) {
	return _Oracle.Contract.TipQuery(&_Oracle.TransactOpts, _queryId, _tip, _queryData)
}

// OracleMiningLockChangedIterator is returned from FilterMiningLockChanged and is used to iterate over the raw logs and unpacked data for MiningLockChanged events raised by the Oracle contract.
type OracleMiningLockChangedIterator struct {
	Event *OracleMiningLockChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleMiningLockChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleMiningLockChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleMiningLockChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleMiningLockChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleMiningLockChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleMiningLockChanged represents a MiningLockChanged event raised by the Oracle contract.
type OracleMiningLockChanged struct {
	NewMiningLock *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMiningLockChanged is a free log retrieval operation binding the contract event 0x095309b0b6e827e9df0e249aa7da4ad3362ad0617052642f3a237eb0da21ce6f.
//
// Solidity: event MiningLockChanged(uint256 _newMiningLock)
func (_Oracle *OracleFilterer) FilterMiningLockChanged(opts *bind.FilterOpts) (*OracleMiningLockChangedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "MiningLockChanged")
	if err != nil {
		return nil, err
	}
	return &OracleMiningLockChangedIterator{contract: _Oracle.contract, event: "MiningLockChanged", logs: logs, sub: sub}, nil
}

// WatchMiningLockChanged is a free log subscription operation binding the contract event 0x095309b0b6e827e9df0e249aa7da4ad3362ad0617052642f3a237eb0da21ce6f.
//
// Solidity: event MiningLockChanged(uint256 _newMiningLock)
func (_Oracle *OracleFilterer) WatchMiningLockChanged(opts *bind.WatchOpts, sink chan<- *OracleMiningLockChanged) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "MiningLockChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleMiningLockChanged)
				if err := _Oracle.contract.UnpackLog(event, "MiningLockChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMiningLockChanged is a log parse operation binding the contract event 0x095309b0b6e827e9df0e249aa7da4ad3362ad0617052642f3a237eb0da21ce6f.
//
// Solidity: event MiningLockChanged(uint256 _newMiningLock)
func (_Oracle *OracleFilterer) ParseMiningLockChanged(log types.Log) (*OracleMiningLockChanged, error) {
	event := new(OracleMiningLockChanged)
	if err := _Oracle.contract.UnpackLog(event, "MiningLockChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleNewReportIterator is returned from FilterNewReport and is used to iterate over the raw logs and unpacked data for NewReport events raised by the Oracle contract.
type OracleNewReportIterator struct {
	Event *OracleNewReport // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleNewReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleNewReport)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleNewReport)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleNewReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleNewReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleNewReport represents a NewReport event raised by the Oracle contract.
type OracleNewReport struct {
	QueryId   [32]byte
	Time      *big.Int
	Value     []byte
	Reward    *big.Int
	Nonce     *big.Int
	QueryData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewReport is a free log retrieval operation binding the contract event 0xb5e84abaa3916191eacb4f925fb0d61505d1485414b50288e5bb52dd63c75cb7.
//
// Solidity: event NewReport(bytes32 _queryId, uint256 _time, bytes _value, uint256 _reward, uint256 _nonce, bytes _queryData)
func (_Oracle *OracleFilterer) FilterNewReport(opts *bind.FilterOpts) (*OracleNewReportIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "NewReport")
	if err != nil {
		return nil, err
	}
	return &OracleNewReportIterator{contract: _Oracle.contract, event: "NewReport", logs: logs, sub: sub}, nil
}

// WatchNewReport is a free log subscription operation binding the contract event 0xb5e84abaa3916191eacb4f925fb0d61505d1485414b50288e5bb52dd63c75cb7.
//
// Solidity: event NewReport(bytes32 _queryId, uint256 _time, bytes _value, uint256 _reward, uint256 _nonce, bytes _queryData)
func (_Oracle *OracleFilterer) WatchNewReport(opts *bind.WatchOpts, sink chan<- *OracleNewReport) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "NewReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleNewReport)
				if err := _Oracle.contract.UnpackLog(event, "NewReport", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewReport is a log parse operation binding the contract event 0xb5e84abaa3916191eacb4f925fb0d61505d1485414b50288e5bb52dd63c75cb7.
//
// Solidity: event NewReport(bytes32 _queryId, uint256 _time, bytes _value, uint256 _reward, uint256 _nonce, bytes _queryData)
func (_Oracle *OracleFilterer) ParseNewReport(log types.Log) (*OracleNewReport, error) {
	event := new(OracleNewReport)
	if err := _Oracle.contract.UnpackLog(event, "NewReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleTimeBasedRewardsChangedIterator is returned from FilterTimeBasedRewardsChanged and is used to iterate over the raw logs and unpacked data for TimeBasedRewardsChanged events raised by the Oracle contract.
type OracleTimeBasedRewardsChangedIterator struct {
	Event *OracleTimeBasedRewardsChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleTimeBasedRewardsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTimeBasedRewardsChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleTimeBasedRewardsChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleTimeBasedRewardsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTimeBasedRewardsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTimeBasedRewardsChanged represents a TimeBasedRewardsChanged event raised by the Oracle contract.
type OracleTimeBasedRewardsChanged struct {
	NewTimeBasedReward *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTimeBasedRewardsChanged is a free log retrieval operation binding the contract event 0x01432c5d54689da3dbfec924c89d1e453ecaec349928ceabd50207e909e8f7de.
//
// Solidity: event TimeBasedRewardsChanged(uint256 _newTimeBasedReward)
func (_Oracle *OracleFilterer) FilterTimeBasedRewardsChanged(opts *bind.FilterOpts) (*OracleTimeBasedRewardsChangedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "TimeBasedRewardsChanged")
	if err != nil {
		return nil, err
	}
	return &OracleTimeBasedRewardsChangedIterator{contract: _Oracle.contract, event: "TimeBasedRewardsChanged", logs: logs, sub: sub}, nil
}

// WatchTimeBasedRewardsChanged is a free log subscription operation binding the contract event 0x01432c5d54689da3dbfec924c89d1e453ecaec349928ceabd50207e909e8f7de.
//
// Solidity: event TimeBasedRewardsChanged(uint256 _newTimeBasedReward)
func (_Oracle *OracleFilterer) WatchTimeBasedRewardsChanged(opts *bind.WatchOpts, sink chan<- *OracleTimeBasedRewardsChanged) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "TimeBasedRewardsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTimeBasedRewardsChanged)
				if err := _Oracle.contract.UnpackLog(event, "TimeBasedRewardsChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimeBasedRewardsChanged is a log parse operation binding the contract event 0x01432c5d54689da3dbfec924c89d1e453ecaec349928ceabd50207e909e8f7de.
//
// Solidity: event TimeBasedRewardsChanged(uint256 _newTimeBasedReward)
func (_Oracle *OracleFilterer) ParseTimeBasedRewardsChanged(log types.Log) (*OracleTimeBasedRewardsChanged, error) {
	event := new(OracleTimeBasedRewardsChanged)
	if err := _Oracle.contract.UnpackLog(event, "TimeBasedRewardsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleTipAddedIterator is returned from FilterTipAdded and is used to iterate over the raw logs and unpacked data for TipAdded events raised by the Oracle contract.
type OracleTipAddedIterator struct {
	Event *OracleTipAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleTipAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTipAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleTipAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleTipAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTipAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTipAdded represents a TipAdded event raised by the Oracle contract.
type OracleTipAdded struct {
	User      common.Address
	QueryId   [32]byte
	Tip       *big.Int
	TotalTip  *big.Int
	QueryData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTipAdded is a free log retrieval operation binding the contract event 0xd951d408a0f5057da5c25b826fb5ce403d56542962b6ac6994dbc6d5432fbff5.
//
// Solidity: event TipAdded(address indexed _user, bytes32 indexed _queryId, uint256 _tip, uint256 _totalTip, bytes _queryData)
func (_Oracle *OracleFilterer) FilterTipAdded(opts *bind.FilterOpts, _user []common.Address, _queryId [][32]byte) (*OracleTipAddedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _queryIdRule []interface{}
	for _, _queryIdItem := range _queryId {
		_queryIdRule = append(_queryIdRule, _queryIdItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "TipAdded", _userRule, _queryIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleTipAddedIterator{contract: _Oracle.contract, event: "TipAdded", logs: logs, sub: sub}, nil
}

// WatchTipAdded is a free log subscription operation binding the contract event 0xd951d408a0f5057da5c25b826fb5ce403d56542962b6ac6994dbc6d5432fbff5.
//
// Solidity: event TipAdded(address indexed _user, bytes32 indexed _queryId, uint256 _tip, uint256 _totalTip, bytes _queryData)
func (_Oracle *OracleFilterer) WatchTipAdded(opts *bind.WatchOpts, sink chan<- *OracleTipAdded, _user []common.Address, _queryId [][32]byte) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _queryIdRule []interface{}
	for _, _queryIdItem := range _queryId {
		_queryIdRule = append(_queryIdRule, _queryIdItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "TipAdded", _userRule, _queryIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTipAdded)
				if err := _Oracle.contract.UnpackLog(event, "TipAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTipAdded is a log parse operation binding the contract event 0xd951d408a0f5057da5c25b826fb5ce403d56542962b6ac6994dbc6d5432fbff5.
//
// Solidity: event TipAdded(address indexed _user, bytes32 indexed _queryId, uint256 _tip, uint256 _totalTip, bytes _queryData)
func (_Oracle *OracleFilterer) ParseTipAdded(log types.Log) (*OracleTipAdded, error) {
	event := new(OracleTipAdded)
	if err := _Oracle.contract.UnpackLog(event, "TipAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorVariablesMetaData contains all meta data concerning the TellorVariables contract.
var TellorVariablesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220a82e3ebe9a894a945cecef6fe31794bb656b2bdb829432ddf51d6f0213f4523664736f6c63430008030033",
}

// TellorVariablesABI is the input ABI used to generate the binding from.
// Deprecated: Use TellorVariablesMetaData.ABI instead.
var TellorVariablesABI = TellorVariablesMetaData.ABI

// TellorVariablesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TellorVariablesMetaData.Bin instead.
var TellorVariablesBin = TellorVariablesMetaData.Bin

// DeployTellorVariables deploys a new Ethereum contract, binding an instance of TellorVariables to it.
func DeployTellorVariables(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorVariables, error) {
	parsed, err := TellorVariablesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TellorVariablesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorVariables{TellorVariablesCaller: TellorVariablesCaller{contract: contract}, TellorVariablesTransactor: TellorVariablesTransactor{contract: contract}, TellorVariablesFilterer: TellorVariablesFilterer{contract: contract}}, nil
}

// TellorVariables is an auto generated Go binding around an Ethereum contract.
type TellorVariables struct {
	TellorVariablesCaller     // Read-only binding to the contract
	TellorVariablesTransactor // Write-only binding to the contract
	TellorVariablesFilterer   // Log filterer for contract events
}

// TellorVariablesCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorVariablesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorVariablesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorVariablesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorVariablesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorVariablesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorVariablesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorVariablesSession struct {
	Contract     *TellorVariables  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorVariablesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorVariablesCallerSession struct {
	Contract *TellorVariablesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TellorVariablesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorVariablesTransactorSession struct {
	Contract     *TellorVariablesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TellorVariablesRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorVariablesRaw struct {
	Contract *TellorVariables // Generic contract binding to access the raw methods on
}

// TellorVariablesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorVariablesCallerRaw struct {
	Contract *TellorVariablesCaller // Generic read-only contract binding to access the raw methods on
}

// TellorVariablesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorVariablesTransactorRaw struct {
	Contract *TellorVariablesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorVariables creates a new instance of TellorVariables, bound to a specific deployed contract.
func NewTellorVariables(address common.Address, backend bind.ContractBackend) (*TellorVariables, error) {
	contract, err := bindTellorVariables(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorVariables{TellorVariablesCaller: TellorVariablesCaller{contract: contract}, TellorVariablesTransactor: TellorVariablesTransactor{contract: contract}, TellorVariablesFilterer: TellorVariablesFilterer{contract: contract}}, nil
}

// NewTellorVariablesCaller creates a new read-only instance of TellorVariables, bound to a specific deployed contract.
func NewTellorVariablesCaller(address common.Address, caller bind.ContractCaller) (*TellorVariablesCaller, error) {
	contract, err := bindTellorVariables(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorVariablesCaller{contract: contract}, nil
}

// NewTellorVariablesTransactor creates a new write-only instance of TellorVariables, bound to a specific deployed contract.
func NewTellorVariablesTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorVariablesTransactor, error) {
	contract, err := bindTellorVariables(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorVariablesTransactor{contract: contract}, nil
}

// NewTellorVariablesFilterer creates a new log filterer instance of TellorVariables, bound to a specific deployed contract.
func NewTellorVariablesFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorVariablesFilterer, error) {
	contract, err := bindTellorVariables(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorVariablesFilterer{contract: contract}, nil
}

// bindTellorVariables binds a generic wrapper to an already deployed contract.
func bindTellorVariables(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorVariablesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorVariables *TellorVariablesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorVariables.Contract.TellorVariablesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorVariables *TellorVariablesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorVariables.Contract.TellorVariablesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorVariables *TellorVariablesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorVariables.Contract.TellorVariablesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorVariables *TellorVariablesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorVariables.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorVariables *TellorVariablesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorVariables.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorVariables *TellorVariablesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorVariables.Contract.contract.Transact(opts, method, params...)
}

// TellorVarsMetaData contains all meta data concerning the TellorVars contract.
var TellorVarsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220c47f9d3f7eb3eb036c286e598828abeebadce5795fa0acd56bad3b44b3326ff164736f6c63430008030033",
}

// TellorVarsABI is the input ABI used to generate the binding from.
// Deprecated: Use TellorVarsMetaData.ABI instead.
var TellorVarsABI = TellorVarsMetaData.ABI

// TellorVarsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TellorVarsMetaData.Bin instead.
var TellorVarsBin = TellorVarsMetaData.Bin

// DeployTellorVars deploys a new Ethereum contract, binding an instance of TellorVars to it.
func DeployTellorVars(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorVars, error) {
	parsed, err := TellorVarsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TellorVarsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorVars{TellorVarsCaller: TellorVarsCaller{contract: contract}, TellorVarsTransactor: TellorVarsTransactor{contract: contract}, TellorVarsFilterer: TellorVarsFilterer{contract: contract}}, nil
}

// TellorVars is an auto generated Go binding around an Ethereum contract.
type TellorVars struct {
	TellorVarsCaller     // Read-only binding to the contract
	TellorVarsTransactor // Write-only binding to the contract
	TellorVarsFilterer   // Log filterer for contract events
}

// TellorVarsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorVarsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorVarsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorVarsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorVarsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorVarsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorVarsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorVarsSession struct {
	Contract     *TellorVars       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorVarsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorVarsCallerSession struct {
	Contract *TellorVarsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TellorVarsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorVarsTransactorSession struct {
	Contract     *TellorVarsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TellorVarsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorVarsRaw struct {
	Contract *TellorVars // Generic contract binding to access the raw methods on
}

// TellorVarsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorVarsCallerRaw struct {
	Contract *TellorVarsCaller // Generic read-only contract binding to access the raw methods on
}

// TellorVarsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorVarsTransactorRaw struct {
	Contract *TellorVarsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorVars creates a new instance of TellorVars, bound to a specific deployed contract.
func NewTellorVars(address common.Address, backend bind.ContractBackend) (*TellorVars, error) {
	contract, err := bindTellorVars(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorVars{TellorVarsCaller: TellorVarsCaller{contract: contract}, TellorVarsTransactor: TellorVarsTransactor{contract: contract}, TellorVarsFilterer: TellorVarsFilterer{contract: contract}}, nil
}

// NewTellorVarsCaller creates a new read-only instance of TellorVars, bound to a specific deployed contract.
func NewTellorVarsCaller(address common.Address, caller bind.ContractCaller) (*TellorVarsCaller, error) {
	contract, err := bindTellorVars(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorVarsCaller{contract: contract}, nil
}

// NewTellorVarsTransactor creates a new write-only instance of TellorVars, bound to a specific deployed contract.
func NewTellorVarsTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorVarsTransactor, error) {
	contract, err := bindTellorVars(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorVarsTransactor{contract: contract}, nil
}

// NewTellorVarsFilterer creates a new log filterer instance of TellorVars, bound to a specific deployed contract.
func NewTellorVarsFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorVarsFilterer, error) {
	contract, err := bindTellorVars(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorVarsFilterer{contract: contract}, nil
}

// bindTellorVars binds a generic wrapper to an already deployed contract.
func bindTellorVars(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorVarsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorVars *TellorVarsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorVars.Contract.TellorVarsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorVars *TellorVarsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorVars.Contract.TellorVarsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorVars *TellorVarsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorVars.Contract.TellorVarsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorVars *TellorVarsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorVars.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorVars *TellorVarsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorVars.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorVars *TellorVarsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorVars.Contract.contract.Transact(opts, method, params...)
}
