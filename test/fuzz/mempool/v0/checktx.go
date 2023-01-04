package v0

import (
	"context"

	abciclient "github.com/HighStakesSwitzerland/tendermint/abci/client"
	"github.com/HighStakesSwitzerland/tendermint/abci/example/kvstore"
	"github.com/HighStakesSwitzerland/tendermint/config"
	"github.com/HighStakesSwitzerland/tendermint/internals/mempool"
	mempoolv0 "github.com/HighStakesSwitzerland/tendermint/internals/mempool/v0"
)

var mp mempool.Mempool

func init() {
	app := kvstore.NewApplication()
	cc := abciclient.NewLocalCreator(app)
	appConnMem, _ := cc()
	err := appConnMem.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false

	mp = mempoolv0.NewCListMempool(cfg, appConnMem, 0)
}

func Fuzz(data []byte) int {
	err := mp.CheckTx(context.Background(), data, nil, mempool.TxInfo{})
	if err != nil {
		return 0
	}

	return 1
}
