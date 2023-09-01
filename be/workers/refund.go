package workers

import (
	"context"
	"fmt"

	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/fox-one/mixin-sdk-go"
)

type RefundWorker struct {
	grp mtg.Group
}

func NewRefundWorker(grp mtg.Group) *RefundWorker {
	return &RefundWorker{
		grp: grp,
	}
}

func (rw *RefundWorker) ProcessOutput(ctx context.Context, out *mtg.Output) {
	fmt.Println("mtg.Output:", out)
	receivers := []string{out.Sender}
	traceId := mixin.UniqueConversationID(out.UTXOID, "refund")
	err := rw.grp.BuildTransaction(ctx, out.AssetID, receivers, int(1), out.Amount.String(), "refund", traceId, "")
	if err != nil {
		panic(err)
	}
}

func (rw *RefundWorker) ProcessCollectibleOutput(ctx context.Context, out *mtg.CollectibleOutput) {
	return
}
