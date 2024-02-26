package main

import (
	"fmt"
	"net/http"
	"ratequotes/internal/app/adapter"
	"ratequotes/internal/app/handler"
	"ratequotes/internal/app/usecase"
	"ratequotes/pkg"
	"time"

	"github.com/gin-gonic/gin"
)

// @title           Plata backend
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /api
func main() {

}

func init() {

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	quotesUseCase := usecase.NewUserUsecase()
	rclient, err := pkg.NewClient()
	if err != nil {
		fmt.Errorf("Error redis connection", err)
	}
	facadeRepos := adapter.NewFacadeApi(&httpClient)
	quotesRepository := adapter.NewQuotesRepository(rclient)
	router := gin.Default()
	handler.Handler(&router.RouterGroup, quotesUseCase, quotesRepository, facadeRepos)

	router.Run()
}
