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
	sdk.RegisterDenom(genState.BaseDenom, sdk.NewDec(1))

	for _, assetInfo := range genState.Assets {
		k.SetPrice(ctx, assetInfo.Denom, assetInfo.Price)

		if err := k.AddSupply(ctx, assetInfo.Denom, assetInfo.TotalSupply); err != nil {
			panic(err)
		}		
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
