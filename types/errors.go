package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrValidatorOwnerExists            = sdkerrors.Register(ModuleName, 4, "validator already exist for this operator address; must use new validator operator address")
	ErrValidatorPubKeyExists           = sdkerrors.Register(ModuleName, 5, "validator already exist for this pubkey; must use new validator pubkey")
	ErrValidatorPubKeyTypeNotSupported = sdkerrors.Register(ModuleName, 6, "validator pubkey type is not supported")
)
