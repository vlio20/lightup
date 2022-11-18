package global

import (
	"lightup/src/modules/feature_flag/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	route := gin.Default()

	v1 := route.Group("api/v1")
	{
		controller.Init(v1)
	}

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.Run(":4321")
}
