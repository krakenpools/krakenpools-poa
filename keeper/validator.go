package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/krakenpools/poa/types"
)

func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetValidatorKey(addr))
	if value == nil {
		return validator, false
	}

	validator = types.MustUnmarshalValidator(k.cdc, value)
	return validator, true
}

// Get the set of all validators
func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		validator := types.MustUnmarshalValidator(k.cdc, iterator.Value())
		validators = append(validators, validator)
	}

	return validators
}

func (k Keeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)

	opAddr := store.Get(types.GetValidatorByConsAddrKey(consAddr))
	if opAddr == nil {
		return validator, false
	}

	return k.GetValidator(ctx, opAddr)
}

func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalValidator(k.cdc, &validator)
	store.Set(types.GetValidatorKey(validator.GetOperator()), bz)
}

func (k Keeper) SetValidatorByConsAddr(ctx sdk.Context, validator types.Validator) error {
	consPk, err := validator.GetConsAddr()
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetValidatorByConsAddrKey(consPk), validator.GetOperator())
	return nil
}

// func (k Keeper) AppendPost(ctx sdk.Context, validator types.Validator) uint64 {
// 	count := k.GetPostCount(ctx)
// 	validator.Id = count
// 	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
// 	appendedValue := k.cdc.MustMarshal(&validator)
// 	store.Set(GetPostIDBytes(validator.Id), appendedValue)
// 	k.SetPostCount(ctx, count+1)
// 	return count
// }
