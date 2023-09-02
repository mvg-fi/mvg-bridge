package providers

import "testing"

func TestMixpayGetPrice(t *testing.T) {
	println("Test pay 1 BTC to estimate receive ETH")
	GetPrice("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "1", "")
	println("\n")

	println("Test receive 1 ETH to estimate pay BTC")
	GetPrice("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "", "1")
}
