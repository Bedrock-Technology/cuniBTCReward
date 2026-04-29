// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package airdrop

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

// AirdropDist is an auto generated low-level Go binding around an user-defined struct.
type AirdropDist struct {
	Root        [32]byte
	ActivatedAt *big.Int
	Duration    *big.Int
	Disabled    bool
	Token       common.Address
}

// AirdropMetaData contains all meta data concerning the Airdrop contract.
var AirdropMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldActivationDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newActivationDelay\",\"type\":\"uint256\"}],\"name\":\"ActivationDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AirdropClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"preStatus\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"DistributionDisabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardsValidTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activatedAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"MerkleRootSubmit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"preRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"MerkleRootUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"preToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"preValidDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validDuration\",\"type\":\"uint256\"}],\"name\":\"ValidDurationUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activationDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amount\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"_proof\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[]\",\"name\":\"_epoch\",\"type\":\"uint256[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"getRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"activatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structAirdrop.Dist\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_users\",\"type\":\"address[]\"}],\"name\":\"hasClaimed\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_epoch\",\"type\":\"uint256[]\"}],\"name\":\"hasClaimed\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_activationDelay\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_disabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"setAirdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_activationDelay\",\"type\":\"uint256\"}],\"name\":\"setDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"submitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"updateDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"updateRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"updateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AirdropABI is the input ABI used to generate the binding from.
// Deprecated: Use AirdropMetaData.ABI instead.
var AirdropABI = AirdropMetaData.ABI

// Airdrop is an auto generated Go binding around an Ethereum contract.
type Airdrop struct {
	AirdropCaller     // Read-only binding to the contract
	AirdropTransactor // Write-only binding to the contract
	AirdropFilterer   // Log filterer for contract events
}

// AirdropCaller is an auto generated read-only Go binding around an Ethereum contract.
type AirdropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AirdropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AirdropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AirdropSession struct {
	Contract     *Airdrop          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AirdropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AirdropCallerSession struct {
	Contract *AirdropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AirdropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AirdropTransactorSession struct {
	Contract     *AirdropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AirdropRaw is an auto generated low-level Go binding around an Ethereum contract.
type AirdropRaw struct {
	Contract *Airdrop // Generic contract binding to access the raw methods on
}

// AirdropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AirdropCallerRaw struct {
	Contract *AirdropCaller // Generic read-only contract binding to access the raw methods on
}

// AirdropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AirdropTransactorRaw struct {
	Contract *AirdropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAirdrop creates a new instance of Airdrop, bound to a specific deployed contract.
func NewAirdrop(address common.Address, backend bind.ContractBackend) (*Airdrop, error) {
	contract, err := bindAirdrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Airdrop{AirdropCaller: AirdropCaller{contract: contract}, AirdropTransactor: AirdropTransactor{contract: contract}, AirdropFilterer: AirdropFilterer{contract: contract}}, nil
}

// NewAirdropCaller creates a new read-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropCaller(address common.Address, caller bind.ContractCaller) (*AirdropCaller, error) {
	contract, err := bindAirdrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropCaller{contract: contract}, nil
}

// NewAirdropTransactor creates a new write-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropTransactor(address common.Address, transactor bind.ContractTransactor) (*AirdropTransactor, error) {
	contract, err := bindAirdrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropTransactor{contract: contract}, nil
}

// NewAirdropFilterer creates a new log filterer instance of Airdrop, bound to a specific deployed contract.
func NewAirdropFilterer(address common.Address, filterer bind.ContractFilterer) (*AirdropFilterer, error) {
	contract, err := bindAirdrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AirdropFilterer{contract: contract}, nil
}

// bindAirdrop binds a generic wrapper to an already deployed contract.
func bindAirdrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AirdropMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.AirdropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Airdrop *AirdropCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Airdrop *AirdropSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Airdrop.Contract.DEFAULTADMINROLE(&_Airdrop.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Airdrop *AirdropCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Airdrop.Contract.DEFAULTADMINROLE(&_Airdrop.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Airdrop *AirdropCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Airdrop *AirdropSession) OPERATORROLE() ([32]byte, error) {
	return _Airdrop.Contract.OPERATORROLE(&_Airdrop.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Airdrop *AirdropCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Airdrop.Contract.OPERATORROLE(&_Airdrop.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Airdrop *AirdropCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Airdrop *AirdropSession) PAUSERROLE() ([32]byte, error) {
	return _Airdrop.Contract.PAUSERROLE(&_Airdrop.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Airdrop *AirdropCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Airdrop.Contract.PAUSERROLE(&_Airdrop.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint256)
func (_Airdrop *AirdropCaller) ActivationDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "activationDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint256)
func (_Airdrop *AirdropSession) ActivationDelay() (*big.Int, error) {
	return _Airdrop.Contract.ActivationDelay(&_Airdrop.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint256)
func (_Airdrop *AirdropCallerSession) ActivationDelay() (*big.Int, error) {
	return _Airdrop.Contract.ActivationDelay(&_Airdrop.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Airdrop *AirdropCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Airdrop *AirdropSession) CurrentEpoch() (*big.Int, error) {
	return _Airdrop.Contract.CurrentEpoch(&_Airdrop.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Airdrop *AirdropCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Airdrop.Contract.CurrentEpoch(&_Airdrop.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Airdrop *AirdropCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Airdrop *AirdropSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Airdrop.Contract.GetRoleAdmin(&_Airdrop.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Airdrop *AirdropCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Airdrop.Contract.GetRoleAdmin(&_Airdrop.CallOpts, role)
}

// GetRoot is a free data retrieval call binding the contract method 0x9b24b3b0.
//
// Solidity: function getRoot(uint256 _epoch) view returns((bytes32,uint256,uint256,bool,address))
func (_Airdrop *AirdropCaller) GetRoot(opts *bind.CallOpts, _epoch *big.Int) (AirdropDist, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "getRoot", _epoch)

	if err != nil {
		return *new(AirdropDist), err
	}

	out0 := *abi.ConvertType(out[0], new(AirdropDist)).(*AirdropDist)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x9b24b3b0.
//
// Solidity: function getRoot(uint256 _epoch) view returns((bytes32,uint256,uint256,bool,address))
func (_Airdrop *AirdropSession) GetRoot(_epoch *big.Int) (AirdropDist, error) {
	return _Airdrop.Contract.GetRoot(&_Airdrop.CallOpts, _epoch)
}

// GetRoot is a free data retrieval call binding the contract method 0x9b24b3b0.
//
// Solidity: function getRoot(uint256 _epoch) view returns((bytes32,uint256,uint256,bool,address))
func (_Airdrop *AirdropCallerSession) GetRoot(_epoch *big.Int) (AirdropDist, error) {
	return _Airdrop.Contract.GetRoot(&_Airdrop.CallOpts, _epoch)
}

// HasClaimed is a free data retrieval call binding the contract method 0x8a8a24e8.
//
// Solidity: function hasClaimed(uint256 _epoch, address[] _users) view returns(bool[])
func (_Airdrop *AirdropCaller) HasClaimed(opts *bind.CallOpts, _epoch *big.Int, _users []common.Address) ([]bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "hasClaimed", _epoch, _users)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// HasClaimed is a free data retrieval call binding the contract method 0x8a8a24e8.
//
// Solidity: function hasClaimed(uint256 _epoch, address[] _users) view returns(bool[])
func (_Airdrop *AirdropSession) HasClaimed(_epoch *big.Int, _users []common.Address) ([]bool, error) {
	return _Airdrop.Contract.HasClaimed(&_Airdrop.CallOpts, _epoch, _users)
}

// HasClaimed is a free data retrieval call binding the contract method 0x8a8a24e8.
//
// Solidity: function hasClaimed(uint256 _epoch, address[] _users) view returns(bool[])
func (_Airdrop *AirdropCallerSession) HasClaimed(_epoch *big.Int, _users []common.Address) ([]bool, error) {
	return _Airdrop.Contract.HasClaimed(&_Airdrop.CallOpts, _epoch, _users)
}

// HasClaimed0 is a free data retrieval call binding the contract method 0xc36e9055.
//
// Solidity: function hasClaimed(address _user, uint256[] _epoch) view returns(bool[])
func (_Airdrop *AirdropCaller) HasClaimed0(opts *bind.CallOpts, _user common.Address, _epoch []*big.Int) ([]bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "hasClaimed0", _user, _epoch)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// HasClaimed0 is a free data retrieval call binding the contract method 0xc36e9055.
//
// Solidity: function hasClaimed(address _user, uint256[] _epoch) view returns(bool[])
func (_Airdrop *AirdropSession) HasClaimed0(_user common.Address, _epoch []*big.Int) ([]bool, error) {
	return _Airdrop.Contract.HasClaimed0(&_Airdrop.CallOpts, _user, _epoch)
}

// HasClaimed0 is a free data retrieval call binding the contract method 0xc36e9055.
//
// Solidity: function hasClaimed(address _user, uint256[] _epoch) view returns(bool[])
func (_Airdrop *AirdropCallerSession) HasClaimed0(_user common.Address, _epoch []*big.Int) ([]bool, error) {
	return _Airdrop.Contract.HasClaimed0(&_Airdrop.CallOpts, _user, _epoch)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Airdrop *AirdropCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Airdrop *AirdropSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Airdrop.Contract.HasRole(&_Airdrop.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Airdrop *AirdropCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Airdrop.Contract.HasRole(&_Airdrop.CallOpts, role, account)
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 _epoch) view returns(bool)
func (_Airdrop *AirdropCaller) IsActive(opts *bind.CallOpts, _epoch *big.Int) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "isActive", _epoch)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 _epoch) view returns(bool)
func (_Airdrop *AirdropSession) IsActive(_epoch *big.Int) (bool, error) {
	return _Airdrop.Contract.IsActive(&_Airdrop.CallOpts, _epoch)
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 _epoch) view returns(bool)
func (_Airdrop *AirdropCallerSession) IsActive(_epoch *big.Int) (bool, error) {
	return _Airdrop.Contract.IsActive(&_Airdrop.CallOpts, _epoch)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Airdrop *AirdropCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Airdrop *AirdropSession) Paused() (bool, error) {
	return _Airdrop.Contract.Paused(&_Airdrop.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Airdrop *AirdropCallerSession) Paused() (bool, error) {
	return _Airdrop.Contract.Paused(&_Airdrop.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Airdrop *AirdropCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Airdrop *AirdropSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Airdrop.Contract.SupportsInterface(&_Airdrop.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Airdrop *AirdropCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Airdrop.Contract.SupportsInterface(&_Airdrop.CallOpts, interfaceId)
}

// Claim is a paid mutator transaction binding the contract method 0x2f52ebb7.
//
// Solidity: function claim(uint256 _amount, bytes32[] _proof) returns()
func (_Airdrop *AirdropTransactor) Claim(opts *bind.TransactOpts, _amount *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "claim", _amount, _proof)
}

// Claim is a paid mutator transaction binding the contract method 0x2f52ebb7.
//
// Solidity: function claim(uint256 _amount, bytes32[] _proof) returns()
func (_Airdrop *AirdropSession) Claim(_amount *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _Airdrop.Contract.Claim(&_Airdrop.TransactOpts, _amount, _proof)
}

// Claim is a paid mutator transaction binding the contract method 0x2f52ebb7.
//
// Solidity: function claim(uint256 _amount, bytes32[] _proof) returns()
func (_Airdrop *AirdropTransactorSession) Claim(_amount *big.Int, _proof [][32]byte) (*types.Transaction, error) {
	return _Airdrop.Contract.Claim(&_Airdrop.TransactOpts, _amount, _proof)
}

// Claim0 is a paid mutator transaction binding the contract method 0x9535be2a.
//
// Solidity: function claim(uint256[] _amount, bytes32[][] _proof, uint256[] _epoch) returns()
func (_Airdrop *AirdropTransactor) Claim0(opts *bind.TransactOpts, _amount []*big.Int, _proof [][][32]byte, _epoch []*big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "claim0", _amount, _proof, _epoch)
}

// Claim0 is a paid mutator transaction binding the contract method 0x9535be2a.
//
// Solidity: function claim(uint256[] _amount, bytes32[][] _proof, uint256[] _epoch) returns()
func (_Airdrop *AirdropSession) Claim0(_amount []*big.Int, _proof [][][32]byte, _epoch []*big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.Claim0(&_Airdrop.TransactOpts, _amount, _proof, _epoch)
}

// Claim0 is a paid mutator transaction binding the contract method 0x9535be2a.
//
// Solidity: function claim(uint256[] _amount, bytes32[][] _proof, uint256[] _epoch) returns()
func (_Airdrop *AirdropTransactorSession) Claim0(_amount []*big.Int, _proof [][][32]byte, _epoch []*big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.Claim0(&_Airdrop.TransactOpts, _amount, _proof, _epoch)
}

// Claim1 is a paid mutator transaction binding the contract method 0xf339f526.
//
// Solidity: function claim(uint256 _amount, bytes32[] _proof, uint256 _epoch) returns()
func (_Airdrop *AirdropTransactor) Claim1(opts *bind.TransactOpts, _amount *big.Int, _proof [][32]byte, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "claim1", _amount, _proof, _epoch)
}

// Claim1 is a paid mutator transaction binding the contract method 0xf339f526.
//
// Solidity: function claim(uint256 _amount, bytes32[] _proof, uint256 _epoch) returns()
func (_Airdrop *AirdropSession) Claim1(_amount *big.Int, _proof [][32]byte, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.Claim1(&_Airdrop.TransactOpts, _amount, _proof, _epoch)
}

// Claim1 is a paid mutator transaction binding the contract method 0xf339f526.
//
// Solidity: function claim(uint256 _amount, bytes32[] _proof, uint256 _epoch) returns()
func (_Airdrop *AirdropTransactorSession) Claim1(_amount *big.Int, _proof [][32]byte, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.Claim1(&_Airdrop.TransactOpts, _amount, _proof, _epoch)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.GrantRole(&_Airdrop.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.GrantRole(&_Airdrop.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 _activationDelay, address _admin) returns()
func (_Airdrop *AirdropTransactor) Initialize(opts *bind.TransactOpts, _activationDelay *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "initialize", _activationDelay, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 _activationDelay, address _admin) returns()
func (_Airdrop *AirdropSession) Initialize(_activationDelay *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.Initialize(&_Airdrop.TransactOpts, _activationDelay, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 _activationDelay, address _admin) returns()
func (_Airdrop *AirdropTransactorSession) Initialize(_activationDelay *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.Initialize(&_Airdrop.TransactOpts, _activationDelay, _admin)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Airdrop *AirdropTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Airdrop *AirdropSession) Pause() (*types.Transaction, error) {
	return _Airdrop.Contract.Pause(&_Airdrop.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Airdrop *AirdropTransactorSession) Pause() (*types.Transaction, error) {
	return _Airdrop.Contract.Pause(&_Airdrop.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.RenounceRole(&_Airdrop.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.RenounceRole(&_Airdrop.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.RevokeRole(&_Airdrop.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Airdrop *AirdropTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.RevokeRole(&_Airdrop.TransactOpts, role, account)
}

// SetAirdrop is a paid mutator transaction binding the contract method 0x5d42dcec.
//
// Solidity: function setAirdrop(bool _disabled, uint256 _epoch) returns()
func (_Airdrop *AirdropTransactor) SetAirdrop(opts *bind.TransactOpts, _disabled bool, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "setAirdrop", _disabled, _epoch)
}

// SetAirdrop is a paid mutator transaction binding the contract method 0x5d42dcec.
//
// Solidity: function setAirdrop(bool _disabled, uint256 _epoch) returns()
func (_Airdrop *AirdropSession) SetAirdrop(_disabled bool, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.SetAirdrop(&_Airdrop.TransactOpts, _disabled, _epoch)
}

// SetAirdrop is a paid mutator transaction binding the contract method 0x5d42dcec.
//
// Solidity: function setAirdrop(bool _disabled, uint256 _epoch) returns()
func (_Airdrop *AirdropTransactorSession) SetAirdrop(_disabled bool, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.SetAirdrop(&_Airdrop.TransactOpts, _disabled, _epoch)
}

// SetDelay is a paid mutator transaction binding the contract method 0xe177246e.
//
// Solidity: function setDelay(uint256 _activationDelay) returns()
func (_Airdrop *AirdropTransactor) SetDelay(opts *bind.TransactOpts, _activationDelay *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "setDelay", _activationDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0xe177246e.
//
// Solidity: function setDelay(uint256 _activationDelay) returns()
func (_Airdrop *AirdropSession) SetDelay(_activationDelay *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.SetDelay(&_Airdrop.TransactOpts, _activationDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0xe177246e.
//
// Solidity: function setDelay(uint256 _activationDelay) returns()
func (_Airdrop *AirdropTransactorSession) SetDelay(_activationDelay *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.SetDelay(&_Airdrop.TransactOpts, _activationDelay)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0xf8b6cf23.
//
// Solidity: function submitRoot(bytes32 _newRoot, uint256 _duration, address _token) returns()
func (_Airdrop *AirdropTransactor) SubmitRoot(opts *bind.TransactOpts, _newRoot [32]byte, _duration *big.Int, _token common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "submitRoot", _newRoot, _duration, _token)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0xf8b6cf23.
//
// Solidity: function submitRoot(bytes32 _newRoot, uint256 _duration, address _token) returns()
func (_Airdrop *AirdropSession) SubmitRoot(_newRoot [32]byte, _duration *big.Int, _token common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.SubmitRoot(&_Airdrop.TransactOpts, _newRoot, _duration, _token)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0xf8b6cf23.
//
// Solidity: function submitRoot(bytes32 _newRoot, uint256 _duration, address _token) returns()
func (_Airdrop *AirdropTransactorSession) SubmitRoot(_newRoot [32]byte, _duration *big.Int, _token common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.SubmitRoot(&_Airdrop.TransactOpts, _newRoot, _duration, _token)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Airdrop *AirdropTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Airdrop *AirdropSession) Unpause() (*types.Transaction, error) {
	return _Airdrop.Contract.Unpause(&_Airdrop.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Airdrop *AirdropTransactorSession) Unpause() (*types.Transaction, error) {
	return _Airdrop.Contract.Unpause(&_Airdrop.TransactOpts)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x2ad1f1db.
//
// Solidity: function updateDuration(uint256 _duration, uint256 _epoch) returns()
func (_Airdrop *AirdropTransactor) UpdateDuration(opts *bind.TransactOpts, _duration *big.Int, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "updateDuration", _duration, _epoch)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x2ad1f1db.
//
// Solidity: function updateDuration(uint256 _duration, uint256 _epoch) returns()
func (_Airdrop *AirdropSession) UpdateDuration(_duration *big.Int, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.UpdateDuration(&_Airdrop.TransactOpts, _duration, _epoch)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x2ad1f1db.
//
// Solidity: function updateDuration(uint256 _duration, uint256 _epoch) returns()
func (_Airdrop *AirdropTransactorSession) UpdateDuration(_duration *big.Int, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.UpdateDuration(&_Airdrop.TransactOpts, _duration, _epoch)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0xc2e3e62b.
//
// Solidity: function updateRoot(bytes32 _newRoot, uint256 _epoch) returns()
func (_Airdrop *AirdropTransactor) UpdateRoot(opts *bind.TransactOpts, _newRoot [32]byte, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "updateRoot", _newRoot, _epoch)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0xc2e3e62b.
//
// Solidity: function updateRoot(bytes32 _newRoot, uint256 _epoch) returns()
func (_Airdrop *AirdropSession) UpdateRoot(_newRoot [32]byte, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.UpdateRoot(&_Airdrop.TransactOpts, _newRoot, _epoch)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0xc2e3e62b.
//
// Solidity: function updateRoot(bytes32 _newRoot, uint256 _epoch) returns()
func (_Airdrop *AirdropTransactorSession) UpdateRoot(_newRoot [32]byte, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.UpdateRoot(&_Airdrop.TransactOpts, _newRoot, _epoch)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x2bce9e7b.
//
// Solidity: function updateToken(address _token, uint256 _epoch) returns()
func (_Airdrop *AirdropTransactor) UpdateToken(opts *bind.TransactOpts, _token common.Address, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "updateToken", _token, _epoch)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x2bce9e7b.
//
// Solidity: function updateToken(address _token, uint256 _epoch) returns()
func (_Airdrop *AirdropSession) UpdateToken(_token common.Address, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.UpdateToken(&_Airdrop.TransactOpts, _token, _epoch)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x2bce9e7b.
//
// Solidity: function updateToken(address _token, uint256 _epoch) returns()
func (_Airdrop *AirdropTransactorSession) UpdateToken(_token common.Address, _epoch *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.UpdateToken(&_Airdrop.TransactOpts, _token, _epoch)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _to, uint256 _amount) returns()
func (_Airdrop *AirdropTransactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "withdraw", _token, _to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _to, uint256 _amount) returns()
func (_Airdrop *AirdropSession) Withdraw(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.Withdraw(&_Airdrop.TransactOpts, _token, _to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _to, uint256 _amount) returns()
func (_Airdrop *AirdropTransactorSession) Withdraw(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.Withdraw(&_Airdrop.TransactOpts, _token, _to, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Airdrop *AirdropTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Airdrop *AirdropSession) Receive() (*types.Transaction, error) {
	return _Airdrop.Contract.Receive(&_Airdrop.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Airdrop *AirdropTransactorSession) Receive() (*types.Transaction, error) {
	return _Airdrop.Contract.Receive(&_Airdrop.TransactOpts)
}

// AirdropActivationDelaySetIterator is returned from FilterActivationDelaySet and is used to iterate over the raw logs and unpacked data for ActivationDelaySet events raised by the Airdrop contract.
type AirdropActivationDelaySetIterator struct {
	Event *AirdropActivationDelaySet // Event containing the contract specifics and raw log

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
func (it *AirdropActivationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropActivationDelaySet)
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
		it.Event = new(AirdropActivationDelaySet)
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
func (it *AirdropActivationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropActivationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropActivationDelaySet represents a ActivationDelaySet event raised by the Airdrop contract.
type AirdropActivationDelaySet struct {
	OldActivationDelay *big.Int
	NewActivationDelay *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterActivationDelaySet is a free log retrieval operation binding the contract event 0x293d72af226bd6f48fc96e954fa80023f4ba3238f6c8803bf4bcd27b97bebf5e.
//
// Solidity: event ActivationDelaySet(uint256 oldActivationDelay, uint256 newActivationDelay)
func (_Airdrop *AirdropFilterer) FilterActivationDelaySet(opts *bind.FilterOpts) (*AirdropActivationDelaySetIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return &AirdropActivationDelaySetIterator{contract: _Airdrop.contract, event: "ActivationDelaySet", logs: logs, sub: sub}, nil
}

// WatchActivationDelaySet is a free log subscription operation binding the contract event 0x293d72af226bd6f48fc96e954fa80023f4ba3238f6c8803bf4bcd27b97bebf5e.
//
// Solidity: event ActivationDelaySet(uint256 oldActivationDelay, uint256 newActivationDelay)
func (_Airdrop *AirdropFilterer) WatchActivationDelaySet(opts *bind.WatchOpts, sink chan<- *AirdropActivationDelaySet) (event.Subscription, error) {

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropActivationDelaySet)
				if err := _Airdrop.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
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

// ParseActivationDelaySet is a log parse operation binding the contract event 0x293d72af226bd6f48fc96e954fa80023f4ba3238f6c8803bf4bcd27b97bebf5e.
//
// Solidity: event ActivationDelaySet(uint256 oldActivationDelay, uint256 newActivationDelay)
func (_Airdrop *AirdropFilterer) ParseActivationDelaySet(log types.Log) (*AirdropActivationDelaySet, error) {
	event := new(AirdropActivationDelaySet)
	if err := _Airdrop.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropAirdropClaimedIterator is returned from FilterAirdropClaimed and is used to iterate over the raw logs and unpacked data for AirdropClaimed events raised by the Airdrop contract.
type AirdropAirdropClaimedIterator struct {
	Event *AirdropAirdropClaimed // Event containing the contract specifics and raw log

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
func (it *AirdropAirdropClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropAirdropClaimed)
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
		it.Event = new(AirdropAirdropClaimed)
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
func (it *AirdropAirdropClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropAirdropClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropAirdropClaimed represents a AirdropClaimed event raised by the Airdrop contract.
type AirdropAirdropClaimed struct {
	Epoch        *big.Int
	User         common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAirdropClaimed is a free log retrieval operation binding the contract event 0x04e4e409a0e62b53ae4ba4894edd5d35b224608622a728788940f6ea03a8b8ac.
//
// Solidity: event AirdropClaimed(uint256 indexed epoch, address indexed user, address tokenAddress, uint256 amount)
func (_Airdrop *AirdropFilterer) FilterAirdropClaimed(opts *bind.FilterOpts, epoch []*big.Int, user []common.Address) (*AirdropAirdropClaimedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "AirdropClaimed", epochRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AirdropAirdropClaimedIterator{contract: _Airdrop.contract, event: "AirdropClaimed", logs: logs, sub: sub}, nil
}

// WatchAirdropClaimed is a free log subscription operation binding the contract event 0x04e4e409a0e62b53ae4ba4894edd5d35b224608622a728788940f6ea03a8b8ac.
//
// Solidity: event AirdropClaimed(uint256 indexed epoch, address indexed user, address tokenAddress, uint256 amount)
func (_Airdrop *AirdropFilterer) WatchAirdropClaimed(opts *bind.WatchOpts, sink chan<- *AirdropAirdropClaimed, epoch []*big.Int, user []common.Address) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "AirdropClaimed", epochRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropAirdropClaimed)
				if err := _Airdrop.contract.UnpackLog(event, "AirdropClaimed", log); err != nil {
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

// ParseAirdropClaimed is a log parse operation binding the contract event 0x04e4e409a0e62b53ae4ba4894edd5d35b224608622a728788940f6ea03a8b8ac.
//
// Solidity: event AirdropClaimed(uint256 indexed epoch, address indexed user, address tokenAddress, uint256 amount)
func (_Airdrop *AirdropFilterer) ParseAirdropClaimed(log types.Log) (*AirdropAirdropClaimed, error) {
	event := new(AirdropAirdropClaimed)
	if err := _Airdrop.contract.UnpackLog(event, "AirdropClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropDistributionDisabledSetIterator is returned from FilterDistributionDisabledSet and is used to iterate over the raw logs and unpacked data for DistributionDisabledSet events raised by the Airdrop contract.
type AirdropDistributionDisabledSetIterator struct {
	Event *AirdropDistributionDisabledSet // Event containing the contract specifics and raw log

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
func (it *AirdropDistributionDisabledSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropDistributionDisabledSet)
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
		it.Event = new(AirdropDistributionDisabledSet)
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
func (it *AirdropDistributionDisabledSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropDistributionDisabledSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropDistributionDisabledSet represents a DistributionDisabledSet event raised by the Airdrop contract.
type AirdropDistributionDisabledSet struct {
	Epoch     *big.Int
	PreStatus bool
	Status    bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributionDisabledSet is a free log retrieval operation binding the contract event 0xc143c2d1f7eddc3d4531772bc8494966bb1ce9886c51498d36756b4ff1c350eb.
//
// Solidity: event DistributionDisabledSet(uint256 indexed epoch, bool preStatus, bool status)
func (_Airdrop *AirdropFilterer) FilterDistributionDisabledSet(opts *bind.FilterOpts, epoch []*big.Int) (*AirdropDistributionDisabledSetIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "DistributionDisabledSet", epochRule)
	if err != nil {
		return nil, err
	}
	return &AirdropDistributionDisabledSetIterator{contract: _Airdrop.contract, event: "DistributionDisabledSet", logs: logs, sub: sub}, nil
}

// WatchDistributionDisabledSet is a free log subscription operation binding the contract event 0xc143c2d1f7eddc3d4531772bc8494966bb1ce9886c51498d36756b4ff1c350eb.
//
// Solidity: event DistributionDisabledSet(uint256 indexed epoch, bool preStatus, bool status)
func (_Airdrop *AirdropFilterer) WatchDistributionDisabledSet(opts *bind.WatchOpts, sink chan<- *AirdropDistributionDisabledSet, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "DistributionDisabledSet", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropDistributionDisabledSet)
				if err := _Airdrop.contract.UnpackLog(event, "DistributionDisabledSet", log); err != nil {
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

// ParseDistributionDisabledSet is a log parse operation binding the contract event 0xc143c2d1f7eddc3d4531772bc8494966bb1ce9886c51498d36756b4ff1c350eb.
//
// Solidity: event DistributionDisabledSet(uint256 indexed epoch, bool preStatus, bool status)
func (_Airdrop *AirdropFilterer) ParseDistributionDisabledSet(log types.Log) (*AirdropDistributionDisabledSet, error) {
	event := new(AirdropDistributionDisabledSet)
	if err := _Airdrop.contract.UnpackLog(event, "DistributionDisabledSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Airdrop contract.
type AirdropInitializedIterator struct {
	Event *AirdropInitialized // Event containing the contract specifics and raw log

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
func (it *AirdropInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropInitialized)
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
		it.Event = new(AirdropInitialized)
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
func (it *AirdropInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropInitialized represents a Initialized event raised by the Airdrop contract.
type AirdropInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Airdrop *AirdropFilterer) FilterInitialized(opts *bind.FilterOpts) (*AirdropInitializedIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AirdropInitializedIterator{contract: _Airdrop.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Airdrop *AirdropFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AirdropInitialized) (event.Subscription, error) {

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropInitialized)
				if err := _Airdrop.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseInitialized(log types.Log) (*AirdropInitialized, error) {
	event := new(AirdropInitialized)
	if err := _Airdrop.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropMerkleRootSubmitIterator is returned from FilterMerkleRootSubmit and is used to iterate over the raw logs and unpacked data for MerkleRootSubmit events raised by the Airdrop contract.
type AirdropMerkleRootSubmitIterator struct {
	Event *AirdropMerkleRootSubmit // Event containing the contract specifics and raw log

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
func (it *AirdropMerkleRootSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropMerkleRootSubmit)
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
		it.Event = new(AirdropMerkleRootSubmit)
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
func (it *AirdropMerkleRootSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropMerkleRootSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropMerkleRootSubmit represents a MerkleRootSubmit event raised by the Airdrop contract.
type AirdropMerkleRootSubmit struct {
	Epoch            *big.Int
	Root             [32]byte
	RewardsValidTime *big.Int
	ActivatedAt      *big.Int
	Token            common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMerkleRootSubmit is a free log retrieval operation binding the contract event 0x7ef8df9ad985c94bf7202817536e2ae484e5541c6ffb8940a37e87734d6968a2.
//
// Solidity: event MerkleRootSubmit(uint256 indexed epoch, bytes32 root, uint256 rewardsValidTime, uint256 activatedAt, address token)
func (_Airdrop *AirdropFilterer) FilterMerkleRootSubmit(opts *bind.FilterOpts, epoch []*big.Int) (*AirdropMerkleRootSubmitIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "MerkleRootSubmit", epochRule)
	if err != nil {
		return nil, err
	}
	return &AirdropMerkleRootSubmitIterator{contract: _Airdrop.contract, event: "MerkleRootSubmit", logs: logs, sub: sub}, nil
}

// WatchMerkleRootSubmit is a free log subscription operation binding the contract event 0x7ef8df9ad985c94bf7202817536e2ae484e5541c6ffb8940a37e87734d6968a2.
//
// Solidity: event MerkleRootSubmit(uint256 indexed epoch, bytes32 root, uint256 rewardsValidTime, uint256 activatedAt, address token)
func (_Airdrop *AirdropFilterer) WatchMerkleRootSubmit(opts *bind.WatchOpts, sink chan<- *AirdropMerkleRootSubmit, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "MerkleRootSubmit", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropMerkleRootSubmit)
				if err := _Airdrop.contract.UnpackLog(event, "MerkleRootSubmit", log); err != nil {
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

// ParseMerkleRootSubmit is a log parse operation binding the contract event 0x7ef8df9ad985c94bf7202817536e2ae484e5541c6ffb8940a37e87734d6968a2.
//
// Solidity: event MerkleRootSubmit(uint256 indexed epoch, bytes32 root, uint256 rewardsValidTime, uint256 activatedAt, address token)
func (_Airdrop *AirdropFilterer) ParseMerkleRootSubmit(log types.Log) (*AirdropMerkleRootSubmit, error) {
	event := new(AirdropMerkleRootSubmit)
	if err := _Airdrop.contract.UnpackLog(event, "MerkleRootSubmit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropMerkleRootUpdateIterator is returned from FilterMerkleRootUpdate and is used to iterate over the raw logs and unpacked data for MerkleRootUpdate events raised by the Airdrop contract.
type AirdropMerkleRootUpdateIterator struct {
	Event *AirdropMerkleRootUpdate // Event containing the contract specifics and raw log

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
func (it *AirdropMerkleRootUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropMerkleRootUpdate)
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
		it.Event = new(AirdropMerkleRootUpdate)
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
func (it *AirdropMerkleRootUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropMerkleRootUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropMerkleRootUpdate represents a MerkleRootUpdate event raised by the Airdrop contract.
type AirdropMerkleRootUpdate struct {
	Epoch   *big.Int
	PreRoot [32]byte
	Root    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMerkleRootUpdate is a free log retrieval operation binding the contract event 0x744fd8c3189403b64f91de414accb2c31f570bd1251a1463981ce3c3b479eb5e.
//
// Solidity: event MerkleRootUpdate(uint256 indexed epoch, bytes32 preRoot, bytes32 root)
func (_Airdrop *AirdropFilterer) FilterMerkleRootUpdate(opts *bind.FilterOpts, epoch []*big.Int) (*AirdropMerkleRootUpdateIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "MerkleRootUpdate", epochRule)
	if err != nil {
		return nil, err
	}
	return &AirdropMerkleRootUpdateIterator{contract: _Airdrop.contract, event: "MerkleRootUpdate", logs: logs, sub: sub}, nil
}

// WatchMerkleRootUpdate is a free log subscription operation binding the contract event 0x744fd8c3189403b64f91de414accb2c31f570bd1251a1463981ce3c3b479eb5e.
//
// Solidity: event MerkleRootUpdate(uint256 indexed epoch, bytes32 preRoot, bytes32 root)
func (_Airdrop *AirdropFilterer) WatchMerkleRootUpdate(opts *bind.WatchOpts, sink chan<- *AirdropMerkleRootUpdate, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "MerkleRootUpdate", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropMerkleRootUpdate)
				if err := _Airdrop.contract.UnpackLog(event, "MerkleRootUpdate", log); err != nil {
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

// ParseMerkleRootUpdate is a log parse operation binding the contract event 0x744fd8c3189403b64f91de414accb2c31f570bd1251a1463981ce3c3b479eb5e.
//
// Solidity: event MerkleRootUpdate(uint256 indexed epoch, bytes32 preRoot, bytes32 root)
func (_Airdrop *AirdropFilterer) ParseMerkleRootUpdate(log types.Log) (*AirdropMerkleRootUpdate, error) {
	event := new(AirdropMerkleRootUpdate)
	if err := _Airdrop.contract.UnpackLog(event, "MerkleRootUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Airdrop contract.
type AirdropPausedIterator struct {
	Event *AirdropPaused // Event containing the contract specifics and raw log

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
func (it *AirdropPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropPaused)
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
		it.Event = new(AirdropPaused)
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
func (it *AirdropPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropPaused represents a Paused event raised by the Airdrop contract.
type AirdropPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Airdrop *AirdropFilterer) FilterPaused(opts *bind.FilterOpts) (*AirdropPausedIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AirdropPausedIterator{contract: _Airdrop.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Airdrop *AirdropFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AirdropPaused) (event.Subscription, error) {

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropPaused)
				if err := _Airdrop.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Airdrop *AirdropFilterer) ParsePaused(log types.Log) (*AirdropPaused, error) {
	event := new(AirdropPaused)
	if err := _Airdrop.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Airdrop contract.
type AirdropRoleAdminChangedIterator struct {
	Event *AirdropRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AirdropRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropRoleAdminChanged)
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
		it.Event = new(AirdropRoleAdminChanged)
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
func (it *AirdropRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropRoleAdminChanged represents a RoleAdminChanged event raised by the Airdrop contract.
type AirdropRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Airdrop *AirdropFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AirdropRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AirdropRoleAdminChangedIterator{contract: _Airdrop.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Airdrop *AirdropFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AirdropRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropRoleAdminChanged)
				if err := _Airdrop.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseRoleAdminChanged(log types.Log) (*AirdropRoleAdminChanged, error) {
	event := new(AirdropRoleAdminChanged)
	if err := _Airdrop.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Airdrop contract.
type AirdropRoleGrantedIterator struct {
	Event *AirdropRoleGranted // Event containing the contract specifics and raw log

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
func (it *AirdropRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropRoleGranted)
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
		it.Event = new(AirdropRoleGranted)
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
func (it *AirdropRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropRoleGranted represents a RoleGranted event raised by the Airdrop contract.
type AirdropRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Airdrop *AirdropFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AirdropRoleGrantedIterator, error) {

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

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AirdropRoleGrantedIterator{contract: _Airdrop.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Airdrop *AirdropFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AirdropRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropRoleGranted)
				if err := _Airdrop.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseRoleGranted(log types.Log) (*AirdropRoleGranted, error) {
	event := new(AirdropRoleGranted)
	if err := _Airdrop.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Airdrop contract.
type AirdropRoleRevokedIterator struct {
	Event *AirdropRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AirdropRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropRoleRevoked)
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
		it.Event = new(AirdropRoleRevoked)
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
func (it *AirdropRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropRoleRevoked represents a RoleRevoked event raised by the Airdrop contract.
type AirdropRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Airdrop *AirdropFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AirdropRoleRevokedIterator, error) {

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

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AirdropRoleRevokedIterator{contract: _Airdrop.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Airdrop *AirdropFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AirdropRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropRoleRevoked)
				if err := _Airdrop.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseRoleRevoked(log types.Log) (*AirdropRoleRevoked, error) {
	event := new(AirdropRoleRevoked)
	if err := _Airdrop.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropTokenUpdateIterator is returned from FilterTokenUpdate and is used to iterate over the raw logs and unpacked data for TokenUpdate events raised by the Airdrop contract.
type AirdropTokenUpdateIterator struct {
	Event *AirdropTokenUpdate // Event containing the contract specifics and raw log

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
func (it *AirdropTokenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropTokenUpdate)
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
		it.Event = new(AirdropTokenUpdate)
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
func (it *AirdropTokenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropTokenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropTokenUpdate represents a TokenUpdate event raised by the Airdrop contract.
type AirdropTokenUpdate struct {
	Epoch    *big.Int
	PreToken common.Address
	Token    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenUpdate is a free log retrieval operation binding the contract event 0xfffa6d0e61dfc97ff8e10ba0aa55e4a3caefca2aee221fe34d63c3d901a1dd93.
//
// Solidity: event TokenUpdate(uint256 indexed epoch, address preToken, address token)
func (_Airdrop *AirdropFilterer) FilterTokenUpdate(opts *bind.FilterOpts, epoch []*big.Int) (*AirdropTokenUpdateIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "TokenUpdate", epochRule)
	if err != nil {
		return nil, err
	}
	return &AirdropTokenUpdateIterator{contract: _Airdrop.contract, event: "TokenUpdate", logs: logs, sub: sub}, nil
}

// WatchTokenUpdate is a free log subscription operation binding the contract event 0xfffa6d0e61dfc97ff8e10ba0aa55e4a3caefca2aee221fe34d63c3d901a1dd93.
//
// Solidity: event TokenUpdate(uint256 indexed epoch, address preToken, address token)
func (_Airdrop *AirdropFilterer) WatchTokenUpdate(opts *bind.WatchOpts, sink chan<- *AirdropTokenUpdate, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "TokenUpdate", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropTokenUpdate)
				if err := _Airdrop.contract.UnpackLog(event, "TokenUpdate", log); err != nil {
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

// ParseTokenUpdate is a log parse operation binding the contract event 0xfffa6d0e61dfc97ff8e10ba0aa55e4a3caefca2aee221fe34d63c3d901a1dd93.
//
// Solidity: event TokenUpdate(uint256 indexed epoch, address preToken, address token)
func (_Airdrop *AirdropFilterer) ParseTokenUpdate(log types.Log) (*AirdropTokenUpdate, error) {
	event := new(AirdropTokenUpdate)
	if err := _Airdrop.contract.UnpackLog(event, "TokenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Airdrop contract.
type AirdropUnpausedIterator struct {
	Event *AirdropUnpaused // Event containing the contract specifics and raw log

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
func (it *AirdropUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropUnpaused)
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
		it.Event = new(AirdropUnpaused)
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
func (it *AirdropUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropUnpaused represents a Unpaused event raised by the Airdrop contract.
type AirdropUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Airdrop *AirdropFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AirdropUnpausedIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AirdropUnpausedIterator{contract: _Airdrop.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Airdrop *AirdropFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AirdropUnpaused) (event.Subscription, error) {

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropUnpaused)
				if err := _Airdrop.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Airdrop *AirdropFilterer) ParseUnpaused(log types.Log) (*AirdropUnpaused, error) {
	event := new(AirdropUnpaused)
	if err := _Airdrop.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropValidDurationUpdateIterator is returned from FilterValidDurationUpdate and is used to iterate over the raw logs and unpacked data for ValidDurationUpdate events raised by the Airdrop contract.
type AirdropValidDurationUpdateIterator struct {
	Event *AirdropValidDurationUpdate // Event containing the contract specifics and raw log

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
func (it *AirdropValidDurationUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropValidDurationUpdate)
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
		it.Event = new(AirdropValidDurationUpdate)
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
func (it *AirdropValidDurationUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropValidDurationUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropValidDurationUpdate represents a ValidDurationUpdate event raised by the Airdrop contract.
type AirdropValidDurationUpdate struct {
	Epoch            *big.Int
	PreValidDuration *big.Int
	ValidDuration    *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidDurationUpdate is a free log retrieval operation binding the contract event 0xd00c826efe8bb0d760cdb78d5e125607c6898f12e59176be76064977c239cff9.
//
// Solidity: event ValidDurationUpdate(uint256 indexed epoch, uint256 preValidDuration, uint256 validDuration)
func (_Airdrop *AirdropFilterer) FilterValidDurationUpdate(opts *bind.FilterOpts, epoch []*big.Int) (*AirdropValidDurationUpdateIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "ValidDurationUpdate", epochRule)
	if err != nil {
		return nil, err
	}
	return &AirdropValidDurationUpdateIterator{contract: _Airdrop.contract, event: "ValidDurationUpdate", logs: logs, sub: sub}, nil
}

// WatchValidDurationUpdate is a free log subscription operation binding the contract event 0xd00c826efe8bb0d760cdb78d5e125607c6898f12e59176be76064977c239cff9.
//
// Solidity: event ValidDurationUpdate(uint256 indexed epoch, uint256 preValidDuration, uint256 validDuration)
func (_Airdrop *AirdropFilterer) WatchValidDurationUpdate(opts *bind.WatchOpts, sink chan<- *AirdropValidDurationUpdate, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "ValidDurationUpdate", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropValidDurationUpdate)
				if err := _Airdrop.contract.UnpackLog(event, "ValidDurationUpdate", log); err != nil {
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

// ParseValidDurationUpdate is a log parse operation binding the contract event 0xd00c826efe8bb0d760cdb78d5e125607c6898f12e59176be76064977c239cff9.
//
// Solidity: event ValidDurationUpdate(uint256 indexed epoch, uint256 preValidDuration, uint256 validDuration)
func (_Airdrop *AirdropFilterer) ParseValidDurationUpdate(log types.Log) (*AirdropValidDurationUpdate, error) {
	event := new(AirdropValidDurationUpdate)
	if err := _Airdrop.contract.UnpackLog(event, "ValidDurationUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
