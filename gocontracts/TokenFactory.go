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

// TokenFactoryABI is the input ABI used to generate the binding from.
const TokenFactoryABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_logoURL\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_hardCap\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"makeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"getToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"_logoURL\",\"type\":\"string\"},{\"name\":\"_hardCap\",\"type\":\"uint256\"}],\"name\":\"createToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenFactory is an auto generated Go binding around an Ethereum contract.
type TokenFactory struct {
	TokenFactoryCaller     // Read-only binding to the contract
	TokenFactoryTransactor // Write-only binding to the contract
	TokenFactoryFilterer   // Log filterer for contract events
}

// TokenFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenFactorySession struct {
	Contract     *TokenFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenFactoryCallerSession struct {
	Contract *TokenFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenFactoryTransactorSession struct {
	Contract     *TokenFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenFactoryRaw struct {
	Contract *TokenFactory // Generic contract binding to access the raw methods on
}

// TokenFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenFactoryCallerRaw struct {
	Contract *TokenFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenFactoryTransactorRaw struct {
	Contract *TokenFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenFactory creates a new instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactory(address common.Address, backend bind.ContractBackend) (*TokenFactory, error) {
	contract, err := bindTokenFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenFactory{TokenFactoryCaller: TokenFactoryCaller{contract: contract}, TokenFactoryTransactor: TokenFactoryTransactor{contract: contract}, TokenFactoryFilterer: TokenFactoryFilterer{contract: contract}}, nil
}

// NewTokenFactoryCaller creates a new read-only instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryCaller(address common.Address, caller bind.ContractCaller) (*TokenFactoryCaller, error) {
	contract, err := bindTokenFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryCaller{contract: contract}, nil
}

// NewTokenFactoryTransactor creates a new write-only instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenFactoryTransactor, error) {
	contract, err := bindTokenFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryTransactor{contract: contract}, nil
}

// NewTokenFactoryFilterer creates a new log filterer instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFactoryFilterer, error) {
	contract, err := bindTokenFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryFilterer{contract: contract}, nil
}

// bindTokenFactory binds a generic wrapper to an already deployed contract.
func bindTokenFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFactory *TokenFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenFactory.Contract.TokenFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFactory *TokenFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFactory.Contract.TokenFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFactory *TokenFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFactory.Contract.TokenFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFactory *TokenFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFactory *TokenFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFactory *TokenFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFactory.Contract.contract.Transact(opts, method, params...)
}

// GetToken is a free data retrieval call binding the contract method 0xc1733f68.
//
// Solidity: function getToken(_symbol string) constant returns(address, string, string, uint8, string, address, uint256)
func (_TokenFactory *TokenFactoryCaller) GetToken(opts *bind.CallOpts, _symbol string) (common.Address, string, string, uint8, string, common.Address, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new(uint8)
		ret4 = new(string)
		ret5 = new(common.Address)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _TokenFactory.contract.Call(opts, out, "getToken", _symbol)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetToken is a free data retrieval call binding the contract method 0xc1733f68.
//
// Solidity: function getToken(_symbol string) constant returns(address, string, string, uint8, string, address, uint256)
func (_TokenFactory *TokenFactorySession) GetToken(_symbol string) (common.Address, string, string, uint8, string, common.Address, *big.Int, error) {
	return _TokenFactory.Contract.GetToken(&_TokenFactory.CallOpts, _symbol)
}

// GetToken is a free data retrieval call binding the contract method 0xc1733f68.
//
// Solidity: function getToken(_symbol string) constant returns(address, string, string, uint8, string, address, uint256)
func (_TokenFactory *TokenFactoryCallerSession) GetToken(_symbol string) (common.Address, string, string, uint8, string, common.Address, *big.Int, error) {
	return _TokenFactory.Contract.GetToken(&_TokenFactory.CallOpts, _symbol)
}

// CreateToken is a paid mutator transaction binding the contract method 0x36985d0d.
//
// Solidity: function createToken(_name string, _symbol string, _decimals uint8, _logoURL string, _hardCap uint256) returns()
func (_TokenFactory *TokenFactoryTransactor) CreateToken(opts *bind.TransactOpts, _name string, _symbol string, _decimals uint8, _logoURL string, _hardCap *big.Int) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "createToken", _name, _symbol, _decimals, _logoURL, _hardCap)
}

// CreateToken is a paid mutator transaction binding the contract method 0x36985d0d.
//
// Solidity: function createToken(_name string, _symbol string, _decimals uint8, _logoURL string, _hardCap uint256) returns()
func (_TokenFactory *TokenFactorySession) CreateToken(_name string, _symbol string, _decimals uint8, _logoURL string, _hardCap *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateToken(&_TokenFactory.TransactOpts, _name, _symbol, _decimals, _logoURL, _hardCap)
}

// CreateToken is a paid mutator transaction binding the contract method 0x36985d0d.
//
// Solidity: function createToken(_name string, _symbol string, _decimals uint8, _logoURL string, _hardCap uint256) returns()
func (_TokenFactory *TokenFactoryTransactorSession) CreateToken(_name string, _symbol string, _decimals uint8, _logoURL string, _hardCap *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateToken(&_TokenFactory.TransactOpts, _name, _symbol, _decimals, _logoURL, _hardCap)
}

// MakeAdmin is a paid mutator transaction binding the contract method 0x472905ca.
//
// Solidity: function makeAdmin(_address address) returns()
func (_TokenFactory *TokenFactoryTransactor) MakeAdmin(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "makeAdmin", _address)
}

// MakeAdmin is a paid mutator transaction binding the contract method 0x472905ca.
//
// Solidity: function makeAdmin(_address address) returns()
func (_TokenFactory *TokenFactorySession) MakeAdmin(_address common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.MakeAdmin(&_TokenFactory.TransactOpts, _address)
}

// MakeAdmin is a paid mutator transaction binding the contract method 0x472905ca.
//
// Solidity: function makeAdmin(_address address) returns()
func (_TokenFactory *TokenFactoryTransactorSession) MakeAdmin(_address common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.MakeAdmin(&_TokenFactory.TransactOpts, _address)
}

// MintToken is a paid mutator transaction binding the contract method 0xcf2a9169.
//
// Solidity: function mintToken(_symbol string, _to address, _amount uint256) returns()
func (_TokenFactory *TokenFactoryTransactor) MintToken(opts *bind.TransactOpts, _symbol string, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "mintToken", _symbol, _to, _amount)
}

// MintToken is a paid mutator transaction binding the contract method 0xcf2a9169.
//
// Solidity: function mintToken(_symbol string, _to address, _amount uint256) returns()
func (_TokenFactory *TokenFactorySession) MintToken(_symbol string, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.MintToken(&_TokenFactory.TransactOpts, _symbol, _to, _amount)
}

// MintToken is a paid mutator transaction binding the contract method 0xcf2a9169.
//
// Solidity: function mintToken(_symbol string, _to address, _amount uint256) returns()
func (_TokenFactory *TokenFactoryTransactorSession) MintToken(_symbol string, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenFactory.Contract.MintToken(&_TokenFactory.TransactOpts, _symbol, _to, _amount)
}

// TokenFactoryAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the TokenFactory contract.
type TokenFactoryAdminAddedIterator struct {
	Event *TokenFactoryAdminAdded // Event containing the contract specifics and raw log

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
func (it *TokenFactoryAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryAdminAdded)
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
		it.Event = new(TokenFactoryAdminAdded)
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
func (it *TokenFactoryAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryAdminAdded represents a AdminAdded event raised by the TokenFactory contract.
type TokenFactoryAdminAdded struct {
	From common.Address
	Who  common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0xbf3f493c772c8c283fd124432c2d0f539ab343faa04258fe88e52912d36b102b.
//
// Solidity: event AdminAdded(_from indexed address, _who indexed address)
func (_TokenFactory *TokenFactoryFilterer) FilterAdminAdded(opts *bind.FilterOpts, _from []common.Address, _who []common.Address) (*TokenFactoryAdminAddedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "AdminAdded", _fromRule, _whoRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryAdminAddedIterator{contract: _TokenFactory.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0xbf3f493c772c8c283fd124432c2d0f539ab343faa04258fe88e52912d36b102b.
//
// Solidity: event AdminAdded(_from indexed address, _who indexed address)
func (_TokenFactory *TokenFactoryFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *TokenFactoryAdminAdded, _from []common.Address, _who []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "AdminAdded", _fromRule, _whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryAdminAdded)
				if err := _TokenFactory.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// TokenFactoryTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the TokenFactory contract.
type TokenFactoryTokenAddedIterator struct {
	Event *TokenFactoryTokenAdded // Event containing the contract specifics and raw log

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
func (it *TokenFactoryTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryTokenAdded)
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
		it.Event = new(TokenFactoryTokenAdded)
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
func (it *TokenFactoryTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryTokenAdded represents a TokenAdded event raised by the TokenFactory contract.
type TokenFactoryTokenAdded struct {
	From            common.Address
	Timestamp       *big.Int
	ContractAddress common.Address
	Name            string
	Symbol          string
	Decimals        uint8
	LogoURL         string
	HardCap         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x9aeaaaeb259d07bb4dd3a7108b98292916901127a8ea11957ac316cffd02c3ae.
//
// Solidity: event TokenAdded(_from indexed address, _timestamp uint256, _contractAddress address, _name string, _symbol string, _decimals uint8, _logoURL string, _hardCap uint256)
func (_TokenFactory *TokenFactoryFilterer) FilterTokenAdded(opts *bind.FilterOpts, _from []common.Address) (*TokenFactoryTokenAddedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "TokenAdded", _fromRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryTokenAddedIterator{contract: _TokenFactory.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x9aeaaaeb259d07bb4dd3a7108b98292916901127a8ea11957ac316cffd02c3ae.
//
// Solidity: event TokenAdded(_from indexed address, _timestamp uint256, _contractAddress address, _name string, _symbol string, _decimals uint8, _logoURL string, _hardCap uint256)
func (_TokenFactory *TokenFactoryFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *TokenFactoryTokenAdded, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "TokenAdded", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryTokenAdded)
				if err := _TokenFactory.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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
