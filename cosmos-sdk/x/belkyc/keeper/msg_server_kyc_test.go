package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/cosmos/cosmos-sdk/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/x/belkyc/keeper"
	"github.com/cosmos/cosmos-sdk/x/belkyc/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestKycMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.BelkycKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateKyc{Creator: creator,
			Address: strconv.Itoa(i),
		}
		_, err := srv.CreateKyc(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetKyc(ctx,
			expected.Address,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestKycMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateKyc
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateKyc{Creator: creator,
				Address: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateKyc{Creator: "B",
				Address: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateKyc{Creator: creator,
				Address: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BelkycKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateKyc{Creator: creator,
				Address: strconv.Itoa(0),
			}
			_, err := srv.CreateKyc(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateKyc(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetKyc(ctx,
					expected.Address,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestKycMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteKyc
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteKyc{Creator: creator,
				Address: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteKyc{Creator: "B",
				Address: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteKyc{Creator: creator,
				Address: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BelkycKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateKyc(wctx, &types.MsgCreateKyc{Creator: creator,
				Address: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteKyc(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetKyc(ctx,
					tc.request.Address,
				)
				require.False(t, found)
			}
		})
	}
}
