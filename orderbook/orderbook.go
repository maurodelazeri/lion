package orderbook

import (
	"strconv"

	"github.com/shopspring/decimal"
	lane "gopkg.in/oleiade/lane.v1"
)

// OrderBook ...
type OrderBook struct {
	deque       *lane.Deque
	buys        *OrderTree
	sells       *OrderTree
	time        int
	nextOrderID int
}

// NewOrderBook ...
func NewOrderBook() *OrderBook {
	deque := lane.NewDeque()
	buys := NewOrderTree()
	sells := NewOrderTree()
	return &OrderBook{deque, buys, sells, 0, 0}
}

// UpdateTime ...
func (orderBook *OrderBook) UpdateTime() {
	orderBook.time++
}

// BestBuy ...
func (orderBook *OrderBook) BestBuy() (value decimal.Decimal) {
	value = orderBook.buys.MaxPrice()
	return
}

// BestSell ...
func (orderBook *OrderBook) BestSell() (value decimal.Decimal) {
	value = orderBook.sells.MinPrice()
	return
}

// WorstBuy ...
func (orderBook *OrderBook) WorstBuy() (value decimal.Decimal) {
	value = orderBook.buys.MinPrice()
	return
}

// WorstSell ...
func (orderBook *OrderBook) WorstSell() (value decimal.Decimal) {
	value = orderBook.sells.MaxPrice()
	return
}

// ProcessMarketOrder ...
func (orderBook *OrderBook) ProcessMarketOrder(quote map[string]string, verbose bool) []map[string]string {
	var trades []map[string]string
	quantityToTrade, _ := decimal.NewFromString(quote["quantity"])
	side := quote["side"]
	var newTrades []map[string]string

	if side == "BUY" {
		for quantityToTrade.GreaterThan(decimal.Zero) && orderBook.sells.Length() > 0 {
			bestPriceBuys := orderBook.sells.MinPriceList()
			quantityToTrade, newTrades = orderBook.ProcessOrderList("SELL", bestPriceBuys, quantityToTrade, quote, verbose)
			trades = append(trades, newTrades...)
		}
	} else if side == "SELL" {
		for quantityToTrade.GreaterThan(decimal.Zero) && orderBook.buys.Length() > 0 {
			bestPriceBuys := orderBook.buys.MaxPriceList()
			quantityToTrade, newTrades = orderBook.ProcessOrderList("BUY", bestPriceBuys, quantityToTrade, quote, verbose)
			trades = append(trades, newTrades...)
		}
	}
	return trades
}

// ProcessLimitOrder ...
func (orderBook *OrderBook) ProcessLimitOrder(quote map[string]string, verbose bool) ([]map[string]string, map[string]string) {
	var trades []map[string]string
	quantityToTrade, _ := decimal.NewFromString(quote["quantity"])
	side := quote["side"]
	//price, _ := decimal.NewFromString(quote["price"])
	//var new_trades []map[string]string
	var orderInBook map[string]string

	if side == "BUY" {
		// minPrice := orderBook.sells.MinPrice()
		// for quantity_to_trade.GreaterThan(decimal.Zero) && orderBook.sells.Length() > 0 && price.GreaterThanOrEqual(minPrice) {
		// 	best_price_sells := orderBook.sells.MinPriceList()
		// 	quantity_to_trade, new_trades = orderBook.ProcessOrderList("SELL", best_price_sells, quantity_to_trade, quote, verbose)
		// 	trades = append(trades, new_trades...)
		// 	minPrice = orderBook.sells.MinPrice()
		// }

		if quantityToTrade.GreaterThan(decimal.Zero) {
			quote["order_id"] = strconv.Itoa(orderBook.nextOrderID)
			quote["quantity"] = quantityToTrade.String()
			orderBook.buys.InsertOrder(quote)
			orderInBook = quote
		}

	} else if side == "SELL" {
		// maxPrice := orderBook.buys.MaxPrice()
		// for quantity_to_trade.GreaterThan(decimal.Zero) && orderBook.buys.Length() > 0 && price.LessThanOrEqual(maxPrice) {
		// 	best_price_buys := orderBook.buys.MaxPriceList()
		// 	quantity_to_trade, new_trades = orderBook.ProcessOrderList("BUY", best_price_buys, quantity_to_trade, quote, verbose)
		// 	trades = append(trades, new_trades...)
		// 	maxPrice = orderBook.buys.MaxPrice()
		// }

		if quantityToTrade.GreaterThan(decimal.Zero) {
			quote["order_id"] = strconv.Itoa(orderBook.nextOrderID)
			quote["quantity"] = quantityToTrade.String()
			orderBook.sells.InsertOrder(quote)
			orderInBook = quote
		}
	}
	return trades, orderInBook
}

// ProcessOrder ...
func (orderBook *OrderBook) ProcessOrder(quote map[string]string, verbose bool) ([]map[string]string, map[string]string) {

	orderType := quote["type"]
	var orderInBook map[string]string
	var trades []map[string]string

	orderBook.UpdateTime()
	quote["timestamp"] = strconv.Itoa(orderBook.time)
	orderBook.nextOrderID++

	if orderType == "MARKET" {
		trades = orderBook.ProcessMarketOrder(quote, verbose)
	} else {
		trades, orderInBook = orderBook.ProcessLimitOrder(quote, verbose)
	}
	return trades, orderInBook
}

// ProcessOrderList ...
func (orderBook *OrderBook) ProcessOrderList(side string, orderList *OrderList, quantityStillToTrade decimal.Decimal, quote map[string]string, verbose bool) (decimal.Decimal, []map[string]string) {
	quantityToTrade := quantityStillToTrade
	var trades []map[string]string

	for orderList.Length() > 0 && quantityToTrade.GreaterThan(decimal.Zero) {
		headOrder := orderList.HeadOrder()
		tradedPrice := headOrder.price
		// counterParty := headOrder.trade_id
		var newBookQuantity decimal.Decimal
		var tradedQuantity decimal.Decimal

		if quantityToTrade.LessThan(headOrder.quantity) {
			tradedQuantity = quantityToTrade
			// Do the transaction
			newBookQuantity = headOrder.quantity.Sub(quantityToTrade)
			headOrder.UpdateQuantity(newBookQuantity, headOrder.timestamp)
			quantityToTrade = decimal.Zero

		} else if quantityToTrade.Equal(headOrder.quantity) {
			tradedQuantity = quantityToTrade
			if side == "BUY" {
				orderBook.buys.RemoveOrderByID(headOrder.orderID)
			} else {
				orderBook.sells.RemoveOrderByID(headOrder.orderID)
			}
			quantityToTrade = decimal.Zero

		} else {
			tradedQuantity = headOrder.quantity
			if side == "BUY" {
				orderBook.buys.RemoveOrderByID(headOrder.orderID)
			} else {
				orderBook.sells.RemoveOrderByID(headOrder.orderID)
			}
		}

		if verbose {
			//	fmt.Println("TRADE: Time - %v, Price - %v, Quantity - %v, TradeID - %v, Matching TradeID - %v", orderBook.time, tradedPrice.String(), tradedQuantity.String(), counterParty, quote["trade_id"])
		}

		transactionRecord := make(map[string]string)
		transactionRecord["timestamp"] = strconv.Itoa(orderBook.time)
		transactionRecord["price"] = tradedPrice.String()
		transactionRecord["quantity"] = tradedQuantity.String()
		transactionRecord["time"] = strconv.Itoa(orderBook.time)

		orderBook.deque.Append(transactionRecord)
		trades = append(trades, transactionRecord)
	}
	return quantityToTrade, trades
}

// CancelOrder ...
func (orderBook *OrderBook) CancelOrder(side string, orderID int) {
	orderBook.UpdateTime()

	if side == "BUY" {
		if orderBook.buys.OrderExist(strconv.Itoa(orderID)) {
			orderBook.buys.RemoveOrderByID(strconv.Itoa(orderID))
		}
	} else {
		if orderBook.sells.OrderExist(strconv.Itoa(orderID)) {
			orderBook.sells.RemoveOrderByID(strconv.Itoa(orderID))
		}
	}
}

// ModifyOrder ...
func (orderBook *OrderBook) ModifyOrder(quoteUpdate map[string]string, orderID int) {
	orderBook.UpdateTime()

	side := quoteUpdate["side"]
	quoteUpdate["order_id"] = strconv.Itoa(orderID)
	quoteUpdate["timestamp"] = strconv.Itoa(orderBook.time)

	if side == "BUY" {
		if orderBook.buys.OrderExist(strconv.Itoa(orderID)) {
			orderBook.buys.UpdateOrder(quoteUpdate)
		}
	} else {
		if orderBook.sells.OrderExist(strconv.Itoa(orderID)) {
			orderBook.sells.UpdateOrder(quoteUpdate)
		}
	}
}

// VolumeAtPrice ...
func (orderBook *OrderBook) VolumeAtPrice(side string, price decimal.Decimal) decimal.Decimal {
	if side == "BUY" {
		volume := decimal.Zero
		if orderBook.buys.PriceExist(price) {
			volume = orderBook.buys.PriceList(price).volume
		}

		return volume

	}
	volume := decimal.Zero
	if orderBook.sells.PriceExist(price) {
		volume = orderBook.sells.PriceList(price).volume
	}
	return volume
}
