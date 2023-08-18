package types

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "belkyc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_belkyc"
)

func StoreKeyval() storetypes.StoreKey {
	return storetypes.NewKVStoreKey(ModuleName)

}
func KeyPrefix(p string) []byte {
	// fmt.Println("I am in key prefix -------------------------")
	return []byte(p)
}
