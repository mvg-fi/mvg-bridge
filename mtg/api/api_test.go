package api

import (
	"os/user"
	"path/filepath"
	"testing"

	"github.com/mvg-fi/mvg-bridge/config"
	"github.com/mvg-fi/mvg-bridge/store"
)

func TestRun(t *testing.T) {
	usr, _ := user.Current()
	cp = filepath.Join(usr.HomeDir, (cp)[2:])
	conf, err := config.ReadConfiguration(cp)
	if err != nil {
		return err
	}
	db, err := store.OpenBadger(ctx, bp)
	if err != nil {
		return err
	}
	defer db.Close()
	api := api.NewAPIWorker(db, conf.API)
	api.Run()
}

// Test active
// curl -i -X POST -H "Content-Type: application/json" --data '{}' http://127.0.0.1:8000/api/status
// Test rate limit
//
