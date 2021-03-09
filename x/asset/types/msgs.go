package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// asset message types
const (
	TypeMsgBuyAsset		= "buy_asset"
	TypeMsgSellAsset	= "sell_asset"
	TypeMsgSetPrice		= "set_price"
	TypeMsgAddSupply	= "add_supply"
)

var (
	_ sdk.Msg = &MsgBuyAsset{}
	_ sdk.Msg = &MsgSellAsset{}
	_ sdk.Msg = &MsgSetPrice{}
	_ sdk.Msg = &MsgAddSupply{}
)

// NewMsgBuyAsset creates a new MsgBuyAsset instance
func NewMsgBuyAsset(buyer string, denom string, amount uint64) *MsgBuyAsset {
	return &MsgBuyAsset{
		Buyer:	buyer,
		Denom:	denom,
		Amount:	amount,
	}
}

// Route implements sdk.Msg interface
func (msg MsgBuyAsset) Route() string { return RouterKey }

// Type implements sdk.Msg interface
func (msg MsgBuyAsset) Type() string { return TypeMsgBuyAsset }

// GetSigners implements sdk.Msg interface. It returns address(es) that
// must sign over msg.GetSignBytes()
func (msg MsgBuyAsset) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// GetSignBytes returns the message bytes to sign over
func (msg MsgBuyAsset) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements sdk.Msg interface
func (msg MsgBuyAsset) ValidateBasic() error {
	addr, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return err
	}
	if addr.Empty() {
		return ErrEmptyAddr
	}

	if msg.Denom == "" {
		return ErrEmptyDenom
	}

	if msg.Amount <= 0 {
		return ErrInvalidAmt
	}

	return nil
}


// NewMsgSellAsset creates a new MsgSellAsset instance
func NewMsgSellAsset(seller string, denom string, amount uint64) *MsgSellAsset {
	return &MsgSellAsset {
		Seller: seller,
		Denom:	denom,
		Amount:	amount,
	}
}

// Route implements sdk.Msg interface
func (msg MsgSellAsset) Route() string { return RouterKey }

// Type implements sdk.Msg interface
func (msg MsgSellAsset) Type() string { return TypeMsgSellAsset }

// GetSigners implements sdk.Msg interface. It return the address(es) that
// must sign over msg.GetSignBytes()
func (msg MsgSellAsset) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// GetSignBytes returns the message byte to sign over
func (msg MsgSellAsset) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface
func (msg MsgSellAsset) ValidateBasic() error {
	addr, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		return err
	}
	if addr.Empty() {
		return ErrEmptyAddr
	}

	if msg.Denom == "" {
		return ErrEmptyDenom
	}

	if msg.Amount <= 0 {
		return ErrInvalidAmt
	}

	return nil
}


// NewMsgSetPrice creates a new MsgSetPrice instance
func NewMsgSetPrice(addr string, denom string, price uint64) *MsgSetPrice {
	return &MsgSetPrice {
		Addr: addr,
		Denom: denom,
		Price: price,
	}
}

// Route implements sdk.Msg interface
func (msg MsgSetPrice) Route() string { return RouterKey }

// Type implements sdk.Msg interface
func (msg MsgSetPrice) Type() string { return TypeMsgSetPrice }

// GetSigners implements sdk.Msg interface. It returns address(es) that
// must sign over msg.GetSignBytes()
func (msg MsgSetPrice) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Addr)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// GetSignBytes returns the message byte to sign over
func (msg MsgSetPrice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements sdk.Msg interface
func (msg MsgSetPrice) ValidateBasic() error {
	if msg.Denom == "" {
		return ErrEmptyDenom
	}

	if msg.Price <= 0 {
		return ErrInvalidAmt
	}

	return nil
}


// NewMsgAddSupply creates a new MsgAddSupply instance
func NewMsgAddSupply(addr string, denom string, amount uint64) *MsgAddSupply {
	return &MsgAddSupply {
		Addr: addr,
		Denom: denom,
		Amount: amount,
	}
}

// Route implements sdk.Msg interface
func (msg MsgAddSupply) Route() string { return RouterKey }

// Type implments sdk.Msg interface
func (msg MsgAddSupply) Type() string { return TypeMsgAddSupply }

// GetSigners implments sdk.Msg interface. It returns address(es) that
// must sign over msg.GetSignBytes()
func (msg MsgAddSupply) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Addr)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// GetSignBytes returns the message byte to sign over
func (msg MsgAddSupply) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements sdk.Msg interface
func (msg MsgAddSupply) ValidateBasic() error {
	if msg.Denom == "" {
		return ErrEmptyDenom
	}

	if msg.Amount <= 0 {
		return ErrInvalidAmt
	}

	return nil
}