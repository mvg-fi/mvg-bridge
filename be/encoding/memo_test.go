package encoding

import (
	"fmt"
	"testing"

	"github.com/mvg-fi/mvg-bridge/constants"
)

var (
	o = &constants.Order{
		FromAssetID: "b106b13f-ab3a-4a7c-bfe1-121e8b177034",
		ToAssetID:   "3afba112-c2ce-4d3c-a6a7-9d9b1fecf32f",
		Amount:      "0.12342567",
		Except:      "",
		Cex:         true,
		TraceID:     "73cd89ac-9610-49d8-8a67-0be375335c86",
		Address:     "0xE6954ef09cDf7ec6526b29a5978ca502F2E50C58",
		Memo:        "12345657890124567890123456789012",
		Expire:      "2012-01-01T18:23:32Z",
		Status:      "Recevied",
	}
	sm = &constants.Swap{
		FromAssetID: "b106b13f-ab3a-4a7c-bfe1-121e8b177034",
		ToAssetID:   "3afba112-c2ce-4d3c-a6a7-9d9b1fecf32f",
		Amount:      "0.12342567",
		Receive:     "",
		OrderID:     "73cd89ac-9610-49d8-8a67-0be375335c86",
		TraceID:     "05bcb6ef-cb50-4ced-9be4-64c0fa421d4a",
		Status:      "SwapSent",
		Provider:    "PandoSwap",
		Type:        "Main",
	}
	sf = &constants.Swap{
		FromAssetID: "b106b13f-ab3a-4a7c-bfe1-121e8b177034",
		ToAssetID:   "3afba112-c2ce-4d3c-a6a7-9d9b1fecf32f",
		Amount:      "0.12342567",
		Receive:     "",
		OrderID:     "73cd89ac-9610-49d8-8a67-0be375335c86",
		TraceID:     "749b3e60-a21b-43fd-836b-4f680a5095be",
		Status:      "SwapSent",
		Provider:    "Exinone",
		Type:        "fee",
	}
)

func TestMsgpackCompressOrder(t *testing.T) {
	memo := MsgpackCompressOrder(o)

	println(memo)
	println(len(memo), "\n")
}

func TestMsgpackOrder(t *testing.T) {
	memo := MsgpackMashalOrder(o)

	println(memo)
	println(len(memo), "\n")
}

func TestMsgpackDecompressOrder(t *testing.T) {
	memo := MsgpackCompressOrder(o)
	o := MsgpackDecompressOrder(memo)
	fmt.Printf("%+v", o)
}

func TestGetWithdrawalMemo(t *testing.T) {
	mm, mf := GetWithdrawalMemo(o, sm, sf)
	println(mm)
	println(len(mm))
	println(mf)
	println(len(mf))
}

func TestRetriveWithdrawalMemo(t *testing.T) {
	fmt.Printf("%+v\n", o)
	mm, mf := GetWithdrawalMemo(o, sm, sf)
	println(mm)
	println(mf)
	wm, err := RetriveWithdrawalMemo(mm)
	if err != nil {
		panic(err)
	}
	wf, err := RetriveWithdrawalMemo(mf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", wm)
	fmt.Printf("%+v\n", wf)
}
