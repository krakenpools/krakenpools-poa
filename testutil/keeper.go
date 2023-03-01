package testutil

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/krakenpools/poa/keeper"
	"github.com/krakenpools/poa/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

var (
	PKs = CreateTestPubKeys(500)
)

func PoaKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	return k, ctx
}

// Create a validator for test
func MockValidator(t testing.TB) types.Validator {
	// Junk description
	// validatorDescription := types.Description{
	// 	Moniker:         "Moniker",
	// 	Identity:        "Identity",
	// 	Website:         "Website",
	// 	SecurityContact: "SecurityContact",
	// 	Details:         "Details",
	// }

	valPubKey := PKs[0]
	valAddr := sdk.ValAddress(valPubKey.Address().Bytes())

	validator := NewValidator(t, valAddr, valPubKey)

	return validator
}
