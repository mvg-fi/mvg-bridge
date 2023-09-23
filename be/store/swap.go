package store

import (
	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/dgraph-io/badger/v4"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/constants"
)

func (bs *BadgerStore) WriteSwap(s *constants.Swap) error {
	logger.Verbosef("BadgerStore.WriteSwap(%v)", s)
	return bs.Badger().Update(func(txn *badger.Txn) error {
		key := []byte(constants.PrefixSwap + s.TraceID)
		val := mtg.MsgpackMarshalPanic(s)
		return txn.Set(key, val)
	})
}

func (bs *BadgerStore) ReadSwap(traceID string) (*constants.Swap, error) {
	txn := bs.Badger().NewTransaction(false)
	defer txn.Discard()

	key := []byte(constants.PrefixSwap + traceID)
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
	var swap constants.Swap
	err = mtg.MsgpackUnmarshal(val, &swap)
	return &swap, err
}

func (bs *BadgerStore) ListSwaps(limit int) ([]*constants.Swap, error) {
	swaps := make([]*constants.Swap, 0)
	txn := bs.Badger().NewTransaction(false)
	defer txn.Discard()

	opts := badger.DefaultIteratorOptions
	opts.Prefix = []byte(constants.PrefixSwap)
	it := txn.NewIterator(opts)
	defer it.Close()
	it.Seek(opts.Prefix)
	for ; it.Valid() && len(swaps) < limit; it.Next() {
		item := it.Item()
		v, err := item.ValueCopy(nil)
		if err != nil {
			return swaps, err
		}
		var swap constants.Swap
		err = mtg.DecompressMsgpackUnmarshal(v, &swap)
		if err != nil {
			return swaps, err
		}
		swaps = append(swaps, &swap)
	}
	return swaps, nil
}

func (bs *BadgerStore) UpdateSwap(traceID string, s *constants.Swap) error {
	return bs.Badger().Update(func(txn *badger.Txn) error {
		key := []byte(constants.PrefixSwap + traceID)
		val := mtg.MsgpackMarshalPanic(s)
		return txn.Set(key, val)
	})
}

func (bs *BadgerStore) RemoveSwap(traceID string) error {
	return bs.Badger().Update(func(txn *badger.Txn) error {
		key := []byte(constants.PrefixSwap + traceID)
		err := txn.Delete(key)
		if err != nil {
			return err
		}
		return nil
	})
}
