package kafkaconsumer

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

// Kafka ...
type Kafka struct {
	Consumer     *kafka.Consumer
	Subscription []string
	GroupID      string
}

// InitKafkaConnection initializes
func InitKafkaConnection(topics []string, groupID string) *Kafka {
	k, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               os.Getenv("STREAMING_SERVERS"),
		"group.id":                        groupID,
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "latest"}})
	if err != nil {
		logrus.Error("No connection with kafka, ", err)
		os.Exit(1)
	}
	err = k.SubscribeTopics(topics, nil)
	if err != nil {
		logrus.Error("Problem to subscribe topics, ", err)
		os.Exit(1)
	}
	connection := new(Kafka)
	connection.Consumer = k
	connection.Subscription = topics
	connection.GroupID = groupID
	return connection
}
