package http

import (
	"fmt"
)

type Error struct {
	StatusCode    int    `json:"statusCode"`
	Message       string `json:"message"`
	OriginalError error  `json:"-"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d:%v: http error", e.StatusCode, e.Message)
}

func GetHttpServerError(originalError error) Error {
	if _, ok := originalError.(Error); ok {
		return originalError.(Error)
	}

	return Error{
		StatusCode:    500,
		Message:       "Internal Server Error",
		OriginalError: originalError,
	}
}
