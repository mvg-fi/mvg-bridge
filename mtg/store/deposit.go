package store

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"github.com/mvg-fi/common/time"
	"github.com/mvg-fi/common/uuid"
	"github.com/mvg-fi/mvg-bridge/constants"
)

const (
// KeyExample           = "{STATUS}|{TRACE UUID}"
// ValueExample         = "{TYPE}|{ADDRESS GENERATED}|{FROM ASSETID}|{TO ASSETID}|{TO ADDRESS}|{MEMO}|{AMOUNT}|{TIMESTAMP}"
)

type Deposit struct {
	Type        string `json:"type"`         // Transaction type (CS|SS|BE)
	Address     string `json:"address"`      // Deposit address for receiving user's fund
	ToAddress   string `json:"to_address"`   // The destination
	FromAssetID string `json:"from_assetid"` // Deposit asset id
	ToAssetID   string `json:"to_assetid"`   // Withdrawal asset id
	Memo        string `json:"memo"`         // Memo
	Amount      string `json:"amount"`       // User final receive amount of to asset
}

func (bs *BadgerStore) WriteDeposit(txn *badger.Txn, d *Deposit) (string, error) {
	if d.Type != "CS" && d.Type != "SS" && d.Type != "BE" {
		return "", fmt.Errorf("Invalid deposit type")
	}
	traceID := uuid.NewV4()
	key := []byte(fmt.Sprintf("%s|%s", constants.PrefixDepositPending, traceID))
	value := []byte(fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s", d.Type, d.Address, d.FromAssetID, d.ToAssetID, d.ToAddress, d.Memo, d.Amount, time.UTCP8Now()))
	err := txn.Set(key, value)
	if err != nil {
		return "", fmt.Errorf("txn.Set(key)=>%w", err)
	}
	return traceID, nil
}
