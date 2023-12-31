package encoding

import (
	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/constants"
)

// The memo from proxy users to MTG
// Return value length must be under 200
func MsgpackCompressOrder(o *constants.Order) string {
	return string(mtg.CompressMsgpackMarshalPanic(&constants.OrderMinium{
		ToAssetID: o.ToAssetID,
		Except:    o.Except,
		Cex:       o.Cex,
		TraceID:   o.TraceID,
		Address:   o.Address,
		Memo:      o.Memo,
	}))
}

func MsgpackDecompressOrder(s string) *constants.Order {
	var minimum constants.OrderMinium
	err := mtg.DecompressMsgpackUnmarshal([]byte(s), &minimum)
	if err != nil {
		logger.Errorf("MsgpackDecompressOrder(%s) => %v", s, err)
	}
	return &constants.Order{
		ToAssetID: minimum.ToAssetID,
		Except:    minimum.Except,
		Cex:       minimum.Cex,
		TraceID:   minimum.TraceID,
		Address:   minimum.Address,
		Memo:      minimum.Memo,
	}
}

func MsgpackMashalOrder(o *constants.Order) string {
	return string(mtg.MsgpackMarshalPanic(o))
}

// The memo from MTG to withdrawal bot
func GetWithdrawalMemo(o *constants.Order, main *constants.Swap, fee *constants.Swap) (string, string) {
	return string(mtg.CompressMsgpackMarshalPanic(&constants.Withdrawal{
			OrderID: o.TraceID,
			Asset:   main.ToAssetID,
			Amount:  main.Receive,
			Address: o.Address,
			Memo:    o.Memo,
		})),
		string(mtg.CompressMsgpackMarshalPanic(&constants.Withdrawal{
			OrderID: o.TraceID,
			Asset:   fee.ToAssetID,
			Amount:  fee.Receive,
			Address: o.Address,
			Memo:    o.Memo,
		}))
}

func RetriveWithdrawalMemo(memo string) (*constants.Withdrawal, error) {
	var withdrawal constants.Withdrawal
	err := mtg.DecompressMsgpackUnmarshal([]byte(memo), &withdrawal)
	if err != nil {
		return nil, err
	}
	return &withdrawal, nil
}
