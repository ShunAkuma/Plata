package model

import (
	"time"
)

type Quotes struct {
	Id         string
	TimeUpdate time.Time
	Base       string
	Rates      map[string]float64
}

type ResponseQuotesModel struct {
	Base       string
	TimeUpdate time.Time
	Rates      map[string]float64
}

type QuotesRequestModel struct {
	CurrencySymbol string `json:"currencycode"`
}
