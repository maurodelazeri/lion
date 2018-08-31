package orderbook

import (
	"strconv"

	"github.com/shopspring/decimal"
	lane "gopkg.in/oleiade/lane.v1"
)

type OrderBook struct {
	deque       *lane.Deque
	buys        *OrderTree
	sells        *OrderTree
	time        int
	nextOrderID int
}

func NewOrderBook() *OrderBook {
	deque := lane.NewDeque()
	buys := NewOrderTree()
	sells := NewOrderTree()
	return &OrderBook{deque, buys, sells, 0, 0}
}

func (orderBook *OrderBook) UpdateTime() {
	orderBook.time++
}

func (orderBook *OrderBook) BestBuy() (value decimal.Decimal) {
	value = orderBook.buys.MaxPrice()
	return
}

func (orderBook *OrderBook) BestSell() (value decimal.Decimal) {
	value = orderBook.sells.MinPrice()
	return
}

func (orderBook *OrderBook) WorstBuy() (value decimal.Decimal) {
	value = orderBook.buys.MinPrice()
	return
}

func (orderBook *OrderBook) WorstSell() (value decimal.Decimal) {
	value = orderBook.sells.MaxPrice()
	return
}

func (orderBook *OrderBook) ProcessMarketOrder(quote map[string]string, verbose bool) []map[string]string {
	var trades []map[string]string
	quantity_to_trade, _ := decimal.NewFromString(quote["quantity"])
	side := quote["side"]
	var new_trades []map[string]string

	if side == "BUY" {
		for quantity_to_trade.GreaterThan(decimal.Zero) && orderBook.sells.Length() > 0 {
			best_price_sells := orderBook.sells.MinPriceList()
			quantity_to_trade, new_trades = orderBook.ProcessOrderList("SELL", best_price_sells, quantity_to_trade, quote, verbose)
			trades = append(trades, new_trades...)
		}
	} else if side == "SELL" {
		for quantity_to_trade.GreaterThan(decimal.Zero) && orderBook.buys.Length() > 0 {
			best_price_buys := orderBook.buys.MaxPriceList()
			quantity_to_trade, new_trades = orderBook.ProcessOrderList("BUY", best_price_buys, quantity_to_trade, quote, verbose)
			trades = append(trades, new_trades...)
		}
	}
	return trades
}

func (orderBook *OrderBook) ProcessLimitOrder(quote map[string]string, verbose bool) ([]map[string]string, map[string]string) {
	var trades []map[string]string
	quantity_to_trade, _ := decimal.NewFromString(quote["quantity"])
	side := quote["side"]
	price, _ := decimal.NewFromString(quote["price"])
	var new_trades []map[string]string

	var order_in_book map[string]string

	if side == "BUY" {
		minPrice := orderBook.sells.MinPrice()
		for quantity_to_trade.GreaterThan(decimal.Zero) && orderBook.sells.Length() > 0 && price.GreaterThanOrEqual(minPrice) {
			best_price_sells := orderBook.sells.MinPriceList()
			quantity_to_trade, new_trades = orderBook.ProcessOrderList("SELL", best_price_sells, quantity_to_trade, quote, verbose)
			trades = append(trades, new_trades...)
			minPrice = orderBook.sells.MinPrice()
		}

		if quantity_to_trade.GreaterThan(decimal.Zero) {
			quote["order_id"] = strconv.Itoa(orderBook.nextOrderID)
			quote["quantity"] = quantity_to_trade.String()
			orderBook.buys.InsertOrder(quote)
			order_in_book = quote
		}

	} else if side == "SELL" {
		maxPrice := orderBook.buys.MaxPrice()
		for quantity_to_trade.GreaterThan(decimal.Zero) && orderBook.buys.Length() > 0 && price.LessThanOrEqual(maxPrice) {
			best_price_buys := orderBook.buys.MaxPriceList()
			quantity_to_trade, new_trades = orderBook.ProcessOrderList("BUY", best_price_buys, quantity_to_trade, quote, verbose)
			trades = append(trades, new_trades...)
			maxPrice = orderBook.buys.MaxPrice()
		}

		if quantity_to_trade.GreaterThan(decimal.Zero) {
			quote["order_id"] = strconv.Itoa(orderBook.nextOrderID)
			quote["quantity"] = quantity_to_trade.String()
			orderBook.sells.InsertOrder(quote)
			order_in_book = quote
		}
	}
	return trades, order_in_book
}

func (orderBook *OrderBook) ProcessOrder(quote map[string]string, verbose bool) ([]map[string]string, map[string]string) {

	order_type := quote["type"]
	var order_in_book map[string]string
	var trades []map[string]string

	orderBook.UpdateTime()
	quote["timestamp"] = strconv.Itoa(orderBook.time)
	orderBook.nextOrderID++

	if order_type == "MARKET" {
		trades = orderBook.ProcessMarketOrder(quote, verbose)
	} else {
		trades, order_in_book = orderBook.ProcessLimitOrder(quote, verbose)
	}
	return trades, order_in_book
}

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
				orderBook.buys.RemoveOrderById(headOrder.orderID)
			} else {
				orderBook.sells.RemoveOrderById(headOrder.orderID)
			}
			quantityToTrade = decimal.Zero

		} else {
			tradedQuantity = headOrder.quantity
			if side == "BUY" {
				orderBook.buys.RemoveOrderById(headOrder.orderID)
			} else {
				orderBook.sells.RemoveOrderById(headOrder.orderID)
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

func (orderBook *OrderBook) CancelOrder(side string, order_id int) {
	orderBook.UpdateTime()

	if side == "BUY" {
		if orderBook.buys.OrderExist(strconv.Itoa(order_id)) {
			orderBook.buys.RemoveOrderById(strconv.Itoa(order_id))
		}
	} else {
		if orderBook.sells.OrderExist(strconv.Itoa(order_id)) {
			orderBook.sells.RemoveOrderById(strconv.Itoa(order_id))
		}
	}
}

func (orderBook *OrderBook) ModifyOrder(quoteUpdate map[string]string, order_id int) {
	orderBook.UpdateTime()

	side := quoteUpdate["side"]
	quoteUpdate["order_id"] = strconv.Itoa(order_id)
	quoteUpdate["timestamp"] = strconv.Itoa(orderBook.time)

	if side == "BUY" {
		if orderBook.buys.OrderExist(strconv.Itoa(order_id)) {
			orderBook.buys.UpdateOrder(quoteUpdate)
		}
	} else {
		if orderBook.sells.OrderExist(strconv.Itoa(order_id)) {
			orderBook.sells.UpdateOrder(quoteUpdate)
		}
	}
}

func (orderBook *OrderBook) VolumeAtPrice(side string, price decimal.Decimal) decimal.Decimal {
	if side == "BUY" {
		volume := decimal.Zero
		if orderBook.buys.PriceExist(price) {
			volume = orderBook.buys.PriceList(price).volume
		}

		return volume

	} else {
		volume := decimal.Zero
		if orderBook.sells.PriceExist(price) {
			volume = orderBook.sells.PriceList(price).volume
		}
		return volume
	}
}
