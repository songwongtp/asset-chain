package keeper

import (
	"github.com/songwongtp/asset-chain/x/asset/types"
)

var _ types.QueryServer = Keeper{}
