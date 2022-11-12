// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package snaili

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

// SnailiStageMintConfig is an auto generated low-level Go binding around an user-defined struct.
type SnailiStageMintConfig struct {
	StageNum              uint64
	MaxPerStage           uint64
	MaxPerAddress         uint64
	IsWhiteListMintActive bool
	MerkleRoot            [32]byte
	IsPublicMintActive    bool
}

// SnailiMetaData contains all meta data concerning the Snaili contract.
var SnailiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalToCurrentOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApproveToCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"stageNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxPerStage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxPerAddress\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isWhiteListMintActive\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isPublicMintActive\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structSnaili.StageMintConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"StageMintConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"quantity\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"airDrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"address_\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"isKYCAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"quantity\",\"type\":\"uint64\"}],\"name\":\"publicMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_kycMerkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setKycMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newMaxPerAddress\",\"type\":\"uint64\"}],\"name\":\"setMaxPerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newMaxPerStage\",\"type\":\"uint64\"}],\"name\":\"setMaxPerStage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"mintStarted\",\"type\":\"bool\"}],\"name\":\"setPublicMintActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"stageNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxPerStage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxPerAddress\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isWhiteListMintActive\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isPublicMintActive\",\"type\":\"bool\"}],\"internalType\":\"structSnaili.StageMintConfig\",\"name\":\"config_\",\"type\":\"tuple\"}],\"name\":\"setStageMintConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"}],\"name\":\"setTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"mintStarted\",\"type\":\"bool\"}],\"name\":\"setWhiteListMintActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stageMintConfig\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"stageNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxPerStage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxPerAddress\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isWhiteListMintActive\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isPublicMintActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"quantity\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"whitelistMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SnailiABI is the input ABI used to generate the binding from.
// Deprecated: Use SnailiMetaData.ABI instead.
var SnailiABI = SnailiMetaData.ABI

// Snaili is an auto generated Go binding around an Ethereum contract.
type Snaili struct {
	SnailiCaller     // Read-only binding to the contract
	SnailiTransactor // Write-only binding to the contract
	SnailiFilterer   // Log filterer for contract events
}

// SnailiCaller is an auto generated read-only Go binding around an Ethereum contract.
type SnailiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnailiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SnailiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnailiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SnailiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnailiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SnailiSession struct {
	Contract     *Snaili           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SnailiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SnailiCallerSession struct {
	Contract *SnailiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SnailiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SnailiTransactorSession struct {
	Contract     *SnailiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SnailiRaw is an auto generated low-level Go binding around an Ethereum contract.
type SnailiRaw struct {
	Contract *Snaili // Generic contract binding to access the raw methods on
}

// SnailiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SnailiCallerRaw struct {
	Contract *SnailiCaller // Generic read-only contract binding to access the raw methods on
}

// SnailiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SnailiTransactorRaw struct {
	Contract *SnailiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSnaili creates a new instance of Snaili, bound to a specific deployed contract.
func NewSnaili(address common.Address, backend bind.ContractBackend) (*Snaili, error) {
	contract, err := bindSnaili(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Snaili{SnailiCaller: SnailiCaller{contract: contract}, SnailiTransactor: SnailiTransactor{contract: contract}, SnailiFilterer: SnailiFilterer{contract: contract}}, nil
}

// NewSnailiCaller creates a new read-only instance of Snaili, bound to a specific deployed contract.
func NewSnailiCaller(address common.Address, caller bind.ContractCaller) (*SnailiCaller, error) {
	contract, err := bindSnaili(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SnailiCaller{contract: contract}, nil
}

// NewSnailiTransactor creates a new write-only instance of Snaili, bound to a specific deployed contract.
func NewSnailiTransactor(address common.Address, transactor bind.ContractTransactor) (*SnailiTransactor, error) {
	contract, err := bindSnaili(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SnailiTransactor{contract: contract}, nil
}

// NewSnailiFilterer creates a new log filterer instance of Snaili, bound to a specific deployed contract.
func NewSnailiFilterer(address common.Address, filterer bind.ContractFilterer) (*SnailiFilterer, error) {
	contract, err := bindSnaili(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SnailiFilterer{contract: contract}, nil
}

// bindSnaili binds a generic wrapper to an already deployed contract.
func bindSnaili(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SnailiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Snaili *SnailiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Snaili.Contract.SnailiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Snaili *SnailiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snaili.Contract.SnailiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Snaili *SnailiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Snaili.Contract.SnailiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Snaili *SnailiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Snaili.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Snaili *SnailiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snaili.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Snaili *SnailiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Snaili.Contract.contract.Transact(opts, method, params...)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Snaili *SnailiCaller) MAXSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "MAX_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Snaili *SnailiSession) MAXSUPPLY() (*big.Int, error) {
	return _Snaili.Contract.MAXSUPPLY(&_Snaili.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Snaili *SnailiCallerSession) MAXSUPPLY() (*big.Int, error) {
	return _Snaili.Contract.MAXSUPPLY(&_Snaili.CallOpts)
}

// AddressMinted is a free data retrieval call binding the contract method 0xf4a66808.
//
// Solidity: function addressMinted(uint256 , address ) view returns(uint256)
func (_Snaili *SnailiCaller) AddressMinted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "addressMinted", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressMinted is a free data retrieval call binding the contract method 0xf4a66808.
//
// Solidity: function addressMinted(uint256 , address ) view returns(uint256)
func (_Snaili *SnailiSession) AddressMinted(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Snaili.Contract.AddressMinted(&_Snaili.CallOpts, arg0, arg1)
}

// AddressMinted is a free data retrieval call binding the contract method 0xf4a66808.
//
// Solidity: function addressMinted(uint256 , address ) view returns(uint256)
func (_Snaili *SnailiCallerSession) AddressMinted(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Snaili.Contract.AddressMinted(&_Snaili.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Snaili *SnailiCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Snaili *SnailiSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Snaili.Contract.BalanceOf(&_Snaili.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Snaili *SnailiCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Snaili.Contract.BalanceOf(&_Snaili.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Snaili *SnailiCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Snaili *SnailiSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Snaili.Contract.GetApproved(&_Snaili.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Snaili *SnailiCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Snaili.Contract.GetApproved(&_Snaili.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Snaili *SnailiCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Snaili *SnailiSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Snaili.Contract.IsApprovedForAll(&_Snaili.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Snaili *SnailiCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Snaili.Contract.IsApprovedForAll(&_Snaili.CallOpts, owner, operator)
}

// IsKYCAddress is a free data retrieval call binding the contract method 0xf102e39e.
//
// Solidity: function isKYCAddress(address address_, bytes32[] merkleProof) view returns(bool)
func (_Snaili *SnailiCaller) IsKYCAddress(opts *bind.CallOpts, address_ common.Address, merkleProof [][32]byte) (bool, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "isKYCAddress", address_, merkleProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKYCAddress is a free data retrieval call binding the contract method 0xf102e39e.
//
// Solidity: function isKYCAddress(address address_, bytes32[] merkleProof) view returns(bool)
func (_Snaili *SnailiSession) IsKYCAddress(address_ common.Address, merkleProof [][32]byte) (bool, error) {
	return _Snaili.Contract.IsKYCAddress(&_Snaili.CallOpts, address_, merkleProof)
}

// IsKYCAddress is a free data retrieval call binding the contract method 0xf102e39e.
//
// Solidity: function isKYCAddress(address address_, bytes32[] merkleProof) view returns(bool)
func (_Snaili *SnailiCallerSession) IsKYCAddress(address_ common.Address, merkleProof [][32]byte) (bool, error) {
	return _Snaili.Contract.IsKYCAddress(&_Snaili.CallOpts, address_, merkleProof)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Snaili *SnailiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Snaili *SnailiSession) Name() (string, error) {
	return _Snaili.Contract.Name(&_Snaili.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Snaili *SnailiCallerSession) Name() (string, error) {
	return _Snaili.Contract.Name(&_Snaili.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Snaili *SnailiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Snaili *SnailiSession) Owner() (common.Address, error) {
	return _Snaili.Contract.Owner(&_Snaili.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Snaili *SnailiCallerSession) Owner() (common.Address, error) {
	return _Snaili.Contract.Owner(&_Snaili.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Snaili *SnailiCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Snaili *SnailiSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Snaili.Contract.OwnerOf(&_Snaili.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Snaili *SnailiCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Snaili.Contract.OwnerOf(&_Snaili.CallOpts, tokenId)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Snaili *SnailiCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Snaili *SnailiSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Snaili.Contract.RoyaltyInfo(&_Snaili.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Snaili *SnailiCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Snaili.Contract.RoyaltyInfo(&_Snaili.CallOpts, tokenId, salePrice)
}

// StageMintConfig is a free data retrieval call binding the contract method 0x3173ea1a.
//
// Solidity: function stageMintConfig() view returns(uint64 stageNum, uint64 maxPerStage, uint64 maxPerAddress, bool isWhiteListMintActive, bytes32 merkleRoot, bool isPublicMintActive)
func (_Snaili *SnailiCaller) StageMintConfig(opts *bind.CallOpts) (struct {
	StageNum              uint64
	MaxPerStage           uint64
	MaxPerAddress         uint64
	IsWhiteListMintActive bool
	MerkleRoot            [32]byte
	IsPublicMintActive    bool
}, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "stageMintConfig")

	outstruct := new(struct {
		StageNum              uint64
		MaxPerStage           uint64
		MaxPerAddress         uint64
		IsWhiteListMintActive bool
		MerkleRoot            [32]byte
		IsPublicMintActive    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StageNum = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.MaxPerStage = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.MaxPerAddress = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.IsWhiteListMintActive = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.MerkleRoot = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.IsPublicMintActive = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// StageMintConfig is a free data retrieval call binding the contract method 0x3173ea1a.
//
// Solidity: function stageMintConfig() view returns(uint64 stageNum, uint64 maxPerStage, uint64 maxPerAddress, bool isWhiteListMintActive, bytes32 merkleRoot, bool isPublicMintActive)
func (_Snaili *SnailiSession) StageMintConfig() (struct {
	StageNum              uint64
	MaxPerStage           uint64
	MaxPerAddress         uint64
	IsWhiteListMintActive bool
	MerkleRoot            [32]byte
	IsPublicMintActive    bool
}, error) {
	return _Snaili.Contract.StageMintConfig(&_Snaili.CallOpts)
}

// StageMintConfig is a free data retrieval call binding the contract method 0x3173ea1a.
//
// Solidity: function stageMintConfig() view returns(uint64 stageNum, uint64 maxPerStage, uint64 maxPerAddress, bool isWhiteListMintActive, bytes32 merkleRoot, bool isPublicMintActive)
func (_Snaili *SnailiCallerSession) StageMintConfig() (struct {
	StageNum              uint64
	MaxPerStage           uint64
	MaxPerAddress         uint64
	IsWhiteListMintActive bool
	MerkleRoot            [32]byte
	IsPublicMintActive    bool
}, error) {
	return _Snaili.Contract.StageMintConfig(&_Snaili.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Snaili *SnailiCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Snaili *SnailiSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Snaili.Contract.SupportsInterface(&_Snaili.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Snaili *SnailiCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Snaili.Contract.SupportsInterface(&_Snaili.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Snaili *SnailiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Snaili *SnailiSession) Symbol() (string, error) {
	return _Snaili.Contract.Symbol(&_Snaili.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Snaili *SnailiCallerSession) Symbol() (string, error) {
	return _Snaili.Contract.Symbol(&_Snaili.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Snaili *SnailiCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Snaili *SnailiSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Snaili.Contract.TokenURI(&_Snaili.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Snaili *SnailiCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Snaili.Contract.TokenURI(&_Snaili.CallOpts, tokenId)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Snaili *SnailiCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Snaili *SnailiSession) TotalMinted() (*big.Int, error) {
	return _Snaili.Contract.TotalMinted(&_Snaili.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Snaili *SnailiCallerSession) TotalMinted() (*big.Int, error) {
	return _Snaili.Contract.TotalMinted(&_Snaili.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Snaili *SnailiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Snaili.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Snaili *SnailiSession) TotalSupply() (*big.Int, error) {
	return _Snaili.Contract.TotalSupply(&_Snaili.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Snaili *SnailiCallerSession) TotalSupply() (*big.Int, error) {
	return _Snaili.Contract.TotalSupply(&_Snaili.CallOpts)
}

// AirDrop is a paid mutator transaction binding the contract method 0xb2c148b2.
//
// Solidity: function airDrop(uint64 quantity, address to) returns()
func (_Snaili *SnailiTransactor) AirDrop(opts *bind.TransactOpts, quantity uint64, to common.Address) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "airDrop", quantity, to)
}

// AirDrop is a paid mutator transaction binding the contract method 0xb2c148b2.
//
// Solidity: function airDrop(uint64 quantity, address to) returns()
func (_Snaili *SnailiSession) AirDrop(quantity uint64, to common.Address) (*types.Transaction, error) {
	return _Snaili.Contract.AirDrop(&_Snaili.TransactOpts, quantity, to)
}

// AirDrop is a paid mutator transaction binding the contract method 0xb2c148b2.
//
// Solidity: function airDrop(uint64 quantity, address to) returns()
func (_Snaili *SnailiTransactorSession) AirDrop(quantity uint64, to common.Address) (*types.Transaction, error) {
	return _Snaili.Contract.AirDrop(&_Snaili.TransactOpts, quantity, to)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Snaili *SnailiTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Snaili *SnailiSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Snaili.Contract.Approve(&_Snaili.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Snaili *SnailiTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Snaili.Contract.Approve(&_Snaili.TransactOpts, to, tokenId)
}

// PublicMint is a paid mutator transaction binding the contract method 0x6afcb7b0.
//
// Solidity: function publicMint(uint64 quantity) returns()
func (_Snaili *SnailiTransactor) PublicMint(opts *bind.TransactOpts, quantity uint64) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "publicMint", quantity)
}

// PublicMint is a paid mutator transaction binding the contract method 0x6afcb7b0.
//
// Solidity: function publicMint(uint64 quantity) returns()
func (_Snaili *SnailiSession) PublicMint(quantity uint64) (*types.Transaction, error) {
	return _Snaili.Contract.PublicMint(&_Snaili.TransactOpts, quantity)
}

// PublicMint is a paid mutator transaction binding the contract method 0x6afcb7b0.
//
// Solidity: function publicMint(uint64 quantity) returns()
func (_Snaili *SnailiTransactorSession) PublicMint(quantity uint64) (*types.Transaction, error) {
	return _Snaili.Contract.PublicMint(&_Snaili.TransactOpts, quantity)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Snaili *SnailiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Snaili *SnailiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Snaili.Contract.RenounceOwnership(&_Snaili.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Snaili *SnailiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Snaili.Contract.RenounceOwnership(&_Snaili.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Snaili *SnailiTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Snaili *SnailiSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Snaili.Contract.SafeTransferFrom(&_Snaili.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Snaili *SnailiTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Snaili.Contract.SafeTransferFrom(&_Snaili.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Snaili *SnailiTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Snaili *SnailiSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Snaili.Contract.SafeTransferFrom0(&_Snaili.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Snaili *SnailiTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Snaili.Contract.SafeTransferFrom0(&_Snaili.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Snaili *SnailiTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Snaili *SnailiSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Snaili.Contract.SetApprovalForAll(&_Snaili.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Snaili *SnailiTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Snaili.Contract.SetApprovalForAll(&_Snaili.TransactOpts, operator, approved)
}

// SetKycMerkleRoot is a paid mutator transaction binding the contract method 0x786867b5.
//
// Solidity: function setKycMerkleRoot(bytes32 _kycMerkleRoot) returns()
func (_Snaili *SnailiTransactor) SetKycMerkleRoot(opts *bind.TransactOpts, _kycMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "setKycMerkleRoot", _kycMerkleRoot)
}

// SetKycMerkleRoot is a paid mutator transaction binding the contract method 0x786867b5.
//
// Solidity: function setKycMerkleRoot(bytes32 _kycMerkleRoot) returns()
func (_Snaili *SnailiSession) SetKycMerkleRoot(_kycMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Snaili.Contract.SetKycMerkleRoot(&_Snaili.TransactOpts, _kycMerkleRoot)
}

// SetKycMerkleRoot is a paid mutator transaction binding the contract method 0x786867b5.
//
// Solidity: function setKycMerkleRoot(bytes32 _kycMerkleRoot) returns()
func (_Snaili *SnailiTransactorSession) SetKycMerkleRoot(_kycMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Snaili.Contract.SetKycMerkleRoot(&_Snaili.TransactOpts, _kycMerkleRoot)
}

// SetMaxPerAddress is a paid mutator transaction binding the contract method 0x430a06fa.
//
// Solidity: function setMaxPerAddress(uint64 newMaxPerAddress) returns()
func (_Snaili *SnailiTransactor) SetMaxPerAddress(opts *bind.TransactOpts, newMaxPerAddress uint64) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "setMaxPerAddress", newMaxPerAddress)
}

// SetMaxPerAddress is a paid mutator transaction binding the contract method 0x430a06fa.
//
// Solidity: function setMaxPerAddress(uint64 newMaxPerAddress) returns()
func (_Snaili *SnailiSession) SetMaxPerAddress(newMaxPerAddress uint64) (*types.Transaction, error) {
	return _Snaili.Contract.SetMaxPerAddress(&_Snaili.TransactOpts, newMaxPerAddress)
}

// SetMaxPerAddress is a paid mutator transaction binding the contract method 0x430a06fa.
//
// Solidity: function setMaxPerAddress(uint64 newMaxPerAddress) returns()
func (_Snaili *SnailiTransactorSession) SetMaxPerAddress(newMaxPerAddress uint64) (*types.Transaction, error) {
	return _Snaili.Contract.SetMaxPerAddress(&_Snaili.TransactOpts, newMaxPerAddress)
}

// SetMaxPerStage is a paid mutator transaction binding the contract method 0x630a3bc1.
//
// Solidity: function setMaxPerStage(uint64 newMaxPerStage) returns()
func (_Snaili *SnailiTransactor) SetMaxPerStage(opts *bind.TransactOpts, newMaxPerStage uint64) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "setMaxPerStage", newMaxPerStage)
}

// SetMaxPerStage is a paid mutator transaction binding the contract method 0x630a3bc1.
//
// Solidity: function setMaxPerStage(uint64 newMaxPerStage) returns()
func (_Snaili *SnailiSession) SetMaxPerStage(newMaxPerStage uint64) (*types.Transaction, error) {
	return _Snaili.Contract.SetMaxPerStage(&_Snaili.TransactOpts, newMaxPerStage)
}

// SetMaxPerStage is a paid mutator transaction binding the contract method 0x630a3bc1.
//
// Solidity: function setMaxPerStage(uint64 newMaxPerStage) returns()
func (_Snaili *SnailiTransactorSession) SetMaxPerStage(newMaxPerStage uint64) (*types.Transaction, error) {
	return _Snaili.Contract.SetMaxPerStage(&_Snaili.TransactOpts, newMaxPerStage)
}

// SetPublicMintActive is a paid mutator transaction binding the contract method 0x2b707c71.
//
// Solidity: function setPublicMintActive(bool mintStarted) returns()
func (_Snaili *SnailiTransactor) SetPublicMintActive(opts *bind.TransactOpts, mintStarted bool) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "setPublicMintActive", mintStarted)
}

// SetPublicMintActive is a paid mutator transaction binding the contract method 0x2b707c71.
//
// Solidity: function setPublicMintActive(bool mintStarted) returns()
func (_Snaili *SnailiSession) SetPublicMintActive(mintStarted bool) (*types.Transaction, error) {
	return _Snaili.Contract.SetPublicMintActive(&_Snaili.TransactOpts, mintStarted)
}

// SetPublicMintActive is a paid mutator transaction binding the contract method 0x2b707c71.
//
// Solidity: function setPublicMintActive(bool mintStarted) returns()
func (_Snaili *SnailiTransactorSession) SetPublicMintActive(mintStarted bool) (*types.Transaction, error) {
	return _Snaili.Contract.SetPublicMintActive(&_Snaili.TransactOpts, mintStarted)
}

// SetStageMintConfig is a paid mutator transaction binding the contract method 0x5b894742.
//
// Solidity: function setStageMintConfig((uint64,uint64,uint64,bool,bytes32,bool) config_) returns()
func (_Snaili *SnailiTransactor) SetStageMintConfig(opts *bind.TransactOpts, config_ SnailiStageMintConfig) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "setStageMintConfig", config_)
}

// SetStageMintConfig is a paid mutator transaction binding the contract method 0x5b894742.
//
// Solidity: function setStageMintConfig((uint64,uint64,uint64,bool,bytes32,bool) config_) returns()
func (_Snaili *SnailiSession) SetStageMintConfig(config_ SnailiStageMintConfig) (*types.Transaction, error) {
	return _Snaili.Contract.SetStageMintConfig(&_Snaili.TransactOpts, config_)
}

// SetStageMintConfig is a paid mutator transaction binding the contract method 0x5b894742.
//
// Solidity: function setStageMintConfig((uint64,uint64,uint64,bool,bytes32,bool) config_) returns()
func (_Snaili *SnailiTransactorSession) SetStageMintConfig(config_ SnailiStageMintConfig) (*types.Transaction, error) {
	return _Snaili.Contract.SetStageMintConfig(&_Snaili.TransactOpts, config_)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0xe0df5b6f.
//
// Solidity: function setTokenURI(string tokenURI_) returns()
func (_Snaili *SnailiTransactor) SetTokenURI(opts *bind.TransactOpts, tokenURI_ string) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "setTokenURI", tokenURI_)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0xe0df5b6f.
//
// Solidity: function setTokenURI(string tokenURI_) returns()
func (_Snaili *SnailiSession) SetTokenURI(tokenURI_ string) (*types.Transaction, error) {
	return _Snaili.Contract.SetTokenURI(&_Snaili.TransactOpts, tokenURI_)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0xe0df5b6f.
//
// Solidity: function setTokenURI(string tokenURI_) returns()
func (_Snaili *SnailiTransactorSession) SetTokenURI(tokenURI_ string) (*types.Transaction, error) {
	return _Snaili.Contract.SetTokenURI(&_Snaili.TransactOpts, tokenURI_)
}

// SetWhiteListMintActive is a paid mutator transaction binding the contract method 0xa4f0ed50.
//
// Solidity: function setWhiteListMintActive(bool mintStarted) returns()
func (_Snaili *SnailiTransactor) SetWhiteListMintActive(opts *bind.TransactOpts, mintStarted bool) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "setWhiteListMintActive", mintStarted)
}

// SetWhiteListMintActive is a paid mutator transaction binding the contract method 0xa4f0ed50.
//
// Solidity: function setWhiteListMintActive(bool mintStarted) returns()
func (_Snaili *SnailiSession) SetWhiteListMintActive(mintStarted bool) (*types.Transaction, error) {
	return _Snaili.Contract.SetWhiteListMintActive(&_Snaili.TransactOpts, mintStarted)
}

// SetWhiteListMintActive is a paid mutator transaction binding the contract method 0xa4f0ed50.
//
// Solidity: function setWhiteListMintActive(bool mintStarted) returns()
func (_Snaili *SnailiTransactorSession) SetWhiteListMintActive(mintStarted bool) (*types.Transaction, error) {
	return _Snaili.Contract.SetWhiteListMintActive(&_Snaili.TransactOpts, mintStarted)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Snaili *SnailiTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Snaili *SnailiSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Snaili.Contract.TransferFrom(&_Snaili.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Snaili *SnailiTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Snaili.Contract.TransferFrom(&_Snaili.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Snaili *SnailiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Snaili *SnailiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Snaili.Contract.TransferOwnership(&_Snaili.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Snaili *SnailiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Snaili.Contract.TransferOwnership(&_Snaili.TransactOpts, newOwner)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0x81e1ced3.
//
// Solidity: function whitelistMint(uint64 quantity, bytes32[] merkleProof) returns()
func (_Snaili *SnailiTransactor) WhitelistMint(opts *bind.TransactOpts, quantity uint64, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "whitelistMint", quantity, merkleProof)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0x81e1ced3.
//
// Solidity: function whitelistMint(uint64 quantity, bytes32[] merkleProof) returns()
func (_Snaili *SnailiSession) WhitelistMint(quantity uint64, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Snaili.Contract.WhitelistMint(&_Snaili.TransactOpts, quantity, merkleProof)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0x81e1ced3.
//
// Solidity: function whitelistMint(uint64 quantity, bytes32[] merkleProof) returns()
func (_Snaili *SnailiTransactorSession) WhitelistMint(quantity uint64, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Snaili.Contract.WhitelistMint(&_Snaili.TransactOpts, quantity, merkleProof)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Snaili *SnailiTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Snaili *SnailiSession) Withdraw() (*types.Transaction, error) {
	return _Snaili.Contract.Withdraw(&_Snaili.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Snaili *SnailiTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Snaili.Contract.Withdraw(&_Snaili.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address token) returns()
func (_Snaili *SnailiTransactor) WithdrawTokens(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Snaili.contract.Transact(opts, "withdrawTokens", token)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address token) returns()
func (_Snaili *SnailiSession) WithdrawTokens(token common.Address) (*types.Transaction, error) {
	return _Snaili.Contract.WithdrawTokens(&_Snaili.TransactOpts, token)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x49df728c.
//
// Solidity: function withdrawTokens(address token) returns()
func (_Snaili *SnailiTransactorSession) WithdrawTokens(token common.Address) (*types.Transaction, error) {
	return _Snaili.Contract.WithdrawTokens(&_Snaili.TransactOpts, token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Snaili *SnailiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snaili.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Snaili *SnailiSession) Receive() (*types.Transaction, error) {
	return _Snaili.Contract.Receive(&_Snaili.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Snaili *SnailiTransactorSession) Receive() (*types.Transaction, error) {
	return _Snaili.Contract.Receive(&_Snaili.TransactOpts)
}

// SnailiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Snaili contract.
type SnailiApprovalIterator struct {
	Event *SnailiApproval // Event containing the contract specifics and raw log

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
func (it *SnailiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnailiApproval)
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
		it.Event = new(SnailiApproval)
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
func (it *SnailiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnailiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnailiApproval represents a Approval event raised by the Snaili contract.
type SnailiApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Snaili *SnailiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SnailiApprovalIterator, error) {

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

	logs, sub, err := _Snaili.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SnailiApprovalIterator{contract: _Snaili.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Snaili *SnailiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SnailiApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Snaili.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnailiApproval)
				if err := _Snaili.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Snaili *SnailiFilterer) ParseApproval(log types.Log) (*SnailiApproval, error) {
	event := new(SnailiApproval)
	if err := _Snaili.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SnailiApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Snaili contract.
type SnailiApprovalForAllIterator struct {
	Event *SnailiApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SnailiApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnailiApprovalForAll)
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
		it.Event = new(SnailiApprovalForAll)
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
func (it *SnailiApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnailiApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnailiApprovalForAll represents a ApprovalForAll event raised by the Snaili contract.
type SnailiApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Snaili *SnailiFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SnailiApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Snaili.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SnailiApprovalForAllIterator{contract: _Snaili.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Snaili *SnailiFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SnailiApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Snaili.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnailiApprovalForAll)
				if err := _Snaili.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Snaili *SnailiFilterer) ParseApprovalForAll(log types.Log) (*SnailiApprovalForAll, error) {
	event := new(SnailiApprovalForAll)
	if err := _Snaili.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SnailiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Snaili contract.
type SnailiOwnershipTransferredIterator struct {
	Event *SnailiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SnailiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnailiOwnershipTransferred)
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
		it.Event = new(SnailiOwnershipTransferred)
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
func (it *SnailiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnailiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnailiOwnershipTransferred represents a OwnershipTransferred event raised by the Snaili contract.
type SnailiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Snaili *SnailiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SnailiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Snaili.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SnailiOwnershipTransferredIterator{contract: _Snaili.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Snaili *SnailiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SnailiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Snaili.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnailiOwnershipTransferred)
				if err := _Snaili.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Snaili *SnailiFilterer) ParseOwnershipTransferred(log types.Log) (*SnailiOwnershipTransferred, error) {
	event := new(SnailiOwnershipTransferred)
	if err := _Snaili.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SnailiReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Snaili contract.
type SnailiReceivedIterator struct {
	Event *SnailiReceived // Event containing the contract specifics and raw log

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
func (it *SnailiReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnailiReceived)
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
		it.Event = new(SnailiReceived)
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
func (it *SnailiReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnailiReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnailiReceived represents a Received event raised by the Snaili contract.
type SnailiReceived struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed arg0, uint256 arg1)
func (_Snaili *SnailiFilterer) FilterReceived(opts *bind.FilterOpts, arg0 []common.Address) (*SnailiReceivedIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Snaili.contract.FilterLogs(opts, "Received", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &SnailiReceivedIterator{contract: _Snaili.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address indexed arg0, uint256 arg1)
func (_Snaili *SnailiFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *SnailiReceived, arg0 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Snaili.contract.WatchLogs(opts, "Received", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnailiReceived)
				if err := _Snaili.contract.UnpackLog(event, "Received", log); err != nil {
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
func (_Snaili *SnailiFilterer) ParseReceived(log types.Log) (*SnailiReceived, error) {
	event := new(SnailiReceived)
	if err := _Snaili.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SnailiStageMintConfigChangedIterator is returned from FilterStageMintConfigChanged and is used to iterate over the raw logs and unpacked data for StageMintConfigChanged events raised by the Snaili contract.
type SnailiStageMintConfigChangedIterator struct {
	Event *SnailiStageMintConfigChanged // Event containing the contract specifics and raw log

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
func (it *SnailiStageMintConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnailiStageMintConfigChanged)
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
		it.Event = new(SnailiStageMintConfigChanged)
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
func (it *SnailiStageMintConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnailiStageMintConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnailiStageMintConfigChanged represents a StageMintConfigChanged event raised by the Snaili contract.
type SnailiStageMintConfigChanged struct {
	Config SnailiStageMintConfig
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStageMintConfigChanged is a free log retrieval operation binding the contract event 0xede207c19078509567746f080e3d6ee42074d36324ccd94330364bd907f73b81.
//
// Solidity: event StageMintConfigChanged((uint64,uint64,uint64,bool,bytes32,bool) config)
func (_Snaili *SnailiFilterer) FilterStageMintConfigChanged(opts *bind.FilterOpts) (*SnailiStageMintConfigChangedIterator, error) {

	logs, sub, err := _Snaili.contract.FilterLogs(opts, "StageMintConfigChanged")
	if err != nil {
		return nil, err
	}
	return &SnailiStageMintConfigChangedIterator{contract: _Snaili.contract, event: "StageMintConfigChanged", logs: logs, sub: sub}, nil
}

// WatchStageMintConfigChanged is a free log subscription operation binding the contract event 0xede207c19078509567746f080e3d6ee42074d36324ccd94330364bd907f73b81.
//
// Solidity: event StageMintConfigChanged((uint64,uint64,uint64,bool,bytes32,bool) config)
func (_Snaili *SnailiFilterer) WatchStageMintConfigChanged(opts *bind.WatchOpts, sink chan<- *SnailiStageMintConfigChanged) (event.Subscription, error) {

	logs, sub, err := _Snaili.contract.WatchLogs(opts, "StageMintConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnailiStageMintConfigChanged)
				if err := _Snaili.contract.UnpackLog(event, "StageMintConfigChanged", log); err != nil {
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

// ParseStageMintConfigChanged is a log parse operation binding the contract event 0xede207c19078509567746f080e3d6ee42074d36324ccd94330364bd907f73b81.
//
// Solidity: event StageMintConfigChanged((uint64,uint64,uint64,bool,bytes32,bool) config)
func (_Snaili *SnailiFilterer) ParseStageMintConfigChanged(log types.Log) (*SnailiStageMintConfigChanged, error) {
	event := new(SnailiStageMintConfigChanged)
	if err := _Snaili.contract.UnpackLog(event, "StageMintConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SnailiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Snaili contract.
type SnailiTransferIterator struct {
	Event *SnailiTransfer // Event containing the contract specifics and raw log

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
func (it *SnailiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnailiTransfer)
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
		it.Event = new(SnailiTransfer)
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
func (it *SnailiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnailiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnailiTransfer represents a Transfer event raised by the Snaili contract.
type SnailiTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Snaili *SnailiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SnailiTransferIterator, error) {

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

	logs, sub, err := _Snaili.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SnailiTransferIterator{contract: _Snaili.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Snaili *SnailiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SnailiTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Snaili.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnailiTransfer)
				if err := _Snaili.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Snaili *SnailiFilterer) ParseTransfer(log types.Log) (*SnailiTransfer, error) {
	event := new(SnailiTransfer)
	if err := _Snaili.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
