package natsevents

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/Jeffail/tunny"
	nats "github.com/nats-io/go-nats"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/sirupsen/logrus"
)

var pool *tunny.Pool

// NatsClient stores the initialization of nats client
var NatsClient *nats.Conn

// Job ...
type Job struct {
	event     interface{}
	topic     string
	partition int32
	verbose   bool
}

func init() {
	pool = tunny.NewFunc(runtime.NumCPU(), worker)

	conn, err := nats.Connect(os.Getenv("NATS_SERVER"))
	if err != nil {
		logrus.Info("No connection with nats server, ", err)
		os.Exit(1)
	}
	NatsClient = conn
}

func worker(work interface{}) interface{} {
	switch w := work.(type) {
	case *Job:
		return w.build()
	}
	return "Couldn't find work type"
}

func (j *Job) build() error {
	eventData, err := ffjson.Marshal(&j.event)
	if j.verbose {
		logrus.Info(string(eventData))
	}
	if err != nil {
		logrus.Error("Problem to Marshal order request ", err)
		return err
	}
	// Produce messages to topic (asynchronously)
	err = NatsClient.Publish(j.topic, eventData)
	if err != nil {
		logrus.Warn("Problem to publish message, ", err)
	}
	return nil
}

func processJob(pool *tunny.Pool, event interface{}, topic string, partition int32, verbose bool) {
	j := &Job{event: event, topic: topic, partition: partition, verbose: verbose}
	_, err := pool.ProcessTimed(j, time.Minute*5)
	if err == tunny.ErrJobTimedOut {
		log.Printf("problem to process job %v", err)
	}
}

// Start ...
func Start(topic string, data []byte, partition int32, verbose bool) {
	// err := Producer.PublishMessageSync(topic, data, partition, verbose)
	// if err != nil {
	// 	logrus.Error("Problem to Marshal order request ", err)
	// }
}

// PublishEvent to nats
func PublishEvent(event interface{}, topic string, partition int32, verbose bool) {
	go func(event interface{}, topic string, partition int32, verbose bool) {
		processJob(pool, event, topic, partition, verbose)
	}(event, topic, partition, verbose)
}
