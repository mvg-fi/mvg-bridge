package api

import (
	"context"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/store"
)

func TestRun(t *testing.T) {
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
	api := NewAPIWorker(db, conf.API)
	api.Run()
}

// Test active
// curl -i -X POST -H "Content-Type: application/json" --data '{}' http://127.0.0.1:8000/api/status
// Test rate limit
