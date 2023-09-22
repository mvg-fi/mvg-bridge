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
