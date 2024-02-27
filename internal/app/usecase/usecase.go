package usecase

import (
	"context"
	"fmt"
	"ratequotes/internal/app/adapter"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type QuotesUseCase interface {
	UpdateQuotes(*gin.Context, string, string, adapter.QuotesRepository, adapter.ExternalApiRepository)
	GetQuotesById(*gin.Context, int) *redis.StringCmd
	GetLastQuotes(*gin.Context, string)
}
type useCase struct {
	rclient *redis.Client
}

func NewUserUsecase() QuotesUseCase {
	return &useCase{}
}

func (uc *useCase) UpdateQuotes(gin *gin.Context, currencyCode string, id string, quotesRepos adapter.QuotesRepository, facadeRep adapter.ExternalApiRepository) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err, exexchangeRatesResponse := facadeRep.GetCurrencyRate(ctx, currencyCode)

	if err != nil {
		panic("ERROR EXTERNAL REQUEST")
	}
	quotesRepos.SetQuotes(exexchangeRatesResponse, id)

}

func (uc *useCase) GetQuotesById(gin *gin.Context, updateId int) *redis.StringCmd {
	fmt.Println(updateId)

	result := uc.rclient.Get(fmt.Sprint(updateId))

	return result
	// panic("work")
}

func (uc *useCase) GetLastQuotes(gin *gin.Context, currencyCode string) {
	fmt.Println(currencyCode)
	// panic("work")
}
