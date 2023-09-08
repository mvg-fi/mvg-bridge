package providers

import (
	"fmt"
	"testing"
)

func TestGetPriceSimple(t *testing.T) {
	price, fee := GetPriceSimple("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "1", "", true)
	fmt.Println("\n1 BTC -> ETH", price, fee)
	price, fee = GetPriceSimple("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "0.1", "", true)
	fmt.Println("\n0.1 BTC -> ETH", price, fee)
}

func TestNoCexSimple(t *testing.T) {
	price, fee := GetPriceSimple("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "1", "", false)
	fmt.Println("\n1 BTC -> ETH", price, fee)
	price, fee = GetPriceSimple("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "0.1", "", false)
	fmt.Println("\n0.1 BTC -> ETH", price, fee)

}

func TestGetPriceAll(t *testing.T) {
	prices := GetPriceAll("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "1", "", true)
	fmt.Println("\n1 BTC -> ETH", prices)
	prices = GetPriceAll("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "0.1", "", true)
	fmt.Println("\n0.1 BTC -> ETH", prices)
}

func TestNoCexAll(t *testing.T) {
	prices := GetPriceAll("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "1", "", false)
	fmt.Println("\n1 BTC -> ETH", prices)
	prices = GetPriceAll("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "0.1", "", false)
	fmt.Println("\n0.1 BTC -> ETH", prices)
}
