package usecase

import (
	"log"
	"ratequotes/internal/app/adapter"
	"ratequotes/internal/app/model"
	"time"

	"github.com/gin-gonic/gin"
)

type QuotesUseCase interface {
	UpdateQuotes(*gin.Context, string, string)
	GetQuotesById(*gin.Context, string) (error, model.ResponseQuotesModel)
	GetLastQuotes(*gin.Context, string) (error, model.Quotes)
}
type useCase struct {
	quotesRepos adapter.QuotesRepository
	facadeRep   adapter.ExternalApiRepository
}

func NewUserUsecase(q adapter.QuotesRepository, f adapter.ExternalApiRepository) QuotesUseCase {
	return &useCase{
		quotesRepos: q,
		facadeRep:   f,
	}
}

func (uc *useCase) UpdateQuotes(gin *gin.Context, currencyCode string, id string) {

	err, exexchangeRatesResponse := uc.facadeRep.GetCurrencyRate(currencyCode)
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

	err = uc.quotesRepos.SetQuotes(quotesModel, id)
	if err != nil {
		log.Println(err)
	}

	return

}

func (uc *useCase) GetQuotesById(gin *gin.Context, updateId string) (error, model.ResponseQuotesModel) {
	var model model.ResponseQuotesModel

	err, quotesData := uc.quotesRepos.GetQuotesById(updateId)
	if err != nil {
		return err, model
	}
	model.TimeUpdate = quotesData.TimeUpdate
	model.Rates = quotesData.Rates
	model.Base = quotesData.Base

	return nil, model

}

func (uc *useCase) GetLastQuotes(gin *gin.Context, currencyCode string) (error, model.Quotes) {
	err, id := uc.quotesRepos.GetQuotesBySymbol(currencyCode)
	if err != nil {
		return err, model.Quotes{}
	}

	err, quotesModel := uc.quotesRepos.GetQuotesById(id)

	if err != nil {
		log.Println(err)
		return err, model.Quotes{}
	}

	return nil, quotesModel
}
