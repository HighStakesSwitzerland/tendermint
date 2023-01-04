package blocksync

import (
	bcproto "github.com/HighStakesSwitzerland/tendermint/proto/tendermint/blocksync"
	"github.com/HighStakesSwitzerland/tendermint/types"
)

const (
	MaxMsgSize = types.MaxBlockSizeBytes +
		bcproto.BlockResponseMessagePrefixSize +
		bcproto.BlockResponseMessageFieldKeySize
)
