package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateValidator = "CreateValidator"

var _ sdk.Msg = &MsgCreateValidator{}

func NewMsgCreateValidator(valAddr sdk.ValAddress, pubKey cryptotypes.PubKey, description Description) (*MsgCreateValidator, error) {
	var pkAny *codectypes.Any
	if pubKey != nil {
		var err error
		if pkAny, err = codectypes.NewAnyWithValue(pubKey); err != nil {
			return nil, err
		}
	}
	return &MsgCreateValidator{
		Description:      description,
		ValidatorAddress: valAddr.String(),
		Pubkey:           pkAny,
	}, nil
}

func (msg *MsgCreateValidator) Route() string {
	return RouterKey
}

func (msg *MsgCreateValidator) Type() string {
	return TypeMsgCreateValidator
}

func (msg *MsgCreateValidator) GetSigners() []sdk.AccAddress {
	validator, err := sdk.AccAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{validator}
}

func (msg *MsgCreateValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
