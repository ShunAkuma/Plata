package pkg

import (
	"fmt"

	"github.com/go-redis/redis"
)

// type Client interface {
// 	Set(key string, value interface{}) error
// 	Get(key string) (string, error)
// 	// Add other Redis operations as needed
// }

// NewClient creates a new Redis client
func NewClient() (*redis.Client, error) {
	// Устанавливаем соединение с Redis сервером
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "admin", // no password set
		DB:       0,       // use default DB
	})
	test := client.Ping()
	fmt.Println("-----------", test, "----------------")
	// Проверяем ошибку подключения к Redis
	if err := client.Ping().Err(); err != nil {
		return nil, fmt.Errorf("error connecting to Redis: %v", err)
	}

	fmt.Println("Connected to Redis")

	return client, nil
}
