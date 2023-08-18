package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cosmos/cosmos-sdk/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/x/belkyc/keeper"
	"github.com/cosmos/cosmos-sdk/x/belkyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BelkycKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
