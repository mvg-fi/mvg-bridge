package workers

import (
	"github.com/fox-one/mixin-sdk-go"
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
		err = sw.update(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sw *StatusWorker) update(s *constants.Swap) error {
	status, amount := providers.GetStatus(s.Provider, s.TraceID)
	s.Status = status
	s.Amount = amount

	err := sw.store.UpdateSwap(s.TraceID, s)
	if err != nil {
		return err
	}

	// Check if both swap success, update order to withdrawal
	var pairTrace string
	switch s.Type {
	case constants.SwapTypeMain:
		pairTrace = mixin.UniqueConversationID(s.TraceID, constants.SwapTypeFeeInit)
	case constants.SwapTypeFee:
		pairTrace = mixin.UniqueConversationID(s.TraceID, constants.SwapTypeMainInit)
	}

	pair, err := sw.store.ReadSwap(pairTrace)
	if err != nil {
		return err
	}
	if pair.Status != constants.StatusSwapSuccess {
		return nil
	}

	if order, err := sw.store.ReadOrder(s.OrderID); err == nil {
		order.Status = constants.StatusSwapSuccess
		err = sw.store.UpdateOrder(s.OrderID, order)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}
