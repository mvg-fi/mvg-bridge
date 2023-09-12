package constants

const ExpirePeriod = 10 * 60 // 10 mins

type Order struct {
	FromAssetID string `json:"from_asset_id"`
	ToAssetID   string `json:"to_asset_id"`
	Amount      string `json:"amount"`
	Except      string `json:"except"`
	Cex         bool   `json:"cex"`
	TraceID     string `json:"trace_id"`
	Address     string `json:"address"`
	Memo        string `json:"memo"`
	Expire      string `json:"expire"`
}

type OrderNewReq struct {
	FromAssetID string `json:"from_asset_id"`
	ToAssetID   string `json:"to_asset_id"`
	Amount      string `json:"amount"`
	Except      string `json:"except"`
	Cex         bool   `json:"cex"`
}

type OrderNewResp struct {
	TraceID string `json:"trace_id"`
	Address string `json:"address"`
	Memo    string `json:"memo"`
	Expire  string `json:"expire"`
}
