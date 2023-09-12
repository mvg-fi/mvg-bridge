package users

import (
	"context"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/constants"
)

type User struct {
	constants.User
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

func (u *User) GetDepositAddr(ctx context.Context, assetID string) *mixin.DepositEntry {
	client, err := mixin.NewFromKeystore(u.Key)
	if err != nil {
		logger.Errorf("mixin.NewFromKeystore(%v) => %v", u.Key, err)
	}

	ast, err := client.ReadAsset(ctx, assetID)
	if err != nil {
		logger.Errorf("u.ReadAsset() =>", err)
	}
	return &ast.DepositEntries[0]
}
