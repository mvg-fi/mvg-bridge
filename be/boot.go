package main

import (
	"context"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/api"
	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/store"
	"github.com/mvg-fi/mvg-bridge/users"
	"github.com/mvg-fi/mvg-bridge/workers"
	"github.com/urfave/cli/v2"
)

func bootCmd(c *cli.Context) error {
	logger.SetLevel(logger.VERBOSE)
	ctx := context.Background()

	cp := c.String("config")
	if strings.HasPrefix(cp, "~/") {
		usr, _ := user.Current()
		cp = filepath.Join(usr.HomeDir, (cp)[2:])
	}
	conf, err := config.ReadConfiguration(cp)
	if err != nil {
		return err
	}

	bp := c.String("dir")
	if strings.HasPrefix(bp, "~/") {
		usr, _ := user.Current()
		bp = filepath.Join(usr.HomeDir, (bp)[2:])
	}
	db, err := store.OpenBadger(ctx, bp)
	if err != nil {
		return err
	}
	defer db.Close()

	group, err := mtg.BuildGroup(ctx, db, conf.MTG)
	if err != nil {
		return err
	}

	/*
		s := &mixin.Keystore{
			ClientID:   conf.Messenger.UserId,
			SessionID:  conf.Messenger.SessionId,
			PrivateKey: conf.Messenger.Key,
		}
		mixin, err := mixin.NewFromKeystore(s)
		if err != nil {
			return err
		}
		messenger, err := messenger.NewMixinMessenger(ctx, conf.Messenger)
		if err != nil {
			return err
		}
		println(mixin, messenger)
	*/

	// group.SetOutputGrouper(machine.OutputGrouper)
	// group.AddWorker()
	group.Run(ctx)

	return nil
}

func runCmd(c *cli.Context) error {
	return nil
}

// Only run proxy users and API
func runProxy(c *cli.Context) error {
	logger.SetLevel(logger.VERBOSE)
	ctx := context.Background()
	cp := c.String("config")
	if strings.HasPrefix(cp, "~/") {
		usr, _ := user.Current()
		cp = filepath.Join(usr.HomeDir, (cp)[2:])
	}
	bp := c.String("dir")
	if strings.HasPrefix(bp, "~/") {
		usr, _ := user.Current()
		bp = filepath.Join(usr.HomeDir, (bp)[2:])
	}

	conf, err := config.ReadConfiguration(cp)
	if err != nil {
		return err
	}

	db, err := store.OpenBadger(ctx, bp)
	if err != nil {
		return err
	}
	defer db.Close()

	p := users.NewProxy(ctx, conf)

	api := api.NewAPIWorker(p, db, conf)
	go api.Run(ctx)

	// Run deposit worker
	dw := workers.NewDepositWorker(p, db, conf)
	dw.Run(ctx)

	// Run withbot worker
	// TODO: replace params with real ones
	wb := workers.NewWithbotWorker(ctx, &mixin.Keystore{}, "", db)
	wb.Run(ctx)
	return nil
}
