package constants

type Swap struct {
	OrderID  string `json:"order_id" msgpack: "o"`
	Status   string `json:"status" msgpack: "s"`
	Provider string `json:"provider" msgpack: "p"`
}
