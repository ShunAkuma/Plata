package model

type ResultObject struct {
	Data map[string]string
}

type Response struct {
	Message   string
	ResultObj interface{}
}
