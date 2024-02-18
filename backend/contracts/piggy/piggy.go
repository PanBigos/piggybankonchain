// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package piggy

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

// PiggyMetaData contains all meta data concerning the Piggy contract.
var PiggyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TokenTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newFeeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"}],\"name\":\"setFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferTokenWithFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PiggyABI is the input ABI used to generate the binding from.
// Deprecated: Use PiggyMetaData.ABI instead.
var PiggyABI = PiggyMetaData.ABI

// Piggy is an auto generated Go binding around an Ethereum contract.
type Piggy struct {
	PiggyCaller     // Read-only binding to the contract
	PiggyTransactor // Write-only binding to the contract
	PiggyFilterer   // Log filterer for contract events
}

// PiggyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PiggyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PiggyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PiggyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PiggyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PiggyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PiggySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PiggySession struct {
	Contract     *Piggy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PiggyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PiggyCallerSession struct {
	Contract *PiggyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PiggyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PiggyTransactorSession struct {
	Contract     *PiggyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PiggyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PiggyRaw struct {
	Contract *Piggy // Generic contract binding to access the raw methods on
}

// PiggyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PiggyCallerRaw struct {
	Contract *PiggyCaller // Generic read-only contract binding to access the raw methods on
}

// PiggyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PiggyTransactorRaw struct {
	Contract *PiggyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPiggy creates a new instance of Piggy, bound to a specific deployed contract.
func NewPiggy(address common.Address, backend bind.ContractBackend) (*Piggy, error) {
	contract, err := bindPiggy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Piggy{PiggyCaller: PiggyCaller{contract: contract}, PiggyTransactor: PiggyTransactor{contract: contract}, PiggyFilterer: PiggyFilterer{contract: contract}}, nil
}

// NewPiggyCaller creates a new read-only instance of Piggy, bound to a specific deployed contract.
func NewPiggyCaller(address common.Address, caller bind.ContractCaller) (*PiggyCaller, error) {
	contract, err := bindPiggy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PiggyCaller{contract: contract}, nil
}

// NewPiggyTransactor creates a new write-only instance of Piggy, bound to a specific deployed contract.
func NewPiggyTransactor(address common.Address, transactor bind.ContractTransactor) (*PiggyTransactor, error) {
	contract, err := bindPiggy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PiggyTransactor{contract: contract}, nil
}

// NewPiggyFilterer creates a new log filterer instance of Piggy, bound to a specific deployed contract.
func NewPiggyFilterer(address common.Address, filterer bind.ContractFilterer) (*PiggyFilterer, error) {
	contract, err := bindPiggy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PiggyFilterer{contract: contract}, nil
}

// bindPiggy binds a generic wrapper to an already deployed contract.
func bindPiggy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PiggyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Piggy *PiggyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Piggy.Contract.PiggyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Piggy *PiggyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Piggy.Contract.PiggyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Piggy *PiggyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Piggy.Contract.PiggyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Piggy *PiggyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Piggy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Piggy *PiggyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Piggy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Piggy *PiggyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Piggy.Contract.contract.Transact(opts, method, params...)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Piggy *PiggyCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Piggy.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Piggy *PiggySession) FeeCollector() (common.Address, error) {
	return _Piggy.Contract.FeeCollector(&_Piggy.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_Piggy *PiggyCallerSession) FeeCollector() (common.Address, error) {
	return _Piggy.Contract.FeeCollector(&_Piggy.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Piggy *PiggyCaller) FeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Piggy.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Piggy *PiggySession) FeeRate() (*big.Int, error) {
	return _Piggy.Contract.FeeRate(&_Piggy.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Piggy *PiggyCallerSession) FeeRate() (*big.Int, error) {
	return _Piggy.Contract.FeeRate(&_Piggy.CallOpts)
}

// MaxFee is a free data retrieval call binding the contract method 0x01f59d16.
//
// Solidity: function maxFee() view returns(uint256)
func (_Piggy *PiggyCaller) MaxFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Piggy.contract.Call(opts, &out, "maxFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFee is a free data retrieval call binding the contract method 0x01f59d16.
//
// Solidity: function maxFee() view returns(uint256)
func (_Piggy *PiggySession) MaxFee() (*big.Int, error) {
	return _Piggy.Contract.MaxFee(&_Piggy.CallOpts)
}

// MaxFee is a free data retrieval call binding the contract method 0x01f59d16.
//
// Solidity: function maxFee() view returns(uint256)
func (_Piggy *PiggyCallerSession) MaxFee() (*big.Int, error) {
	return _Piggy.Contract.MaxFee(&_Piggy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Piggy *PiggyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Piggy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Piggy *PiggySession) Owner() (common.Address, error) {
	return _Piggy.Contract.Owner(&_Piggy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Piggy *PiggyCallerSession) Owner() (common.Address, error) {
	return _Piggy.Contract.Owner(&_Piggy.CallOpts)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _newFeeCollector) returns()
func (_Piggy *PiggyTransactor) SetFeeCollector(opts *bind.TransactOpts, _newFeeCollector common.Address) (*types.Transaction, error) {
	return _Piggy.contract.Transact(opts, "setFeeCollector", _newFeeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _newFeeCollector) returns()
func (_Piggy *PiggySession) SetFeeCollector(_newFeeCollector common.Address) (*types.Transaction, error) {
	return _Piggy.Contract.SetFeeCollector(&_Piggy.TransactOpts, _newFeeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _newFeeCollector) returns()
func (_Piggy *PiggyTransactorSession) SetFeeCollector(_newFeeCollector common.Address) (*types.Transaction, error) {
	return _Piggy.Contract.SetFeeCollector(&_Piggy.TransactOpts, _newFeeCollector)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _newRate) returns()
func (_Piggy *PiggyTransactor) SetFeeRate(opts *bind.TransactOpts, _newRate *big.Int) (*types.Transaction, error) {
	return _Piggy.contract.Transact(opts, "setFeeRate", _newRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _newRate) returns()
func (_Piggy *PiggySession) SetFeeRate(_newRate *big.Int) (*types.Transaction, error) {
	return _Piggy.Contract.SetFeeRate(&_Piggy.TransactOpts, _newRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _newRate) returns()
func (_Piggy *PiggyTransactorSession) SetFeeRate(_newRate *big.Int) (*types.Transaction, error) {
	return _Piggy.Contract.SetFeeRate(&_Piggy.TransactOpts, _newRate)
}

// TransferTokenWithFee is a paid mutator transaction binding the contract method 0x2be168db.
//
// Solidity: function transferTokenWithFee(address token, address _to, uint256 _amount) returns()
func (_Piggy *PiggyTransactor) TransferTokenWithFee(opts *bind.TransactOpts, token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Piggy.contract.Transact(opts, "transferTokenWithFee", token, _to, _amount)
}

// TransferTokenWithFee is a paid mutator transaction binding the contract method 0x2be168db.
//
// Solidity: function transferTokenWithFee(address token, address _to, uint256 _amount) returns()
func (_Piggy *PiggySession) TransferTokenWithFee(token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Piggy.Contract.TransferTokenWithFee(&_Piggy.TransactOpts, token, _to, _amount)
}

// TransferTokenWithFee is a paid mutator transaction binding the contract method 0x2be168db.
//
// Solidity: function transferTokenWithFee(address token, address _to, uint256 _amount) returns()
func (_Piggy *PiggyTransactorSession) TransferTokenWithFee(token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Piggy.Contract.TransferTokenWithFee(&_Piggy.TransactOpts, token, _to, _amount)
}

// TransferWithFee is a paid mutator transaction binding the contract method 0xe73de08d.
//
// Solidity: function transferWithFee(address _to) payable returns()
func (_Piggy *PiggyTransactor) TransferWithFee(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Piggy.contract.Transact(opts, "transferWithFee", _to)
}

// TransferWithFee is a paid mutator transaction binding the contract method 0xe73de08d.
//
// Solidity: function transferWithFee(address _to) payable returns()
func (_Piggy *PiggySession) TransferWithFee(_to common.Address) (*types.Transaction, error) {
	return _Piggy.Contract.TransferWithFee(&_Piggy.TransactOpts, _to)
}

// TransferWithFee is a paid mutator transaction binding the contract method 0xe73de08d.
//
// Solidity: function transferWithFee(address _to) payable returns()
func (_Piggy *PiggyTransactorSession) TransferWithFee(_to common.Address) (*types.Transaction, error) {
	return _Piggy.Contract.TransferWithFee(&_Piggy.TransactOpts, _to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Piggy *PiggyTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Piggy.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Piggy *PiggySession) Withdraw() (*types.Transaction, error) {
	return _Piggy.Contract.Withdraw(&_Piggy.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Piggy *PiggyTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Piggy.Contract.Withdraw(&_Piggy.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address token) returns()
func (_Piggy *PiggyTransactor) WithdrawTokens(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Piggy.contract.Transact(opts, "withdrawTokens", token)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address token) returns()
func (_Piggy *PiggySession) WithdrawTokens(token common.Address) (*types.Transaction, error) {
	return _Piggy.Contract.WithdrawTokens(&_Piggy.TransactOpts, token)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address token) returns()
func (_Piggy *PiggyTransactorSession) WithdrawTokens(token common.Address) (*types.Transaction, error) {
	return _Piggy.Contract.WithdrawTokens(&_Piggy.TransactOpts, token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Piggy *PiggyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Piggy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Piggy *PiggySession) Receive() (*types.Transaction, error) {
	return _Piggy.Contract.Receive(&_Piggy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Piggy *PiggyTransactorSession) Receive() (*types.Transaction, error) {
	return _Piggy.Contract.Receive(&_Piggy.TransactOpts)
}

// PiggyTokenTransferIterator is returned from FilterTokenTransfer and is used to iterate over the raw logs and unpacked data for TokenTransfer events raised by the Piggy contract.
type PiggyTokenTransferIterator struct {
	Event *PiggyTokenTransfer // Event containing the contract specifics and raw log

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
func (it *PiggyTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PiggyTokenTransfer)
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
		it.Event = new(PiggyTokenTransfer)
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
func (it *PiggyTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PiggyTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PiggyTokenTransfer represents a TokenTransfer event raised by the Piggy contract.
type PiggyTokenTransfer struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenTransfer is a free log retrieval operation binding the contract event 0xbba8a6f1ace6d0ccb2089d879d1bf044d9153802c1d010c514711798d413828c.
//
// Solidity: event TokenTransfer(address indexed token, address indexed to, uint256 amount, uint256 fee)
func (_Piggy *PiggyFilterer) FilterTokenTransfer(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*PiggyTokenTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Piggy.contract.FilterLogs(opts, "TokenTransfer", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PiggyTokenTransferIterator{contract: _Piggy.contract, event: "TokenTransfer", logs: logs, sub: sub}, nil
}

// WatchTokenTransfer is a free log subscription operation binding the contract event 0xbba8a6f1ace6d0ccb2089d879d1bf044d9153802c1d010c514711798d413828c.
//
// Solidity: event TokenTransfer(address indexed token, address indexed to, uint256 amount, uint256 fee)
func (_Piggy *PiggyFilterer) WatchTokenTransfer(opts *bind.WatchOpts, sink chan<- *PiggyTokenTransfer, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Piggy.contract.WatchLogs(opts, "TokenTransfer", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PiggyTokenTransfer)
				if err := _Piggy.contract.UnpackLog(event, "TokenTransfer", log); err != nil {
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

// ParseTokenTransfer is a log parse operation binding the contract event 0xbba8a6f1ace6d0ccb2089d879d1bf044d9153802c1d010c514711798d413828c.
//
// Solidity: event TokenTransfer(address indexed token, address indexed to, uint256 amount, uint256 fee)
func (_Piggy *PiggyFilterer) ParseTokenTransfer(log types.Log) (*PiggyTokenTransfer, error) {
	event := new(PiggyTokenTransfer)
	if err := _Piggy.contract.UnpackLog(event, "TokenTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PiggyTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Piggy contract.
type PiggyTransferIterator struct {
	Event *PiggyTransfer // Event containing the contract specifics and raw log

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
func (it *PiggyTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PiggyTransfer)
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
		it.Event = new(PiggyTransfer)
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
func (it *PiggyTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PiggyTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PiggyTransfer represents a Transfer event raised by the Piggy contract.
type PiggyTransfer struct {
	To     common.Address
	Amount *big.Int
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x7fa9aafeb8bb803d77de5d84bc2f2edbd842ca91b20cd5020aa21dfe26ab0be9.
//
// Solidity: event Transfer(address indexed to, uint256 amount, uint256 fee)
func (_Piggy *PiggyFilterer) FilterTransfer(opts *bind.FilterOpts, to []common.Address) (*PiggyTransferIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Piggy.contract.FilterLogs(opts, "Transfer", toRule)
	if err != nil {
		return nil, err
	}
	return &PiggyTransferIterator{contract: _Piggy.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x7fa9aafeb8bb803d77de5d84bc2f2edbd842ca91b20cd5020aa21dfe26ab0be9.
//
// Solidity: event Transfer(address indexed to, uint256 amount, uint256 fee)
func (_Piggy *PiggyFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PiggyTransfer, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Piggy.contract.WatchLogs(opts, "Transfer", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PiggyTransfer)
				if err := _Piggy.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x7fa9aafeb8bb803d77de5d84bc2f2edbd842ca91b20cd5020aa21dfe26ab0be9.
//
// Solidity: event Transfer(address indexed to, uint256 amount, uint256 fee)
func (_Piggy *PiggyFilterer) ParseTransfer(log types.Log) (*PiggyTransfer, error) {
	event := new(PiggyTransfer)
	if err := _Piggy.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
