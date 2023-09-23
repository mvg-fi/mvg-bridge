package store

import (
	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/dgraph-io/badger/v4"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/constants"
)

func (bs *BadgerStore) WriteWithdrawal(s *constants.Withdrawal) error {
	logger.Verbosef("BadgerStore.WriteWithdrawal(%v)", s)
	return bs.Badger().Update(func(txn *badger.Txn) error {
		key := []byte(constants.PrefixWithdrawal + s.TraceID)
		val := mtg.MsgpackMarshalPanic(s)
		return txn.Set(key, val)
	})
}

func (bs *BadgerStore) ReadWithdrawal(traceID string) (*constants.Withdrawal, error) {
	txn := bs.Badger().NewTransaction(false)
	defer txn.Discard()

	key := []byte(constants.PrefixWithdrawal + traceID)
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
	var swap constants.Withdrawal
	err = mtg.MsgpackUnmarshal(val, &swap)
	return &swap, err
}

func (bs *BadgerStore) ListWithdrawals(limit int) ([]*constants.Withdrawal, error) {
	swaps := make([]*constants.Withdrawal, 0)
	txn := bs.Badger().NewTransaction(false)
	defer txn.Discard()

	opts := badger.DefaultIteratorOptions
	opts.Prefix = []byte(constants.PrefixWithdrawal)
	it := txn.NewIterator(opts)
	defer it.Close()
	it.Seek(opts.Prefix)
	for ; it.Valid() && len(swaps) < limit; it.Next() {
		item := it.Item()
		v, err := item.ValueCopy(nil)
		if err != nil {
			return swaps, err
		}
		var swap constants.Withdrawal
		err = mtg.DecompressMsgpackUnmarshal(v, &swap)
		if err != nil {
			return swaps, err
		}
		swaps = append(swaps, &swap)
	}
	return swaps, nil
}

func (bs *BadgerStore) UpdateWithdrawal(traceID string, s *constants.Withdrawal) error {
	return bs.Badger().Update(func(txn *badger.Txn) error {
		key := []byte(constants.PrefixWithdrawal + traceID)
		val := mtg.MsgpackMarshalPanic(s)
		return txn.Set(key, val)
	})
}

func (bs *BadgerStore) RemoveWithdrawal(traceID string) error {
	return bs.Badger().Update(func(txn *badger.Txn) error {
		key := []byte(constants.PrefixWithdrawal + traceID)
		err := txn.Delete(key)
		if err != nil {
			return err
		}
		return nil
	})
}
