package v1

import (
	"context"

	abciclient "github.com/HighStakesSwitzerland/tendermint/abci/client"
	"github.com/HighStakesSwitzerland/tendermint/abci/example/kvstore"
	"github.com/HighStakesSwitzerland/tendermint/config"
	"github.com/HighStakesSwitzerland/tendermint/internals/mempool"
	mempoolv1 "github.com/HighStakesSwitzerland/tendermint/internals/mempool/v1"
	"github.com/HighStakesSwitzerland/tendermint/libs/log"
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

	mp = mempoolv1.NewTxMempool(log.NewNopLogger(), cfg, appConnMem, 0)
}

func Fuzz(data []byte) int {
	err := mp.CheckTx(context.Background(), data, nil, mempool.TxInfo{})
	if err != nil {
		return 0
	}

	return 1
}
