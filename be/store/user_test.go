package store

import (
	"context"
	"fmt"
	"os/user"
	"testing"
)

func TestListUsers(t *testing.T) {
	ctx := context.Background()
	usr, _ := user.Current()
	fmt.Println("Opening badger:", usr.HomeDir)
	db, err := OpenBadger(ctx, usr.HomeDir+"/.mvg/bridge")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Opened badger")
	defer db.Close()

	users, err := db.ListUsers(100)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("users:", users)
}
