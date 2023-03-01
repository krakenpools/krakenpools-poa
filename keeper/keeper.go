package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/krakenpools/poa/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	storeKey storetypes.StoreKey
	memKey   storetypes.StoreKey
	cdc      codec.BinaryCodec
}

func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	memKey storetypes.StoreKey,
) *Keeper {

	return &Keeper{
		storeKey: key,
		cdc:      cdc,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}
