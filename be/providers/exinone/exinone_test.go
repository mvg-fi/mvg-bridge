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
	ip := Swap("main", uuid.NewV4(), "c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "0.1")
	fmt.Printf("%+v\n", ip)
}

// Transfer to Exinone bot
// Keep querying
// mixin-cli -f config.json transfer --receivers 61103d28-3ac2-44a2-ae34-bd956070dab1 --threold 1  --asset 43d61dcd-e413-450d-80b8-101d5e903357 --memo RVgjQ08jYzZkMGM3MjgtMjYyNC00MjliLThlMGQtZDlkMTliNjU5MmZhI20uNzAwMDEwNDIzMgo= --amount 0.0001 --trace 01246500-1caf-4ef7-9639-5b38711251f1

func TestStatus(t *testing.T) {
	traceID := "01246500-1caf-4ef7-9639-5b38711251f1"
	status := GetStatus(traceID)
	println(status)
}
