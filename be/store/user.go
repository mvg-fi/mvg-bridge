package store

import (
	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/dgraph-io/badger/v4"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/constants"
)

func (bs *BadgerStore) WriteUser(u *constants.User) error {
	logger.Verbosef("BadgerStore.WriteUser(%s)", u.UserID)
	return bs.Badger().Update(func(txn *badger.Txn) error {
		key := []byte(constants.PrefixUser + u.UserID)
		val := mtg.MsgpackMarshalPanic(u)
		return txn.Set(key, val)
	})
}

func (bs *BadgerStore) ReadUser(userID string) (*constants.User, error) {
	txn := bs.Badger().NewTransaction(false)
	defer txn.Discard()

	key := []byte(constants.PrefixUser + userID)
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
	var user constants.User
	err = mtg.MsgpackUnmarshal(val, &user)
	return &user, err
}

func (bs *BadgerStore) ListUsers(limit int) ([]*constants.User, error) {
	users := make([]*constants.User, 0)
	txn := bs.Badger().NewTransaction(false)
	defer txn.Discard()

	opts := badger.DefaultIteratorOptions
	opts.Prefix = []byte(constants.PrefixUser)
	it := txn.NewIterator(opts)
	defer it.Close()
	it.Seek(opts.Prefix)
	for ; it.Valid() && len(users) < limit; it.Next() {
		item := it.Item()
		v, err := item.ValueCopy(nil)
		if err != nil {
			return users, err
		}
		var user constants.User
		err = mtg.DecompressMsgpackUnmarshal(v, &user)
		if err != nil {
			return users, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (bs *BadgerStore) LockSet(userID, assetID, orderID string) error {
	return bs.Badger().Update(func(txn *badger.Txn) error {
		key := []byte(constants.PrefixLock + userID + ":" + assetID)
		val := mtg.MsgpackMarshalPanic(orderID)
		return txn.Set(key, val)
	})
}

func (bs *BadgerStore) LockCheck(userID, assetID string) (bool, error) {
	txn := bs.Badger().NewTransaction(false)
	defer txn.Discard()

	key := []byte(constants.PrefixLock + userID + ":" + assetID)
	_, err := txn.Get(key)
	if err == badger.ErrKeyNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (bs *BadgerStore) LockGet(userID, assetID string) (string, error) {
	txn := bs.Badger().NewTransaction(false)
	defer txn.Discard()

	key := []byte(constants.PrefixLock + userID + ":" + assetID)
	value, err := txn.Get(key)
	if err != nil {
		return "", err
	}
	val, err := value.ValueCopy(nil)
	if err != nil {
		return "", err
	}
	var orderID string
	err = mtg.MsgpackUnmarshal(val, &orderID)
	if err != nil {
		return "", err
	}

	return string(orderID), nil
}

func (bs *BadgerStore) LockRemove(userID, assetID string) error {
	return bs.Badger().Update(func(txn *badger.Txn) error {
		key := []byte(constants.PrefixLock + userID + ":" + assetID)
		err := txn.Delete(key)
		if err != nil {
			return err
		}
		return nil
	})
}
