package keeper

import (
	"belkyc/x/belkyc/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetKyc set a specific kyc in the store from its index
func (k Keeper) SetKyc(ctx sdk.Context, kyc types.Kyc) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KycKeyPrefix))
	b := k.cdc.MustMarshal(&kyc)
	store.Set(types.KycKey(
		kyc.Address,
	), b)
}

// GetKyc returns a kyc from its index
func (k Keeper) GetKyc(
	ctx sdk.Context,
	address string,

) (val types.Kyc, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KycKeyPrefix))

	b := store.Get(types.KycKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKyc removes a kyc from the store
func (k Keeper) RemoveKyc(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KycKeyPrefix))
	store.Delete(types.KycKey(
		address,
	))
}

// GetAllKyc returns all kyc
func (k Keeper) GetAllKyc(ctx sdk.Context) (list []types.Kyc) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KycKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Kyc
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
