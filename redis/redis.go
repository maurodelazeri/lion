package redis

import (
	"os"

	"github.com/go-redis/redis"
)

// RedisClient ...
var RedisClient *redis.Client

func init() {
	InitEngine()
}

// InitEngine initializes our Database Connection
func InitEngine() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_CONNECTION"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})
}
