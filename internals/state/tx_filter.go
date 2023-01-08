package state

import (
	"github.com/HighStakesSwitzerland/tendermint/internals/mempool"
	"github.com/HighStakesSwitzerland/tendermint/types"
)

// TxPreCheck returns a function to filter transactions before processing.
// The function limits the size of a transaction to the block's maximum data size.
func TxPreCheck(state State) mempool.PreCheckFunc {
	maxDataBytes := types.MaxDataBytesNoEvidence(
		state.ConsensusParams.Block.MaxBytes,
		state.Validators.Size(),
	)
	return mempool.PreCheckMaxBytes(maxDataBytes)
}

// TxPostCheck returns a function to filter transactions after processing.
// The function limits the gas wanted by a transaction to the block's maximum total gas.
func TxPostCheck(state State) mempool.PostCheckFunc {
	return mempool.PostCheckMaxGas(state.ConsensusParams.Block.MaxGas)
}