package workers

import (
	"github.com/MixinNetwork/trusted-group/mtg"
)

type CleanerWorker struct {
	grp mtg.Group
}

func NewCleanerWorker(grp mtg.Group) *CleanerWorker {
	return &CleanerWorker{
		grp: grp,
	}
}

// 1. Remove all completed swaps after a period of time
// 2. Remove all expire orders
// 3. Remove all outdated withdrawals
