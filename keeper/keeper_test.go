package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/krakenpools/poa/types"
)

func TestKeeper(t *testing.T) {
	keys := sdk.NewKVStoreKeys(types.StoreKey)
	encCfg := moduletestutil.MakeTestEncodingConfig()

	NewKeeper(encCfg.Codec, keys[types.StoreKey], keys[types.MemStoreKey])
}
