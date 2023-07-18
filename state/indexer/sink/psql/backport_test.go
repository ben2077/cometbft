package psql

import (
	"github.com/ben2077/cometbft/state/indexer"
	"github.com/ben2077/cometbft/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
