package keeper

import (
	"fmt"
	"encoding/binary"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/songwongtp/asset-chain/x/asset/types"
)

type (
	// Keeper of asset module
	Keeper struct {
		cdc			codec.Marshaler
		storeKey	sdk.StoreKey
		memKey		sdk.StoreKey
		bankKeeper	types.BankKeeper
	}
)

// NewKeeper returns asset keeper give parameters
func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey, bankKeeper types.BankKeeper) *Keeper {
	return &Keeper{
		cdc:		cdc,
		storeKey:	storeKey,
		memKey:		memKey,
		bankKeeper: bankKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Buy exchanges coins from the given addr with the denom asset
func (k Keeper) Buy(ctx sdk.Context, addr sdk.AccAddress, denom string, amount uint64) error {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefix(denom))
	if bz == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not a valid asset type", denom)
	}
	price := uint64(binary.BigEndian.Uint64(bz))

	baseDenom, err := sdk.GetBaseDenom()
	if err != nil {
		return err
	}
	baseDenomAmt := sdk.NewCoin(baseDenom, sdk.NewIntFromUint64(amount * price))
	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(baseDenomAmt)); err != nil {
		return err
	}

	denomAmt := sdk.NewCoin(denom, sdk.NewIntFromUint64(amount))
	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(denomAmt)); err != nil {
		return err
	}

	return nil
}

// Sell exchanges asset from the given addr with coins
func (k Keeper) Sell(ctx sdk.Context, addr sdk.AccAddress, denom string, amount uint64) error {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefix(denom))
	if bz == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not a valid asset type", denom)
	}
	price := uint64(binary.BigEndian.Uint64(bz))

	denomAmt := sdk.NewCoin(denom, sdk.NewIntFromUint64(amount))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(denomAmt)); err != nil {
		return err
	}

	baseDenom, err := sdk.GetBaseDenom()
	if err != nil {
		return err
	}
	baseDenomAmt := sdk.NewCoin(baseDenom, sdk.NewIntFromUint64(amount * price))
	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(baseDenomAmt)); err != nil {
		return err
	}

	return nil
}

// SetPrice sets the price per unit for the given asset denom
func (k Keeper) SetPrice(ctx sdk.Context, denom string, price uint64) {
	store := ctx.KVStore(k.storeKey)

	key := types.KeyPrefix(denom)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, price)
	store.Set(key, bz)
}

// AddSupply adds supply amount of the denom asset
func (k Keeper) AddSupply(ctx sdk.Context, denom string, supply uint64) error {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefix(denom))
	if bz == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not a valid asset type", denom)
	}

	amt := sdk.NewCoin(denom, sdk.NewIntFromUint64(supply))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amt)); err != nil {
		return err
	}

	return nil
}

// GetAssetInfo returns the info of the given denom in Asset type
func (k Keeper) GetAssetInfo(ctx sdk.Context, denom string) types.Asset {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefix(denom))
	if bz == nil {
		panic(sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not a valid asset type", denom))
	}
	price := uint64(binary.BigEndian.Uint64(bz))
	
	totalSupply := k.bankKeeper.GetSupply(ctx, denom).Amount.Uint64()

	return types.Asset {
		Denom: denom,
		TotalSupply: totalSupply,
		Price: price,
	}
}

// GetAllAssetDenoms returns all registered asset denoms
func (k Keeper) GetAllAssetDenoms(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	
	denoms := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		denoms = append(denoms, string(iterator.Key()))
	}
	return denoms
}