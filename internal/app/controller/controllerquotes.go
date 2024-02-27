package controller

import (
	"net/http"
	"ratequotes/internal/app/adapter"
	"ratequotes/internal/app/model"
	"ratequotes/internal/app/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type controller interface {
	UpdateQuotesContolller(*gin.Context, usecase.QuotesUseCase, adapter.QuotesRepository, adapter.ExternalApiRepository)
	GetQuotes(*gin.Context, usecase.QuotesUseCase, adapter.QuotesRepository)
	GetLastQuotes(*gin.Context, usecase.QuotesUseCase, adapter.QuotesRepository)
}

type QuotesController struct {
	symbols map[string]bool
}

func NewController(symbolsMap map[string]bool) controller {
	return &QuotesController{
		symbols: symbolsMap,
	}
}

// UdapteQuotesRate godoc
// @Summary      Update Quotes Rate
// @Description  Update
// @Tags         Quotes
// @Accept       json
// @Produce      json
// @Param        code   path      int  true  "Сurrency code"
// @Success      200  {array}   model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /updatequotes [post]
func (c QuotesController) UpdateQuotesContolller(ctx *gin.Context, quotesUseCase usecase.QuotesUseCase, quotesRepos adapter.QuotesRepository, facadeRep adapter.ExternalApiRepository) {
	var currencyCode string = ctx.Query("CurrencyCode")

	if currencyCode == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Error: not enough input parameters", ResultObj: nil})
		return
	}

	if err := ctx.ShouldBindQuery(&currencyCode); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Error: bind value", ResultObj: nil})
		return
	}

	if _, ok := c.symbols[currencyCode]; !ok {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Bad currency code", ResultObj: nil})
		return
	}

	requestId := uuid.New()

	go quotesUseCase.UpdateQuotes(ctx, currencyCode, requestId.String(), quotesRepos, facadeRep)

	ctx.JSON(http.StatusOK, model.Response{Message: "Request Id", ResultObj: requestId})
	return
}

// GetQuotesById godoc
// @Summary      Get quotes rate by id
// @Description  Get
// @Tags         Quotes
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "quotes"
// @Success      200  {array}   model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /updatequotes [get]
func (c QuotesController) GetQuotes(ctx *gin.Context, quotesUseCase usecase.QuotesUseCase, quotesRepos adapter.QuotesRepository) {

	var updateId string = ctx.Query("UpdateId")

	if updateId == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Error: not enough input parameters", ResultObj: nil})
		return
	}

	if err := ctx.ShouldBindQuery(&updateId); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Error: bind value", ResultObj: nil})
		return
	}

	err, quotesResponse := quotesUseCase.GetQuotesById(ctx, updateId, quotesRepos)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Something went wrong", ResultObj: err})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Message: "Data", ResultObj: quotesResponse})
	return
}

// GetLastQuotes godoc
// @Summary      Get last quotes rate with time
// @Description  Get
// @Tags         Quotes
// @Accept       json
// @Produce      json
// @Param        code   path      string  true  "Сurrency code"
// @Success      200  {array}   model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /lastquotes [get]
func (c QuotesController) GetLastQuotes(ctx *gin.Context, quotesUseCase usecase.QuotesUseCase, quotesRepos adapter.QuotesRepository) {
	var currencyCode string

	if ctx.Query("CurrencyCode") == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Error: not enough input parameters", ResultObj: nil})
		return
	}
	if err := ctx.ShouldBindQuery(&currencyCode); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Error: bind value", ResultObj: nil})
		return
	}

	quotesUseCase.GetLastQuotes(ctx, currencyCode)
}
