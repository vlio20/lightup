package router

import (
	"lightup/src/common/http"
	"lightup/src/common/log"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReqContext struct {
	*gin.Context
	AccountID primitive.ObjectID
}

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

func HandleRequest[T interface{}](inv func(*ReqContext) (T, error)) func(*gin.Context) {
	logger := log.GetLogger("router")

	return func(c *gin.Context) {
		appContext := getRequestContext(c)
		dto, err := inv(appContext)
		handleReturn(logger, c, &dto, err)
	}
}

func HandleBodyBounding[T interface{}, R interface{}](inv func(*ReqContext, *T) (*R, error)) func(*gin.Context) {
	logger := log.GetLogger("router")

	return func(c *gin.Context) {
		var dto T
		appContext := getRequestContext(c)

		if err := c.ShouldBind(&dto); err != nil {
			handleReturn[R](logger, c, nil, &http.HttpError{
				StatusCode:    400,
				Message:       extractValidationError(err.Error()),
				OriginalError: err,
			})

			return
		}

		resultDto, err := inv(appContext, &dto)
		handleReturn(logger, c, &resultDto, err)
	}
}

func HandleQueryBounding[T interface{}, R interface{}](inv func(*ReqContext, *T) (*R, error)) func(*gin.Context) {
	logger := log.GetLogger("router")

	return func(c *gin.Context) {
		var queryDto T
		appContext := getRequestContext(c)

		if err := c.BindQuery(&queryDto); err != nil {
			handleReturn[R](logger, c, nil, &http.HttpError{
				StatusCode:    400,
				Message:       extractValidationError(err.Error()),
				OriginalError: err,
			})

			return
		}

		resultDto, err := inv(appContext, &queryDto)
		handleReturn(logger, c, &resultDto, err)
	}
}

func GetParamAsObjectID(c *ReqContext, key string) (*primitive.ObjectID, *http.HttpError) {
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

func getRequestContext(c *gin.Context) *ReqContext {
	return &ReqContext{
		Context:   c,
		AccountID: primitive.NewObjectID(),
	}
}
