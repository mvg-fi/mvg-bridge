package providers

type Provider interface {
	GetPrice(from, to, amount, except string)
	Swap(from, to, amount string)
	GetStatus(traceId string)
}
