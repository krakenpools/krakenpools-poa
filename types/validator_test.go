package types

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var (
	pk1      = ed25519.GenPrivKey().PubKey()
	pk2      = ed25519.GenPrivKey().PubKey()
	valAddr1 = sdk.ValAddress(pk1.Address())
	valAddr2 = sdk.ValAddress(pk2.Address())
)

func TestValidatorTestEquivalent(t *testing.T) {

	val1 := newValidator(t, valAddr1, pk1)
	val2 := newValidator(t, valAddr1, pk1)

	require.Equal(t, val1, val2)

	val2 = newValidator(t, valAddr2, pk2)
	require.NotEqual(t, val1, val2)
}

func TestABCIValidatorUpdate(t *testing.T) {
	validator := newValidator(t, valAddr1, pk1)
	abciVal := validator.ABCIValidatorUpdate()
	pk, err := validator.TmConsPublicKey()
	require.NoError(t, err)
	require.Equal(t, pk, abciVal.PubKey)
	require.Equal(t, int64(1), abciVal.Power)
}

func TestABCIValidatorUpdateZero(t *testing.T) {
	validator := newValidator(t, valAddr1, pk1)
	abciVal := validator.ABCIValidatorUpdateZero()
	pk, err := validator.TmConsPublicKey()
	require.NoError(t, err)
	require.Equal(t, pk, abciVal.PubKey)
	require.Equal(t, int64(0), abciVal.Power)
}

func newValidator(t *testing.T, operator sdk.ValAddress, pubKey cryptotypes.PubKey) Validator {
	v, err := NewValidator(operator, pubKey, Description{Moniker: "a"})
	require.NoError(t, err)
	return v
}
