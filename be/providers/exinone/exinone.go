package exinone

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/constants"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

const (
	NAME     = "Exinone"
	CLIENTID = "61103d28-3ac2-44a2-ae34-bd956070dab1"
	REFCODE  = "m.7000104232"
	EPCUUID  = "7aee3b93-742b-324c-a59e-46a8427435f9"

	Endpoint      = "https://app.eiduwejdk.com/api/v2"
	GetPricePath  = "/convert/estimate/amount"
	GetStatusPath = "/convert/order/detail"
)

func GetPrice(payAsset, receiveAsset, amount string, ch chan<- float64) {
	if len(amount) == 0 {
		ch <- 0.0
		return
	}

	params := "payAssetUuid=" + payAsset + "&" + "payAssetAmount=" + amount + "&" + "receiveAssetUuid=" + receiveAsset + "&" + "withRange=" + "0"
	path := fmt.Sprintf("%s%s?%s", Endpoint, GetPricePath, params)
	resp, err := http.Get(path)
	if err != nil {
		logger.Errorf("http.Get(%s) => %v", path, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("%v", err)
	}
	//fmt.Println("body:", string(body))
	if gjson.Get(string(body), "success").String() != "true" {
		logger.Println("exinone.GetPrice() =>", string(body))
		ch <- 0.0
		return
	}

	ch <- gjson.Get(string(body), "data.calcSum").Float()
}

func Swap(typee, orderId, payAsset, receiveAsset, amount string) *mixin.TransferInput {
	var typeee string
	if typee == constants.SwapTypeMain {
		typeee = constants.SwapTypeMainInit
	} else {
		typeee = constants.SwapTypeFeeInit
	}
	memo := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("EX#CO#%s#%s", receiveAsset, REFCODE)))
	amt, _ := decimal.NewFromString(amount)
	input := &mixin.TransferInput{
		AssetID: payAsset,
		TraceID: mixin.UniqueConversationID(orderId, typeee),
		Amount:  amt,
		Memo:    memo,
	}
	input.OpponentMultisig.Receivers = []string{CLIENTID}
	input.OpponentMultisig.Threshold = uint8(1)
	return input
}

func GetStatus(traceID string) (string, string) {
	subpath := fmt.Sprintf("payTraceId=%s", traceID)
	path := fmt.Sprintf("%s%s?%s", Endpoint, GetStatusPath, subpath)
	fmt.Println(path)
	resp, err := http.Get(path)
	if err != nil {
		logger.Errorf("http.Get(%s) => %v", path, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("%v", err)
	}

	if gjson.Get(string(body), "code").String() == "30004" {
		return constants.StatusSwapSent, ""
	}
	switch gjson.Get(string(body), "data.orderStatus").String() {
	case "trading":
		return constants.StatusSwapSent, ""
	case "done":
		return constants.StatusSwapSuccess, gjson.Get(string(body), "data.receiveAmount").String()
	}
	if gjson.Get(string(body), "data.refundStatus").String() == "yes" {
		return constants.StatusSwapFailed, ""
	}

	return constants.StatusSwapSent, ""
}

/*

Not Found:
{"success":false,"message":"\u9009\u5b9a\u7684 pay trace id \u662f\u65e0\u6548\u7684","data":{"payTraceId":["\u9009\u5b9a\u7684 pay trace id \u662f\u65e0\u6548\u7684"]},"code":"30004"}

Trading:
{"code":"0","success":true,"message":"","data":{"id":10804976,"source":"snapshot","payAssetUuid":"43d61dcd-e413-450d-80b8-101d5e903357","payAssetSymbol":"ETH","payAmount":"0.00010000","payTraceId":"01246500-1caf-4ef7-9639-5b38711251f1","receiveTraceId":"64e1a2fd-f7c7-4e57-844a-6d29ac1c221e","receiveAssetUuid":"c6d0c728-2624-429b-8e0d-d9d19b6592fa","receiveAssetSymbol":"BTC","receiveAmount":null,"estimateReceiveAmount":null,"feeAmount":null,"feeAssetUuid":null,"feeAssetSymbol":null,"orderStatus":"trading","refundStatus":"no","payWalletUuid":"a13f4c77-5cfc-4368-a2d6-33f07037ae9e","payWalletType":"mixin","receiveWalletUuid":"a13f4c77-5cfc-4368-a2d6-33f07037ae9e","receiveWalletType":"mixin","createdAt":1695464944,"updatedAt":1695464944},"timestampMs":1695464945361}

Done:
{"code":"0","success":true,"message":"","data":{"id":10804976,"source":"snapshot","payAssetUuid":"43d61dcd-e413-450d-80b8-101d5e903357","payAssetSymbol":"ETH","payAmount":"0.00010000","payTraceId":"01246500-1caf-4ef7-9639-5b38711251f1","receiveTraceId":"64e1a2fd-f7c7-4e57-844a-6d29ac1c221e","receiveAssetUuid":"c6d0c728-2624-429b-8e0d-d9d19b6592fa","receiveAssetSymbol":"BTC","receiveAmount":"0.00000595","estimateReceiveAmount":null,"feeAmount":"0.00000002","feeAssetUuid":"c6d0c728-2624-429b-8e0d-d9d19b6592fa","feeAssetSymbol":"BTC","orderStatus":"done","refundStatus":"no","payWalletUuid":"a13f4c77-5cfc-4368-a2d6-33f07037ae9e","payWalletType":"mixin","receiveWalletUuid":"a13f4c77-5cfc-4368-a2d6-33f07037ae9e","receiveWalletType":"mixin","createdAt":1695464944,"updatedAt":1695464948},"timestampMs":1695464949295}

*/
