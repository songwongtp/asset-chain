package types

const (
	TypeOrderBuy  = "order_buy"
	TypeOrderSell = "order_sell"

	StatusOrderPending = "status_pending"
	StatusOrderSuccess = "status_success"
	StatusOrderFail    = "status_fail"
)

const Multiplier = 1000000

// NewOrderAsset creates a new OrderAsset instance
func NewOrderAsset(orderID string, orderType string, addr string, denom string, amount uint64) AssetOrder {
	assetOrder := AssetOrder{
		OrderId:   orderID,
		OrderType: orderType,
		Addr:      addr,
		Denom:     denom,
		Amount:    amount,
		Status:    StatusOrderPending,
	}
	return assetOrder
}
