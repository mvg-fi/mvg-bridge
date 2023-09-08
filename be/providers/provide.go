package providers

import (
	"fmt"

	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/providers/mixpay"
	"github.com/mvg-fi/mvg-bridge/providers/pandoswap"
)

func GetPriceSimple(payAsset, receiveAsset, amount, except string, cex bool) (float64, float64) {
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
		return findLargest(mixpayPrice, pandoswapPrice), findSmallest(mixpayFee, pandoswapFee)
	} else {
		return pandoswapPrice, pandoswapFee
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
				Name:   "Mixpay",
				Amount: fmt.Sprintf("%v", mixpayPrice),
				Fee:    fmt.Sprintf("%v", mixpayFee),
			},
			{
				Name:   "PandoSwap",
				Amount: fmt.Sprintf("%v", pandoswapPrice),
				Fee:    fmt.Sprintf("%v", pandoswapFee),
			},
		}
	} else {
		return []constants.PriceAllResp{
			{
				Name:   "PandoSwap",
				Amount: fmt.Sprintf("%v", pandoswapPrice),
				Fee:    fmt.Sprintf("%v", pandoswapFee),
			},
		}
	}
}

func GetStatus() {

}
