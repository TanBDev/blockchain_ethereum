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

// ImoDollarABI is the input ABI used to generate the binding from.
const ImoDollarABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sellPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newSellPrice\",\"type\":\"uint256\"}],\"name\":\"setPrices\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"name\":\"tokenInitialBalance\",\"type\":\"uint256\"},{\"name\":\"tokenDevAddress\",\"type\":\"address\"},{\"name\":\"setSellPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ImoDollarBin is the compiled bytecode used for deploying new contracts.
const ImoDollarBin = `608060405234801561001057600080fd5b5060405162000ee838038062000ee8833981810160405260a081101561003557600080fd5b81019080805164010000000081111561004d57600080fd5b8281019050602081018481111561006357600080fd5b815185600182028301116401000000008211171561008057600080fd5b5050929190602001805164010000000081111561009c57600080fd5b828101905060208101848111156100b257600080fd5b81518560018202830111640100000000821117156100cf57600080fd5b5050929190602001805190602001909291908051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550846001908051906020019061014e929190610234565b508360029080519060200190610165929190610234565b5081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008390508060076000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080600460008282540192505081905550816006819055505050505050506102d9565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061027557805160ff19168380011785556102a3565b828001600101855582156102a3579182015b828111156102a2578251825591602001919060010190610287565b5b5090506102b091906102b4565b5090565b6102d691905b808211156102d25760008160009055506001016102ba565b5090565b90565b610bff80620002e96000396000f3fe6080604052600436106100c25760003560e01c80638da5cb5b1161007f578063a6f2ae3a11610059578063a6f2ae3a146103d8578063a9059cbb146103e2578063e4849b3214610455578063f2fde38b14610490576100c2565b80638da5cb5b146102b657806395d89b411461030d578063a3201daa1461039d576100c2565b806306fdde03146100c7578063095ea7b31461015757806318160ddd146101ca578063313ce567146101f55780634b7503341461022657806370a0823114610251575b600080fd5b3480156100d357600080fd5b506100dc6104e1565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561011c578082015181840152602081019050610101565b50505050905090810190601f1680156101495780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561016357600080fd5b506101b06004803603604081101561017a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061057f565b604051808215151515815260200191505060405180910390f35b3480156101d657600080fd5b506101df610671565b6040518082815260200191505060405180910390f35b34801561020157600080fd5b5061020a610677565b604051808260ff1660ff16815260200191505060405180910390f35b34801561023257600080fd5b5061023b61068a565b6040518082815260200191505060405180910390f35b34801561025d57600080fd5b506102a06004803603602081101561027457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610690565b6040518082815260200191505060405180910390f35b3480156102c257600080fd5b506102cb6106d9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561031957600080fd5b506103226106fe565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610362578082015181840152602081019050610347565b50505050905090810190601f16801561038f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103a957600080fd5b506103d6600480360360208110156103c057600080fd5b810190808035906020019092919050505061079c565b005b6103e06107ff565b005b3480156103ee57600080fd5b5061043b6004803603604081101561040557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610816565b604051808215151515815260200191505060405180910390f35b34801561046157600080fd5b5061048e6004803603602081101561047857600080fd5b810190808035906020019092919050505061082d565b005b34801561049c57600080fd5b506104df600480360360208110156104b357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108bb565b005b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105775780601f1061054c57610100808354040283529160200191610577565b820191906000526020600020905b81548152906001019060200180831161055a57829003601f168201915b505050505081565b600081600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b60045481565b600360009054906101000a900460ff1681565b60065481565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107945780601f1061076957610100808354040283529160200191610794565b820191906000526020600020905b81548152906001019060200180831161077757829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107f557600080fd5b8060068190555050565b6000600654340290506108123082610816565b5050565b6000610823338484610957565b6001905092915050565b6000309050600654828161083d57fe5b048173ffffffffffffffffffffffffffffffffffffffff1631101561086157600080fd5b61086c333084610957565b3373ffffffffffffffffffffffffffffffffffffffff166108fc60065484029081150290604051600060405180830381858888f193505050501580156108b6573d6000803e3d6000fd5b505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461091457600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b80600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156109a357600080fd5b600081116109b057600080fd5b600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540111610a3c57600080fd5b80600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555080600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555080600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a350505056fea265627a7a723058207fa94cee5d13b85f1c9b4a1b8d5d0c2bc0352b65bba53f6b8d5621e166cbea6b64736f6c63430005090032`

// DeployImoDollar deploys a new Ethereum contract, binding an instance of ImoDollar to it.
func DeployImoDollar(auth *bind.TransactOpts, backend bind.ContractBackend, tokenName string, tokenSymbol string, tokenInitialBalance *big.Int, tokenDevAddress common.Address, setSellPrice *big.Int) (common.Address, *types.Transaction, *ImoDollar, error) {
	parsed, err := abi.JSON(strings.NewReader(ImoDollarABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ImoDollarBin), backend, tokenName, tokenSymbol, tokenInitialBalance, tokenDevAddress, setSellPrice)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ImoDollar{ImoDollarCaller: ImoDollarCaller{contract: contract}, ImoDollarTransactor: ImoDollarTransactor{contract: contract}, ImoDollarFilterer: ImoDollarFilterer{contract: contract}}, nil
}

// ImoDollar is an auto generated Go binding around an Ethereum contract.
type ImoDollar struct {
	ImoDollarCaller     // Read-only binding to the contract
	ImoDollarTransactor // Write-only binding to the contract
	ImoDollarFilterer   // Log filterer for contract events
}

// ImoDollarCaller is an auto generated read-only Go binding around an Ethereum contract.
type ImoDollarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImoDollarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ImoDollarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImoDollarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ImoDollarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImoDollarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ImoDollarSession struct {
	Contract     *ImoDollar        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ImoDollarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ImoDollarCallerSession struct {
	Contract *ImoDollarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ImoDollarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ImoDollarTransactorSession struct {
	Contract     *ImoDollarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ImoDollarRaw is an auto generated low-level Go binding around an Ethereum contract.
type ImoDollarRaw struct {
	Contract *ImoDollar // Generic contract binding to access the raw methods on
}

// ImoDollarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ImoDollarCallerRaw struct {
	Contract *ImoDollarCaller // Generic read-only contract binding to access the raw methods on
}

// ImoDollarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ImoDollarTransactorRaw struct {
	Contract *ImoDollarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewImoDollar creates a new instance of ImoDollar, bound to a specific deployed contract.
func NewImoDollar(address common.Address, backend bind.ContractBackend) (*ImoDollar, error) {
	contract, err := bindImoDollar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ImoDollar{ImoDollarCaller: ImoDollarCaller{contract: contract}, ImoDollarTransactor: ImoDollarTransactor{contract: contract}, ImoDollarFilterer: ImoDollarFilterer{contract: contract}}, nil
}

// NewImoDollarCaller creates a new read-only instance of ImoDollar, bound to a specific deployed contract.
func NewImoDollarCaller(address common.Address, caller bind.ContractCaller) (*ImoDollarCaller, error) {
	contract, err := bindImoDollar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ImoDollarCaller{contract: contract}, nil
}

// NewImoDollarTransactor creates a new write-only instance of ImoDollar, bound to a specific deployed contract.
func NewImoDollarTransactor(address common.Address, transactor bind.ContractTransactor) (*ImoDollarTransactor, error) {
	contract, err := bindImoDollar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ImoDollarTransactor{contract: contract}, nil
}

// NewImoDollarFilterer creates a new log filterer instance of ImoDollar, bound to a specific deployed contract.
func NewImoDollarFilterer(address common.Address, filterer bind.ContractFilterer) (*ImoDollarFilterer, error) {
	contract, err := bindImoDollar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ImoDollarFilterer{contract: contract}, nil
}

// bindImoDollar binds a generic wrapper to an already deployed contract.
func bindImoDollar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ImoDollarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ImoDollar *ImoDollarRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ImoDollar.Contract.ImoDollarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ImoDollar *ImoDollarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImoDollar.Contract.ImoDollarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ImoDollar *ImoDollarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ImoDollar.Contract.ImoDollarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ImoDollar *ImoDollarCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ImoDollar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ImoDollar *ImoDollarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImoDollar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ImoDollar *ImoDollarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ImoDollar.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_ImoDollar *ImoDollarCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ImoDollar.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_ImoDollar *ImoDollarSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ImoDollar.Contract.BalanceOf(&_ImoDollar.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_ImoDollar *ImoDollarCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ImoDollar.Contract.BalanceOf(&_ImoDollar.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ImoDollar *ImoDollarCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ImoDollar.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ImoDollar *ImoDollarSession) Decimals() (uint8, error) {
	return _ImoDollar.Contract.Decimals(&_ImoDollar.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ImoDollar *ImoDollarCallerSession) Decimals() (uint8, error) {
	return _ImoDollar.Contract.Decimals(&_ImoDollar.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ImoDollar *ImoDollarCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ImoDollar.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ImoDollar *ImoDollarSession) Name() (string, error) {
	return _ImoDollar.Contract.Name(&_ImoDollar.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ImoDollar *ImoDollarCallerSession) Name() (string, error) {
	return _ImoDollar.Contract.Name(&_ImoDollar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ImoDollar *ImoDollarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ImoDollar.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ImoDollar *ImoDollarSession) Owner() (common.Address, error) {
	return _ImoDollar.Contract.Owner(&_ImoDollar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ImoDollar *ImoDollarCallerSession) Owner() (common.Address, error) {
	return _ImoDollar.Contract.Owner(&_ImoDollar.CallOpts)
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_ImoDollar *ImoDollarCaller) SellPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ImoDollar.contract.Call(opts, out, "sellPrice")
	return *ret0, err
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_ImoDollar *ImoDollarSession) SellPrice() (*big.Int, error) {
	return _ImoDollar.Contract.SellPrice(&_ImoDollar.CallOpts)
}

// SellPrice is a free data retrieval call binding the contract method 0x4b750334.
//
// Solidity: function sellPrice() constant returns(uint256)
func (_ImoDollar *ImoDollarCallerSession) SellPrice() (*big.Int, error) {
	return _ImoDollar.Contract.SellPrice(&_ImoDollar.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ImoDollar *ImoDollarCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ImoDollar.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ImoDollar *ImoDollarSession) Symbol() (string, error) {
	return _ImoDollar.Contract.Symbol(&_ImoDollar.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ImoDollar *ImoDollarCallerSession) Symbol() (string, error) {
	return _ImoDollar.Contract.Symbol(&_ImoDollar.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ImoDollar *ImoDollarCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ImoDollar.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ImoDollar *ImoDollarSession) TotalSupply() (*big.Int, error) {
	return _ImoDollar.Contract.TotalSupply(&_ImoDollar.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ImoDollar *ImoDollarCallerSession) TotalSupply() (*big.Int, error) {
	return _ImoDollar.Contract.TotalSupply(&_ImoDollar.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool success)
func (_ImoDollar *ImoDollarTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ImoDollar.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool success)
func (_ImoDollar *ImoDollarSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ImoDollar.Contract.Approve(&_ImoDollar.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool success)
func (_ImoDollar *ImoDollarTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ImoDollar.Contract.Approve(&_ImoDollar.TransactOpts, _spender, _amount)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_ImoDollar *ImoDollarTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImoDollar.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_ImoDollar *ImoDollarSession) Buy() (*types.Transaction, error) {
	return _ImoDollar.Contract.Buy(&_ImoDollar.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() returns()
func (_ImoDollar *ImoDollarTransactorSession) Buy() (*types.Transaction, error) {
	return _ImoDollar.Contract.Buy(&_ImoDollar.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 amount) returns()
func (_ImoDollar *ImoDollarTransactor) Sell(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ImoDollar.contract.Transact(opts, "sell", amount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 amount) returns()
func (_ImoDollar *ImoDollarSession) Sell(amount *big.Int) (*types.Transaction, error) {
	return _ImoDollar.Contract.Sell(&_ImoDollar.TransactOpts, amount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 amount) returns()
func (_ImoDollar *ImoDollarTransactorSession) Sell(amount *big.Int) (*types.Transaction, error) {
	return _ImoDollar.Contract.Sell(&_ImoDollar.TransactOpts, amount)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa3201daa.
//
// Solidity: function setPrices(uint256 newSellPrice) returns()
func (_ImoDollar *ImoDollarTransactor) SetPrices(opts *bind.TransactOpts, newSellPrice *big.Int) (*types.Transaction, error) {
	return _ImoDollar.contract.Transact(opts, "setPrices", newSellPrice)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa3201daa.
//
// Solidity: function setPrices(uint256 newSellPrice) returns()
func (_ImoDollar *ImoDollarSession) SetPrices(newSellPrice *big.Int) (*types.Transaction, error) {
	return _ImoDollar.Contract.SetPrices(&_ImoDollar.TransactOpts, newSellPrice)
}

// SetPrices is a paid mutator transaction binding the contract method 0xa3201daa.
//
// Solidity: function setPrices(uint256 newSellPrice) returns()
func (_ImoDollar *ImoDollarTransactorSession) SetPrices(newSellPrice *big.Int) (*types.Transaction, error) {
	return _ImoDollar.Contract.SetPrices(&_ImoDollar.TransactOpts, newSellPrice)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_ImoDollar *ImoDollarTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ImoDollar.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_ImoDollar *ImoDollarSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ImoDollar.Contract.Transfer(&_ImoDollar.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_ImoDollar *ImoDollarTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ImoDollar.Contract.Transfer(&_ImoDollar.TransactOpts, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ImoDollar *ImoDollarTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ImoDollar.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ImoDollar *ImoDollarSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ImoDollar.Contract.TransferOwnership(&_ImoDollar.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ImoDollar *ImoDollarTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ImoDollar.Contract.TransferOwnership(&_ImoDollar.TransactOpts, newOwner)
}

// ImoDollarApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ImoDollar contract.
type ImoDollarApprovalIterator struct {
	Event *ImoDollarApproval // Event containing the contract specifics and raw log

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
func (it *ImoDollarApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImoDollarApproval)
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
		it.Event = new(ImoDollarApproval)
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
func (it *ImoDollarApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImoDollarApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImoDollarApproval represents a Approval event raised by the ImoDollar contract.
type ImoDollarApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ImoDollar *ImoDollarFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ImoDollarApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ImoDollar.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ImoDollarApprovalIterator{contract: _ImoDollar.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ImoDollar *ImoDollarFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ImoDollarApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ImoDollar.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImoDollarApproval)
				if err := _ImoDollar.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ImoDollarTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ImoDollar contract.
type ImoDollarTransferIterator struct {
	Event *ImoDollarTransfer // Event containing the contract specifics and raw log

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
func (it *ImoDollarTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImoDollarTransfer)
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
		it.Event = new(ImoDollarTransfer)
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
func (it *ImoDollarTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImoDollarTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImoDollarTransfer represents a Transfer event raised by the ImoDollar contract.
type ImoDollarTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ImoDollar *ImoDollarFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ImoDollarTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ImoDollar.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ImoDollarTransferIterator{contract: _ImoDollar.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ImoDollar *ImoDollarFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ImoDollarTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ImoDollar.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImoDollarTransfer)
				if err := _ImoDollar.contract.UnpackLog(event, "Transfer", log); err != nil {
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
