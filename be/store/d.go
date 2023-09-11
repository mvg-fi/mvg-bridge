package store

const (
// KeyExample           = "{STATUS}|{TRACE UUID}"
// ValueExample         = "{TYPE}|{ADDRESS GENERATED}|{FROM ASSETID}|{TO ASSETID}|{TO ADDRESS}|{MEMO}|{AMOUNT}|{TIMESTAMP}"
)

/*
func (bs *BadgerStore) WriteDeposit(d *constants.Deposit) (string, error) {
	if d.Type != "CS" && d.Type != "SS" && d.Type != "BE" {
		return "", fmt.Errorf("Invalid deposit type")
	}
	txn := bs.Badger().NewTransaction(true)
	defer txn.Discard()

	traceID := uuid.NewV4()
	key := []byte(fmt.Sprintf("%s|%s", constants.PrefixDepositPending, traceID))
	value := []byte(fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s", d.Type, d.Address, d.FromAssetID, d.ToAssetID, d.ToAddress, d.Memo, d.Amount, time.UTCP8Now()))
	err := txn.Set(key, value)
	if err != nil {
		return "", fmt.Errorf("txn.Set(key)=>%w", err)
	}
	return traceID, nil
}
*/
