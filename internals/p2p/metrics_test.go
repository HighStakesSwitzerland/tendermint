package p2p

import (
	"testing"

	"github.com/HighStakesSwitzerland/tendermint/proto/tendermint/p2p"
	"github.com/stretchr/testify/assert"
)

func TestValueToMetricsLabel(t *testing.T) {
	m := NopMetrics()
	r := &p2p.PexResponse{}
	str := m.ValueToMetricLabel(r)
	assert.Equal(t, "p2p_PexResponse", str)

	// subsequent calls to the function should produce the same result
	str = m.ValueToMetricLabel(r)
	assert.Equal(t, "p2p_PexResponse", str)
}
