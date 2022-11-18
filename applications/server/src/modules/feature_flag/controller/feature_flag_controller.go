package controller

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/router"
	"lightup/src/modules/feature_flag/api"
	"lightup/src/modules/feature_flag/dto"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup) {
	r.GET("/featureFlags/:id", router.HandleRequest(getFeatureFlagById))
	r.POST("/featureFlags", router.HandleBounding(createFeatureFlag))
}

func createFeatureFlag(c *gin.Context, createDto *dto.CreateFeatureFlagDto) (*app_dto.CreatedEntityDto, error) {
	return api.CreateFeatureFlag(createDto)
}

func getFeatureFlagById(c *gin.Context) (*dto.FeatureFlagDto, error) {
	id := c.Param("id")
	return api.GetFeatureFlagById(id)
}
