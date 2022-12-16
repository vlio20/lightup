package ctrl

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/log"
	app_model "lightup/src/common/model"
	"lightup/src/common/router"
	"lightup/src/modules/account/api"
	"lightup/src/modules/account/dto"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountController struct {
	api    *api.AccountApi
	logger log.Logger
}

func New() *AccountController {
	return &AccountController{
		api:    api.New(),
		logger: log.GetLogger("account_ctrl"),
	}
}

func (ctrl *AccountController) Init(r *gin.RouterGroup) {
	r.GET("/accounts/:id", router.HandleRequest(ctrl.getById))
	r.POST("/accounts", router.HandleBodyBounding(ctrl.create))
}

func (ctrl *AccountController) create(c *app_model.ReqContext, createDto *dto.CreateAccountDto) (*app_dto.CreatedEntityDto, error) {
	accountID := primitive.NewObjectID()

	return ctrl.api.CreateAccount(accountID, createDto)
}

func (ctrl *AccountController) getById(c *app_model.ReqContext) (*dto.AccountDto, error) {
	id, err := router.GetParamAsObjectID(c, "id")
	if err != nil {
		return nil, err
	}

	return ctrl.api.GetAccountById(*id)
}
