package core

import (
	"github.com/HighStakesSwitzerland/tendermint/rpc/coretypes"
	rpctypes "github.com/HighStakesSwitzerland/tendermint/rpc/jsonrpc/types"
)

// UnsafeFlushMempool removes all transactions from the mempool.
func (env *Environment) UnsafeFlushMempool(ctx *rpctypes.Context) (*coretypes.ResultUnsafeFlushMempool, error) {
	env.Mempool.Flush()
	return &coretypes.ResultUnsafeFlushMempool{}, nil
}
