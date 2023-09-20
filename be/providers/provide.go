package providers

import (
	"fmt"
	"slices"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/providers/exinone"
	"github.com/mvg-fi/mvg-bridge/providers/pandoswap"
)

var (
	PROVIDERS = []string{exinone.NAME, pandoswap.NAME}
	CEX       = []string{exinone.NAME}
	DEX       = []string{pandoswap.NAME}
	DISABLED  = []string{}
)

// bestPrice is the amount of asset will be received
// bestFee is the amount of original asset will be charged
func GetPriceSimple(payAsset, receiveAsset, amount, except string, cex bool) *constants.PriceSimpleResp {
	var exinoneChan, exinoneFeeChan chan float64
	var exinonePrice, exinoneFee float64
	var pandoswapChan, pandoswapfeeChan chan float64
	var pandoswapPrice, pandoswapFee float64
	chainAsset, fee := AssetFee(receiveAsset)
	if cex {
		exinoneChan = make(chan float64)
		exinoneFeeChan = make(chan float64)
		go exinone.GetPrice(chainAsset, payAsset, fee, exinoneFeeChan)
		if len(except) == 0 {
			go exinone.GetPrice(payAsset, receiveAsset, amount, exinoneChan)
		} else {
			go exinone.GetPrice(chainAsset, payAsset, except, exinoneChan)
		}
	}
	pandoswapChan = make(chan float64)
	pandoswapfeeChan = make(chan float64)
	go pandoswap.GetPrice(payAsset, chainAsset, "", fee, pandoswapfeeChan)
	go pandoswap.GetPrice(payAsset, receiveAsset, amount, except, pandoswapChan)

	pandoswapPrice = <-pandoswapChan
	pandoswapFee = <-pandoswapfeeChan
	if cex {
		exinonePrice = <-exinoneChan
		exinoneFee = <-exinoneFeeChan
	}
	bestPrice, bestFee := findLargest(exinonePrice, pandoswapPrice), findSmallest(exinoneFee, pandoswapFee)

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
	var exinoneChan, exinoneFeeChan chan float64
	var exinonePrice, exinoneFee float64
	var pandoswapChan, pandoswapfeeChan chan float64
	var pandoswapPrice, pandoswapFee float64
	chainAsset, fee := AssetFee(receiveAsset)
	if cex {
		exinoneChan = make(chan float64)
		exinoneFeeChan = make(chan float64)
		go exinone.GetPrice(chainAsset, payAsset, fee, exinoneFeeChan)
		if len(except) == 0 {
			go exinone.GetPrice(payAsset, receiveAsset, amount, exinoneChan)
		} else {
			go exinone.GetPrice(chainAsset, payAsset, except, exinoneChan)
		}
	}
	pandoswapChan = make(chan float64)
	pandoswapfeeChan = make(chan float64)
	go pandoswap.GetPrice(payAsset, chainAsset, "", fee, pandoswapfeeChan)
	go pandoswap.GetPrice(payAsset, receiveAsset, amount, except, pandoswapChan)

	if cex {
		exinonePrice = <-exinoneChan
		exinoneFee = <-exinoneFeeChan
	}
	pandoswapPrice = <-pandoswapChan
	pandoswapFee = <-pandoswapfeeChan

	if cex {
		return []constants.PriceAllResp{
			{
				Name:   exinone.NAME,
				Amount: fmt.Sprintf("%v", exinonePrice),
				Fee:    fmt.Sprintf("%v", exinoneFee),
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
	var i0, i1 *mixin.TransferInput
	priceAll := GetPriceAll(payAsset, receiveAsset, amount, "", cex)
	priceAll = filterDisabledProvider(priceAll, DISABLED)
	best := findMaxAmountAndMinFee(priceAll)
	amount = AfterBridgeFee(best.Amount, best.Fee)

	// TODO: withdrawal fee * 1.05 to avoid price change
	if cex && slices.Contains(CEX, best.Name) {
		switch best.Name {
		case CEX[0]:
			i0 = exinone.Swap(constants.SwapTypeMain, orderId, payAsset, receiveAsset, amount)
		}

	} else if slices.Contains(DEX, best.Name) {
		switch best.Name {
		case DEX[0]:
			i0 = pandoswap.Swap(constants.SwapTypeMain, orderId, payAsset, receiveAsset, amount)
		}
	}

	if cex && slices.Contains(CEX, best.FeeName) {
		switch best.FeeName {
		case CEX[0]:
			i1 = exinone.Swap(constants.SwapTypeFee, orderId, payAsset, receiveAsset, best.Fee)
		}
	} else if slices.Contains(DEX, best.FeeName) {
		switch best.FeeName {
		case DEX[0]:
			i1 = pandoswap.Swap(constants.SwapTypeFee, orderId, payAsset, receiveAsset, best.Fee)
		}

	}

	return i0, i1
}

func GetStatus() {

}
