package ctrl

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/log"
	app_model "lightup/src/common/model"
	"lightup/src/common/router"
	"lightup/src/global/guard"
	"lightup/src/modules/tag/api"
	"lightup/src/modules/tag/dto"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TagController struct {
	api    *api.TagApi
	logger log.Logger
}

func New() *TagController {
	return &TagController{
		api:    api.New(),
		logger: log.GetLogger("tag_ctrl"),
	}
}

func (ctrl *TagController) Init(r *gin.RouterGroup) {
	guards := []guard.Guard{guard.NewAuthGuard()}
	r.GET("/tags/:id", router.HandleRequest(ctrl.getById))
	r.POST("/tags", router.HandleBodyBounding(ctrl.create, guards))
}

func (ctrl *TagController) create(_ *app_model.ReqContext, createDto *dto.CreateTagDto) (*app_dto.CreatedEntityDto, error) {
	accountID := primitive.NewObjectID()

	return ctrl.api.CreateTag(accountID, createDto)
}

func (ctrl *TagController) getById(c *app_model.ReqContext) (*dto.TagDto, error) {
	id, err := router.GetParamAsObjectID(c, "id")
	if err != nil {
		return nil, err
	}

	return ctrl.api.GetTagById(*id)
}
