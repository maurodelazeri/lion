package config

import (
	"database/sql"
	"fmt"

	"github.com/maurodelazeri/concurrency-map-slice"
	"github.com/maurodelazeri/lion/postgres"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
)

// Cfg stores a config
var Cfg Config

// Config holds the venues individual config
type Config struct {
	Venues *utils.ConcurrentMap
}

// VenueConfig holds all the information needed for each enabled Venue.
type VenueConfig struct {
	Venue    *pbAPI.Venue
	Products map[string]*pbAPI.Product
}

// LoadConfig loads your configuration file into your configuration object
func (c *Config) LoadConfig() error {
	var rows *sql.Rows
	var err error

	query := fmt.Sprintf(`SELECT venue_id,name,venue_description,api_key,api_secret,passphrase,spot,futures,options,enabled, FROM venues`)
	if rows, err = postgres.PostgresDB.Query(query); err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		venue := &pbAPI.Venue{}
		args := []interface{}{
			&venue.VenueId, &venue.Name, &venue.VenueDescription, &venue.ApiKey, &venue.ApiSecret, &venue.Passphrase, &venue.Spot, &venue.Futures, &venue.Options, &venue.Enabled}
		if err = rows.Scan(args...); err != nil {
			return err
		}
		venueConf := &VenueConfig{}
		venueConf.Venue = venue
		venueConf.Products = make(map[string]*pbAPI.Product)
	}

	query = fmt.Sprintf(`SELECT product_id,venue_id, base_currency,quote_currency,venue_symbol_identifier,kind,individual_connection,streaming_save,
	 minimum_orders_size,step_size,price_precision,taker_fee,maker_fee,settlement,expiration,enabled, COALESCE((SELECT name FROM currencies WHERE currency_id=products.base_currency), '') || '-' || COALESCE((SELECT name FROM currencies
			WHERE currency_id=products.quote_currency), '') as system_symbol_identifier FROM products`)
	if rows, err = postgres.PostgresDB.Query(query); err != nil {
		return err
	}
	defer rows.Close()
	products := make(map[int32]*pbAPI.Product)
	for rows.Next() {
		product := &pbAPI.Product{}
		args := []interface{}{&product.ProductId, &product.VenueId, &product.BaseCurrency, &product.QuoteCurrency, &product.VenueSymbolIdentifier, &product.Kind, &product.IndividualConnection,
			&product.StreamingSave, &product.MinimumOrdersSize, &product.StepSize, &product.PricePrecision, &product.TakerFee, &product.MakerFee, &product.Settlement, &product.Expiration, &product.Enabled, &product.SystemSymbolIdentifier}

		if err = rows.Scan(args...); err != nil {
			return err
		}
		products[product.VenueId] = product
	}

	// query = fmt.Sprintf(`SELECT venue_id,name,venue_description,api_key,api_secret,passphrase,spot,futures,options,enabled, FROM venues`)
	// if rows, err := postgres.PostgresDB.Query(query); err != nil {
	// 	return err
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	venue := &pbAPI.Venue{}
	// 	args := []interface{}{
	// 		&venue.VenueId, &venue.Name, &venue.VenueDescription, &venue.APIKey, &venue.APISecret, &venue.Passphrase, &venue.Spot, &venue.Futures, &venue.Options, &venue.Enabled}
	// 	if err = rows.Scan(args...); err != nil {
	// 		return err
	// 	}
	// 	c.Venue[venue.Name] = venue
	// }

	// coll := mongodb.MongoDB.Collection("venue_products")
	// filter := bson.NewDocument(bson.EC.Boolean("enabled", true))
	// cursor, err := coll.Find(context.Background(), filter)
	// if err != nil {
	// 	return err
	// }

	// for cursor.Next(context.Background()) {
	// 	var item DBconfig
	// 	venueConf := make(map[string]Product)
	// 	if err := cursor.Decode(&item); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	for _, prods := range item.Product {
	// 		if prods.Enabled {
	// 			venueConf[prods.APIName] = Product{
	// 				IndividualConnection: prods.IndividualConnection,
	// 				VenueName:            prods.VenueName,
	// 				APIName:              prods.APIName,
	// 				MinimumOrdersSize:    prods.MinimumOrdersSize,
	// 				StepSize:             prods.StepSize,
	// 				MakerFee:             prods.MakerFee,
	// 				TakerFee:             prods.TakerFee,
	// 				Enabled:              prods.Enabled,
	// 			}
	// 		}
	// 	}
	// 	finalConf := VenueConfig{
	// 		Protobuf:   item.Protobuf,
	// 		Enabled:    item.Enabled,
	// 		APIKey:     item.APIKey,
	// 		APISecret:  item.APISecret,
	// 		Passphrase: item.Passphrase,
	// 		Name:       item.Name,
	// 		Products:   venueConf,
	// 	}
	// 	Cfg.Venues.Put(item.Name, finalConf)
	// }
	return nil
}
