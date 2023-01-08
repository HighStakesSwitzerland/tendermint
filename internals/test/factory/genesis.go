package factory

import (
	"sort"

	"github.com/HighStakesSwitzerland/tendermint/config"
	tmtime "github.com/HighStakesSwitzerland/tendermint/libs/time"
	"github.com/HighStakesSwitzerland/tendermint/types"
)

func RandGenesisDoc(
	cfg *config.Config,
	numValidators int,
	randPower bool,
	minPower int64) (*types.GenesisDoc, []types.PrivValidator) {

	validators := make([]types.GenesisValidator, numValidators)
	privValidators := make([]types.PrivValidator, numValidators)
	for i := 0; i < numValidators; i++ {
		val, privVal := RandValidator(randPower, minPower)
		validators[i] = types.GenesisValidator{
			PubKey: val.PubKey,
			Power:  val.VotingPower,
		}
		privValidators[i] = privVal
	}
	sort.Sort(types.PrivValidatorsByAddress(privValidators))

	return &types.GenesisDoc{
		GenesisTime:   tmtime.Now(),
		InitialHeight: 1,
		ChainID:       cfg.ChainID(),
		Validators:    validators,
	}, privValidators
}