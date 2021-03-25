package keeper

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"strings"
	"time"

	"github.com/bandprotocol/chain/pkg/obi"
	oracletypes "github.com/bandprotocol/chain/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/songwongtp/asset-chain/x/asset/types"
)

func (k Keeper) SetOrder(ctx sdk.Context, order types.AssetOrder) {
	store := ctx.KVStore(k.storeKey)
	orderStore := prefix.NewStore(store, types.OrderKey)

	key := types.KeyPrefix(order.OrderId)
	value := k.cdc.MustMarshalBinaryBare(&order)

	orderStore.Set(key, value)
}

func (k Keeper) GetOrder(ctx sdk.Context, orderID string) types.AssetOrder {
	store := ctx.KVStore(k.storeKey)
	orderStore := prefix.NewStore(store, types.OrderKey)

	key := types.KeyPrefix(orderID)
	var order types.AssetOrder
	k.cdc.MustUnmarshalBinaryBare(orderStore.Get(key), &order)

	return order
}

func (k Keeper) CreatePendingOrder(ctx sdk.Context, orderType string, addr sdk.AccAddress, denom string, amount uint64) string {
	transactionHash := sha256.Sum256(ctx.TxBytes())
	order := types.NewOrderAsset(
		hex.EncodeToString(transactionHash[:]),
		orderType,
		addr.String(),
		denom,
		amount,
	)
	k.SetOrder(ctx, order)

	return order.OrderId
}

func (k Keeper) ProcessOrder(ctx sdk.Context, orderID string, AssetPricePerUnit uint64) error {
	order := k.GetOrder(ctx, orderID)
	order.PricePerUnit = AssetPricePerUnit
	switch order.OrderType {
	case types.TypeOrderBuy:
		return k.Buy(ctx, order)
	case types.TypeOrderSell:
		return k.Sell(ctx, order)
	}
	return nil
}

func (k Keeper) RequestAssetPrice(ctx sdk.Context, orderID string, denom string, sourceChannel string) error {
	store := ctx.KVStore(k.storeKey)
	oracleScriptIDStore := prefix.NewStore(store, types.OracleScriptKey)
	bz := oracleScriptIDStore.Get(types.KeyPrefix(denom))
	if bz == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not a valid asset type", denom)
	}
	oracleScriptID := uint64(binary.BigEndian.Uint64(bz))

	sourcePort := types.PortID
	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown channel %s port %s",
			sourceChannel, sourcePort,
		)
	}
	destinationPort := sourceChannelEnd.Counterparty.PortId
	destinationChannel := sourceChannelEnd.Counterparty.ChannelId
	sequence, found := k.channelKeeper.GetNextSequenceSend(
		ctx, sourcePort, sourceChannel,
	)
	if !found {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown sequence number for channel %s port %s",
			sourceChannel, sourcePort,
		)
	}
	clientID := strings.Join([]string{types.ModuleName, orderID}, ":")
	oracleScript := oracletypes.OracleScriptID(oracleScriptID)
	callData := obi.MustEncode(types.AssetPriceOBIInput{
		Multiplier: types.Multiplier,
	})
	askCount := uint64(4)
	minCount := uint64(3)

	packet := oracletypes.NewOracleRequestPacketData(
		clientID, oracleScript, callData, askCount, minCount,
	)

	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	if err := k.channelKeeper.SendPacket(ctx, channelCap, channeltypes.NewPacket(
		packet.GetBytes(),
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		clienttypes.ZeroHeight(),
		uint64(ctx.BlockTime().UnixNano()+int64(20*time.Minute)),
	)); err != nil {
		return err
	}
	return nil
}
