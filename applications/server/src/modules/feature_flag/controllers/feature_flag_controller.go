package feature_flag_ctrl

import (
	router_utils "lightup/src/utils/router"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {
	router.POST("/featureFlags", createFeatureFlag)
}

func createFeatureFlag(c *gin.Context) {
	feature_flag, err := router_utils.HandleBounding[CreateFeatureFlagDto](c)

	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"name":        feature_flag.Name,
		"description": feature_flag.Description,
	})
}
