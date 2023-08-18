package keeper

import (
	"context"

	"belkyc/x/belkyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ChangeAdmin(goCtx context.Context, msg *types.MsgChangeAdmin) (*types.MsgChangeAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgChangeAdminResponse{}, nil
}
