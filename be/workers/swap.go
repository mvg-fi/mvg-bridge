package workers

import (
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/providers"
	"github.com/mvg-fi/mvg-bridge/store"
)

type SwapWorker struct {
	store *store.BadgerStore
	conf  *config.Configuration
}

func NewSwapWorker(store *store.BadgerStore, conf *config.Configuration) *SwapWorker {
	return &SwapWorker{
		store,
		conf,
	}
}

func (sw *SwapWorker) process() error {
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

func (sw *SwapWorker) swap(o *constants.Order) error {
	// get amount swap
	input := providers.Swap(o.TraceID, o.FromAssetID, o.ToAssetID, o.Amount, o.Cex)
	// get fee swap

	// set status sent
	o.Status = constants.StatusSwapSent
	err := sw.store.UpdateOrder(o.OrderID, o)
	if err != nil {
		return err
	}
}
