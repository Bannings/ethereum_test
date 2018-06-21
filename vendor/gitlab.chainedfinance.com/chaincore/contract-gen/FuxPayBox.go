// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_gen

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// FuxPayBoxABI is the input ABI used to generate the binding from.
const FuxPayBoxABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isUsingRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sysRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CONFIG\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CFRBAC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRouterAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_HUB\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_router\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FuxPayBoxBin is the compiled bytecode used for deploying new contracts.
const FuxPayBoxBin = `0x608060405234801561001057600080fd5b506040516020806112b38339810180604052810190808051906020019092919050505033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061008c81610092640100000000026401000000009004565b50610111565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156100ce57600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611193806101206000396000f3006080604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630d20aa48146100d557806341c0e1b5146101045780634e71e0c81461011b57806366b9852b146101325780636ed56bce146101895780638da5cb5b14610219578063a9059cbb14610270578063c427d38d146102bd578063d54f7d5e1461034d578063d6b419fb146103a4578063d89e6c4f14610434578063e30c3978146104c4578063f0b9e5ba1461051b578063f2fde38b14610600575b600080fd5b3480156100e157600080fd5b506100ea610643565b604051808215151515815260200191505060405180910390f35b34801561011057600080fd5b5061011961064c565b005b34801561012757600080fd5b506101306107e9565b005b34801561013e57600080fd5b50610147610a58565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561019557600080fd5b5061019e610a7d565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101de5780820151818401526020810190506101c3565b50505050905090810190601f16801561020b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561022557600080fd5b5061022e610ab6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561027c57600080fd5b506102bb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610adc565b005b3480156102c957600080fd5b506102d2610c32565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103125780820151818401526020810190506102f7565b50505050905090810190601f16801561033f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561035957600080fd5b50610362610c6b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103b057600080fd5b506103b9610c94565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103f95780820151818401526020810190506103de565b50505050905090810190601f1680156104265780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561044057600080fd5b50610449610ccd565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561048957808201518184015260208101905061046e565b50505050905090810190601f1680156104b65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104d057600080fd5b506104d9610d06565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561052757600080fd5b506105ac600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610d2c565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b34801561060c57600080fd5b50610641600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d5b565b005b60006001905090565b61068a6040805190810160405280601081526020017f467578506179426f78466163746f727900000000000000000000000000000000815250610d8b565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156106c357600080fd5b60006106cd610ec4565b73ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561076757600080fd5b505af115801561077b573d6000803e3d6000fd5b505050506040513d602081101561079157600080fd5b81019080805190602001909291905050501415156107ae57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b6108276040805190810160405280601081526020017f467578506179426f78466163746f727900000000000000000000000000000000815250610d8b565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480156108b05750600073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b806109085750600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b151561091357600080fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600c81526020017f726f75746572436f6e666967000000000000000000000000000000000000000081525081565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b3857600080fd5b610b40610ec4565b73ffffffffffffffffffffffffffffffffffffffff1663beabacc83084846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015610c1657600080fd5b505af1158015610c2a573d6000803e3d6000fd5b505050505050565b6040805190810160405280600c81526020017f726f75746572434652424143000000000000000000000000000000000000000081525081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6040805190810160405280600681526020017f667578487562000000000000000000000000000000000000000000000000000081525081565b6040805190810160405280600c81526020017f726f6c6546757847726f7570000000000000000000000000000000000000000081525081565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600063f0b9e5ba7c01000000000000000000000000000000000000000000000000000000000290509392505050565b610d6481610f09565b80610d745750610d7333610f09565b5b1515610d7f57600080fd5b610d88816110c7565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c2511025836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610e36578082015181840152602081019050610e1b565b50505050905090810190601f168015610e635780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b158015610e8257600080fd5b505af1158015610e96573d6000803e3d6000fd5b505050506040513d6020811015610eac57600080fd5b81019080805190602001909291905050509050919050565b6000610f046040805190810160405280600881526020017f467578546f6b656e000000000000000000000000000000000000000000000000815250610d8b565b905090565b6000610f496040805190810160405280600c81526020017f726f757465724346524241430000000000000000000000000000000000000000815250610d8b565b73ffffffffffffffffffffffffffffffffffffffff1663217fe6c6836040805190810160405280600681526020017f66757848756200000000000000000000000000000000000000000000000000008152506040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561103857808201518184015260208101905061101d565b50505050905090810190601f1680156110655780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561108557600080fd5b505af1158015611099573d6000803e3d6000fd5b505050506040513d60208110156110af57600080fd5b81019080805190602001909291905050509050919050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561112357600080fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a7230582016b04d762763cc24dc866be2bd034d5e67eb51d0d04d82e210a8370769dbafa80029`

// DeployFuxPayBox deploys a new Ethereum contract, binding an instance of FuxPayBox to it.
func DeployFuxPayBox(auth *bind.TransactOpts, backend bind.ContractBackend, _router common.Address) (common.Address, *types.Transaction, *FuxPayBox, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxPayBoxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FuxPayBoxBin), backend, _router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FuxPayBox{FuxPayBoxCaller: FuxPayBoxCaller{contract: contract}, FuxPayBoxTransactor: FuxPayBoxTransactor{contract: contract}, FuxPayBoxFilterer: FuxPayBoxFilterer{contract: contract}}, nil
}

// FuxPayBox is an auto generated Go binding around an Ethereum contract.
type FuxPayBox struct {
	FuxPayBoxCaller     // Read-only binding to the contract
	FuxPayBoxTransactor // Write-only binding to the contract
	FuxPayBoxFilterer   // Log filterer for contract events
}

// FuxPayBoxCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxPayBoxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxPayBoxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxPayBoxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxPayBoxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FuxPayBoxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxPayBoxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuxPayBoxSession struct {
	Contract     *FuxPayBox        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuxPayBoxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuxPayBoxCallerSession struct {
	Contract *FuxPayBoxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FuxPayBoxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuxPayBoxTransactorSession struct {
	Contract     *FuxPayBoxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FuxPayBoxRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuxPayBoxRaw struct {
	Contract *FuxPayBox // Generic contract binding to access the raw methods on
}

// FuxPayBoxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuxPayBoxCallerRaw struct {
	Contract *FuxPayBoxCaller // Generic read-only contract binding to access the raw methods on
}

// FuxPayBoxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuxPayBoxTransactorRaw struct {
	Contract *FuxPayBoxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFuxPayBox creates a new instance of FuxPayBox, bound to a specific deployed contract.
func NewFuxPayBox(address common.Address, backend bind.ContractBackend) (*FuxPayBox, error) {
	contract, err := bindFuxPayBox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxPayBox{FuxPayBoxCaller: FuxPayBoxCaller{contract: contract}, FuxPayBoxTransactor: FuxPayBoxTransactor{contract: contract}, FuxPayBoxFilterer: FuxPayBoxFilterer{contract: contract}}, nil
}

// NewFuxPayBoxCaller creates a new read-only instance of FuxPayBox, bound to a specific deployed contract.
func NewFuxPayBoxCaller(address common.Address, caller bind.ContractCaller) (*FuxPayBoxCaller, error) {
	contract, err := bindFuxPayBox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FuxPayBoxCaller{contract: contract}, nil
}

// NewFuxPayBoxTransactor creates a new write-only instance of FuxPayBox, bound to a specific deployed contract.
func NewFuxPayBoxTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxPayBoxTransactor, error) {
	contract, err := bindFuxPayBox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FuxPayBoxTransactor{contract: contract}, nil
}

// NewFuxPayBoxFilterer creates a new log filterer instance of FuxPayBox, bound to a specific deployed contract.
func NewFuxPayBoxFilterer(address common.Address, filterer bind.ContractFilterer) (*FuxPayBoxFilterer, error) {
	contract, err := bindFuxPayBox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FuxPayBoxFilterer{contract: contract}, nil
}

// bindFuxPayBox binds a generic wrapper to an already deployed contract.
func bindFuxPayBox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxPayBoxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxPayBox *FuxPayBoxRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxPayBox.Contract.FuxPayBoxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxPayBox *FuxPayBoxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxPayBox.Contract.FuxPayBoxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxPayBox *FuxPayBoxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxPayBox.Contract.FuxPayBoxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxPayBox *FuxPayBoxCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxPayBox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxPayBox *FuxPayBoxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxPayBox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxPayBox *FuxPayBoxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxPayBox.Contract.contract.Transact(opts, method, params...)
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROLEFUXGROUP(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROLE_FUX_GROUP")
	return *ret0, err
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROLEFUXGROUP() (string, error) {
	return _FuxPayBox.Contract.ROLEFUXGROUP(&_FuxPayBox.CallOpts)
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROLEFUXGROUP() (string, error) {
	return _FuxPayBox.Contract.ROLEFUXGROUP(&_FuxPayBox.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROLEFUXHUB(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROLE_FUX_HUB")
	return *ret0, err
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROLEFUXHUB() (string, error) {
	return _FuxPayBox.Contract.ROLEFUXHUB(&_FuxPayBox.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROLEFUXHUB() (string, error) {
	return _FuxPayBox.Contract.ROLEFUXHUB(&_FuxPayBox.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROUTECFRBAC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROUTE_CFRBAC")
	return *ret0, err
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROUTECFRBAC() (string, error) {
	return _FuxPayBox.Contract.ROUTECFRBAC(&_FuxPayBox.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROUTECFRBAC() (string, error) {
	return _FuxPayBox.Contract.ROUTECFRBAC(&_FuxPayBox.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROUTECONFIG(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROUTE_CONFIG")
	return *ret0, err
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROUTECONFIG() (string, error) {
	return _FuxPayBox.Contract.ROUTECONFIG(&_FuxPayBox.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROUTECONFIG() (string, error) {
	return _FuxPayBox.Contract.ROUTECONFIG(&_FuxPayBox.CallOpts)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxPayBox *FuxPayBoxCaller) GetRouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "getRouterAddress")
	return *ret0, err
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxPayBox *FuxPayBoxSession) GetRouterAddress() (common.Address, error) {
	return _FuxPayBox.Contract.GetRouterAddress(&_FuxPayBox.CallOpts)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxPayBox *FuxPayBoxCallerSession) GetRouterAddress() (common.Address, error) {
	return _FuxPayBox.Contract.GetRouterAddress(&_FuxPayBox.CallOpts)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxPayBox *FuxPayBoxCaller) IsUsingRouter(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "isUsingRouter")
	return *ret0, err
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxPayBox *FuxPayBoxSession) IsUsingRouter() (bool, error) {
	return _FuxPayBox.Contract.IsUsingRouter(&_FuxPayBox.CallOpts)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxPayBox *FuxPayBoxCallerSession) IsUsingRouter() (bool, error) {
	return _FuxPayBox.Contract.IsUsingRouter(&_FuxPayBox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FuxPayBox *FuxPayBoxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FuxPayBox *FuxPayBoxSession) Owner() (common.Address, error) {
	return _FuxPayBox.Contract.Owner(&_FuxPayBox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FuxPayBox *FuxPayBoxCallerSession) Owner() (common.Address, error) {
	return _FuxPayBox.Contract.Owner(&_FuxPayBox.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_FuxPayBox *FuxPayBoxCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_FuxPayBox *FuxPayBoxSession) PendingOwner() (common.Address, error) {
	return _FuxPayBox.Contract.PendingOwner(&_FuxPayBox.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_FuxPayBox *FuxPayBoxCallerSession) PendingOwner() (common.Address, error) {
	return _FuxPayBox.Contract.PendingOwner(&_FuxPayBox.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxPayBox *FuxPayBoxCaller) SysRouter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "sysRouter")
	return *ret0, err
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxPayBox *FuxPayBoxSession) SysRouter() (common.Address, error) {
	return _FuxPayBox.Contract.SysRouter(&_FuxPayBox.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxPayBox *FuxPayBoxCallerSession) SysRouter() (common.Address, error) {
	return _FuxPayBox.Contract.SysRouter(&_FuxPayBox.CallOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_FuxPayBox *FuxPayBoxTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxPayBox.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_FuxPayBox *FuxPayBoxSession) ClaimOwnership() (*types.Transaction, error) {
	return _FuxPayBox.Contract.ClaimOwnership(&_FuxPayBox.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_FuxPayBox *FuxPayBoxTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _FuxPayBox.Contract.ClaimOwnership(&_FuxPayBox.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_FuxPayBox *FuxPayBoxTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxPayBox.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_FuxPayBox *FuxPayBoxSession) Kill() (*types.Transaction, error) {
	return _FuxPayBox.Contract.Kill(&_FuxPayBox.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_FuxPayBox *FuxPayBoxTransactorSession) Kill() (*types.Transaction, error) {
	return _FuxPayBox.Contract.Kill(&_FuxPayBox.TransactOpts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xf0b9e5ba.
//
// Solidity: function onERC721Received( address,  uint256,  bytes) returns(bytes4)
func (_FuxPayBox *FuxPayBoxTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _FuxPayBox.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xf0b9e5ba.
//
// Solidity: function onERC721Received( address,  uint256,  bytes) returns(bytes4)
func (_FuxPayBox *FuxPayBoxSession) OnERC721Received(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _FuxPayBox.Contract.OnERC721Received(&_FuxPayBox.TransactOpts, arg0, arg1, arg2)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xf0b9e5ba.
//
// Solidity: function onERC721Received( address,  uint256,  bytes) returns(bytes4)
func (_FuxPayBox *FuxPayBoxTransactorSession) OnERC721Received(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _FuxPayBox.Contract.OnERC721Received(&_FuxPayBox.TransactOpts, arg0, arg1, arg2)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_FuxPayBox *FuxPayBoxTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxPayBox.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_FuxPayBox *FuxPayBoxSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxPayBox.Contract.Transfer(&_FuxPayBox.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_FuxPayBox *FuxPayBoxTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxPayBox.Contract.Transfer(&_FuxPayBox.TransactOpts, _to, _tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FuxPayBox *FuxPayBoxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FuxPayBox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FuxPayBox *FuxPayBoxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FuxPayBox.Contract.TransferOwnership(&_FuxPayBox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FuxPayBox *FuxPayBoxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FuxPayBox.Contract.TransferOwnership(&_FuxPayBox.TransactOpts, newOwner)
}

// FuxPayBoxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FuxPayBox contract.
type FuxPayBoxOwnershipTransferredIterator struct {
	Event *FuxPayBoxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FuxPayBoxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxPayBoxOwnershipTransferred)
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
		it.Event = new(FuxPayBoxOwnershipTransferred)
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
func (it *FuxPayBoxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxPayBoxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxPayBoxOwnershipTransferred represents a OwnershipTransferred event raised by the FuxPayBox contract.
type FuxPayBoxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_FuxPayBox *FuxPayBoxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FuxPayBoxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FuxPayBox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FuxPayBoxOwnershipTransferredIterator{contract: _FuxPayBox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_FuxPayBox *FuxPayBoxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FuxPayBoxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FuxPayBox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxPayBoxOwnershipTransferred)
				if err := _FuxPayBox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
