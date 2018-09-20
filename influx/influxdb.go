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
		//InsertInflux("trade",)
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

// Create a point and add to batch
// tags := map[string]string{"cpu": "cpu-total"}
// fields := map[string]interface{}{
// 	"idle":   10.1,
// 	"system": 53.3,
// 	"user":   46.6,
// }

// Product              Product   `protobuf:"varint,1,opt,name=product,proto3,enum=api.Product" json:"product,omitempty"`
// 	Venue                Venue     `protobuf:"varint,2,opt,name=venue,proto3,enum=api.Venue" json:"venue,omitempty"`
// 	Timestamp            uint64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
// 	Price                float64   `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
// 	Size                 float64   `protobuf:"fixed64,5,opt,name=size,proto3" json:"size,omitempty"`
// 	OrderSide            OrderType `protobuf:"varint,6,opt,name=order_side,json=orderSide,proto3,enum=api.OrderType" json:"order_side,omitempty"`
// 	VenueType            VenueType `protobuf:"varint,7,opt,name=venue_type,json=venueType,proto3,enum=api.VenueType" json:"venue_type,omitempty"`
