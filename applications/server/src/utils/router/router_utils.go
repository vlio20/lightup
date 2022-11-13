package router_utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleBounding[T interface{}](inv func(*gin.Context, T)) func(*gin.Context) {
	return func(c *gin.Context) {
		var dto T

		if err := c.ShouldBind(&dto); err != nil {
			c.JSON(400, gin.H{
				"message": extractValidationError(err.Error()),
			})
			return
		}

		inv(c, dto)
	}
}

func extractValidationError(msg string) string {
	parts := strings.Split(msg, "Error:")

	return parts[len(parts)-1]
}
