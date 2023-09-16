package workers

import (
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
	// list orders
	orders, err := sw.store.ListOrders(100)
	if err != nil {
		return err
	}
	// if status == constants.PrefixReceived
	for _, o := range orders {
		if o.Status == constants.PrefixReceived {
			// swap
			sw.swap()
		}
	}
}

func (sw *SwapWorker) swap(o *constants.Order) {
	// send swap
	providers.Swap()
	// set status sent
	// set swap trace
}
