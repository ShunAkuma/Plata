package handler

import (
	"ratequotes/internal/app/adapter"
	"ratequotes/internal/app/usecase"

	"github.com/gin-gonic/gin"

	_ "ratequotes/docs"

	"ratequotes/internal/app/controller"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Handler(router *gin.RouterGroup, quotesUseCase usecase.QuotesUseCase, quotesRepos adapter.QuotesRepository, facadeRepos adapter.ExternalApiRepository, symbolsMap map[string]bool) {
	quotesController := controller.NewController(symbolsMap)

	router.POST("/updatequotes", func(ctx *gin.Context) {
		quotesController.UpdateQuotesContolller(ctx, quotesUseCase, quotesRepos, facadeRepos)
	})
	router.GET("/quotesbyid", func(ctx *gin.Context) {
		quotesController.GetQuotes(ctx, quotesUseCase, quotesRepos)
	})
	router.GET("/lastquotes", func(ctx *gin.Context) {
		quotesController.GetLastQuotes(ctx, quotesUseCase, quotesRepos)
	})

	//Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
