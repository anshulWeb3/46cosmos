package keeper

import (
	"context"

	"belkyc/x/belkyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateValidatorKYC(goCtx context.Context, msg *types.MsgCreateValidatorKYC) (*types.MsgCreateValidatorKYCResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetValidatorKYC(
		ctx,
		msg.Address,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var validatorKYC = types.ValidatorKYC{
		Creator: msg.Creator,
		Address: msg.Address,
		Value:   msg.Value,
	}

	k.SetValidatorKYC(
		ctx,
		validatorKYC,
	)
	return &types.MsgCreateValidatorKYCResponse{}, nil
}

func (k msgServer) UpdateValidatorKYC(goCtx context.Context, msg *types.MsgUpdateValidatorKYC) (*types.MsgUpdateValidatorKYCResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetValidatorKYC(
		ctx,
		msg.Address,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var validatorKYC = types.ValidatorKYC{
		Creator: msg.Creator,
		Address: msg.Address,
		Value:   msg.Value,
	}

	k.SetValidatorKYC(ctx, validatorKYC)

	return &types.MsgUpdateValidatorKYCResponse{}, nil
}

func (k msgServer) DeleteValidatorKYC(goCtx context.Context, msg *types.MsgDeleteValidatorKYC) (*types.MsgDeleteValidatorKYCResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetValidatorKYC(
		ctx,
		msg.Address,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveValidatorKYC(
		ctx,
		msg.Address,
	)

	return &types.MsgDeleteValidatorKYCResponse{}, nil
}
