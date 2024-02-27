package model

// @Summary Server response model
// @Description Response model from quotes server
type Response struct {
	// Message from server
	// example: "Something went wrong"
	Message string

	// Generalized Response Object from Server
	// Response model that can store an error or response
	ResultObj interface{}
}
