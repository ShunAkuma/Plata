package controller

import (
	"net/http"
	"ratequotes/internal/app/adapter"
	"ratequotes/internal/app/model"
	"ratequotes/internal/app/usecase"

	"github.com/gin-gonic/gin"
)

type controller interface {
	UpdateQuotesContolller(*gin.Context, usecase.QuotesUseCase, adapter.QuotesRepository, adapter.ExternalApiRepository)
	GetQuotes(*gin.Context, usecase.QuotesUseCase, adapter.QuotesRepository)
	GetLastQuotes(*gin.Context, usecase.QuotesUseCase, adapter.QuotesRepository)
}

type QuotesController struct{}

func NewController() controller {
	return &QuotesController{}
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
	var currencyCode string

	if ctx.Query("CurrencyCode") == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{StatusCode: 400, Message: "Error: not enough input parameters", ResultObj: nil})
		return
	}
	if err := ctx.ShouldBindQuery(&currencyCode); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{StatusCode: 400, Message: "Error: bind value", ResultObj: nil})
		return
	}
	quotesUseCase.UpdateQuotes(ctx, currencyCode)
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

	var updateId int

	if ctx.Query("UpdateId") == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{StatusCode: 400, Message: "Error: not enough input parameters", ResultObj: nil})
		return
	}
	if err := ctx.ShouldBindQuery(&updateId); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{StatusCode: 400, Message: "Error: bind value", ResultObj: nil})
		return
	}
	quotesUseCase.GetQuotesById(ctx, updateId)
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
		ctx.JSON(http.StatusBadRequest, model.Response{StatusCode: 400, Message: "Error: not enough input parameters", ResultObj: nil})
		return
	}
	if err := ctx.ShouldBindQuery(&currencyCode); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{StatusCode: 400, Message: "Error: bind value", ResultObj: nil})
		return
	}

	quotesUseCase.GetLastQuotes(ctx, currencyCode)
}
