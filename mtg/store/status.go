package store

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

func (bs *BadgerStore) ReadStatus(status, id string) ([]byte, error) {
	txn := bs.Badger().NewTransaction(false)
	defer txn.Discard()

	key := []byte(fmt.Sprintf("%s|%s", status, id))
	item, err := txn.Get(key)
	if err == badger.ErrKeyNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return item.ValueCopy(nil)
}

/*
func (bs *BadgerStore) ReadStatusBySuffix(id string) ([]byte, error) {
	txn := bs.Badger().NewTransaction(false)
	defer txn.Discard()
	return
}
*/

// Remove old prefix item, create a new item
func (bs *BadgerStore) FinishStatus(originalPrefix, newPrefix, id string) (string, error) {
	txn := bs.Badger().NewTransaction(true)
	defer txn.Discard()

	key := []byte(fmt.Sprintf("%s|%s", originalPrefix, id))
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
	newKey := []byte(fmt.Sprintf("%s|%s", newPrefix, id))
	txn.Set(newKey, value)
	return string(newKey), nil
}
