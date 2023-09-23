package constants

type Swap struct {
	OrderID  string `json:"order_id" msgpack: "o"`
	TraceID  string `json:"trace_id" msgpack: "t"`
	Status   string `json:"status" msgpack: "s"`
	Provider string `json:"provider" msgpack: "p"`
	Type     string `json:"type" msgpack:"e"`
}
