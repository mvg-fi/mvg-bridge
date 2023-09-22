package exinone

type ExinonePriceReq struct {
	// Method is GET so no need
}

type ExinonePriceResp struct {
	Code    string `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		CalcSum               string `json:"calcSum"`
		CalcPrice             string `json:"calcPrice"`
		IsAllowEpc            bool   `json:"isAllowEpc"`
		Fee                   string `json:"fee"`
		IsReverse             bool   `json:"isReverse"`
		RoutePayAssetAmount   string `json:"routePayAssetAmount"`
		RouteExchangeMax      string `json:"routeExchangeMax"`
		IsDelay               bool   `json:"isDelay"`
		DelayReleaseTime      string `json:"delayReleaseTime"`
		ReceiveAssetPriceUsdt string `json:"receiveAssetPriceUsdt"`
		PayAssetPriceUsdt     string `json:"payAssetPriceUsdt"`
		IsTradeAvailable      bool   `json:"isTradeAvailable"`
		Route                 struct {
			Asset []struct {
				ID           int    `json:"id"`
				MixinID      string `json:"mixinId"`
				MixinChainID string `json:"mixinChainId"`
				Name         string `json:"name"`
				Symbol       string `json:"symbol"`
				IconURL      string `json:"iconUrl"`
			} `json:"asset"`
			Exchange []struct {
				ID      int    `json:"id"`
				Name    string `json:"name"`
				En      string `json:"en"`
				Cn      string `json:"cn"`
				Favicon string `json:"favicon"`
				Logo32  string `json:"logo_32"`
			} `json:"exchange"`
		} `json:"route"`
		Range struct {
			Min             string `json:"min"`
			Max             string `json:"max"`
			WithOutDelayMax string `json:"withOutDelayMax"`
		} `json:"range"`
	} `json:"data"`
	TimestampMs int64 `json:"timestampMs"`
}

type ExinoneStatusReq struct {
	// Method is GET so no need
}

type ExinoneStatusResp struct {
	Code    string `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		ID                    int    `json:"id"`
		Source                string `json:"source"`
		PayAssetUUID          string `json:"payAssetUuid"`
		PayAssetSymbol        string `json:"payAssetSymbol"`
		PayAmount             string `json:"payAmount"`
		PayTraceID            string `json:"payTraceId"`
		ReceiveAssetUUID      string `json:"receiveAssetUuid"`
		ReceiveAssetSymbol    string `json:"receiveAssetSymbol"`
		ReceiveAmount         string `json:"receiveAmount"`
		EstimateReceiveAmount any    `json:"estimateReceiveAmount"`
		FeeAmount             string `json:"feeAmount"`
		FeeAssetUUID          string `json:"feeAssetUuid"`
		FeeAssetSymbol        string `json:"feeAssetSymbol"`
		OrderStatus           string `json:"orderStatus"`
		RefundStatus          string `json:"refundStatus"`
		PayWalletUUID         string `json:"payWalletUuid"`
		PayWalletType         string `json:"payWalletType"`
		ReceiveWalletUUID     string `json:"receiveWalletUuid"`
		ReceiveWalletType     any    `json:"receiveWalletType"`
		CreatedAt             int    `json:"createdAt"`
		UpdatedAt             int    `json:"updatedAt"`
	} `json:"data"`
	TimestampMs int64 `json:"timestampMs"`
}
