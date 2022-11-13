package feature_flag_ctrl

import (
	router_utils "lightup/src/utils/router"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {
	router.POST("/featureFlags", router_utils.HandleBounding(createFeatureFlag))
}

func createFeatureFlag(c *gin.Context, dto CreateFeatureFlagDto) {
	c.JSON(200, gin.H{
		"name":        dto.Name,
		"description": dto.Description,
	})
}
