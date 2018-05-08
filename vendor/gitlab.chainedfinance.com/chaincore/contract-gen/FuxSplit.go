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

// FuxSplitABI is the input ABI used to generate the binding from.
const FuxSplitABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isUsingRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sysRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CONFIG\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_isFuxHub\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CFRBAC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRouterAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_HUB\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_isInFuxGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_idKids\",\"type\":\"uint256[2]\"},{\"name\":\"_coinKids\",\"type\":\"uint256[2]\"},{\"name\":\"_stateKids\",\"type\":\"uint256[2]\"}],\"name\":\"splitFux\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FuxSplitBin is the compiled bytecode used for deploying new contracts.
const FuxSplitBin = `0x60806040523480156200001157600080fd5b506040516020806200191783398101806040528101908080519060200190929190505050620000846040805190810160405280600c81526020017f726f75746572436f6e6669670000000000000000000000000000000000000000815250620000f4640100000000026401000000009004565b620000d36040805190810160405280600981526020017f467578436f6e6669670000000000000000000000000000000000000000000000815250620000f4640100000000026401000000009004565b620000ed8162000110640100000000026401000000009004565b506200023f565b80600190805190602001906200010c92919062000190565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156200014d57600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001d357805160ff191683800117855562000204565b8280016001018555821562000204579182015b8281111562000203578251825591602001919060010190620001e6565b5b50905062000213919062000217565b5090565b6200023c91905b80821115620002385760008160009055506001016200021e565b5090565b90565b6116c8806200024f6000396000f3006080604052600436106100af576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630d20aa48146100b45780635209e984146100e357806366b9852b1461019d5780636ed56bce146101f4578063ade6172414610284578063b8d43d27146102df578063c427d38d1461036f578063d54f7d5e146103ff578063d6b419fb14610456578063d89e6c4f146104e6578063f0b15e8214610576575b600080fd5b3480156100c057600080fd5b506100c96105d1565b604051808215151515815260200191505060405180910390f35b3480156100ef57600080fd5b5061019b600480360381019080803590602001909291908060400190600280602002604051908101604052809291908260026020028082843782019150505050509192919290806040019060028060200260405190810160405280929190826002602002808284378201915050505050919291929080604001906002806020026040519081016040528092919082600260200280828437820191505050505091929192905050506105da565b005b3480156101a957600080fd5b506101b2610e82565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561020057600080fd5b50610209610ea7565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561024957808201518184015260208101905061022e565b50505050905090810190601f1680156102765780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561029057600080fd5b506102c5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ee0565b604051808215151515815260200191505060405180910390f35b3480156102eb57600080fd5b506102f4611068565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610334578082015181840152602081019050610319565b50505050905090810190601f1680156103615780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561037b57600080fd5b506103846110a1565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103c45780820151818401526020810190506103a9565b50505050905090810190601f1680156103f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561040b57600080fd5b506104146110da565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561046257600080fd5b5061046b611103565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104ab578082015181840152602081019050610490565b50505050905090810190601f1680156104d85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104f257600080fd5b506104fb61113c565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561053b578082015181840152602081019050610520565b50505050905090810190601f1680156105685780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561058257600080fd5b506105b7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611175565b604051808215151515815260200191505060405180910390f35b60006001905090565b600080600080600080896000806106256040805190810160405280600e81526020017f467578455243373231546f6b656e0000000000000000000000000000000000008152506112fd565b91508173ffffffffffffffffffffffffffffffffffffffff16636352211e846040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561069657600080fd5b505af11580156106aa573d6000803e3d6000fd5b505050506040513d60208110156106c057600080fd5b810190808051906020019092919050505090503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161480610712575061071133611175565b5b80610722575061072133611436565b5b151561072d57600080fd5b61076b6040805190810160405280600e81526020017f467578455243373231546f6b656e0000000000000000000000000000000000008152506112fd565b98506107ab6040805190810160405280600a81526020017f46757853746f72616765000000000000000000000000000000000000000000008152506112fd565b97508873ffffffffffffffffffffffffffffffffffffffff16636352211e8e6040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561081c57600080fd5b505af1158015610830573d6000803e3d6000fd5b505050506040513d602081101561084657600080fd5b810190808051906020019092919050505096508773ffffffffffffffffffffffffffffffffffffffff16631c61984b8e6040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050606060405180830381600087803b1580156108c857600080fd5b505af11580156108dc573d6000803e3d6000fd5b505050506040513d60608110156108f257600080fd5b810190808051906020019092919080519060200190929190805190602001909291905050508097508198508296505050508361095d8c600160028110151561093657fe5b60200201518d600060028110151561094a57fe5b602002015161151c90919063ffffffff16565b14151561096957600080fd5b8873ffffffffffffffffffffffffffffffffffffffff16639dc29fac888f6040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610a0c57600080fd5b505af1158015610a20573d6000803e3d6000fd5b505050508773ffffffffffffffffffffffffffffffffffffffff1663e234f61d8e6040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b158015610a9357600080fd5b505af1158015610aa7573d6000803e3d6000fd5b50505050610ab433611175565b158015610ac75750610ac533611436565b155b15610aff57848a6000600281101515610adc57fe5b602002018181525050848a6001600281101515610af557fe5b6020020181815250505b8773ffffffffffffffffffffffffffffffffffffffff166386b27f4d8d6000600281101515610b2a57fe5b60200201518d6000600281101515610b3e57fe5b6020020151898e6000600281101515610b5357fe5b60200201516040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808581526020018481526020018381526020018281526020018060200182810382526000815260200160200195505050505050600060405180830381600087803b158015610bd357600080fd5b505af1158015610be7573d6000803e3d6000fd5b505050508773ffffffffffffffffffffffffffffffffffffffff166386b27f4d8d6001600281101515610c1657fe5b60200201518d6001600281101515610c2a57fe5b6020020151898e6001600281101515610c3f57fe5b60200201516040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808581526020018481526020018381526020018281526020018060200182810382526000815260200160200195505050505050600060405180830381600087803b158015610cbf57600080fd5b505af1158015610cd3573d6000803e3d6000fd5b505050508873ffffffffffffffffffffffffffffffffffffffff166340c10f19888e6000600281101515610d0357fe5b60200201516040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610d8d57600080fd5b505af1158015610da1573d6000803e3d6000fd5b505050508873ffffffffffffffffffffffffffffffffffffffff166340c10f19888e6001600281101515610dd157fe5b60200201516040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610e5b57600080fd5b505af1158015610e6f573d6000803e3d6000fd5b5050505050505050505050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600c81526020017f726f75746572436f6e666967000000000000000000000000000000000000000081525081565b6000610eea611538565b73ffffffffffffffffffffffffffffffffffffffff1663217fe6c6836040805190810160405280600681526020017f66757848756200000000000000000000000000000000000000000000000000008152506040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610fd9578082015181840152602081019050610fbe565b50505050905090810190601f1680156110065780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561102657600080fd5b505af115801561103a573d6000803e3d6000fd5b505050506040513d602081101561105057600080fd5b81019080805190602001909291905050509050919050565b6040805190810160405280600b81526020017f726f6c65436c757374657200000000000000000000000000000000000000000081525081565b6040805190810160405280600c81526020017f726f75746572434652424143000000000000000000000000000000000000000081525081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6040805190810160405280600681526020017f667578487562000000000000000000000000000000000000000000000000000081525081565b6040805190810160405280600c81526020017f726f6c6546757847726f7570000000000000000000000000000000000000000081525081565b600061117f611538565b73ffffffffffffffffffffffffffffffffffffffff1663217fe6c6836040805190810160405280600c81526020017f726f6c6546757847726f757000000000000000000000000000000000000000008152506040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561126e578082015181840152602081019050611253565b50505050905090810190601f16801561129b5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b1580156112bb57600080fd5b505af11580156112cf573d6000803e3d6000fd5b505050506040513d60208110156112e557600080fd5b81019080805190602001909291905050509050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c2511025836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156113a857808201518184015260208101905061138d565b50505050905090810190601f1680156113d55780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1580156113f457600080fd5b505af1158015611408573d6000803e3d6000fd5b505050506040513d602081101561141e57600080fd5b81019080805190602001909291905050509050919050565b6000611440611538565b73ffffffffffffffffffffffffffffffffffffffff166324d7806c836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156114da57600080fd5b505af11580156114ee573d6000803e3d6000fd5b505050506040513d602081101561150457600080fd5b81019080805190602001909291905050509050919050565b6000818301905082811015151561152f57fe5b80905092915050565b6000611542611547565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c251102560016040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200182810382528381815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561163c5780601f106116115761010080835404028352916020019161163c565b820191906000526020600020905b81548152906001019060200180831161161f57829003601f168201915b505092505050602060405180830381600087803b15801561165c57600080fd5b505af1158015611670573d6000803e3d6000fd5b505050506040513d602081101561168657600080fd5b81019080805190602001909291905050509050905600a165627a7a72305820717e2e7d445206b2e2720f0044d4c99cc76a8054475f1c118ea9665ae47a60f40029`

// DeployFuxSplit deploys a new Ethereum contract, binding an instance of FuxSplit to it.
func DeployFuxSplit(auth *bind.TransactOpts, backend bind.ContractBackend, _addr common.Address) (common.Address, *types.Transaction, *FuxSplit, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxSplitABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FuxSplitBin), backend, _addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FuxSplit{FuxSplitCaller: FuxSplitCaller{contract: contract}, FuxSplitTransactor: FuxSplitTransactor{contract: contract}, FuxSplitFilterer: FuxSplitFilterer{contract: contract}}, nil
}

// FuxSplit is an auto generated Go binding around an Ethereum contract.
type FuxSplit struct {
	FuxSplitCaller     // Read-only binding to the contract
	FuxSplitTransactor // Write-only binding to the contract
	FuxSplitFilterer   // Log filterer for contract events
}

// FuxSplitCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxSplitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxSplitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxSplitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxSplitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FuxSplitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxSplitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuxSplitSession struct {
	Contract     *FuxSplit         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuxSplitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuxSplitCallerSession struct {
	Contract *FuxSplitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FuxSplitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuxSplitTransactorSession struct {
	Contract     *FuxSplitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FuxSplitRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuxSplitRaw struct {
	Contract *FuxSplit // Generic contract binding to access the raw methods on
}

// FuxSplitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuxSplitCallerRaw struct {
	Contract *FuxSplitCaller // Generic read-only contract binding to access the raw methods on
}

// FuxSplitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuxSplitTransactorRaw struct {
	Contract *FuxSplitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFuxSplit creates a new instance of FuxSplit, bound to a specific deployed contract.
func NewFuxSplit(address common.Address, backend bind.ContractBackend) (*FuxSplit, error) {
	contract, err := bindFuxSplit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxSplit{FuxSplitCaller: FuxSplitCaller{contract: contract}, FuxSplitTransactor: FuxSplitTransactor{contract: contract}, FuxSplitFilterer: FuxSplitFilterer{contract: contract}}, nil
}

// NewFuxSplitCaller creates a new read-only instance of FuxSplit, bound to a specific deployed contract.
func NewFuxSplitCaller(address common.Address, caller bind.ContractCaller) (*FuxSplitCaller, error) {
	contract, err := bindFuxSplit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FuxSplitCaller{contract: contract}, nil
}

// NewFuxSplitTransactor creates a new write-only instance of FuxSplit, bound to a specific deployed contract.
func NewFuxSplitTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxSplitTransactor, error) {
	contract, err := bindFuxSplit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FuxSplitTransactor{contract: contract}, nil
}

// NewFuxSplitFilterer creates a new log filterer instance of FuxSplit, bound to a specific deployed contract.
func NewFuxSplitFilterer(address common.Address, filterer bind.ContractFilterer) (*FuxSplitFilterer, error) {
	contract, err := bindFuxSplit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FuxSplitFilterer{contract: contract}, nil
}

// bindFuxSplit binds a generic wrapper to an already deployed contract.
func bindFuxSplit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxSplitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxSplit *FuxSplitRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxSplit.Contract.FuxSplitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxSplit *FuxSplitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxSplit.Contract.FuxSplitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxSplit *FuxSplitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxSplit.Contract.FuxSplitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxSplit *FuxSplitCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxSplit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxSplit *FuxSplitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxSplit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxSplit *FuxSplitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxSplit.Contract.contract.Transact(opts, method, params...)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROLECLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROLECLUSTER() (string, error) {
	return _FuxSplit.Contract.ROLECLUSTER(&_FuxSplit.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROLECLUSTER() (string, error) {
	return _FuxSplit.Contract.ROLECLUSTER(&_FuxSplit.CallOpts)
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROLEFUXGROUP(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROLE_FUX_GROUP")
	return *ret0, err
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROLEFUXGROUP() (string, error) {
	return _FuxSplit.Contract.ROLEFUXGROUP(&_FuxSplit.CallOpts)
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROLEFUXGROUP() (string, error) {
	return _FuxSplit.Contract.ROLEFUXGROUP(&_FuxSplit.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROLEFUXHUB(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROLE_FUX_HUB")
	return *ret0, err
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROLEFUXHUB() (string, error) {
	return _FuxSplit.Contract.ROLEFUXHUB(&_FuxSplit.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROLEFUXHUB() (string, error) {
	return _FuxSplit.Contract.ROLEFUXHUB(&_FuxSplit.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROUTECFRBAC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROUTE_CFRBAC")
	return *ret0, err
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROUTECFRBAC() (string, error) {
	return _FuxSplit.Contract.ROUTECFRBAC(&_FuxSplit.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROUTECFRBAC() (string, error) {
	return _FuxSplit.Contract.ROUTECFRBAC(&_FuxSplit.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxSplit *FuxSplitCaller) ROUTECONFIG(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "ROUTE_CONFIG")
	return *ret0, err
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxSplit *FuxSplitSession) ROUTECONFIG() (string, error) {
	return _FuxSplit.Contract.ROUTECONFIG(&_FuxSplit.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxSplit *FuxSplitCallerSession) ROUTECONFIG() (string, error) {
	return _FuxSplit.Contract.ROUTECONFIG(&_FuxSplit.CallOpts)
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxSplit *FuxSplitCaller) IsFuxHub(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "_isFuxHub", _addr)
	return *ret0, err
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxSplit *FuxSplitSession) IsFuxHub(_addr common.Address) (bool, error) {
	return _FuxSplit.Contract.IsFuxHub(&_FuxSplit.CallOpts, _addr)
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxSplit *FuxSplitCallerSession) IsFuxHub(_addr common.Address) (bool, error) {
	return _FuxSplit.Contract.IsFuxHub(&_FuxSplit.CallOpts, _addr)
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxSplit *FuxSplitCaller) IsInFuxGroup(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "_isInFuxGroup", _addr)
	return *ret0, err
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxSplit *FuxSplitSession) IsInFuxGroup(_addr common.Address) (bool, error) {
	return _FuxSplit.Contract.IsInFuxGroup(&_FuxSplit.CallOpts, _addr)
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxSplit *FuxSplitCallerSession) IsInFuxGroup(_addr common.Address) (bool, error) {
	return _FuxSplit.Contract.IsInFuxGroup(&_FuxSplit.CallOpts, _addr)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxSplit *FuxSplitCaller) GetRouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "getRouterAddress")
	return *ret0, err
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxSplit *FuxSplitSession) GetRouterAddress() (common.Address, error) {
	return _FuxSplit.Contract.GetRouterAddress(&_FuxSplit.CallOpts)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxSplit *FuxSplitCallerSession) GetRouterAddress() (common.Address, error) {
	return _FuxSplit.Contract.GetRouterAddress(&_FuxSplit.CallOpts)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxSplit *FuxSplitCaller) IsUsingRouter(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "isUsingRouter")
	return *ret0, err
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxSplit *FuxSplitSession) IsUsingRouter() (bool, error) {
	return _FuxSplit.Contract.IsUsingRouter(&_FuxSplit.CallOpts)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxSplit *FuxSplitCallerSession) IsUsingRouter() (bool, error) {
	return _FuxSplit.Contract.IsUsingRouter(&_FuxSplit.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxSplit *FuxSplitCaller) SysRouter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxSplit.contract.Call(opts, out, "sysRouter")
	return *ret0, err
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxSplit *FuxSplitSession) SysRouter() (common.Address, error) {
	return _FuxSplit.Contract.SysRouter(&_FuxSplit.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxSplit *FuxSplitCallerSession) SysRouter() (common.Address, error) {
	return _FuxSplit.Contract.SysRouter(&_FuxSplit.CallOpts)
}

// SplitFux is a paid mutator transaction binding the contract method 0x5209e984.
//
// Solidity: function splitFux(_tokenId uint256, _idKids uint256[2], _coinKids uint256[2], _stateKids uint256[2]) returns()
func (_FuxSplit *FuxSplitTransactor) SplitFux(opts *bind.TransactOpts, _tokenId *big.Int, _idKids [2]*big.Int, _coinKids [2]*big.Int, _stateKids [2]*big.Int) (*types.Transaction, error) {
	return _FuxSplit.contract.Transact(opts, "splitFux", _tokenId, _idKids, _coinKids, _stateKids)
}

// SplitFux is a paid mutator transaction binding the contract method 0x5209e984.
//
// Solidity: function splitFux(_tokenId uint256, _idKids uint256[2], _coinKids uint256[2], _stateKids uint256[2]) returns()
func (_FuxSplit *FuxSplitSession) SplitFux(_tokenId *big.Int, _idKids [2]*big.Int, _coinKids [2]*big.Int, _stateKids [2]*big.Int) (*types.Transaction, error) {
	return _FuxSplit.Contract.SplitFux(&_FuxSplit.TransactOpts, _tokenId, _idKids, _coinKids, _stateKids)
}

// SplitFux is a paid mutator transaction binding the contract method 0x5209e984.
//
// Solidity: function splitFux(_tokenId uint256, _idKids uint256[2], _coinKids uint256[2], _stateKids uint256[2]) returns()
func (_FuxSplit *FuxSplitTransactorSession) SplitFux(_tokenId *big.Int, _idKids [2]*big.Int, _coinKids [2]*big.Int, _stateKids [2]*big.Int) (*types.Transaction, error) {
	return _FuxSplit.Contract.SplitFux(&_FuxSplit.TransactOpts, _tokenId, _idKids, _coinKids, _stateKids)
}
