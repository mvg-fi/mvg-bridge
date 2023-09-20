package constants

import "github.com/shopspring/decimal"

// MTG
var (
	MTGMembers   = []string{"aca77da7-450c-4e34-867d-92ee07c1cfee", "3fb68263-4f06-476e-83db-503d25d56b93", "51186d7e-d488-417d-a031-b4e34f4fdf86", "1e92e114-bfa9-4989-b8a8-8e728bf432ef"}
	MTGThreshold = uint8(3)
)

// Values
var (
	BridgeFee = decimal.NewFromFloat(0.001) // 0.1%
)

// Strings
const (
	PrefixLock               = "LOCK:"
	PrefixUser               = "USER:"
	PrefixOrder              = "ORDER:"
	PrefixWithdrawal         = "WITHDRAWAL:"
	PrefixSnapshotList       = "SNAPSHOT:LIST:"
	PrefixSnapshotCheckpoint = "SNAPSHOT:CHECKPOINT"

	StatusReceived     = "RECEIVED"
	StatusSwapSent     = "SWAPSENT"
	StatusSwapReceived = "SWAPRECEIVED"
	StatusWithdrawSent = "WITHDRAWSENT"

	SwapTypeMain     = "main"
	SwapTypeFee      = "fee"
	SwapTypeMainInit = "swap:main:init"
	SwapTypeFeeInit  = "swap:fee:init"

	MTGGroupIDSwap = "SWAP"
)