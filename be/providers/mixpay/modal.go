package mixpay

type MixpayPaymentReq struct {
	PayeeID           string `json:"payeeId"`
	PaymentAssetID    string `json:"paymentAssetId"`
	SettlementAssetID string `json:"settlementAssetId"`
	QuoteAssetID      string `json:"quoteAssetId"`
	TraceID           string `json:"traceId"`
	PaymentAmount     string `json:"paymentAmount"`
	Remark            string `json:"remark,omitempty"`
	IsChain           int    `json:"isChain,omitempty"`
}

type MixpayPaymentResp struct {
	Code        int                   `json:"code"`
	Data        MixpayPaymentRespData `json:"data"`
	Message     string                `json:"message"`
	Success     bool                  `json:"success"`
	TimestampMs int64                 `json:"timestampMs"`
}

type MixpayPaymentRespData struct {
	CancelURL                 string `json:"cancelUrl"`
	CheckoutURL               string `json:"checkoutUrl"`
	ClientID                  string `json:"clientId"`
	DeepLink                  string `json:"deepLink"`
	Destination               string `json:"destination"`
	EstimatedSettlementAmount string `json:"estimatedSettlementAmount"`
	Expire                    int64  `json:"expire"`
	IsChain                   bool   `json:"isChain"`
	Memo                      string `json:"memo"`
	MVMMemo                   string `json:"mvmMemo"`
	Note                      string `json:"note"`
	OrderID                   string `json:"orderId"`
	PayeeID                   string `json:"payeeId"`
	PayerEmail                string `json:"payerEmail"`
	PaymentAmount             string `json:"paymentAmount"`
	PaymentAssetIcon          string `json:"paymentAssetIcon"`
	PaymentAssetID            string `json:"paymentAssetId"`
	PaymentAssetSymbol        string `json:"paymentAssetSymbol"`
	QRString                  string `json:"qrString"`
	QuoteAmount               string `json:"quoteAmount"`
	QuoteAssetIcon            string `json:"quoteAssetIcon"`
	QuoteAssetID              string `json:"quoteAssetId"`
	QuoteAssetSymbol          string `json:"quoteAssetSymbol"`
	Recipient                 string `json:"recipient"`
	Seconds                   int    `json:"seconds"`
	SettlementAssetIcon       string `json:"settlementAssetIcon"`
	SettlementAssetID         string `json:"settlementAssetId"`
	SettlementAssetSymbol     string `json:"settlementAssetSymbol"`
	Tag                       string `json:"tag"`
	TraceID                   string `json:"traceId"`
	UniversalURL              string `json:"universalUrl"`
}
