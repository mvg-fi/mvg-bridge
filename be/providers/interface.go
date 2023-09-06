package providers

type Provider interface {
	GetPrice(payAsset, receiveAsset, amount, except string) float64
	Swap(from, to, amount string)
	GetStatus(traceId string)
}

type Status struct {
}
