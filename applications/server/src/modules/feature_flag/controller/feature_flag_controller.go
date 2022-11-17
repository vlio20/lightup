package controller

import (
	"lightup/src/modules/feature_flag/api"
	"lightup/src/modules/feature_flag/dto"
	router_utils "lightup/src/utils/router"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {
	router.GET("/featureFlags/:id", getFeatureById)
	router.POST("/featureFlags", router_utils.HandleBounding(createFeatureFlag))
}

func createFeatureFlag(c *gin.Context, dto dto.CreateFeatureFlagDto) {
	c.JSON(200, gin.H{
		"name":        dto.Name,
		"description": dto.Description,
	})
}

func getFeatureById(c *gin.Context) {
	id := c.Param("name")
	dto, err := api.GetFeatureFlagById(id)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(200, dto)
}
