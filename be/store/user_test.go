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

func TestLockGet(t *testing.T) {
	ctx := context.Background()
	usr, _ := user.Current()
	fmt.Println("Opening badger:", usr.HomeDir)
	db, err := OpenBadger(ctx, usr.HomeDir+"/.mvg/bridge")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Opened badger")
	defer db.Close()

	userID := "1f58f052-f1ab-44ad-9e78-22126780eacd"
	assetID := "3b3c8339-c926-4a0f-817f-748d4e605b85"
	orderID := "be8b1281-5fd8-423f-becc-a12e545021b0"
	err = db.LockSet(userID, assetID, orderID)
	ord, err := db.LockGet(userID, assetID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("orderID:", ord)
}
