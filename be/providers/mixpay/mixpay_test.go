package mixpay

import "testing"

func TestMixpayGetPrice(t *testing.T) {
	println("Test pay 1 BTC to estimate receive ETH")
	println(GetPrice("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "1", ""))
	println("\n")

	println("Test receive 1 ETH to estimate pay BTC")
	println(GetPrice("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "", "1"))
}

func TestMixpayGetStatus(t *testing.T) {
	println("Test get order status")
	GetStatus("a0d7791408776b47eb1dd3f94ed15d6a", "", "")
}

func TestMixpayPayment(t *testing.T) {
	println("Test swap 1 BTC to ETH")
	CreatePayment("a13f4c77-5cfc-4368-a2d6-33f07037ae9e", "35f2aa7a-d3d9-4f9f-98d0-aee1a17a13e2", "c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "1", true)
}
