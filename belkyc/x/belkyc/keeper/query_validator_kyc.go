package keeper

import (
	"context"

	"belkyc/x/belkyc/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ValidatorKYCAll(goCtx context.Context, req *types.QueryAllValidatorKYCRequest) (*types.QueryAllValidatorKYCResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var validatorKYCs []types.ValidatorKYC
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	validatorKYCStore := prefix.NewStore(store, types.KeyPrefix(types.ValidatorKYCKeyPrefix))

	pageRes, err := query.Paginate(validatorKYCStore, req.Pagination, func(key []byte, value []byte) error {
		var validatorKYC types.ValidatorKYC
		if err := k.cdc.Unmarshal(value, &validatorKYC); err != nil {
			return err
		}

		validatorKYCs = append(validatorKYCs, validatorKYC)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllValidatorKYCResponse{ValidatorKYC: validatorKYCs, Pagination: pageRes}, nil
}

func (k Keeper) ValidatorKYC(goCtx context.Context, req *types.QueryGetValidatorKYCRequest) (*types.QueryGetValidatorKYCResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetValidatorKYC(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetValidatorKYCResponse{ValidatorKYC: val}, nil
}
