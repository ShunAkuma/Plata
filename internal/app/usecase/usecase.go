package usecase

import (
	"fmt"
	"time"

	"ratequotes/internal/app/usecase/externalapi"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type QuotesUseCase interface {
	UpdateQuotes(*gin.Context, string)
	GetQuotesById(*gin.Context, int) *redis.StringCmd
	GetLastQuotes(*gin.Context, string)
	workerUpdater(string, *redis.Client, chan error)
}
type useCase struct {
	rclient *redis.Client
}

func NewUserUsecase() QuotesUseCase {
	return &useCase{}
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
func (uc *useCase) UpdateQuotes(gin *gin.Context, currencyCode string) {
	//!! Логика
	// TODO

	done := make(chan error)

	// Запуск "worker" в фоновом режиме
	go uc.workerUpdater(currencyCode, &redis.Client{}, done)

	// Ожидание завершения работы "worker"
	go func() {
		err := <-done
		if err != nil {
			panic("=======================WORK==========================  ERROR")
			// Обработка ошибки, например, отправка сообщения об ошибке в лог или оповещение администратора
		} else {
			panic("=======================WORK==========================")
		}
	}()
	gin.JSON(200, "BLALALALALALALLALA")
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
func (uc *useCase) GetQuotesById(gin *gin.Context, updateId int) *redis.StringCmd {
	fmt.Println(updateId)

	result := uc.rclient.Get(fmt.Sprint(updateId))

	return result
	// panic("work")
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
func (uc *useCase) GetLastQuotes(gin *gin.Context, currencyCode string) {
	fmt.Println(currencyCode)
	// panic("work")
}

func (uc *useCase) workerUpdater(currencyCode string, rclient *redis.Client, ch chan error) {
	// var err error
	exApiClient := externalapi.NewClient()
	exApiClient.ExternalRequest("EUR")

	time.Sleep(5 * time.Second)
	ch <- fmt.Errorf("error")
}
