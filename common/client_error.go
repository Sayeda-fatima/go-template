package common

import "net/http"

type (
	HttpError interface {
		Error() string
		ErrorStatusCode() int
	}
	httpError struct {
		StatusCode int    `json:"status_code"`
		Key        string `json:"key"`
		Message    string `json:"message"`
		Err        string `json:"error"`
	}
)

func NewHTTPError(statusCode int, msg string, err string) *httpError {
	return &httpError{
		StatusCode: statusCode,
		Key:        http.StatusText(statusCode),
		Message:    msg,
		Err:        err,
	}
}

func (e *httpError) Error() string {
	return e.Err
}

func (e *httpError) ErrorStatusCode() int {
	return e.StatusCode
}
