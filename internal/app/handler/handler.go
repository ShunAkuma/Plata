package handler

import (
	"github.com/gin-gonic/gin"

	_ "ratequotes/docs"

	"ratequotes/internal/app/controller"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type IHandler interface {
	Handler(router *gin.RouterGroup)
}
type Handler struct {
	controller controller.Controller
}

func NewHandler(conroller controller.Controller) IHandler {
	return &Handler{
		controller: conroller,
	}
}

func (h Handler) Handler(router *gin.RouterGroup) {
	router.PATCH("/api/quotes", h.controller.UpdateQuotesContolller)
	router.GET("/api/quotes/:id", h.controller.GetQuotes)
	router.GET("/api/quotes/currency/:currency", h.controller.GetLastQuotes)

	//Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
