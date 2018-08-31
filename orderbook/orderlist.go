package orderbook

import (
	"github.com/HuKeping/rbtree"
	"github.com/shopspring/decimal"
)

// Item ...
type Item interface {
	Less(than Item) bool
}

// OrderList ...
type OrderList struct {
	headOrder *Order
	tailOrder *Order
	length    int
	volume    decimal.Decimal
	lastOrder *Order
	price     decimal.Decimal
}

// NewOrderList ...
func NewOrderList(price decimal.Decimal) *OrderList {
	return &OrderList{headOrder: nil, tailOrder: nil, length: 0, volume: decimal.Zero,
		lastOrder: nil, price: price}
}

// Less ...
func (orderlist *OrderList) Less(than rbtree.Item) bool {
	return orderlist.price.LessThan(than.(*OrderList).price)
}

// Length ...
func (orderlist *OrderList) Length() int {
	return orderlist.length
}

// HeadOrder ...
func (orderlist *OrderList) HeadOrder() *Order {
	return orderlist.headOrder
}

// AppendOrder ...
func (orderlist *OrderList) AppendOrder(order *Order) {
	if orderlist.Length() == 0 {
		order.nextOrder = nil
		order.prevOrder = nil
		orderlist.headOrder = order
		orderlist.tailOrder = order
	} else {
		order.prevOrder = orderlist.tailOrder
		order.nextOrder = nil
		orderlist.tailOrder.nextOrder = order
		orderlist.tailOrder = order
	}
	orderlist.length = orderlist.length + 1
	orderlist.volume = orderlist.volume.Add(order.quantity)
}

// RemoveOrder ...
func (orderlist *OrderList) RemoveOrder(order *Order) {
	orderlist.volume = orderlist.volume.Sub(order.quantity)
	orderlist.length = orderlist.length - 1
	if orderlist.length == 0 {
		return
	}

	nextOrder := order.nextOrder
	prevOrder := order.prevOrder

	if nextOrder != nil && prevOrder != nil {
		nextOrder.prevOrder = prevOrder
		prevOrder.nextOrder = nextOrder
	} else if nextOrder != nil {
		nextOrder.prevOrder = nil
		orderlist.headOrder = nextOrder
	} else if prevOrder != nil {
		prevOrder.nextOrder = nil
		orderlist.tailOrder = prevOrder
	}
}

// MoveToTail ...
func (orderlist *OrderList) MoveToTail(order *Order) {
	if order.prevOrder != nil { // This Order is not the first Order in the OrderList
		order.prevOrder.nextOrder = order.nextOrder // Link the previous Order to the next Order, then move the Order to tail
	} else { // This Order is the first Order in the OrderList
		orderlist.headOrder = order.nextOrder // Make next order the first
	}
	order.nextOrder.prevOrder = order.prevOrder

	// Move Order to the last position. Link up the previous last position Order.
	orderlist.tailOrder.nextOrder = order
	orderlist.tailOrder = order
}
