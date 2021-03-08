package types

// Querier path constants
const (
	QueryAssetInfo		= "asset_info"
	QueryAllAssetInfos	= "all_asset_info"
)

// NewQueryAssetInfoRequest creates a new instance of QueryAssetInfoRequest
func NewQueryAssetInfoRequest(denom string) *QueryAssetInfoRequest {
	return &QueryAssetInfoRequest{Denom: denom}
}

// NewQueryAllAssetInfosRequest creates a new instance of QueryAllAssetInfosRequest
func NewQueryAllAssetInfosRequest() *QueryAllAssetInfosRequest {
	return &QueryAllAssetInfosRequest{}
}