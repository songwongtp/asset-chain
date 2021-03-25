package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/asset module sentinel errors
var (
	ErrEmptyAddr  = sdkerrors.Register(ModuleName, 2, "empty address")
	ErrEmptyDenom = sdkerrors.Register(ModuleName, 3, "empty asset type")
	ErrInvalidAmt = sdkerrors.Register(ModuleName, 4, "non positive asset amount")

	ErrInvalidVersion      = sdkerrors.Register(ModuleName, 5, "invalid ICS20 version")
	ErrMaxTransferChannels = sdkerrors.Register(ModuleName, 6, "max tranfer channels")
)
