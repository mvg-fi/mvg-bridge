package mixpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/tidwall/gjson"
)

const (
	Endpoint = "https://api.mixpay.me"

	GetPricePath  = "/v1/payments_estimated"
	GetStatusPath = "/v1/payments_result"
	PaymentPath   = "/v1/payments"
)

func GetPrice(payAsset, receiveAsset, amount, except string) float64 {
	if len(amount) == 0 && len(except) == 0 {
		return 0.0
	}
	var params string
	if len(except) == 0 {
		params = "paymentAssetId=" + url.QueryEscape(payAsset) + "&" + "settlementAssetId=" + url.QueryEscape(payAsset) + "&" + "quoteAssetId=" + url.QueryEscape(receiveAsset) + "&" + "paymentAmount=" + url.QueryEscape(amount)
	} else {
		params = "paymentAssetId=" + url.QueryEscape(payAsset) + "&" + "settlementAssetId=" + url.QueryEscape(payAsset) + "&" + "quoteAssetId=" + url.QueryEscape(receiveAsset) + "&" + "quoteAmount=" + url.QueryEscape(except)
	}
	path := fmt.Sprintf("%s?%s", Endpoint+GetPricePath, params)
	resp, err := http.Get(path)
	if err != nil {
		log.Println(fmt.Sprintf("http.Get(%s)", path), err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	if gjson.Get(string(body), "success").String() != "true" {
		log.Println("mixpay.GetPrice() failed")
		return 0.0
	}

	if len(except) == 0 {
		return gjson.Get(string(body), "data.quoteAmount").Float()
	} else {
		return gjson.Get(string(body), "data.estimatedSettlementAmount").Float()
	}
}

func GetStatus(traceId, orderId, payeeId string) {
	var subpath string
	if len(traceId) != 0 {
		subpath = "traceId=" + traceId
	} else if len(orderId) != 0 {
		subpath = "orderId=" + orderId
	} else if len(payeeId) != 0 {
		subpath = "payeeId=" + payeeId
	}
	path := fmt.Sprintf("%s?%s", Endpoint+GetStatusPath, subpath)
	fmt.Println(path)
	resp, err := http.Get(path)
	if err != nil {
		log.Println(fmt.Sprintf("http.Get(%s)", path), err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(body))
	if gjson.Get(string(body), "success").String() != "true" {
		return
	}
	//_ := gjson.Get(string(body), "").String()
}

// For swaps
func CreatePayment(payeeId, traceId, payAsset, receiveAsset, amount string, onChain bool) {
	var chain int
	if onChain {
		chain = 1
	} else {
		chain = 0
	}
	js, err := json.Marshal(map[string]string{
		"payeeId":           payeeId,
		"orderId":           traceId,
		"paymentAssetId":    payAsset,
		"settlementAssetId": payAsset,
		"quoteAssetId":      receiveAsset,
		"paymentAmount":     amount,
		"isChain":           strconv.Itoa(chain),
	})
	if err != nil {
		log.Println(err)
	}

	resp, err := http.Post(Endpoint+PaymentPath, "application/json", bytes.NewBuffer(js))

	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(body))
	if gjson.Get(string(body), "success").String() != "true" {
		log.Println("mixpay.CreatePayment() failed")
		return
	}
	dest := gjson.Get(string(body), "data.destination").String()
	fmt.Println("dest:", dest)
}
