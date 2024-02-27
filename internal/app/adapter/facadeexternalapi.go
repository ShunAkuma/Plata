package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"ratequotes/internal/app/model"
)

type ExternalApiRepository interface {
	GetCurrencyRate(ctx context.Context, currencyCode string) (error, model.ExchangeRatesResponse)
}

type FacadeApiRepository struct {
	httpClient *http.Client
}

func NewFacadeApi(httpClient *http.Client) ExternalApiRepository {
	return &FacadeApiRepository{
		httpClient: httpClient,
	}
}

func (fa FacadeApiRepository) GetCurrencyRate(ctx context.Context, currencyCode string) (error, model.ExchangeRatesResponse) {
	var exchangeRates model.ExchangeRatesResponse

	// baseURL := os.Getenv("EXTERNAL_API_URL")
	baseURL := "https://api.vatcomply.com/rates"
	fmt.Println(baseURL)
	params := url.Values{}
	params.Set("base", currencyCode)

	// Создаем URL с параметрами запроса
	reqURL, err := url.Parse(baseURL)
	if err != nil {
		return err, model.ExchangeRatesResponse{}
	}
	reqURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return err, model.ExchangeRatesResponse{}
	}
	fmt.Println(reqURL)
	// Выполняем запрос с помощью клиента
	resp, err := fa.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err, model.ExchangeRatesResponse{}
	}

	err = json.NewDecoder(resp.Body).Decode(&exchangeRates)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return err, model.ExchangeRatesResponse{}
	}

	defer resp.Body.Close()
	return nil, exchangeRates
}
