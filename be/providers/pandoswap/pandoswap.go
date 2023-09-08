package pandoswap

import (
	"context"
	"fmt"
	"log"

	fswap "github.com/fox-one/4swap-sdk-go"
	"github.com/shopspring/decimal"
)

func GetPrice(payAsset, receiveAsset, amount, except string, ch chan<- float64) {
	ctx := context.Background()
	pairs, err := fswap.ListPairs(ctx)
	if err != nil {
		log.Println("fswap.ListPairs() => ", err)
	}
	if len(amount) > 0 {
		amt, err := decimal.NewFromString(amount)
		if err != nil {
			log.Println("decimal.NewFromString() => ", err)
		}

		route, err := fswap.Route(pairs, payAsset, receiveAsset, amt)
		if err != nil {
			log.Println("fswap.Route() => ", err)
		}
		val, _ := route.FillAmount.Float64()
		ch <- val
		return
	}

	amt, err := decimal.NewFromString(except)
	if err != nil {
		log.Println("decimal.NewFromString() => ", err)
	}
	route, err := fswap.ReverseRoute(pairs, payAsset, receiveAsset, amt)
	if err != nil {
		log.Println("fswap.Route() => ", err)
	}
	val, _ := route.PayAmount.Float64()
	ch <- val
}

func GetStatus(traceId string) {
	ctx := context.Background()
	order, err := fswap.ReadOrder(ctx, traceId)
	if err != nil {
		log.Println("fswap.ReadOrder() => ", err)
	}
	fmt.Printf("%+v", order)
}
