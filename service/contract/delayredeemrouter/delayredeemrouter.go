// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package delayredeemrouter

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

// DelayRedeemRouterDebtTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type DelayRedeemRouterDebtTokenAmount struct {
	Token  common.Address
	Amount *big.Int
	Fee    *big.Int
}

// DelayRedeemRouterDelayedRedeem is an auto generated low-level Go binding around an user-defined struct.
type DelayRedeemRouterDelayedRedeem struct {
	Amount    *big.Int
	CreatedAt uint32
	Token     common.Address
}

// DelayRedeemRouterUserDelayedRedeems is an auto generated low-level Go binding around an user-defined struct.
type DelayRedeemRouterUserDelayedRedeems struct {
	DelayedRedeemsCompleted *big.Int
	DelayedRedeems          []DelayRedeemRouterDelayedRedeem
}

// DelayredeemrouterMetaData contains all meta data concerning the Delayredeemrouter contract.
var DelayredeemrouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"BlacklistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"BlacklistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"BtclistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"BtclistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"CancelFeeRateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemFee\",\"type\":\"uint256\"}],\"name\":\"DelayedRedeemCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedAmount\",\"type\":\"uint256\"}],\"name\":\"DelayedRedeemsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delayedRedeemsCompleted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"}],\"name\":\"DelayedRedeemsCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"fastLaneStatus\",\"type\":\"bool[]\"}],\"name\":\"FastLaneSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ManagementFeeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousQuota\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuota\",\"type\":\"uint256\"}],\"name\":\"MaxQuotaSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousQuota\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuota\",\"type\":\"uint256\"}],\"name\":\"RateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"RedeemDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"RedeemFeeRateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"name\":\"RetainAmountsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokensPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokensUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"WhitelistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"WhitelistEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"WhitelistRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_REDEEM_FEE_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXCHANGE_RATE_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DAILY_REDEEM_CAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REDEEM_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_BTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REDEEM_FEE_RATE_RANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_IN_A_DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"addToBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"addToBtclist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"canClaimDelayedRedeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimDelayedRedeems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxNumberOfDelayedRedeemsToClaim\",\"type\":\"uint256\"}],\"name\":\"claimDelayedRedeems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createDelayedRedeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fastLanes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getClaimableUserDelayedRedeems\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"createdAt\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structDelayRedeemRouter.DelayedRedeem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getCurrentCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserDelayedRedeems\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"createdAt\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structDelayRedeemRouter.DelayedRedeem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxNumberOfDelayedRedeemsToClaim\",\"type\":\"uint256\"}],\"name\":\"getUserRedeem\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structDelayRedeemRouter.DebtTokenAmount[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uniBTC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_redeemDelay\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_whitelistEnabled\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isBtclisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isPausedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastRebaseTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managementFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxQuotas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"pauseTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"quotaBases\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"quotaRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"redeemFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"removeFromBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"removeFromBtclist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"retainAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setCancelFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_fastLaneStatus\",\"type\":\"bool[]\"}],\"name\":\"setFastLane\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_quotas\",\"type\":\"uint256[]\"}],\"name\":\"setMaxQuotaForTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_quotas\",\"type\":\"uint256[]\"}],\"name\":\"setQuotaRates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDelay\",\"type\":\"uint256\"}],\"name\":\"setRedeemDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setRedeemFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_balances\",\"type\":\"uint256[]\"}],\"name\":\"setRetainAmounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setWhitelistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenDebts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalDebts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCleared\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniBTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"unpauseTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"userDelayedRedeemByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"createdAt\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structDelayRedeemRouter.DelayedRedeem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userRedeems\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayedRedeemsCompleted\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"createdAt\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structDelayRedeemRouter.DelayedRedeem[]\",\"name\":\"delayedRedeems\",\"type\":\"tuple[]\"}],\"internalType\":\"structDelayRedeemRouter.UserDelayedRedeems\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userRedeemsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"withdrawManagementFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DelayredeemrouterABI is the input ABI used to generate the binding from.
// Deprecated: Use DelayredeemrouterMetaData.ABI instead.
var DelayredeemrouterABI = DelayredeemrouterMetaData.ABI

// Delayredeemrouter is an auto generated Go binding around an Ethereum contract.
type Delayredeemrouter struct {
	DelayredeemrouterCaller     // Read-only binding to the contract
	DelayredeemrouterTransactor // Write-only binding to the contract
	DelayredeemrouterFilterer   // Log filterer for contract events
}

// DelayredeemrouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelayredeemrouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayredeemrouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelayredeemrouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayredeemrouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelayredeemrouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayredeemrouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelayredeemrouterSession struct {
	Contract     *Delayredeemrouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DelayredeemrouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelayredeemrouterCallerSession struct {
	Contract *DelayredeemrouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DelayredeemrouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelayredeemrouterTransactorSession struct {
	Contract     *DelayredeemrouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DelayredeemrouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelayredeemrouterRaw struct {
	Contract *Delayredeemrouter // Generic contract binding to access the raw methods on
}

// DelayredeemrouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelayredeemrouterCallerRaw struct {
	Contract *DelayredeemrouterCaller // Generic read-only contract binding to access the raw methods on
}

// DelayredeemrouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelayredeemrouterTransactorRaw struct {
	Contract *DelayredeemrouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelayredeemrouter creates a new instance of Delayredeemrouter, bound to a specific deployed contract.
func NewDelayredeemrouter(address common.Address, backend bind.ContractBackend) (*Delayredeemrouter, error) {
	contract, err := bindDelayredeemrouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Delayredeemrouter{DelayredeemrouterCaller: DelayredeemrouterCaller{contract: contract}, DelayredeemrouterTransactor: DelayredeemrouterTransactor{contract: contract}, DelayredeemrouterFilterer: DelayredeemrouterFilterer{contract: contract}}, nil
}

// NewDelayredeemrouterCaller creates a new read-only instance of Delayredeemrouter, bound to a specific deployed contract.
func NewDelayredeemrouterCaller(address common.Address, caller bind.ContractCaller) (*DelayredeemrouterCaller, error) {
	contract, err := bindDelayredeemrouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterCaller{contract: contract}, nil
}

// NewDelayredeemrouterTransactor creates a new write-only instance of Delayredeemrouter, bound to a specific deployed contract.
func NewDelayredeemrouterTransactor(address common.Address, transactor bind.ContractTransactor) (*DelayredeemrouterTransactor, error) {
	contract, err := bindDelayredeemrouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterTransactor{contract: contract}, nil
}

// NewDelayredeemrouterFilterer creates a new log filterer instance of Delayredeemrouter, bound to a specific deployed contract.
func NewDelayredeemrouterFilterer(address common.Address, filterer bind.ContractFilterer) (*DelayredeemrouterFilterer, error) {
	contract, err := bindDelayredeemrouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterFilterer{contract: contract}, nil
}

// bindDelayredeemrouter binds a generic wrapper to an already deployed contract.
func bindDelayredeemrouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DelayredeemrouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delayredeemrouter *DelayredeemrouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Delayredeemrouter.Contract.DelayredeemrouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delayredeemrouter *DelayredeemrouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.DelayredeemrouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delayredeemrouter *DelayredeemrouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.DelayredeemrouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delayredeemrouter *DelayredeemrouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Delayredeemrouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delayredeemrouter *DelayredeemrouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delayredeemrouter *DelayredeemrouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Delayredeemrouter *DelayredeemrouterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Delayredeemrouter *DelayredeemrouterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Delayredeemrouter.Contract.DEFAULTADMINROLE(&_Delayredeemrouter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Delayredeemrouter.Contract.DEFAULTADMINROLE(&_Delayredeemrouter.CallOpts)
}

// DEFAULTREDEEMFEERATE is a free data retrieval call binding the contract method 0xc71aba6c.
//
// Solidity: function DEFAULT_REDEEM_FEE_RATE() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) DEFAULTREDEEMFEERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "DEFAULT_REDEEM_FEE_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTREDEEMFEERATE is a free data retrieval call binding the contract method 0xc71aba6c.
//
// Solidity: function DEFAULT_REDEEM_FEE_RATE() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) DEFAULTREDEEMFEERATE() (*big.Int, error) {
	return _Delayredeemrouter.Contract.DEFAULTREDEEMFEERATE(&_Delayredeemrouter.CallOpts)
}

// DEFAULTREDEEMFEERATE is a free data retrieval call binding the contract method 0xc71aba6c.
//
// Solidity: function DEFAULT_REDEEM_FEE_RATE() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) DEFAULTREDEEMFEERATE() (*big.Int, error) {
	return _Delayredeemrouter.Contract.DEFAULTREDEEMFEERATE(&_Delayredeemrouter.CallOpts)
}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) EXCHANGERATEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "EXCHANGE_RATE_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) EXCHANGERATEBASE() (*big.Int, error) {
	return _Delayredeemrouter.Contract.EXCHANGERATEBASE(&_Delayredeemrouter.CallOpts)
}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) EXCHANGERATEBASE() (*big.Int, error) {
	return _Delayredeemrouter.Contract.EXCHANGERATEBASE(&_Delayredeemrouter.CallOpts)
}

// MAXDAILYREDEEMCAP is a free data retrieval call binding the contract method 0xb0d381c7.
//
// Solidity: function MAX_DAILY_REDEEM_CAP() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) MAXDAILYREDEEMCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "MAX_DAILY_REDEEM_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDAILYREDEEMCAP is a free data retrieval call binding the contract method 0xb0d381c7.
//
// Solidity: function MAX_DAILY_REDEEM_CAP() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) MAXDAILYREDEEMCAP() (*big.Int, error) {
	return _Delayredeemrouter.Contract.MAXDAILYREDEEMCAP(&_Delayredeemrouter.CallOpts)
}

// MAXDAILYREDEEMCAP is a free data retrieval call binding the contract method 0xb0d381c7.
//
// Solidity: function MAX_DAILY_REDEEM_CAP() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) MAXDAILYREDEEMCAP() (*big.Int, error) {
	return _Delayredeemrouter.Contract.MAXDAILYREDEEMCAP(&_Delayredeemrouter.CallOpts)
}

// MAXREDEEMDELAY is a free data retrieval call binding the contract method 0xc5a56664.
//
// Solidity: function MAX_REDEEM_DELAY() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) MAXREDEEMDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "MAX_REDEEM_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXREDEEMDELAY is a free data retrieval call binding the contract method 0xc5a56664.
//
// Solidity: function MAX_REDEEM_DELAY() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) MAXREDEEMDELAY() (*big.Int, error) {
	return _Delayredeemrouter.Contract.MAXREDEEMDELAY(&_Delayredeemrouter.CallOpts)
}

// MAXREDEEMDELAY is a free data retrieval call binding the contract method 0xc5a56664.
//
// Solidity: function MAX_REDEEM_DELAY() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) MAXREDEEMDELAY() (*big.Int, error) {
	return _Delayredeemrouter.Contract.MAXREDEEMDELAY(&_Delayredeemrouter.CallOpts)
}

// NATIVEBTC is a free data retrieval call binding the contract method 0x3af02ff3.
//
// Solidity: function NATIVE_BTC() view returns(address)
func (_Delayredeemrouter *DelayredeemrouterCaller) NATIVEBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "NATIVE_BTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEBTC is a free data retrieval call binding the contract method 0x3af02ff3.
//
// Solidity: function NATIVE_BTC() view returns(address)
func (_Delayredeemrouter *DelayredeemrouterSession) NATIVEBTC() (common.Address, error) {
	return _Delayredeemrouter.Contract.NATIVEBTC(&_Delayredeemrouter.CallOpts)
}

// NATIVEBTC is a free data retrieval call binding the contract method 0x3af02ff3.
//
// Solidity: function NATIVE_BTC() view returns(address)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) NATIVEBTC() (common.Address, error) {
	return _Delayredeemrouter.Contract.NATIVEBTC(&_Delayredeemrouter.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Delayredeemrouter *DelayredeemrouterCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Delayredeemrouter *DelayredeemrouterSession) OPERATORROLE() ([32]byte, error) {
	return _Delayredeemrouter.Contract.OPERATORROLE(&_Delayredeemrouter.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Delayredeemrouter.Contract.OPERATORROLE(&_Delayredeemrouter.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Delayredeemrouter *DelayredeemrouterCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Delayredeemrouter *DelayredeemrouterSession) PAUSERROLE() ([32]byte, error) {
	return _Delayredeemrouter.Contract.PAUSERROLE(&_Delayredeemrouter.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Delayredeemrouter.Contract.PAUSERROLE(&_Delayredeemrouter.CallOpts)
}

// REDEEMFEERATERANGE is a free data retrieval call binding the contract method 0x58f7e664.
//
// Solidity: function REDEEM_FEE_RATE_RANGE() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) REDEEMFEERATERANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "REDEEM_FEE_RATE_RANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REDEEMFEERATERANGE is a free data retrieval call binding the contract method 0x58f7e664.
//
// Solidity: function REDEEM_FEE_RATE_RANGE() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) REDEEMFEERATERANGE() (*big.Int, error) {
	return _Delayredeemrouter.Contract.REDEEMFEERATERANGE(&_Delayredeemrouter.CallOpts)
}

// REDEEMFEERATERANGE is a free data retrieval call binding the contract method 0x58f7e664.
//
// Solidity: function REDEEM_FEE_RATE_RANGE() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) REDEEMFEERATERANGE() (*big.Int, error) {
	return _Delayredeemrouter.Contract.REDEEMFEERATERANGE(&_Delayredeemrouter.CallOpts)
}

// SECONDSINADAY is a free data retrieval call binding the contract method 0xf9cfa06f.
//
// Solidity: function SECONDS_IN_A_DAY() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) SECONDSINADAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "SECONDS_IN_A_DAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSINADAY is a free data retrieval call binding the contract method 0xf9cfa06f.
//
// Solidity: function SECONDS_IN_A_DAY() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) SECONDSINADAY() (*big.Int, error) {
	return _Delayredeemrouter.Contract.SECONDSINADAY(&_Delayredeemrouter.CallOpts)
}

// SECONDSINADAY is a free data retrieval call binding the contract method 0xf9cfa06f.
//
// Solidity: function SECONDS_IN_A_DAY() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) SECONDSINADAY() (*big.Int, error) {
	return _Delayredeemrouter.Contract.SECONDSINADAY(&_Delayredeemrouter.CallOpts)
}

// CanClaimDelayedRedeem is a free data retrieval call binding the contract method 0x8b745eae.
//
// Solidity: function canClaimDelayedRedeem(address user, uint256 index) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCaller) CanClaimDelayedRedeem(opts *bind.CallOpts, user common.Address, index *big.Int) (bool, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "canClaimDelayedRedeem", user, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaimDelayedRedeem is a free data retrieval call binding the contract method 0x8b745eae.
//
// Solidity: function canClaimDelayedRedeem(address user, uint256 index) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterSession) CanClaimDelayedRedeem(user common.Address, index *big.Int) (bool, error) {
	return _Delayredeemrouter.Contract.CanClaimDelayedRedeem(&_Delayredeemrouter.CallOpts, user, index)
}

// CanClaimDelayedRedeem is a free data retrieval call binding the contract method 0x8b745eae.
//
// Solidity: function canClaimDelayedRedeem(address user, uint256 index) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) CanClaimDelayedRedeem(user common.Address, index *big.Int) (bool, error) {
	return _Delayredeemrouter.Contract.CanClaimDelayedRedeem(&_Delayredeemrouter.CallOpts, user, index)
}

// CancelFeeRate is a free data retrieval call binding the contract method 0x3c4da553.
//
// Solidity: function cancelFeeRate() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) CancelFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "cancelFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CancelFeeRate is a free data retrieval call binding the contract method 0x3c4da553.
//
// Solidity: function cancelFeeRate() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) CancelFeeRate() (*big.Int, error) {
	return _Delayredeemrouter.Contract.CancelFeeRate(&_Delayredeemrouter.CallOpts)
}

// CancelFeeRate is a free data retrieval call binding the contract method 0x3c4da553.
//
// Solidity: function cancelFeeRate() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) CancelFeeRate() (*big.Int, error) {
	return _Delayredeemrouter.Contract.CancelFeeRate(&_Delayredeemrouter.CallOpts)
}

// FastLanes is a free data retrieval call binding the contract method 0xa29efe37.
//
// Solidity: function fastLanes(address ) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCaller) FastLanes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "fastLanes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FastLanes is a free data retrieval call binding the contract method 0xa29efe37.
//
// Solidity: function fastLanes(address ) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterSession) FastLanes(arg0 common.Address) (bool, error) {
	return _Delayredeemrouter.Contract.FastLanes(&_Delayredeemrouter.CallOpts, arg0)
}

// FastLanes is a free data retrieval call binding the contract method 0xa29efe37.
//
// Solidity: function fastLanes(address ) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) FastLanes(arg0 common.Address) (bool, error) {
	return _Delayredeemrouter.Contract.FastLanes(&_Delayredeemrouter.CallOpts, arg0)
}

// GetClaimableUserDelayedRedeems is a free data retrieval call binding the contract method 0x4929b254.
//
// Solidity: function getClaimableUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_Delayredeemrouter *DelayredeemrouterCaller) GetClaimableUserDelayedRedeems(opts *bind.CallOpts, user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "getClaimableUserDelayedRedeems", user)

	if err != nil {
		return *new([]DelayRedeemRouterDelayedRedeem), err
	}

	out0 := *abi.ConvertType(out[0], new([]DelayRedeemRouterDelayedRedeem)).(*[]DelayRedeemRouterDelayedRedeem)

	return out0, err

}

// GetClaimableUserDelayedRedeems is a free data retrieval call binding the contract method 0x4929b254.
//
// Solidity: function getClaimableUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_Delayredeemrouter *DelayredeemrouterSession) GetClaimableUserDelayedRedeems(user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	return _Delayredeemrouter.Contract.GetClaimableUserDelayedRedeems(&_Delayredeemrouter.CallOpts, user)
}

// GetClaimableUserDelayedRedeems is a free data retrieval call binding the contract method 0x4929b254.
//
// Solidity: function getClaimableUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_Delayredeemrouter *DelayredeemrouterCallerSession) GetClaimableUserDelayedRedeems(user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	return _Delayredeemrouter.Contract.GetClaimableUserDelayedRedeems(&_Delayredeemrouter.CallOpts, user)
}

// GetCurrentCap is a free data retrieval call binding the contract method 0x2c612832.
//
// Solidity: function getCurrentCap(address token) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) GetCurrentCap(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "getCurrentCap", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentCap is a free data retrieval call binding the contract method 0x2c612832.
//
// Solidity: function getCurrentCap(address token) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) GetCurrentCap(token common.Address) (*big.Int, error) {
	return _Delayredeemrouter.Contract.GetCurrentCap(&_Delayredeemrouter.CallOpts, token)
}

// GetCurrentCap is a free data retrieval call binding the contract method 0x2c612832.
//
// Solidity: function getCurrentCap(address token) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) GetCurrentCap(token common.Address) (*big.Int, error) {
	return _Delayredeemrouter.Contract.GetCurrentCap(&_Delayredeemrouter.CallOpts, token)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Delayredeemrouter *DelayredeemrouterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Delayredeemrouter *DelayredeemrouterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Delayredeemrouter.Contract.GetRoleAdmin(&_Delayredeemrouter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Delayredeemrouter.Contract.GetRoleAdmin(&_Delayredeemrouter.CallOpts, role)
}

// GetUserDelayedRedeems is a free data retrieval call binding the contract method 0x99a49065.
//
// Solidity: function getUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_Delayredeemrouter *DelayredeemrouterCaller) GetUserDelayedRedeems(opts *bind.CallOpts, user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "getUserDelayedRedeems", user)

	if err != nil {
		return *new([]DelayRedeemRouterDelayedRedeem), err
	}

	out0 := *abi.ConvertType(out[0], new([]DelayRedeemRouterDelayedRedeem)).(*[]DelayRedeemRouterDelayedRedeem)

	return out0, err

}

// GetUserDelayedRedeems is a free data retrieval call binding the contract method 0x99a49065.
//
// Solidity: function getUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_Delayredeemrouter *DelayredeemrouterSession) GetUserDelayedRedeems(user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	return _Delayredeemrouter.Contract.GetUserDelayedRedeems(&_Delayredeemrouter.CallOpts, user)
}

// GetUserDelayedRedeems is a free data retrieval call binding the contract method 0x99a49065.
//
// Solidity: function getUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_Delayredeemrouter *DelayredeemrouterCallerSession) GetUserDelayedRedeems(user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	return _Delayredeemrouter.Contract.GetUserDelayedRedeems(&_Delayredeemrouter.CallOpts, user)
}

// GetUserRedeem is a free data retrieval call binding the contract method 0xcc6c5423.
//
// Solidity: function getUserRedeem(address recipient, uint256 maxNumberOfDelayedRedeemsToClaim) view returns((address,uint256,uint256)[])
func (_Delayredeemrouter *DelayredeemrouterCaller) GetUserRedeem(opts *bind.CallOpts, recipient common.Address, maxNumberOfDelayedRedeemsToClaim *big.Int) ([]DelayRedeemRouterDebtTokenAmount, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "getUserRedeem", recipient, maxNumberOfDelayedRedeemsToClaim)

	if err != nil {
		return *new([]DelayRedeemRouterDebtTokenAmount), err
	}

	out0 := *abi.ConvertType(out[0], new([]DelayRedeemRouterDebtTokenAmount)).(*[]DelayRedeemRouterDebtTokenAmount)

	return out0, err

}

// GetUserRedeem is a free data retrieval call binding the contract method 0xcc6c5423.
//
// Solidity: function getUserRedeem(address recipient, uint256 maxNumberOfDelayedRedeemsToClaim) view returns((address,uint256,uint256)[])
func (_Delayredeemrouter *DelayredeemrouterSession) GetUserRedeem(recipient common.Address, maxNumberOfDelayedRedeemsToClaim *big.Int) ([]DelayRedeemRouterDebtTokenAmount, error) {
	return _Delayredeemrouter.Contract.GetUserRedeem(&_Delayredeemrouter.CallOpts, recipient, maxNumberOfDelayedRedeemsToClaim)
}

// GetUserRedeem is a free data retrieval call binding the contract method 0xcc6c5423.
//
// Solidity: function getUserRedeem(address recipient, uint256 maxNumberOfDelayedRedeemsToClaim) view returns((address,uint256,uint256)[])
func (_Delayredeemrouter *DelayredeemrouterCallerSession) GetUserRedeem(recipient common.Address, maxNumberOfDelayedRedeemsToClaim *big.Int) ([]DelayRedeemRouterDebtTokenAmount, error) {
	return _Delayredeemrouter.Contract.GetUserRedeem(&_Delayredeemrouter.CallOpts, recipient, maxNumberOfDelayedRedeemsToClaim)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Delayredeemrouter.Contract.HasRole(&_Delayredeemrouter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Delayredeemrouter.Contract.HasRole(&_Delayredeemrouter.CallOpts, role, account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCaller) IsBlacklisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "isBlacklisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterSession) IsBlacklisted(account common.Address) (bool, error) {
	return _Delayredeemrouter.Contract.IsBlacklisted(&_Delayredeemrouter.CallOpts, account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) IsBlacklisted(account common.Address) (bool, error) {
	return _Delayredeemrouter.Contract.IsBlacklisted(&_Delayredeemrouter.CallOpts, account)
}

// IsBtclisted is a free data retrieval call binding the contract method 0x362aada3.
//
// Solidity: function isBtclisted(address token) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCaller) IsBtclisted(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "isBtclisted", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBtclisted is a free data retrieval call binding the contract method 0x362aada3.
//
// Solidity: function isBtclisted(address token) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterSession) IsBtclisted(token common.Address) (bool, error) {
	return _Delayredeemrouter.Contract.IsBtclisted(&_Delayredeemrouter.CallOpts, token)
}

// IsBtclisted is a free data retrieval call binding the contract method 0x362aada3.
//
// Solidity: function isBtclisted(address token) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) IsBtclisted(token common.Address) (bool, error) {
	return _Delayredeemrouter.Contract.IsBtclisted(&_Delayredeemrouter.CallOpts, token)
}

// IsPausedToken is a free data retrieval call binding the contract method 0x7de26f35.
//
// Solidity: function isPausedToken(address token) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCaller) IsPausedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "isPausedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPausedToken is a free data retrieval call binding the contract method 0x7de26f35.
//
// Solidity: function isPausedToken(address token) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterSession) IsPausedToken(token common.Address) (bool, error) {
	return _Delayredeemrouter.Contract.IsPausedToken(&_Delayredeemrouter.CallOpts, token)
}

// IsPausedToken is a free data retrieval call binding the contract method 0x7de26f35.
//
// Solidity: function isPausedToken(address token) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) IsPausedToken(token common.Address) (bool, error) {
	return _Delayredeemrouter.Contract.IsPausedToken(&_Delayredeemrouter.CallOpts, token)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "isWhitelisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Delayredeemrouter.Contract.IsWhitelisted(&_Delayredeemrouter.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Delayredeemrouter.Contract.IsWhitelisted(&_Delayredeemrouter.CallOpts, account)
}

// LastRebaseTimestamps is a free data retrieval call binding the contract method 0x7364a714.
//
// Solidity: function lastRebaseTimestamps(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) LastRebaseTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "lastRebaseTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRebaseTimestamps is a free data retrieval call binding the contract method 0x7364a714.
//
// Solidity: function lastRebaseTimestamps(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) LastRebaseTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Delayredeemrouter.Contract.LastRebaseTimestamps(&_Delayredeemrouter.CallOpts, arg0)
}

// LastRebaseTimestamps is a free data retrieval call binding the contract method 0x7364a714.
//
// Solidity: function lastRebaseTimestamps(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) LastRebaseTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Delayredeemrouter.Contract.LastRebaseTimestamps(&_Delayredeemrouter.CallOpts, arg0)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) ManagementFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "managementFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) ManagementFee() (*big.Int, error) {
	return _Delayredeemrouter.Contract.ManagementFee(&_Delayredeemrouter.CallOpts)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) ManagementFee() (*big.Int, error) {
	return _Delayredeemrouter.Contract.ManagementFee(&_Delayredeemrouter.CallOpts)
}

// MaxQuotas is a free data retrieval call binding the contract method 0x4fa581f9.
//
// Solidity: function maxQuotas(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) MaxQuotas(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "maxQuotas", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxQuotas is a free data retrieval call binding the contract method 0x4fa581f9.
//
// Solidity: function maxQuotas(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) MaxQuotas(arg0 common.Address) (*big.Int, error) {
	return _Delayredeemrouter.Contract.MaxQuotas(&_Delayredeemrouter.CallOpts, arg0)
}

// MaxQuotas is a free data retrieval call binding the contract method 0x4fa581f9.
//
// Solidity: function maxQuotas(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) MaxQuotas(arg0 common.Address) (*big.Int, error) {
	return _Delayredeemrouter.Contract.MaxQuotas(&_Delayredeemrouter.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterSession) Paused() (bool, error) {
	return _Delayredeemrouter.Contract.Paused(&_Delayredeemrouter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) Paused() (bool, error) {
	return _Delayredeemrouter.Contract.Paused(&_Delayredeemrouter.CallOpts)
}

// QuotaBases is a free data retrieval call binding the contract method 0x7bd8ba9d.
//
// Solidity: function quotaBases(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) QuotaBases(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "quotaBases", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuotaBases is a free data retrieval call binding the contract method 0x7bd8ba9d.
//
// Solidity: function quotaBases(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) QuotaBases(arg0 common.Address) (*big.Int, error) {
	return _Delayredeemrouter.Contract.QuotaBases(&_Delayredeemrouter.CallOpts, arg0)
}

// QuotaBases is a free data retrieval call binding the contract method 0x7bd8ba9d.
//
// Solidity: function quotaBases(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) QuotaBases(arg0 common.Address) (*big.Int, error) {
	return _Delayredeemrouter.Contract.QuotaBases(&_Delayredeemrouter.CallOpts, arg0)
}

// QuotaRates is a free data retrieval call binding the contract method 0xc0bcde5d.
//
// Solidity: function quotaRates(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) QuotaRates(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "quotaRates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuotaRates is a free data retrieval call binding the contract method 0xc0bcde5d.
//
// Solidity: function quotaRates(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) QuotaRates(arg0 common.Address) (*big.Int, error) {
	return _Delayredeemrouter.Contract.QuotaRates(&_Delayredeemrouter.CallOpts, arg0)
}

// QuotaRates is a free data retrieval call binding the contract method 0xc0bcde5d.
//
// Solidity: function quotaRates(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) QuotaRates(arg0 common.Address) (*big.Int, error) {
	return _Delayredeemrouter.Contract.QuotaRates(&_Delayredeemrouter.CallOpts, arg0)
}

// RedeemDelay is a free data retrieval call binding the contract method 0xd2adf402.
//
// Solidity: function redeemDelay() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) RedeemDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "redeemDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemDelay is a free data retrieval call binding the contract method 0xd2adf402.
//
// Solidity: function redeemDelay() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) RedeemDelay() (*big.Int, error) {
	return _Delayredeemrouter.Contract.RedeemDelay(&_Delayredeemrouter.CallOpts)
}

// RedeemDelay is a free data retrieval call binding the contract method 0xd2adf402.
//
// Solidity: function redeemDelay() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) RedeemDelay() (*big.Int, error) {
	return _Delayredeemrouter.Contract.RedeemDelay(&_Delayredeemrouter.CallOpts)
}

// RedeemFeeRate is a free data retrieval call binding the contract method 0x5872e6fa.
//
// Solidity: function redeemFeeRate() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) RedeemFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "redeemFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemFeeRate is a free data retrieval call binding the contract method 0x5872e6fa.
//
// Solidity: function redeemFeeRate() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) RedeemFeeRate() (*big.Int, error) {
	return _Delayredeemrouter.Contract.RedeemFeeRate(&_Delayredeemrouter.CallOpts)
}

// RedeemFeeRate is a free data retrieval call binding the contract method 0x5872e6fa.
//
// Solidity: function redeemFeeRate() view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) RedeemFeeRate() (*big.Int, error) {
	return _Delayredeemrouter.Contract.RedeemFeeRate(&_Delayredeemrouter.CallOpts)
}

// RedeemFees is a free data retrieval call binding the contract method 0x6b725ce9.
//
// Solidity: function redeemFees(address , uint256 ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) RedeemFees(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "redeemFees", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemFees is a free data retrieval call binding the contract method 0x6b725ce9.
//
// Solidity: function redeemFees(address , uint256 ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) RedeemFees(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Delayredeemrouter.Contract.RedeemFees(&_Delayredeemrouter.CallOpts, arg0, arg1)
}

// RedeemFees is a free data retrieval call binding the contract method 0x6b725ce9.
//
// Solidity: function redeemFees(address , uint256 ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) RedeemFees(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Delayredeemrouter.Contract.RedeemFees(&_Delayredeemrouter.CallOpts, arg0, arg1)
}

// RetainAmounts is a free data retrieval call binding the contract method 0xe5b01c35.
//
// Solidity: function retainAmounts(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) RetainAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "retainAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetainAmounts is a free data retrieval call binding the contract method 0xe5b01c35.
//
// Solidity: function retainAmounts(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) RetainAmounts(arg0 common.Address) (*big.Int, error) {
	return _Delayredeemrouter.Contract.RetainAmounts(&_Delayredeemrouter.CallOpts, arg0)
}

// RetainAmounts is a free data retrieval call binding the contract method 0xe5b01c35.
//
// Solidity: function retainAmounts(address ) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) RetainAmounts(arg0 common.Address) (*big.Int, error) {
	return _Delayredeemrouter.Contract.RetainAmounts(&_Delayredeemrouter.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Delayredeemrouter.Contract.SupportsInterface(&_Delayredeemrouter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Delayredeemrouter.Contract.SupportsInterface(&_Delayredeemrouter.CallOpts, interfaceId)
}

// TokenDebts is a free data retrieval call binding the contract method 0xf190439e.
//
// Solidity: function tokenDebts(address ) view returns(uint256 totalDebts, uint256 totalCleared)
func (_Delayredeemrouter *DelayredeemrouterCaller) TokenDebts(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalDebts   *big.Int
	TotalCleared *big.Int
}, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "tokenDebts", arg0)

	outstruct := new(struct {
		TotalDebts   *big.Int
		TotalCleared *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalDebts = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalCleared = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenDebts is a free data retrieval call binding the contract method 0xf190439e.
//
// Solidity: function tokenDebts(address ) view returns(uint256 totalDebts, uint256 totalCleared)
func (_Delayredeemrouter *DelayredeemrouterSession) TokenDebts(arg0 common.Address) (struct {
	TotalDebts   *big.Int
	TotalCleared *big.Int
}, error) {
	return _Delayredeemrouter.Contract.TokenDebts(&_Delayredeemrouter.CallOpts, arg0)
}

// TokenDebts is a free data retrieval call binding the contract method 0xf190439e.
//
// Solidity: function tokenDebts(address ) view returns(uint256 totalDebts, uint256 totalCleared)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) TokenDebts(arg0 common.Address) (struct {
	TotalDebts   *big.Int
	TotalCleared *big.Int
}, error) {
	return _Delayredeemrouter.Contract.TokenDebts(&_Delayredeemrouter.CallOpts, arg0)
}

// UniBTC is a free data retrieval call binding the contract method 0x59f3d39b.
//
// Solidity: function uniBTC() view returns(address)
func (_Delayredeemrouter *DelayredeemrouterCaller) UniBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "uniBTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniBTC is a free data retrieval call binding the contract method 0x59f3d39b.
//
// Solidity: function uniBTC() view returns(address)
func (_Delayredeemrouter *DelayredeemrouterSession) UniBTC() (common.Address, error) {
	return _Delayredeemrouter.Contract.UniBTC(&_Delayredeemrouter.CallOpts)
}

// UniBTC is a free data retrieval call binding the contract method 0x59f3d39b.
//
// Solidity: function uniBTC() view returns(address)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) UniBTC() (common.Address, error) {
	return _Delayredeemrouter.Contract.UniBTC(&_Delayredeemrouter.CallOpts)
}

// UserDelayedRedeemByIndex is a free data retrieval call binding the contract method 0xc1541631.
//
// Solidity: function userDelayedRedeemByIndex(address user, uint256 index) view returns((uint224,uint32,address))
func (_Delayredeemrouter *DelayredeemrouterCaller) UserDelayedRedeemByIndex(opts *bind.CallOpts, user common.Address, index *big.Int) (DelayRedeemRouterDelayedRedeem, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "userDelayedRedeemByIndex", user, index)

	if err != nil {
		return *new(DelayRedeemRouterDelayedRedeem), err
	}

	out0 := *abi.ConvertType(out[0], new(DelayRedeemRouterDelayedRedeem)).(*DelayRedeemRouterDelayedRedeem)

	return out0, err

}

// UserDelayedRedeemByIndex is a free data retrieval call binding the contract method 0xc1541631.
//
// Solidity: function userDelayedRedeemByIndex(address user, uint256 index) view returns((uint224,uint32,address))
func (_Delayredeemrouter *DelayredeemrouterSession) UserDelayedRedeemByIndex(user common.Address, index *big.Int) (DelayRedeemRouterDelayedRedeem, error) {
	return _Delayredeemrouter.Contract.UserDelayedRedeemByIndex(&_Delayredeemrouter.CallOpts, user, index)
}

// UserDelayedRedeemByIndex is a free data retrieval call binding the contract method 0xc1541631.
//
// Solidity: function userDelayedRedeemByIndex(address user, uint256 index) view returns((uint224,uint32,address))
func (_Delayredeemrouter *DelayredeemrouterCallerSession) UserDelayedRedeemByIndex(user common.Address, index *big.Int) (DelayRedeemRouterDelayedRedeem, error) {
	return _Delayredeemrouter.Contract.UserDelayedRedeemByIndex(&_Delayredeemrouter.CallOpts, user, index)
}

// UserRedeems is a free data retrieval call binding the contract method 0x8d18e24b.
//
// Solidity: function userRedeems(address user) view returns((uint256,(uint224,uint32,address)[]))
func (_Delayredeemrouter *DelayredeemrouterCaller) UserRedeems(opts *bind.CallOpts, user common.Address) (DelayRedeemRouterUserDelayedRedeems, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "userRedeems", user)

	if err != nil {
		return *new(DelayRedeemRouterUserDelayedRedeems), err
	}

	out0 := *abi.ConvertType(out[0], new(DelayRedeemRouterUserDelayedRedeems)).(*DelayRedeemRouterUserDelayedRedeems)

	return out0, err

}

// UserRedeems is a free data retrieval call binding the contract method 0x8d18e24b.
//
// Solidity: function userRedeems(address user) view returns((uint256,(uint224,uint32,address)[]))
func (_Delayredeemrouter *DelayredeemrouterSession) UserRedeems(user common.Address) (DelayRedeemRouterUserDelayedRedeems, error) {
	return _Delayredeemrouter.Contract.UserRedeems(&_Delayredeemrouter.CallOpts, user)
}

// UserRedeems is a free data retrieval call binding the contract method 0x8d18e24b.
//
// Solidity: function userRedeems(address user) view returns((uint256,(uint224,uint32,address)[]))
func (_Delayredeemrouter *DelayredeemrouterCallerSession) UserRedeems(user common.Address) (DelayRedeemRouterUserDelayedRedeems, error) {
	return _Delayredeemrouter.Contract.UserRedeems(&_Delayredeemrouter.CallOpts, user)
}

// UserRedeemsLength is a free data retrieval call binding the contract method 0x5449c33c.
//
// Solidity: function userRedeemsLength(address user) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCaller) UserRedeemsLength(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "userRedeemsLength", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRedeemsLength is a free data retrieval call binding the contract method 0x5449c33c.
//
// Solidity: function userRedeemsLength(address user) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterSession) UserRedeemsLength(user common.Address) (*big.Int, error) {
	return _Delayredeemrouter.Contract.UserRedeemsLength(&_Delayredeemrouter.CallOpts, user)
}

// UserRedeemsLength is a free data retrieval call binding the contract method 0x5449c33c.
//
// Solidity: function userRedeemsLength(address user) view returns(uint256)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) UserRedeemsLength(user common.Address) (*big.Int, error) {
	return _Delayredeemrouter.Contract.UserRedeemsLength(&_Delayredeemrouter.CallOpts, user)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Delayredeemrouter *DelayredeemrouterCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Delayredeemrouter *DelayredeemrouterSession) Vault() (common.Address, error) {
	return _Delayredeemrouter.Contract.Vault(&_Delayredeemrouter.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) Vault() (common.Address, error) {
	return _Delayredeemrouter.Contract.Vault(&_Delayredeemrouter.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCaller) WhitelistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Delayredeemrouter.contract.Call(opts, &out, "whitelistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterSession) WhitelistEnabled() (bool, error) {
	return _Delayredeemrouter.Contract.WhitelistEnabled(&_Delayredeemrouter.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Delayredeemrouter *DelayredeemrouterCallerSession) WhitelistEnabled() (bool, error) {
	return _Delayredeemrouter.Contract.WhitelistEnabled(&_Delayredeemrouter.CallOpts)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x935eb35f.
//
// Solidity: function addToBlacklist(address[] _accounts) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) AddToBlacklist(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "addToBlacklist", _accounts)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x935eb35f.
//
// Solidity: function addToBlacklist(address[] _accounts) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) AddToBlacklist(_accounts []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.AddToBlacklist(&_Delayredeemrouter.TransactOpts, _accounts)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x935eb35f.
//
// Solidity: function addToBlacklist(address[] _accounts) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) AddToBlacklist(_accounts []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.AddToBlacklist(&_Delayredeemrouter.TransactOpts, _accounts)
}

// AddToBtclist is a paid mutator transaction binding the contract method 0x0ba15afb.
//
// Solidity: function addToBtclist(address[] _tokens) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) AddToBtclist(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "addToBtclist", _tokens)
}

// AddToBtclist is a paid mutator transaction binding the contract method 0x0ba15afb.
//
// Solidity: function addToBtclist(address[] _tokens) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) AddToBtclist(_tokens []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.AddToBtclist(&_Delayredeemrouter.TransactOpts, _tokens)
}

// AddToBtclist is a paid mutator transaction binding the contract method 0x0ba15afb.
//
// Solidity: function addToBtclist(address[] _tokens) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) AddToBtclist(_tokens []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.AddToBtclist(&_Delayredeemrouter.TransactOpts, _tokens)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _accounts) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) AddToWhitelist(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "addToWhitelist", _accounts)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _accounts) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) AddToWhitelist(_accounts []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.AddToWhitelist(&_Delayredeemrouter.TransactOpts, _accounts)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _accounts) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) AddToWhitelist(_accounts []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.AddToWhitelist(&_Delayredeemrouter.TransactOpts, _accounts)
}

// ClaimDelayedRedeems is a paid mutator transaction binding the contract method 0xf2881130.
//
// Solidity: function claimDelayedRedeems() returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) ClaimDelayedRedeems(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "claimDelayedRedeems")
}

// ClaimDelayedRedeems is a paid mutator transaction binding the contract method 0xf2881130.
//
// Solidity: function claimDelayedRedeems() returns()
func (_Delayredeemrouter *DelayredeemrouterSession) ClaimDelayedRedeems() (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.ClaimDelayedRedeems(&_Delayredeemrouter.TransactOpts)
}

// ClaimDelayedRedeems is a paid mutator transaction binding the contract method 0xf2881130.
//
// Solidity: function claimDelayedRedeems() returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) ClaimDelayedRedeems() (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.ClaimDelayedRedeems(&_Delayredeemrouter.TransactOpts)
}

// ClaimDelayedRedeems0 is a paid mutator transaction binding the contract method 0xf33fca17.
//
// Solidity: function claimDelayedRedeems(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) ClaimDelayedRedeems0(opts *bind.TransactOpts, maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "claimDelayedRedeems0", maxNumberOfDelayedRedeemsToClaim)
}

// ClaimDelayedRedeems0 is a paid mutator transaction binding the contract method 0xf33fca17.
//
// Solidity: function claimDelayedRedeems(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) ClaimDelayedRedeems0(maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.ClaimDelayedRedeems0(&_Delayredeemrouter.TransactOpts, maxNumberOfDelayedRedeemsToClaim)
}

// ClaimDelayedRedeems0 is a paid mutator transaction binding the contract method 0xf33fca17.
//
// Solidity: function claimDelayedRedeems(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) ClaimDelayedRedeems0(maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.ClaimDelayedRedeems0(&_Delayredeemrouter.TransactOpts, maxNumberOfDelayedRedeemsToClaim)
}

// CreateDelayedRedeem is a paid mutator transaction binding the contract method 0xddc1bdea.
//
// Solidity: function createDelayedRedeem(address token, uint256 amount) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) CreateDelayedRedeem(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "createDelayedRedeem", token, amount)
}

// CreateDelayedRedeem is a paid mutator transaction binding the contract method 0xddc1bdea.
//
// Solidity: function createDelayedRedeem(address token, uint256 amount) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) CreateDelayedRedeem(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.CreateDelayedRedeem(&_Delayredeemrouter.TransactOpts, token, amount)
}

// CreateDelayedRedeem is a paid mutator transaction binding the contract method 0xddc1bdea.
//
// Solidity: function createDelayedRedeem(address token, uint256 amount) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) CreateDelayedRedeem(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.CreateDelayedRedeem(&_Delayredeemrouter.TransactOpts, token, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.GrantRole(&_Delayredeemrouter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.GrantRole(&_Delayredeemrouter.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xeea39f79.
//
// Solidity: function initialize(address _defaultAdmin, address _uniBTC, address _vault, uint256 _redeemDelay, bool _whitelistEnabled) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) Initialize(opts *bind.TransactOpts, _defaultAdmin common.Address, _uniBTC common.Address, _vault common.Address, _redeemDelay *big.Int, _whitelistEnabled bool) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "initialize", _defaultAdmin, _uniBTC, _vault, _redeemDelay, _whitelistEnabled)
}

// Initialize is a paid mutator transaction binding the contract method 0xeea39f79.
//
// Solidity: function initialize(address _defaultAdmin, address _uniBTC, address _vault, uint256 _redeemDelay, bool _whitelistEnabled) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) Initialize(_defaultAdmin common.Address, _uniBTC common.Address, _vault common.Address, _redeemDelay *big.Int, _whitelistEnabled bool) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.Initialize(&_Delayredeemrouter.TransactOpts, _defaultAdmin, _uniBTC, _vault, _redeemDelay, _whitelistEnabled)
}

// Initialize is a paid mutator transaction binding the contract method 0xeea39f79.
//
// Solidity: function initialize(address _defaultAdmin, address _uniBTC, address _vault, uint256 _redeemDelay, bool _whitelistEnabled) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) Initialize(_defaultAdmin common.Address, _uniBTC common.Address, _vault common.Address, _redeemDelay *big.Int, _whitelistEnabled bool) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.Initialize(&_Delayredeemrouter.TransactOpts, _defaultAdmin, _uniBTC, _vault, _redeemDelay, _whitelistEnabled)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Delayredeemrouter *DelayredeemrouterSession) Pause() (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.Pause(&_Delayredeemrouter.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) Pause() (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.Pause(&_Delayredeemrouter.TransactOpts)
}

// PauseTokens is a paid mutator transaction binding the contract method 0xc609684a.
//
// Solidity: function pauseTokens(address[] _tokens) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) PauseTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "pauseTokens", _tokens)
}

// PauseTokens is a paid mutator transaction binding the contract method 0xc609684a.
//
// Solidity: function pauseTokens(address[] _tokens) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) PauseTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.PauseTokens(&_Delayredeemrouter.TransactOpts, _tokens)
}

// PauseTokens is a paid mutator transaction binding the contract method 0xc609684a.
//
// Solidity: function pauseTokens(address[] _tokens) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) PauseTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.PauseTokens(&_Delayredeemrouter.TransactOpts, _tokens)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x89daf799.
//
// Solidity: function removeFromBlacklist(address[] _accounts) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) RemoveFromBlacklist(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "removeFromBlacklist", _accounts)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x89daf799.
//
// Solidity: function removeFromBlacklist(address[] _accounts) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) RemoveFromBlacklist(_accounts []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.RemoveFromBlacklist(&_Delayredeemrouter.TransactOpts, _accounts)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x89daf799.
//
// Solidity: function removeFromBlacklist(address[] _accounts) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) RemoveFromBlacklist(_accounts []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.RemoveFromBlacklist(&_Delayredeemrouter.TransactOpts, _accounts)
}

// RemoveFromBtclist is a paid mutator transaction binding the contract method 0xdd954363.
//
// Solidity: function removeFromBtclist(address[] _tokens) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) RemoveFromBtclist(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "removeFromBtclist", _tokens)
}

// RemoveFromBtclist is a paid mutator transaction binding the contract method 0xdd954363.
//
// Solidity: function removeFromBtclist(address[] _tokens) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) RemoveFromBtclist(_tokens []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.RemoveFromBtclist(&_Delayredeemrouter.TransactOpts, _tokens)
}

// RemoveFromBtclist is a paid mutator transaction binding the contract method 0xdd954363.
//
// Solidity: function removeFromBtclist(address[] _tokens) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) RemoveFromBtclist(_tokens []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.RemoveFromBtclist(&_Delayredeemrouter.TransactOpts, _tokens)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _accounts) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "removeFromWhitelist", _accounts)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _accounts) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) RemoveFromWhitelist(_accounts []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.RemoveFromWhitelist(&_Delayredeemrouter.TransactOpts, _accounts)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _accounts) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) RemoveFromWhitelist(_accounts []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.RemoveFromWhitelist(&_Delayredeemrouter.TransactOpts, _accounts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.RenounceRole(&_Delayredeemrouter.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.RenounceRole(&_Delayredeemrouter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.RevokeRole(&_Delayredeemrouter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.RevokeRole(&_Delayredeemrouter.TransactOpts, role, account)
}

// SetCancelFeeRate is a paid mutator transaction binding the contract method 0xfb00d01f.
//
// Solidity: function setCancelFeeRate(uint256 _newFeeRate) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) SetCancelFeeRate(opts *bind.TransactOpts, _newFeeRate *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "setCancelFeeRate", _newFeeRate)
}

// SetCancelFeeRate is a paid mutator transaction binding the contract method 0xfb00d01f.
//
// Solidity: function setCancelFeeRate(uint256 _newFeeRate) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) SetCancelFeeRate(_newFeeRate *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetCancelFeeRate(&_Delayredeemrouter.TransactOpts, _newFeeRate)
}

// SetCancelFeeRate is a paid mutator transaction binding the contract method 0xfb00d01f.
//
// Solidity: function setCancelFeeRate(uint256 _newFeeRate) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) SetCancelFeeRate(_newFeeRate *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetCancelFeeRate(&_Delayredeemrouter.TransactOpts, _newFeeRate)
}

// SetFastLane is a paid mutator transaction binding the contract method 0xc376f608.
//
// Solidity: function setFastLane(address[] _accounts, bool[] _fastLaneStatus) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) SetFastLane(opts *bind.TransactOpts, _accounts []common.Address, _fastLaneStatus []bool) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "setFastLane", _accounts, _fastLaneStatus)
}

// SetFastLane is a paid mutator transaction binding the contract method 0xc376f608.
//
// Solidity: function setFastLane(address[] _accounts, bool[] _fastLaneStatus) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) SetFastLane(_accounts []common.Address, _fastLaneStatus []bool) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetFastLane(&_Delayredeemrouter.TransactOpts, _accounts, _fastLaneStatus)
}

// SetFastLane is a paid mutator transaction binding the contract method 0xc376f608.
//
// Solidity: function setFastLane(address[] _accounts, bool[] _fastLaneStatus) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) SetFastLane(_accounts []common.Address, _fastLaneStatus []bool) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetFastLane(&_Delayredeemrouter.TransactOpts, _accounts, _fastLaneStatus)
}

// SetMaxQuotaForTokens is a paid mutator transaction binding the contract method 0x3a413394.
//
// Solidity: function setMaxQuotaForTokens(address[] _tokens, uint256[] _quotas) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) SetMaxQuotaForTokens(opts *bind.TransactOpts, _tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "setMaxQuotaForTokens", _tokens, _quotas)
}

// SetMaxQuotaForTokens is a paid mutator transaction binding the contract method 0x3a413394.
//
// Solidity: function setMaxQuotaForTokens(address[] _tokens, uint256[] _quotas) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) SetMaxQuotaForTokens(_tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetMaxQuotaForTokens(&_Delayredeemrouter.TransactOpts, _tokens, _quotas)
}

// SetMaxQuotaForTokens is a paid mutator transaction binding the contract method 0x3a413394.
//
// Solidity: function setMaxQuotaForTokens(address[] _tokens, uint256[] _quotas) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) SetMaxQuotaForTokens(_tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetMaxQuotaForTokens(&_Delayredeemrouter.TransactOpts, _tokens, _quotas)
}

// SetQuotaRates is a paid mutator transaction binding the contract method 0x377ef35a.
//
// Solidity: function setQuotaRates(address[] _tokens, uint256[] _quotas) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) SetQuotaRates(opts *bind.TransactOpts, _tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "setQuotaRates", _tokens, _quotas)
}

// SetQuotaRates is a paid mutator transaction binding the contract method 0x377ef35a.
//
// Solidity: function setQuotaRates(address[] _tokens, uint256[] _quotas) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) SetQuotaRates(_tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetQuotaRates(&_Delayredeemrouter.TransactOpts, _tokens, _quotas)
}

// SetQuotaRates is a paid mutator transaction binding the contract method 0x377ef35a.
//
// Solidity: function setQuotaRates(address[] _tokens, uint256[] _quotas) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) SetQuotaRates(_tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetQuotaRates(&_Delayredeemrouter.TransactOpts, _tokens, _quotas)
}

// SetRedeemDelay is a paid mutator transaction binding the contract method 0x668282a0.
//
// Solidity: function setRedeemDelay(uint256 _newDelay) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) SetRedeemDelay(opts *bind.TransactOpts, _newDelay *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "setRedeemDelay", _newDelay)
}

// SetRedeemDelay is a paid mutator transaction binding the contract method 0x668282a0.
//
// Solidity: function setRedeemDelay(uint256 _newDelay) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) SetRedeemDelay(_newDelay *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetRedeemDelay(&_Delayredeemrouter.TransactOpts, _newDelay)
}

// SetRedeemDelay is a paid mutator transaction binding the contract method 0x668282a0.
//
// Solidity: function setRedeemDelay(uint256 _newDelay) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) SetRedeemDelay(_newDelay *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetRedeemDelay(&_Delayredeemrouter.TransactOpts, _newDelay)
}

// SetRedeemFeeRate is a paid mutator transaction binding the contract method 0x21e822c5.
//
// Solidity: function setRedeemFeeRate(uint256 _newFeeRate) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) SetRedeemFeeRate(opts *bind.TransactOpts, _newFeeRate *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "setRedeemFeeRate", _newFeeRate)
}

// SetRedeemFeeRate is a paid mutator transaction binding the contract method 0x21e822c5.
//
// Solidity: function setRedeemFeeRate(uint256 _newFeeRate) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) SetRedeemFeeRate(_newFeeRate *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetRedeemFeeRate(&_Delayredeemrouter.TransactOpts, _newFeeRate)
}

// SetRedeemFeeRate is a paid mutator transaction binding the contract method 0x21e822c5.
//
// Solidity: function setRedeemFeeRate(uint256 _newFeeRate) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) SetRedeemFeeRate(_newFeeRate *big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetRedeemFeeRate(&_Delayredeemrouter.TransactOpts, _newFeeRate)
}

// SetRetainAmounts is a paid mutator transaction binding the contract method 0x34101858.
//
// Solidity: function setRetainAmounts(address[] _tokens, uint256[] _balances) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) SetRetainAmounts(opts *bind.TransactOpts, _tokens []common.Address, _balances []*big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "setRetainAmounts", _tokens, _balances)
}

// SetRetainAmounts is a paid mutator transaction binding the contract method 0x34101858.
//
// Solidity: function setRetainAmounts(address[] _tokens, uint256[] _balances) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) SetRetainAmounts(_tokens []common.Address, _balances []*big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetRetainAmounts(&_Delayredeemrouter.TransactOpts, _tokens, _balances)
}

// SetRetainAmounts is a paid mutator transaction binding the contract method 0x34101858.
//
// Solidity: function setRetainAmounts(address[] _tokens, uint256[] _balances) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) SetRetainAmounts(_tokens []common.Address, _balances []*big.Int) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetRetainAmounts(&_Delayredeemrouter.TransactOpts, _tokens, _balances)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _enabled) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) SetWhitelistEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "setWhitelistEnabled", _enabled)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _enabled) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) SetWhitelistEnabled(_enabled bool) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetWhitelistEnabled(&_Delayredeemrouter.TransactOpts, _enabled)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _enabled) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) SetWhitelistEnabled(_enabled bool) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.SetWhitelistEnabled(&_Delayredeemrouter.TransactOpts, _enabled)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Delayredeemrouter *DelayredeemrouterSession) Unpause() (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.Unpause(&_Delayredeemrouter.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) Unpause() (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.Unpause(&_Delayredeemrouter.TransactOpts)
}

// UnpauseTokens is a paid mutator transaction binding the contract method 0x7a4df6eb.
//
// Solidity: function unpauseTokens(address[] _tokens) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) UnpauseTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "unpauseTokens", _tokens)
}

// UnpauseTokens is a paid mutator transaction binding the contract method 0x7a4df6eb.
//
// Solidity: function unpauseTokens(address[] _tokens) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) UnpauseTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.UnpauseTokens(&_Delayredeemrouter.TransactOpts, _tokens)
}

// UnpauseTokens is a paid mutator transaction binding the contract method 0x7a4df6eb.
//
// Solidity: function unpauseTokens(address[] _tokens) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) UnpauseTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.UnpauseTokens(&_Delayredeemrouter.TransactOpts, _tokens)
}

// WithdrawManagementFee is a paid mutator transaction binding the contract method 0xb175968e.
//
// Solidity: function withdrawManagementFee(uint256 _amount, address _recipient) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) WithdrawManagementFee(opts *bind.TransactOpts, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.Transact(opts, "withdrawManagementFee", _amount, _recipient)
}

// WithdrawManagementFee is a paid mutator transaction binding the contract method 0xb175968e.
//
// Solidity: function withdrawManagementFee(uint256 _amount, address _recipient) returns()
func (_Delayredeemrouter *DelayredeemrouterSession) WithdrawManagementFee(_amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.WithdrawManagementFee(&_Delayredeemrouter.TransactOpts, _amount, _recipient)
}

// WithdrawManagementFee is a paid mutator transaction binding the contract method 0xb175968e.
//
// Solidity: function withdrawManagementFee(uint256 _amount, address _recipient) returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) WithdrawManagementFee(_amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.WithdrawManagementFee(&_Delayredeemrouter.TransactOpts, _amount, _recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Delayredeemrouter *DelayredeemrouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delayredeemrouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Delayredeemrouter *DelayredeemrouterSession) Receive() (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.Receive(&_Delayredeemrouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Delayredeemrouter *DelayredeemrouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Delayredeemrouter.Contract.Receive(&_Delayredeemrouter.TransactOpts)
}

// DelayredeemrouterBlacklistAddedIterator is returned from FilterBlacklistAdded and is used to iterate over the raw logs and unpacked data for BlacklistAdded events raised by the Delayredeemrouter contract.
type DelayredeemrouterBlacklistAddedIterator struct {
	Event *DelayredeemrouterBlacklistAdded // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterBlacklistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterBlacklistAdded)
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
		it.Event = new(DelayredeemrouterBlacklistAdded)
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
func (it *DelayredeemrouterBlacklistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterBlacklistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterBlacklistAdded represents a BlacklistAdded event raised by the Delayredeemrouter contract.
type DelayredeemrouterBlacklistAdded struct {
	Accounts []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlacklistAdded is a free log retrieval operation binding the contract event 0x065786e2f100ecf1a39fd27fb1e6042658f97ff4d9093657ae121f534e46c038.
//
// Solidity: event BlacklistAdded(address[] accounts)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterBlacklistAdded(opts *bind.FilterOpts) (*DelayredeemrouterBlacklistAddedIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "BlacklistAdded")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterBlacklistAddedIterator{contract: _Delayredeemrouter.contract, event: "BlacklistAdded", logs: logs, sub: sub}, nil
}

// WatchBlacklistAdded is a free log subscription operation binding the contract event 0x065786e2f100ecf1a39fd27fb1e6042658f97ff4d9093657ae121f534e46c038.
//
// Solidity: event BlacklistAdded(address[] accounts)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchBlacklistAdded(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterBlacklistAdded) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "BlacklistAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterBlacklistAdded)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "BlacklistAdded", log); err != nil {
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

// ParseBlacklistAdded is a log parse operation binding the contract event 0x065786e2f100ecf1a39fd27fb1e6042658f97ff4d9093657ae121f534e46c038.
//
// Solidity: event BlacklistAdded(address[] accounts)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseBlacklistAdded(log types.Log) (*DelayredeemrouterBlacklistAdded, error) {
	event := new(DelayredeemrouterBlacklistAdded)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "BlacklistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterBlacklistRemovedIterator is returned from FilterBlacklistRemoved and is used to iterate over the raw logs and unpacked data for BlacklistRemoved events raised by the Delayredeemrouter contract.
type DelayredeemrouterBlacklistRemovedIterator struct {
	Event *DelayredeemrouterBlacklistRemoved // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterBlacklistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterBlacklistRemoved)
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
		it.Event = new(DelayredeemrouterBlacklistRemoved)
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
func (it *DelayredeemrouterBlacklistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterBlacklistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterBlacklistRemoved represents a BlacklistRemoved event raised by the Delayredeemrouter contract.
type DelayredeemrouterBlacklistRemoved struct {
	Accounts []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlacklistRemoved is a free log retrieval operation binding the contract event 0xb1a383a26b5d809f3cf5b9b022fbfe3e4896e6f1ff310ce38c785d5638bc31bf.
//
// Solidity: event BlacklistRemoved(address[] accounts)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterBlacklistRemoved(opts *bind.FilterOpts) (*DelayredeemrouterBlacklistRemovedIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "BlacklistRemoved")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterBlacklistRemovedIterator{contract: _Delayredeemrouter.contract, event: "BlacklistRemoved", logs: logs, sub: sub}, nil
}

// WatchBlacklistRemoved is a free log subscription operation binding the contract event 0xb1a383a26b5d809f3cf5b9b022fbfe3e4896e6f1ff310ce38c785d5638bc31bf.
//
// Solidity: event BlacklistRemoved(address[] accounts)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchBlacklistRemoved(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterBlacklistRemoved) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "BlacklistRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterBlacklistRemoved)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "BlacklistRemoved", log); err != nil {
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

// ParseBlacklistRemoved is a log parse operation binding the contract event 0xb1a383a26b5d809f3cf5b9b022fbfe3e4896e6f1ff310ce38c785d5638bc31bf.
//
// Solidity: event BlacklistRemoved(address[] accounts)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseBlacklistRemoved(log types.Log) (*DelayredeemrouterBlacklistRemoved, error) {
	event := new(DelayredeemrouterBlacklistRemoved)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "BlacklistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterBtclistAddedIterator is returned from FilterBtclistAdded and is used to iterate over the raw logs and unpacked data for BtclistAdded events raised by the Delayredeemrouter contract.
type DelayredeemrouterBtclistAddedIterator struct {
	Event *DelayredeemrouterBtclistAdded // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterBtclistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterBtclistAdded)
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
		it.Event = new(DelayredeemrouterBtclistAdded)
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
func (it *DelayredeemrouterBtclistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterBtclistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterBtclistAdded represents a BtclistAdded event raised by the Delayredeemrouter contract.
type DelayredeemrouterBtclistAdded struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBtclistAdded is a free log retrieval operation binding the contract event 0xa8f863280802460b807e3f117dda44a95e056aee58b5ac4ac75c254003130e73.
//
// Solidity: event BtclistAdded(address[] tokens)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterBtclistAdded(opts *bind.FilterOpts) (*DelayredeemrouterBtclistAddedIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "BtclistAdded")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterBtclistAddedIterator{contract: _Delayredeemrouter.contract, event: "BtclistAdded", logs: logs, sub: sub}, nil
}

// WatchBtclistAdded is a free log subscription operation binding the contract event 0xa8f863280802460b807e3f117dda44a95e056aee58b5ac4ac75c254003130e73.
//
// Solidity: event BtclistAdded(address[] tokens)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchBtclistAdded(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterBtclistAdded) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "BtclistAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterBtclistAdded)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "BtclistAdded", log); err != nil {
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

// ParseBtclistAdded is a log parse operation binding the contract event 0xa8f863280802460b807e3f117dda44a95e056aee58b5ac4ac75c254003130e73.
//
// Solidity: event BtclistAdded(address[] tokens)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseBtclistAdded(log types.Log) (*DelayredeemrouterBtclistAdded, error) {
	event := new(DelayredeemrouterBtclistAdded)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "BtclistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterBtclistRemovedIterator is returned from FilterBtclistRemoved and is used to iterate over the raw logs and unpacked data for BtclistRemoved events raised by the Delayredeemrouter contract.
type DelayredeemrouterBtclistRemovedIterator struct {
	Event *DelayredeemrouterBtclistRemoved // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterBtclistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterBtclistRemoved)
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
		it.Event = new(DelayredeemrouterBtclistRemoved)
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
func (it *DelayredeemrouterBtclistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterBtclistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterBtclistRemoved represents a BtclistRemoved event raised by the Delayredeemrouter contract.
type DelayredeemrouterBtclistRemoved struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBtclistRemoved is a free log retrieval operation binding the contract event 0x90c77a347cd0b5c7f69208f770965d69e3af347c0e5ef344bfe02c2edcd3ccdc.
//
// Solidity: event BtclistRemoved(address[] tokens)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterBtclistRemoved(opts *bind.FilterOpts) (*DelayredeemrouterBtclistRemovedIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "BtclistRemoved")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterBtclistRemovedIterator{contract: _Delayredeemrouter.contract, event: "BtclistRemoved", logs: logs, sub: sub}, nil
}

// WatchBtclistRemoved is a free log subscription operation binding the contract event 0x90c77a347cd0b5c7f69208f770965d69e3af347c0e5ef344bfe02c2edcd3ccdc.
//
// Solidity: event BtclistRemoved(address[] tokens)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchBtclistRemoved(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterBtclistRemoved) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "BtclistRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterBtclistRemoved)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "BtclistRemoved", log); err != nil {
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

// ParseBtclistRemoved is a log parse operation binding the contract event 0x90c77a347cd0b5c7f69208f770965d69e3af347c0e5ef344bfe02c2edcd3ccdc.
//
// Solidity: event BtclistRemoved(address[] tokens)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseBtclistRemoved(log types.Log) (*DelayredeemrouterBtclistRemoved, error) {
	event := new(DelayredeemrouterBtclistRemoved)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "BtclistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterCancelFeeRateSetIterator is returned from FilterCancelFeeRateSet and is used to iterate over the raw logs and unpacked data for CancelFeeRateSet events raised by the Delayredeemrouter contract.
type DelayredeemrouterCancelFeeRateSetIterator struct {
	Event *DelayredeemrouterCancelFeeRateSet // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterCancelFeeRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterCancelFeeRateSet)
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
		it.Event = new(DelayredeemrouterCancelFeeRateSet)
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
func (it *DelayredeemrouterCancelFeeRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterCancelFeeRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterCancelFeeRateSet represents a CancelFeeRateSet event raised by the Delayredeemrouter contract.
type DelayredeemrouterCancelFeeRateSet struct {
	PreviousFeeRate *big.Int
	NewFeeRate      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCancelFeeRateSet is a free log retrieval operation binding the contract event 0x8bea0dd2ef3ce75d3862a08fc5c63017ae74acd1adf71aaa5887559ae088b528.
//
// Solidity: event CancelFeeRateSet(uint256 previousFeeRate, uint256 newFeeRate)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterCancelFeeRateSet(opts *bind.FilterOpts) (*DelayredeemrouterCancelFeeRateSetIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "CancelFeeRateSet")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterCancelFeeRateSetIterator{contract: _Delayredeemrouter.contract, event: "CancelFeeRateSet", logs: logs, sub: sub}, nil
}

// WatchCancelFeeRateSet is a free log subscription operation binding the contract event 0x8bea0dd2ef3ce75d3862a08fc5c63017ae74acd1adf71aaa5887559ae088b528.
//
// Solidity: event CancelFeeRateSet(uint256 previousFeeRate, uint256 newFeeRate)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchCancelFeeRateSet(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterCancelFeeRateSet) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "CancelFeeRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterCancelFeeRateSet)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "CancelFeeRateSet", log); err != nil {
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

// ParseCancelFeeRateSet is a log parse operation binding the contract event 0x8bea0dd2ef3ce75d3862a08fc5c63017ae74acd1adf71aaa5887559ae088b528.
//
// Solidity: event CancelFeeRateSet(uint256 previousFeeRate, uint256 newFeeRate)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseCancelFeeRateSet(log types.Log) (*DelayredeemrouterCancelFeeRateSet, error) {
	event := new(DelayredeemrouterCancelFeeRateSet)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "CancelFeeRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterDelayedRedeemCreatedIterator is returned from FilterDelayedRedeemCreated and is used to iterate over the raw logs and unpacked data for DelayedRedeemCreated events raised by the Delayredeemrouter contract.
type DelayredeemrouterDelayedRedeemCreatedIterator struct {
	Event *DelayredeemrouterDelayedRedeemCreated // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterDelayedRedeemCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterDelayedRedeemCreated)
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
		it.Event = new(DelayredeemrouterDelayedRedeemCreated)
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
func (it *DelayredeemrouterDelayedRedeemCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterDelayedRedeemCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterDelayedRedeemCreated represents a DelayedRedeemCreated event raised by the Delayredeemrouter contract.
type DelayredeemrouterDelayedRedeemCreated struct {
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	Index     *big.Int
	RedeemFee *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayedRedeemCreated is a free log retrieval operation binding the contract event 0xe2ede624e6a605222e831f4f91f2930b35c4d10323da5b68923c3086d4f0b3c0.
//
// Solidity: event DelayedRedeemCreated(address recipient, address token, uint256 amount, uint256 index, uint256 redeemFee)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterDelayedRedeemCreated(opts *bind.FilterOpts) (*DelayredeemrouterDelayedRedeemCreatedIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "DelayedRedeemCreated")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterDelayedRedeemCreatedIterator{contract: _Delayredeemrouter.contract, event: "DelayedRedeemCreated", logs: logs, sub: sub}, nil
}

// WatchDelayedRedeemCreated is a free log subscription operation binding the contract event 0xe2ede624e6a605222e831f4f91f2930b35c4d10323da5b68923c3086d4f0b3c0.
//
// Solidity: event DelayedRedeemCreated(address recipient, address token, uint256 amount, uint256 index, uint256 redeemFee)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchDelayedRedeemCreated(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterDelayedRedeemCreated) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "DelayedRedeemCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterDelayedRedeemCreated)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "DelayedRedeemCreated", log); err != nil {
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

// ParseDelayedRedeemCreated is a log parse operation binding the contract event 0xe2ede624e6a605222e831f4f91f2930b35c4d10323da5b68923c3086d4f0b3c0.
//
// Solidity: event DelayedRedeemCreated(address recipient, address token, uint256 amount, uint256 index, uint256 redeemFee)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseDelayedRedeemCreated(log types.Log) (*DelayredeemrouterDelayedRedeemCreated, error) {
	event := new(DelayredeemrouterDelayedRedeemCreated)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "DelayedRedeemCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterDelayedRedeemsClaimedIterator is returned from FilterDelayedRedeemsClaimed and is used to iterate over the raw logs and unpacked data for DelayedRedeemsClaimed events raised by the Delayredeemrouter contract.
type DelayredeemrouterDelayedRedeemsClaimedIterator struct {
	Event *DelayredeemrouterDelayedRedeemsClaimed // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterDelayedRedeemsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterDelayedRedeemsClaimed)
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
		it.Event = new(DelayredeemrouterDelayedRedeemsClaimed)
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
func (it *DelayredeemrouterDelayedRedeemsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterDelayedRedeemsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterDelayedRedeemsClaimed represents a DelayedRedeemsClaimed event raised by the Delayredeemrouter contract.
type DelayredeemrouterDelayedRedeemsClaimed struct {
	Recipient     common.Address
	Token         common.Address
	ClaimedAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelayedRedeemsClaimed is a free log retrieval operation binding the contract event 0xea08d9fa9c1ac98b666df5fdb636c7cda43a9b75c4d0e84088b1510d5d2396ea.
//
// Solidity: event DelayedRedeemsClaimed(address recipient, address token, uint256 claimedAmount)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterDelayedRedeemsClaimed(opts *bind.FilterOpts) (*DelayredeemrouterDelayedRedeemsClaimedIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "DelayedRedeemsClaimed")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterDelayedRedeemsClaimedIterator{contract: _Delayredeemrouter.contract, event: "DelayedRedeemsClaimed", logs: logs, sub: sub}, nil
}

// WatchDelayedRedeemsClaimed is a free log subscription operation binding the contract event 0xea08d9fa9c1ac98b666df5fdb636c7cda43a9b75c4d0e84088b1510d5d2396ea.
//
// Solidity: event DelayedRedeemsClaimed(address recipient, address token, uint256 claimedAmount)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchDelayedRedeemsClaimed(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterDelayedRedeemsClaimed) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "DelayedRedeemsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterDelayedRedeemsClaimed)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "DelayedRedeemsClaimed", log); err != nil {
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

// ParseDelayedRedeemsClaimed is a log parse operation binding the contract event 0xea08d9fa9c1ac98b666df5fdb636c7cda43a9b75c4d0e84088b1510d5d2396ea.
//
// Solidity: event DelayedRedeemsClaimed(address recipient, address token, uint256 claimedAmount)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseDelayedRedeemsClaimed(log types.Log) (*DelayredeemrouterDelayedRedeemsClaimed, error) {
	event := new(DelayredeemrouterDelayedRedeemsClaimed)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "DelayedRedeemsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterDelayedRedeemsCompletedIterator is returned from FilterDelayedRedeemsCompleted and is used to iterate over the raw logs and unpacked data for DelayedRedeemsCompleted events raised by the Delayredeemrouter contract.
type DelayredeemrouterDelayedRedeemsCompletedIterator struct {
	Event *DelayredeemrouterDelayedRedeemsCompleted // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterDelayedRedeemsCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterDelayedRedeemsCompleted)
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
		it.Event = new(DelayredeemrouterDelayedRedeemsCompleted)
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
func (it *DelayredeemrouterDelayedRedeemsCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterDelayedRedeemsCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterDelayedRedeemsCompleted represents a DelayedRedeemsCompleted event raised by the Delayredeemrouter contract.
type DelayredeemrouterDelayedRedeemsCompleted struct {
	Recipient               common.Address
	BurnedAmount            *big.Int
	DelayedRedeemsCompleted *big.Int
	TotalFee                *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterDelayedRedeemsCompleted is a free log retrieval operation binding the contract event 0x3e3e78b8ae4f7de657f7dfdf7ab56dbd00fa2fe773d425bd6911a47450e4c3c5.
//
// Solidity: event DelayedRedeemsCompleted(address recipient, uint256 burnedAmount, uint256 delayedRedeemsCompleted, uint256 totalFee)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterDelayedRedeemsCompleted(opts *bind.FilterOpts) (*DelayredeemrouterDelayedRedeemsCompletedIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "DelayedRedeemsCompleted")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterDelayedRedeemsCompletedIterator{contract: _Delayredeemrouter.contract, event: "DelayedRedeemsCompleted", logs: logs, sub: sub}, nil
}

// WatchDelayedRedeemsCompleted is a free log subscription operation binding the contract event 0x3e3e78b8ae4f7de657f7dfdf7ab56dbd00fa2fe773d425bd6911a47450e4c3c5.
//
// Solidity: event DelayedRedeemsCompleted(address recipient, uint256 burnedAmount, uint256 delayedRedeemsCompleted, uint256 totalFee)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchDelayedRedeemsCompleted(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterDelayedRedeemsCompleted) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "DelayedRedeemsCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterDelayedRedeemsCompleted)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "DelayedRedeemsCompleted", log); err != nil {
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

// ParseDelayedRedeemsCompleted is a log parse operation binding the contract event 0x3e3e78b8ae4f7de657f7dfdf7ab56dbd00fa2fe773d425bd6911a47450e4c3c5.
//
// Solidity: event DelayedRedeemsCompleted(address recipient, uint256 burnedAmount, uint256 delayedRedeemsCompleted, uint256 totalFee)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseDelayedRedeemsCompleted(log types.Log) (*DelayredeemrouterDelayedRedeemsCompleted, error) {
	event := new(DelayredeemrouterDelayedRedeemsCompleted)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "DelayedRedeemsCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterFastLaneSetIterator is returned from FilterFastLaneSet and is used to iterate over the raw logs and unpacked data for FastLaneSet events raised by the Delayredeemrouter contract.
type DelayredeemrouterFastLaneSetIterator struct {
	Event *DelayredeemrouterFastLaneSet // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterFastLaneSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterFastLaneSet)
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
		it.Event = new(DelayredeemrouterFastLaneSet)
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
func (it *DelayredeemrouterFastLaneSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterFastLaneSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterFastLaneSet represents a FastLaneSet event raised by the Delayredeemrouter contract.
type DelayredeemrouterFastLaneSet struct {
	Accounts       []common.Address
	FastLaneStatus []bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFastLaneSet is a free log retrieval operation binding the contract event 0x13ecb27a093a9554bc7f0862f15885686d9b5db6454fe1bfb686395523382d86.
//
// Solidity: event FastLaneSet(address[] accounts, bool[] fastLaneStatus)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterFastLaneSet(opts *bind.FilterOpts) (*DelayredeemrouterFastLaneSetIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "FastLaneSet")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterFastLaneSetIterator{contract: _Delayredeemrouter.contract, event: "FastLaneSet", logs: logs, sub: sub}, nil
}

// WatchFastLaneSet is a free log subscription operation binding the contract event 0x13ecb27a093a9554bc7f0862f15885686d9b5db6454fe1bfb686395523382d86.
//
// Solidity: event FastLaneSet(address[] accounts, bool[] fastLaneStatus)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchFastLaneSet(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterFastLaneSet) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "FastLaneSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterFastLaneSet)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "FastLaneSet", log); err != nil {
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

// ParseFastLaneSet is a log parse operation binding the contract event 0x13ecb27a093a9554bc7f0862f15885686d9b5db6454fe1bfb686395523382d86.
//
// Solidity: event FastLaneSet(address[] accounts, bool[] fastLaneStatus)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseFastLaneSet(log types.Log) (*DelayredeemrouterFastLaneSet, error) {
	event := new(DelayredeemrouterFastLaneSet)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "FastLaneSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Delayredeemrouter contract.
type DelayredeemrouterInitializedIterator struct {
	Event *DelayredeemrouterInitialized // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterInitialized)
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
		it.Event = new(DelayredeemrouterInitialized)
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
func (it *DelayredeemrouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterInitialized represents a Initialized event raised by the Delayredeemrouter contract.
type DelayredeemrouterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*DelayredeemrouterInitializedIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterInitializedIterator{contract: _Delayredeemrouter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterInitialized) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterInitialized)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseInitialized(log types.Log) (*DelayredeemrouterInitialized, error) {
	event := new(DelayredeemrouterInitialized)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterManagementFeeWithdrawnIterator is returned from FilterManagementFeeWithdrawn and is used to iterate over the raw logs and unpacked data for ManagementFeeWithdrawn events raised by the Delayredeemrouter contract.
type DelayredeemrouterManagementFeeWithdrawnIterator struct {
	Event *DelayredeemrouterManagementFeeWithdrawn // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterManagementFeeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterManagementFeeWithdrawn)
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
		it.Event = new(DelayredeemrouterManagementFeeWithdrawn)
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
func (it *DelayredeemrouterManagementFeeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterManagementFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterManagementFeeWithdrawn represents a ManagementFeeWithdrawn event raised by the Delayredeemrouter contract.
type DelayredeemrouterManagementFeeWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterManagementFeeWithdrawn is a free log retrieval operation binding the contract event 0x583d1744b1f2b01833f7f10ff38436dd7a76ff50695a487bfb9f0f3d07749b49.
//
// Solidity: event ManagementFeeWithdrawn(address recipient, uint256 amount)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterManagementFeeWithdrawn(opts *bind.FilterOpts) (*DelayredeemrouterManagementFeeWithdrawnIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "ManagementFeeWithdrawn")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterManagementFeeWithdrawnIterator{contract: _Delayredeemrouter.contract, event: "ManagementFeeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchManagementFeeWithdrawn is a free log subscription operation binding the contract event 0x583d1744b1f2b01833f7f10ff38436dd7a76ff50695a487bfb9f0f3d07749b49.
//
// Solidity: event ManagementFeeWithdrawn(address recipient, uint256 amount)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchManagementFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterManagementFeeWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "ManagementFeeWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterManagementFeeWithdrawn)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "ManagementFeeWithdrawn", log); err != nil {
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

// ParseManagementFeeWithdrawn is a log parse operation binding the contract event 0x583d1744b1f2b01833f7f10ff38436dd7a76ff50695a487bfb9f0f3d07749b49.
//
// Solidity: event ManagementFeeWithdrawn(address recipient, uint256 amount)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseManagementFeeWithdrawn(log types.Log) (*DelayredeemrouterManagementFeeWithdrawn, error) {
	event := new(DelayredeemrouterManagementFeeWithdrawn)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "ManagementFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterMaxQuotaSetIterator is returned from FilterMaxQuotaSet and is used to iterate over the raw logs and unpacked data for MaxQuotaSet events raised by the Delayredeemrouter contract.
type DelayredeemrouterMaxQuotaSetIterator struct {
	Event *DelayredeemrouterMaxQuotaSet // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterMaxQuotaSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterMaxQuotaSet)
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
		it.Event = new(DelayredeemrouterMaxQuotaSet)
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
func (it *DelayredeemrouterMaxQuotaSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterMaxQuotaSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterMaxQuotaSet represents a MaxQuotaSet event raised by the Delayredeemrouter contract.
type DelayredeemrouterMaxQuotaSet struct {
	Token         common.Address
	PreviousQuota *big.Int
	NewQuota      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxQuotaSet is a free log retrieval operation binding the contract event 0x96fa1888abd5b7a22236027bf87904c3f144b181c9c8016594b72c9b47f94d79.
//
// Solidity: event MaxQuotaSet(address token, uint256 previousQuota, uint256 newQuota)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterMaxQuotaSet(opts *bind.FilterOpts) (*DelayredeemrouterMaxQuotaSetIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "MaxQuotaSet")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterMaxQuotaSetIterator{contract: _Delayredeemrouter.contract, event: "MaxQuotaSet", logs: logs, sub: sub}, nil
}

// WatchMaxQuotaSet is a free log subscription operation binding the contract event 0x96fa1888abd5b7a22236027bf87904c3f144b181c9c8016594b72c9b47f94d79.
//
// Solidity: event MaxQuotaSet(address token, uint256 previousQuota, uint256 newQuota)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchMaxQuotaSet(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterMaxQuotaSet) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "MaxQuotaSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterMaxQuotaSet)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "MaxQuotaSet", log); err != nil {
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

// ParseMaxQuotaSet is a log parse operation binding the contract event 0x96fa1888abd5b7a22236027bf87904c3f144b181c9c8016594b72c9b47f94d79.
//
// Solidity: event MaxQuotaSet(address token, uint256 previousQuota, uint256 newQuota)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseMaxQuotaSet(log types.Log) (*DelayredeemrouterMaxQuotaSet, error) {
	event := new(DelayredeemrouterMaxQuotaSet)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "MaxQuotaSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Delayredeemrouter contract.
type DelayredeemrouterPausedIterator struct {
	Event *DelayredeemrouterPaused // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterPaused)
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
		it.Event = new(DelayredeemrouterPaused)
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
func (it *DelayredeemrouterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterPaused represents a Paused event raised by the Delayredeemrouter contract.
type DelayredeemrouterPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterPaused(opts *bind.FilterOpts) (*DelayredeemrouterPausedIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterPausedIterator{contract: _Delayredeemrouter.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterPaused) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterPaused)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParsePaused(log types.Log) (*DelayredeemrouterPaused, error) {
	event := new(DelayredeemrouterPaused)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterRateSetIterator is returned from FilterRateSet and is used to iterate over the raw logs and unpacked data for RateSet events raised by the Delayredeemrouter contract.
type DelayredeemrouterRateSetIterator struct {
	Event *DelayredeemrouterRateSet // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterRateSet)
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
		it.Event = new(DelayredeemrouterRateSet)
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
func (it *DelayredeemrouterRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterRateSet represents a RateSet event raised by the Delayredeemrouter contract.
type DelayredeemrouterRateSet struct {
	Token         common.Address
	PreviousQuota *big.Int
	NewQuota      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRateSet is a free log retrieval operation binding the contract event 0x9e31cca092b9e764bfc6b1b552d55ad4b035e609318fecc26cd38b34e8dd08bb.
//
// Solidity: event RateSet(address token, uint256 previousQuota, uint256 newQuota)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterRateSet(opts *bind.FilterOpts) (*DelayredeemrouterRateSetIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "RateSet")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterRateSetIterator{contract: _Delayredeemrouter.contract, event: "RateSet", logs: logs, sub: sub}, nil
}

// WatchRateSet is a free log subscription operation binding the contract event 0x9e31cca092b9e764bfc6b1b552d55ad4b035e609318fecc26cd38b34e8dd08bb.
//
// Solidity: event RateSet(address token, uint256 previousQuota, uint256 newQuota)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchRateSet(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterRateSet) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "RateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterRateSet)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "RateSet", log); err != nil {
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

// ParseRateSet is a log parse operation binding the contract event 0x9e31cca092b9e764bfc6b1b552d55ad4b035e609318fecc26cd38b34e8dd08bb.
//
// Solidity: event RateSet(address token, uint256 previousQuota, uint256 newQuota)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseRateSet(log types.Log) (*DelayredeemrouterRateSet, error) {
	event := new(DelayredeemrouterRateSet)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "RateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterRedeemDelaySetIterator is returned from FilterRedeemDelaySet and is used to iterate over the raw logs and unpacked data for RedeemDelaySet events raised by the Delayredeemrouter contract.
type DelayredeemrouterRedeemDelaySetIterator struct {
	Event *DelayredeemrouterRedeemDelaySet // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterRedeemDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterRedeemDelaySet)
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
		it.Event = new(DelayredeemrouterRedeemDelaySet)
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
func (it *DelayredeemrouterRedeemDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterRedeemDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterRedeemDelaySet represents a RedeemDelaySet event raised by the Delayredeemrouter contract.
type DelayredeemrouterRedeemDelaySet struct {
	PreviousDelay *big.Int
	NewDelay      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRedeemDelaySet is a free log retrieval operation binding the contract event 0x1431e66d652872e09a583cd3d9bb280e0f743dca287cc9344a1d80a1596add3a.
//
// Solidity: event RedeemDelaySet(uint256 previousDelay, uint256 newDelay)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterRedeemDelaySet(opts *bind.FilterOpts) (*DelayredeemrouterRedeemDelaySetIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "RedeemDelaySet")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterRedeemDelaySetIterator{contract: _Delayredeemrouter.contract, event: "RedeemDelaySet", logs: logs, sub: sub}, nil
}

// WatchRedeemDelaySet is a free log subscription operation binding the contract event 0x1431e66d652872e09a583cd3d9bb280e0f743dca287cc9344a1d80a1596add3a.
//
// Solidity: event RedeemDelaySet(uint256 previousDelay, uint256 newDelay)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchRedeemDelaySet(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterRedeemDelaySet) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "RedeemDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterRedeemDelaySet)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "RedeemDelaySet", log); err != nil {
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

// ParseRedeemDelaySet is a log parse operation binding the contract event 0x1431e66d652872e09a583cd3d9bb280e0f743dca287cc9344a1d80a1596add3a.
//
// Solidity: event RedeemDelaySet(uint256 previousDelay, uint256 newDelay)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseRedeemDelaySet(log types.Log) (*DelayredeemrouterRedeemDelaySet, error) {
	event := new(DelayredeemrouterRedeemDelaySet)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "RedeemDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterRedeemFeeRateSetIterator is returned from FilterRedeemFeeRateSet and is used to iterate over the raw logs and unpacked data for RedeemFeeRateSet events raised by the Delayredeemrouter contract.
type DelayredeemrouterRedeemFeeRateSetIterator struct {
	Event *DelayredeemrouterRedeemFeeRateSet // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterRedeemFeeRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterRedeemFeeRateSet)
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
		it.Event = new(DelayredeemrouterRedeemFeeRateSet)
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
func (it *DelayredeemrouterRedeemFeeRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterRedeemFeeRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterRedeemFeeRateSet represents a RedeemFeeRateSet event raised by the Delayredeemrouter contract.
type DelayredeemrouterRedeemFeeRateSet struct {
	PreviousFeeRate *big.Int
	NewFeeRate      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRedeemFeeRateSet is a free log retrieval operation binding the contract event 0x4d9ff9777a971177b4810d0c671707ff05f2469b58efd13a676a0eb42fe53528.
//
// Solidity: event RedeemFeeRateSet(uint256 previousFeeRate, uint256 newFeeRate)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterRedeemFeeRateSet(opts *bind.FilterOpts) (*DelayredeemrouterRedeemFeeRateSetIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "RedeemFeeRateSet")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterRedeemFeeRateSetIterator{contract: _Delayredeemrouter.contract, event: "RedeemFeeRateSet", logs: logs, sub: sub}, nil
}

// WatchRedeemFeeRateSet is a free log subscription operation binding the contract event 0x4d9ff9777a971177b4810d0c671707ff05f2469b58efd13a676a0eb42fe53528.
//
// Solidity: event RedeemFeeRateSet(uint256 previousFeeRate, uint256 newFeeRate)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchRedeemFeeRateSet(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterRedeemFeeRateSet) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "RedeemFeeRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterRedeemFeeRateSet)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "RedeemFeeRateSet", log); err != nil {
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

// ParseRedeemFeeRateSet is a log parse operation binding the contract event 0x4d9ff9777a971177b4810d0c671707ff05f2469b58efd13a676a0eb42fe53528.
//
// Solidity: event RedeemFeeRateSet(uint256 previousFeeRate, uint256 newFeeRate)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseRedeemFeeRateSet(log types.Log) (*DelayredeemrouterRedeemFeeRateSet, error) {
	event := new(DelayredeemrouterRedeemFeeRateSet)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "RedeemFeeRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterRetainAmountsSetIterator is returned from FilterRetainAmountsSet and is used to iterate over the raw logs and unpacked data for RetainAmountsSet events raised by the Delayredeemrouter contract.
type DelayredeemrouterRetainAmountsSetIterator struct {
	Event *DelayredeemrouterRetainAmountsSet // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterRetainAmountsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterRetainAmountsSet)
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
		it.Event = new(DelayredeemrouterRetainAmountsSet)
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
func (it *DelayredeemrouterRetainAmountsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterRetainAmountsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterRetainAmountsSet represents a RetainAmountsSet event raised by the Delayredeemrouter contract.
type DelayredeemrouterRetainAmountsSet struct {
	Tokens   []common.Address
	Balances []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRetainAmountsSet is a free log retrieval operation binding the contract event 0xca1add6d45e7e4d6159a5bfaab5635b7411fe4fb9b15f10301ed3d66273d1e2a.
//
// Solidity: event RetainAmountsSet(address[] tokens, uint256[] balances)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterRetainAmountsSet(opts *bind.FilterOpts) (*DelayredeemrouterRetainAmountsSetIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "RetainAmountsSet")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterRetainAmountsSetIterator{contract: _Delayredeemrouter.contract, event: "RetainAmountsSet", logs: logs, sub: sub}, nil
}

// WatchRetainAmountsSet is a free log subscription operation binding the contract event 0xca1add6d45e7e4d6159a5bfaab5635b7411fe4fb9b15f10301ed3d66273d1e2a.
//
// Solidity: event RetainAmountsSet(address[] tokens, uint256[] balances)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchRetainAmountsSet(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterRetainAmountsSet) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "RetainAmountsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterRetainAmountsSet)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "RetainAmountsSet", log); err != nil {
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

// ParseRetainAmountsSet is a log parse operation binding the contract event 0xca1add6d45e7e4d6159a5bfaab5635b7411fe4fb9b15f10301ed3d66273d1e2a.
//
// Solidity: event RetainAmountsSet(address[] tokens, uint256[] balances)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseRetainAmountsSet(log types.Log) (*DelayredeemrouterRetainAmountsSet, error) {
	event := new(DelayredeemrouterRetainAmountsSet)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "RetainAmountsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Delayredeemrouter contract.
type DelayredeemrouterRoleAdminChangedIterator struct {
	Event *DelayredeemrouterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterRoleAdminChanged)
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
		it.Event = new(DelayredeemrouterRoleAdminChanged)
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
func (it *DelayredeemrouterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterRoleAdminChanged represents a RoleAdminChanged event raised by the Delayredeemrouter contract.
type DelayredeemrouterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DelayredeemrouterRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterRoleAdminChangedIterator{contract: _Delayredeemrouter.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterRoleAdminChanged)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseRoleAdminChanged(log types.Log) (*DelayredeemrouterRoleAdminChanged, error) {
	event := new(DelayredeemrouterRoleAdminChanged)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Delayredeemrouter contract.
type DelayredeemrouterRoleGrantedIterator struct {
	Event *DelayredeemrouterRoleGranted // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterRoleGranted)
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
		it.Event = new(DelayredeemrouterRoleGranted)
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
func (it *DelayredeemrouterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterRoleGranted represents a RoleGranted event raised by the Delayredeemrouter contract.
type DelayredeemrouterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DelayredeemrouterRoleGrantedIterator, error) {

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

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterRoleGrantedIterator{contract: _Delayredeemrouter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterRoleGranted)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseRoleGranted(log types.Log) (*DelayredeemrouterRoleGranted, error) {
	event := new(DelayredeemrouterRoleGranted)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Delayredeemrouter contract.
type DelayredeemrouterRoleRevokedIterator struct {
	Event *DelayredeemrouterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterRoleRevoked)
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
		it.Event = new(DelayredeemrouterRoleRevoked)
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
func (it *DelayredeemrouterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterRoleRevoked represents a RoleRevoked event raised by the Delayredeemrouter contract.
type DelayredeemrouterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DelayredeemrouterRoleRevokedIterator, error) {

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

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterRoleRevokedIterator{contract: _Delayredeemrouter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterRoleRevoked)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseRoleRevoked(log types.Log) (*DelayredeemrouterRoleRevoked, error) {
	event := new(DelayredeemrouterRoleRevoked)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterTokensPausedIterator is returned from FilterTokensPaused and is used to iterate over the raw logs and unpacked data for TokensPaused events raised by the Delayredeemrouter contract.
type DelayredeemrouterTokensPausedIterator struct {
	Event *DelayredeemrouterTokensPaused // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterTokensPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterTokensPaused)
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
		it.Event = new(DelayredeemrouterTokensPaused)
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
func (it *DelayredeemrouterTokensPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterTokensPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterTokensPaused represents a TokensPaused event raised by the Delayredeemrouter contract.
type DelayredeemrouterTokensPaused struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensPaused is a free log retrieval operation binding the contract event 0xa4fbb323c047ef2555d72263081bbb4280ee59d506657303bf5fc991b54204be.
//
// Solidity: event TokensPaused(address[] tokens)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterTokensPaused(opts *bind.FilterOpts) (*DelayredeemrouterTokensPausedIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "TokensPaused")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterTokensPausedIterator{contract: _Delayredeemrouter.contract, event: "TokensPaused", logs: logs, sub: sub}, nil
}

// WatchTokensPaused is a free log subscription operation binding the contract event 0xa4fbb323c047ef2555d72263081bbb4280ee59d506657303bf5fc991b54204be.
//
// Solidity: event TokensPaused(address[] tokens)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchTokensPaused(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterTokensPaused) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "TokensPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterTokensPaused)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "TokensPaused", log); err != nil {
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

// ParseTokensPaused is a log parse operation binding the contract event 0xa4fbb323c047ef2555d72263081bbb4280ee59d506657303bf5fc991b54204be.
//
// Solidity: event TokensPaused(address[] tokens)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseTokensPaused(log types.Log) (*DelayredeemrouterTokensPaused, error) {
	event := new(DelayredeemrouterTokensPaused)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "TokensPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterTokensUnpausedIterator is returned from FilterTokensUnpaused and is used to iterate over the raw logs and unpacked data for TokensUnpaused events raised by the Delayredeemrouter contract.
type DelayredeemrouterTokensUnpausedIterator struct {
	Event *DelayredeemrouterTokensUnpaused // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterTokensUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterTokensUnpaused)
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
		it.Event = new(DelayredeemrouterTokensUnpaused)
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
func (it *DelayredeemrouterTokensUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterTokensUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterTokensUnpaused represents a TokensUnpaused event raised by the Delayredeemrouter contract.
type DelayredeemrouterTokensUnpaused struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensUnpaused is a free log retrieval operation binding the contract event 0x4dd04d346e64df7bcc65df29a0f1f1f84815ff758717f30587cabe38792d7c31.
//
// Solidity: event TokensUnpaused(address[] tokens)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterTokensUnpaused(opts *bind.FilterOpts) (*DelayredeemrouterTokensUnpausedIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "TokensUnpaused")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterTokensUnpausedIterator{contract: _Delayredeemrouter.contract, event: "TokensUnpaused", logs: logs, sub: sub}, nil
}

// WatchTokensUnpaused is a free log subscription operation binding the contract event 0x4dd04d346e64df7bcc65df29a0f1f1f84815ff758717f30587cabe38792d7c31.
//
// Solidity: event TokensUnpaused(address[] tokens)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchTokensUnpaused(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterTokensUnpaused) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "TokensUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterTokensUnpaused)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "TokensUnpaused", log); err != nil {
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

// ParseTokensUnpaused is a log parse operation binding the contract event 0x4dd04d346e64df7bcc65df29a0f1f1f84815ff758717f30587cabe38792d7c31.
//
// Solidity: event TokensUnpaused(address[] tokens)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseTokensUnpaused(log types.Log) (*DelayredeemrouterTokensUnpaused, error) {
	event := new(DelayredeemrouterTokensUnpaused)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "TokensUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Delayredeemrouter contract.
type DelayredeemrouterUnpausedIterator struct {
	Event *DelayredeemrouterUnpaused // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterUnpaused)
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
		it.Event = new(DelayredeemrouterUnpaused)
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
func (it *DelayredeemrouterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterUnpaused represents a Unpaused event raised by the Delayredeemrouter contract.
type DelayredeemrouterUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterUnpaused(opts *bind.FilterOpts) (*DelayredeemrouterUnpausedIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterUnpausedIterator{contract: _Delayredeemrouter.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterUnpaused) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterUnpaused)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseUnpaused(log types.Log) (*DelayredeemrouterUnpaused, error) {
	event := new(DelayredeemrouterUnpaused)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterWhitelistAddedIterator is returned from FilterWhitelistAdded and is used to iterate over the raw logs and unpacked data for WhitelistAdded events raised by the Delayredeemrouter contract.
type DelayredeemrouterWhitelistAddedIterator struct {
	Event *DelayredeemrouterWhitelistAdded // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterWhitelistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterWhitelistAdded)
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
		it.Event = new(DelayredeemrouterWhitelistAdded)
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
func (it *DelayredeemrouterWhitelistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterWhitelistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterWhitelistAdded represents a WhitelistAdded event raised by the Delayredeemrouter contract.
type DelayredeemrouterWhitelistAdded struct {
	Accounts []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdded is a free log retrieval operation binding the contract event 0xf74f148a4f930a0f67a2c33ba932a14e3e91b4e6468f21e545932fd825111538.
//
// Solidity: event WhitelistAdded(address[] accounts)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterWhitelistAdded(opts *bind.FilterOpts) (*DelayredeemrouterWhitelistAddedIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "WhitelistAdded")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterWhitelistAddedIterator{contract: _Delayredeemrouter.contract, event: "WhitelistAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdded is a free log subscription operation binding the contract event 0xf74f148a4f930a0f67a2c33ba932a14e3e91b4e6468f21e545932fd825111538.
//
// Solidity: event WhitelistAdded(address[] accounts)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchWhitelistAdded(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterWhitelistAdded) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "WhitelistAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterWhitelistAdded)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "WhitelistAdded", log); err != nil {
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

// ParseWhitelistAdded is a log parse operation binding the contract event 0xf74f148a4f930a0f67a2c33ba932a14e3e91b4e6468f21e545932fd825111538.
//
// Solidity: event WhitelistAdded(address[] accounts)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseWhitelistAdded(log types.Log) (*DelayredeemrouterWhitelistAdded, error) {
	event := new(DelayredeemrouterWhitelistAdded)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "WhitelistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterWhitelistEnabledSetIterator is returned from FilterWhitelistEnabledSet and is used to iterate over the raw logs and unpacked data for WhitelistEnabledSet events raised by the Delayredeemrouter contract.
type DelayredeemrouterWhitelistEnabledSetIterator struct {
	Event *DelayredeemrouterWhitelistEnabledSet // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterWhitelistEnabledSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterWhitelistEnabledSet)
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
		it.Event = new(DelayredeemrouterWhitelistEnabledSet)
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
func (it *DelayredeemrouterWhitelistEnabledSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterWhitelistEnabledSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterWhitelistEnabledSet represents a WhitelistEnabledSet event raised by the Delayredeemrouter contract.
type DelayredeemrouterWhitelistEnabledSet struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistEnabledSet is a free log retrieval operation binding the contract event 0x411283ae1b0e68089790510eb77ccad9b761295be576637799607c8ae066fe9f.
//
// Solidity: event WhitelistEnabledSet(bool enabled)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterWhitelistEnabledSet(opts *bind.FilterOpts) (*DelayredeemrouterWhitelistEnabledSetIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "WhitelistEnabledSet")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterWhitelistEnabledSetIterator{contract: _Delayredeemrouter.contract, event: "WhitelistEnabledSet", logs: logs, sub: sub}, nil
}

// WatchWhitelistEnabledSet is a free log subscription operation binding the contract event 0x411283ae1b0e68089790510eb77ccad9b761295be576637799607c8ae066fe9f.
//
// Solidity: event WhitelistEnabledSet(bool enabled)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchWhitelistEnabledSet(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterWhitelistEnabledSet) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "WhitelistEnabledSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterWhitelistEnabledSet)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "WhitelistEnabledSet", log); err != nil {
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

// ParseWhitelistEnabledSet is a log parse operation binding the contract event 0x411283ae1b0e68089790510eb77ccad9b761295be576637799607c8ae066fe9f.
//
// Solidity: event WhitelistEnabledSet(bool enabled)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseWhitelistEnabledSet(log types.Log) (*DelayredeemrouterWhitelistEnabledSet, error) {
	event := new(DelayredeemrouterWhitelistEnabledSet)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "WhitelistEnabledSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayredeemrouterWhitelistRemovedIterator is returned from FilterWhitelistRemoved and is used to iterate over the raw logs and unpacked data for WhitelistRemoved events raised by the Delayredeemrouter contract.
type DelayredeemrouterWhitelistRemovedIterator struct {
	Event *DelayredeemrouterWhitelistRemoved // Event containing the contract specifics and raw log

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
func (it *DelayredeemrouterWhitelistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayredeemrouterWhitelistRemoved)
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
		it.Event = new(DelayredeemrouterWhitelistRemoved)
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
func (it *DelayredeemrouterWhitelistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayredeemrouterWhitelistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayredeemrouterWhitelistRemoved represents a WhitelistRemoved event raised by the Delayredeemrouter contract.
type DelayredeemrouterWhitelistRemoved struct {
	Accounts []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistRemoved is a free log retrieval operation binding the contract event 0x1d474f57a5c483b47a8bf6006e39086f96dd040a00cb348e22f80a4ca2c6f222.
//
// Solidity: event WhitelistRemoved(address[] accounts)
func (_Delayredeemrouter *DelayredeemrouterFilterer) FilterWhitelistRemoved(opts *bind.FilterOpts) (*DelayredeemrouterWhitelistRemovedIterator, error) {

	logs, sub, err := _Delayredeemrouter.contract.FilterLogs(opts, "WhitelistRemoved")
	if err != nil {
		return nil, err
	}
	return &DelayredeemrouterWhitelistRemovedIterator{contract: _Delayredeemrouter.contract, event: "WhitelistRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistRemoved is a free log subscription operation binding the contract event 0x1d474f57a5c483b47a8bf6006e39086f96dd040a00cb348e22f80a4ca2c6f222.
//
// Solidity: event WhitelistRemoved(address[] accounts)
func (_Delayredeemrouter *DelayredeemrouterFilterer) WatchWhitelistRemoved(opts *bind.WatchOpts, sink chan<- *DelayredeemrouterWhitelistRemoved) (event.Subscription, error) {

	logs, sub, err := _Delayredeemrouter.contract.WatchLogs(opts, "WhitelistRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayredeemrouterWhitelistRemoved)
				if err := _Delayredeemrouter.contract.UnpackLog(event, "WhitelistRemoved", log); err != nil {
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

// ParseWhitelistRemoved is a log parse operation binding the contract event 0x1d474f57a5c483b47a8bf6006e39086f96dd040a00cb348e22f80a4ca2c6f222.
//
// Solidity: event WhitelistRemoved(address[] accounts)
func (_Delayredeemrouter *DelayredeemrouterFilterer) ParseWhitelistRemoved(log types.Log) (*DelayredeemrouterWhitelistRemoved, error) {
	event := new(DelayredeemrouterWhitelistRemoved)
	if err := _Delayredeemrouter.contract.UnpackLog(event, "WhitelistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
