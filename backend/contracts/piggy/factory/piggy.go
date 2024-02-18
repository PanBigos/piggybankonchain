// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FactoryMetaData contains all meta data concerning the Factory contract.
var FactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"piggyBank\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"}],\"name\":\"CreatedPiggyBank\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDate\",\"type\":\"uint256\"}],\"name\":\"createPiggyBank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"piggyBank\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getCreatedPiggyBanks\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getPiggyBanks\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use FactoryMetaData.ABI instead.
var FactoryABI = FactoryMetaData.ABI

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// GetCreatedPiggyBanks is a free data retrieval call binding the contract method 0x60a071a8.
//
// Solidity: function getCreatedPiggyBanks(address _user) view returns(address[])
func (_Factory *FactoryCaller) GetCreatedPiggyBanks(opts *bind.CallOpts, _user common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getCreatedPiggyBanks", _user)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCreatedPiggyBanks is a free data retrieval call binding the contract method 0x60a071a8.
//
// Solidity: function getCreatedPiggyBanks(address _user) view returns(address[])
func (_Factory *FactorySession) GetCreatedPiggyBanks(_user common.Address) ([]common.Address, error) {
	return _Factory.Contract.GetCreatedPiggyBanks(&_Factory.CallOpts, _user)
}

// GetCreatedPiggyBanks is a free data retrieval call binding the contract method 0x60a071a8.
//
// Solidity: function getCreatedPiggyBanks(address _user) view returns(address[])
func (_Factory *FactoryCallerSession) GetCreatedPiggyBanks(_user common.Address) ([]common.Address, error) {
	return _Factory.Contract.GetCreatedPiggyBanks(&_Factory.CallOpts, _user)
}

// GetPiggyBanks is a free data retrieval call binding the contract method 0x4f7b628a.
//
// Solidity: function getPiggyBanks(address _user) view returns(address[])
func (_Factory *FactoryCaller) GetPiggyBanks(opts *bind.CallOpts, _user common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getPiggyBanks", _user)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPiggyBanks is a free data retrieval call binding the contract method 0x4f7b628a.
//
// Solidity: function getPiggyBanks(address _user) view returns(address[])
func (_Factory *FactorySession) GetPiggyBanks(_user common.Address) ([]common.Address, error) {
	return _Factory.Contract.GetPiggyBanks(&_Factory.CallOpts, _user)
}

// GetPiggyBanks is a free data retrieval call binding the contract method 0x4f7b628a.
//
// Solidity: function getPiggyBanks(address _user) view returns(address[])
func (_Factory *FactoryCallerSession) GetPiggyBanks(_user common.Address) ([]common.Address, error) {
	return _Factory.Contract.GetPiggyBanks(&_Factory.CallOpts, _user)
}

// CreatePiggyBank is a paid mutator transaction binding the contract method 0x05c9e96e.
//
// Solidity: function createPiggyBank(address _owner, uint256 _unlockDate) returns(address piggyBank)
func (_Factory *FactoryTransactor) CreatePiggyBank(opts *bind.TransactOpts, _owner common.Address, _unlockDate *big.Int) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createPiggyBank", _owner, _unlockDate)
}

// CreatePiggyBank is a paid mutator transaction binding the contract method 0x05c9e96e.
//
// Solidity: function createPiggyBank(address _owner, uint256 _unlockDate) returns(address piggyBank)
func (_Factory *FactorySession) CreatePiggyBank(_owner common.Address, _unlockDate *big.Int) (*types.Transaction, error) {
	return _Factory.Contract.CreatePiggyBank(&_Factory.TransactOpts, _owner, _unlockDate)
}

// CreatePiggyBank is a paid mutator transaction binding the contract method 0x05c9e96e.
//
// Solidity: function createPiggyBank(address _owner, uint256 _unlockDate) returns(address piggyBank)
func (_Factory *FactoryTransactorSession) CreatePiggyBank(_owner common.Address, _unlockDate *big.Int) (*types.Transaction, error) {
	return _Factory.Contract.CreatePiggyBank(&_Factory.TransactOpts, _owner, _unlockDate)
}

// FactoryCreatedPiggyBankIterator is returned from FilterCreatedPiggyBank and is used to iterate over the raw logs and unpacked data for CreatedPiggyBank events raised by the Factory contract.
type FactoryCreatedPiggyBankIterator struct {
	Event *FactoryCreatedPiggyBank // Event containing the contract specifics and raw log

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
func (it *FactoryCreatedPiggyBankIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryCreatedPiggyBank)
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
		it.Event = new(FactoryCreatedPiggyBank)
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
func (it *FactoryCreatedPiggyBankIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryCreatedPiggyBankIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryCreatedPiggyBank represents a CreatedPiggyBank event raised by the Factory contract.
type FactoryCreatedPiggyBank struct {
	PiggyBank  common.Address
	Creator    common.Address
	Owner      common.Address
	CreatedAt  *big.Int
	UnlockDate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreatedPiggyBank is a free log retrieval operation binding the contract event 0x3dfd958f9baa46677a0ef1cf62438393326bb7ea0d1c04ba8198864b930f7afc.
//
// Solidity: event CreatedPiggyBank(address piggyBank, address creator, address owner, uint256 createdAt, uint256 unlockDate)
func (_Factory *FactoryFilterer) FilterCreatedPiggyBank(opts *bind.FilterOpts) (*FactoryCreatedPiggyBankIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "CreatedPiggyBank")
	if err != nil {
		return nil, err
	}
	return &FactoryCreatedPiggyBankIterator{contract: _Factory.contract, event: "CreatedPiggyBank", logs: logs, sub: sub}, nil
}

// WatchCreatedPiggyBank is a free log subscription operation binding the contract event 0x3dfd958f9baa46677a0ef1cf62438393326bb7ea0d1c04ba8198864b930f7afc.
//
// Solidity: event CreatedPiggyBank(address piggyBank, address creator, address owner, uint256 createdAt, uint256 unlockDate)
func (_Factory *FactoryFilterer) WatchCreatedPiggyBank(opts *bind.WatchOpts, sink chan<- *FactoryCreatedPiggyBank) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "CreatedPiggyBank")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryCreatedPiggyBank)
				if err := _Factory.contract.UnpackLog(event, "CreatedPiggyBank", log); err != nil {
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

// ParseCreatedPiggyBank is a log parse operation binding the contract event 0x3dfd958f9baa46677a0ef1cf62438393326bb7ea0d1c04ba8198864b930f7afc.
//
// Solidity: event CreatedPiggyBank(address piggyBank, address creator, address owner, uint256 createdAt, uint256 unlockDate)
func (_Factory *FactoryFilterer) ParseCreatedPiggyBank(log types.Log) (*FactoryCreatedPiggyBank, error) {
	event := new(FactoryCreatedPiggyBank)
	if err := _Factory.contract.UnpackLog(event, "CreatedPiggyBank", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
