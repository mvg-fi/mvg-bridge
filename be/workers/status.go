package workers

import (
	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/providers"
	"github.com/mvg-fi/mvg-bridge/store"
)

type StatusWorker struct {
	store *store.BadgerStore
	conf  *config.Configuration
}

func NewStatusWorker(store *store.BadgerStore, conf *config.Configuration) *StatusWorker {
	return &StatusWorker{
		store,
		conf,
	}
}

func (sw *StatusWorker) loopSwaps() error {
	// Get Swap Orders
	swaps, err := sw.store.ListSwaps(100)
	if err != nil {
		return err
	}
	for _, s := range swaps {
		sw.update(s)
	}
	return nil
}

func (sw *StatusWorker) update(s *constants.Swap) error {
	status := providers.GetStatus(s.Provider, s.OrderID)

	err := sw.store.UpdateSwap(s.OrderID, &constants.Swap{
		OrderID:  s.OrderID,
		Status:   status,
		Provider: s.Provider,
	})
	if err != nil {
		return err
	}
	return nil
}
