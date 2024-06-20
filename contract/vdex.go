// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// BaseOrderUtilsCancelOrderParams is an auto generated low-level Go binding around an user-defined struct.
type BaseOrderUtilsCancelOrderParams struct {
	Order           Order
	WalletSignature []byte
	CancellationFee *big.Int
}

// BaseOrderUtilsCreateOrderParams is an auto generated low-level Go binding around an user-defined struct.
type BaseOrderUtilsCreateOrderParams struct {
	SwapPath             []common.Address
	Nonce                *big.Int
	QuantityInAssetUnits *big.Int
	TriggerPrice         *big.Int
	AcceptablePrice      *big.Int
	Expiration           *big.Int
	OrderType            uint8
	Side                 uint8
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Trader           common.Address
	BaseAsset        common.Address
	QuoteAsset       common.Address
	BaseAssetAmount  *big.Int
	QuoteAssetAmount *big.Int
	OrderType        uint8
	Side             uint8
	Expiration       *big.Int
	Nonce            *big.Int
}

// OrderBookTrade is an auto generated low-level Go binding around an user-defined struct.
type OrderBookTrade struct {
	BaseAssetAddress     common.Address
	QuoteAssetAddress    common.Address
	GrossBaseQuantity    *big.Int
	GrossQuoteQuantity   *big.Int
	NetBaseQuantity      *big.Int
	NetQuoteQuantity     *big.Int
	MakerFeeAssetAddress common.Address
	TakerFeeAssetAddress common.Address
	MakerFeeQuantity     *big.Int
	TakerFeeQuantity     *big.Int
	Price                *big.Int
	MakerSide            uint8
}

// TypesOrderParam is an auto generated low-level Go binding around an user-defined struct.
type TypesOrderParam struct {
	Trader           common.Address
	BaseAssetAmount  *big.Int
	QuoteAssetAmount *big.Int
	OrderType        uint8
	Side             uint8
	Expiration       *big.Int
	Nonce            *big.Int
	WalletSignature  []byte
}

// VdexMetaData contains all meta data concerning the Vdex contract.
var VdexMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdvanceNonceFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpiredOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GaslessInvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GaslessOrderInvalidAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderType\",\"type\":\"uint256\"}],\"name\":\"OrderTypeCannotBeCreated\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAssetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAssetAmount\",\"type\":\"uint256\"}],\"name\":\"CancelSwapOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sizeDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockTime\",\"type\":\"uint256\"}],\"name\":\"CreateDecreasePosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sizeDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"CreateIncreasePosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"quantityInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newExchangeBalanceInPips\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExchangeBalanceInAssetUnits\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sizeDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockGap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeGap\",\"type\":\"uint256\"}],\"name\":\"ExecuteDecreasePosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sizeDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockGap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeGap\",\"type\":\"uint256\"}],\"name\":\"ExecuteIncreasePosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"NewFactoryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sellWallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyerInputUpdate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellerInputUpdate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyerOutputUpdate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellerOutputUpdate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumOrderSide\",\"name\":\"takerSide\",\"type\":\"uint8\"}],\"name\":\"OrderBookTradeExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"makerAccountOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAssetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAssetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderSide\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"WalletExitCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"effectiveBlockNumber\",\"type\":\"uint256\"}],\"name\":\"WalletExited\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_walletExits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"effectiveBlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDispatcherWallet\",\"type\":\"address\"}],\"name\":\"addDispatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wETH\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"advanceNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callCoinPermit\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumOrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"walletSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"cancellationFee\",\"type\":\"uint256\"}],\"internalType\":\"structBaseOrderUtils.CancelOrderParams[]\",\"name\":\"_swapOrders\",\"type\":\"tuple[]\"}],\"name\":\"cancelMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"swapPath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityInAssetUnits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumOrderSide\",\"name\":\"side\",\"type\":\"uint8\"}],\"internalType\":\"structBaseOrderUtils.CreateOrderParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"swapPath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityInAssetUnits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumOrderSide\",\"name\":\"side\",\"type\":\"uint8\"}],\"internalType\":\"structBaseOrderUtils.CreateOrderParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"createOrderETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumOrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structOrder\",\"name\":\"order_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"createOrderWithPermit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"dispatcherWallets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumOrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"walletSignature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.OrderParam\",\"name\":\"buy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"enumOrderSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"walletSignature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.OrderParam\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"baseAssetAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteAssetAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"grossBaseQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"grossQuoteQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netBaseQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netQuoteQuantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerFeeAssetAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerFeeAssetAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerFeeQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeeQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"enumOrderSide\",\"name\":\"makerSide\",\"type\":\"uint8\"}],\"internalType\":\"structOrderBookTrade\",\"name\":\"orderBookTrade\",\"type\":\"tuple\"}],\"name\":\"executeOrderBookTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filledLimitBuyQuoteAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gaslessOrderNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increaseNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dispatcherWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callCoinPermit\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loadFeeWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"contractLPoolStorage\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"marginLimit\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceUpdater\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pool0Insurance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pool1Insurance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerNonce\",\"type\":\"uint256\"}],\"name\":\"nonceEquals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numPairs\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dispatcherWallet\",\"type\":\"address\"}],\"name\":\"reverseDispatcherState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"setFactoryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeWallet\",\"type\":\"address\"}],\"name\":\"setFeeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalHelds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// VdexABI is the input ABI used to generate the binding from.
// Deprecated: Use VdexMetaData.ABI instead.
var VdexABI = VdexMetaData.ABI

// Vdex is an auto generated Go binding around an Ethereum contract.
type Vdex struct {
	VdexCaller     // Read-only binding to the contract
	VdexTransactor // Write-only binding to the contract
	VdexFilterer   // Log filterer for contract events
}

// VdexCaller is an auto generated read-only Go binding around an Ethereum contract.
type VdexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VdexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VdexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VdexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VdexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VdexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VdexSession struct {
	Contract     *Vdex             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VdexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VdexCallerSession struct {
	Contract *VdexCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VdexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VdexTransactorSession struct {
	Contract     *VdexTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VdexRaw is an auto generated low-level Go binding around an Ethereum contract.
type VdexRaw struct {
	Contract *Vdex // Generic contract binding to access the raw methods on
}

// VdexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VdexCallerRaw struct {
	Contract *VdexCaller // Generic read-only contract binding to access the raw methods on
}

// VdexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VdexTransactorRaw struct {
	Contract *VdexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVdex creates a new instance of Vdex, bound to a specific deployed contract.
func NewVdex(address common.Address, backend bind.ContractBackend) (*Vdex, error) {
	contract, err := bindVdex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vdex{VdexCaller: VdexCaller{contract: contract}, VdexTransactor: VdexTransactor{contract: contract}, VdexFilterer: VdexFilterer{contract: contract}}, nil
}

// NewVdexCaller creates a new read-only instance of Vdex, bound to a specific deployed contract.
func NewVdexCaller(address common.Address, caller bind.ContractCaller) (*VdexCaller, error) {
	contract, err := bindVdex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VdexCaller{contract: contract}, nil
}

// NewVdexTransactor creates a new write-only instance of Vdex, bound to a specific deployed contract.
func NewVdexTransactor(address common.Address, transactor bind.ContractTransactor) (*VdexTransactor, error) {
	contract, err := bindVdex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VdexTransactor{contract: contract}, nil
}

// NewVdexFilterer creates a new log filterer instance of Vdex, bound to a specific deployed contract.
func NewVdexFilterer(address common.Address, filterer bind.ContractFilterer) (*VdexFilterer, error) {
	contract, err := bindVdex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VdexFilterer{contract: contract}, nil
}

// bindVdex binds a generic wrapper to an already deployed contract.
func bindVdex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VdexABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vdex *VdexRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vdex.Contract.VdexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vdex *VdexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vdex.Contract.VdexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vdex *VdexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vdex.Contract.VdexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vdex *VdexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vdex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vdex *VdexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vdex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vdex *VdexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vdex.Contract.contract.Transact(opts, method, params...)
}

// WalletExits is a free data retrieval call binding the contract method 0x98166c0d.
//
// Solidity: function _walletExits(address ) view returns(bool exists, uint256 effectiveBlockNumber)
func (_Vdex *VdexCaller) WalletExits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Exists               bool
	EffectiveBlockNumber *big.Int
}, error) {
	var out []interface{}
	err := _Vdex.contract.Call(opts, &out, "_walletExits", arg0)

	outstruct := new(struct {
		Exists               bool
		EffectiveBlockNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.EffectiveBlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// WalletExits is a free data retrieval call binding the contract method 0x98166c0d.
//
// Solidity: function _walletExits(address ) view returns(bool exists, uint256 effectiveBlockNumber)
func (_Vdex *VdexSession) WalletExits(arg0 common.Address) (struct {
	Exists               bool
	EffectiveBlockNumber *big.Int
}, error) {
	return _Vdex.Contract.WalletExits(&_Vdex.CallOpts, arg0)
}

// WalletExits is a free data retrieval call binding the contract method 0x98166c0d.
//
// Solidity: function _walletExits(address ) view returns(bool exists, uint256 effectiveBlockNumber)
func (_Vdex *VdexCallerSession) WalletExits(arg0 common.Address) (struct {
	Exists               bool
	EffectiveBlockNumber *big.Int
}, error) {
	return _Vdex.Contract.WalletExits(&_Vdex.CallOpts, arg0)
}

// AddressConfig is a free data retrieval call binding the contract method 0xc275fe54.
//
// Solidity: function addressConfig() view returns(address factory, address wETH)
func (_Vdex *VdexCaller) AddressConfig(opts *bind.CallOpts) (struct {
	Factory common.Address
	WETH    common.Address
}, error) {
	var out []interface{}
	err := _Vdex.contract.Call(opts, &out, "addressConfig")

	outstruct := new(struct {
		Factory common.Address
		WETH    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Factory = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.WETH = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AddressConfig is a free data retrieval call binding the contract method 0xc275fe54.
//
// Solidity: function addressConfig() view returns(address factory, address wETH)
func (_Vdex *VdexSession) AddressConfig() (struct {
	Factory common.Address
	WETH    common.Address
}, error) {
	return _Vdex.Contract.AddressConfig(&_Vdex.CallOpts)
}

// AddressConfig is a free data retrieval call binding the contract method 0xc275fe54.
//
// Solidity: function addressConfig() view returns(address factory, address wETH)
func (_Vdex *VdexCallerSession) AddressConfig() (struct {
	Factory common.Address
	WETH    common.Address
}, error) {
	return _Vdex.Contract.AddressConfig(&_Vdex.CallOpts)
}

// CallCoinPermit is a free data retrieval call binding the contract method 0x5fa0e962.
//
// Solidity: function callCoinPermit() view returns(address)
func (_Vdex *VdexCaller) CallCoinPermit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vdex.contract.Call(opts, &out, "callCoinPermit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CallCoinPermit is a free data retrieval call binding the contract method 0x5fa0e962.
//
// Solidity: function callCoinPermit() view returns(address)
func (_Vdex *VdexSession) CallCoinPermit() (common.Address, error) {
	return _Vdex.Contract.CallCoinPermit(&_Vdex.CallOpts)
}

// CallCoinPermit is a free data retrieval call binding the contract method 0x5fa0e962.
//
// Solidity: function callCoinPermit() view returns(address)
func (_Vdex *VdexCallerSession) CallCoinPermit() (common.Address, error) {
	return _Vdex.Contract.CallCoinPermit(&_Vdex.CallOpts)
}

// DispatcherWallets is a free data retrieval call binding the contract method 0x43f32683.
//
// Solidity: function dispatcherWallets(address ) view returns(bool)
func (_Vdex *VdexCaller) DispatcherWallets(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Vdex.contract.Call(opts, &out, "dispatcherWallets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DispatcherWallets is a free data retrieval call binding the contract method 0x43f32683.
//
// Solidity: function dispatcherWallets(address ) view returns(bool)
func (_Vdex *VdexSession) DispatcherWallets(arg0 common.Address) (bool, error) {
	return _Vdex.Contract.DispatcherWallets(&_Vdex.CallOpts, arg0)
}

// DispatcherWallets is a free data retrieval call binding the contract method 0x43f32683.
//
// Solidity: function dispatcherWallets(address ) view returns(bool)
func (_Vdex *VdexCallerSession) DispatcherWallets(arg0 common.Address) (bool, error) {
	return _Vdex.Contract.DispatcherWallets(&_Vdex.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Vdex *VdexCaller) Filled(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Vdex.contract.Call(opts, &out, "filled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Vdex *VdexSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Vdex.Contract.Filled(&_Vdex.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Vdex *VdexCallerSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Vdex.Contract.Filled(&_Vdex.CallOpts, arg0)
}

// FilledLimitBuyQuoteAmount is a free data retrieval call binding the contract method 0x5b627cd1.
//
// Solidity: function filledLimitBuyQuoteAmount(bytes32 ) view returns(uint256)
func (_Vdex *VdexCaller) FilledLimitBuyQuoteAmount(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Vdex.contract.Call(opts, &out, "filledLimitBuyQuoteAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FilledLimitBuyQuoteAmount is a free data retrieval call binding the contract method 0x5b627cd1.
//
// Solidity: function filledLimitBuyQuoteAmount(bytes32 ) view returns(uint256)
func (_Vdex *VdexSession) FilledLimitBuyQuoteAmount(arg0 [32]byte) (*big.Int, error) {
	return _Vdex.Contract.FilledLimitBuyQuoteAmount(&_Vdex.CallOpts, arg0)
}

// FilledLimitBuyQuoteAmount is a free data retrieval call binding the contract method 0x5b627cd1.
//
// Solidity: function filledLimitBuyQuoteAmount(bytes32 ) view returns(uint256)
func (_Vdex *VdexCallerSession) FilledLimitBuyQuoteAmount(arg0 [32]byte) (*big.Int, error) {
	return _Vdex.Contract.FilledLimitBuyQuoteAmount(&_Vdex.CallOpts, arg0)
}

// GaslessOrderNonces is a free data retrieval call binding the contract method 0xa2f7ba73.
//
// Solidity: function gaslessOrderNonces(address ) view returns(uint256)
func (_Vdex *VdexCaller) GaslessOrderNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vdex.contract.Call(opts, &out, "gaslessOrderNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GaslessOrderNonces is a free data retrieval call binding the contract method 0xa2f7ba73.
//
// Solidity: function gaslessOrderNonces(address ) view returns(uint256)
func (_Vdex *VdexSession) GaslessOrderNonces(arg0 common.Address) (*big.Int, error) {
	return _Vdex.Contract.GaslessOrderNonces(&_Vdex.CallOpts, arg0)
}

// GaslessOrderNonces is a free data retrieval call binding the contract method 0xa2f7ba73.
//
// Solidity: function gaslessOrderNonces(address ) view returns(uint256)
func (_Vdex *VdexCallerSession) GaslessOrderNonces(arg0 common.Address) (*big.Int, error) {
	return _Vdex.Contract.GaslessOrderNonces(&_Vdex.CallOpts, arg0)
}

// LoadFeeWallet is a free data retrieval call binding the contract method 0x02ca6002.
//
// Solidity: function loadFeeWallet() view returns(address)
func (_Vdex *VdexCaller) LoadFeeWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vdex.contract.Call(opts, &out, "loadFeeWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LoadFeeWallet is a free data retrieval call binding the contract method 0x02ca6002.
//
// Solidity: function loadFeeWallet() view returns(address)
func (_Vdex *VdexSession) LoadFeeWallet() (common.Address, error) {
	return _Vdex.Contract.LoadFeeWallet(&_Vdex.CallOpts)
}

// LoadFeeWallet is a free data retrieval call binding the contract method 0x02ca6002.
//
// Solidity: function loadFeeWallet() view returns(address)
func (_Vdex *VdexCallerSession) LoadFeeWallet() (common.Address, error) {
	return _Vdex.Contract.LoadFeeWallet(&_Vdex.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x284f4302.
//
// Solidity: function markets(uint16 ) view returns(address pool, address token0, address token1, uint16 marginLimit, address priceUpdater, uint256 pool0Insurance, uint256 pool1Insurance)
func (_Vdex *VdexCaller) Markets(opts *bind.CallOpts, arg0 uint16) (struct {
	Pool           common.Address
	Token0         common.Address
	Token1         common.Address
	MarginLimit    uint16
	PriceUpdater   common.Address
	Pool0Insurance *big.Int
	Pool1Insurance *big.Int
}, error) {
	var out []interface{}
	err := _Vdex.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		Pool           common.Address
		Token0         common.Address
		Token1         common.Address
		MarginLimit    uint16
		PriceUpdater   common.Address
		Pool0Insurance *big.Int
		Pool1Insurance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pool = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.MarginLimit = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.PriceUpdater = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Pool0Insurance = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Pool1Insurance = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x284f4302.
//
// Solidity: function markets(uint16 ) view returns(address pool, address token0, address token1, uint16 marginLimit, address priceUpdater, uint256 pool0Insurance, uint256 pool1Insurance)
func (_Vdex *VdexSession) Markets(arg0 uint16) (struct {
	Pool           common.Address
	Token0         common.Address
	Token1         common.Address
	MarginLimit    uint16
	PriceUpdater   common.Address
	Pool0Insurance *big.Int
	Pool1Insurance *big.Int
}, error) {
	return _Vdex.Contract.Markets(&_Vdex.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x284f4302.
//
// Solidity: function markets(uint16 ) view returns(address pool, address token0, address token1, uint16 marginLimit, address priceUpdater, uint256 pool0Insurance, uint256 pool1Insurance)
func (_Vdex *VdexCallerSession) Markets(arg0 uint16) (struct {
	Pool           common.Address
	Token0         common.Address
	Token1         common.Address
	MarginLimit    uint16
	PriceUpdater   common.Address
	Pool0Insurance *big.Int
	Pool1Insurance *big.Int
}, error) {
	return _Vdex.Contract.Markets(&_Vdex.CallOpts, arg0)
}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_Vdex *VdexCaller) NonceEquals(opts *bind.CallOpts, makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	var out []interface{}
	err := _Vdex.contract.Call(opts, &out, "nonceEquals", makerAddress, makerNonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_Vdex *VdexSession) NonceEquals(makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	return _Vdex.Contract.NonceEquals(&_Vdex.CallOpts, makerAddress, makerNonce)
}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_Vdex *VdexCallerSession) NonceEquals(makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	return _Vdex.Contract.NonceEquals(&_Vdex.CallOpts, makerAddress, makerNonce)
}

// NumPairs is a free data retrieval call binding the contract method 0xf03e8adc.
//
// Solidity: function numPairs() view returns(uint16)
func (_Vdex *VdexCaller) NumPairs(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Vdex.contract.Call(opts, &out, "numPairs")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// NumPairs is a free data retrieval call binding the contract method 0xf03e8adc.
//
// Solidity: function numPairs() view returns(uint16)
func (_Vdex *VdexSession) NumPairs() (uint16, error) {
	return _Vdex.Contract.NumPairs(&_Vdex.CallOpts)
}

// NumPairs is a free data retrieval call binding the contract method 0xf03e8adc.
//
// Solidity: function numPairs() view returns(uint16)
func (_Vdex *VdexCallerSession) NumPairs() (uint16, error) {
	return _Vdex.Contract.NumPairs(&_Vdex.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vdex *VdexCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vdex.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vdex *VdexSession) ProxiableUUID() ([32]byte, error) {
	return _Vdex.Contract.ProxiableUUID(&_Vdex.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vdex *VdexCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Vdex.Contract.ProxiableUUID(&_Vdex.CallOpts)
}

// TotalHelds is a free data retrieval call binding the contract method 0xec36be82.
//
// Solidity: function totalHelds(address ) view returns(uint256)
func (_Vdex *VdexCaller) TotalHelds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vdex.contract.Call(opts, &out, "totalHelds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalHelds is a free data retrieval call binding the contract method 0xec36be82.
//
// Solidity: function totalHelds(address ) view returns(uint256)
func (_Vdex *VdexSession) TotalHelds(arg0 common.Address) (*big.Int, error) {
	return _Vdex.Contract.TotalHelds(&_Vdex.CallOpts, arg0)
}

// TotalHelds is a free data retrieval call binding the contract method 0xec36be82.
//
// Solidity: function totalHelds(address ) view returns(uint256)
func (_Vdex *VdexCallerSession) TotalHelds(arg0 common.Address) (*big.Int, error) {
	return _Vdex.Contract.TotalHelds(&_Vdex.CallOpts, arg0)
}

// AddDispatcher is a paid mutator transaction binding the contract method 0x75348d35.
//
// Solidity: function addDispatcher(address newDispatcherWallet) returns()
func (_Vdex *VdexTransactor) AddDispatcher(opts *bind.TransactOpts, newDispatcherWallet common.Address) (*types.Transaction, error) {
	return _Vdex.contract.Transact(opts, "addDispatcher", newDispatcherWallet)
}

// AddDispatcher is a paid mutator transaction binding the contract method 0x75348d35.
//
// Solidity: function addDispatcher(address newDispatcherWallet) returns()
func (_Vdex *VdexSession) AddDispatcher(newDispatcherWallet common.Address) (*types.Transaction, error) {
	return _Vdex.Contract.AddDispatcher(&_Vdex.TransactOpts, newDispatcherWallet)
}

// AddDispatcher is a paid mutator transaction binding the contract method 0x75348d35.
//
// Solidity: function addDispatcher(address newDispatcherWallet) returns()
func (_Vdex *VdexTransactorSession) AddDispatcher(newDispatcherWallet common.Address) (*types.Transaction, error) {
	return _Vdex.Contract.AddDispatcher(&_Vdex.TransactOpts, newDispatcherWallet)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0xf0e0d64a.
//
// Solidity: function advanceNonce(uint256 amount) returns()
func (_Vdex *VdexTransactor) AdvanceNonce(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Vdex.contract.Transact(opts, "advanceNonce", amount)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0xf0e0d64a.
//
// Solidity: function advanceNonce(uint256 amount) returns()
func (_Vdex *VdexSession) AdvanceNonce(amount *big.Int) (*types.Transaction, error) {
	return _Vdex.Contract.AdvanceNonce(&_Vdex.TransactOpts, amount)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0xf0e0d64a.
//
// Solidity: function advanceNonce(uint256 amount) returns()
func (_Vdex *VdexTransactorSession) AdvanceNonce(amount *big.Int) (*types.Transaction, error) {
	return _Vdex.Contract.AdvanceNonce(&_Vdex.TransactOpts, amount)
}

// CancelMultiple is a paid mutator transaction binding the contract method 0x6f883299.
//
// Solidity: function cancelMultiple(((address,address,address,uint256,uint256,uint8,uint8,uint256,uint256),bytes,uint256)[] _swapOrders) returns()
func (_Vdex *VdexTransactor) CancelMultiple(opts *bind.TransactOpts, _swapOrders []BaseOrderUtilsCancelOrderParams) (*types.Transaction, error) {
	return _Vdex.contract.Transact(opts, "cancelMultiple", _swapOrders)
}

// CancelMultiple is a paid mutator transaction binding the contract method 0x6f883299.
//
// Solidity: function cancelMultiple(((address,address,address,uint256,uint256,uint8,uint8,uint256,uint256),bytes,uint256)[] _swapOrders) returns()
func (_Vdex *VdexSession) CancelMultiple(_swapOrders []BaseOrderUtilsCancelOrderParams) (*types.Transaction, error) {
	return _Vdex.Contract.CancelMultiple(&_Vdex.TransactOpts, _swapOrders)
}

// CancelMultiple is a paid mutator transaction binding the contract method 0x6f883299.
//
// Solidity: function cancelMultiple(((address,address,address,uint256,uint256,uint8,uint8,uint256,uint256),bytes,uint256)[] _swapOrders) returns()
func (_Vdex *VdexTransactorSession) CancelMultiple(_swapOrders []BaseOrderUtilsCancelOrderParams) (*types.Transaction, error) {
	return _Vdex.Contract.CancelMultiple(&_Vdex.TransactOpts, _swapOrders)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x2aec3988.
//
// Solidity: function createOrder((address[],uint256,uint256,uint256,uint256,uint256,uint8,uint8) params) returns()
func (_Vdex *VdexTransactor) CreateOrder(opts *bind.TransactOpts, params BaseOrderUtilsCreateOrderParams) (*types.Transaction, error) {
	return _Vdex.contract.Transact(opts, "createOrder", params)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x2aec3988.
//
// Solidity: function createOrder((address[],uint256,uint256,uint256,uint256,uint256,uint8,uint8) params) returns()
func (_Vdex *VdexSession) CreateOrder(params BaseOrderUtilsCreateOrderParams) (*types.Transaction, error) {
	return _Vdex.Contract.CreateOrder(&_Vdex.TransactOpts, params)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x2aec3988.
//
// Solidity: function createOrder((address[],uint256,uint256,uint256,uint256,uint256,uint8,uint8) params) returns()
func (_Vdex *VdexTransactorSession) CreateOrder(params BaseOrderUtilsCreateOrderParams) (*types.Transaction, error) {
	return _Vdex.Contract.CreateOrder(&_Vdex.TransactOpts, params)
}

// CreateOrderETH is a paid mutator transaction binding the contract method 0x59de6ce4.
//
// Solidity: function createOrderETH((address[],uint256,uint256,uint256,uint256,uint256,uint8,uint8) params) payable returns()
func (_Vdex *VdexTransactor) CreateOrderETH(opts *bind.TransactOpts, params BaseOrderUtilsCreateOrderParams) (*types.Transaction, error) {
	return _Vdex.contract.Transact(opts, "createOrderETH", params)
}

// CreateOrderETH is a paid mutator transaction binding the contract method 0x59de6ce4.
//
// Solidity: function createOrderETH((address[],uint256,uint256,uint256,uint256,uint256,uint8,uint8) params) payable returns()
func (_Vdex *VdexSession) CreateOrderETH(params BaseOrderUtilsCreateOrderParams) (*types.Transaction, error) {
	return _Vdex.Contract.CreateOrderETH(&_Vdex.TransactOpts, params)
}

// CreateOrderETH is a paid mutator transaction binding the contract method 0x59de6ce4.
//
// Solidity: function createOrderETH((address[],uint256,uint256,uint256,uint256,uint256,uint8,uint8) params) payable returns()
func (_Vdex *VdexTransactorSession) CreateOrderETH(params BaseOrderUtilsCreateOrderParams) (*types.Transaction, error) {
	return _Vdex.Contract.CreateOrderETH(&_Vdex.TransactOpts, params)
}

// CreateOrderWithPermit is a paid mutator transaction binding the contract method 0xbb5e2de6.
//
// Solidity: function createOrderWithPermit((address,address,address,uint256,uint256,uint8,uint8,uint256,uint256) order_, bytes signature, bytes permit) returns(bytes32)
func (_Vdex *VdexTransactor) CreateOrderWithPermit(opts *bind.TransactOpts, order_ Order, signature []byte, permit []byte) (*types.Transaction, error) {
	return _Vdex.contract.Transact(opts, "createOrderWithPermit", order_, signature, permit)
}

// CreateOrderWithPermit is a paid mutator transaction binding the contract method 0xbb5e2de6.
//
// Solidity: function createOrderWithPermit((address,address,address,uint256,uint256,uint8,uint8,uint256,uint256) order_, bytes signature, bytes permit) returns(bytes32)
func (_Vdex *VdexSession) CreateOrderWithPermit(order_ Order, signature []byte, permit []byte) (*types.Transaction, error) {
	return _Vdex.Contract.CreateOrderWithPermit(&_Vdex.TransactOpts, order_, signature, permit)
}

// CreateOrderWithPermit is a paid mutator transaction binding the contract method 0xbb5e2de6.
//
// Solidity: function createOrderWithPermit((address,address,address,uint256,uint256,uint8,uint8,uint256,uint256) order_, bytes signature, bytes permit) returns(bytes32)
func (_Vdex *VdexTransactorSession) CreateOrderWithPermit(order_ Order, signature []byte, permit []byte) (*types.Transaction, error) {
	return _Vdex.Contract.CreateOrderWithPermit(&_Vdex.TransactOpts, order_, signature, permit)
}

// ExecuteOrderBookTrade is a paid mutator transaction binding the contract method 0xbc9623a2.
//
// Solidity: function executeOrderBookTrade((address,uint256,uint256,uint8,uint8,uint256,uint256,bytes) buy, (address,uint256,uint256,uint8,uint8,uint256,uint256,bytes) sell, (address,address,uint256,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8) orderBookTrade) returns()
func (_Vdex *VdexTransactor) ExecuteOrderBookTrade(opts *bind.TransactOpts, buy TypesOrderParam, sell TypesOrderParam, orderBookTrade OrderBookTrade) (*types.Transaction, error) {
	return _Vdex.contract.Transact(opts, "executeOrderBookTrade", buy, sell, orderBookTrade)
}

// ExecuteOrderBookTrade is a paid mutator transaction binding the contract method 0xbc9623a2.
//
// Solidity: function executeOrderBookTrade((address,uint256,uint256,uint8,uint8,uint256,uint256,bytes) buy, (address,uint256,uint256,uint8,uint8,uint256,uint256,bytes) sell, (address,address,uint256,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8) orderBookTrade) returns()
func (_Vdex *VdexSession) ExecuteOrderBookTrade(buy TypesOrderParam, sell TypesOrderParam, orderBookTrade OrderBookTrade) (*types.Transaction, error) {
	return _Vdex.Contract.ExecuteOrderBookTrade(&_Vdex.TransactOpts, buy, sell, orderBookTrade)
}

// ExecuteOrderBookTrade is a paid mutator transaction binding the contract method 0xbc9623a2.
//
// Solidity: function executeOrderBookTrade((address,uint256,uint256,uint8,uint8,uint256,uint256,bytes) buy, (address,uint256,uint256,uint8,uint8,uint256,uint256,bytes) sell, (address,address,uint256,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint8) orderBookTrade) returns()
func (_Vdex *VdexTransactorSession) ExecuteOrderBookTrade(buy TypesOrderParam, sell TypesOrderParam, orderBookTrade OrderBookTrade) (*types.Transaction, error) {
	return _Vdex.Contract.ExecuteOrderBookTrade(&_Vdex.TransactOpts, buy, sell, orderBookTrade)
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_Vdex *VdexTransactor) IncreaseNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vdex.contract.Transact(opts, "increaseNonce")
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_Vdex *VdexSession) IncreaseNonce() (*types.Transaction, error) {
	return _Vdex.Contract.IncreaseNonce(&_Vdex.TransactOpts)
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_Vdex *VdexTransactorSession) IncreaseNonce() (*types.Transaction, error) {
	return _Vdex.Contract.IncreaseNonce(&_Vdex.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _wETH, address _dispatcherWallet, address _callCoinPermit) returns()
func (_Vdex *VdexTransactor) Initialize(opts *bind.TransactOpts, _wETH common.Address, _dispatcherWallet common.Address, _callCoinPermit common.Address) (*types.Transaction, error) {
	return _Vdex.contract.Transact(opts, "initialize", _wETH, _dispatcherWallet, _callCoinPermit)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _wETH, address _dispatcherWallet, address _callCoinPermit) returns()
func (_Vdex *VdexSession) Initialize(_wETH common.Address, _dispatcherWallet common.Address, _callCoinPermit common.Address) (*types.Transaction, error) {
	return _Vdex.Contract.Initialize(&_Vdex.TransactOpts, _wETH, _dispatcherWallet, _callCoinPermit)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _wETH, address _dispatcherWallet, address _callCoinPermit) returns()
func (_Vdex *VdexTransactorSession) Initialize(_wETH common.Address, _dispatcherWallet common.Address, _callCoinPermit common.Address) (*types.Transaction, error) {
	return _Vdex.Contract.Initialize(&_Vdex.TransactOpts, _wETH, _dispatcherWallet, _callCoinPermit)
}

// ReverseDispatcherState is a paid mutator transaction binding the contract method 0x21d43873.
//
// Solidity: function reverseDispatcherState(address dispatcherWallet) returns()
func (_Vdex *VdexTransactor) ReverseDispatcherState(opts *bind.TransactOpts, dispatcherWallet common.Address) (*types.Transaction, error) {
	return _Vdex.contract.Transact(opts, "reverseDispatcherState", dispatcherWallet)
}

// ReverseDispatcherState is a paid mutator transaction binding the contract method 0x21d43873.
//
// Solidity: function reverseDispatcherState(address dispatcherWallet) returns()
func (_Vdex *VdexSession) ReverseDispatcherState(dispatcherWallet common.Address) (*types.Transaction, error) {
	return _Vdex.Contract.ReverseDispatcherState(&_Vdex.TransactOpts, dispatcherWallet)
}

// ReverseDispatcherState is a paid mutator transaction binding the contract method 0x21d43873.
//
// Solidity: function reverseDispatcherState(address dispatcherWallet) returns()
func (_Vdex *VdexTransactorSession) ReverseDispatcherState(dispatcherWallet common.Address) (*types.Transaction, error) {
	return _Vdex.Contract.ReverseDispatcherState(&_Vdex.TransactOpts, dispatcherWallet)
}

// SetFactoryAddress is a paid mutator transaction binding the contract method 0x83c17c55.
//
// Solidity: function setFactoryAddress(address factory) returns()
func (_Vdex *VdexTransactor) SetFactoryAddress(opts *bind.TransactOpts, factory common.Address) (*types.Transaction, error) {
	return _Vdex.contract.Transact(opts, "setFactoryAddress", factory)
}

// SetFactoryAddress is a paid mutator transaction binding the contract method 0x83c17c55.
//
// Solidity: function setFactoryAddress(address factory) returns()
func (_Vdex *VdexSession) SetFactoryAddress(factory common.Address) (*types.Transaction, error) {
	return _Vdex.Contract.SetFactoryAddress(&_Vdex.TransactOpts, factory)
}

// SetFactoryAddress is a paid mutator transaction binding the contract method 0x83c17c55.
//
// Solidity: function setFactoryAddress(address factory) returns()
func (_Vdex *VdexTransactorSession) SetFactoryAddress(factory common.Address) (*types.Transaction, error) {
	return _Vdex.Contract.SetFactoryAddress(&_Vdex.TransactOpts, factory)
}

// SetFeeWallet is a paid mutator transaction binding the contract method 0x90d49b9d.
//
// Solidity: function setFeeWallet(address newFeeWallet) returns()
func (_Vdex *VdexTransactor) SetFeeWallet(opts *bind.TransactOpts, newFeeWallet common.Address) (*types.Transaction, error) {
	return _Vdex.contract.Transact(opts, "setFeeWallet", newFeeWallet)
}

// SetFeeWallet is a paid mutator transaction binding the contract method 0x90d49b9d.
//
// Solidity: function setFeeWallet(address newFeeWallet) returns()
func (_Vdex *VdexSession) SetFeeWallet(newFeeWallet common.Address) (*types.Transaction, error) {
	return _Vdex.Contract.SetFeeWallet(&_Vdex.TransactOpts, newFeeWallet)
}

// SetFeeWallet is a paid mutator transaction binding the contract method 0x90d49b9d.
//
// Solidity: function setFeeWallet(address newFeeWallet) returns()
func (_Vdex *VdexTransactorSession) SetFeeWallet(newFeeWallet common.Address) (*types.Transaction, error) {
	return _Vdex.Contract.SetFeeWallet(&_Vdex.TransactOpts, newFeeWallet)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vdex *VdexTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Vdex.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vdex *VdexSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Vdex.Contract.UpgradeTo(&_Vdex.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vdex *VdexTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Vdex.Contract.UpgradeTo(&_Vdex.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vdex *VdexTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vdex.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vdex *VdexSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vdex.Contract.UpgradeToAndCall(&_Vdex.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vdex *VdexTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vdex.Contract.UpgradeToAndCall(&_Vdex.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vdex *VdexTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vdex.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vdex *VdexSession) Receive() (*types.Transaction, error) {
	return _Vdex.Contract.Receive(&_Vdex.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vdex *VdexTransactorSession) Receive() (*types.Transaction, error) {
	return _Vdex.Contract.Receive(&_Vdex.TransactOpts)
}

// VdexAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Vdex contract.
type VdexAdminChangedIterator struct {
	Event *VdexAdminChanged // Event containing the contract specifics and raw log

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
func (it *VdexAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexAdminChanged)
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
		it.Event = new(VdexAdminChanged)
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
func (it *VdexAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexAdminChanged represents a AdminChanged event raised by the Vdex contract.
type VdexAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Vdex *VdexFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VdexAdminChangedIterator, error) {

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VdexAdminChangedIterator{contract: _Vdex.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Vdex *VdexFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VdexAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexAdminChanged)
				if err := _Vdex.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Vdex *VdexFilterer) ParseAdminChanged(log types.Log) (*VdexAdminChanged, error) {
	event := new(VdexAdminChanged)
	if err := _Vdex.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Vdex contract.
type VdexBeaconUpgradedIterator struct {
	Event *VdexBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *VdexBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexBeaconUpgraded)
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
		it.Event = new(VdexBeaconUpgraded)
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
func (it *VdexBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexBeaconUpgraded represents a BeaconUpgraded event raised by the Vdex contract.
type VdexBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Vdex *VdexFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*VdexBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &VdexBeaconUpgradedIterator{contract: _Vdex.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Vdex *VdexFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *VdexBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexBeaconUpgraded)
				if err := _Vdex.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Vdex *VdexFilterer) ParseBeaconUpgraded(log types.Log) (*VdexBeaconUpgraded, error) {
	event := new(VdexBeaconUpgraded)
	if err := _Vdex.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexCancelSwapOrderIterator is returned from FilterCancelSwapOrder and is used to iterate over the raw logs and unpacked data for CancelSwapOrder events raised by the Vdex contract.
type VdexCancelSwapOrderIterator struct {
	Event *VdexCancelSwapOrder // Event containing the contract specifics and raw log

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
func (it *VdexCancelSwapOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexCancelSwapOrder)
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
		it.Event = new(VdexCancelSwapOrder)
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
func (it *VdexCancelSwapOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexCancelSwapOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexCancelSwapOrder represents a CancelSwapOrder event raised by the Vdex contract.
type VdexCancelSwapOrder struct {
	OrderHash        [32]byte
	Account          common.Address
	BaseAsset        common.Address
	QuoteAsset       common.Address
	BaseAssetAmount  *big.Int
	QuoteAssetAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCancelSwapOrder is a free log retrieval operation binding the contract event 0x989f18930b2038a1e3e175b58f88e8b67aeed561f2df3832c84c696723f9e20b.
//
// Solidity: event CancelSwapOrder(bytes32 orderHash, address indexed account, address baseAsset, address quoteAsset, uint256 baseAssetAmount, uint256 quoteAssetAmount)
func (_Vdex *VdexFilterer) FilterCancelSwapOrder(opts *bind.FilterOpts, account []common.Address) (*VdexCancelSwapOrderIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "CancelSwapOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return &VdexCancelSwapOrderIterator{contract: _Vdex.contract, event: "CancelSwapOrder", logs: logs, sub: sub}, nil
}

// WatchCancelSwapOrder is a free log subscription operation binding the contract event 0x989f18930b2038a1e3e175b58f88e8b67aeed561f2df3832c84c696723f9e20b.
//
// Solidity: event CancelSwapOrder(bytes32 orderHash, address indexed account, address baseAsset, address quoteAsset, uint256 baseAssetAmount, uint256 quoteAssetAmount)
func (_Vdex *VdexFilterer) WatchCancelSwapOrder(opts *bind.WatchOpts, sink chan<- *VdexCancelSwapOrder, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "CancelSwapOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexCancelSwapOrder)
				if err := _Vdex.contract.UnpackLog(event, "CancelSwapOrder", log); err != nil {
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

// ParseCancelSwapOrder is a log parse operation binding the contract event 0x989f18930b2038a1e3e175b58f88e8b67aeed561f2df3832c84c696723f9e20b.
//
// Solidity: event CancelSwapOrder(bytes32 orderHash, address indexed account, address baseAsset, address quoteAsset, uint256 baseAssetAmount, uint256 quoteAssetAmount)
func (_Vdex *VdexFilterer) ParseCancelSwapOrder(log types.Log) (*VdexCancelSwapOrder, error) {
	event := new(VdexCancelSwapOrder)
	if err := _Vdex.contract.UnpackLog(event, "CancelSwapOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexCreateDecreasePositionIterator is returned from FilterCreateDecreasePosition and is used to iterate over the raw logs and unpacked data for CreateDecreasePosition events raised by the Vdex contract.
type VdexCreateDecreasePositionIterator struct {
	Event *VdexCreateDecreasePosition // Event containing the contract specifics and raw log

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
func (it *VdexCreateDecreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexCreateDecreasePosition)
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
		it.Event = new(VdexCreateDecreasePosition)
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
func (it *VdexCreateDecreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexCreateDecreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexCreateDecreasePosition represents a CreateDecreasePosition event raised by the Vdex contract.
type VdexCreateDecreasePosition struct {
	Account         common.Address
	DepositToken    common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Receiver        common.Address
	AcceptablePrice *big.Int
	MinOut          *big.Int
	ExecutionFee    *big.Int
	Index           *big.Int
	QueueIndex      *big.Int
	BlockNumber     *big.Int
	BlockTime       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateDecreasePosition is a free log retrieval operation binding the contract event 0x5cec8d41a6544687bc446a7cce13f2b1af6ff7339b6c3bf19bb5e6eac9174c1e.
//
// Solidity: event CreateDecreasePosition(address indexed account, address depositToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 index, uint256 queueIndex, uint256 blockNumber, uint256 blockTime)
func (_Vdex *VdexFilterer) FilterCreateDecreasePosition(opts *bind.FilterOpts, account []common.Address) (*VdexCreateDecreasePositionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "CreateDecreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return &VdexCreateDecreasePositionIterator{contract: _Vdex.contract, event: "CreateDecreasePosition", logs: logs, sub: sub}, nil
}

// WatchCreateDecreasePosition is a free log subscription operation binding the contract event 0x5cec8d41a6544687bc446a7cce13f2b1af6ff7339b6c3bf19bb5e6eac9174c1e.
//
// Solidity: event CreateDecreasePosition(address indexed account, address depositToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 index, uint256 queueIndex, uint256 blockNumber, uint256 blockTime)
func (_Vdex *VdexFilterer) WatchCreateDecreasePosition(opts *bind.WatchOpts, sink chan<- *VdexCreateDecreasePosition, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "CreateDecreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexCreateDecreasePosition)
				if err := _Vdex.contract.UnpackLog(event, "CreateDecreasePosition", log); err != nil {
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

// ParseCreateDecreasePosition is a log parse operation binding the contract event 0x5cec8d41a6544687bc446a7cce13f2b1af6ff7339b6c3bf19bb5e6eac9174c1e.
//
// Solidity: event CreateDecreasePosition(address indexed account, address depositToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 index, uint256 queueIndex, uint256 blockNumber, uint256 blockTime)
func (_Vdex *VdexFilterer) ParseCreateDecreasePosition(log types.Log) (*VdexCreateDecreasePosition, error) {
	event := new(VdexCreateDecreasePosition)
	if err := _Vdex.contract.UnpackLog(event, "CreateDecreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexCreateIncreasePositionIterator is returned from FilterCreateIncreasePosition and is used to iterate over the raw logs and unpacked data for CreateIncreasePosition events raised by the Vdex contract.
type VdexCreateIncreasePositionIterator struct {
	Event *VdexCreateIncreasePosition // Event containing the contract specifics and raw log

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
func (it *VdexCreateIncreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexCreateIncreasePosition)
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
		it.Event = new(VdexCreateIncreasePosition)
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
func (it *VdexCreateIncreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexCreateIncreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexCreateIncreasePosition represents a CreateIncreasePosition event raised by the Vdex contract.
type VdexCreateIncreasePosition struct {
	Account         common.Address
	DepositToken    common.Address
	IndexToken      common.Address
	AmountIn        *big.Int
	MinOut          *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	AcceptablePrice *big.Int
	ExecutionFee    *big.Int
	Index           *big.Int
	QueueIndex      *big.Int
	BlockNumber     *big.Int
	BlockTime       *big.Int
	GasPrice        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateIncreasePosition is a free log retrieval operation binding the contract event 0xa91a05eaddf7d709821b201999ac6f971c896481c96a00421dd28fe2b929d29f.
//
// Solidity: event CreateIncreasePosition(address indexed account, address depositToken, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 index, uint256 queueIndex, uint256 blockNumber, uint256 blockTime, uint256 gasPrice)
func (_Vdex *VdexFilterer) FilterCreateIncreasePosition(opts *bind.FilterOpts, account []common.Address) (*VdexCreateIncreasePositionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "CreateIncreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return &VdexCreateIncreasePositionIterator{contract: _Vdex.contract, event: "CreateIncreasePosition", logs: logs, sub: sub}, nil
}

// WatchCreateIncreasePosition is a free log subscription operation binding the contract event 0xa91a05eaddf7d709821b201999ac6f971c896481c96a00421dd28fe2b929d29f.
//
// Solidity: event CreateIncreasePosition(address indexed account, address depositToken, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 index, uint256 queueIndex, uint256 blockNumber, uint256 blockTime, uint256 gasPrice)
func (_Vdex *VdexFilterer) WatchCreateIncreasePosition(opts *bind.WatchOpts, sink chan<- *VdexCreateIncreasePosition, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "CreateIncreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexCreateIncreasePosition)
				if err := _Vdex.contract.UnpackLog(event, "CreateIncreasePosition", log); err != nil {
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

// ParseCreateIncreasePosition is a log parse operation binding the contract event 0xa91a05eaddf7d709821b201999ac6f971c896481c96a00421dd28fe2b929d29f.
//
// Solidity: event CreateIncreasePosition(address indexed account, address depositToken, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 index, uint256 queueIndex, uint256 blockNumber, uint256 blockTime, uint256 gasPrice)
func (_Vdex *VdexFilterer) ParseCreateIncreasePosition(log types.Log) (*VdexCreateIncreasePosition, error) {
	event := new(VdexCreateIncreasePosition)
	if err := _Vdex.contract.UnpackLog(event, "CreateIncreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Vdex contract.
type VdexDepositedIterator struct {
	Event *VdexDeposited // Event containing the contract specifics and raw log

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
func (it *VdexDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexDeposited)
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
		it.Event = new(VdexDeposited)
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
func (it *VdexDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexDeposited represents a Deposited event raised by the Vdex contract.
type VdexDeposited struct {
	Wallet                         common.Address
	AssetAddress                   common.Address
	QuantityInPips                 uint64
	NewExchangeBalanceInPips       uint64
	NewExchangeBalanceInAssetUnits *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x6350fec42d5eeb4ce59964286bef86ead82f4d029d1d4bfe13208f819b5109f5.
//
// Solidity: event Deposited(address wallet, address assetAddress, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_Vdex *VdexFilterer) FilterDeposited(opts *bind.FilterOpts) (*VdexDepositedIterator, error) {

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &VdexDepositedIterator{contract: _Vdex.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x6350fec42d5eeb4ce59964286bef86ead82f4d029d1d4bfe13208f819b5109f5.
//
// Solidity: event Deposited(address wallet, address assetAddress, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_Vdex *VdexFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *VdexDeposited) (event.Subscription, error) {

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexDeposited)
				if err := _Vdex.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x6350fec42d5eeb4ce59964286bef86ead82f4d029d1d4bfe13208f819b5109f5.
//
// Solidity: event Deposited(address wallet, address assetAddress, uint64 quantityInPips, uint64 newExchangeBalanceInPips, uint256 newExchangeBalanceInAssetUnits)
func (_Vdex *VdexFilterer) ParseDeposited(log types.Log) (*VdexDeposited, error) {
	event := new(VdexDeposited)
	if err := _Vdex.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexExecuteDecreasePositionIterator is returned from FilterExecuteDecreasePosition and is used to iterate over the raw logs and unpacked data for ExecuteDecreasePosition events raised by the Vdex contract.
type VdexExecuteDecreasePositionIterator struct {
	Event *VdexExecuteDecreasePosition // Event containing the contract specifics and raw log

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
func (it *VdexExecuteDecreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexExecuteDecreasePosition)
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
		it.Event = new(VdexExecuteDecreasePosition)
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
func (it *VdexExecuteDecreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexExecuteDecreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexExecuteDecreasePosition represents a ExecuteDecreasePosition event raised by the Vdex contract.
type VdexExecuteDecreasePosition struct {
	Account         common.Address
	Path            []common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Receiver        common.Address
	AcceptablePrice *big.Int
	MinOut          *big.Int
	ExecutionFee    *big.Int
	BlockGap        *big.Int
	TimeGap         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExecuteDecreasePosition is a free log retrieval operation binding the contract event 0x21435c5b618d77ff3657140cd3318e2cffaebc5e0e1b7318f56a9ba4044c3ed2.
//
// Solidity: event ExecuteDecreasePosition(address indexed account, address[] path, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_Vdex *VdexFilterer) FilterExecuteDecreasePosition(opts *bind.FilterOpts, account []common.Address) (*VdexExecuteDecreasePositionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "ExecuteDecreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return &VdexExecuteDecreasePositionIterator{contract: _Vdex.contract, event: "ExecuteDecreasePosition", logs: logs, sub: sub}, nil
}

// WatchExecuteDecreasePosition is a free log subscription operation binding the contract event 0x21435c5b618d77ff3657140cd3318e2cffaebc5e0e1b7318f56a9ba4044c3ed2.
//
// Solidity: event ExecuteDecreasePosition(address indexed account, address[] path, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_Vdex *VdexFilterer) WatchExecuteDecreasePosition(opts *bind.WatchOpts, sink chan<- *VdexExecuteDecreasePosition, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "ExecuteDecreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexExecuteDecreasePosition)
				if err := _Vdex.contract.UnpackLog(event, "ExecuteDecreasePosition", log); err != nil {
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

// ParseExecuteDecreasePosition is a log parse operation binding the contract event 0x21435c5b618d77ff3657140cd3318e2cffaebc5e0e1b7318f56a9ba4044c3ed2.
//
// Solidity: event ExecuteDecreasePosition(address indexed account, address[] path, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_Vdex *VdexFilterer) ParseExecuteDecreasePosition(log types.Log) (*VdexExecuteDecreasePosition, error) {
	event := new(VdexExecuteDecreasePosition)
	if err := _Vdex.contract.UnpackLog(event, "ExecuteDecreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexExecuteIncreasePositionIterator is returned from FilterExecuteIncreasePosition and is used to iterate over the raw logs and unpacked data for ExecuteIncreasePosition events raised by the Vdex contract.
type VdexExecuteIncreasePositionIterator struct {
	Event *VdexExecuteIncreasePosition // Event containing the contract specifics and raw log

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
func (it *VdexExecuteIncreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexExecuteIncreasePosition)
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
		it.Event = new(VdexExecuteIncreasePosition)
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
func (it *VdexExecuteIncreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexExecuteIncreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexExecuteIncreasePosition represents a ExecuteIncreasePosition event raised by the Vdex contract.
type VdexExecuteIncreasePosition struct {
	Account         common.Address
	Path            []common.Address
	IndexToken      common.Address
	AmountIn        *big.Int
	MinOut          *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	AcceptablePrice *big.Int
	ExecutionFee    *big.Int
	BlockGap        *big.Int
	TimeGap         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExecuteIncreasePosition is a free log retrieval operation binding the contract event 0x1be316b94d38c07bd41cdb4913772d0a0a82802786a2f8b657b6e85dbcdfc641.
//
// Solidity: event ExecuteIncreasePosition(address indexed account, address[] path, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_Vdex *VdexFilterer) FilterExecuteIncreasePosition(opts *bind.FilterOpts, account []common.Address) (*VdexExecuteIncreasePositionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "ExecuteIncreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return &VdexExecuteIncreasePositionIterator{contract: _Vdex.contract, event: "ExecuteIncreasePosition", logs: logs, sub: sub}, nil
}

// WatchExecuteIncreasePosition is a free log subscription operation binding the contract event 0x1be316b94d38c07bd41cdb4913772d0a0a82802786a2f8b657b6e85dbcdfc641.
//
// Solidity: event ExecuteIncreasePosition(address indexed account, address[] path, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_Vdex *VdexFilterer) WatchExecuteIncreasePosition(opts *bind.WatchOpts, sink chan<- *VdexExecuteIncreasePosition, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "ExecuteIncreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexExecuteIncreasePosition)
				if err := _Vdex.contract.UnpackLog(event, "ExecuteIncreasePosition", log); err != nil {
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

// ParseExecuteIncreasePosition is a log parse operation binding the contract event 0x1be316b94d38c07bd41cdb4913772d0a0a82802786a2f8b657b6e85dbcdfc641.
//
// Solidity: event ExecuteIncreasePosition(address indexed account, address[] path, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_Vdex *VdexFilterer) ParseExecuteIncreasePosition(log types.Log) (*VdexExecuteIncreasePosition, error) {
	event := new(VdexExecuteIncreasePosition)
	if err := _Vdex.contract.UnpackLog(event, "ExecuteIncreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Vdex contract.
type VdexInitializedIterator struct {
	Event *VdexInitialized // Event containing the contract specifics and raw log

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
func (it *VdexInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexInitialized)
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
		it.Event = new(VdexInitialized)
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
func (it *VdexInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexInitialized represents a Initialized event raised by the Vdex contract.
type VdexInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Vdex *VdexFilterer) FilterInitialized(opts *bind.FilterOpts) (*VdexInitializedIterator, error) {

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VdexInitializedIterator{contract: _Vdex.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Vdex *VdexFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VdexInitialized) (event.Subscription, error) {

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexInitialized)
				if err := _Vdex.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Vdex *VdexFilterer) ParseInitialized(log types.Log) (*VdexInitialized, error) {
	event := new(VdexInitialized)
	if err := _Vdex.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexNewFactoryAddressIterator is returned from FilterNewFactoryAddress and is used to iterate over the raw logs and unpacked data for NewFactoryAddress events raised by the Vdex contract.
type VdexNewFactoryAddressIterator struct {
	Event *VdexNewFactoryAddress // Event containing the contract specifics and raw log

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
func (it *VdexNewFactoryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexNewFactoryAddress)
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
		it.Event = new(VdexNewFactoryAddress)
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
func (it *VdexNewFactoryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexNewFactoryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexNewFactoryAddress represents a NewFactoryAddress event raised by the Vdex contract.
type VdexNewFactoryAddress struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewFactoryAddress is a free log retrieval operation binding the contract event 0xd75366e5d64c01e3af9734d8317e7c2c1a42cf0474a7568e68f3c1d0aa3a67ba.
//
// Solidity: event NewFactoryAddress(address factory)
func (_Vdex *VdexFilterer) FilterNewFactoryAddress(opts *bind.FilterOpts) (*VdexNewFactoryAddressIterator, error) {

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "NewFactoryAddress")
	if err != nil {
		return nil, err
	}
	return &VdexNewFactoryAddressIterator{contract: _Vdex.contract, event: "NewFactoryAddress", logs: logs, sub: sub}, nil
}

// WatchNewFactoryAddress is a free log subscription operation binding the contract event 0xd75366e5d64c01e3af9734d8317e7c2c1a42cf0474a7568e68f3c1d0aa3a67ba.
//
// Solidity: event NewFactoryAddress(address factory)
func (_Vdex *VdexFilterer) WatchNewFactoryAddress(opts *bind.WatchOpts, sink chan<- *VdexNewFactoryAddress) (event.Subscription, error) {

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "NewFactoryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexNewFactoryAddress)
				if err := _Vdex.contract.UnpackLog(event, "NewFactoryAddress", log); err != nil {
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

// ParseNewFactoryAddress is a log parse operation binding the contract event 0xd75366e5d64c01e3af9734d8317e7c2c1a42cf0474a7568e68f3c1d0aa3a67ba.
//
// Solidity: event NewFactoryAddress(address factory)
func (_Vdex *VdexFilterer) ParseNewFactoryAddress(log types.Log) (*VdexNewFactoryAddress, error) {
	event := new(VdexNewFactoryAddress)
	if err := _Vdex.contract.UnpackLog(event, "NewFactoryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexNonceIncreasedIterator is returned from FilterNonceIncreased and is used to iterate over the raw logs and unpacked data for NonceIncreased events raised by the Vdex contract.
type VdexNonceIncreasedIterator struct {
	Event *VdexNonceIncreased // Event containing the contract specifics and raw log

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
func (it *VdexNonceIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexNonceIncreased)
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
		it.Event = new(VdexNonceIncreased)
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
func (it *VdexNonceIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexNonceIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexNonceIncreased represents a NonceIncreased event raised by the Vdex contract.
type VdexNonceIncreased struct {
	Maker    common.Address
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncreased is a free log retrieval operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_Vdex *VdexFilterer) FilterNonceIncreased(opts *bind.FilterOpts, maker []common.Address) (*VdexNonceIncreasedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "NonceIncreased", makerRule)
	if err != nil {
		return nil, err
	}
	return &VdexNonceIncreasedIterator{contract: _Vdex.contract, event: "NonceIncreased", logs: logs, sub: sub}, nil
}

// WatchNonceIncreased is a free log subscription operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_Vdex *VdexFilterer) WatchNonceIncreased(opts *bind.WatchOpts, sink chan<- *VdexNonceIncreased, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "NonceIncreased", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexNonceIncreased)
				if err := _Vdex.contract.UnpackLog(event, "NonceIncreased", log); err != nil {
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

// ParseNonceIncreased is a log parse operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_Vdex *VdexFilterer) ParseNonceIncreased(log types.Log) (*VdexNonceIncreased, error) {
	event := new(VdexNonceIncreased)
	if err := _Vdex.contract.UnpackLog(event, "NonceIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexOrderBookTradeExecutedIterator is returned from FilterOrderBookTradeExecuted and is used to iterate over the raw logs and unpacked data for OrderBookTradeExecuted events raised by the Vdex contract.
type VdexOrderBookTradeExecutedIterator struct {
	Event *VdexOrderBookTradeExecuted // Event containing the contract specifics and raw log

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
func (it *VdexOrderBookTradeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexOrderBookTradeExecuted)
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
		it.Event = new(VdexOrderBookTradeExecuted)
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
func (it *VdexOrderBookTradeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexOrderBookTradeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexOrderBookTradeExecuted represents a OrderBookTradeExecuted event raised by the Vdex contract.
type VdexOrderBookTradeExecuted struct {
	BuyWallet          common.Address
	SellWallet         common.Address
	BuyerInputUpdate   *big.Int
	SellerInputUpdate  *big.Int
	BuyerOutputUpdate  *big.Int
	SellerOutputUpdate *big.Int
	TakerSide          uint8
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOrderBookTradeExecuted is a free log retrieval operation binding the contract event 0x08cba63b0336877d8d806f5a9f58ca558fc6831babc138965dbafe5f854ece06.
//
// Solidity: event OrderBookTradeExecuted(address buyWallet, address sellWallet, uint256 buyerInputUpdate, uint256 sellerInputUpdate, uint256 buyerOutputUpdate, uint256 sellerOutputUpdate, uint8 takerSide)
func (_Vdex *VdexFilterer) FilterOrderBookTradeExecuted(opts *bind.FilterOpts) (*VdexOrderBookTradeExecutedIterator, error) {

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "OrderBookTradeExecuted")
	if err != nil {
		return nil, err
	}
	return &VdexOrderBookTradeExecutedIterator{contract: _Vdex.contract, event: "OrderBookTradeExecuted", logs: logs, sub: sub}, nil
}

// WatchOrderBookTradeExecuted is a free log subscription operation binding the contract event 0x08cba63b0336877d8d806f5a9f58ca558fc6831babc138965dbafe5f854ece06.
//
// Solidity: event OrderBookTradeExecuted(address buyWallet, address sellWallet, uint256 buyerInputUpdate, uint256 sellerInputUpdate, uint256 buyerOutputUpdate, uint256 sellerOutputUpdate, uint8 takerSide)
func (_Vdex *VdexFilterer) WatchOrderBookTradeExecuted(opts *bind.WatchOpts, sink chan<- *VdexOrderBookTradeExecuted) (event.Subscription, error) {

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "OrderBookTradeExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexOrderBookTradeExecuted)
				if err := _Vdex.contract.UnpackLog(event, "OrderBookTradeExecuted", log); err != nil {
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

// ParseOrderBookTradeExecuted is a log parse operation binding the contract event 0x08cba63b0336877d8d806f5a9f58ca558fc6831babc138965dbafe5f854ece06.
//
// Solidity: event OrderBookTradeExecuted(address buyWallet, address sellWallet, uint256 buyerInputUpdate, uint256 sellerInputUpdate, uint256 buyerOutputUpdate, uint256 sellerOutputUpdate, uint8 takerSide)
func (_Vdex *VdexFilterer) ParseOrderBookTradeExecuted(log types.Log) (*VdexOrderBookTradeExecuted, error) {
	event := new(VdexOrderBookTradeExecuted)
	if err := _Vdex.contract.UnpackLog(event, "OrderBookTradeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the Vdex contract.
type VdexOrderCreatedIterator struct {
	Event *VdexOrderCreated // Event containing the contract specifics and raw log

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
func (it *VdexOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexOrderCreated)
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
		it.Event = new(VdexOrderCreated)
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
func (it *VdexOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexOrderCreated represents a OrderCreated event raised by the Vdex contract.
type VdexOrderCreated struct {
	OrderHash         [32]byte
	MakerAccountOwner common.Address
	BaseAsset         common.Address
	QuoteAsset        common.Address
	OrderType         *big.Int
	BaseAssetAmount   *big.Int
	QuoteAssetAmount  *big.Int
	OrderSide         *big.Int
	Expiration        *big.Int
	Nonce             *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x1f7b2798b5acb13b767ac02db1a6178f42af1afe4228c23518bc5fb58743366b.
//
// Solidity: event OrderCreated(bytes32 indexed orderHash, address makerAccountOwner, address baseAsset, address quoteAsset, uint256 orderType, uint256 baseAssetAmount, uint256 quoteAssetAmount, uint256 orderSide, uint256 expiration, uint256 nonce)
func (_Vdex *VdexFilterer) FilterOrderCreated(opts *bind.FilterOpts, orderHash [][32]byte) (*VdexOrderCreatedIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "OrderCreated", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &VdexOrderCreatedIterator{contract: _Vdex.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x1f7b2798b5acb13b767ac02db1a6178f42af1afe4228c23518bc5fb58743366b.
//
// Solidity: event OrderCreated(bytes32 indexed orderHash, address makerAccountOwner, address baseAsset, address quoteAsset, uint256 orderType, uint256 baseAssetAmount, uint256 quoteAssetAmount, uint256 orderSide, uint256 expiration, uint256 nonce)
func (_Vdex *VdexFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *VdexOrderCreated, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "OrderCreated", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexOrderCreated)
				if err := _Vdex.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0x1f7b2798b5acb13b767ac02db1a6178f42af1afe4228c23518bc5fb58743366b.
//
// Solidity: event OrderCreated(bytes32 indexed orderHash, address makerAccountOwner, address baseAsset, address quoteAsset, uint256 orderType, uint256 baseAssetAmount, uint256 quoteAssetAmount, uint256 orderSide, uint256 expiration, uint256 nonce)
func (_Vdex *VdexFilterer) ParseOrderCreated(log types.Log) (*VdexOrderCreated, error) {
	event := new(VdexOrderCreated)
	if err := _Vdex.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Vdex contract.
type VdexUpgradedIterator struct {
	Event *VdexUpgraded // Event containing the contract specifics and raw log

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
func (it *VdexUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexUpgraded)
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
		it.Event = new(VdexUpgraded)
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
func (it *VdexUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexUpgraded represents a Upgraded event raised by the Vdex contract.
type VdexUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vdex *VdexFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VdexUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VdexUpgradedIterator{contract: _Vdex.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vdex *VdexFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VdexUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexUpgraded)
				if err := _Vdex.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vdex *VdexFilterer) ParseUpgraded(log types.Log) (*VdexUpgraded, error) {
	event := new(VdexUpgraded)
	if err := _Vdex.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexWalletExitClearedIterator is returned from FilterWalletExitCleared and is used to iterate over the raw logs and unpacked data for WalletExitCleared events raised by the Vdex contract.
type VdexWalletExitClearedIterator struct {
	Event *VdexWalletExitCleared // Event containing the contract specifics and raw log

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
func (it *VdexWalletExitClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexWalletExitCleared)
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
		it.Event = new(VdexWalletExitCleared)
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
func (it *VdexWalletExitClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexWalletExitClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexWalletExitCleared represents a WalletExitCleared event raised by the Vdex contract.
type VdexWalletExitCleared struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWalletExitCleared is a free log retrieval operation binding the contract event 0xb771d4b2a83beca38f442c8903629e0e8ab1a07cf76e94eb2977153167e20936.
//
// Solidity: event WalletExitCleared(address wallet)
func (_Vdex *VdexFilterer) FilterWalletExitCleared(opts *bind.FilterOpts) (*VdexWalletExitClearedIterator, error) {

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "WalletExitCleared")
	if err != nil {
		return nil, err
	}
	return &VdexWalletExitClearedIterator{contract: _Vdex.contract, event: "WalletExitCleared", logs: logs, sub: sub}, nil
}

// WatchWalletExitCleared is a free log subscription operation binding the contract event 0xb771d4b2a83beca38f442c8903629e0e8ab1a07cf76e94eb2977153167e20936.
//
// Solidity: event WalletExitCleared(address wallet)
func (_Vdex *VdexFilterer) WatchWalletExitCleared(opts *bind.WatchOpts, sink chan<- *VdexWalletExitCleared) (event.Subscription, error) {

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "WalletExitCleared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexWalletExitCleared)
				if err := _Vdex.contract.UnpackLog(event, "WalletExitCleared", log); err != nil {
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

// ParseWalletExitCleared is a log parse operation binding the contract event 0xb771d4b2a83beca38f442c8903629e0e8ab1a07cf76e94eb2977153167e20936.
//
// Solidity: event WalletExitCleared(address wallet)
func (_Vdex *VdexFilterer) ParseWalletExitCleared(log types.Log) (*VdexWalletExitCleared, error) {
	event := new(VdexWalletExitCleared)
	if err := _Vdex.contract.UnpackLog(event, "WalletExitCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VdexWalletExitedIterator is returned from FilterWalletExited and is used to iterate over the raw logs and unpacked data for WalletExited events raised by the Vdex contract.
type VdexWalletExitedIterator struct {
	Event *VdexWalletExited // Event containing the contract specifics and raw log

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
func (it *VdexWalletExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VdexWalletExited)
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
		it.Event = new(VdexWalletExited)
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
func (it *VdexWalletExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VdexWalletExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VdexWalletExited represents a WalletExited event raised by the Vdex contract.
type VdexWalletExited struct {
	Wallet               common.Address
	EffectiveBlockNumber *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterWalletExited is a free log retrieval operation binding the contract event 0xd60f9f7b2f1a208268475a927bd727c4e198fc8b40aab3004ebcc2bc78ca8480.
//
// Solidity: event WalletExited(address wallet, uint256 effectiveBlockNumber)
func (_Vdex *VdexFilterer) FilterWalletExited(opts *bind.FilterOpts) (*VdexWalletExitedIterator, error) {

	logs, sub, err := _Vdex.contract.FilterLogs(opts, "WalletExited")
	if err != nil {
		return nil, err
	}
	return &VdexWalletExitedIterator{contract: _Vdex.contract, event: "WalletExited", logs: logs, sub: sub}, nil
}

// WatchWalletExited is a free log subscription operation binding the contract event 0xd60f9f7b2f1a208268475a927bd727c4e198fc8b40aab3004ebcc2bc78ca8480.
//
// Solidity: event WalletExited(address wallet, uint256 effectiveBlockNumber)
func (_Vdex *VdexFilterer) WatchWalletExited(opts *bind.WatchOpts, sink chan<- *VdexWalletExited) (event.Subscription, error) {

	logs, sub, err := _Vdex.contract.WatchLogs(opts, "WalletExited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VdexWalletExited)
				if err := _Vdex.contract.UnpackLog(event, "WalletExited", log); err != nil {
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

// ParseWalletExited is a log parse operation binding the contract event 0xd60f9f7b2f1a208268475a927bd727c4e198fc8b40aab3004ebcc2bc78ca8480.
//
// Solidity: event WalletExited(address wallet, uint256 effectiveBlockNumber)
func (_Vdex *VdexFilterer) ParseWalletExited(log types.Log) (*VdexWalletExited, error) {
	event := new(VdexWalletExited)
	if err := _Vdex.contract.UnpackLog(event, "WalletExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
