package store

import (
	"context"
	"encoding/binary"

	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/mvg-fi/common/store"
)

type BadgerStore struct {
	store.BadgerStore
}

func OpenBadger(ctx context.Context, path string) (*BadgerStore, error) {
	bs, err := store.OpenBadger(ctx, path)
	if err != nil {
		return nil, err
	}
	return &BadgerStore{*bs}, err
}

func uint64Bytes(i uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

func (bs *BadgerStore) ListIterations() ([]*mtg.Iteration, error) {
	return nil, nil
}
func (bs *BadgerStore) WriteIteration(ir *mtg.Iteration) error {
	return nil
}
