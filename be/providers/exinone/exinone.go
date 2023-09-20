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

func GetStatus(traceId string) string {
	subpath := fmt.Sprintf("payTraceId=%s", traceId)
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
	fmt.Println(string(body))
	if gjson.Get(string(body), "success").String() != "true" {
		return ""
	}
	return gjson.Get(string(body), "data.orderStatus").String()
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
	return &mixin.TransferInput{
		AssetID:    payAsset,
		OpponentID: CLIENTID,
		TraceID:    mixin.UniqueConversationID(orderId, typeee),
		Amount:     amt,
		Memo:       memo,
	}
}
