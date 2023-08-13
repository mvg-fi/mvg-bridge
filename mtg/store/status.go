package store

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

func (bs *BadgerStore) ReadStatus(txn *badger.Txn, status, id string) ([]byte, error) {
	key := []byte(fmt.Sprintf("%s|%s", status, id))
	item, err := txn.Get(key)
	if err == badger.ErrKeyNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return item.ValueCopy(nil)
}

// Remove old prefix item, create a new item
func (bs *BadgerStore) FinishStatus(txn *badger.Txn, originalPrefix, newPrefix, id string) (string, error) {
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
