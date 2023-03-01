package testutil

import (
	"testing"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/krakenpools/poa/types"
	"github.com/stretchr/testify/require"
)

func NewValidator(t testing.TB, operator sdk.ValAddress, pubKey cryptotypes.PubKey) types.Validator {
	v, err := types.NewValidator(operator, pubKey, types.Description{})
	require.NoError(t, err)
	return v
}
