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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_queryData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_reporter\",\"type\":\"address\"}],\"name\":\"NewReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newReportingLock\",\"type\":\"uint256\"}],\"name\":\"ReportingLockChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newTimeBasedReward\",\"type\":\"uint256\"}],\"name\":\"TimeBasedRewardsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_queryData\",\"type\":\"bytes\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newReportingLock\",\"type\":\"uint256\"}],\"name\":\"changeReportingLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newTimeBasedReward\",\"type\":\"uint256\"}],\"name\":\"changeTimeBasedReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getBlockNumberByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"}],\"name\":\"getCurrentReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"}],\"name\":\"getCurrentValue\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReportTimestampByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getReporterByTimestamp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter\",\"type\":\"address\"}],\"name\":\"getReporterLastTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportingLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reporter\",\"type\":\"address\"}],\"name\":\"getReportsSubmittedByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeBasedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeOfLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"}],\"name\":\"getTimestampCountById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getTimestampIndexByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"}],\"name\":\"getTipsById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getTipsByUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getValueByTimestamp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"removeValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reportingLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_queryData\",\"type\":\"bytes\"}],\"name\":\"submitValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeBasedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeOfLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_queryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_queryData\",\"type\":\"bytes\"}],\"name\":\"tipQuery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tipsInContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5d183cfa": "changeReportingLock(uint256)",
		"6d53585f": "changeTimeBasedReward(uint256)",
		"935408d0": "getBlockNumberByTimestamp(bytes32,uint256)",
		"a1e588a5": "getCurrentReward(bytes32)",
		"adf1639d": "getCurrentValue(bytes32)",
		"7c37b8b4": "getReportTimestampByIndex(bytes32,uint256)",
		"e07c5486": "getReporterByTimestamp(bytes32,uint256)",
		"50005b83": "getReporterLastTimestamp(address)",
		"460c33a2": "getReportingLock()",
		"3878293e": "getReportsSubmittedByAddress(address)",
		"14d66b9a": "getTimeBasedReward()",
		"c0f95d52": "getTimeOfLastNewValue()",
		"35e72432": "getTimestampCountById(bytes32)",
		"9d9b16ed": "getTimestampIndexByTimestamp(bytes32,uint256)",
		"ef4c262d": "getTipsById(bytes32)",
		"b736ec36": "getTipsByUser(address)",
		"0b2d2b0d": "getValueByTimestamp(bytes32,uint256)",
		"5b5edcfc": "removeValue(bytes32,uint256)",
		"3321fc41": "reportingLock()",
		"5eaa9ced": "submitValue(bytes32,bytes,uint256,bytes)",
		"96426d97": "timeBasedReward()",
		"6fd4f229": "timeOfLastNewValue()",
		"ef0234ad": "tipQuery(bytes32,uint256,bytes)",
		"602bf227": "tips(bytes32)",
		"69d43bd3": "tipsInContract()",
		"fc735e99": "verify()",
	},
	Bin: "0x608060405261a8c06000556706f05b59d3b200006001554260025534801561002657600080fd5b50611df5806100366000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c80636fd4f229116100de578063adf1639d11610097578063e07c548611610071578063e07c5486146103b7578063ef0234ad14610406578063ef4c262d14610419578063fc735e99146104395761018e565b8063adf1639d14610373578063b736ec3614610386578063c0f95d52146103af5761018e565b80636fd4f229146102c85780637c37b8b4146102d1578063935408d0146102e457806396426d97146103135780639d9b16ed1461031c578063a1e588a51461034b5761018e565b806350005b831161014b5780635eaa9ced116101255780635eaa9ced14610279578063602bf2271461028c57806369d43bd3146102ac5780636d53585f146102b55761018e565b806350005b83146102285780635b5edcfc146102515780635d183cfa146102665761018e565b80630b2d2b0d1461019357806314d66b9a146101bc5780633321fc41146101ce57806335e72432146101d75780633878293e146101f7578063460c33a214610220575b600080fd5b6101a66101a1366004611af4565b610441565b6040516101b39190611c5c565b60405180910390f35b6001545b6040519081526020016101b3565b6101c060005481565b6101c06101e5366004611a37565b60009081526004602052604090205490565b6101c06102053660046119d8565b6001600160a01b031660009081526007602052604090205490565b6000546101c0565b6101c06102363660046119d8565b6001600160a01b031660009081526006602052604090205490565b61026461025f366004611af4565b6104f2565b005b610264610274366004611a37565b610884565b610264610287366004611a4f565b610a20565b6101c061029a366004611a37565b60056020526000908152604090205481565b6101c060035481565b6102646102c3366004611a37565b611102565b6101c060025481565b6101c06102df366004611af4565b611248565b6101c06102f2366004611af4565b60009182526004602090815260408084209284526002909201905290205490565b6101c060015481565b6101c061032a366004611af4565b60009182526004602090815260408084209284526001909201905290205490565b61035e610359366004611a37565b611289565b604080519283526020830191909152016101b3565b6101a6610381366004611a37565b611400565b6101c06103943660046119d8565b6001600160a01b031660009081526008602052604090205490565b6002546101c0565b6103ee6103c5366004611af4565b60009182526004602081815260408085209385529290910190529020546001600160a01b031690565b6040516001600160a01b0390911681526020016101b3565b610264610414366004611b15565b6114ed565b6101c0610427366004611a37565b60009081526005602052604090205490565b61270f6101c0565b6000828152600460209081526040808320848452600301909152902080546060919061046c90611d05565b80601f016020809104026020016040519081016040528092919081815260200182805461049890611d05565b80156104e55780601f106104ba576101008083540402835291602001916104e5565b820191906000526020600020905b8154815290600101906020018083116104c857829003601f168201915b5050505050905092915050565b60405163699f200f60e01b8152600080516020611da083398151915260048201527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a09063699f200f90602401602060405180830381600087803b15801561054c57600080fd5b505af1158015610560573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061058491906119fb565b6001600160a01b0316336001600160a01b03161480610657575060405163699f200f60e01b81527ffa522e460446113e8fd353d7fa015625a68bc0369712213a42e006346440891e60048201527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a09063699f200f90602401602060405180830381600087803b15801561060a57600080fd5b505af115801561061e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064291906119fb565b6001600160a01b0316336001600160a01b0316145b6106ce5760405162461bcd60e51b815260206004820152603d60248201527f63616c6c6572206d7573742062652074686520676f7665726e616e636520636f60448201527f6e7472616374206f7220746865206f7261636c6520636f6e747261637400000060648201526084015b60405180910390fd5b60008281526004602090815260408083208484526001810190925290912054805b82546106fd90600190611cee565b8110156107d55782610710826001611c97565b8154811061072e57634e487b7160e01b600052603260045260246000fd5b906000526020600020015483600001828154811061075c57634e487b7160e01b600052603260045260246000fd5b9060005260206000200181905550600183600101600085600001848154811061079557634e487b7160e01b600052603260045260246000fd5b9060005260206000200154815260200190815260200160002060008282546107bd9190611cee565b909155508190506107cd81611d40565b9150506106ef565b50815482906107e690600190611cee565b8154811061080457634e487b7160e01b600052603260045260246000fd5b6000918252602082200155815482908061082e57634e487b7160e01b600052603160045260246000fd5b60008281526020808220830160001990810183905590920190925560408051808301808352848252878552600387019093529220915161086f929190611844565b50506000918252600101602052604081205550565b60405163699f200f60e01b8152600080516020611da083398151915260048201527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a09063699f200f90602401602060405180830381600087803b1580156108de57600080fd5b505af11580156108f2573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061091691906119fb565b6001600160a01b0316336001600160a01b0316146109925760405162461bcd60e51b815260206004820152603360248201527f4f6e6c7920676f7665726e616e636520636f6e74726163742063616e206368616044820152723733b2903932b837b93a34b733903637b1b59760691b60648201526084016106c5565b6283d60081106109e45760405162461bcd60e51b815260206004820152601f60248201527f496e76616c6964205f6e65775265706f7274696e674c6f636b2076616c75650060448201526064016106c5565b60008190556040518181527f8867dac9fd763aeaa408825321cc53f6fd67a46ea326adadfd2a378d17fedf22906020015b60405180910390a150565b600085815260046020526040902080548314610a7e5760405162461bcd60e51b815260206004820181905260248201527f6e6f6e6365206d757374206d617463682074696d657374616d7020696e64657860448201526064016106c5565b600080543382526006602052604090912054610a9a9042611cee565b11610af95760405162461bcd60e51b815260206004820152602960248201527f7374696c6c20696e207265706f727465722074696d65206c6f636b2c20706c6560448201526861736520776169742160b81b60648201526084016106c5565b60405163699f200f60e01b81527ffa522e460446113e8fd353d7fa015625a68bc0369712213a42e006346440891e60048201527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a09063699f200f90602401602060405180830381600087803b158015610b6557600080fd5b505af1158015610b79573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9d91906119fb565b6001600160a01b0316306001600160a01b031614610c105760405162461bcd60e51b815260206004820152602a60248201527f63616e206f6e6c79207375626d697420746f2063757272656e74206f7261636c604482015269194818dbdb9d1c9858dd60b21b60648201526084016106c5565b81516020830120861480610c25575060648611155b610c715760405162461bcd60e51b815260206004820152601d60248201527f6964206d7573742062652068617368206f66206279746573206461746100000060448201526064016106c5565b3360008181526006602052604080822042905551630733bdef60e41b815260048101929092527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a091829063733bdef090602401604080518083038186803b158015610ccf57600080fd5b505afa158015610ce3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d079190611b7b565b50905080600114610d5a5760405162461bcd60e51b815260206004820152601d60248201527f5265706f7274657220737461747573206973206e6f74207374616b657200000060448201526064016106c5565b604051632d67853560e21b81527f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc9760048201526001600160a01b0383169063b59e14d490602401602060405180830381600087803b158015610dbb57600080fd5b505af1158015610dcf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df39190611b63565b6040516370a0823160e01b81523360048201526001600160a01b038416906370a082319060240160206040518083038186803b158015610e3257600080fd5b505afa158015610e46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6a9190611b63565b1015610eca5760405162461bcd60e51b815260206004820152602960248201527f62616c616e6365206d7573742062652067726561746572207468616e207374616044820152681ad948185b5bdd5b9d60ba1b60648201526084016106c5565b4260009081526004840160205260409020546001600160a01b031615610f325760405162461bcd60e51b815260206004820152601e60248201527f74696d657374616d7020616c7265616479207265706f7274656420666f72000060448201526064016106c5565b82544260008181526001808701602090815260408084208690559185018855878352808320909401839055918152600286018352818120439055600386019092529020610f809088886118c8565b50426000908152600484016020526040812080546001600160a01b0319163317905580610fac8a611289565b915091508160036000828254610fc29190611cee565b9091555060009050610fd48383611c97565b1115611074576001600160a01b03841663a9059cbb33610ff48585611c97565b6040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604401602060405180830381600087803b15801561103a57600080fd5b505af115801561104e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110729190611a17565b505b60008a815260056020908152604080832083905542600255338352600790915281208054916110a283611d40565b909155507fab1d593f8e2ecb165106e30a39db6769078d35c3fdbb110f24932f0d7af68c2990508a428b8b6110d78688611c97565b8c8c336040516110ee989796959493929190611be9565b60405180910390a150505050505050505050565b60405163699f200f60e01b8152600080516020611da083398151915260048201527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a09063699f200f90602401602060405180830381600087803b15801561115c57600080fd5b505af1158015611170573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061119491906119fb565b6001600160a01b0316336001600160a01b0316146112135760405162461bcd60e51b815260206004820152603660248201527f4f6e6c7920676f7665726e616e636520636f6e74726163742063616e206368616044820152753733b2903a34b6b2903130b9b2b2103932bbb0b9321760511b60648201526084016106c5565b60018190556040518181527f01432c5d54689da3dbfec924c89d1e453ecaec349928ceabd50207e909e8f7de90602001610a15565b600082815260046020526040812080548390811061127657634e487b7160e01b600052603260045260246000fd5b9060005260206000200154905092915050565b60008060007388df592f8eb5d7bd38bfef7deb0fbc02cf3778a090506000600254426112b59190611cee565b9050600061012c600154836112ca9190611ccf565b6112d49190611caf565b9050600354816112e49190611c97565b6040516370a0823160e01b81523060048201526001600160a01b038516906370a082319060240160206040518083038186803b15801561132357600080fd5b505afa158015611337573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061135b9190611b63565b10156113e8576003546040516370a0823160e01b81523060048201526001600160a01b038516906370a082319060240160206040518083038186803b1580156113a357600080fd5b505afa1580156113b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113db9190611b63565b6113e59190611cee565b90505b60009586526005602052604090952054959350505050565b600081815260046020526040812080546060926003830192909161142690600190611cee565b8154811061144457634e487b7160e01b600052603260045260246000fd5b90600052602060002001548152602001908152602001600020805461146890611d05565b80601f016020809104026020016040519081016040528092919081815260200182805461149490611d05565b80156114e15780601f106114b6576101008083540402835291602001916114e1565b820191906000526020600020905b8154815290600101906020018083116114c457829003601f168201915b50505050509050919050565b6001821161153d5760405162461bcd60e51b815260206004820152601c60248201527f5469702073686f756c642062652067726561746572207468616e20310000000060448201526064016106c5565b60405163288c9c9d60e01b8152336004820152306024820152604481018390527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a09063288c9c9d90606401602060405180830381600087803b15801561159657600080fd5b505af11580156115aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ce9190611a17565b61160d5760405162461bcd60e51b815260206004820152601060248201526f1d1a5c081b5d5cdd081899481c185a5960821b60448201526064016106c5565b80516020820120831480611622575060648311155b806116cf575060405163699f200f60e01b8152600080516020611da083398151915260048201527388df592f8eb5d7bd38bfef7deb0fbc02cf3778a09063699f200f90602401602060405180830381600087803b15801561168257600080fd5b505af1158015611696573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116ba91906119fb565b6001600160a01b0316336001600160a01b0316145b61171b5760405162461bcd60e51b815260206004820152601d60248201527f6964206d7573742062652068617368206f66206279746573206461746100000060448201526064016106c5565b611726600283611caf565b604051630852cd8d60e31b8152600481018290529092507388df592f8eb5d7bd38bfef7deb0fbc02cf3778a0906342966c6890602401600060405180830381600087803b15801561177657600080fd5b505af115801561178a573d6000803e3d6000fd5b505050600084815260056020526040812080548593509091906117ae908490611c97565b909155505033600090815260086020526040812080548492906117d2908490611c97565b9250508190555081600360008282546117eb9190611c97565b909155505060008381526005602052604090819020549051849133917fd951d408a0f5057da5c25b826fb5ce403d56542962b6ac6994dbc6d5432fbff591611837918791908790611c6f565b60405180910390a3505050565b82805461185090611d05565b90600052602060002090601f01602090048101928261187257600085556118b8565b82601f1061188b57805160ff19168380011785556118b8565b828001600101855582156118b8579182015b828111156118b857825182559160200191906001019061189d565b506118c492915061193c565b5090565b8280546118d490611d05565b90600052602060002090601f0160209004810192826118f657600085556118b8565b82601f1061190f5782800160ff198235161785556118b8565b828001600101855582156118b8579182015b828111156118b8578235825591602001919060010190611921565b5b808211156118c4576000815560010161193d565b600082601f830112611961578081fd5b813567ffffffffffffffff8082111561197c5761197c611d71565b604051601f8301601f19908116603f011681019082821181831017156119a4576119a4611d71565b816040528381528660208588010111156119bc578485fd5b8360208701602083013792830160200193909352509392505050565b6000602082840312156119e9578081fd5b81356119f481611d87565b9392505050565b600060208284031215611a0c578081fd5b81516119f481611d87565b600060208284031215611a28578081fd5b815180151581146119f4578182fd5b600060208284031215611a48578081fd5b5035919050565b600080600080600060808688031215611a66578081fd5b85359450602086013567ffffffffffffffff80821115611a84578283fd5b818801915088601f830112611a97578283fd5b813581811115611aa5578384fd5b896020828501011115611ab6578384fd5b60208301965080955050604088013593506060880135915080821115611ada578283fd5b50611ae788828901611951565b9150509295509295909350565b60008060408385031215611b06578182fd5b50508035926020909101359150565b600080600060608486031215611b29578283fd5b8335925060208401359150604084013567ffffffffffffffff811115611b4d578182fd5b611b5986828701611951565b9150509250925092565b600060208284031215611b74578081fd5b5051919050565b60008060408385031215611b8d578182fd5b505080516020909101519092909150565b60008151808452815b81811015611bc357602081850181015186830182015201611ba7565b81811115611bd45782602083870101525b50601f01601f19169290920160200192915050565b600089825288602083015260e060408301528660e083015261010087898285013781818985010152601f19601f8901168301876060850152866080850152818482030160a0850152611c3d82820187611b9e565b9250505060018060a01b03831660c08301529998505050505050505050565b6000602082526119f46020830184611b9e565b600084825283602083015260606040830152611c8e6060830184611b9e565b95945050505050565b60008219821115611caa57611caa611d5b565b500190565b600082611cca57634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615611ce957611ce9611d5b565b500290565b600082821015611d0057611d00611d5b565b500390565b600181811c90821680611d1957607f821691505b60208210811415611d3a57634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415611d5457611d54611d5b565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114611d9c57600080fd5b5056feefa19baa864049f50491093580c5433e97e8d5e41f8db1a61108b4fa44cacd93a26469706673582212209b0a392f41201f6046b72df4df84466dd6f7d0a96a927f8781e6d3955d2bb03164736f6c63430008030033",
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

// GetReporterLastTimestamp is a free data retrieval call binding the contract method 0x50005b83.
//
// Solidity: function getReporterLastTimestamp(address _reporter) view returns(uint256)
func (_Oracle *OracleCaller) GetReporterLastTimestamp(opts *bind.CallOpts, _reporter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getReporterLastTimestamp", _reporter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReporterLastTimestamp is a free data retrieval call binding the contract method 0x50005b83.
//
// Solidity: function getReporterLastTimestamp(address _reporter) view returns(uint256)
func (_Oracle *OracleSession) GetReporterLastTimestamp(_reporter common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetReporterLastTimestamp(&_Oracle.CallOpts, _reporter)
}

// GetReporterLastTimestamp is a free data retrieval call binding the contract method 0x50005b83.
//
// Solidity: function getReporterLastTimestamp(address _reporter) view returns(uint256)
func (_Oracle *OracleCallerSession) GetReporterLastTimestamp(_reporter common.Address) (*big.Int, error) {
	return _Oracle.Contract.GetReporterLastTimestamp(&_Oracle.CallOpts, _reporter)
}

// GetReportingLock is a free data retrieval call binding the contract method 0x460c33a2.
//
// Solidity: function getReportingLock() view returns(uint256)
func (_Oracle *OracleCaller) GetReportingLock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getReportingLock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportingLock is a free data retrieval call binding the contract method 0x460c33a2.
//
// Solidity: function getReportingLock() view returns(uint256)
func (_Oracle *OracleSession) GetReportingLock() (*big.Int, error) {
	return _Oracle.Contract.GetReportingLock(&_Oracle.CallOpts)
}

// GetReportingLock is a free data retrieval call binding the contract method 0x460c33a2.
//
// Solidity: function getReportingLock() view returns(uint256)
func (_Oracle *OracleCallerSession) GetReportingLock() (*big.Int, error) {
	return _Oracle.Contract.GetReportingLock(&_Oracle.CallOpts)
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

// ReportingLock is a free data retrieval call binding the contract method 0x3321fc41.
//
// Solidity: function reportingLock() view returns(uint256)
func (_Oracle *OracleCaller) ReportingLock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "reportingLock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReportingLock is a free data retrieval call binding the contract method 0x3321fc41.
//
// Solidity: function reportingLock() view returns(uint256)
func (_Oracle *OracleSession) ReportingLock() (*big.Int, error) {
	return _Oracle.Contract.ReportingLock(&_Oracle.CallOpts)
}

// ReportingLock is a free data retrieval call binding the contract method 0x3321fc41.
//
// Solidity: function reportingLock() view returns(uint256)
func (_Oracle *OracleCallerSession) ReportingLock() (*big.Int, error) {
	return _Oracle.Contract.ReportingLock(&_Oracle.CallOpts)
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

// ChangeReportingLock is a paid mutator transaction binding the contract method 0x5d183cfa.
//
// Solidity: function changeReportingLock(uint256 _newReportingLock) returns()
func (_Oracle *OracleTransactor) ChangeReportingLock(opts *bind.TransactOpts, _newReportingLock *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "changeReportingLock", _newReportingLock)
}

// ChangeReportingLock is a paid mutator transaction binding the contract method 0x5d183cfa.
//
// Solidity: function changeReportingLock(uint256 _newReportingLock) returns()
func (_Oracle *OracleSession) ChangeReportingLock(_newReportingLock *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.ChangeReportingLock(&_Oracle.TransactOpts, _newReportingLock)
}

// ChangeReportingLock is a paid mutator transaction binding the contract method 0x5d183cfa.
//
// Solidity: function changeReportingLock(uint256 _newReportingLock) returns()
func (_Oracle *OracleTransactorSession) ChangeReportingLock(_newReportingLock *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.ChangeReportingLock(&_Oracle.TransactOpts, _newReportingLock)
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
	Reporter  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewReport is a free log retrieval operation binding the contract event 0xab1d593f8e2ecb165106e30a39db6769078d35c3fdbb110f24932f0d7af68c29.
//
// Solidity: event NewReport(bytes32 _queryId, uint256 _time, bytes _value, uint256 _reward, uint256 _nonce, bytes _queryData, address _reporter)
func (_Oracle *OracleFilterer) FilterNewReport(opts *bind.FilterOpts) (*OracleNewReportIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "NewReport")
	if err != nil {
		return nil, err
	}
	return &OracleNewReportIterator{contract: _Oracle.contract, event: "NewReport", logs: logs, sub: sub}, nil
}

// WatchNewReport is a free log subscription operation binding the contract event 0xab1d593f8e2ecb165106e30a39db6769078d35c3fdbb110f24932f0d7af68c29.
//
// Solidity: event NewReport(bytes32 _queryId, uint256 _time, bytes _value, uint256 _reward, uint256 _nonce, bytes _queryData, address _reporter)
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

// ParseNewReport is a log parse operation binding the contract event 0xab1d593f8e2ecb165106e30a39db6769078d35c3fdbb110f24932f0d7af68c29.
//
// Solidity: event NewReport(bytes32 _queryId, uint256 _time, bytes _value, uint256 _reward, uint256 _nonce, bytes _queryData, address _reporter)
func (_Oracle *OracleFilterer) ParseNewReport(log types.Log) (*OracleNewReport, error) {
	event := new(OracleNewReport)
	if err := _Oracle.contract.UnpackLog(event, "NewReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleReportingLockChangedIterator is returned from FilterReportingLockChanged and is used to iterate over the raw logs and unpacked data for ReportingLockChanged events raised by the Oracle contract.
type OracleReportingLockChangedIterator struct {
	Event *OracleReportingLockChanged // Event containing the contract specifics and raw log

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
func (it *OracleReportingLockChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleReportingLockChanged)
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
		it.Event = new(OracleReportingLockChanged)
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
func (it *OracleReportingLockChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleReportingLockChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleReportingLockChanged represents a ReportingLockChanged event raised by the Oracle contract.
type OracleReportingLockChanged struct {
	NewReportingLock *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReportingLockChanged is a free log retrieval operation binding the contract event 0x8867dac9fd763aeaa408825321cc53f6fd67a46ea326adadfd2a378d17fedf22.
//
// Solidity: event ReportingLockChanged(uint256 _newReportingLock)
func (_Oracle *OracleFilterer) FilterReportingLockChanged(opts *bind.FilterOpts) (*OracleReportingLockChangedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "ReportingLockChanged")
	if err != nil {
		return nil, err
	}
	return &OracleReportingLockChangedIterator{contract: _Oracle.contract, event: "ReportingLockChanged", logs: logs, sub: sub}, nil
}

// WatchReportingLockChanged is a free log subscription operation binding the contract event 0x8867dac9fd763aeaa408825321cc53f6fd67a46ea326adadfd2a378d17fedf22.
//
// Solidity: event ReportingLockChanged(uint256 _newReportingLock)
func (_Oracle *OracleFilterer) WatchReportingLockChanged(opts *bind.WatchOpts, sink chan<- *OracleReportingLockChanged) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "ReportingLockChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleReportingLockChanged)
				if err := _Oracle.contract.UnpackLog(event, "ReportingLockChanged", log); err != nil {
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

// ParseReportingLockChanged is a log parse operation binding the contract event 0x8867dac9fd763aeaa408825321cc53f6fd67a46ea326adadfd2a378d17fedf22.
//
// Solidity: event ReportingLockChanged(uint256 _newReportingLock)
func (_Oracle *OracleFilterer) ParseReportingLockChanged(log types.Log) (*OracleReportingLockChanged, error) {
	event := new(OracleReportingLockChanged)
	if err := _Oracle.contract.UnpackLog(event, "ReportingLockChanged", log); err != nil {
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
