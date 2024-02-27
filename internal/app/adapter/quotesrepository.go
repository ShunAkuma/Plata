package adapter

import (
	"encoding/json"
	"ratequotes/internal/app/model"
	"time"

	"github.com/go-redis/redis"
)

type QuotesRepository interface {
	SetQuotes(model.ExchangeRatesResponse, string)
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

func (rc QuotesRedisRepository) SetQuotes(exchangeRates model.ExchangeRatesResponse, id string) {
	data, err := json.Marshal(exchangeRates)
	if err != nil {
		panic(err)
	}
	status := rc.redis.Set(id, data, 5*time.Second)

}

func (rc QuotesRedisRepository) GetQuotesById() {

}
