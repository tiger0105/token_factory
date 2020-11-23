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

// TokenTemplateABI is the input ABI used to generate the binding from.
const TokenTemplateABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addPermissionToMint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canMint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"logoURL\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"_logoURL\",\"type\":\"string\"},{\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// TokenTemplate is an auto generated Go binding around an Ethereum contract.
type TokenTemplate struct {
	TokenTemplateCaller     // Read-only binding to the contract
	TokenTemplateTransactor // Write-only binding to the contract
	TokenTemplateFilterer   // Log filterer for contract events
}

// TokenTemplateCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenTemplateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTemplateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTemplateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTemplateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenTemplateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTemplateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenTemplateSession struct {
	Contract     *TokenTemplate    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenTemplateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenTemplateCallerSession struct {
	Contract *TokenTemplateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenTemplateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTemplateTransactorSession struct {
	Contract     *TokenTemplateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenTemplateRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenTemplateRaw struct {
	Contract *TokenTemplate // Generic contract binding to access the raw methods on
}

// TokenTemplateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenTemplateCallerRaw struct {
	Contract *TokenTemplateCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTemplateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTemplateTransactorRaw struct {
	Contract *TokenTemplateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenTemplate creates a new instance of TokenTemplate, bound to a specific deployed contract.
func NewTokenTemplate(address common.Address, backend bind.ContractBackend) (*TokenTemplate, error) {
	contract, err := bindTokenTemplate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenTemplate{TokenTemplateCaller: TokenTemplateCaller{contract: contract}, TokenTemplateTransactor: TokenTemplateTransactor{contract: contract}, TokenTemplateFilterer: TokenTemplateFilterer{contract: contract}}, nil
}

// NewTokenTemplateCaller creates a new read-only instance of TokenTemplate, bound to a specific deployed contract.
func NewTokenTemplateCaller(address common.Address, caller bind.ContractCaller) (*TokenTemplateCaller, error) {
	contract, err := bindTokenTemplate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTemplateCaller{contract: contract}, nil
}

// NewTokenTemplateTransactor creates a new write-only instance of TokenTemplate, bound to a specific deployed contract.
func NewTokenTemplateTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTemplateTransactor, error) {
	contract, err := bindTokenTemplate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTemplateTransactor{contract: contract}, nil
}

// NewTokenTemplateFilterer creates a new log filterer instance of TokenTemplate, bound to a specific deployed contract.
func NewTokenTemplateFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenTemplateFilterer, error) {
	contract, err := bindTokenTemplate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenTemplateFilterer{contract: contract}, nil
}

// bindTokenTemplate binds a generic wrapper to an already deployed contract.
func bindTokenTemplate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenTemplateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenTemplate *TokenTemplateRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenTemplate.Contract.TokenTemplateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenTemplate *TokenTemplateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenTemplate.Contract.TokenTemplateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenTemplate *TokenTemplateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenTemplate.Contract.TokenTemplateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenTemplate *TokenTemplateCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenTemplate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenTemplate *TokenTemplateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenTemplate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenTemplate *TokenTemplateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenTemplate.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_TokenTemplate *TokenTemplateCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenTemplate.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_TokenTemplate *TokenTemplateSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _TokenTemplate.Contract.Allowance(&_TokenTemplate.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_TokenTemplate *TokenTemplateCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _TokenTemplate.Contract.Allowance(&_TokenTemplate.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_TokenTemplate *TokenTemplateCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenTemplate.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_TokenTemplate *TokenTemplateSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _TokenTemplate.Contract.BalanceOf(&_TokenTemplate.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_TokenTemplate *TokenTemplateCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _TokenTemplate.Contract.BalanceOf(&_TokenTemplate.CallOpts, _owner)
}

// CanMint is a free data retrieval call binding the contract method 0xbeb9716d.
//
// Solidity: function canMint() constant returns(bool)
func (_TokenTemplate *TokenTemplateCaller) CanMint(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenTemplate.contract.Call(opts, out, "canMint")
	return *ret0, err
}

// CanMint is a free data retrieval call binding the contract method 0xbeb9716d.
//
// Solidity: function canMint() constant returns(bool)
func (_TokenTemplate *TokenTemplateSession) CanMint() (bool, error) {
	return _TokenTemplate.Contract.CanMint(&_TokenTemplate.CallOpts)
}

// CanMint is a free data retrieval call binding the contract method 0xbeb9716d.
//
// Solidity: function canMint() constant returns(bool)
func (_TokenTemplate *TokenTemplateCallerSession) CanMint() (bool, error) {
	return _TokenTemplate.Contract.CanMint(&_TokenTemplate.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TokenTemplate *TokenTemplateCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _TokenTemplate.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TokenTemplate *TokenTemplateSession) Decimals() (uint8, error) {
	return _TokenTemplate.Contract.Decimals(&_TokenTemplate.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TokenTemplate *TokenTemplateCallerSession) Decimals() (uint8, error) {
	return _TokenTemplate.Contract.Decimals(&_TokenTemplate.CallOpts)
}

// LogoURL is a free data retrieval call binding the contract method 0xd2ce89e5.
//
// Solidity: function logoURL() constant returns(string)
func (_TokenTemplate *TokenTemplateCaller) LogoURL(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenTemplate.contract.Call(opts, out, "logoURL")
	return *ret0, err
}

// LogoURL is a free data retrieval call binding the contract method 0xd2ce89e5.
//
// Solidity: function logoURL() constant returns(string)
func (_TokenTemplate *TokenTemplateSession) LogoURL() (string, error) {
	return _TokenTemplate.Contract.LogoURL(&_TokenTemplate.CallOpts)
}

// LogoURL is a free data retrieval call binding the contract method 0xd2ce89e5.
//
// Solidity: function logoURL() constant returns(string)
func (_TokenTemplate *TokenTemplateCallerSession) LogoURL() (string, error) {
	return _TokenTemplate.Contract.LogoURL(&_TokenTemplate.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_TokenTemplate *TokenTemplateCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenTemplate.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_TokenTemplate *TokenTemplateSession) MintingFinished() (bool, error) {
	return _TokenTemplate.Contract.MintingFinished(&_TokenTemplate.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_TokenTemplate *TokenTemplateCallerSession) MintingFinished() (bool, error) {
	return _TokenTemplate.Contract.MintingFinished(&_TokenTemplate.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TokenTemplate *TokenTemplateCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenTemplate.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TokenTemplate *TokenTemplateSession) Name() (string, error) {
	return _TokenTemplate.Contract.Name(&_TokenTemplate.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TokenTemplate *TokenTemplateCallerSession) Name() (string, error) {
	return _TokenTemplate.Contract.Name(&_TokenTemplate.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TokenTemplate *TokenTemplateCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TokenTemplate.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TokenTemplate *TokenTemplateSession) Symbol() (string, error) {
	return _TokenTemplate.Contract.Symbol(&_TokenTemplate.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TokenTemplate *TokenTemplateCallerSession) Symbol() (string, error) {
	return _TokenTemplate.Contract.Symbol(&_TokenTemplate.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TokenTemplate *TokenTemplateCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenTemplate.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TokenTemplate *TokenTemplateSession) TotalSupply() (*big.Int, error) {
	return _TokenTemplate.Contract.TotalSupply(&_TokenTemplate.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TokenTemplate *TokenTemplateCallerSession) TotalSupply() (*big.Int, error) {
	return _TokenTemplate.Contract.TotalSupply(&_TokenTemplate.CallOpts)
}

// AddPermissionToMint is a paid mutator transaction binding the contract method 0x6b9ffb0d.
//
// Solidity: function addPermissionToMint(_address address) returns()
func (_TokenTemplate *TokenTemplateTransactor) AddPermissionToMint(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _TokenTemplate.contract.Transact(opts, "addPermissionToMint", _address)
}

// AddPermissionToMint is a paid mutator transaction binding the contract method 0x6b9ffb0d.
//
// Solidity: function addPermissionToMint(_address address) returns()
func (_TokenTemplate *TokenTemplateSession) AddPermissionToMint(_address common.Address) (*types.Transaction, error) {
	return _TokenTemplate.Contract.AddPermissionToMint(&_TokenTemplate.TransactOpts, _address)
}

// AddPermissionToMint is a paid mutator transaction binding the contract method 0x6b9ffb0d.
//
// Solidity: function addPermissionToMint(_address address) returns()
func (_TokenTemplate *TokenTemplateTransactorSession) AddPermissionToMint(_address common.Address) (*types.Transaction, error) {
	return _TokenTemplate.Contract.AddPermissionToMint(&_TokenTemplate.TransactOpts, _address)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_TokenTemplate *TokenTemplateTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_TokenTemplate *TokenTemplateSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.Contract.Approve(&_TokenTemplate.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_TokenTemplate *TokenTemplateTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.Contract.Approve(&_TokenTemplate.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_TokenTemplate *TokenTemplateTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_TokenTemplate *TokenTemplateSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.Contract.DecreaseApproval(&_TokenTemplate.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_TokenTemplate *TokenTemplateTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.Contract.DecreaseApproval(&_TokenTemplate.TransactOpts, _spender, _subtractedValue)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_TokenTemplate *TokenTemplateTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenTemplate.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_TokenTemplate *TokenTemplateSession) FinishMinting() (*types.Transaction, error) {
	return _TokenTemplate.Contract.FinishMinting(&_TokenTemplate.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_TokenTemplate *TokenTemplateTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _TokenTemplate.Contract.FinishMinting(&_TokenTemplate.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_TokenTemplate *TokenTemplateTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_TokenTemplate *TokenTemplateSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.Contract.IncreaseApproval(&_TokenTemplate.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_TokenTemplate *TokenTemplateTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.Contract.IncreaseApproval(&_TokenTemplate.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_TokenTemplate *TokenTemplateTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_TokenTemplate *TokenTemplateSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.Contract.Mint(&_TokenTemplate.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_TokenTemplate *TokenTemplateTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.Contract.Mint(&_TokenTemplate.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_TokenTemplate *TokenTemplateTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_TokenTemplate *TokenTemplateSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.Contract.Transfer(&_TokenTemplate.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_TokenTemplate *TokenTemplateTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.Contract.Transfer(&_TokenTemplate.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_TokenTemplate *TokenTemplateTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_TokenTemplate *TokenTemplateSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.Contract.TransferFrom(&_TokenTemplate.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_TokenTemplate *TokenTemplateTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenTemplate.Contract.TransferFrom(&_TokenTemplate.TransactOpts, _from, _to, _value)
}

// TokenTemplateApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TokenTemplate contract.
type TokenTemplateApprovalIterator struct {
	Event *TokenTemplateApproval // Event containing the contract specifics and raw log

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
func (it *TokenTemplateApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTemplateApproval)
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
		it.Event = new(TokenTemplateApproval)
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
func (it *TokenTemplateApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTemplateApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTemplateApproval represents a Approval event raised by the TokenTemplate contract.
type TokenTemplateApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_TokenTemplate *TokenTemplateFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenTemplateApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenTemplate.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenTemplateApprovalIterator{contract: _TokenTemplate.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_TokenTemplate *TokenTemplateFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenTemplateApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenTemplate.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTemplateApproval)
				if err := _TokenTemplate.contract.UnpackLog(event, "Approval", log); err != nil {
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

// TokenTemplateMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the TokenTemplate contract.
type TokenTemplateMintIterator struct {
	Event *TokenTemplateMint // Event containing the contract specifics and raw log

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
func (it *TokenTemplateMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTemplateMint)
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
		it.Event = new(TokenTemplateMint)
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
func (it *TokenTemplateMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTemplateMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTemplateMint represents a Mint event raised by the TokenTemplate contract.
type TokenTemplateMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(to indexed address, amount uint256)
func (_TokenTemplate *TokenTemplateFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*TokenTemplateMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenTemplate.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTemplateMintIterator{contract: _TokenTemplate.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(to indexed address, amount uint256)
func (_TokenTemplate *TokenTemplateFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *TokenTemplateMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenTemplate.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTemplateMint)
				if err := _TokenTemplate.contract.UnpackLog(event, "Mint", log); err != nil {
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

// TokenTemplateMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the TokenTemplate contract.
type TokenTemplateMintFinishedIterator struct {
	Event *TokenTemplateMintFinished // Event containing the contract specifics and raw log

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
func (it *TokenTemplateMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTemplateMintFinished)
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
		it.Event = new(TokenTemplateMintFinished)
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
func (it *TokenTemplateMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTemplateMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTemplateMintFinished represents a MintFinished event raised by the TokenTemplate contract.
type TokenTemplateMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_TokenTemplate *TokenTemplateFilterer) FilterMintFinished(opts *bind.FilterOpts) (*TokenTemplateMintFinishedIterator, error) {

	logs, sub, err := _TokenTemplate.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &TokenTemplateMintFinishedIterator{contract: _TokenTemplate.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_TokenTemplate *TokenTemplateFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *TokenTemplateMintFinished) (event.Subscription, error) {

	logs, sub, err := _TokenTemplate.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTemplateMintFinished)
				if err := _TokenTemplate.contract.UnpackLog(event, "MintFinished", log); err != nil {
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

// TokenTemplateTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenTemplate contract.
type TokenTemplateTransferIterator struct {
	Event *TokenTemplateTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTemplateTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTemplateTransfer)
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
		it.Event = new(TokenTemplateTransfer)
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
func (it *TokenTemplateTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTemplateTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTemplateTransfer represents a Transfer event raised by the TokenTemplate contract.
type TokenTemplateTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_TokenTemplate *TokenTemplateFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTemplateTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenTemplate.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTemplateTransferIterator{contract: _TokenTemplate.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_TokenTemplate *TokenTemplateFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTemplateTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenTemplate.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTemplateTransfer)
				if err := _TokenTemplate.contract.UnpackLog(event, "Transfer", log); err != nil {
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
