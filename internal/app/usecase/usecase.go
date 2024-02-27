package usecase

import (
	"fmt"
	"log"
	"ratequotes/internal/app/adapter"
	"ratequotes/internal/app/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type QuotesUseCase interface {
	UpdateQuotes(*gin.Context, string, string, adapter.QuotesRepository, adapter.ExternalApiRepository)
	GetQuotesById(*gin.Context, string, adapter.QuotesRepository) (error, model.ResponseQuotesModel)
	GetLastQuotes(*gin.Context, string)
}
type useCase struct {
	rclient *redis.Client
}

func NewUserUsecase() QuotesUseCase {
	return &useCase{}
}

func (uc *useCase) UpdateQuotes(gin *gin.Context, currencyCode string, id string, quotesRepos adapter.QuotesRepository, facadeRep adapter.ExternalApiRepository) {

	err, exexchangeRatesResponse := facadeRep.GetCurrencyRate(currencyCode)
	if err != nil {
		log.Println("Error: external request")
		return
	}
	quotesModel := model.Quotes{
		Id:         id,
		TimeUpdate: time.Now(),
		Base:       exexchangeRatesResponse.Base,
		Rates:      exexchangeRatesResponse.Rates,
	}

	err = quotesRepos.SetQuotes(quotesModel, id)
	if err != nil {
		log.Println(err)
	}

	return

}

func (uc *useCase) GetQuotesById(gin *gin.Context, updateId string, quotesRepos adapter.QuotesRepository) (error, model.ResponseQuotesModel) {
	var model model.ResponseQuotesModel

	err, quotesData := quotesRepos.GetQuotesById(updateId)
	if err != nil {
		return err, model
	}
	model.TimeUpdate = quotesData.TimeUpdate
	model.Rates = quotesData.Rates
	model.Base = quotesData.Base

	return nil, model

}

func (uc *useCase) GetLastQuotes(gin *gin.Context, currencyCode string) {
	fmt.Println(currencyCode)
	// panic("work")
}
