package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/songwongtp/asset-chain/x/asset/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the asset MsgServer interface
// for the provided keeper
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) BuyAsset(goCtx context.Context, msg *types.MsgBuyAsset) (*types.MsgBuyAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return nil, err
	}

	err = k.Buy(ctx, addr, msg.Denom, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgBuyAssetResponse{}, nil
}

func (k msgServer) SellAsset(goCtx context.Context, msg *types.MsgSellAsset) (*types.MsgSellAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		return nil, err
	}

	err = k.Sell(ctx, addr, msg.Denom, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgSellAssetResponse{}, nil
}
