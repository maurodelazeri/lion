package orderbook

import (
	"strconv"

	"github.com/shopspring/decimal"
)

// Order ...
type Order struct {
	timestamp int
	quantity  decimal.Decimal
	price     decimal.Decimal
	orderID   string
	tradeID   string
	nextOrder *Order
	prevOrder *Order
	orderList *OrderList
}

// NewOrder ...
func NewOrder(quote map[string]string, orderList *OrderList) *Order {
	timestamp, _ := strconv.Atoi(quote["timestamp"])
	quantity, _ := decimal.NewFromString(quote["quantity"])
	price, _ := decimal.NewFromString(quote["price"])
	orderID := quote["order_id"]
	tradeID := quote["trade_id"]
	return &Order{timestamp: timestamp, quantity: quantity, price: price, orderID: orderID,
		tradeID: tradeID, nextOrder: nil, prevOrder: nil, orderList: orderList}
}

// NextOrder ...
func (o *Order) NextOrder() *Order {
	return o.nextOrder
}

// PrevOrder ...
func (o *Order) PrevOrder() *Order {
	return o.prevOrder
}

// UpdateQuantity ...
func (o *Order) UpdateQuantity(newQuantity decimal.Decimal, newTimestamp int) {
	if newQuantity.GreaterThan(o.quantity) && o.orderList.tailOrder != o {
		o.orderList.MoveToTail(o)
	}
	o.orderList.volume = o.orderList.volume.Sub(o.quantity.Sub(newQuantity))
	o.timestamp = newTimestamp
	o.quantity = newQuantity
}
