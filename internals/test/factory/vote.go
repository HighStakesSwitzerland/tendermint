package factory

import (
	"context"
	"time"

	tmproto "github.com/HighStakesSwitzerland/tendermint/proto/tendermint/types"
	"github.com/HighStakesSwitzerland/tendermint/types"
)

func MakeVote(
	val types.PrivValidator,
	chainID string,
	valIndex int32,
	height int64,
	round int32,
	step int,
	blockID types.BlockID,
	time time.Time,
) (*types.Vote, error) {
	pubKey, err := val.GetPubKey(context.Background())
	if err != nil {
		return nil, err
	}
	v := &types.Vote{
		ValidatorAddress: pubKey.Address(),
		ValidatorIndex:   valIndex,
		Height:           height,
		Round:            round,
		Type:             tmproto.SignedMsgType(step),
		BlockID:          blockID,
		Timestamp:        time,
	}

	vpb := v.ToProto()
	err = val.SignVote(context.Background(), chainID, vpb)
	if err != nil {
		panic(err)
	}
	v.Signature = vpb.Signature
	return v, nil
}