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

// FuxTokenABI is the input ABI used to generate the binding from.
const FuxTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isUsingRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sysRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CONFIG\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_isFuxHub\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CFRBAC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRouterAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_HUB\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_isInFuxGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_coin\",\"type\":\"uint256\"},{\"name\":\"_expire\",\"type\":\"uint256\"},{\"name\":\"_state\",\"type\":\"uint256\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"fuxBalanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FuxTokenBin is the compiled bytecode used for deploying new contracts.
const FuxTokenBin = `0x60806040523480156200001157600080fd5b5060405160208062001f2c83398101806040528101908080519060200190929190505050620000846040805190810160405280600c81526020017f726f75746572436f6e6669670000000000000000000000000000000000000000815250620000f4640100000000026401000000009004565b620000d36040805190810160405280600981526020017f467578436f6e6669670000000000000000000000000000000000000000000000815250620000f4640100000000026401000000009004565b620000ed8162000110640100000000026401000000009004565b506200023f565b80600190805190602001906200010c92919062000190565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156200014d57600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001d357805160ff191683800117855562000204565b8280016001018555821562000204579182015b8281111562000203578251825591602001919060010190620001e6565b5b50905062000213919062000217565b5090565b6200023c91905b80821115620002385760008160009055506001016200021e565b5090565b90565b611cdd806200024f6000396000f3006080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630d20aa48146100e057806366b9852b1461010f5780636ed56bce1461016657806370a08231146101f65780639dc29fac1461024d578063ade617241461029a578063b8d43d27146102f5578063beabacc814610385578063c427d38d146103f2578063cf21977c14610482578063d54f7d5e14610533578063d6b419fb1461058a578063d75ab6161461061a578063d89e6c4f14610671578063f0b15e8214610701575b600080fd5b3480156100ec57600080fd5b506100f561075c565b604051808215151515815260200191505060405180910390f35b34801561011b57600080fd5b50610124610765565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561017257600080fd5b5061017b61078a565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101bb5780820151818401526020810190506101a0565b50505050905090810190601f1680156101e85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561020257600080fd5b50610237600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107c3565b6040518082815260200191505060405180910390f35b34801561025957600080fd5b50610298600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506108a9565b005b3480156102a657600080fd5b506102db600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b99565b604051808215151515815260200191505060405180910390f35b34801561030157600080fd5b5061030a610d21565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561034a57808201518184015260208101905061032f565b50505050905090810190601f1680156103775780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561039157600080fd5b506103f0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d5a565b005b3480156103fe57600080fd5b506104076110a8565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561044757808201518184015260208101905061042c565b50505050905090810190601f1680156104745780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561048e57600080fd5b50610531600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919080359060200190929190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506110e1565b005b34801561053f57600080fd5b506105486113a4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561059657600080fd5b5061059f6113cd565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105df5780820151818401526020810190506105c4565b50505050905090810190601f16801561060c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561062657600080fd5b5061065b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611406565b6040518082815260200191505060405180910390f35b34801561067d57600080fd5b506106866114ec565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106c65780820151818401526020810190506106ab565b50505050905090810190601f1680156106f35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561070d57600080fd5b50610742600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611525565b604051808215151515815260200191505060405180910390f35b60006001905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600c81526020017f726f75746572436f6e666967000000000000000000000000000000000000000081525081565b60006107cd6116ad565b73ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561086757600080fd5b505af115801561087b573d6000803e3d6000fd5b505050506040513d602081101561089157600080fd5b81019080805190602001909291905050509050919050565b60006108b433611525565b806108c457506108c3336116f2565b5b15156108cf57600080fd5b6108d76117d8565b73ffffffffffffffffffffffffffffffffffffffff16632639508d836040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561094557600080fd5b505af1158015610959573d6000803e3d6000fd5b505050506040513d602081101561096f57600080fd5b8101908080519060200190929190505050905061098a61181d565b73ffffffffffffffffffffffffffffffffffffffff166326ffee0884836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610a2c57600080fd5b505af1158015610a40573d6000803e3d6000fd5b50505050610a4c6116ad565b73ffffffffffffffffffffffffffffffffffffffff16639dc29fac84846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610aee57600080fd5b505af1158015610b02573d6000803e3d6000fd5b50505050610b0e6117d8565b73ffffffffffffffffffffffffffffffffffffffff1663e234f61d836040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b158015610b7c57600080fd5b505af1158015610b90573d6000803e3d6000fd5b50505050505050565b6000610ba3611862565b73ffffffffffffffffffffffffffffffffffffffff1663217fe6c6836040805190810160405280600681526020017f66757848756200000000000000000000000000000000000000000000000000008152506040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610c92578082015181840152602081019050610c77565b50505050905090810190601f168015610cbf5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015610cdf57600080fd5b505af1158015610cf3573d6000803e3d6000fd5b505050506040513d6020811015610d0957600080fd5b81019080805190602001909291905050509050919050565b6040805190810160405280600b81526020017f726f6c65436c757374657200000000000000000000000000000000000000000081525081565b60008383610d688282611871565b1515610d7357600080fd5b610d7b6117d8565b73ffffffffffffffffffffffffffffffffffffffff16632639508d856040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b158015610de957600080fd5b505af1158015610dfd573d6000803e3d6000fd5b505050506040513d6020811015610e1357600080fd5b81019080805190602001909291905050509250610e2e61181d565b73ffffffffffffffffffffffffffffffffffffffff1663f5d82b6b86856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610ed057600080fd5b505af1158015610ee4573d6000803e3d6000fd5b50505050610ef061181d565b73ffffffffffffffffffffffffffffffffffffffff166326ffee0887856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610f9257600080fd5b505af1158015610fa6573d6000803e3d6000fd5b50505050610fb26116ad565b73ffffffffffffffffffffffffffffffffffffffff1663beabacc88787876040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561108857600080fd5b505af115801561109c573d6000803e3d6000fd5b50505050505050505050565b6040805190810160405280600c81526020017f726f75746572434652424143000000000000000000000000000000000000000081525081565b6110ea33611525565b806110fa57506110f9336116f2565b5b151561110557600080fd5b61110d6117d8565b73ffffffffffffffffffffffffffffffffffffffff166386b27f4d86868686866040518663ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018086815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156111b0578082015181840152602081019050611195565b50505050905090810190601f1680156111dd5780820380516001836020036101000a031916815260200191505b509650505050505050600060405180830381600087803b15801561120057600080fd5b505af1158015611214573d6000803e3d6000fd5b505050506112206116ad565b73ffffffffffffffffffffffffffffffffffffffff166340c10f1987876040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156112c257600080fd5b505af11580156112d6573d6000803e3d6000fd5b505050506112e261181d565b73ffffffffffffffffffffffffffffffffffffffff1663f5d82b6b87866040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561138457600080fd5b505af1158015611398573d6000803e3d6000fd5b50505050505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6040805190810160405280600681526020017f667578487562000000000000000000000000000000000000000000000000000081525081565b600061141061181d565b73ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156114aa57600080fd5b505af11580156114be573d6000803e3d6000fd5b505050506040513d60208110156114d457600080fd5b81019080805190602001909291905050509050919050565b6040805190810160405280600c81526020017f726f6c6546757847726f7570000000000000000000000000000000000000000081525081565b600061152f611862565b73ffffffffffffffffffffffffffffffffffffffff1663217fe6c6836040805190810160405280600c81526020017f726f6c6546757847726f757000000000000000000000000000000000000000008152506040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561161e578082015181840152602081019050611603565b50505050905090810190601f16801561164b5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561166b57600080fd5b505af115801561167f573d6000803e3d6000fd5b505050506040513d602081101561169557600080fd5b81019080805190602001909291905050509050919050565b60006116ed6040805190810160405280600e81526020017f467578455243373231546f6b656e0000000000000000000000000000000000008152506119bc565b905090565b60006116fc611862565b73ffffffffffffffffffffffffffffffffffffffff166324d7806c836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561179657600080fd5b505af11580156117aa573d6000803e3d6000fd5b505050506040513d60208110156117c057600080fd5b81019080805190602001909291905050509050919050565b60006118186040805190810160405280600a81526020017f46757853746f72616765000000000000000000000000000000000000000000008152506119bc565b905090565b600061185d6040805190810160405280600781526020017f467578436f696e000000000000000000000000000000000000000000000000008152506119bc565b905090565b600061186c611af5565b905090565b600061187b611c4a565b73ffffffffffffffffffffffffffffffffffffffff16631de1ed6484846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561194957600080fd5b505af115801561195d573d6000803e3d6000fd5b505050506040513d602081101561197357600080fd5b810190808051906020019092919050505080611994575061199383611c8f565b5b806119a457506119a382611c8f565b5b806119b457506119b333611c8f565b5b905092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c2511025836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611a67578082015181840152602081019050611a4c565b50505050905090810190601f168015611a945780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b158015611ab357600080fd5b505af1158015611ac7573d6000803e3d6000fd5b505050506040513d6020811015611add57600080fd5b81019080805190602001909291905050509050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c251102560016040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015611bea5780601f10611bbf57610100808354040283529160200191611bea565b820191906000526020600020905b815481529060010190602001808311611bcd57829003601f168201915b505092505050602060405180830381600087803b158015611c0a57600080fd5b505af1158015611c1e573d6000803e3d6000fd5b505050506040513d6020811015611c3457600080fd5b8101908080519060200190929190505050905090565b6000611c8a6040805190810160405280601181526020017f4675785472616e7366657246696c7465720000000000000000000000000000008152506119bc565b905090565b6000611c9a826116f2565b80611caa5750611ca982610b99565b5b90509190505600a165627a7a723058200be4909be3c648c1e4463090c223ad6c3937bcd3325d7ee2ffed53627e4c51600029`

// DeployFuxToken deploys a new Ethereum contract, binding an instance of FuxToken to it.
func DeployFuxToken(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address) (common.Address, *types.Transaction, *FuxToken, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FuxTokenBin), backend, _proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FuxToken{FuxTokenCaller: FuxTokenCaller{contract: contract}, FuxTokenTransactor: FuxTokenTransactor{contract: contract}, FuxTokenFilterer: FuxTokenFilterer{contract: contract}}, nil
}

// FuxToken is an auto generated Go binding around an Ethereum contract.
type FuxToken struct {
	FuxTokenCaller     // Read-only binding to the contract
	FuxTokenTransactor // Write-only binding to the contract
	FuxTokenFilterer   // Log filterer for contract events
}

// FuxTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuxTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuxTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FuxTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuxTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuxTokenSession struct {
	Contract     *FuxToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuxTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuxTokenCallerSession struct {
	Contract *FuxTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FuxTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuxTokenTransactorSession struct {
	Contract     *FuxTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FuxTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuxTokenRaw struct {
	Contract *FuxToken // Generic contract binding to access the raw methods on
}

// FuxTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuxTokenCallerRaw struct {
	Contract *FuxTokenCaller // Generic read-only contract binding to access the raw methods on
}

// FuxTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuxTokenTransactorRaw struct {
	Contract *FuxTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFuxToken creates a new instance of FuxToken, bound to a specific deployed contract.
func NewFuxToken(address common.Address, backend bind.ContractBackend) (*FuxToken, error) {
	contract, err := bindFuxToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FuxToken{FuxTokenCaller: FuxTokenCaller{contract: contract}, FuxTokenTransactor: FuxTokenTransactor{contract: contract}, FuxTokenFilterer: FuxTokenFilterer{contract: contract}}, nil
}

// NewFuxTokenCaller creates a new read-only instance of FuxToken, bound to a specific deployed contract.
func NewFuxTokenCaller(address common.Address, caller bind.ContractCaller) (*FuxTokenCaller, error) {
	contract, err := bindFuxToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FuxTokenCaller{contract: contract}, nil
}

// NewFuxTokenTransactor creates a new write-only instance of FuxToken, bound to a specific deployed contract.
func NewFuxTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*FuxTokenTransactor, error) {
	contract, err := bindFuxToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FuxTokenTransactor{contract: contract}, nil
}

// NewFuxTokenFilterer creates a new log filterer instance of FuxToken, bound to a specific deployed contract.
func NewFuxTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*FuxTokenFilterer, error) {
	contract, err := bindFuxToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FuxTokenFilterer{contract: contract}, nil
}

// bindFuxToken binds a generic wrapper to an already deployed contract.
func bindFuxToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxToken *FuxTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxToken.Contract.FuxTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxToken *FuxTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxToken.Contract.FuxTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxToken *FuxTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxToken.Contract.FuxTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FuxToken *FuxTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FuxToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FuxToken *FuxTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FuxToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FuxToken *FuxTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FuxToken.Contract.contract.Transact(opts, method, params...)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROLECLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxToken *FuxTokenSession) ROLECLUSTER() (string, error) {
	return _FuxToken.Contract.ROLECLUSTER(&_FuxToken.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROLECLUSTER() (string, error) {
	return _FuxToken.Contract.ROLECLUSTER(&_FuxToken.CallOpts)
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROLEFUXGROUP(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROLE_FUX_GROUP")
	return *ret0, err
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxToken *FuxTokenSession) ROLEFUXGROUP() (string, error) {
	return _FuxToken.Contract.ROLEFUXGROUP(&_FuxToken.CallOpts)
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROLEFUXGROUP() (string, error) {
	return _FuxToken.Contract.ROLEFUXGROUP(&_FuxToken.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROLEFUXHUB(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROLE_FUX_HUB")
	return *ret0, err
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxToken *FuxTokenSession) ROLEFUXHUB() (string, error) {
	return _FuxToken.Contract.ROLEFUXHUB(&_FuxToken.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROLEFUXHUB() (string, error) {
	return _FuxToken.Contract.ROLEFUXHUB(&_FuxToken.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROUTECFRBAC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROUTE_CFRBAC")
	return *ret0, err
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxToken *FuxTokenSession) ROUTECFRBAC() (string, error) {
	return _FuxToken.Contract.ROUTECFRBAC(&_FuxToken.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROUTECFRBAC() (string, error) {
	return _FuxToken.Contract.ROUTECFRBAC(&_FuxToken.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxToken *FuxTokenCaller) ROUTECONFIG(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ROUTE_CONFIG")
	return *ret0, err
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxToken *FuxTokenSession) ROUTECONFIG() (string, error) {
	return _FuxToken.Contract.ROUTECONFIG(&_FuxToken.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) ROUTECONFIG() (string, error) {
	return _FuxToken.Contract.ROUTECONFIG(&_FuxToken.CallOpts)
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxToken *FuxTokenCaller) IsFuxHub(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "_isFuxHub", _addr)
	return *ret0, err
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxToken *FuxTokenSession) IsFuxHub(_addr common.Address) (bool, error) {
	return _FuxToken.Contract.IsFuxHub(&_FuxToken.CallOpts, _addr)
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxToken *FuxTokenCallerSession) IsFuxHub(_addr common.Address) (bool, error) {
	return _FuxToken.Contract.IsFuxHub(&_FuxToken.CallOpts, _addr)
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxToken *FuxTokenCaller) IsInFuxGroup(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "_isInFuxGroup", _addr)
	return *ret0, err
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxToken *FuxTokenSession) IsInFuxGroup(_addr common.Address) (bool, error) {
	return _FuxToken.Contract.IsInFuxGroup(&_FuxToken.CallOpts, _addr)
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxToken *FuxTokenCallerSession) IsInFuxGroup(_addr common.Address) (bool, error) {
	return _FuxToken.Contract.IsInFuxGroup(&_FuxToken.CallOpts, _addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_FuxToken *FuxTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_FuxToken *FuxTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _FuxToken.Contract.BalanceOf(&_FuxToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_FuxToken *FuxTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _FuxToken.Contract.BalanceOf(&_FuxToken.CallOpts, _owner)
}

// FuxBalanceOf is a free data retrieval call binding the contract method 0xd75ab616.
//
// Solidity: function fuxBalanceOf(_owner address) constant returns(uint256)
func (_FuxToken *FuxTokenCaller) FuxBalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "fuxBalanceOf", _owner)
	return *ret0, err
}

// FuxBalanceOf is a free data retrieval call binding the contract method 0xd75ab616.
//
// Solidity: function fuxBalanceOf(_owner address) constant returns(uint256)
func (_FuxToken *FuxTokenSession) FuxBalanceOf(_owner common.Address) (*big.Int, error) {
	return _FuxToken.Contract.FuxBalanceOf(&_FuxToken.CallOpts, _owner)
}

// FuxBalanceOf is a free data retrieval call binding the contract method 0xd75ab616.
//
// Solidity: function fuxBalanceOf(_owner address) constant returns(uint256)
func (_FuxToken *FuxTokenCallerSession) FuxBalanceOf(_owner common.Address) (*big.Int, error) {
	return _FuxToken.Contract.FuxBalanceOf(&_FuxToken.CallOpts, _owner)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxToken *FuxTokenCaller) GetRouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "getRouterAddress")
	return *ret0, err
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxToken *FuxTokenSession) GetRouterAddress() (common.Address, error) {
	return _FuxToken.Contract.GetRouterAddress(&_FuxToken.CallOpts)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxToken *FuxTokenCallerSession) GetRouterAddress() (common.Address, error) {
	return _FuxToken.Contract.GetRouterAddress(&_FuxToken.CallOpts)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxToken *FuxTokenCaller) IsUsingRouter(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "isUsingRouter")
	return *ret0, err
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxToken *FuxTokenSession) IsUsingRouter() (bool, error) {
	return _FuxToken.Contract.IsUsingRouter(&_FuxToken.CallOpts)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxToken *FuxTokenCallerSession) IsUsingRouter() (bool, error) {
	return _FuxToken.Contract.IsUsingRouter(&_FuxToken.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxToken *FuxTokenCaller) SysRouter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "sysRouter")
	return *ret0, err
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxToken *FuxTokenSession) SysRouter() (common.Address, error) {
	return _FuxToken.Contract.SysRouter(&_FuxToken.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxToken *FuxTokenCallerSession) SysRouter() (common.Address, error) {
	return _FuxToken.Contract.SysRouter(&_FuxToken.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_owner address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactor) Burn(opts *bind.TransactOpts, _owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "burn", _owner, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_owner address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenSession) Burn(_owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.Burn(&_FuxToken.TransactOpts, _owner, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(_owner address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) Burn(_owner common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.Burn(&_FuxToken.TransactOpts, _owner, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xcf21977c.
//
// Solidity: function mint(_to address, _tokenId uint256, _coin uint256, _expire uint256, _state uint256, _otherInfo string) returns()
func (_FuxToken *FuxTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int, _coin *big.Int, _expire *big.Int, _state *big.Int, _otherInfo string) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "mint", _to, _tokenId, _coin, _expire, _state, _otherInfo)
}

// Mint is a paid mutator transaction binding the contract method 0xcf21977c.
//
// Solidity: function mint(_to address, _tokenId uint256, _coin uint256, _expire uint256, _state uint256, _otherInfo string) returns()
func (_FuxToken *FuxTokenSession) Mint(_to common.Address, _tokenId *big.Int, _coin *big.Int, _expire *big.Int, _state *big.Int, _otherInfo string) (*types.Transaction, error) {
	return _FuxToken.Contract.Mint(&_FuxToken.TransactOpts, _to, _tokenId, _coin, _expire, _state, _otherInfo)
}

// Mint is a paid mutator transaction binding the contract method 0xcf21977c.
//
// Solidity: function mint(_to address, _tokenId uint256, _coin uint256, _expire uint256, _state uint256, _otherInfo string) returns()
func (_FuxToken *FuxTokenTransactorSession) Mint(_to common.Address, _tokenId *big.Int, _coin *big.Int, _expire *big.Int, _state *big.Int, _otherInfo string) (*types.Transaction, error) {
	return _FuxToken.Contract.Mint(&_FuxToken.TransactOpts, _to, _tokenId, _coin, _expire, _state, _otherInfo)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_from address, _to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactor) Transfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "transfer", _from, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_from address, _to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenSession) Transfer(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.Transfer(&_FuxToken.TransactOpts, _from, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(_from address, _to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) Transfer(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.Transfer(&_FuxToken.TransactOpts, _from, _to, _tokenId)
}
