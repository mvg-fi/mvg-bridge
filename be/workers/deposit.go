package workers

import (
	"context"
	"time"

	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/store"
	"github.com/mvg-fi/mvg-bridge/users"
	"github.com/shopspring/decimal"
)

type DepositWorker struct {
	*users.Proxy
	store *store.BadgerStore
	conf  *config.Configuration
}

func NewDepositWorker(proxy *users.Proxy, store *store.BadgerStore, conf *config.Configuration) *DepositWorker {
	return &DepositWorker{
		proxy,
		store,
		conf,
	}
}

// Collect all snapshots
func (dw *DepositWorker) loopSnapshots(ctx context.Context) error {
	ckpt, err := dw.store.ReadSnapshotsCheckpoint(ctx)
	if err != nil {
		return err
	}
	snapshots, err := dw.ReadNetworkSnapshots(ctx, "", ckpt, "asc", 500)
	if err != nil {
		logger.Verbosef("dw.ReadNetworkSnapshots(%s) => %d %v", ckpt, len(snapshots), err)
		return err
	}
	for _, s := range snapshots {
		ckpt = s.CreatedAt
		if s.UserID == "" {
			continue
		}
		if s.Amount.Cmp(decimal.NewFromFloat(0.00000001)) < 0 {
			continue
		}
		logger.Verbosef("dw.loopSnapshots(%s) => %d %v => %v", ckpt, len(snapshots), err, *s)
		err = dw.store.WriteSnapshot(s)
		if err != nil {
			return err
		}
	}

	err = dw.store.WriteSnapshotsCheckpoint(ctx, ckpt)
	if err != nil {
		return err
	}
	if len(snapshots) < 500 {
		time.Sleep(time.Millisecond * 300)
	}
	return nil
}

// Process all snapshots and delete after process
func (dw *DepositWorker) processSnapshots(ctx context.Context) {
	snapshots, err := dw.store.ListSnapshots(100)
	if err != nil {
		logger.Errorf("dw.store.ListSnapshots(100) => %v", err)
	}

	for _, s := range snapshots {
		user, err := dw.store.ReadUser(s.UserID)
		if err != nil {
			logger.Errorf("dw.store.ListSnapshots(100) => %v", err)
		}
		if user == nil {
			continue
		}
		u := users.User{*user}
		err = u.Handle(ctx, dw.store, dw.conf, s)
		if err != nil {
			logger.Errorf("user.Handle() => %v", err)
		}
	}

	err = dw.store.DeleteSnapshots(snapshots)
	if err != nil {
		logger.Errorf("dw.store.DeleteSnapshots(%v) => %v", snapshots, err)
	}
	if len(snapshots) < 100 {
		time.Sleep(time.Millisecond * 200)
	}
}

func (dw *DepositWorker) Run(ctx context.Context) {
	logger.Println("DepositWorker started")

	go func() {
		for {
			dw.loopSnapshots(ctx)
		}
	}()
	go func() {
		for {
			dw.processSnapshots(ctx)
		}
	}()
}
