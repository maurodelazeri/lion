package orderbook

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

var limitOrders = make([]map[string]string, 0)

func TestNewOrderBook(t *testing.T) {
	orderBook := NewOrderBook()

	if !(orderBook.VolumeAtPrice("BUY", decimal.Zero).Equal(decimal.Zero)) {
		t.Errorf("orderBook.VolumeAtPrice incorrect, got: %d, want: %d.", orderBook.VolumeAtPrice("BUY", decimal.Zero), decimal.Zero)
	}

	if !(orderBook.BestSell().Equal(decimal.Zero)) {
		t.Errorf("orderBook.BestSell incorrect, got: %d, want: %d.", orderBook.BestSell(), decimal.Zero)
	}

	if !(orderBook.WorstBuy().Equal(decimal.Zero)) {
		t.Errorf("orderBook.WorstBuy incorrect, got: %d, want: %d.", orderBook.WorstBuy(), decimal.Zero)
	}

	if !(orderBook.WorstSell().Equal(decimal.Zero)) {
		t.Errorf("orderBook.WorstSell incorrect, got: %d, want: %d.", orderBook.WorstSell(), decimal.Zero)
	}

	if !(orderBook.BestBuy().Equal(decimal.Zero)) {
		t.Errorf("orderBook.BestBuy incorrect, got: %d, want: %d.", orderBook.BestBuy(), decimal.Zero)
	}
}

func TestOrderBook(t *testing.T) {
	orderBook := NewOrderBook()

	fmt.Println(orderBook.BestSell())

	dummyOrder := make(map[string]string)
	dummyOrder["type"] = "LIMIT"
	dummyOrder["side"] = "SELL"
	dummyOrder["quantity"] = "5"
	dummyOrder["price"] = "101"
	dummyOrder["trade_id"] = "100"

	limitOrders = append(limitOrders, dummyOrder)

	dummyOrder1 := make(map[string]string)
	dummyOrder1["type"] = "LIMIT"
	dummyOrder1["side"] = "SELL"
	dummyOrder1["quantity"] = "5"
	dummyOrder1["price"] = "103"
	dummyOrder1["trade_id"] = "101"

	limitOrders = append(limitOrders, dummyOrder1)

	dummyOrder2 := make(map[string]string)
	dummyOrder2["type"] = "LIMIT"
	dummyOrder2["side"] = "SELL"
	dummyOrder2["quantity"] = "5"
	dummyOrder2["price"] = "101"
	dummyOrder2["trade_id"] = "102"

	limitOrders = append(limitOrders, dummyOrder2)

	dummyOrder7 := make(map[string]string)
	dummyOrder7["type"] = "LIMIT"
	dummyOrder7["side"] = "SELL"
	dummyOrder7["quantity"] = "5"
	dummyOrder7["price"] = "101"
	dummyOrder7["trade_id"] = "103"

	limitOrders = append(limitOrders, dummyOrder7)

	dummyOrder3 := make(map[string]string)
	dummyOrder3["type"] = "LIMIT"
	dummyOrder3["side"] = "BUY"
	dummyOrder3["quantity"] = "5"
	dummyOrder3["price"] = "99"
	dummyOrder3["trade_id"] = "100"

	limitOrders = append(limitOrders, dummyOrder3)

	dummyOrder4 := make(map[string]string)
	dummyOrder4["type"] = "LIMIT"
	dummyOrder4["side"] = "BUY"
	dummyOrder4["quantity"] = "5"
	dummyOrder4["price"] = "98"
	dummyOrder4["trade_id"] = "101"

	limitOrders = append(limitOrders, dummyOrder4)

	dummyOrder5 := make(map[string]string)
	dummyOrder5["type"] = "LIMIT"
	dummyOrder5["side"] = "BUY"
	dummyOrder5["quantity"] = "5"
	dummyOrder5["price"] = "99"
	dummyOrder5["trade_id"] = "102"

	limitOrders = append(limitOrders, dummyOrder5)

	dummyOrder6 := make(map[string]string)
	dummyOrder6["type"] = "LIMIT"
	dummyOrder6["side"] = "BUY"
	dummyOrder6["quantity"] = "5"
	dummyOrder6["price"] = "97"
	dummyOrder6["trade_id"] = "103"

	limitOrders = append(limitOrders, dummyOrder6)

	for _, order := range limitOrders {
		orderBook.ProcessOrder(order, true)
	}

	value, _ := decimal.NewFromString("101")
	if !(orderBook.BestSell().Equal(value)) {
		t.Errorf("orderBook.BestSell incorrect, got: %d, want: %d.", orderBook.BestSell(), value)
	}

	value, _ = decimal.NewFromString("103")
	if !(orderBook.WorstSell().Equal(value)) {
		t.Errorf("orderBook.WorstBuy incorrect, got: %d, want: %d.", orderBook.WorstSell(), value)
	}

	value, _ = decimal.NewFromString("99")
	if !(orderBook.BestBuy().Equal(value)) {
		t.Errorf("orderBook.BestBuy incorrect, got: %d, want: %d.", orderBook.BestBuy(), value)
	}

	value, _ = decimal.NewFromString("97")
	if !(orderBook.WorstBuy().Equal(value)) {
		t.Errorf("orderBook.BestBuy incorrect, got: %d, want: %d.", orderBook.WorstBuy(), value)
	}

	value, _ = decimal.NewFromString("15")
	pricePoint, _ := decimal.NewFromString("101")
	if !(orderBook.VolumeAtPrice("SELL", pricePoint).Equal(value)) {
		t.Errorf("orderBook.VolumeAtPrice incorrect, got: %d, want: %d.", orderBook.VolumeAtPrice("BUY", decimal.Zero), decimal.Zero)
	}

	//Submitting a limit order that crosses the opposing best price will result in a trade
	marketOrder := make(map[string]string)
	marketOrder["type"] = "LIMIT"
	marketOrder["side"] = "BUY"
	marketOrder["quantity"] = "2"
	marketOrder["price"] = "102"
	marketOrder["trade_id"] = "109"

	trades, order_in_book := orderBook.ProcessOrder(marketOrder, true)

	tradedPrice := trades[0]["price"]
	tradedQuantity := trades[0]["quantity"]

	if !(tradedPrice == "101" && tradedQuantity == "2" && len(order_in_book) == 0) {
		t.Errorf("orderBook.ProcessOrder incorrect")
	}

	// If a limit crosses but is only partially matched, the remaning volume will
	// be placed in the book as an outstanding order
	bigOrder := make(map[string]string)
	bigOrder["type"] = "LIMIT"
	bigOrder["side"] = "BUY"
	bigOrder["quantity"] = "50"
	bigOrder["price"] = "102"
	bigOrder["trade_id"] = "110"

	trades, order_in_book = orderBook.ProcessOrder(bigOrder, true)

	fmt.Println(trades)
	fmt.Println(order_in_book)

	if !(len(order_in_book) != 0) {
		t.Errorf("orderBook.ProcessOrder incorrect")
	}

	// Market orders only require that a user specifies a side (buy or sell), a quantity, and their unique trade id
	marketOrder = make(map[string]string)
	marketOrder["type"] = "MARKET"
	marketOrder["side"] = "SELL"
	marketOrder["quantity"] = "20"
	marketOrder["trade_id"] = "111"

	trades, order_in_book = orderBook.ProcessOrder(marketOrder, true)

}
