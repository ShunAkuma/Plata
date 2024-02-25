package externalapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Client представляет клиент для взаимодействия с внешним API.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient создает новый экземпляр клиента для взаимодействия с внешним API.
func NewClient() *Client {
	return &Client{
		BaseURL: os.Getenv("EXTERNAL_API_URL"),
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) ExternalRequest(currencyCode string) {
	ctx := context.Background()
	var response interface{}
	var path string = "currencies"
	response, err := c.do(context.WithValue(ctx, "code", currencyCode), path)
	if err != nil {
		// TODO
		fmt.Println(err)
		panic("")
	}
	fmt.Println(response)

	// currency Code in response json, if not return err or something for indicato to error

}

// Do запрос HTTP к внешнему API.
func (c *Client) do(ctx context.Context, path string) (interface{}, error) {
	// currencyCode := ctx.Value("code")

	var response interface{}
	// Формирование URL для запроса
	// url := fmt.Sprintf("%s%s", c.BaseURL, path)
	urltest := "https://api.frankfurter.app/currencies"
	// Создание HTTP-запроса
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urltest, nil)
	if err != nil {
		return nil, err
	}

	// Установка заголовка Content-Type для JSON
	req.Header.Set("Content-Type", "application/json")

	// Выполнение HTTP-запроса
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Проверка кода статуса ответа
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	// Декодирование ответа JSON в указанный объект
	if response != nil {
		err = json.NewDecoder(resp.Body).Decode(response)
		if err != nil {
			return nil, err
		}
	}
	fmt.Println(response)
	return response, nil
}
