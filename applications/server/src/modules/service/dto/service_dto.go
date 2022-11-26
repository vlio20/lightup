package dto

import (
	"lightup/src/common/dto"
	"lightup/src/modules/service/dal"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateServiceDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ServiceDto struct {
	dto.BaseEntityDto
	AccountID   primitive.ObjectID `json:"accountId"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Archived    bool               `json:"archived"`
}

func CreateFromEntity(entity *dal.ServiceEntity) *ServiceDto {
	return &ServiceDto{
		BaseEntityDto: *dto.GetResourceFromEntity(&entity.BaseEntity),
		AccountID:     entity.AccountID,
		Name:          entity.Name,
		Description:   entity.Description,
		Archived:      entity.Archived,
	}
}
