package constants

type Deposit struct {
	Type        string `json:"type"`         // Transaction type (CS|SS|BE)
	Address     string `json:"address"`      // Deposit address for receiving user's fund
	ToAddress   string `json:"to_address"`   // The destination
	FromAssetID string `json:"from_assetid"` // Deposit asset id
	ToAssetID   string `json:"to_assetid"`   // Withdrawal asset id
	Memo        string `json:"memo"`         // Memo
	Amount      string `json:"amount"`       // User final receive amount of to asset
}

type DepositReq struct {
	Type        string `json:"type"`
	ToAddress   string `json:"to_address"`
	FromAssetID string `json:"from_assetid"`
	ToAssetID   string `json:"to_assetid"`
	Memo        string `json:"memo"`
	Amount      string `json:"amount"`
}

type DepositResp struct {
	Address string `json:address`
	Memo    string `json:memo,omitempty`
	TraceID string `json:trace_id`
}

type StatusReq struct {
	TraceID string `json:trace_id`
}

type StatusResp struct {
	Status string `json:status`
}
