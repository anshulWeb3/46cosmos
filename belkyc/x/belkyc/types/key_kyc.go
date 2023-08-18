package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// KycKeyPrefix is the prefix to retrieve all Kyc
	KycKeyPrefix = "Kyc/value/"
)

// KycKey returns the store key to retrieve a Kyc from the index fields
func KycKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
