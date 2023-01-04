package client_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/HighStakesSwitzerland/tendermint/abci/example/kvstore"
	"github.com/HighStakesSwitzerland/tendermint/config"
	"github.com/HighStakesSwitzerland/tendermint/libs/service"
	rpctest "github.com/HighStakesSwitzerland/tendermint/rpc/test"
	"github.com/stretchr/testify/require"
)

func NodeSuite(t *testing.T) (service.Service, *config.Config) {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())

	conf, err := rpctest.CreateConfig(t.Name())
	require.NoError(t, err)

	// start a tendermint node in the background to test against
	dir, err := ioutil.TempDir("/tmp", fmt.Sprint("rpc-client-test-", t.Name()))
	require.NoError(t, err)

	app := kvstore.NewPersistentKVStoreApplication(dir)

	node, closer, err := rpctest.StartTendermint(ctx, conf, app, rpctest.SuppressStdout)
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = closer(ctx)
		cancel()
		app.Close()
		_ = os.RemoveAll(dir)
	})
	return node, conf
}
