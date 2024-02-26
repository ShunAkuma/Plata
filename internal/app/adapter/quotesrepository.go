package adapter

import "github.com/go-redis/redis"

type QuotesRepository interface {
	SetQuotes()
	GetQuotesById()
}
type QuotesRedisRepository struct {
	redis *redis.Client
}

func NewQuotesRepository(rclient *redis.Client) QuotesRepository {
	return &QuotesRedisRepository{
		redis: rclient,
	}
}

func (rc QuotesRedisRepository) SetQuotes() {

}

func (rc QuotesRedisRepository) GetQuotesById() {

}
