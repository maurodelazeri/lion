package marketdata

import (
	"log"
	"runtime"
	"time"

	"github.com/maurodelazeri/lion/streaming/producer"

	"github.com/Jeffail/tunny"
	"github.com/sirupsen/logrus"
)

var pool *tunny.Pool

// Job ...
type Job struct {
	data      []byte
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
	err := kafkaproducer.PublishMessageSync(j.topic, j.data, j.partition, j.verbose)
	logrus.Info("KAFKA DONE")
	if err != nil {
		logrus.Error("Problem PublishMessageSync request ", err)
		return err
	}
	return nil
}

func processJob(pool *tunny.Pool, data []byte, topic string, partition int64, verbose bool) {
	j := &Job{data: data, topic: topic, partition: partition, verbose: verbose}
	_, err := pool.ProcessTimed(j, time.Minute*5)
	if err == tunny.ErrJobTimedOut {
		log.Printf("problem to process job %v", err)
	}
}

// PublishMarketData to kafka
func PublishMarketData(data []byte, topic string, partition int64, verbose bool) {
	go func(data []byte, topic string, partition int64, verbose bool) {
		processJob(pool, data, topic, partition, verbose)
	}(data, topic, partition, verbose)
}
