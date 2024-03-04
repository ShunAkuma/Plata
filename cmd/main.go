package main

import (
	"log"
	"net/http"
	"ratequotes/internal/app/adapter"
	"ratequotes/internal/app/controller"
	"ratequotes/internal/app/handler"
	"ratequotes/internal/app/usecase"
	"ratequotes/pkg"
	"time"

	"github.com/gin-gonic/gin"
)

// @title           Plata backend
// @version         1.0
// @description     Quotation server

// @host      localhost:8080
// @BasePath  /api
func main() {

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}
	symbols := map[string]bool{
		"EUR": true,
		"MXN": true,
		"USD": true,
		"RUB": true,
	}
	rclient, err := pkg.NewClient()
	if err != nil {
		log.Println("Error redis connection", err)
	}
	facadeRepos := adapter.NewFacadeApi(&httpClient)
	quotesRepository := adapter.NewQuotesRepository(rclient)
	quotesUseCase := usecase.NewUserUsecase(quotesRepository, facadeRepos)
	quotesController := controller.NewController(symbols, quotesUseCase)

	router := gin.Default()

	handler := handler.NewHandler(quotesController)
	handler.Handler(&router.RouterGroup)

	router.Run()
}
