package providers

import (
	"fmt"
	"slices"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/providers/mixpay"
	"github.com/mvg-fi/mvg-bridge/providers/pandoswap"
)

var (
	PROVIDERS = []string{mixpay.NAME, pandoswap.NAME}
	CEX       = []string{mixpay.NAME}
	DEX       = []string{pandoswap.NAME}
	DISABLED  = []string{""}
)

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

func Swap(orderId, payAsset, receiveAsset, amount string, cex bool) *mixin.TransferInput {
	// Find best provider and filter disabled
	priceAll := GetPriceAll(payAsset, receiveAsset, amount, "", cex)
	priceAll = filterDisabledProvider(priceAll, DISABLED)
	best := findMaxAmountAndMinFee(priceAll)

	// Swap
	if cex && slices.Contains(CEX, best.Name) {
		switch best.Name {
		case CEX[0]:
			return mixpay.Swap(orderId, payAsset, receiveAsset, amount, false)
		}
	}
	if slices.Contains(DEX, best.Name) {
		switch best.Name {
		case DEX[0]:
			return pandoswap.Swap(orderId, payAsset, receiveAsset, amount)
		}
	}
	return &mixin.TransferInput{}
}

func GetStatus() {

}
