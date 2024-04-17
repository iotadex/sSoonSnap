// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cont

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

// StakeNFTMetaData contains all meta data concerning the StakeNFT contract.
var StakeNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maxWeeks\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxScale\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"lockWeeks\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"beginTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickMax\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"no\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"k\",\"type\":\"uint8\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mergeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token0Fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token1Fee\",\"type\":\"uint256\"}],\"name\":\"WithdrawAndMerge\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BEGIN_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"END_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOCK_WEEKNUM\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SCALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TICK\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WEEKS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_TICK\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WEEK_SECONDS\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"canClaimAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"k\",\"type\":\"uint8\"}],\"name\":\"getScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserNFTs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"lockedRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftToken\",\"outputs\":[{\"internalType\":\"contractIIotabeeSwapNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"no\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"setReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nos\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"setRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"k\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakingNFTs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"beginNo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endNo\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalScores\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userCanClaimWeeks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userNFTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userScores\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mergeTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"withdrawAndMerge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakeNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeNFTMetaData.ABI instead.
var StakeNFTABI = StakeNFTMetaData.ABI

// StakeNFT is an auto generated Go binding around an Ethereum contract.
type StakeNFT struct {
	StakeNFTCaller     // Read-only binding to the contract
	StakeNFTTransactor // Write-only binding to the contract
	StakeNFTFilterer   // Log filterer for contract events
}

// StakeNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeNFTSession struct {
	Contract     *StakeNFT         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeNFTCallerSession struct {
	Contract *StakeNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StakeNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeNFTTransactorSession struct {
	Contract     *StakeNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StakeNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeNFTRaw struct {
	Contract *StakeNFT // Generic contract binding to access the raw methods on
}

// StakeNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeNFTCallerRaw struct {
	Contract *StakeNFTCaller // Generic read-only contract binding to access the raw methods on
}

// StakeNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeNFTTransactorRaw struct {
	Contract *StakeNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeNFT creates a new instance of StakeNFT, bound to a specific deployed contract.
func NewStakeNFT(address common.Address, backend bind.ContractBackend) (*StakeNFT, error) {
	contract, err := bindStakeNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeNFT{StakeNFTCaller: StakeNFTCaller{contract: contract}, StakeNFTTransactor: StakeNFTTransactor{contract: contract}, StakeNFTFilterer: StakeNFTFilterer{contract: contract}}, nil
}

// NewStakeNFTCaller creates a new read-only instance of StakeNFT, bound to a specific deployed contract.
func NewStakeNFTCaller(address common.Address, caller bind.ContractCaller) (*StakeNFTCaller, error) {
	contract, err := bindStakeNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeNFTCaller{contract: contract}, nil
}

// NewStakeNFTTransactor creates a new write-only instance of StakeNFT, bound to a specific deployed contract.
func NewStakeNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeNFTTransactor, error) {
	contract, err := bindStakeNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeNFTTransactor{contract: contract}, nil
}

// NewStakeNFTFilterer creates a new log filterer instance of StakeNFT, bound to a specific deployed contract.
func NewStakeNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeNFTFilterer, error) {
	contract, err := bindStakeNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeNFTFilterer{contract: contract}, nil
}

// bindStakeNFT binds a generic wrapper to an already deployed contract.
func bindStakeNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeNFT *StakeNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeNFT.Contract.StakeNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeNFT *StakeNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeNFT.Contract.StakeNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeNFT *StakeNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeNFT.Contract.StakeNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeNFT *StakeNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeNFT *StakeNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeNFT *StakeNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeNFT.Contract.contract.Transact(opts, method, params...)
}

// BEGINTIME is a free data retrieval call binding the contract method 0xc75363b6.
//
// Solidity: function BEGIN_TIME() view returns(uint256)
func (_StakeNFT *StakeNFTCaller) BEGINTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "BEGIN_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BEGINTIME is a free data retrieval call binding the contract method 0xc75363b6.
//
// Solidity: function BEGIN_TIME() view returns(uint256)
func (_StakeNFT *StakeNFTSession) BEGINTIME() (*big.Int, error) {
	return _StakeNFT.Contract.BEGINTIME(&_StakeNFT.CallOpts)
}

// BEGINTIME is a free data retrieval call binding the contract method 0xc75363b6.
//
// Solidity: function BEGIN_TIME() view returns(uint256)
func (_StakeNFT *StakeNFTCallerSession) BEGINTIME() (*big.Int, error) {
	return _StakeNFT.Contract.BEGINTIME(&_StakeNFT.CallOpts)
}

// ENDTIME is a free data retrieval call binding the contract method 0x37ba682d.
//
// Solidity: function END_TIME() view returns(uint256)
func (_StakeNFT *StakeNFTCaller) ENDTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "END_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ENDTIME is a free data retrieval call binding the contract method 0x37ba682d.
//
// Solidity: function END_TIME() view returns(uint256)
func (_StakeNFT *StakeNFTSession) ENDTIME() (*big.Int, error) {
	return _StakeNFT.Contract.ENDTIME(&_StakeNFT.CallOpts)
}

// ENDTIME is a free data retrieval call binding the contract method 0x37ba682d.
//
// Solidity: function END_TIME() view returns(uint256)
func (_StakeNFT *StakeNFTCallerSession) ENDTIME() (*big.Int, error) {
	return _StakeNFT.Contract.ENDTIME(&_StakeNFT.CallOpts)
}

// LOCKWEEKNUM is a free data retrieval call binding the contract method 0x0d437d40.
//
// Solidity: function LOCK_WEEKNUM() view returns(uint8)
func (_StakeNFT *StakeNFTCaller) LOCKWEEKNUM(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "LOCK_WEEKNUM")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LOCKWEEKNUM is a free data retrieval call binding the contract method 0x0d437d40.
//
// Solidity: function LOCK_WEEKNUM() view returns(uint8)
func (_StakeNFT *StakeNFTSession) LOCKWEEKNUM() (uint8, error) {
	return _StakeNFT.Contract.LOCKWEEKNUM(&_StakeNFT.CallOpts)
}

// LOCKWEEKNUM is a free data retrieval call binding the contract method 0x0d437d40.
//
// Solidity: function LOCK_WEEKNUM() view returns(uint8)
func (_StakeNFT *StakeNFTCallerSession) LOCKWEEKNUM() (uint8, error) {
	return _StakeNFT.Contract.LOCKWEEKNUM(&_StakeNFT.CallOpts)
}

// MAXSCALE is a free data retrieval call binding the contract method 0xf6def7cf.
//
// Solidity: function MAX_SCALE() view returns(uint256)
func (_StakeNFT *StakeNFTCaller) MAXSCALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "MAX_SCALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSCALE is a free data retrieval call binding the contract method 0xf6def7cf.
//
// Solidity: function MAX_SCALE() view returns(uint256)
func (_StakeNFT *StakeNFTSession) MAXSCALE() (*big.Int, error) {
	return _StakeNFT.Contract.MAXSCALE(&_StakeNFT.CallOpts)
}

// MAXSCALE is a free data retrieval call binding the contract method 0xf6def7cf.
//
// Solidity: function MAX_SCALE() view returns(uint256)
func (_StakeNFT *StakeNFTCallerSession) MAXSCALE() (*big.Int, error) {
	return _StakeNFT.Contract.MAXSCALE(&_StakeNFT.CallOpts)
}

// MAXTICK is a free data retrieval call binding the contract method 0x6882a888.
//
// Solidity: function MAX_TICK() view returns(int24)
func (_StakeNFT *StakeNFTCaller) MAXTICK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "MAX_TICK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTICK is a free data retrieval call binding the contract method 0x6882a888.
//
// Solidity: function MAX_TICK() view returns(int24)
func (_StakeNFT *StakeNFTSession) MAXTICK() (*big.Int, error) {
	return _StakeNFT.Contract.MAXTICK(&_StakeNFT.CallOpts)
}

// MAXTICK is a free data retrieval call binding the contract method 0x6882a888.
//
// Solidity: function MAX_TICK() view returns(int24)
func (_StakeNFT *StakeNFTCallerSession) MAXTICK() (*big.Int, error) {
	return _StakeNFT.Contract.MAXTICK(&_StakeNFT.CallOpts)
}

// MAXWEEKS is a free data retrieval call binding the contract method 0xe11400bb.
//
// Solidity: function MAX_WEEKS() view returns(uint8)
func (_StakeNFT *StakeNFTCaller) MAXWEEKS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "MAX_WEEKS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXWEEKS is a free data retrieval call binding the contract method 0xe11400bb.
//
// Solidity: function MAX_WEEKS() view returns(uint8)
func (_StakeNFT *StakeNFTSession) MAXWEEKS() (uint8, error) {
	return _StakeNFT.Contract.MAXWEEKS(&_StakeNFT.CallOpts)
}

// MAXWEEKS is a free data retrieval call binding the contract method 0xe11400bb.
//
// Solidity: function MAX_WEEKS() view returns(uint8)
func (_StakeNFT *StakeNFTCallerSession) MAXWEEKS() (uint8, error) {
	return _StakeNFT.Contract.MAXWEEKS(&_StakeNFT.CallOpts)
}

// MINTICK is a free data retrieval call binding the contract method 0xa1634b14.
//
// Solidity: function MIN_TICK() view returns(int24)
func (_StakeNFT *StakeNFTCaller) MINTICK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "MIN_TICK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTICK is a free data retrieval call binding the contract method 0xa1634b14.
//
// Solidity: function MIN_TICK() view returns(int24)
func (_StakeNFT *StakeNFTSession) MINTICK() (*big.Int, error) {
	return _StakeNFT.Contract.MINTICK(&_StakeNFT.CallOpts)
}

// MINTICK is a free data retrieval call binding the contract method 0xa1634b14.
//
// Solidity: function MIN_TICK() view returns(int24)
func (_StakeNFT *StakeNFTCallerSession) MINTICK() (*big.Int, error) {
	return _StakeNFT.Contract.MINTICK(&_StakeNFT.CallOpts)
}

// WEEKSECONDS is a free data retrieval call binding the contract method 0xcd671a58.
//
// Solidity: function WEEK_SECONDS() view returns(uint24)
func (_StakeNFT *StakeNFTCaller) WEEKSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "WEEK_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WEEKSECONDS is a free data retrieval call binding the contract method 0xcd671a58.
//
// Solidity: function WEEK_SECONDS() view returns(uint24)
func (_StakeNFT *StakeNFTSession) WEEKSECONDS() (*big.Int, error) {
	return _StakeNFT.Contract.WEEKSECONDS(&_StakeNFT.CallOpts)
}

// WEEKSECONDS is a free data retrieval call binding the contract method 0xcd671a58.
//
// Solidity: function WEEK_SECONDS() view returns(uint24)
func (_StakeNFT *StakeNFTCallerSession) WEEKSECONDS() (*big.Int, error) {
	return _StakeNFT.Contract.WEEKSECONDS(&_StakeNFT.CallOpts)
}

// CanClaimAmount is a free data retrieval call binding the contract method 0xfa60c77c.
//
// Solidity: function canClaimAmount(address user) view returns(uint256)
func (_StakeNFT *StakeNFTCaller) CanClaimAmount(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "canClaimAmount", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CanClaimAmount is a free data retrieval call binding the contract method 0xfa60c77c.
//
// Solidity: function canClaimAmount(address user) view returns(uint256)
func (_StakeNFT *StakeNFTSession) CanClaimAmount(user common.Address) (*big.Int, error) {
	return _StakeNFT.Contract.CanClaimAmount(&_StakeNFT.CallOpts, user)
}

// CanClaimAmount is a free data retrieval call binding the contract method 0xfa60c77c.
//
// Solidity: function canClaimAmount(address user) view returns(uint256)
func (_StakeNFT *StakeNFTCallerSession) CanClaimAmount(user common.Address) (*big.Int, error) {
	return _StakeNFT.Contract.CanClaimAmount(&_StakeNFT.CallOpts, user)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_StakeNFT *StakeNFTCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_StakeNFT *StakeNFTSession) Fee() (*big.Int, error) {
	return _StakeNFT.Contract.Fee(&_StakeNFT.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_StakeNFT *StakeNFTCallerSession) Fee() (*big.Int, error) {
	return _StakeNFT.Contract.Fee(&_StakeNFT.CallOpts)
}

// GetScore is a free data retrieval call binding the contract method 0xee3453ed.
//
// Solidity: function getScore(uint256 amount, uint8 k) view returns(uint256 score)
func (_StakeNFT *StakeNFTCaller) GetScore(opts *bind.CallOpts, amount *big.Int, k uint8) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "getScore", amount, k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetScore is a free data retrieval call binding the contract method 0xee3453ed.
//
// Solidity: function getScore(uint256 amount, uint8 k) view returns(uint256 score)
func (_StakeNFT *StakeNFTSession) GetScore(amount *big.Int, k uint8) (*big.Int, error) {
	return _StakeNFT.Contract.GetScore(&_StakeNFT.CallOpts, amount, k)
}

// GetScore is a free data retrieval call binding the contract method 0xee3453ed.
//
// Solidity: function getScore(uint256 amount, uint8 k) view returns(uint256 score)
func (_StakeNFT *StakeNFTCallerSession) GetScore(amount *big.Int, k uint8) (*big.Int, error) {
	return _StakeNFT.Contract.GetScore(&_StakeNFT.CallOpts, amount, k)
}

// GetUserNFTs is a free data retrieval call binding the contract method 0xfdc5ab8e.
//
// Solidity: function getUserNFTs() view returns(uint256[], uint256)
func (_StakeNFT *StakeNFTCaller) GetUserNFTs(opts *bind.CallOpts) ([]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "getUserNFTs")

	if err != nil {
		return *new([]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUserNFTs is a free data retrieval call binding the contract method 0xfdc5ab8e.
//
// Solidity: function getUserNFTs() view returns(uint256[], uint256)
func (_StakeNFT *StakeNFTSession) GetUserNFTs() ([]*big.Int, *big.Int, error) {
	return _StakeNFT.Contract.GetUserNFTs(&_StakeNFT.CallOpts)
}

// GetUserNFTs is a free data retrieval call binding the contract method 0xfdc5ab8e.
//
// Solidity: function getUserNFTs() view returns(uint256[], uint256)
func (_StakeNFT *StakeNFTCallerSession) GetUserNFTs() ([]*big.Int, *big.Int, error) {
	return _StakeNFT.Contract.GetUserNFTs(&_StakeNFT.CallOpts)
}

// LatestNo is a free data retrieval call binding the contract method 0x29ed1cfd.
//
// Solidity: function latestNo() view returns(uint256)
func (_StakeNFT *StakeNFTCaller) LatestNo(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "latestNo")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestNo is a free data retrieval call binding the contract method 0x29ed1cfd.
//
// Solidity: function latestNo() view returns(uint256)
func (_StakeNFT *StakeNFTSession) LatestNo() (*big.Int, error) {
	return _StakeNFT.Contract.LatestNo(&_StakeNFT.CallOpts)
}

// LatestNo is a free data retrieval call binding the contract method 0x29ed1cfd.
//
// Solidity: function latestNo() view returns(uint256)
func (_StakeNFT *StakeNFTCallerSession) LatestNo() (*big.Int, error) {
	return _StakeNFT.Contract.LatestNo(&_StakeNFT.CallOpts)
}

// LockedRewardAmount is a free data retrieval call binding the contract method 0x02e9e495.
//
// Solidity: function lockedRewardAmount(address user) view returns(uint256, uint256[])
func (_StakeNFT *StakeNFTCaller) LockedRewardAmount(opts *bind.CallOpts, user common.Address) (*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "lockedRewardAmount", user)

	if err != nil {
		return *new(*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// LockedRewardAmount is a free data retrieval call binding the contract method 0x02e9e495.
//
// Solidity: function lockedRewardAmount(address user) view returns(uint256, uint256[])
func (_StakeNFT *StakeNFTSession) LockedRewardAmount(user common.Address) (*big.Int, []*big.Int, error) {
	return _StakeNFT.Contract.LockedRewardAmount(&_StakeNFT.CallOpts, user)
}

// LockedRewardAmount is a free data retrieval call binding the contract method 0x02e9e495.
//
// Solidity: function lockedRewardAmount(address user) view returns(uint256, uint256[])
func (_StakeNFT *StakeNFTCallerSession) LockedRewardAmount(user common.Address) (*big.Int, []*big.Int, error) {
	return _StakeNFT.Contract.LockedRewardAmount(&_StakeNFT.CallOpts, user)
}

// NftToken is a free data retrieval call binding the contract method 0xd06fcba8.
//
// Solidity: function nftToken() view returns(address)
func (_StakeNFT *StakeNFTCaller) NftToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "nftToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftToken is a free data retrieval call binding the contract method 0xd06fcba8.
//
// Solidity: function nftToken() view returns(address)
func (_StakeNFT *StakeNFTSession) NftToken() (common.Address, error) {
	return _StakeNFT.Contract.NftToken(&_StakeNFT.CallOpts)
}

// NftToken is a free data retrieval call binding the contract method 0xd06fcba8.
//
// Solidity: function nftToken() view returns(address)
func (_StakeNFT *StakeNFTCallerSession) NftToken() (common.Address, error) {
	return _StakeNFT.Contract.NftToken(&_StakeNFT.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_StakeNFT *StakeNFTCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_StakeNFT *StakeNFTSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _StakeNFT.Contract.OnERC721Received(&_StakeNFT.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_StakeNFT *StakeNFTCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _StakeNFT.Contract.OnERC721Received(&_StakeNFT.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeNFT *StakeNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeNFT *StakeNFTSession) Owner() (common.Address, error) {
	return _StakeNFT.Contract.Owner(&_StakeNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeNFT *StakeNFTCallerSession) Owner() (common.Address, error) {
	return _StakeNFT.Contract.Owner(&_StakeNFT.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_StakeNFT *StakeNFTCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_StakeNFT *StakeNFTSession) RewardToken() (common.Address, error) {
	return _StakeNFT.Contract.RewardToken(&_StakeNFT.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_StakeNFT *StakeNFTCallerSession) RewardToken() (common.Address, error) {
	return _StakeNFT.Contract.RewardToken(&_StakeNFT.CallOpts)
}

// RewardsOf is a free data retrieval call binding the contract method 0x0bce1284.
//
// Solidity: function rewardsOf(uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTCaller) RewardsOf(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "rewardsOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsOf is a free data retrieval call binding the contract method 0x0bce1284.
//
// Solidity: function rewardsOf(uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTSession) RewardsOf(arg0 *big.Int) (*big.Int, error) {
	return _StakeNFT.Contract.RewardsOf(&_StakeNFT.CallOpts, arg0)
}

// RewardsOf is a free data retrieval call binding the contract method 0x0bce1284.
//
// Solidity: function rewardsOf(uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTCallerSession) RewardsOf(arg0 *big.Int) (*big.Int, error) {
	return _StakeNFT.Contract.RewardsOf(&_StakeNFT.CallOpts, arg0)
}

// StakingNFTs is a free data retrieval call binding the contract method 0x10c6c1c2.
//
// Solidity: function stakingNFTs(uint256 ) view returns(address owner, uint256 score, uint256 beginNo, uint256 endNo)
func (_StakeNFT *StakeNFTCaller) StakingNFTs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner   common.Address
	Score   *big.Int
	BeginNo *big.Int
	EndNo   *big.Int
}, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "stakingNFTs", arg0)

	outstruct := new(struct {
		Owner   common.Address
		Score   *big.Int
		BeginNo *big.Int
		EndNo   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Score = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BeginNo = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndNo = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StakingNFTs is a free data retrieval call binding the contract method 0x10c6c1c2.
//
// Solidity: function stakingNFTs(uint256 ) view returns(address owner, uint256 score, uint256 beginNo, uint256 endNo)
func (_StakeNFT *StakeNFTSession) StakingNFTs(arg0 *big.Int) (struct {
	Owner   common.Address
	Score   *big.Int
	BeginNo *big.Int
	EndNo   *big.Int
}, error) {
	return _StakeNFT.Contract.StakingNFTs(&_StakeNFT.CallOpts, arg0)
}

// StakingNFTs is a free data retrieval call binding the contract method 0x10c6c1c2.
//
// Solidity: function stakingNFTs(uint256 ) view returns(address owner, uint256 score, uint256 beginNo, uint256 endNo)
func (_StakeNFT *StakeNFTCallerSession) StakingNFTs(arg0 *big.Int) (struct {
	Owner   common.Address
	Score   *big.Int
	BeginNo *big.Int
	EndNo   *big.Int
}, error) {
	return _StakeNFT.Contract.StakingNFTs(&_StakeNFT.CallOpts, arg0)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_StakeNFT *StakeNFTCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_StakeNFT *StakeNFTSession) Token0() (common.Address, error) {
	return _StakeNFT.Contract.Token0(&_StakeNFT.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_StakeNFT *StakeNFTCallerSession) Token0() (common.Address, error) {
	return _StakeNFT.Contract.Token0(&_StakeNFT.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_StakeNFT *StakeNFTCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_StakeNFT *StakeNFTSession) Token1() (common.Address, error) {
	return _StakeNFT.Contract.Token1(&_StakeNFT.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_StakeNFT *StakeNFTCallerSession) Token1() (common.Address, error) {
	return _StakeNFT.Contract.Token1(&_StakeNFT.CallOpts)
}

// TotalScores is a free data retrieval call binding the contract method 0x157775df.
//
// Solidity: function totalScores(uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTCaller) TotalScores(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "totalScores", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalScores is a free data retrieval call binding the contract method 0x157775df.
//
// Solidity: function totalScores(uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTSession) TotalScores(arg0 *big.Int) (*big.Int, error) {
	return _StakeNFT.Contract.TotalScores(&_StakeNFT.CallOpts, arg0)
}

// TotalScores is a free data retrieval call binding the contract method 0x157775df.
//
// Solidity: function totalScores(uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTCallerSession) TotalScores(arg0 *big.Int) (*big.Int, error) {
	return _StakeNFT.Contract.TotalScores(&_StakeNFT.CallOpts, arg0)
}

// UserCanClaimWeeks is a free data retrieval call binding the contract method 0xbb57852d.
//
// Solidity: function userCanClaimWeeks(address , uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTCaller) UserCanClaimWeeks(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "userCanClaimWeeks", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserCanClaimWeeks is a free data retrieval call binding the contract method 0xbb57852d.
//
// Solidity: function userCanClaimWeeks(address , uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTSession) UserCanClaimWeeks(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _StakeNFT.Contract.UserCanClaimWeeks(&_StakeNFT.CallOpts, arg0, arg1)
}

// UserCanClaimWeeks is a free data retrieval call binding the contract method 0xbb57852d.
//
// Solidity: function userCanClaimWeeks(address , uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTCallerSession) UserCanClaimWeeks(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _StakeNFT.Contract.UserCanClaimWeeks(&_StakeNFT.CallOpts, arg0, arg1)
}

// UserNFTs is a free data retrieval call binding the contract method 0x291c73c3.
//
// Solidity: function userNFTs(address , uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTCaller) UserNFTs(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "userNFTs", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserNFTs is a free data retrieval call binding the contract method 0x291c73c3.
//
// Solidity: function userNFTs(address , uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTSession) UserNFTs(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _StakeNFT.Contract.UserNFTs(&_StakeNFT.CallOpts, arg0, arg1)
}

// UserNFTs is a free data retrieval call binding the contract method 0x291c73c3.
//
// Solidity: function userNFTs(address , uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTCallerSession) UserNFTs(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _StakeNFT.Contract.UserNFTs(&_StakeNFT.CallOpts, arg0, arg1)
}

// UserScores is a free data retrieval call binding the contract method 0x5b680a45.
//
// Solidity: function userScores(address , uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTCaller) UserScores(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeNFT.contract.Call(opts, &out, "userScores", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserScores is a free data retrieval call binding the contract method 0x5b680a45.
//
// Solidity: function userScores(address , uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTSession) UserScores(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _StakeNFT.Contract.UserScores(&_StakeNFT.CallOpts, arg0, arg1)
}

// UserScores is a free data retrieval call binding the contract method 0x5b680a45.
//
// Solidity: function userScores(address , uint256 ) view returns(uint256)
func (_StakeNFT *StakeNFTCallerSession) UserScores(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _StakeNFT.Contract.UserScores(&_StakeNFT.CallOpts, arg0, arg1)
}

// AcceptSetter is a paid mutator transaction binding the contract method 0x5f8e422f.
//
// Solidity: function acceptSetter() returns()
func (_StakeNFT *StakeNFTTransactor) AcceptSetter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeNFT.contract.Transact(opts, "acceptSetter")
}

// AcceptSetter is a paid mutator transaction binding the contract method 0x5f8e422f.
//
// Solidity: function acceptSetter() returns()
func (_StakeNFT *StakeNFTSession) AcceptSetter() (*types.Transaction, error) {
	return _StakeNFT.Contract.AcceptSetter(&_StakeNFT.TransactOpts)
}

// AcceptSetter is a paid mutator transaction binding the contract method 0x5f8e422f.
//
// Solidity: function acceptSetter() returns()
func (_StakeNFT *StakeNFTTransactorSession) AcceptSetter() (*types.Transaction, error) {
	return _StakeNFT.Contract.AcceptSetter(&_StakeNFT.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 deadline) returns()
func (_StakeNFT *StakeNFTTransactor) ClaimReward(opts *bind.TransactOpts, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.contract.Transact(opts, "claimReward", deadline)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 deadline) returns()
func (_StakeNFT *StakeNFTSession) ClaimReward(deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.Contract.ClaimReward(&_StakeNFT.TransactOpts, deadline)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 deadline) returns()
func (_StakeNFT *StakeNFTTransactorSession) ClaimReward(deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.Contract.ClaimReward(&_StakeNFT.TransactOpts, deadline)
}

// SetReward is a paid mutator transaction binding the contract method 0xe1c05d93.
//
// Solidity: function setReward(uint256 no, uint256 amount, uint256 deadline) returns()
func (_StakeNFT *StakeNFTTransactor) SetReward(opts *bind.TransactOpts, no *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.contract.Transact(opts, "setReward", no, amount, deadline)
}

// SetReward is a paid mutator transaction binding the contract method 0xe1c05d93.
//
// Solidity: function setReward(uint256 no, uint256 amount, uint256 deadline) returns()
func (_StakeNFT *StakeNFTSession) SetReward(no *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.Contract.SetReward(&_StakeNFT.TransactOpts, no, amount, deadline)
}

// SetReward is a paid mutator transaction binding the contract method 0xe1c05d93.
//
// Solidity: function setReward(uint256 no, uint256 amount, uint256 deadline) returns()
func (_StakeNFT *StakeNFTTransactorSession) SetReward(no *big.Int, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.Contract.SetReward(&_StakeNFT.TransactOpts, no, amount, deadline)
}

// SetRewards is a paid mutator transaction binding the contract method 0x07c6e2fa.
//
// Solidity: function setRewards(uint256[] nos, uint256[] amounts, uint256 deadline) returns(uint256 total)
func (_StakeNFT *StakeNFTTransactor) SetRewards(opts *bind.TransactOpts, nos []*big.Int, amounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.contract.Transact(opts, "setRewards", nos, amounts, deadline)
}

// SetRewards is a paid mutator transaction binding the contract method 0x07c6e2fa.
//
// Solidity: function setRewards(uint256[] nos, uint256[] amounts, uint256 deadline) returns(uint256 total)
func (_StakeNFT *StakeNFTSession) SetRewards(nos []*big.Int, amounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.Contract.SetRewards(&_StakeNFT.TransactOpts, nos, amounts, deadline)
}

// SetRewards is a paid mutator transaction binding the contract method 0x07c6e2fa.
//
// Solidity: function setRewards(uint256[] nos, uint256[] amounts, uint256 deadline) returns(uint256 total)
func (_StakeNFT *StakeNFTTransactorSession) SetRewards(nos []*big.Int, amounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.Contract.SetRewards(&_StakeNFT.TransactOpts, nos, amounts, deadline)
}

// Stake is a paid mutator transaction binding the contract method 0xf99b7d75.
//
// Solidity: function stake(uint256 tokenId, uint8 k, uint256 deadline) returns()
func (_StakeNFT *StakeNFTTransactor) Stake(opts *bind.TransactOpts, tokenId *big.Int, k uint8, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.contract.Transact(opts, "stake", tokenId, k, deadline)
}

// Stake is a paid mutator transaction binding the contract method 0xf99b7d75.
//
// Solidity: function stake(uint256 tokenId, uint8 k, uint256 deadline) returns()
func (_StakeNFT *StakeNFTSession) Stake(tokenId *big.Int, k uint8, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.Contract.Stake(&_StakeNFT.TransactOpts, tokenId, k, deadline)
}

// Stake is a paid mutator transaction binding the contract method 0xf99b7d75.
//
// Solidity: function stake(uint256 tokenId, uint8 k, uint256 deadline) returns()
func (_StakeNFT *StakeNFTTransactorSession) Stake(tokenId *big.Int, k uint8, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.Contract.Stake(&_StakeNFT.TransactOpts, tokenId, k, deadline)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address _owner) returns()
func (_StakeNFT *StakeNFTTransactor) TransferOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _StakeNFT.contract.Transact(opts, "transferOwner", _owner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address _owner) returns()
func (_StakeNFT *StakeNFTSession) TransferOwner(_owner common.Address) (*types.Transaction, error) {
	return _StakeNFT.Contract.TransferOwner(&_StakeNFT.TransactOpts, _owner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(address _owner) returns()
func (_StakeNFT *StakeNFTTransactorSession) TransferOwner(_owner common.Address) (*types.Transaction, error) {
	return _StakeNFT.Contract.TransferOwner(&_StakeNFT.TransactOpts, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 tokenId, uint256 deadline) returns()
func (_StakeNFT *StakeNFTTransactor) Withdraw(opts *bind.TransactOpts, tokenId *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.contract.Transact(opts, "withdraw", tokenId, deadline)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 tokenId, uint256 deadline) returns()
func (_StakeNFT *StakeNFTSession) Withdraw(tokenId *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.Contract.Withdraw(&_StakeNFT.TransactOpts, tokenId, deadline)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 tokenId, uint256 deadline) returns()
func (_StakeNFT *StakeNFTTransactorSession) Withdraw(tokenId *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.Contract.Withdraw(&_StakeNFT.TransactOpts, tokenId, deadline)
}

// WithdrawAndMerge is a paid mutator transaction binding the contract method 0x1a906d65.
//
// Solidity: function withdrawAndMerge(uint256 tokenId, uint256 mergeTokenId, uint256 deadline) returns()
func (_StakeNFT *StakeNFTTransactor) WithdrawAndMerge(opts *bind.TransactOpts, tokenId *big.Int, mergeTokenId *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.contract.Transact(opts, "withdrawAndMerge", tokenId, mergeTokenId, deadline)
}

// WithdrawAndMerge is a paid mutator transaction binding the contract method 0x1a906d65.
//
// Solidity: function withdrawAndMerge(uint256 tokenId, uint256 mergeTokenId, uint256 deadline) returns()
func (_StakeNFT *StakeNFTSession) WithdrawAndMerge(tokenId *big.Int, mergeTokenId *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.Contract.WithdrawAndMerge(&_StakeNFT.TransactOpts, tokenId, mergeTokenId, deadline)
}

// WithdrawAndMerge is a paid mutator transaction binding the contract method 0x1a906d65.
//
// Solidity: function withdrawAndMerge(uint256 tokenId, uint256 mergeTokenId, uint256 deadline) returns()
func (_StakeNFT *StakeNFTTransactorSession) WithdrawAndMerge(tokenId *big.Int, mergeTokenId *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.Contract.WithdrawAndMerge(&_StakeNFT.TransactOpts, tokenId, mergeTokenId, deadline)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x52cb5133.
//
// Solidity: function withdrawRewards(uint256 amount, uint256 deadline) returns()
func (_StakeNFT *StakeNFTTransactor) WithdrawRewards(opts *bind.TransactOpts, amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.contract.Transact(opts, "withdrawRewards", amount, deadline)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x52cb5133.
//
// Solidity: function withdrawRewards(uint256 amount, uint256 deadline) returns()
func (_StakeNFT *StakeNFTSession) WithdrawRewards(amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.Contract.WithdrawRewards(&_StakeNFT.TransactOpts, amount, deadline)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x52cb5133.
//
// Solidity: function withdrawRewards(uint256 amount, uint256 deadline) returns()
func (_StakeNFT *StakeNFTTransactorSession) WithdrawRewards(amount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakeNFT.Contract.WithdrawRewards(&_StakeNFT.TransactOpts, amount, deadline)
}

// StakeNFTClaimRewardIterator is returned from FilterClaimReward and is used to iterate over the raw logs and unpacked data for ClaimReward events raised by the StakeNFT contract.
type StakeNFTClaimRewardIterator struct {
	Event *StakeNFTClaimReward // Event containing the contract specifics and raw log

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
func (it *StakeNFTClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNFTClaimReward)
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
		it.Event = new(StakeNFTClaimReward)
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
func (it *StakeNFTClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNFTClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNFTClaimReward represents a ClaimReward event raised by the StakeNFT contract.
type StakeNFTClaimReward struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimReward is a free log retrieval operation binding the contract event 0xba8de60c3403ec381d1d484652ea1980e3c3e56359195c92525bff4ce47ad98e.
//
// Solidity: event ClaimReward(address indexed user, uint256 amount)
func (_StakeNFT *StakeNFTFilterer) FilterClaimReward(opts *bind.FilterOpts, user []common.Address) (*StakeNFTClaimRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNFT.contract.FilterLogs(opts, "ClaimReward", userRule)
	if err != nil {
		return nil, err
	}
	return &StakeNFTClaimRewardIterator{contract: _StakeNFT.contract, event: "ClaimReward", logs: logs, sub: sub}, nil
}

// WatchClaimReward is a free log subscription operation binding the contract event 0xba8de60c3403ec381d1d484652ea1980e3c3e56359195c92525bff4ce47ad98e.
//
// Solidity: event ClaimReward(address indexed user, uint256 amount)
func (_StakeNFT *StakeNFTFilterer) WatchClaimReward(opts *bind.WatchOpts, sink chan<- *StakeNFTClaimReward, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNFT.contract.WatchLogs(opts, "ClaimReward", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNFTClaimReward)
				if err := _StakeNFT.contract.UnpackLog(event, "ClaimReward", log); err != nil {
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

// ParseClaimReward is a log parse operation binding the contract event 0xba8de60c3403ec381d1d484652ea1980e3c3e56359195c92525bff4ce47ad98e.
//
// Solidity: event ClaimReward(address indexed user, uint256 amount)
func (_StakeNFT *StakeNFTFilterer) ParseClaimReward(log types.Log) (*StakeNFTClaimReward, error) {
	event := new(StakeNFTClaimReward)
	if err := _StakeNFT.contract.UnpackLog(event, "ClaimReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNFTSetRewardIterator is returned from FilterSetReward and is used to iterate over the raw logs and unpacked data for SetReward events raised by the StakeNFT contract.
type StakeNFTSetRewardIterator struct {
	Event *StakeNFTSetReward // Event containing the contract specifics and raw log

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
func (it *StakeNFTSetRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNFTSetReward)
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
		it.Event = new(StakeNFTSetReward)
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
func (it *StakeNFTSetRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNFTSetRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNFTSetReward represents a SetReward event raised by the StakeNFT contract.
type StakeNFTSetReward struct {
	User   common.Address
	No     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetReward is a free log retrieval operation binding the contract event 0xc3991418cacd70a010fce559a44fcbb47534ca43bf7543964bdca2639f6e6e5a.
//
// Solidity: event SetReward(address indexed user, uint256 no, uint256 amount)
func (_StakeNFT *StakeNFTFilterer) FilterSetReward(opts *bind.FilterOpts, user []common.Address) (*StakeNFTSetRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNFT.contract.FilterLogs(opts, "SetReward", userRule)
	if err != nil {
		return nil, err
	}
	return &StakeNFTSetRewardIterator{contract: _StakeNFT.contract, event: "SetReward", logs: logs, sub: sub}, nil
}

// WatchSetReward is a free log subscription operation binding the contract event 0xc3991418cacd70a010fce559a44fcbb47534ca43bf7543964bdca2639f6e6e5a.
//
// Solidity: event SetReward(address indexed user, uint256 no, uint256 amount)
func (_StakeNFT *StakeNFTFilterer) WatchSetReward(opts *bind.WatchOpts, sink chan<- *StakeNFTSetReward, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNFT.contract.WatchLogs(opts, "SetReward", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNFTSetReward)
				if err := _StakeNFT.contract.UnpackLog(event, "SetReward", log); err != nil {
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

// ParseSetReward is a log parse operation binding the contract event 0xc3991418cacd70a010fce559a44fcbb47534ca43bf7543964bdca2639f6e6e5a.
//
// Solidity: event SetReward(address indexed user, uint256 no, uint256 amount)
func (_StakeNFT *StakeNFTFilterer) ParseSetReward(log types.Log) (*StakeNFTSetReward, error) {
	event := new(StakeNFTSetReward)
	if err := _StakeNFT.contract.UnpackLog(event, "SetReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNFTStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the StakeNFT contract.
type StakeNFTStakeIterator struct {
	Event *StakeNFTStake // Event containing the contract specifics and raw log

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
func (it *StakeNFTStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNFTStake)
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
		it.Event = new(StakeNFTStake)
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
func (it *StakeNFTStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNFTStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNFTStake represents a Stake event raised by the StakeNFT contract.
type StakeNFTStake struct {
	User    common.Address
	TokenId *big.Int
	Amount  *big.Int
	K       uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0x6fa2f6021ee337436cdf2b2570285886da18bc06402ce284d296504a0fe8f875.
//
// Solidity: event Stake(address indexed user, uint256 tokenId, uint256 amount, uint8 k)
func (_StakeNFT *StakeNFTFilterer) FilterStake(opts *bind.FilterOpts, user []common.Address) (*StakeNFTStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNFT.contract.FilterLogs(opts, "Stake", userRule)
	if err != nil {
		return nil, err
	}
	return &StakeNFTStakeIterator{contract: _StakeNFT.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0x6fa2f6021ee337436cdf2b2570285886da18bc06402ce284d296504a0fe8f875.
//
// Solidity: event Stake(address indexed user, uint256 tokenId, uint256 amount, uint8 k)
func (_StakeNFT *StakeNFTFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *StakeNFTStake, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNFT.contract.WatchLogs(opts, "Stake", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNFTStake)
				if err := _StakeNFT.contract.UnpackLog(event, "Stake", log); err != nil {
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

// ParseStake is a log parse operation binding the contract event 0x6fa2f6021ee337436cdf2b2570285886da18bc06402ce284d296504a0fe8f875.
//
// Solidity: event Stake(address indexed user, uint256 tokenId, uint256 amount, uint8 k)
func (_StakeNFT *StakeNFTFilterer) ParseStake(log types.Log) (*StakeNFTStake, error) {
	event := new(StakeNFTStake)
	if err := _StakeNFT.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNFTWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the StakeNFT contract.
type StakeNFTWithdrawIterator struct {
	Event *StakeNFTWithdraw // Event containing the contract specifics and raw log

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
func (it *StakeNFTWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNFTWithdraw)
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
		it.Event = new(StakeNFTWithdraw)
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
func (it *StakeNFTWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNFTWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNFTWithdraw represents a Withdraw event raised by the StakeNFT contract.
type StakeNFTWithdraw struct {
	User    common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 tokenId)
func (_StakeNFT *StakeNFTFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*StakeNFTWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNFT.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &StakeNFTWithdrawIterator{contract: _StakeNFT.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 tokenId)
func (_StakeNFT *StakeNFTFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *StakeNFTWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNFT.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNFTWithdraw)
				if err := _StakeNFT.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 tokenId)
func (_StakeNFT *StakeNFTFilterer) ParseWithdraw(log types.Log) (*StakeNFTWithdraw, error) {
	event := new(StakeNFTWithdraw)
	if err := _StakeNFT.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNFTWithdrawAndMergeIterator is returned from FilterWithdrawAndMerge and is used to iterate over the raw logs and unpacked data for WithdrawAndMerge events raised by the StakeNFT contract.
type StakeNFTWithdrawAndMergeIterator struct {
	Event *StakeNFTWithdrawAndMerge // Event containing the contract specifics and raw log

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
func (it *StakeNFTWithdrawAndMergeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNFTWithdrawAndMerge)
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
		it.Event = new(StakeNFTWithdrawAndMerge)
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
func (it *StakeNFTWithdrawAndMergeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNFTWithdrawAndMergeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNFTWithdrawAndMerge represents a WithdrawAndMerge event raised by the StakeNFT contract.
type StakeNFTWithdrawAndMerge struct {
	User      common.Address
	Liquidity *big.Int
	TokenId   *big.Int
	MergeId   *big.Int
	Token0Fee *big.Int
	Token1Fee *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndMerge is a free log retrieval operation binding the contract event 0x9292b1de86834b12fcdc5d8fc52cca04b266ebc75a61d49888ec43218a7e6029.
//
// Solidity: event WithdrawAndMerge(address indexed user, uint128 liquidity, uint256 tokenId, uint256 mergeId, uint256 token0Fee, uint256 token1Fee)
func (_StakeNFT *StakeNFTFilterer) FilterWithdrawAndMerge(opts *bind.FilterOpts, user []common.Address) (*StakeNFTWithdrawAndMergeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNFT.contract.FilterLogs(opts, "WithdrawAndMerge", userRule)
	if err != nil {
		return nil, err
	}
	return &StakeNFTWithdrawAndMergeIterator{contract: _StakeNFT.contract, event: "WithdrawAndMerge", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndMerge is a free log subscription operation binding the contract event 0x9292b1de86834b12fcdc5d8fc52cca04b266ebc75a61d49888ec43218a7e6029.
//
// Solidity: event WithdrawAndMerge(address indexed user, uint128 liquidity, uint256 tokenId, uint256 mergeId, uint256 token0Fee, uint256 token1Fee)
func (_StakeNFT *StakeNFTFilterer) WatchWithdrawAndMerge(opts *bind.WatchOpts, sink chan<- *StakeNFTWithdrawAndMerge, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNFT.contract.WatchLogs(opts, "WithdrawAndMerge", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNFTWithdrawAndMerge)
				if err := _StakeNFT.contract.UnpackLog(event, "WithdrawAndMerge", log); err != nil {
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

// ParseWithdrawAndMerge is a log parse operation binding the contract event 0x9292b1de86834b12fcdc5d8fc52cca04b266ebc75a61d49888ec43218a7e6029.
//
// Solidity: event WithdrawAndMerge(address indexed user, uint128 liquidity, uint256 tokenId, uint256 mergeId, uint256 token0Fee, uint256 token1Fee)
func (_StakeNFT *StakeNFTFilterer) ParseWithdrawAndMerge(log types.Log) (*StakeNFTWithdrawAndMerge, error) {
	event := new(StakeNFTWithdrawAndMerge)
	if err := _StakeNFT.contract.UnpackLog(event, "WithdrawAndMerge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
