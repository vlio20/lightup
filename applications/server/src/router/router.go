package router

import (
	feature_flag_ctrl "lightup/src/modules/feature_flag/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	route := gin.Default()

	v1 := route.Group("api/v1")
	{
		feature_flag_ctrl.Init(v1)
	}

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.Run(":4321")
}
