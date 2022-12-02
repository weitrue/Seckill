// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package delegate

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
)

// DelegatefusionSignInfo is an auto generated low-level Go binding around an user-defined struct.
type DelegatefusionSignInfo struct {
	TokenIDs       []*big.Int
	ChainID        *big.Int
	FusionContract common.Address
	R              [32]byte
	S              [32]byte
	V              uint8
}

// DelegateMetaData contains all meta data concerning the Delegate contract.
var DelegateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fusionContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structDelegate.fusionSignInfo\",\"name\":\"signInfo\",\"type\":\"tuple\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// DelegateABI is the input ABI used to generate the binding from.
// Deprecated: Use DelegateMetaData.ABI instead.
var DelegateABI = DelegateMetaData.ABI

// Delegate is an auto generated Go binding around an Ethereum contract.
type Delegate struct {
	DelegateCaller     // Read-only binding to the contract
	DelegateTransactor // Write-only binding to the contract
	DelegateFilterer   // Log filterer for contract events
}

// DelegateCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegateSession struct {
	Contract     *Delegate         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegateCallerSession struct {
	Contract *DelegateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DelegateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegateTransactorSession struct {
	Contract     *DelegateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DelegateRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegateRaw struct {
	Contract *Delegate // Generic contract binding to access the raw methods on
}

// DelegateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegateCallerRaw struct {
	Contract *DelegateCaller // Generic read-only contract binding to access the raw methods on
}

// DelegateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegateTransactorRaw struct {
	Contract *DelegateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegate creates a new instance of Delegate, bound to a specific deployed contract.
func NewDelegate(address common.Address, backend bind.ContractBackend) (*Delegate, error) {
	contract, err := bindDelegate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Delegate{DelegateCaller: DelegateCaller{contract: contract}, DelegateTransactor: DelegateTransactor{contract: contract}, DelegateFilterer: DelegateFilterer{contract: contract}}, nil
}

// NewDelegateCaller creates a new read-only instance of Delegate, bound to a specific deployed contract.
func NewDelegateCaller(address common.Address, caller bind.ContractCaller) (*DelegateCaller, error) {
	contract, err := bindDelegate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegateCaller{contract: contract}, nil
}

// NewDelegateTransactor creates a new write-only instance of Delegate, bound to a specific deployed contract.
func NewDelegateTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegateTransactor, error) {
	contract, err := bindDelegate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegateTransactor{contract: contract}, nil
}

// NewDelegateFilterer creates a new log filterer instance of Delegate, bound to a specific deployed contract.
func NewDelegateFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegateFilterer, error) {
	contract, err := bindDelegate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegateFilterer{contract: contract}, nil
}

// bindDelegate binds a generic wrapper to an already deployed contract.
func bindDelegate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegate *DelegateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Delegate.Contract.DelegateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegate *DelegateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegate.Contract.DelegateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegate *DelegateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegate.Contract.DelegateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegate *DelegateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Delegate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegate *DelegateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegate *DelegateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegate.Contract.contract.Transact(opts, method, params...)
}

// VerifySignature is a free data retrieval call binding the contract method 0x39781e38.
//
// Solidity: function verifySignature((uint256[],uint256,address,bytes32,bytes32,uint8) signInfo) pure returns(bool)
func (_Delegate *DelegateCaller) VerifySignature(opts *bind.CallOpts, signInfo DelegatefusionSignInfo) (bool, error) {
	var out []interface{}
	err := _Delegate.contract.Call(opts, &out, "verifySignature", signInfo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySignature is a free data retrieval call binding the contract method 0x39781e38.
//
// Solidity: function verifySignature((uint256[],uint256,address,bytes32,bytes32,uint8) signInfo) pure returns(bool)
func (_Delegate *DelegateSession) VerifySignature(signInfo DelegatefusionSignInfo) (bool, error) {
	return _Delegate.Contract.VerifySignature(&_Delegate.CallOpts, signInfo)
}

// VerifySignature is a free data retrieval call binding the contract method 0x39781e38.
//
// Solidity: function verifySignature((uint256[],uint256,address,bytes32,bytes32,uint8) signInfo) pure returns(bool)
func (_Delegate *DelegateCallerSession) VerifySignature(signInfo DelegatefusionSignInfo) (bool, error) {
	return _Delegate.Contract.VerifySignature(&_Delegate.CallOpts, signInfo)
}
