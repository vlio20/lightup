package ctrl

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/log"
	app_model "lightup/src/common/model"
	"lightup/src/common/router"
	"lightup/src/global/guard"
	"lightup/src/modules/feature_flag/api"
	"lightup/src/modules/feature_flag/dto"

	"github.com/gin-gonic/gin"
)

type FeatureFlagController struct {
	api    *api.FeatureFlagApi
	logger log.Logger
}

func New() *FeatureFlagController {
	return &FeatureFlagController{
		api:    api.New(),
		logger: log.GetLogger("feature_flag_ctrl"),
	}
}

func (ctrl *FeatureFlagController) Init(r *gin.RouterGroup) {
	guards := []guard.Guard{guard.NewAuthGuard()}
	r.GET("/featureFlags/:id", router.HandleRequest(ctrl.getFeatureFlagById))
	r.POST("/featureFlags", router.HandleBodyBounding(ctrl.createFeatureFlag, guards))
}

func (ctrl *FeatureFlagController) createFeatureFlag(
	c *app_model.ReqContext,
	createDto *dto.CreateFeatureFlagDto,
) (*app_dto.CreatedEntityDto, error) {
	return ctrl.api.CreateFeatureFlag(c.AccountID, createDto)
}

func (ctrl *FeatureFlagController) getFeatureFlagById(c *app_model.ReqContext) (*dto.FeatureFlagDto, error) {
	id, err := router.GetParamAsObjectID(c, "id")

	if err != nil {
		return nil, err
	}

	return ctrl.api.GetFeatureFlagById(*id)
}
