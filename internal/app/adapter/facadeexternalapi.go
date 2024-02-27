package adapter

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"ratequotes/internal/app/model"
)

type ExternalApiRepository interface {
	GetCurrencyRate(currencyCode string) (error, model.ExchangeRatesResponse)
}

type FacadeApiRepository struct {
	httpClient *http.Client
}

func NewFacadeApi(httpClient *http.Client) ExternalApiRepository {
	return &FacadeApiRepository{
		httpClient: httpClient,
	}
}

func (fa FacadeApiRepository) GetCurrencyRate(currencyCode string) (error, model.ExchangeRatesResponse) {
	var exchangeRates model.ExchangeRatesResponse

	baseURL := os.Getenv("EXTERNAL_API_URL")
	params := url.Values{}
	params.Set("base", currencyCode)

	reqURL, err := url.Parse(baseURL)
	if err != nil {
		log.Println(err)
		return err, model.ExchangeRatesResponse{}
	}
	reqURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		log.Println(err)
		return err, model.ExchangeRatesResponse{}
	}
	resp, err := fa.httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return err, model.ExchangeRatesResponse{}
	}

	err = json.NewDecoder(resp.Body).Decode(&exchangeRates)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return err, model.ExchangeRatesResponse{}
	}

	defer resp.Body.Close()
	return nil, exchangeRates
}
