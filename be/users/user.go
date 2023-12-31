package users

import (
	"context"
	"strings"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/encoding"
	"github.com/mvg-fi/mvg-bridge/store"
)

type User struct {
	constants.User
}

func (u *User) Handle(ctx context.Context, store *store.BadgerStore, conf *config.Configuration, s *mixin.Snapshot) error {
	orderID, err := store.LockGet(s.UserID, s.AssetID)
	if err != nil {
		logger.Errorf("store.LockGet(%s, %s) => %v", s.UserID, s.AssetID, err)
	}
	order, err := store.ReadOrder(orderID)
	if err != nil {
		logger.Errorf("store.ReadOrder(%s) => %v", orderID, err)
	}

	// Change order state to swap
	order.Status = constants.StatusReceived
	err = store.UpdateOrder(orderID, order)
	if err != nil {
		logger.Errorf("store.UpdateOrder(%s) => %v", orderID, err)
	}

	// Transfer to MTG
	input := &mixin.TransferInput{
		AssetID: s.AssetID,
		Amount:  s.Amount,
		TraceID: mixin.UniqueConversationID(s.SnapshotID, "HANDLE||TRANSFER"),
	}
	input.OpponentMultisig.Receivers = conf.MTG.Genesis.Members
	input.OpponentMultisig.Threshold = uint8(conf.MTG.Genesis.Threshold)
	input.Memo = encoding.MsgpackCompressOrder(order)

	user, err := store.ReadUser(s.UserID)
	if err != nil {
		logger.Errorf("store.ReadUser(%s) => %v", s.UserID, err)
	}
	if user == nil {
		logger.Errorf("store.ReadUser(%s) => User doesn't exist", s.UserID)
		return nil
	}

	logger.Verbosef("User.handle(%v)", *s)
	u.SendUntilSufficient(ctx, input)

	// Remove lock
	err = store.LockRemove(s.UserID, s.AssetID)
	if err != nil {
		logger.Errorf("store.LockRemove(%s, %s) => %v", s.UserID, s.AssetID, err)
	}
	return store.RemoveOrder(orderID)
}

func (u *User) Send(ctx context.Context, in *mixin.TransferInput) error {
	uc, err := mixin.NewFromKeystore(u.Key)
	if err != nil {
		panic(err)
	}
	if len(in.OpponentMultisig.Receivers) > 0 {
		_, err = uc.Transaction(ctx, in, u.PIN)
	} else {
		_, err = uc.Transfer(ctx, in, u.PIN)
	}
	logger.Verbosef("User.send(%v) => %v", *in, err)
	return err
}

func (u *User) SendUntilSufficient(ctx context.Context, in *mixin.TransferInput) error {
	for {
		err := u.Send(ctx, in)
		if mixin.IsErrorCodes(err, 30103) {
			time.Sleep(3 * time.Second)
			continue
		}
		if err != nil && strings.Contains(err.Error(), "Client.Timeout exceeded") {
			time.Sleep(1 * time.Second)
			continue
		}
		return err
	}
}

func (u *User) GetDepositAddr(ctx context.Context, assetID string) *mixin.DepositEntry {
	client, err := mixin.NewFromKeystore(u.Key)
	if err != nil {
		logger.Errorf("mixin.NewFromKeystore(%v) => %v", u.Key, err)
	}

	ast, err := client.ReadAsset(ctx, assetID)
	if err != nil {
		logger.Errorf("u.ReadAsset() => %v", err)
	}
	return &ast.DepositEntries[0]
}
