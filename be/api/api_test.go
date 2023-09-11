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

// Test Price Simple
// curl -i -X POST -H "Content-Type: application/json" --data '{"from_asset_id": "c6d0c728-2624-429b-8e0d-d9d19b6592fa", "to_asset_id":"43d61dcd-e413-450d-80b8-101d5e903357", "amount": "1", "except": "", "cex": true}' http://127.0.0.1:8000/price/simple

// Test Price All
// curl -i -X POST -H "Content-Type: application/json" --data '{"from_asset_id": "c6d0c728-2624-429b-8e0d-d9d19b6592fa", "to_asset_id":"43d61dcd-e413-450d-80b8-101d5e903357", "amount": "1", "except": "", "cex": true}' http://127.0.0.1:8000/price/all

// Test Create Order

// Test Rate limit
// for i in {0..61}
// do
// curl -X POST 127.0.0.1:8000/address
// done
