package http

import (
	"fmt"
)

type HttpError struct {
	StatusCode    int    `json:"statusCode"`
	Message       string `json:"message"`
	OriginalError error  `json:"-"`
}

func (e HttpError) Error() string {
	return fmt.Sprintf("%d:%v: http error", e.StatusCode, e.Message)
}

func GetHttpServerError(originalError error) *HttpError {
	if _, ok := originalError.(*HttpError); ok {
		return originalError.(*HttpError)
	}

	return &HttpError{
		StatusCode:    500,
		Message:       "Internal Server Error",
		OriginalError: originalError,
	}
}
