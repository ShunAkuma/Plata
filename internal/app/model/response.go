package model

// @Summary Server response model
type Response struct {
	// Message from server
	// example: "Something went wrong"
	Message string

	// Response model that can store an error or response
	ResultObj interface{}
}
