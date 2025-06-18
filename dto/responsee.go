package dto

type ErrorResponse struct {
	Message string `json:"message"`
	Err     string `json:"error"`
}

type SuccessResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}
