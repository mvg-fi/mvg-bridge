package constants

// KeyExample = "{PREFIX}|{DEPOSIT TYPES}|{ADDRESS GENERATED}|{TIMESTAMP}|{SHORTEN TRACE UUID}"
// ValueExample = "{SHORTEN FROM ASSETID}|{SHORTEN TO ASSETID}|{TO ADDRESS}|{MEMO}|{AMOUNT}"

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
