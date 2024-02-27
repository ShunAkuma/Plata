package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

func NewClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PAS"), // no password set
		DB:       0,                      // use default DB
	})
	if err := client.Ping().Err(); err != nil {
		return nil, fmt.Errorf("error connecting to Redis: %v", err)
	}

	log.Println("Connected to Redis")

	return client, nil
}
