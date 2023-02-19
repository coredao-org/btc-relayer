// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relayerhub

import (
	"fmt"
	"math/big"
	"strconv"
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
	_ = abi.MaxUint256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RelayerhubABI is the input ABI used to generate the binding from.
const RelayerhubABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"paramChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"relayerRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"relayerUnRegister\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CANDIDATE_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_DUES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_REQUIRED_DEPOSIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PLEDGE_AGENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Relayerhub is an auto generated Go binding around an Ethereum contract.
type Relayerhub struct {
	RelayerhubCaller     // Read-only binding to the contract
	RelayerhubTransactor // Write-only binding to the contract
	RelayerhubFilterer   // Log filterer for contract events
}

// RelayerhubCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayerhubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerhubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayerhubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerhubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayerhubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerhubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayerhubSession struct {
	Contract     *Relayerhub       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayerhubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayerhubCallerSession struct {
	Contract *RelayerhubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RelayerhubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayerhubTransactorSession struct {
	Contract     *RelayerhubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RelayerhubRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayerhubRaw struct {
	Contract *Relayerhub // Generic contract binding to access the raw methods on
}

// RelayerhubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayerhubCallerRaw struct {
	Contract *RelayerhubCaller // Generic read-only contract binding to access the raw methods on
}

// RelayerhubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayerhubTransactorRaw struct {
	Contract *RelayerhubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayerhub creates a new instance of Relayerhub, bound to a specific deployed contract.
func NewRelayerhub(address common.Address, backend bind.ContractBackend) (*Relayerhub, error) {
	contract, err := bindRelayerhub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relayerhub{RelayerhubCaller: RelayerhubCaller{contract: contract}, RelayerhubTransactor: RelayerhubTransactor{contract: contract}, RelayerhubFilterer: RelayerhubFilterer{contract: contract}}, nil
}

// NewRelayerhubCaller creates a new read-only instance of Relayerhub, bound to a specific deployed contract.
func NewRelayerhubCaller(address common.Address, caller bind.ContractCaller) (*RelayerhubCaller, error) {
	contract, err := bindRelayerhub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerhubCaller{contract: contract}, nil
}

// NewRelayerhubTransactor creates a new write-only instance of Relayerhub, bound to a specific deployed contract.
func NewRelayerhubTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayerhubTransactor, error) {
	contract, err := bindRelayerhub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerhubTransactor{contract: contract}, nil
}

// NewRelayerhubFilterer creates a new log filterer instance of Relayerhub, bound to a specific deployed contract.
func NewRelayerhubFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayerhubFilterer, error) {
	contract, err := bindRelayerhub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayerhubFilterer{contract: contract}, nil
}

// bindRelayerhub binds a generic wrapper to an already deployed contract.
func bindRelayerhub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayerhubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
//func (_Relayerhub *RelayerhubRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
//	return _Relayerhub.Contract.RelayerhubCaller.contract.Call(opts, result, method, params...)
//}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayerhub *RelayerhubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.Contract.RelayerhubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayerhub *RelayerhubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayerhub.Contract.RelayerhubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
//func (_Relayerhub *RelayerhubCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
//	return _Relayerhub.Contract.contract.Call(opts, result, method, params...)
//}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayerhub *RelayerhubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayerhub *RelayerhubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayerhub.Contract.contract.Transact(opts, method, params...)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) constant returns(bool)
func (_Relayerhub *RelayerhubCaller) IsRelayer(opts *bind.CallOpts, sender common.Address) (bool, error) {
	retval := make([]interface{}, 0, 1)
	err := _Relayerhub.contract.Call(opts, &retval, "isRelayer", sender)
	if err == nil {
		ret0, err := strconv.ParseBool(fmt.Sprint(retval[0]))
		return ret0, err
	}
	return false, err
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) constant returns(bool)
func (_Relayerhub *RelayerhubSession) IsRelayer(sender common.Address) (bool, error) {
	return _Relayerhub.Contract.IsRelayer(&_Relayerhub.CallOpts, sender)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) constant returns(bool)
func (_Relayerhub *RelayerhubCallerSession) IsRelayer(sender common.Address) (bool, error) {
	return _Relayerhub.Contract.IsRelayer(&_Relayerhub.CallOpts, sender)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Relayerhub *RelayerhubTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Relayerhub *RelayerhubSession) Init() (*types.Transaction, error) {
	return _Relayerhub.Contract.Init(&_Relayerhub.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Relayerhub *RelayerhubTransactorSession) Init() (*types.Transaction, error) {
	return _Relayerhub.Contract.Init(&_Relayerhub.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Relayerhub *RelayerhubTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Relayerhub *RelayerhubSession) Register() (*types.Transaction, error) {
	return _Relayerhub.Contract.Register(&_Relayerhub.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Relayerhub *RelayerhubTransactorSession) Register() (*types.Transaction, error) {
	return _Relayerhub.Contract.Register(&_Relayerhub.TransactOpts)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_Relayerhub *RelayerhubTransactor) Unregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "unregister")
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_Relayerhub *RelayerhubSession) Unregister() (*types.Transaction, error) {
	return _Relayerhub.Contract.Unregister(&_Relayerhub.TransactOpts)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_Relayerhub *RelayerhubTransactorSession) Unregister() (*types.Transaction, error) {
	return _Relayerhub.Contract.Unregister(&_Relayerhub.TransactOpts)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Relayerhub *RelayerhubTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Relayerhub *RelayerhubSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Relayerhub.Contract.UpdateParam(&_Relayerhub.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Relayerhub *RelayerhubTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Relayerhub.Contract.UpdateParam(&_Relayerhub.TransactOpts, key, value)
}

// RelayerhubParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Relayerhub contract.
type RelayerhubParamChangeIterator struct {
	Event *RelayerhubParamChange // Event containing the contract specifics and raw log

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
func (it *RelayerhubParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerhubParamChange)
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
		it.Event = new(RelayerhubParamChange)
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
func (it *RelayerhubParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerhubParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerhubParamChange represents a ParamChange event raised by the Relayerhub contract.
type RelayerhubParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Relayerhub *RelayerhubFilterer) FilterParamChange(opts *bind.FilterOpts) (*RelayerhubParamChangeIterator, error) {

	logs, sub, err := _Relayerhub.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &RelayerhubParamChangeIterator{contract: _Relayerhub.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Relayerhub *RelayerhubFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *RelayerhubParamChange) (event.Subscription, error) {

	logs, sub, err := _Relayerhub.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerhubParamChange)
				if err := _Relayerhub.contract.UnpackLog(event, "paramChange", log); err != nil {
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

// ParseParamChange is a log parse operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Relayerhub *RelayerhubFilterer) ParseParamChange(log types.Log) (*RelayerhubParamChange, error) {
	event := new(RelayerhubParamChange)
	if err := _Relayerhub.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerhubRelayerRegisterIterator is returned from FilterRelayerRegister and is used to iterate over the raw logs and unpacked data for RelayerRegister events raised by the Relayerhub contract.
type RelayerhubRelayerRegisterIterator struct {
	Event *RelayerhubRelayerRegister // Event containing the contract specifics and raw log

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
func (it *RelayerhubRelayerRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerhubRelayerRegister)
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
		it.Event = new(RelayerhubRelayerRegister)
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
func (it *RelayerhubRelayerRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerhubRelayerRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerhubRelayerRegister represents a RelayerRegister event raised by the Relayerhub contract.
type RelayerhubRelayerRegister struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRegister is a free log retrieval operation binding the contract event 0xdb33a09d38b59a8fa8b7d92a1d82c8015e99f05f67ae9c9ae623157767959496.
//
// Solidity: event relayerRegister(address _relayer)
func (_Relayerhub *RelayerhubFilterer) FilterRelayerRegister(opts *bind.FilterOpts) (*RelayerhubRelayerRegisterIterator, error) {

	logs, sub, err := _Relayerhub.contract.FilterLogs(opts, "relayerRegister")
	if err != nil {
		return nil, err
	}
	return &RelayerhubRelayerRegisterIterator{contract: _Relayerhub.contract, event: "relayerRegister", logs: logs, sub: sub}, nil
}

// WatchRelayerRegister is a free log subscription operation binding the contract event 0xdb33a09d38b59a8fa8b7d92a1d82c8015e99f05f67ae9c9ae623157767959496.
//
// Solidity: event relayerRegister(address _relayer)
func (_Relayerhub *RelayerhubFilterer) WatchRelayerRegister(opts *bind.WatchOpts, sink chan<- *RelayerhubRelayerRegister) (event.Subscription, error) {

	logs, sub, err := _Relayerhub.contract.WatchLogs(opts, "relayerRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerhubRelayerRegister)
				if err := _Relayerhub.contract.UnpackLog(event, "relayerRegister", log); err != nil {
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

// ParseRelayerRegister is a log parse operation binding the contract event 0xdb33a09d38b59a8fa8b7d92a1d82c8015e99f05f67ae9c9ae623157767959496.
//
// Solidity: event relayerRegister(address _relayer)
func (_Relayerhub *RelayerhubFilterer) ParseRelayerRegister(log types.Log) (*RelayerhubRelayerRegister, error) {
	event := new(RelayerhubRelayerRegister)
	if err := _Relayerhub.contract.UnpackLog(event, "relayerRegister", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerhubRelayerUnRegisterIterator is returned from FilterRelayerUnRegister and is used to iterate over the raw logs and unpacked data for RelayerUnRegister events raised by the Relayerhub contract.
type RelayerhubRelayerUnRegisterIterator struct {
	Event *RelayerhubRelayerUnRegister // Event containing the contract specifics and raw log

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
func (it *RelayerhubRelayerUnRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerhubRelayerUnRegister)
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
		it.Event = new(RelayerhubRelayerUnRegister)
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
func (it *RelayerhubRelayerUnRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerhubRelayerUnRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerhubRelayerUnRegister represents a RelayerUnRegister event raised by the Relayerhub contract.
type RelayerhubRelayerUnRegister struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerUnRegister is a free log retrieval operation binding the contract event 0xd17202129b83db7880d6b9f25df81c58ad46f7e0e2c92236b1aa10663a487667.
//
// Solidity: event relayerUnRegister(address _relayer)
func (_Relayerhub *RelayerhubFilterer) FilterRelayerUnRegister(opts *bind.FilterOpts) (*RelayerhubRelayerUnRegisterIterator, error) {

	logs, sub, err := _Relayerhub.contract.FilterLogs(opts, "relayerUnRegister")
	if err != nil {
		return nil, err
	}
	return &RelayerhubRelayerUnRegisterIterator{contract: _Relayerhub.contract, event: "relayerUnRegister", logs: logs, sub: sub}, nil
}

// WatchRelayerUnRegister is a free log subscription operation binding the contract event 0xd17202129b83db7880d6b9f25df81c58ad46f7e0e2c92236b1aa10663a487667.
//
// Solidity: event relayerUnRegister(address _relayer)
func (_Relayerhub *RelayerhubFilterer) WatchRelayerUnRegister(opts *bind.WatchOpts, sink chan<- *RelayerhubRelayerUnRegister) (event.Subscription, error) {

	logs, sub, err := _Relayerhub.contract.WatchLogs(opts, "relayerUnRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerhubRelayerUnRegister)
				if err := _Relayerhub.contract.UnpackLog(event, "relayerUnRegister", log); err != nil {
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

// ParseRelayerUnRegister is a log parse operation binding the contract event 0xd17202129b83db7880d6b9f25df81c58ad46f7e0e2c92236b1aa10663a487667.
//
// Solidity: event relayerUnRegister(address _relayer)
func (_Relayerhub *RelayerhubFilterer) ParseRelayerUnRegister(log types.Log) (*RelayerhubRelayerUnRegister, error) {
	event := new(RelayerhubRelayerUnRegister)
	if err := _Relayerhub.contract.UnpackLog(event, "relayerUnRegister", log); err != nil {
		return nil, err
	}
	return event, nil
}
