package txindex_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	db "github.com/cometbft/cometbft-db"

	abci "github.com/ben2077/cometbft/abci/types"
	"github.com/ben2077/cometbft/libs/log"
	blockidxkv "github.com/ben2077/cometbft/state/indexer/block/kv"
	"github.com/ben2077/cometbft/state/txindex"
	"github.com/ben2077/cometbft/state/txindex/kv"
	"github.com/ben2077/cometbft/types"
)

func TestIndexerServiceIndexesBlocks(t *testing.T) {
	// event bus
	eventBus := types.NewEventBus()
	eventBus.SetLogger(log.TestingLogger())
	err := eventBus.Start()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := eventBus.Stop(); err != nil {
			t.Error(err)
		}
	})

	// tx indexer
	store := db.NewMemDB()
	txIndexer := kv.NewTxIndex(store)
	blockIndexer := blockidxkv.New(db.NewPrefixDB(store, []byte("block_events")))

	service := txindex.NewIndexerService(txIndexer, blockIndexer, eventBus, false)
	service.SetLogger(log.TestingLogger())
	err = service.Start()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := service.Stop(); err != nil {
			t.Error(err)
		}
	})

	// publish block with events
	err = eventBus.PublishEventNewBlockEvents(types.EventDataNewBlockEvents{
		Height: 1,
		Events: []abci.Event{
			{
				Type: "begin_event",
				Attributes: []abci.EventAttribute{
					{
						Key:   "proposer",
						Value: "FCAA001",
						Index: true,
					},
				},
			},
		},
		NumTxs: int64(2),
	})
	require.NoError(t, err)
	txResult1 := &abci.TxResult{
		Height: 1,
		Index:  uint32(0),
		Tx:     types.Tx("foo"),
		Result: abci.ExecTxResult{Code: 0},
	}
	err = eventBus.PublishEventTx(types.EventDataTx{TxResult: *txResult1})
	require.NoError(t, err)
	txResult2 := &abci.TxResult{
		Height: 1,
		Index:  uint32(1),
		Tx:     types.Tx("bar"),
		Result: abci.ExecTxResult{Code: 0},
	}
	err = eventBus.PublishEventTx(types.EventDataTx{TxResult: *txResult2})
	require.NoError(t, err)

	time.Sleep(100 * time.Millisecond)

	res, err := txIndexer.Get(types.Tx("foo").Hash())
	require.NoError(t, err)
	require.Equal(t, txResult1, res)

	ok, err := blockIndexer.Has(1)
	require.NoError(t, err)
	require.True(t, ok)

	res, err = txIndexer.Get(types.Tx("bar").Hash())
	require.NoError(t, err)
	require.Equal(t, txResult2, res)
}
