package ctrl

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/router"
	"lightup/src/modules/feature_flag/api"
	"lightup/src/modules/feature_flag/dto"

	"github.com/gin-gonic/gin"
)

type FeatureFlagController struct {
	api *api.FeatureFlagApi
}

func New() *FeatureFlagController {
	return &FeatureFlagController{
		api: api.New(),
	}
}

func (ctrl *FeatureFlagController) Init(r *gin.RouterGroup) {
	r.GET("/featureFlags/:id", router.HandleRequest(ctrl.getFeatureFlagById))
	r.POST("/featureFlags", router.HandleBounding(ctrl.createFeatureFlag))
}

func (ctrl *FeatureFlagController) createFeatureFlag(c *gin.Context, createDto *dto.CreateFeatureFlagDto) (*app_dto.CreatedEntityDto, error) {
	return ctrl.api.CreateFeatureFlag(createDto)
}

func (ctrl *FeatureFlagController) getFeatureFlagById(c *gin.Context) (*dto.FeatureFlagDto, error) {
	id := c.Param("id")
	return ctrl.api.GetFeatureFlagById(id)
}
