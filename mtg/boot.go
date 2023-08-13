package main

import (
	"context"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/api"
	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/store"
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

	// TODO: Figure out what this step is used for
	// group.SetOutputGrouper(machine.OutputGrouper)
	// group.AddWorker()
	api := api.NewAPIWorker(conf.API)
	api.Loop()
	group.Run(ctx)

	return nil
}
