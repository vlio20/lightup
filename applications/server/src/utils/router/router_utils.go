package router_utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleBounding(c *gin.Context, obj any) any {
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"message": extractValidationError(err.Error()),
		})
		return err
	}

	return nil
}

func extractValidationError(msg string) string {
	parts := strings.Split(msg, "Error:")

	return parts[len(parts)-1]
}
