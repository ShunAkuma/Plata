package adapter

import (
	"fmt"
	"ratequotes/internal/app/model"

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
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "admin", // no password set
		DB:       0,       // use default DB
	})

	client.Set(id, "test", 0)
	test := client.Ping()
	fmt.Println("-------------------", test, "------------------")
	// _, err := json.Marshal(exchangeRates)
	// if err != nil {
	// 	panic(err)
	// }
	// add := rc.redis.Options().Addr
	// fmt.Println(add)
	// status := rc.redis.Set(id, 123232, 5*time.Second)
	// fmt.Println(status)
	// fmt.Println(err, "-----------------------------")
}

func (rc QuotesRedisRepository) GetQuotesById() {

}
