package adapter

import "github.com/go-redis/redis"

type RedisClient struct {
	redis *redis.Client
}

func (rc RedisClient) Set() {

}

func (rc RedisClient) Get() {

}
