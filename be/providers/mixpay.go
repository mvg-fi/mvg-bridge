package providers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const (
	Endpoint = "https://api.mixpay.me"

	GetPricePath  = "/v1/payments_estimated"
	SwapPath      = "/v1/payments"
	GetStatusPath = "/v1/payments_result"
)

func GetPrice(payAsset, receiveAsset, amount, except string) {
	if len(amount) == 0 && len(except) == 0 {
		return
	}
	params := ""
	if except == "" {
		params = "paymentAssetId=" + url.QueryEscape(payAsset) + "&" + "settlementAssetId=" + url.QueryEscape(payAsset) + "&" + "quoteAssetId=" + url.QueryEscape(receiveAsset) + "&" + "paymentAmount=" + url.QueryEscape(amount)
	} else {
		params = "paymentAssetId=" + url.QueryEscape(payAsset) + "&" + "settlementAssetId=" + url.QueryEscape(payAsset) + "&" + "quoteAssetId=" + url.QueryEscape(receiveAsset) + "&" + "quoteAmount=" + url.QueryEscape(except)
	}
	path := fmt.Sprintf("%s?%s", Endpoint+GetPricePath, params)
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var res map[string]any

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(Endpoint + GetPricePath)
	if except == "" {
		fmt.Println("Estimated receive", res["data"])
		//["quoteAmount"]
	} else {
		fmt.Println("Estimated pay", res["data"])
		//["estimatedSettlementAmount"]
	}
}

func GetStatus() {

}
