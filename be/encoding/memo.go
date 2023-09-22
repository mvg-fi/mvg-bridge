package encoding

import (
	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/mvg-fi/common/logger"
	"github.com/mvg-fi/mvg-bridge/constants"
)

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
