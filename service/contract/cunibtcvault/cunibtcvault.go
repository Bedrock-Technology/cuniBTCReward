// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cunibtcvault

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

// CunibtcvaultMetaData contains all meta data concerning the Cunibtcvault contract.
var CunibtcvaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"CapSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatePeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupPeriod\",\"type\":\"uint256\"}],\"name\":\"PeriodSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StartService\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StopService\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"}],\"name\":\"TargetAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"}],\"name\":\"TargetDenied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"}],\"name\":\"TokenAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"}],\"name\":\"TokenDenied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"}],\"name\":\"TokenPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"}],\"name\":\"TokenUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"TotalSupplySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXCHANGE_RATE_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"}],\"name\":\"allowTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_token\",\"type\":\"address[]\"}],\"name\":\"allowToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedTargetList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedTokenList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cuniBTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"}],\"name\":\"denyTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_token\",\"type\":\"address[]\"}],\"name\":\"denyToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cuniBTC\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLockupPeriod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOperatePeriod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockupPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outOfService\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"pauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodRemain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"setCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_operatePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockupPeriod\",\"type\":\"uint256\"}],\"name\":\"setPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"name\":\"setTotalSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startGenesis\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"unpauseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// CunibtcvaultABI is the input ABI used to generate the binding from.
// Deprecated: Use CunibtcvaultMetaData.ABI instead.
var CunibtcvaultABI = CunibtcvaultMetaData.ABI

// Cunibtcvault is an auto generated Go binding around an Ethereum contract.
type Cunibtcvault struct {
	CunibtcvaultCaller     // Read-only binding to the contract
	CunibtcvaultTransactor // Write-only binding to the contract
	CunibtcvaultFilterer   // Log filterer for contract events
}

// CunibtcvaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type CunibtcvaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CunibtcvaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CunibtcvaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CunibtcvaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CunibtcvaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CunibtcvaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CunibtcvaultSession struct {
	Contract     *Cunibtcvault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CunibtcvaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CunibtcvaultCallerSession struct {
	Contract *CunibtcvaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CunibtcvaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CunibtcvaultTransactorSession struct {
	Contract     *CunibtcvaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CunibtcvaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type CunibtcvaultRaw struct {
	Contract *Cunibtcvault // Generic contract binding to access the raw methods on
}

// CunibtcvaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CunibtcvaultCallerRaw struct {
	Contract *CunibtcvaultCaller // Generic read-only contract binding to access the raw methods on
}

// CunibtcvaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CunibtcvaultTransactorRaw struct {
	Contract *CunibtcvaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCunibtcvault creates a new instance of Cunibtcvault, bound to a specific deployed contract.
func NewCunibtcvault(address common.Address, backend bind.ContractBackend) (*Cunibtcvault, error) {
	contract, err := bindCunibtcvault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cunibtcvault{CunibtcvaultCaller: CunibtcvaultCaller{contract: contract}, CunibtcvaultTransactor: CunibtcvaultTransactor{contract: contract}, CunibtcvaultFilterer: CunibtcvaultFilterer{contract: contract}}, nil
}

// NewCunibtcvaultCaller creates a new read-only instance of Cunibtcvault, bound to a specific deployed contract.
func NewCunibtcvaultCaller(address common.Address, caller bind.ContractCaller) (*CunibtcvaultCaller, error) {
	contract, err := bindCunibtcvault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultCaller{contract: contract}, nil
}

// NewCunibtcvaultTransactor creates a new write-only instance of Cunibtcvault, bound to a specific deployed contract.
func NewCunibtcvaultTransactor(address common.Address, transactor bind.ContractTransactor) (*CunibtcvaultTransactor, error) {
	contract, err := bindCunibtcvault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultTransactor{contract: contract}, nil
}

// NewCunibtcvaultFilterer creates a new log filterer instance of Cunibtcvault, bound to a specific deployed contract.
func NewCunibtcvaultFilterer(address common.Address, filterer bind.ContractFilterer) (*CunibtcvaultFilterer, error) {
	contract, err := bindCunibtcvault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultFilterer{contract: contract}, nil
}

// bindCunibtcvault binds a generic wrapper to an already deployed contract.
func bindCunibtcvault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CunibtcvaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cunibtcvault *CunibtcvaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cunibtcvault.Contract.CunibtcvaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cunibtcvault *CunibtcvaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.CunibtcvaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cunibtcvault *CunibtcvaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.CunibtcvaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cunibtcvault *CunibtcvaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cunibtcvault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cunibtcvault *CunibtcvaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cunibtcvault *CunibtcvaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Cunibtcvault *CunibtcvaultCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Cunibtcvault *CunibtcvaultSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Cunibtcvault.Contract.DEFAULTADMINROLE(&_Cunibtcvault.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Cunibtcvault *CunibtcvaultCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Cunibtcvault.Contract.DEFAULTADMINROLE(&_Cunibtcvault.CallOpts)
}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCaller) EXCHANGERATEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "EXCHANGE_RATE_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultSession) EXCHANGERATEBASE() (*big.Int, error) {
	return _Cunibtcvault.Contract.EXCHANGERATEBASE(&_Cunibtcvault.CallOpts)
}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCallerSession) EXCHANGERATEBASE() (*big.Int, error) {
	return _Cunibtcvault.Contract.EXCHANGERATEBASE(&_Cunibtcvault.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Cunibtcvault *CunibtcvaultCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Cunibtcvault *CunibtcvaultSession) OPERATORROLE() ([32]byte, error) {
	return _Cunibtcvault.Contract.OPERATORROLE(&_Cunibtcvault.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Cunibtcvault *CunibtcvaultCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Cunibtcvault.Contract.OPERATORROLE(&_Cunibtcvault.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Cunibtcvault *CunibtcvaultCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Cunibtcvault *CunibtcvaultSession) PAUSERROLE() ([32]byte, error) {
	return _Cunibtcvault.Contract.PAUSERROLE(&_Cunibtcvault.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Cunibtcvault *CunibtcvaultCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Cunibtcvault.Contract.PAUSERROLE(&_Cunibtcvault.CallOpts)
}

// AllowedTargetList is a free data retrieval call binding the contract method 0xc6de4dab.
//
// Solidity: function allowedTargetList(address ) view returns(bool)
func (_Cunibtcvault *CunibtcvaultCaller) AllowedTargetList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "allowedTargetList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedTargetList is a free data retrieval call binding the contract method 0xc6de4dab.
//
// Solidity: function allowedTargetList(address ) view returns(bool)
func (_Cunibtcvault *CunibtcvaultSession) AllowedTargetList(arg0 common.Address) (bool, error) {
	return _Cunibtcvault.Contract.AllowedTargetList(&_Cunibtcvault.CallOpts, arg0)
}

// AllowedTargetList is a free data retrieval call binding the contract method 0xc6de4dab.
//
// Solidity: function allowedTargetList(address ) view returns(bool)
func (_Cunibtcvault *CunibtcvaultCallerSession) AllowedTargetList(arg0 common.Address) (bool, error) {
	return _Cunibtcvault.Contract.AllowedTargetList(&_Cunibtcvault.CallOpts, arg0)
}

// AllowedTokenList is a free data retrieval call binding the contract method 0x71da85f9.
//
// Solidity: function allowedTokenList(address ) view returns(bool)
func (_Cunibtcvault *CunibtcvaultCaller) AllowedTokenList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "allowedTokenList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedTokenList is a free data retrieval call binding the contract method 0x71da85f9.
//
// Solidity: function allowedTokenList(address ) view returns(bool)
func (_Cunibtcvault *CunibtcvaultSession) AllowedTokenList(arg0 common.Address) (bool, error) {
	return _Cunibtcvault.Contract.AllowedTokenList(&_Cunibtcvault.CallOpts, arg0)
}

// AllowedTokenList is a free data retrieval call binding the contract method 0x71da85f9.
//
// Solidity: function allowedTokenList(address ) view returns(bool)
func (_Cunibtcvault *CunibtcvaultCallerSession) AllowedTokenList(arg0 common.Address) (bool, error) {
	return _Cunibtcvault.Contract.AllowedTokenList(&_Cunibtcvault.CallOpts, arg0)
}

// CuniBTC is a free data retrieval call binding the contract method 0x5ab99b63.
//
// Solidity: function cuniBTC() view returns(address)
func (_Cunibtcvault *CunibtcvaultCaller) CuniBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "cuniBTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CuniBTC is a free data retrieval call binding the contract method 0x5ab99b63.
//
// Solidity: function cuniBTC() view returns(address)
func (_Cunibtcvault *CunibtcvaultSession) CuniBTC() (common.Address, error) {
	return _Cunibtcvault.Contract.CuniBTC(&_Cunibtcvault.CallOpts)
}

// CuniBTC is a free data retrieval call binding the contract method 0x5ab99b63.
//
// Solidity: function cuniBTC() view returns(address)
func (_Cunibtcvault *CunibtcvaultCallerSession) CuniBTC() (common.Address, error) {
	return _Cunibtcvault.Contract.CuniBTC(&_Cunibtcvault.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Cunibtcvault *CunibtcvaultCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Cunibtcvault *CunibtcvaultSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Cunibtcvault.Contract.GetRoleAdmin(&_Cunibtcvault.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Cunibtcvault *CunibtcvaultCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Cunibtcvault.Contract.GetRoleAdmin(&_Cunibtcvault.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Cunibtcvault *CunibtcvaultCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Cunibtcvault *CunibtcvaultSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Cunibtcvault.Contract.HasRole(&_Cunibtcvault.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Cunibtcvault *CunibtcvaultCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Cunibtcvault.Contract.HasRole(&_Cunibtcvault.CallOpts, role, account)
}

// IsLockupPeriod is a free data retrieval call binding the contract method 0x07aa621b.
//
// Solidity: function isLockupPeriod() view returns(bool)
func (_Cunibtcvault *CunibtcvaultCaller) IsLockupPeriod(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "isLockupPeriod")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLockupPeriod is a free data retrieval call binding the contract method 0x07aa621b.
//
// Solidity: function isLockupPeriod() view returns(bool)
func (_Cunibtcvault *CunibtcvaultSession) IsLockupPeriod() (bool, error) {
	return _Cunibtcvault.Contract.IsLockupPeriod(&_Cunibtcvault.CallOpts)
}

// IsLockupPeriod is a free data retrieval call binding the contract method 0x07aa621b.
//
// Solidity: function isLockupPeriod() view returns(bool)
func (_Cunibtcvault *CunibtcvaultCallerSession) IsLockupPeriod() (bool, error) {
	return _Cunibtcvault.Contract.IsLockupPeriod(&_Cunibtcvault.CallOpts)
}

// IsOperatePeriod is a free data retrieval call binding the contract method 0xae25b72a.
//
// Solidity: function isOperatePeriod() view returns(bool)
func (_Cunibtcvault *CunibtcvaultCaller) IsOperatePeriod(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "isOperatePeriod")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatePeriod is a free data retrieval call binding the contract method 0xae25b72a.
//
// Solidity: function isOperatePeriod() view returns(bool)
func (_Cunibtcvault *CunibtcvaultSession) IsOperatePeriod() (bool, error) {
	return _Cunibtcvault.Contract.IsOperatePeriod(&_Cunibtcvault.CallOpts)
}

// IsOperatePeriod is a free data retrieval call binding the contract method 0xae25b72a.
//
// Solidity: function isOperatePeriod() view returns(bool)
func (_Cunibtcvault *CunibtcvaultCallerSession) IsOperatePeriod() (bool, error) {
	return _Cunibtcvault.Contract.IsOperatePeriod(&_Cunibtcvault.CallOpts)
}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCaller) LockupPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "lockupPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultSession) LockupPeriod() (*big.Int, error) {
	return _Cunibtcvault.Contract.LockupPeriod(&_Cunibtcvault.CallOpts)
}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCallerSession) LockupPeriod() (*big.Int, error) {
	return _Cunibtcvault.Contract.LockupPeriod(&_Cunibtcvault.CallOpts)
}

// OperatePeriod is a free data retrieval call binding the contract method 0x9fdedd9e.
//
// Solidity: function operatePeriod() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCaller) OperatePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "operatePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatePeriod is a free data retrieval call binding the contract method 0x9fdedd9e.
//
// Solidity: function operatePeriod() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultSession) OperatePeriod() (*big.Int, error) {
	return _Cunibtcvault.Contract.OperatePeriod(&_Cunibtcvault.CallOpts)
}

// OperatePeriod is a free data retrieval call binding the contract method 0x9fdedd9e.
//
// Solidity: function operatePeriod() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCallerSession) OperatePeriod() (*big.Int, error) {
	return _Cunibtcvault.Contract.OperatePeriod(&_Cunibtcvault.CallOpts)
}

// OutOfService is a free data retrieval call binding the contract method 0x30d3c2db.
//
// Solidity: function outOfService() view returns(bool)
func (_Cunibtcvault *CunibtcvaultCaller) OutOfService(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "outOfService")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OutOfService is a free data retrieval call binding the contract method 0x30d3c2db.
//
// Solidity: function outOfService() view returns(bool)
func (_Cunibtcvault *CunibtcvaultSession) OutOfService() (bool, error) {
	return _Cunibtcvault.Contract.OutOfService(&_Cunibtcvault.CallOpts)
}

// OutOfService is a free data retrieval call binding the contract method 0x30d3c2db.
//
// Solidity: function outOfService() view returns(bool)
func (_Cunibtcvault *CunibtcvaultCallerSession) OutOfService() (bool, error) {
	return _Cunibtcvault.Contract.OutOfService(&_Cunibtcvault.CallOpts)
}

// PeriodRemain is a free data retrieval call binding the contract method 0x9a26ccae.
//
// Solidity: function periodRemain() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCaller) PeriodRemain(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "periodRemain")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodRemain is a free data retrieval call binding the contract method 0x9a26ccae.
//
// Solidity: function periodRemain() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultSession) PeriodRemain() (*big.Int, error) {
	return _Cunibtcvault.Contract.PeriodRemain(&_Cunibtcvault.CallOpts)
}

// PeriodRemain is a free data retrieval call binding the contract method 0x9a26ccae.
//
// Solidity: function periodRemain() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCallerSession) PeriodRemain() (*big.Int, error) {
	return _Cunibtcvault.Contract.PeriodRemain(&_Cunibtcvault.CallOpts)
}

// StartGenesis is a free data retrieval call binding the contract method 0xf09e906a.
//
// Solidity: function startGenesis() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCaller) StartGenesis(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "startGenesis")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartGenesis is a free data retrieval call binding the contract method 0xf09e906a.
//
// Solidity: function startGenesis() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultSession) StartGenesis() (*big.Int, error) {
	return _Cunibtcvault.Contract.StartGenesis(&_Cunibtcvault.CallOpts)
}

// StartGenesis is a free data retrieval call binding the contract method 0xf09e906a.
//
// Solidity: function startGenesis() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCallerSession) StartGenesis() (*big.Int, error) {
	return _Cunibtcvault.Contract.StartGenesis(&_Cunibtcvault.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cunibtcvault *CunibtcvaultCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cunibtcvault *CunibtcvaultSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cunibtcvault.Contract.SupportsInterface(&_Cunibtcvault.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Cunibtcvault *CunibtcvaultCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Cunibtcvault.Contract.SupportsInterface(&_Cunibtcvault.CallOpts, interfaceId)
}

// TokenCaps is a free data retrieval call binding the contract method 0xe0f3871f.
//
// Solidity: function tokenCaps(address ) view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCaller) TokenCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "tokenCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCaps is a free data retrieval call binding the contract method 0xe0f3871f.
//
// Solidity: function tokenCaps(address ) view returns(uint256)
func (_Cunibtcvault *CunibtcvaultSession) TokenCaps(arg0 common.Address) (*big.Int, error) {
	return _Cunibtcvault.Contract.TokenCaps(&_Cunibtcvault.CallOpts, arg0)
}

// TokenCaps is a free data retrieval call binding the contract method 0xe0f3871f.
//
// Solidity: function tokenCaps(address ) view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCallerSession) TokenCaps(arg0 common.Address) (*big.Int, error) {
	return _Cunibtcvault.Contract.TokenCaps(&_Cunibtcvault.CallOpts, arg0)
}

// TokenMinted is a free data retrieval call binding the contract method 0x656000a8.
//
// Solidity: function tokenMinted(address ) view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCaller) TokenMinted(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "tokenMinted", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMinted is a free data retrieval call binding the contract method 0x656000a8.
//
// Solidity: function tokenMinted(address ) view returns(uint256)
func (_Cunibtcvault *CunibtcvaultSession) TokenMinted(arg0 common.Address) (*big.Int, error) {
	return _Cunibtcvault.Contract.TokenMinted(&_Cunibtcvault.CallOpts, arg0)
}

// TokenMinted is a free data retrieval call binding the contract method 0x656000a8.
//
// Solidity: function tokenMinted(address ) view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCallerSession) TokenMinted(arg0 common.Address) (*big.Int, error) {
	return _Cunibtcvault.Contract.TokenMinted(&_Cunibtcvault.CallOpts, arg0)
}

// TokenPaused is a free data retrieval call binding the contract method 0x8995881d.
//
// Solidity: function tokenPaused(address ) view returns(bool)
func (_Cunibtcvault *CunibtcvaultCaller) TokenPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "tokenPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenPaused is a free data retrieval call binding the contract method 0x8995881d.
//
// Solidity: function tokenPaused(address ) view returns(bool)
func (_Cunibtcvault *CunibtcvaultSession) TokenPaused(arg0 common.Address) (bool, error) {
	return _Cunibtcvault.Contract.TokenPaused(&_Cunibtcvault.CallOpts, arg0)
}

// TokenPaused is a free data retrieval call binding the contract method 0x8995881d.
//
// Solidity: function tokenPaused(address ) view returns(bool)
func (_Cunibtcvault *CunibtcvaultCallerSession) TokenPaused(arg0 common.Address) (bool, error) {
	return _Cunibtcvault.Contract.TokenPaused(&_Cunibtcvault.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cunibtcvault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultSession) TotalSupply() (*big.Int, error) {
	return _Cunibtcvault.Contract.TotalSupply(&_Cunibtcvault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cunibtcvault *CunibtcvaultCallerSession) TotalSupply() (*big.Int, error) {
	return _Cunibtcvault.Contract.TotalSupply(&_Cunibtcvault.CallOpts)
}

// AllowTarget is a paid mutator transaction binding the contract method 0xd30f0bec.
//
// Solidity: function allowTarget(address[] _targets) returns()
func (_Cunibtcvault *CunibtcvaultTransactor) AllowTarget(opts *bind.TransactOpts, _targets []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "allowTarget", _targets)
}

// AllowTarget is a paid mutator transaction binding the contract method 0xd30f0bec.
//
// Solidity: function allowTarget(address[] _targets) returns()
func (_Cunibtcvault *CunibtcvaultSession) AllowTarget(_targets []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.AllowTarget(&_Cunibtcvault.TransactOpts, _targets)
}

// AllowTarget is a paid mutator transaction binding the contract method 0xd30f0bec.
//
// Solidity: function allowTarget(address[] _targets) returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) AllowTarget(_targets []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.AllowTarget(&_Cunibtcvault.TransactOpts, _targets)
}

// AllowToken is a paid mutator transaction binding the contract method 0x136f2bf4.
//
// Solidity: function allowToken(address[] _token) returns()
func (_Cunibtcvault *CunibtcvaultTransactor) AllowToken(opts *bind.TransactOpts, _token []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "allowToken", _token)
}

// AllowToken is a paid mutator transaction binding the contract method 0x136f2bf4.
//
// Solidity: function allowToken(address[] _token) returns()
func (_Cunibtcvault *CunibtcvaultSession) AllowToken(_token []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.AllowToken(&_Cunibtcvault.TransactOpts, _token)
}

// AllowToken is a paid mutator transaction binding the contract method 0x136f2bf4.
//
// Solidity: function allowToken(address[] _token) returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) AllowToken(_token []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.AllowToken(&_Cunibtcvault.TransactOpts, _token)
}

// DenyTarget is a paid mutator transaction binding the contract method 0xb588dd46.
//
// Solidity: function denyTarget(address[] _targets) returns()
func (_Cunibtcvault *CunibtcvaultTransactor) DenyTarget(opts *bind.TransactOpts, _targets []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "denyTarget", _targets)
}

// DenyTarget is a paid mutator transaction binding the contract method 0xb588dd46.
//
// Solidity: function denyTarget(address[] _targets) returns()
func (_Cunibtcvault *CunibtcvaultSession) DenyTarget(_targets []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.DenyTarget(&_Cunibtcvault.TransactOpts, _targets)
}

// DenyTarget is a paid mutator transaction binding the contract method 0xb588dd46.
//
// Solidity: function denyTarget(address[] _targets) returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) DenyTarget(_targets []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.DenyTarget(&_Cunibtcvault.TransactOpts, _targets)
}

// DenyToken is a paid mutator transaction binding the contract method 0xb81766da.
//
// Solidity: function denyToken(address[] _token) returns()
func (_Cunibtcvault *CunibtcvaultTransactor) DenyToken(opts *bind.TransactOpts, _token []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "denyToken", _token)
}

// DenyToken is a paid mutator transaction binding the contract method 0xb81766da.
//
// Solidity: function denyToken(address[] _token) returns()
func (_Cunibtcvault *CunibtcvaultSession) DenyToken(_token []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.DenyToken(&_Cunibtcvault.TransactOpts, _token)
}

// DenyToken is a paid mutator transaction binding the contract method 0xb81766da.
//
// Solidity: function denyToken(address[] _token) returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) DenyToken(_token []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.DenyToken(&_Cunibtcvault.TransactOpts, _token)
}

// Execute is a paid mutator transaction binding the contract method 0xa04a0908.
//
// Solidity: function execute(address _target, bytes _data, uint256 _value) returns(bytes)
func (_Cunibtcvault *CunibtcvaultTransactor) Execute(opts *bind.TransactOpts, _target common.Address, _data []byte, _value *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "execute", _target, _data, _value)
}

// Execute is a paid mutator transaction binding the contract method 0xa04a0908.
//
// Solidity: function execute(address _target, bytes _data, uint256 _value) returns(bytes)
func (_Cunibtcvault *CunibtcvaultSession) Execute(_target common.Address, _data []byte, _value *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.Execute(&_Cunibtcvault.TransactOpts, _target, _data, _value)
}

// Execute is a paid mutator transaction binding the contract method 0xa04a0908.
//
// Solidity: function execute(address _target, bytes _data, uint256 _value) returns(bytes)
func (_Cunibtcvault *CunibtcvaultTransactorSession) Execute(_target common.Address, _data []byte, _value *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.Execute(&_Cunibtcvault.TransactOpts, _target, _data, _value)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Cunibtcvault *CunibtcvaultTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Cunibtcvault *CunibtcvaultSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.GrantRole(&_Cunibtcvault.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.GrantRole(&_Cunibtcvault.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _defaultAdmin, address _cuniBTC, uint256 _totalSupply) returns()
func (_Cunibtcvault *CunibtcvaultTransactor) Initialize(opts *bind.TransactOpts, _defaultAdmin common.Address, _cuniBTC common.Address, _totalSupply *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "initialize", _defaultAdmin, _cuniBTC, _totalSupply)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _defaultAdmin, address _cuniBTC, uint256 _totalSupply) returns()
func (_Cunibtcvault *CunibtcvaultSession) Initialize(_defaultAdmin common.Address, _cuniBTC common.Address, _totalSupply *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.Initialize(&_Cunibtcvault.TransactOpts, _defaultAdmin, _cuniBTC, _totalSupply)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _defaultAdmin, address _cuniBTC, uint256 _totalSupply) returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) Initialize(_defaultAdmin common.Address, _cuniBTC common.Address, _totalSupply *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.Initialize(&_Cunibtcvault.TransactOpts, _defaultAdmin, _cuniBTC, _totalSupply)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _token, uint256 _amount) returns()
func (_Cunibtcvault *CunibtcvaultTransactor) Mint(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "mint", _token, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _token, uint256 _amount) returns()
func (_Cunibtcvault *CunibtcvaultSession) Mint(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.Mint(&_Cunibtcvault.TransactOpts, _token, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _token, uint256 _amount) returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) Mint(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.Mint(&_Cunibtcvault.TransactOpts, _token, _amount)
}

// PauseToken is a paid mutator transaction binding the contract method 0x0d2c7ea4.
//
// Solidity: function pauseToken(address[] _tokens) returns()
func (_Cunibtcvault *CunibtcvaultTransactor) PauseToken(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "pauseToken", _tokens)
}

// PauseToken is a paid mutator transaction binding the contract method 0x0d2c7ea4.
//
// Solidity: function pauseToken(address[] _tokens) returns()
func (_Cunibtcvault *CunibtcvaultSession) PauseToken(_tokens []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.PauseToken(&_Cunibtcvault.TransactOpts, _tokens)
}

// PauseToken is a paid mutator transaction binding the contract method 0x0d2c7ea4.
//
// Solidity: function pauseToken(address[] _tokens) returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) PauseToken(_tokens []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.PauseToken(&_Cunibtcvault.TransactOpts, _tokens)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Cunibtcvault *CunibtcvaultTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Cunibtcvault *CunibtcvaultSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.RenounceRole(&_Cunibtcvault.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.RenounceRole(&_Cunibtcvault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Cunibtcvault *CunibtcvaultTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Cunibtcvault *CunibtcvaultSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.RevokeRole(&_Cunibtcvault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.RevokeRole(&_Cunibtcvault.TransactOpts, role, account)
}

// SetCap is a paid mutator transaction binding the contract method 0x80ad2cf3.
//
// Solidity: function setCap(address _token, uint256 _cap) returns()
func (_Cunibtcvault *CunibtcvaultTransactor) SetCap(opts *bind.TransactOpts, _token common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "setCap", _token, _cap)
}

// SetCap is a paid mutator transaction binding the contract method 0x80ad2cf3.
//
// Solidity: function setCap(address _token, uint256 _cap) returns()
func (_Cunibtcvault *CunibtcvaultSession) SetCap(_token common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.SetCap(&_Cunibtcvault.TransactOpts, _token, _cap)
}

// SetCap is a paid mutator transaction binding the contract method 0x80ad2cf3.
//
// Solidity: function setCap(address _token, uint256 _cap) returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) SetCap(_token common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.SetCap(&_Cunibtcvault.TransactOpts, _token, _cap)
}

// SetPeriod is a paid mutator transaction binding the contract method 0x8feadcb7.
//
// Solidity: function setPeriod(uint256 _start, uint256 _operatePeriod, uint256 _lockupPeriod) returns()
func (_Cunibtcvault *CunibtcvaultTransactor) SetPeriod(opts *bind.TransactOpts, _start *big.Int, _operatePeriod *big.Int, _lockupPeriod *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "setPeriod", _start, _operatePeriod, _lockupPeriod)
}

// SetPeriod is a paid mutator transaction binding the contract method 0x8feadcb7.
//
// Solidity: function setPeriod(uint256 _start, uint256 _operatePeriod, uint256 _lockupPeriod) returns()
func (_Cunibtcvault *CunibtcvaultSession) SetPeriod(_start *big.Int, _operatePeriod *big.Int, _lockupPeriod *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.SetPeriod(&_Cunibtcvault.TransactOpts, _start, _operatePeriod, _lockupPeriod)
}

// SetPeriod is a paid mutator transaction binding the contract method 0x8feadcb7.
//
// Solidity: function setPeriod(uint256 _start, uint256 _operatePeriod, uint256 _lockupPeriod) returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) SetPeriod(_start *big.Int, _operatePeriod *big.Int, _lockupPeriod *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.SetPeriod(&_Cunibtcvault.TransactOpts, _start, _operatePeriod, _lockupPeriod)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(uint256 _totalSupply) returns()
func (_Cunibtcvault *CunibtcvaultTransactor) SetTotalSupply(opts *bind.TransactOpts, _totalSupply *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "setTotalSupply", _totalSupply)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(uint256 _totalSupply) returns()
func (_Cunibtcvault *CunibtcvaultSession) SetTotalSupply(_totalSupply *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.SetTotalSupply(&_Cunibtcvault.TransactOpts, _totalSupply)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(uint256 _totalSupply) returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) SetTotalSupply(_totalSupply *big.Int) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.SetTotalSupply(&_Cunibtcvault.TransactOpts, _totalSupply)
}

// StartService is a paid mutator transaction binding the contract method 0x789b5771.
//
// Solidity: function startService() returns()
func (_Cunibtcvault *CunibtcvaultTransactor) StartService(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "startService")
}

// StartService is a paid mutator transaction binding the contract method 0x789b5771.
//
// Solidity: function startService() returns()
func (_Cunibtcvault *CunibtcvaultSession) StartService() (*types.Transaction, error) {
	return _Cunibtcvault.Contract.StartService(&_Cunibtcvault.TransactOpts)
}

// StartService is a paid mutator transaction binding the contract method 0x789b5771.
//
// Solidity: function startService() returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) StartService() (*types.Transaction, error) {
	return _Cunibtcvault.Contract.StartService(&_Cunibtcvault.TransactOpts)
}

// StopService is a paid mutator transaction binding the contract method 0x1e54e974.
//
// Solidity: function stopService() returns()
func (_Cunibtcvault *CunibtcvaultTransactor) StopService(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "stopService")
}

// StopService is a paid mutator transaction binding the contract method 0x1e54e974.
//
// Solidity: function stopService() returns()
func (_Cunibtcvault *CunibtcvaultSession) StopService() (*types.Transaction, error) {
	return _Cunibtcvault.Contract.StopService(&_Cunibtcvault.TransactOpts)
}

// StopService is a paid mutator transaction binding the contract method 0x1e54e974.
//
// Solidity: function stopService() returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) StopService() (*types.Transaction, error) {
	return _Cunibtcvault.Contract.StopService(&_Cunibtcvault.TransactOpts)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x277af3a0.
//
// Solidity: function unpauseToken(address[] _tokens) returns()
func (_Cunibtcvault *CunibtcvaultTransactor) UnpauseToken(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.contract.Transact(opts, "unpauseToken", _tokens)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x277af3a0.
//
// Solidity: function unpauseToken(address[] _tokens) returns()
func (_Cunibtcvault *CunibtcvaultSession) UnpauseToken(_tokens []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.UnpauseToken(&_Cunibtcvault.TransactOpts, _tokens)
}

// UnpauseToken is a paid mutator transaction binding the contract method 0x277af3a0.
//
// Solidity: function unpauseToken(address[] _tokens) returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) UnpauseToken(_tokens []common.Address) (*types.Transaction, error) {
	return _Cunibtcvault.Contract.UnpauseToken(&_Cunibtcvault.TransactOpts, _tokens)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cunibtcvault *CunibtcvaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cunibtcvault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cunibtcvault *CunibtcvaultSession) Receive() (*types.Transaction, error) {
	return _Cunibtcvault.Contract.Receive(&_Cunibtcvault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cunibtcvault *CunibtcvaultTransactorSession) Receive() (*types.Transaction, error) {
	return _Cunibtcvault.Contract.Receive(&_Cunibtcvault.TransactOpts)
}

// CunibtcvaultCapSetIterator is returned from FilterCapSet and is used to iterate over the raw logs and unpacked data for CapSet events raised by the Cunibtcvault contract.
type CunibtcvaultCapSetIterator struct {
	Event *CunibtcvaultCapSet // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultCapSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultCapSet)
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
		it.Event = new(CunibtcvaultCapSet)
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
func (it *CunibtcvaultCapSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultCapSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultCapSet represents a CapSet event raised by the Cunibtcvault contract.
type CunibtcvaultCapSet struct {
	Token common.Address
	Cap   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCapSet is a free log retrieval operation binding the contract event 0x1c9182333abe71fddddd188f90cabb1c34182de0b0ac74c87944a196460a2d6c.
//
// Solidity: event CapSet(address token, uint256 cap)
func (_Cunibtcvault *CunibtcvaultFilterer) FilterCapSet(opts *bind.FilterOpts) (*CunibtcvaultCapSetIterator, error) {

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "CapSet")
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultCapSetIterator{contract: _Cunibtcvault.contract, event: "CapSet", logs: logs, sub: sub}, nil
}

// WatchCapSet is a free log subscription operation binding the contract event 0x1c9182333abe71fddddd188f90cabb1c34182de0b0ac74c87944a196460a2d6c.
//
// Solidity: event CapSet(address token, uint256 cap)
func (_Cunibtcvault *CunibtcvaultFilterer) WatchCapSet(opts *bind.WatchOpts, sink chan<- *CunibtcvaultCapSet) (event.Subscription, error) {

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "CapSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultCapSet)
				if err := _Cunibtcvault.contract.UnpackLog(event, "CapSet", log); err != nil {
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

// ParseCapSet is a log parse operation binding the contract event 0x1c9182333abe71fddddd188f90cabb1c34182de0b0ac74c87944a196460a2d6c.
//
// Solidity: event CapSet(address token, uint256 cap)
func (_Cunibtcvault *CunibtcvaultFilterer) ParseCapSet(log types.Log) (*CunibtcvaultCapSet, error) {
	event := new(CunibtcvaultCapSet)
	if err := _Cunibtcvault.contract.UnpackLog(event, "CapSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Cunibtcvault contract.
type CunibtcvaultInitializedIterator struct {
	Event *CunibtcvaultInitialized // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultInitialized)
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
		it.Event = new(CunibtcvaultInitialized)
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
func (it *CunibtcvaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultInitialized represents a Initialized event raised by the Cunibtcvault contract.
type CunibtcvaultInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Cunibtcvault *CunibtcvaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*CunibtcvaultInitializedIterator, error) {

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultInitializedIterator{contract: _Cunibtcvault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Cunibtcvault *CunibtcvaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CunibtcvaultInitialized) (event.Subscription, error) {

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultInitialized)
				if err := _Cunibtcvault.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Cunibtcvault *CunibtcvaultFilterer) ParseInitialized(log types.Log) (*CunibtcvaultInitialized, error) {
	event := new(CunibtcvaultInitialized)
	if err := _Cunibtcvault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the Cunibtcvault contract.
type CunibtcvaultMintedIterator struct {
	Event *CunibtcvaultMinted // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultMinted)
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
		it.Event = new(CunibtcvaultMinted)
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
func (it *CunibtcvaultMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultMinted represents a Minted event raised by the Cunibtcvault contract.
type CunibtcvaultMinted struct {
	Sender common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address sender, address token, uint256 amount)
func (_Cunibtcvault *CunibtcvaultFilterer) FilterMinted(opts *bind.FilterOpts) (*CunibtcvaultMintedIterator, error) {

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultMintedIterator{contract: _Cunibtcvault.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address sender, address token, uint256 amount)
func (_Cunibtcvault *CunibtcvaultFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *CunibtcvaultMinted) (event.Subscription, error) {

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultMinted)
				if err := _Cunibtcvault.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address sender, address token, uint256 amount)
func (_Cunibtcvault *CunibtcvaultFilterer) ParseMinted(log types.Log) (*CunibtcvaultMinted, error) {
	event := new(CunibtcvaultMinted)
	if err := _Cunibtcvault.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultPeriodSetIterator is returned from FilterPeriodSet and is used to iterate over the raw logs and unpacked data for PeriodSet events raised by the Cunibtcvault contract.
type CunibtcvaultPeriodSetIterator struct {
	Event *CunibtcvaultPeriodSet // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultPeriodSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultPeriodSet)
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
		it.Event = new(CunibtcvaultPeriodSet)
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
func (it *CunibtcvaultPeriodSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultPeriodSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultPeriodSet represents a PeriodSet event raised by the Cunibtcvault contract.
type CunibtcvaultPeriodSet struct {
	Start         *big.Int
	OperatePeriod *big.Int
	LockupPeriod  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPeriodSet is a free log retrieval operation binding the contract event 0x883b392e4889d0f8c1208781c80d3995a9935c1312f28f7584e611d00aec1359.
//
// Solidity: event PeriodSet(uint256 start, uint256 operatePeriod, uint256 lockupPeriod)
func (_Cunibtcvault *CunibtcvaultFilterer) FilterPeriodSet(opts *bind.FilterOpts) (*CunibtcvaultPeriodSetIterator, error) {

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "PeriodSet")
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultPeriodSetIterator{contract: _Cunibtcvault.contract, event: "PeriodSet", logs: logs, sub: sub}, nil
}

// WatchPeriodSet is a free log subscription operation binding the contract event 0x883b392e4889d0f8c1208781c80d3995a9935c1312f28f7584e611d00aec1359.
//
// Solidity: event PeriodSet(uint256 start, uint256 operatePeriod, uint256 lockupPeriod)
func (_Cunibtcvault *CunibtcvaultFilterer) WatchPeriodSet(opts *bind.WatchOpts, sink chan<- *CunibtcvaultPeriodSet) (event.Subscription, error) {

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "PeriodSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultPeriodSet)
				if err := _Cunibtcvault.contract.UnpackLog(event, "PeriodSet", log); err != nil {
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

// ParsePeriodSet is a log parse operation binding the contract event 0x883b392e4889d0f8c1208781c80d3995a9935c1312f28f7584e611d00aec1359.
//
// Solidity: event PeriodSet(uint256 start, uint256 operatePeriod, uint256 lockupPeriod)
func (_Cunibtcvault *CunibtcvaultFilterer) ParsePeriodSet(log types.Log) (*CunibtcvaultPeriodSet, error) {
	event := new(CunibtcvaultPeriodSet)
	if err := _Cunibtcvault.contract.UnpackLog(event, "PeriodSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Cunibtcvault contract.
type CunibtcvaultRoleAdminChangedIterator struct {
	Event *CunibtcvaultRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultRoleAdminChanged)
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
		it.Event = new(CunibtcvaultRoleAdminChanged)
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
func (it *CunibtcvaultRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultRoleAdminChanged represents a RoleAdminChanged event raised by the Cunibtcvault contract.
type CunibtcvaultRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Cunibtcvault *CunibtcvaultFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CunibtcvaultRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultRoleAdminChangedIterator{contract: _Cunibtcvault.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Cunibtcvault *CunibtcvaultFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CunibtcvaultRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultRoleAdminChanged)
				if err := _Cunibtcvault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Cunibtcvault *CunibtcvaultFilterer) ParseRoleAdminChanged(log types.Log) (*CunibtcvaultRoleAdminChanged, error) {
	event := new(CunibtcvaultRoleAdminChanged)
	if err := _Cunibtcvault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Cunibtcvault contract.
type CunibtcvaultRoleGrantedIterator struct {
	Event *CunibtcvaultRoleGranted // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultRoleGranted)
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
		it.Event = new(CunibtcvaultRoleGranted)
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
func (it *CunibtcvaultRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultRoleGranted represents a RoleGranted event raised by the Cunibtcvault contract.
type CunibtcvaultRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cunibtcvault *CunibtcvaultFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CunibtcvaultRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultRoleGrantedIterator{contract: _Cunibtcvault.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cunibtcvault *CunibtcvaultFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CunibtcvaultRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultRoleGranted)
				if err := _Cunibtcvault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cunibtcvault *CunibtcvaultFilterer) ParseRoleGranted(log types.Log) (*CunibtcvaultRoleGranted, error) {
	event := new(CunibtcvaultRoleGranted)
	if err := _Cunibtcvault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Cunibtcvault contract.
type CunibtcvaultRoleRevokedIterator struct {
	Event *CunibtcvaultRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultRoleRevoked)
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
		it.Event = new(CunibtcvaultRoleRevoked)
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
func (it *CunibtcvaultRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultRoleRevoked represents a RoleRevoked event raised by the Cunibtcvault contract.
type CunibtcvaultRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cunibtcvault *CunibtcvaultFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CunibtcvaultRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultRoleRevokedIterator{contract: _Cunibtcvault.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cunibtcvault *CunibtcvaultFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CunibtcvaultRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultRoleRevoked)
				if err := _Cunibtcvault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Cunibtcvault *CunibtcvaultFilterer) ParseRoleRevoked(log types.Log) (*CunibtcvaultRoleRevoked, error) {
	event := new(CunibtcvaultRoleRevoked)
	if err := _Cunibtcvault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultStartServiceIterator is returned from FilterStartService and is used to iterate over the raw logs and unpacked data for StartService events raised by the Cunibtcvault contract.
type CunibtcvaultStartServiceIterator struct {
	Event *CunibtcvaultStartService // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultStartServiceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultStartService)
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
		it.Event = new(CunibtcvaultStartService)
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
func (it *CunibtcvaultStartServiceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultStartServiceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultStartService represents a StartService event raised by the Cunibtcvault contract.
type CunibtcvaultStartService struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStartService is a free log retrieval operation binding the contract event 0xcff364222ca8c055192e4db077cf0091c6c013260449218d7ceb39149bffe2e9.
//
// Solidity: event StartService()
func (_Cunibtcvault *CunibtcvaultFilterer) FilterStartService(opts *bind.FilterOpts) (*CunibtcvaultStartServiceIterator, error) {

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "StartService")
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultStartServiceIterator{contract: _Cunibtcvault.contract, event: "StartService", logs: logs, sub: sub}, nil
}

// WatchStartService is a free log subscription operation binding the contract event 0xcff364222ca8c055192e4db077cf0091c6c013260449218d7ceb39149bffe2e9.
//
// Solidity: event StartService()
func (_Cunibtcvault *CunibtcvaultFilterer) WatchStartService(opts *bind.WatchOpts, sink chan<- *CunibtcvaultStartService) (event.Subscription, error) {

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "StartService")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultStartService)
				if err := _Cunibtcvault.contract.UnpackLog(event, "StartService", log); err != nil {
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

// ParseStartService is a log parse operation binding the contract event 0xcff364222ca8c055192e4db077cf0091c6c013260449218d7ceb39149bffe2e9.
//
// Solidity: event StartService()
func (_Cunibtcvault *CunibtcvaultFilterer) ParseStartService(log types.Log) (*CunibtcvaultStartService, error) {
	event := new(CunibtcvaultStartService)
	if err := _Cunibtcvault.contract.UnpackLog(event, "StartService", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultStopServiceIterator is returned from FilterStopService and is used to iterate over the raw logs and unpacked data for StopService events raised by the Cunibtcvault contract.
type CunibtcvaultStopServiceIterator struct {
	Event *CunibtcvaultStopService // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultStopServiceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultStopService)
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
		it.Event = new(CunibtcvaultStopService)
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
func (it *CunibtcvaultStopServiceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultStopServiceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultStopService represents a StopService event raised by the Cunibtcvault contract.
type CunibtcvaultStopService struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopService is a free log retrieval operation binding the contract event 0xc06db5c826a32f4aec9251d2fad227a2ff4c899fcb2248170d460fdab8e05ff2.
//
// Solidity: event StopService()
func (_Cunibtcvault *CunibtcvaultFilterer) FilterStopService(opts *bind.FilterOpts) (*CunibtcvaultStopServiceIterator, error) {

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "StopService")
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultStopServiceIterator{contract: _Cunibtcvault.contract, event: "StopService", logs: logs, sub: sub}, nil
}

// WatchStopService is a free log subscription operation binding the contract event 0xc06db5c826a32f4aec9251d2fad227a2ff4c899fcb2248170d460fdab8e05ff2.
//
// Solidity: event StopService()
func (_Cunibtcvault *CunibtcvaultFilterer) WatchStopService(opts *bind.WatchOpts, sink chan<- *CunibtcvaultStopService) (event.Subscription, error) {

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "StopService")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultStopService)
				if err := _Cunibtcvault.contract.UnpackLog(event, "StopService", log); err != nil {
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

// ParseStopService is a log parse operation binding the contract event 0xc06db5c826a32f4aec9251d2fad227a2ff4c899fcb2248170d460fdab8e05ff2.
//
// Solidity: event StopService()
func (_Cunibtcvault *CunibtcvaultFilterer) ParseStopService(log types.Log) (*CunibtcvaultStopService, error) {
	event := new(CunibtcvaultStopService)
	if err := _Cunibtcvault.contract.UnpackLog(event, "StopService", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultTargetAllowedIterator is returned from FilterTargetAllowed and is used to iterate over the raw logs and unpacked data for TargetAllowed events raised by the Cunibtcvault contract.
type CunibtcvaultTargetAllowedIterator struct {
	Event *CunibtcvaultTargetAllowed // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultTargetAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultTargetAllowed)
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
		it.Event = new(CunibtcvaultTargetAllowed)
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
func (it *CunibtcvaultTargetAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultTargetAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultTargetAllowed represents a TargetAllowed event raised by the Cunibtcvault contract.
type CunibtcvaultTargetAllowed struct {
	Token []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTargetAllowed is a free log retrieval operation binding the contract event 0x3e06a05d4c3951224df6e48c13ba18c9858f48529bc1b25dc6f6b6e7e62d3beb.
//
// Solidity: event TargetAllowed(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) FilterTargetAllowed(opts *bind.FilterOpts) (*CunibtcvaultTargetAllowedIterator, error) {

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "TargetAllowed")
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultTargetAllowedIterator{contract: _Cunibtcvault.contract, event: "TargetAllowed", logs: logs, sub: sub}, nil
}

// WatchTargetAllowed is a free log subscription operation binding the contract event 0x3e06a05d4c3951224df6e48c13ba18c9858f48529bc1b25dc6f6b6e7e62d3beb.
//
// Solidity: event TargetAllowed(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) WatchTargetAllowed(opts *bind.WatchOpts, sink chan<- *CunibtcvaultTargetAllowed) (event.Subscription, error) {

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "TargetAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultTargetAllowed)
				if err := _Cunibtcvault.contract.UnpackLog(event, "TargetAllowed", log); err != nil {
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

// ParseTargetAllowed is a log parse operation binding the contract event 0x3e06a05d4c3951224df6e48c13ba18c9858f48529bc1b25dc6f6b6e7e62d3beb.
//
// Solidity: event TargetAllowed(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) ParseTargetAllowed(log types.Log) (*CunibtcvaultTargetAllowed, error) {
	event := new(CunibtcvaultTargetAllowed)
	if err := _Cunibtcvault.contract.UnpackLog(event, "TargetAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultTargetDeniedIterator is returned from FilterTargetDenied and is used to iterate over the raw logs and unpacked data for TargetDenied events raised by the Cunibtcvault contract.
type CunibtcvaultTargetDeniedIterator struct {
	Event *CunibtcvaultTargetDenied // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultTargetDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultTargetDenied)
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
		it.Event = new(CunibtcvaultTargetDenied)
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
func (it *CunibtcvaultTargetDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultTargetDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultTargetDenied represents a TargetDenied event raised by the Cunibtcvault contract.
type CunibtcvaultTargetDenied struct {
	Token []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTargetDenied is a free log retrieval operation binding the contract event 0x5242d62d33cff3cfb1444011cd3b283c67cf174a088c07c3c761aab2c0b80a38.
//
// Solidity: event TargetDenied(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) FilterTargetDenied(opts *bind.FilterOpts) (*CunibtcvaultTargetDeniedIterator, error) {

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "TargetDenied")
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultTargetDeniedIterator{contract: _Cunibtcvault.contract, event: "TargetDenied", logs: logs, sub: sub}, nil
}

// WatchTargetDenied is a free log subscription operation binding the contract event 0x5242d62d33cff3cfb1444011cd3b283c67cf174a088c07c3c761aab2c0b80a38.
//
// Solidity: event TargetDenied(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) WatchTargetDenied(opts *bind.WatchOpts, sink chan<- *CunibtcvaultTargetDenied) (event.Subscription, error) {

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "TargetDenied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultTargetDenied)
				if err := _Cunibtcvault.contract.UnpackLog(event, "TargetDenied", log); err != nil {
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

// ParseTargetDenied is a log parse operation binding the contract event 0x5242d62d33cff3cfb1444011cd3b283c67cf174a088c07c3c761aab2c0b80a38.
//
// Solidity: event TargetDenied(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) ParseTargetDenied(log types.Log) (*CunibtcvaultTargetDenied, error) {
	event := new(CunibtcvaultTargetDenied)
	if err := _Cunibtcvault.contract.UnpackLog(event, "TargetDenied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultTokenAllowedIterator is returned from FilterTokenAllowed and is used to iterate over the raw logs and unpacked data for TokenAllowed events raised by the Cunibtcvault contract.
type CunibtcvaultTokenAllowedIterator struct {
	Event *CunibtcvaultTokenAllowed // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultTokenAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultTokenAllowed)
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
		it.Event = new(CunibtcvaultTokenAllowed)
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
func (it *CunibtcvaultTokenAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultTokenAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultTokenAllowed represents a TokenAllowed event raised by the Cunibtcvault contract.
type CunibtcvaultTokenAllowed struct {
	Token []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAllowed is a free log retrieval operation binding the contract event 0x272dfe25b5b0c656af0461e366b26532af938a45b7c4946ff15be2bf78ba8e88.
//
// Solidity: event TokenAllowed(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) FilterTokenAllowed(opts *bind.FilterOpts) (*CunibtcvaultTokenAllowedIterator, error) {

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "TokenAllowed")
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultTokenAllowedIterator{contract: _Cunibtcvault.contract, event: "TokenAllowed", logs: logs, sub: sub}, nil
}

// WatchTokenAllowed is a free log subscription operation binding the contract event 0x272dfe25b5b0c656af0461e366b26532af938a45b7c4946ff15be2bf78ba8e88.
//
// Solidity: event TokenAllowed(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) WatchTokenAllowed(opts *bind.WatchOpts, sink chan<- *CunibtcvaultTokenAllowed) (event.Subscription, error) {

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "TokenAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultTokenAllowed)
				if err := _Cunibtcvault.contract.UnpackLog(event, "TokenAllowed", log); err != nil {
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

// ParseTokenAllowed is a log parse operation binding the contract event 0x272dfe25b5b0c656af0461e366b26532af938a45b7c4946ff15be2bf78ba8e88.
//
// Solidity: event TokenAllowed(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) ParseTokenAllowed(log types.Log) (*CunibtcvaultTokenAllowed, error) {
	event := new(CunibtcvaultTokenAllowed)
	if err := _Cunibtcvault.contract.UnpackLog(event, "TokenAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultTokenDeniedIterator is returned from FilterTokenDenied and is used to iterate over the raw logs and unpacked data for TokenDenied events raised by the Cunibtcvault contract.
type CunibtcvaultTokenDeniedIterator struct {
	Event *CunibtcvaultTokenDenied // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultTokenDeniedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultTokenDenied)
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
		it.Event = new(CunibtcvaultTokenDenied)
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
func (it *CunibtcvaultTokenDeniedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultTokenDeniedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultTokenDenied represents a TokenDenied event raised by the Cunibtcvault contract.
type CunibtcvaultTokenDenied struct {
	Token []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenDenied is a free log retrieval operation binding the contract event 0x68d2d2120aa64293dc680d52802c724abab22ec31f5b6aa303e8d60e445b75f6.
//
// Solidity: event TokenDenied(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) FilterTokenDenied(opts *bind.FilterOpts) (*CunibtcvaultTokenDeniedIterator, error) {

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "TokenDenied")
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultTokenDeniedIterator{contract: _Cunibtcvault.contract, event: "TokenDenied", logs: logs, sub: sub}, nil
}

// WatchTokenDenied is a free log subscription operation binding the contract event 0x68d2d2120aa64293dc680d52802c724abab22ec31f5b6aa303e8d60e445b75f6.
//
// Solidity: event TokenDenied(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) WatchTokenDenied(opts *bind.WatchOpts, sink chan<- *CunibtcvaultTokenDenied) (event.Subscription, error) {

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "TokenDenied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultTokenDenied)
				if err := _Cunibtcvault.contract.UnpackLog(event, "TokenDenied", log); err != nil {
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

// ParseTokenDenied is a log parse operation binding the contract event 0x68d2d2120aa64293dc680d52802c724abab22ec31f5b6aa303e8d60e445b75f6.
//
// Solidity: event TokenDenied(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) ParseTokenDenied(log types.Log) (*CunibtcvaultTokenDenied, error) {
	event := new(CunibtcvaultTokenDenied)
	if err := _Cunibtcvault.contract.UnpackLog(event, "TokenDenied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultTokenPausedIterator is returned from FilterTokenPaused and is used to iterate over the raw logs and unpacked data for TokenPaused events raised by the Cunibtcvault contract.
type CunibtcvaultTokenPausedIterator struct {
	Event *CunibtcvaultTokenPaused // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultTokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultTokenPaused)
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
		it.Event = new(CunibtcvaultTokenPaused)
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
func (it *CunibtcvaultTokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultTokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultTokenPaused represents a TokenPaused event raised by the Cunibtcvault contract.
type CunibtcvaultTokenPaused struct {
	Token []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenPaused is a free log retrieval operation binding the contract event 0x424453f3372b3b94183712f7b3e6532da132633b2279015cc4c17cea0a3d965d.
//
// Solidity: event TokenPaused(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) FilterTokenPaused(opts *bind.FilterOpts) (*CunibtcvaultTokenPausedIterator, error) {

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "TokenPaused")
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultTokenPausedIterator{contract: _Cunibtcvault.contract, event: "TokenPaused", logs: logs, sub: sub}, nil
}

// WatchTokenPaused is a free log subscription operation binding the contract event 0x424453f3372b3b94183712f7b3e6532da132633b2279015cc4c17cea0a3d965d.
//
// Solidity: event TokenPaused(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) WatchTokenPaused(opts *bind.WatchOpts, sink chan<- *CunibtcvaultTokenPaused) (event.Subscription, error) {

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "TokenPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultTokenPaused)
				if err := _Cunibtcvault.contract.UnpackLog(event, "TokenPaused", log); err != nil {
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

// ParseTokenPaused is a log parse operation binding the contract event 0x424453f3372b3b94183712f7b3e6532da132633b2279015cc4c17cea0a3d965d.
//
// Solidity: event TokenPaused(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) ParseTokenPaused(log types.Log) (*CunibtcvaultTokenPaused, error) {
	event := new(CunibtcvaultTokenPaused)
	if err := _Cunibtcvault.contract.UnpackLog(event, "TokenPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultTokenUnpausedIterator is returned from FilterTokenUnpaused and is used to iterate over the raw logs and unpacked data for TokenUnpaused events raised by the Cunibtcvault contract.
type CunibtcvaultTokenUnpausedIterator struct {
	Event *CunibtcvaultTokenUnpaused // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultTokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultTokenUnpaused)
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
		it.Event = new(CunibtcvaultTokenUnpaused)
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
func (it *CunibtcvaultTokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultTokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultTokenUnpaused represents a TokenUnpaused event raised by the Cunibtcvault contract.
type CunibtcvaultTokenUnpaused struct {
	Token []common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenUnpaused is a free log retrieval operation binding the contract event 0x9c6609137c6c2d0073a208dfe728a5f0178b08ead915514ea3fe8bba8f5fc365.
//
// Solidity: event TokenUnpaused(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) FilterTokenUnpaused(opts *bind.FilterOpts) (*CunibtcvaultTokenUnpausedIterator, error) {

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "TokenUnpaused")
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultTokenUnpausedIterator{contract: _Cunibtcvault.contract, event: "TokenUnpaused", logs: logs, sub: sub}, nil
}

// WatchTokenUnpaused is a free log subscription operation binding the contract event 0x9c6609137c6c2d0073a208dfe728a5f0178b08ead915514ea3fe8bba8f5fc365.
//
// Solidity: event TokenUnpaused(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) WatchTokenUnpaused(opts *bind.WatchOpts, sink chan<- *CunibtcvaultTokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "TokenUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultTokenUnpaused)
				if err := _Cunibtcvault.contract.UnpackLog(event, "TokenUnpaused", log); err != nil {
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

// ParseTokenUnpaused is a log parse operation binding the contract event 0x9c6609137c6c2d0073a208dfe728a5f0178b08ead915514ea3fe8bba8f5fc365.
//
// Solidity: event TokenUnpaused(address[] token)
func (_Cunibtcvault *CunibtcvaultFilterer) ParseTokenUnpaused(log types.Log) (*CunibtcvaultTokenUnpaused, error) {
	event := new(CunibtcvaultTokenUnpaused)
	if err := _Cunibtcvault.contract.UnpackLog(event, "TokenUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CunibtcvaultTotalSupplySetIterator is returned from FilterTotalSupplySet and is used to iterate over the raw logs and unpacked data for TotalSupplySet events raised by the Cunibtcvault contract.
type CunibtcvaultTotalSupplySetIterator struct {
	Event *CunibtcvaultTotalSupplySet // Event containing the contract specifics and raw log

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
func (it *CunibtcvaultTotalSupplySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CunibtcvaultTotalSupplySet)
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
		it.Event = new(CunibtcvaultTotalSupplySet)
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
func (it *CunibtcvaultTotalSupplySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CunibtcvaultTotalSupplySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CunibtcvaultTotalSupplySet represents a TotalSupplySet event raised by the Cunibtcvault contract.
type CunibtcvaultTotalSupplySet struct {
	TotalSupply *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTotalSupplySet is a free log retrieval operation binding the contract event 0xe876c27307b9e6aa2b15a7730a0dfd41be0f2c7a7c39bc08e89df77a548fb794.
//
// Solidity: event TotalSupplySet(uint256 totalSupply)
func (_Cunibtcvault *CunibtcvaultFilterer) FilterTotalSupplySet(opts *bind.FilterOpts) (*CunibtcvaultTotalSupplySetIterator, error) {

	logs, sub, err := _Cunibtcvault.contract.FilterLogs(opts, "TotalSupplySet")
	if err != nil {
		return nil, err
	}
	return &CunibtcvaultTotalSupplySetIterator{contract: _Cunibtcvault.contract, event: "TotalSupplySet", logs: logs, sub: sub}, nil
}

// WatchTotalSupplySet is a free log subscription operation binding the contract event 0xe876c27307b9e6aa2b15a7730a0dfd41be0f2c7a7c39bc08e89df77a548fb794.
//
// Solidity: event TotalSupplySet(uint256 totalSupply)
func (_Cunibtcvault *CunibtcvaultFilterer) WatchTotalSupplySet(opts *bind.WatchOpts, sink chan<- *CunibtcvaultTotalSupplySet) (event.Subscription, error) {

	logs, sub, err := _Cunibtcvault.contract.WatchLogs(opts, "TotalSupplySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CunibtcvaultTotalSupplySet)
				if err := _Cunibtcvault.contract.UnpackLog(event, "TotalSupplySet", log); err != nil {
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

// ParseTotalSupplySet is a log parse operation binding the contract event 0xe876c27307b9e6aa2b15a7730a0dfd41be0f2c7a7c39bc08e89df77a548fb794.
//
// Solidity: event TotalSupplySet(uint256 totalSupply)
func (_Cunibtcvault *CunibtcvaultFilterer) ParseTotalSupplySet(log types.Log) (*CunibtcvaultTotalSupplySet, error) {
	event := new(CunibtcvaultTotalSupplySet)
	if err := _Cunibtcvault.contract.UnpackLog(event, "TotalSupplySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
