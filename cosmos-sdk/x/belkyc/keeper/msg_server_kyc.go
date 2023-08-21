package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/belkyc/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var admin = "cosmos1vqd2k39umk0t8drn2nzeftv57e3px7sxgnf7ll"

func (k msgServer) CreateKyc(goCtx context.Context, msg *types.MsgCreateKyc) (*types.MsgCreateKycResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != admin {
		return &types.MsgCreateKycResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Creator is not an admin")
	}
	// Check if the value already exists
	_, isFound := k.GetKyc(
		ctx,
		msg.Address,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var kyc = types.Kyc{
		Creator: msg.Creator,
		Address: msg.Address,
		Value:   msg.Value,
	}
	k.SetKyc(
		ctx,
		kyc,
	)
	// Check if the creator has KYC done

	return &types.MsgCreateKycResponse{}, nil
}

func (k msgServer) UpdateKyc(goCtx context.Context, msg *types.MsgUpdateKyc) (*types.MsgUpdateKycResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKyc(
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

	var kyc = types.Kyc{
		Creator: msg.Creator,
		Address: msg.Address,
		Value:   msg.Value,
	}

	k.SetKyc(ctx, kyc)

	return &types.MsgUpdateKycResponse{}, nil
}

func (k msgServer) DeleteKyc(goCtx context.Context, msg *types.MsgDeleteKyc) (*types.MsgDeleteKycResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKyc(
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

	k.RemoveKyc(
		ctx,
		msg.Address,
	)

	return &types.MsgDeleteKycResponse{}, nil
}

func (k msgServer) ChangeAdmin(goCtx context.Context, msg *types.MsgChangeAdmin) (*types.MsgChangeAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != admin {
		return &types.MsgChangeAdminResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Only admin is authorised to change admin")
	}
	admin = msg.Address
	k.ak.ChangeAdmin(ctx, admin)

	fmt.Println("The changed admin address is: ", admin)
	_ = ctx

	return &types.MsgChangeAdminResponse{}, nil
}

func (k Keeper) GetAdmin(goCtx context.Context, req *types.QueryGetAdminRequest) (*types.QueryGetAdminResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx
	
	return &types.QueryGetAdminResponse{Address: admin}, nil
}


