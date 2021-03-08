package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/songwongtp/asset-chain/x/asset/types"
)

// QueryAssetInfosRequestHandlerFn returns a REST handler that queries for
// all asset infos or a specific asset info
func QueryAssetInfosRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		var (
			params interface{}
			route string
		)

		denom := r.FormValue("denom")
		if denom == "" {
			params = types.NewQueryAllAssetInfosRequest()
			route = fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAllAssetInfos)
		} else {
			params = types.NewQueryAssetInfoRequest(denom)
			route = fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAssetInfo)
		}

		bz, err := ctx.LegacyAmino.MarshalJSON(params)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		res, height, err := ctx.QueryWithData(route, bz)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		ctx = ctx.WithHeight(height)
		rest.PostProcessResponse(w, ctx, res)
	}
}