package store

import (
	"encoding/binary"
	"time"

	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/dgraph-io/badger/v4"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/mvg-bridge/constants"
)

func (s *BadgerStore) ListSnapshots(limit int) ([]*mixin.Snapshot, error) {
	snapshots := make([]*mixin.Snapshot, 0)
	txn := s.Badger().NewTransaction(false)
	defer txn.Discard()

	opts := badger.DefaultIteratorOptions
	opts.Prefix = []byte(constants.PrefixSnapshotList)
	it := txn.NewIterator(opts)
	defer it.Close()

	it.Seek(opts.Prefix)
	for ; it.Valid() && len(snapshots) < limit; it.Next() {
		item := it.Item()
		v, err := item.ValueCopy(nil)
		if err != nil {
			return snapshots, err
		}
		var snap mixin.Snapshot
		err = mtg.DecompressMsgpackUnmarshal(v, &snap)
		if err != nil {
			return snapshots, err
		}
		snapshots = append(snapshots, &snap)
	}

	return snapshots, nil
}

func (s *BadgerStore) DeleteSnapshots(snaps []*mixin.Snapshot) error {
	return s.Badger().Update(func(txn *badger.Txn) error {
		for _, snap := range snaps {
			key := snapshotKey(snap)
			err := txn.Delete(key)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

/*
func (s *BadgerStore) ListWithdrawals(limit int) ([]*Withdrawal, error) {
	withdrawals := make([]*Withdrawal, 0)
	txn := s.NewTransaction(false)
	defer txn.Discard()

	opts := badger.DefaultIteratorOptions
	opts.Prefix = []byte(storePrefixWithdrawalPair)
	it := txn.NewIterator(opts)
	defer it.Close()

	it.Seek(opts.Prefix)
	for ; it.Valid() && len(withdrawals) < limit; it.Next() {
		item := it.Item()
		v, err := item.ValueCopy(nil)
		if err != nil {
			return withdrawals, err
		}
		var w Withdrawal
		err = mtg.DecompressMsgpackUnmarshal(v, &w)
		if err != nil {
			return withdrawals, err
		}
		withdrawals = append(withdrawals, &w)
	}

	return withdrawals, nil
}

func (s *BadgerStore) DeleteWitdrawals(withdrawals []*Withdrawal) error {
	return s.Badger().Update(func(txn *badger.Txn) error {
		for _, w := range withdrawals {
			key := []byte(storePrefixWithdrawalPair + w.TraceId)
			err := txn.Delete(key)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *BadgerStore) WriteWithdrawal(w *Withdrawal) error {
	logger.Verbosef("BadgerStore.writeWithdrawal(%v)", *w)
	return s.Badger().Update(func(txn *badger.Txn) error {
		old, err := s.readWithdrawal(txn, w.TraceId)
		if err != nil {
			return err
		}
		if old != nil && w.Asset == nil {
			panic(old.Asset)
		}
		if old != nil && w.Fee == nil {
			panic(old.Fee)
		}
		if old != nil && old.CreatedAt != w.CreatedAt {
			panic(old.CreatedAt)
		}
		if old != nil && old.UserId != w.UserId {
			panic(old.UserId)
		}
		key := []byte(storePrefixWithdrawalPair + w.TraceId)
		val := mtg.CompressMsgpackMarshalPanic(w)
		return txn.Set(key, val)
	})
}

func (s *BadgerStore) ReadWithdrawalById(id string) (*Withdrawal, error) {
	txn := s.NewTransaction(false)
	defer txn.Discard()

	return s.readWithdrawal(txn, id)
}

func (s *BadgerStore) ReadWithdrawal(txn *badger.Txn, id string) (*Withdrawal, error) {
	key := []byte(storePrefixWithdrawalPair + id)
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
	var w Withdrawal
	err = mtg.DecompressMsgpackUnmarshal(val, &w)
	return &w, err
}
*/

func snapshotKey(s *mixin.Snapshot) []byte {
	key := []byte(constants.PrefixSnapshotList)
	buf := timeToBytes(s.CreatedAt)
	key = append(key, buf...)
	key = append(key, s.SnapshotID...)
	return key
}
func timeToBytes(t time.Time) []byte {
	buf := make([]byte, 8)
	now := uint64(t.UnixNano())
	binary.BigEndian.PutUint64(buf, now)
	return buf
}
