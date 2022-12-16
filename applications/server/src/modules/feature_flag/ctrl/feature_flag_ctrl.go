package ctrl

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/log"
	app_model "lightup/src/common/model"
	"lightup/src/common/router"
	"lightup/src/modules/feature_flag/api"
	"lightup/src/modules/feature_flag/dto"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	r.GET("/featureFlags/:id", router.HandleRequest(ctrl.getFeatureFlagById))
	r.POST("/featureFlags", router.HandleBodyBounding(ctrl.createFeatureFlag))
}

func (ctrl *FeatureFlagController) createFeatureFlag(c *app_model.ReqContext, createDto *dto.CreateFeatureFlagDto) (*app_dto.CreatedEntityDto, error) {
	return ctrl.api.CreateFeatureFlag(primitive.NewObjectID(), createDto)
}

func (ctrl *FeatureFlagController) getFeatureFlagById(c *app_model.ReqContext) (*dto.FeatureFlagDto, error) {
	id, err := router.GetParamAsObjectID(c, "id")

	if err != nil {
		return nil, err
	}

	return ctrl.api.GetFeatureFlagById(*id)
}
