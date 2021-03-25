package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/songwongtp/asset-chain/x/asset/types"
)

type (
	// Keeper of asset module
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		channelKeeper types.ChannelKeeper
		portKeeper    types.PortKeeper
		authKeeper    types.AccountKeeper
		bankKeeper    types.BankKeeper
		scopedKeeper  capabilitykeeper.ScopedKeeper
	}
)

// NewKeeper returns asset keeper give parameters
func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey,
	channelKeeper types.ChannelKeeper, portKeeper types.PortKeeper,
	authKeeper types.AccountKeeper, bankKeeper types.BankKeeper, scopedKeeper capabilitykeeper.ScopedKeeper,
) Keeper {

	// ensure ibc asset module account is set
	if addr := authKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the IBC asset module account has not been set")
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		authKeeper:    authKeeper,
		bankKeeper:    bankKeeper,
		scopedKeeper:  scopedKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

//____________________________________________________________________________
// IBC functions

// ChanCloseInit defines a wrapper function for the channel keeper's function
// in order to expose it to the ICS20 transfer handler.
func (k Keeper) ChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	capName := host.ChannelCapabilityPath(portID, channelID)
	chanCap, ok := k.scopedKeeper.GetCapability(ctx, capName)
	if !ok {
		return sdkerrors.Wrapf(channeltypes.ErrChannelCapabilityNotFound, "could not retrieve channel capability at : %s", capName)
	}
	return k.channelKeeper.ChanCloseInit(ctx, portID, channelID, chanCap)
}

// IsBound checks if the asset module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.scopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}

// BindPort defines a wrapper function for the port Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.portKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// GetPort returns the portID for the asset module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the asset module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PortKey, []byte(portID))
}

// AuthenticateCapablity wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	return k.scopedKeeper.AuthenticateCapability(ctx, cap, name)
}

// ClaimCapability allows the asset module that can claim a capability that IBC module
// passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

//____________________________________________________________________________

// Buy exchanges coins from the given addr with the denom asset
func (k Keeper) Buy(ctx sdk.Context, order types.AssetOrder) error {
	accAddr, err := sdk.AccAddressFromBech32(order.Addr)
	if err != nil {
		err := sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "buyer address is invalid: %w", err)
		order.Status = types.StatusOrderFail
		order.StatusLog = err.Error()
		k.SetOrder(ctx, order)
		return err
	}

	baseDenomAmt := sdk.NewCoin(types.BaseDenom, sdk.NewIntFromUint64(order.Amount*order.PricePerUnit))
	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accAddr, types.ModuleName, sdk.NewCoins(baseDenomAmt)); err != nil {
		err := sdkerrors.Wrap(err, "fail to transfer coins to the asset module")
		order.Status = types.StatusOrderFail
		order.StatusLog = err.Error()
		k.SetOrder(ctx, order)
		return err
	}

	denomAmt := sdk.NewCoin(order.Denom, sdk.NewIntFromUint64(order.Amount))
	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddr, sdk.NewCoins(denomAmt)); err != nil {
		err := sdkerrors.Wrap(err, "fail to transfer asset to the buyer")
		order.Status = types.StatusOrderFail
		order.StatusLog = err.Error()
		k.SetOrder(ctx, order)
		return err
	}

	order.Status = types.StatusOrderSuccess
	k.SetOrder(ctx, order)
	return nil
}

// Sell exchanges asset from the given addr with coins
func (k Keeper) Sell(ctx sdk.Context, order types.AssetOrder) error {
	accAddr, err := sdk.AccAddressFromBech32(order.Addr)
	if err != nil {
		err := sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "buyer address is invalid: %w", err)
		order.Status = types.StatusOrderFail
		order.StatusLog = err.Error()
		k.SetOrder(ctx, order)
		return err
	}

	denomAmt := sdk.NewCoin(order.Denom, sdk.NewIntFromUint64(order.Amount))
	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddr, sdk.NewCoins(denomAmt)); err != nil {
		err := sdkerrors.Wrap(err, "fail to transfer coins to the seller")
		order.Status = types.StatusOrderFail
		order.StatusLog = err.Error()
		k.SetOrder(ctx, order)
		return err
	}

	baseDenomAmt := sdk.NewCoin(types.BaseDenom, sdk.NewIntFromUint64(order.Amount*order.PricePerUnit))
	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accAddr, types.ModuleName, sdk.NewCoins(baseDenomAmt)); err != nil {
		err := sdkerrors.Wrap(err, "fail to transfer assets to the asset module")
		order.Status = types.StatusOrderFail
		order.StatusLog = err.Error()
		k.SetOrder(ctx, order)
		return err
	}

	order.Status = types.StatusOrderSuccess
	k.SetOrder(ctx, order)
	return nil
}

// SetOracleScriptID sets the oracle script ID for the given asset denom
func (k Keeper) SetOracleScriptID(ctx sdk.Context, denom string, oracleScriptID uint64) {
	store := ctx.KVStore(k.storeKey)
	oracleScriptIDStore := prefix.NewStore(store, types.OracleScriptKey)

	key := types.KeyPrefix(denom)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, oracleScriptID)
	oracleScriptIDStore.Set(key, bz)
}

// AddSupply adds supply amount of the denom asset
func (k Keeper) AddSupply(ctx sdk.Context, denom string, supply uint64) error {
	store := ctx.KVStore(k.storeKey)
	oracleScriptIDStore := prefix.NewStore(store, types.OracleScriptKey)
	bz := oracleScriptIDStore.Get(types.KeyPrefix(denom))
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
	oracleScriptIDStore := prefix.NewStore(store, types.OracleScriptKey)
	bz := oracleScriptIDStore.Get(types.KeyPrefix(denom))
	if bz == nil {
		panic(sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not a valid asset type", denom))
	}
	oracleScriptID := uint64(binary.BigEndian.Uint64(bz))

	totalSupply := k.bankKeeper.GetSupply(ctx).GetTotal().AmountOf(denom).Uint64()

	return types.Asset{
		Denom:          denom,
		TotalSupply:    totalSupply,
		OracleScriptId: oracleScriptID,
	}
}

// GetAllAssetDenoms returns all registered asset denoms
func (k Keeper) GetAllAssetDenoms(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	oracleScriptIDStore := prefix.NewStore(store, types.OracleScriptKey)
	iterator := oracleScriptIDStore.Iterator(nil, nil)
	defer iterator.Close()

	denoms := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		denoms = append(denoms, string(iterator.Key()))
	}
	return denoms
}

// GetAllAssetInfos returns all asset infos
func (k Keeper) GetAllAssetInfos(ctx sdk.Context) []types.Asset {
	denoms := k.GetAllAssetDenoms(ctx)

	assets := make([]types.Asset, 0)
	for _, denom := range denoms {
		assets = append(assets, k.GetAssetInfo(ctx, denom))
	}

	return assets
}
