package workers

import (
	"context"
	"fmt"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/mvg-fi/mvg-bridge/store"
	"github.com/mvg-fi/mvg-bridge/users"
)

func TestProcess(t *testing.T) {
	ctx := context.Background()
	usr, _ := user.Current()
	cp := "../config/config1.toml"
	bp := filepath.Join(usr.HomeDir, "/mvg/bridge")
	conf, err := config.ReadConfiguration(cp)
	if err != nil {
		panic(err)
	}
	db, err := store.OpenBadger(ctx, bp)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	group, err := mtg.BuildGroup(ctx, db, conf.MTG)
	if err != nil {
		panic(err)
	}

	proxy := users.NewProxy(ctx, conf)
	dw := NewDepositWorker(proxy, db, conf)
	go dw.Run(ctx)

	sw := NewSwapWorker(group, db, conf)
	sw.Run(ctx)
}

func TestUniqueConversationID(t *testing.T) {
	a := "4cd5f7a2-c80b-455e-be05-7f37843de2f4"
	b := constants.SwapTypeFeeInit
	c := mixin.UniqueConversationID(a, b)
	fmt.Printf("a:%s\nb:%s\nc:%s\n", a, b, c)

	fmt.Printf("b:%s\n", mixin.UniqueConversationID(a, c))
}
