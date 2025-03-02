package indexer

import (
	"context"

	"github.com/ben2077/cometbft/libs/log"
	"github.com/ben2077/cometbft/libs/pubsub/query"
	"github.com/ben2077/cometbft/types"
)

//go:generate ../../scripts/mockery_generate.sh BlockIndexer

// BlockIndexer defines an interface contract for indexing block events.
type BlockIndexer interface {
	// Has returns true if the given height has been indexed. An error is returned
	// upon database query failure.
	Has(height int64) (bool, error)

	// Index indexes FinalizeBlock events for a given block by its height.
	Index(types.EventDataNewBlockEvents) error

	// Search performs a query for block heights that match a given FinalizeBlock
	// event search criteria.
	Search(ctx context.Context, q *query.Query) ([]int64, error)

	SetLogger(l log.Logger)
}
