package keeper_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	test "github.com/krakenpools/poa/testutil"
)

func TestValidator(t *testing.T) {
	k, ctx := test.PoaKeeper(t)
	validator := test.MockValidator(t)

	k.SetValidator(ctx, validator)

	operator := validator.GetOperator()
	retrievedValidator, found := k.GetValidator(ctx, operator)
	if !found {
		t.Errorf("GetValidator should find validator if it has been set")
	}

	if !cmp.Equal(validator, retrievedValidator) {
		t.Errorf("GetValidator should find %v, found %v", validator, retrievedValidator)
	}
}
