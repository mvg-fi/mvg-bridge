package workers

import (
	"context"

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
		err := sw.swap(o)
		if err != nil {
			logger.Errorf("sw.swap(%v) => %v", o, err)
		}
	}
}

func (sw *SwapWorker) swap(ctx context.Context, o *constants.Order) error {
	// get amount swap
	i0, i1 := providers.Swap(o.TraceID, o.FromAssetID, o.ToAssetID, o.Amount, o.Cex)
	// main
	err := sw.group.BuildTransaction(ctx, i0.AssetID, i0.OpponentMultisig.Receivers, i0.OpponentMultisig.Threshold, i0.Amount, i0.Memo, i0.TraceID, constants.SwapTypeMainInit)
	// fee
	err := sw.group.BuildTransaction(ctx, i1.AssetID, i1.OpponentMultisig.Receivers, i1.OpponentMultisig.Threshold, i1.Amount, i1.Memo, i1.TraceID, constants.SwapTypeFeeInit)

	//(ctx, evt.Asset, evt.Members, evt.Threshold, amount, memo, traceId, p.Identifier
	// get fee swap

	// set status sent
	o.Status = constants.StatusSwapSent
	err := sw.store.UpdateOrder(o.OrderID, o)
	if err != nil {
		return err
	}
}
