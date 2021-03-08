package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

// RegisterRoutes registers asset-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/asset/infos", QueryAssetInfosRequestHandlerFn(clientCtx)).Methods("GET")
}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/asset/buy/{denom}", NewBuyAssetRequestHandlerFn(clientCtx)).Methods("POST")
	r.HandleFunc("/asset/sell/{denom}", NewSellAssetRequestHandlerFn(clientCtx)).Methods("POST")
}
