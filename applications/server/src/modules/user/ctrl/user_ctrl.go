package ctrl

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/log"
	"lightup/src/common/router"
	"lightup/src/modules/user/api"
	"lightup/src/modules/user/dto"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	api    *api.UserApi
	logger log.Logger
}

func New() *UserController {
	return &UserController{
		api:    api.New(),
		logger: log.GetLogger("user_ctrl"),
	}
}

func (ctrl *UserController) Init(r *gin.RouterGroup) {
	r.GET("/users/:id", router.HandleRequest(ctrl.getById))
	r.POST("/users", router.HandleBounding(ctrl.create))
}

func (ctrl *UserController) create(c *router.ReqContext, createDto *dto.CreateUserDto) (*app_dto.CreatedEntityDto, error) {
	accountID := c.AccountID

	return ctrl.api.CreateUser(accountID, createDto)
}

func (ctrl *UserController) getById(c *router.ReqContext) (*dto.UserDto, error) {
	id, err := router.GetParamAsObjectID(c, "id")
	if err != nil {
		return nil, err
	}

	return ctrl.api.GetUserById(*id)
}
