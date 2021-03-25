package types

import (
	"crypto/sha256"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "asset"

	// Version dfines the current version the IBC asset
	// module supports
	Version = "ics20-1"

	// PortID is the default port id that asset module binds to
	PortID = "asset"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	// base denom for buying/selling asset
	BaseDenom = "uusd"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey         = []byte{0x01}
	BaseDenomKey    = []byte{0x00}
	OracleScriptKey = []byte{0x10}
	OrderKey        = []byte{0x11}
)

// GetEscrowAddress reurns the escrow address for the specified channel.
// The escrow address follows the format as outlined in ADR 028:
// https://github.com/cosmos/cosmos-sdk/blob/master/docs/architecture/adr-028-public-key-addresses.md
func GetEscrowAddress(portID, channelID string) sdk.AccAddress {
	// a slash is used to create domain separation between port and channel identifiers to
	// prevent address collisions between escrow addresses created for different channels
	contents := fmt.Sprintf("%s/%s", portID, channelID)

	// ADR 028 AddressHash construction
	preImage := []byte(Version)
	preImage = append(preImage, 0)
	preImage = append(preImage, contents...)
	hash := sha256.Sum256(preImage)
	return hash[:20]
}

// KeyPrefix converts string to KVStore's key
func KeyPrefix(p string) []byte {
	return []byte(p)
}
