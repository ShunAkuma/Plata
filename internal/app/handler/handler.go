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
	conroller controller.Controller
}

func NewHandler(conroller controller.Controller) IHandler {
	return &Handler{
		conroller: conroller,
	}
}

func (h Handler) Handler(router *gin.RouterGroup) {

	router.POST("/updatequotes", h.conroller.UpdateQuotesContolller)
	router.GET("/quotesbyid", h.conroller.GetQuotes)
	router.GET("/lastquotes", h.conroller.GetLastQuotes)

	//Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
