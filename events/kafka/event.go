package kafkaevents

import (
	"log"
	"runtime"
	"time"

	"github.com/Jeffail/tunny"
	"github.com/maurodelazeri/lion/common"
	"github.com/maurodelazeri/lion/protobuf/events"
	"github.com/maurodelazeri/lion/streaming/kafka/producer"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/sirupsen/logrus"
)

var pool *tunny.Pool

// Producer stores the initialization of kafka
var Producer *producer.Producer

// Job ...
type Job struct {
	event     events.Event
	topic     string
	partition int32
	verbose   bool
}

func init() {
	pool = tunny.NewFunc(runtime.NumCPU(), worker)

	Producer = new(producer.Producer)
	Producer.InitializeProducer()
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
	Producer.PublishMessageAsync(j.topic, eventData, j.partition, j.verbose)
	return nil
}

func processJob(pool *tunny.Pool, event events.Event, topic string, partition int32, verbose bool) {
	j := &Job{event: event, topic: topic, partition: partition, verbose: verbose}
	_, err := pool.ProcessTimed(j, time.Minute*5)
	if err == tunny.ErrJobTimedOut {
		log.Printf("problem to process job %v", err)
	}
}

// Start ...
func Start(topic string, data []byte, partition int32, verbose bool) {
	err := Producer.PublishMessageSync(topic, data, partition, verbose)
	if err != nil {
		logrus.Error("Problem to Marshal order request ", err)
	}
}

// CreateBaseEvent create a initial event
func CreateBaseEvent(id, event, account, container, user, strategy string) *events.Event {
	return &events.Event{
		Event:     event,
		Id:        id,
		Account:   account,
		User:      user,
		Container: container,
		Strategy:  strategy,
		Error:     false,
		Timestamp: common.MakeTimestamp(),
	}
}

// PublishEvent to kafka
func PublishEvent(event *events.Event, topic string, partition int32, verbose bool) {
	go func(event *events.Event, topic string, partition int32, verbose bool) {
		processJob(pool, *event, topic, partition, verbose)
	}(event, topic, partition, verbose)
}
