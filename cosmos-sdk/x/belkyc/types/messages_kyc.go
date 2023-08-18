package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateKyc = "create_kyc"
	TypeMsgUpdateKyc = "update_kyc"
	TypeMsgDeleteKyc = "delete_kyc"
)

var _ sdk.Msg = &MsgCreateKyc{}

func NewMsgCreateKyc(
	creator string,
	address string,
	value bool,

) *MsgCreateKyc {
	return &MsgCreateKyc{
		Creator: creator,
		Address: address,
		Value:   value,
	}
}

func (msg *MsgCreateKyc) Route() string {
	return RouterKey
}

func (msg *MsgCreateKyc) Type() string {
	return TypeMsgCreateKyc
}

func (msg *MsgCreateKyc) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateKyc) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateKyc) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateKyc{}

func NewMsgUpdateKyc(
	creator string,
	address string,
	value bool,

) *MsgUpdateKyc {
	return &MsgUpdateKyc{
		Creator: creator,
		Address: address,
		Value:   value,
	}
}

func (msg *MsgUpdateKyc) Route() string {
	return RouterKey
}

func (msg *MsgUpdateKyc) Type() string {
	return TypeMsgUpdateKyc
}

func (msg *MsgUpdateKyc) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateKyc) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateKyc) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteKyc{}

func NewMsgDeleteKyc(
	creator string,
	address string,

) *MsgDeleteKyc {
	return &MsgDeleteKyc{
		Creator: creator,
		Address: address,
	}
}
func (msg *MsgDeleteKyc) Route() string {
	return RouterKey
}

func (msg *MsgDeleteKyc) Type() string {
	return TypeMsgDeleteKyc
}

func (msg *MsgDeleteKyc) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteKyc) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteKyc) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
