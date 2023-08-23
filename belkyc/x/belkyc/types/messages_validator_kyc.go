package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateValidatorKYC = "create_validator_kyc"
	TypeMsgUpdateValidatorKYC = "update_validator_kyc"
	TypeMsgDeleteValidatorKYC = "delete_validator_kyc"
)

var _ sdk.Msg = &MsgCreateValidatorKYC{}

func NewMsgCreateValidatorKYC(
	creator string,
	address string,
	value bool,

) *MsgCreateValidatorKYC {
	return &MsgCreateValidatorKYC{
		Creator: creator,
		Address: address,
		Value:   value,
	}
}

func (msg *MsgCreateValidatorKYC) Route() string {
	return RouterKey
}

func (msg *MsgCreateValidatorKYC) Type() string {
	return TypeMsgCreateValidatorKYC
}

func (msg *MsgCreateValidatorKYC) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateValidatorKYC) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateValidatorKYC) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateValidatorKYC{}

func NewMsgUpdateValidatorKYC(
	creator string,
	address string,
	value bool,

) *MsgUpdateValidatorKYC {
	return &MsgUpdateValidatorKYC{
		Creator: creator,
		Address: address,
		Value:   value,
	}
}

func (msg *MsgUpdateValidatorKYC) Route() string {
	return RouterKey
}

func (msg *MsgUpdateValidatorKYC) Type() string {
	return TypeMsgUpdateValidatorKYC
}

func (msg *MsgUpdateValidatorKYC) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateValidatorKYC) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateValidatorKYC) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteValidatorKYC{}

func NewMsgDeleteValidatorKYC(
	creator string,
	address string,

) *MsgDeleteValidatorKYC {
	return &MsgDeleteValidatorKYC{
		Creator: creator,
		Address: address,
	}
}
func (msg *MsgDeleteValidatorKYC) Route() string {
	return RouterKey
}

func (msg *MsgDeleteValidatorKYC) Type() string {
	return TypeMsgDeleteValidatorKYC
}

func (msg *MsgDeleteValidatorKYC) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteValidatorKYC) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteValidatorKYC) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
