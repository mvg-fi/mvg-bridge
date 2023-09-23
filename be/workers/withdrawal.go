package workers

import (
	"context"

	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/store"
)

type WithdrawalWorker struct {
	store *store.BadgerStore
	conf  *config.Configuration
}

func NewWithdrawalWorker(store *store.BadgerStore, conf *config.Configuration) *WithdrawalWorker {
	return &WithdrawalWorker{
		store,
		conf,
	}
}

func (ww *WithdrawalWorker) processSnapshots(ctx context.Context) {
	// Loop all swap received
	// Send them to MVG Bridge with proper memo
	// Add to withdrawal store
}

func (ww *WithdrawalWorker) Run(ctx context.Context) {
	logger.Println("WithdrawalWorker started")
	go func() {
		for {
			ww.processSnapshots(ctx)
		}
	}()
}
