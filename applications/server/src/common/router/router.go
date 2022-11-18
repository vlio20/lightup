package router

import (
	"lightup/src/common/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func handleReturn[T interface{}](c *gin.Context, dto *T, err error) {
	var httpError *http.HttpError

	if err != nil {
		if _, ok := err.(*http.HttpError); ok {
			httpError = err.(*http.HttpError)
		} else {
			httpError = http.GetHttpServerError(err)
		}

		c.JSON(httpError.StatusCode, httpError)
		return
	}

	c.JSON(200, dto)
}

func HandleRequest[T interface{}](inv func(*gin.Context) (T, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		dto, err := inv(c)
		handleReturn(c, &dto, err)
	}
}

func HandleBounding[T interface{}, R interface{}](inv func(*gin.Context, *T) (*R, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		var dto T

		if err := c.ShouldBind(&dto); err != nil {
			handleReturn[R](c, nil, &http.HttpError{
				StatusCode:    400,
				Message:       extractValidationError(err.Error()),
				OriginalError: err,
			})

			return
		}

		resultDto, err := inv(c, &dto)
		handleReturn(c, &resultDto, err)
	}
}

func extractValidationError(msg string) string {
	parts := strings.Split(msg, "Error:")

	return parts[len(parts)-1]
}
