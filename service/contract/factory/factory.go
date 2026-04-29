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

// FactoryStrategy is an auto generated low-level Go binding around an user-defined struct.
type FactoryStrategy struct {
	Name              string
	Symbol            string
	Vault             common.Address
	CuniBTC           common.Address
	DelayRedeemRouter common.Address
	Airdrop           common.Address
}

// FactoryMetaData contains all meta data concerning the Factory contract.
var FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImpl\",\"type\":\"address\"}],\"name\":\"BeaconUpgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cuniBTC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delayRedeemRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"airdrop\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structFactory.Strategy\",\"name\":\"strategy\",\"type\":\"tuple\"}],\"name\":\"StrategyCreate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"airdropBeacon\",\"outputs\":[{\"internalType\":\"contractUpgradeableBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"airdropImpl\",\"outputs\":[{\"internalType\":\"contractAirdrop\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uniBTC\",\"type\":\"address\"}],\"name\":\"createStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cuniBTCBeacon\",\"outputs\":[{\"internalType\":\"contractUpgradeableBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cuniBTCImpl\",\"outputs\":[{\"internalType\":\"contractcuniBTC\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayRedeemRouterBeacon\",\"outputs\":[{\"internalType\":\"contractUpgradeableBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayRedeemRouterImpl\",\"outputs\":[{\"internalType\":\"contractDelayRedeemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cuniBTCImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_airdropImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delayRedeemRouterImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"strategies\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cuniBTC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delayRedeemRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"airdrop\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beacon\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"upgradeBeacon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultBeacon\",\"outputs\":[{\"internalType\":\"contractUpgradeableBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultImpl\",\"outputs\":[{\"internalType\":\"contractVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// AirdropBeacon is a free data retrieval call binding the contract method 0x30152558.
//
// Solidity: function airdropBeacon() view returns(address)
func (_Factory *FactoryCaller) AirdropBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "airdropBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AirdropBeacon is a free data retrieval call binding the contract method 0x30152558.
//
// Solidity: function airdropBeacon() view returns(address)
func (_Factory *FactorySession) AirdropBeacon() (common.Address, error) {
	return _Factory.Contract.AirdropBeacon(&_Factory.CallOpts)
}

// AirdropBeacon is a free data retrieval call binding the contract method 0x30152558.
//
// Solidity: function airdropBeacon() view returns(address)
func (_Factory *FactoryCallerSession) AirdropBeacon() (common.Address, error) {
	return _Factory.Contract.AirdropBeacon(&_Factory.CallOpts)
}

// AirdropImpl is a free data retrieval call binding the contract method 0x4ed4f1da.
//
// Solidity: function airdropImpl() view returns(address)
func (_Factory *FactoryCaller) AirdropImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "airdropImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AirdropImpl is a free data retrieval call binding the contract method 0x4ed4f1da.
//
// Solidity: function airdropImpl() view returns(address)
func (_Factory *FactorySession) AirdropImpl() (common.Address, error) {
	return _Factory.Contract.AirdropImpl(&_Factory.CallOpts)
}

// AirdropImpl is a free data retrieval call binding the contract method 0x4ed4f1da.
//
// Solidity: function airdropImpl() view returns(address)
func (_Factory *FactoryCallerSession) AirdropImpl() (common.Address, error) {
	return _Factory.Contract.AirdropImpl(&_Factory.CallOpts)
}

// CuniBTCBeacon is a free data retrieval call binding the contract method 0x929c1ca3.
//
// Solidity: function cuniBTCBeacon() view returns(address)
func (_Factory *FactoryCaller) CuniBTCBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "cuniBTCBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CuniBTCBeacon is a free data retrieval call binding the contract method 0x929c1ca3.
//
// Solidity: function cuniBTCBeacon() view returns(address)
func (_Factory *FactorySession) CuniBTCBeacon() (common.Address, error) {
	return _Factory.Contract.CuniBTCBeacon(&_Factory.CallOpts)
}

// CuniBTCBeacon is a free data retrieval call binding the contract method 0x929c1ca3.
//
// Solidity: function cuniBTCBeacon() view returns(address)
func (_Factory *FactoryCallerSession) CuniBTCBeacon() (common.Address, error) {
	return _Factory.Contract.CuniBTCBeacon(&_Factory.CallOpts)
}

// CuniBTCImpl is a free data retrieval call binding the contract method 0x07945a79.
//
// Solidity: function cuniBTCImpl() view returns(address)
func (_Factory *FactoryCaller) CuniBTCImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "cuniBTCImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CuniBTCImpl is a free data retrieval call binding the contract method 0x07945a79.
//
// Solidity: function cuniBTCImpl() view returns(address)
func (_Factory *FactorySession) CuniBTCImpl() (common.Address, error) {
	return _Factory.Contract.CuniBTCImpl(&_Factory.CallOpts)
}

// CuniBTCImpl is a free data retrieval call binding the contract method 0x07945a79.
//
// Solidity: function cuniBTCImpl() view returns(address)
func (_Factory *FactoryCallerSession) CuniBTCImpl() (common.Address, error) {
	return _Factory.Contract.CuniBTCImpl(&_Factory.CallOpts)
}

// DelayRedeemRouterBeacon is a free data retrieval call binding the contract method 0xd4f89e96.
//
// Solidity: function delayRedeemRouterBeacon() view returns(address)
func (_Factory *FactoryCaller) DelayRedeemRouterBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "delayRedeemRouterBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayRedeemRouterBeacon is a free data retrieval call binding the contract method 0xd4f89e96.
//
// Solidity: function delayRedeemRouterBeacon() view returns(address)
func (_Factory *FactorySession) DelayRedeemRouterBeacon() (common.Address, error) {
	return _Factory.Contract.DelayRedeemRouterBeacon(&_Factory.CallOpts)
}

// DelayRedeemRouterBeacon is a free data retrieval call binding the contract method 0xd4f89e96.
//
// Solidity: function delayRedeemRouterBeacon() view returns(address)
func (_Factory *FactoryCallerSession) DelayRedeemRouterBeacon() (common.Address, error) {
	return _Factory.Contract.DelayRedeemRouterBeacon(&_Factory.CallOpts)
}

// DelayRedeemRouterImpl is a free data retrieval call binding the contract method 0x1e8f51f7.
//
// Solidity: function delayRedeemRouterImpl() view returns(address)
func (_Factory *FactoryCaller) DelayRedeemRouterImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "delayRedeemRouterImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayRedeemRouterImpl is a free data retrieval call binding the contract method 0x1e8f51f7.
//
// Solidity: function delayRedeemRouterImpl() view returns(address)
func (_Factory *FactorySession) DelayRedeemRouterImpl() (common.Address, error) {
	return _Factory.Contract.DelayRedeemRouterImpl(&_Factory.CallOpts)
}

// DelayRedeemRouterImpl is a free data retrieval call binding the contract method 0x1e8f51f7.
//
// Solidity: function delayRedeemRouterImpl() view returns(address)
func (_Factory *FactoryCallerSession) DelayRedeemRouterImpl() (common.Address, error) {
	return _Factory.Contract.DelayRedeemRouterImpl(&_Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Factory *FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Factory *FactorySession) Owner() (common.Address, error) {
	return _Factory.Contract.Owner(&_Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Factory *FactoryCallerSession) Owner() (common.Address, error) {
	return _Factory.Contract.Owner(&_Factory.CallOpts)
}

// Strategies is a free data retrieval call binding the contract method 0x780f1acd.
//
// Solidity: function strategies(string ) view returns(string name, string symbol, address vault, address cuniBTC, address delayRedeemRouter, address airdrop)
func (_Factory *FactoryCaller) Strategies(opts *bind.CallOpts, arg0 string) (struct {
	Name              string
	Symbol            string
	Vault             common.Address
	CuniBTC           common.Address
	DelayRedeemRouter common.Address
	Airdrop           common.Address
}, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "strategies", arg0)

	outstruct := new(struct {
		Name              string
		Symbol            string
		Vault             common.Address
		CuniBTC           common.Address
		DelayRedeemRouter common.Address
		Airdrop           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Vault = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CuniBTC = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.DelayRedeemRouter = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Airdrop = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Strategies is a free data retrieval call binding the contract method 0x780f1acd.
//
// Solidity: function strategies(string ) view returns(string name, string symbol, address vault, address cuniBTC, address delayRedeemRouter, address airdrop)
func (_Factory *FactorySession) Strategies(arg0 string) (struct {
	Name              string
	Symbol            string
	Vault             common.Address
	CuniBTC           common.Address
	DelayRedeemRouter common.Address
	Airdrop           common.Address
}, error) {
	return _Factory.Contract.Strategies(&_Factory.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0x780f1acd.
//
// Solidity: function strategies(string ) view returns(string name, string symbol, address vault, address cuniBTC, address delayRedeemRouter, address airdrop)
func (_Factory *FactoryCallerSession) Strategies(arg0 string) (struct {
	Name              string
	Symbol            string
	Vault             common.Address
	CuniBTC           common.Address
	DelayRedeemRouter common.Address
	Airdrop           common.Address
}, error) {
	return _Factory.Contract.Strategies(&_Factory.CallOpts, arg0)
}

// VaultBeacon is a free data retrieval call binding the contract method 0x9d343be1.
//
// Solidity: function vaultBeacon() view returns(address)
func (_Factory *FactoryCaller) VaultBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "vaultBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultBeacon is a free data retrieval call binding the contract method 0x9d343be1.
//
// Solidity: function vaultBeacon() view returns(address)
func (_Factory *FactorySession) VaultBeacon() (common.Address, error) {
	return _Factory.Contract.VaultBeacon(&_Factory.CallOpts)
}

// VaultBeacon is a free data retrieval call binding the contract method 0x9d343be1.
//
// Solidity: function vaultBeacon() view returns(address)
func (_Factory *FactoryCallerSession) VaultBeacon() (common.Address, error) {
	return _Factory.Contract.VaultBeacon(&_Factory.CallOpts)
}

// VaultImpl is a free data retrieval call binding the contract method 0xec1e31a0.
//
// Solidity: function vaultImpl() view returns(address)
func (_Factory *FactoryCaller) VaultImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "vaultImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultImpl is a free data retrieval call binding the contract method 0xec1e31a0.
//
// Solidity: function vaultImpl() view returns(address)
func (_Factory *FactorySession) VaultImpl() (common.Address, error) {
	return _Factory.Contract.VaultImpl(&_Factory.CallOpts)
}

// VaultImpl is a free data retrieval call binding the contract method 0xec1e31a0.
//
// Solidity: function vaultImpl() view returns(address)
func (_Factory *FactoryCallerSession) VaultImpl() (common.Address, error) {
	return _Factory.Contract.VaultImpl(&_Factory.CallOpts)
}

// CreateStrategy is a paid mutator transaction binding the contract method 0xb462eda8.
//
// Solidity: function createStrategy(string _name, string _symbol, address _defaultAdmin, address _uniBTC) returns()
func (_Factory *FactoryTransactor) CreateStrategy(opts *bind.TransactOpts, _name string, _symbol string, _defaultAdmin common.Address, _uniBTC common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createStrategy", _name, _symbol, _defaultAdmin, _uniBTC)
}

// CreateStrategy is a paid mutator transaction binding the contract method 0xb462eda8.
//
// Solidity: function createStrategy(string _name, string _symbol, address _defaultAdmin, address _uniBTC) returns()
func (_Factory *FactorySession) CreateStrategy(_name string, _symbol string, _defaultAdmin common.Address, _uniBTC common.Address) (*types.Transaction, error) {
	return _Factory.Contract.CreateStrategy(&_Factory.TransactOpts, _name, _symbol, _defaultAdmin, _uniBTC)
}

// CreateStrategy is a paid mutator transaction binding the contract method 0xb462eda8.
//
// Solidity: function createStrategy(string _name, string _symbol, address _defaultAdmin, address _uniBTC) returns()
func (_Factory *FactoryTransactorSession) CreateStrategy(_name string, _symbol string, _defaultAdmin common.Address, _uniBTC common.Address) (*types.Transaction, error) {
	return _Factory.Contract.CreateStrategy(&_Factory.TransactOpts, _name, _symbol, _defaultAdmin, _uniBTC)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _cuniBTCImpl, address _vaultImpl, address _airdropImpl, address _delayRedeemRouterImpl) returns()
func (_Factory *FactoryTransactor) Initialize(opts *bind.TransactOpts, _cuniBTCImpl common.Address, _vaultImpl common.Address, _airdropImpl common.Address, _delayRedeemRouterImpl common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "initialize", _cuniBTCImpl, _vaultImpl, _airdropImpl, _delayRedeemRouterImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _cuniBTCImpl, address _vaultImpl, address _airdropImpl, address _delayRedeemRouterImpl) returns()
func (_Factory *FactorySession) Initialize(_cuniBTCImpl common.Address, _vaultImpl common.Address, _airdropImpl common.Address, _delayRedeemRouterImpl common.Address) (*types.Transaction, error) {
	return _Factory.Contract.Initialize(&_Factory.TransactOpts, _cuniBTCImpl, _vaultImpl, _airdropImpl, _delayRedeemRouterImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _cuniBTCImpl, address _vaultImpl, address _airdropImpl, address _delayRedeemRouterImpl) returns()
func (_Factory *FactoryTransactorSession) Initialize(_cuniBTCImpl common.Address, _vaultImpl common.Address, _airdropImpl common.Address, _delayRedeemRouterImpl common.Address) (*types.Transaction, error) {
	return _Factory.Contract.Initialize(&_Factory.TransactOpts, _cuniBTCImpl, _vaultImpl, _airdropImpl, _delayRedeemRouterImpl)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Factory *FactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Factory *FactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Factory.Contract.RenounceOwnership(&_Factory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Factory *FactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Factory.Contract.RenounceOwnership(&_Factory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Factory *FactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Factory *FactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Factory.Contract.TransferOwnership(&_Factory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Factory *FactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Factory.Contract.TransferOwnership(&_Factory.TransactOpts, newOwner)
}

// UpgradeBeacon is a paid mutator transaction binding the contract method 0x848bf918.
//
// Solidity: function upgradeBeacon(address _beacon, address _newImpl) returns()
func (_Factory *FactoryTransactor) UpgradeBeacon(opts *bind.TransactOpts, _beacon common.Address, _newImpl common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "upgradeBeacon", _beacon, _newImpl)
}

// UpgradeBeacon is a paid mutator transaction binding the contract method 0x848bf918.
//
// Solidity: function upgradeBeacon(address _beacon, address _newImpl) returns()
func (_Factory *FactorySession) UpgradeBeacon(_beacon common.Address, _newImpl common.Address) (*types.Transaction, error) {
	return _Factory.Contract.UpgradeBeacon(&_Factory.TransactOpts, _beacon, _newImpl)
}

// UpgradeBeacon is a paid mutator transaction binding the contract method 0x848bf918.
//
// Solidity: function upgradeBeacon(address _beacon, address _newImpl) returns()
func (_Factory *FactoryTransactorSession) UpgradeBeacon(_beacon common.Address, _newImpl common.Address) (*types.Transaction, error) {
	return _Factory.Contract.UpgradeBeacon(&_Factory.TransactOpts, _beacon, _newImpl)
}

// FactoryBeaconUpgradeIterator is returned from FilterBeaconUpgrade and is used to iterate over the raw logs and unpacked data for BeaconUpgrade events raised by the Factory contract.
type FactoryBeaconUpgradeIterator struct {
	Event *FactoryBeaconUpgrade // Event containing the contract specifics and raw log

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
func (it *FactoryBeaconUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryBeaconUpgrade)
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
		it.Event = new(FactoryBeaconUpgrade)
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
func (it *FactoryBeaconUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryBeaconUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryBeaconUpgrade represents a BeaconUpgrade event raised by the Factory contract.
type FactoryBeaconUpgrade struct {
	Beacon  common.Address
	NewImpl common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgrade is a free log retrieval operation binding the contract event 0x9b7ca51e4c4c7938ffb6cf03b53172b43f052d2f3ba1fea4c2056de92f4996c0.
//
// Solidity: event BeaconUpgrade(address beacon, address newImpl)
func (_Factory *FactoryFilterer) FilterBeaconUpgrade(opts *bind.FilterOpts) (*FactoryBeaconUpgradeIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "BeaconUpgrade")
	if err != nil {
		return nil, err
	}
	return &FactoryBeaconUpgradeIterator{contract: _Factory.contract, event: "BeaconUpgrade", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgrade is a free log subscription operation binding the contract event 0x9b7ca51e4c4c7938ffb6cf03b53172b43f052d2f3ba1fea4c2056de92f4996c0.
//
// Solidity: event BeaconUpgrade(address beacon, address newImpl)
func (_Factory *FactoryFilterer) WatchBeaconUpgrade(opts *bind.WatchOpts, sink chan<- *FactoryBeaconUpgrade) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "BeaconUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryBeaconUpgrade)
				if err := _Factory.contract.UnpackLog(event, "BeaconUpgrade", log); err != nil {
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

// ParseBeaconUpgrade is a log parse operation binding the contract event 0x9b7ca51e4c4c7938ffb6cf03b53172b43f052d2f3ba1fea4c2056de92f4996c0.
//
// Solidity: event BeaconUpgrade(address beacon, address newImpl)
func (_Factory *FactoryFilterer) ParseBeaconUpgrade(log types.Log) (*FactoryBeaconUpgrade, error) {
	event := new(FactoryBeaconUpgrade)
	if err := _Factory.contract.UnpackLog(event, "BeaconUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Factory contract.
type FactoryInitializedIterator struct {
	Event *FactoryInitialized // Event containing the contract specifics and raw log

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
func (it *FactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryInitialized)
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
		it.Event = new(FactoryInitialized)
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
func (it *FactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryInitialized represents a Initialized event raised by the Factory contract.
type FactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Factory *FactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*FactoryInitializedIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FactoryInitializedIterator{contract: _Factory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Factory *FactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryInitialized)
				if err := _Factory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Factory *FactoryFilterer) ParseInitialized(log types.Log) (*FactoryInitialized, error) {
	event := new(FactoryInitialized)
	if err := _Factory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Factory contract.
type FactoryOwnershipTransferredIterator struct {
	Event *FactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryOwnershipTransferred)
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
		it.Event = new(FactoryOwnershipTransferred)
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
func (it *FactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Factory contract.
type FactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Factory *FactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FactoryOwnershipTransferredIterator{contract: _Factory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Factory *FactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryOwnershipTransferred)
				if err := _Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Factory *FactoryFilterer) ParseOwnershipTransferred(log types.Log) (*FactoryOwnershipTransferred, error) {
	event := new(FactoryOwnershipTransferred)
	if err := _Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryStrategyCreateIterator is returned from FilterStrategyCreate and is used to iterate over the raw logs and unpacked data for StrategyCreate events raised by the Factory contract.
type FactoryStrategyCreateIterator struct {
	Event *FactoryStrategyCreate // Event containing the contract specifics and raw log

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
func (it *FactoryStrategyCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryStrategyCreate)
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
		it.Event = new(FactoryStrategyCreate)
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
func (it *FactoryStrategyCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryStrategyCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryStrategyCreate represents a StrategyCreate event raised by the Factory contract.
type FactoryStrategyCreate struct {
	Symbol   string
	Strategy FactoryStrategy
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyCreate is a free log retrieval operation binding the contract event 0x89147275d4bf3d28e2a437be0d4eb5ee4c4cfcc957e43a14c63f604ba7df6bd9.
//
// Solidity: event StrategyCreate(string symbol, (string,string,address,address,address,address) strategy)
func (_Factory *FactoryFilterer) FilterStrategyCreate(opts *bind.FilterOpts) (*FactoryStrategyCreateIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "StrategyCreate")
	if err != nil {
		return nil, err
	}
	return &FactoryStrategyCreateIterator{contract: _Factory.contract, event: "StrategyCreate", logs: logs, sub: sub}, nil
}

// WatchStrategyCreate is a free log subscription operation binding the contract event 0x89147275d4bf3d28e2a437be0d4eb5ee4c4cfcc957e43a14c63f604ba7df6bd9.
//
// Solidity: event StrategyCreate(string symbol, (string,string,address,address,address,address) strategy)
func (_Factory *FactoryFilterer) WatchStrategyCreate(opts *bind.WatchOpts, sink chan<- *FactoryStrategyCreate) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "StrategyCreate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryStrategyCreate)
				if err := _Factory.contract.UnpackLog(event, "StrategyCreate", log); err != nil {
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

// ParseStrategyCreate is a log parse operation binding the contract event 0x89147275d4bf3d28e2a437be0d4eb5ee4c4cfcc957e43a14c63f604ba7df6bd9.
//
// Solidity: event StrategyCreate(string symbol, (string,string,address,address,address,address) strategy)
func (_Factory *FactoryFilterer) ParseStrategyCreate(log types.Log) (*FactoryStrategyCreate, error) {
	event := new(FactoryStrategyCreate)
	if err := _Factory.contract.UnpackLog(event, "StrategyCreate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
