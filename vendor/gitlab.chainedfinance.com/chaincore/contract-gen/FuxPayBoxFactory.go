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

// FuxPayBoxFactoryABI is the input ABI used to generate the binding from.
const FuxPayBoxFactoryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isUsingRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sysRouter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CONFIG\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_isFuxHub\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_CLUSTER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROUTE_CFRBAC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRouterAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_HUB\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_FUX_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"_isInFuxGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_router\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_boxId\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"createPayBox\",\"outputs\":[{\"name\":\"boxAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_boxId\",\"type\":\"uint256\"}],\"name\":\"getPayBoxAddress\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getBoxId\",\"outputs\":[{\"name\":\"boxId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_boxAddr\",\"type\":\"address\"}],\"name\":\"closeBox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FuxPayBoxFactoryBin is the compiled bytecode used for deploying new contracts.
const FuxPayBoxFactoryBin = `0x60806040523480156200001157600080fd5b5060405160208062002cc783398101806040528101908080519060200190929190505050620000846040805190810160405280600c81526020017f726f75746572436f6e6669670000000000000000000000000000000000000000815250620000f4640100000000026401000000009004565b620000d36040805190810160405280600981526020017f467578436f6e6669670000000000000000000000000000000000000000000000815250620000f4640100000000026401000000009004565b620000ed8162000110640100000000026401000000009004565b506200023f565b80600190805190602001906200010c92919062000190565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156200014d57600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001d357805160ff191683800117855562000204565b8280016001018555821562000204579182015b8281111562000203578251825591602001919060010190620001e6565b5b50905062000213919062000217565b5090565b6200023c91905b80821115620002385760008160009055506001016200021e565b5090565b90565b612a78806200024f6000396000f3006080604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630d20aa48146100d5578063606c56321461010457806366b9852b146101475780636ed56bce1461019e578063946f01671461022e578063ade6172414610285578063b8d43d27146102e0578063c427d38d14610370578063d54f7d5e14610400578063d6b419fb14610457578063d89e6c4f146104e7578063f0b15e8214610577578063fcfc43a7146105d2578063fe07cf8e1461063f575b600080fd5b3480156100e157600080fd5b506100ea6106cc565b604051808215151515815260200191505060405180910390f35b34801561011057600080fd5b50610145600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106d5565b005b34801561015357600080fd5b5061015c61098b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101aa57600080fd5b506101b36109b0565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101f35780820151818401526020810190506101d8565b50505050905090810190601f1680156102205780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561023a57600080fd5b5061026f600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109e9565b6040518082815260200191505060405180910390f35b34801561029157600080fd5b506102c6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a32565b604051808215151515815260200191505060405180910390f35b3480156102ec57600080fd5b506102f5610bba565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561033557808201518184015260208101905061031a565b50505050905090810190601f1680156103625780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561037c57600080fd5b50610385610bf3565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103c55780820151818401526020810190506103aa565b50505050905090810190601f1680156103f25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561040c57600080fd5b50610415610c2c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561046357600080fd5b5061046c610c55565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104ac578082015181840152602081019050610491565b50505050905090810190601f1680156104d95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104f357600080fd5b506104fc610c8e565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561053c578082015181840152602081019050610521565b50505050905090810190601f1680156105695780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561058357600080fd5b506105b8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610cc7565b604051808215151515815260200191505060405180910390f35b3480156105de57600080fd5b506105fd60048036038101908080359060200190929190505050610e4f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561064b57600080fd5b5061068a60048036038101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e8c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60006001905090565b6000806106e1336111dd565b15156106ec57600080fd5b6000600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541415151561073b57600080fd5b8291508173ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156107a257600080fd5b505af11580156107b6573d6000803e3d6000fd5b505050506040513d60208110156107cc57600080fd5b810190808051906020019092919050505090508173ffffffffffffffffffffffffffffffffffffffff166341c0e1b56040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561084357600080fd5b505af1158015610857573d6000803e3d6000fd5b5050505060026000600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905561097c600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846112c3565b61098683826112c3565b505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600c81526020017f726f75746572436f6e666967000000000000000000000000000000000000000081525081565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000610a3c6113b5565b73ffffffffffffffffffffffffffffffffffffffff1663217fe6c6836040805190810160405280600681526020017f66757848756200000000000000000000000000000000000000000000000000008152506040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610b2b578082015181840152602081019050610b10565b50505050905090810190601f168015610b585780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015610b7857600080fd5b505af1158015610b8c573d6000803e3d6000fd5b505050506040513d6020811015610ba257600080fd5b81019080805190602001909291905050509050919050565b6040805190810160405280600b81526020017f726f6c65436c757374657200000000000000000000000000000000000000000081525081565b6040805190810160405280600c81526020017f726f75746572434652424143000000000000000000000000000000000000000081525081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6040805190810160405280600681526020017f667578487562000000000000000000000000000000000000000000000000000081525081565b6040805190810160405280600c81526020017f726f6c6546757847726f7570000000000000000000000000000000000000000081525081565b6000610cd16113b5565b73ffffffffffffffffffffffffffffffffffffffff1663217fe6c6836040805190810160405280600c81526020017f726f6c6546757847726f757000000000000000000000000000000000000000008152506040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610dc0578082015181840152602081019050610da5565b50505050905090810190601f168015610ded5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015610e0d57600080fd5b505af1158015610e21573d6000803e3d6000fd5b505050506040513d6020811015610e3757600080fd5b81019080805190602001909291905050509050919050565b60006002600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600080600073ffffffffffffffffffffffffffffffffffffffff166002600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610efd57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610f27611789565b808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050604051809103906000f080158015610f79573d6000803e3d6000fd5b5090508073ffffffffffffffffffffffffffffffffffffffff1663f2fde38b336040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561101757600080fd5b505af115801561102b573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16634e71e0c86040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561109357600080fd5b505af11580156110a7573d6000803e3d6000fd5b50505050809150816002600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555033600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506111cc33836113c4565b6111d682846113c4565b5092915050565b60006111e76113b5565b73ffffffffffffffffffffffffffffffffffffffff166324d7806c836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561128157600080fd5b505af1158015611295573d6000803e3d6000fd5b505050506040513d60208110156112ab57600080fd5b81019080805190602001909291905050509050919050565b6112cb6114b6565b73ffffffffffffffffffffffffffffffffffffffff16638e69029c83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561139957600080fd5b505af11580156113ad573d6000803e3d6000fd5b505050505050565b60006113bf6114fb565b905090565b6113cc6114b6565b73ffffffffffffffffffffffffffffffffffffffff1663bb88d88d83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050600060405180830381600087803b15801561149a57600080fd5b505af11580156114ae573d6000803e3d6000fd5b505050505050565b60006114f66040805190810160405280601181526020017f4675785472616e7366657246696c746572000000000000000000000000000000815250611650565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c251102560016040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156115f05780601f106115c5576101008083540402835291602001916115f0565b820191906000526020600020905b8154815290600101906020018083116115d357829003601f168201915b505092505050602060405180830381600087803b15801561161057600080fd5b505af1158015611624573d6000803e3d6000fd5b505050506040513d602081101561163a57600080fd5b8101908080519060200190929190505050905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c2511025836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156116fb5780820151818401526020810190506116e0565b50505050905090810190601f1680156117285780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b15801561174757600080fd5b505af115801561175b573d6000803e3d6000fd5b505050506040513d602081101561177157600080fd5b81019080805190602001909291905050509050919050565b6040516112b38061179a833901905600608060405234801561001057600080fd5b506040516020806112b38339810180604052810190808051906020019092919050505033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061008c81610092640100000000026401000000009004565b50610111565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156100ce57600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611193806101206000396000f3006080604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630d20aa48146100d557806341c0e1b5146101045780634e71e0c81461011b57806366b9852b146101325780636ed56bce146101895780638da5cb5b14610219578063a9059cbb14610270578063c427d38d146102bd578063d54f7d5e1461034d578063d6b419fb146103a4578063d89e6c4f14610434578063e30c3978146104c4578063f0b9e5ba1461051b578063f2fde38b14610600575b600080fd5b3480156100e157600080fd5b506100ea610643565b604051808215151515815260200191505060405180910390f35b34801561011057600080fd5b5061011961064c565b005b34801561012757600080fd5b506101306107e9565b005b34801561013e57600080fd5b50610147610a58565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561019557600080fd5b5061019e610a7d565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101de5780820151818401526020810190506101c3565b50505050905090810190601f16801561020b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561022557600080fd5b5061022e610ab6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561027c57600080fd5b506102bb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610adc565b005b3480156102c957600080fd5b506102d2610c32565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103125780820151818401526020810190506102f7565b50505050905090810190601f16801561033f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561035957600080fd5b50610362610c6b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103b057600080fd5b506103b9610c94565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103f95780820151818401526020810190506103de565b50505050905090810190601f1680156104265780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561044057600080fd5b50610449610ccd565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561048957808201518184015260208101905061046e565b50505050905090810190601f1680156104b65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104d057600080fd5b506104d9610d06565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561052757600080fd5b506105ac600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610d2c565b60405180827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b34801561060c57600080fd5b50610641600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d5b565b005b60006001905090565b61068a6040805190810160405280601081526020017f467578506179426f78466163746f727900000000000000000000000000000000815250610d8b565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156106c357600080fd5b60006106cd610ec4565b73ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561076757600080fd5b505af115801561077b573d6000803e3d6000fd5b505050506040513d602081101561079157600080fd5b81019080805190602001909291905050501415156107ae57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b6108276040805190810160405280601081526020017f467578506179426f78466163746f727900000000000000000000000000000000815250610d8b565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480156108b05750600073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b806109085750600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b151561091357600080fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6040805190810160405280600c81526020017f726f75746572436f6e666967000000000000000000000000000000000000000081525081565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b3857600080fd5b610b40610ec4565b73ffffffffffffffffffffffffffffffffffffffff1663beabacc83084846040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015610c1657600080fd5b505af1158015610c2a573d6000803e3d6000fd5b505050505050565b6040805190810160405280600c81526020017f726f75746572434652424143000000000000000000000000000000000000000081525081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6040805190810160405280600681526020017f667578487562000000000000000000000000000000000000000000000000000081525081565b6040805190810160405280600c81526020017f726f6c6546757847726f7570000000000000000000000000000000000000000081525081565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600063f0b9e5ba7c01000000000000000000000000000000000000000000000000000000000290509392505050565b610d6481610f09565b80610d745750610d7333610f09565b5b1515610d7f57600080fd5b610d88816110c7565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c2511025836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610e36578082015181840152602081019050610e1b565b50505050905090810190601f168015610e635780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b158015610e8257600080fd5b505af1158015610e96573d6000803e3d6000fd5b505050506040513d6020811015610eac57600080fd5b81019080805190602001909291905050509050919050565b6000610f046040805190810160405280600881526020017f467578546f6b656e000000000000000000000000000000000000000000000000815250610d8b565b905090565b6000610f496040805190810160405280600c81526020017f726f757465724346524241430000000000000000000000000000000000000000815250610d8b565b73ffffffffffffffffffffffffffffffffffffffff1663217fe6c6836040805190810160405280600681526020017f66757848756200000000000000000000000000000000000000000000000000008152506040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561103857808201518184015260208101905061101d565b50505050905090810190601f1680156110655780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561108557600080fd5b505af1158015611099573d6000803e3d6000fd5b505050506040513d60208110156110af57600080fd5b81019080805190602001909291905050509050919050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561112357600080fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a7230582016b04d762763cc24dc866be2bd034d5e67eb51d0d04d82e210a8370769dbafa80029a165627a7a72305820591d89b58dc1a61c2c4f1fe04609ac1db0aa896593a20d36938b40437e103bfd0029`

// DeployFuxPayBoxFactory deploys a new Ethereum contract, binding an instance of FuxPayBoxFactory to it.
func DeployFuxPayBoxFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _router common.Address) (common.Address, *types.Transaction, *FuxPayBoxFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxPayBoxFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FuxPayBoxFactoryBin), backend, _router)
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

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) ROLECLUSTER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "ROLE_CLUSTER")
	return *ret0, err
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) ROLECLUSTER() (string, error) {
	return _FuxPayBoxFactory.Contract.ROLECLUSTER(&_FuxPayBoxFactory.CallOpts)
}

// ROLECLUSTER is a free data retrieval call binding the contract method 0xb8d43d27.
//
// Solidity: function ROLE_CLUSTER() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) ROLECLUSTER() (string, error) {
	return _FuxPayBoxFactory.Contract.ROLECLUSTER(&_FuxPayBoxFactory.CallOpts)
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) ROLEFUXGROUP(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "ROLE_FUX_GROUP")
	return *ret0, err
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) ROLEFUXGROUP() (string, error) {
	return _FuxPayBoxFactory.Contract.ROLEFUXGROUP(&_FuxPayBoxFactory.CallOpts)
}

// ROLEFUXGROUP is a free data retrieval call binding the contract method 0xd89e6c4f.
//
// Solidity: function ROLE_FUX_GROUP() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) ROLEFUXGROUP() (string, error) {
	return _FuxPayBoxFactory.Contract.ROLEFUXGROUP(&_FuxPayBoxFactory.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) ROLEFUXHUB(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "ROLE_FUX_HUB")
	return *ret0, err
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) ROLEFUXHUB() (string, error) {
	return _FuxPayBoxFactory.Contract.ROLEFUXHUB(&_FuxPayBoxFactory.CallOpts)
}

// ROLEFUXHUB is a free data retrieval call binding the contract method 0xd6b419fb.
//
// Solidity: function ROLE_FUX_HUB() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) ROLEFUXHUB() (string, error) {
	return _FuxPayBoxFactory.Contract.ROLEFUXHUB(&_FuxPayBoxFactory.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) ROUTECFRBAC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "ROUTE_CFRBAC")
	return *ret0, err
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) ROUTECFRBAC() (string, error) {
	return _FuxPayBoxFactory.Contract.ROUTECFRBAC(&_FuxPayBoxFactory.CallOpts)
}

// ROUTECFRBAC is a free data retrieval call binding the contract method 0xc427d38d.
//
// Solidity: function ROUTE_CFRBAC() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) ROUTECFRBAC() (string, error) {
	return _FuxPayBoxFactory.Contract.ROUTECFRBAC(&_FuxPayBoxFactory.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) ROUTECONFIG(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "ROUTE_CONFIG")
	return *ret0, err
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) ROUTECONFIG() (string, error) {
	return _FuxPayBoxFactory.Contract.ROUTECONFIG(&_FuxPayBoxFactory.CallOpts)
}

// ROUTECONFIG is a free data retrieval call binding the contract method 0x6ed56bce.
//
// Solidity: function ROUTE_CONFIG() constant returns(string)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) ROUTECONFIG() (string, error) {
	return _FuxPayBoxFactory.Contract.ROUTECONFIG(&_FuxPayBoxFactory.CallOpts)
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) IsFuxHub(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "_isFuxHub", _addr)
	return *ret0, err
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) IsFuxHub(_addr common.Address) (bool, error) {
	return _FuxPayBoxFactory.Contract.IsFuxHub(&_FuxPayBoxFactory.CallOpts, _addr)
}

// IsFuxHub is a free data retrieval call binding the contract method 0xade61724.
//
// Solidity: function _isFuxHub(_addr address) constant returns(bool)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) IsFuxHub(_addr common.Address) (bool, error) {
	return _FuxPayBoxFactory.Contract.IsFuxHub(&_FuxPayBoxFactory.CallOpts, _addr)
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) IsInFuxGroup(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "_isInFuxGroup", _addr)
	return *ret0, err
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) IsInFuxGroup(_addr common.Address) (bool, error) {
	return _FuxPayBoxFactory.Contract.IsInFuxGroup(&_FuxPayBoxFactory.CallOpts, _addr)
}

// IsInFuxGroup is a free data retrieval call binding the contract method 0xf0b15e82.
//
// Solidity: function _isInFuxGroup(_addr address) constant returns(bool)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) IsInFuxGroup(_addr common.Address) (bool, error) {
	return _FuxPayBoxFactory.Contract.IsInFuxGroup(&_FuxPayBoxFactory.CallOpts, _addr)
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

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) GetRouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "getRouterAddress")
	return *ret0, err
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) GetRouterAddress() (common.Address, error) {
	return _FuxPayBoxFactory.Contract.GetRouterAddress(&_FuxPayBoxFactory.CallOpts)
}

// GetRouterAddress is a free data retrieval call binding the contract method 0xd54f7d5e.
//
// Solidity: function getRouterAddress() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) GetRouterAddress() (common.Address, error) {
	return _FuxPayBoxFactory.Contract.GetRouterAddress(&_FuxPayBoxFactory.CallOpts)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) IsUsingRouter(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "isUsingRouter")
	return *ret0, err
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) IsUsingRouter() (bool, error) {
	return _FuxPayBoxFactory.Contract.IsUsingRouter(&_FuxPayBoxFactory.CallOpts)
}

// IsUsingRouter is a free data retrieval call binding the contract method 0x0d20aa48.
//
// Solidity: function isUsingRouter() constant returns(bool)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) IsUsingRouter() (bool, error) {
	return _FuxPayBoxFactory.Contract.IsUsingRouter(&_FuxPayBoxFactory.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCaller) SysRouter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxPayBoxFactory.contract.Call(opts, out, "sysRouter")
	return *ret0, err
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) SysRouter() (common.Address, error) {
	return _FuxPayBoxFactory.Contract.SysRouter(&_FuxPayBoxFactory.CallOpts)
}

// SysRouter is a free data retrieval call binding the contract method 0x66b9852b.
//
// Solidity: function sysRouter() constant returns(address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryCallerSession) SysRouter() (common.Address, error) {
	return _FuxPayBoxFactory.Contract.SysRouter(&_FuxPayBoxFactory.CallOpts)
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

// CreatePayBox is a paid mutator transaction binding the contract method 0xfe07cf8e.
//
// Solidity: function createPayBox(_boxId uint256, _to address) returns(boxAddr address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactor) CreatePayBox(opts *bind.TransactOpts, _boxId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.contract.Transact(opts, "createPayBox", _boxId, _to)
}

// CreatePayBox is a paid mutator transaction binding the contract method 0xfe07cf8e.
//
// Solidity: function createPayBox(_boxId uint256, _to address) returns(boxAddr address)
func (_FuxPayBoxFactory *FuxPayBoxFactorySession) CreatePayBox(_boxId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.CreatePayBox(&_FuxPayBoxFactory.TransactOpts, _boxId, _to)
}

// CreatePayBox is a paid mutator transaction binding the contract method 0xfe07cf8e.
//
// Solidity: function createPayBox(_boxId uint256, _to address) returns(boxAddr address)
func (_FuxPayBoxFactory *FuxPayBoxFactoryTransactorSession) CreatePayBox(_boxId *big.Int, _to common.Address) (*types.Transaction, error) {
	return _FuxPayBoxFactory.Contract.CreatePayBox(&_FuxPayBoxFactory.TransactOpts, _boxId, _to)
}
