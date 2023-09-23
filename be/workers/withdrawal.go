package workers

import (
	"context"
	"time"

	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/common/uuid"
	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/constants"
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

func getWithdrawalParams(o *constants.Order) (string, string) {
	//TODO: Get amount by reading swap orders
	// Get memo by designing a withdrawal memo format
	mainMemo := ""
	feeMemo := ""
	return mainMemo, feeMemo
}

func (ww *WithdrawalWorker) withdraw(ctx context.Context, o *constants.Order) error {
	mainAmount, feeAmount := "", ""
	mainMemo, feeMemo := getWithdrawalParams(o)
	err := ww.group.BuildTransaction(ctx, o.ToAssetID, []string{constants.MVGBridgeUUID}, int(1), mainAmount, mainMemo, uuid.NewV4(), constants.WithdrawMainInit)
	if err != nil {
		return err
	}
	err = ww.group.BuildTransaction(ctx, o.ToAssetID, []string{constants.MVGBridgeUUID}, int(1), feeAmount, feeMemo, uuid.NewV4(), constants.WithdrawFeeInit)
	if err != nil {
		return err
	}
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
		err = store.UpdateOrder(o.TraceID, o)
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
