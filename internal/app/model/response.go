package model

type ResultObject struct {
	Data map[string]string
}

type Response struct {
	StatusCode int
	Message    string
	ResultObj  []ResultObject
}
