package ctrl

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/log"
	app_model "lightup/src/common/model"
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
	r.POST("/users", router.HandleBodyBounding(ctrl.create))
	r.POST("/users/tokens", router.HandleBodyBounding(ctrl.createToken))
}

func (ctrl *UserController) create(c *app_model.ReqContext, createDto *dto.CreateUserDto) (*app_dto.CreatedEntityDto, error) {
	return ctrl.api.CreateUser(createDto)
}

func (ctrl *UserController) getById(c *app_model.ReqContext) (*dto.UserDto, error) {
	id, err := router.GetParamAsObjectID(c, "id")
	if err != nil {
		return nil, err
	}

	return ctrl.api.GetUserById(*id)
}

func (ctrl *UserController) createToken(context *app_model.ReqContext, t *dto.CreateTokenDto) (*dto.CreatedTokenDto, error) {
	return ctrl.api.CreateToken(t)
}
