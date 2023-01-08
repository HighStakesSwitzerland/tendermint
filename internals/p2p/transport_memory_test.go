package p2p_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/HighStakesSwitzerland/tendermint/internals/p2p"
	"github.com/HighStakesSwitzerland/tendermint/libs/log"
	"github.com/HighStakesSwitzerland/tendermint/types"
	"github.com/stretchr/testify/require"
)

// Transports are mainly tested by common tests in transport_test.go, we
// register a transport factory here to get included in those tests.
func init() {
	var network *p2p.MemoryNetwork // shared by transports in the same test

	testTransports["memory"] = func(t *testing.T) p2p.Transport {
		if network == nil {
			network = p2p.NewMemoryNetwork(log.TestingLogger(), 1)
		}
		i := byte(network.Size())
		nodeID, err := types.NewNodeID(hex.EncodeToString(bytes.Repeat([]byte{i<<4 + i}, 20)))
		require.NoError(t, err)
		transport := network.CreateTransport(nodeID)

		t.Cleanup(func() {
			require.NoError(t, transport.Close())
			network = nil // set up a new memory network for the next test
		})

		return transport
	}
}