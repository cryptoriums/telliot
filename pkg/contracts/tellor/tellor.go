// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tellor

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

// ExtensionMetaData contains all meta data concerning the Extension contract.
var ExtensionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_result\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reportedMiner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_reportingParty\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_passed\",\"type\":\"bool\"}],\"name\":\"DisputeVoteTallied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newTellor\",\"type\":\"address\"}],\"name\":\"NewTellorAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"fromBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentMiners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"disputesById\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"tally\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disputeVotePassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPropFork\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"reportedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reportingParty\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedForkAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_diff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"idsOnDeck\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"tipsOnDeck\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopRequestIDs\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minersByChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newValueTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateMinDisputeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"updateTellor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"024c2ddd": "_allowances(address,address)",
		"699f200f": "addresses(bytes32)",
		"cbf1304d": "balances(address,uint256)",
		"62dd1d2a": "bytesVars(bytes32)",
		"1fd22364": "currentMiners(uint256)",
		"313ce567": "decimals()",
		"0d2d76a2": "depositStake()",
		"63bb82ad": "didMine(bytes32,address)",
		"a7c438bc": "didVote(uint256,address)",
		"d01f4d9e": "disputeIdByDisputeHash(bytes32)",
		"db085beb": "disputesById(uint256)",
		"133bee5e": "getAddressVars(bytes32)",
		"af0b1327": "getAllDisputeVars(uint256)",
		"da379941": "getDisputeIdByDisputeHash(bytes32)",
		"7f6fd5d9": "getDisputeUintVars(uint256,bytes32)",
		"3180f8df": "getLastNewValueById(uint256)",
		"c775b542": "getMinedBlockNum(uint256,uint256)",
		"69026d63": "getMinersByRequestIdAndTimestamp(uint256,uint256)",
		"4049f198": "getNewCurrentVariables()",
		"46eee1c4": "getNewValueCountbyRequestId(uint256)",
		"9a7077ab": "getNewVariablesOnDeck()",
		"6173c0b8": "getRequestIdByRequestQIndex(uint256)",
		"b5413029": "getRequestQ()",
		"e0ae93c1": "getRequestUintVars(uint256,bytes32)",
		"e1eee6d6": "getRequestVars(uint256)",
		"733bdef0": "getStakerInfo(address)",
		"11c98512": "getSubmissionsByTimestamp(uint256,uint256)",
		"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
		"fe1cd15d": "getTopRequestIDs()",
		"612c8f7f": "getUintVar(bytes32)",
		"3df0777b": "isInDispute(uint256,uint256)",
		"4ba0a5ee": "migrated(address)",
		"48b18e54": "minersByChallenge(bytes32,address)",
		"06fdde03": "name()",
		"438c0aa3": "newValueTimestamps(uint256)",
		"5700242c": "requestIdByQueryHash(bytes32)",
		"28449c3a": "requestStakingWithdraw()",
		"93fa4915": "retrieveData(uint256,uint256)",
		"95d89b41": "symbol()",
		"4d318b0e": "tallyVotes(uint256)",
		"18160ddd": "totalSupply()",
		"b59e14d4": "uints(bytes32)",
		"90e5b235": "updateMinDisputeFee()",
		"f458ab98": "updateTellor(uint256)",
		"bed9d861": "withdrawStake()",
	},
	Bin: "0x608060405234801561001057600080fd5b506122d5806100206000396000f3fe608060405234801561001057600080fd5b506004361061027f5760003560e01c806369026d631161015c578063b5413029116100ce578063da37994111610087578063da37994114610975578063db085beb14610992578063e0ae93c1146109ff578063e1eee6d614610a22578063f458ab9814610a3f578063fe1cd15d14610a5c5761027f565b8063b541302914610897578063b59e14d4146108b5578063bed9d861146108d2578063c775b542146108da578063cbf1304d146108fd578063d01f4d9e146109585761027f565b806390e5b2351161012057806390e5b2351461071b57806393fa49151461072357806395d89b41146107465780639a7077ab1461074e578063a7c438bc146107ba578063af0b1327146107e65761027f565b806369026d6314610656578063699f200f14610679578063733bdef01461069657806377fbb663146106d55780637f6fd5d9146106f85761027f565b80634049f198116101f55780634d318b0e116101b95780634d318b0e146105995780635700242c146105b6578063612c8f7f146105d35780636173c0b8146105f057806362dd1d2a1461060d57806363bb82ad1461062a5761027f565b80634049f198146104b8578063438c0aa31461050d57806346eee1c41461052a57806348b18e54146105475780634ba0a5ee146105735761027f565b806318160ddd1161024757806318160ddd146103df5780631fd22364146103e757806328449c3a14610425578063313ce5671461042d5780633180f8df1461044b5780633df0777b146104815761027f565b8063024c2ddd1461028457806306fdde03146102c45780630d2d76a21461034157806311c985121461034b578063133bee5e146103a6575b600080fd5b6102b26004803603604081101561029a57600080fd5b506001600160a01b0381358116916020013516610a64565b60408051918252519081900360200190f35b6102cc610a81565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103065781810151838201526020016102ee565b50505050905090810190601f1680156103335780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610349610aaa565b005b61036e6004803603604081101561036157600080fd5b5080359060200135610abd565b604051808260a080838360005b8381101561039357818101518382015260200161037b565b5050505090500191505060405180910390f35b6103c3600480360360208110156103bc57600080fd5b5035610b14565b604080516001600160a01b039092168252519081900360200190f35b6102b2610b2f565b610404600480360360208110156103fd57600080fd5b5035610b7d565b604080519283526001600160a01b0390911660208301528051918290030190f35b610349610ba8565b610435610c9c565b6040805160ff9092168252519081900360200190f35b6104686004803603602081101561046157600080fd5b5035610ca1565b6040805192835290151560208301528051918290030190f35b6104a46004803603604081101561049757600080fd5b5080359060200135610cfb565b604080519115158252519081900360200190f35b6104c0610d1f565b604051848152602081018460a080838360005b838110156104eb5781810151838201526020016104d3565b5050505090500183815260200182815260200194505050505060405180910390f35b6102b26004803603602081101561052357600080fd5b5035610dff565b6102b26004803603602081101561054057600080fd5b5035610e20565b6104a46004803603604081101561055d57600080fd5b50803590602001356001600160a01b0316610e32565b6104a46004803603602081101561058957600080fd5b50356001600160a01b0316610e52565b610349600480360360208110156105af57600080fd5b5035610e67565b6102b2600480360360208110156105cc57600080fd5b5035611123565b6102b2600480360360208110156105e957600080fd5b5035611135565b6102b26004803603602081101561060657600080fd5b5035611147565b6102b26004803603602081101561062357600080fd5b50356111b2565b6104a46004803603604081101561064057600080fd5b50803590602001356001600160a01b03166111c4565b61036e6004803603604081101561066c57600080fd5b50803590602001356111ef565b6103c36004803603602081101561068f57600080fd5b5035611252565b6106bc600480360360208110156106ac57600080fd5b50356001600160a01b031661126d565b6040805192835260208301919091528051918290030190f35b6102b2600480360360408110156106eb57600080fd5b5080359060200135611290565b6102b26004803603604081101561070e57600080fd5b50803590602001356112bd565b6103496112de565b6102b26004803603604081101561073957600080fd5b50803590602001356113f8565b6102cc611419565b610756611436565b604051808360a080838360005b8381101561077b578181015183820152602001610763565b5050505090500182600560200280838360005b838110156107a657818101518382015260200161078e565b505050509050019250505060405180910390f35b6104a4600480360360408110156107d057600080fd5b50803590602001356001600160a01b03166114d5565b610803600480360360208110156107fc57600080fd5b5035611504565b604051808a8152602001891515815260200188151581526020018715158152602001866001600160a01b03168152602001856001600160a01b03168152602001846001600160a01b0316815260200183600960200280838360005b8381101561087657818101518382015260200161085e565b50505050905001828152602001995050505050505050505060405180910390f35b61089f61172f565b604051815181528082610660808383602061037b565b6102b2600480360360208110156108cb57600080fd5b503561176b565b61034961177d565b6102b2600480360360408110156108f057600080fd5b5080359060200135611857565b6109296004803603604081101561091357600080fd5b506001600160a01b038135169060200135611878565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b6102b26004803603602081101561096e57600080fd5b50356118bb565b6102b26004803603602081101561098b57600080fd5b50356118cd565b6109af600480360360208110156109a857600080fd5b50356118df565b60408051988952602089019790975294151587870152921515606087015290151560808601526001600160a01b0390811660a086015290811660c08501521660e083015251908190036101000190f35b6102b260048036036040811015610a1557600080fd5b508035906020013561193b565b6106bc60048036036020811015610a3857600080fd5b503561195c565b61034960048036036020811015610a5557600080fd5b50356119c4565b61036e611d68565b604a60209081526000928352604080842090915290825290205481565b60408051808201909152600f81526e54656c6c6f7220547269627574657360881b602082015290565b610ab333611e4d565b610abb6112de565b565b610ac5612188565b600083815260456020908152604080832085845260060190915290819020815160a08101928390529160059082845b815481526020019060010190808311610af4575050505050905092915050565b6000908152604760205260409020546001600160a01b031690565b7fe6148e7230ca038d456350e69a91b66968b222bfac9ebfbea6ff0a1fb738016060005260466020527ffffeead1ec15181fd57b4590d95e0c076bccb59e311315e8b38f23c710aa7c3e5490565b603a8160058110610b8d57600080fd5b6002020180546001909101549091506001600160a01b031682565b3360009081526044602052604090208054600114610c03576040805162461bcd60e51b8152602060048201526013602482015272135a5b995c881a5cc81b9bdd081cdd185ad959606a1b604482015290519081900360640190fd5b60028155620151804206420360018201557f10c168823622203e4057b65015ff4d95b4c650b308918e8c92dc32ab5a0a034b60005260466020527fa5ae3e2b97d73fb849ea855d27f073b72815b38452d976bd57e4a157827dadd38054600019019055610c6e6112de565b60405133907f453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf90600090a250565b601290565b6000818152604560205260408120805482919015610ced578054610ce190859083906000198101908110610cd157fe5b90600052602060002001546113f8565b60019250925050610cf6565b60008092509250505b915091565b60009182526045602090815260408084209284526004909201905290205460ff1690565b6000610d29612188565b60008060005b6005811015610d6357603a8160058110610d4557fe5b6002020154848260058110610d5657fe5b6020020152600101610d2f565b50507f52cb9007c7c6068f8ef37039d4f232cbf5a28ff8d93a5983c4c0c27cd2f9bc0d5460466020527f5bccd7373734898281f858d7562320d2cdfc0b17bd72f779686937174d150025547f09659d32f99e50ac728058418d38174fe83a137c455ff1847e6fb8e15f78f77a6000527f38b16d06a20ab673b01c748aff938df6a38f81640035f4ce8bd9abb03aae5b7254919450915090919293565b60338181548110610e0f57600080fd5b600091825260209091200154905081565b60009081526045602052604090205490565b603960209081526000928352604080842090915290825290205460ff1681565b604b6020526000908152604090205460ff1681565b6000818152603660205260409020600281015460ff1615610eb95760405162461bcd60e51b815260040180806020018281038252602181526020018061224c6021913960400191505060405180910390fd5b7f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d286000908152600582016020526040902054421015610f3f576040805162461bcd60e51b815260206004820152601f60248201527f54696d6520666f7220766f74696e6720686176656e277420656c617073656400604482015290519081900360640190fd5b60038101546001600160a01b0316610f9e576040805162461bcd60e51b815260206004820152601c60248201527f7265706f7274696e672050617274792069732061646472657373203000000000604482015290519081900360640190fd5b6001810154600081131561106f57600282015462010000900460ff166110055760028201805461ff00191661010017908190556001600160a01b036301000000909104166000908152604460205260409020805460031415610fff57600481555b5061106f565b7fe6148e7230ca038d456350e69a91b66968b222bfac9ebfbea6ff0a1fb738016060005260466020527ffffeead1ec15181fd57b4590d95e0c076bccb59e311315e8b38f23c710aa7c3e5460649060050204811061106f5760028201805461ff0019166101001790555b7ff9e1ae10923bfc79f52e309baf8c7699edb821f91ef5b5bd07be29545917b3a6600090815260058301602090815260409182902042905560028401805460ff191660011790819055600385015483518581526001600160a01b0391821693810193909352610100820460ff1615158385015292516301000000909104929092169185917f21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739919081900360600190a3505050565b60376020526000908152604090205481565b60009081526046602052604090205490565b6000603282111561119f576040805162461bcd60e51b815260206004820152601a60248201527f526571756573745120696e6465782069732061626f7665203530000000000000604482015290519081900360640190fd5b5060009081526035602052604090205490565b60486020526000908152604090205481565b60009182526039602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6111f7612188565b6000838152604560209081526040808320858452600590810190925291829020825160a08101938490529290919082845b81546001600160a01b03168152600190910190602001808311611228575050505050905092915050565b6047602052600090815260409020546001600160a01b031681565b6001600160a01b0316600090815260446020526040902080546001909101549091565b60008281526045602052604081208054839081106112aa57fe5b9060005260206000200154905092915050565b60009182526036602090815260408084209284526005909201905290205490565b60466020527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be547f2e2f0a18eb55ef91e37921b3810d7feeef7a855ddc7f4f4249ef03d7b887ae31547f10c168823622203e4057b65015ff4d95b4c650b308918e8c92dc32ab5a0a034b6000527fa5ae3e2b97d73fb849ea855d27f073b72815b38452d976bd57e4a157827dadd3546113a99067d02ab486cedc0000906103e890849061138c90829061200d565b6103e80286028161139957fe5b04816113a157fe5b048403612025565b7f675d2171f68d6f5545d54fb9b1fb61a0e6897e6188ca1cd664e7c9530d91ecfc60005260466020527f3e5522f19747f0f285b96ded572ac4128c3a764aea9f44058dc0afc9dda44986555050565b60009182526045602090815260408084209284526003909201905290205490565b6040805180820190915260038152622a292160e91b602082015290565b61143e612188565b611446612188565b61144e611d68565b915060005b60058110156114d0576045600084836005811061146c57fe5b6020020151815260200190815260200160002060010160007f1590276b7f31dd8e2a06f9a92867333eeb3eddbc91e73b9833e3e55d8e34f77d60001b8152602001908152602001600020548282600581106114c357fe5b6020020152600101611453565b509091565b60008281526036602090815260408083206001600160a01b038516845260060190915290205460ff1692915050565b60008060008060008060006115176121a6565b5050506000958652505060366020908152604080862080546002820154600383015460048401548551610120810187527f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f8c5260058601808952878d205482527f2f9328a9c75282bec25bb04befad06926366736e0030c985108445fa728335e58d52808952878d2054828a01527f9147231ab14efb72c38117f68521ddef8de64f092c18c69dbfb602ffc4de7f478d52808952878d2054828901527f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d288d52808952878d205460608301527f1da378694063870452ce03b189f48e04c1aa026348e74e6c86e10738514ad2c48d52808952878d205460808301527f4b4cefd5ced7569ef0d091282b4bca9c52a034c56471a6061afd1bf307a2de7c8d52808952878d205460a08301527f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d8d52808952878d205460c08301527f30e85ae205656781c1a951cba9f9f53f884833c049d377a2a7046eb5e6d14b268d52808952878d205460e08301527f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc4098d52909752949099205461010080870191909152600190930154919960ff8083169a948304811699506201000083041697506001600160a01b036301000000909204821696509281169493169291565b6117376121c5565b604080516106608101918290529060009060339082845b81548152602001906001019080831161174e575050505050905090565b60466020526000908152604090205481565b336000908152604460205260409020600181015462093a809062015180420642030310156117e7576040805162461bcd60e51b8152602060048201526012602482015271372064617973206469646e2774207061737360701b604482015290519081900360640190fd5b80546002146118275760405162461bcd60e51b81526004018080602001828103825260238152602001806121e56023913960400191505060405180910390fd5b600080825560405133917f4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec91a250565b60009182526045602090815260408084209284526002909201905290205490565b6049602052816000526040600020818154811061189457600080fd5b6000918252602090912001546001600160801b038082169350600160801b90910416905082565b60386020526000908152604090205481565b60009081526038602052604090205490565b603660205260009081526040902080546001820154600283015460038401546004909401549293919260ff808316936101008404821693620100008104909216926001600160a01b036301000000909304831692918216911688565b60009182526045602090815260408084209284526001909201905290205490565b60009081526045602090815260408083207ff68d680ab3160f1aa5d9c3a1383c49e3e60bf3c0c031245cbb036f5ce99afaa18452600101909152808220547f1590276b7f31dd8e2a06f9a92867333eeb3eddbc91e73b9833e3e55d8e34f77d83529120549091565b60008181526036602081815260408084205480855260388352818520548086528484528286207f6ab2b18aafe78fd59c6a4092015bddd9fcacb8170f72b299074f74d76a91a92387526005018085528387205484518087019190915284518082038701815290850185528051908601208752845282862054808752949093529320600281015491929162010000900460ff16611aa7576040805162461bcd60e51b815260206004820152601760248201527f6d757374206265206120666f726b2070726f706f73616c000000000000000000604482015290519081900360640190fd5b7fda571dfc0b95cdc4a3835f5982cfdf36f73258bee7cb8eb797b4af8b17329875600090815260058201602052604090205415611b155760405162461bcd60e51b815260040180806020018281038252602281526020018061222a6022913960400191505060405180910390fd5b600281015460ff610100909104161515600114611b6e576040805162461bcd60e51b8152602060048201526012602482015271766f7465206e6565647320746f207061737360701b604482015290519081900360640190fd5b7ff9e1ae10923bfc79f52e309baf8c7699edb821f91ef5b5bd07be29545917b3a66000908152600582016020526040902054611bf1576040805162461bcd60e51b815260206004820152601860248201527f766f7465206e6565647320746f2062652074616c6c6965640000000000000000604482015290519081900360640190fd5b7ff9e1ae10923bfc79f52e309baf8c7699edb821f91ef5b5bd07be29545917b3a6600090815260058201602052604090205462015180429190910311611c685760405162461bcd60e51b815260040180806020018281038252603381526020018061226d6033913960400191505060405180910390fd5b7fda571dfc0b95cdc4a3835f5982cfdf36f73258bee7cb8eb797b4af8b17329875600090815260058201602090815260408083206001905560048401547f0f1293c916694ac6af4daa2f866f0448d0c2ce8847074a7896d397c961914a08909352604782527ffe10c9a395cce5a324df121072934b83aa2f3aa5f594428b2a75cf926b73fae880546001600160a01b0319166001600160a01b0390941693841790557f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c3839055805183815290517fc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d929181900390910190a1505050505050565b611d70612188565b611d78612188565b611d80612188565b60408051610660810191829052611db79160009060339082845b815481526020019060010190808311611d9a575050505050612034565b909250905060005b6005811015611e4757828160058110611dd457fe5b602002015115611e185760356000838360058110611dee57fe5b6020020151815260200190815260200160002054848260058110611e0e57fe5b6020020152611e3f565b603a8160040360058110611e2857fe5b6002020154848260058110611e3957fe5b60200201525b600101611dbf565b50505090565b7f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be546001600160a01b038216600090815260496020526040902080546000198101908110611e9757fe5b600091825260209091200154600160801b90046001600160801b03161015611ef05760405162461bcd60e51b81526004018080602001828103825260228152602001806122086022913960400191505060405180910390fd5b6001600160a01b0381166000908152604460205260409020541580611f2d57506001600160a01b0381166000908152604460205260409020546002145b611f7e576040805162461bcd60e51b815260206004820152601b60248201527f4d696e657220697320696e207468652077726f6e672073746174650000000000604482015290519081900360640190fd5b7fa5ae3e2b97d73fb849ea855d27f073b72815b38452d976bd57e4a157827dadd3805460019081019091556040805180820182528281524260208281019182526001600160a01b038616600081815260449092528482209351845591519290940191909155905190917f46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e291a250565b600081831061201c578161201e565b825b9392505050565b600081831161201c578161201e565b61203c612188565b612044612188565b60208301516000805b60058110156120c75785816001016033811061206557fe5b602002015185826005811061207657fe5b60200201526001810184826005811061208b57fe5b60200201528285826005811061209d57fe5b602002015110156120bf578481600581106120b457fe5b602002015192508091505b60010161204d565b5060065b603381101561218057828682603381106120e157fe5b60200201511115612178578581603381106120f857fe5b602002015185836005811061210957fe5b60200201528084836005811061211b57fe5b602002015285816033811061212c57fe5b6020020151925060005b6005811015612176578386826005811061214c57fe5b6020020151101561216e5785816005811061216357fe5b602002015193508092505b600101612136565b505b6001016120cb565b505050915091565b6040518060a001604052806005906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b604051806106600160405280603390602082028036833750919291505056fe4d696e657220776173206e6f74206c6f636b656420666f72207769746864726177616c42616c616e6365206973206c6f776572207468616e207374616b6520616d6f756e747570646174652054656c6c6f722068617320616c7265616479206265656e2072756e4469737075746520686173206265656e20616c726561647920657865637574656454696d6520666f7220766f74696e6720666f72206675727468657220646973707574657320686173206e6f7420706173736564a2646970667358221220480e18e07e88a2ec3beff7e2b13f995de3289ed71dfe0240ee635a483dadbab164736f6c63430007040033",
}

// ExtensionABI is the input ABI used to generate the binding from.
// Deprecated: Use ExtensionMetaData.ABI instead.
var ExtensionABI = ExtensionMetaData.ABI

// Deprecated: Use ExtensionMetaData.Sigs instead.
// ExtensionFuncSigs maps the 4-byte function signature to its string representation.
var ExtensionFuncSigs = ExtensionMetaData.Sigs

// ExtensionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExtensionMetaData.Bin instead.
var ExtensionBin = ExtensionMetaData.Bin

// DeployExtension deploys a new Ethereum contract, binding an instance of Extension to it.
func DeployExtension(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Extension, error) {
	parsed, err := ExtensionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExtensionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Extension{ExtensionCaller: ExtensionCaller{contract: contract}, ExtensionTransactor: ExtensionTransactor{contract: contract}, ExtensionFilterer: ExtensionFilterer{contract: contract}}, nil
}

// Extension is an auto generated Go binding around an Ethereum contract.
type Extension struct {
	ExtensionCaller     // Read-only binding to the contract
	ExtensionTransactor // Write-only binding to the contract
	ExtensionFilterer   // Log filterer for contract events
}

// ExtensionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtensionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtensionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtensionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtensionSession struct {
	Contract     *Extension        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExtensionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtensionCallerSession struct {
	Contract *ExtensionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ExtensionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtensionTransactorSession struct {
	Contract     *ExtensionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ExtensionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtensionRaw struct {
	Contract *Extension // Generic contract binding to access the raw methods on
}

// ExtensionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtensionCallerRaw struct {
	Contract *ExtensionCaller // Generic read-only contract binding to access the raw methods on
}

// ExtensionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtensionTransactorRaw struct {
	Contract *ExtensionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtension creates a new instance of Extension, bound to a specific deployed contract.
func NewExtension(address common.Address, backend bind.ContractBackend) (*Extension, error) {
	contract, err := bindExtension(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Extension{ExtensionCaller: ExtensionCaller{contract: contract}, ExtensionTransactor: ExtensionTransactor{contract: contract}, ExtensionFilterer: ExtensionFilterer{contract: contract}}, nil
}

// NewExtensionCaller creates a new read-only instance of Extension, bound to a specific deployed contract.
func NewExtensionCaller(address common.Address, caller bind.ContractCaller) (*ExtensionCaller, error) {
	contract, err := bindExtension(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtensionCaller{contract: contract}, nil
}

// NewExtensionTransactor creates a new write-only instance of Extension, bound to a specific deployed contract.
func NewExtensionTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtensionTransactor, error) {
	contract, err := bindExtension(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtensionTransactor{contract: contract}, nil
}

// NewExtensionFilterer creates a new log filterer instance of Extension, bound to a specific deployed contract.
func NewExtensionFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtensionFilterer, error) {
	contract, err := bindExtension(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtensionFilterer{contract: contract}, nil
}

// bindExtension binds a generic wrapper to an already deployed contract.
func bindExtension(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExtensionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Extension *ExtensionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Extension.Contract.ExtensionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Extension *ExtensionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extension.Contract.ExtensionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Extension *ExtensionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Extension.Contract.ExtensionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Extension *ExtensionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Extension.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Extension *ExtensionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extension.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Extension *ExtensionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Extension.Contract.contract.Transact(opts, method, params...)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_Extension *ExtensionCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "_allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_Extension *ExtensionSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Extension.Contract.Allowances(&_Extension.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_Extension *ExtensionCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Extension.Contract.Allowances(&_Extension.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_Extension *ExtensionCaller) Addresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_Extension *ExtensionSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _Extension.Contract.Addresses(&_Extension.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_Extension *ExtensionCallerSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _Extension.Contract.Addresses(&_Extension.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_Extension *ExtensionCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "balances", arg0, arg1)

	outstruct := new(struct {
		FromBlock *big.Int
		Value     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_Extension *ExtensionSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _Extension.Contract.Balances(&_Extension.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_Extension *ExtensionCallerSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _Extension.Contract.Balances(&_Extension.CallOpts, arg0, arg1)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_Extension *ExtensionCaller) BytesVars(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "bytesVars", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_Extension *ExtensionSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _Extension.Contract.BytesVars(&_Extension.CallOpts, arg0)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_Extension *ExtensionCallerSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _Extension.Contract.BytesVars(&_Extension.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_Extension *ExtensionCaller) CurrentMiners(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "currentMiners", arg0)

	outstruct := new(struct {
		Value *big.Int
		Miner common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Miner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_Extension *ExtensionSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _Extension.Contract.CurrentMiners(&_Extension.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_Extension *ExtensionCallerSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _Extension.Contract.CurrentMiners(&_Extension.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Extension *ExtensionCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Extension *ExtensionSession) Decimals() (uint8, error) {
	return _Extension.Contract.Decimals(&_Extension.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Extension *ExtensionCallerSession) Decimals() (uint8, error) {
	return _Extension.Contract.Decimals(&_Extension.CallOpts)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_Extension *ExtensionCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "didMine", _challenge, _miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_Extension *ExtensionSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _Extension.Contract.DidMine(&_Extension.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_Extension *ExtensionCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _Extension.Contract.DidMine(&_Extension.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_Extension *ExtensionCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "didVote", _disputeId, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_Extension *ExtensionSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _Extension.Contract.DidVote(&_Extension.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_Extension *ExtensionCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _Extension.Contract.DidVote(&_Extension.CallOpts, _disputeId, _address)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_Extension *ExtensionCaller) DisputeIdByDisputeHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "disputeIdByDisputeHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_Extension *ExtensionSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _Extension.Contract.DisputeIdByDisputeHash(&_Extension.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_Extension *ExtensionCallerSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _Extension.Contract.DisputeIdByDisputeHash(&_Extension.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_Extension *ExtensionCaller) DisputesById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "disputesById", arg0)

	outstruct := new(struct {
		Hash                [32]byte
		Tally               *big.Int
		Executed            bool
		DisputeVotePassed   bool
		IsPropFork          bool
		ReportedMiner       common.Address
		ReportingParty      common.Address
		ProposedForkAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Tally = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Executed = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.DisputeVotePassed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.IsPropFork = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.ReportedMiner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.ReportingParty = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.ProposedForkAddress = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_Extension *ExtensionSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _Extension.Contract.DisputesById(&_Extension.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_Extension *ExtensionCallerSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _Extension.Contract.DisputesById(&_Extension.CallOpts, arg0)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_Extension *ExtensionCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getAddressVars", _data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_Extension *ExtensionSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _Extension.Contract.GetAddressVars(&_Extension.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_Extension *ExtensionCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _Extension.Contract.GetAddressVars(&_Extension.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_Extension *ExtensionCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getAllDisputeVars", _disputeId)

	if err != nil {
		return *new([32]byte), *new(bool), *new(bool), *new(bool), *new(common.Address), *new(common.Address), *new(common.Address), *new([9]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	out6 := *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	out7 := *abi.ConvertType(out[7], new([9]*big.Int)).(*[9]*big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_Extension *ExtensionSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _Extension.Contract.GetAllDisputeVars(&_Extension.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_Extension *ExtensionCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _Extension.Contract.GetAllDisputeVars(&_Extension.CallOpts, _disputeId)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_Extension *ExtensionCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getDisputeIdByDisputeHash", _hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_Extension *ExtensionSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetDisputeIdByDisputeHash(&_Extension.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetDisputeIdByDisputeHash(&_Extension.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_Extension *ExtensionCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getDisputeUintVars", _disputeId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_Extension *ExtensionSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetDisputeUintVars(&_Extension.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetDisputeUintVars(&_Extension.CallOpts, _disputeId, _data)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_Extension *ExtensionCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getLastNewValueById", _requestId)

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
func (_Extension *ExtensionSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _Extension.Contract.GetLastNewValueById(&_Extension.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_Extension *ExtensionCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _Extension.Contract.GetLastNewValueById(&_Extension.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getMinedBlockNum", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetMinedBlockNum(&_Extension.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetMinedBlockNum(&_Extension.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_Extension *ExtensionCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([5]common.Address)).(*[5]common.Address)

	return out0, err

}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_Extension *ExtensionSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _Extension.Contract.GetMinersByRequestIdAndTimestamp(&_Extension.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_Extension *ExtensionCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _Extension.Contract.GetMinersByRequestIdAndTimestamp(&_Extension.CallOpts, _requestId, _timestamp)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _diff, uint256 _tip)
func (_Extension *ExtensionCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Diff       *big.Int
	Tip        *big.Int
}, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getNewCurrentVariables")

	outstruct := new(struct {
		Challenge  [32]byte
		RequestIds [5]*big.Int
		Diff       *big.Int
		Tip        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Challenge = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RequestIds = *abi.ConvertType(out[1], new([5]*big.Int)).(*[5]*big.Int)
	outstruct.Diff = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Tip = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _diff, uint256 _tip)
func (_Extension *ExtensionSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Diff       *big.Int
	Tip        *big.Int
}, error) {
	return _Extension.Contract.GetNewCurrentVariables(&_Extension.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _diff, uint256 _tip)
func (_Extension *ExtensionCallerSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Diff       *big.Int
	Tip        *big.Int
}, error) {
	return _Extension.Contract.GetNewCurrentVariables(&_Extension.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_Extension *ExtensionCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_Extension *ExtensionSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetNewValueCountbyRequestId(&_Extension.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetNewValueCountbyRequestId(&_Extension.CallOpts, _requestId)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_Extension *ExtensionCaller) GetNewVariablesOnDeck(opts *bind.CallOpts) (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getNewVariablesOnDeck")

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
func (_Extension *ExtensionSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _Extension.Contract.GetNewVariablesOnDeck(&_Extension.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_Extension *ExtensionCallerSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _Extension.Contract.GetNewVariablesOnDeck(&_Extension.CallOpts)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_Extension *ExtensionCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getRequestIdByRequestQIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_Extension *ExtensionSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetRequestIdByRequestQIndex(&_Extension.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetRequestIdByRequestQIndex(&_Extension.CallOpts, _index)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_Extension *ExtensionCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getRequestQ")

	if err != nil {
		return *new([51]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([51]*big.Int)).(*[51]*big.Int)

	return out0, err

}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_Extension *ExtensionSession) GetRequestQ() ([51]*big.Int, error) {
	return _Extension.Contract.GetRequestQ(&_Extension.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_Extension *ExtensionCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _Extension.Contract.GetRequestQ(&_Extension.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_Extension *ExtensionCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getRequestUintVars", _requestId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_Extension *ExtensionSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetRequestUintVars(&_Extension.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetRequestUintVars(&_Extension.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_Extension *ExtensionCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getRequestVars", _requestId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_Extension *ExtensionSession) GetRequestVars(_requestId *big.Int) (*big.Int, *big.Int, error) {
	return _Extension.Contract.GetRequestVars(&_Extension.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_Extension *ExtensionCallerSession) GetRequestVars(_requestId *big.Int) (*big.Int, *big.Int, error) {
	return _Extension.Contract.GetRequestVars(&_Extension.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_Extension *ExtensionCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getStakerInfo", _staker)

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
func (_Extension *ExtensionSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _Extension.Contract.GetStakerInfo(&_Extension.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_Extension *ExtensionCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _Extension.Contract.GetStakerInfo(&_Extension.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_Extension *ExtensionCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getSubmissionsByTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_Extension *ExtensionSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _Extension.Contract.GetSubmissionsByTimestamp(&_Extension.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_Extension *ExtensionCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _Extension.Contract.GetSubmissionsByTimestamp(&_Extension.CallOpts, _requestId, _timestamp)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_Extension *ExtensionCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestID, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_Extension *ExtensionSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetTimestampbyRequestIDandIndex(&_Extension.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _Extension.Contract.GetTimestampbyRequestIDandIndex(&_Extension.CallOpts, _requestID, _index)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_Extension *ExtensionCaller) GetTopRequestIDs(opts *bind.CallOpts) ([5]*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getTopRequestIDs")

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_Extension *ExtensionSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _Extension.Contract.GetTopRequestIDs(&_Extension.CallOpts)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_Extension *ExtensionCallerSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _Extension.Contract.GetTopRequestIDs(&_Extension.CallOpts)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_Extension *ExtensionCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "getUintVar", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_Extension *ExtensionSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetUintVar(&_Extension.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_Extension *ExtensionCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _Extension.Contract.GetUintVar(&_Extension.CallOpts, _data)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_Extension *ExtensionCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_Extension *ExtensionSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _Extension.Contract.IsInDispute(&_Extension.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_Extension *ExtensionCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _Extension.Contract.IsInDispute(&_Extension.CallOpts, _requestId, _timestamp)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_Extension *ExtensionCaller) Migrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "migrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_Extension *ExtensionSession) Migrated(arg0 common.Address) (bool, error) {
	return _Extension.Contract.Migrated(&_Extension.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_Extension *ExtensionCallerSession) Migrated(arg0 common.Address) (bool, error) {
	return _Extension.Contract.Migrated(&_Extension.CallOpts, arg0)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_Extension *ExtensionCaller) MinersByChallenge(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "minersByChallenge", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_Extension *ExtensionSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Extension.Contract.MinersByChallenge(&_Extension.CallOpts, arg0, arg1)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_Extension *ExtensionCallerSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Extension.Contract.MinersByChallenge(&_Extension.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Extension *ExtensionCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Extension *ExtensionSession) Name() (string, error) {
	return _Extension.Contract.Name(&_Extension.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Extension *ExtensionCallerSession) Name() (string, error) {
	return _Extension.Contract.Name(&_Extension.CallOpts)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_Extension *ExtensionCaller) NewValueTimestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "newValueTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_Extension *ExtensionSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _Extension.Contract.NewValueTimestamps(&_Extension.CallOpts, arg0)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_Extension *ExtensionCallerSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _Extension.Contract.NewValueTimestamps(&_Extension.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_Extension *ExtensionCaller) RequestIdByQueryHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "requestIdByQueryHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_Extension *ExtensionSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _Extension.Contract.RequestIdByQueryHash(&_Extension.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_Extension *ExtensionCallerSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _Extension.Contract.RequestIdByQueryHash(&_Extension.CallOpts, arg0)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _Extension.Contract.RetrieveData(&_Extension.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_Extension *ExtensionCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _Extension.Contract.RetrieveData(&_Extension.CallOpts, _requestId, _timestamp)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Extension *ExtensionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Extension *ExtensionSession) Symbol() (string, error) {
	return _Extension.Contract.Symbol(&_Extension.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Extension *ExtensionCallerSession) Symbol() (string, error) {
	return _Extension.Contract.Symbol(&_Extension.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Extension *ExtensionCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Extension *ExtensionSession) TotalSupply() (*big.Int, error) {
	return _Extension.Contract.TotalSupply(&_Extension.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Extension *ExtensionCallerSession) TotalSupply() (*big.Int, error) {
	return _Extension.Contract.TotalSupply(&_Extension.CallOpts)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_Extension *ExtensionCaller) Uints(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Extension.contract.Call(opts, &out, "uints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_Extension *ExtensionSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _Extension.Contract.Uints(&_Extension.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_Extension *ExtensionCallerSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _Extension.Contract.Uints(&_Extension.CallOpts, arg0)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Extension *ExtensionTransactor) DepositStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extension.contract.Transact(opts, "depositStake")
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Extension *ExtensionSession) DepositStake() (*types.Transaction, error) {
	return _Extension.Contract.DepositStake(&_Extension.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Extension *ExtensionTransactorSession) DepositStake() (*types.Transaction, error) {
	return _Extension.Contract.DepositStake(&_Extension.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Extension *ExtensionTransactor) RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extension.contract.Transact(opts, "requestStakingWithdraw")
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Extension *ExtensionSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Extension.Contract.RequestStakingWithdraw(&_Extension.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Extension *ExtensionTransactorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Extension.Contract.RequestStakingWithdraw(&_Extension.TransactOpts)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Extension *ExtensionTransactor) TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _Extension.contract.Transact(opts, "tallyVotes", _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Extension *ExtensionSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Extension.Contract.TallyVotes(&_Extension.TransactOpts, _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Extension *ExtensionTransactorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Extension.Contract.TallyVotes(&_Extension.TransactOpts, _disputeId)
}

// UpdateMinDisputeFee is a paid mutator transaction binding the contract method 0x90e5b235.
//
// Solidity: function updateMinDisputeFee() returns()
func (_Extension *ExtensionTransactor) UpdateMinDisputeFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extension.contract.Transact(opts, "updateMinDisputeFee")
}

// UpdateMinDisputeFee is a paid mutator transaction binding the contract method 0x90e5b235.
//
// Solidity: function updateMinDisputeFee() returns()
func (_Extension *ExtensionSession) UpdateMinDisputeFee() (*types.Transaction, error) {
	return _Extension.Contract.UpdateMinDisputeFee(&_Extension.TransactOpts)
}

// UpdateMinDisputeFee is a paid mutator transaction binding the contract method 0x90e5b235.
//
// Solidity: function updateMinDisputeFee() returns()
func (_Extension *ExtensionTransactorSession) UpdateMinDisputeFee() (*types.Transaction, error) {
	return _Extension.Contract.UpdateMinDisputeFee(&_Extension.TransactOpts)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_Extension *ExtensionTransactor) UpdateTellor(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _Extension.contract.Transact(opts, "updateTellor", _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_Extension *ExtensionSession) UpdateTellor(_disputeId *big.Int) (*types.Transaction, error) {
	return _Extension.Contract.UpdateTellor(&_Extension.TransactOpts, _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_Extension *ExtensionTransactorSession) UpdateTellor(_disputeId *big.Int) (*types.Transaction, error) {
	return _Extension.Contract.UpdateTellor(&_Extension.TransactOpts, _disputeId)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Extension *ExtensionTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Extension.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Extension *ExtensionSession) WithdrawStake() (*types.Transaction, error) {
	return _Extension.Contract.WithdrawStake(&_Extension.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Extension *ExtensionTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Extension.Contract.WithdrawStake(&_Extension.TransactOpts)
}

// ExtensionDisputeVoteTalliedIterator is returned from FilterDisputeVoteTallied and is used to iterate over the raw logs and unpacked data for DisputeVoteTallied events raised by the Extension contract.
type ExtensionDisputeVoteTalliedIterator struct {
	Event *ExtensionDisputeVoteTallied // Event containing the contract specifics and raw log

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
func (it *ExtensionDisputeVoteTalliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionDisputeVoteTallied)
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
		it.Event = new(ExtensionDisputeVoteTallied)
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
func (it *ExtensionDisputeVoteTalliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionDisputeVoteTalliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionDisputeVoteTallied represents a DisputeVoteTallied event raised by the Extension contract.
type ExtensionDisputeVoteTallied struct {
	DisputeID      *big.Int
	Result         *big.Int
	ReportedMiner  common.Address
	ReportingParty common.Address
	Passed         bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDisputeVoteTallied is a free log retrieval operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _passed)
func (_Extension *ExtensionFilterer) FilterDisputeVoteTallied(opts *bind.FilterOpts, _disputeID []*big.Int, _reportedMiner []common.Address) (*ExtensionDisputeVoteTalliedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _Extension.contract.FilterLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionDisputeVoteTalliedIterator{contract: _Extension.contract, event: "DisputeVoteTallied", logs: logs, sub: sub}, nil
}

// WatchDisputeVoteTallied is a free log subscription operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _passed)
func (_Extension *ExtensionFilterer) WatchDisputeVoteTallied(opts *bind.WatchOpts, sink chan<- *ExtensionDisputeVoteTallied, _disputeID []*big.Int, _reportedMiner []common.Address) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _Extension.contract.WatchLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionDisputeVoteTallied)
				if err := _Extension.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
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

// ParseDisputeVoteTallied is a log parse operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _passed)
func (_Extension *ExtensionFilterer) ParseDisputeVoteTallied(log types.Log) (*ExtensionDisputeVoteTallied, error) {
	event := new(ExtensionDisputeVoteTallied)
	if err := _Extension.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionNewStakeIterator is returned from FilterNewStake and is used to iterate over the raw logs and unpacked data for NewStake events raised by the Extension contract.
type ExtensionNewStakeIterator struct {
	Event *ExtensionNewStake // Event containing the contract specifics and raw log

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
func (it *ExtensionNewStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionNewStake)
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
		it.Event = new(ExtensionNewStake)
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
func (it *ExtensionNewStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionNewStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionNewStake represents a NewStake event raised by the Extension contract.
type ExtensionNewStake struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewStake is a free log retrieval operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_Extension *ExtensionFilterer) FilterNewStake(opts *bind.FilterOpts, _sender []common.Address) (*ExtensionNewStakeIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Extension.contract.FilterLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionNewStakeIterator{contract: _Extension.contract, event: "NewStake", logs: logs, sub: sub}, nil
}

// WatchNewStake is a free log subscription operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_Extension *ExtensionFilterer) WatchNewStake(opts *bind.WatchOpts, sink chan<- *ExtensionNewStake, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Extension.contract.WatchLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionNewStake)
				if err := _Extension.contract.UnpackLog(event, "NewStake", log); err != nil {
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

// ParseNewStake is a log parse operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_Extension *ExtensionFilterer) ParseNewStake(log types.Log) (*ExtensionNewStake, error) {
	event := new(ExtensionNewStake)
	if err := _Extension.contract.UnpackLog(event, "NewStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionNewTellorAddressIterator is returned from FilterNewTellorAddress and is used to iterate over the raw logs and unpacked data for NewTellorAddress events raised by the Extension contract.
type ExtensionNewTellorAddressIterator struct {
	Event *ExtensionNewTellorAddress // Event containing the contract specifics and raw log

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
func (it *ExtensionNewTellorAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionNewTellorAddress)
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
		it.Event = new(ExtensionNewTellorAddress)
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
func (it *ExtensionNewTellorAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionNewTellorAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionNewTellorAddress represents a NewTellorAddress event raised by the Extension contract.
type ExtensionNewTellorAddress struct {
	NewTellor common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTellorAddress is a free log retrieval operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: event NewTellorAddress(address _newTellor)
func (_Extension *ExtensionFilterer) FilterNewTellorAddress(opts *bind.FilterOpts) (*ExtensionNewTellorAddressIterator, error) {

	logs, sub, err := _Extension.contract.FilterLogs(opts, "NewTellorAddress")
	if err != nil {
		return nil, err
	}
	return &ExtensionNewTellorAddressIterator{contract: _Extension.contract, event: "NewTellorAddress", logs: logs, sub: sub}, nil
}

// WatchNewTellorAddress is a free log subscription operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: event NewTellorAddress(address _newTellor)
func (_Extension *ExtensionFilterer) WatchNewTellorAddress(opts *bind.WatchOpts, sink chan<- *ExtensionNewTellorAddress) (event.Subscription, error) {

	logs, sub, err := _Extension.contract.WatchLogs(opts, "NewTellorAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionNewTellorAddress)
				if err := _Extension.contract.UnpackLog(event, "NewTellorAddress", log); err != nil {
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

// ParseNewTellorAddress is a log parse operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: event NewTellorAddress(address _newTellor)
func (_Extension *ExtensionFilterer) ParseNewTellorAddress(log types.Log) (*ExtensionNewTellorAddress, error) {
	event := new(ExtensionNewTellorAddress)
	if err := _Extension.contract.UnpackLog(event, "NewTellorAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionStakeWithdrawRequestedIterator is returned from FilterStakeWithdrawRequested and is used to iterate over the raw logs and unpacked data for StakeWithdrawRequested events raised by the Extension contract.
type ExtensionStakeWithdrawRequestedIterator struct {
	Event *ExtensionStakeWithdrawRequested // Event containing the contract specifics and raw log

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
func (it *ExtensionStakeWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionStakeWithdrawRequested)
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
		it.Event = new(ExtensionStakeWithdrawRequested)
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
func (it *ExtensionStakeWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionStakeWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionStakeWithdrawRequested represents a StakeWithdrawRequested event raised by the Extension contract.
type ExtensionStakeWithdrawRequested struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawRequested is a free log retrieval operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_Extension *ExtensionFilterer) FilterStakeWithdrawRequested(opts *bind.FilterOpts, _sender []common.Address) (*ExtensionStakeWithdrawRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Extension.contract.FilterLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionStakeWithdrawRequestedIterator{contract: _Extension.contract, event: "StakeWithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawRequested is a free log subscription operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_Extension *ExtensionFilterer) WatchStakeWithdrawRequested(opts *bind.WatchOpts, sink chan<- *ExtensionStakeWithdrawRequested, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Extension.contract.WatchLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionStakeWithdrawRequested)
				if err := _Extension.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
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

// ParseStakeWithdrawRequested is a log parse operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_Extension *ExtensionFilterer) ParseStakeWithdrawRequested(log types.Log) (*ExtensionStakeWithdrawRequested, error) {
	event := new(ExtensionStakeWithdrawRequested)
	if err := _Extension.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the Extension contract.
type ExtensionStakeWithdrawnIterator struct {
	Event *ExtensionStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *ExtensionStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionStakeWithdrawn)
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
		it.Event = new(ExtensionStakeWithdrawn)
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
func (it *ExtensionStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionStakeWithdrawn represents a StakeWithdrawn event raised by the Extension contract.
type ExtensionStakeWithdrawn struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_Extension *ExtensionFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, _sender []common.Address) (*ExtensionStakeWithdrawnIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Extension.contract.FilterLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionStakeWithdrawnIterator{contract: _Extension.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_Extension *ExtensionFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *ExtensionStakeWithdrawn, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Extension.contract.WatchLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionStakeWithdrawn)
				if err := _Extension.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_Extension *ExtensionFilterer) ParseStakeWithdrawn(log types.Log) (*ExtensionStakeWithdrawn, error) {
	event := new(ExtensionStakeWithdrawn)
	if err := _Extension.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorMetaData contains all meta data concerning the ITellor contract.
var ITellorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_result\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reportedMiner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_reportingParty\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_passed\",\"type\":\"bool\"}],\"name\":\"DisputeVoteTallied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_currentRequestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newTellor\",\"type\":\"address\"}],\"name\":\"NewTellorAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_slot\",\"type\":\"uint256\"}],\"name\":\"NonceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newDeity\",\"type\":\"address\"}],\"name\":\"changeDeity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_extension\",\"type\":\"address\"}],\"name\":\"changeExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"changeMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tellorContract\",\"type\":\"address\"}],\"name\":\"changeTellorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"idsOnDeck\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"tipsOnDeck\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopRequestIDs\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_diff\",\"type\":\"uint256\"}],\"name\":\"manuallySetDifficulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_bypass\",\"type\":\"bool\"}],\"name\":\"migrateFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_destination\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amount\",\"type\":\"uint256[]\"}],\"name\":\"migrateForBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_origin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_bypass\",\"type\":\"bool\"}],\"name\":\"migrateFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_origin\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_destination\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amount\",\"type\":\"uint256[]\"}],\"name\":\"migrateFromBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_propNewTellorAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_values\",\"type\":\"uint256[5]\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"}],\"name\":\"testSubmitMiningSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"requests\",\"type\":\"uint256[51]\"}],\"name\":\"testgetMax5\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_max\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_index\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"theLazyCoon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"unlockDisputeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateMinDisputeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"updateTellor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"752d49a1": "addTip(uint256,uint256)",
		"dd62ed3e": "allowance(address,address)",
		"999cf26c": "allowedToTrade(address,uint256)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"4ee2cd7e": "balanceOfAt(address,uint256)",
		"8581af19": "beginDispute(uint256,uint256,uint256)",
		"47abd7f1": "changeDeity(address)",
		"b69a363f": "changeExtension(address)",
		"141e13fa": "changeMigrator(address)",
		"a6f9dae1": "changeOwner(address)",
		"ae0a8279": "changeTellorContract(address)",
		"313ce567": "decimals()",
		"0d2d76a2": "depositStake()",
		"63bb82ad": "didMine(bytes32,address)",
		"a7c438bc": "didVote(uint256,address)",
		"133bee5e": "getAddressVars(bytes32)",
		"af0b1327": "getAllDisputeVars(uint256)",
		"da379941": "getDisputeIdByDisputeHash(bytes32)",
		"7f6fd5d9": "getDisputeUintVars(uint256,bytes32)",
		"fc7cf0a0": "getLastNewValue()",
		"3180f8df": "getLastNewValueById(uint256)",
		"c775b542": "getMinedBlockNum(uint256,uint256)",
		"69026d63": "getMinersByRequestIdAndTimestamp(uint256,uint256)",
		"4049f198": "getNewCurrentVariables()",
		"46eee1c4": "getNewValueCountbyRequestId(uint256)",
		"9a7077ab": "getNewVariablesOnDeck()",
		"6173c0b8": "getRequestIdByRequestQIndex(uint256)",
		"0f0b424d": "getRequestIdByTimestamp(uint256)",
		"b5413029": "getRequestQ()",
		"e0ae93c1": "getRequestUintVars(uint256,bytes32)",
		"e1eee6d6": "getRequestVars(uint256)",
		"733bdef0": "getStakerInfo(address)",
		"11c98512": "getSubmissionsByTimestamp(uint256,uint256)",
		"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
		"fe1cd15d": "getTopRequestIDs()",
		"612c8f7f": "getUintVar(bytes32)",
		"3df0777b": "isInDispute(uint256,uint256)",
		"c52e9539": "manuallySetDifficulty(uint256)",
		"8fd3ab80": "migrate()",
		"a9fa7d34": "migrateFor(address,uint256,bool)",
		"42a89bd6": "migrateForBatch(address[],uint256[])",
		"121dd372": "migrateFrom(address,address,uint256,bool)",
		"8c0f4076": "migrateFromBatch(address[],address[],uint256[])",
		"06fdde03": "name()",
		"26b7d9f6": "proposeFork(address)",
		"28449c3a": "requestStakingWithdraw()",
		"93fa4915": "retrieveData(uint256,uint256)",
		"4350283e": "submitMiningSolution(string,uint256[5],uint256[5])",
		"95d89b41": "symbol()",
		"4d318b0e": "tallyVotes(uint256)",
		"d47f0dd4": "testSubmitMiningSolution(string,uint256[5],uint256[5])",
		"5e93d863": "testgetMax5(uint256[51])",
		"b079f64a": "theLazyCoon(address,uint256)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"9a01ca13": "unlockDisputeFee(uint256)",
		"90e5b235": "updateMinDisputeFee()",
		"f458ab98": "updateTellor(uint256)",
		"fc735e99": "verify()",
		"c9d27afe": "vote(uint256,bool)",
		"bed9d861": "withdrawStake()",
	},
}

// ITellorABI is the input ABI used to generate the binding from.
// Deprecated: Use ITellorMetaData.ABI instead.
var ITellorABI = ITellorMetaData.ABI

// Deprecated: Use ITellorMetaData.Sigs instead.
// ITellorFuncSigs maps the 4-byte function signature to its string representation.
var ITellorFuncSigs = ITellorMetaData.Sigs

// ITellor is an auto generated Go binding around an Ethereum contract.
type ITellor struct {
	ITellorCaller     // Read-only binding to the contract
	ITellorTransactor // Write-only binding to the contract
	ITellorFilterer   // Log filterer for contract events
}

// ITellorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITellorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITellorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITellorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITellorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITellorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITellorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITellorSession struct {
	Contract     *ITellor          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITellorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITellorCallerSession struct {
	Contract *ITellorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ITellorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITellorTransactorSession struct {
	Contract     *ITellorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ITellorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITellorRaw struct {
	Contract *ITellor // Generic contract binding to access the raw methods on
}

// ITellorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITellorCallerRaw struct {
	Contract *ITellorCaller // Generic read-only contract binding to access the raw methods on
}

// ITellorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITellorTransactorRaw struct {
	Contract *ITellorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITellor creates a new instance of ITellor, bound to a specific deployed contract.
func NewITellor(address common.Address, backend bind.ContractBackend) (*ITellor, error) {
	contract, err := bindITellor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITellor{ITellorCaller: ITellorCaller{contract: contract}, ITellorTransactor: ITellorTransactor{contract: contract}, ITellorFilterer: ITellorFilterer{contract: contract}}, nil
}

// NewITellorCaller creates a new read-only instance of ITellor, bound to a specific deployed contract.
func NewITellorCaller(address common.Address, caller bind.ContractCaller) (*ITellorCaller, error) {
	contract, err := bindITellor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITellorCaller{contract: contract}, nil
}

// NewITellorTransactor creates a new write-only instance of ITellor, bound to a specific deployed contract.
func NewITellorTransactor(address common.Address, transactor bind.ContractTransactor) (*ITellorTransactor, error) {
	contract, err := bindITellor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITellorTransactor{contract: contract}, nil
}

// NewITellorFilterer creates a new log filterer instance of ITellor, bound to a specific deployed contract.
func NewITellorFilterer(address common.Address, filterer bind.ContractFilterer) (*ITellorFilterer, error) {
	contract, err := bindITellor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITellorFilterer{contract: contract}, nil
}

// bindITellor binds a generic wrapper to an already deployed contract.
func bindITellor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITellorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITellor *ITellorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITellor.Contract.ITellorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITellor *ITellorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.Contract.ITellorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITellor *ITellorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITellor.Contract.ITellorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITellor *ITellorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITellor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITellor *ITellorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITellor *ITellorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITellor.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ITellor *ITellorCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "allowance", _user, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ITellor *ITellorSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ITellor.Contract.Allowance(&_ITellor.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_ITellor *ITellorCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ITellor.Contract.Allowance(&_ITellor.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_ITellor *ITellorCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "allowedToTrade", _user, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_ITellor *ITellorSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _ITellor.Contract.AllowedToTrade(&_ITellor.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_ITellor *ITellorCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _ITellor.Contract.AllowedToTrade(&_ITellor.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ITellor *ITellorCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ITellor *ITellorSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ITellor.Contract.BalanceOf(&_ITellor.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_ITellor *ITellorCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ITellor.Contract.BalanceOf(&_ITellor.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_ITellor *ITellorCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "balanceOfAt", _user, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_ITellor *ITellorSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _ITellor.Contract.BalanceOfAt(&_ITellor.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_ITellor *ITellorCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _ITellor.Contract.BalanceOfAt(&_ITellor.CallOpts, _user, _blockNumber)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ITellor *ITellorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ITellor *ITellorSession) Decimals() (uint8, error) {
	return _ITellor.Contract.Decimals(&_ITellor.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ITellor *ITellorCallerSession) Decimals() (uint8, error) {
	return _ITellor.Contract.Decimals(&_ITellor.CallOpts)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ITellor *ITellorCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "didMine", _challenge, _miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ITellor *ITellorSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ITellor.Contract.DidMine(&_ITellor.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_ITellor *ITellorCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ITellor.Contract.DidMine(&_ITellor.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ITellor *ITellorCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "didVote", _disputeId, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ITellor *ITellorSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ITellor.Contract.DidVote(&_ITellor.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_ITellor *ITellorCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ITellor.Contract.DidVote(&_ITellor.CallOpts, _disputeId, _address)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ITellor *ITellorCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getAddressVars", _data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ITellor *ITellorSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ITellor.Contract.GetAddressVars(&_ITellor.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_ITellor *ITellorCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ITellor.Contract.GetAddressVars(&_ITellor.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ITellor *ITellorCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getAllDisputeVars", _disputeId)

	if err != nil {
		return *new([32]byte), *new(bool), *new(bool), *new(bool), *new(common.Address), *new(common.Address), *new(common.Address), *new([9]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	out6 := *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	out7 := *abi.ConvertType(out[7], new([9]*big.Int)).(*[9]*big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ITellor *ITellorSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ITellor.Contract.GetAllDisputeVars(&_ITellor.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ITellor *ITellorCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ITellor.Contract.GetAllDisputeVars(&_ITellor.CallOpts, _disputeId)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ITellor *ITellorCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getDisputeIdByDisputeHash", _hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ITellor *ITellorSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetDisputeIdByDisputeHash(&_ITellor.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetDisputeIdByDisputeHash(&_ITellor.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ITellor *ITellorCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getDisputeUintVars", _disputeId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ITellor *ITellorSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetDisputeUintVars(&_ITellor.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetDisputeUintVars(&_ITellor.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ITellor *ITellorCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getLastNewValue")

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ITellor *ITellorSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ITellor.Contract.GetLastNewValue(&_ITellor.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() view returns(uint256, bool)
func (_ITellor *ITellorCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ITellor.Contract.GetLastNewValue(&_ITellor.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ITellor *ITellorCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getLastNewValueById", _requestId)

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
func (_ITellor *ITellorSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ITellor.Contract.GetLastNewValueById(&_ITellor.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_ITellor *ITellorCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ITellor.Contract.GetLastNewValueById(&_ITellor.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getMinedBlockNum", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetMinedBlockNum(&_ITellor.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetMinedBlockNum(&_ITellor.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ITellor *ITellorCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([5]common.Address)).(*[5]common.Address)

	return out0, err

}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ITellor *ITellorSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ITellor.Contract.GetMinersByRequestIdAndTimestamp(&_ITellor.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_ITellor *ITellorCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ITellor.Contract.GetMinersByRequestIdAndTimestamp(&_ITellor.CallOpts, _requestId, _timestamp)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficulty, uint256 _tip)
func (_ITellor *ITellorCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficulty *big.Int
	Tip        *big.Int
}, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getNewCurrentVariables")

	outstruct := new(struct {
		Challenge  [32]byte
		RequestIds [5]*big.Int
		Difficulty *big.Int
		Tip        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Challenge = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RequestIds = *abi.ConvertType(out[1], new([5]*big.Int)).(*[5]*big.Int)
	outstruct.Difficulty = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Tip = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficulty, uint256 _tip)
func (_ITellor *ITellorSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficulty *big.Int
	Tip        *big.Int
}, error) {
	return _ITellor.Contract.GetNewCurrentVariables(&_ITellor.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficulty, uint256 _tip)
func (_ITellor *ITellorCallerSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficulty *big.Int
	Tip        *big.Int
}, error) {
	return _ITellor.Contract.GetNewCurrentVariables(&_ITellor.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ITellor *ITellorCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ITellor *ITellorSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetNewValueCountbyRequestId(&_ITellor.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetNewValueCountbyRequestId(&_ITellor.CallOpts, _requestId)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_ITellor *ITellorCaller) GetNewVariablesOnDeck(opts *bind.CallOpts) (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getNewVariablesOnDeck")

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
func (_ITellor *ITellorSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _ITellor.Contract.GetNewVariablesOnDeck(&_ITellor.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_ITellor *ITellorCallerSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _ITellor.Contract.GetNewVariablesOnDeck(&_ITellor.CallOpts)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ITellor *ITellorCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getRequestIdByRequestQIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ITellor *ITellorSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetRequestIdByRequestQIndex(&_ITellor.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetRequestIdByRequestQIndex(&_ITellor.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getRequestIdByTimestamp", _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetRequestIdByTimestamp(&_ITellor.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetRequestIdByTimestamp(&_ITellor.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ITellor *ITellorCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getRequestQ")

	if err != nil {
		return *new([51]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([51]*big.Int)).(*[51]*big.Int)

	return out0, err

}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ITellor *ITellorSession) GetRequestQ() ([51]*big.Int, error) {
	return _ITellor.Contract.GetRequestQ(&_ITellor.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_ITellor *ITellorCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _ITellor.Contract.GetRequestQ(&_ITellor.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ITellor *ITellorCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getRequestUintVars", _requestId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ITellor *ITellorSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetRequestUintVars(&_ITellor.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetRequestUintVars(&_ITellor.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_ITellor *ITellorCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getRequestVars", _requestId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_ITellor *ITellorSession) GetRequestVars(_requestId *big.Int) (*big.Int, *big.Int, error) {
	return _ITellor.Contract.GetRequestVars(&_ITellor.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_ITellor *ITellorCallerSession) GetRequestVars(_requestId *big.Int) (*big.Int, *big.Int, error) {
	return _ITellor.Contract.GetRequestVars(&_ITellor.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ITellor *ITellorCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getStakerInfo", _staker)

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
func (_ITellor *ITellorSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ITellor.Contract.GetStakerInfo(&_ITellor.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_ITellor *ITellorCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ITellor.Contract.GetStakerInfo(&_ITellor.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ITellor *ITellorCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getSubmissionsByTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ITellor *ITellorSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ITellor.Contract.GetSubmissionsByTimestamp(&_ITellor.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_ITellor *ITellorCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ITellor.Contract.GetSubmissionsByTimestamp(&_ITellor.CallOpts, _requestId, _timestamp)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ITellor *ITellorCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestID, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ITellor *ITellorSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetTimestampbyRequestIDandIndex(&_ITellor.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ITellor.Contract.GetTimestampbyRequestIDandIndex(&_ITellor.CallOpts, _requestID, _index)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_ITellor *ITellorCaller) GetTopRequestIDs(opts *bind.CallOpts) ([5]*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getTopRequestIDs")

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_ITellor *ITellorSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _ITellor.Contract.GetTopRequestIDs(&_ITellor.CallOpts)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_ITellor *ITellorCallerSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _ITellor.Contract.GetTopRequestIDs(&_ITellor.CallOpts)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ITellor *ITellorCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "getUintVar", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ITellor *ITellorSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetUintVar(&_ITellor.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_ITellor *ITellorCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ITellor.Contract.GetUintVar(&_ITellor.CallOpts, _data)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ITellor *ITellorCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ITellor *ITellorSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ITellor.Contract.IsInDispute(&_ITellor.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_ITellor *ITellorCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ITellor.Contract.IsInDispute(&_ITellor.CallOpts, _requestId, _timestamp)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ITellor *ITellorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ITellor *ITellorSession) Name() (string, error) {
	return _ITellor.Contract.Name(&_ITellor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_ITellor *ITellorCallerSession) Name() (string, error) {
	return _ITellor.Contract.Name(&_ITellor.CallOpts)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ITellor.Contract.RetrieveData(&_ITellor.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_ITellor *ITellorCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ITellor.Contract.RetrieveData(&_ITellor.CallOpts, _requestId, _timestamp)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_ITellor *ITellorCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_ITellor *ITellorSession) Symbol() (string, error) {
	return _ITellor.Contract.Symbol(&_ITellor.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_ITellor *ITellorCallerSession) Symbol() (string, error) {
	return _ITellor.Contract.Symbol(&_ITellor.CallOpts)
}

// TestgetMax5 is a free data retrieval call binding the contract method 0x5e93d863.
//
// Solidity: function testgetMax5(uint256[51] requests) view returns(uint256[5] _max, uint256[5] _index)
func (_ITellor *ITellorCaller) TestgetMax5(opts *bind.CallOpts, requests [51]*big.Int) (struct {
	Max   [5]*big.Int
	Index [5]*big.Int
}, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "testgetMax5", requests)

	outstruct := new(struct {
		Max   [5]*big.Int
		Index [5]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Max = *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)
	outstruct.Index = *abi.ConvertType(out[1], new([5]*big.Int)).(*[5]*big.Int)

	return *outstruct, err

}

// TestgetMax5 is a free data retrieval call binding the contract method 0x5e93d863.
//
// Solidity: function testgetMax5(uint256[51] requests) view returns(uint256[5] _max, uint256[5] _index)
func (_ITellor *ITellorSession) TestgetMax5(requests [51]*big.Int) (struct {
	Max   [5]*big.Int
	Index [5]*big.Int
}, error) {
	return _ITellor.Contract.TestgetMax5(&_ITellor.CallOpts, requests)
}

// TestgetMax5 is a free data retrieval call binding the contract method 0x5e93d863.
//
// Solidity: function testgetMax5(uint256[51] requests) view returns(uint256[5] _max, uint256[5] _index)
func (_ITellor *ITellorCallerSession) TestgetMax5(requests [51]*big.Int) (struct {
	Max   [5]*big.Int
	Index [5]*big.Int
}, error) {
	return _ITellor.Contract.TestgetMax5(&_ITellor.CallOpts, requests)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ITellor *ITellorCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ITellor.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ITellor *ITellorSession) TotalSupply() (*big.Int, error) {
	return _ITellor.Contract.TotalSupply(&_ITellor.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ITellor *ITellorCallerSession) TotalSupply() (*big.Int, error) {
	return _ITellor.Contract.TotalSupply(&_ITellor.CallOpts)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_ITellor *ITellorTransactor) AddTip(opts *bind.TransactOpts, _requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "addTip", _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_ITellor *ITellorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.AddTip(&_ITellor.TransactOpts, _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_ITellor *ITellorTransactorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.AddTip(&_ITellor.TransactOpts, _requestId, _tip)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ITellor *ITellorTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ITellor *ITellorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.Approve(&_ITellor.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ITellor *ITellorTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.Approve(&_ITellor.TransactOpts, _spender, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_ITellor *ITellorTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_ITellor *ITellorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.BeginDispute(&_ITellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_ITellor *ITellorTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.BeginDispute(&_ITellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_ITellor *ITellorTransactor) ChangeDeity(opts *bind.TransactOpts, _newDeity common.Address) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "changeDeity", _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_ITellor *ITellorSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeDeity(&_ITellor.TransactOpts, _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(address _newDeity) returns()
func (_ITellor *ITellorTransactorSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeDeity(&_ITellor.TransactOpts, _newDeity)
}

// ChangeExtension is a paid mutator transaction binding the contract method 0xb69a363f.
//
// Solidity: function changeExtension(address _extension) returns()
func (_ITellor *ITellorTransactor) ChangeExtension(opts *bind.TransactOpts, _extension common.Address) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "changeExtension", _extension)
}

// ChangeExtension is a paid mutator transaction binding the contract method 0xb69a363f.
//
// Solidity: function changeExtension(address _extension) returns()
func (_ITellor *ITellorSession) ChangeExtension(_extension common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeExtension(&_ITellor.TransactOpts, _extension)
}

// ChangeExtension is a paid mutator transaction binding the contract method 0xb69a363f.
//
// Solidity: function changeExtension(address _extension) returns()
func (_ITellor *ITellorTransactorSession) ChangeExtension(_extension common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeExtension(&_ITellor.TransactOpts, _extension)
}

// ChangeMigrator is a paid mutator transaction binding the contract method 0x141e13fa.
//
// Solidity: function changeMigrator(address _migrator) returns()
func (_ITellor *ITellorTransactor) ChangeMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "changeMigrator", _migrator)
}

// ChangeMigrator is a paid mutator transaction binding the contract method 0x141e13fa.
//
// Solidity: function changeMigrator(address _migrator) returns()
func (_ITellor *ITellorSession) ChangeMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeMigrator(&_ITellor.TransactOpts, _migrator)
}

// ChangeMigrator is a paid mutator transaction binding the contract method 0x141e13fa.
//
// Solidity: function changeMigrator(address _migrator) returns()
func (_ITellor *ITellorTransactorSession) ChangeMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeMigrator(&_ITellor.TransactOpts, _migrator)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_ITellor *ITellorTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_ITellor *ITellorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeOwner(&_ITellor.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_ITellor *ITellorTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeOwner(&_ITellor.TransactOpts, _newOwner)
}

// ChangeTellorContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeTellorContract(address _tellorContract) returns()
func (_ITellor *ITellorTransactor) ChangeTellorContract(opts *bind.TransactOpts, _tellorContract common.Address) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "changeTellorContract", _tellorContract)
}

// ChangeTellorContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeTellorContract(address _tellorContract) returns()
func (_ITellor *ITellorSession) ChangeTellorContract(_tellorContract common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeTellorContract(&_ITellor.TransactOpts, _tellorContract)
}

// ChangeTellorContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeTellorContract(address _tellorContract) returns()
func (_ITellor *ITellorTransactorSession) ChangeTellorContract(_tellorContract common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ChangeTellorContract(&_ITellor.TransactOpts, _tellorContract)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_ITellor *ITellorTransactor) DepositStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "depositStake")
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_ITellor *ITellorSession) DepositStake() (*types.Transaction, error) {
	return _ITellor.Contract.DepositStake(&_ITellor.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_ITellor *ITellorTransactorSession) DepositStake() (*types.Transaction, error) {
	return _ITellor.Contract.DepositStake(&_ITellor.TransactOpts)
}

// ManuallySetDifficulty is a paid mutator transaction binding the contract method 0xc52e9539.
//
// Solidity: function manuallySetDifficulty(uint256 _diff) returns()
func (_ITellor *ITellorTransactor) ManuallySetDifficulty(opts *bind.TransactOpts, _diff *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "manuallySetDifficulty", _diff)
}

// ManuallySetDifficulty is a paid mutator transaction binding the contract method 0xc52e9539.
//
// Solidity: function manuallySetDifficulty(uint256 _diff) returns()
func (_ITellor *ITellorSession) ManuallySetDifficulty(_diff *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.ManuallySetDifficulty(&_ITellor.TransactOpts, _diff)
}

// ManuallySetDifficulty is a paid mutator transaction binding the contract method 0xc52e9539.
//
// Solidity: function manuallySetDifficulty(uint256 _diff) returns()
func (_ITellor *ITellorTransactorSession) ManuallySetDifficulty(_diff *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.ManuallySetDifficulty(&_ITellor.TransactOpts, _diff)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_ITellor *ITellorTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_ITellor *ITellorSession) Migrate() (*types.Transaction, error) {
	return _ITellor.Contract.Migrate(&_ITellor.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_ITellor *ITellorTransactorSession) Migrate() (*types.Transaction, error) {
	return _ITellor.Contract.Migrate(&_ITellor.TransactOpts)
}

// MigrateFor is a paid mutator transaction binding the contract method 0xa9fa7d34.
//
// Solidity: function migrateFor(address _destination, uint256 _amount, bool _bypass) returns()
func (_ITellor *ITellorTransactor) MigrateFor(opts *bind.TransactOpts, _destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "migrateFor", _destination, _amount, _bypass)
}

// MigrateFor is a paid mutator transaction binding the contract method 0xa9fa7d34.
//
// Solidity: function migrateFor(address _destination, uint256 _amount, bool _bypass) returns()
func (_ITellor *ITellorSession) MigrateFor(_destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateFor(&_ITellor.TransactOpts, _destination, _amount, _bypass)
}

// MigrateFor is a paid mutator transaction binding the contract method 0xa9fa7d34.
//
// Solidity: function migrateFor(address _destination, uint256 _amount, bool _bypass) returns()
func (_ITellor *ITellorTransactorSession) MigrateFor(_destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateFor(&_ITellor.TransactOpts, _destination, _amount, _bypass)
}

// MigrateForBatch is a paid mutator transaction binding the contract method 0x42a89bd6.
//
// Solidity: function migrateForBatch(address[] _destination, uint256[] _amount) returns()
func (_ITellor *ITellorTransactor) MigrateForBatch(opts *bind.TransactOpts, _destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "migrateForBatch", _destination, _amount)
}

// MigrateForBatch is a paid mutator transaction binding the contract method 0x42a89bd6.
//
// Solidity: function migrateForBatch(address[] _destination, uint256[] _amount) returns()
func (_ITellor *ITellorSession) MigrateForBatch(_destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateForBatch(&_ITellor.TransactOpts, _destination, _amount)
}

// MigrateForBatch is a paid mutator transaction binding the contract method 0x42a89bd6.
//
// Solidity: function migrateForBatch(address[] _destination, uint256[] _amount) returns()
func (_ITellor *ITellorTransactorSession) MigrateForBatch(_destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateForBatch(&_ITellor.TransactOpts, _destination, _amount)
}

// MigrateFrom is a paid mutator transaction binding the contract method 0x121dd372.
//
// Solidity: function migrateFrom(address _origin, address _destination, uint256 _amount, bool _bypass) returns()
func (_ITellor *ITellorTransactor) MigrateFrom(opts *bind.TransactOpts, _origin common.Address, _destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "migrateFrom", _origin, _destination, _amount, _bypass)
}

// MigrateFrom is a paid mutator transaction binding the contract method 0x121dd372.
//
// Solidity: function migrateFrom(address _origin, address _destination, uint256 _amount, bool _bypass) returns()
func (_ITellor *ITellorSession) MigrateFrom(_origin common.Address, _destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateFrom(&_ITellor.TransactOpts, _origin, _destination, _amount, _bypass)
}

// MigrateFrom is a paid mutator transaction binding the contract method 0x121dd372.
//
// Solidity: function migrateFrom(address _origin, address _destination, uint256 _amount, bool _bypass) returns()
func (_ITellor *ITellorTransactorSession) MigrateFrom(_origin common.Address, _destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateFrom(&_ITellor.TransactOpts, _origin, _destination, _amount, _bypass)
}

// MigrateFromBatch is a paid mutator transaction binding the contract method 0x8c0f4076.
//
// Solidity: function migrateFromBatch(address[] _origin, address[] _destination, uint256[] _amount) returns()
func (_ITellor *ITellorTransactor) MigrateFromBatch(opts *bind.TransactOpts, _origin []common.Address, _destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "migrateFromBatch", _origin, _destination, _amount)
}

// MigrateFromBatch is a paid mutator transaction binding the contract method 0x8c0f4076.
//
// Solidity: function migrateFromBatch(address[] _origin, address[] _destination, uint256[] _amount) returns()
func (_ITellor *ITellorSession) MigrateFromBatch(_origin []common.Address, _destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateFromBatch(&_ITellor.TransactOpts, _origin, _destination, _amount)
}

// MigrateFromBatch is a paid mutator transaction binding the contract method 0x8c0f4076.
//
// Solidity: function migrateFromBatch(address[] _origin, address[] _destination, uint256[] _amount) returns()
func (_ITellor *ITellorTransactorSession) MigrateFromBatch(_origin []common.Address, _destination []common.Address, _amount []*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.MigrateFromBatch(&_ITellor.TransactOpts, _origin, _destination, _amount)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_ITellor *ITellorTransactor) ProposeFork(opts *bind.TransactOpts, _propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "proposeFork", _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_ITellor *ITellorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ProposeFork(&_ITellor.TransactOpts, _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_ITellor *ITellorTransactorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _ITellor.Contract.ProposeFork(&_ITellor.TransactOpts, _propNewTellorAddress)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_ITellor *ITellorTransactor) RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "requestStakingWithdraw")
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_ITellor *ITellorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _ITellor.Contract.RequestStakingWithdraw(&_ITellor.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_ITellor *ITellorTransactorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _ITellor.Contract.RequestStakingWithdraw(&_ITellor.TransactOpts)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestIds, uint256[5] _values) returns()
func (_ITellor *ITellorTransactor) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestIds [5]*big.Int, _values [5]*big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "submitMiningSolution", _nonce, _requestIds, _values)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestIds, uint256[5] _values) returns()
func (_ITellor *ITellorSession) SubmitMiningSolution(_nonce string, _requestIds [5]*big.Int, _values [5]*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.SubmitMiningSolution(&_ITellor.TransactOpts, _nonce, _requestIds, _values)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestIds, uint256[5] _values) returns()
func (_ITellor *ITellorTransactorSession) SubmitMiningSolution(_nonce string, _requestIds [5]*big.Int, _values [5]*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.SubmitMiningSolution(&_ITellor.TransactOpts, _nonce, _requestIds, _values)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_ITellor *ITellorTransactor) TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "tallyVotes", _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_ITellor *ITellorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TallyVotes(&_ITellor.TransactOpts, _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_ITellor *ITellorTransactorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TallyVotes(&_ITellor.TransactOpts, _disputeId)
}

// TestSubmitMiningSolution is a paid mutator transaction binding the contract method 0xd47f0dd4.
//
// Solidity: function testSubmitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_ITellor *ITellorTransactor) TestSubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "testSubmitMiningSolution", _nonce, _requestId, _value)
}

// TestSubmitMiningSolution is a paid mutator transaction binding the contract method 0xd47f0dd4.
//
// Solidity: function testSubmitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_ITellor *ITellorSession) TestSubmitMiningSolution(_nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TestSubmitMiningSolution(&_ITellor.TransactOpts, _nonce, _requestId, _value)
}

// TestSubmitMiningSolution is a paid mutator transaction binding the contract method 0xd47f0dd4.
//
// Solidity: function testSubmitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_ITellor *ITellorTransactorSession) TestSubmitMiningSolution(_nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TestSubmitMiningSolution(&_ITellor.TransactOpts, _nonce, _requestId, _value)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(address _address, uint256 _amount) returns()
func (_ITellor *ITellorTransactor) TheLazyCoon(opts *bind.TransactOpts, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "theLazyCoon", _address, _amount)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(address _address, uint256 _amount) returns()
func (_ITellor *ITellorSession) TheLazyCoon(_address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TheLazyCoon(&_ITellor.TransactOpts, _address, _amount)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(address _address, uint256 _amount) returns()
func (_ITellor *ITellorTransactorSession) TheLazyCoon(_address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TheLazyCoon(&_ITellor.TransactOpts, _address, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_ITellor *ITellorTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_ITellor *ITellorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.Transfer(&_ITellor.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_ITellor *ITellorTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.Transfer(&_ITellor.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_ITellor *ITellorTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_ITellor *ITellorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TransferFrom(&_ITellor.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_ITellor *ITellorTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.TransferFrom(&_ITellor.TransactOpts, _from, _to, _amount)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_ITellor *ITellorTransactor) UnlockDisputeFee(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "unlockDisputeFee", _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_ITellor *ITellorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.UnlockDisputeFee(&_ITellor.TransactOpts, _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_ITellor *ITellorTransactorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.UnlockDisputeFee(&_ITellor.TransactOpts, _disputeId)
}

// UpdateMinDisputeFee is a paid mutator transaction binding the contract method 0x90e5b235.
//
// Solidity: function updateMinDisputeFee() returns()
func (_ITellor *ITellorTransactor) UpdateMinDisputeFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "updateMinDisputeFee")
}

// UpdateMinDisputeFee is a paid mutator transaction binding the contract method 0x90e5b235.
//
// Solidity: function updateMinDisputeFee() returns()
func (_ITellor *ITellorSession) UpdateMinDisputeFee() (*types.Transaction, error) {
	return _ITellor.Contract.UpdateMinDisputeFee(&_ITellor.TransactOpts)
}

// UpdateMinDisputeFee is a paid mutator transaction binding the contract method 0x90e5b235.
//
// Solidity: function updateMinDisputeFee() returns()
func (_ITellor *ITellorTransactorSession) UpdateMinDisputeFee() (*types.Transaction, error) {
	return _ITellor.Contract.UpdateMinDisputeFee(&_ITellor.TransactOpts)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_ITellor *ITellorTransactor) UpdateTellor(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "updateTellor", _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_ITellor *ITellorSession) UpdateTellor(_disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.UpdateTellor(&_ITellor.TransactOpts, _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_ITellor *ITellorTransactorSession) UpdateTellor(_disputeId *big.Int) (*types.Transaction, error) {
	return _ITellor.Contract.UpdateTellor(&_ITellor.TransactOpts, _disputeId)
}

// Verify is a paid mutator transaction binding the contract method 0xfc735e99.
//
// Solidity: function verify() returns(uint256)
func (_ITellor *ITellorTransactor) Verify(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "verify")
}

// Verify is a paid mutator transaction binding the contract method 0xfc735e99.
//
// Solidity: function verify() returns(uint256)
func (_ITellor *ITellorSession) Verify() (*types.Transaction, error) {
	return _ITellor.Contract.Verify(&_ITellor.TransactOpts)
}

// Verify is a paid mutator transaction binding the contract method 0xfc735e99.
//
// Solidity: function verify() returns(uint256)
func (_ITellor *ITellorTransactorSession) Verify() (*types.Transaction, error) {
	return _ITellor.Contract.Verify(&_ITellor.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_ITellor *ITellorTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_ITellor *ITellorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _ITellor.Contract.Vote(&_ITellor.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_ITellor *ITellorTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _ITellor.Contract.Vote(&_ITellor.TransactOpts, _disputeId, _supportsDispute)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_ITellor *ITellorTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITellor.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_ITellor *ITellorSession) WithdrawStake() (*types.Transaction, error) {
	return _ITellor.Contract.WithdrawStake(&_ITellor.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_ITellor *ITellorTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _ITellor.Contract.WithdrawStake(&_ITellor.TransactOpts)
}

// ITellorApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ITellor contract.
type ITellorApprovalIterator struct {
	Event *ITellorApproval // Event containing the contract specifics and raw log

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
func (it *ITellorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorApproval)
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
		it.Event = new(ITellorApproval)
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
func (it *ITellorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorApproval represents a Approval event raised by the ITellor contract.
type ITellorApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ITellor *ITellorFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ITellorApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ITellorApprovalIterator{contract: _ITellor.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ITellor *ITellorFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ITellorApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorApproval)
				if err := _ITellor.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ITellor *ITellorFilterer) ParseApproval(log types.Log) (*ITellorApproval, error) {
	event := new(ITellorApproval)
	if err := _ITellor.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorDisputeVoteTalliedIterator is returned from FilterDisputeVoteTallied and is used to iterate over the raw logs and unpacked data for DisputeVoteTallied events raised by the ITellor contract.
type ITellorDisputeVoteTalliedIterator struct {
	Event *ITellorDisputeVoteTallied // Event containing the contract specifics and raw log

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
func (it *ITellorDisputeVoteTalliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorDisputeVoteTallied)
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
		it.Event = new(ITellorDisputeVoteTallied)
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
func (it *ITellorDisputeVoteTalliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorDisputeVoteTalliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorDisputeVoteTallied represents a DisputeVoteTallied event raised by the ITellor contract.
type ITellorDisputeVoteTallied struct {
	DisputeID      *big.Int
	Result         *big.Int
	ReportedMiner  common.Address
	ReportingParty common.Address
	Passed         bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDisputeVoteTallied is a free log retrieval operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _passed)
func (_ITellor *ITellorFilterer) FilterDisputeVoteTallied(opts *bind.FilterOpts, _disputeID []*big.Int, _reportedMiner []common.Address) (*ITellorDisputeVoteTalliedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return &ITellorDisputeVoteTalliedIterator{contract: _ITellor.contract, event: "DisputeVoteTallied", logs: logs, sub: sub}, nil
}

// WatchDisputeVoteTallied is a free log subscription operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _passed)
func (_ITellor *ITellorFilterer) WatchDisputeVoteTallied(opts *bind.WatchOpts, sink chan<- *ITellorDisputeVoteTallied, _disputeID []*big.Int, _reportedMiner []common.Address) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorDisputeVoteTallied)
				if err := _ITellor.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
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

// ParseDisputeVoteTallied is a log parse operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: event DisputeVoteTallied(uint256 indexed _disputeID, int256 _result, address indexed _reportedMiner, address _reportingParty, bool _passed)
func (_ITellor *ITellorFilterer) ParseDisputeVoteTallied(log types.Log) (*ITellorDisputeVoteTallied, error) {
	event := new(ITellorDisputeVoteTallied)
	if err := _ITellor.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorNewChallengeIterator is returned from FilterNewChallenge and is used to iterate over the raw logs and unpacked data for NewChallenge events raised by the ITellor contract.
type ITellorNewChallengeIterator struct {
	Event *ITellorNewChallenge // Event containing the contract specifics and raw log

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
func (it *ITellorNewChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorNewChallenge)
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
		it.Event = new(ITellorNewChallenge)
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
func (it *ITellorNewChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorNewChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorNewChallenge represents a NewChallenge event raised by the ITellor contract.
type ITellorNewChallenge struct {
	CurrentChallenge [32]byte
	CurrentRequestId [5]*big.Int
	Difficulty       *big.Int
	TotalTips        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewChallenge is a free log retrieval operation binding the contract event 0x1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c1408.
//
// Solidity: event NewChallenge(bytes32 indexed _currentChallenge, uint256[5] _currentRequestId, uint256 _difficulty, uint256 _totalTips)
func (_ITellor *ITellorFilterer) FilterNewChallenge(opts *bind.FilterOpts, _currentChallenge [][32]byte) (*ITellorNewChallengeIterator, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "NewChallenge", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &ITellorNewChallengeIterator{contract: _ITellor.contract, event: "NewChallenge", logs: logs, sub: sub}, nil
}

// WatchNewChallenge is a free log subscription operation binding the contract event 0x1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c1408.
//
// Solidity: event NewChallenge(bytes32 indexed _currentChallenge, uint256[5] _currentRequestId, uint256 _difficulty, uint256 _totalTips)
func (_ITellor *ITellorFilterer) WatchNewChallenge(opts *bind.WatchOpts, sink chan<- *ITellorNewChallenge, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "NewChallenge", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorNewChallenge)
				if err := _ITellor.contract.UnpackLog(event, "NewChallenge", log); err != nil {
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

// ParseNewChallenge is a log parse operation binding the contract event 0x1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c1408.
//
// Solidity: event NewChallenge(bytes32 indexed _currentChallenge, uint256[5] _currentRequestId, uint256 _difficulty, uint256 _totalTips)
func (_ITellor *ITellorFilterer) ParseNewChallenge(log types.Log) (*ITellorNewChallenge, error) {
	event := new(ITellorNewChallenge)
	if err := _ITellor.contract.UnpackLog(event, "NewChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorNewDisputeIterator is returned from FilterNewDispute and is used to iterate over the raw logs and unpacked data for NewDispute events raised by the ITellor contract.
type ITellorNewDisputeIterator struct {
	Event *ITellorNewDispute // Event containing the contract specifics and raw log

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
func (it *ITellorNewDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorNewDispute)
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
		it.Event = new(ITellorNewDispute)
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
func (it *ITellorNewDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorNewDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorNewDispute represents a NewDispute event raised by the ITellor contract.
type ITellorNewDispute struct {
	DisputeId *big.Int
	RequestId *big.Int
	Timestamp *big.Int
	Miner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewDispute is a free log retrieval operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_ITellor *ITellorFilterer) FilterNewDispute(opts *bind.FilterOpts, _disputeId []*big.Int, _requestId []*big.Int) (*ITellorNewDisputeIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ITellorNewDisputeIterator{contract: _ITellor.contract, event: "NewDispute", logs: logs, sub: sub}, nil
}

// WatchNewDispute is a free log subscription operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_ITellor *ITellorFilterer) WatchNewDispute(opts *bind.WatchOpts, sink chan<- *ITellorNewDispute, _disputeId []*big.Int, _requestId []*big.Int) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorNewDispute)
				if err := _ITellor.contract.UnpackLog(event, "NewDispute", log); err != nil {
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

// ParseNewDispute is a log parse operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_ITellor *ITellorFilterer) ParseNewDispute(log types.Log) (*ITellorNewDispute, error) {
	event := new(ITellorNewDispute)
	if err := _ITellor.contract.UnpackLog(event, "NewDispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorNewStakeIterator is returned from FilterNewStake and is used to iterate over the raw logs and unpacked data for NewStake events raised by the ITellor contract.
type ITellorNewStakeIterator struct {
	Event *ITellorNewStake // Event containing the contract specifics and raw log

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
func (it *ITellorNewStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorNewStake)
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
		it.Event = new(ITellorNewStake)
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
func (it *ITellorNewStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorNewStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorNewStake represents a NewStake event raised by the ITellor contract.
type ITellorNewStake struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewStake is a free log retrieval operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_ITellor *ITellorFilterer) FilterNewStake(opts *bind.FilterOpts, _sender []common.Address) (*ITellorNewStakeIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ITellorNewStakeIterator{contract: _ITellor.contract, event: "NewStake", logs: logs, sub: sub}, nil
}

// WatchNewStake is a free log subscription operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_ITellor *ITellorFilterer) WatchNewStake(opts *bind.WatchOpts, sink chan<- *ITellorNewStake, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorNewStake)
				if err := _ITellor.contract.UnpackLog(event, "NewStake", log); err != nil {
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

// ParseNewStake is a log parse operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: event NewStake(address indexed _sender)
func (_ITellor *ITellorFilterer) ParseNewStake(log types.Log) (*ITellorNewStake, error) {
	event := new(ITellorNewStake)
	if err := _ITellor.contract.UnpackLog(event, "NewStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorNewTellorAddressIterator is returned from FilterNewTellorAddress and is used to iterate over the raw logs and unpacked data for NewTellorAddress events raised by the ITellor contract.
type ITellorNewTellorAddressIterator struct {
	Event *ITellorNewTellorAddress // Event containing the contract specifics and raw log

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
func (it *ITellorNewTellorAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorNewTellorAddress)
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
		it.Event = new(ITellorNewTellorAddress)
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
func (it *ITellorNewTellorAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorNewTellorAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorNewTellorAddress represents a NewTellorAddress event raised by the ITellor contract.
type ITellorNewTellorAddress struct {
	NewTellor common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTellorAddress is a free log retrieval operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: event NewTellorAddress(address _newTellor)
func (_ITellor *ITellorFilterer) FilterNewTellorAddress(opts *bind.FilterOpts) (*ITellorNewTellorAddressIterator, error) {

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "NewTellorAddress")
	if err != nil {
		return nil, err
	}
	return &ITellorNewTellorAddressIterator{contract: _ITellor.contract, event: "NewTellorAddress", logs: logs, sub: sub}, nil
}

// WatchNewTellorAddress is a free log subscription operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: event NewTellorAddress(address _newTellor)
func (_ITellor *ITellorFilterer) WatchNewTellorAddress(opts *bind.WatchOpts, sink chan<- *ITellorNewTellorAddress) (event.Subscription, error) {

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "NewTellorAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorNewTellorAddress)
				if err := _ITellor.contract.UnpackLog(event, "NewTellorAddress", log); err != nil {
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

// ParseNewTellorAddress is a log parse operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: event NewTellorAddress(address _newTellor)
func (_ITellor *ITellorFilterer) ParseNewTellorAddress(log types.Log) (*ITellorNewTellorAddress, error) {
	event := new(ITellorNewTellorAddress)
	if err := _ITellor.contract.UnpackLog(event, "NewTellorAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorNewValueIterator is returned from FilterNewValue and is used to iterate over the raw logs and unpacked data for NewValue events raised by the ITellor contract.
type ITellorNewValueIterator struct {
	Event *ITellorNewValue // Event containing the contract specifics and raw log

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
func (it *ITellorNewValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorNewValue)
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
		it.Event = new(ITellorNewValue)
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
func (it *ITellorNewValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorNewValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorNewValue represents a NewValue event raised by the ITellor contract.
type ITellorNewValue struct {
	RequestId        [5]*big.Int
	Time             *big.Int
	Value            [5]*big.Int
	TotalTips        *big.Int
	CurrentChallenge [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewValue is a free log retrieval operation binding the contract event 0xbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc45.
//
// Solidity: event NewValue(uint256[5] _requestId, uint256 _time, uint256[5] _value, uint256 _totalTips, bytes32 indexed _currentChallenge)
func (_ITellor *ITellorFilterer) FilterNewValue(opts *bind.FilterOpts, _currentChallenge [][32]byte) (*ITellorNewValueIterator, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "NewValue", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &ITellorNewValueIterator{contract: _ITellor.contract, event: "NewValue", logs: logs, sub: sub}, nil
}

// WatchNewValue is a free log subscription operation binding the contract event 0xbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc45.
//
// Solidity: event NewValue(uint256[5] _requestId, uint256 _time, uint256[5] _value, uint256 _totalTips, bytes32 indexed _currentChallenge)
func (_ITellor *ITellorFilterer) WatchNewValue(opts *bind.WatchOpts, sink chan<- *ITellorNewValue, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "NewValue", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorNewValue)
				if err := _ITellor.contract.UnpackLog(event, "NewValue", log); err != nil {
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

// ParseNewValue is a log parse operation binding the contract event 0xbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc45.
//
// Solidity: event NewValue(uint256[5] _requestId, uint256 _time, uint256[5] _value, uint256 _totalTips, bytes32 indexed _currentChallenge)
func (_ITellor *ITellorFilterer) ParseNewValue(log types.Log) (*ITellorNewValue, error) {
	event := new(ITellorNewValue)
	if err := _ITellor.contract.UnpackLog(event, "NewValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorNonceSubmittedIterator is returned from FilterNonceSubmitted and is used to iterate over the raw logs and unpacked data for NonceSubmitted events raised by the ITellor contract.
type ITellorNonceSubmittedIterator struct {
	Event *ITellorNonceSubmitted // Event containing the contract specifics and raw log

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
func (it *ITellorNonceSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorNonceSubmitted)
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
		it.Event = new(ITellorNonceSubmitted)
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
func (it *ITellorNonceSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorNonceSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorNonceSubmitted represents a NonceSubmitted event raised by the ITellor contract.
type ITellorNonceSubmitted struct {
	Miner            common.Address
	Nonce            string
	RequestId        [5]*big.Int
	Value            [5]*big.Int
	CurrentChallenge [32]byte
	Slot             *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNonceSubmitted is a free log retrieval operation binding the contract event 0x9d2e5f03fc65aff196e0f3a8dd924b24099de487e8cffc888921d420ab196e39.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256[5] _requestId, uint256[5] _value, bytes32 indexed _currentChallenge, uint256 _slot)
func (_ITellor *ITellorFilterer) FilterNonceSubmitted(opts *bind.FilterOpts, _miner []common.Address, _currentChallenge [][32]byte) (*ITellorNonceSubmittedIterator, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "NonceSubmitted", _minerRule, _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &ITellorNonceSubmittedIterator{contract: _ITellor.contract, event: "NonceSubmitted", logs: logs, sub: sub}, nil
}

// WatchNonceSubmitted is a free log subscription operation binding the contract event 0x9d2e5f03fc65aff196e0f3a8dd924b24099de487e8cffc888921d420ab196e39.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256[5] _requestId, uint256[5] _value, bytes32 indexed _currentChallenge, uint256 _slot)
func (_ITellor *ITellorFilterer) WatchNonceSubmitted(opts *bind.WatchOpts, sink chan<- *ITellorNonceSubmitted, _miner []common.Address, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "NonceSubmitted", _minerRule, _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorNonceSubmitted)
				if err := _ITellor.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
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

// ParseNonceSubmitted is a log parse operation binding the contract event 0x9d2e5f03fc65aff196e0f3a8dd924b24099de487e8cffc888921d420ab196e39.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256[5] _requestId, uint256[5] _value, bytes32 indexed _currentChallenge, uint256 _slot)
func (_ITellor *ITellorFilterer) ParseNonceSubmitted(log types.Log) (*ITellorNonceSubmitted, error) {
	event := new(ITellorNonceSubmitted)
	if err := _ITellor.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorStakeWithdrawRequestedIterator is returned from FilterStakeWithdrawRequested and is used to iterate over the raw logs and unpacked data for StakeWithdrawRequested events raised by the ITellor contract.
type ITellorStakeWithdrawRequestedIterator struct {
	Event *ITellorStakeWithdrawRequested // Event containing the contract specifics and raw log

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
func (it *ITellorStakeWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorStakeWithdrawRequested)
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
		it.Event = new(ITellorStakeWithdrawRequested)
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
func (it *ITellorStakeWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorStakeWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorStakeWithdrawRequested represents a StakeWithdrawRequested event raised by the ITellor contract.
type ITellorStakeWithdrawRequested struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawRequested is a free log retrieval operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_ITellor *ITellorFilterer) FilterStakeWithdrawRequested(opts *bind.FilterOpts, _sender []common.Address) (*ITellorStakeWithdrawRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ITellorStakeWithdrawRequestedIterator{contract: _ITellor.contract, event: "StakeWithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawRequested is a free log subscription operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_ITellor *ITellorFilterer) WatchStakeWithdrawRequested(opts *bind.WatchOpts, sink chan<- *ITellorStakeWithdrawRequested, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorStakeWithdrawRequested)
				if err := _ITellor.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
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

// ParseStakeWithdrawRequested is a log parse operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: event StakeWithdrawRequested(address indexed _sender)
func (_ITellor *ITellorFilterer) ParseStakeWithdrawRequested(log types.Log) (*ITellorStakeWithdrawRequested, error) {
	event := new(ITellorStakeWithdrawRequested)
	if err := _ITellor.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the ITellor contract.
type ITellorStakeWithdrawnIterator struct {
	Event *ITellorStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *ITellorStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorStakeWithdrawn)
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
		it.Event = new(ITellorStakeWithdrawn)
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
func (it *ITellorStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorStakeWithdrawn represents a StakeWithdrawn event raised by the ITellor contract.
type ITellorStakeWithdrawn struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_ITellor *ITellorFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, _sender []common.Address) (*ITellorStakeWithdrawnIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ITellorStakeWithdrawnIterator{contract: _ITellor.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_ITellor *ITellorFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *ITellorStakeWithdrawn, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorStakeWithdrawn)
				if err := _ITellor.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: event StakeWithdrawn(address indexed _sender)
func (_ITellor *ITellorFilterer) ParseStakeWithdrawn(log types.Log) (*ITellorStakeWithdrawn, error) {
	event := new(ITellorStakeWithdrawn)
	if err := _ITellor.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorTipAddedIterator is returned from FilterTipAdded and is used to iterate over the raw logs and unpacked data for TipAdded events raised by the ITellor contract.
type ITellorTipAddedIterator struct {
	Event *ITellorTipAdded // Event containing the contract specifics and raw log

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
func (it *ITellorTipAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorTipAdded)
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
		it.Event = new(ITellorTipAdded)
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
func (it *ITellorTipAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorTipAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorTipAdded represents a TipAdded event raised by the ITellor contract.
type ITellorTipAdded struct {
	Sender    common.Address
	RequestId *big.Int
	Tip       *big.Int
	TotalTips *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTipAdded is a free log retrieval operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_ITellor *ITellorFilterer) FilterTipAdded(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*ITellorTipAddedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ITellorTipAddedIterator{contract: _ITellor.contract, event: "TipAdded", logs: logs, sub: sub}, nil
}

// WatchTipAdded is a free log subscription operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_ITellor *ITellorFilterer) WatchTipAdded(opts *bind.WatchOpts, sink chan<- *ITellorTipAdded, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorTipAdded)
				if err := _ITellor.contract.UnpackLog(event, "TipAdded", log); err != nil {
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

// ParseTipAdded is a log parse operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_ITellor *ITellorFilterer) ParseTipAdded(log types.Log) (*ITellorTipAdded, error) {
	event := new(ITellorTipAdded)
	if err := _ITellor.contract.UnpackLog(event, "TipAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the ITellor contract.
type ITellorTransferredIterator struct {
	Event *ITellorTransferred // Event containing the contract specifics and raw log

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
func (it *ITellorTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorTransferred)
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
		it.Event = new(ITellorTransferred)
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
func (it *ITellorTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorTransferred represents a Transferred event raised by the ITellor contract.
type ITellorTransferred struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ITellor *ITellorFilterer) FilterTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ITellorTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ITellorTransferredIterator{contract: _ITellor.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ITellor *ITellorFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *ITellorTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorTransferred)
				if err := _ITellor.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransferred is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ITellor *ITellorFilterer) ParseTransferred(log types.Log) (*ITellorTransferred, error) {
	event := new(ITellorTransferred)
	if err := _ITellor.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITellorVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the ITellor contract.
type ITellorVotedIterator struct {
	Event *ITellorVoted // Event containing the contract specifics and raw log

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
func (it *ITellorVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITellorVoted)
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
		it.Event = new(ITellorVoted)
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
func (it *ITellorVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITellorVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITellorVoted represents a Voted event raised by the ITellor contract.
type ITellorVoted struct {
	DisputeID  *big.Int
	Position   bool
	Voter      common.Address
	VoteWeight *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_ITellor *ITellorFilterer) FilterVoted(opts *bind.FilterOpts, _disputeID []*big.Int, _voter []common.Address, _voteWeight []*big.Int) (*ITellorVotedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}
	var _voteWeightRule []interface{}
	for _, _voteWeightItem := range _voteWeight {
		_voteWeightRule = append(_voteWeightRule, _voteWeightItem)
	}

	logs, sub, err := _ITellor.contract.FilterLogs(opts, "Voted", _disputeIDRule, _voterRule, _voteWeightRule)
	if err != nil {
		return nil, err
	}
	return &ITellorVotedIterator{contract: _ITellor.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_ITellor *ITellorFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *ITellorVoted, _disputeID []*big.Int, _voter []common.Address, _voteWeight []*big.Int) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}
	var _voteWeightRule []interface{}
	for _, _voteWeightItem := range _voteWeight {
		_voteWeightRule = append(_voteWeightRule, _voteWeightItem)
	}

	logs, sub, err := _ITellor.contract.WatchLogs(opts, "Voted", _disputeIDRule, _voterRule, _voteWeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITellorVoted)
				if err := _ITellor.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_ITellor *ITellorFilterer) ParseVoted(log types.Log) (*ITellorVoted, error) {
	event := new(ITellorVoted)
	if err := _ITellor.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e64e60812cc978977df6d5cee5ae4f2040c2a98a6540e103509b6ed18aaad34e64736f6c63430007040033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TellorMetaData contains all meta data concerning the Tellor contract.
var TellorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ext\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_currentRequestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"indexed\":false,\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_slot\",\"type\":\"uint256\"}],\"name\":\"NonceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"fromBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentMiners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"disputesById\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"tally\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disputeVotePassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPropFork\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"reportedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reportingParty\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedForkAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_bypass\",\"type\":\"bool\"}],\"name\":\"migrateFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minersByChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newValueTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_propNewTellorAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_values\",\"type\":\"uint256[5]\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"unlockDisputeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"024c2ddd": "_allowances(address,address)",
		"752d49a1": "addTip(uint256,uint256)",
		"699f200f": "addresses(bytes32)",
		"dd62ed3e": "allowance(address,address)",
		"999cf26c": "allowedToTrade(address,uint256)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"4ee2cd7e": "balanceOfAt(address,uint256)",
		"cbf1304d": "balances(address,uint256)",
		"8581af19": "beginDispute(uint256,uint256,uint256)",
		"62dd1d2a": "bytesVars(bytes32)",
		"1fd22364": "currentMiners(uint256)",
		"d01f4d9e": "disputeIdByDisputeHash(bytes32)",
		"db085beb": "disputesById(uint256)",
		"8fd3ab80": "migrate()",
		"a9fa7d34": "migrateFor(address,uint256,bool)",
		"4ba0a5ee": "migrated(address)",
		"48b18e54": "minersByChallenge(bytes32,address)",
		"438c0aa3": "newValueTimestamps(uint256)",
		"26b7d9f6": "proposeFork(address)",
		"5700242c": "requestIdByQueryHash(bytes32)",
		"4350283e": "submitMiningSolution(string,uint256[5],uint256[5])",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"b59e14d4": "uints(bytes32)",
		"9a01ca13": "unlockDisputeFee(uint256)",
		"fc735e99": "verify()",
		"c9d27afe": "vote(uint256,bool)",
	},
	Bin: "0x60a060405234801561001057600080fd5b5060405162004e1538038062004e158339818101604052602081101561003557600080fd5b5051606081901b6001600160601b0319166080526001600160a01b0316614dab6200006a600039806101c85250614dab6000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c8063752d49a1116100f9578063b59e14d411610097578063d01f4d9e11610071578063d01f4d9e146106c4578063db085beb146106e1578063dd62ed3e1461074e578063fc735e991461077c576101c4565b8063b59e14d414610627578063c9d27afe14610644578063cbf1304d14610669576101c4565b8063999cf26c116100d3578063999cf26c1461057e5780639a01ca13146105aa578063a9059cbb146105c7578063a9fa7d34146105f3576101c4565b8063752d49a11461052a5780638581af191461054d5780638fd3ab8014610576576101c4565b806348b18e54116101665780635700242c116101405780635700242c1461049157806362dd1d2a146104ae578063699f200f146104cb57806370a0823114610504576101c4565b806348b18e54146104135780634ba0a5ee1461043f5780634ee2cd7e14610465576101c4565b806323b872dd116101a257806323b872dd1461032457806326b7d9f61461035a5780634350283e14610382578063438c0aa3146103f6576101c4565b8063024c2ddd14610266578063095ea7b3146102a65780631fd22364146102e6575b60007f000000000000000000000000000000000000000000000000000000000000000090506000816001600160a01b03166000366040518083838082843760405192019450600093509091505080830381855af49150503d8060008114610247576040519150601f19603f3d011682016040523d82523d6000602084013e61024c565b606091505b505090503d6000803e808015610261573d6000f35b3d6000fd5b6102946004803603604081101561027c57600080fd5b506001600160a01b0381358116916020013516610784565b60408051918252519081900360200190f35b6102d2600480360360408110156102bc57600080fd5b506001600160a01b0381351690602001356107a1565b604080519115158252519081900360200190f35b610303600480360360208110156102fc57600080fd5b5035610850565b604080519283526001600160a01b0390911660208301528051918290030190f35b6102d26004803603606081101561033a57600080fd5b506001600160a01b0381358116916020810135909116906040013561087b565b6103806004803603602081101561037057600080fd5b50356001600160a01b0316610927565b005b610380600480360361016081101561039957600080fd5b8101906020810181356401000000008111156103b457600080fd5b8201836020820111156103c657600080fd5b803590602001918460018302840111640100000000831117156103e857600080fd5b919350915060a08101610c2c565b6102946004803603602081101561040c57600080fd5b5035610dd9565b6102d26004803603604081101561042957600080fd5b50803590602001356001600160a01b0316610dfa565b6102d26004803603602081101561045557600080fd5b50356001600160a01b0316610e1a565b6102946004803603604081101561047b57600080fd5b506001600160a01b038135169060200135610e2f565b610294600480360360208110156104a757600080fd5b5035610fda565b610294600480360360208110156104c457600080fd5b5035610fec565b6104e8600480360360208110156104e157600080fd5b5035610ffe565b604080516001600160a01b039092168252519081900360200190f35b6102946004803603602081101561051a57600080fd5b50356001600160a01b0316611019565b6103806004803603604081101561054057600080fd5b5080359060200135611025565b6103806004803603606081101561056357600080fd5b5080359060208101359060400135611235565b6103806117a4565b6102d26004803603604081101561059457600080fd5b506001600160a01b0381351690602001356117af565b610380600480360360208110156105c057600080fd5b503561187b565b6102d2600480360360408110156105dd57600080fd5b506001600160a01b038135169060200135612111565b6103806004803603606081101561060957600080fd5b506001600160a01b0381351690602081013590604001351515612127565b6102946004803603602081101561063d57600080fd5b50356121cc565b6103806004803603604081101561065a57600080fd5b508035906020013515156121de565b6106956004803603604081101561067f57600080fd5b506001600160a01b0381351690602001356124c0565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b610294600480360360208110156106da57600080fd5b5035612503565b6106fe600480360360208110156106f757600080fd5b5035612515565b60408051988952602089019790975294151587870152921515606087015290151560808601526001600160a01b0390811660a086015290811660c08501521660e083015251908190036101000190f35b6102946004803603604081101561076457600080fd5b506001600160a01b0381358116916020013516612571565b61029461259c565b604a60209081526000928352604080842090915290825290205481565b60006001600160a01b0383166107e85760405162461bcd60e51b8152600401808060200182810382526022815260200180614ac36022913960400191505060405180910390fd5b336000818152604a602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b603a816005811061086057600080fd5b6002020180546001909101549091506001600160a01b031682565b6001600160a01b0383166000908152604a602090815260408083203384529091528120548211156108e8576040805162461bcd60e51b8152602060048201526012602482015271416c6c6f77616e63652069732077726f6e6760701b604482015290519081900360640190fd5b6001600160a01b0384166000908152604a6020908152604080832033845290915290208054839003905561091d8484846125a2565b5060019392505050565b7fd051321aa26ce60d202f153d0c0e67687e975532ab88ce92d84f18e39895d90760005260466020527f7042aca2505aa9a8428c1710298edbac96e027637d6417651c8aa8f88ed1ca0a54156109b3576040805162461bcd60e51b815260206004820152600c60248201526b6e6f2072656e7472616e637960a01b604482015290519081900360640190fd5b7fd051321aa26ce60d202f153d0c0e67687e975532ab88ce92d84f18e39895d907600052604660205260017f7042aca2505aa9a8428c1710298edbac96e027637d6417651c8aa8f88ed1ca0a55610a0981612756565b7fd051321aa26ce60d202f153d0c0e67687e975532ab88ce92d84f18e39895d9076000908152604660209081527f7042aca2505aa9a8428c1710298edbac96e027637d6417651c8aa8f88ed1ca0a829055604080516001600160a01b038516818401528151808203840181529082018252805190830120600080516020614c39833981519152805460010190819055818552603890935292205415610af657600082815260386020908152604080832054848452603683528184207fed92b4c1e0a9e559a31171d487ecbec963526662038ecfa3a71160bd62fb8733855260050190925290912055610b08565b60008281526038602052604090208190555b600082815260386020526040812054610b229083906128ab565b6000838152603660205260408120858155600280820180546201000062ff000019909116176301000000600160b81b0319163363010000008102919091179091556003830180546001600160a01b031990811683179091556004840180549091166001600160a01b038b16179055600190920192909255919250610bba919030906000198501900a68056bc75e2d63100000026125a2565b5060009081526036602090815260408083207f4b4cefd5ced7569ef0d091282b4bca9c52a034c56471a6061afd1bf307a2de7c84526005019091528082204390557f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d28825290204262093a800190555050565b6040805133602080830191909152825180830382018152918301835281519181019190912060008181526046909252919020541580610c7f57506000818152604660205260409020546103844291909103115b610cba5760405162461bcd60e51b815260040180806020018281038252602a815260200180614a50602a913960400191505060405180910390fd5b7fdfbec46864bc123768f0d134913175d9577a55bb71b9b2595fda21e21f36b0826000526046602052600080516020614b5183398151915254600414610d3957610d3985858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612a9d92505050565b6000818152604660209081526040918290204290558151601f8701829004820281018201909252858252610dd291908790879081908401838280828437600092019190915250506040805160a081810190925292508791506005908390839080828437600092019190915250506040805160a0818101909252915086906005908390839080828437600092019190915250612d62915050565b5050505050565b60338181548110610de957600080fd5b600091825260209091200154905081565b603960209081526000928352604080842090915290825290205460ff1681565b604b6020526000908152604090205460ff1681565b6001600160a01b038216600090815260496020526040812080541580610e7557508281600081548110610e5e57fe5b6000918252602090912001546001600160801b0316115b15610e8457600091505061084a565b805481906000198101908110610e9657fe5b6000918252602090912001546001600160801b03168310610ee857805481906000198101908110610ec357fe5b600091825260209091200154600160801b90046001600160801b0316915061084a9050565b8054600090600119015b81811115610fa0576000600260018385010104905085848281548110610f1457fe5b6000918252602090912001546001600160801b03161415610f6357838181548110610f3b57fe5b600091825260209091200154600160801b90046001600160801b0316945061084a9350505050565b85848281548110610f7057fe5b6000918252602090912001546001600160801b03161015610f9357809250610f9a565b6001810391505b50610ef2565b828281548110610fac57fe5b600091825260209091200154600160801b90046001600160801b0316935061084a92505050565b5092915050565b60376020526000908152604090205481565b60486020526000908152604090205481565b6047602052600090815260409020546001600160a01b031681565b600061084a8243610e2f565b81611068576040805162461bcd60e51b815260206004820152600e60248201526d052657175657374496420697320360941b604482015290519081900360640190fd5b806110ba576040805162461bcd60e51b815260206004820152601c60248201527f5469702073686f756c642062652067726561746572207468616e203000000000604482015290519081900360640190fd5b7f3f8b5616fa9e7f2ce4a868fde15c58b92e77bc1acd6769bf1567629a3dc4c86560005260466020527f7119b9afaa3bda0901ffe121c1535f50cd6d0d09df5d29eb1cb16c8ab47a55d65460010182811415611162577f3f8b5616fa9e7f2ce4a868fde15c58b92e77bc1acd6769bf1567629a3dc4c86560005260466020527f7119b9afaa3bda0901ffe121c1535f50cd6d0d09df5d29eb1cb16c8ab47a55d68190556111b6565b8083106111b6576040805162461bcd60e51b815260206004820181905260248201527f526571756573744964206973206e6f74206c657373207468616e20636f756e74604482015290519081900360640190fd5b6111c033836133eb565b6111ca8383613553565b6000838152604560209081526040808320600080516020614d568339815191528452600101825291829020548251858152918201528151859233927fd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820929081900390910190a3505050565b60008381526045602090815260408083208584526002810190925290912054611298576040805162461bcd60e51b815260206004820152601060248201526f04d696e656420626c6f636b20697320360841b604482015290519081900360640190fd5b600582106112e4576040805162461bcd60e51b81526020600482015260146024820152734d696e657220696e6465782069732077726f6e6760601b604482015290519081900360640190fd5b6000838152600580830160205260408220908490811061130057fe5b015460408051606083811b6001600160601b03191660208084018290526b111a5cdc1d5d1950dbdd5b9d60a21b603485015284518085038201815284860186528051908201206000908152604682528581208054600190810190915593850192909252607484018b905260948085018b90528551808603909101815260b49094018552835193810193909320600080516020614c39833981519152805490930192839055808252603890935292909220546001600160a01b0390931693509180156114095760008281526036602090815260408083207fed92b4c1e0a9e559a31171d487ecbec963526662038ecfa3a71160bd62fb873384526005019091529020819055611460565b62093a808742031061144c5760405162461bcd60e51b8152600401808060200182810382526032815260200180614cb96032913960400191505060405180910390fd5b506000828152603860205260409020819055805b600061146c83836128ab565b9050600087600214156114ff57506000898152604560209081526040808320600080516020614b108339815191528452600190810183529083208054909101908190557f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc97909252604690527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be540261154e565b507f675d2171f68d6f5545d54fb9b1fb61a0e6897e6188ca1cd664e7c9530d91ecfc60005260466020527f3e5522f19747f0f285b96ded572ac4128c3a764aea9f44058dc0afc9dda449865481025b6000848152603660209081526040808320888155600281018054600383018054336001600160a01b031991821617909155600484018054909116905562010000600160b81b03191663010000006001600160a01b038d16021761ffff19169055600181018490557f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f8452600590810183528184208e9055600080516020614bb883398151915284528184208d90558c845260068b01909252909120908990811061161457fe5b015460008581526036602090815260408083207f9147231ab14efb72c38117f68521ddef8de64f092c18c69dbfb602ffc4de7f478452600501909152808220929092557f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d2881528181206202a3008502420190557f4b4cefd5ced7569ef0d091282b4bca9c52a034c56471a6061afd1bf307a2de7c81528181204390557f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d81528181208a90557f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc40981522081905561170b3330836125a2565b876002141561173d5760008981526004880160209081526040808320805460ff1916600117905560038a019091528120555b6001600160a01b0386166000818152604460209081526040918290206003905581518c81529081019290925280518c9287927feceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da6492918290030190a350505050505050505050565b6117ad33613761565b565b6001600160a01b038216600090815260446020526040812054158015906117ee57506001600160a01b0383166000908152604460205260409020546005115b15611868577f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc9760005260466020527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be5482906118539061184d86611019565b906138b0565b106118605750600161084a565b50600061084a565b8161187284611019565b10159392505050565b600080516020614b108339815191526000526046602052600080516020614c39833981519152548111156118ef576040805162461bcd60e51b8152602060048201526016602482015275191a5cdc1d5d1948191bd95cc81b9bdd08195e1a5cdd60521b604482015290519081900360640190fd5b600081815260366020818152604080842054845260388252808420548085529282528084207f6ab2b18aafe78fd59c6a4092015bddd9fcacb8170f72b299074f74d76a91a92385526005018083528185205482518085019190915282518082038501815290830183528051908401208552909152909120548061196f5750805b60008281526036602090815260408083208484528184207f6ab2b18aafe78fd59c6a4092015bddd9fcacb8170f72b299074f74d76a91a923855260058201909352922054806119bc575060015b7f29169706298d2b6df50a532e958b56426de1465348b93650fca42d456eaec5fc600090815260058401602052604081205415611a33576040805162461bcd60e51b815260206004820152601060248201526f185b1c9958591e481c185a59081bdd5d60821b604482015290519081900360640190fd5b600284015462010000900460ff1615611a7d5760405162461bcd60e51b8152600401808060200182810382526029815260200180614a9a6029913960400191505060405180910390fd5b7ff9e1ae10923bfc79f52e309baf8c7699edb821f91ef5b5bd07be29545917b3a66000908152600584016020526040902054611b00576040805162461bcd60e51b815260206004820152601860248201527f766f7465206e6565647320746f2062652074616c6c6965640000000000000000604482015290519081900360640190fd5b7ff9e1ae10923bfc79f52e309baf8c7699edb821f91ef5b5bd07be29545917b3a6600090815260058401602052604090205462015180429190910311611b775760405162461bcd60e51b815260040180806020018281038252602b815260200180614ae5602b913960400191505060405180910390fd5b600284810154630100000090046001600160a01b031660009081526044602090815260408083207f29169706298d2b6df50a532e958b56426de1465348b93650fca42d456eaec5fc84526005890190925290912060019081905591850154909161010090910460ff1615151415611da0576201518042064203600182015560466020527fa5ae3e2b97d73fb849ea855d27f073b72815b38452d976bd57e4a157827dadd380546000190190557f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc976000527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be5460028601548190611c8990630100000090046001600160a01b0316611019565b1015611cb0576002860154611cad90630100000090046001600160a01b0316611019565b90505b815460041415611ceb576005825560028601546003870154611ce6916001600160a01b03630100000090910481169116836125a2565b600082555b60005b84811015611d9957604080518287036020808301919091528251808303820181529183018352815191810191909120600090815260058a019091522054935083611d36578893505b600084815260366020908152604080832060038101547f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc409855260058201909352922054611d909130916001600160a01b03909116906125a2565b50600101611cee565b505061200c565b600284015460408051630100000090920460601b6001600160601b0319166020808401919091526b111a5cdc1d5d1950dbdd5b9d60a21b603484015281518084038201815292820182528251928101929092206000908152604690925290205460011415611e0d57600181555b7f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f60009081526005860160208181526040808420548452604582528084207f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d8552929091529091205460021415611ed6577f9147231ab14efb72c38117f68521ddef8de64f092c18c69dbfb602ffc4de7f476000908152600587016020908152604080832054600080516020614bb8833981519152845281842054845260038501909252909120555b600080516020614bb8833981519152600090815260058701602090815260408083205483526004840190915290205460ff16151560011415611f4957600080516020614bb883398151915260009081526005870160209081526040808320548352600484019091529020805460ff191690555b60005b8481101561200957604080518287036020808301919091528251808303820181529183018352815191810191909120600090815260058a01909152205493508315611fa257600084815260366020526040902095505b600286015460008581526036602090815260408083207f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc409845260050190915290205461200191309163010000009091046001600160a01b0316906125a2565b600101611f4c565b50505b60028481015460408051630100000090920460601b6001600160601b0319166020808401919091526b111a5cdc1d5d1950dbdd5b9d60a21b6034840152815180840382018152928201825282519281019290922060009081526046835281812080546000190190557f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d8152600589019092529020541415612107577f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f6000908152600586016020908152604080832054835260458252808320600080516020614b108339815191528452600101909152902080546000190190555b5050505050505050565b600061211e3384846125a2565b50600192915050565b7f5fc094d10c65bc33cc842217b2eccca0191ff24148319da094e540a55989896160005260476020527f437dd27c2043efdfef03344e9331c924985f7bd1752abef5ea93bdbfed685100546001600160a01b031633146121bc576040805162461bcd60e51b815260206004820152600b60248201526a1b9bdd08185b1b1bddd95960aa1b604482015290519081900360640190fd5b6121c78383836138c2565b505050565b60466020526000908152604090205481565b600080516020614b108339815191526000526046602052600080516020614c3983398151915254821115612252576040805162461bcd60e51b8152602060048201526016602482015275191a5cdc1d5d1948191bd95cc81b9bdd08195e1a5cdd60521b604482015290519081900360640190fd5b6000828152603660205260409020600281015460ff16156122a45760405162461bcd60e51b8152600401808060200182810382526025815260200180614d0c6025913960400191505060405180910390fd5b7f4b4cefd5ced7569ef0d091282b4bca9c52a034c56471a6061afd1bf307a2de7c60009081526005820160205260408120546122e1903390610e2f565b33600090815260068401602052604090205490915060ff16151560011415612350576040805162461bcd60e51b815260206004820152601860248201527f53656e6465722068617320616c726561647920766f7465640000000000000000604482015290519081900360640190fd5b80612396576040805162461bcd60e51b81526020600482015260116024820152700557365722062616c616e6365206973203607c1b604482015290519081900360640190fd5b33600090815260446020526040902054600314156123f4576040805162461bcd60e51b81526020600482015260166024820152754d696e657220697320756e646572206469737075746560501b604482015290519081900360640190fd5b3360009081526006830160209081526040808320805460ff191660019081179091557f1da378694063870452ce03b189f48e04c1aa026348e74e6c86e10738514ad2c484526005860190925290912080549091019055821561246957600182015461245f9082613958565b600183015561247e565b60018201546124789082613983565b60018301555b60408051841515815290518291339187917f911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e919081900360200190a450505050565b604960205281600052604060002081815481106124dc57600080fd5b6000918252602090912001546001600160801b038082169350600160801b90910416905082565b60386020526000908152604090205481565b603660205260009081526040902080546001820154600283015460038401546004909401549293919260ff808316936101008404821693620100008104909216926001600160a01b036301000000909304831692918216911688565b6001600160a01b039182166000908152604a6020908152604080832093909416825291909152205490565b610bb790565b806125de5760405162461bcd60e51b8152600401808060200182810382526021815260200180614bf86021913960400191505060405180910390fd5b6001600160a01b038216612631576040805162461bcd60e51b815260206004820152601560248201527452656365697665722069732030206164647265737360581b604482015290519081900360640190fd5b61263b83826117af565b6126765760405162461bcd60e51b8152600401808060200182810382526027815260200180614b716027913960400191505060405180910390fd5b600061268184611019565b905081612690858284036139a9565b61269984611019565b9150816001600160801b03168183016001600160801b031610156126f8576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b612704848284016139a9565b836001600160a01b0316856001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a35050505050565b60408051602060248201819052600060448301819052835180840360640181526084909301845290820180516001600160e01b031663fc735e9960e01b1781529251825191936060936001600160a01b0387169390929182918083835b602083106127d25780518252601f1990920191602091820191016127b3565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114612834576040519150601f19603f3d011682016040523d82523d6000602084013e612839565b606091505b50915091508180156128625750610bb781806020019051602081101561285e57600080fd5b5051115b6121c7576040805162461bcd60e51b81526020600482015260156024820152741b995dc81d195b1b1bdc881a5cc81a5b9d985b1a59605a1b604482015290519081900360640190fd5b60008181526036602081815260408084207f6ab2b18aafe78fd59c6a4092015bddd9fcacb8170f72b299074f74d76a91a923855260050180835281852080546001019081905586865293835281518084018590528251808203850181529083018352805190840120855290915290912083905582821461084a57600082815260366020818152604080842081516000198701818501528251808203850181529083018352805190840120855260059081018352818520548086529383528185207f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d28865201909152909120544210156129ea576040805162461bcd60e51b815260206004820152601760248201527f4469737075746520697320616c7265616479206f70656e000000000000000000604482015290519081900360640190fd5b60008181526036602052604090206002015460ff1615610fd35760008181526036602090815260408083207ff9e1ae10923bfc79f52e309baf8c7699edb821f91ef5b5bd07be29545917b3a684526005019091529020546201518042919091031115610fd3576040805162461bcd60e51b815260206004820152601f60248201527f54696d6520666f7220766f74696e6720686176656e277420656c617073656400604482015290519081900360640190fd5b600080516020614c79833981519152547fd54702836c9d21d0727ffacc3e39f57c92b5ae0f50177e593bfb5ec66e3de28060005260486020908152600080516020614c99833981519152546040805180840183815233606081901b93830193909352865160029560039594938993926054909101918401908083835b60208310612b385780518252601f199092019160209182019101612b19565b6001836020036101000a038019825116818451168082178552505050505050905001935050505060405160208183030381529060405280519060200120604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b60208310612bc35780518252601f199092019160209182019101612ba4565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015612c02573d6000803e3d6000fd5b5050506040515160601b60405160200180826bffffffffffffffffffffffff191681526014019150506040516020818303038152906040526040518082805190602001908083835b60208310612c695780518252601f199092019160209182019101612c4a565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015612ca8573d6000803e3d6000fd5b5050506040513d6020811015612cbd57600080fd5b505181612cc657fe5b061580612d2457507f2c8b528fbaf48aaf13162a5a0519a7ad5a612da8ff8783465c17e076660a59f160005260466020527f231bb0dc207f13dd4e565ebc32496c470e35391bd8d3b6649269ee2328e0311854610384429190910310155b612d5f5760405162461bcd60e51b8152600401808060200182810382526025815260200180614d316025913960400191505060405180910390fd5b50565b604080513360208083018290528351808403820181529284018452825192810192909220600091825260449092529190912054600114612de9576040805162461bcd60e51b815260206004820152601a60248201527f4d696e657220737461747573206973206e6f74207374616b6572000000000000604482015290519081900360640190fd5b603a54835114612e36576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b603c54602084015114612e86576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b603e54604084015114612ed6576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b604054606084015114612f26576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b604254608084015114612f76576040805162461bcd60e51b8152602060048201526013602482015272526571756573742049442069732077726f6e6760681b604482015290519081900360640190fd5b6000818152604660209081526040808320429055600080516020614c9983398151915254600080516020614b51833981519152548185526039845282852033865290935292205460ff1615612ffc5760405162461bcd60e51b8152600401808060200182810382526021815260200180614ceb6021913960400191505060405180910390fd5b60008281526039602090815260408083203384528252808320805460ff191660011790557fe97d205f7d20bf394e3813033d2203b4733acb28b351c8d2a771647ab0d41c3c548352604582528083208751848052600682019093529220836005811061306457fe5b015560208086015160016000908152600684019092526040909120836005811061308a57fe5b015560408086015160026000908152600684016020529190912083600581106130af57fe5b0155606085015160036000908152600683016020526040902083600581106130d357fe5b0155608085015160046000908152600683016020526040902083600581106130f757fe5b015560008080526005808301602052604090912033918490811061311757fe5b0180546001600160a01b0319166001600160a01b0392909216919091179055600160009081526005828101602052604090912033918490811061315657fe5b0180546001600160a01b0319166001600160a01b0392909216919091179055600260009081526005828101602052604090912033918490811061319557fe5b0180546001600160a01b0319166001600160a01b039290921691909117905560036000908152600582810160205260409091203391849081106131d457fe5b0180546001600160a01b0319166001600160a01b0392909216919091179055600460009081526005828101602052604090912033918490811061321357fe5b0180546001600160a01b0319166001600160a01b0392909216919091179055600182016004141561324657613246613aa5565b82336001600160a01b03167f9d2e5f03fc65aff196e0f3a8dd924b24099de487e8cffc888921d420ab196e3989898987604051808060200185600560200280838360005b838110156132a257818101518382015260200161328a565b5050505090500184600560200280838360005b838110156132cd5781810151838201526020016132b5565b50505050905001838152602001828103825286818151815260200191508051906020019080838360005b8381101561330f5781810151838201526020016132f7565b50505050905090810190601f16801561333c5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a381600101600514156133a2576133628787613bbf565b7fdfbec46864bc123768f0d134913175d9577a55bb71b9b2595fda21e21f36b08260009081526046602052600080516020614b51833981519152556133e2565b7fdfbec46864bc123768f0d134913175d9577a55bb71b9b2595fda21e21f36b0826000526046602052600080516020614b51833981519152805460010190555b50505050505050565b806133f55761354f565b6133ff82826117af565b61343a5760405162461bcd60e51b8152600401808060200182810382526027815260200180614b716027913960400191505060405180910390fd5b600061344583611019565b9050816001600160801b0380831682840390911611156134a0576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b600080516020614c198339815191526000526046602052600080516020614a7a83398151915254838103811015613512576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b61351e858385036139a9565b5050600080516020614c19833981519152600052506046602052600080516020614a7a833981519152805482900390555b5050565b6000828152604560209081526040808320600080516020614d568339815191528452600181019092529091205461358a90836143e2565b600080516020614d568339815191526000908152600183016020526040902055603a548314806135bb5750603c5483145b806135c75750603e5483145b806135d3575060405483145b806135df575060425483145b1561361557600080516020614c598339815191526000526046602052600080516020614b988339815191528054830190556121c7565b600080516020614bd883398151915260009081526001820160205260409020546137275760408051610660810191829052600091829161367491839060339082845b8154815260200190600101908083116136575750505050506143f8565b600080516020614d56833981519152600090815260018601602052604090205491935091508210806136a4575081155b1561372057600080516020614d5683398151915260009081526001840160205260408120549082603381106136d557fe5b015560008181526035602090815260408083208054845260458352818420600080516020614bd883398151915285526001908101845282852085905590899055860190915290208190555b50506121c7565b600080516020614bd883398151915260009081526001820160205260408120548391906033811061375457fe5b0180549091019055505050565b6001600160a01b0381166000908152604b602052604090205460ff16156137c2576040805162461bcd60e51b815260206004820152601060248201526f105b1c9958591e481b5a59dc985d195960821b604482015290519081900360640190fd5b7f56e0987db9eaec01ed9e0af003a0fd5c062371f9d23722eb4a3ebc74f16ea371600052604760209081527fc930326aab6c1874fc004d856083a6ed34e057e064970b7effb48e8e6e8ca12754604080516370a0823160e01b81526001600160a01b038086166004830152915161388c94869493909316926370a082319260248082019391829003018186803b15801561385b57600080fd5b505afa15801561386f573d6000803e3d6000fd5b505050506040513d602081101561388557600080fd5b5051614447565b6001600160a01b03166000908152604b60205260409020805460ff19166001179055565b6000828211156138bc57fe5b50900390565b80613928576001600160a01b0383166000908152604b602052604090205460ff1615613928576040805162461bcd60e51b815260206004820152601060248201526f185b1c9958591e481b5a59dc985d195960821b604482015290519081900360640190fd5b6139328383614447565b50506001600160a01b03166000908152604b60205260409020805460ff19166001179055565b60008082131561397557508181018281121561397057fe5b61084a565b508181018281131561084a57fe5b60008082131561399b57508082038281131561397057fe5b508082038281121561084a57fe5b6001600160a01b0382166000908152604960205260409020805415806139f6575080544390829060001981019081106139de57fe5b6000918252602090912001546001600160801b031614155b15613a675760408051808201909152436001600160801b0390811682528381166020808401918252845460018101865560008681529190912093519301805491516fffffffffffffffffffffffffffffffff19909216938316939093178216600160801b91909216021790556121c7565b805460009082906000198101908110613a7c57fe5b600091825260209091200180546001600160801b03808616600160801b02911617905550505050565b7f2c8b528fbaf48aaf13162a5a0519a7ad5a612da8ff8783465c17e076660a59f1600090815260466020527f231bb0dc207f13dd4e565ebc32496c470e35391bd8d3b6649269ee2328e0311854420390613b016104b08361462e565b6046602052600080516020614c79833981519152547fd4f87b8d0f3d3b7e665df74631f6100b2695daa0e30e40eeac02172e15a999e16000527f8156e704072c396780f8253d0562de28216b73a1503daa96e259b9cdd951d71c54610fa0929003810291909105915081613b7457600191505b613b818282016001614644565b7ff758978fc1647996a3d9992f611883adc442931dc49488312360acc90601759b6000526046602052600080516020614c7983398151915255505050565b7fe97d205f7d20bf394e3813033d2203b4733acb28b351c8d2a771647ab0d41c3c5460009081526045602090815260408220600080516020614c99833981519152547f2c8b528fbaf48aaf13162a5a0519a7ad5a612da8ff8783465c17e076660a59f190935260469091527f231bb0dc207f13dd4e565ebc32496c470e35391bd8d3b6649269ee2328e031188054429182905591929190613c5e6149be565b613c666149be565b60005b6005811015613fb75760015b6005811015613e2457600082815260068901602052604081208260058110613c9957fe5b0154905060008960050160008581526020019081526020016000208360058110613cbf57fe5b01546001600160a01b03169050825b600081118015613cfd5750600085815260068c0160205260409020600019820160058110613cf857fe5b015483105b15613db457600085815260068c0160205260409020600019820160058110613d2157fe5b0154600086815260068d01602052604090208260058110613d3e57fe5b015560008581526005808d016020526040909120906000198301908110613d6157fe5b015460008681526005808e0160205260409091206001600160a01b03909216919083908110613d8c57fe5b0180546001600160a01b0319166001600160a01b039290921691909117905560001901613cce565b83811015613e1957600085815260068c016020526040902083908260058110613dd957fe5b015560008581526005808d016020526040909120839183908110613df957fe5b0180546001600160a01b0319166001600160a01b03929092169190911790555b505050600101613c75565b506000604560008a8460058110613e3757fe5b602002015181526020019081526020016000209050876006016000838152602001908152602001600020600580602002604051908101604052809291908260058015613e98576020028201915b815481526020019060010190808311613e84575b5050505050935083600260058110613eac57fe5b6020908102919091015160008781526003840190925260409182902055840151838360058110613ed857fe5b6020908102919091019190915260008381526005808b0183526040808320898452858301909452909120613f0d9290916149dc565b5060008281526006808a0160209081526040808420898552928501909152909120613f399160056149dc565b5060008281526005890160205260408120613f5391614a17565b60008281526006890160205260408120613f6c91614a17565b805460018181018355600083815260208082209093018890558781526002840183526040808220439055600080516020614d5683398151915282529382019092529181205501613c69565b50847fbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc4588858460466000600080516020614c5983398151915260001b8152602001908152602001600020546040518085600560200280838360005b8381101561402a578181015183820152602001614012565b5050505090500184815260200183600560200280838360005b8381101561405b578181015183820152602001614043565b5050505090500182815260200194505050505060405180910390a2603380546001810182556000919091527f82a75bdeeae8604d839476ae9efd8b0e15aa447e21bfd7f41283bb54e22c9a82018390556140b36149be565b87516000908152604560209081526040808320878452600590810190925291829020825160a08101938490529290919082845b81546001600160a01b031681526001909101906020018083116140e657505050505090506141148186614653565b7ff3b93531fa65b3a18680d9ea49df06d96fbd883c4889dc7db866f8b131602dfb60005260466020527fe97d205f7d20bf394e3813033d2203b4733acb28b351c8d2a771647ab0d41c3c8054600101905561416d6149be565b614175614785565b905060005b600581101561427a5781816005811061418f57fe5b6020020151603a82600581106141a157fe5b60020201556000806045818585600581106141b857fe5b602002015181526020019081526020016000206001016000600080516020614bd883398151915260001b815260200190815260200160002054603381106141fb57fe5b01556045600083836005811061420d57fe5b6020908102919091015182528181019290925260409081016000908120600080516020614d5683398151915282526001908101845291812054600080516020614c598339815191529091526046909252600080516020614b9883398151915280549092019091550161417a565b50898760014303406040516020018080602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b838110156142ce5781810151838201526020016142b6565b50505050905090810190601f1680156142fb5780820380516001836020036101000a031916815260200191505b5060408051601f19818403018152908290528051602091820120600080516020614c998339815191528190556046909152600080516020614c7983398151915254600080516020614c598339815191526000908152600080516020614b9883398151915254929f508f98507f1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c140897508996509094509092508190859060a0908190849084905b838110156143b85781810151838201526020016143a0565b5050505091909101938452505060208201526040805191829003019150a250505050505050505050565b6000828201838110156143f157fe5b9392505050565b610640810151603260315b8015614441578284826033811061441657fe5b602002015110156144385783816033811061442d57fe5b602002015192508091505b60001901614403565b50915091565b806144835760405162461bcd60e51b8152600401808060200182810382526021815260200180614b306021913960400191505060405180910390fd5b6001600160a01b0382166144d6576040805162461bcd60e51b815260206004820152601560248201527452656365697665722069732030206164647265737360581b604482015290519081900360640190fd5b60006144e183611019565b9050816001600160801b03808316828401909116101561453c576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b600080516020614c198339815191526000526046602052600080516020614a7a833981519152548381018111156145ae576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b600080516020614c198339815191526000526046602052600080516020614a7a8339815191528054850190556145e6858484016139a9565b6040805185815290516001600160a01b038716916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050505050565b600081831061463d57816143f1565b5090919050565b600081831361463d57816143f1565b60466020527fc2c579d641b643400780d5c7ce967b420034b9f66962a5ee405cf70e4cbed6bb54600080516020614c598339815191526000908152600080516020614b9883398151915254428490039261012c9084020491600a9091049060028304906146ca9087905b6020020151838501614447565b6146d58660016146bd565b6146e08660026146bd565b6146eb8660036146bd565b6146f68660046146bd565b7f7a39905194de50bde334d18b76bbb36dddd11641d4d50b470cb837cf3bae5def60005260476020527fb5f7e7387e8e977cc9c4c9513388b0d7224264b9a0159cd8e8bdd84a9ed504c354614754906001600160a01b031682614447565b5050600080516020614c5983398151915260009081526046602052600080516020614b988339815191525550505050565b61478d6149be565b6147956149be565b61479d6149be565b604080516106608101918290526147d49160009060339082845b8154815260200190600101908083116147b757505050505061486a565b909250905060005b6005811015614864578281600581106147f157fe5b602002015115614835576035600083836005811061480b57fe5b602002015181526020019081526020016000205484826005811061482b57fe5b602002015261485c565b603a816004036005811061484557fe5b600202015484826005811061485657fe5b60200201525b6001016147dc565b50505090565b6148726149be565b61487a6149be565b60208301516000805b60058110156148fd5785816001016033811061489b57fe5b60200201518582600581106148ac57fe5b6020020152600181018482600581106148c157fe5b6020020152828582600581106148d357fe5b602002015110156148f5578481600581106148ea57fe5b602002015192508091505b600101614883565b5060065b60338110156149b6578286826033811061491757fe5b602002015111156149ae5785816033811061492e57fe5b602002015185836005811061493f57fe5b60200201528084836005811061495157fe5b602002015285816033811061496257fe5b6020020151925060005b60058110156149ac578386826005811061498257fe5b602002015110156149a45785816005811061499957fe5b602002015193508092505b60010161496c565b505b600101614901565b505050915091565b6040518060a001604052806005906020820280368337509192915050565b8260058101928215614a07579182015b82811115614a075782548255916001019190600101906149ec565b50614a13929150614a3a565b5090565b506000815560010160008155600101600081556001016000815560010160009055565b5b80821115614a135760008155600101614a3b56fe4d696e65722063616e206f6e6c792077696e2072657761726473206f6e636520706572203135206d696efffeead1ec15181fd57b4590d95e0c076bccb59e311315e8b38f23c710aa7c3e66756e6374696f6e206e6f742063616c6c61626c6520666f726b20666f726b2070726f706f73616c7345524332303a20617070726f766520746f20746865207a65726f206164647265737354696d6520666f72206120666f6c6c6f772075702064697370757465206861736e277420656c6170736564310199159a20c50879ffb440b45802138b5b162ec9426720e9dd3ee8bbcdb9d7547269656420746f206d696e74206e6f6e2d706f73697469766520616d6f756e747df1eb1754bc067736ff3d89af41d339bf906d31b0f5978e3c78f402d4ed249253686f756c6420686176652073756666696369656e742062616c616e636520746f20747261646538b16d06a20ab673b01c748aff938df6a38f81640035f4ce8bd9abb03aae5b722f9328a9c75282bec25bb04befad06926366736e0030c985108445fa728335e5f68d680ab3160f1aa5d9c3a1383c49e3e60bf3c0c031245cbb036f5ce99afaa1547269656420746f2073656e64206e6f6e2d706f73697469766520616d6f756e74e6148e7230ca038d456350e69a91b66968b222bfac9ebfbea6ff0a1fb73801601ce2382bc92689b00ba121fa5a411aa976168affdd8ac143a69035dd984b3b6a09659d32f99e50ac728058418d38174fe83a137c455ff1847e6fb8e15f78f77a5bccd7373734898281f858d7562320d2cdfc0b17bd72f779686937174d15002552cb9007c7c6068f8ef37039d4f232cbf5a28ff8d93a5983c4c0c27cd2f9bc0d44697370757465206d75737420626520737461727465642077697468696e2061207765656b206f66206261642076616c75654d696e657220616c7265616479207375626d6974746564207468652076616c756574686520646973707574652068617320616c7265616479206265656e206578656375746564496e636f7272656374206e6f6e636520666f722063757272656e74206368616c6c656e67651590276b7f31dd8e2a06f9a92867333eeb3eddbc91e73b9833e3e55d8e34f77da2646970667358221220166e386438f97c378d85d543a06fcb3e921cb40ef9478ff65ca54ae3af3ef1de64736f6c63430007040033",
}

// TellorABI is the input ABI used to generate the binding from.
// Deprecated: Use TellorMetaData.ABI instead.
var TellorABI = TellorMetaData.ABI

// Deprecated: Use TellorMetaData.Sigs instead.
// TellorFuncSigs maps the 4-byte function signature to its string representation.
var TellorFuncSigs = TellorMetaData.Sigs

// TellorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TellorMetaData.Bin instead.
var TellorBin = TellorMetaData.Bin

// DeployTellor deploys a new Ethereum contract, binding an instance of Tellor to it.
func DeployTellor(auth *bind.TransactOpts, backend bind.ContractBackend, _ext common.Address) (common.Address, *types.Transaction, *Tellor, error) {
	parsed, err := TellorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TellorBin), backend, _ext)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tellor{TellorCaller: TellorCaller{contract: contract}, TellorTransactor: TellorTransactor{contract: contract}, TellorFilterer: TellorFilterer{contract: contract}}, nil
}

// Tellor is an auto generated Go binding around an Ethereum contract.
type Tellor struct {
	TellorCaller     // Read-only binding to the contract
	TellorTransactor // Write-only binding to the contract
	TellorFilterer   // Log filterer for contract events
}

// TellorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorSession struct {
	Contract     *Tellor           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorCallerSession struct {
	Contract *TellorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TellorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorTransactorSession struct {
	Contract     *TellorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorRaw struct {
	Contract *Tellor // Generic contract binding to access the raw methods on
}

// TellorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorCallerRaw struct {
	Contract *TellorCaller // Generic read-only contract binding to access the raw methods on
}

// TellorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorTransactorRaw struct {
	Contract *TellorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellor creates a new instance of Tellor, bound to a specific deployed contract.
func NewTellor(address common.Address, backend bind.ContractBackend) (*Tellor, error) {
	contract, err := bindTellor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tellor{TellorCaller: TellorCaller{contract: contract}, TellorTransactor: TellorTransactor{contract: contract}, TellorFilterer: TellorFilterer{contract: contract}}, nil
}

// NewTellorCaller creates a new read-only instance of Tellor, bound to a specific deployed contract.
func NewTellorCaller(address common.Address, caller bind.ContractCaller) (*TellorCaller, error) {
	contract, err := bindTellor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorCaller{contract: contract}, nil
}

// NewTellorTransactor creates a new write-only instance of Tellor, bound to a specific deployed contract.
func NewTellorTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorTransactor, error) {
	contract, err := bindTellor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorTransactor{contract: contract}, nil
}

// NewTellorFilterer creates a new log filterer instance of Tellor, bound to a specific deployed contract.
func NewTellorFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorFilterer, error) {
	contract, err := bindTellor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorFilterer{contract: contract}, nil
}

// bindTellor binds a generic wrapper to an already deployed contract.
func bindTellor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tellor *TellorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tellor.Contract.TellorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tellor *TellorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.Contract.TellorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tellor *TellorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tellor.Contract.TellorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tellor *TellorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tellor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tellor *TellorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tellor *TellorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tellor.Contract.contract.Transact(opts, method, params...)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_Tellor *TellorCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "_allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_Tellor *TellorSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Tellor.Contract.Allowances(&_Tellor.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_Tellor *TellorCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Tellor.Contract.Allowances(&_Tellor.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_Tellor *TellorCaller) Addresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_Tellor *TellorSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _Tellor.Contract.Addresses(&_Tellor.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_Tellor *TellorCallerSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _Tellor.Contract.Addresses(&_Tellor.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_Tellor *TellorCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "allowance", _user, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_Tellor *TellorSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _Tellor.Contract.Allowance(&_Tellor.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_Tellor *TellorCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _Tellor.Contract.Allowance(&_Tellor.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_Tellor *TellorCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "allowedToTrade", _user, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_Tellor *TellorSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _Tellor.Contract.AllowedToTrade(&_Tellor.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_Tellor *TellorCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _Tellor.Contract.AllowedToTrade(&_Tellor.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_Tellor *TellorCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_Tellor *TellorSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _Tellor.Contract.BalanceOf(&_Tellor.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_Tellor *TellorCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _Tellor.Contract.BalanceOf(&_Tellor.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_Tellor *TellorCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "balanceOfAt", _user, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_Tellor *TellorSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _Tellor.Contract.BalanceOfAt(&_Tellor.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_Tellor *TellorCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _Tellor.Contract.BalanceOfAt(&_Tellor.CallOpts, _user, _blockNumber)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_Tellor *TellorCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "balances", arg0, arg1)

	outstruct := new(struct {
		FromBlock *big.Int
		Value     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_Tellor *TellorSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _Tellor.Contract.Balances(&_Tellor.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_Tellor *TellorCallerSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _Tellor.Contract.Balances(&_Tellor.CallOpts, arg0, arg1)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_Tellor *TellorCaller) BytesVars(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "bytesVars", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_Tellor *TellorSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _Tellor.Contract.BytesVars(&_Tellor.CallOpts, arg0)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_Tellor *TellorCallerSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _Tellor.Contract.BytesVars(&_Tellor.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_Tellor *TellorCaller) CurrentMiners(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "currentMiners", arg0)

	outstruct := new(struct {
		Value *big.Int
		Miner common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Miner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_Tellor *TellorSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _Tellor.Contract.CurrentMiners(&_Tellor.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_Tellor *TellorCallerSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _Tellor.Contract.CurrentMiners(&_Tellor.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_Tellor *TellorCaller) DisputeIdByDisputeHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "disputeIdByDisputeHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_Tellor *TellorSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _Tellor.Contract.DisputeIdByDisputeHash(&_Tellor.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_Tellor *TellorCallerSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _Tellor.Contract.DisputeIdByDisputeHash(&_Tellor.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_Tellor *TellorCaller) DisputesById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "disputesById", arg0)

	outstruct := new(struct {
		Hash                [32]byte
		Tally               *big.Int
		Executed            bool
		DisputeVotePassed   bool
		IsPropFork          bool
		ReportedMiner       common.Address
		ReportingParty      common.Address
		ProposedForkAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Tally = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Executed = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.DisputeVotePassed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.IsPropFork = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.ReportedMiner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.ReportingParty = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.ProposedForkAddress = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_Tellor *TellorSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _Tellor.Contract.DisputesById(&_Tellor.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_Tellor *TellorCallerSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _Tellor.Contract.DisputesById(&_Tellor.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_Tellor *TellorCaller) Migrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "migrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_Tellor *TellorSession) Migrated(arg0 common.Address) (bool, error) {
	return _Tellor.Contract.Migrated(&_Tellor.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_Tellor *TellorCallerSession) Migrated(arg0 common.Address) (bool, error) {
	return _Tellor.Contract.Migrated(&_Tellor.CallOpts, arg0)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_Tellor *TellorCaller) MinersByChallenge(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "minersByChallenge", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_Tellor *TellorSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Tellor.Contract.MinersByChallenge(&_Tellor.CallOpts, arg0, arg1)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_Tellor *TellorCallerSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Tellor.Contract.MinersByChallenge(&_Tellor.CallOpts, arg0, arg1)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_Tellor *TellorCaller) NewValueTimestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "newValueTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_Tellor *TellorSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _Tellor.Contract.NewValueTimestamps(&_Tellor.CallOpts, arg0)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_Tellor *TellorCallerSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _Tellor.Contract.NewValueTimestamps(&_Tellor.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_Tellor *TellorCaller) RequestIdByQueryHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "requestIdByQueryHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_Tellor *TellorSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _Tellor.Contract.RequestIdByQueryHash(&_Tellor.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_Tellor *TellorCallerSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _Tellor.Contract.RequestIdByQueryHash(&_Tellor.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_Tellor *TellorCaller) Uints(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "uints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_Tellor *TellorSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _Tellor.Contract.Uints(&_Tellor.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_Tellor *TellorCallerSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _Tellor.Contract.Uints(&_Tellor.CallOpts, arg0)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Tellor *TellorTransactor) AddTip(opts *bind.TransactOpts, _requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "addTip", _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Tellor *TellorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.AddTip(&_Tellor.TransactOpts, _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Tellor *TellorTransactorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.AddTip(&_Tellor.TransactOpts, _requestId, _tip)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Tellor *TellorTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Tellor *TellorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.Approve(&_Tellor.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Tellor *TellorTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.Approve(&_Tellor.TransactOpts, _spender, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Tellor *TellorTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Tellor *TellorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.BeginDispute(&_Tellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Tellor *TellorTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.BeginDispute(&_Tellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Tellor *TellorTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Tellor *TellorSession) Migrate() (*types.Transaction, error) {
	return _Tellor.Contract.Migrate(&_Tellor.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_Tellor *TellorTransactorSession) Migrate() (*types.Transaction, error) {
	return _Tellor.Contract.Migrate(&_Tellor.TransactOpts)
}

// MigrateFor is a paid mutator transaction binding the contract method 0xa9fa7d34.
//
// Solidity: function migrateFor(address _destination, uint256 _amount, bool _bypass) returns()
func (_Tellor *TellorTransactor) MigrateFor(opts *bind.TransactOpts, _destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "migrateFor", _destination, _amount, _bypass)
}

// MigrateFor is a paid mutator transaction binding the contract method 0xa9fa7d34.
//
// Solidity: function migrateFor(address _destination, uint256 _amount, bool _bypass) returns()
func (_Tellor *TellorSession) MigrateFor(_destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _Tellor.Contract.MigrateFor(&_Tellor.TransactOpts, _destination, _amount, _bypass)
}

// MigrateFor is a paid mutator transaction binding the contract method 0xa9fa7d34.
//
// Solidity: function migrateFor(address _destination, uint256 _amount, bool _bypass) returns()
func (_Tellor *TellorTransactorSession) MigrateFor(_destination common.Address, _amount *big.Int, _bypass bool) (*types.Transaction, error) {
	return _Tellor.Contract.MigrateFor(&_Tellor.TransactOpts, _destination, _amount, _bypass)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_Tellor *TellorTransactor) ProposeFork(opts *bind.TransactOpts, _propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "proposeFork", _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_Tellor *TellorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _Tellor.Contract.ProposeFork(&_Tellor.TransactOpts, _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_Tellor *TellorTransactorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _Tellor.Contract.ProposeFork(&_Tellor.TransactOpts, _propNewTellorAddress)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestIds, uint256[5] _values) returns()
func (_Tellor *TellorTransactor) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestIds [5]*big.Int, _values [5]*big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "submitMiningSolution", _nonce, _requestIds, _values)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestIds, uint256[5] _values) returns()
func (_Tellor *TellorSession) SubmitMiningSolution(_nonce string, _requestIds [5]*big.Int, _values [5]*big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.SubmitMiningSolution(&_Tellor.TransactOpts, _nonce, _requestIds, _values)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestIds, uint256[5] _values) returns()
func (_Tellor *TellorTransactorSession) SubmitMiningSolution(_nonce string, _requestIds [5]*big.Int, _values [5]*big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.SubmitMiningSolution(&_Tellor.TransactOpts, _nonce, _requestIds, _values)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_Tellor *TellorTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_Tellor *TellorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.Transfer(&_Tellor.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_Tellor *TellorTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.Transfer(&_Tellor.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_Tellor *TellorTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_Tellor *TellorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.TransferFrom(&_Tellor.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_Tellor *TellorTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.TransferFrom(&_Tellor.TransactOpts, _from, _to, _amount)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_Tellor *TellorTransactor) UnlockDisputeFee(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "unlockDisputeFee", _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_Tellor *TellorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.UnlockDisputeFee(&_Tellor.TransactOpts, _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_Tellor *TellorTransactorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.UnlockDisputeFee(&_Tellor.TransactOpts, _disputeId)
}

// Verify is a paid mutator transaction binding the contract method 0xfc735e99.
//
// Solidity: function verify() returns(uint256)
func (_Tellor *TellorTransactor) Verify(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "verify")
}

// Verify is a paid mutator transaction binding the contract method 0xfc735e99.
//
// Solidity: function verify() returns(uint256)
func (_Tellor *TellorSession) Verify() (*types.Transaction, error) {
	return _Tellor.Contract.Verify(&_Tellor.TransactOpts)
}

// Verify is a paid mutator transaction binding the contract method 0xfc735e99.
//
// Solidity: function verify() returns(uint256)
func (_Tellor *TellorTransactorSession) Verify() (*types.Transaction, error) {
	return _Tellor.Contract.Verify(&_Tellor.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Tellor *TellorTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Tellor *TellorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Tellor.Contract.Vote(&_Tellor.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Tellor *TellorTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Tellor.Contract.Vote(&_Tellor.TransactOpts, _disputeId, _supportsDispute)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Tellor *TellorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Tellor.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Tellor *TellorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Tellor.Contract.Fallback(&_Tellor.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Tellor *TellorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Tellor.Contract.Fallback(&_Tellor.TransactOpts, calldata)
}

// TellorApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Tellor contract.
type TellorApprovalIterator struct {
	Event *TellorApproval // Event containing the contract specifics and raw log

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
func (it *TellorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorApproval)
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
		it.Event = new(TellorApproval)
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
func (it *TellorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorApproval represents a Approval event raised by the Tellor contract.
type TellorApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Tellor *TellorFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*TellorApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &TellorApprovalIterator{contract: _Tellor.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Tellor *TellorFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TellorApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorApproval)
				if err := _Tellor.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Tellor *TellorFilterer) ParseApproval(log types.Log) (*TellorApproval, error) {
	event := new(TellorApproval)
	if err := _Tellor.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorNewChallengeIterator is returned from FilterNewChallenge and is used to iterate over the raw logs and unpacked data for NewChallenge events raised by the Tellor contract.
type TellorNewChallengeIterator struct {
	Event *TellorNewChallenge // Event containing the contract specifics and raw log

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
func (it *TellorNewChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorNewChallenge)
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
		it.Event = new(TellorNewChallenge)
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
func (it *TellorNewChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorNewChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorNewChallenge represents a NewChallenge event raised by the Tellor contract.
type TellorNewChallenge struct {
	CurrentChallenge [32]byte
	CurrentRequestId [5]*big.Int
	Difficulty       *big.Int
	TotalTips        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewChallenge is a free log retrieval operation binding the contract event 0x1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c1408.
//
// Solidity: event NewChallenge(bytes32 indexed _currentChallenge, uint256[5] _currentRequestId, uint256 _difficulty, uint256 _totalTips)
func (_Tellor *TellorFilterer) FilterNewChallenge(opts *bind.FilterOpts, _currentChallenge [][32]byte) (*TellorNewChallengeIterator, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "NewChallenge", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &TellorNewChallengeIterator{contract: _Tellor.contract, event: "NewChallenge", logs: logs, sub: sub}, nil
}

// WatchNewChallenge is a free log subscription operation binding the contract event 0x1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c1408.
//
// Solidity: event NewChallenge(bytes32 indexed _currentChallenge, uint256[5] _currentRequestId, uint256 _difficulty, uint256 _totalTips)
func (_Tellor *TellorFilterer) WatchNewChallenge(opts *bind.WatchOpts, sink chan<- *TellorNewChallenge, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "NewChallenge", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorNewChallenge)
				if err := _Tellor.contract.UnpackLog(event, "NewChallenge", log); err != nil {
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

// ParseNewChallenge is a log parse operation binding the contract event 0x1d85ce10456e29b67de37887496d3f1fcf1b64c79c4d07484038703a9f5c1408.
//
// Solidity: event NewChallenge(bytes32 indexed _currentChallenge, uint256[5] _currentRequestId, uint256 _difficulty, uint256 _totalTips)
func (_Tellor *TellorFilterer) ParseNewChallenge(log types.Log) (*TellorNewChallenge, error) {
	event := new(TellorNewChallenge)
	if err := _Tellor.contract.UnpackLog(event, "NewChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorNewDisputeIterator is returned from FilterNewDispute and is used to iterate over the raw logs and unpacked data for NewDispute events raised by the Tellor contract.
type TellorNewDisputeIterator struct {
	Event *TellorNewDispute // Event containing the contract specifics and raw log

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
func (it *TellorNewDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorNewDispute)
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
		it.Event = new(TellorNewDispute)
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
func (it *TellorNewDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorNewDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorNewDispute represents a NewDispute event raised by the Tellor contract.
type TellorNewDispute struct {
	DisputeId *big.Int
	RequestId *big.Int
	Timestamp *big.Int
	Miner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewDispute is a free log retrieval operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_Tellor *TellorFilterer) FilterNewDispute(opts *bind.FilterOpts, _disputeId []*big.Int, _requestId []*big.Int) (*TellorNewDisputeIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &TellorNewDisputeIterator{contract: _Tellor.contract, event: "NewDispute", logs: logs, sub: sub}, nil
}

// WatchNewDispute is a free log subscription operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_Tellor *TellorFilterer) WatchNewDispute(opts *bind.WatchOpts, sink chan<- *TellorNewDispute, _disputeId []*big.Int, _requestId []*big.Int) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorNewDispute)
				if err := _Tellor.contract.UnpackLog(event, "NewDispute", log); err != nil {
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

// ParseNewDispute is a log parse operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_Tellor *TellorFilterer) ParseNewDispute(log types.Log) (*TellorNewDispute, error) {
	event := new(TellorNewDispute)
	if err := _Tellor.contract.UnpackLog(event, "NewDispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorNewValueIterator is returned from FilterNewValue and is used to iterate over the raw logs and unpacked data for NewValue events raised by the Tellor contract.
type TellorNewValueIterator struct {
	Event *TellorNewValue // Event containing the contract specifics and raw log

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
func (it *TellorNewValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorNewValue)
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
		it.Event = new(TellorNewValue)
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
func (it *TellorNewValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorNewValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorNewValue represents a NewValue event raised by the Tellor contract.
type TellorNewValue struct {
	RequestId        [5]*big.Int
	Time             *big.Int
	Value            [5]*big.Int
	TotalTips        *big.Int
	CurrentChallenge [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewValue is a free log retrieval operation binding the contract event 0xbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc45.
//
// Solidity: event NewValue(uint256[5] _requestId, uint256 _time, uint256[5] _value, uint256 _totalTips, bytes32 indexed _currentChallenge)
func (_Tellor *TellorFilterer) FilterNewValue(opts *bind.FilterOpts, _currentChallenge [][32]byte) (*TellorNewValueIterator, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "NewValue", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &TellorNewValueIterator{contract: _Tellor.contract, event: "NewValue", logs: logs, sub: sub}, nil
}

// WatchNewValue is a free log subscription operation binding the contract event 0xbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc45.
//
// Solidity: event NewValue(uint256[5] _requestId, uint256 _time, uint256[5] _value, uint256 _totalTips, bytes32 indexed _currentChallenge)
func (_Tellor *TellorFilterer) WatchNewValue(opts *bind.WatchOpts, sink chan<- *TellorNewValue, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "NewValue", _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorNewValue)
				if err := _Tellor.contract.UnpackLog(event, "NewValue", log); err != nil {
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

// ParseNewValue is a log parse operation binding the contract event 0xbeb3b9f46c8d7bb00c873fca10d307538df350866d25f891ffb395147ddbdc45.
//
// Solidity: event NewValue(uint256[5] _requestId, uint256 _time, uint256[5] _value, uint256 _totalTips, bytes32 indexed _currentChallenge)
func (_Tellor *TellorFilterer) ParseNewValue(log types.Log) (*TellorNewValue, error) {
	event := new(TellorNewValue)
	if err := _Tellor.contract.UnpackLog(event, "NewValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorNonceSubmittedIterator is returned from FilterNonceSubmitted and is used to iterate over the raw logs and unpacked data for NonceSubmitted events raised by the Tellor contract.
type TellorNonceSubmittedIterator struct {
	Event *TellorNonceSubmitted // Event containing the contract specifics and raw log

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
func (it *TellorNonceSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorNonceSubmitted)
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
		it.Event = new(TellorNonceSubmitted)
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
func (it *TellorNonceSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorNonceSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorNonceSubmitted represents a NonceSubmitted event raised by the Tellor contract.
type TellorNonceSubmitted struct {
	Miner            common.Address
	Nonce            string
	RequestId        [5]*big.Int
	Value            [5]*big.Int
	CurrentChallenge [32]byte
	Slot             *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNonceSubmitted is a free log retrieval operation binding the contract event 0x9d2e5f03fc65aff196e0f3a8dd924b24099de487e8cffc888921d420ab196e39.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256[5] _requestId, uint256[5] _value, bytes32 indexed _currentChallenge, uint256 _slot)
func (_Tellor *TellorFilterer) FilterNonceSubmitted(opts *bind.FilterOpts, _miner []common.Address, _currentChallenge [][32]byte) (*TellorNonceSubmittedIterator, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "NonceSubmitted", _minerRule, _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return &TellorNonceSubmittedIterator{contract: _Tellor.contract, event: "NonceSubmitted", logs: logs, sub: sub}, nil
}

// WatchNonceSubmitted is a free log subscription operation binding the contract event 0x9d2e5f03fc65aff196e0f3a8dd924b24099de487e8cffc888921d420ab196e39.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256[5] _requestId, uint256[5] _value, bytes32 indexed _currentChallenge, uint256 _slot)
func (_Tellor *TellorFilterer) WatchNonceSubmitted(opts *bind.WatchOpts, sink chan<- *TellorNonceSubmitted, _miner []common.Address, _currentChallenge [][32]byte) (event.Subscription, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _currentChallengeRule []interface{}
	for _, _currentChallengeItem := range _currentChallenge {
		_currentChallengeRule = append(_currentChallengeRule, _currentChallengeItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "NonceSubmitted", _minerRule, _currentChallengeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorNonceSubmitted)
				if err := _Tellor.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
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

// ParseNonceSubmitted is a log parse operation binding the contract event 0x9d2e5f03fc65aff196e0f3a8dd924b24099de487e8cffc888921d420ab196e39.
//
// Solidity: event NonceSubmitted(address indexed _miner, string _nonce, uint256[5] _requestId, uint256[5] _value, bytes32 indexed _currentChallenge, uint256 _slot)
func (_Tellor *TellorFilterer) ParseNonceSubmitted(log types.Log) (*TellorNonceSubmitted, error) {
	event := new(TellorNonceSubmitted)
	if err := _Tellor.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorTipAddedIterator is returned from FilterTipAdded and is used to iterate over the raw logs and unpacked data for TipAdded events raised by the Tellor contract.
type TellorTipAddedIterator struct {
	Event *TellorTipAdded // Event containing the contract specifics and raw log

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
func (it *TellorTipAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorTipAdded)
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
		it.Event = new(TellorTipAdded)
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
func (it *TellorTipAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorTipAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorTipAdded represents a TipAdded event raised by the Tellor contract.
type TellorTipAdded struct {
	Sender    common.Address
	RequestId *big.Int
	Tip       *big.Int
	TotalTips *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTipAdded is a free log retrieval operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_Tellor *TellorFilterer) FilterTipAdded(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*TellorTipAddedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &TellorTipAddedIterator{contract: _Tellor.contract, event: "TipAdded", logs: logs, sub: sub}, nil
}

// WatchTipAdded is a free log subscription operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_Tellor *TellorFilterer) WatchTipAdded(opts *bind.WatchOpts, sink chan<- *TellorTipAdded, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorTipAdded)
				if err := _Tellor.contract.UnpackLog(event, "TipAdded", log); err != nil {
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

// ParseTipAdded is a log parse operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip, uint256 _totalTips)
func (_Tellor *TellorFilterer) ParseTipAdded(log types.Log) (*TellorTipAdded, error) {
	event := new(TellorTipAdded)
	if err := _Tellor.contract.UnpackLog(event, "TipAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the Tellor contract.
type TellorTransferredIterator struct {
	Event *TellorTransferred // Event containing the contract specifics and raw log

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
func (it *TellorTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorTransferred)
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
		it.Event = new(TellorTransferred)
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
func (it *TellorTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorTransferred represents a Transferred event raised by the Tellor contract.
type TellorTransferred struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Tellor *TellorFilterer) FilterTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*TellorTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TellorTransferredIterator{contract: _Tellor.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Tellor *TellorFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *TellorTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorTransferred)
				if err := _Tellor.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransferred is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Tellor *TellorFilterer) ParseTransferred(log types.Log) (*TellorTransferred, error) {
	event := new(TellorTransferred)
	if err := _Tellor.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Tellor contract.
type TellorVotedIterator struct {
	Event *TellorVoted // Event containing the contract specifics and raw log

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
func (it *TellorVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorVoted)
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
		it.Event = new(TellorVoted)
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
func (it *TellorVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorVoted represents a Voted event raised by the Tellor contract.
type TellorVoted struct {
	DisputeID  *big.Int
	Position   bool
	Voter      common.Address
	VoteWeight *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_Tellor *TellorFilterer) FilterVoted(opts *bind.FilterOpts, _disputeID []*big.Int, _voter []common.Address, _voteWeight []*big.Int) (*TellorVotedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}
	var _voteWeightRule []interface{}
	for _, _voteWeightItem := range _voteWeight {
		_voteWeightRule = append(_voteWeightRule, _voteWeightItem)
	}

	logs, sub, err := _Tellor.contract.FilterLogs(opts, "Voted", _disputeIDRule, _voterRule, _voteWeightRule)
	if err != nil {
		return nil, err
	}
	return &TellorVotedIterator{contract: _Tellor.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_Tellor *TellorFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *TellorVoted, _disputeID []*big.Int, _voter []common.Address, _voteWeight []*big.Int) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}
	var _voteWeightRule []interface{}
	for _, _voteWeightItem := range _voteWeight {
		_voteWeightRule = append(_voteWeightRule, _voteWeightItem)
	}

	logs, sub, err := _Tellor.contract.WatchLogs(opts, "Voted", _disputeIDRule, _voterRule, _voteWeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorVoted)
				if err := _Tellor.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_Tellor *TellorFilterer) ParseVoted(log types.Log) (*TellorVoted, error) {
	event := new(TellorVoted)
	if err := _Tellor.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorGettersMetaData contains all meta data concerning the TellorGetters contract.
var TellorGettersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"fromBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentMiners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"disputesById\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"tally\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disputeVotePassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPropFork\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"reportedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reportingParty\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedForkAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[9]\",\"name\":\"\",\"type\":\"uint256[9]\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_diff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNewVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"idsOnDeck\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"tipsOnDeck\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"internalType\":\"uint256[51]\",\"name\":\"\",\"type\":\"uint256[51]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopRequestIDs\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minersByChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newValueTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"024c2ddd": "_allowances(address,address)",
		"699f200f": "addresses(bytes32)",
		"cbf1304d": "balances(address,uint256)",
		"62dd1d2a": "bytesVars(bytes32)",
		"1fd22364": "currentMiners(uint256)",
		"313ce567": "decimals()",
		"63bb82ad": "didMine(bytes32,address)",
		"a7c438bc": "didVote(uint256,address)",
		"d01f4d9e": "disputeIdByDisputeHash(bytes32)",
		"db085beb": "disputesById(uint256)",
		"133bee5e": "getAddressVars(bytes32)",
		"af0b1327": "getAllDisputeVars(uint256)",
		"da379941": "getDisputeIdByDisputeHash(bytes32)",
		"7f6fd5d9": "getDisputeUintVars(uint256,bytes32)",
		"3180f8df": "getLastNewValueById(uint256)",
		"c775b542": "getMinedBlockNum(uint256,uint256)",
		"69026d63": "getMinersByRequestIdAndTimestamp(uint256,uint256)",
		"4049f198": "getNewCurrentVariables()",
		"46eee1c4": "getNewValueCountbyRequestId(uint256)",
		"9a7077ab": "getNewVariablesOnDeck()",
		"6173c0b8": "getRequestIdByRequestQIndex(uint256)",
		"b5413029": "getRequestQ()",
		"e0ae93c1": "getRequestUintVars(uint256,bytes32)",
		"e1eee6d6": "getRequestVars(uint256)",
		"733bdef0": "getStakerInfo(address)",
		"11c98512": "getSubmissionsByTimestamp(uint256,uint256)",
		"77fbb663": "getTimestampbyRequestIDandIndex(uint256,uint256)",
		"fe1cd15d": "getTopRequestIDs()",
		"612c8f7f": "getUintVar(bytes32)",
		"3df0777b": "isInDispute(uint256,uint256)",
		"4ba0a5ee": "migrated(address)",
		"48b18e54": "minersByChallenge(bytes32,address)",
		"06fdde03": "name()",
		"438c0aa3": "newValueTimestamps(uint256)",
		"5700242c": "requestIdByQueryHash(bytes32)",
		"93fa4915": "retrieveData(uint256,uint256)",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"b59e14d4": "uints(bytes32)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061163a806100206000396000f3fe608060405234801561001057600080fd5b506004361061023d5760003560e01c806369026d631161013b578063b5413029116100b8578063da3799411161007c578063da379941146108f4578063db085beb14610911578063e0ae93c11461097e578063e1eee6d6146109a1578063fe1cd15d146109be5761023d565b8063b54130291461081e578063b59e14d41461083c578063c775b54214610859578063cbf1304d1461087c578063d01f4d9e146108d75761023d565b806393fa4915116100ff57806393fa4915146106aa57806395d89b41146106cd5780639a7077ab146106d5578063a7c438bc14610741578063af0b13271461076d5761023d565b806369026d63146105e5578063699f200f14610608578063733bdef01461062557806377fbb663146106645780637f6fd5d9146106875761023d565b80634049f198116101c95780635700242c1161018d5780635700242c14610545578063612c8f7f146105625780636173c0b81461057f57806362dd1d2a1461059c57806363bb82ad146105b95761023d565b80634049f19814610464578063438c0aa3146104b957806346eee1c4146104d657806348b18e54146104f35780634ba0a5ee1461051f5761023d565b806318160ddd1161021057806318160ddd146103935780631fd223641461039b578063313ce567146103d95780633180f8df146103f75780633df0777b1461042d5761023d565b8063024c2ddd1461024257806306fdde031461028257806311c98512146102ff578063133bee5e1461035a575b600080fd5b6102706004803603604081101561025857600080fd5b506001600160a01b03813581169160200135166109c6565b60408051918252519081900360200190f35b61028a6109e3565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102c45781810151838201526020016102ac565b50505050905090810190601f1680156102f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103226004803603604081101561031557600080fd5b5080359060200135610a0c565b604051808260a080838360005b8381101561034757818101518382015260200161032f565b5050505090500191505060405180910390f35b6103776004803603602081101561037057600080fd5b5035610a63565b604080516001600160a01b039092168252519081900360200190f35b610270610a7e565b6103b8600480360360208110156103b157600080fd5b5035610acc565b604080519283526001600160a01b0390911660208301528051918290030190f35b6103e1610af7565b6040805160ff9092168252519081900360200190f35b6104146004803603602081101561040d57600080fd5b5035610afc565b6040805192835290151560208301528051918290030190f35b6104506004803603604081101561044357600080fd5b5080359060200135610b56565b604080519115158252519081900360200190f35b61046c610b7a565b604051848152602081018460a080838360005b8381101561049757818101518382015260200161047f565b5050505090500183815260200182815260200194505050505060405180910390f35b610270600480360360208110156104cf57600080fd5b5035610c5a565b610270600480360360208110156104ec57600080fd5b5035610c7b565b6104506004803603604081101561050957600080fd5b50803590602001356001600160a01b0316610c8d565b6104506004803603602081101561053557600080fd5b50356001600160a01b0316610cad565b6102706004803603602081101561055b57600080fd5b5035610cc2565b6102706004803603602081101561057857600080fd5b5035610cd4565b6102706004803603602081101561059557600080fd5b5035610ce6565b610270600480360360208110156105b257600080fd5b5035610d51565b610450600480360360408110156105cf57600080fd5b50803590602001356001600160a01b0316610d63565b610322600480360360408110156105fb57600080fd5b5080359060200135610d8e565b6103776004803603602081101561061e57600080fd5b5035610df1565b61064b6004803603602081101561063b57600080fd5b50356001600160a01b0316610e0c565b6040805192835260208301919091528051918290030190f35b6102706004803603604081101561067a57600080fd5b5080359060200135610e2f565b6102706004803603604081101561069d57600080fd5b5080359060200135610e5c565b610270600480360360408110156106c057600080fd5b5080359060200135610e7d565b61028a610e9e565b6106dd610ebb565b604051808360a080838360005b838110156107025781810151838201526020016106ea565b5050505090500182600560200280838360005b8381101561072d578181015183820152602001610715565b505050509050019250505060405180910390f35b6104506004803603604081101561075757600080fd5b50803590602001356001600160a01b0316610f5a565b61078a6004803603602081101561078357600080fd5b5035610f89565b604051808a8152602001891515815260200188151581526020018715158152602001866001600160a01b03168152602001856001600160a01b03168152602001846001600160a01b0316815260200183600960200280838360005b838110156107fd5781810151838201526020016107e5565b50505050905001828152602001995050505050505050505060405180910390f35b6108266111b4565b604051815181528082610660808383602061032f565b6102706004803603602081101561085257600080fd5b50356111f0565b6102706004803603604081101561086f57600080fd5b5080359060200135611202565b6108a86004803603604081101561089257600080fd5b506001600160a01b038135169060200135611223565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b610270600480360360208110156108ed57600080fd5b5035611266565b6102706004803603602081101561090a57600080fd5b5035611278565b61092e6004803603602081101561092757600080fd5b503561128a565b60408051988952602089019790975294151587870152921515606087015290151560808601526001600160a01b0390811660a086015290811660c08501521660e083015251908190036101000190f35b6102706004803603604081101561099457600080fd5b50803590602001356112e6565b61064b600480360360208110156109b757600080fd5b5035611307565b61032261136f565b604a60209081526000928352604080842090915290825290205481565b60408051808201909152600f81526e54656c6c6f7220547269627574657360881b602082015290565b610a146115a8565b600083815260456020908152604080832085845260060190915290819020815160a08101928390529160059082845b815481526020019060010190808311610a43575050505050905092915050565b6000908152604760205260409020546001600160a01b031690565b7fe6148e7230ca038d456350e69a91b66968b222bfac9ebfbea6ff0a1fb738016060005260466020527ffffeead1ec15181fd57b4590d95e0c076bccb59e311315e8b38f23c710aa7c3e5490565b603a8160058110610adc57600080fd5b6002020180546001909101549091506001600160a01b031682565b601290565b6000818152604560205260408120805482919015610b48578054610b3c90859083906000198101908110610b2c57fe5b9060005260206000200154610e7d565b60019250925050610b51565b60008092509250505b915091565b60009182526045602090815260408084209284526004909201905290205460ff1690565b6000610b846115a8565b60008060005b6005811015610bbe57603a8160058110610ba057fe5b6002020154848260058110610bb157fe5b6020020152600101610b8a565b50507f52cb9007c7c6068f8ef37039d4f232cbf5a28ff8d93a5983c4c0c27cd2f9bc0d5460466020527f5bccd7373734898281f858d7562320d2cdfc0b17bd72f779686937174d150025547f09659d32f99e50ac728058418d38174fe83a137c455ff1847e6fb8e15f78f77a6000527f38b16d06a20ab673b01c748aff938df6a38f81640035f4ce8bd9abb03aae5b7254919450915090919293565b60338181548110610c6a57600080fd5b600091825260209091200154905081565b60009081526045602052604090205490565b603960209081526000928352604080842090915290825290205460ff1681565b604b6020526000908152604090205460ff1681565b60376020526000908152604090205481565b60009081526046602052604090205490565b60006032821115610d3e576040805162461bcd60e51b815260206004820152601a60248201527f526571756573745120696e6465782069732061626f7665203530000000000000604482015290519081900360640190fd5b5060009081526035602052604090205490565b60486020526000908152604090205481565b60009182526039602090815260408084206001600160a01b0393909316845291905290205460ff1690565b610d966115a8565b6000838152604560209081526040808320858452600590810190925291829020825160a08101938490529290919082845b81546001600160a01b03168152600190910190602001808311610dc7575050505050905092915050565b6047602052600090815260409020546001600160a01b031681565b6001600160a01b0316600090815260446020526040902080546001909101549091565b6000828152604560205260408120805483908110610e4957fe5b9060005260206000200154905092915050565b60009182526036602090815260408084209284526005909201905290205490565b60009182526045602090815260408084209284526003909201905290205490565b6040805180820190915260038152622a292160e91b602082015290565b610ec36115a8565b610ecb6115a8565b610ed361136f565b915060005b6005811015610f555760456000848360058110610ef157fe5b6020020151815260200190815260200160002060010160007f1590276b7f31dd8e2a06f9a92867333eeb3eddbc91e73b9833e3e55d8e34f77d60001b815260200190815260200160002054828260058110610f4857fe5b6020020152600101610ed8565b509091565b60008281526036602090815260408083206001600160a01b038516845260060190915290205460ff1692915050565b6000806000806000806000610f9c6115c6565b5050506000958652505060366020908152604080862080546002820154600383015460048401548551610120810187527f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f8c5260058601808952878d205482527f2f9328a9c75282bec25bb04befad06926366736e0030c985108445fa728335e58d52808952878d2054828a01527f9147231ab14efb72c38117f68521ddef8de64f092c18c69dbfb602ffc4de7f478d52808952878d2054828901527f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d288d52808952878d205460608301527f1da378694063870452ce03b189f48e04c1aa026348e74e6c86e10738514ad2c48d52808952878d205460808301527f4b4cefd5ced7569ef0d091282b4bca9c52a034c56471a6061afd1bf307a2de7c8d52808952878d205460a08301527f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d8d52808952878d205460c08301527f30e85ae205656781c1a951cba9f9f53f884833c049d377a2a7046eb5e6d14b268d52808952878d205460e08301527f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc4098d52909752949099205461010080870191909152600190930154919960ff8083169a948304811699506201000083041697506001600160a01b036301000000909204821696509281169493169291565b6111bc6115e5565b604080516106608101918290529060009060339082845b8154815260200190600101908083116111d3575050505050905090565b60466020526000908152604090205481565b60009182526045602090815260408084209284526002909201905290205490565b6049602052816000526040600020818154811061123f57600080fd5b6000918252602090912001546001600160801b038082169350600160801b90910416905082565b60386020526000908152604090205481565b60009081526038602052604090205490565b603660205260009081526040902080546001820154600283015460038401546004909401549293919260ff808316936101008404821693620100008104909216926001600160a01b036301000000909304831692918216911688565b60009182526045602090815260408084209284526001909201905290205490565b60009081526045602090815260408083207ff68d680ab3160f1aa5d9c3a1383c49e3e60bf3c0c031245cbb036f5ce99afaa18452600101909152808220547f1590276b7f31dd8e2a06f9a92867333eeb3eddbc91e73b9833e3e55d8e34f77d83529120549091565b6113776115a8565b61137f6115a8565b6113876115a8565b604080516106608101918290526113be9160009060339082845b8154815260200190600101908083116113a1575050505050611454565b909250905060005b600581101561144e578281600581106113db57fe5b60200201511561141f57603560008383600581106113f557fe5b602002015181526020019081526020016000205484826005811061141557fe5b6020020152611446565b603a816004036005811061142f57fe5b600202015484826005811061144057fe5b60200201525b6001016113c6565b50505090565b61145c6115a8565b6114646115a8565b60208301516000805b60058110156114e75785816001016033811061148557fe5b602002015185826005811061149657fe5b6020020152600181018482600581106114ab57fe5b6020020152828582600581106114bd57fe5b602002015110156114df578481600581106114d457fe5b602002015192508091505b60010161146d565b5060065b60338110156115a0578286826033811061150157fe5b602002015111156115985785816033811061151857fe5b602002015185836005811061152957fe5b60200201528084836005811061153b57fe5b602002015285816033811061154c57fe5b6020020151925060005b6005811015611596578386826005811061156c57fe5b6020020151101561158e5785816005811061158357fe5b602002015193508092505b600101611556565b505b6001016114eb565b505050915091565b6040518060a001604052806005906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b604051806106600160405280603390602082028036833750919291505056fea2646970667358221220b857a56e66c83d39061e4cf6e37ff3d6abeb8fddf5cf8be62f907d97e46b603164736f6c63430007040033",
}

// TellorGettersABI is the input ABI used to generate the binding from.
// Deprecated: Use TellorGettersMetaData.ABI instead.
var TellorGettersABI = TellorGettersMetaData.ABI

// Deprecated: Use TellorGettersMetaData.Sigs instead.
// TellorGettersFuncSigs maps the 4-byte function signature to its string representation.
var TellorGettersFuncSigs = TellorGettersMetaData.Sigs

// TellorGettersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TellorGettersMetaData.Bin instead.
var TellorGettersBin = TellorGettersMetaData.Bin

// DeployTellorGetters deploys a new Ethereum contract, binding an instance of TellorGetters to it.
func DeployTellorGetters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorGetters, error) {
	parsed, err := TellorGettersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TellorGettersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorGetters{TellorGettersCaller: TellorGettersCaller{contract: contract}, TellorGettersTransactor: TellorGettersTransactor{contract: contract}, TellorGettersFilterer: TellorGettersFilterer{contract: contract}}, nil
}

// TellorGetters is an auto generated Go binding around an Ethereum contract.
type TellorGetters struct {
	TellorGettersCaller     // Read-only binding to the contract
	TellorGettersTransactor // Write-only binding to the contract
	TellorGettersFilterer   // Log filterer for contract events
}

// TellorGettersCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorGettersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorGettersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorGettersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorGettersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorGettersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorGettersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorGettersSession struct {
	Contract     *TellorGetters    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorGettersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorGettersCallerSession struct {
	Contract *TellorGettersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TellorGettersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorGettersTransactorSession struct {
	Contract     *TellorGettersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TellorGettersRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorGettersRaw struct {
	Contract *TellorGetters // Generic contract binding to access the raw methods on
}

// TellorGettersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorGettersCallerRaw struct {
	Contract *TellorGettersCaller // Generic read-only contract binding to access the raw methods on
}

// TellorGettersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorGettersTransactorRaw struct {
	Contract *TellorGettersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorGetters creates a new instance of TellorGetters, bound to a specific deployed contract.
func NewTellorGetters(address common.Address, backend bind.ContractBackend) (*TellorGetters, error) {
	contract, err := bindTellorGetters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorGetters{TellorGettersCaller: TellorGettersCaller{contract: contract}, TellorGettersTransactor: TellorGettersTransactor{contract: contract}, TellorGettersFilterer: TellorGettersFilterer{contract: contract}}, nil
}

// NewTellorGettersCaller creates a new read-only instance of TellorGetters, bound to a specific deployed contract.
func NewTellorGettersCaller(address common.Address, caller bind.ContractCaller) (*TellorGettersCaller, error) {
	contract, err := bindTellorGetters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorGettersCaller{contract: contract}, nil
}

// NewTellorGettersTransactor creates a new write-only instance of TellorGetters, bound to a specific deployed contract.
func NewTellorGettersTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorGettersTransactor, error) {
	contract, err := bindTellorGetters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorGettersTransactor{contract: contract}, nil
}

// NewTellorGettersFilterer creates a new log filterer instance of TellorGetters, bound to a specific deployed contract.
func NewTellorGettersFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorGettersFilterer, error) {
	contract, err := bindTellorGetters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorGettersFilterer{contract: contract}, nil
}

// bindTellorGetters binds a generic wrapper to an already deployed contract.
func bindTellorGetters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorGettersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorGetters *TellorGettersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorGetters.Contract.TellorGettersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorGetters *TellorGettersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorGetters.Contract.TellorGettersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorGetters *TellorGettersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorGetters.Contract.TellorGettersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorGetters *TellorGettersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorGetters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorGetters *TellorGettersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorGetters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorGetters *TellorGettersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorGetters.Contract.contract.Transact(opts, method, params...)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "_allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorGetters *TellorGettersSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorGetters.Contract.Allowances(&_TellorGetters.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorGetters.Contract.Allowances(&_TellorGetters.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorGetters *TellorGettersCaller) Addresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorGetters *TellorGettersSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorGetters.Contract.Addresses(&_TellorGetters.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorGetters *TellorGettersCallerSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorGetters.Contract.Addresses(&_TellorGetters.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorGetters *TellorGettersCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "balances", arg0, arg1)

	outstruct := new(struct {
		FromBlock *big.Int
		Value     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorGetters *TellorGettersSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorGetters.Contract.Balances(&_TellorGetters.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorGetters *TellorGettersCallerSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorGetters.Contract.Balances(&_TellorGetters.CallOpts, arg0, arg1)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorGetters *TellorGettersCaller) BytesVars(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "bytesVars", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorGetters *TellorGettersSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorGetters.Contract.BytesVars(&_TellorGetters.CallOpts, arg0)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorGetters *TellorGettersCallerSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorGetters.Contract.BytesVars(&_TellorGetters.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorGetters *TellorGettersCaller) CurrentMiners(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "currentMiners", arg0)

	outstruct := new(struct {
		Value *big.Int
		Miner common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Miner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorGetters *TellorGettersSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorGetters.Contract.CurrentMiners(&_TellorGetters.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorGetters *TellorGettersCallerSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorGetters.Contract.CurrentMiners(&_TellorGetters.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_TellorGetters *TellorGettersCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_TellorGetters *TellorGettersSession) Decimals() (uint8, error) {
	return _TellorGetters.Contract.Decimals(&_TellorGetters.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_TellorGetters *TellorGettersCallerSession) Decimals() (uint8, error) {
	return _TellorGetters.Contract.Decimals(&_TellorGetters.CallOpts)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_TellorGetters *TellorGettersCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "didMine", _challenge, _miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_TellorGetters *TellorGettersSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _TellorGetters.Contract.DidMine(&_TellorGetters.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(bytes32 _challenge, address _miner) view returns(bool)
func (_TellorGetters *TellorGettersCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _TellorGetters.Contract.DidMine(&_TellorGetters.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_TellorGetters *TellorGettersCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "didVote", _disputeId, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_TellorGetters *TellorGettersSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _TellorGetters.Contract.DidVote(&_TellorGetters.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(uint256 _disputeId, address _address) view returns(bool)
func (_TellorGetters *TellorGettersCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _TellorGetters.Contract.DidVote(&_TellorGetters.CallOpts, _disputeId, _address)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) DisputeIdByDisputeHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "disputeIdByDisputeHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.DisputeIdByDisputeHash(&_TellorGetters.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.DisputeIdByDisputeHash(&_TellorGetters.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorGetters *TellorGettersCaller) DisputesById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "disputesById", arg0)

	outstruct := new(struct {
		Hash                [32]byte
		Tally               *big.Int
		Executed            bool
		DisputeVotePassed   bool
		IsPropFork          bool
		ReportedMiner       common.Address
		ReportingParty      common.Address
		ProposedForkAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Tally = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Executed = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.DisputeVotePassed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.IsPropFork = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.ReportedMiner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.ReportingParty = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.ProposedForkAddress = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorGetters *TellorGettersSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorGetters.Contract.DisputesById(&_TellorGetters.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorGetters *TellorGettersCallerSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorGetters.Contract.DisputesById(&_TellorGetters.CallOpts, arg0)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_TellorGetters *TellorGettersCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getAddressVars", _data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_TellorGetters *TellorGettersSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _TellorGetters.Contract.GetAddressVars(&_TellorGetters.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(bytes32 _data) view returns(address)
func (_TellorGetters *TellorGettersCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _TellorGetters.Contract.GetAddressVars(&_TellorGetters.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_TellorGetters *TellorGettersCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getAllDisputeVars", _disputeId)

	if err != nil {
		return *new([32]byte), *new(bool), *new(bool), *new(bool), *new(common.Address), *new(common.Address), *new(common.Address), *new([9]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	out6 := *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	out7 := *abi.ConvertType(out[7], new([9]*big.Int)).(*[9]*big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_TellorGetters *TellorGettersSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetAllDisputeVars(&_TellorGetters.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(uint256 _disputeId) view returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_TellorGetters *TellorGettersCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetAllDisputeVars(&_TellorGetters.CallOpts, _disputeId)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getDisputeIdByDisputeHash", _hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeIdByDisputeHash(&_TellorGetters.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(bytes32 _hash) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeIdByDisputeHash(&_TellorGetters.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getDisputeUintVars", _disputeId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeUintVars(&_TellorGetters.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(uint256 _disputeId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetDisputeUintVars(&_TellorGetters.CallOpts, _disputeId, _data)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_TellorGetters *TellorGettersCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getLastNewValueById", _requestId)

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
func (_TellorGetters *TellorGettersSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _TellorGetters.Contract.GetLastNewValueById(&_TellorGetters.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(uint256 _requestId) view returns(uint256, bool)
func (_TellorGetters *TellorGettersCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _TellorGetters.Contract.GetLastNewValueById(&_TellorGetters.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getMinedBlockNum", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetMinedBlockNum(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetMinedBlockNum(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_TellorGetters *TellorGettersCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([5]common.Address)).(*[5]common.Address)

	return out0, err

}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_TellorGetters *TellorGettersSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _TellorGetters.Contract.GetMinersByRequestIdAndTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(uint256 _requestId, uint256 _timestamp) view returns(address[5])
func (_TellorGetters *TellorGettersCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _TellorGetters.Contract.GetMinersByRequestIdAndTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _diff, uint256 _tip)
func (_TellorGetters *TellorGettersCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Diff       *big.Int
	Tip        *big.Int
}, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getNewCurrentVariables")

	outstruct := new(struct {
		Challenge  [32]byte
		RequestIds [5]*big.Int
		Diff       *big.Int
		Tip        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Challenge = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RequestIds = *abi.ConvertType(out[1], new([5]*big.Int)).(*[5]*big.Int)
	outstruct.Diff = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Tip = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _diff, uint256 _tip)
func (_TellorGetters *TellorGettersSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Diff       *big.Int
	Tip        *big.Int
}, error) {
	return _TellorGetters.Contract.GetNewCurrentVariables(&_TellorGetters.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _diff, uint256 _tip)
func (_TellorGetters *TellorGettersCallerSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Diff       *big.Int
	Tip        *big.Int
}, error) {
	return _TellorGetters.Contract.GetNewCurrentVariables(&_TellorGetters.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetNewValueCountbyRequestId(&_TellorGetters.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetNewValueCountbyRequestId(&_TellorGetters.CallOpts, _requestId)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_TellorGetters *TellorGettersCaller) GetNewVariablesOnDeck(opts *bind.CallOpts) (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getNewVariablesOnDeck")

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
func (_TellorGetters *TellorGettersSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _TellorGetters.Contract.GetNewVariablesOnDeck(&_TellorGetters.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_TellorGetters *TellorGettersCallerSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _TellorGetters.Contract.GetNewVariablesOnDeck(&_TellorGetters.CallOpts)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestIdByRequestQIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByRequestQIndex(&_TellorGetters.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestIdByRequestQIndex(&_TellorGetters.CallOpts, _index)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_TellorGetters *TellorGettersCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestQ")

	if err != nil {
		return *new([51]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([51]*big.Int)).(*[51]*big.Int)

	return out0, err

}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_TellorGetters *TellorGettersSession) GetRequestQ() ([51]*big.Int, error) {
	return _TellorGetters.Contract.GetRequestQ(&_TellorGetters.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() view returns(uint256[51])
func (_TellorGetters *TellorGettersCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _TellorGetters.Contract.GetRequestQ(&_TellorGetters.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestUintVars", _requestId, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestUintVars(&_TellorGetters.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(uint256 _requestId, bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetRequestUintVars(&_TellorGetters.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_TellorGetters *TellorGettersCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getRequestVars", _requestId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_TellorGetters *TellorGettersSession) GetRequestVars(_requestId *big.Int) (*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetRequestVars(&_TellorGetters.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(uint256 _requestId) view returns(uint256, uint256)
func (_TellorGetters *TellorGettersCallerSession) GetRequestVars(_requestId *big.Int) (*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetRequestVars(&_TellorGetters.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_TellorGetters *TellorGettersCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getStakerInfo", _staker)

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
func (_TellorGetters *TellorGettersSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetStakerInfo(&_TellorGetters.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address _staker) view returns(uint256, uint256)
func (_TellorGetters *TellorGettersCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _TellorGetters.Contract.GetStakerInfo(&_TellorGetters.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_TellorGetters *TellorGettersCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getSubmissionsByTimestamp", _requestId, _timestamp)

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_TellorGetters *TellorGettersSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _TellorGetters.Contract.GetSubmissionsByTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(uint256 _requestId, uint256 _timestamp) view returns(uint256[5])
func (_TellorGetters *TellorGettersCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _TellorGetters.Contract.GetSubmissionsByTimestamp(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestID, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetTimestampbyRequestIDandIndex(&_TellorGetters.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestID, uint256 _index) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.GetTimestampbyRequestIDandIndex(&_TellorGetters.CallOpts, _requestID, _index)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_TellorGetters *TellorGettersCaller) GetTopRequestIDs(opts *bind.CallOpts) ([5]*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getTopRequestIDs")

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_TellorGetters *TellorGettersSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _TellorGetters.Contract.GetTopRequestIDs(&_TellorGetters.CallOpts)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_TellorGetters *TellorGettersCallerSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _TellorGetters.Contract.GetTopRequestIDs(&_TellorGetters.CallOpts)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "getUintVar", _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetUintVar(&_TellorGetters.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(bytes32 _data) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.GetUintVar(&_TellorGetters.CallOpts, _data)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_TellorGetters *TellorGettersCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_TellorGetters *TellorGettersSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _TellorGetters.Contract.IsInDispute(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_TellorGetters *TellorGettersCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _TellorGetters.Contract.IsInDispute(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorGetters *TellorGettersCaller) Migrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "migrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorGetters *TellorGettersSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorGetters.Contract.Migrated(&_TellorGetters.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorGetters *TellorGettersCallerSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorGetters.Contract.Migrated(&_TellorGetters.CallOpts, arg0)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorGetters *TellorGettersCaller) MinersByChallenge(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "minersByChallenge", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorGetters *TellorGettersSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorGetters.Contract.MinersByChallenge(&_TellorGetters.CallOpts, arg0, arg1)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorGetters *TellorGettersCallerSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorGetters.Contract.MinersByChallenge(&_TellorGetters.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_TellorGetters *TellorGettersCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_TellorGetters *TellorGettersSession) Name() (string, error) {
	return _TellorGetters.Contract.Name(&_TellorGetters.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_TellorGetters *TellorGettersCallerSession) Name() (string, error) {
	return _TellorGetters.Contract.Name(&_TellorGetters.CallOpts)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) NewValueTimestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "newValueTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorGetters *TellorGettersSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.NewValueTimestamps(&_TellorGetters.CallOpts, arg0)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.NewValueTimestamps(&_TellorGetters.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) RequestIdByQueryHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "requestIdByQueryHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.RequestIdByQueryHash(&_TellorGetters.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.RequestIdByQueryHash(&_TellorGetters.CallOpts, arg0)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.RetrieveData(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorGetters.Contract.RetrieveData(&_TellorGetters.CallOpts, _requestId, _timestamp)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_TellorGetters *TellorGettersCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_TellorGetters *TellorGettersSession) Symbol() (string, error) {
	return _TellorGetters.Contract.Symbol(&_TellorGetters.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_TellorGetters *TellorGettersCallerSession) Symbol() (string, error) {
	return _TellorGetters.Contract.Symbol(&_TellorGetters.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TellorGetters *TellorGettersCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TellorGetters *TellorGettersSession) TotalSupply() (*big.Int, error) {
	return _TellorGetters.Contract.TotalSupply(&_TellorGetters.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) TotalSupply() (*big.Int, error) {
	return _TellorGetters.Contract.TotalSupply(&_TellorGetters.CallOpts)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersCaller) Uints(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorGetters.contract.Call(opts, &out, "uints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.Uints(&_TellorGetters.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorGetters *TellorGettersCallerSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorGetters.Contract.Uints(&_TellorGetters.CallOpts, arg0)
}

// TellorStakeMetaData contains all meta data concerning the TellorStake contract.
var TellorStakeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"fromBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentMiners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"disputesById\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"tally\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disputeVotePassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPropFork\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"reportedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reportingParty\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedForkAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minersByChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newValueTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_propNewTellorAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"unlockDisputeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"024c2ddd": "_allowances(address,address)",
		"699f200f": "addresses(bytes32)",
		"dd62ed3e": "allowance(address,address)",
		"999cf26c": "allowedToTrade(address,uint256)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"4ee2cd7e": "balanceOfAt(address,uint256)",
		"cbf1304d": "balances(address,uint256)",
		"8581af19": "beginDispute(uint256,uint256,uint256)",
		"62dd1d2a": "bytesVars(bytes32)",
		"1fd22364": "currentMiners(uint256)",
		"d01f4d9e": "disputeIdByDisputeHash(bytes32)",
		"db085beb": "disputesById(uint256)",
		"4ba0a5ee": "migrated(address)",
		"48b18e54": "minersByChallenge(bytes32,address)",
		"438c0aa3": "newValueTimestamps(uint256)",
		"26b7d9f6": "proposeFork(address)",
		"5700242c": "requestIdByQueryHash(bytes32)",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"b59e14d4": "uints(bytes32)",
		"9a01ca13": "unlockDisputeFee(uint256)",
		"fc735e99": "verify()",
		"c9d27afe": "vote(uint256,bool)",
	},
	Bin: "0x608060405234801561001057600080fd5b50612763806100206000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c806370a08231116100c3578063c9d27afe1161007c578063c9d27afe14610468578063cbf1304d1461048d578063d01f4d9e146104e8578063db085beb14610505578063dd62ed3e14610572578063fc735e99146105a057610158565b806370a08231146103875780638581af19146103ad578063999cf26c146103d65780639a01ca1314610402578063a9059cbb1461041f578063b59e14d41461044b57610158565b806348b18e541161011557806348b18e54146102965780634ba0a5ee146102c25780634ee2cd7e146102e85780635700242c1461031457806362dd1d2a14610331578063699f200f1461034e57610158565b8063024c2ddd1461015d578063095ea7b31461019d5780631fd22364146101dd57806323b872dd1461021b57806326b7d9f614610251578063438c0aa314610279575b600080fd5b61018b6004803603604081101561017357600080fd5b506001600160a01b03813581169160200135166105a8565b60408051918252519081900360200190f35b6101c9600480360360408110156101b357600080fd5b506001600160a01b0381351690602001356105c5565b604080519115158252519081900360200190f35b6101fa600480360360208110156101f357600080fd5b5035610674565b604080519283526001600160a01b0390911660208301528051918290030190f35b6101c96004803603606081101561023157600080fd5b506001600160a01b0381358116916020810135909116906040013561069f565b6102776004803603602081101561026757600080fd5b50356001600160a01b031661074b565b005b61018b6004803603602081101561028f57600080fd5b5035610a50565b6101c9600480360360408110156102ac57600080fd5b50803590602001356001600160a01b0316610a71565b6101c9600480360360208110156102d857600080fd5b50356001600160a01b0316610a91565b61018b600480360360408110156102fe57600080fd5b506001600160a01b038135169060200135610aa6565b61018b6004803603602081101561032a57600080fd5b5035610c51565b61018b6004803603602081101561034757600080fd5b5035610c63565b61036b6004803603602081101561036457600080fd5b5035610c75565b604080516001600160a01b039092168252519081900360200190f35b61018b6004803603602081101561039d57600080fd5b50356001600160a01b0316610c90565b610277600480360360608110156103c357600080fd5b5080359060208101359060400135610c9c565b6101c9600480360360408110156103ec57600080fd5b506001600160a01b03813516906020013561120b565b6102776004803603602081101561041857600080fd5b50356112d7565b6101c96004803603604081101561043557600080fd5b506001600160a01b038135169060200135611b6d565b61018b6004803603602081101561046157600080fd5b5035611b83565b6102776004803603604081101561047e57600080fd5b50803590602001351515611b95565b6104b9600480360360408110156104a357600080fd5b506001600160a01b038135169060200135611e77565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b61018b600480360360208110156104fe57600080fd5b5035611eba565b6105226004803603602081101561051b57600080fd5b5035611ecc565b60408051988952602089019790975294151587870152921515606087015290151560808601526001600160a01b0390811660a086015290811660c08501521660e083015251908190036101000190f35b61018b6004803603604081101561058857600080fd5b506001600160a01b0381358116916020013516611f28565b61018b611f53565b604a60209081526000928352604080842090915290825290205481565b60006001600160a01b03831661060c5760405162461bcd60e51b81526004018080602001828103825260228152602001806125e26022913960400191505060405180910390fd5b336000818152604a602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b603a816005811061068457600080fd5b6002020180546001909101549091506001600160a01b031682565b6001600160a01b0383166000908152604a6020908152604080832033845290915281205482111561070c576040805162461bcd60e51b8152602060048201526012602482015271416c6c6f77616e63652069732077726f6e6760701b604482015290519081900360640190fd5b6001600160a01b0384166000908152604a60209081526040808320338452909152902080548390039055610741848484611f59565b5060019392505050565b7fd051321aa26ce60d202f153d0c0e67687e975532ab88ce92d84f18e39895d90760005260466020527f7042aca2505aa9a8428c1710298edbac96e027637d6417651c8aa8f88ed1ca0a54156107d7576040805162461bcd60e51b815260206004820152600c60248201526b6e6f2072656e7472616e637960a01b604482015290519081900360640190fd5b7fd051321aa26ce60d202f153d0c0e67687e975532ab88ce92d84f18e39895d907600052604660205260017f7042aca2505aa9a8428c1710298edbac96e027637d6417651c8aa8f88ed1ca0a5561082d8161210d565b7fd051321aa26ce60d202f153d0c0e67687e975532ab88ce92d84f18e39895d9076000908152604660209081527f7042aca2505aa9a8428c1710298edbac96e027637d6417651c8aa8f88ed1ca0a829055604080516001600160a01b0385168184015281518082038401815290820182528051908301206000805160206126b783398151915280546001019081905581855260389093529220541561091a57600082815260386020908152604080832054848452603683528184207fed92b4c1e0a9e559a31171d487ecbec963526662038ecfa3a71160bd62fb873385526005019092529091205561092c565b60008281526038602052604090208190555b600082815260386020526040812054610946908390612267565b6000838152603660205260408120858155600280820180546201000062ff000019909116176301000000600160b81b0319163363010000008102919091179091556003830180546001600160a01b031990811683179091556004840180549091166001600160a01b038b161790556001909201929092559192506109de919030906000198501900a68056bc75e2d6310000002611f59565b5060009081526036602090815260408083207f4b4cefd5ced7569ef0d091282b4bca9c52a034c56471a6061afd1bf307a2de7c84526005019091528082204390557f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d28825290204262093a800190555050565b60338181548110610a6057600080fd5b600091825260209091200154905081565b603960209081526000928352604080842090915290825290205460ff1681565b604b6020526000908152604090205460ff1681565b6001600160a01b038216600090815260496020526040812080541580610aec57508281600081548110610ad557fe5b6000918252602090912001546001600160801b0316115b15610afb57600091505061066e565b805481906000198101908110610b0d57fe5b6000918252602090912001546001600160801b03168310610b5f57805481906000198101908110610b3a57fe5b600091825260209091200154600160801b90046001600160801b0316915061066e9050565b8054600090600119015b81811115610c17576000600260018385010104905085848281548110610b8b57fe5b6000918252602090912001546001600160801b03161415610bda57838181548110610bb257fe5b600091825260209091200154600160801b90046001600160801b0316945061066e9350505050565b85848281548110610be757fe5b6000918252602090912001546001600160801b03161015610c0a57809250610c11565b6001810391505b50610b69565b828281548110610c2357fe5b600091825260209091200154600160801b90046001600160801b0316935061066e92505050565b5092915050565b60376020526000908152604090205481565b60486020526000908152604090205481565b6047602052600090815260409020546001600160a01b031681565b600061066e8243610aa6565b60008381526045602090815260408083208584526002810190925290912054610cff576040805162461bcd60e51b815260206004820152601060248201526f04d696e656420626c6f636b20697320360841b604482015290519081900360640190fd5b60058210610d4b576040805162461bcd60e51b81526020600482015260146024820152734d696e657220696e6465782069732077726f6e6760601b604482015290519081900360640190fd5b60008381526005808301602052604082209084908110610d6757fe5b015460408051606083811b6001600160601b03191660208084018290526b111a5cdc1d5d1950dbdd5b9d60a21b603485015284518085038201815284860186528051908201206000908152604682528581208054600190810190915593850192909252607484018b905260948085018b90528551808603909101815260b490940185528351938101939093206000805160206126b7833981519152805490930192839055808252603890935292909220546001600160a01b039093169350918015610e705760008281526036602090815260408083207fed92b4c1e0a9e559a31171d487ecbec963526662038ecfa3a71160bd62fb873384526005019091529020819055610ec7565b62093a8087420310610eb35760405162461bcd60e51b81526004018080602001828103825260328152602001806126d76032913960400191505060405180910390fd5b506000828152603860205260409020819055805b6000610ed38383612267565b905060008760021415610f665750600089815260456020908152604080832060008051602061262f8339815191528452600190810183529083208054909101908190557f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc97909252604690527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be5402610fb5565b507f675d2171f68d6f5545d54fb9b1fb61a0e6897e6188ca1cd664e7c9530d91ecfc60005260466020527f3e5522f19747f0f285b96ded572ac4128c3a764aea9f44058dc0afc9dda449865481025b6000848152603660209081526040808320888155600281018054600383018054336001600160a01b031991821617909155600484018054909116905562010000600160b81b03191663010000006001600160a01b038d16021761ffff19169055600181018490557f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f8452600590810183528184208e905560008051602061267683398151915284528184208d90558c845260068b01909252909120908990811061107b57fe5b015460008581526036602090815260408083207f9147231ab14efb72c38117f68521ddef8de64f092c18c69dbfb602ffc4de7f478452600501909152808220929092557f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d2881528181206202a3008502420190557f4b4cefd5ced7569ef0d091282b4bca9c52a034c56471a6061afd1bf307a2de7c81528181204390557f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d81528181208a90557f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc409815220819055611172333083611f59565b87600214156111a45760008981526004880160209081526040808320805460ff1916600117905560038a019091528120555b6001600160a01b0386166000818152604460209081526040918290206003905581518c81529081019290925280518c9287927feceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da6492918290030190a350505050505050505050565b6001600160a01b0382166000908152604460205260408120541580159061124a57506001600160a01b0383166000908152604460205260409020546005115b156112c4577f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc9760005260466020527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be5482906112af906112a986610c90565b90612459565b106112bc5750600161066e565b50600061066e565b816112ce84610c90565b10159392505050565b60008051602061262f83398151915260005260466020526000805160206126b78339815191525481111561134b576040805162461bcd60e51b8152602060048201526016602482015275191a5cdc1d5d1948191bd95cc81b9bdd08195e1a5cdd60521b604482015290519081900360640190fd5b600081815260366020818152604080842054845260388252808420548085529282528084207f6ab2b18aafe78fd59c6a4092015bddd9fcacb8170f72b299074f74d76a91a9238552600501808352818520548251808501919091528251808203850181529083018352805190840120855290915290912054806113cb5750805b60008281526036602090815260408083208484528184207f6ab2b18aafe78fd59c6a4092015bddd9fcacb8170f72b299074f74d76a91a92385526005820190935292205480611418575060015b7f29169706298d2b6df50a532e958b56426de1465348b93650fca42d456eaec5fc60009081526005840160205260408120541561148f576040805162461bcd60e51b815260206004820152601060248201526f185b1c9958591e481c185a59081bdd5d60821b604482015290519081900360640190fd5b600284015462010000900460ff16156114d95760405162461bcd60e51b81526004018080602001828103825260298152602001806125b96029913960400191505060405180910390fd5b7ff9e1ae10923bfc79f52e309baf8c7699edb821f91ef5b5bd07be29545917b3a6600090815260058401602052604090205461155c576040805162461bcd60e51b815260206004820152601860248201527f766f7465206e6565647320746f2062652074616c6c6965640000000000000000604482015290519081900360640190fd5b7ff9e1ae10923bfc79f52e309baf8c7699edb821f91ef5b5bd07be29545917b3a66000908152600584016020526040902054620151804291909103116115d35760405162461bcd60e51b815260040180806020018281038252602b815260200180612604602b913960400191505060405180910390fd5b600284810154630100000090046001600160a01b031660009081526044602090815260408083207f29169706298d2b6df50a532e958b56426de1465348b93650fca42d456eaec5fc84526005890190925290912060019081905591850154909161010090910460ff16151514156117fc576201518042064203600182015560466020527fa5ae3e2b97d73fb849ea855d27f073b72815b38452d976bd57e4a157827dadd380546000190190557f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc976000527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be54600286015481906116e590630100000090046001600160a01b0316610c90565b101561170c57600286015461170990630100000090046001600160a01b0316610c90565b90505b815460041415611747576005825560028601546003870154611742916001600160a01b0363010000009091048116911683611f59565b600082555b60005b848110156117f557604080518287036020808301919091528251808303820181529183018352815191810191909120600090815260058a019091522054935083611792578893505b600084815260366020908152604080832060038101547f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc4098552600582019093529220546117ec9130916001600160a01b0390911690611f59565b5060010161174a565b5050611a68565b600284015460408051630100000090920460601b6001600160601b0319166020808401919091526b111a5cdc1d5d1950dbdd5b9d60a21b60348401528151808403820181529282018252825192810192909220600090815260469092529020546001141561186957600181555b7f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f60009081526005860160208181526040808420548452604582528084207f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d8552929091529091205460021415611932577f9147231ab14efb72c38117f68521ddef8de64f092c18c69dbfb602ffc4de7f476000908152600587016020908152604080832054600080516020612676833981519152845281842054845260038501909252909120555b600080516020612676833981519152600090815260058701602090815260408083205483526004840190915290205460ff161515600114156119a55760008051602061267683398151915260009081526005870160209081526040808320548352600484019091529020805460ff191690555b60005b84811015611a6557604080518287036020808301919091528251808303820181529183018352815191810191909120600090815260058a019091522054935083156119fe57600084815260366020526040902095505b600286015460008581526036602090815260408083207f1da95f11543c9b03927178e07951795dfc95c7501a9d1cf00e13414ca33bc4098452600501909152902054611a5d91309163010000009091046001600160a01b031690611f59565b6001016119a8565b50505b60028481015460408051630100000090920460601b6001600160601b0319166020808401919091526b111a5cdc1d5d1950dbdd5b9d60a21b6034840152815180840382018152928201825282519281019290922060009081526046835281812080546000190190557f6de96ee4d33a0617f40a846309c8759048857f51b9d59a12d3c3786d4778883d8152600589019092529020541415611b63577f9f47a2659c3d32b749ae717d975e7962959890862423c4318cf86e4ec220291f600090815260058601602090815260408083205483526045825280832060008051602061262f8339815191528452600101909152902080546000190190555b5050505050505050565b6000611b7a338484611f59565b50600192915050565b60466020526000908152604090205481565b60008051602061262f83398151915260005260466020526000805160206126b783398151915254821115611c09576040805162461bcd60e51b8152602060048201526016602482015275191a5cdc1d5d1948191bd95cc81b9bdd08195e1a5cdd60521b604482015290519081900360640190fd5b6000828152603660205260409020600281015460ff1615611c5b5760405162461bcd60e51b81526004018080602001828103825260258152602001806127096025913960400191505060405180910390fd5b7f4b4cefd5ced7569ef0d091282b4bca9c52a034c56471a6061afd1bf307a2de7c6000908152600582016020526040812054611c98903390610aa6565b33600090815260068401602052604090205490915060ff16151560011415611d07576040805162461bcd60e51b815260206004820152601860248201527f53656e6465722068617320616c726561647920766f7465640000000000000000604482015290519081900360640190fd5b80611d4d576040805162461bcd60e51b81526020600482015260116024820152700557365722062616c616e6365206973203607c1b604482015290519081900360640190fd5b3360009081526044602052604090205460031415611dab576040805162461bcd60e51b81526020600482015260166024820152754d696e657220697320756e646572206469737075746560501b604482015290519081900360640190fd5b3360009081526006830160209081526040808320805460ff191660019081179091557f1da378694063870452ce03b189f48e04c1aa026348e74e6c86e10738514ad2c4845260058601909252909120805490910190558215611e20576001820154611e16908261246b565b6001830155611e35565b6001820154611e2f9082612496565b60018301555b60408051841515815290518291339187917f911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e919081900360200190a450505050565b60496020528160005260406000208181548110611e9357600080fd5b6000918252602090912001546001600160801b038082169350600160801b90910416905082565b60386020526000908152604090205481565b603660205260009081526040902080546001820154600283015460038401546004909401549293919260ff808316936101008404821693620100008104909216926001600160a01b036301000000909304831692918216911688565b6001600160a01b039182166000908152604a6020908152604080832093909416825291909152205490565b610bb790565b80611f955760405162461bcd60e51b81526004018080602001828103825260218152602001806126966021913960400191505060405180910390fd5b6001600160a01b038216611fe8576040805162461bcd60e51b815260206004820152601560248201527452656365697665722069732030206164647265737360581b604482015290519081900360640190fd5b611ff2838261120b565b61202d5760405162461bcd60e51b815260040180806020018281038252602781526020018061264f6027913960400191505060405180910390fd5b600061203884610c90565b905081612047858284036124bc565b61205084610c90565b9150816001600160801b03168183016001600160801b031610156120af576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b6120bb848284016124bc565b836001600160a01b0316856001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a35050505050565b60408051602060248201819052600060448301819052835180840360640181526084909301845290820180516001600160e01b031663fc735e9960e01b1781529251825191936060936001600160a01b0387169390929182918083835b602083106121895780518252601f19909201916020918201910161216a565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146121eb576040519150601f19603f3d011682016040523d82523d6000602084013e6121f0565b606091505b50915091508180156122195750610bb781806020019051602081101561221557600080fd5b5051115b612262576040805162461bcd60e51b81526020600482015260156024820152741b995dc81d195b1b1bdc881a5cc81a5b9d985b1a59605a1b604482015290519081900360640190fd5b505050565b60008181526036602081815260408084207f6ab2b18aafe78fd59c6a4092015bddd9fcacb8170f72b299074f74d76a91a923855260050180835281852080546001019081905586865293835281518084018590528251808203850181529083018352805190840120855290915290912083905582821461066e57600082815260366020818152604080842081516000198701818501528251808203850181529083018352805190840120855260059081018352818520548086529383528185207f46f7d53798d31923f6952572c6a19ad2d1a8238d26649c2f3493a6d69e425d28865201909152909120544210156123a6576040805162461bcd60e51b815260206004820152601760248201527f4469737075746520697320616c7265616479206f70656e000000000000000000604482015290519081900360640190fd5b60008181526036602052604090206002015460ff1615610c4a5760008181526036602090815260408083207ff9e1ae10923bfc79f52e309baf8c7699edb821f91ef5b5bd07be29545917b3a684526005019091529020546201518042919091031115610c4a576040805162461bcd60e51b815260206004820152601f60248201527f54696d6520666f7220766f74696e6720686176656e277420656c617073656400604482015290519081900360640190fd5b60008282111561246557fe5b50900390565b60008082131561248857508181018281121561248357fe5b61066e565b508181018281131561066e57fe5b6000808213156124ae57508082038281131561248357fe5b508082038281121561066e57fe5b6001600160a01b038216600090815260496020526040902080541580612509575080544390829060001981019081106124f157fe5b6000918252602090912001546001600160801b031614155b1561257a5760408051808201909152436001600160801b0390811682528381166020808401918252845460018101865560008681529190912093519301805491516fffffffffffffffffffffffffffffffff19909216938316939093178216600160801b9190921602179055612262565b80546000908290600019810190811061258f57fe5b600091825260209091200180546001600160801b03808616600160801b0291161790555050505056fe66756e6374696f6e206e6f742063616c6c61626c6520666f726b20666f726b2070726f706f73616c7345524332303a20617070726f766520746f20746865207a65726f206164647265737354696d6520666f72206120666f6c6c6f772075702064697370757465206861736e277420656c6170736564310199159a20c50879ffb440b45802138b5b162ec9426720e9dd3ee8bbcdb9d753686f756c6420686176652073756666696369656e742062616c616e636520746f2074726164652f9328a9c75282bec25bb04befad06926366736e0030c985108445fa728335e5547269656420746f2073656e64206e6f6e2d706f73697469766520616d6f756e741ce2382bc92689b00ba121fa5a411aa976168affdd8ac143a69035dd984b3b6a44697370757465206d75737420626520737461727465642077697468696e2061207765656b206f66206261642076616c756574686520646973707574652068617320616c7265616479206265656e206578656375746564a2646970667358221220c2c9c7925204e5f8e35be029a5158670f2ca3f6381133bfd71cf71a148b56b9364736f6c63430007040033",
}

// TellorStakeABI is the input ABI used to generate the binding from.
// Deprecated: Use TellorStakeMetaData.ABI instead.
var TellorStakeABI = TellorStakeMetaData.ABI

// Deprecated: Use TellorStakeMetaData.Sigs instead.
// TellorStakeFuncSigs maps the 4-byte function signature to its string representation.
var TellorStakeFuncSigs = TellorStakeMetaData.Sigs

// TellorStakeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TellorStakeMetaData.Bin instead.
var TellorStakeBin = TellorStakeMetaData.Bin

// DeployTellorStake deploys a new Ethereum contract, binding an instance of TellorStake to it.
func DeployTellorStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorStake, error) {
	parsed, err := TellorStakeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TellorStakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorStake{TellorStakeCaller: TellorStakeCaller{contract: contract}, TellorStakeTransactor: TellorStakeTransactor{contract: contract}, TellorStakeFilterer: TellorStakeFilterer{contract: contract}}, nil
}

// TellorStake is an auto generated Go binding around an Ethereum contract.
type TellorStake struct {
	TellorStakeCaller     // Read-only binding to the contract
	TellorStakeTransactor // Write-only binding to the contract
	TellorStakeFilterer   // Log filterer for contract events
}

// TellorStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorStakeSession struct {
	Contract     *TellorStake      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorStakeCallerSession struct {
	Contract *TellorStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TellorStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorStakeTransactorSession struct {
	Contract     *TellorStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TellorStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorStakeRaw struct {
	Contract *TellorStake // Generic contract binding to access the raw methods on
}

// TellorStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorStakeCallerRaw struct {
	Contract *TellorStakeCaller // Generic read-only contract binding to access the raw methods on
}

// TellorStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorStakeTransactorRaw struct {
	Contract *TellorStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorStake creates a new instance of TellorStake, bound to a specific deployed contract.
func NewTellorStake(address common.Address, backend bind.ContractBackend) (*TellorStake, error) {
	contract, err := bindTellorStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorStake{TellorStakeCaller: TellorStakeCaller{contract: contract}, TellorStakeTransactor: TellorStakeTransactor{contract: contract}, TellorStakeFilterer: TellorStakeFilterer{contract: contract}}, nil
}

// NewTellorStakeCaller creates a new read-only instance of TellorStake, bound to a specific deployed contract.
func NewTellorStakeCaller(address common.Address, caller bind.ContractCaller) (*TellorStakeCaller, error) {
	contract, err := bindTellorStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorStakeCaller{contract: contract}, nil
}

// NewTellorStakeTransactor creates a new write-only instance of TellorStake, bound to a specific deployed contract.
func NewTellorStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorStakeTransactor, error) {
	contract, err := bindTellorStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorStakeTransactor{contract: contract}, nil
}

// NewTellorStakeFilterer creates a new log filterer instance of TellorStake, bound to a specific deployed contract.
func NewTellorStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorStakeFilterer, error) {
	contract, err := bindTellorStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorStakeFilterer{contract: contract}, nil
}

// bindTellorStake binds a generic wrapper to an already deployed contract.
func bindTellorStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorStakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorStake *TellorStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorStake.Contract.TellorStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorStake *TellorStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorStake.Contract.TellorStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorStake *TellorStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorStake.Contract.TellorStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorStake *TellorStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorStake *TellorStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorStake *TellorStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorStake.Contract.contract.Transact(opts, method, params...)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorStake *TellorStakeCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "_allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorStake *TellorStakeSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorStake.Contract.Allowances(&_TellorStake.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorStake.Contract.Allowances(&_TellorStake.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorStake *TellorStakeCaller) Addresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorStake *TellorStakeSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorStake.Contract.Addresses(&_TellorStake.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorStake *TellorStakeCallerSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorStake.Contract.Addresses(&_TellorStake.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorStake *TellorStakeCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "allowance", _user, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorStake *TellorStakeSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _TellorStake.Contract.Allowance(&_TellorStake.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _TellorStake.Contract.Allowance(&_TellorStake.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorStake *TellorStakeCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "allowedToTrade", _user, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorStake *TellorStakeSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _TellorStake.Contract.AllowedToTrade(&_TellorStake.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorStake *TellorStakeCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _TellorStake.Contract.AllowedToTrade(&_TellorStake.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorStake *TellorStakeCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorStake *TellorStakeSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _TellorStake.Contract.BalanceOf(&_TellorStake.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _TellorStake.Contract.BalanceOf(&_TellorStake.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorStake *TellorStakeCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "balanceOfAt", _user, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorStake *TellorStakeSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _TellorStake.Contract.BalanceOfAt(&_TellorStake.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _TellorStake.Contract.BalanceOfAt(&_TellorStake.CallOpts, _user, _blockNumber)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorStake *TellorStakeCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "balances", arg0, arg1)

	outstruct := new(struct {
		FromBlock *big.Int
		Value     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorStake *TellorStakeSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorStake.Contract.Balances(&_TellorStake.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorStake *TellorStakeCallerSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorStake.Contract.Balances(&_TellorStake.CallOpts, arg0, arg1)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorStake *TellorStakeCaller) BytesVars(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "bytesVars", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorStake *TellorStakeSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorStake.Contract.BytesVars(&_TellorStake.CallOpts, arg0)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorStake *TellorStakeCallerSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorStake.Contract.BytesVars(&_TellorStake.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorStake *TellorStakeCaller) CurrentMiners(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "currentMiners", arg0)

	outstruct := new(struct {
		Value *big.Int
		Miner common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Miner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorStake *TellorStakeSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorStake.Contract.CurrentMiners(&_TellorStake.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorStake *TellorStakeCallerSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorStake.Contract.CurrentMiners(&_TellorStake.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeCaller) DisputeIdByDisputeHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "disputeIdByDisputeHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStake.Contract.DisputeIdByDisputeHash(&_TellorStake.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStake.Contract.DisputeIdByDisputeHash(&_TellorStake.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorStake *TellorStakeCaller) DisputesById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "disputesById", arg0)

	outstruct := new(struct {
		Hash                [32]byte
		Tally               *big.Int
		Executed            bool
		DisputeVotePassed   bool
		IsPropFork          bool
		ReportedMiner       common.Address
		ReportingParty      common.Address
		ProposedForkAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Tally = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Executed = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.DisputeVotePassed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.IsPropFork = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.ReportedMiner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.ReportingParty = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.ProposedForkAddress = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorStake *TellorStakeSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorStake.Contract.DisputesById(&_TellorStake.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorStake *TellorStakeCallerSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorStake.Contract.DisputesById(&_TellorStake.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorStake *TellorStakeCaller) Migrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "migrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorStake *TellorStakeSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorStake.Contract.Migrated(&_TellorStake.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorStake *TellorStakeCallerSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorStake.Contract.Migrated(&_TellorStake.CallOpts, arg0)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorStake *TellorStakeCaller) MinersByChallenge(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "minersByChallenge", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorStake *TellorStakeSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorStake.Contract.MinersByChallenge(&_TellorStake.CallOpts, arg0, arg1)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorStake *TellorStakeCallerSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorStake.Contract.MinersByChallenge(&_TellorStake.CallOpts, arg0, arg1)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorStake *TellorStakeCaller) NewValueTimestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "newValueTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorStake *TellorStakeSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorStake.Contract.NewValueTimestamps(&_TellorStake.CallOpts, arg0)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorStake.Contract.NewValueTimestamps(&_TellorStake.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeCaller) RequestIdByQueryHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "requestIdByQueryHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStake.Contract.RequestIdByQueryHash(&_TellorStake.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStake.Contract.RequestIdByQueryHash(&_TellorStake.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeCaller) Uints(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorStake.contract.Call(opts, &out, "uints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorStake.Contract.Uints(&_TellorStake.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorStake *TellorStakeCallerSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorStake.Contract.Uints(&_TellorStake.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_TellorStake *TellorStakeTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_TellorStake *TellorStakeSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.Approve(&_TellorStake.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_TellorStake *TellorStakeTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.Approve(&_TellorStake.TransactOpts, _spender, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_TellorStake *TellorStakeTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _TellorStake.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_TellorStake *TellorStakeSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.BeginDispute(&_TellorStake.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_TellorStake *TellorStakeTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.BeginDispute(&_TellorStake.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_TellorStake *TellorStakeTransactor) ProposeFork(opts *bind.TransactOpts, _propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _TellorStake.contract.Transact(opts, "proposeFork", _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_TellorStake *TellorStakeSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _TellorStake.Contract.ProposeFork(&_TellorStake.TransactOpts, _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_TellorStake *TellorStakeTransactorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _TellorStake.Contract.ProposeFork(&_TellorStake.TransactOpts, _propNewTellorAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_TellorStake *TellorStakeTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_TellorStake *TellorStakeSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.Transfer(&_TellorStake.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_TellorStake *TellorStakeTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.Transfer(&_TellorStake.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_TellorStake *TellorStakeTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_TellorStake *TellorStakeSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.TransferFrom(&_TellorStake.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_TellorStake *TellorStakeTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.TransferFrom(&_TellorStake.TransactOpts, _from, _to, _amount)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_TellorStake *TellorStakeTransactor) UnlockDisputeFee(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _TellorStake.contract.Transact(opts, "unlockDisputeFee", _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_TellorStake *TellorStakeSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.UnlockDisputeFee(&_TellorStake.TransactOpts, _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_TellorStake *TellorStakeTransactorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _TellorStake.Contract.UnlockDisputeFee(&_TellorStake.TransactOpts, _disputeId)
}

// Verify is a paid mutator transaction binding the contract method 0xfc735e99.
//
// Solidity: function verify() returns(uint256)
func (_TellorStake *TellorStakeTransactor) Verify(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorStake.contract.Transact(opts, "verify")
}

// Verify is a paid mutator transaction binding the contract method 0xfc735e99.
//
// Solidity: function verify() returns(uint256)
func (_TellorStake *TellorStakeSession) Verify() (*types.Transaction, error) {
	return _TellorStake.Contract.Verify(&_TellorStake.TransactOpts)
}

// Verify is a paid mutator transaction binding the contract method 0xfc735e99.
//
// Solidity: function verify() returns(uint256)
func (_TellorStake *TellorStakeTransactorSession) Verify() (*types.Transaction, error) {
	return _TellorStake.Contract.Verify(&_TellorStake.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_TellorStake *TellorStakeTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _TellorStake.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_TellorStake *TellorStakeSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _TellorStake.Contract.Vote(&_TellorStake.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_TellorStake *TellorStakeTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _TellorStake.Contract.Vote(&_TellorStake.TransactOpts, _disputeId, _supportsDispute)
}

// TellorStakeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TellorStake contract.
type TellorStakeApprovalIterator struct {
	Event *TellorStakeApproval // Event containing the contract specifics and raw log

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
func (it *TellorStakeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorStakeApproval)
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
		it.Event = new(TellorStakeApproval)
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
func (it *TellorStakeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorStakeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorStakeApproval represents a Approval event raised by the TellorStake contract.
type TellorStakeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_TellorStake *TellorStakeFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*TellorStakeApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _TellorStake.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &TellorStakeApprovalIterator{contract: _TellorStake.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_TellorStake *TellorStakeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TellorStakeApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _TellorStake.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorStakeApproval)
				if err := _TellorStake.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_TellorStake *TellorStakeFilterer) ParseApproval(log types.Log) (*TellorStakeApproval, error) {
	event := new(TellorStakeApproval)
	if err := _TellorStake.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorStakeNewDisputeIterator is returned from FilterNewDispute and is used to iterate over the raw logs and unpacked data for NewDispute events raised by the TellorStake contract.
type TellorStakeNewDisputeIterator struct {
	Event *TellorStakeNewDispute // Event containing the contract specifics and raw log

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
func (it *TellorStakeNewDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorStakeNewDispute)
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
		it.Event = new(TellorStakeNewDispute)
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
func (it *TellorStakeNewDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorStakeNewDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorStakeNewDispute represents a NewDispute event raised by the TellorStake contract.
type TellorStakeNewDispute struct {
	DisputeId *big.Int
	RequestId *big.Int
	Timestamp *big.Int
	Miner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewDispute is a free log retrieval operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_TellorStake *TellorStakeFilterer) FilterNewDispute(opts *bind.FilterOpts, _disputeId []*big.Int, _requestId []*big.Int) (*TellorStakeNewDisputeIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorStake.contract.FilterLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &TellorStakeNewDisputeIterator{contract: _TellorStake.contract, event: "NewDispute", logs: logs, sub: sub}, nil
}

// WatchNewDispute is a free log subscription operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_TellorStake *TellorStakeFilterer) WatchNewDispute(opts *bind.WatchOpts, sink chan<- *TellorStakeNewDispute, _disputeId []*big.Int, _requestId []*big.Int) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorStake.contract.WatchLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorStakeNewDispute)
				if err := _TellorStake.contract.UnpackLog(event, "NewDispute", log); err != nil {
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

// ParseNewDispute is a log parse operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: event NewDispute(uint256 indexed _disputeId, uint256 indexed _requestId, uint256 _timestamp, address _miner)
func (_TellorStake *TellorStakeFilterer) ParseNewDispute(log types.Log) (*TellorStakeNewDispute, error) {
	event := new(TellorStakeNewDispute)
	if err := _TellorStake.contract.UnpackLog(event, "NewDispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorStakeTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the TellorStake contract.
type TellorStakeTransferredIterator struct {
	Event *TellorStakeTransferred // Event containing the contract specifics and raw log

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
func (it *TellorStakeTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorStakeTransferred)
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
		it.Event = new(TellorStakeTransferred)
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
func (it *TellorStakeTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorStakeTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorStakeTransferred represents a Transferred event raised by the TellorStake contract.
type TellorStakeTransferred struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_TellorStake *TellorStakeFilterer) FilterTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*TellorStakeTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _TellorStake.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TellorStakeTransferredIterator{contract: _TellorStake.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_TellorStake *TellorStakeFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *TellorStakeTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _TellorStake.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorStakeTransferred)
				if err := _TellorStake.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransferred is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_TellorStake *TellorStakeFilterer) ParseTransferred(log types.Log) (*TellorStakeTransferred, error) {
	event := new(TellorStakeTransferred)
	if err := _TellorStake.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorStakeVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the TellorStake contract.
type TellorStakeVotedIterator struct {
	Event *TellorStakeVoted // Event containing the contract specifics and raw log

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
func (it *TellorStakeVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorStakeVoted)
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
		it.Event = new(TellorStakeVoted)
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
func (it *TellorStakeVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorStakeVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorStakeVoted represents a Voted event raised by the TellorStake contract.
type TellorStakeVoted struct {
	DisputeID  *big.Int
	Position   bool
	Voter      common.Address
	VoteWeight *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_TellorStake *TellorStakeFilterer) FilterVoted(opts *bind.FilterOpts, _disputeID []*big.Int, _voter []common.Address, _voteWeight []*big.Int) (*TellorStakeVotedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}
	var _voteWeightRule []interface{}
	for _, _voteWeightItem := range _voteWeight {
		_voteWeightRule = append(_voteWeightRule, _voteWeightItem)
	}

	logs, sub, err := _TellorStake.contract.FilterLogs(opts, "Voted", _disputeIDRule, _voterRule, _voteWeightRule)
	if err != nil {
		return nil, err
	}
	return &TellorStakeVotedIterator{contract: _TellorStake.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_TellorStake *TellorStakeFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *TellorStakeVoted, _disputeID []*big.Int, _voter []common.Address, _voteWeight []*big.Int) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}
	var _voteWeightRule []interface{}
	for _, _voteWeightItem := range _voteWeight {
		_voteWeightRule = append(_voteWeightRule, _voteWeightItem)
	}

	logs, sub, err := _TellorStake.contract.WatchLogs(opts, "Voted", _disputeIDRule, _voterRule, _voteWeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorStakeVoted)
				if err := _TellorStake.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x911ef2e98570b1d96c0e8ef81107a33d5b8e844aeb8f9710f9bc76c3b3fef40e.
//
// Solidity: event Voted(uint256 indexed _disputeID, bool _position, address indexed _voter, uint256 indexed _voteWeight)
func (_TellorStake *TellorStakeFilterer) ParseVoted(log types.Log) (*TellorStakeVoted, error) {
	event := new(TellorStakeVoted)
	if err := _TellorStake.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorStorageMetaData contains all meta data concerning the TellorStorage contract.
var TellorStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"fromBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentMiners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"disputesById\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"tally\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disputeVotePassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPropFork\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"reportedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reportingParty\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedForkAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minersByChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newValueTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"024c2ddd": "_allowances(address,address)",
		"699f200f": "addresses(bytes32)",
		"cbf1304d": "balances(address,uint256)",
		"62dd1d2a": "bytesVars(bytes32)",
		"1fd22364": "currentMiners(uint256)",
		"d01f4d9e": "disputeIdByDisputeHash(bytes32)",
		"db085beb": "disputesById(uint256)",
		"4ba0a5ee": "migrated(address)",
		"48b18e54": "minersByChallenge(bytes32,address)",
		"438c0aa3": "newValueTimestamps(uint256)",
		"5700242c": "requestIdByQueryHash(bytes32)",
		"b59e14d4": "uints(bytes32)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610505806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806362dd1d2a1161007157806362dd1d2a146101d7578063699f200f146101f4578063b59e14d41461022d578063cbf1304d1461024a578063d01f4d9e146102a5578063db085beb146102c2576100b4565b8063024c2ddd146100b95780631fd22364146100f9578063438c0aa31461013757806348b18e54146101545780634ba0a5ee146101945780635700242c146101ba575b600080fd5b6100e7600480360360408110156100cf57600080fd5b506001600160a01b038135811691602001351661032f565b60408051918252519081900360200190f35b6101166004803603602081101561010f57600080fd5b503561034c565b604080519283526001600160a01b0390911660208301528051918290030190f35b6100e76004803603602081101561014d57600080fd5b5035610377565b6101806004803603604081101561016a57600080fd5b50803590602001356001600160a01b0316610398565b604080519115158252519081900360200190f35b610180600480360360208110156101aa57600080fd5b50356001600160a01b03166103b8565b6100e7600480360360208110156101d057600080fd5b50356103cd565b6100e7600480360360208110156101ed57600080fd5b50356103df565b6102116004803603602081101561020a57600080fd5b50356103f1565b604080516001600160a01b039092168252519081900360200190f35b6100e76004803603602081101561024357600080fd5b503561040c565b6102766004803603604081101561026057600080fd5b506001600160a01b03813516906020013561041e565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b6100e7600480360360208110156102bb57600080fd5b5035610461565b6102df600480360360208110156102d857600080fd5b5035610473565b60408051988952602089019790975294151587870152921515606087015290151560808601526001600160a01b0390811660a086015290811660c08501521660e083015251908190036101000190f35b604a60209081526000928352604080842090915290825290205481565b603a816005811061035c57600080fd5b6002020180546001909101549091506001600160a01b031682565b6033818154811061038757600080fd5b600091825260209091200154905081565b603960209081526000928352604080842090915290825290205460ff1681565b604b6020526000908152604090205460ff1681565b60376020526000908152604090205481565b60486020526000908152604090205481565b6047602052600090815260409020546001600160a01b031681565b60466020526000908152604090205481565b6049602052816000526040600020818154811061043a57600080fd5b6000918252602090912001546001600160801b038082169350600160801b90910416905082565b60386020526000908152604090205481565b603660205260009081526040902080546001820154600283015460038401546004909401549293919260ff808316936101008404821693620100008104909216926001600160a01b03630100000090930483169291821691168856fea26469706673582212200a686b53ea501bbaf6d1d8c5f295fd35e767b4c8b5f68f67a84e56bb4944c91964736f6c63430007040033",
}

// TellorStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use TellorStorageMetaData.ABI instead.
var TellorStorageABI = TellorStorageMetaData.ABI

// Deprecated: Use TellorStorageMetaData.Sigs instead.
// TellorStorageFuncSigs maps the 4-byte function signature to its string representation.
var TellorStorageFuncSigs = TellorStorageMetaData.Sigs

// TellorStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TellorStorageMetaData.Bin instead.
var TellorStorageBin = TellorStorageMetaData.Bin

// DeployTellorStorage deploys a new Ethereum contract, binding an instance of TellorStorage to it.
func DeployTellorStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorStorage, error) {
	parsed, err := TellorStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TellorStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorStorage{TellorStorageCaller: TellorStorageCaller{contract: contract}, TellorStorageTransactor: TellorStorageTransactor{contract: contract}, TellorStorageFilterer: TellorStorageFilterer{contract: contract}}, nil
}

// TellorStorage is an auto generated Go binding around an Ethereum contract.
type TellorStorage struct {
	TellorStorageCaller     // Read-only binding to the contract
	TellorStorageTransactor // Write-only binding to the contract
	TellorStorageFilterer   // Log filterer for contract events
}

// TellorStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorStorageSession struct {
	Contract     *TellorStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorStorageCallerSession struct {
	Contract *TellorStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TellorStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorStorageTransactorSession struct {
	Contract     *TellorStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TellorStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorStorageRaw struct {
	Contract *TellorStorage // Generic contract binding to access the raw methods on
}

// TellorStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorStorageCallerRaw struct {
	Contract *TellorStorageCaller // Generic read-only contract binding to access the raw methods on
}

// TellorStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorStorageTransactorRaw struct {
	Contract *TellorStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorStorage creates a new instance of TellorStorage, bound to a specific deployed contract.
func NewTellorStorage(address common.Address, backend bind.ContractBackend) (*TellorStorage, error) {
	contract, err := bindTellorStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorStorage{TellorStorageCaller: TellorStorageCaller{contract: contract}, TellorStorageTransactor: TellorStorageTransactor{contract: contract}, TellorStorageFilterer: TellorStorageFilterer{contract: contract}}, nil
}

// NewTellorStorageCaller creates a new read-only instance of TellorStorage, bound to a specific deployed contract.
func NewTellorStorageCaller(address common.Address, caller bind.ContractCaller) (*TellorStorageCaller, error) {
	contract, err := bindTellorStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorStorageCaller{contract: contract}, nil
}

// NewTellorStorageTransactor creates a new write-only instance of TellorStorage, bound to a specific deployed contract.
func NewTellorStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorStorageTransactor, error) {
	contract, err := bindTellorStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorStorageTransactor{contract: contract}, nil
}

// NewTellorStorageFilterer creates a new log filterer instance of TellorStorage, bound to a specific deployed contract.
func NewTellorStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorStorageFilterer, error) {
	contract, err := bindTellorStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorStorageFilterer{contract: contract}, nil
}

// bindTellorStorage binds a generic wrapper to an already deployed contract.
func bindTellorStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorStorage *TellorStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorStorage.Contract.TellorStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorStorage *TellorStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorStorage.Contract.TellorStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorStorage *TellorStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorStorage.Contract.TellorStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorStorage *TellorStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorStorage *TellorStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorStorage *TellorStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorStorage.Contract.contract.Transact(opts, method, params...)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorStorage *TellorStorageCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "_allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorStorage *TellorStorageSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorStorage.Contract.Allowances(&_TellorStorage.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorStorage *TellorStorageCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorStorage.Contract.Allowances(&_TellorStorage.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorStorage *TellorStorageCaller) Addresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorStorage *TellorStorageSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorStorage.Contract.Addresses(&_TellorStorage.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorStorage *TellorStorageCallerSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorStorage.Contract.Addresses(&_TellorStorage.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorStorage *TellorStorageCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "balances", arg0, arg1)

	outstruct := new(struct {
		FromBlock *big.Int
		Value     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorStorage *TellorStorageSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorStorage.Contract.Balances(&_TellorStorage.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorStorage *TellorStorageCallerSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorStorage.Contract.Balances(&_TellorStorage.CallOpts, arg0, arg1)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorStorage *TellorStorageCaller) BytesVars(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "bytesVars", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorStorage *TellorStorageSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorStorage.Contract.BytesVars(&_TellorStorage.CallOpts, arg0)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorStorage *TellorStorageCallerSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorStorage.Contract.BytesVars(&_TellorStorage.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorStorage *TellorStorageCaller) CurrentMiners(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "currentMiners", arg0)

	outstruct := new(struct {
		Value *big.Int
		Miner common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Miner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorStorage *TellorStorageSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorStorage.Contract.CurrentMiners(&_TellorStorage.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorStorage *TellorStorageCallerSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorStorage.Contract.CurrentMiners(&_TellorStorage.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageCaller) DisputeIdByDisputeHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "disputeIdByDisputeHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStorage.Contract.DisputeIdByDisputeHash(&_TellorStorage.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageCallerSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStorage.Contract.DisputeIdByDisputeHash(&_TellorStorage.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorStorage *TellorStorageCaller) DisputesById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "disputesById", arg0)

	outstruct := new(struct {
		Hash                [32]byte
		Tally               *big.Int
		Executed            bool
		DisputeVotePassed   bool
		IsPropFork          bool
		ReportedMiner       common.Address
		ReportingParty      common.Address
		ProposedForkAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Tally = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Executed = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.DisputeVotePassed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.IsPropFork = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.ReportedMiner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.ReportingParty = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.ProposedForkAddress = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorStorage *TellorStorageSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorStorage.Contract.DisputesById(&_TellorStorage.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorStorage *TellorStorageCallerSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorStorage.Contract.DisputesById(&_TellorStorage.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorStorage *TellorStorageCaller) Migrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "migrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorStorage *TellorStorageSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorStorage.Contract.Migrated(&_TellorStorage.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorStorage *TellorStorageCallerSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorStorage.Contract.Migrated(&_TellorStorage.CallOpts, arg0)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorStorage *TellorStorageCaller) MinersByChallenge(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "minersByChallenge", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorStorage *TellorStorageSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorStorage.Contract.MinersByChallenge(&_TellorStorage.CallOpts, arg0, arg1)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorStorage *TellorStorageCallerSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorStorage.Contract.MinersByChallenge(&_TellorStorage.CallOpts, arg0, arg1)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorStorage *TellorStorageCaller) NewValueTimestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "newValueTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorStorage *TellorStorageSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorStorage.Contract.NewValueTimestamps(&_TellorStorage.CallOpts, arg0)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorStorage *TellorStorageCallerSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorStorage.Contract.NewValueTimestamps(&_TellorStorage.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageCaller) RequestIdByQueryHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "requestIdByQueryHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStorage.Contract.RequestIdByQueryHash(&_TellorStorage.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageCallerSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorStorage.Contract.RequestIdByQueryHash(&_TellorStorage.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageCaller) Uints(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorStorage.contract.Call(opts, &out, "uints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorStorage.Contract.Uints(&_TellorStorage.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorStorage *TellorStorageCallerSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorStorage.Contract.Uints(&_TellorStorage.CallOpts, arg0)
}

// TellorTransferMetaData contains all meta data concerning the TellorTransfer contract.
var TellorTransferMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"fromBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bytesVars\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentMiners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputeIdByDisputeHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"disputesById\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"tally\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disputeVotePassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPropFork\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"reportedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reportingParty\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedForkAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minersByChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"newValueTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdByQueryHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"uints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"024c2ddd": "_allowances(address,address)",
		"699f200f": "addresses(bytes32)",
		"dd62ed3e": "allowance(address,address)",
		"999cf26c": "allowedToTrade(address,uint256)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"4ee2cd7e": "balanceOfAt(address,uint256)",
		"cbf1304d": "balances(address,uint256)",
		"62dd1d2a": "bytesVars(bytes32)",
		"1fd22364": "currentMiners(uint256)",
		"d01f4d9e": "disputeIdByDisputeHash(bytes32)",
		"db085beb": "disputesById(uint256)",
		"4ba0a5ee": "migrated(address)",
		"48b18e54": "minersByChallenge(bytes32,address)",
		"438c0aa3": "newValueTimestamps(uint256)",
		"5700242c": "requestIdByQueryHash(bytes32)",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"b59e14d4": "uints(bytes32)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610df1806100206000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806362dd1d2a116100ad578063b59e14d411610071578063b59e14d4146103a6578063cbf1304d146103c3578063d01f4d9e1461041e578063db085beb1461043b578063dd62ed3e146104a857610121565b806362dd1d2a146102d2578063699f200f146102ef57806370a0823114610328578063999cf26c1461034e578063a9059cbb1461037a57610121565b8063438c0aa3116100f4578063438c0aa31461021a57806348b18e54146102375780634ba0a5ee146102635780634ee2cd7e146102895780635700242c146102b557610121565b8063024c2ddd14610126578063095ea7b3146101665780631fd22364146101a657806323b872dd146101e4575b600080fd5b6101546004803603604081101561013c57600080fd5b506001600160a01b03813581169160200135166104d6565b60408051918252519081900360200190f35b6101926004803603604081101561017c57600080fd5b506001600160a01b0381351690602001356104f3565b604080519115158252519081900360200190f35b6101c3600480360360208110156101bc57600080fd5b50356105a2565b604080519283526001600160a01b0390911660208301528051918290030190f35b610192600480360360608110156101fa57600080fd5b506001600160a01b038135811691602081013590911690604001356105cd565b6101546004803603602081101561023057600080fd5b5035610679565b6101926004803603604081101561024d57600080fd5b50803590602001356001600160a01b031661069a565b6101926004803603602081101561027957600080fd5b50356001600160a01b03166106ba565b6101546004803603604081101561029f57600080fd5b506001600160a01b0381351690602001356106cf565b610154600480360360208110156102cb57600080fd5b5035610873565b610154600480360360208110156102e857600080fd5b5035610885565b61030c6004803603602081101561030557600080fd5b5035610897565b604080516001600160a01b039092168252519081900360200190f35b6101546004803603602081101561033e57600080fd5b50356001600160a01b03166108b2565b6101926004803603604081101561036457600080fd5b506001600160a01b0381351690602001356108be565b6101926004803603604081101561039057600080fd5b506001600160a01b03813516906020013561098a565b610154600480360360208110156103bc57600080fd5b50356109a0565b6103ef600480360360408110156103d957600080fd5b506001600160a01b0381351690602001356109b2565b60405180836001600160801b03168152602001826001600160801b031681526020019250505060405180910390f35b6101546004803603602081101561043457600080fd5b50356109f5565b6104586004803603602081101561045157600080fd5b5035610a07565b60408051988952602089019790975294151587870152921515606087015290151560808601526001600160a01b0390811660a086015290811660c08501521660e083015251908190036101000190f35b610154600480360360408110156104be57600080fd5b506001600160a01b0381358116916020013516610a63565b604a60209081526000928352604080842090915290825290205481565b60006001600160a01b03831661053a5760405162461bcd60e51b8152600401808060200182810382526022815260200180610d526022913960400191505060405180910390fd5b336000818152604a602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b603a81600581106105b257600080fd5b6002020180546001909101549091506001600160a01b031682565b6001600160a01b0383166000908152604a6020908152604080832033845290915281205482111561063a576040805162461bcd60e51b8152602060048201526012602482015271416c6c6f77616e63652069732077726f6e6760701b604482015290519081900360640190fd5b6001600160a01b0384166000908152604a6020908152604080832033845290915290208054839003905561066f848484610a8e565b5060019392505050565b6033818154811061068957600080fd5b600091825260209091200154905081565b603960209081526000928352604080842090915290825290205460ff1681565b604b6020526000908152604090205460ff1681565b6001600160a01b038216600090815260496020526040812080541580610715575082816000815481106106fe57fe5b6000918252602090912001546001600160801b0316115b1561072457600091505061059c565b80548190600019810190811061073657fe5b6000918252602090912001546001600160801b031683106107885780548190600019810190811061076357fe5b600091825260209091200154600160801b90046001600160801b0316915061059c9050565b8054600090600119015b818111156108405760006002600183850101049050858482815481106107b457fe5b6000918252602090912001546001600160801b03161415610803578381815481106107db57fe5b600091825260209091200154600160801b90046001600160801b0316945061059c9350505050565b8584828154811061081057fe5b6000918252602090912001546001600160801b031610156108335780925061083a565b6001810391505b50610792565b82828154811061084c57fe5b600091825260209091200154600160801b90046001600160801b0316935061059c92505050565b60376020526000908152604090205481565b60486020526000908152604090205481565b6047602052600090815260409020546001600160a01b031681565b600061059c82436106cf565b6001600160a01b038216600090815260446020526040812054158015906108fd57506001600160a01b0383166000908152604460205260409020546005115b15610977577f5d9fadfc729fd027e395e5157ef1b53ef9fa4a8f053043c5f159307543e7cc9760005260466020527f167af83a0768d27540775cfef6d996eb63f8a61fcdfb26e654c18fb50960e3be5482906109629061095c866108b2565b90610c42565b1061096f5750600161059c565b50600061059c565b81610981846108b2565b10159392505050565b6000610997338484610a8e565b50600192915050565b60466020526000908152604090205481565b604960205281600052604060002081815481106109ce57600080fd5b6000918252602090912001546001600160801b038082169350600160801b90910416905082565b60386020526000908152604090205481565b603660205260009081526040902080546001820154600283015460038401546004909401549293919260ff808316936101008404821693620100008104909216926001600160a01b036301000000909304831692918216911688565b6001600160a01b039182166000908152604a6020908152604080832093909416825291909152205490565b80610aca5760405162461bcd60e51b8152600401808060200182810382526021815260200180610d9b6021913960400191505060405180910390fd5b6001600160a01b038216610b1d576040805162461bcd60e51b815260206004820152601560248201527452656365697665722069732030206164647265737360581b604482015290519081900360640190fd5b610b2783826108be565b610b625760405162461bcd60e51b8152600401808060200182810382526027815260200180610d746027913960400191505060405180910390fd5b6000610b6d846108b2565b905081610b7c85828403610c54565b610b85846108b2565b9150816001600160801b03168183016001600160801b03161015610be4576040805162461bcd60e51b815260206004820152601160248201527013dd995c999b1bddc81a185c1c195b9959607a1b604482015290519081900360640190fd5b610bf084828401610c54565b836001600160a01b0316856001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a35050505050565b600082821115610c4e57fe5b50900390565b6001600160a01b038216600090815260496020526040902080541580610ca157508054439082906000198101908110610c8957fe5b6000918252602090912001546001600160801b031614155b15610d125760408051808201909152436001600160801b0390811682528381166020808401918252845460018101865560008681529190912093519301805491516fffffffffffffffffffffffffffffffff19909216938316939093178216600160801b9190921602179055610d4c565b805460009082906000198101908110610d2757fe5b600091825260209091200180546001600160801b03808616600160801b029116179055505b50505056fe45524332303a20617070726f766520746f20746865207a65726f206164647265737353686f756c6420686176652073756666696369656e742062616c616e636520746f207472616465547269656420746f2073656e64206e6f6e2d706f73697469766520616d6f756e74a2646970667358221220058774e0adf3697ed569a90f0bb7e8e8a9b13a626f2f06ebf7764b0d31114add64736f6c63430007040033",
}

// TellorTransferABI is the input ABI used to generate the binding from.
// Deprecated: Use TellorTransferMetaData.ABI instead.
var TellorTransferABI = TellorTransferMetaData.ABI

// Deprecated: Use TellorTransferMetaData.Sigs instead.
// TellorTransferFuncSigs maps the 4-byte function signature to its string representation.
var TellorTransferFuncSigs = TellorTransferMetaData.Sigs

// TellorTransferBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TellorTransferMetaData.Bin instead.
var TellorTransferBin = TellorTransferMetaData.Bin

// DeployTellorTransfer deploys a new Ethereum contract, binding an instance of TellorTransfer to it.
func DeployTellorTransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TellorTransfer, error) {
	parsed, err := TellorTransferMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TellorTransferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TellorTransfer{TellorTransferCaller: TellorTransferCaller{contract: contract}, TellorTransferTransactor: TellorTransferTransactor{contract: contract}, TellorTransferFilterer: TellorTransferFilterer{contract: contract}}, nil
}

// TellorTransfer is an auto generated Go binding around an Ethereum contract.
type TellorTransfer struct {
	TellorTransferCaller     // Read-only binding to the contract
	TellorTransferTransactor // Write-only binding to the contract
	TellorTransferFilterer   // Log filterer for contract events
}

// TellorTransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorTransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorTransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorTransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorTransferSession struct {
	Contract     *TellorTransfer   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorTransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorTransferCallerSession struct {
	Contract *TellorTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TellorTransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorTransferTransactorSession struct {
	Contract     *TellorTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TellorTransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorTransferRaw struct {
	Contract *TellorTransfer // Generic contract binding to access the raw methods on
}

// TellorTransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorTransferCallerRaw struct {
	Contract *TellorTransferCaller // Generic read-only contract binding to access the raw methods on
}

// TellorTransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorTransferTransactorRaw struct {
	Contract *TellorTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorTransfer creates a new instance of TellorTransfer, bound to a specific deployed contract.
func NewTellorTransfer(address common.Address, backend bind.ContractBackend) (*TellorTransfer, error) {
	contract, err := bindTellorTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorTransfer{TellorTransferCaller: TellorTransferCaller{contract: contract}, TellorTransferTransactor: TellorTransferTransactor{contract: contract}, TellorTransferFilterer: TellorTransferFilterer{contract: contract}}, nil
}

// NewTellorTransferCaller creates a new read-only instance of TellorTransfer, bound to a specific deployed contract.
func NewTellorTransferCaller(address common.Address, caller bind.ContractCaller) (*TellorTransferCaller, error) {
	contract, err := bindTellorTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorTransferCaller{contract: contract}, nil
}

// NewTellorTransferTransactor creates a new write-only instance of TellorTransfer, bound to a specific deployed contract.
func NewTellorTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorTransferTransactor, error) {
	contract, err := bindTellorTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorTransferTransactor{contract: contract}, nil
}

// NewTellorTransferFilterer creates a new log filterer instance of TellorTransfer, bound to a specific deployed contract.
func NewTellorTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorTransferFilterer, error) {
	contract, err := bindTellorTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorTransferFilterer{contract: contract}, nil
}

// bindTellorTransfer binds a generic wrapper to an already deployed contract.
func bindTellorTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorTransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorTransfer *TellorTransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorTransfer.Contract.TellorTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorTransfer *TellorTransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorTransfer.Contract.TellorTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorTransfer *TellorTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorTransfer.Contract.TellorTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorTransfer *TellorTransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorTransfer *TellorTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorTransfer *TellorTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorTransfer.Contract.contract.Transact(opts, method, params...)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "_allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorTransfer.Contract.Allowances(&_TellorTransfer.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x024c2ddd.
//
// Solidity: function _allowances(address , address ) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TellorTransfer.Contract.Allowances(&_TellorTransfer.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorTransfer *TellorTransferCaller) Addresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorTransfer *TellorTransferSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorTransfer.Contract.Addresses(&_TellorTransfer.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_TellorTransfer *TellorTransferCallerSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _TellorTransfer.Contract.Addresses(&_TellorTransfer.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "allowance", _user, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _TellorTransfer.Contract.Allowance(&_TellorTransfer.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _user, address _spender) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _TellorTransfer.Contract.Allowance(&_TellorTransfer.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorTransfer *TellorTransferCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "allowedToTrade", _user, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorTransfer *TellorTransferSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _TellorTransfer.Contract.AllowedToTrade(&_TellorTransfer.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(address _user, uint256 _amount) view returns(bool)
func (_TellorTransfer *TellorTransferCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _TellorTransfer.Contract.AllowedToTrade(&_TellorTransfer.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _TellorTransfer.Contract.BalanceOf(&_TellorTransfer.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _TellorTransfer.Contract.BalanceOf(&_TellorTransfer.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "balanceOfAt", _user, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _TellorTransfer.Contract.BalanceOfAt(&_TellorTransfer.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address _user, uint256 _blockNumber) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _TellorTransfer.Contract.BalanceOfAt(&_TellorTransfer.CallOpts, _user, _blockNumber)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorTransfer *TellorTransferCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "balances", arg0, arg1)

	outstruct := new(struct {
		FromBlock *big.Int
		Value     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorTransfer *TellorTransferSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorTransfer.Contract.Balances(&_TellorTransfer.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xcbf1304d.
//
// Solidity: function balances(address , uint256 ) view returns(uint128 fromBlock, uint128 value)
func (_TellorTransfer *TellorTransferCallerSession) Balances(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock *big.Int
	Value     *big.Int
}, error) {
	return _TellorTransfer.Contract.Balances(&_TellorTransfer.CallOpts, arg0, arg1)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorTransfer *TellorTransferCaller) BytesVars(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "bytesVars", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorTransfer *TellorTransferSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorTransfer.Contract.BytesVars(&_TellorTransfer.CallOpts, arg0)
}

// BytesVars is a free data retrieval call binding the contract method 0x62dd1d2a.
//
// Solidity: function bytesVars(bytes32 ) view returns(bytes32)
func (_TellorTransfer *TellorTransferCallerSession) BytesVars(arg0 [32]byte) ([32]byte, error) {
	return _TellorTransfer.Contract.BytesVars(&_TellorTransfer.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorTransfer *TellorTransferCaller) CurrentMiners(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "currentMiners", arg0)

	outstruct := new(struct {
		Value *big.Int
		Miner common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Miner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorTransfer *TellorTransferSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorTransfer.Contract.CurrentMiners(&_TellorTransfer.CallOpts, arg0)
}

// CurrentMiners is a free data retrieval call binding the contract method 0x1fd22364.
//
// Solidity: function currentMiners(uint256 ) view returns(uint256 value, address miner)
func (_TellorTransfer *TellorTransferCallerSession) CurrentMiners(arg0 *big.Int) (struct {
	Value *big.Int
	Miner common.Address
}, error) {
	return _TellorTransfer.Contract.CurrentMiners(&_TellorTransfer.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) DisputeIdByDisputeHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "disputeIdByDisputeHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorTransfer.Contract.DisputeIdByDisputeHash(&_TellorTransfer.CallOpts, arg0)
}

// DisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xd01f4d9e.
//
// Solidity: function disputeIdByDisputeHash(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) DisputeIdByDisputeHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorTransfer.Contract.DisputeIdByDisputeHash(&_TellorTransfer.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorTransfer *TellorTransferCaller) DisputesById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "disputesById", arg0)

	outstruct := new(struct {
		Hash                [32]byte
		Tally               *big.Int
		Executed            bool
		DisputeVotePassed   bool
		IsPropFork          bool
		ReportedMiner       common.Address
		ReportingParty      common.Address
		ProposedForkAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Tally = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Executed = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.DisputeVotePassed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.IsPropFork = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.ReportedMiner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.ReportingParty = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.ProposedForkAddress = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorTransfer *TellorTransferSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorTransfer.Contract.DisputesById(&_TellorTransfer.CallOpts, arg0)
}

// DisputesById is a free data retrieval call binding the contract method 0xdb085beb.
//
// Solidity: function disputesById(uint256 ) view returns(bytes32 hash, int256 tally, bool executed, bool disputeVotePassed, bool isPropFork, address reportedMiner, address reportingParty, address proposedForkAddress)
func (_TellorTransfer *TellorTransferCallerSession) DisputesById(arg0 *big.Int) (struct {
	Hash                [32]byte
	Tally               *big.Int
	Executed            bool
	DisputeVotePassed   bool
	IsPropFork          bool
	ReportedMiner       common.Address
	ReportingParty      common.Address
	ProposedForkAddress common.Address
}, error) {
	return _TellorTransfer.Contract.DisputesById(&_TellorTransfer.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorTransfer *TellorTransferCaller) Migrated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "migrated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorTransfer *TellorTransferSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorTransfer.Contract.Migrated(&_TellorTransfer.CallOpts, arg0)
}

// Migrated is a free data retrieval call binding the contract method 0x4ba0a5ee.
//
// Solidity: function migrated(address ) view returns(bool)
func (_TellorTransfer *TellorTransferCallerSession) Migrated(arg0 common.Address) (bool, error) {
	return _TellorTransfer.Contract.Migrated(&_TellorTransfer.CallOpts, arg0)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorTransfer *TellorTransferCaller) MinersByChallenge(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "minersByChallenge", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorTransfer *TellorTransferSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorTransfer.Contract.MinersByChallenge(&_TellorTransfer.CallOpts, arg0, arg1)
}

// MinersByChallenge is a free data retrieval call binding the contract method 0x48b18e54.
//
// Solidity: function minersByChallenge(bytes32 , address ) view returns(bool)
func (_TellorTransfer *TellorTransferCallerSession) MinersByChallenge(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _TellorTransfer.Contract.MinersByChallenge(&_TellorTransfer.CallOpts, arg0, arg1)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) NewValueTimestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "newValueTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorTransfer.Contract.NewValueTimestamps(&_TellorTransfer.CallOpts, arg0)
}

// NewValueTimestamps is a free data retrieval call binding the contract method 0x438c0aa3.
//
// Solidity: function newValueTimestamps(uint256 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) NewValueTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _TellorTransfer.Contract.NewValueTimestamps(&_TellorTransfer.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) RequestIdByQueryHash(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "requestIdByQueryHash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorTransfer.Contract.RequestIdByQueryHash(&_TellorTransfer.CallOpts, arg0)
}

// RequestIdByQueryHash is a free data retrieval call binding the contract method 0x5700242c.
//
// Solidity: function requestIdByQueryHash(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) RequestIdByQueryHash(arg0 [32]byte) (*big.Int, error) {
	return _TellorTransfer.Contract.RequestIdByQueryHash(&_TellorTransfer.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCaller) Uints(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TellorTransfer.contract.Call(opts, &out, "uints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorTransfer.Contract.Uints(&_TellorTransfer.CallOpts, arg0)
}

// Uints is a free data retrieval call binding the contract method 0xb59e14d4.
//
// Solidity: function uints(bytes32 ) view returns(uint256)
func (_TellorTransfer *TellorTransferCallerSession) Uints(arg0 [32]byte) (*big.Int, error) {
	return _TellorTransfer.Contract.Uints(&_TellorTransfer.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_TellorTransfer *TellorTransferTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_TellorTransfer *TellorTransferSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.Contract.Approve(&_TellorTransfer.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_TellorTransfer *TellorTransferTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.Contract.Approve(&_TellorTransfer.TransactOpts, _spender, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_TellorTransfer *TellorTransferTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_TellorTransfer *TellorTransferSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.Contract.Transfer(&_TellorTransfer.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_TellorTransfer *TellorTransferTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.Contract.Transfer(&_TellorTransfer.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_TellorTransfer *TellorTransferTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_TellorTransfer *TellorTransferSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.Contract.TransferFrom(&_TellorTransfer.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_TellorTransfer *TellorTransferTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TellorTransfer.Contract.TransferFrom(&_TellorTransfer.TransactOpts, _from, _to, _amount)
}

// TellorTransferApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TellorTransfer contract.
type TellorTransferApprovalIterator struct {
	Event *TellorTransferApproval // Event containing the contract specifics and raw log

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
func (it *TellorTransferApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorTransferApproval)
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
		it.Event = new(TellorTransferApproval)
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
func (it *TellorTransferApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorTransferApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorTransferApproval represents a Approval event raised by the TellorTransfer contract.
type TellorTransferApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_TellorTransfer *TellorTransferFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*TellorTransferApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _TellorTransfer.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &TellorTransferApprovalIterator{contract: _TellorTransfer.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_TellorTransfer *TellorTransferFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TellorTransferApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _TellorTransfer.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorTransferApproval)
				if err := _TellorTransfer.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_TellorTransfer *TellorTransferFilterer) ParseApproval(log types.Log) (*TellorTransferApproval, error) {
	event := new(TellorTransferApproval)
	if err := _TellorTransfer.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorTransferTransferredIterator is returned from FilterTransferred and is used to iterate over the raw logs and unpacked data for Transferred events raised by the TellorTransfer contract.
type TellorTransferTransferredIterator struct {
	Event *TellorTransferTransferred // Event containing the contract specifics and raw log

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
func (it *TellorTransferTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorTransferTransferred)
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
		it.Event = new(TellorTransferTransferred)
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
func (it *TellorTransferTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorTransferTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorTransferTransferred represents a Transferred event raised by the TellorTransfer contract.
type TellorTransferTransferred struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferred is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_TellorTransfer *TellorTransferFilterer) FilterTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*TellorTransferTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _TellorTransfer.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TellorTransferTransferredIterator{contract: _TellorTransfer.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransferred is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_TellorTransfer *TellorTransferFilterer) WatchTransferred(opts *bind.WatchOpts, sink chan<- *TellorTransferTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _TellorTransfer.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorTransferTransferred)
				if err := _TellorTransfer.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransferred is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_TellorTransfer *TellorTransferFilterer) ParseTransferred(log types.Log) (*TellorTransferTransferred, error) {
	event := new(TellorTransferTransferred)
	if err := _TellorTransfer.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorVariablesMetaData contains all meta data concerning the TellorVariables contract.
var TellorVariablesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220645ce8ea99bd96538fc86213e648dab4cbf3ca298646eae01bcc71725bf07f4d64736f6c63430007040033",
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

// UtilitiesMetaData contains all meta data concerning the Utilities contract.
var UtilitiesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212200df04770b16a7d56fe88f44d5dabece30921705725c348b2ba0537247578c99d64736f6c63430007040033",
}

// UtilitiesABI is the input ABI used to generate the binding from.
// Deprecated: Use UtilitiesMetaData.ABI instead.
var UtilitiesABI = UtilitiesMetaData.ABI

// UtilitiesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UtilitiesMetaData.Bin instead.
var UtilitiesBin = UtilitiesMetaData.Bin

// DeployUtilities deploys a new Ethereum contract, binding an instance of Utilities to it.
func DeployUtilities(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utilities, error) {
	parsed, err := UtilitiesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UtilitiesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utilities{UtilitiesCaller: UtilitiesCaller{contract: contract}, UtilitiesTransactor: UtilitiesTransactor{contract: contract}, UtilitiesFilterer: UtilitiesFilterer{contract: contract}}, nil
}

// Utilities is an auto generated Go binding around an Ethereum contract.
type Utilities struct {
	UtilitiesCaller     // Read-only binding to the contract
	UtilitiesTransactor // Write-only binding to the contract
	UtilitiesFilterer   // Log filterer for contract events
}

// UtilitiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilitiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilitiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilitiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilitiesSession struct {
	Contract     *Utilities        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilitiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilitiesCallerSession struct {
	Contract *UtilitiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UtilitiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilitiesTransactorSession struct {
	Contract     *UtilitiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UtilitiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilitiesRaw struct {
	Contract *Utilities // Generic contract binding to access the raw methods on
}

// UtilitiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilitiesCallerRaw struct {
	Contract *UtilitiesCaller // Generic read-only contract binding to access the raw methods on
}

// UtilitiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilitiesTransactorRaw struct {
	Contract *UtilitiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtilities creates a new instance of Utilities, bound to a specific deployed contract.
func NewUtilities(address common.Address, backend bind.ContractBackend) (*Utilities, error) {
	contract, err := bindUtilities(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utilities{UtilitiesCaller: UtilitiesCaller{contract: contract}, UtilitiesTransactor: UtilitiesTransactor{contract: contract}, UtilitiesFilterer: UtilitiesFilterer{contract: contract}}, nil
}

// NewUtilitiesCaller creates a new read-only instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesCaller(address common.Address, caller bind.ContractCaller) (*UtilitiesCaller, error) {
	contract, err := bindUtilities(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilitiesCaller{contract: contract}, nil
}

// NewUtilitiesTransactor creates a new write-only instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilitiesTransactor, error) {
	contract, err := bindUtilities(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilitiesTransactor{contract: contract}, nil
}

// NewUtilitiesFilterer creates a new log filterer instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilitiesFilterer, error) {
	contract, err := bindUtilities(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilitiesFilterer{contract: contract}, nil
}

// bindUtilities binds a generic wrapper to an already deployed contract.
func bindUtilities(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilitiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utilities *UtilitiesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Utilities.Contract.UtilitiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utilities *UtilitiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utilities.Contract.UtilitiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utilities *UtilitiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utilities.Contract.UtilitiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utilities *UtilitiesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Utilities.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utilities *UtilitiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utilities.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utilities *UtilitiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utilities.Contract.contract.Transact(opts, method, params...)
}
