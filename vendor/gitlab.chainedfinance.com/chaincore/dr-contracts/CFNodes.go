// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dr_contracts

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

// CFNodesABI is the input ABI used to generate the binding from.
const CFNodesABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isUsingRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sysRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CONFIG\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_isInDRGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CFRBAC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRouterAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_router\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"otherNodeInfo\",\"type\":\"string\"}],\"name\":\"AddNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"UpdateNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeName\",\"type\":\"string\"}],\"name\":\"DeleteNode\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeName\",\"type\":\"string\"},{\"name\":\"_publicKey\",\"type\":\"string\"},{\"name\":\"_otherNodeInfo\",\"type\":\"string\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeName\",\"type\":\"string\"}],\"name\":\"isNodeExisted\",\"outputs\":[{\"name\":\"existed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeName\",\"type\":\"string\"}],\"name\":\"deleteNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeName\",\"type\":\"string\"}],\"name\":\"getNodeKey\",\"outputs\":[{\"name\":\"publicKey\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeCount\",\"outputs\":[{\"name\":\"totalNodes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeName\",\"type\":\"string\"},{\"name\":\"_publicKey\",\"type\":\"string\"}],\"name\":\"updateNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CFNodes is an auto generated Go binding around an Ethereum contract.
type CFNodes struct {
	CFNodesCaller     // Read-only binding to the contract
	CFNodesTransactor // Write-only binding to the contract
	CFNodesFilterer   // Log filterer for contract events
}

// CFNodesCaller is an auto generated read-only Go binding around an Ethereum contract.
type CFNodesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CFNodesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CFNodesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CFNodesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CFNodesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CFNodesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CFNodesSession struct {
	Contract     *CFNodes          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CFNodesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CFNodesCallerSession struct {
	Contract *CFNodesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CFNodesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CFNodesTransactorSession struct {
	Contract     *CFNodesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CFNodesRaw is an auto generated low-level Go binding around an Ethereum contract.
type CFNodesRaw struct {
	Contract *CFNodes // Generic contract binding to access the raw methods on
}

// CFNodesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CFNodesCallerRaw struct {
	Contract *CFNodesCaller // Generic read-only contract binding to access the raw methods on
}

// CFNodesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CFNodesTransactorRaw struct {
	Contract *CFNodesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCFNodes creates a new instance of CFNodes, bound to a specific deployed contract.
func NewCFNodes(address common.Address, backend bind.ContractBackend) (*CFNodes, error) {
	contract, err := bindCFNodes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CFNodes{CFNodesCaller: CFNodesCaller{contract: contract}, CFNodesTransactor: CFNodesTransactor{contract: contract}, CFNodesFilterer: CFNodesFilterer{contract: contract}}, nil
}

// NewCFNodesCaller creates a new read-only instance of CFNodes, bound to a specific deployed contract.
func NewCFNodesCaller(address common.Address, caller bind.ContractCaller) (*CFNodesCaller, error) {
	contract, err := bindCFNodes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CFNodesCaller{contract: contract}, nil
}

// NewCFNodesTransactor creates a new write-only instance of CFNodes, bound to a specific deployed contract.
func NewCFNodesTransactor(address common.Address, transactor bind.ContractTransactor) (*CFNodesTransactor, error) {
	contract, err := bindCFNodes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CFNodesTransactor{contract: contract}, nil
}

// NewCFNodesFilterer creates a new log filterer instance of CFNodes, bound to a specific deployed contract.
func NewCFNodesFilterer(address common.Address, filterer bind.ContractFilterer) (*CFNodesFilterer, error) {
	contract, err := bindCFNodes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CFNodesFilterer{contract: contract}, nil
}

// bindCFNodes binds a generic wrapper to an already deployed contract.
func bindCFNodes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CFNodesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CFNodes *CFNodesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CFNodes.Contract.CFNodesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CFNodes *CFNodesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CFNodes.Contract.CFNodesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CFNodes *CFNodesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CFNodes.Contract.CFNodesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CFNodes *CFNodesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CFNodes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CFNodes *CFNodesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CFNodes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CFNodes *CFNodesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CFNodes.Contract.contract.Transact(opts, method, params...)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_CFNodes *CFNodesCaller) ROUTECFRBAC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CFNodes.contract.Call(opts, out, "ROUTE_CFRBAC")
	return *ret0, err
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_CFNodes *CFNodesSession) ROUTECFRBAC() (string, error) {
	return _CFNodes.Contract.ROUTECFRBAC(&_CFNodes.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_CFNodes *CFNodesCallerSession) ROUTECFRBAC() (string, error) {
	return _CFNodes.Contract.ROUTECFRBAC(&_CFNodes.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_CFNodes *CFNodesCaller) ROUTECONFIG(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CFNodes.contract.Call(opts, out, "ROUTE_CONFIG")
	return *ret0, err
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_CFNodes *CFNodesSession) ROUTECONFIG() (string, error) {
	return _CFNodes.Contract.ROUTECONFIG(&_CFNodes.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_CFNodes *CFNodesCallerSession) ROUTECONFIG() (string, error) {
	return _CFNodes.Contract.ROUTECONFIG(&_CFNodes.CallOpts)
}

// IsInDRGroup is a free data retrieval call binding the contract method 0xacd11b60.
//
// Solidity: function _isInDRGroup(_addr address) constant returns(bool)
func (_CFNodes *CFNodesCaller) IsInDRGroup(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CFNodes.contract.Call(opts, out, "_isInDRGroup", _addr)
	return *ret0, err
}

// IsInDRGroup is a free data retrieval call binding the contract method 0xacd11b60.
//
// Solidity: function _isInDRGroup(_addr address) constant returns(bool)
func (_CFNodes *CFNodesSession) IsInDRGroup(_addr common.Address) (bool, error) {
	return _CFNodes.Contract.IsInDRGroup(&_CFNodes.CallOpts, _addr)
}

// IsInDRGroup is a free data retrieval call binding the contract method 0xacd11b60.
//
// Solidity: function _isInDRGroup(_addr address) constant returns(bool)
func (_CFNodes *CFNodesCallerSession) IsInDRGroup(_addr common.Address) (bool, error) {
	return _CFNodes.Contract.IsInDRGroup(&_CFNodes.CallOpts, _addr)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() constant returns(totalNodes uint256)
func (_CFNodes *CFNodesCaller) GetNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CFNodes.contract.Call(opts, out, "getNodeCount")
	return *ret0, err
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() constant returns(totalNodes uint256)
func (_CFNodes *CFNodesSession) GetNodeCount() (*big.Int, error) {
	return _CFNodes.Contract.GetNodeCount(&_CFNodes.CallOpts)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() constant returns(totalNodes uint256)
func (_CFNodes *CFNodesCallerSession) GetNodeCount() (*big.Int, error) {
	return _CFNodes.Contract.GetNodeCount(&_CFNodes.CallOpts)
}

// GetNodeKey is a free data retrieval call binding the contract method 0x93d5badf.
//
// Solidity: function getNodeKey(_nodeName string) constant returns(publicKey string)
func (_CFNodes *CFNodesCaller) GetNodeKey(opts *bind.CallOpts, _nodeName string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CFNodes.contract.Call(opts, out, "getNodeKey", _nodeName)
	return *ret0, err
}

// GetNodeKey is a free data retrieval call binding the contract method 0x93d5badf.
//
// Solidity: function getNodeKey(_nodeName string) constant returns(publicKey string)
func (_CFNodes *CFNodesSession) GetNodeKey(_nodeName string) (string, error) {
	return _CFNodes.Contract.GetNodeKey(&_CFNodes.CallOpts, _nodeName)
}

// GetNodeKey is a free data retrieval call binding the contract method 0x93d5badf.
//
// Solidity: function getNodeKey(_nodeName string) constant returns(publicKey string)
func (_CFNodes *CFNodesCallerSession) GetNodeKey(_nodeName string) (string, error) {
	return _CFNodes.Contract.GetNodeKey(&_CFNodes.CallOpts, _nodeName)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_CFNodes *CFNodesCaller) GetRouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CFNodes.contract.Call(opts, out, "getRouterAddress")
	return *ret0, err
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_CFNodes *CFNodesSession) GetRouterAddress() (common.Address, error) {
	return _CFNodes.Contract.GetRouterAddress(&_CFNodes.CallOpts)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_CFNodes *CFNodesCallerSession) GetRouterAddress() (common.Address, error) {
	return _CFNodes.Contract.GetRouterAddress(&_CFNodes.CallOpts)
}

// IsNodeExisted is a free data retrieval call binding the contract method 0xf99101f5.
//
// Solidity: function isNodeExisted(_nodeName string) constant returns(existed bool)
func (_CFNodes *CFNodesCaller) IsNodeExisted(opts *bind.CallOpts, _nodeName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CFNodes.contract.Call(opts, out, "isNodeExisted", _nodeName)
	return *ret0, err
}

// IsNodeExisted is a free data retrieval call binding the contract method 0xf99101f5.
//
// Solidity: function isNodeExisted(_nodeName string) constant returns(existed bool)
func (_CFNodes *CFNodesSession) IsNodeExisted(_nodeName string) (bool, error) {
	return _CFNodes.Contract.IsNodeExisted(&_CFNodes.CallOpts, _nodeName)
}

// IsNodeExisted is a free data retrieval call binding the contract method 0xf99101f5.
//
// Solidity: function isNodeExisted(_nodeName string) constant returns(existed bool)
func (_CFNodes *CFNodesCallerSession) IsNodeExisted(_nodeName string) (bool, error) {
	return _CFNodes.Contract.IsNodeExisted(&_CFNodes.CallOpts, _nodeName)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_CFNodes *CFNodesCaller) IsUsingRouter(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CFNodes.contract.Call(opts, out, "isUsingRouter")
	return *ret0, err
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_CFNodes *CFNodesSession) IsUsingRouter() (bool, error) {
	return _CFNodes.Contract.IsUsingRouter(&_CFNodes.CallOpts)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_CFNodes *CFNodesCallerSession) IsUsingRouter() (bool, error) {
	return _CFNodes.Contract.IsUsingRouter(&_CFNodes.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_CFNodes *CFNodesCaller) SysRouter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CFNodes.contract.Call(opts, out, "sysRouter")
	return *ret0, err
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_CFNodes *CFNodesSession) SysRouter() (common.Address, error) {
	return _CFNodes.Contract.SysRouter(&_CFNodes.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_CFNodes *CFNodesCallerSession) SysRouter() (common.Address, error) {
	return _CFNodes.Contract.SysRouter(&_CFNodes.CallOpts)
}

// AddNode is a paid mutator transaction binding the contract method 0x03f1bc4f.
//
// Solidity: function addNode(_nodeName string, _publicKey string, _otherNodeInfo string) returns()
func (_CFNodes *CFNodesTransactor) AddNode(opts *bind.TransactOpts, _nodeName string, _publicKey string, _otherNodeInfo string) (*types.Transaction, error) {
	return _CFNodes.contract.Transact(opts, "addNode", _nodeName, _publicKey, _otherNodeInfo)
}

// AddNode is a paid mutator transaction binding the contract method 0x03f1bc4f.
//
// Solidity: function addNode(_nodeName string, _publicKey string, _otherNodeInfo string) returns()
func (_CFNodes *CFNodesSession) AddNode(_nodeName string, _publicKey string, _otherNodeInfo string) (*types.Transaction, error) {
	return _CFNodes.Contract.AddNode(&_CFNodes.TransactOpts, _nodeName, _publicKey, _otherNodeInfo)
}

// AddNode is a paid mutator transaction binding the contract method 0x03f1bc4f.
//
// Solidity: function addNode(_nodeName string, _publicKey string, _otherNodeInfo string) returns()
func (_CFNodes *CFNodesTransactorSession) AddNode(_nodeName string, _publicKey string, _otherNodeInfo string) (*types.Transaction, error) {
	return _CFNodes.Contract.AddNode(&_CFNodes.TransactOpts, _nodeName, _publicKey, _otherNodeInfo)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(_nodeName string) returns()
func (_CFNodes *CFNodesTransactor) DeleteNode(opts *bind.TransactOpts, _nodeName string) (*types.Transaction, error) {
	return _CFNodes.contract.Transact(opts, "deleteNode", _nodeName)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(_nodeName string) returns()
func (_CFNodes *CFNodesSession) DeleteNode(_nodeName string) (*types.Transaction, error) {
	return _CFNodes.Contract.DeleteNode(&_CFNodes.TransactOpts, _nodeName)
}

// DeleteNode is a paid mutator transaction binding the contract method 0xa0c15b77.
//
// Solidity: function deleteNode(_nodeName string) returns()
func (_CFNodes *CFNodesTransactorSession) DeleteNode(_nodeName string) (*types.Transaction, error) {
	return _CFNodes.Contract.DeleteNode(&_CFNodes.TransactOpts, _nodeName)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x517fb203.
//
// Solidity: function updateNode(_nodeName string, _publicKey string) returns()
func (_CFNodes *CFNodesTransactor) UpdateNode(opts *bind.TransactOpts, _nodeName string, _publicKey string) (*types.Transaction, error) {
	return _CFNodes.contract.Transact(opts, "updateNode", _nodeName, _publicKey)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x517fb203.
//
// Solidity: function updateNode(_nodeName string, _publicKey string) returns()
func (_CFNodes *CFNodesSession) UpdateNode(_nodeName string, _publicKey string) (*types.Transaction, error) {
	return _CFNodes.Contract.UpdateNode(&_CFNodes.TransactOpts, _nodeName, _publicKey)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x517fb203.
//
// Solidity: function updateNode(_nodeName string, _publicKey string) returns()
func (_CFNodes *CFNodesTransactorSession) UpdateNode(_nodeName string, _publicKey string) (*types.Transaction, error) {
	return _CFNodes.Contract.UpdateNode(&_CFNodes.TransactOpts, _nodeName, _publicKey)
}

// CFNodesAddNodeIterator is returned from FilterAddNode and is used to iterate over the raw logs and unpacked data for AddNode events raised by the CFNodes contract.
type CFNodesAddNodeIterator struct {
	Event *CFNodesAddNode // Event containing the contract specifics and raw log

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
func (it *CFNodesAddNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CFNodesAddNode)
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
		it.Event = new(CFNodesAddNode)
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
func (it *CFNodesAddNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CFNodesAddNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CFNodesAddNode represents a AddNode event raised by the CFNodes contract.
type CFNodesAddNode struct {
	NodeName      string
	PublicKey     string
	OtherNodeInfo string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddNode is a free log retrieval operation binding the contract event 0x92be66c0964e32dbb66cfb7c427ebd70418b6786b3d95b18ef07611a349edacd.
//
// Solidity: event AddNode(nodeName string, publicKey string, otherNodeInfo string)
func (_CFNodes *CFNodesFilterer) FilterAddNode(opts *bind.FilterOpts) (*CFNodesAddNodeIterator, error) {

	logs, sub, err := _CFNodes.contract.FilterLogs(opts, "AddNode")
	if err != nil {
		return nil, err
	}
	return &CFNodesAddNodeIterator{contract: _CFNodes.contract, event: "AddNode", logs: logs, sub: sub}, nil
}

// WatchAddNode is a free log subscription operation binding the contract event 0x92be66c0964e32dbb66cfb7c427ebd70418b6786b3d95b18ef07611a349edacd.
//
// Solidity: event AddNode(nodeName string, publicKey string, otherNodeInfo string)
func (_CFNodes *CFNodesFilterer) WatchAddNode(opts *bind.WatchOpts, sink chan<- *CFNodesAddNode) (event.Subscription, error) {

	logs, sub, err := _CFNodes.contract.WatchLogs(opts, "AddNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CFNodesAddNode)
				if err := _CFNodes.contract.UnpackLog(event, "AddNode", log); err != nil {
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

// CFNodesDeleteNodeIterator is returned from FilterDeleteNode and is used to iterate over the raw logs and unpacked data for DeleteNode events raised by the CFNodes contract.
type CFNodesDeleteNodeIterator struct {
	Event *CFNodesDeleteNode // Event containing the contract specifics and raw log

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
func (it *CFNodesDeleteNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CFNodesDeleteNode)
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
		it.Event = new(CFNodesDeleteNode)
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
func (it *CFNodesDeleteNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CFNodesDeleteNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CFNodesDeleteNode represents a DeleteNode event raised by the CFNodes contract.
type CFNodesDeleteNode struct {
	NodeName string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteNode is a free log retrieval operation binding the contract event 0xba443e89754ccbbfbdc44ab21d223373615fd710d0b11376d860814b15e7cd0a.
//
// Solidity: event DeleteNode(nodeName string)
func (_CFNodes *CFNodesFilterer) FilterDeleteNode(opts *bind.FilterOpts) (*CFNodesDeleteNodeIterator, error) {

	logs, sub, err := _CFNodes.contract.FilterLogs(opts, "DeleteNode")
	if err != nil {
		return nil, err
	}
	return &CFNodesDeleteNodeIterator{contract: _CFNodes.contract, event: "DeleteNode", logs: logs, sub: sub}, nil
}

// WatchDeleteNode is a free log subscription operation binding the contract event 0xba443e89754ccbbfbdc44ab21d223373615fd710d0b11376d860814b15e7cd0a.
//
// Solidity: event DeleteNode(nodeName string)
func (_CFNodes *CFNodesFilterer) WatchDeleteNode(opts *bind.WatchOpts, sink chan<- *CFNodesDeleteNode) (event.Subscription, error) {

	logs, sub, err := _CFNodes.contract.WatchLogs(opts, "DeleteNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CFNodesDeleteNode)
				if err := _CFNodes.contract.UnpackLog(event, "DeleteNode", log); err != nil {
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

// CFNodesUpdateNodeIterator is returned from FilterUpdateNode and is used to iterate over the raw logs and unpacked data for UpdateNode events raised by the CFNodes contract.
type CFNodesUpdateNodeIterator struct {
	Event *CFNodesUpdateNode // Event containing the contract specifics and raw log

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
func (it *CFNodesUpdateNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CFNodesUpdateNode)
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
		it.Event = new(CFNodesUpdateNode)
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
func (it *CFNodesUpdateNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CFNodesUpdateNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CFNodesUpdateNode represents a UpdateNode event raised by the CFNodes contract.
type CFNodesUpdateNode struct {
	NodeName  string
	PublicKey string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateNode is a free log retrieval operation binding the contract event 0xddd749469df4ef4f7716c20ad7266424f3217042b222980d3f2b62655124a76a.
//
// Solidity: event UpdateNode(nodeName string, publicKey string)
func (_CFNodes *CFNodesFilterer) FilterUpdateNode(opts *bind.FilterOpts) (*CFNodesUpdateNodeIterator, error) {

	logs, sub, err := _CFNodes.contract.FilterLogs(opts, "UpdateNode")
	if err != nil {
		return nil, err
	}
	return &CFNodesUpdateNodeIterator{contract: _CFNodes.contract, event: "UpdateNode", logs: logs, sub: sub}, nil
}

// WatchUpdateNode is a free log subscription operation binding the contract event 0xddd749469df4ef4f7716c20ad7266424f3217042b222980d3f2b62655124a76a.
//
// Solidity: event UpdateNode(nodeName string, publicKey string)
func (_CFNodes *CFNodesFilterer) WatchUpdateNode(opts *bind.WatchOpts, sink chan<- *CFNodesUpdateNode) (event.Subscription, error) {

	logs, sub, err := _CFNodes.contract.WatchLogs(opts, "UpdateNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CFNodesUpdateNode)
				if err := _CFNodes.contract.UnpackLog(event, "UpdateNode", log); err != nil {
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
