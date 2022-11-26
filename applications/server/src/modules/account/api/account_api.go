package api

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/http"
	"lightup/src/modules/service/bl"
	"lightup/src/modules/service/dto"
	"lightup/src/modules/service/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceApi struct {
	serviceBl bl.ServiceBl
}

func New() *ServiceApi {
	return &ServiceApi{
		serviceBl: bl.New(),
	}
}

func (api *ServiceApi) GetServiceById(id primitive.ObjectID) (*dto.ServiceDto, error) {
	entity, err := api.serviceBl.GetServiceById(id)

	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, &http.HttpError{StatusCode: 404, Message: "Feature flag not found"}
	}

	return dto.CreateFromEntity(entity), nil
}

func (api *ServiceApi) CreateService(accountID primitive.ObjectID, createDto *dto.CreateServiceDto) (*app_dto.CreatedEntityDto, error) {
	exisistingService, err := api.serviceBl.GetService(accountID, createDto.Name)

	if err != nil {
		return nil, err
	}

	if exisistingService != nil {
		return nil, &http.HttpError{StatusCode: 409, Message: "Feature flag already exists"}
	}

	input := model.CreateServiceModel{
		AccountID:   accountID,
		Name:        createDto.Name,
		Description: createDto.Description,
	}

	entity, err := api.serviceBl.CreateService(&input)

	if err != nil {
		return nil, err
	}

	return &app_dto.CreatedEntityDto{
		ID: entity.ID,
	}, nil
}
