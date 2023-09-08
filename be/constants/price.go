package constants

type PriceSimpleReq struct {
	FromAssetID string `json:"from_asset_id"`
	ToAssetID   string `json:"to_asset_id"`
	Amount      string `json:"amount"`
	Except      string `json:"except"`
	Cex         bool   `json:"cex"`
}

type PriceSimpleResp struct {
	Amount string `json:"amount"`
	Fee    string `json:"fee"`
}

type PriceAllReq struct {
	FromAssetID string `json:"from_asset_id"`
	ToAssetID   string `json:"to_asset_id"`
	Amount      string `json:"amount"`
	Except      string `json:"except"`
	Cex         bool   `json:"cex"`
}

type PriceAllResp struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
	Fee    string `json:"fee"`
}
