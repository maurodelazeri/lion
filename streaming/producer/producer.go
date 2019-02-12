package kafkaproducer

// advertised.host.name=<Kafka Running Machine IP>
// server.properties
import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

/*
./kafka-console-consumer --bootstrap-server 192.168.3.100:9092 --topic BTC_USD.COINBASEPRO.trade --from-beginning
./kafka-console-producer --broker-list 192.168.3.100:9092 --topic BTC_USD.COINBASEPRO.trade
./kafka-topics --list --zookeeper 192.168.3.100:2181
*/

// Producer ...
var Producer *kafka.Producer

func init() {
	InitEngine()
}

// InitEngine ...
func InitEngine() {
	brokers := os.Getenv("KAFKA_BROKERS")
	client, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":            brokers,
		"message.max.bytes":            100857600,
		"queue.buffering.max.messages": 1000000,
		"queue.buffering.max.kbytes":   1000000,
		//"queue.buffering.max.ms":       5000,
		//"batch.num.messages":           0,
		"log.connection.close": false})
	if err != nil {
		logrus.Error("No connection with kafka, ", err)
		os.Exit(1)
	}
	Producer = client
}

// PublishMessageAsync send a message to kafka server (asynchronous)
func PublishMessageAsync(topic string, message []byte, partition int64, verbose bool) error {
	// err := Producer.Produce(&kafka.Message{
	// 	TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	// 	Value:          message,
	// }, nil)
	// if err != nil {
	// 	logrus.Warn("Problem to publish message topic:"+topic+" ", err)
	// }
	// //Producer.Flush(15 * 1000)
	// return err
	deliveryChan := make(chan kafka.Event)

	err := Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, deliveryChan)
	if err != nil {
		logrus.Warn("Problem with producer message, ", err)
		close(deliveryChan)
		return err
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)

	if verbose {
		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
			close(deliveryChan)
			return m.TopicPartition.Error
		} else {
			fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
	}
	close(deliveryChan)
	return nil
}

// PublishMessageSync send a message to kafka server (synchronous)
func PublishMessageSync(topic string, message []byte, partition int64, verbose bool) error {
	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	err := Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, deliveryChan)
	if err != nil {
		logrus.Warn("Problem with producer message, ", err)
		close(deliveryChan)
		return err
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)

	if verbose {
		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
			close(deliveryChan)
			return m.TopicPartition.Error
		} else {
			fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
	}
	close(deliveryChan)
	return nil
}
