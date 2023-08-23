package keeper_test

import (
	"strconv"
	"testing"

	keepertest "belkyc/testutil/keeper"
	"belkyc/testutil/nullify"
	"belkyc/x/belkyc/keeper"
	"belkyc/x/belkyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNValidatorKYC(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ValidatorKYC {
	items := make([]types.ValidatorKYC, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetValidatorKYC(ctx, items[i])
	}
	return items
}

func TestValidatorKYCGet(t *testing.T) {
	keeper, ctx := keepertest.BelkycKeeper(t)
	items := createNValidatorKYC(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetValidatorKYC(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestValidatorKYCRemove(t *testing.T) {
	keeper, ctx := keepertest.BelkycKeeper(t)
	items := createNValidatorKYC(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValidatorKYC(ctx,
			item.Address,
		)
		_, found := keeper.GetValidatorKYC(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestValidatorKYCGetAll(t *testing.T) {
	keeper, ctx := keepertest.BelkycKeeper(t)
	items := createNValidatorKYC(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllValidatorKYC(ctx)),
	)
}
