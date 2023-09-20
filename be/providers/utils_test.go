package providers

import "testing"

func TestAfterBridgeFee(t *testing.T) {
	// (10-0.005)/1000*999
	println(AfterBridgeFee("10", "0.005"))
}
