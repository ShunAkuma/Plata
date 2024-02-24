package handler

import (
	"net/http"
	"ratequotes/internal/app/usecase"

	"github.com/gin-gonic/gin"

	_ "ratequotes/docs"

	"ratequotes/internal/app/model"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Handler(router *gin.RouterGroup) {
	quotesUseCase := usecase.NewUserUsecase()

	router.POST("/updatequotes", func(ctx *gin.Context) {
		var currencyCode string

		if ctx.Query("CurrencyCode") == "" {
			ctx.JSON(http.StatusBadRequest, model.Response{StatusCode: 400, Message: "Error: not enough input parameters", ResultObj: nil})
			return
		}
		if err := ctx.ShouldBindQuery(&currencyCode); err != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{StatusCode: 400, Message: "Error: bind value", ResultObj: nil})
			return
		}
		// !! Return update ID (maybe UUID)
		quotesUseCase.UpdateQuotes(ctx, currencyCode)
	})

	router.GET("/quotesbyid", func(ctx *gin.Context) {
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
	})
	router.GET("/lastquotes", func(ctx *gin.Context) {
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
	})

	//Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
