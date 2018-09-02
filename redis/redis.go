package redis

import (
	"os"

	"github.com/oleiade/lane"
	"github.com/sirupsen/logrus"

	"github.com/go-redis/redis"
)

// Redis ...
var (
	RedisClient *redis.Client
	RedisQueue  *lane.Queue
)

func init() {
	InitEngine()
	RedisQueue = lane.NewQueue()
	InitQueue()
}

// InitEngine initializes our Database Connection
func InitEngine() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_CONNECTION"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}

// InitQueue to update the database, the operations are in a queue to guarantee the correct execution order
func InitQueue() {
	// Let's handle the clients asynchronously
	go func() {
		for {
			for RedisQueue.Head() != nil {
				item := RedisQueue.Dequeue()
				Worker(item)
			}
		}
	}()
}

// Worker execute sequencial execution based on the received array of strings with type and instructions
func Worker(item interface{}) {
	switch t := item.(type) {
	case []string:
		switch t[0] {
		case "SET":
		case "HSET":
		case "HMSET":
		default:
			return
		}
	default:
		logrus.Error("Data is not an array ", t)
	}
}
