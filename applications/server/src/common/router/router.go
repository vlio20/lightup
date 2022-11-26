package router

import (
	"lightup/src/common/http"
	"lightup/src/common/log"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func handleReturn[T interface{}](logger log.Logger, c *gin.Context, dto *T, err error) {
	var httpError *http.HttpError

	if err != nil {
		if _, ok := err.(*http.HttpError); ok {
			httpError = err.(*http.HttpError)
		} else {
			httpError = http.GetHttpServerError(err)
		}

		c.JSON(httpError.StatusCode, httpError)
		logger.Error(httpError.Message, "error", err)

		return
	}

	c.JSON(200, dto)
}

func HandleRequest[T interface{}](inv func(*gin.Context) (T, error)) func(*gin.Context) {
	logger := log.GetLogger("router")

	return func(c *gin.Context) {
		dto, err := inv(c)
		handleReturn(logger, c, &dto, err)
	}
}

func HandleBounding[T interface{}, R interface{}](inv func(*gin.Context, *T) (*R, error)) func(*gin.Context) {
	logger := log.GetLogger("router")
	return func(c *gin.Context) {
		var dto T

		if err := c.ShouldBind(&dto); err != nil {
			handleReturn[R](logger, c, nil, &http.HttpError{
				StatusCode:    400,
				Message:       extractValidationError(err.Error()),
				OriginalError: err,
			})

			return
		}

		resultDto, err := inv(c, &dto)
		handleReturn(logger, c, &resultDto, err)
	}
}

func GetParamAsObjectID(c *gin.Context, key string) (*primitive.ObjectID, *http.HttpError) {
	id := c.Param(key)

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, &http.HttpError{
			StatusCode:    400,
			Message:       "Invalid ID: " + id,
			OriginalError: nil,
		}
	}

	return &objId, nil
}

func extractValidationError(msg string) string {
	parts := strings.Split(msg, "Error:")

	return parts[len(parts)-1]
}
