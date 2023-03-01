package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName is the name of the staking module
	ModuleName = "poa"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// RouterKey is the msg router key for the staking module
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_poa"
)

var (
	ValidatorsKey           = []byte{0x21}
	ValidatorsByConsAddrKey = []byte{0x22}
)

func GetValidatorKey(operatorAddr sdk.ValAddress) []byte {
	return append(ValidatorsKey, address.MustLengthPrefix(operatorAddr)...)
}

func GetValidatorByConsAddrKey(addr sdk.ConsAddress) []byte {
	return append(ValidatorsByConsAddrKey, address.MustLengthPrefix(addr)...)
}
