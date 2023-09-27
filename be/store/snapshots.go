package store

import (
	"context"
	"encoding/binary"
	"time"

	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/dgraph-io/badger/v4"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/constants"
)

func (bs *BadgerStore) ReadSnapshotsCheckpoint(ctx context.Context) (time.Time, error) {
	key := []byte(constants.PrefixSnapshotCheckpoint)
	return bs.ReadCheckpoint(ctx, key)
}

func (bs *BadgerStore) WriteSnapshotsCheckpoint(ctx context.Context, ckpt time.Time) error {
	key := []byte(constants.PrefixSnapshotCheckpoint)
	return bs.WriteCheckpoint(ctx, key, ckpt)
}

func (bs *BadgerStore) ReadCheckpoint(ctx context.Context, key []byte) (time.Time, error) {
	txn := bs.Badger().NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get(key)
	if err == badger.ErrKeyNotFound {
		return time.Now(), nil
	} else if err != nil {
		return time.Time{}, err
	}
	val, err := item.ValueCopy(nil)
	if err != nil {
		return time.Time{}, err
	}
	ckpt := binary.BigEndian.Uint64(val)
	return time.Unix(0, int64(ckpt)), nil
}

func (bs *BadgerStore) WriteCheckpoint(ctx context.Context, key []byte, ckpt time.Time) error {
	return bs.Badger().Update(func(txn *badger.Txn) error {
		val := timeToBytes(ckpt)
		return txn.Set(key, val)
	})
}

func (bs *BadgerStore) WriteSnapshot(snap *mixin.Snapshot) error {
	logger.Verbosef("BadgerStore.WriteSnapshot(%v)", *snap)
	return bs.Badger().Update(func(txn *badger.Txn) error {
		key := snapshotKey(snap)
		val := mtg.CompressMsgpackMarshalPanic(snap)
		return txn.Set(key, val)
	})
}

func (bs *BadgerStore) ListSnapshots(limit int) ([]*mixin.Snapshot, error) {
	snapshots := make([]*mixin.Snapshot, 0)
	txn := bs.Badger().NewTransaction(false)
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
