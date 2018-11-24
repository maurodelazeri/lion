package streaming

import (
	"time"

	"github.com/maurodelazeri/lion/common"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/sirupsen/logrus"
)

// InitInfluxQueue ....
func (s *Streaming) InitInfluxQueue() {
	// Let's handle the clients asynchronously
	go func() {
		for {
			for s.InfluxQueue.Head() != nil {
				item := s.InfluxQueue.Dequeue()
				s.InfluxWorker(item)
			}
		}
	}()
}

// InfluxWorker execute sequencial execution based on the received instructions
func (s *Streaming) InfluxWorker(item interface{}) {
	switch t := item.(type) {
	case *pbAPI.Trade:
		// Create a point and add to batch
		tags := map[string]string{
			"venue":   t.GetVenue().String(),
			"product": t.GetProduct().String(),
			"side":    t.GetOrderSide().String(),
		}
		fields := map[string]interface{}{
			"price":      t.GetPrice(),
			"size":       t.GetVolume(),
			"side":       pbAPI.Side_value[t.GetOrderSide().String()],
			"venue_type": pbAPI.VenueType_value[t.GetVenueType().String()],
		}
		s.InsertInflux("trade", tags, fields, common.MakeTimestampFromInt64(t.Timestamp))
	default:
		logrus.Error("Influx not found a correct type ", t)
	}
}

// InsertInflux ...
func (s *Streaming) InsertInflux(name string, tags map[string]string, fields map[string]interface{}, t time.Time) {
	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "ns",
	})
	pt, err := client.NewPoint(name, tags, fields, t)
	if err != nil {
		panic(err.Error())
	}
	bp.AddPoint(pt)
	// Write the batch
	s.InfluxUDPClinet.Write(bp)
}
