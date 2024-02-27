package adapter

import (
	"encoding/json"
	"errors"
	"log"
	"ratequotes/internal/app/model"

	"github.com/go-redis/redis"
)

type QuotesRepository interface {
	SetQuotes(model.Quotes, string) error
	GetQuotesById(string) (error, model.Quotes)
}
type QuotesRedisRepository struct {
	redis *redis.Client
}

func NewQuotesRepository(rclient *redis.Client) QuotesRepository {
	return &QuotesRedisRepository{
		redis: rclient,
	}
}

func (rc QuotesRedisRepository) SetQuotes(exchangeRates model.Quotes, id string) error {
	data, err := json.Marshal(exchangeRates)
	if err != nil {
		log.Panicln(err)
		return err
	}
	statusCmd := rc.redis.Set(id, data, 0)
	if statusCmd.Err() != nil {
		log.Println("Ошибка при установке значения в Redis:", statusCmd.Err())
		return err
		// Дополнительная обработка ошибки здесь, если необходимо
	} else {
		log.Println("Значение успешно установлено в Redis.")
	}

	return nil

}

func (rc QuotesRedisRepository) GetQuotesById(updateId string) (error, model.Quotes) {
	result, err := rc.redis.Get(updateId).Result()
	if err != nil {
		log.Println(err)
		return err, model.Quotes{}
	}
	if result == "" {
		err = errors.New("No data in Redis with this key")
		log.Println("No data in Redis with this key")
		return err, model.Quotes{}
	}

	var quotes model.Quotes
	err = json.Unmarshal([]byte(result), &quotes)
	if err != nil {
		log.Println("Ошибка при декодировании данных из Redis:", err)
		return err, model.Quotes{}
	}

	return nil, quotes

}