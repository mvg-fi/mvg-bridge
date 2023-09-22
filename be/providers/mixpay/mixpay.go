package mixpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/mvg-fi/common/logger"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

const (
	NAME     = "Mixpay"
	Endpoint = "https://api.mixpay.me"

	GetPricePath  = "/v1/payments_estimated"
	GetStatusPath = "/v1/payments_result"
	PaymentPath   = "/v1/payments"

	MultisigID = "d5af86f7-0ca7-46f7-aea1-bf916d536f87" //4311 4313-4315
	Threshold  = 2
)

func GetPrice(payAsset, receiveAsset, amount, except string, ch chan<- float64) {
	if len(amount) == 0 && len(except) == 0 {
		ch <- 0.0
		return
	}
	var params string
	if len(amount) > 0 {
		params = "paymentAssetId=" + url.QueryEscape(payAsset) + "&" + "settlementAssetId=" + url.QueryEscape(payAsset) + "&" + "quoteAssetId=" + url.QueryEscape(receiveAsset) + "&" + "paymentAmount=" + url.QueryEscape(amount)
	} else {
		params = "paymentAssetId=" + url.QueryEscape(payAsset) + "&" + "settlementAssetId=" + url.QueryEscape(payAsset) + "&" + "quoteAssetId=" + url.QueryEscape(receiveAsset) + "&" + "quoteAmount=" + url.QueryEscape(except)
	}
	path := fmt.Sprintf("%s?%s", Endpoint+GetPricePath, params)
	resp, err := http.Get(path)
	if err != nil {
		logger.Errorf("http.Get(%s) => %v", path, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("%v", err)
	}
	if gjson.Get(string(body), "success").String() != "true" {
		logger.Println("mixpay.GetPrice() =>", string(body))
		ch <- 0.0
		return
	}

	if len(amount) > 0 {
		ch <- gjson.Get(string(body), "data.quoteAmount").Float()
	} else {
		ch <- gjson.Get(string(body), "data.estimatedSettlementAmount").Float()
	}
}

func GetStatus(traceId, orderId, payeeId string) {
	var subpath string
	if len(traceId) > 0 {
		subpath = "traceId=" + traceId
	} else if len(orderId) > 0 {
		subpath = "orderId=" + orderId
	} else if len(payeeId) > 0 {
		subpath = "payeeId=" + payeeId
	}
	path := fmt.Sprintf("%s?%s", Endpoint+GetStatusPath, subpath)
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
		return
	}
	//_ := gjson.Get(string(body), "").String()
}

// For swaps
func createPayment(orderId, payAsset, receiveAsset, amount string, onChain bool) *MixpayPaymentResp {
	var chain int
	if onChain {
		chain = 1
	} else {
		chain = 0
	}
	js, err := json.Marshal(MixpayPaymentReq{
		PayeeID:           MultisigID,
		PaymentAssetID:    payAsset,
		SettlementAssetID: payAsset,
		QuoteAssetID:      receiveAsset,
		TraceID:           mixin.UniqueConversationID(orderId, "swap:init"),
		PaymentAmount:     amount,
		Remark:            "",
		IsChain:           chain,
	})
	if err != nil {
		logger.Errorf("%v", err)
	}

	resp, err := http.Post(Endpoint+PaymentPath, "application/json", bytes.NewBuffer(js))
	if err != nil {
		logger.Errorf("%v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("%v", err)
	}

	var mixpayResp MixpayPaymentResp
	err = json.Unmarshal(body, &mixpayResp)
	if err != nil {
		logger.Errorf("%v", err)
	}

	if !mixpayResp.Success {
		logger.Println("mixpay.createPayment() failed")
	}
	return &mixpayResp
}

func Swap(orderId, payAsset, receiveAsset, amount string, onChain bool) *mixin.TransferInput {
	// Create mixin payment
	mixpayResp := createPayment(orderId, payAsset, receiveAsset, amount, onChain)
	paymentAmount, _ := decimal.NewFromString(amount)
	return &mixin.TransferInput{
		AssetID:    mixpayResp.Data.PaymentAssetID,
		OpponentID: mixpayResp.Data.Recipient,
		TraceID:    mixin.UniqueConversationID(orderId, "swap:init"),
		Amount:     paymentAmount,
		Memo:       mixpayResp.Data.Memo,
	}
}

func GetMultisigId() {
	// https://mixpay.me/developers/api/multisig/get-multisig-id
}
