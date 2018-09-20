package influx

import (
	"os"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

var influxClinet client.Client

func init() {
	// Make client
	c, err := client.NewUDPClient(client.UDPConfig{Addr: os.Getenv("INFLUX_UDP_ADDR")})
	if err != nil {
		panic(err.Error())
	}
	influxClinet = c
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
