package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "belkyc/testutil/keeper"
	"belkyc/x/belkyc/keeper"
	"belkyc/x/belkyc/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestValidatorKYCMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.BelkycKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateValidatorKYC{Creator: creator,
			Address: strconv.Itoa(i),
		}
		_, err := srv.CreateValidatorKYC(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetValidatorKYC(ctx,
			expected.Address,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestValidatorKYCMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateValidatorKYC
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateValidatorKYC{Creator: creator,
				Address: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateValidatorKYC{Creator: "B",
				Address: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateValidatorKYC{Creator: creator,
				Address: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BelkycKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateValidatorKYC{Creator: creator,
				Address: strconv.Itoa(0),
			}
			_, err := srv.CreateValidatorKYC(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateValidatorKYC(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetValidatorKYC(ctx,
					expected.Address,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestValidatorKYCMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteValidatorKYC
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteValidatorKYC{Creator: creator,
				Address: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteValidatorKYC{Creator: "B",
				Address: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteValidatorKYC{Creator: creator,
				Address: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BelkycKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateValidatorKYC(wctx, &types.MsgCreateValidatorKYC{Creator: creator,
				Address: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteValidatorKYC(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetValidatorKYC(ctx,
					tc.request.Address,
				)
				require.False(t, found)
			}
		})
	}
}
