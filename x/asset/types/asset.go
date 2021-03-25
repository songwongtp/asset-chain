package types

//-----------------------------------------------------------------------------
// Asset

// NewAsset returns a new asset info containing the denom name, totalsupply, and price
func NewAsset(denom string, totalSupply uint64, oracleScriptID uint64) Asset {
	asset := Asset{
		Denom:          denom,
		TotalSupply:    totalSupply,
		OracleScriptId: oracleScriptID,
	}
	return asset
}

//-----------------------------------------------------------------------------
// CANNOT USE: import cycle not allowed
// Assets

// Assets is a set of Assets, one per asset type(denom)
type Assets []Asset

// NewAssets construct a new asset set
func NewAssets(assets ...Asset) Assets {
	result := make([]Asset, 0, len(assets))
	result = append(result, assets...)
	return result
}
