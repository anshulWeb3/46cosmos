package keeper

import (
	"belkyc/x/belkyc/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetValidatorKYC set a specific validatorKYC in the store from its index
func (k Keeper) SetValidatorKYC(ctx sdk.Context, validatorKYC types.ValidatorKYC) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKYCKeyPrefix))
	b := k.cdc.MustMarshal(&validatorKYC)
	store.Set(types.ValidatorKYCKey(
		validatorKYC.Address,
	), b)
}

// GetValidatorKYC returns a validatorKYC from its index
func (k Keeper) GetValidatorKYC(
	ctx sdk.Context,
	address string,

) (val types.ValidatorKYC, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKYCKeyPrefix))

	b := store.Get(types.ValidatorKYCKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveValidatorKYC removes a validatorKYC from the store
func (k Keeper) RemoveValidatorKYC(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKYCKeyPrefix))
	store.Delete(types.ValidatorKYCKey(
		address,
	))
}

// GetAllValidatorKYC returns all validatorKYC
func (k Keeper) GetAllValidatorKYC(ctx sdk.Context) (list []types.ValidatorKYC) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKYCKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ValidatorKYC
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
