package constants

const ExpirePeriod = 10 * 60 // 10 mins

type Order struct {
	FromAssetID string `json:"from_asset_id" msgpack:"f"`
	ToAssetID   string `json:"to_asset_id" msgpack:"t"`
	Amount      string `json:"amount" msgpack:"a"`
	Except      string `json:"except" msgpack:"x"`
	Cex         bool   `json:"cex" msgpack:"c"`
	TraceID     string `json:"trace_id" msgpack:"r"`
	Address     string `json:"address" msgpack:"d"`
	Memo        string `json:"memo" msgpack:"m"`
	Expire      string `json:"expire" msgpack:"e"`
	Status      string `json:"status" msgpack:"s"`
}

type OrderMinium struct {
	ToAssetID string `json:"to_asset_id" msgpack:"t"`
	Except    string `json:"except" msgpack:"e"`
	Cex       bool   `json:"cex" msgpack:"c"`
	TraceID   string `json:"trace_id" msgpack:"r"`
	Address   string `json:"address" msgpack:"d"`
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

type Swap struct {
	OrderID     string `json:"order_id" msgpack: "o"`
	TraceID     string `json:"trace_id" msgpack: "a"`
	FromAssetID string `json:"asset" msgpack: "f"`
	ToAssetID   string `json:"to_asset_id" msgpack:"t"`
	Amount      string `json:"amount" msgpack: "m"`
	Receive     string `json:"receive" msgpack: "r"`
	Status      string `json:"status" msgpack: "s"`
	Provider    string `json:"provider" msgpack: "p"`
	Type        string `json:"type" msgpack:"e"`
}

type Withdrawal struct {
	Asset   string `json:"asset" msgpack:"a"`
	Amount  string `json:"amount" msgpack:"o"`
	Address string `json:"address" msgpack:"d"`
	Memo    string `json:"memo" msgpack:"m"`
}

type WithdrawalFull struct {
	OrderID   string `json:"order_id" msgpack:"r"`
	MainTrace string `json:"main_trace" msgpack:"mt"`
	FeeTrace  string `json:"fee_trace" msgpack:"ft"`
	Asset     string `json:"asset" msgpack:"a"`
	Amount    string `json:"amount" msgpack:"o"`
	Address   string `json:"address" msgpack:"d"`
	Memo      string `json:"memo" msgpack:"m"`
}
