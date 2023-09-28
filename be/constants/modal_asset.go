package constants

type Asset struct {
	AssetID   string `json:"asset_id"`
	AssetKey  string `json:"asset_key"`
	ChainIcon string `json:"chain_icon"`
	ChainID   string `json:"chain_id"`
	ChainName string `json:"chain_name"`
	Icon      string `json:"icon"`
	Name      string `json:"name"`
	Symbol    string `json:"symbol"`
}
