package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/songwongtp/asset-chain/x/asset/types"
)

// Req defines the property of a buy/sell/add_supply request's body
type Req struct {
	BaseReq	rest.BaseReq	`json:"base_req" yaml:"base_req"`
	Amount	uint64			`json:"amount" yaml:"amount"`
}

// PriceReq defines the property of a price_set request's body
type PriceReq struct {
	BaseReq	rest.BaseReq	`json:"base_req" yaml:"base_req"`
	Price	uint64			`json:"price" yaml:"price"`
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

// NewSetPriceRequestHandlerFn returns an HTTP REST handler for creating
// a MsgSetPrice transaction
func NewSetPriceRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		denom := vars["denom"]

		var req PriceReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgSetPrice(req.BaseReq.From, denom, req.Price)
		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

// NewAddSupplyRequestHandlerFn returns an HTTP REST handler for creating
// a MsgAddSupply transaction
func NewAddSupplyRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
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

		msg := types.NewMsgAddSupply(req.BaseReq.From, denom, req.Amount)
		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}