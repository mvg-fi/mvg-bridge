package pandoswap

import (
	"context"
	"fmt"
	"testing"
)

func TestReadAssets(t *testing.T) {
	assets := ReadAssets(context.TODO())
	fmt.Println(assets)
}

func TestFswapPrice(t *testing.T) {
	//println(GetPrice("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "1", ""))
	//println(GetPrice("c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "", "1"))
}

func TestFswapStatus(t *testing.T) {
	GetStatus("")
}

func TestFsSwap(t *testing.T) {
	//fmt.Printf("%v", Swap([]string{"c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357"}, 1, "c6d0c728-2624-429b-8e0d-d9d19b6592fa", "c6d0c728-2624-429b-8e0d-d9d19b6592fa", "43d61dcd-e413-450d-80b8-101d5e903357", "1", ""))
}
