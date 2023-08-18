package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/cosmos/cosmos-sdk/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/testutil/nullify"
	"github.com/cosmos/cosmos-sdk/x/belkyc/keeper"
	"github.com/cosmos/cosmos-sdk/x/belkyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNKyc(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Kyc {
	items := make([]types.Kyc, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetKyc(ctx, items[i])
	}
	return items
}

func TestKycGet(t *testing.T) {
	keeper, ctx := keepertest.BelkycKeeper(t)
	items := createNKyc(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetKyc(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestKycRemove(t *testing.T) {
	keeper, ctx := keepertest.BelkycKeeper(t)
	items := createNKyc(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveKyc(ctx,
			item.Address,
		)
		_, found := keeper.GetKyc(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestKycGetAll(t *testing.T) {
	keeper, ctx := keepertest.BelkycKeeper(t)
	items := createNKyc(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllKyc(ctx)),
	)
}
