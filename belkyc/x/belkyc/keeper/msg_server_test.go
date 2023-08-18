package keeper_test

import (
	"context"
	"testing"

	keepertest "belkyc/testutil/keeper"
	"belkyc/x/belkyc/keeper"
	"belkyc/x/belkyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BelkycKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
