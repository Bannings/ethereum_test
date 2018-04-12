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

// FuxTokenABI is the input ABI used to generate the binding from.
const FuxTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getFuxAttrs\",\"outputs\":[{\"name\":\"coin\",\"type\":\"uint256\"},{\"name\":\"lockTime\",\"type\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approvedFor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFuxIssuer\",\"outputs\":[{\"name\":\"issuer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"takeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fuxIssuer\",\"type\":\"address\"}],\"name\":\"setFuxIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferEx\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_coin\",\"type\":\"uint256\"},{\"name\":\"_lockTime\",\"type\":\"uint256\"},{\"name\":\"_state\",\"type\":\"uint256\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"createFux\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_otherInfo\",\"type\":\"string\"}],\"name\":\"updateFuxInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_newNum\",\"type\":\"uint256\"}],\"name\":\"updateFuxLockTime\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_newNum\",\"type\":\"uint256\"}],\"name\":\"updateFuxState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_newNum\",\"type\":\"uint256\"}],\"name\":\"updateFuxCoin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFux\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"removeFux\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"queryFuxInfo\",\"outputs\":[{\"name\":\"info\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"queryHisFuxInfo\",\"outputs\":[{\"name\":\"info\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_newIds\",\"type\":\"uint256[2]\"},{\"name\":\"_newCoins\",\"type\":\"uint256[2]\"},{\"name\":\"_states\",\"type\":\"uint256[2]\"}],\"name\":\"splitFux\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FuxTokenBin is the compiled bytecode used for deploying new contracts.
const FuxTokenBin = `0x60606040526040805190810160405280600381526020017f4675780000000000000000000000000000000000000000000000000000000000815250600990805190602001906200005192919062000099565b5033600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062000148565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000dc57805160ff19168380011785556200010d565b828001600101855582156200010d579182015b828111156200010c578251825591602001919060010190620000ef565b5b5090506200011c919062000120565b5090565b6200014591905b808211156200014157600081600090555060010162000127565b5090565b90565b612ae780620001586000396000f30060606040526004361061015f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde0314610164578063095ea7b3146101f25780630d18fc1f1461023457806318160ddd146102d05780631c61984b146102f95780632a6dd48f1461033e5780632f745c59146103a1578063462b7df5146103f75780635209e984146104235780635a3f2672146104ca5780636352211e146105585780636eab2e63146105bb57806370a082311461062157806372fb7d0b1461066e578063736cd3641461069a5780638da5cb5b146106c6578063904c00c11461071b57806395d89b41146107705780639a43d9d1146107fe578063a9059cbb1461089e578063b2e6ceeb146108e0578063c0b43a4514610903578063d9cacff21461093c578063e234f61d1461099d578063ea9724ae146109c0578063f2fde38b14610a02578063f449b17114610a3b575b600080fd5b341561016f57600080fd5b610177610ad7565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101b757808201518184015260208101905061019c565b50505050905090810190601f1680156101e45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101fd57600080fd5b610232600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610b7f565b005b341561023f57600080fd5b6102556004808035906020019091905050610bd1565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561029557808201518184015260208101905061027a565b50505050905090810190601f1680156102c25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102db57600080fd5b6102e3610c6a565b6040518082815260200191505060405180910390f35b341561030457600080fd5b61031a6004808035906020019091905050610c73565b60405180848152602001838152602001828152602001935050505060405180910390f35b341561034957600080fd5b61035f6004808035906020019091905050610db7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103ac57600080fd5b6103e1600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610df4565b6040518082815260200191505060405180910390f35b341561040257600080fd5b6104216004808035906020019091908035906020019091905050610e34565b005b341561042e57600080fd5b6104c86004808035906020019091908060400190600280602002604051908101604052809291908260026020028082843782019150505050509190806040019060028060200260405190810160405280929190826002602002808284378201915050505050919080604001906002806020026040519081016040528092919082600260200280828437820191505050505091905050610e83565b005b34156104d557600080fd5b610501600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611029565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610544578082015181840152602081019050610529565b505050509050019250505060405180910390f35b341561056357600080fd5b61057960048080359060200190919050506110c6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156105c657600080fd5b61061f600480803590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050611144565b005b341561062c57600080fd5b610658600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611193565b6040518082815260200191505060405180910390f35b341561067957600080fd5b61069860048080359060200190919080359060200190919050506111df565b005b34156106a557600080fd5b6106c4600480803590602001909190803590602001909190505061122e565b005b34156106d157600080fd5b6106d961127d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561072657600080fd5b61072e6112a3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561077b57600080fd5b61078361134e565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156107c35780820151818401526020810190506107a8565b50505050905090810190601f1680156107f05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561080957600080fd5b61089c600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001909190803590602001909190803590602001909190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506113f6565b005b34156108a957600080fd5b6108de600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050611496565b005b34156108eb57600080fd5b61090160048080359060200190919050506114e8565b005b341561090e57600080fd5b61093a600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611513565b005b341561094757600080fd5b61099b600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506115b3565b005b34156109a857600080fd5b6109be6004808035906020019091905050611604565b005b34156109cb57600080fd5b610a00600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190505061165a565b005b3415610a0d57600080fd5b610a39600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061167c565b005b3415610a4657600080fd5b610a5c60048080359060200190919050506117d4565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610a9c578082015181840152602081019050610a81565b50505050905090810190601f168015610ac95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610adf61287b565b60098054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b755780601f10610b4a57610100808354040283529160200191610b75565b820191906000526020600020905b815481529060010190602001808311610b5857829003601f168201915b5050505050905090565b803373ffffffffffffffffffffffffffffffffffffffff16610ba0826110c6565b73ffffffffffffffffffffffffffffffffffffffff16141515610bc257600080fd5b610bcc838361182d565b505050565b610bd961287b565b81610be26112a3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610c4e57503373ffffffffffffffffffffffffffffffffffffffff16610c36826110c6565b73ffffffffffffffffffffffffffffffffffffffff16145b1515610c5957600080fd5b610c62836119c0565b915050919050565b60008054905090565b6000806000610c8061288f565b610c8985611a92565b1515610c9457600080fd5b6007600086815260200190815260200160002060c0604051908101604052908160008201548152602001600182015481526020016002820154815260200160038201548152602001600482018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d745780601f10610d4957610100808354040283529160200191610d74565b820191906000526020600020905b815481529060010190602001808311610d5757829003601f168201915b505050505081526020016005820160009054906101000a900460ff1615151515815250509050806020015181604001518260600151935093509350509193909250565b60006002600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000610dff83611193565b82101515610e0c57600080fd5b610e1583611029565b82815181101515610e2257fe5b90602001906020020151905092915050565b610e3c6112a3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610e7557600080fd5b610e7f8282611abf565b5050565b600080600080610e916112a3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610eca57600080fd5b610ed3886110c6565b9350610ede88610c73565b80945081955082935050505080610f24876001600281101515610efd57fe5b6020020151886000600281101515610f1157fe5b6020020151611af290919063ffffffff16565b141515610f3057600080fd5b610f3988611604565b610f5684886000600281101515610f4c57fe5b6020020151611b10565b610f7384886001600281101515610f6957fe5b6020020151611b10565b610fc9876000600281101515610f8557fe5b6020020151876000600281101515610f9957fe5b602002015185886000600281101515610fae57fe5b60200201516020604051908101604052806000815250611b73565b61101f876001600281101515610fdb57fe5b6020020151876001600281101515610fef57fe5b60200201518588600160028110151561100457fe5b60200201516020604051908101604052806000815250611b73565b5050505050505050565b6110316128ce565b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156110ba57602002820191906000526020600020905b8154815260200190600101908083116110a6575b50505050509050919050565b6000806001600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561113b57600080fd5b80915050919050565b61114c6112a3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561118557600080fd5b61118f8282611c47565b5050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490509050919050565b6111e76112a3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561122057600080fd5b61122a8282611c8a565b5050565b6112366112a3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561126f57600080fd5b6112798282611abf565b5050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561132557600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905061134b565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b90565b61135661287b565b60098054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113ec5780601f106113c1576101008083540402835291602001916113ec565b820191906000526020600020905b8154815290600101906020018083116113cf57829003601f168201915b5050505050905090565b6113fe6112a3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561143757600080fd5b6114448585858585611b73565b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415611484576114816112a3565b95505b61148e8686611b10565b505050505050565b803373ffffffffffffffffffffffffffffffffffffffff166114b7826110c6565b73ffffffffffffffffffffffffffffffffffffffff161415156114d957600080fd5b6114e38383611cbd565b505050565b6114f23382611d10565b15156114fd57600080fd5b611510611509826110c6565b3383611d51565b50565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561156f57600080fd5b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6115bb6112a3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156115f457600080fd5b6115ff838383611d51565b505050565b61160c6112a3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561164557600080fd5b61164e81611e9a565b61165781611efb565b50565b61166381611a92565b151561166e57600080fd5b6116788282611496565b5050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156116d857600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561171457600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6117dc61287b565b6117e46112a3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561181d57600080fd5b611826826120f7565b9050919050565b6000813373ffffffffffffffffffffffffffffffffffffffff16611850826110c6565b73ffffffffffffffffffffffffffffffffffffffff1614151561187257600080fd5b61187b836110c6565b91508173ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141515156118b857600080fd5b60006118c384610db7565b73ffffffffffffffffffffffffffffffffffffffff161415806118fd575060008473ffffffffffffffffffffffffffffffffffffffff1614155b156119ba57836002600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925856040518082815260200191505060405180910390a35b50505050565b6119c861287b565b6119d182611a92565b15156119dc57600080fd5b600760008381526020019081526020016000206004018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611a865780601f10611a5b57610100808354040283529160200191611a86565b820191906000526020600020905b815481529060010190602001808311611a6957829003601f168201915b50505050509050919050565b60006007600083815260200190815260200160002060050160009054906101000a900460ff169050919050565b611ac882611a92565b1515611ad357600080fd5b8060076000848152602001908152602001600020600301819055505050565b6000808284019050838110151515611b0657fe5b8091505092915050565b611b186112a3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611b5157600080fd5b611b5a81611a92565b1515611b6557600080fd5b611b6f82826121c9565b5050565b611b7b61288f565b611b8486612263565b151515611b9057600080fd5b60c0604051908101604052808781526020018681526020018581526020018481526020018381526020016001151581525090508060076000888152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004019080519060200190611c1b9291906128e2565b5060a08201518160050160006101000a81548160ff021916908315150217905550905050505050505050565b611c5082611a92565b1515611c5b57600080fd5b80600760008481526020019081526020016000206004019080519060200190611c85929190612962565b505050565b611c9382611a92565b1515611c9e57600080fd5b8060076000848152602001908152602001600020600201819055505050565b803373ffffffffffffffffffffffffffffffffffffffff16611cde826110c6565b73ffffffffffffffffffffffffffffffffffffffff16141515611d0057600080fd5b611d0b338484611d51565b505050565b60008273ffffffffffffffffffffffffffffffffffffffff16611d3283610db7565b73ffffffffffffffffffffffffffffffffffffffff1614905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515611d8d57600080fd5b611d96816110c6565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515611dd057600080fd5b8273ffffffffffffffffffffffffffffffffffffffff16611df0826110c6565b73ffffffffffffffffffffffffffffffffffffffff16141515611e1257600080fd5b611e1c8382612285565b611e26838261236e565b611e3082826125f6565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b611ea26112a3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611edb57600080fd5b611ee481611a92565b1515611eef57600080fd5b611ef881612761565b50565b611f0361288f565b611f0c82611a92565b1515611f1757600080fd5b6007600083815260200190815260200160002060c0604051908101604052908160008201548152602001600182015481526020016002820154815260200160038201548152602001600482018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611ff75780601f10611fcc57610100808354040283529160200191611ff7565b820191906000526020600020905b815481529060010190602001808311611fda57829003601f168201915b505050505081526020016005820160009054906101000a900460ff161515151581525050905080600860008481526020019081526020016000206000820151816000015560208201518160010155604082015181600201556060820151816003015560808201518160040190805190602001906120759291906128e2565b5060a08201518160050160006101000a81548160ff021916908315150217905550905050600760008381526020019081526020016000206000808201600090556001820160009055600282016000905560038201600090556004820160006120dd91906129e2565b6005820160006101000a81549060ff021916905550505050565b6120ff61287b565b61210882612835565b151561211357600080fd5b600860008381526020019081526020016000206004018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156121bd5780601f10612192576101008083540402835291602001916121bd565b820191906000526020600020905b8154815290600101906020018083116121a057829003601f168201915b50505050509050919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561220557600080fd5b61220f82826125f6565b8173ffffffffffffffffffffffffffffffffffffffff1660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600061226e82611a92565b8061227e575061227d82612835565b5b9050919050565b8173ffffffffffffffffffffffffffffffffffffffff166122a5826110c6565b73ffffffffffffffffffffffffffffffffffffffff161415156122c757600080fd5b60006002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a35050565b60008060008473ffffffffffffffffffffffffffffffffffffffff16612393856110c6565b73ffffffffffffffffffffffffffffffffffffffff161415156123b557600080fd5b600460008581526020019081526020016000205492506123e760016123d987611193565b61286290919063ffffffff16565b9150600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208281548110151561243557fe5b906000526020600020900154905060006001600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020848154811015156124e357fe5b9060005260206000209001819055506000600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208381548110151561254057fe5b906000526020600020900181905550600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054809190600190036125a19190612a2a565b50600060046000868152602001908152602001600020819055508260046000838152602001908152602001600020819055506125e9600160005461286290919063ffffffff16565b6000819055505050505050565b60008073ffffffffffffffffffffffffffffffffffffffff166001600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561266557600080fd5b826001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506126c083611193565b9050600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060010182816127139190612a56565b9160005260206000209001600084909190915055508060046000848152602001908152602001600020819055506127566001600054611af290919063ffffffff16565b600081905550505050565b803373ffffffffffffffffffffffffffffffffffffffff16612782826110c6565b73ffffffffffffffffffffffffffffffffffffffff161415156127a457600080fd5b60006127af83610db7565b73ffffffffffffffffffffffffffffffffffffffff161415156127d7576127d63383612285565b5b6127e1338361236e565b60003373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a35050565b60006008600083815260200190815260200160002060050160009054906101000a900460ff169050919050565b600082821115151561287057fe5b818303905092915050565b602060405190810160405280600081525090565b60c060405190810160405280600081526020016000815260200160008152602001600081526020016128bf612a82565b81526020016000151581525090565b602060405190810160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061292357805160ff1916838001178555612951565b82800160010185558215612951579182015b82811115612950578251825591602001919060010190612935565b5b50905061295e9190612a96565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106129a357805160ff19168380011785556129d1565b828001600101855582156129d1579182015b828111156129d05782518255916020019190600101906129b5565b5b5090506129de9190612a96565b5090565b50805460018160011615610100020316600290046000825580601f10612a085750612a27565b601f016020900490600052602060002090810190612a269190612a96565b5b50565b815481835581811511612a5157818360005260206000209182019101612a509190612a96565b5b505050565b815481835581811511612a7d57818360005260206000209182019101612a7c9190612a96565b5b505050565b602060405190810160405280600081525090565b612ab891905b80821115612ab4576000816000905550600101612a9c565b5090565b905600a165627a7a7230582095c813245505025373854ee3f7d51588bf44134b1885e2c53764cc4d05f33c370029`

// DeployFuxToken deploys a new Ethereum contract, binding an instance of FuxToken to it.
func DeployFuxToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FuxToken, error) {
	parsed, err := abi.JSON(strings.NewReader(FuxTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FuxTokenBin), backend)
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

// ApprovedFor is a free data retrieval call binding the contract method 0x2a6dd48f.
//
// Solidity: function approvedFor(_tokenId uint256) constant returns(address)
func (_FuxToken *FuxTokenCaller) ApprovedFor(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "approvedFor", _tokenId)
	return *ret0, err
}

// ApprovedFor is a free data retrieval call binding the contract method 0x2a6dd48f.
//
// Solidity: function approvedFor(_tokenId uint256) constant returns(address)
func (_FuxToken *FuxTokenSession) ApprovedFor(_tokenId *big.Int) (common.Address, error) {
	return _FuxToken.Contract.ApprovedFor(&_FuxToken.CallOpts, _tokenId)
}

// ApprovedFor is a free data retrieval call binding the contract method 0x2a6dd48f.
//
// Solidity: function approvedFor(_tokenId uint256) constant returns(address)
func (_FuxToken *FuxTokenCallerSession) ApprovedFor(_tokenId *big.Int) (common.Address, error) {
	return _FuxToken.Contract.ApprovedFor(&_FuxToken.CallOpts, _tokenId)
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

// GetFuxAttrs is a free data retrieval call binding the contract method 0x1c61984b.
//
// Solidity: function getFuxAttrs(_tokenId uint256) constant returns(coin uint256, lockTime uint256, state uint256)
func (_FuxToken *FuxTokenCaller) GetFuxAttrs(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Coin     *big.Int
	LockTime *big.Int
	State    *big.Int
}, error) {
	ret := new(struct {
		Coin     *big.Int
		LockTime *big.Int
		State    *big.Int
	})
	out := ret
	err := _FuxToken.contract.Call(opts, out, "getFuxAttrs", _tokenId)
	return *ret, err
}

// GetFuxAttrs is a free data retrieval call binding the contract method 0x1c61984b.
//
// Solidity: function getFuxAttrs(_tokenId uint256) constant returns(coin uint256, lockTime uint256, state uint256)
func (_FuxToken *FuxTokenSession) GetFuxAttrs(_tokenId *big.Int) (struct {
	Coin     *big.Int
	LockTime *big.Int
	State    *big.Int
}, error) {
	return _FuxToken.Contract.GetFuxAttrs(&_FuxToken.CallOpts, _tokenId)
}

// GetFuxAttrs is a free data retrieval call binding the contract method 0x1c61984b.
//
// Solidity: function getFuxAttrs(_tokenId uint256) constant returns(coin uint256, lockTime uint256, state uint256)
func (_FuxToken *FuxTokenCallerSession) GetFuxAttrs(_tokenId *big.Int) (struct {
	Coin     *big.Int
	LockTime *big.Int
	State    *big.Int
}, error) {
	return _FuxToken.Contract.GetFuxAttrs(&_FuxToken.CallOpts, _tokenId)
}

// GetFuxIssuer is a free data retrieval call binding the contract method 0x904c00c1.
//
// Solidity: function getFuxIssuer() constant returns(issuer address)
func (_FuxToken *FuxTokenCaller) GetFuxIssuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "getFuxIssuer")
	return *ret0, err
}

// GetFuxIssuer is a free data retrieval call binding the contract method 0x904c00c1.
//
// Solidity: function getFuxIssuer() constant returns(issuer address)
func (_FuxToken *FuxTokenSession) GetFuxIssuer() (common.Address, error) {
	return _FuxToken.Contract.GetFuxIssuer(&_FuxToken.CallOpts)
}

// GetFuxIssuer is a free data retrieval call binding the contract method 0x904c00c1.
//
// Solidity: function getFuxIssuer() constant returns(issuer address)
func (_FuxToken *FuxTokenCallerSession) GetFuxIssuer() (common.Address, error) {
	return _FuxToken.Contract.GetFuxIssuer(&_FuxToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FuxToken *FuxTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FuxToken *FuxTokenSession) Name() (string, error) {
	return _FuxToken.Contract.Name(&_FuxToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) Name() (string, error) {
	return _FuxToken.Contract.Name(&_FuxToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FuxToken *FuxTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FuxToken *FuxTokenSession) Owner() (common.Address, error) {
	return _FuxToken.Contract.Owner(&_FuxToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FuxToken *FuxTokenCallerSession) Owner() (common.Address, error) {
	return _FuxToken.Contract.Owner(&_FuxToken.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_FuxToken *FuxTokenCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_FuxToken *FuxTokenSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _FuxToken.Contract.OwnerOf(&_FuxToken.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_FuxToken *FuxTokenCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _FuxToken.Contract.OwnerOf(&_FuxToken.CallOpts, _tokenId)
}

// QueryFuxInfo is a free data retrieval call binding the contract method 0x0d18fc1f.
//
// Solidity: function queryFuxInfo(_tokenId uint256) constant returns(info string)
func (_FuxToken *FuxTokenCaller) QueryFuxInfo(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "queryFuxInfo", _tokenId)
	return *ret0, err
}

// QueryFuxInfo is a free data retrieval call binding the contract method 0x0d18fc1f.
//
// Solidity: function queryFuxInfo(_tokenId uint256) constant returns(info string)
func (_FuxToken *FuxTokenSession) QueryFuxInfo(_tokenId *big.Int) (string, error) {
	return _FuxToken.Contract.QueryFuxInfo(&_FuxToken.CallOpts, _tokenId)
}

// QueryFuxInfo is a free data retrieval call binding the contract method 0x0d18fc1f.
//
// Solidity: function queryFuxInfo(_tokenId uint256) constant returns(info string)
func (_FuxToken *FuxTokenCallerSession) QueryFuxInfo(_tokenId *big.Int) (string, error) {
	return _FuxToken.Contract.QueryFuxInfo(&_FuxToken.CallOpts, _tokenId)
}

// QueryHisFuxInfo is a free data retrieval call binding the contract method 0xf449b171.
//
// Solidity: function queryHisFuxInfo(_tokenId uint256) constant returns(info string)
func (_FuxToken *FuxTokenCaller) QueryHisFuxInfo(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "queryHisFuxInfo", _tokenId)
	return *ret0, err
}

// QueryHisFuxInfo is a free data retrieval call binding the contract method 0xf449b171.
//
// Solidity: function queryHisFuxInfo(_tokenId uint256) constant returns(info string)
func (_FuxToken *FuxTokenSession) QueryHisFuxInfo(_tokenId *big.Int) (string, error) {
	return _FuxToken.Contract.QueryHisFuxInfo(&_FuxToken.CallOpts, _tokenId)
}

// QueryHisFuxInfo is a free data retrieval call binding the contract method 0xf449b171.
//
// Solidity: function queryHisFuxInfo(_tokenId uint256) constant returns(info string)
func (_FuxToken *FuxTokenCallerSession) QueryHisFuxInfo(_tokenId *big.Int) (string, error) {
	return _FuxToken.Contract.QueryHisFuxInfo(&_FuxToken.CallOpts, _tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FuxToken *FuxTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FuxToken *FuxTokenSession) Symbol() (string, error) {
	return _FuxToken.Contract.Symbol(&_FuxToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FuxToken *FuxTokenCallerSession) Symbol() (string, error) {
	return _FuxToken.Contract.Symbol(&_FuxToken.CallOpts)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(tokenId uint256)
func (_FuxToken *FuxTokenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(tokenId uint256)
func (_FuxToken *FuxTokenSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _FuxToken.Contract.TokenOfOwnerByIndex(&_FuxToken.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(tokenId uint256)
func (_FuxToken *FuxTokenCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _FuxToken.Contract.TokenOfOwnerByIndex(&_FuxToken.CallOpts, _owner, _index)
}

// TokensOf is a free data retrieval call binding the contract method 0x5a3f2672.
//
// Solidity: function tokensOf(_owner address) constant returns(uint256[])
func (_FuxToken *FuxTokenCaller) TokensOf(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "tokensOf", _owner)
	return *ret0, err
}

// TokensOf is a free data retrieval call binding the contract method 0x5a3f2672.
//
// Solidity: function tokensOf(_owner address) constant returns(uint256[])
func (_FuxToken *FuxTokenSession) TokensOf(_owner common.Address) ([]*big.Int, error) {
	return _FuxToken.Contract.TokensOf(&_FuxToken.CallOpts, _owner)
}

// TokensOf is a free data retrieval call binding the contract method 0x5a3f2672.
//
// Solidity: function tokensOf(_owner address) constant returns(uint256[])
func (_FuxToken *FuxTokenCallerSession) TokensOf(_owner common.Address) ([]*big.Int, error) {
	return _FuxToken.Contract.TokensOf(&_FuxToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_FuxToken *FuxTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FuxToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_FuxToken *FuxTokenSession) TotalSupply() (*big.Int, error) {
	return _FuxToken.Contract.TotalSupply(&_FuxToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_FuxToken *FuxTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _FuxToken.Contract.TotalSupply(&_FuxToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.Approve(&_FuxToken.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.Approve(&_FuxToken.TransactOpts, _to, _tokenId)
}

// CreateFux is a paid mutator transaction binding the contract method 0x9a43d9d1.
//
// Solidity: function createFux(_to address, _tokenId uint256, _coin uint256, _lockTime uint256, _state uint256, _otherInfo string) returns()
func (_FuxToken *FuxTokenTransactor) CreateFux(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int, _coin *big.Int, _lockTime *big.Int, _state *big.Int, _otherInfo string) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "createFux", _to, _tokenId, _coin, _lockTime, _state, _otherInfo)
}

// CreateFux is a paid mutator transaction binding the contract method 0x9a43d9d1.
//
// Solidity: function createFux(_to address, _tokenId uint256, _coin uint256, _lockTime uint256, _state uint256, _otherInfo string) returns()
func (_FuxToken *FuxTokenSession) CreateFux(_to common.Address, _tokenId *big.Int, _coin *big.Int, _lockTime *big.Int, _state *big.Int, _otherInfo string) (*types.Transaction, error) {
	return _FuxToken.Contract.CreateFux(&_FuxToken.TransactOpts, _to, _tokenId, _coin, _lockTime, _state, _otherInfo)
}

// CreateFux is a paid mutator transaction binding the contract method 0x9a43d9d1.
//
// Solidity: function createFux(_to address, _tokenId uint256, _coin uint256, _lockTime uint256, _state uint256, _otherInfo string) returns()
func (_FuxToken *FuxTokenTransactorSession) CreateFux(_to common.Address, _tokenId *big.Int, _coin *big.Int, _lockTime *big.Int, _state *big.Int, _otherInfo string) (*types.Transaction, error) {
	return _FuxToken.Contract.CreateFux(&_FuxToken.TransactOpts, _to, _tokenId, _coin, _lockTime, _state, _otherInfo)
}

// RemoveFux is a paid mutator transaction binding the contract method 0xe234f61d.
//
// Solidity: function removeFux(_tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactor) RemoveFux(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "removeFux", _tokenId)
}

// RemoveFux is a paid mutator transaction binding the contract method 0xe234f61d.
//
// Solidity: function removeFux(_tokenId uint256) returns()
func (_FuxToken *FuxTokenSession) RemoveFux(_tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.RemoveFux(&_FuxToken.TransactOpts, _tokenId)
}

// RemoveFux is a paid mutator transaction binding the contract method 0xe234f61d.
//
// Solidity: function removeFux(_tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) RemoveFux(_tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.RemoveFux(&_FuxToken.TransactOpts, _tokenId)
}

// SetFuxIssuer is a paid mutator transaction binding the contract method 0xc0b43a45.
//
// Solidity: function setFuxIssuer(_fuxIssuer address) returns()
func (_FuxToken *FuxTokenTransactor) SetFuxIssuer(opts *bind.TransactOpts, _fuxIssuer common.Address) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "setFuxIssuer", _fuxIssuer)
}

// SetFuxIssuer is a paid mutator transaction binding the contract method 0xc0b43a45.
//
// Solidity: function setFuxIssuer(_fuxIssuer address) returns()
func (_FuxToken *FuxTokenSession) SetFuxIssuer(_fuxIssuer common.Address) (*types.Transaction, error) {
	return _FuxToken.Contract.SetFuxIssuer(&_FuxToken.TransactOpts, _fuxIssuer)
}

// SetFuxIssuer is a paid mutator transaction binding the contract method 0xc0b43a45.
//
// Solidity: function setFuxIssuer(_fuxIssuer address) returns()
func (_FuxToken *FuxTokenTransactorSession) SetFuxIssuer(_fuxIssuer common.Address) (*types.Transaction, error) {
	return _FuxToken.Contract.SetFuxIssuer(&_FuxToken.TransactOpts, _fuxIssuer)
}

// SplitFux is a paid mutator transaction binding the contract method 0x5209e984.
//
// Solidity: function splitFux(_tokenId uint256, _newIds uint256[2], _newCoins uint256[2], _states uint256[2]) returns()
func (_FuxToken *FuxTokenTransactor) SplitFux(opts *bind.TransactOpts, _tokenId *big.Int, _newIds [2]*big.Int, _newCoins [2]*big.Int, _states [2]*big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "splitFux", _tokenId, _newIds, _newCoins, _states)
}

// SplitFux is a paid mutator transaction binding the contract method 0x5209e984.
//
// Solidity: function splitFux(_tokenId uint256, _newIds uint256[2], _newCoins uint256[2], _states uint256[2]) returns()
func (_FuxToken *FuxTokenSession) SplitFux(_tokenId *big.Int, _newIds [2]*big.Int, _newCoins [2]*big.Int, _states [2]*big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.SplitFux(&_FuxToken.TransactOpts, _tokenId, _newIds, _newCoins, _states)
}

// SplitFux is a paid mutator transaction binding the contract method 0x5209e984.
//
// Solidity: function splitFux(_tokenId uint256, _newIds uint256[2], _newCoins uint256[2], _states uint256[2]) returns()
func (_FuxToken *FuxTokenTransactorSession) SplitFux(_tokenId *big.Int, _newIds [2]*big.Int, _newCoins [2]*big.Int, _states [2]*big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.SplitFux(&_FuxToken.TransactOpts, _tokenId, _newIds, _newCoins, _states)
}

// TakeOwnership is a paid mutator transaction binding the contract method 0xb2e6ceeb.
//
// Solidity: function takeOwnership(_tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactor) TakeOwnership(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "takeOwnership", _tokenId)
}

// TakeOwnership is a paid mutator transaction binding the contract method 0xb2e6ceeb.
//
// Solidity: function takeOwnership(_tokenId uint256) returns()
func (_FuxToken *FuxTokenSession) TakeOwnership(_tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.TakeOwnership(&_FuxToken.TransactOpts, _tokenId)
}

// TakeOwnership is a paid mutator transaction binding the contract method 0xb2e6ceeb.
//
// Solidity: function takeOwnership(_tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) TakeOwnership(_tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.TakeOwnership(&_FuxToken.TransactOpts, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.Transfer(&_FuxToken.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.Transfer(&_FuxToken.TransactOpts, _to, _tokenId)
}

// TransferEx is a paid mutator transaction binding the contract method 0xd9cacff2.
//
// Solidity: function transferEx(_from address, _to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactor) TransferEx(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "transferEx", _from, _to, _tokenId)
}

// TransferEx is a paid mutator transaction binding the contract method 0xd9cacff2.
//
// Solidity: function transferEx(_from address, _to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenSession) TransferEx(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.TransferEx(&_FuxToken.TransactOpts, _from, _to, _tokenId)
}

// TransferEx is a paid mutator transaction binding the contract method 0xd9cacff2.
//
// Solidity: function transferEx(_from address, _to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) TransferEx(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.TransferEx(&_FuxToken.TransactOpts, _from, _to, _tokenId)
}

// TransferFux is a paid mutator transaction binding the contract method 0xea9724ae.
//
// Solidity: function transferFux(_to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactor) TransferFux(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "transferFux", _to, _tokenId)
}

// TransferFux is a paid mutator transaction binding the contract method 0xea9724ae.
//
// Solidity: function transferFux(_to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenSession) TransferFux(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.TransferFux(&_FuxToken.TransactOpts, _to, _tokenId)
}

// TransferFux is a paid mutator transaction binding the contract method 0xea9724ae.
//
// Solidity: function transferFux(_to address, _tokenId uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) TransferFux(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.TransferFux(&_FuxToken.TransactOpts, _to, _tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FuxToken *FuxTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FuxToken *FuxTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FuxToken.Contract.TransferOwnership(&_FuxToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_FuxToken *FuxTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FuxToken.Contract.TransferOwnership(&_FuxToken.TransactOpts, newOwner)
}

// UpdateFuxCoin is a paid mutator transaction binding the contract method 0x462b7df5.
//
// Solidity: function updateFuxCoin(_tokenId uint256, _newNum uint256) returns()
func (_FuxToken *FuxTokenTransactor) UpdateFuxCoin(opts *bind.TransactOpts, _tokenId *big.Int, _newNum *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "updateFuxCoin", _tokenId, _newNum)
}

// UpdateFuxCoin is a paid mutator transaction binding the contract method 0x462b7df5.
//
// Solidity: function updateFuxCoin(_tokenId uint256, _newNum uint256) returns()
func (_FuxToken *FuxTokenSession) UpdateFuxCoin(_tokenId *big.Int, _newNum *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.UpdateFuxCoin(&_FuxToken.TransactOpts, _tokenId, _newNum)
}

// UpdateFuxCoin is a paid mutator transaction binding the contract method 0x462b7df5.
//
// Solidity: function updateFuxCoin(_tokenId uint256, _newNum uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) UpdateFuxCoin(_tokenId *big.Int, _newNum *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.UpdateFuxCoin(&_FuxToken.TransactOpts, _tokenId, _newNum)
}

// UpdateFuxInfo is a paid mutator transaction binding the contract method 0x6eab2e63.
//
// Solidity: function updateFuxInfo(_tokenId uint256, _otherInfo string) returns()
func (_FuxToken *FuxTokenTransactor) UpdateFuxInfo(opts *bind.TransactOpts, _tokenId *big.Int, _otherInfo string) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "updateFuxInfo", _tokenId, _otherInfo)
}

// UpdateFuxInfo is a paid mutator transaction binding the contract method 0x6eab2e63.
//
// Solidity: function updateFuxInfo(_tokenId uint256, _otherInfo string) returns()
func (_FuxToken *FuxTokenSession) UpdateFuxInfo(_tokenId *big.Int, _otherInfo string) (*types.Transaction, error) {
	return _FuxToken.Contract.UpdateFuxInfo(&_FuxToken.TransactOpts, _tokenId, _otherInfo)
}

// UpdateFuxInfo is a paid mutator transaction binding the contract method 0x6eab2e63.
//
// Solidity: function updateFuxInfo(_tokenId uint256, _otherInfo string) returns()
func (_FuxToken *FuxTokenTransactorSession) UpdateFuxInfo(_tokenId *big.Int, _otherInfo string) (*types.Transaction, error) {
	return _FuxToken.Contract.UpdateFuxInfo(&_FuxToken.TransactOpts, _tokenId, _otherInfo)
}

// UpdateFuxLockTime is a paid mutator transaction binding the contract method 0x72fb7d0b.
//
// Solidity: function updateFuxLockTime(_tokenId uint256, _newNum uint256) returns()
func (_FuxToken *FuxTokenTransactor) UpdateFuxLockTime(opts *bind.TransactOpts, _tokenId *big.Int, _newNum *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "updateFuxLockTime", _tokenId, _newNum)
}

// UpdateFuxLockTime is a paid mutator transaction binding the contract method 0x72fb7d0b.
//
// Solidity: function updateFuxLockTime(_tokenId uint256, _newNum uint256) returns()
func (_FuxToken *FuxTokenSession) UpdateFuxLockTime(_tokenId *big.Int, _newNum *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.UpdateFuxLockTime(&_FuxToken.TransactOpts, _tokenId, _newNum)
}

// UpdateFuxLockTime is a paid mutator transaction binding the contract method 0x72fb7d0b.
//
// Solidity: function updateFuxLockTime(_tokenId uint256, _newNum uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) UpdateFuxLockTime(_tokenId *big.Int, _newNum *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.UpdateFuxLockTime(&_FuxToken.TransactOpts, _tokenId, _newNum)
}

// UpdateFuxState is a paid mutator transaction binding the contract method 0x736cd364.
//
// Solidity: function updateFuxState(_tokenId uint256, _newNum uint256) returns()
func (_FuxToken *FuxTokenTransactor) UpdateFuxState(opts *bind.TransactOpts, _tokenId *big.Int, _newNum *big.Int) (*types.Transaction, error) {
	return _FuxToken.contract.Transact(opts, "updateFuxState", _tokenId, _newNum)
}

// UpdateFuxState is a paid mutator transaction binding the contract method 0x736cd364.
//
// Solidity: function updateFuxState(_tokenId uint256, _newNum uint256) returns()
func (_FuxToken *FuxTokenSession) UpdateFuxState(_tokenId *big.Int, _newNum *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.UpdateFuxState(&_FuxToken.TransactOpts, _tokenId, _newNum)
}

// UpdateFuxState is a paid mutator transaction binding the contract method 0x736cd364.
//
// Solidity: function updateFuxState(_tokenId uint256, _newNum uint256) returns()
func (_FuxToken *FuxTokenTransactorSession) UpdateFuxState(_tokenId *big.Int, _newNum *big.Int) (*types.Transaction, error) {
	return _FuxToken.Contract.UpdateFuxState(&_FuxToken.TransactOpts, _tokenId, _newNum)
}

// FuxTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FuxToken contract.
type FuxTokenApprovalIterator struct {
	Event *FuxTokenApproval // Event containing the contract specifics and raw log

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
func (it *FuxTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxTokenApproval)
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
		it.Event = new(FuxTokenApproval)
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
func (it *FuxTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxTokenApproval represents a Approval event raised by the FuxToken contract.
type FuxTokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_FuxToken *FuxTokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*FuxTokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _FuxToken.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &FuxTokenApprovalIterator{contract: _FuxToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(_owner indexed address, _approved indexed address, _tokenId uint256)
func (_FuxToken *FuxTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FuxTokenApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _FuxToken.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxTokenApproval)
				if err := _FuxToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// FuxTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FuxToken contract.
type FuxTokenOwnershipTransferredIterator struct {
	Event *FuxTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FuxTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxTokenOwnershipTransferred)
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
		it.Event = new(FuxTokenOwnershipTransferred)
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
func (it *FuxTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxTokenOwnershipTransferred represents a OwnershipTransferred event raised by the FuxToken contract.
type FuxTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_FuxToken *FuxTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FuxTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FuxToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FuxTokenOwnershipTransferredIterator{contract: _FuxToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_FuxToken *FuxTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FuxTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FuxToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxTokenOwnershipTransferred)
				if err := _FuxToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// FuxTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FuxToken contract.
type FuxTokenTransferIterator struct {
	Event *FuxTokenTransfer // Event containing the contract specifics and raw log

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
func (it *FuxTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuxTokenTransfer)
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
		it.Event = new(FuxTokenTransfer)
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
func (it *FuxTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuxTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuxTokenTransfer represents a Transfer event raised by the FuxToken contract.
type FuxTokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_FuxToken *FuxTokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*FuxTokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _FuxToken.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &FuxTokenTransferIterator{contract: _FuxToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _tokenId uint256)
func (_FuxToken *FuxTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FuxTokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _FuxToken.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuxTokenTransfer)
				if err := _FuxToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
