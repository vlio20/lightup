package router

import (
	"lightup/src/common/http"
	"lightup/src/common/log"
	app_model "lightup/src/common/model"
	guard "lightup/src/global/guard"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func handleReturn[T interface{}](logger log.Logger, c *gin.Context, dto *T, err error) {
	var httpError http.Error

	if err != nil {
		if _, ok := err.(http.Error); ok {
			httpError = err.(http.Error)
		} else {
			httpError = http.GetHttpServerError(err)
			logger.Error(httpError.Message, ". error: ", err)
		}

		c.JSON(httpError.StatusCode, httpError)

		return
	}

	c.JSON(200, dto)
}

func HandleRequest[T interface{}](inv func(ctx *app_model.ReqContext) (T, error)) func(*gin.Context) {
	logger := log.GetLogger("router")

	return func(c *gin.Context) {
		appContext := getRequestContext(c)
		dto, err := inv(appContext)
		handleReturn(logger, c, &dto, err)
	}
}

func HandleBodyBounding[T interface{}, R interface{}](
	inv func(*app_model.ReqContext, *T) (*R, error),
	guards []guard.Guard,
) func(*gin.Context) {
	logger := log.GetLogger("router")

	return func(c *gin.Context) {
		var dto T
		appContext := getRequestContext(c)
		err := validateGuards(appContext, guards)

		if err != nil {
			handleReturn[R](logger, c, nil, err)
			return
		}

		if err := c.ShouldBind(&dto); err != nil {
			handleReturn[R](logger, c, nil, &http.Error{
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

func HandleQueryBounding[T interface{}, R interface{}](
	inv func(*app_model.ReqContext, *T) (*R, error),
	guards []guard.Guard,
) func(*gin.Context) {
	logger := log.GetLogger("router")

	return func(c *gin.Context) {
		var queryDto T
		appContext := getRequestContext(c)
		err := validateGuards(appContext, guards)

		if err != nil {
			handleReturn[R](logger, c, nil, err)
			return
		}

		if err := c.BindQuery(&queryDto); err != nil {
			handleReturn[R](logger, c, nil, http.Error{
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

func GetParamAsObjectID(c *app_model.ReqContext, key string) (*primitive.ObjectID, error) {
	id := c.Param(key)

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, http.Error{
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

func getRequestContext(c *gin.Context) *app_model.ReqContext {
	return &app_model.ReqContext{
		Context:   c,
		AccountID: primitive.NilObjectID,
	}
}

func validateGuards(c *app_model.ReqContext, guards []guard.Guard) error {
	if guards == nil {
		return nil
	}

	for _, g := range guards {
		if err := g.IsActive(c); err != nil {
			return err
		}
	}

	return nil
}
