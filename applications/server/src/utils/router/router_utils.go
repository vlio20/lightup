package router_utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleBounding[T interface{}](c *gin.Context) (T, error) {
	var dto T

	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(400, gin.H{
			"message": extractValidationError(err.Error()),
		})
		return dto, err
	}

	return dto, nil
}

func extractValidationError(msg string) string {
	parts := strings.Split(msg, "Error:")

	return parts[len(parts)-1]
}
