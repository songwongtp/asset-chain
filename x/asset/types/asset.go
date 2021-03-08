package types


//-----------------------------------------------------------------------------
// Asset

// NewAsset returns a new asset info containing the denom name, totalsupply, and price
func NewAsset(denom string, totalSupply uint64, price uint64) Asset {
	asset := Asset {
		Denom: denom,
		TotalSupply: totalSupply,
		Price: price,
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

	for _, asset := range assets {
		result = append(result, asset)
	}

	return result
}