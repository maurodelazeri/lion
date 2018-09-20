package influx

import (
	"os"
	"time"

	pbAPI "github.com/maurodelazeri/lion/protobuf/api"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/oleiade/lane"
	"github.com/sirupsen/logrus"
)

var (
	//InfluxQueue ...
	InfluxQueue  *lane.Queue
	influxClinet client.Client
)

func init() {
	// Make client
	c, err := client.NewUDPClient(client.UDPConfig{Addr: os.Getenv("INFLUX_UDP_ADDR")})
	if err != nil {
		panic(err.Error())
	}
	influxClinet = c
	InitQueue()
}

// InitQueue the operations are in a queue to guarantee the correct execution order
func InitQueue() {
	// Let's handle the clients asynchronously
	go func() {
		for {
			for InfluxQueue.Head() != nil {
				item := InfluxQueue.Dequeue()
				Worker(item)
			}
		}
	}()
}

// Worker execute sequencial execution based on the received instructions
func Worker(item interface{}) {
	switch t := item.(type) {
	case pbAPI.Trade:
		// Create a point and add to batch
		tags := map[string]string{
			"venue":   t.GetVenue().String(),
			"product": t.GetProduct().String(),
			"side":    t.GetOrderSide().String(),
		}
		fields := map[string]interface{}{
			"price": t.GetPrice(),
			"size":  t.GetSize(),
			"side":  pbAPI.OrderType_value[t.GetOrderSide().String()],
		}
		InsertInflux("trade", tags, fields)
	default:
		logrus.Error("Data is not an array ", t)
	}
}

// InsertInflux ...
func InsertInflux(name string, tags map[string]string, fields map[string]interface{}, t ...time.Time) {
	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "ns",
	})
	pt, err := client.NewPoint(name, tags, fields, time.Now())
	if err != nil {
		panic(err.Error())
	}
	bp.AddPoint(pt)
	// Write the batch
	influxClinet.Write(bp)
}
