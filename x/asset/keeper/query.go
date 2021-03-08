package keeper

import (
	// this line is used by starport scaffolding # 1

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/songwongtp/asset-chain/x/asset/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier returns a new sdk.Keeper instance
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)

		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryAssetInfo:
			res, err = queryAssetInfo(ctx, req, k, legacyQuerierCdc)
		case types.QueryAllAssetInfos:
			res, err = queryAllAssetInfos(ctx, req, k, legacyQuerierCdc)
		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
}

func queryAssetInfo(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAssetInfoRequest

	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	asset := k.GetAssetInfo(ctx, params.Denom)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, asset)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryAllAssetInfos(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	denoms := k.GetAllAssetDenoms(ctx)

	assets := make([]types.Asset, 0)
	for _, denom := range denoms {
		assets = append(assets, k.GetAssetInfo(ctx, denom))
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, assets)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}