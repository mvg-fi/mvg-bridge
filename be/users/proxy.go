package users

import (
	"context"
	"crypto/ed25519"

	"github.com/MixinNetwork/mixin/crypto"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/common/time"
	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/store"
)

type Proxy struct {
	*mixin.Client
	config *config.Configuration
}

func NewProxy(ctx context.Context, conf *config.Configuration) *Proxy {
	client, err := mixin.NewFromKeystore(&mixin.Keystore{
		ClientID:   conf.ProxyRoot.ClientID,
		SessionID:  conf.ProxyRoot.SessionID,
		PrivateKey: conf.ProxyRoot.PrivateKey,
		PinToken:   conf.ProxyRoot.PinToken,
		Scope:      conf.ProxyRoot.Scope,
	})
	if err != nil {
		panic(err)
	}
	return &Proxy{
		client,
		conf,
	}
}

func (p *Proxy) NewUser(ctx context.Context, store *store.BadgerStore) (*constants.User, error) {
	now := time.UTCNow()
	seed := crypto.NewHash([]byte(p.config.Proxy.ProxyUserSecret + now))
	signer := ed25519.NewKeyFromSeed(seed[:])
	u, ks, err := p.CreateUser(ctx, signer, now)
	if err != nil {
		return nil, err
	}

	user := &constants.User{u, ks, p.config.Proxy.ProxyPIN}

	if !user.HasPin {
		uc, err := mixin.NewFromKeystore(ks)
		if err != nil {
			return nil, err
		}
		err = uc.ModifyPin(ctx, "", p.config.Proxy.ProxyPIN)
		if err != nil {
			return nil, err
		}
		user.HasPin = true
	}

	err = store.WriteUser(user)
	if err != nil {
		return nil, err
	}
	return store.ReadUser(user.UserID)
}

func (p *Proxy) GetRandomDepositUser(ctx context.Context, store *store.BadgerStore, assetID string) *User {
	users, err := store.ListUsers(100)
	if err != nil {
		logger.Errorf("proxy.GetRandomDepositUser() => %v", err)
	}

	for i := 0; i < len(users); i++ {
		locked, err := store.LockCheck(users[i].UserID, assetID)
		if err != nil {
			logger.Errorf("store.LockCheck(users[i], assetID) => %v", err)
			continue
		}
		if !locked {
			return &User{*users[i]}
		}
	}
	u, err := p.NewUser(ctx, store)
	if err != nil {
		logger.Errorf("proxy.NewUser() => %v", err)
	}
	err = store.WriteUser(u)
	if err != nil {
		logger.Errorf("store.Write(u) => %v", err)
	}
	return &User{*u}
}

func (p *Proxy) GetADeposit(ctx context.Context, store *store.BadgerStore, assetID, orderID string) *mixin.DepositEntry {
	user := p.GetRandomDepositUser(ctx, store, assetID)
	err := store.LockSet(user.UserID, assetID, orderID)
	if err != nil {
		logger.Errorf("store.LockSet(user.UserID, assetID) => %v", err)
	}
	return user.GetDepositAddr(ctx, assetID)
}
