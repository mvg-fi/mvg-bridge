package constants

// KeyExample           = "{STATUS}|{TRACE UUID}"
// ValueExample         = "{TYPE}|{ADDRESS GENERATED}|{FROM ASSETID}|{TO ASSETID}|{TO ADDRESS}|{MEMO}|{AMOUNT}|{TIMESTAMP}"

const (
	// Prefix
	PrefixDepositPending  = "PENDING"
	PrefixDepositTimeout  = "TIMEOUT"
	PrefixDepositReceived = "RECEIVED"
	PrefixSwapInited      = "SWAPINITED"
	PrefixSwapFailed      = "SWAPFAILED"
	PrefixWithdrawInited  = "WITHDRAWINITED"
	PrefixWithdrawSucess  = "WITHDRAWSUCCESS"
	PrefixWithdrawFailed  = "WITHDRAWFAILED"

	// Deposit types
	CrossChainSwap = "CS"
	SameChainSwap  = "SS"
	Bridge         = "BE"
)
