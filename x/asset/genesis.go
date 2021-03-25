package asset

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/songwongtp/asset-chain/x/asset/keeper"
	"github.com/songwongtp/asset-chain/x/asset/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	for _, assetInfo := range genState.Assets {
		// change to IBC
		k.SetOracleScriptID(ctx, assetInfo.Denom, assetInfo.OracleScriptId)

		if err := k.AddSupply(ctx, assetInfo.Denom, assetInfo.TotalSupply); err != nil {
			panic(err)
		}

		k.SetPort(ctx, types.PortID)
		k.BindPort(ctx, types.PortID)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	// genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	assets := k.GetAllAssetInfos(ctx)

	return &types.GenesisState{
		Assets: assets,
	}
}
