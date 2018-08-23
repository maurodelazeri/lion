package producer
//aaa
import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

/*
./kafka-console-consumer --bootstrap-server localhost:9092 --topic test --from-beginning
./kafka-console-producer --broker-list localhost:9092 --topic test
./kafka-topics --list --zookeeper localhost:2181
*/

// Producer stores the main session
type Producer struct {
	kafkalient *kafka.Producer
}

// InitializeProducer a kafka instance
func (k *Producer) InitializeProducer() {
	brokers := os.Getenv("BROKERS")
	client, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":            brokers,
		"message.max.bytes":            100857600,
		"queue.buffering.max.messages": 1000000,
		//"queue.buffering.max.ms":       5000,
		//"batch.num.messages":           0,
		"log.connection.close": false})
	//	"client.id":                    socket.gethostname(),
	//"default.topic.config": {"acks": "all"}})

	if err != nil {
		logrus.Info("No connection with kafka, ", err)
		os.Exit(1)
	}

	k.kafkalient = client
}

// PublishMessage send a message to kafka server
func (k *Producer) PublishMessage(topic string, message []byte, partition int32, verbose bool) error {
	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	// deliveryChan := make(chan kafka.Event)

	// err := k.kafkalient.Produce(&kafka.Message{
	// 	TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: partition},
	// 	Value:          message,
	// }, deliveryChan)

	// if err != nil {
	// 	fmt.Printf("Produce failed: %v\n", err)
	// }

	// e := <-deliveryChan
	// m := e.(*kafka.Message)

	// if m.TopicPartition.Error != nil {
	// 	fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	// 	return m.TopicPartition.Error
	// }

	// if verbose {
	// 	fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
	// 		*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	// }

	// close(deliveryChan)
	// return nil

	// Produce messages to topic (asynchronously)

	err := k.kafkalient.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)
	if err != nil {
		logrus.Warn("Problem to publish message, ", err)
	}

	return err
}
