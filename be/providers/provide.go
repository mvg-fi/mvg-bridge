package providers

import (
	"fmt"
	"slices"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/providers/mixpay"
	"github.com/mvg-fi/mvg-bridge/providers/pandoswap"
	"github.com/shopspring/decimal"
)

var (
	PROVIDERS = []string{mixpay.NAME, pandoswap.NAME}
	CEX       = []string{mixpay.NAME}
	DEX       = []string{pandoswap.NAME}
	DISABLED  = []string{}
)

// bestPrice is the amount of asset will be received
// bestFee is the amount of original asset will be charged
func GetPriceSimple(payAsset, receiveAsset, amount, except string, cex bool) *constants.PriceSimpleResp {
	var mixpayChan, mixpayFeeChan chan float64
	var mixpayPrice, mixpayFee float64
	var pandoswapChan, pandoswapfeeChan chan float64
	var pandoswapPrice, pandoswapFee float64
	chainAsset, fee := AssetFee(receiveAsset)
	if cex {
		mixpayChan = make(chan float64)
		mixpayFeeChan = make(chan float64)
		go mixpay.GetPrice(payAsset, chainAsset, "", fee, mixpayFeeChan)
		go mixpay.GetPrice(payAsset, receiveAsset, amount, except, mixpayChan)
	}
	pandoswapChan = make(chan float64)
	pandoswapfeeChan = make(chan float64)
	go pandoswap.GetPrice(payAsset, chainAsset, "", fee, pandoswapfeeChan)
	go pandoswap.GetPrice(payAsset, receiveAsset, amount, except, pandoswapChan)

	pandoswapPrice = <-pandoswapChan
	pandoswapFee = <-pandoswapfeeChan
	if cex {
		mixpayPrice = <-mixpayChan
		mixpayFee = <-mixpayFeeChan
	}
	bestPrice, bestFee := findLargest(mixpayPrice, pandoswapPrice), findSmallest(mixpayFee, pandoswapFee)

	if cex {
		return &constants.PriceSimpleResp{
			Amount: bestPrice,
			Fee:    bestFee,
		}
	} else {
		return &constants.PriceSimpleResp{
			Amount: pandoswapPrice,
			Fee:    pandoswapFee,
		}
	}
}

func GetPriceAll(payAsset, receiveAsset, amount, except string, cex bool) []constants.PriceAllResp {
	var mixpayChan, mixpayFeeChan chan float64
	var mixpayPrice, mixpayFee float64
	var pandoswapChan, pandoswapfeeChan chan float64
	var pandoswapPrice, pandoswapFee float64
	chainAsset, fee := AssetFee(receiveAsset)
	if cex {
		mixpayChan = make(chan float64)
		mixpayFeeChan = make(chan float64)
		go mixpay.GetPrice(payAsset, chainAsset, "", fee, mixpayFeeChan)
		go mixpay.GetPrice(payAsset, receiveAsset, amount, except, mixpayChan)
	}
	pandoswapChan = make(chan float64)
	pandoswapfeeChan = make(chan float64)
	go pandoswap.GetPrice(payAsset, chainAsset, "", fee, pandoswapfeeChan)
	go pandoswap.GetPrice(payAsset, receiveAsset, amount, except, pandoswapChan)

	if cex {
		mixpayPrice = <-mixpayChan
		mixpayFee = <-mixpayFeeChan
	}
	pandoswapPrice = <-pandoswapChan
	pandoswapFee = <-pandoswapfeeChan

	if cex {
		return []constants.PriceAllResp{
			{
				Name:   mixpay.NAME,
				Amount: fmt.Sprintf("%v", mixpayPrice),
				Fee:    fmt.Sprintf("%v", mixpayFee),
			},
			{
				Name:   pandoswap.NAME,
				Amount: fmt.Sprintf("%v", pandoswapPrice),
				Fee:    fmt.Sprintf("%v", pandoswapFee),
			},
		}
	} else {
		return []constants.PriceAllResp{
			{
				Name:   pandoswap.NAME,
				Amount: fmt.Sprintf("%v", pandoswapPrice),
				Fee:    fmt.Sprintf("%v", pandoswapFee),
			},
		}
	}
}

func Swap(orderId, payAsset, receiveAsset, amount string, cex bool) (*mixin.TransferInput, *mixin.TransferInput) {
	// Find best provider and filter disabled
	priceAll := GetPriceAll(payAsset, receiveAsset, amount, "", cex)
	priceAll = filterDisabledProvider(priceAll, DISABLED)
	best := findMaxAmountAndMinFee(priceAll)
	var i0, i1 *mixin.TransferInput
	fromAmount, _ := decimal.NewFromString(best.Amount)
	fromFeeAmount, _ := decimal.NewFromString(best.Fee)
	amount = fromAmount.Sub(fromFeeAmount).String()

	//TODO also sub bridge fee
	//fmt.Printf("Amount:%s\nFee:%s\namount: %s\n", best.Amount, best.Fee, amount)

	if cex && slices.Contains(CEX, best.Name) {
		switch best.Name {
		case CEX[0]:
			i0 = mixpay.Swap(orderId, payAsset, receiveAsset, amount, false)
		}

	} else if slices.Contains(DEX, best.Name) {
		switch best.Name {
		case DEX[0]:
			i0 = pandoswap.Swap(orderId, payAsset, receiveAsset, amount)
		}
	}

	if cex && slices.Contains(CEX, best.FeeName) {
		switch best.FeeName {
		case CEX[0]:
			i1 = mixpay.Swap(orderId, payAsset, receiveAsset, best.Fee, false)
		}
	} else if slices.Contains(DEX, best.FeeName) {
		switch best.FeeName {
		case DEX[0]:
			i1 = pandoswap.Swap(orderId, payAsset, receiveAsset, best.Fee)
		}

	}

	return i0, i1
}

func GetStatus() {

}
