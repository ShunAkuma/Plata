package main

import (
	"log"
	"ratequotes/internal/app/handler"
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

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func init() {

	router := gin.Default()
	router.Use(Logger())
	handler.Handler(&router.RouterGroup)

	router.Run()
}
