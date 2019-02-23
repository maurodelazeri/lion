package event

import (
	"log"
	"runtime"
	"time"

	"github.com/maurodelazeri/lion/streaming/producer"

	"github.com/Jeffail/tunny"
	eventAPI "github.com/maurodelazeri/lion/protobuf/heraldsquareAPI"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/sirupsen/logrus"
)

var pool *tunny.Pool

// Job ...
type Job struct {
	event     eventAPI.Event
	topic     string
	partition int64
	verbose   bool
}

func init() {
	pool = tunny.NewFunc(runtime.NumCPU(), worker)
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
	err = kafkaproducer.PublishMessageSync(j.topic, eventData, j.partition, j.verbose)
	if err != nil {
		logrus.Error("Problem PublishMessageSync request ", err)
		return err
	}
	return nil
}

func processJob(pool *tunny.Pool, event eventAPI.Event, topic string, partition int64, verbose bool) {
	j := &Job{event: event, topic: topic, partition: partition, verbose: verbose}
	_, err := pool.ProcessTimed(j, time.Minute*5)
	if err == tunny.ErrJobTimedOut {
		log.Printf("problem to process job %v", err)
	}
}

// CreateBaseEvent create a initial event
func CreateBaseEvent(systemEventID, function, message, payload string, err bool, UserID int64, system eventAPI.System) *eventAPI.Event {
	return &eventAPI.Event{
		SystemEventId: systemEventID,
		System:        system,
		Function:      function,
		UserId:        UserID,
		Message:       message,
		Payload:       payload,
		Error:         err,
		Timestamp:     time.Now().UTC().String(),
	}
}

// PublishEvent to kafka
func PublishEvent(event *eventAPI.Event, topic string, partition int64, verbose bool) {
	go func(event *eventAPI.Event, topic string, partition int64, verbose bool) {
		processJob(pool, *event, topic, partition, verbose)
	}(event, topic, partition, verbose)
}
