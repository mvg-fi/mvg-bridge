package pandoswap

import "testing"

func TestFswapPrice(t *testing.T) {
	println(GetPrice("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "1", ""))
	println(GetPrice("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "", "1"))
}

func TestFswapStatus(t *testing.T) {
	GetStatus("")
}
