// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_gen

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// FuxBatchABI is the input ABI used to generate the binding from.
const FuxBatchABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isUsingRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sysRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CONFIG\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_isFuxHub\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CFRBAC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRouterAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_HUB\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_isInFuxGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_router\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_len\",\"type\":\"uint256\"}],\"name\":\"setMaxLength\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokens\",\"type\":\"uint256[]\"}],\"name\":\"addJob\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"runJob\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isCompile\",\"outputs\":[{\"name\":\"compile\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FuxBatchBin is the compiled bytecode used for deploying new contracts.
const FuxBatchBin = `0x60806040523480156200001157600080fd5b506040516020806200166f83398101806040528101908080519060200190929190505050620000846040805190810160405280600c81526020017f726f75746572436f6e6669670000000000000000000000000000000000000000815250620000fc640100000000026401000000009004565b620000d36040805190810160405280600981526020017f467578436f6e6669670000000000000000000000000000000000000000000000815250620000fc640100000000026401000000009004565b620000ed8162000118640100000000026401000000009004565b601e6002819055505062000247565b80600190805190602001906200011492919062000198565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156200015557600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001db57805160ff19168380011785556200020c565b828001600101855582156200020c579182015b828111156200020b578251825591602001919060010190620001ee565b5b5090506200021b91906200021f565b5090565b6200024491905b808211156200024057600081600090555060010162000226565b5090565b90565b61141880620002576000396000f3006080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630d20aa48146100e057806366b9852b1461010f5780636ed56bce146101665780639a1e4822146101f6578063a3c332ab14610223578063abf836fc14610268578063ade6172414610318578063b8d43d2714610373578063c427d38d14610403578063d06a89a414610493578063d54f7d5e146104be578063d6b419fb14610515578063d89e6c4f146105a5578063dc2f786714610635578063f0b15e8214610662575b600080fd5b3480156100ec57600080fd5b506100f56106bd565b604051808215151515815260200191505060405180910390f35b34801561011b57600080fd5b506101246106c6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561017257600080fd5b5061017b6106eb565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101bb5780820151818401526020810190506101a0565b50505050905090810190601f1680156101e85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561020257600080fd5b5061022160048036038101908080359060200190929190505050610724565b005b34801561022f57600080fd5b5061024e6004803603810190808035906020019092919050505061094d565b604051808215151515815260200191505060405180910390f35b34801561027457600080fd5b5061031660048036038101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050610977565b005b34801561032457600080fd5b50610359600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b71565b604051808215151515815260200191505060405180910390f35b34801561037f57600080fd5b50610388610cf9565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103c85780820151818401526020810190506103ad565b50505050905090810190601f1680156103f55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561040f57600080fd5b50610418610d32565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561045857808201518184015260208101905061043d565b50505050905090810190601f1680156104855780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561049f57600080fd5b506104a8610d6b565b6040518082815260200191505060405180910390f35b3480156104ca57600080fd5b506104d3610d71565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561052157600080fd5b5061052a610d9a565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561056a57808201518184015260208101905061054f565b50505050905090810190601f1680156105975780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156105b157600080fd5b506105ba610dd3565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105fa5780820151818401526020810190506105df565b50505050905090810190601f1680156106275780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561064157600080fd5b5061066060048036038101908080359060200190929190505050610e0c565b005b34801561066e57600080fd5b506106a3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e2a565b604051808215151515815260200191505060405180910390f35b60006001905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600c81526020017f726f75746572436f6e666967000000000000000000000000000000000000000081525081565b600080600061073233610e2a565b80610742575061074133610fb2565b5b151561074d57600080fd5b6107568461094d565b1561076057610947565b6004600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1692506005600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16915060016006600086815260200190815260200160002060006101000a81548160ff021916908315150217905550600090505b600360008581526020019081526020016000208054905081101561094657610823611098565b73ffffffffffffffffffffffffffffffffffffffff1663beabacc88484600360008981526020019081526020016000208581548110151561086057fe5b90600052602060002001546040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561092357600080fd5b505af1158015610937573d6000803e3d6000fd5b505050508060010190506107fd565b5b50505050565b60006006600083815260200190815260200160002060009054906101000a900460ff169050919050565b61098033610e2a565b80610990575061098f33610fb2565b5b151561099b57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141515156109d757600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515610a1357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166004600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610a8157600080fd5b6002548151108015610a94575060008151115b1515610a9f57600080fd5b80600360008681526020019081526020016000209080519060200190610ac692919061137a565b50826004600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816005600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b6000610b7b6110dd565b73ffffffffffffffffffffffffffffffffffffffff1663217fe6c6836040805190810160405280600681526020017f66757848756200000000000000000000000000000000000000000000000000008152506040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610c6a578082015181840152602081019050610c4f565b50505050905090810190601f168015610c975780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015610cb757600080fd5b505af1158015610ccb573d6000803e3d6000fd5b505050506040513d6020811015610ce157600080fd5b81019080805190602001909291905050509050919050565b6040805190810160405280600b81526020017f726f6c65436c757374657200000000000000000000000000000000000000000081525081565b6040805190810160405280600c81526020017f726f75746572434652424143000000000000000000000000000000000000000081525081565b60025481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6040805190810160405280600681526020017f667578487562000000000000000000000000000000000000000000000000000081525081565b6040805190810160405280600c81526020017f726f6c6546757847726f7570000000000000000000000000000000000000000081525081565b610e1533610fb2565b1515610e2057600080fd5b8060028190555050565b6000610e346110dd565b73ffffffffffffffffffffffffffffffffffffffff1663217fe6c6836040805190810160405280600c81526020017f726f6c6546757847726f757000000000000000000000000000000000000000008152506040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610f23578082015181840152602081019050610f08565b50505050905090810190601f168015610f505780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015610f7057600080fd5b505af1158015610f84573d6000803e3d6000fd5b505050506040513d6020811015610f9a57600080fd5b81019080805190602001909291905050509050919050565b6000610fbc6110dd565b73ffffffffffffffffffffffffffffffffffffffff166324d7806c836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561105657600080fd5b505af115801561106a573d6000803e3d6000fd5b505050506040513d602081101561108057600080fd5b81019080805190602001909291905050509050919050565b60006110d86040805190810160405280600881526020017f467578546f6b656e0000000000000000000000000000000000000000000000008152506110ec565b905090565b60006110e7611225565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c2511025836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561119757808201518184015260208101905061117c565b50505050905090810190601f1680156111c45780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1580156111e357600080fd5b505af11580156111f7573d6000803e3d6000fd5b505050506040513d602081101561120d57600080fd5b81019080805190602001909291905050509050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c251102560016040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200182810382528381815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561131a5780601f106112ef5761010080835404028352916020019161131a565b820191906000526020600020905b8154815290600101906020018083116112fd57829003601f168201915b505092505050602060405180830381600087803b15801561133a57600080fd5b505af115801561134e573d6000803e3d6000fd5b505050506040513d602081101561136457600080fd5b8101908080519060200190929190505050905090565b8280548282559060005260206000209081019282156113b6579160200282015b828111156113b557825182559160200191906001019061139a565b5b5090506113c391906113c7565b5090565b6113e991905b808211156113e55760008160009055506001016113cd565b5090565b905600a165627a7a72305820413fa7679fdcbf49c34f0deab4f61e89c3ae8284a5b934236193743ecd696cff0029`

// DeployFuxBatch deploys a new Ethereum contract, binding an instance of FuxBatch to it.
func DeployFuxBatch(auth *bind.TransactOpts, backend bind.ContractBackend, _router common.Address) (common.Address, *types.Transaction, *FuxBatch, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxBatchABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FuxBatchBin), backend, _router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FuxBatch{FuxBatchCaller: FuxBatchCaller{contract: contract}, FuxBatchTransactor: FuxBatchTransactor{contract: contract}, FuxBatchFilterer: FuxBatchFilterer{contract: contract}}, nil
}

// FuxBatch is an auto generated Go binding around an Ethereum contract.
type FuxBatch struct {
	FuxBatchCaller     // Read-only binding to the contract
	FuxBatchTransactor // Write-only binding to the contract
	FuxBatchFilterer   // Log filterer for contract events
}

// FuxBatchCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxBatchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxBatchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxBatchTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxBatchFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FuxBatchFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxBatchSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuxBatchSession struct {
	Contract     *FuxBatch         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuxBatchCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuxBatchCallerSession struct {
	Contract *FuxBatchCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FuxBatchTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuxBatchTransactorSession struct {
	Contract     *FuxBatchTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FuxBatchRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuxBatchRaw struct {
	Contract *FuxBatch // Generic contract binding to access the raw methods on
}

// FuxBatchCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuxBatchCallerRaw struct {
	Contract *FuxBatchCaller // Generic read-only contract binding to access the raw methods on
}

// FuxBatchTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuxBatchTransactorRaw struct {
	Contract *FuxBatchTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFuxBatch creates a new instance of FuxBatch, bound to a specific deployed contract.
func NewFuxBatch(address common.Address, backend bind.ContractBackend) (*FuxBatch, error) {
	contract, err := bindFuxBatch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxBatch{FuxBatchCaller: FuxBatchCaller{contract: contract}, FuxBatchTransactor: FuxBatchTransactor{contract: contract}, FuxBatchFilterer: FuxBatchFilterer{contract: contract}}, nil
}

// NewFuxBatchCaller creates a new read-only instance of FuxBatch, bound to a specific deployed contract.
func NewFuxBatchCaller(address common.Address, caller bind.ContractCaller) (*FuxBatchCaller, error) {
	contract, err := bindFuxBatch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FuxBatchCaller{contract: contract}, nil
}

// NewFuxBatchTransactor creates a new write-only instance of FuxBatch, bound to a specific deployed contract.
func NewFuxBatchTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxBatchTransactor, error) {
	contract, err := bindFuxBatch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FuxBatchTransactor{contract: contract}, nil
}

// NewFuxBatchFilterer creates a new log filterer instance of FuxBatch, bound to a specific deployed contract.
func NewFuxBatchFilterer(address common.Address, filterer bind.ContractFilterer) (*FuxBatchFilterer, error) {
	contract, err := bindFuxBatch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FuxBatchFilterer{contract: contract}, nil
}

// bindFuxBatch binds a generic wrapper to an already deployed contract.
func bindFuxBatch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxBatchABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxBatch *FuxBatchRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxBatch.Contract.FuxBatchCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxBatch *FuxBatchRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxBatch.Contract.FuxBatchTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxBatch *FuxBatchRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxBatch.Contract.FuxBatchTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxBatch *FuxBatchCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxBatch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxBatch *FuxBatchTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxBatch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxBatch *FuxBatchTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxBatch.Contract.contract.Transact(opts, method, params...)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROLECLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROLECLUSTER() (string, error) {
	return _FuxBatch.Contract.ROLECLUSTER(&_FuxBatch.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROLECLUSTER() (string, error) {
	return _FuxBatch.Contract.ROLECLUSTER(&_FuxBatch.CallOpts)
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROLEFUXGROUP(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROLE_FUX_GROUP")
	return *ret0, err
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROLEFUXGROUP() (string, error) {
	return _FuxBatch.Contract.ROLEFUXGROUP(&_FuxBatch.CallOpts)
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROLEFUXGROUP() (string, error) {
	return _FuxBatch.Contract.ROLEFUXGROUP(&_FuxBatch.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROLEFUXHUB(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROLE_FUX_HUB")
	return *ret0, err
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROLEFUXHUB() (string, error) {
	return _FuxBatch.Contract.ROLEFUXHUB(&_FuxBatch.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROLEFUXHUB() (string, error) {
	return _FuxBatch.Contract.ROLEFUXHUB(&_FuxBatch.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTECFRBAC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_CFRBAC")
	return *ret0, err
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTECFRBAC() (string, error) {
	return _FuxBatch.Contract.ROUTECFRBAC(&_FuxBatch.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTECFRBAC() (string, error) {
	return _FuxBatch.Contract.ROUTECFRBAC(&_FuxBatch.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxBatch *FuxBatchCaller) ROUTECONFIG(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "ROUTE_CONFIG")
	return *ret0, err
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxBatch *FuxBatchSession) ROUTECONFIG() (string, error) {
	return _FuxBatch.Contract.ROUTECONFIG(&_FuxBatch.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxBatch *FuxBatchCallerSession) ROUTECONFIG() (string, error) {
	return _FuxBatch.Contract.ROUTECONFIG(&_FuxBatch.CallOpts)
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxBatch *FuxBatchCaller) IsFuxHub(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "_isFuxHub", _addr)
	return *ret0, err
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxBatch *FuxBatchSession) IsFuxHub(_addr common.Address) (bool, error) {
	return _FuxBatch.Contract.IsFuxHub(&_FuxBatch.CallOpts, _addr)
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxBatch *FuxBatchCallerSession) IsFuxHub(_addr common.Address) (bool, error) {
	return _FuxBatch.Contract.IsFuxHub(&_FuxBatch.CallOpts, _addr)
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxBatch *FuxBatchCaller) IsInFuxGroup(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "_isInFuxGroup", _addr)
	return *ret0, err
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxBatch *FuxBatchSession) IsInFuxGroup(_addr common.Address) (bool, error) {
	return _FuxBatch.Contract.IsInFuxGroup(&_FuxBatch.CallOpts, _addr)
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxBatch *FuxBatchCallerSession) IsInFuxGroup(_addr common.Address) (bool, error) {
	return _FuxBatch.Contract.IsInFuxGroup(&_FuxBatch.CallOpts, _addr)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxBatch *FuxBatchCaller) GetRouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "getRouterAddress")
	return *ret0, err
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxBatch *FuxBatchSession) GetRouterAddress() (common.Address, error) {
	return _FuxBatch.Contract.GetRouterAddress(&_FuxBatch.CallOpts)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxBatch *FuxBatchCallerSession) GetRouterAddress() (common.Address, error) {
	return _FuxBatch.Contract.GetRouterAddress(&_FuxBatch.CallOpts)
}

// IsCompile is a free data retrieval call binding the contract method 0xa3c332ab.
//
// Solidity: function isCompile(_id uint256) constant returns(compile bool)
func (_FuxBatch *FuxBatchCaller) IsCompile(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "isCompile", _id)
	return *ret0, err
}

// IsCompile is a free data retrieval call binding the contract method 0xa3c332ab.
//
// Solidity: function isCompile(_id uint256) constant returns(compile bool)
func (_FuxBatch *FuxBatchSession) IsCompile(_id *big.Int) (bool, error) {
	return _FuxBatch.Contract.IsCompile(&_FuxBatch.CallOpts, _id)
}

// IsCompile is a free data retrieval call binding the contract method 0xa3c332ab.
//
// Solidity: function isCompile(_id uint256) constant returns(compile bool)
func (_FuxBatch *FuxBatchCallerSession) IsCompile(_id *big.Int) (bool, error) {
	return _FuxBatch.Contract.IsCompile(&_FuxBatch.CallOpts, _id)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxBatch *FuxBatchCaller) IsUsingRouter(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "isUsingRouter")
	return *ret0, err
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxBatch *FuxBatchSession) IsUsingRouter() (bool, error) {
	return _FuxBatch.Contract.IsUsingRouter(&_FuxBatch.CallOpts)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxBatch *FuxBatchCallerSession) IsUsingRouter() (bool, error) {
	return _FuxBatch.Contract.IsUsingRouter(&_FuxBatch.CallOpts)
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_FuxBatch *FuxBatchCaller) MaxLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "maxLength")
	return *ret0, err
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_FuxBatch *FuxBatchSession) MaxLength() (*big.Int, error) {
	return _FuxBatch.Contract.MaxLength(&_FuxBatch.CallOpts)
}

// MaxLength is a free data retrieval call binding the contract method 0xd06a89a4.
//
// Solidity: function maxLength() constant returns(uint256)
func (_FuxBatch *FuxBatchCallerSession) MaxLength() (*big.Int, error) {
	return _FuxBatch.Contract.MaxLength(&_FuxBatch.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxBatch *FuxBatchCaller) SysRouter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxBatch.contract.Call(opts, out, "sysRouter")
	return *ret0, err
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxBatch *FuxBatchSession) SysRouter() (common.Address, error) {
	return _FuxBatch.Contract.SysRouter(&_FuxBatch.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxBatch *FuxBatchCallerSession) SysRouter() (common.Address, error) {
	return _FuxBatch.Contract.SysRouter(&_FuxBatch.CallOpts)
}

// AddJob is a paid mutator transaction binding the contract method 0xabf836fc.
//
// Solidity: function addJob(_id uint256, _from address, _to address, _tokens uint256[]) returns()
func (_FuxBatch *FuxBatchTransactor) AddJob(opts *bind.TransactOpts, _id *big.Int, _from common.Address, _to common.Address, _tokens []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "addJob", _id, _from, _to, _tokens)
}

// AddJob is a paid mutator transaction binding the contract method 0xabf836fc.
//
// Solidity: function addJob(_id uint256, _from address, _to address, _tokens uint256[]) returns()
func (_FuxBatch *FuxBatchSession) AddJob(_id *big.Int, _from common.Address, _to common.Address, _tokens []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.AddJob(&_FuxBatch.TransactOpts, _id, _from, _to, _tokens)
}

// AddJob is a paid mutator transaction binding the contract method 0xabf836fc.
//
// Solidity: function addJob(_id uint256, _from address, _to address, _tokens uint256[]) returns()
func (_FuxBatch *FuxBatchTransactorSession) AddJob(_id *big.Int, _from common.Address, _to common.Address, _tokens []*big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.AddJob(&_FuxBatch.TransactOpts, _id, _from, _to, _tokens)
}

// RunJob is a paid mutator transaction binding the contract method 0x9a1e4822.
//
// Solidity: function runJob(_id uint256) returns()
func (_FuxBatch *FuxBatchTransactor) RunJob(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "runJob", _id)
}

// RunJob is a paid mutator transaction binding the contract method 0x9a1e4822.
//
// Solidity: function runJob(_id uint256) returns()
func (_FuxBatch *FuxBatchSession) RunJob(_id *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.RunJob(&_FuxBatch.TransactOpts, _id)
}

// RunJob is a paid mutator transaction binding the contract method 0x9a1e4822.
//
// Solidity: function runJob(_id uint256) returns()
func (_FuxBatch *FuxBatchTransactorSession) RunJob(_id *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.RunJob(&_FuxBatch.TransactOpts, _id)
}

// SetMaxLength is a paid mutator transaction binding the contract method 0xdc2f7867.
//
// Solidity: function setMaxLength(_len uint256) returns()
func (_FuxBatch *FuxBatchTransactor) SetMaxLength(opts *bind.TransactOpts, _len *big.Int) (*types.Transaction, error) {
	return _FuxBatch.contract.Transact(opts, "setMaxLength", _len)
}

// SetMaxLength is a paid mutator transaction binding the contract method 0xdc2f7867.
//
// Solidity: function setMaxLength(_len uint256) returns()
func (_FuxBatch *FuxBatchSession) SetMaxLength(_len *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.SetMaxLength(&_FuxBatch.TransactOpts, _len)
}

// SetMaxLength is a paid mutator transaction binding the contract method 0xdc2f7867.
//
// Solidity: function setMaxLength(_len uint256) returns()
func (_FuxBatch *FuxBatchTransactorSession) SetMaxLength(_len *big.Int) (*types.Transaction, error) {
	return _FuxBatch.Contract.SetMaxLength(&_FuxBatch.TransactOpts, _len)
}
