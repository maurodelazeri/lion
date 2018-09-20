package influx

import (
	"log"
	"os"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

func aaaaa(name string,tags map[string]string, fields map[string]interface{},t ...time.Time) {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     os.Getenv("INFLUX_HOST"),
		Username: os.Getenv("INFLUX_USERNAME"),
		Password: os.Getenv("INFLUX_PASSWORD"),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  os.Getenv("INFLUX_DATABASE"),
		Precision: "ns",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

export INFLUX_HOST=""
export INFLUX_USERNAME=""