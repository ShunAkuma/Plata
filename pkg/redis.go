package pkg

import (
	"errors"
	"os"

	"github.com/go-redis/redis"
)

// type Client interface {
// 	Set(key string, value interface{}) error
// 	Get(key string) (string, error)
// 	// Add other Redis operations as needed
// }

// NewClient creates a new Redis client
func NewClient() (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password:   os.Getenv("REDIS_PASSWORD"), // no password set
		DB:         0,                           // use default DB
		MaxRetries: 3,
	})

	if client == nil {
		return nil, errors.New("Error connection to redis")
	}
	return client, nil
}
