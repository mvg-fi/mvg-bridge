package pandoswap

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	fswap "github.com/fox-one/4swap-sdk-go"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/google/uuid"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/pandodao/mtg/mtgpack"
	"github.com/pandodao/mtg/protocol"
	"github.com/shopspring/decimal"
)

var (
	NAME                  = "PandoSwap"
	PandoSwapMTGMembers   = []string{"a753e0eb-3010-4c4a-a7b2-a7bda4063f62", "099627f8-4031-42e3-a846-006ee598c56e", "aefbfd62-727d-4424-89db-ae41f75d2e04", "d68ca71f-0e2c-458a-bb9c-1d6c2eed2497", "e4bc0740-f8fe-418c-ae1b-32d9926f5863"}
	PandoSwapMTGThreshold = uint8(3)
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

func Swap(orderId, payAsset, receiveAsset, amount string) *mixin.TransferInput {
	amt, _ := decimal.NewFromString(amount)
	memo := generateSwapMemo(constants.MTGMembers, constants.MTGThreshold, orderId, receiveAsset, amount, "")
	input := &mixin.TransferInput{
		AssetID: payAsset,
		TraceID: mixin.UniqueConversationID(orderId, "swap:init"),
		Amount:  amt,
		Memo:    memo,
	}
	input.OpponentMultisig.Receivers = PandoSwapMTGMembers
	input.OpponentMultisig.Threshold = PandoSwapMTGThreshold
	return input
}

func generateSwapMemo(members []string, threshold uint8, orderId, receiveAsset, amount, minReceive string) string {
	header := protocol.Header{
		Version:    1,
		ProtocolID: protocol.ProtocolFswap,
		FollowID:   uuid.MustParse(mixin.UniqueConversationID(orderId, "4swap:followID")),
		Action:     3,
	}

	mbrs := parseUUIDs(members)
	receiver := protocol.MultisigReceiver{
		Version:   1,
		Members:   mbrs,
		Threshold: threshold,
	}

	min, _ := decimal.NewFromString(minReceive)
	if minReceive == "" {
		min = decimal.NewFromInt(0)
	}

	enc := mtgpack.NewEncoder()
	if err := enc.EncodeValues(header, receiver, uuid.MustParse(receiveAsset), min); err != nil {
		logger.Errorf("%v", err)
	}

	return base64.StdEncoding.EncodeToString(enc.Bytes())
}

func parseUUIDs(uuids []string) []uuid.UUID {
	var a []uuid.UUID
	for i := 0; i < len(uuids); i++ {
		a = append(a, uuid.MustParse(uuids[i]))
	}
	return a
}
