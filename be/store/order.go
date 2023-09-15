package store

import (
	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/dgraph-io/badger/v4"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/constants"
)

func (bs *BadgerStore) WriteOrder(o *constants.Order) error {
	logger.Verbosef("BadgerStore.WriteOrder(%v)", o)
	return bs.Badger().Update(func(txn *badger.Txn) error {
		key := []byte(constants.PrefixOrder + o.TraceID)
		val := mtg.MsgpackMarshalPanic(o)
		return txn.Set(key, val)
	})
}

func (bs *BadgerStore) ReadOrder(traceID string) (*constants.Order, error) {
	txn := bs.Badger().NewTransaction(false)
	defer txn.Discard()

	key := []byte(constants.PrefixOrder + traceID)
	item, err := txn.Get(key)
	if err == badger.ErrKeyNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	val, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}
	var order constants.Order
	err = mtg.MsgpackUnmarshal(val, &order)
	return &order, err
}

func (bs *BadgerStore) ListOrders(limit int) ([]*constants.Order, error) {
	orders := make([]*constants.Order, 0)
	txn := bs.Badger().NewTransaction(false)
	defer txn.Discard()

	opts := badger.DefaultIteratorOptions
	opts.Prefix = []byte(constants.PrefixOrder)
	it := txn.NewIterator(opts)
	defer it.Close()
	it.Seek(opts.Prefix)
	for ; it.Valid() && len(orders) < limit; it.Next() {
		item := it.Item()
		v, err := item.ValueCopy(nil)
		if err != nil {
			return orders, err
		}
		var order constants.Order
		err = mtg.DecompressMsgpackUnmarshal(v, &order)
		if err != nil {
			return orders, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

func (bs *BadgerStore) UpdateOrder(orderID string, o *constants.Order) error {
	return bs.Badger().Update(func(txn *badger.Txn) error {
		key := []byte(constants.PrefixOrder + orderID)
		val := mtg.MsgpackMarshalPanic(o)
		return txn.Set(key, val)
	})
}

func (bs *BadgerStore) RemoveOrder(traceID string) error {
	return bs.Badger().Update(func(txn *badger.Txn) error {
		key := []byte(constants.PrefixOrder + traceID)
		err := txn.Delete(key)
		if err != nil {
			return err
		}
		return nil
	})
}
