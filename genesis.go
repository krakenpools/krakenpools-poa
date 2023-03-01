package poa

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/krakenpools/poa/keeper"
	"github.com/krakenpools/poa/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// WriteValidators returns a slice of bonded genesis validators.
// func WriteValidators(ctx sdk.Context, keeper *keeper.Keeper) (vals []tmtypes.GenesisValidator, returnErr error) {
// 	keeper.IterateLastValidators(ctx, func(_ int64, validator types.ValidatorI) (stop bool) {
// 		pk, err := validator.ConsPubKey()
// 		if err != nil {
// 			returnErr = err
// 			return true
// 		}
// 		tmPk, err := cryptocodec.ToTmPubKeyInterface(pk)
// 		if err != nil {
// 			returnErr = err
// 			return true
// 		}

// 		vals = append(vals, tmtypes.GenesisValidator{
// 			Address: sdk.ConsAddress(tmPk.Address()).Bytes(),
// 			PubKey:  tmPk,
// 			Power:   validator.GetConsensusPower(keeper.PowerReduction(ctx)),
// 			Name:    validator.GetMoniker(),
// 		})

// 		return false
// 	})

// 	return
// }

// ValidateGenesis validates the provided staking genesis state to ensure the
// expected invariants holds. (i.e. params in correct bounds, no duplicate validators)
func ValidateGenesis(data *types.GenesisState) error {
	if err := validateGenesisStateValidators(data.Validators); err != nil {
		return err
	}

	return nil
}

func validateGenesisStateValidators(validators []types.Validator) error {
	addrMap := make(map[string]bool, len(validators))

	for i := 0; i < len(validators); i++ {
		val := validators[i]
		consPk, err := val.ConsPubKey()
		if err != nil {
			return err
		}

		strKey := string(consPk.Bytes())

		if _, ok := addrMap[strKey]; ok {
			consAddr, err := val.GetConsAddr()
			if err != nil {
				return err
			}
			return fmt.Errorf("duplicate validator in genesis state: moniker %v, address %v", val.Description.Moniker, consAddr)
		}

		addrMap[strKey] = true
	}

	return nil
}

func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) (res []abci.ValidatorUpdate) {

	for _, validator := range data.Validators {
		k.SetValidator(ctx, validator)
		k.SetValidatorByConsAddr(ctx, validator)
		// res = append(res, validator.ABCIValidatorUpdate())
	}

	return res
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data types.GenesisState) {
	return types.GenesisState{
		Validators: k.GetAllValidators(ctx),
	}
}
