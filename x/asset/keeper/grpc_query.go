package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/songwongtp/asset-chain/x/asset/types"
)

var _ types.QueryServer = Keeper{}

// AssetInfo implements the Query/AssetInfo gRPC method
func (k Keeper) AssetInfo(ctx context.Context, req *types.QueryAssetInfoRequest) (*types.QueryAssetInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.Denom == "" {
		return nil, status.Error(codes.InvalidArgument, "denom cannot be empty")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	asset := k.GetAssetInfo(sdkCtx, req.Denom)
	return &types.QueryAssetInfoResponse{Asset: &asset}, nil
}

// AllAssetInfos implements the Query/AllAssetInfos gRPC method
func (k Keeper) AllAssetInfos(ctx context.Context, req *types.QueryAllAssetInfosRequest) (*types.QueryAllAssetInfosResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	denoms := k.GetAllAssetDenoms(sdkCtx)

	assets := make([]types.Asset, 0)
	for _, denom := range denoms {
		assets = append(assets, k.GetAssetInfo(sdkCtx, denom))
	}
	return &types.QueryAllAssetInfosResponse{Assets: assets}, nil
}