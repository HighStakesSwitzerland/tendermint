package null

import (
	"context"
	"errors"

	"github.com/HighStakesSwitzerland/tendermint/internals/state/indexer"
	"github.com/HighStakesSwitzerland/tendermint/libs/pubsub/query"
	"github.com/HighStakesSwitzerland/tendermint/types"
)

var _ indexer.BlockIndexer = (*BlockerIndexer)(nil)

// TxIndex implements a no-op block indexer.
type BlockerIndexer struct{}

func (idx *BlockerIndexer) Has(height int64) (bool, error) {
	return false, errors.New(`indexing is disabled (set 'tx_index = "kv"' in config)`)
}

func (idx *BlockerIndexer) Index(types.EventDataNewBlockHeader) error {
	return nil
}

func (idx *BlockerIndexer) Search(ctx context.Context, q *query.Query) ([]int64, error) {
	return []int64{}, nil
}