package dto

import (
	"lightup/src/common/dto"
	app_model "lightup/src/common/model"
	"lightup/src/modules/feature_flag/dal"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateFeatureFlagDto struct {
	Name        string                      `json:"name" binding:"required"`
	Description string                      `json:"description" binding:"required"`
	ServiceID   primitive.ObjectID          `json:"serviceId"`
	Config      app_model.FeatureFlagConfig `json:"config" binding:"required"`
}

type FeatureFlagDto struct {
	dto.BaseEntityDto
	AccountID   primitive.ObjectID          `json:"accountId"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	ServiceID   primitive.ObjectID          `json:"serviceId"`
	Archived    bool                        `json:"archived"`
	Config      app_model.FeatureFlagConfig `json:"config"`
}

func CreateFromEntity(entity *dal.FeatureFlagEntity) *FeatureFlagDto {
	return &FeatureFlagDto{
		BaseEntityDto: *dto.GetResourceFromEntity(&entity.BaseEntity),
		AccountID:     entity.AccountID,
		Name:          entity.Name,
		Description:   entity.Description,
		Config:        entity.Config,
		ServiceID:     entity.ServiceID,
		Archived:      entity.Archived,
	}
}
