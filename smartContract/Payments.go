// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PaymentsABI is the input ABI used to generate the binding from.
const PaymentsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"externalOwnerId\",\"type\":\"uint256\"}],\"name\":\"getPaymentsByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"clientId\",\"type\":\"uint256\"}],\"name\":\"getAClient\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"clientId\",\"type\":\"uint256\"}],\"name\":\"checkIfClientExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"externalPaymentId\",\"type\":\"uint256\"}],\"name\":\"payPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"externalPaymentId\",\"type\":\"uint256\"}],\"name\":\"getPaymentById\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sellPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payments\",\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"receiverId\",\"type\":\"uint256\"},{\"name\":\"payerId\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newSellPrice\",\"type\":\"uint256\"}],\"name\":\"setPrices\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userName\",\"type\":\"string\"},{\"name\":\"file\",\"type\":\"string\"},{\"name\":\"clientAccountAddress\",\"type\":\"string\"}],\"name\":\"aNewClient\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"externalPayerId\",\"type\":\"uint256\"},{\"name\":\"externalReceiverId\",\"type\":\"uint256\"}],\"name\":\"startNewPayment\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"clients\",\"outputs\":[{\"name\":\"userName\",\"type\":\"string\"},{\"name\":\"file\",\"type\":\"string\"},{\"name\":\"clientAccountAddress\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"name\":\"tokenInitialBalance\",\"type\":\"uint256\"},{\"name\":\"tokenDevAddress\",\"type\":\"address\"},{\"name\":\"setSellPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"externalPaymentId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"externalPayerId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"receiverId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"NewPaymentEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"externalClientId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"userName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"file\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"clientAccountAddress\",\"type\":\"string\"}],\"name\":\"NewClientEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// PaymentsBin is the compiled bytecode used for deploying new contracts.
const PaymentsBin = `60806040523480156200001157600080fd5b506040516200284b3803806200284b833981810160405260a08110156200003757600080fd5b8101908080516401000000008111156200005057600080fd5b828101905060208101848111156200006757600080fd5b81518560018202830111640100000000821117156200008557600080fd5b50509291906020018051640100000000811115620000a257600080fd5b82810190506020810184811115620000b957600080fd5b8151856001820283011164010000000082111715620000d757600080fd5b50509291906020018051906020019092919080519060200190929190805190602001909291905050508484848484336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600190805190602001906200015d9291906200024b565b508360029080519060200190620001769291906200024b565b5081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008390508060076000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080600460008282540192505081905550816006819055505050505050505050505050620002fa565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200028e57805160ff1916838001178555620002bf565b82800160010185558215620002bf579182015b82811115620002be578251825591602001919060010190620002a1565b5b509050620002ce9190620002d2565b5090565b620002f791905b80821115620002f3576000816000905550600101620002d9565b5090565b90565b612541806200030a6000396000f3fe6080604052600436106101345760003560e01c806387d81789116100ab578063a9059cbb1161006f578063a9059cbb1461080e578063bdf3ed9314610881578063e4849b3214610a8b578063f2fde38b14610ac6578063f48cef2014610b17578063f88af21d14610b7a57610134565b806387d817891461067e5780638da5cb5b146106e257806395d89b4114610739578063a3201daa146107c9578063a6f2ae3a1461080457610134565b80631e18ea8e116100fd5780631e18ea8e146104cb57806323ed5d001461051e578063313ce567146105595780633de73ebe1461058a5780634b750334146105ee57806370a082311461061957610134565b8062394bb01461013957806302ae172d1461021157806306fdde031461039d578063095ea7b31461042d57806318160ddd146104a0575b600080fd5b34801561014557600080fd5b506101726004803603602081101561015c57600080fd5b8101908080359060200190929190505050610d06565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156101b957808201518184015260208101905061019e565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156101fb5780820151818401526020810190506101e0565b5050505090500194505050505060405180910390f35b34801561021d57600080fd5b5061024a6004803603602081101561023457600080fd5b8101908080359060200190929190505050610ef2565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b83811015610292578082015181840152602081019050610277565b50505050905090810190601f1680156102bf5780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b838110156102f85780820151818401526020810190506102dd565b50505050905090810190601f1680156103255780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b8381101561035e578082015181840152602081019050610343565b50505050905090810190601f16801561038b5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b3480156103a957600080fd5b506103b261113a565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103f25780820151818401526020810190506103d7565b50505050905090810190601f16801561041f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561043957600080fd5b506104866004803603604081101561045057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506111d8565b604051808215151515815260200191505060405180910390f35b3480156104ac57600080fd5b506104b56112ca565b6040518082815260200191505060405180910390f35b3480156104d757600080fd5b50610504600480360360208110156104ee57600080fd5b81019080803590602001909291905050506112d0565b604051808215151515815260200191505060405180910390f35b34801561052a57600080fd5b506105576004803603602081101561054157600080fd5b81019080803590602001909291905050506113e7565b005b34801561056557600080fd5b5061056e61157f565b604051808260ff1660ff16815260200191505060405180910390f35b34801561059657600080fd5b506105c3600480360360208110156105ad57600080fd5b8101908080359060200190929190505050611592565b6040518085815260200184815260200183815260200182815260200194505050505060405180910390f35b3480156105fa57600080fd5b5061060361163b565b6040518082815260200191505060405180910390f35b34801561062557600080fd5b506106686004803603602081101561063c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611641565b6040518082815260200191505060405180910390f35b34801561068a57600080fd5b506106b7600480360360208110156106a157600080fd5b810190808035906020019092919050505061168a565b6040518085815260200184815260200183815260200182815260200194505050505060405180910390f35b3480156106ee57600080fd5b506106f76116c7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561074557600080fd5b5061074e6116ec565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561078e578082015181840152602081019050610773565b50505050905090810190601f1680156107bb5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156107d557600080fd5b50610802600480360360208110156107ec57600080fd5b810190808035906020019092919050505061178a565b005b61080c6117ed565b005b34801561081a57600080fd5b506108676004803603604081101561083157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611804565b604051808215151515815260200191505060405180910390f35b34801561088d57600080fd5b50610a75600480360360608110156108a457600080fd5b81019080803590602001906401000000008111156108c157600080fd5b8201836020820111156108d357600080fd5b803590602001918460018302840111640100000000831117156108f557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561095857600080fd5b82018360208201111561096a57600080fd5b8035906020019184600183028401116401000000008311171561098c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156109ef57600080fd5b820183602082011115610a0157600080fd5b80359060200191846001830284011164010000000083111715610a2357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061181b565b6040518082815260200191505060405180910390f35b348015610a9757600080fd5b50610ac460048036036020811015610aae57600080fd5b8101908080359060200190929190505050611a4a565b005b348015610ad257600080fd5b50610b1560048036036020811015610ae957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611ad8565b005b348015610b2357600080fd5b50610b6460048036036060811015610b3a57600080fd5b81019080803590602001909291908035906020019092919080359060200190929190505050611b74565b6040518082815260200191505060405180910390f35b348015610b8657600080fd5b50610bb360048036036020811015610b9d57600080fd5b8101908080359060200190929190505050611d72565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b83811015610bfb578082015181840152602081019050610be0565b50505050905090810190601f168015610c285780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b83811015610c61578082015181840152602081019050610c46565b50505050905090810190601f168015610c8e5780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b83811015610cc7578082015181840152602081019050610cac565b50505050905090810190601f168015610cf45780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b6060806000600d600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506060600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051908082528060200260200182016040528015610db15781602001602082028038833980820191505090505b5090506060600c60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051908082528060200260200182016040528015610e245781602001602082028038833980820191505090505b5090506000809050600080905060008090505b600980549050811015610ee1578860098281548110610e5257fe5b9060005260206000209060040201600201541415610e8c5780858481518110610e7757fe5b60200260200101818152505082806001019350505b8860098281548110610e9a57fe5b9060005260206000209060040201600101541415610ed45780848381518110610ebf57fe5b60200260200101818152505081806001019250505b8080600101915050610e37565b508383965096505050505050915091565b606080606080600a8581548110610f0557fe5b90600052602060002090600302016000018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610faa5780601f10610f7f57610100808354040283529160200191610faa565b820191906000526020600020905b815481529060010190602001808311610f8d57829003601f168201915b505050505090506060600a8681548110610fc057fe5b90600052602060002090600302016001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110655780601f1061103a57610100808354040283529160200191611065565b820191906000526020600020905b81548152906001019060200180831161104857829003601f168201915b505050505090506060600a878154811061107b57fe5b90600052602060002090600302016002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111205780601f106110f557610100808354040283529160200191611120565b820191906000526020600020905b81548152906001019060200180831161110357829003601f168201915b505050505090508282829550955095505050509193909250565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111d05780601f106111a5576101008083540402835291602001916111d0565b820191906000526020600020905b8154815290600101906020018083116111b357829003601f168201915b505050505081565b600081600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b60045481565b60006112f36040518060600160405280602a81526020016124e3602a9139611f71565b73ffffffffffffffffffffffffffffffffffffffff166113c8600a848154811061131957fe5b90600052602060002090600302016002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113be5780601f10611393576101008083540402835291602001916113be565b820191906000526020600020905b8154815290600101906020018083116113a157829003601f168201915b5050505050611f71565b73ffffffffffffffffffffffffffffffffffffffff1614159050919050565b6000600982815481106113f657fe5b906000526020600020906004020160000154905060006009838154811061141957fe5b906000526020600020906004020160010154905060006114ee600a838154811061143f57fe5b90600052602060002090600302016002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156114e45780601f106114b9576101008083540402835291602001916114e4565b820191906000526020600020905b8154815290600101906020018083116114c757829003601f168201915b5050505050611f71565b90506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561154a57600080fd5b6115548184611804565b5060026009858154811061156457fe5b90600052602060002090600402016003018190555050505050565b600360009054906101000a900460ff1681565b6000806000806000600986815481106115a757fe5b90600052602060002090600402016000015490506000600987815481106115ca57fe5b90600052602060002090600402016002015490506000600988815481106115ed57fe5b906000526020600020906004020160010154905060006009898154811061161057fe5b9060005260206000209060040201600301549050838383839750975097509750505050509193509193565b60065481565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6009818154811061169757fe5b90600052602060002090600402016000915090508060000154908060010154908060020154908060030154905084565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156117825780601f1061175757610100808354040283529160200191611782565b820191906000526020600020905b81548152906001019060200180831161176557829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146117e357600080fd5b8060068190555050565b6000600654340290506118003082611804565b5050565b60006118113384846121ca565b6001905092915050565b600080600a60405180606001604052808781526020018681526020018581525090806001815401808255809150509060018203906000526020600020906003020160009091929091909150600082015181600001908051906020019061188292919061243d565b50602082015181600101908051906020019061189f92919061243d565b5060408201518160020190805190602001906118bc92919061243d565b50505090507f8f93562878197ea1bc0b4849240129cbe4babdc145f734e6a441b02d891167818186868660405180858152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b83811015611934578082015181840152602081019050611919565b50505050905090810190601f1680156119615780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b8381101561199a57808201518184015260208101905061197f565b50505050905090810190601f1680156119c75780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b83811015611a005780820151818401526020810190506119e5565b50505050905090810190601f168015611a2d5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a1809150509392505050565b60003090506006548281611a5a57fe5b048173ffffffffffffffffffffffffffffffffffffffff16311015611a7e57600080fd5b611a893330846121ca565b3373ffffffffffffffffffffffffffffffffffffffff166108fc60065484029081150290604051600060405180830381858888f19350505050158015611ad3573d6000803e3d6000fd5b505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611b3157600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000808411611b8257600080fd5b60006009604051806080016040528087815260200185815260200186815260200160018152509080600181540180825580915050906001820390600052602060002090600402016000909192909190915060008201518160000155602082015181600101556040820151816002015560608201518160030155505090506000600d600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000600d600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008154809291906001019190505550600c60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081548092919060010191905055507f2f31e803cd06ad25ed25b55a829d47dd89067f047a4257b568d63b1ad5701057838888886001604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390a18293505050509392505050565b600a8181548110611d7f57fe5b9060005260206000209060030201600091509050806000018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611e2b5780601f10611e0057610100808354040283529160200191611e2b565b820191906000526020600020905b815481529060010190602001808311611e0e57829003601f168201915b505050505090806001018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611ec95780601f10611e9e57610100808354040283529160200191611ec9565b820191906000526020600020905b815481529060010190602001808311611eac57829003601f168201915b505050505090806002018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611f675780601f10611f3c57610100808354040283529160200191611f67565b820191906000526020600020905b815481529060010190602001808311611f4a57829003601f168201915b5050505050905083565b6000606082905060008090506000806000600290505b602a8110156121bd5761010084029350848181518110611fa357fe5b602001015160f81c60f81b60f81c60ff169250846001820181518110611fc557fe5b602001015160f81c60f81b60f81c60ff16915060618373ffffffffffffffffffffffffffffffffffffffff1610158015612016575060668373ffffffffffffffffffffffffffffffffffffffff1611155b15612026576057830392506120c0565b60418373ffffffffffffffffffffffffffffffffffffffff1610158015612064575060468373ffffffffffffffffffffffffffffffffffffffff1611155b15612074576037830392506120bf565b60308373ffffffffffffffffffffffffffffffffffffffff16101580156120b2575060398373ffffffffffffffffffffffffffffffffffffffff1611155b156120be576030830392505b5b5b60618273ffffffffffffffffffffffffffffffffffffffff16101580156120fe575060668273ffffffffffffffffffffffffffffffffffffffff1611155b1561210e576057820391506121a8565b60418273ffffffffffffffffffffffffffffffffffffffff161015801561214c575060468273ffffffffffffffffffffffffffffffffffffffff1611155b1561215c576037820391506121a7565b60308273ffffffffffffffffffffffffffffffffffffffff161015801561219a575060398273ffffffffffffffffffffffffffffffffffffffff1611155b156121a6576030820391505b5b5b81601084020184019350600281019050611f87565b5082945050505050919050565b80600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561221657600080fd5b6000811161222357600080fd5b600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401116122af57600080fd5b80600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555080600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555080600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061247e57805160ff19168380011785556124ac565b828001600101855582156124ac579182015b828111156124ab578251825591602001919060010190612490565b5b5090506124b991906124bd565b5090565b6124df91905b808211156124db5760008160009055506001016124c3565b5090565b9056fe307830303030303030303030303030303030303030303030303030303030303030303030303030303030a265627a7a723058200bc060e209f666ea1ff2cc12e37d529d1146511fec6f775db458ed5279c80e9664736f6c63430005090032`

// DeployPayments deploys a new Ethereum contract, binding an instance of Payments to it.
func DeployPayments(auth *bind.TransactOpts, backend bind.ContractBackend, tokenName string, tokenSymbol string, tokenInitialBalance *big.Int, tokenDevAddress common.Address, setSellPrice *big.Int) (common.Address, *types.Transaction, *Payments, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PaymentsBin), backend, tokenName, tokenSymbol, tokenInitialBalance, tokenDevAddress, setSellPrice)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Payments{PaymentsCaller: PaymentsCaller{contract: contract}, PaymentsTransactor: PaymentsTransactor{contract: contract}, PaymentsFilterer: PaymentsFilterer{contract: contract}}, nil
}

// Payments is an auto generated Go binding around an Ethereum contract.
type Payments struct {
	PaymentsCaller     // Read-only binding to the contract
	PaymentsTransactor // Write-only binding to the contract
	PaymentsFilterer   // Log filterer for contract events
}

// PaymentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentsSession struct {
	Contract     *Payments         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentsCallerSession struct {
	Contract *PaymentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PaymentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentsTransactorSession struct {
	Contract     *PaymentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PaymentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentsRaw struct {
	Contract *Payments // Generic contract binding to access the raw methods on
}

// PaymentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentsCallerRaw struct {
	Contract *PaymentsCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentsTransactorRaw struct {
	Contract *PaymentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayments creates a new instance of Payments, bound to a specific deployed contract.
func NewPayments(address common.Address, backend bind.ContractBackend) (*Payments, error) {
	contract, err := bindPayments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Payments{PaymentsCaller: PaymentsCaller{contract: contract}, PaymentsTransactor: PaymentsTransactor{contract: contract}, PaymentsFilterer: PaymentsFilterer{contract: contract}}, nil
}

// NewPaymentsCaller creates a new read-only instance of Payments, bound to a specific deployed contract.
func NewPaymentsCaller(address common.Address, caller bind.ContractCaller) (*PaymentsCaller, error) {
	contract, err := bindPayments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentsCaller{contract: contract}, nil
}

// NewPaymentsTransactor creates a new write-only instance of Payments, bound to a specific deployed contract.
func NewPaymentsTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentsTransactor, error) {
	contract, err := bindPayments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentsTransactor{contract: contract}, nil
}

// NewPaymentsFilterer creates a new log filterer instance of Payments, bound to a specific deployed contract.
func NewPaymentsFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentsFilterer, error) {
	contract, err := bindPayments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentsFilterer{contract: contract}, nil
}

// bindPayments binds a generic wrapper to an already deployed contract.
func bindPayments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payments *PaymentsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Payments.Contract.PaymentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payments *PaymentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payments.Contract.PaymentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payments *PaymentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payments.Contract.PaymentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payments *PaymentsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Payments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payments *PaymentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payments *PaymentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payments.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_Payments *PaymentsCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Payments.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_Payments *PaymentsSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Payments.Contract.BalanceOf(&_Payments.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_Payments *PaymentsCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Payments.Contract.BalanceOf(&_Payments.CallOpts, _owner)
}

// Clients is a free data retrieval call binding the contract method 0xf88af21d.
//
// Solidity: function clients(uint256 ) constant returns(string userName, string file, string clientAccountAddress)
func (_Payments *PaymentsCaller) Clients(opts *bind.CallOpts, arg0 *big.Int) (struct {
	UserName             string
	File                 string
	ClientAccountAddress string
}, error) {
	ret := new(struct {
		UserName             string
		File                 string
		ClientAccountAddress string
	})
	out := ret
	err := _Payments.contract.Call(opts, out, "clients", arg0)
	return *ret, err
}

// Clients is a free data retrieval call binding the contract method 0xf88af21d.
//
// Solidity: function clients(uint256 ) constant returns(string userName, string file, string clientAccountAddress)
func (_Payments *PaymentsSession) Clients(arg0 *big.Int) (struct {
	UserName             string
	File                 string
	ClientAccountAddress string
}, error) {
	return _Payments.Contract.Clients(&_Payments.CallOpts, arg0)
}

// Clients is a free data retrieval call binding the contract method 0xf88af21d.
//
// Solidity: function clients(uint256 ) constant returns(string userName, string file, string clientAccountAddress)
func (_Payments *PaymentsCallerSession) Clients(arg0 *big.Int) (struct {
	UserName             string
	File                 string
	ClientAccountAddress string
}, error) {
	return _Payments.Contract.Clients(&_Payments.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Payments *PaymentsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Payments.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Payments *PaymentsSession) Decimals() (uint8, error) {
	return _Payments.Contract.Decimals(&_Payments.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Payments *PaymentsCallerSession) Decimals() (uint8, error) {
	return _Payments.Contract.Decimals(&_Payments.CallOpts)
}

// GetAClient is a free data retrieval call binding the contract method 0x02ae172d.
//
// Solidity: function getAClient(uint256 clientId) constant returns(string, string, string)
func (_Payments *PaymentsCaller) GetAClient(opts *bind.CallOpts, clientId *big.Int) (string, string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Payments.contract.Call(opts, out, "getAClient", clientId)
	return *ret0, *ret1, *ret2, err
}

// GetAClient is a free data retrieval call binding the contract method 0x02ae172d.
//
// Solidity: function getAClient(uint256 clientId) constant returns(string, string, string)
func (_Payments *PaymentsSession) GetAClient(clientId *big.Int) (string, string, string, error) {
	return _Payments.Contract.GetAClient(&_Payments.CallOpts, clientId)
}

// GetAClient is a free data retrieval call binding the contract method 0x02ae172d.
//
// Solidity: function getAClient(uint256 clientId) constant returns(string, string, string)
func (_Payments *PaymentsCallerSession) GetAClient(clientId *big.Int) (string, string, string, error) {
	return _Payments.Contract.GetAClient(&_Payments.CallOpts, clientId)
}

// GetPaymentById is a free data retrieval call binding the contract method 0x3de73ebe.
//
// Solidity: function getPaymentById(uint256 externalPaymentId) constant returns(uint256, uint256, uint256, uint256)
func (_Payments *PaymentsCaller) GetPaymentById(opts *bind.CallOpts, externalPaymentId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Payments.contract.Call(opts, out, "getPaymentById", externalPaymentId)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetPaymentById is a free data retrieval call binding the contract method 0x3de73ebe.
//
// Solidity: function getPaymentById(uint256 externalPaymentId) constant returns(uint256, uint256, uint256, uint256)
func (_Payments *PaymentsSession) GetPaymentById(externalPaymentId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Payments.Contract.GetPaymentById(&_Payments.CallOpts, externalPaymentId)
}

// GetPaymentById is a free data retrieval call binding the contract method 0x3de73ebe.
//
// Solidity: function getPaymentById(uint256 externalPaymentId) constant returns(uint256, uint256, uint256, uint256)
func (_Payments *PaymentsCallerSession) GetPaymentById(externalPaymentId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Payments.Contract.GetPaymentById(&_Payments.CallOpts, externalPaymentId)
}

// GetPaymentsByOwner is a free data retrieval call binding the contract method 0x00394bb0.
//
// Solidity: function getPaymentsByOwner(uint256 externalOwnerId) constant returns(uint256[], uint256[])
func (_Payments *PaymentsCaller) GetPaymentsByOwner(opts *bind.CallOpts, externalOwnerId *big.Int) ([]*big.Int, []*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Payments.contract.Call(opts, out, "getPaymentsByOwner", externalOwnerId)
	return *ret0, *ret1, err
}

// GetPaymentsByOwner is a free data retrieval call binding the contract method 0x00394bb0.
//
// Solidity: function getPaymentsByOwner(uint256 externalOwnerId) constant returns(uint256[], uint256[])
func (_Payments *PaymentsSession) GetPaymentsByOwner(externalOwnerId *big.Int) ([]*big.Int, []*big.Int, error) {
	return _Payments.Contract.GetPaymentsByOwner(&_Payments.CallOpts, externalOwnerId)
}

// GetPaymentsByOwner is a free data retrieval call binding the contract method 0x00394bb0.
//
// Solidity: function getPaymentsByOwner(uint256 externalOwnerId) constant returns(uint256[], uint256[])
func (_Payments *PaymentsCallerSession) GetPaymentsByOwner(externalOwnerId *big.Int) ([]*big.Int, []*big.Int, error) {
	return _Payments.Contract.GetPaymentsByOwner(&_Payments.CallOpts, externalOwnerId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Payments *PaymentsCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Payments.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Payments *PaymentsSession) Name() (string, error) {
	return _Payments.Contract.Name(&_Payments.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Payments *PaymentsCallerSession) Name() (string, error) {
	return _Payments.Contract.Name(&_Payments.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Payments *PaymentsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Payments.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Payments *PaymentsSession) Owner() (common.Address, error) {
	return _Payments.Contract.Owner(&_Payments.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Payments *PaymentsCallerSession) Owner() (common.Address, error) {
	return _Payments.Contract.Owner(&_Payments.CallOpts)
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) constant returns(uint256 price, uint256 receiverId, uint256 payerId, uint256 status)
func (_Payments *PaymentsCaller) Payments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Price      *big.Int
	ReceiverId *big.Int
	PayerId    *big.Int
	Status     *big.Int
}, error) {
	ret := new(struct {
		Price      *big.Int
		ReceiverId *big.Int
		PayerId    *big.Int
		Status     *big.Int
	})
	out := ret
	err := _Payments.contract.Call(opts, out, "payments", arg0)
	return *ret, err
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) constant returns(uint256 price, uint256 receiverId, uint256 payerId, uint256 status)
func (_Payments *PaymentsSession) Payments(arg0 *big.Int) (struct {
	Price      *big.Int
	ReceiverId *big.Int
	PayerId    *big.Int
	Status     *big.Int
}, error) {
	return _Payments.Contract.Payments(&_Payments.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) constant returns(uint256 price, uint256 receiverId, uint256 payerId, uint256 status)
func (_Payments *PaymentsCallerSession) Payments(arg0 *big.Int) (struct {
	Price      *big.Int
	ReceiverId *big.Int
	PayerId    *big.Int
	Status     *big.Int
}, error) {
	return _Payments.Contract.Payments(&_Payments.CallOpts, arg0)
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_Payments *PaymentsCaller) SellPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Payments.contract.Call(opts, out, "sellPrice")
	return *ret0, err
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_Payments *PaymentsSession) SellPrice() (*big.Int, error) {
	return _Payments.Contract.SellPrice(&_Payments.CallOpts)
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_Payments *PaymentsCallerSession) SellPrice() (*big.Int, error) {
	return _Payments.Contract.SellPrice(&_Payments.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Payments *PaymentsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Payments.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Payments *PaymentsSession) Symbol() (string, error) {
	return _Payments.Contract.Symbol(&_Payments.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Payments *PaymentsCallerSession) Symbol() (string, error) {
	return _Payments.Contract.Symbol(&_Payments.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Payments *PaymentsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Payments.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Payments *PaymentsSession) TotalSupply() (*big.Int, error) {
	return _Payments.Contract.TotalSupply(&_Payments.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Payments *PaymentsCallerSession) TotalSupply() (*big.Int, error) {
	return _Payments.Contract.TotalSupply(&_Payments.CallOpts)
}

// ANewClient is a paid mutator transaction binding the contract method 0xbdf3ed93.
//
// Solidity: function aNewClient(string userName, string file, string clientAccountAddress) returns(uint256)
func (_Payments *PaymentsTransactor) ANewClient(opts *bind.TransactOpts, userName string, file string, clientAccountAddress string) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "aNewClient", userName, file, clientAccountAddress)
}

// ANewClient is a paid mutator transaction binding the contract method 0xbdf3ed93.
//
// Solidity: function aNewClient(string userName, string file, string clientAccountAddress) returns(uint256)
func (_Payments *PaymentsSession) ANewClient(userName string, file string, clientAccountAddress string) (*types.Transaction, error) {
	return _Payments.Contract.ANewClient(&_Payments.TransactOpts, userName, file, clientAccountAddress)
}

// ANewClient is a paid mutator transaction binding the contract method 0xbdf3ed93.
//
// Solidity: function aNewClient(string userName, string file, string clientAccountAddress) returns(uint256)
func (_Payments *PaymentsTransactorSession) ANewClient(userName string, file string, clientAccountAddress string) (*types.Transaction, error) {
	return _Payments.Contract.ANewClient(&_Payments.TransactOpts, userName, file, clientAccountAddress)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool success)
func (_Payments *PaymentsTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool success)
func (_Payments *PaymentsSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.Approve(&_Payments.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool success)
func (_Payments *PaymentsTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.Approve(&_Payments.TransactOpts, _spender, _amount)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_Payments *PaymentsTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_Payments *PaymentsSession) Buy() (*types.Transaction, error) {
	return _Payments.Contract.Buy(&_Payments.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_Payments *PaymentsTransactorSession) Buy() (*types.Transaction, error) {
	return _Payments.Contract.Buy(&_Payments.TransactOpts)
}

// CheckIfClientExists is a paid mutator transaction binding the contract method 0x1e18ea8e.
//
// Solidity: function checkIfClientExists(uint256 clientId) returns(bool)
func (_Payments *PaymentsTransactor) CheckIfClientExists(opts *bind.TransactOpts, clientId *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "checkIfClientExists", clientId)
}

// CheckIfClientExists is a paid mutator transaction binding the contract method 0x1e18ea8e.
//
// Solidity: function checkIfClientExists(uint256 clientId) returns(bool)
func (_Payments *PaymentsSession) CheckIfClientExists(clientId *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.CheckIfClientExists(&_Payments.TransactOpts, clientId)
}

// CheckIfClientExists is a paid mutator transaction binding the contract method 0x1e18ea8e.
//
// Solidity: function checkIfClientExists(uint256 clientId) returns(bool)
func (_Payments *PaymentsTransactorSession) CheckIfClientExists(clientId *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.CheckIfClientExists(&_Payments.TransactOpts, clientId)
}

// PayPayment is a paid mutator transaction binding the contract method 0x23ed5d00.
//
// Solidity: function payPayment(uint256 externalPaymentId) returns()
func (_Payments *PaymentsTransactor) PayPayment(opts *bind.TransactOpts, externalPaymentId *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "payPayment", externalPaymentId)
}

// PayPayment is a paid mutator transaction binding the contract method 0x23ed5d00.
//
// Solidity: function payPayment(uint256 externalPaymentId) returns()
func (_Payments *PaymentsSession) PayPayment(externalPaymentId *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.PayPayment(&_Payments.TransactOpts, externalPaymentId)
}

// PayPayment is a paid mutator transaction binding the contract method 0x23ed5d00.
//
// Solidity: function payPayment(uint256 externalPaymentId) returns()
func (_Payments *PaymentsTransactorSession) PayPayment(externalPaymentId *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.PayPayment(&_Payments.TransactOpts, externalPaymentId)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 amount) returns()
func (_Payments *PaymentsTransactor) Sell(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "sell", amount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 amount) returns()
func (_Payments *PaymentsSession) Sell(amount *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.Sell(&_Payments.TransactOpts, amount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 amount) returns()
func (_Payments *PaymentsTransactorSession) Sell(amount *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.Sell(&_Payments.TransactOpts, amount)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa3201daa.
//
// Solidity: function setPrices(uint256 newSellPrice) returns()
func (_Payments *PaymentsTransactor) SetPrices(opts *bind.TransactOpts, newSellPrice *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "setPrices", newSellPrice)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa3201daa.
//
// Solidity: function setPrices(uint256 newSellPrice) returns()
func (_Payments *PaymentsSession) SetPrices(newSellPrice *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.SetPrices(&_Payments.TransactOpts, newSellPrice)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa3201daa.
//
// Solidity: function setPrices(uint256 newSellPrice) returns()
func (_Payments *PaymentsTransactorSession) SetPrices(newSellPrice *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.SetPrices(&_Payments.TransactOpts, newSellPrice)
}

// StartNewPayment is a paid mutator transaction binding the contract method 0xf48cef20.
//
// Solidity: function startNewPayment(uint256 price, uint256 externalPayerId, uint256 externalReceiverId) returns(uint256)
func (_Payments *PaymentsTransactor) StartNewPayment(opts *bind.TransactOpts, price *big.Int, externalPayerId *big.Int, externalReceiverId *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "startNewPayment", price, externalPayerId, externalReceiverId)
}

// StartNewPayment is a paid mutator transaction binding the contract method 0xf48cef20.
//
// Solidity: function startNewPayment(uint256 price, uint256 externalPayerId, uint256 externalReceiverId) returns(uint256)
func (_Payments *PaymentsSession) StartNewPayment(price *big.Int, externalPayerId *big.Int, externalReceiverId *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.StartNewPayment(&_Payments.TransactOpts, price, externalPayerId, externalReceiverId)
}

// StartNewPayment is a paid mutator transaction binding the contract method 0xf48cef20.
//
// Solidity: function startNewPayment(uint256 price, uint256 externalPayerId, uint256 externalReceiverId) returns(uint256)
func (_Payments *PaymentsTransactorSession) StartNewPayment(price *big.Int, externalPayerId *big.Int, externalReceiverId *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.StartNewPayment(&_Payments.TransactOpts, price, externalPayerId, externalReceiverId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Payments *PaymentsTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Payments *PaymentsSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.Transfer(&_Payments.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Payments *PaymentsTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.Transfer(&_Payments.TransactOpts, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payments *PaymentsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payments *PaymentsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Payments.Contract.TransferOwnership(&_Payments.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payments *PaymentsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Payments.Contract.TransferOwnership(&_Payments.TransactOpts, newOwner)
}

// PaymentsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Payments contract.
type PaymentsApprovalIterator struct {
	Event *PaymentsApproval // Event containing the contract specifics and raw log

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
func (it *PaymentsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsApproval)
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
		it.Event = new(PaymentsApproval)
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
func (it *PaymentsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsApproval represents a Approval event raised by the Payments contract.
type PaymentsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Payments *PaymentsFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*PaymentsApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Payments.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &PaymentsApprovalIterator{contract: _Payments.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Payments *PaymentsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PaymentsApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Payments.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsApproval)
				if err := _Payments.contract.UnpackLog(event, "Approval", log); err != nil {
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

// PaymentsNewClientEventIterator is returned from FilterNewClientEvent and is used to iterate over the raw logs and unpacked data for NewClientEvent events raised by the Payments contract.
type PaymentsNewClientEventIterator struct {
	Event *PaymentsNewClientEvent // Event containing the contract specifics and raw log

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
func (it *PaymentsNewClientEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsNewClientEvent)
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
		it.Event = new(PaymentsNewClientEvent)
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
func (it *PaymentsNewClientEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsNewClientEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsNewClientEvent represents a NewClientEvent event raised by the Payments contract.
type PaymentsNewClientEvent struct {
	ExternalClientId     *big.Int
	UserName             string
	File                 string
	ClientAccountAddress string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewClientEvent is a free log retrieval operation binding the contract event 0x8f93562878197ea1bc0b4849240129cbe4babdc145f734e6a441b02d89116781.
//
// Solidity: event NewClientEvent(uint256 externalClientId, string userName, string file, string clientAccountAddress)
func (_Payments *PaymentsFilterer) FilterNewClientEvent(opts *bind.FilterOpts) (*PaymentsNewClientEventIterator, error) {

	logs, sub, err := _Payments.contract.FilterLogs(opts, "NewClientEvent")
	if err != nil {
		return nil, err
	}
	return &PaymentsNewClientEventIterator{contract: _Payments.contract, event: "NewClientEvent", logs: logs, sub: sub}, nil
}

// WatchNewClientEvent is a free log subscription operation binding the contract event 0x8f93562878197ea1bc0b4849240129cbe4babdc145f734e6a441b02d89116781.
//
// Solidity: event NewClientEvent(uint256 externalClientId, string userName, string file, string clientAccountAddress)
func (_Payments *PaymentsFilterer) WatchNewClientEvent(opts *bind.WatchOpts, sink chan<- *PaymentsNewClientEvent) (event.Subscription, error) {

	logs, sub, err := _Payments.contract.WatchLogs(opts, "NewClientEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsNewClientEvent)
				if err := _Payments.contract.UnpackLog(event, "NewClientEvent", log); err != nil {
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

// PaymentsNewPaymentEventIterator is returned from FilterNewPaymentEvent and is used to iterate over the raw logs and unpacked data for NewPaymentEvent events raised by the Payments contract.
type PaymentsNewPaymentEventIterator struct {
	Event *PaymentsNewPaymentEvent // Event containing the contract specifics and raw log

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
func (it *PaymentsNewPaymentEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsNewPaymentEvent)
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
		it.Event = new(PaymentsNewPaymentEvent)
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
func (it *PaymentsNewPaymentEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsNewPaymentEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsNewPaymentEvent represents a NewPaymentEvent event raised by the Payments contract.
type PaymentsNewPaymentEvent struct {
	ExternalPaymentId *big.Int
	Price             *big.Int
	ExternalPayerId   *big.Int
	ReceiverId        *big.Int
	Status            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewPaymentEvent is a free log retrieval operation binding the contract event 0x2f31e803cd06ad25ed25b55a829d47dd89067f047a4257b568d63b1ad5701057.
//
// Solidity: event NewPaymentEvent(uint256 externalPaymentId, uint256 price, uint256 externalPayerId, uint256 receiverId, uint256 status)
func (_Payments *PaymentsFilterer) FilterNewPaymentEvent(opts *bind.FilterOpts) (*PaymentsNewPaymentEventIterator, error) {

	logs, sub, err := _Payments.contract.FilterLogs(opts, "NewPaymentEvent")
	if err != nil {
		return nil, err
	}
	return &PaymentsNewPaymentEventIterator{contract: _Payments.contract, event: "NewPaymentEvent", logs: logs, sub: sub}, nil
}

// WatchNewPaymentEvent is a free log subscription operation binding the contract event 0x2f31e803cd06ad25ed25b55a829d47dd89067f047a4257b568d63b1ad5701057.
//
// Solidity: event NewPaymentEvent(uint256 externalPaymentId, uint256 price, uint256 externalPayerId, uint256 receiverId, uint256 status)
func (_Payments *PaymentsFilterer) WatchNewPaymentEvent(opts *bind.WatchOpts, sink chan<- *PaymentsNewPaymentEvent) (event.Subscription, error) {

	logs, sub, err := _Payments.contract.WatchLogs(opts, "NewPaymentEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsNewPaymentEvent)
				if err := _Payments.contract.UnpackLog(event, "NewPaymentEvent", log); err != nil {
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

// PaymentsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Payments contract.
type PaymentsTransferIterator struct {
	Event *PaymentsTransfer // Event containing the contract specifics and raw log

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
func (it *PaymentsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsTransfer)
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
		it.Event = new(PaymentsTransfer)
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
func (it *PaymentsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsTransfer represents a Transfer event raised by the Payments contract.
type PaymentsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Payments *PaymentsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PaymentsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Payments.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PaymentsTransferIterator{contract: _Payments.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Payments *PaymentsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PaymentsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Payments.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsTransfer)
				if err := _Payments.contract.UnpackLog(event, "Transfer", log); err != nil {
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
