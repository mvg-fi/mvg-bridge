package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/encoding"
	"github.com/mvg-fi/mvg-bridge/store"
	"github.com/shopspring/decimal"
)

// withdrawal worker for bot, in charge of withdrawing the assets sent by the MTG

type WithbotWorker struct {
	store  *store.BadgerStore
	client *mixin.Client
	ka     *mixin.KeystoreAuth
	pin    string
}

func NewWithbotWorker(ctx context.Context, key *mixin.Keystore, pin string, store *store.BadgerStore) *WithbotWorker {
	client, err := mixin.NewFromKeystore(key)
	if err != nil {
		panic(err)
	}
	ka, err := mixin.AuthFromKeystore(key)
	if err != nil {
		panic(err)
	}

	return &WithbotWorker{
		client: client,
		ka:     ka,
		store:  store,
		pin:    pin,
	}
}

func decodeMemo(memo string) (*constants.Withdrawal, error) {
	w, err := encoding.RetriveWithdrawalMemo(memo)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (wb *WithbotWorker) withdraw(ctx context.Context, w *constants.Withdrawal) error {
	ain := mixin.CreateAddressInput{
		AssetID:     w.Asset,
		Destination: w.Address,
		Label:       w.OrderID,
		Tag:         "",
	}
	addr, err := wb.client.CreateAddress(ctx, ain, wb.pin)
	if err != nil {
		return err
	}

	amt, _ := decimal.NewFromString(w.Amount)
	win := mixin.WithdrawInput{
		AddressID: addr.AddressID,
		TraceID:   mixin.UniqueConversationID(w.OrderID, constants.WithdrawFinal),
		Amount:    amt,
		Memo:      w.Memo,
	}
	_, err = wb.client.Withdraw(ctx, win, wb.pin)
	if err != nil {
		return err
	}
	return wb.client.DeleteAddress(ctx, ain.AssetID, wb.pin)
}

func (wb *WithbotWorker) loopSnapshots(ctx context.Context) error {
	ckpt, err := wb.store.ReadSnapshotsCheckpoint(ctx)
	if err != nil {
		return err
	}
	snapshots, err := wb.client.ReadNetworkSnapshots(ctx, "", ckpt, "asc", 500)
	if err != nil {
		logger.Verbosef("wb.ReadNetworkSnapshots(%s) => %d %v", ckpt, len(snapshots), err)
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
		logger.Verbosef("wb.loopSnapshots(%s) => %d %v => %v", ckpt, len(snapshots), err, *s)
		err = wb.store.WriteSnapshot(s)
		if err != nil {
			return err
		}
	}

	err = wb.store.WriteSnapshotsCheckpoint(ctx, ckpt)
	if err != nil {
		return err
	}
	if len(snapshots) < 500 {
		time.Sleep(time.Millisecond * 300)
	}
	return nil
}

func (wb *WithbotWorker) process(ctx context.Context, s *mixin.Snapshot) error {
	// skip all none withdrawal memo, refund if possible
	w, err := decodeMemo(s.Memo)
	if err != nil {
		return fmt.Errorf("invalid withdrawal memo %s => %v", s.Memo, err)
	}

	if w.Asset == "" || w.Amount == "" || w.Address == "" {
		return fmt.Errorf("invalid withdrawal params")
	}

	old, err := wb.store.ReadWithdrawal(w.OrderID)
	if err != nil {
		return fmt.Errorf("wb.store.ReadWithdrawal(%s)", w.OrderID)
	}

	wf := constants.WithdrawalFull{
		OrderID: w.OrderID,
		Asset:   w.Asset,
		Amount:  w.Amount,
		Address: w.Address,
		Memo:    w.Memo,
	}
	mainTrace := mixin.UniqueConversationID(w.OrderID, constants.WithdrawMainInit)
	feeTrace := mixin.UniqueConversationID(w.OrderID, constants.WithdrawFeeInit)
	if mainTrace == s.TraceID {
		wf.MainTrace = mainTrace
	}
	if feeTrace == s.TraceID {
		wf.FeeTrace = feeTrace
	}
	if old == nil {
		return wb.store.WriteWithdrawalFull(&wf)
	}
	if (old.MainTrace != "" && wf.FeeTrace != "") || (wf.MainTrace != "" && old.FeeTrace != "") {
		err = wb.withdraw(ctx, w)
		if err != nil {
			return fmt.Errorf("wb.withdraw(ctx, %+v) => %v", w, err)
		}
	}
	return nil
}

func (wb *WithbotWorker) processSnapshots(ctx context.Context) {
	snapshots, err := wb.store.ListSnapshots(100)
	if err != nil {
		logger.Errorf("wb.store.ListSnapshots(100) => %v", err)
	}

	for _, s := range snapshots {
		err = wb.process(ctx, s)
		if err != nil {
			logger.Errorf("wb.process(%+v) => %v", s, err)
			continue
		}
	}

	if len(snapshots) < 100 {
		time.Sleep(time.Millisecond * 200)
	}
}

func (wb *WithbotWorker) Run(ctx context.Context) {
	go func() {
		for {
			wb.loopSnapshots(ctx)
		}
	}()

	go func() {
		for {
			wb.processSnapshots(ctx)
		}
	}()
}
