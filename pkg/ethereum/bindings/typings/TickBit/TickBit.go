// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TickBit

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

// TickBitBlockHeader is an auto generated low-level Go binding around an user-defined struct.
type TickBitBlockHeader struct {
	MerkleRootHash    [32]byte
	NBits             [4]byte
	Nonce             [4]byte
	PreviousBlockHash [32]byte
	Timestamp         [4]byte
	Version           [4]byte
}

// TickBitMetaData contains all meta data concerning the TickBit contract.
var TickBitMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_verifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_Px\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_Py\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_tickSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"Px\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"Py\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bet\",\"inputs\":[{\"name\":\"timestamps\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blockBets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"bettedAt\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"bettor\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blockTimestamps\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"convertToBigEndian\",\"inputs\":[{\"name\":\"bytesLE\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"convertToBytes32\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"doubleHash\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"latestBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"parseBlockHeader\",\"inputs\":[{\"name\":\"blockHeader\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"parsedHeader\",\"type\":\"tuple\",\"internalType\":\"structTickBit.BlockHeader\",\"components\":[{\"name\":\"merkleRootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nBits\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"nonce\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"previousBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"pools\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"acruedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settledAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tickSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifiedBlocks\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"merkleRootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nBits\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"nonce\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"previousBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractSXGVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyAndSettleBlock\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHeader\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyBlock\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHeader\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structTickBit.BlockHeader\",\"components\":[{\"name\":\"merkleRootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nBits\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"nonce\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"previousBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BetPlaced\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"timestamps\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// TickBitABI is the input ABI used to generate the binding from.
// Deprecated: Use TickBitMetaData.ABI instead.
var TickBitABI = TickBitMetaData.ABI

// TickBit is an auto generated Go binding around an Ethereum contract.
type TickBit struct {
	TickBitCaller     // Read-only binding to the contract
	TickBitTransactor // Write-only binding to the contract
	TickBitFilterer   // Log filterer for contract events
}

// TickBitCaller is an auto generated read-only Go binding around an Ethereum contract.
type TickBitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickBitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TickBitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickBitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TickBitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TickBitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TickBitSession struct {
	Contract     *TickBit          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TickBitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TickBitCallerSession struct {
	Contract *TickBitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TickBitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TickBitTransactorSession struct {
	Contract     *TickBitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TickBitRaw is an auto generated low-level Go binding around an Ethereum contract.
type TickBitRaw struct {
	Contract *TickBit // Generic contract binding to access the raw methods on
}

// TickBitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TickBitCallerRaw struct {
	Contract *TickBitCaller // Generic read-only contract binding to access the raw methods on
}

// TickBitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TickBitTransactorRaw struct {
	Contract *TickBitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTickBit creates a new instance of TickBit, bound to a specific deployed contract.
func NewTickBit(address common.Address, backend bind.ContractBackend) (*TickBit, error) {
	contract, err := bindTickBit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TickBit{TickBitCaller: TickBitCaller{contract: contract}, TickBitTransactor: TickBitTransactor{contract: contract}, TickBitFilterer: TickBitFilterer{contract: contract}}, nil
}

// NewTickBitCaller creates a new read-only instance of TickBit, bound to a specific deployed contract.
func NewTickBitCaller(address common.Address, caller bind.ContractCaller) (*TickBitCaller, error) {
	contract, err := bindTickBit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TickBitCaller{contract: contract}, nil
}

// NewTickBitTransactor creates a new write-only instance of TickBit, bound to a specific deployed contract.
func NewTickBitTransactor(address common.Address, transactor bind.ContractTransactor) (*TickBitTransactor, error) {
	contract, err := bindTickBit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TickBitTransactor{contract: contract}, nil
}

// NewTickBitFilterer creates a new log filterer instance of TickBit, bound to a specific deployed contract.
func NewTickBitFilterer(address common.Address, filterer bind.ContractFilterer) (*TickBitFilterer, error) {
	contract, err := bindTickBit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TickBitFilterer{contract: contract}, nil
}

// bindTickBit binds a generic wrapper to an already deployed contract.
func bindTickBit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TickBitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TickBit *TickBitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TickBit.Contract.TickBitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TickBit *TickBitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TickBit.Contract.TickBitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TickBit *TickBitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TickBit.Contract.TickBitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TickBit *TickBitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TickBit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TickBit *TickBitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TickBit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TickBit *TickBitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TickBit.Contract.contract.Transact(opts, method, params...)
}

// Px is a free data retrieval call binding the contract method 0x13a22b77.
//
// Solidity: function Px() view returns(bytes32)
func (_TickBit *TickBitCaller) Px(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "Px")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Px is a free data retrieval call binding the contract method 0x13a22b77.
//
// Solidity: function Px() view returns(bytes32)
func (_TickBit *TickBitSession) Px() ([32]byte, error) {
	return _TickBit.Contract.Px(&_TickBit.CallOpts)
}

// Px is a free data retrieval call binding the contract method 0x13a22b77.
//
// Solidity: function Px() view returns(bytes32)
func (_TickBit *TickBitCallerSession) Px() ([32]byte, error) {
	return _TickBit.Contract.Px(&_TickBit.CallOpts)
}

// Py is a free data retrieval call binding the contract method 0x99806ec1.
//
// Solidity: function Py() view returns(bytes32)
func (_TickBit *TickBitCaller) Py(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "Py")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Py is a free data retrieval call binding the contract method 0x99806ec1.
//
// Solidity: function Py() view returns(bytes32)
func (_TickBit *TickBitSession) Py() ([32]byte, error) {
	return _TickBit.Contract.Py(&_TickBit.CallOpts)
}

// Py is a free data retrieval call binding the contract method 0x99806ec1.
//
// Solidity: function Py() view returns(bytes32)
func (_TickBit *TickBitCallerSession) Py() ([32]byte, error) {
	return _TickBit.Contract.Py(&_TickBit.CallOpts)
}

// BlockBets is a free data retrieval call binding the contract method 0xa9ead5c7.
//
// Solidity: function blockBets(uint256 , uint256 , uint256 ) view returns(uint96 bettedAt, address bettor)
func (_TickBit *TickBitCaller) BlockBets(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	BettedAt *big.Int
	Bettor   common.Address
}, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "blockBets", arg0, arg1, arg2)

	outstruct := new(struct {
		BettedAt *big.Int
		Bettor   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BettedAt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Bettor = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// BlockBets is a free data retrieval call binding the contract method 0xa9ead5c7.
//
// Solidity: function blockBets(uint256 , uint256 , uint256 ) view returns(uint96 bettedAt, address bettor)
func (_TickBit *TickBitSession) BlockBets(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	BettedAt *big.Int
	Bettor   common.Address
}, error) {
	return _TickBit.Contract.BlockBets(&_TickBit.CallOpts, arg0, arg1, arg2)
}

// BlockBets is a free data retrieval call binding the contract method 0xa9ead5c7.
//
// Solidity: function blockBets(uint256 , uint256 , uint256 ) view returns(uint96 bettedAt, address bettor)
func (_TickBit *TickBitCallerSession) BlockBets(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	BettedAt *big.Int
	Bettor   common.Address
}, error) {
	return _TickBit.Contract.BlockBets(&_TickBit.CallOpts, arg0, arg1, arg2)
}

// BlockTimestamps is a free data retrieval call binding the contract method 0x976903d2.
//
// Solidity: function blockTimestamps(uint256 , uint256 ) view returns(uint256)
func (_TickBit *TickBitCaller) BlockTimestamps(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "blockTimestamps", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockTimestamps is a free data retrieval call binding the contract method 0x976903d2.
//
// Solidity: function blockTimestamps(uint256 , uint256 ) view returns(uint256)
func (_TickBit *TickBitSession) BlockTimestamps(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TickBit.Contract.BlockTimestamps(&_TickBit.CallOpts, arg0, arg1)
}

// BlockTimestamps is a free data retrieval call binding the contract method 0x976903d2.
//
// Solidity: function blockTimestamps(uint256 , uint256 ) view returns(uint256)
func (_TickBit *TickBitCallerSession) BlockTimestamps(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TickBit.Contract.BlockTimestamps(&_TickBit.CallOpts, arg0, arg1)
}

// ConvertToBigEndian is a free data retrieval call binding the contract method 0xb6c6d7b4.
//
// Solidity: function convertToBigEndian(bytes bytesLE) pure returns(bytes)
func (_TickBit *TickBitCaller) ConvertToBigEndian(opts *bind.CallOpts, bytesLE []byte) ([]byte, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "convertToBigEndian", bytesLE)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ConvertToBigEndian is a free data retrieval call binding the contract method 0xb6c6d7b4.
//
// Solidity: function convertToBigEndian(bytes bytesLE) pure returns(bytes)
func (_TickBit *TickBitSession) ConvertToBigEndian(bytesLE []byte) ([]byte, error) {
	return _TickBit.Contract.ConvertToBigEndian(&_TickBit.CallOpts, bytesLE)
}

// ConvertToBigEndian is a free data retrieval call binding the contract method 0xb6c6d7b4.
//
// Solidity: function convertToBigEndian(bytes bytesLE) pure returns(bytes)
func (_TickBit *TickBitCallerSession) ConvertToBigEndian(bytesLE []byte) ([]byte, error) {
	return _TickBit.Contract.ConvertToBigEndian(&_TickBit.CallOpts, bytesLE)
}

// ConvertToBytes32 is a free data retrieval call binding the contract method 0x721f75ca.
//
// Solidity: function convertToBytes32(bytes data) pure returns(bytes32)
func (_TickBit *TickBitCaller) ConvertToBytes32(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "convertToBytes32", data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConvertToBytes32 is a free data retrieval call binding the contract method 0x721f75ca.
//
// Solidity: function convertToBytes32(bytes data) pure returns(bytes32)
func (_TickBit *TickBitSession) ConvertToBytes32(data []byte) ([32]byte, error) {
	return _TickBit.Contract.ConvertToBytes32(&_TickBit.CallOpts, data)
}

// ConvertToBytes32 is a free data retrieval call binding the contract method 0x721f75ca.
//
// Solidity: function convertToBytes32(bytes data) pure returns(bytes32)
func (_TickBit *TickBitCallerSession) ConvertToBytes32(data []byte) ([32]byte, error) {
	return _TickBit.Contract.ConvertToBytes32(&_TickBit.CallOpts, data)
}

// DoubleHash is a free data retrieval call binding the contract method 0x10e635a5.
//
// Solidity: function doubleHash(bytes data) pure returns(bytes32)
func (_TickBit *TickBitCaller) DoubleHash(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "doubleHash", data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DoubleHash is a free data retrieval call binding the contract method 0x10e635a5.
//
// Solidity: function doubleHash(bytes data) pure returns(bytes32)
func (_TickBit *TickBitSession) DoubleHash(data []byte) ([32]byte, error) {
	return _TickBit.Contract.DoubleHash(&_TickBit.CallOpts, data)
}

// DoubleHash is a free data retrieval call binding the contract method 0x10e635a5.
//
// Solidity: function doubleHash(bytes data) pure returns(bytes32)
func (_TickBit *TickBitCallerSession) DoubleHash(data []byte) ([32]byte, error) {
	return _TickBit.Contract.DoubleHash(&_TickBit.CallOpts, data)
}

// LatestBlock is a free data retrieval call binding the contract method 0x07e2da96.
//
// Solidity: function latestBlock() view returns(uint256)
func (_TickBit *TickBitCaller) LatestBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "latestBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlock is a free data retrieval call binding the contract method 0x07e2da96.
//
// Solidity: function latestBlock() view returns(uint256)
func (_TickBit *TickBitSession) LatestBlock() (*big.Int, error) {
	return _TickBit.Contract.LatestBlock(&_TickBit.CallOpts)
}

// LatestBlock is a free data retrieval call binding the contract method 0x07e2da96.
//
// Solidity: function latestBlock() view returns(uint256)
func (_TickBit *TickBitCallerSession) LatestBlock() (*big.Int, error) {
	return _TickBit.Contract.LatestBlock(&_TickBit.CallOpts)
}

// ParseBlockHeader is a free data retrieval call binding the contract method 0x849d926b.
//
// Solidity: function parseBlockHeader(bytes blockHeader) pure returns((bytes32,bytes4,bytes4,bytes32,bytes4,bytes4) parsedHeader)
func (_TickBit *TickBitCaller) ParseBlockHeader(opts *bind.CallOpts, blockHeader []byte) (TickBitBlockHeader, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "parseBlockHeader", blockHeader)

	if err != nil {
		return *new(TickBitBlockHeader), err
	}

	out0 := *abi.ConvertType(out[0], new(TickBitBlockHeader)).(*TickBitBlockHeader)

	return out0, err

}

// ParseBlockHeader is a free data retrieval call binding the contract method 0x849d926b.
//
// Solidity: function parseBlockHeader(bytes blockHeader) pure returns((bytes32,bytes4,bytes4,bytes32,bytes4,bytes4) parsedHeader)
func (_TickBit *TickBitSession) ParseBlockHeader(blockHeader []byte) (TickBitBlockHeader, error) {
	return _TickBit.Contract.ParseBlockHeader(&_TickBit.CallOpts, blockHeader)
}

// ParseBlockHeader is a free data retrieval call binding the contract method 0x849d926b.
//
// Solidity: function parseBlockHeader(bytes blockHeader) pure returns((bytes32,bytes4,bytes4,bytes32,bytes4,bytes4) parsedHeader)
func (_TickBit *TickBitCallerSession) ParseBlockHeader(blockHeader []byte) (TickBitBlockHeader, error) {
	return _TickBit.Contract.ParseBlockHeader(&_TickBit.CallOpts, blockHeader)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 acruedAmount, uint256 settledAt)
func (_TickBit *TickBitCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (struct {
	AcruedAmount *big.Int
	SettledAt    *big.Int
}, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "pools", arg0)

	outstruct := new(struct {
		AcruedAmount *big.Int
		SettledAt    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AcruedAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SettledAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 acruedAmount, uint256 settledAt)
func (_TickBit *TickBitSession) Pools(arg0 *big.Int) (struct {
	AcruedAmount *big.Int
	SettledAt    *big.Int
}, error) {
	return _TickBit.Contract.Pools(&_TickBit.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 acruedAmount, uint256 settledAt)
func (_TickBit *TickBitCallerSession) Pools(arg0 *big.Int) (struct {
	AcruedAmount *big.Int
	SettledAt    *big.Int
}, error) {
	return _TickBit.Contract.Pools(&_TickBit.CallOpts, arg0)
}

// TickSize is a free data retrieval call binding the contract method 0xf210d087.
//
// Solidity: function tickSize() view returns(uint256)
func (_TickBit *TickBitCaller) TickSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "tickSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSize is a free data retrieval call binding the contract method 0xf210d087.
//
// Solidity: function tickSize() view returns(uint256)
func (_TickBit *TickBitSession) TickSize() (*big.Int, error) {
	return _TickBit.Contract.TickSize(&_TickBit.CallOpts)
}

// TickSize is a free data retrieval call binding the contract method 0xf210d087.
//
// Solidity: function tickSize() view returns(uint256)
func (_TickBit *TickBitCallerSession) TickSize() (*big.Int, error) {
	return _TickBit.Contract.TickSize(&_TickBit.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TickBit *TickBitCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TickBit *TickBitSession) Token() (common.Address, error) {
	return _TickBit.Contract.Token(&_TickBit.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TickBit *TickBitCallerSession) Token() (common.Address, error) {
	return _TickBit.Contract.Token(&_TickBit.CallOpts)
}

// VerifiedBlocks is a free data retrieval call binding the contract method 0xb1d3e159.
//
// Solidity: function verifiedBlocks(uint256 ) view returns(bytes32 merkleRootHash, bytes4 nBits, bytes4 nonce, bytes32 previousBlockHash, bytes4 timestamp, bytes4 version)
func (_TickBit *TickBitCaller) VerifiedBlocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MerkleRootHash    [32]byte
	NBits             [4]byte
	Nonce             [4]byte
	PreviousBlockHash [32]byte
	Timestamp         [4]byte
	Version           [4]byte
}, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "verifiedBlocks", arg0)

	outstruct := new(struct {
		MerkleRootHash    [32]byte
		NBits             [4]byte
		Nonce             [4]byte
		PreviousBlockHash [32]byte
		Timestamp         [4]byte
		Version           [4]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MerkleRootHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.NBits = *abi.ConvertType(out[1], new([4]byte)).(*[4]byte)
	outstruct.Nonce = *abi.ConvertType(out[2], new([4]byte)).(*[4]byte)
	outstruct.PreviousBlockHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.Timestamp = *abi.ConvertType(out[4], new([4]byte)).(*[4]byte)
	outstruct.Version = *abi.ConvertType(out[5], new([4]byte)).(*[4]byte)

	return *outstruct, err

}

// VerifiedBlocks is a free data retrieval call binding the contract method 0xb1d3e159.
//
// Solidity: function verifiedBlocks(uint256 ) view returns(bytes32 merkleRootHash, bytes4 nBits, bytes4 nonce, bytes32 previousBlockHash, bytes4 timestamp, bytes4 version)
func (_TickBit *TickBitSession) VerifiedBlocks(arg0 *big.Int) (struct {
	MerkleRootHash    [32]byte
	NBits             [4]byte
	Nonce             [4]byte
	PreviousBlockHash [32]byte
	Timestamp         [4]byte
	Version           [4]byte
}, error) {
	return _TickBit.Contract.VerifiedBlocks(&_TickBit.CallOpts, arg0)
}

// VerifiedBlocks is a free data retrieval call binding the contract method 0xb1d3e159.
//
// Solidity: function verifiedBlocks(uint256 ) view returns(bytes32 merkleRootHash, bytes4 nBits, bytes4 nonce, bytes32 previousBlockHash, bytes4 timestamp, bytes4 version)
func (_TickBit *TickBitCallerSession) VerifiedBlocks(arg0 *big.Int) (struct {
	MerkleRootHash    [32]byte
	NBits             [4]byte
	Nonce             [4]byte
	PreviousBlockHash [32]byte
	Timestamp         [4]byte
	Version           [4]byte
}, error) {
	return _TickBit.Contract.VerifiedBlocks(&_TickBit.CallOpts, arg0)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_TickBit *TickBitCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_TickBit *TickBitSession) Verifier() (common.Address, error) {
	return _TickBit.Contract.Verifier(&_TickBit.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_TickBit *TickBitCallerSession) Verifier() (common.Address, error) {
	return _TickBit.Contract.Verifier(&_TickBit.CallOpts)
}

// VerifyBlock is a free data retrieval call binding the contract method 0xf1eac5d0.
//
// Solidity: function verifyBlock(uint256 blockNumber, bytes blockHeader, bytes proof) view returns((bytes32,bytes4,bytes4,bytes32,bytes4,bytes4) header)
func (_TickBit *TickBitCaller) VerifyBlock(opts *bind.CallOpts, blockNumber *big.Int, blockHeader []byte, proof []byte) (TickBitBlockHeader, error) {
	var out []interface{}
	err := _TickBit.contract.Call(opts, &out, "verifyBlock", blockNumber, blockHeader, proof)

	if err != nil {
		return *new(TickBitBlockHeader), err
	}

	out0 := *abi.ConvertType(out[0], new(TickBitBlockHeader)).(*TickBitBlockHeader)

	return out0, err

}

// VerifyBlock is a free data retrieval call binding the contract method 0xf1eac5d0.
//
// Solidity: function verifyBlock(uint256 blockNumber, bytes blockHeader, bytes proof) view returns((bytes32,bytes4,bytes4,bytes32,bytes4,bytes4) header)
func (_TickBit *TickBitSession) VerifyBlock(blockNumber *big.Int, blockHeader []byte, proof []byte) (TickBitBlockHeader, error) {
	return _TickBit.Contract.VerifyBlock(&_TickBit.CallOpts, blockNumber, blockHeader, proof)
}

// VerifyBlock is a free data retrieval call binding the contract method 0xf1eac5d0.
//
// Solidity: function verifyBlock(uint256 blockNumber, bytes blockHeader, bytes proof) view returns((bytes32,bytes4,bytes4,bytes32,bytes4,bytes4) header)
func (_TickBit *TickBitCallerSession) VerifyBlock(blockNumber *big.Int, blockHeader []byte, proof []byte) (TickBitBlockHeader, error) {
	return _TickBit.Contract.VerifyBlock(&_TickBit.CallOpts, blockNumber, blockHeader, proof)
}

// Bet is a paid mutator transaction binding the contract method 0x0c872d92.
//
// Solidity: function bet(uint256[] timestamps, uint256 blockNumber) returns()
func (_TickBit *TickBitTransactor) Bet(opts *bind.TransactOpts, timestamps []*big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _TickBit.contract.Transact(opts, "bet", timestamps, blockNumber)
}

// Bet is a paid mutator transaction binding the contract method 0x0c872d92.
//
// Solidity: function bet(uint256[] timestamps, uint256 blockNumber) returns()
func (_TickBit *TickBitSession) Bet(timestamps []*big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _TickBit.Contract.Bet(&_TickBit.TransactOpts, timestamps, blockNumber)
}

// Bet is a paid mutator transaction binding the contract method 0x0c872d92.
//
// Solidity: function bet(uint256[] timestamps, uint256 blockNumber) returns()
func (_TickBit *TickBitTransactorSession) Bet(timestamps []*big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _TickBit.Contract.Bet(&_TickBit.TransactOpts, timestamps, blockNumber)
}

// VerifyAndSettleBlock is a paid mutator transaction binding the contract method 0x1c6c39cc.
//
// Solidity: function verifyAndSettleBlock(uint256 blockNumber, bytes blockHeader, bytes proof) returns()
func (_TickBit *TickBitTransactor) VerifyAndSettleBlock(opts *bind.TransactOpts, blockNumber *big.Int, blockHeader []byte, proof []byte) (*types.Transaction, error) {
	return _TickBit.contract.Transact(opts, "verifyAndSettleBlock", blockNumber, blockHeader, proof)
}

// VerifyAndSettleBlock is a paid mutator transaction binding the contract method 0x1c6c39cc.
//
// Solidity: function verifyAndSettleBlock(uint256 blockNumber, bytes blockHeader, bytes proof) returns()
func (_TickBit *TickBitSession) VerifyAndSettleBlock(blockNumber *big.Int, blockHeader []byte, proof []byte) (*types.Transaction, error) {
	return _TickBit.Contract.VerifyAndSettleBlock(&_TickBit.TransactOpts, blockNumber, blockHeader, proof)
}

// VerifyAndSettleBlock is a paid mutator transaction binding the contract method 0x1c6c39cc.
//
// Solidity: function verifyAndSettleBlock(uint256 blockNumber, bytes blockHeader, bytes proof) returns()
func (_TickBit *TickBitTransactorSession) VerifyAndSettleBlock(blockNumber *big.Int, blockHeader []byte, proof []byte) (*types.Transaction, error) {
	return _TickBit.Contract.VerifyAndSettleBlock(&_TickBit.TransactOpts, blockNumber, blockHeader, proof)
}

// TickBitBetPlacedIterator is returned from FilterBetPlaced and is used to iterate over the raw logs and unpacked data for BetPlaced events raised by the TickBit contract.
type TickBitBetPlacedIterator struct {
	Event *TickBitBetPlaced // Event containing the contract specifics and raw log

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
func (it *TickBitBetPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TickBitBetPlaced)
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
		it.Event = new(TickBitBetPlaced)
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
func (it *TickBitBetPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TickBitBetPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TickBitBetPlaced represents a BetPlaced event raised by the TickBit contract.
type TickBitBetPlaced struct {
	Addr        common.Address
	BlockNumber *big.Int
	Timestamps  []*big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBetPlaced is a free log retrieval operation binding the contract event 0xf05a1393f3f7a79fabc31d47f94d644ed6ec5f80c12258ac176f612ed97360e3.
//
// Solidity: event BetPlaced(address indexed addr, uint256 indexed blockNumber, uint256[] timestamps, uint256 amount)
func (_TickBit *TickBitFilterer) FilterBetPlaced(opts *bind.FilterOpts, addr []common.Address, blockNumber []*big.Int) (*TickBitBetPlacedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _TickBit.contract.FilterLogs(opts, "BetPlaced", addrRule, blockNumberRule)
	if err != nil {
		return nil, err
	}
	return &TickBitBetPlacedIterator{contract: _TickBit.contract, event: "BetPlaced", logs: logs, sub: sub}, nil
}

// WatchBetPlaced is a free log subscription operation binding the contract event 0xf05a1393f3f7a79fabc31d47f94d644ed6ec5f80c12258ac176f612ed97360e3.
//
// Solidity: event BetPlaced(address indexed addr, uint256 indexed blockNumber, uint256[] timestamps, uint256 amount)
func (_TickBit *TickBitFilterer) WatchBetPlaced(opts *bind.WatchOpts, sink chan<- *TickBitBetPlaced, addr []common.Address, blockNumber []*big.Int) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _TickBit.contract.WatchLogs(opts, "BetPlaced", addrRule, blockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TickBitBetPlaced)
				if err := _TickBit.contract.UnpackLog(event, "BetPlaced", log); err != nil {
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

// ParseBetPlaced is a log parse operation binding the contract event 0xf05a1393f3f7a79fabc31d47f94d644ed6ec5f80c12258ac176f612ed97360e3.
//
// Solidity: event BetPlaced(address indexed addr, uint256 indexed blockNumber, uint256[] timestamps, uint256 amount)
func (_TickBit *TickBitFilterer) ParseBetPlaced(log types.Log) (*TickBitBetPlaced, error) {
	event := new(TickBitBetPlaced)
	if err := _TickBit.contract.UnpackLog(event, "BetPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
