package controller

import (
	"lightup/src/common/router"
	"lightup/src/modules/feature_flag/api"
	"lightup/src/modules/feature_flag/dto"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup) {
	r.GET("/featureFlags/:id", router.HandleRequest(getFeatureById))
	// router.POST("/featureFlags", router.HandleBounding(createFeatureFlag))
}

// func createFeatureFlag(c *gin.Context, dto dto.CreateFeatureFlagDto) {
// 	dto = c.JSON(200, gin.H{
// 		"name":        dto.Name,
// 		"description": dto.Description,
// 	})
// }

func getFeatureById(c *gin.Context) (*dto.FeatureFlagDto, error) {
	id := c.Param("name")
	return api.GetFeatureFlagById(id)

}
