package store

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"github.com/mvg-fi/common/time"
	"github.com/mvg-fi/common/uuid"
	"github.com/mvg-fi/mvg-bridge/constants"
)

const (
// KeyExample           = "{STATUS}|{TRACE UUID}"
// ValueExample         = "{TYPE}|{ADDRESS GENERATED}|{FROM ASSETID}|{TO ASSETID}|{TO ADDRESS}|{MEMO}|{AMOUNT}|{TIMESTAMP}"
)

type Deposit struct {
	Type        string // Transaction type (CS|SS|BE)
	Address     string // Deposit address for receiving user's fund
	ToAddress   string // The destination
	FromAssetID string // Deposit asset id
	ToAssetID   string // Withdrawal asset id
	Memo        string // Memo
	Amount      string // User final receive amount of to asset
}

func (bs *BadgerStore) WriteDeposit(txn *badger.Txn, d *Deposit) (string, error) {
	if d.Type != "CS" && d.Type != "SS" && d.Type != "BE" {
		return "", fmt.Errorf("Invalid deposit type")
	}
	traceID := uuid.NewV4()
	key := []byte(fmt.Sprintf("%s|%s", constants.PrefixDepositPending, traceID))
	value := []byte(fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s", d.Type, d.Address, d.FromAssetID, d.ToAssetID, d.ToAddress, d.Memo, d.Amount, time.UTCP8Now()))
	err := txn.Set(key, value)
	if err != nil {
		return "", fmt.Errorf("txn.Set(key)=>%w", err)
	}
	return traceID, nil
}

func (bs *BadgerStore) ReadPendingDeposit(txn *badger.Txn, id string) ([]byte, error) {
	key := []byte(fmt.Sprintf("%s|%s", constants.PrefixDepositPending, id))
	item, err := txn.Get(key)
	if err == badger.ErrKeyNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return item.ValueCopy(nil)
}

func (bs *BadgerStore) FinishPendingDeposit(txn *badger.Txn, id string) (string, error) {
	// Delete pending, create a new tx
	key := []byte(fmt.Sprintf("%s|%s", constants.PrefixDepositPending, id))
	item, err := txn.Get(key)
	if err != nil {
		return "", fmt.Errorf("txn.Get(key)=>%w", err)
	}
	value, err := item.ValueCopy(nil)
	if err != nil {
		return "", fmt.Errorf("item.ValueCopy(nil)=>%w", err)
	}
	err = txn.Delete(key)
	if err != nil {
		return "", fmt.Errorf("txn.Delete(key)=>%w", err)
	}
	newKey := []byte(fmt.Sprintf("%s|%s", constants.PrefixDepositReceived, id))
	txn.Set(newKey, value)
	return string(newKey), nil
}
