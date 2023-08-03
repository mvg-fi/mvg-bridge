package store

import (
	"github.com/MixinNetwork/trusted-group/mvm/encoding"
	"github.com/dgraph-io/badger"
)

const (
	prefixDeposit        = "MVG:DEPOSIT:"
	prefixDepositPending = "MVG:DEPOSIT:PENDING:"
)

func (bs *BadgerStore) writeDeposit(txn *badger.Txn, id string, event *encoding.Event) error {
	buf := encoding.JSONMarshalPanic(event)
	key := append([]byte(prefixDeposit), id...)
	return txn.Set(key, buf)
}

func (bs *BadgerStore) readDeposit(txn *badger.Txn, id string) ([]byte, error) {
	key := append([]byte(prefixDeposit), id...)
	item, err := txn.Get(key)
	if err == badger.ErrKeyNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return item.ValueCopy(nil)
}
