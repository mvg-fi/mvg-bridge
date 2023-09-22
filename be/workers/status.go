package workers

import (
	"github.com/mvg-fi/mvg-bridge/config"
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

func (sw *StatusWorker) update() {
	// Get Swap Orders
}
