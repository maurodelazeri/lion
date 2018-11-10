package bitmex

import (
	"errors"
	"log"
	"time"

	"github.com/maurodelazeri/zion/common"
	"github.com/maurodelazeri/zion/currency/pair"
	exchange "github.com/maurodelazeri/zion/exchanges"
	"github.com/maurodelazeri/zion/exchanges/orderbook"
)

// GetExchangePullingType return the type of pulling data the exchange is using
func (b *Bitmex) GetExchangePullingType() string {
	if b.Websocket {
		return "socket"
	}
	return "rest"
}

// Start starts the bitmex go routine
func (b *Bitmex) Start() {
	go b.Run()
}

// Run implements the bitmex wrapper
func (b *Bitmex) Run() {
	if b.Verbose {
		log.Printf("%s Websocket: %s. (url: %s).\n", b.GetName(), common.IsEnabled(b.Websocket), bitmexWebsocketURL)
		log.Printf("%s polling delay: %ds.\n", b.GetName(), b.RESTPollingDelay)
		log.Printf("%s %d currencies enabled: %s.\n", b.GetName(), len(b.APIEnabledPairs), b.APIEnabledPairs)
	}

	if b.Websocket {
		go b.WebsocketClient()
	}

	b.UpdateExchangePairsInternals()
}

// UpdateExchangePairsInternals set the coin config
func (b *Bitmex) UpdateExchangePairsInternals() {
	go func() {
		for {

			// pairData, err := b.GetCurrencies()
			// if err != nil {
			// 	logrus.Error("Problem to get the exchange pair data ", b.Name, " ", err)
			// 	time.Sleep(5 * time.Minute)
			// 	continue
			// }

			// for _, data := range pairData {
			// 	var enabled bool

			// 	if data.Status == "online" {
			// 		enabled = true
			// 	} else {
			// 		enabled = false
			// 	}

			// 	var pairconfig exchange.PairExchangeInternalConfig
			// 	pairconfig.Product = strings.ToUpper(data.ID)
			// 	pairconfig.MakerFee = b.MakerFee
			// 	pairconfig.TakerFee = b.TakerFee
			// 	pairconfig.Enabled = enabled
			// 	pairconfig.LastUpdate = time.Now().UTC()

			// 	var depositConfig exchange.DepositConfig
			// 	depositConfig.Enabled = enabled
			// 	depositConfig.Confirmations = -1
			// 	depositConfig.ExchangeFee = 0.0
			// 	pairconfig.Deposit = depositConfig

			// 	var withdrawalConfig exchange.WithdrawalConfig
			// 	withdrawalConfig.Enabled = enabled
			// 	withdrawalConfig.NetworkFee = -1
			// 	withdrawalConfig.ExchangeFee = 0.0
			// 	pairconfig.Withdrawal = withdrawalConfig

			// 	serialized, _ := json.Marshal(pairconfig)
			// 	g.redisSession.HSet("exchange:"+g.Name+":pairconfig", strings.ToUpper(data.ID), serialized)
			// 	g.redisSession.Expire("exchange:"+g.Name+":pairconfig", 6*time.Minute)
			// }

			time.Sleep(5 * time.Minute)
		}
	}()
}

// GetExchangeAccountInfo retrieves balances for all enabled currencies for the
// bitmex exchange
func (b *Bitmex) GetExchangeAccountInfo() (exchange.AccountInfo, error) {
	var response exchange.AccountInfo
	return response, nil
}

// UpdateTicker updates and returns the ticker for a currency pair
func (b *Bitmex) UpdateTicker() {

}

// UpdateOrderbook updates and returns the orderbook for a currency pair
func (b *Bitmex) UpdateOrderbook(normalized, currency string) (orderbook.Base, error) {
	var orderBook orderbook.Base
	return orderBook, nil
}

// GetExchangeHistory returns historic trade data since exchange opening.
func (b *Bitmex) GetExchangeHistory(p pair.CurrencyPair, assetType string) ([]exchange.TradeHistory, error) {
	var resp []exchange.TradeHistory

	return resp, errors.New("trade history not yet implemented")
}

// NewTriangularOrder execute thee order on the exchange
func (b *Bitmex) NewTriangularOrder(orders []exchange.NewOrder) (exchange.ResponseTriangular, error) {
	var response exchange.ResponseTriangular
	return response, nil
}
