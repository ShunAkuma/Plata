package adapter

import "net/http"

type ExternalApiRepository interface {
}

type FacadeApiRepository struct {
	httpClient *http.Client
}

func NewFacadeApi(httpClient *http.Client) ExternalApiRepository {
	return &FacadeApiRepository{
		httpClient: httpClient,
	}
}
