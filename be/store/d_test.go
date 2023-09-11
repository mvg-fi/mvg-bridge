package store

/*
func TestDeposit(t *testing.T) {
	ctx := context.Background()
	usr, _ := user.Current()
	fmt.Println("Opening badger:", usr.HomeDir)
	db, err := OpenBadger(ctx, usr.HomeDir+"/.mvg/bridge")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Opened badger")
	defer db.Close()

	fmt.Println("Writing deposit")

	traceID, err := db.WriteDeposit(&Deposit{
		Type:        "CS",
		Address:     "0x2CFD0582C121B1a84BA795Dd701f798966605598",
		ToAddress:   "0x1AE60D36412a6745fce4d4935FF5Bf2b8139a371",
		FromAssetID: "1faaf87d-cc0d-4828-9ca6-ae89f239c533",
		ToAssetID:   "47b74129-64eb-4998-ac73-b2cf2f52099c",
		Memo:        "just a test",
		Amount:      "0.01",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("TraceID:", traceID)
	result, err := db.ReadStatus(txn, constants.PrefixDepositPending, traceID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ReadPendingDeposit:", string(result))

	newKey, err := db.FinishStatus(txn, constants.PrefixDepositPending, constants.PrefixDepositReceived, traceID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("NewKey:", newKey)
}
*/
