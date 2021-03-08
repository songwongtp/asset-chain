package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/songwongtp/asset-chain/x/asset/types"
)

// Req defines the property of a buy/sell request's body
type Req struct {
	BaseReq	rest.BaseReq	`json:"base_req" yaml:"base_req"`
	Amount	uint64			`json:"amount" yaml:"amount"`
}

// NewBuyAssetRequestHandlerFn returns an HTTP REST handler for creating
// a MsgBuyAsset transaction
func NewBuyAssetRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		denom := vars["denom"]

		var req Req
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgBuyAsset(req.BaseReq.From, denom, req.Amount)
		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

// NewSellAssetRequestHandlerFn returns an HTTP REST handler for creating
// a MsgSellAsset transaction
func NewSellAssetRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		denom := vars["denom"]

		var req Req
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgSellAsset(req.BaseReq.From, denom, req.Amount)
		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}