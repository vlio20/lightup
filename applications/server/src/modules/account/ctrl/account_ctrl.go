package ctrl

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/log"
	"lightup/src/common/router"
	"lightup/src/modules/service/api"
	"lightup/src/modules/service/dto"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceController struct {
	api    *api.ServiceApi
	logger log.Logger
}

func New() *ServiceController {
	return &ServiceController{
		api:    api.New(),
		logger: log.GetLogger("service_ctrl"),
	}
}

func (ctrl *ServiceController) Init(r *gin.RouterGroup) {
	r.GET("/services/:id", router.HandleRequest(ctrl.getById))
	r.POST("/services", router.HandleBounding(ctrl.create))
}

func (ctrl *ServiceController) create(c *gin.Context, createDto *dto.CreateServiceDto) (*app_dto.CreatedEntityDto, error) {
	accountID := primitive.NewObjectID()

	return ctrl.api.CreateService(accountID, createDto)
}

func (ctrl *ServiceController) getById(c *gin.Context) (*dto.ServiceDto, error) {
	id, err := router.GetParamAsObjectID(c, "id")
	if err != nil {
		return nil, err
	}

	return ctrl.api.GetServiceById(*id)
}
