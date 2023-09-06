package providers

import (
	"math"

	"github.com/mvg-fi/mvg-bridge/providers/mixpay"
	"github.com/mvg-fi/mvg-bridge/providers/pandoswap"
)

func GetPrice(payAsset, receiveAsset, amount, except string) float64 {
	mpPrice := mixpay.GetPrice(payAsset, receiveAsset, amount, except)
	psPrice := pandoswap.GetPrice(payAsset, receiveAsset, amount, except)
	if len(except) == 0 {
		return findLargest(mpPrice, psPrice)
	}
	return findSmallest(mpPrice, psPrice)
}

func findLargest(numbers ...float64) float64 {
	largest := math.Inf(-1)
	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}
	return largest
}

func findSmallest(numbers ...float64) float64 {
	smallest := math.Inf(1)
	for _, num := range numbers {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}
