//go:build gofuzz || go1.20

package tests

import (
	"testing"

	abciclient "github.com/ben2077/cometbft/abci/client"
	"github.com/ben2077/cometbft/abci/example/kvstore"
	"github.com/ben2077/cometbft/config"
	cmtsync "github.com/ben2077/cometbft/libs/sync"
	mempool "github.com/ben2077/cometbft/mempool"
)

func FuzzMempool(f *testing.F) {
	app := kvstore.NewInMemoryApplication()
	mtx := new(cmtsync.Mutex)
	conn := abciclient.NewLocalClient(mtx, app)
	err := conn.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false

	mp := mempool.NewCListMempool(cfg, conn, 0)

	f.Fuzz(func(t *testing.T, data []byte) {
		_ = mp.CheckTx(data, nil, mempool.TxInfo{})
	})
}
