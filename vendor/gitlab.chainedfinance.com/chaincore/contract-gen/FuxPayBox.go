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

// FuxPayBoxABI is the input ABI used to generate the binding from.
const FuxPayBoxABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_TRANSFER_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_BC_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_RBAC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_ISSUER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"staticAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sysProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_COIN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProxyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_TOKEN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isUsingProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_SPLIT\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_STORAGE\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_HUB\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_PAYBOX_FACTORY\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_ERC721_TOKEN\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_FUX_CONFIG\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FuxPayBoxBin is the compiled bytecode used for deploying new contracts.
const FuxPayBoxBin = `0x606060405260008060146101000a81548160ff021916908315150217905550341561002957600080fd5b604051602080611a0a8339810160405280805190602001909190505033600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506100a2816100a864010000000002611827176401000000009004565b50610128565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156100e457600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6118d3806101376000396000f30060606040526004361061013e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063071db5611461014357806309c4b0fa146101d15780630c7bba781461025f5780630daaf6e4146102ed5780632bc3a44c1461037b57806332061899146103a857806333e653e8146103fd57806341c0e1b51461048b57806343a73d9a146104a05780634aa03a35146104f55780634e71e0c81461058357806375ae7a8d14610598578063781bed44146105ed57806388f7748d1461061a5780638da5cb5b146106a8578063a9059cbb146106fd578063c0d193251461073f578063d6b419fb146107cd578063e30c39781461085b578063e941ed61146108b0578063f0b9e5ba1461093e578063f2fde38b14610a15578063f43a6d2414610a4e578063f5ac8fb914610adc575b600080fd5b341561014e57600080fd5b610156610b6a565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561019657808201518184015260208101905061017b565b50505050905090810190601f1680156101c35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101dc57600080fd5b6101e4610ba3565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610224578082015181840152602081019050610209565b50505050905090810190601f1680156102515780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561026a57600080fd5b610272610bdc565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102b2578082015181840152602081019050610297565b50505050905090810190601f1680156102df5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102f857600080fd5b610300610c15565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610340578082015181840152602081019050610325565b50505050905090810190601f16801561036d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561038657600080fd5b61038e610c4e565b604051808215151515815260200191505060405180910390f35b34156103b357600080fd5b6103bb610c61565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561040857600080fd5b610410610c87565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610450578082015181840152602081019050610435565b50505050905090810190601f16801561047d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561049657600080fd5b61049e610cc0565b005b34156104ab57600080fd5b6104b3610e38565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561050057600080fd5b610508610e62565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561054857808201518184015260208101905061052d565b50505050905090810190601f1680156105755780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561058e57600080fd5b610596610e9b565b005b34156105a357600080fd5b6105ab61110a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156105f857600080fd5b61060061112f565b604051808215151515815260200191505060405180910390f35b341561062557600080fd5b61062d611146565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561066d578082015181840152602081019050610652565b50505050905090810190601f16801561069a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156106b357600080fd5b6106bb61117f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561070857600080fd5b61073d600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506111a5565b005b341561074a57600080fd5b6107526112f2565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610792578082015181840152602081019050610777565b50505050905090810190601f1680156107bf5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156107d857600080fd5b6107e061132b565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610820578082015181840152602081019050610805565b50505050905090810190601f16801561084d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561086657600080fd5b61086e611364565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156108bb57600080fd5b6108c361138a565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156109035780820151818401526020810190506108e8565b50505050905090810190601f1680156109305780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561094957600080fd5b6109c1600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506113c3565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b3415610a2057600080fd5b610a4c600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506113f2565b005b3415610a5957600080fd5b610a61611422565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610aa1578082015181840152602081019050610a86565b50505050905090810190601f168015610ace5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3415610ae757600080fd5b610aef61145b565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610b2f578082015181840152602081019050610b14565b50505050905090810190601f168015610b5c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6040805190810160405280601381526020017f5472616e73666572526f6c65576974684675780000000000000000000000000081525081565b6040805190810160405280600a81526020017f667578426347726f75700000000000000000000000000000000000000000000081525081565b6040805190810160405280600b81526020017f524241435769746846757800000000000000000000000000000000000000000081525081565b6040805190810160405280600981526020017f667578497373756572000000000000000000000000000000000000000000000081525081565b600060149054906101000a900460ff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600781526020017f467578436f696e0000000000000000000000000000000000000000000000000081525081565b610cfe6040805190810160405280601081526020017f467578506179426f78466163746f727900000000000000000000000000000000815250611494565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610d3757600080fd5b6000610d416115a9565b73ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1515610dda57600080fd5b5af11515610de757600080fd5b50505060405180519050141515610dfd57600080fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6040805190810160405280600881526020017f467578546f6b656e00000000000000000000000000000000000000000000000081525081565b610ed96040805190810160405280601081526020017f467578506179426f78466163746f727900000000000000000000000000000000815250611494565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148015610f625750600073ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b80610fba5750600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515610fc557600080fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060149054906101000a900460ff1615905090565b6040805190810160405280600881526020017f46757853706c697400000000000000000000000000000000000000000000000081525081565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561120157600080fd5b6112096115a9565b73ffffffffffffffffffffffffffffffffffffffff1663beabacc83084846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15156112de57600080fd5b5af115156112eb57600080fd5b5050505050565b6040805190810160405280600a81526020017f46757853746f726167650000000000000000000000000000000000000000000081525081565b6040805190810160405280600681526020017f667578487562000000000000000000000000000000000000000000000000000081525081565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280601081526020017f467578506179426f78466163746f72790000000000000000000000000000000081525081565b600063f0b9e5ba7c01000000000000000000000000000000000000000000000000000000000290509392505050565b6113fb816115ee565b8061140b575061140a336115ee565b5b151561141657600080fd5b61141f81611787565b50565b6040805190810160405280600e81526020017f467578455243373231546f6b656e00000000000000000000000000000000000081525081565b6040805190810160405280600981526020017f467578436f6e666967000000000000000000000000000000000000000000000081525081565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c2511025836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611540578082015181840152602081019050611525565b50505050905090810190601f16801561156d5780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b151561158b57600080fd5b5af1151561159857600080fd5b505050604051805190509050919050565b60006115e96040805190810160405280600881526020017f467578546f6b656e000000000000000000000000000000000000000000000000815250611494565b905090565b600061162e6040805190810160405280600b81526020017f5242414357697468467578000000000000000000000000000000000000000000815250611494565b73ffffffffffffffffffffffffffffffffffffffff1663217fe6c6836040805190810160405280600681526020017f66757848756200000000000000000000000000000000000000000000000000008152506040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561171d578082015181840152602081019050611702565b50505050905090810190601f16801561174a5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b151561176957600080fd5b5af1151561177657600080fd5b505050604051805190509050919050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156117e357600080fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561186357600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a72305820965c50717d50ae3d0fc539832d0759939e700b2e4aecfb08f73edfb527776e9e0029`

// DeployFuxPayBox deploys a new Ethereum contract, binding an instance of FuxPayBox to it.
func DeployFuxPayBox(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address) (common.Address, *types.Transaction, *FuxPayBox, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxPayBoxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FuxPayBoxBin), backend, _proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FuxPayBox{FuxPayBoxCaller: FuxPayBoxCaller{contract: contract}, FuxPayBoxTransactor: FuxPayBoxTransactor{contract: contract}}, nil
}

// FuxPayBox is an auto generated Go binding around an Ethereum contract.
type FuxPayBox struct {
	FuxPayBoxCaller     // Read-only binding to the contract
	FuxPayBoxTransactor // Write-only binding to the contract
}

// FuxPayBoxCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxPayBoxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxPayBoxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxPayBoxTransactor struct {
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
	contract, err := bindFuxPayBox(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxPayBox{FuxPayBoxCaller: FuxPayBoxCaller{contract: contract}, FuxPayBoxTransactor: FuxPayBoxTransactor{contract: contract}}, nil
}

// NewFuxPayBoxCaller creates a new read-only instance of FuxPayBox, bound to a specific deployed contract.
func NewFuxPayBoxCaller(address common.Address, caller bind.ContractCaller) (*FuxPayBoxCaller, error) {
	contract, err := bindFuxPayBox(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &FuxPayBoxCaller{contract: contract}, nil
}

// NewFuxPayBoxTransactor creates a new write-only instance of FuxPayBox, bound to a specific deployed contract.
func NewFuxPayBoxTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxPayBoxTransactor, error) {
	contract, err := bindFuxPayBox(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &FuxPayBoxTransactor{contract: contract}, nil
}

// bindFuxPayBox binds a generic wrapper to an already deployed contract.
func bindFuxPayBox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxPayBoxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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

// ROLE_FUX_BC_GROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROLE_FUX_BC_GROUP(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROLE_FUX_BC_GROUP")
	return *ret0, err
}

// ROLE_FUX_BC_GROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROLE_FUX_BC_GROUP() (string, error) {
	return _FuxPayBox.Contract.ROLE_FUX_BC_GROUP(&_FuxPayBox.CallOpts)
}

// ROLE_FUX_BC_GROUP is a free data retrieval call binding the contract method 0x09c4b0fa.
//
// Solidity: function ROLE_FUX_BC_GROUP() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROLE_FUX_BC_GROUP() (string, error) {
	return _FuxPayBox.Contract.ROLE_FUX_BC_GROUP(&_FuxPayBox.CallOpts)
}

// ROLE_FUX_HUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROLE_FUX_HUB(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROLE_FUX_HUB")
	return *ret0, err
}

// ROLE_FUX_HUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROLE_FUX_HUB() (string, error) {
	return _FuxPayBox.Contract.ROLE_FUX_HUB(&_FuxPayBox.CallOpts)
}

// ROLE_FUX_HUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROLE_FUX_HUB() (string, error) {
	return _FuxPayBox.Contract.ROLE_FUX_HUB(&_FuxPayBox.CallOpts)
}

// ROLE_FUX_ISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROLE_FUX_ISSUER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROLE_FUX_ISSUER")
	return *ret0, err
}

// ROLE_FUX_ISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROLE_FUX_ISSUER() (string, error) {
	return _FuxPayBox.Contract.ROLE_FUX_ISSUER(&_FuxPayBox.CallOpts)
}

// ROLE_FUX_ISSUER is a free data retrieval call binding the contract method 0x0daaf6e4.
//
// Solidity: function ROLE_FUX_ISSUER() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROLE_FUX_ISSUER() (string, error) {
	return _FuxPayBox.Contract.ROLE_FUX_ISSUER(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_COIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROUTE_FUX_COIN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROUTE_FUX_COIN")
	return *ret0, err
}

// ROUTE_FUX_COIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROUTE_FUX_COIN() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_COIN(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_COIN is a free data retrieval call binding the contract method 0x33e653e8.
//
// Solidity: function ROUTE_FUX_COIN() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROUTE_FUX_COIN() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_COIN(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_CONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROUTE_FUX_CONFIG(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROUTE_FUX_CONFIG")
	return *ret0, err
}

// ROUTE_FUX_CONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROUTE_FUX_CONFIG() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_CONFIG(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_CONFIG is a free data retrieval call binding the contract method 0xf5ac8fb9.
//
// Solidity: function ROUTE_FUX_CONFIG() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROUTE_FUX_CONFIG() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_CONFIG(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_ERC721_TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROUTE_FUX_ERC721_TOKEN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROUTE_FUX_ERC721_TOKEN")
	return *ret0, err
}

// ROUTE_FUX_ERC721_TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROUTE_FUX_ERC721_TOKEN() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_ERC721_TOKEN(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_ERC721_TOKEN is a free data retrieval call binding the contract method 0xf43a6d24.
//
// Solidity: function ROUTE_FUX_ERC721_TOKEN() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROUTE_FUX_ERC721_TOKEN() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_ERC721_TOKEN(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_PAYBOX_FACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROUTE_FUX_PAYBOX_FACTORY(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROUTE_FUX_PAYBOX_FACTORY")
	return *ret0, err
}

// ROUTE_FUX_PAYBOX_FACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROUTE_FUX_PAYBOX_FACTORY() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_PAYBOX_FACTORY(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_PAYBOX_FACTORY is a free data retrieval call binding the contract method 0xe941ed61.
//
// Solidity: function ROUTE_FUX_PAYBOX_FACTORY() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROUTE_FUX_PAYBOX_FACTORY() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_PAYBOX_FACTORY(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_RBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROUTE_FUX_RBAC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROUTE_FUX_RBAC")
	return *ret0, err
}

// ROUTE_FUX_RBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROUTE_FUX_RBAC() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_RBAC(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_RBAC is a free data retrieval call binding the contract method 0x0c7bba78.
//
// Solidity: function ROUTE_FUX_RBAC() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROUTE_FUX_RBAC() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_RBAC(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_SPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROUTE_FUX_SPLIT(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROUTE_FUX_SPLIT")
	return *ret0, err
}

// ROUTE_FUX_SPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROUTE_FUX_SPLIT() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_SPLIT(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_SPLIT is a free data retrieval call binding the contract method 0x88f7748d.
//
// Solidity: function ROUTE_FUX_SPLIT() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROUTE_FUX_SPLIT() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_SPLIT(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_STORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROUTE_FUX_STORAGE(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROUTE_FUX_STORAGE")
	return *ret0, err
}

// ROUTE_FUX_STORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROUTE_FUX_STORAGE() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_STORAGE(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_STORAGE is a free data retrieval call binding the contract method 0xc0d19325.
//
// Solidity: function ROUTE_FUX_STORAGE() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROUTE_FUX_STORAGE() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_STORAGE(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_TOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROUTE_FUX_TOKEN(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROUTE_FUX_TOKEN")
	return *ret0, err
}

// ROUTE_FUX_TOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROUTE_FUX_TOKEN() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_TOKEN(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_TOKEN is a free data retrieval call binding the contract method 0x4aa03a35.
//
// Solidity: function ROUTE_FUX_TOKEN() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROUTE_FUX_TOKEN() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_TOKEN(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_TRANSFER_ROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxPayBox *FuxPayBoxCaller) ROUTE_FUX_TRANSFER_ROLE(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "ROUTE_FUX_TRANSFER_ROLE")
	return *ret0, err
}

// ROUTE_FUX_TRANSFER_ROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxPayBox *FuxPayBoxSession) ROUTE_FUX_TRANSFER_ROLE() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_TRANSFER_ROLE(&_FuxPayBox.CallOpts)
}

// ROUTE_FUX_TRANSFER_ROLE is a free data retrieval call binding the contract method 0x071db561.
//
// Solidity: function ROUTE_FUX_TRANSFER_ROLE() constant returns(string)
func (_FuxPayBox *FuxPayBoxCallerSession) ROUTE_FUX_TRANSFER_ROLE() (string, error) {
	return _FuxPayBox.Contract.ROUTE_FUX_TRANSFER_ROLE(&_FuxPayBox.CallOpts)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_FuxPayBox *FuxPayBoxCaller) GetProxyAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "getProxyAddress")
	return *ret0, err
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_FuxPayBox *FuxPayBoxSession) GetProxyAddress() (common.Address, error) {
	return _FuxPayBox.Contract.GetProxyAddress(&_FuxPayBox.CallOpts)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_FuxPayBox *FuxPayBoxCallerSession) GetProxyAddress() (common.Address, error) {
	return _FuxPayBox.Contract.GetProxyAddress(&_FuxPayBox.CallOpts)
}

// IsUsingProxy is a free data retrieval call binding the contract method 0x781bed44.
//
// Solidity: function isUsingProxy() constant returns(bool)
func (_FuxPayBox *FuxPayBoxCaller) IsUsingProxy(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "isUsingProxy")
	return *ret0, err
}

// IsUsingProxy is a free data retrieval call binding the contract method 0x781bed44.
//
// Solidity: function isUsingProxy() constant returns(bool)
func (_FuxPayBox *FuxPayBoxSession) IsUsingProxy() (bool, error) {
	return _FuxPayBox.Contract.IsUsingProxy(&_FuxPayBox.CallOpts)
}

// IsUsingProxy is a free data retrieval call binding the contract method 0x781bed44.
//
// Solidity: function isUsingProxy() constant returns(bool)
func (_FuxPayBox *FuxPayBoxCallerSession) IsUsingProxy() (bool, error) {
	return _FuxPayBox.Contract.IsUsingProxy(&_FuxPayBox.CallOpts)
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

// SAddr is a free data retrieval call binding the contract method 0x75ae7a8d.
//
// Solidity: function sAddr() constant returns(address)
func (_FuxPayBox *FuxPayBoxCaller) SAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "sAddr")
	return *ret0, err
}

// SAddr is a free data retrieval call binding the contract method 0x75ae7a8d.
//
// Solidity: function sAddr() constant returns(address)
func (_FuxPayBox *FuxPayBoxSession) SAddr() (common.Address, error) {
	return _FuxPayBox.Contract.SAddr(&_FuxPayBox.CallOpts)
}

// SAddr is a free data retrieval call binding the contract method 0x75ae7a8d.
//
// Solidity: function sAddr() constant returns(address)
func (_FuxPayBox *FuxPayBoxCallerSession) SAddr() (common.Address, error) {
	return _FuxPayBox.Contract.SAddr(&_FuxPayBox.CallOpts)
}

// StaticAddr is a free data retrieval call binding the contract method 0x2bc3a44c.
//
// Solidity: function staticAddr() constant returns(bool)
func (_FuxPayBox *FuxPayBoxCaller) StaticAddr(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "staticAddr")
	return *ret0, err
}

// StaticAddr is a free data retrieval call binding the contract method 0x2bc3a44c.
//
// Solidity: function staticAddr() constant returns(bool)
func (_FuxPayBox *FuxPayBoxSession) StaticAddr() (bool, error) {
	return _FuxPayBox.Contract.StaticAddr(&_FuxPayBox.CallOpts)
}

// StaticAddr is a free data retrieval call binding the contract method 0x2bc3a44c.
//
// Solidity: function staticAddr() constant returns(bool)
func (_FuxPayBox *FuxPayBoxCallerSession) StaticAddr() (bool, error) {
	return _FuxPayBox.Contract.StaticAddr(&_FuxPayBox.CallOpts)
}

// SysProxy is a free data retrieval call binding the contract method 0x32061899.
//
// Solidity: function sysProxy() constant returns(address)
func (_FuxPayBox *FuxPayBoxCaller) SysProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxPayBox.contract.Call(opts, out, "sysProxy")
	return *ret0, err
}

// SysProxy is a free data retrieval call binding the contract method 0x32061899.
//
// Solidity: function sysProxy() constant returns(address)
func (_FuxPayBox *FuxPayBoxSession) SysProxy() (common.Address, error) {
	return _FuxPayBox.Contract.SysProxy(&_FuxPayBox.CallOpts)
}

// SysProxy is a free data retrieval call binding the contract method 0x32061899.
//
// Solidity: function sysProxy() constant returns(address)
func (_FuxPayBox *FuxPayBoxCallerSession) SysProxy() (common.Address, error) {
	return _FuxPayBox.Contract.SysProxy(&_FuxPayBox.CallOpts)
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
