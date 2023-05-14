// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VenusERC4626Reinvest

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

// VenusERC4626ReinvestMetaData contains all meta data concerning the VenusERC4626Reinvest contract.
var VenusERC4626ReinvestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"asset_\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"reward_\",\"type\":\"address\"},{\"internalType\":\"contractICERC20\",\"name\":\"cToken_\",\"type\":\"address\"},{\"internalType\":\"contractIComptroller\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorCode\",\"type\":\"uint256\"}],\"name\":\"CompoundERC4626__CompoundError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SwapInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pair1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pair2\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cToken\",\"outputs\":[{\"internalType\":\"contractICERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractIComptroller\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardsAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reward\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pair1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pair2\",\"type\":\"address\"}],\"name\":\"setRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VenusERC4626ReinvestABI is the input ABI used to generate the binding from.
// Deprecated: Use VenusERC4626ReinvestMetaData.ABI instead.
var VenusERC4626ReinvestABI = VenusERC4626ReinvestMetaData.ABI

// VenusERC4626Reinvest is an auto generated Go binding around an Ethereum contract.
type VenusERC4626Reinvest struct {
	VenusERC4626ReinvestCaller     // Read-only binding to the contract
	VenusERC4626ReinvestTransactor // Write-only binding to the contract
	VenusERC4626ReinvestFilterer   // Log filterer for contract events
}

// VenusERC4626ReinvestCaller is an auto generated read-only Go binding around an Ethereum contract.
type VenusERC4626ReinvestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusERC4626ReinvestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VenusERC4626ReinvestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusERC4626ReinvestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VenusERC4626ReinvestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusERC4626ReinvestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VenusERC4626ReinvestSession struct {
	Contract     *VenusERC4626Reinvest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VenusERC4626ReinvestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VenusERC4626ReinvestCallerSession struct {
	Contract *VenusERC4626ReinvestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// VenusERC4626ReinvestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VenusERC4626ReinvestTransactorSession struct {
	Contract     *VenusERC4626ReinvestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// VenusERC4626ReinvestRaw is an auto generated low-level Go binding around an Ethereum contract.
type VenusERC4626ReinvestRaw struct {
	Contract *VenusERC4626Reinvest // Generic contract binding to access the raw methods on
}

// VenusERC4626ReinvestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VenusERC4626ReinvestCallerRaw struct {
	Contract *VenusERC4626ReinvestCaller // Generic read-only contract binding to access the raw methods on
}

// VenusERC4626ReinvestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VenusERC4626ReinvestTransactorRaw struct {
	Contract *VenusERC4626ReinvestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVenusERC4626Reinvest creates a new instance of VenusERC4626Reinvest, bound to a specific deployed contract.
func NewVenusERC4626Reinvest(address common.Address, backend bind.ContractBackend) (*VenusERC4626Reinvest, error) {
	contract, err := bindVenusERC4626Reinvest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VenusERC4626Reinvest{VenusERC4626ReinvestCaller: VenusERC4626ReinvestCaller{contract: contract}, VenusERC4626ReinvestTransactor: VenusERC4626ReinvestTransactor{contract: contract}, VenusERC4626ReinvestFilterer: VenusERC4626ReinvestFilterer{contract: contract}}, nil
}

// NewVenusERC4626ReinvestCaller creates a new read-only instance of VenusERC4626Reinvest, bound to a specific deployed contract.
func NewVenusERC4626ReinvestCaller(address common.Address, caller bind.ContractCaller) (*VenusERC4626ReinvestCaller, error) {
	contract, err := bindVenusERC4626Reinvest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VenusERC4626ReinvestCaller{contract: contract}, nil
}

// NewVenusERC4626ReinvestTransactor creates a new write-only instance of VenusERC4626Reinvest, bound to a specific deployed contract.
func NewVenusERC4626ReinvestTransactor(address common.Address, transactor bind.ContractTransactor) (*VenusERC4626ReinvestTransactor, error) {
	contract, err := bindVenusERC4626Reinvest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VenusERC4626ReinvestTransactor{contract: contract}, nil
}

// NewVenusERC4626ReinvestFilterer creates a new log filterer instance of VenusERC4626Reinvest, bound to a specific deployed contract.
func NewVenusERC4626ReinvestFilterer(address common.Address, filterer bind.ContractFilterer) (*VenusERC4626ReinvestFilterer, error) {
	contract, err := bindVenusERC4626Reinvest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VenusERC4626ReinvestFilterer{contract: contract}, nil
}

// bindVenusERC4626Reinvest binds a generic wrapper to an already deployed contract.
func bindVenusERC4626Reinvest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VenusERC4626ReinvestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VenusERC4626Reinvest *VenusERC4626ReinvestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VenusERC4626Reinvest.Contract.VenusERC4626ReinvestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VenusERC4626Reinvest *VenusERC4626ReinvestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.VenusERC4626ReinvestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VenusERC4626Reinvest *VenusERC4626ReinvestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.VenusERC4626ReinvestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VenusERC4626Reinvest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _VenusERC4626Reinvest.Contract.DOMAINSEPARATOR(&_VenusERC4626Reinvest.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _VenusERC4626Reinvest.Contract.DOMAINSEPARATOR(&_VenusERC4626Reinvest.CallOpts)
}

// SwapInfo is a free data retrieval call binding the contract method 0x53104b8e.
//
// Solidity: function SwapInfo() view returns(address token, address pair1, address pair2)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) SwapInfo(opts *bind.CallOpts) (struct {
	Token common.Address
	Pair1 common.Address
	Pair2 common.Address
}, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "SwapInfo")

	outstruct := new(struct {
		Token common.Address
		Pair1 common.Address
		Pair2 common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Pair1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Pair2 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SwapInfo is a free data retrieval call binding the contract method 0x53104b8e.
//
// Solidity: function SwapInfo() view returns(address token, address pair1, address pair2)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) SwapInfo() (struct {
	Token common.Address
	Pair1 common.Address
	Pair2 common.Address
}, error) {
	return _VenusERC4626Reinvest.Contract.SwapInfo(&_VenusERC4626Reinvest.CallOpts)
}

// SwapInfo is a free data retrieval call binding the contract method 0x53104b8e.
//
// Solidity: function SwapInfo() view returns(address token, address pair1, address pair2)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) SwapInfo() (struct {
	Token common.Address
	Pair1 common.Address
	Pair2 common.Address
}, error) {
	return _VenusERC4626Reinvest.Contract.SwapInfo(&_VenusERC4626Reinvest.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.Allowance(&_VenusERC4626Reinvest.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.Allowance(&_VenusERC4626Reinvest.CallOpts, arg0, arg1)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Asset() (common.Address, error) {
	return _VenusERC4626Reinvest.Contract.Asset(&_VenusERC4626Reinvest.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) Asset() (common.Address, error) {
	return _VenusERC4626Reinvest.Contract.Asset(&_VenusERC4626Reinvest.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.BalanceOf(&_VenusERC4626Reinvest.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.BalanceOf(&_VenusERC4626Reinvest.CallOpts, arg0)
}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) CToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "cToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) CToken() (common.Address, error) {
	return _VenusERC4626Reinvest.Contract.CToken(&_VenusERC4626Reinvest.CallOpts)
}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) CToken() (common.Address, error) {
	return _VenusERC4626Reinvest.Contract.CToken(&_VenusERC4626Reinvest.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Comptroller() (common.Address, error) {
	return _VenusERC4626Reinvest.Contract.Comptroller(&_VenusERC4626Reinvest.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) Comptroller() (common.Address, error) {
	return _VenusERC4626Reinvest.Contract.Comptroller(&_VenusERC4626Reinvest.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.ConvertToAssets(&_VenusERC4626Reinvest.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.ConvertToAssets(&_VenusERC4626Reinvest.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.ConvertToShares(&_VenusERC4626Reinvest.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.ConvertToShares(&_VenusERC4626Reinvest.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Decimals() (uint8, error) {
	return _VenusERC4626Reinvest.Contract.Decimals(&_VenusERC4626Reinvest.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) Decimals() (uint8, error) {
	return _VenusERC4626Reinvest.Contract.Decimals(&_VenusERC4626Reinvest.CallOpts)
}

// GetRewardsAccrued is a free data retrieval call binding the contract method 0x609355b2.
//
// Solidity: function getRewardsAccrued() view returns(uint256 amount)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) GetRewardsAccrued(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "getRewardsAccrued")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardsAccrued is a free data retrieval call binding the contract method 0x609355b2.
//
// Solidity: function getRewardsAccrued() view returns(uint256 amount)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) GetRewardsAccrued() (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.GetRewardsAccrued(&_VenusERC4626Reinvest.CallOpts)
}

// GetRewardsAccrued is a free data retrieval call binding the contract method 0x609355b2.
//
// Solidity: function getRewardsAccrued() view returns(uint256 amount)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) GetRewardsAccrued() (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.GetRewardsAccrued(&_VenusERC4626Reinvest.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Manager() (common.Address, error) {
	return _VenusERC4626Reinvest.Contract.Manager(&_VenusERC4626Reinvest.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) Manager() (common.Address, error) {
	return _VenusERC4626Reinvest.Contract.Manager(&_VenusERC4626Reinvest.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.MaxDeposit(&_VenusERC4626Reinvest.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.MaxDeposit(&_VenusERC4626Reinvest.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.MaxMint(&_VenusERC4626Reinvest.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.MaxMint(&_VenusERC4626Reinvest.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.MaxRedeem(&_VenusERC4626Reinvest.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.MaxRedeem(&_VenusERC4626Reinvest.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.MaxWithdraw(&_VenusERC4626Reinvest.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.MaxWithdraw(&_VenusERC4626Reinvest.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Name() (string, error) {
	return _VenusERC4626Reinvest.Contract.Name(&_VenusERC4626Reinvest.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) Name() (string, error) {
	return _VenusERC4626Reinvest.Contract.Name(&_VenusERC4626Reinvest.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.Nonces(&_VenusERC4626Reinvest.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.Nonces(&_VenusERC4626Reinvest.CallOpts, arg0)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.PreviewDeposit(&_VenusERC4626Reinvest.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.PreviewDeposit(&_VenusERC4626Reinvest.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.PreviewMint(&_VenusERC4626Reinvest.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.PreviewMint(&_VenusERC4626Reinvest.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.PreviewRedeem(&_VenusERC4626Reinvest.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.PreviewRedeem(&_VenusERC4626Reinvest.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.PreviewWithdraw(&_VenusERC4626Reinvest.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.PreviewWithdraw(&_VenusERC4626Reinvest.CallOpts, assets)
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) Reward(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "reward")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Reward() (common.Address, error) {
	return _VenusERC4626Reinvest.Contract.Reward(&_VenusERC4626Reinvest.CallOpts)
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() view returns(address)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) Reward() (common.Address, error) {
	return _VenusERC4626Reinvest.Contract.Reward(&_VenusERC4626Reinvest.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Symbol() (string, error) {
	return _VenusERC4626Reinvest.Contract.Symbol(&_VenusERC4626Reinvest.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) Symbol() (string, error) {
	return _VenusERC4626Reinvest.Contract.Symbol(&_VenusERC4626Reinvest.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) TotalAssets() (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.TotalAssets(&_VenusERC4626Reinvest.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) TotalAssets() (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.TotalAssets(&_VenusERC4626Reinvest.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusERC4626Reinvest.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) TotalSupply() (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.TotalSupply(&_VenusERC4626Reinvest.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestCallerSession) TotalSupply() (*big.Int, error) {
	return _VenusERC4626Reinvest.Contract.TotalSupply(&_VenusERC4626Reinvest.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Approve(&_VenusERC4626Reinvest.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Approve(&_VenusERC4626Reinvest.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Deposit(&_VenusERC4626Reinvest.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Deposit(&_VenusERC4626Reinvest.TransactOpts, assets, receiver)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactor) Harvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.contract.Transact(opts, "harvest")
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Harvest() (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Harvest(&_VenusERC4626Reinvest.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns()
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactorSession) Harvest() (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Harvest(&_VenusERC4626Reinvest.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Mint(&_VenusERC4626Reinvest.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Mint(&_VenusERC4626Reinvest.TransactOpts, shares, receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Permit(&_VenusERC4626Reinvest.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Permit(&_VenusERC4626Reinvest.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Redeem(&_VenusERC4626Reinvest.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Redeem(&_VenusERC4626Reinvest.TransactOpts, shares, receiver, owner)
}

// SetRoute is a paid mutator transaction binding the contract method 0xec42be77.
//
// Solidity: function setRoute(address token, address pair1, address pair2) returns()
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactor) SetRoute(opts *bind.TransactOpts, token common.Address, pair1 common.Address, pair2 common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.contract.Transact(opts, "setRoute", token, pair1, pair2)
}

// SetRoute is a paid mutator transaction binding the contract method 0xec42be77.
//
// Solidity: function setRoute(address token, address pair1, address pair2) returns()
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) SetRoute(token common.Address, pair1 common.Address, pair2 common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.SetRoute(&_VenusERC4626Reinvest.TransactOpts, token, pair1, pair2)
}

// SetRoute is a paid mutator transaction binding the contract method 0xec42be77.
//
// Solidity: function setRoute(address token, address pair1, address pair2) returns()
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactorSession) SetRoute(token common.Address, pair1 common.Address, pair2 common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.SetRoute(&_VenusERC4626Reinvest.TransactOpts, token, pair1, pair2)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Transfer(&_VenusERC4626Reinvest.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Transfer(&_VenusERC4626Reinvest.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.TransferFrom(&_VenusERC4626Reinvest.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.TransferFrom(&_VenusERC4626Reinvest.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Withdraw(&_VenusERC4626Reinvest.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _VenusERC4626Reinvest.Contract.Withdraw(&_VenusERC4626Reinvest.TransactOpts, assets, receiver, owner)
}

// VenusERC4626ReinvestApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VenusERC4626Reinvest contract.
type VenusERC4626ReinvestApprovalIterator struct {
	Event *VenusERC4626ReinvestApproval // Event containing the contract specifics and raw log

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
func (it *VenusERC4626ReinvestApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusERC4626ReinvestApproval)
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
		it.Event = new(VenusERC4626ReinvestApproval)
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
func (it *VenusERC4626ReinvestApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusERC4626ReinvestApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusERC4626ReinvestApproval represents a Approval event raised by the VenusERC4626Reinvest contract.
type VenusERC4626ReinvestApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VenusERC4626ReinvestApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VenusERC4626Reinvest.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VenusERC4626ReinvestApprovalIterator{contract: _VenusERC4626Reinvest.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VenusERC4626ReinvestApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VenusERC4626Reinvest.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusERC4626ReinvestApproval)
				if err := _VenusERC4626Reinvest.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestFilterer) ParseApproval(log types.Log) (*VenusERC4626ReinvestApproval, error) {
	event := new(VenusERC4626ReinvestApproval)
	if err := _VenusERC4626Reinvest.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusERC4626ReinvestDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the VenusERC4626Reinvest contract.
type VenusERC4626ReinvestDepositIterator struct {
	Event *VenusERC4626ReinvestDeposit // Event containing the contract specifics and raw log

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
func (it *VenusERC4626ReinvestDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusERC4626ReinvestDeposit)
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
		it.Event = new(VenusERC4626ReinvestDeposit)
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
func (it *VenusERC4626ReinvestDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusERC4626ReinvestDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusERC4626ReinvestDeposit represents a Deposit event raised by the VenusERC4626Reinvest contract.
type VenusERC4626ReinvestDeposit struct {
	Caller common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestFilterer) FilterDeposit(opts *bind.FilterOpts, caller []common.Address, owner []common.Address) (*VenusERC4626ReinvestDepositIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VenusERC4626Reinvest.contract.FilterLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &VenusERC4626ReinvestDepositIterator{contract: _VenusERC4626Reinvest.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VenusERC4626ReinvestDeposit, caller []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VenusERC4626Reinvest.contract.WatchLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusERC4626ReinvestDeposit)
				if err := _VenusERC4626Reinvest.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestFilterer) ParseDeposit(log types.Log) (*VenusERC4626ReinvestDeposit, error) {
	event := new(VenusERC4626ReinvestDeposit)
	if err := _VenusERC4626Reinvest.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusERC4626ReinvestTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VenusERC4626Reinvest contract.
type VenusERC4626ReinvestTransferIterator struct {
	Event *VenusERC4626ReinvestTransfer // Event containing the contract specifics and raw log

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
func (it *VenusERC4626ReinvestTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusERC4626ReinvestTransfer)
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
		it.Event = new(VenusERC4626ReinvestTransfer)
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
func (it *VenusERC4626ReinvestTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusERC4626ReinvestTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusERC4626ReinvestTransfer represents a Transfer event raised by the VenusERC4626Reinvest contract.
type VenusERC4626ReinvestTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VenusERC4626ReinvestTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VenusERC4626Reinvest.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VenusERC4626ReinvestTransferIterator{contract: _VenusERC4626Reinvest.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VenusERC4626ReinvestTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VenusERC4626Reinvest.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusERC4626ReinvestTransfer)
				if err := _VenusERC4626Reinvest.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestFilterer) ParseTransfer(log types.Log) (*VenusERC4626ReinvestTransfer, error) {
	event := new(VenusERC4626ReinvestTransfer)
	if err := _VenusERC4626Reinvest.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusERC4626ReinvestWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the VenusERC4626Reinvest contract.
type VenusERC4626ReinvestWithdrawIterator struct {
	Event *VenusERC4626ReinvestWithdraw // Event containing the contract specifics and raw log

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
func (it *VenusERC4626ReinvestWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusERC4626ReinvestWithdraw)
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
		it.Event = new(VenusERC4626ReinvestWithdraw)
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
func (it *VenusERC4626ReinvestWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusERC4626ReinvestWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusERC4626ReinvestWithdraw represents a Withdraw event raised by the VenusERC4626Reinvest contract.
type VenusERC4626ReinvestWithdraw struct {
	Caller   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestFilterer) FilterWithdraw(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, owner []common.Address) (*VenusERC4626ReinvestWithdrawIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VenusERC4626Reinvest.contract.FilterLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &VenusERC4626ReinvestWithdrawIterator{contract: _VenusERC4626Reinvest.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *VenusERC4626ReinvestWithdraw, caller []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VenusERC4626Reinvest.contract.WatchLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusERC4626ReinvestWithdraw)
				if err := _VenusERC4626Reinvest.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_VenusERC4626Reinvest *VenusERC4626ReinvestFilterer) ParseWithdraw(log types.Log) (*VenusERC4626ReinvestWithdraw, error) {
	event := new(VenusERC4626ReinvestWithdraw)
	if err := _VenusERC4626Reinvest.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
