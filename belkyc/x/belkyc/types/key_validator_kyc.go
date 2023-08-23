package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ValidatorKYCKeyPrefix is the prefix to retrieve all ValidatorKYC
	ValidatorKYCKeyPrefix = "ValidatorKYC/value/"
)

// ValidatorKYCKey returns the store key to retrieve a ValidatorKYC from the index fields
func ValidatorKYCKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
