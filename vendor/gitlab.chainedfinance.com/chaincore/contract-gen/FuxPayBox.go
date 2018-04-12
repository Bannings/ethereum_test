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
const FuxPayBoxABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"},{\"name\":\"_fuxToken\",\"type\":\"address\"},{\"name\":\"_factory\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"superFox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FuxPayBoxBin is the compiled bytecode used for deploying new contracts.
const FuxPayBoxBin = `0x6060604052341561000f57600080fd5b6040516060806108f783398101604052808051906020019091908051906020019091908051906020019091905050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506107a5806101526000396000f30060606040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633d8d3fd91461007257806341c0e1b5146100ab5780638da5cb5b146100c0578063a9059cbb14610115578063f2fde38b14610157575b600080fd5b341561007d57600080fd5b6100a9600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610190565b005b34156100b657600080fd5b6100be61022f565b005b34156100cb57600080fd5b6100d36103ab565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561012057600080fd5b610155600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506103d0565b005b341561016257600080fd5b61018e600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610509565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156101ec57600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561028d57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b151561034e57600080fd5b5af1151561035b57600080fd5b5050506040518051905014151561037157600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561042d57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15156104f457600080fd5b5af1151561050157600080fd5b505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561056457600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16148061060d5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b151561061857600080fd5b61062181610624565b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561067f57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156106bb57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a723058202f4fe665dcc4a3ad0fb0c7215d36c5e527b474663eba9f2fc1bbd41798ce202b0029`

// DeployFuxPayBox deploys a new Ethereum contract, binding an instance of FuxPayBox to it.
func DeployFuxPayBox(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _fuxToken common.Address, _factory common.Address) (common.Address, *types.Transaction, *FuxPayBox, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxPayBoxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FuxPayBoxBin), backend, _admin, _fuxToken, _factory)
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

// SuperFox is a paid mutator transaction binding the contract method 0x3d8d3fd9.
//
// Solidity: function superFox(_owner address) returns()
func (_FuxPayBox *FuxPayBoxTransactor) SuperFox(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _FuxPayBox.contract.Transact(opts, "superFox", _owner)
}

// SuperFox is a paid mutator transaction binding the contract method 0x3d8d3fd9.
//
// Solidity: function superFox(_owner address) returns()
func (_FuxPayBox *FuxPayBoxSession) SuperFox(_owner common.Address) (*types.Transaction, error) {
	return _FuxPayBox.Contract.SuperFox(&_FuxPayBox.TransactOpts, _owner)
}

// SuperFox is a paid mutator transaction binding the contract method 0x3d8d3fd9.
//
// Solidity: function superFox(_owner address) returns()
func (_FuxPayBox *FuxPayBoxTransactorSession) SuperFox(_owner common.Address) (*types.Transaction, error) {
	return _FuxPayBox.Contract.SuperFox(&_FuxPayBox.TransactOpts, _owner)
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
