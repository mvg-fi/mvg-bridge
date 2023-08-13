package workers

import (
	"context"

	"github.com/MixinNetwork/trusted-group/mtg"
	"github.com/fox-one/mixin-sdk-go"
)

type TimeoutWorker struct {
	grp mtg.Group
}

func NewTimeoutWorker(grp mtg.Group) *TimeoutWorker {
	return &TimeoutWorker{
		grp: grp,
	}
}

func (rw *TimeoutWorker) ProcessOutput(ctx context.Context, out *mtg.Output) {
	receivers := []string{out.Sender}
	traceId := mixin.UniqueConversationID(out.UTXOID, "refund")
	err := rw.grp.BuildTransaction(ctx, out.AssetID, receivers, int(1), out.Amount.String(), "refund", traceId, "")
	if err != nil {
		panic(err)
	}
}

func (rw *TimeoutWorker) ProcessCollectibleOutput(ctx context.Context, out *mtg.CollectibleOutput) {
	return
}
