package streaming

import (
	"context"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gogo/protobuf/proto"
	client "github.com/influxdata/influxdb/client/v2"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/mongodb/mongo-go-driver/core/connstring"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/oleiade/lane"
	"github.com/sirupsen/logrus"
)

// Streaming ...
type Streaming struct {
	Consumer        *kafka.Consumer
	Subscription    []string
	GroupID         string
	InfluxQueue     *lane.Queue
	MongoQueue      *lane.Queue
	InfluxUDPClinet client.Client
	MongoDB         *mongo.Database
	MongoClient     *mongo.Client
	Verbose         bool
}

// InitKafkaConnection initializes
func InitKafkaConnection(topics []string, groupID string) *Streaming {
	k, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               os.Getenv("KAFKA_BROKERS"),
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
	connection := new(Streaming)
	connection.Consumer = k
	connection.Subscription = topics
	connection.GroupID = groupID
	return connection
}

// StartReading ...
func (s *Streaming) StartReading() {
	s.StartMongo()
	s.StartInflux()
	s.ReadStreaming()
}

// ReadStreaming ...
func (s *Streaming) ReadStreaming() {
	go func() {
		for {
			select {
			case ev := <-s.Consumer.Events():
				switch e := ev.(type) {
				case kafka.AssignedPartitions:
					if s.Verbose {
						fmt.Fprintf(os.Stderr, "Kafka AssignedPartitions - %% %v\n", e)
					}
					s.Consumer.Assign(e.Partitions)
				case kafka.RevokedPartitions:
					if s.Verbose {
						fmt.Fprintf(os.Stderr, "Kafka RevokedPartitions %% %v\n", e)
					}
					s.Consumer.Unassign()
				case *kafka.Message:
					switch e.Value[0] {
					case 0:
						var trade = &pbAPI.Trade{}
						err := proto.Unmarshal(e.Value[4:], trade)
						if err != nil {
							logrus.Error("proto.Marshal error: ", err)
						}
						s.InfluxQueue.Enqueue(trade)
						s.MongoQueue.Enqueue(trade)
					case 1:
						var orderbook = &pbAPI.Orderbook{}
						err := proto.Unmarshal(e.Value[4:], orderbook)
						if err != nil {
							logrus.Error("proto.Marshal error: ", err)
						}
						logrus.Info(orderbook)
					default:

					}
				case kafka.PartitionEOF:
					if s.Verbose {
						fmt.Printf("%% Reached %v\n", e)
					}
				case kafka.Error:
					if s.Verbose {
						fmt.Fprintf(os.Stderr, "%% Kafka Error: %v\n", e)
					}
				}
			}
		}
	}()
}

// StartMongo ...
func (s *Streaming) StartMongo() {
	var err error
	var CS = connstring.ConnString{
		Hosts:    []string{os.Getenv("MONGODB_CONNECTION")},
		Username: os.Getenv("MONGODB_USERNAME"),
		Password: os.Getenv("MONGODB_PASSWORD"),
	}
	s.MongoClient, err = mongo.NewClientFromConnString(CS)
	if err != nil {
		logrus.Error("Problem to connect with mongo ", err)
		os.Exit(1)
	}
	// connect to mongo
	err = s.MongoClient.Connect(context.Background())
	if err != nil {
		logrus.Error("Mongo ", err)
		os.Exit(1)
	}
	s.MongoDB = s.MongoClient.Database(os.Getenv("MONGODB_DATABASE_NAME"), nil)
	s.MongoQueue = lane.NewQueue()
	s.InitQueueMongo()
}

// StartInflux ...
func (s *Streaming) StartInflux() {
	var err error
	s.InfluxUDPClinet, err = client.NewUDPClient(client.UDPConfig{Addr: os.Getenv("INFLUX_UDP_ADDR")})
	if err != nil {
		logrus.Error("Problem to connect on influxdb ", err.Error())
		os.Exit(1)
	}
	s.InfluxQueue = lane.NewQueue()
	s.InitInfluxQueue()
}
