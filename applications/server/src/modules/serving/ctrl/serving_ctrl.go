package ctrl

import (
	"lightup/src/common/log"
	app_model "lightup/src/common/model"
	"lightup/src/common/router"
	guard "lightup/src/global/auth"
	"lightup/src/modules/serving/api"
	"lightup/src/modules/serving/dto"
	"lightup/src/modules/serving/model"

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
	guards := []guard.Guard{guard.NewAuthGuard()}
	r.GET("/servings/featureFlags", router.HandleQueryBounding(ctrl.getFeatureFlagState, guards))
}

func (ctrl *ServingController) getFeatureFlagState(
	c *app_model.ReqContext,
	query *model.FlagStateParams,
) (*dto.FeatureFlagStateDto, error) {
	return ctrl.api.GetFeatureFlagState(c.AccountID, query, c.Request.URL.Query())
}
