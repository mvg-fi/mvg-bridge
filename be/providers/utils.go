package providers

import (
	"math"
	"slices"
	"strconv"

	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/shopspring/decimal"
)

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

func findMaxAmountAndMinFee(prices []constants.PriceAllResp) *constants.PriceBest {
	maxAmount := prices[0].Amount
	minFee := prices[0].Fee
	bestAmountName := prices[0].Name
	bestFeeName := prices[0].Name

	for _, price := range prices {
		// Compare the Amount with the current maximum
		if price.Amount > maxAmount {
			maxAmount = price.Amount
			bestAmountName = price.Name
		}

		// Convert Amount and Fee strings to floating-point numbers for comparison
		feeValue := parseAmount(price.Fee)

		// Compare the Fee with the current minimum
		if feeValue < parseAmount(minFee) {
			minFee = price.Fee
			bestFeeName = price.Name
		}
	}

	return &constants.PriceBest{
		Name:    bestAmountName,
		Amount:  maxAmount,
		FeeName: bestFeeName,
		Fee:     minFee,
	}
}

func parseAmount(amount string) float64 {
	value, _ := strconv.ParseFloat(amount, 64)
	return value
}

func filterDisabledProvider(all []constants.PriceAllResp, disabled []string) []constants.PriceAllResp {
	var left []constants.PriceAllResp
	for _, source := range all {
		if !slices.Contains(DISABLED, source.Name) {
			left = append(left, source)
		}
	}
	return left
}

func AfterBridgeFee(amount, fee string) string {
	am, _ := decimal.NewFromString(amount)
	fe, _ := decimal.NewFromString(fee)
	after := am.Sub(fe)
	return after.Sub(after.Mul(constants.BridgeFee)).RoundUp(8).String()
}

func fillAssetBlankField() {

}
