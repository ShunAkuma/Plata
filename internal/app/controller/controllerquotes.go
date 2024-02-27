package controller

import (
	"net/http"
	"ratequotes/internal/app/model"
	"ratequotes/internal/app/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller interface {
	UpdateQuotesContolller(*gin.Context)
	GetQuotes(*gin.Context)
	GetLastQuotes(*gin.Context)
}

type QuotesController struct {
	symbols       map[string]bool
	quotesUseCase usecase.QuotesUseCase
}

func NewController(symbolsMap map[string]bool, quotesUseCase usecase.QuotesUseCase) Controller {
	return &QuotesController{
		symbols:       symbolsMap,
		quotesUseCase: quotesUseCase,
	}
}

// UdapteQuotesRate godoc
// @Summary      Update Quotes Rate
// @Description  Updating the quote in the background
// @Tags         Quotes
// @Accept       json
// @Produce      json
// @Param        CurrencyCode   query      string  true  "Сurrency code"
// @Success      200  {array}   model.Response
// @Failure      400,404  {object}  model.Response
// @Router       /updatequotes [post]
func (c QuotesController) UpdateQuotesContolller(ctx *gin.Context) {
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

	go c.quotesUseCase.UpdateQuotes(ctx, currencyCode, requestId.String())

	ctx.JSON(http.StatusOK, model.Response{Message: "Request Id", ResultObj: requestId})
	return
}

// GetQuotesById godoc
// @Summary      Get quotes rate by id
// @Description  Get Get quotes rate by id from redis
// @Tags         Quotes
// @Accept       json
// @Produce      json
// @Param        UpdateId   query      string  true  "quotes"
// @Success      200  {array}   model.Response
// @Failure      400,404  {object}  model.Response
// @Router       /quotesbyid [get]
func (c QuotesController) GetQuotes(ctx *gin.Context) {

	var updateId string = ctx.Query("UpdateId")

	if updateId == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Error: not enough input parameters", ResultObj: nil})
		return
	}

	if err := ctx.ShouldBindQuery(&updateId); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Error: bind value", ResultObj: nil})
		return
	}

	err, quotesResponse := c.quotesUseCase.GetQuotesById(ctx, updateId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Something went wrong", ResultObj: err})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Message: "Data", ResultObj: quotesResponse})
	return
}

// GetLastQuotes godoc
// @Summary      Get last quotes rate with time
// @Description  Get last quotes with time and rate
// @Tags         Quotes
// @Accept       json
// @Produce      json
// @Param        CurrencyCode   query      string  true  "Сurrency code"
// @Success      200  {array}   model.Response
// @Failure      400,404  {object}  model.Response
// @Router       /lastquotes [get]
func (c QuotesController) GetLastQuotes(ctx *gin.Context) {
	currencyCode := ctx.Query("CurrencyCode")
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

	err, quotes := c.quotesUseCase.GetLastQuotes(ctx, currencyCode)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Something went wrong", ResultObj: err})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Message: "Data", ResultObj: quotes})
	return
}
