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
	InfluxQueue     *lane.Queue
	influxUDPClinet client.Client
)

func init() {
	// Make client
	c, err := client.NewUDPClient(client.UDPConfig{Addr: os.Getenv("INFLUX_UDP_ADDR")})
	if err != nil {
		logrus.Error("Problem to connect on influxdb ", err.Error())
		os.Exit(1)
	}
	influxUDPClinet = c
	InfluxQueue = lane.NewQueue()
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
	case *pbAPI.Trade:
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
		InsertInflux("trade", tags, fields, time.Unix(0, int64(t.Timestamp)*int64(time.Microsecond)))
	default:
		logrus.Error("Influx not found a correct type ", t)
	}
}

// InsertInflux ...
func InsertInflux(name string, tags map[string]string, fields map[string]interface{}, t time.Time) {
	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "us",
	})
	pt, err := client.NewPoint(name, tags, fields, t)
	if err != nil {
		panic(err.Error())
	}
	bp.AddPoint(pt)
	// Write the batch
	influxUDPClinet.Write(bp)
}
