package constants

const ExpirePeriod = 10 * 60 // 10 mins

type Order struct {
	FromAssetID string `json:"from_asset_id" msgpack:"f"`
	ToAssetID   string `json:"to_asset_id" msgpack:"t"`
	Amount      string `json:"amount" msgpack:"at"`
	Except      string `json:"except" msgpack:"et"`
	Cex         bool   `json:"cex" msgpack:"c"`
	TraceID     string `json:"trace_id" msgpack:"td"`
	Address     string `json:"address" msgpack:"ad"`
	Memo        string `json:"memo" msgpack:"m"`
	Expire      string `json:"expire" msgpack:"ee"`
	Status      string `json:"status" msgpack:"s"`
}

type OrderMinium struct {
	ToAssetID string `json:"to_asset_id" msgpack:"t"`
	Except    string `json:"except" msgpack:"et"`
	Cex       bool   `json:"cex" msgpack:"c"`
	TraceID   string `json:"trace_id" msgpack:"td"`
	Address   string `json:"address" msgpack:"ad"`
	Memo      string `json:"memo" msgpack:"m"`
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
