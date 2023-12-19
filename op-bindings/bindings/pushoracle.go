// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// PushOracleMetaData contains all meta data concerning the PushOracle contract.
var PushOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DEPOSITOR_ACCOUNT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"data\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506101cf806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80631ab06ee514610046578063e591b2821461005b578063f0ba8440146100a0575b600080fd5b610059610054366004610187565b6100ce565b005b61007673deaddeaddeaddeaddeaddeaddeaddeaddead000181565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100c06100ae3660046101a9565b60006020819052908152604090205481565b604051908152602001610097565b3373deaddeaddeaddeaddeaddeaddeaddeaddead000114610175576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f507573684f7261636c653a206f6e6c7920746865206465706f7369746f72206160448201527f63636f756e740000000000000000000000000000000000000000000000000000606482015260840160405180910390fd5b60009182526020829052604090912055565b6000806040838503121561019a57600080fd5b50508035926020909101359150565b6000602082840312156101bb57600080fd5b503591905056fea164736f6c634300080f000a",
}

// PushOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use PushOracleMetaData.ABI instead.
var PushOracleABI = PushOracleMetaData.ABI

// PushOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PushOracleMetaData.Bin instead.
var PushOracleBin = PushOracleMetaData.Bin

// DeployPushOracle deploys a new Ethereum contract, binding an instance of PushOracle to it.
func DeployPushOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PushOracle, error) {
	parsed, err := PushOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PushOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PushOracle{PushOracleCaller: PushOracleCaller{contract: contract}, PushOracleTransactor: PushOracleTransactor{contract: contract}, PushOracleFilterer: PushOracleFilterer{contract: contract}}, nil
}

// PushOracle is an auto generated Go binding around an Ethereum contract.
type PushOracle struct {
	PushOracleCaller     // Read-only binding to the contract
	PushOracleTransactor // Write-only binding to the contract
	PushOracleFilterer   // Log filterer for contract events
}

// PushOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PushOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PushOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PushOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PushOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PushOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PushOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PushOracleSession struct {
	Contract     *PushOracle       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PushOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PushOracleCallerSession struct {
	Contract *PushOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PushOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PushOracleTransactorSession struct {
	Contract     *PushOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PushOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PushOracleRaw struct {
	Contract *PushOracle // Generic contract binding to access the raw methods on
}

// PushOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PushOracleCallerRaw struct {
	Contract *PushOracleCaller // Generic read-only contract binding to access the raw methods on
}

// PushOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PushOracleTransactorRaw struct {
	Contract *PushOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPushOracle creates a new instance of PushOracle, bound to a specific deployed contract.
func NewPushOracle(address common.Address, backend bind.ContractBackend) (*PushOracle, error) {
	contract, err := bindPushOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PushOracle{PushOracleCaller: PushOracleCaller{contract: contract}, PushOracleTransactor: PushOracleTransactor{contract: contract}, PushOracleFilterer: PushOracleFilterer{contract: contract}}, nil
}

// NewPushOracleCaller creates a new read-only instance of PushOracle, bound to a specific deployed contract.
func NewPushOracleCaller(address common.Address, caller bind.ContractCaller) (*PushOracleCaller, error) {
	contract, err := bindPushOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PushOracleCaller{contract: contract}, nil
}

// NewPushOracleTransactor creates a new write-only instance of PushOracle, bound to a specific deployed contract.
func NewPushOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*PushOracleTransactor, error) {
	contract, err := bindPushOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PushOracleTransactor{contract: contract}, nil
}

// NewPushOracleFilterer creates a new log filterer instance of PushOracle, bound to a specific deployed contract.
func NewPushOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*PushOracleFilterer, error) {
	contract, err := bindPushOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PushOracleFilterer{contract: contract}, nil
}

// bindPushOracle binds a generic wrapper to an already deployed contract.
func bindPushOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PushOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PushOracle *PushOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PushOracle.Contract.PushOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PushOracle *PushOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PushOracle.Contract.PushOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PushOracle *PushOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PushOracle.Contract.PushOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PushOracle *PushOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PushOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PushOracle *PushOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PushOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PushOracle *PushOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PushOracle.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_PushOracle *PushOracleCaller) DEPOSITORACCOUNT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PushOracle.contract.Call(opts, &out, "DEPOSITOR_ACCOUNT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_PushOracle *PushOracleSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _PushOracle.Contract.DEPOSITORACCOUNT(&_PushOracle.CallOpts)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_PushOracle *PushOracleCallerSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _PushOracle.Contract.DEPOSITORACCOUNT(&_PushOracle.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0xf0ba8440.
//
// Solidity: function data(uint256 ) view returns(uint256)
func (_PushOracle *PushOracleCaller) Data(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PushOracle.contract.Call(opts, &out, "data", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Data is a free data retrieval call binding the contract method 0xf0ba8440.
//
// Solidity: function data(uint256 ) view returns(uint256)
func (_PushOracle *PushOracleSession) Data(arg0 *big.Int) (*big.Int, error) {
	return _PushOracle.Contract.Data(&_PushOracle.CallOpts, arg0)
}

// Data is a free data retrieval call binding the contract method 0xf0ba8440.
//
// Solidity: function data(uint256 ) view returns(uint256)
func (_PushOracle *PushOracleCallerSession) Data(arg0 *big.Int) (*big.Int, error) {
	return _PushOracle.Contract.Data(&_PushOracle.CallOpts, arg0)
}

// Set is a paid mutator transaction binding the contract method 0x1ab06ee5.
//
// Solidity: function set(uint256 slot, uint256 value) returns()
func (_PushOracle *PushOracleTransactor) Set(opts *bind.TransactOpts, slot *big.Int, value *big.Int) (*types.Transaction, error) {
	return _PushOracle.contract.Transact(opts, "set", slot, value)
}

// Set is a paid mutator transaction binding the contract method 0x1ab06ee5.
//
// Solidity: function set(uint256 slot, uint256 value) returns()
func (_PushOracle *PushOracleSession) Set(slot *big.Int, value *big.Int) (*types.Transaction, error) {
	return _PushOracle.Contract.Set(&_PushOracle.TransactOpts, slot, value)
}

// Set is a paid mutator transaction binding the contract method 0x1ab06ee5.
//
// Solidity: function set(uint256 slot, uint256 value) returns()
func (_PushOracle *PushOracleTransactorSession) Set(slot *big.Int, value *big.Int) (*types.Transaction, error) {
	return _PushOracle.Contract.Set(&_PushOracle.TransactOpts, slot, value)
}
