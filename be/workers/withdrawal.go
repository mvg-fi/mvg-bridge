package workers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/encoding"
	"github.com/mvg-fi/mvg-bridge/store"
)

type WithdrawalWorker struct {
	group *mtg.Group
	store *store.BadgerStore
	conf  *config.Configuration
}

func NewWithdrawalWorker(group *mtg.Group, store *store.BadgerStore, conf *config.Configuration) *WithdrawalWorker {
	return &WithdrawalWorker{
		group,
		store,
		conf,
	}
}

func getWithdrawalParams(store *store.BadgerStore, o *constants.Order) (string, string, string, string) {
	//TODO: Get amount by reading swap orders
	// Get memo by designing a withdrawal memo format
	mainSwapTrace := mixin.UniqueConversationID(o.TraceID, constants.SwapTypeMainInit)
	feeSwapTrace := mixin.UniqueConversationID(o.TraceID, constants.SwapTypeFeeInit)
	main, err := store.ReadSwap(mainSwapTrace)
	if err != nil {
		logger.Errorf("store.ReadSwap(%s) => %v", mainSwapTrace, err)
		return "", "", "", ""
	}
	fee, err := store.ReadSwap(feeSwapTrace)
	if err != nil {
		logger.Errorf("store.ReadSwap(%s) => %v", feeSwapTrace, err)
		return "", "", "", ""
	}

	mainMemo, feeMemo := encoding.GetWithdrawalMemo(o, main, fee)
	return mainMemo, feeMemo, main.Receive, fee.Receive
}

func (ww *WithdrawalWorker) withdraw(ctx context.Context, o *constants.Order) error {
	mainMemo, feeMemo, mainAmount, feeAmount := getWithdrawalParams(ww.store, o)
	if mainAmount == "" || feeAmount == "" {
		return errors.New(fmt.Sprintf("Swap not completed:%s", o.TraceID))
	}
	mainTrace := mixin.UniqueConversationID(o.TraceID, constants.WithdrawMainInit)
	feeTrace := mixin.UniqueConversationID(o.TraceID, constants.WithdrawFeeInit)
	err := ww.group.BuildTransaction(ctx, o.ToAssetID, []string{constants.MVGBridgeUUID}, int(1), mainAmount, mainMemo, mainTrace, constants.WithdrawMainInit)
	if err != nil {
		return err
	}
	err = ww.group.BuildTransaction(ctx, o.ToAssetID, []string{constants.MVGBridgeUUID}, int(1), feeAmount, feeMemo, feeTrace, constants.WithdrawFeeInit)
	if err != nil {
		return err
	}

	ww.store.WriteWithdrawalFull(&constants.WithdrawalFull{
		OrderID:   o.TraceID,
		MainTrace: mainTrace,
		FeeTrace:  feeTrace,
		Asset:     o.ToAssetID,
		Amount:    mainAmount,
		Address:   o.Address,
		Memo:      o.Memo,
	})
	return nil
}

func (ww *WithdrawalWorker) run(ctx context.Context) error {
	// Loop all order swap success
	orders, err := ww.store.ListOrders(100)
	if err != nil {
		return err
	}
	for _, o := range orders {
		if o.Status != constants.StatusSwapSuccess {
			continue
		}

		// Send them to MVG Bridge with proper memo
		err = ww.withdraw(ctx, o)
		if err != nil {
			logger.Printf("ww.withdraw(%+v) => %v", o, err)
			continue
		}

		// Set order status withdrawal sent
		o.Status = constants.StatusWithdrawSent
		err = ww.store.UpdateOrder(o.TraceID, o)
		if err != nil {
			logger.Printf("store.UpdateOrder(%s, %+v) => %v", o.TraceID, o, err)
			continue
		}
	}

	return nil
}

func (ww *WithdrawalWorker) Run(ctx context.Context) {
	logger.Println("WithdrawalWorker started")
	go func() {
		for {
			err := ww.run(ctx)
			if err != nil {
				logger.Printf("ww.run() => %v", err)
				continue
			}
			time.Sleep(3 * time.Second)
		}
	}()
}
