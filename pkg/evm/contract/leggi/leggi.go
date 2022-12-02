// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package service

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LeggiNFTStageMintConfig is an auto generated low-level Go binding around an user-defined struct.
type LeggiNFTStageMintConfig struct {
	StageNum              uint64
	MaxPerStage           uint64
	MaxPerAddress         uint64
	IsWhiteListMintActive bool
	IsPublicMintActive    bool
	BeginTime             uint64
	EndTime               uint64
}

// LeggiABI is the input ABI used to generate the binding from.
const LeggiABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"stageNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxPerStage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxPerAddress\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isWhiteListMintActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPublicMintActive\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"beginTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structLeggiNFT.StageMintConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"StageMintConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"quantity\",\"type\":\"uint64\"}],\"name\":\"publicMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"stageNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxPerStage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxPerAddress\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isWhiteListMintActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPublicMintActive\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"beginTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"internalType\":\"structLeggiNFT.StageMintConfig\",\"name\":\"config_\",\"type\":\"tuple\"}],\"name\":\"setStageMintConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"}],\"name\":\"setTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stageMintConfig\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"stageNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxPerStage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxPerAddress\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isWhiteListMintActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPublicMintActive\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"beginTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Leggi is an auto generated Go binding around an Ethereum contract.
type Leggi struct {
	LeggiCaller     // Read-only binding to the contract
	LeggiTransactor // Write-only binding to the contract
	LeggiFilterer   // Log filterer for contract events
}

// LeggiCaller is an auto generated read-only Go binding around an Ethereum contract.
type LeggiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeggiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LeggiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeggiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LeggiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeggiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LeggiSession struct {
	Contract     *Leggi            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeggiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LeggiCallerSession struct {
	Contract *LeggiCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LeggiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LeggiTransactorSession struct {
	Contract     *LeggiTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeggiRaw is an auto generated low-level Go binding around an Ethereum contract.
type LeggiRaw struct {
	Contract *Leggi // Generic contract binding to access the raw methods on
}

// LeggiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LeggiCallerRaw struct {
	Contract *LeggiCaller // Generic read-only contract binding to access the raw methods on
}

// LeggiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LeggiTransactorRaw struct {
	Contract *LeggiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLeggi creates a new instance of Leggi, bound to a specific deployed contract.
func NewLeggi(address common.Address, backend bind.ContractBackend) (*Leggi, error) {
	contract, err := bindLeggi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Leggi{LeggiCaller: LeggiCaller{contract: contract}, LeggiTransactor: LeggiTransactor{contract: contract}, LeggiFilterer: LeggiFilterer{contract: contract}}, nil
}

// NewLeggiCaller creates a new read-only instance of Leggi, bound to a specific deployed contract.
func NewLeggiCaller(address common.Address, caller bind.ContractCaller) (*LeggiCaller, error) {
	contract, err := bindLeggi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LeggiCaller{contract: contract}, nil
}

// NewLeggiTransactor creates a new write-only instance of Leggi, bound to a specific deployed contract.
func NewLeggiTransactor(address common.Address, transactor bind.ContractTransactor) (*LeggiTransactor, error) {
	contract, err := bindLeggi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LeggiTransactor{contract: contract}, nil
}

// NewLeggiFilterer creates a new log filterer instance of Leggi, bound to a specific deployed contract.
func NewLeggiFilterer(address common.Address, filterer bind.ContractFilterer) (*LeggiFilterer, error) {
	contract, err := bindLeggi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LeggiFilterer{contract: contract}, nil
}

// bindLeggi binds a generic wrapper to an already deployed contract.
func bindLeggi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LeggiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Leggi *LeggiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Leggi.Contract.LeggiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Leggi *LeggiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leggi.Contract.LeggiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Leggi *LeggiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Leggi.Contract.LeggiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Leggi *LeggiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Leggi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Leggi *LeggiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leggi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Leggi *LeggiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Leggi.Contract.contract.Transact(opts, method, params...)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Leggi *LeggiCaller) MAXSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "MAX_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Leggi *LeggiSession) MAXSUPPLY() (*big.Int, error) {
	return _Leggi.Contract.MAXSUPPLY(&_Leggi.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Leggi *LeggiCallerSession) MAXSUPPLY() (*big.Int, error) {
	return _Leggi.Contract.MAXSUPPLY(&_Leggi.CallOpts)
}

// AddressMinted is a free data retrieval call binding the contract method 0xf4a66808.
//
// Solidity: function addressMinted(uint256 , address ) view returns(uint256)
func (_Leggi *LeggiCaller) AddressMinted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "addressMinted", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressMinted is a free data retrieval call binding the contract method 0xf4a66808.
//
// Solidity: function addressMinted(uint256 , address ) view returns(uint256)
func (_Leggi *LeggiSession) AddressMinted(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Leggi.Contract.AddressMinted(&_Leggi.CallOpts, arg0, arg1)
}

// AddressMinted is a free data retrieval call binding the contract method 0xf4a66808.
//
// Solidity: function addressMinted(uint256 , address ) view returns(uint256)
func (_Leggi *LeggiCallerSession) AddressMinted(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Leggi.Contract.AddressMinted(&_Leggi.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Leggi *LeggiCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Leggi *LeggiSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Leggi.Contract.BalanceOf(&_Leggi.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Leggi *LeggiCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Leggi.Contract.BalanceOf(&_Leggi.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Leggi *LeggiCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Leggi *LeggiSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Leggi.Contract.GetApproved(&_Leggi.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Leggi *LeggiCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Leggi.Contract.GetApproved(&_Leggi.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Leggi *LeggiCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Leggi *LeggiSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Leggi.Contract.IsApprovedForAll(&_Leggi.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Leggi *LeggiCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Leggi.Contract.IsApprovedForAll(&_Leggi.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Leggi *LeggiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Leggi *LeggiSession) Name() (string, error) {
	return _Leggi.Contract.Name(&_Leggi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Leggi *LeggiCallerSession) Name() (string, error) {
	return _Leggi.Contract.Name(&_Leggi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Leggi *LeggiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Leggi *LeggiSession) Owner() (common.Address, error) {
	return _Leggi.Contract.Owner(&_Leggi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Leggi *LeggiCallerSession) Owner() (common.Address, error) {
	return _Leggi.Contract.Owner(&_Leggi.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Leggi *LeggiCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Leggi *LeggiSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Leggi.Contract.OwnerOf(&_Leggi.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Leggi *LeggiCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Leggi.Contract.OwnerOf(&_Leggi.CallOpts, tokenId)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Leggi *LeggiCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})

	outstruct.Receiver = out[0].(common.Address)
	outstruct.RoyaltyAmount = out[1].(*big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Leggi *LeggiSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Leggi.Contract.RoyaltyInfo(&_Leggi.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Leggi *LeggiCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Leggi.Contract.RoyaltyInfo(&_Leggi.CallOpts, tokenId, salePrice)
}

// StageMintConfig is a free data retrieval call binding the contract method 0x3173ea1a.
//
// Solidity: function stageMintConfig() view returns(uint64 stageNum, uint64 maxPerStage, uint64 maxPerAddress, bool isWhiteListMintActive, bool isPublicMintActive, uint64 beginTime, uint64 endTime)
func (_Leggi *LeggiCaller) StageMintConfig(opts *bind.CallOpts) (struct {
	StageNum              uint64
	MaxPerStage           uint64
	MaxPerAddress         uint64
	IsWhiteListMintActive bool
	IsPublicMintActive    bool
	BeginTime             uint64
	EndTime               uint64
}, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "stageMintConfig")

	outstruct := new(struct {
		StageNum              uint64
		MaxPerStage           uint64
		MaxPerAddress         uint64
		IsWhiteListMintActive bool
		IsPublicMintActive    bool
		BeginTime             uint64
		EndTime               uint64
	})

	outstruct.StageNum = out[0].(uint64)
	outstruct.MaxPerStage = out[1].(uint64)
	outstruct.MaxPerAddress = out[2].(uint64)
	outstruct.IsWhiteListMintActive = out[3].(bool)
	outstruct.IsPublicMintActive = out[4].(bool)
	outstruct.BeginTime = out[5].(uint64)
	outstruct.EndTime = out[6].(uint64)

	return *outstruct, err

}

// StageMintConfig is a free data retrieval call binding the contract method 0x3173ea1a.
//
// Solidity: function stageMintConfig() view returns(uint64 stageNum, uint64 maxPerStage, uint64 maxPerAddress, bool isWhiteListMintActive, bool isPublicMintActive, uint64 beginTime, uint64 endTime)
func (_Leggi *LeggiSession) StageMintConfig() (struct {
	StageNum              uint64
	MaxPerStage           uint64
	MaxPerAddress         uint64
	IsWhiteListMintActive bool
	IsPublicMintActive    bool
	BeginTime             uint64
	EndTime               uint64
}, error) {
	return _Leggi.Contract.StageMintConfig(&_Leggi.CallOpts)
}

// StageMintConfig is a free data retrieval call binding the contract method 0x3173ea1a.
//
// Solidity: function stageMintConfig() view returns(uint64 stageNum, uint64 maxPerStage, uint64 maxPerAddress, bool isWhiteListMintActive, bool isPublicMintActive, uint64 beginTime, uint64 endTime)
func (_Leggi *LeggiCallerSession) StageMintConfig() (struct {
	StageNum              uint64
	MaxPerStage           uint64
	MaxPerAddress         uint64
	IsWhiteListMintActive bool
	IsPublicMintActive    bool
	BeginTime             uint64
	EndTime               uint64
}, error) {
	return _Leggi.Contract.StageMintConfig(&_Leggi.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Leggi *LeggiCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Leggi *LeggiSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Leggi.Contract.SupportsInterface(&_Leggi.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Leggi *LeggiCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Leggi.Contract.SupportsInterface(&_Leggi.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Leggi *LeggiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Leggi *LeggiSession) Symbol() (string, error) {
	return _Leggi.Contract.Symbol(&_Leggi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Leggi *LeggiCallerSession) Symbol() (string, error) {
	return _Leggi.Contract.Symbol(&_Leggi.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Leggi *LeggiCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Leggi *LeggiSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Leggi.Contract.TokenURI(&_Leggi.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Leggi *LeggiCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Leggi.Contract.TokenURI(&_Leggi.CallOpts, tokenId)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Leggi *LeggiCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Leggi *LeggiSession) TotalMinted() (*big.Int, error) {
	return _Leggi.Contract.TotalMinted(&_Leggi.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Leggi *LeggiCallerSession) TotalMinted() (*big.Int, error) {
	return _Leggi.Contract.TotalMinted(&_Leggi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Leggi *LeggiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leggi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Leggi *LeggiSession) TotalSupply() (*big.Int, error) {
	return _Leggi.Contract.TotalSupply(&_Leggi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Leggi *LeggiCallerSession) TotalSupply() (*big.Int, error) {
	return _Leggi.Contract.TotalSupply(&_Leggi.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Leggi *LeggiTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leggi.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Leggi *LeggiSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leggi.Contract.Approve(&_Leggi.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Leggi *LeggiTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leggi.Contract.Approve(&_Leggi.TransactOpts, to, tokenId)
}

// PublicMint is a paid mutator transaction binding the contract method 0x6afcb7b0.
//
// Solidity: function publicMint(uint64 quantity) returns()
func (_Leggi *LeggiTransactor) PublicMint(opts *bind.TransactOpts, quantity uint64) (*types.Transaction, error) {
	return _Leggi.contract.Transact(opts, "publicMint", quantity)
}

// PublicMint is a paid mutator transaction binding the contract method 0x6afcb7b0.
//
// Solidity: function publicMint(uint64 quantity) returns()
func (_Leggi *LeggiSession) PublicMint(quantity uint64) (*types.Transaction, error) {
	return _Leggi.Contract.PublicMint(&_Leggi.TransactOpts, quantity)
}

// PublicMint is a paid mutator transaction binding the contract method 0x6afcb7b0.
//
// Solidity: function publicMint(uint64 quantity) returns()
func (_Leggi *LeggiTransactorSession) PublicMint(quantity uint64) (*types.Transaction, error) {
	return _Leggi.Contract.PublicMint(&_Leggi.TransactOpts, quantity)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Leggi *LeggiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leggi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Leggi *LeggiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Leggi.Contract.RenounceOwnership(&_Leggi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Leggi *LeggiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Leggi.Contract.RenounceOwnership(&_Leggi.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Leggi *LeggiTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leggi.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Leggi *LeggiSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leggi.Contract.SafeTransferFrom(&_Leggi.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Leggi *LeggiTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leggi.Contract.SafeTransferFrom(&_Leggi.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Leggi *LeggiTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Leggi.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Leggi *LeggiSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Leggi.Contract.SafeTransferFrom0(&_Leggi.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Leggi *LeggiTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Leggi.Contract.SafeTransferFrom0(&_Leggi.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Leggi *LeggiTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Leggi.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Leggi *LeggiSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Leggi.Contract.SetApprovalForAll(&_Leggi.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Leggi *LeggiTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Leggi.Contract.SetApprovalForAll(&_Leggi.TransactOpts, operator, approved)
}

// SetStageMintConfig is a paid mutator transaction binding the contract method 0x96d5e0af.
//
// Solidity: function setStageMintConfig((uint64,uint64,uint64,bool,bool,uint64,uint64) config_) returns()
func (_Leggi *LeggiTransactor) SetStageMintConfig(opts *bind.TransactOpts, config_ LeggiNFTStageMintConfig) (*types.Transaction, error) {
	return _Leggi.contract.Transact(opts, "setStageMintConfig", config_)
}

// SetStageMintConfig is a paid mutator transaction binding the contract method 0x96d5e0af.
//
// Solidity: function setStageMintConfig((uint64,uint64,uint64,bool,bool,uint64,uint64) config_) returns()
func (_Leggi *LeggiSession) SetStageMintConfig(config_ LeggiNFTStageMintConfig) (*types.Transaction, error) {
	return _Leggi.Contract.SetStageMintConfig(&_Leggi.TransactOpts, config_)
}

// SetStageMintConfig is a paid mutator transaction binding the contract method 0x96d5e0af.
//
// Solidity: function setStageMintConfig((uint64,uint64,uint64,bool,bool,uint64,uint64) config_) returns()
func (_Leggi *LeggiTransactorSession) SetStageMintConfig(config_ LeggiNFTStageMintConfig) (*types.Transaction, error) {
	return _Leggi.Contract.SetStageMintConfig(&_Leggi.TransactOpts, config_)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0xe0df5b6f.
//
// Solidity: function setTokenURI(string tokenURI_) returns()
func (_Leggi *LeggiTransactor) SetTokenURI(opts *bind.TransactOpts, tokenURI_ string) (*types.Transaction, error) {
	return _Leggi.contract.Transact(opts, "setTokenURI", tokenURI_)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0xe0df5b6f.
//
// Solidity: function setTokenURI(string tokenURI_) returns()
func (_Leggi *LeggiSession) SetTokenURI(tokenURI_ string) (*types.Transaction, error) {
	return _Leggi.Contract.SetTokenURI(&_Leggi.TransactOpts, tokenURI_)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0xe0df5b6f.
//
// Solidity: function setTokenURI(string tokenURI_) returns()
func (_Leggi *LeggiTransactorSession) SetTokenURI(tokenURI_ string) (*types.Transaction, error) {
	return _Leggi.Contract.SetTokenURI(&_Leggi.TransactOpts, tokenURI_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Leggi *LeggiTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leggi.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Leggi *LeggiSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leggi.Contract.TransferFrom(&_Leggi.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Leggi *LeggiTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Leggi.Contract.TransferFrom(&_Leggi.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Leggi *LeggiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Leggi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Leggi *LeggiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Leggi.Contract.TransferOwnership(&_Leggi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Leggi *LeggiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Leggi.Contract.TransferOwnership(&_Leggi.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Leggi *LeggiTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leggi.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Leggi *LeggiSession) Withdraw() (*types.Transaction, error) {
	return _Leggi.Contract.Withdraw(&_Leggi.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Leggi *LeggiTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Leggi.Contract.Withdraw(&_Leggi.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address token) returns()
func (_Leggi *LeggiTransactor) WithdrawTokens(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Leggi.contract.Transact(opts, "withdrawTokens", token)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address token) returns()
func (_Leggi *LeggiSession) WithdrawTokens(token common.Address) (*types.Transaction, error) {
	return _Leggi.Contract.WithdrawTokens(&_Leggi.TransactOpts, token)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address token) returns()
func (_Leggi *LeggiTransactorSession) WithdrawTokens(token common.Address) (*types.Transaction, error) {
	return _Leggi.Contract.WithdrawTokens(&_Leggi.TransactOpts, token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Leggi *LeggiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leggi.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Leggi *LeggiSession) Receive() (*types.Transaction, error) {
	return _Leggi.Contract.Receive(&_Leggi.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Leggi *LeggiTransactorSession) Receive() (*types.Transaction, error) {
	return _Leggi.Contract.Receive(&_Leggi.TransactOpts)
}

// LeggiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Leggi contract.
type LeggiApprovalIterator struct {
	Event *LeggiApproval // Event containing the contract specifics and raw log

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
func (it *LeggiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeggiApproval)
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
		it.Event = new(LeggiApproval)
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
func (it *LeggiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeggiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeggiApproval represents a Approval event raised by the Leggi contract.
type LeggiApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Leggi *LeggiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LeggiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Leggi.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LeggiApprovalIterator{contract: _Leggi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Leggi *LeggiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LeggiApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Leggi.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeggiApproval)
				if err := _Leggi.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Leggi *LeggiFilterer) ParseApproval(log types.Log) (*LeggiApproval, error) {
	event := new(LeggiApproval)
	if err := _Leggi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeggiApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Leggi contract.
type LeggiApprovalForAllIterator struct {
	Event *LeggiApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LeggiApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeggiApprovalForAll)
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
		it.Event = new(LeggiApprovalForAll)
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
func (it *LeggiApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeggiApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeggiApprovalForAll represents a ApprovalForAll event raised by the Leggi contract.
type LeggiApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Leggi *LeggiFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LeggiApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Leggi.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LeggiApprovalForAllIterator{contract: _Leggi.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Leggi *LeggiFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LeggiApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Leggi.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeggiApprovalForAll)
				if err := _Leggi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Leggi *LeggiFilterer) ParseApprovalForAll(log types.Log) (*LeggiApprovalForAll, error) {
	event := new(LeggiApprovalForAll)
	if err := _Leggi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeggiConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Leggi contract.
type LeggiConsecutiveTransferIterator struct {
	Event *LeggiConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *LeggiConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeggiConsecutiveTransfer)
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
		it.Event = new(LeggiConsecutiveTransfer)
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
func (it *LeggiConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeggiConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeggiConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Leggi contract.
type LeggiConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Leggi *LeggiFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*LeggiConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Leggi.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LeggiConsecutiveTransferIterator{contract: _Leggi.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Leggi *LeggiFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *LeggiConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Leggi.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeggiConsecutiveTransfer)
				if err := _Leggi.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Leggi *LeggiFilterer) ParseConsecutiveTransfer(log types.Log) (*LeggiConsecutiveTransfer, error) {
	event := new(LeggiConsecutiveTransfer)
	if err := _Leggi.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeggiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Leggi contract.
type LeggiOwnershipTransferredIterator struct {
	Event *LeggiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LeggiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeggiOwnershipTransferred)
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
		it.Event = new(LeggiOwnershipTransferred)
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
func (it *LeggiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeggiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeggiOwnershipTransferred represents a OwnershipTransferred event raised by the Leggi contract.
type LeggiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Leggi *LeggiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LeggiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Leggi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LeggiOwnershipTransferredIterator{contract: _Leggi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Leggi *LeggiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LeggiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Leggi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeggiOwnershipTransferred)
				if err := _Leggi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Leggi *LeggiFilterer) ParseOwnershipTransferred(log types.Log) (*LeggiOwnershipTransferred, error) {
	event := new(LeggiOwnershipTransferred)
	if err := _Leggi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeggiReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Leggi contract.
type LeggiReceivedIterator struct {
	Event *LeggiReceived // Event containing the contract specifics and raw log

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
func (it *LeggiReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeggiReceived)
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
		it.Event = new(LeggiReceived)
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
func (it *LeggiReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeggiReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeggiReceived represents a Received event raised by the Leggi contract.
type LeggiReceived struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed arg0, uint256 arg1)
func (_Leggi *LeggiFilterer) FilterReceived(opts *bind.FilterOpts, arg0 []common.Address) (*LeggiReceivedIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Leggi.contract.FilterLogs(opts, "Received", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &LeggiReceivedIterator{contract: _Leggi.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed arg0, uint256 arg1)
func (_Leggi *LeggiFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *LeggiReceived, arg0 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Leggi.contract.WatchLogs(opts, "Received", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeggiReceived)
				if err := _Leggi.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed arg0, uint256 arg1)
func (_Leggi *LeggiFilterer) ParseReceived(log types.Log) (*LeggiReceived, error) {
	event := new(LeggiReceived)
	if err := _Leggi.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeggiStageMintConfigChangedIterator is returned from FilterStageMintConfigChanged and is used to iterate over the raw logs and unpacked data for StageMintConfigChanged events raised by the Leggi contract.
type LeggiStageMintConfigChangedIterator struct {
	Event *LeggiStageMintConfigChanged // Event containing the contract specifics and raw log

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
func (it *LeggiStageMintConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeggiStageMintConfigChanged)
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
		it.Event = new(LeggiStageMintConfigChanged)
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
func (it *LeggiStageMintConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeggiStageMintConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeggiStageMintConfigChanged represents a StageMintConfigChanged event raised by the Leggi contract.
type LeggiStageMintConfigChanged struct {
	Config LeggiNFTStageMintConfig
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStageMintConfigChanged is a free log retrieval operation binding the contract event 0xd37212ab2e34cdb607238f80545d9144ce37490ecac6d116d15c1ce74bd06fee.
//
// Solidity: event StageMintConfigChanged((uint64,uint64,uint64,bool,bool,uint64,uint64) config)
func (_Leggi *LeggiFilterer) FilterStageMintConfigChanged(opts *bind.FilterOpts) (*LeggiStageMintConfigChangedIterator, error) {

	logs, sub, err := _Leggi.contract.FilterLogs(opts, "StageMintConfigChanged")
	if err != nil {
		return nil, err
	}
	return &LeggiStageMintConfigChangedIterator{contract: _Leggi.contract, event: "StageMintConfigChanged", logs: logs, sub: sub}, nil
}

// WatchStageMintConfigChanged is a free log subscription operation binding the contract event 0xd37212ab2e34cdb607238f80545d9144ce37490ecac6d116d15c1ce74bd06fee.
//
// Solidity: event StageMintConfigChanged((uint64,uint64,uint64,bool,bool,uint64,uint64) config)
func (_Leggi *LeggiFilterer) WatchStageMintConfigChanged(opts *bind.WatchOpts, sink chan<- *LeggiStageMintConfigChanged) (event.Subscription, error) {

	logs, sub, err := _Leggi.contract.WatchLogs(opts, "StageMintConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeggiStageMintConfigChanged)
				if err := _Leggi.contract.UnpackLog(event, "StageMintConfigChanged", log); err != nil {
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

// ParseStageMintConfigChanged is a log parse operation binding the contract event 0xd37212ab2e34cdb607238f80545d9144ce37490ecac6d116d15c1ce74bd06fee.
//
// Solidity: event StageMintConfigChanged((uint64,uint64,uint64,bool,bool,uint64,uint64) config)
func (_Leggi *LeggiFilterer) ParseStageMintConfigChanged(log types.Log) (*LeggiStageMintConfigChanged, error) {
	event := new(LeggiStageMintConfigChanged)
	if err := _Leggi.contract.UnpackLog(event, "StageMintConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeggiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Leggi contract.
type LeggiTransferIterator struct {
	Event *LeggiTransfer // Event containing the contract specifics and raw log

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
func (it *LeggiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeggiTransfer)
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
		it.Event = new(LeggiTransfer)
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
func (it *LeggiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeggiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeggiTransfer represents a Transfer event raised by the Leggi contract.
type LeggiTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Leggi *LeggiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LeggiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Leggi.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LeggiTransferIterator{contract: _Leggi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Leggi *LeggiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LeggiTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Leggi.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeggiTransfer)
				if err := _Leggi.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Leggi *LeggiFilterer) ParseTransfer(log types.Log) (*LeggiTransfer, error) {
	event := new(LeggiTransfer)
	if err := _Leggi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
