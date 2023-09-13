package workers

import (
	"context"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/common/uuid"
	"github.com/mvg-fi/mvg-bridge/store"
	"github.com/mvg-fi/mvg-bridge/users"
	"github.com/shopspring/decimal"
)

// Monitor all deposit pending prefix addresses
// If arrived, complete the rest job
// If passed, refund to the sender

type DepositWorker struct {
	*users.Proxy
	store *store.BadgerStore
}

func NewDepositWorker(proxy *users.Proxy, store *store.BadgerStore) *DepositWorker {
	return &DepositWorker{
		proxy,
		store,
	}
}

func (dw *DepositWorker) loopSnapshots(ctx context.Context) error {
	ckpt, err := dw.store.ReadSnapshotsCheckpoint(ctx)
	if err != nil {
		return err
	}
	snapshots, err := dw.ReadNetworkSnapshots(ctx, "", ckpt, "asc", 500)
	logger.Verbosef("dw.ReadNetworkSnapshots(%s) => %d %v", ckpt, len(snapshots), err)
	if err != nil {
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
		logger.Verbosef("dw.loopsnapshots(%s) => %d %v => %v", ckpt, len(snapshots), err, *s)
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

func (dw *DepositWorker) ProcessSnapshots(ctx context.Context) {
	snapshots, err := dw.store.ListSnapshots(100)
	if err != nil {
		panic(err)
	}

	for _, s := range snapshots {
		// Get order ID by Lock
		orderID, err := dw.store.LockGet(s.UserID, s.AssetID)
		if err != nil {
			logger.Errorf("LockGet(%s, %s) => %v", s.UserID, s.AssetID, err)
		}
		order, err := dw.store.ReadOrder(orderID)
		if err != nil {
			logger.Errorf("dw.store.ReadOrder(%s) => %v", traceID, err)
		}

		// Transfer to MTG
		input := mixin.TransferInput{
			AssetID: s.AssetID,
			Amount:  s.Amount,
			TraceID: uuid.NewV4(),
		}
		input.OpponentMultisig.Receivers = string[10]{""}
		input.OpponentMultisig.Threshold = uint8(MVMThreshold)
		input.Memo = orderID
		u.sendUntilSufficient(ctx, &input)
		// Remove lock

		println(s)
	}

	err = dw.store.DeleteSnapshots(snapshots)
	if err != nil {
		panic(err)
	}
	if len(snapshots) < 100 {
		time.Sleep(time.Millisecond * 200)
	}
}
