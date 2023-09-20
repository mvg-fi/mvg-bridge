package exinone

import (
	"fmt"
	"testing"

	"github.com/mvg-fi/common/uuid"
)

func TestGetPrice(t *testing.T) {
	ep := make(chan float64)
	go GetPrice("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "0.1", ep)
	epi := <-ep
	println(epi)
}

func TestSwap(t *testing.T) {
	ip := Swap(uuid.NewV4(), "c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "0.1")
	fmt.Printf("%+v\n", ip)
}
