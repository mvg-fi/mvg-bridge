package workers

import (
	"context"
	"time"

	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/providers"
	"github.com/mvg-fi/mvg-bridge/store"
)

type SwapWorker struct {
	group *mtg.Group
	store *store.BadgerStore
	conf  *config.Configuration
}

func NewSwapWorker(group *mtg.Group, store *store.BadgerStore, conf *config.Configuration) *SwapWorker {
	return &SwapWorker{
		group,
		store,
		conf,
	}
}

func (sw *SwapWorker) process(ctx context.Context) error {
	orders, err := sw.store.ListOrders(100)
	if err != nil {
		return err
	}
	for _, o := range orders {
		if o.Status != constants.StatusReceived {
			continue
		}
		err := sw.swap(ctx, o)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sw *SwapWorker) swap(ctx context.Context, o *constants.Order) error {
	// get amount swap
	p0, p1, i0, i1 := providers.Swap(o.TraceID, o.FromAssetID, o.ToAssetID, o.Amount, o.Cex)
	// main
	err := sw.group.BuildTransaction(ctx, i0.AssetID, i0.OpponentMultisig.Receivers, int(i0.OpponentMultisig.Threshold), i0.Amount.String(), i0.Memo, i0.TraceID, constants.SwapTypeMainInit)
	// create a new swap store
	err = sw.store.WriteSwap(&constants.Swap{
		OrderID:  o.TraceID,
		TraceID:  i0.TraceID,
		Status:   constants.StatusSwapSent,
		Provider: p0,
		Type:     constants.SwapTypeMain,
	})
	if err != nil {
		return err
	}

	// fee
	err = sw.group.BuildTransaction(ctx, i1.AssetID, i1.OpponentMultisig.Receivers, int(i1.OpponentMultisig.Threshold), i1.Amount.String(), i1.Memo, i1.TraceID, constants.SwapTypeFeeInit)
	// create a fee swap store
	err = sw.store.WriteSwap(&constants.Swap{
		OrderID:  o.TraceID,
		TraceID:  i1.TraceID,
		Status:   constants.StatusSwapSent,
		Provider: p1,
		Type:     constants.SwapTypeFee,
	})
	if err != nil {
		return err
	}

	// set status sent
	o.Status = constants.StatusSwapSent
	err = sw.store.UpdateOrder(o.TraceID, o)
	if err != nil {
		return err
	}
	return nil
}

func (sw *SwapWorker) Run(ctx context.Context) {
	logger.Println("SwapWorker started")
	for {
		err := sw.process(ctx)
		if err != nil {
			logger.Errorf("sw.swap() => %v", err)
		}
		time.Sleep(3 * time.Second)
	}
}
