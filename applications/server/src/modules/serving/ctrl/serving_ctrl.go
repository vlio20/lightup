package ctrl

import (
	"lightup/src/common/log"
	"lightup/src/common/router"
	"lightup/src/modules/serving/api"
	"lightup/src/modules/serving/dto"

	"github.com/gin-gonic/gin"
)

type ServingController struct {
	api    *api.ServingApi
	logger log.Logger
}

func New() *ServingController {
	return &ServingController{
		api:    api.New(),
		logger: log.GetLogger("serving_ctrl"),
	}
}

func (ctrl *ServingController) Init(r *gin.RouterGroup) {
	r.GET("/servings/featureFlags", router.HandleQueryBounding(ctrl.getFeatureFlagState))
}

func (ctrl *ServingController) getFeatureFlagState(c *router.ReqContext, query *dto.GetFeatureFlagStateParams) (*dto.FeatureFlagStateDto, error) {
	return nil, nil
}
