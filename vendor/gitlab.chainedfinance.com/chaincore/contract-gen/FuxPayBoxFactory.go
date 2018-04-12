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

// FuxPayBoxFactoryABI is the input ABI used to generate the binding from.
const FuxPayBoxFactoryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFuxIssuer\",\"outputs\":[{\"name\":\"issuer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fuxIssuer\",\"type\":\"address\"}],\"name\":\"setFuxIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fuxToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setFuxToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_boxId\",\"type\":\"uint256\"}],\"name\":\"createPayBox\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_boxId\",\"type\":\"uint256\"}],\"name\":\"getPayBoxAddress\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getBoxId\",\"outputs\":[{\"name\":\"boxId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_boxAddr\",\"type\":\"address\"}],\"name\":\"closeBox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FuxPayBoxFactoryBin is the compiled bytecode used for deploying new contracts.
const FuxPayBoxFactoryBin = `0x6060604052341561000f57600080fd5b6040516020806116d383398101604052808051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050611617806100bc6000396000f3006060604052600436106100ba576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806328baafe1146100bf5780634c982e57146100f8578063606c56321461015b578063704b6c02146101945780638da5cb5b146101cd578063904c00c114610222578063946f016714610277578063c0b43a45146102c4578063d91f5a6a146102fd578063f2fde38b14610352578063f851a4401461038b578063fcfc43a7146103e0575b600080fd5b34156100ca57600080fd5b6100f6600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610443565b005b341561010357600080fd5b61011960048080359060200190919050506104e2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561016657600080fd5b610192600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506107a4565b005b341561019f57600080fd5b6101cb600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506108b0565b005b34156101d857600080fd5b6101e061094f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561022d57600080fd5b610235610974565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561028257600080fd5b6102ae600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610a1e565b6040518082815260200191505060405180910390f35b34156102cf57600080fd5b6102fb600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610a67565b005b341561030857600080fd5b610310610b06565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561035d57600080fd5b610389600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610b2c565b005b341561039657600080fd5b61039e610c81565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103eb57600080fd5b6104016004808035906020019091905050610ca7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561049e57600080fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060008073ffffffffffffffffffffffffffffffffffffffff166004600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561055457600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16306105a3610ce4565b808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019350505050604051809103906000f080151561065557600080fd5b91508173ffffffffffffffffffffffffffffffffffffffff1663f2fde38b336040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15156106f157600080fd5b5af115156106fe57600080fd5b505050819050806004600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508092505050919050565b60006107ae610974565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156107e757600080fd5b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541415151561083657600080fd5b8190508073ffffffffffffffffffffffffffffffffffffffff166341c0e1b56040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b151561089c57600080fd5b5af115156108a957600080fd5b5050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561090b57600080fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156109f5576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050610a1b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b90565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ac257600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b8757600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610bc357600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6040516108f780610cf58339019056006060604052341561000f57600080fd5b6040516060806108f783398101604052808051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506107a5806101526000396000f30060606040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633d8d3fd91461007257806341c0e1b5146100ab5780638da5cb5b146100c0578063a9059cbb14610115578063f2fde38b14610157575b600080fd5b341561007d57600080fd5b6100a9600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610190565b005b34156100b657600080fd5b6100be61022f565b005b34156100cb57600080fd5b6100d36103ab565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561012057600080fd5b610155600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506103d0565b005b341561016257600080fd5b61018e600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610509565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156101ec57600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561028d57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561034e57600080fd5b5af1151561035b57600080fd5b5050506040518051905014151561037157600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561042d57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15156104f457600080fd5b5af1151561050157600080fd5b505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561056457600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16148061060d5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b151561061857600080fd5b61062181610624565b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561067f57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156106bb57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a723058202f4fe665dcc4a3ad0fb0c7215d36c5e527b474663eba9f2fc1bbd41798ce202b0029a165627a7a72305820c3bfc505797f9cfadfe5d2b65a977934505dba6b7c0f22ccb1cf3bf0c68320a90029`

// DeployFuxPayBoxFactory deploys a new Ethereum contract, binding an instance of FuxPayBoxFactory to it.
func DeployFuxPayBoxFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address) (common.Address, *types.Transaction, *FuxPayBoxFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxPayBoxFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FuxPayBoxFactoryBin), backend, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FuxPayBoxFactory{FuxPayBoxFactoryCaller: FuxPayBoxFactoryCaller{contract: contract}, FuxPayBoxFactoryTransactor: FuxPayBoxFactoryTransactor{contract: contract}, FuxPayBoxFactoryFilterer: FuxPayBoxFactoryFilterer{contract: contract}}, nil
}

// FuxPayBoxFactory is an auto generated Go binding around an Ethereum contract.
type FuxPayBoxFactory struct {
	FuxPayBoxFactoryCaller     // Read-only binding to the contract
	FuxPayBoxFactoryTransactor // Write-only binding to the contract
	FuxPayBoxFactoryFilterer   // Log filterer for contract events
}

// FuxPayBoxFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxPayBoxFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxPayBoxFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxPayBoxFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxPayBoxFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FuxPayBoxFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxPayBoxFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuxPayBoxFactorySession struct {
	Contract     *FuxPayBoxFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuxPayBoxFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuxPayBoxFactoryCallerSession struct {
	Contract *FuxPayBoxFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FuxPayBoxFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuxPayBoxFactoryTransactorSession struct {
	Contract     *FuxPayBoxFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FuxPayBoxFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuxPayBoxFactoryRaw struct {
	Contract *FuxPayBoxFactory // Generic contract binding to access the raw methods on
}

// FuxPayBoxFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuxPayBoxFactoryCallerRaw struct {
	Contract *FuxPayBoxFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FuxPayBoxFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuxPayBoxFactoryTransactorRaw struct {
	Contract *FuxPayBoxFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFuxPayBoxFactory creates a new instance of FuxPayBoxFactory, bound to a specific deployed contract.
func NewFuxPayBoxFactory(address common.Address, backend bind.ContractBackend) (*FuxPayBoxFactory, error) {
	contract, err := bindFuxPayBoxFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxPayBoxFactory{FuxPayBoxFactoryCaller: FuxPayBoxFactoryCaller{contract: contract}, FuxPayBoxFactoryTransactor: FuxPayBoxFactoryTransactor{contract: contract}, FuxPayBoxFactoryFilterer: FuxPayBoxFactoryFilterer{contract: contract}}, nil
}

// NewFuxPayBoxFactoryCaller creates a new read-only instance of FuxPayBoxFactory, bound to a specific deployed contract.
func NewFuxPayBoxFactoryCaller(address common.Address, caller bind.ContractCaller) (*FuxPayBoxFactoryCaller, error) {
	contract, err := bindFuxPayBoxFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FuxPayBoxFactoryCaller{contract: contract}, nil
}

// NewFuxPayBoxFactoryTransactor creates a new write-only instance of FuxPayBoxFactory, bound to a specific deployed contract.
func NewFuxPayBoxFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxPayBoxFactoryTransactor, error) {
	contract, err := bindFuxPayBoxFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FuxPayBoxFactoryTransactor{contract: contract}, nil
}

// NewFuxPayBoxFactoryFilterer creates a new log filterer instance of FuxPayBoxFactory, bound to a specific deployed contract.
func NewFuxPayBoxFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FuxPayBoxFactoryFilterer, error) {
	contract, err := bindFuxPayBoxFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FuxPayBoxFactoryFilterer{contract: contract}, nil
}

// bindFuxPayBoxFactory binds a generic wrapper to an already deployed contract.
func bindFuxPayBoxFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxPayBoxFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxPayBoxFactory *FuxPayBoxFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxPayBoxFactory.Contract.FuxPayBoxFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxPayBoxFactory *FuxPayBoxFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.FuxPayBoxFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxPayBoxFactory *FuxPayBoxFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.FuxPayBoxFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxPayBoxFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) Admin() (common.Address, error) {
	return _FuxPayBoxFactory.Contract.Admin(&_FuxPayBoxFactory.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) Admin() (common.Address, error) {
	return _FuxPayBoxFactory.Contract.Admin(&_FuxPayBoxFactory.CallOpts)
}

// FuxToken is a free data retrieval call binding the contract method 0xd91f5a6a.
//
// Solidity: function fuxToken() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) FuxToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "fuxToken")
	return *ret0, err
}

// FuxToken is a free data retrieval call binding the contract method 0xd91f5a6a.
//
// Solidity: function fuxToken() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) FuxToken() (common.Address, error) {
	return _FuxPayBoxFactory.Contract.FuxToken(&_FuxPayBoxFactory.CallOpts)
}

// FuxToken is a free data retrieval call binding the contract method 0xd91f5a6a.
//
// Solidity: function fuxToken() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) FuxToken() (common.Address, error) {
	return _FuxPayBoxFactory.Contract.FuxToken(&_FuxPayBoxFactory.CallOpts)
}

// GetBoxId is a free data retrieval call binding the contract method 0x946f0167.
//
// Solidity: function getBoxId(_addr address) constant returns(boxId uint256)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) GetBoxId(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "getBoxId", _addr)
	return *ret0, err
}

// GetBoxId is a free data retrieval call binding the contract method 0x946f0167.
//
// Solidity: function getBoxId(_addr address) constant returns(boxId uint256)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) GetBoxId(_addr common.Address) (*big.Int, error) {
	return _FuxPayBoxFactory.Contract.GetBoxId(&_FuxPayBoxFactory.CallOpts, _addr)
}

// GetBoxId is a free data retrieval call binding the contract method 0x946f0167.
//
// Solidity: function getBoxId(_addr address) constant returns(boxId uint256)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) GetBoxId(_addr common.Address) (*big.Int, error) {
	return _FuxPayBoxFactory.Contract.GetBoxId(&_FuxPayBoxFactory.CallOpts, _addr)
}

// GetFuxIssuer is a free data retrieval call binding the contract method 0x904c00c1.
//
// Solidity: function getFuxIssuer() constant returns(issuer address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) GetFuxIssuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "getFuxIssuer")
	return *ret0, err
}

// GetFuxIssuer is a free data retrieval call binding the contract method 0x904c00c1.
//
// Solidity: function getFuxIssuer() constant returns(issuer address)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) GetFuxIssuer() (common.Address, error) {
	return _FuxPayBoxFactory.Contract.GetFuxIssuer(&_FuxPayBoxFactory.CallOpts)
}

// GetFuxIssuer is a free data retrieval call binding the contract method 0x904c00c1.
//
// Solidity: function getFuxIssuer() constant returns(issuer address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) GetFuxIssuer() (common.Address, error) {
	return _FuxPayBoxFactory.Contract.GetFuxIssuer(&_FuxPayBoxFactory.CallOpts)
}

// GetPayBoxAddress is a free data retrieval call binding the contract method 0xfcfc43a7.
//
// Solidity: function getPayBoxAddress(_boxId uint256) constant returns(addr address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) GetPayBoxAddress(opts *bind.CallOpts, _boxId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "getPayBoxAddress", _boxId)
	return *ret0, err
}

// GetPayBoxAddress is a free data retrieval call binding the contract method 0xfcfc43a7.
//
// Solidity: function getPayBoxAddress(_boxId uint256) constant returns(addr address)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) GetPayBoxAddress(_boxId *big.Int) (common.Address, error) {
	return _FuxPayBoxFactory.Contract.GetPayBoxAddress(&_FuxPayBoxFactory.CallOpts, _boxId)
}

// GetPayBoxAddress is a free data retrieval call binding the contract method 0xfcfc43a7.
//
// Solidity: function getPayBoxAddress(_boxId uint256) constant returns(addr address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) GetPayBoxAddress(_boxId *big.Int) (common.Address, error) {
	return _FuxPayBoxFactory.Contract.GetPayBoxAddress(&_FuxPayBoxFactory.CallOpts, _boxId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) Owner() (common.Address, error) {
	return _FuxPayBoxFactory.Contract.Owner(&_FuxPayBoxFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) Owner() (common.Address, error) {
	return _FuxPayBoxFactory.Contract.Owner(&_FuxPayBoxFactory.CallOpts)
}

// CloseBox is a paid mutator transaction binding the contract method 0x606c5632.
//
// Solidity: function closeBox(_boxAddr address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactor) CloseBox(opts *bind.TransactOpts, _boxAddr common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.contract.Transact(opts, "closeBox", _boxAddr)
}

// CloseBox is a paid mutator transaction binding the contract method 0x606c5632.
//
// Solidity: function closeBox(_boxAddr address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) CloseBox(_boxAddr common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.CloseBox(&_FuxPayBoxFactory.TransactOpts, _boxAddr)
}

// CloseBox is a paid mutator transaction binding the contract method 0x606c5632.
//
// Solidity: function closeBox(_boxAddr address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactorSession) CloseBox(_boxAddr common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.CloseBox(&_FuxPayBoxFactory.TransactOpts, _boxAddr)
}

// CreatePayBox is a paid mutator transaction binding the contract method 0x4c982e57.
//
// Solidity: function createPayBox(_boxId uint256) returns(addr address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactor) CreatePayBox(opts *bind.TransactOpts, _boxId *big.Int) (*types.Transaction, error) {
	return _FuxPayBoxFactory.contract.Transact(opts, "createPayBox", _boxId)
}

// CreatePayBox is a paid mutator transaction binding the contract method 0x4c982e57.
//
// Solidity: function createPayBox(_boxId uint256) returns(addr address)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) CreatePayBox(_boxId *big.Int) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.CreatePayBox(&_FuxPayBoxFactory.TransactOpts, _boxId)
}

// CreatePayBox is a paid mutator transaction binding the contract method 0x4c982e57.
//
// Solidity: function createPayBox(_boxId uint256) returns(addr address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactorSession) CreatePayBox(_boxId *big.Int) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.CreatePayBox(&_FuxPayBoxFactory.TransactOpts, _boxId)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_addr address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactor) SetAdmin(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.contract.Transact(opts, "setAdmin", _addr)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_addr address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) SetAdmin(_addr common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.SetAdmin(&_FuxPayBoxFactory.TransactOpts, _addr)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(_addr address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactorSession) SetAdmin(_addr common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.SetAdmin(&_FuxPayBoxFactory.TransactOpts, _addr)
}

// SetFuxIssuer is a paid mutator transaction binding the contract method 0xc0b43a45.
//
// Solidity: function setFuxIssuer(_fuxIssuer address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactor) SetFuxIssuer(opts *bind.TransactOpts, _fuxIssuer common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.contract.Transact(opts, "setFuxIssuer", _fuxIssuer)
}

// SetFuxIssuer is a paid mutator transaction binding the contract method 0xc0b43a45.
//
// Solidity: function setFuxIssuer(_fuxIssuer address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) SetFuxIssuer(_fuxIssuer common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.SetFuxIssuer(&_FuxPayBoxFactory.TransactOpts, _fuxIssuer)
}

// SetFuxIssuer is a paid mutator transaction binding the contract method 0xc0b43a45.
//
// Solidity: function setFuxIssuer(_fuxIssuer address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactorSession) SetFuxIssuer(_fuxIssuer common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.SetFuxIssuer(&_FuxPayBoxFactory.TransactOpts, _fuxIssuer)
}

// SetFuxToken is a paid mutator transaction binding the contract method 0x28baafe1.
//
// Solidity: function setFuxToken(_addr address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactor) SetFuxToken(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.contract.Transact(opts, "setFuxToken", _addr)
}

// SetFuxToken is a paid mutator transaction binding the contract method 0x28baafe1.
//
// Solidity: function setFuxToken(_addr address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) SetFuxToken(_addr common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.SetFuxToken(&_FuxPayBoxFactory.TransactOpts, _addr)
}

// SetFuxToken is a paid mutator transaction binding the contract method 0x28baafe1.
//
// Solidity: function setFuxToken(_addr address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactorSession) SetFuxToken(_addr common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.SetFuxToken(&_FuxPayBoxFactory.TransactOpts, _addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.TransferOwnership(&_FuxPayBoxFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.TransferOwnership(&_FuxPayBoxFactory.TransactOpts, newOwner)
}

// FuxPayBoxFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FuxPayBoxFactory contract.
type FuxPayBoxFactoryOwnershipTransferredIterator struct {
	Event *FuxPayBoxFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FuxPayBoxFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxPayBoxFactoryOwnershipTransferred)
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
		it.Event = new(FuxPayBoxFactoryOwnershipTransferred)
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
func (it *FuxPayBoxFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxPayBoxFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxPayBoxFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the FuxPayBoxFactory contract.
type FuxPayBoxFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FuxPayBoxFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FuxPayBoxFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FuxPayBoxFactoryOwnershipTransferredIterator{contract: _FuxPayBoxFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FuxPayBoxFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FuxPayBoxFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxPayBoxFactoryOwnershipTransferred)
				if err := _FuxPayBoxFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
