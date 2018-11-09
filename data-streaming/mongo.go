package streaming

import (
	"context"

	"github.com/maurodelazeri/lion/common"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/sirupsen/logrus"
)

// InitQueueMongo ....
func (s *Streaming) InitQueueMongo() {
	// Let's handle the clients asynchronously
	go func() {
		for {
			for s.MongoQueue.Head() != nil {
				item := s.MongoQueue.Dequeue()
				s.WorkerTrades(item)
			}
		}
	}()
}

// WorkerTrades execute sequencial execution based on the received instructions
func (s *Streaming) WorkerTrades(item interface{}) {
	switch t := item.(type) {
	case *pbAPI.Trade:
		arrBids := bson.NewArray()
		for _, values := range t.GetBids() {
			value := bson.VC.DocumentFromElements(
				bson.EC.Double("price", values.GetPrice()),
				bson.EC.Double("volume", values.GetVolume()),
			)
			arrBids.Append(value)
		}
		arrAsks := bson.NewArray()
		for _, values := range t.GetAsks() {
			value := bson.VC.DocumentFromElements(
				bson.EC.Double("price", values.GetPrice()),
				bson.EC.Double("volume", values.GetVolume()),
			)
			arrAsks.Append(value)
		}
		coll := s.MongoDB.Collection("trades")
		_, err := coll.InsertOne(
			context.Background(),
			bson.NewDocument(
				bson.EC.Int32("venue", pbAPI.Venue_value[t.GetVenue().String()]),
				bson.EC.Int32("product", pbAPI.Product_value[t.GetProduct().String()]),
				bson.EC.Time("timestamp", common.MakeTimestampFromInt64(t.GetTimestamp())),
				bson.EC.Double("price", t.GetPrice()),
				bson.EC.Double("volume", t.GetVolume()),
				bson.EC.Int32("side", pbAPI.Side_value[t.GetOrderSide().String()]),
				bson.EC.Int32("venue_type", pbAPI.VenueType_value[t.GetVenueType().String()]),
				bson.EC.Array("bids", arrBids),
				bson.EC.Array("asks", arrAsks),
			))
		if err != nil {
			logrus.Error("Problem to insert on mongo (trade)", err)
		}
	}
}
