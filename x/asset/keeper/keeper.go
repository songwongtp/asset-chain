package keeper

import (
	"fmt"

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
func (k Keeper) Buy (ctx sdk.Context, addr sdk.AccAddress, denom string, amount uint64) error {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefix(denom))
	if bz == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not a valid asset type", denom)
	}

	var assetInfo types.Asset
	k.cdc.MustUnmarshalBinaryBare(bz, &assetInfo)

	baseDenom, err := sdk.GetBaseDenom()
	if err != nil {
		return err
	}
	baseDenomAmt := sdk.NewCoin(baseDenom, sdk.NewIntFromUint64(amount * assetInfo.Price))
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
func (k Keeper) Sell (ctx sdk.Context, addr sdk.AccAddress, denom string, amount uint64) error {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefix(denom))
	if bz == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not a valid asset type", denom)
	}

	var assetInfo types.Asset
	k.cdc.MustUnmarshalBinaryBare(bz, &assetInfo)

	denomAmt := sdk.NewCoin(denom, sdk.NewIntFromUint64(amount))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(denomAmt)); err != nil {
		return err
	}

	baseDenom, err := sdk.GetBaseDenom()
	if err != nil {
		return err
	}
	baseDenomAmt := sdk.NewCoin(baseDenom, sdk.NewIntFromUint64(amount * assetInfo.Price))
	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(baseDenomAmt)); err != nil {
		return err
	}

	return nil
}

func (k Keeper) setAsset(ctx sdk.Context, assetInfo types.Asset) {
	store := ctx.KVStore(k.storeKey)

	key := types.KeyPrefix(assetInfo.Denom)
	info := k.cdc.MustMarshalBinaryBare(&assetInfo)
	store.Set(key, info)
}