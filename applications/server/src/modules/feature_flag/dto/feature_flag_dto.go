package dto

import (
	"lightup/src/common/dto"
	app_model "lightup/src/common/model"
	"lightup/src/modules/feature_flag/ff_dal"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateFeatureFlagDto struct {
	Name        string                      `json:"name" binding:"required"`
	Description string                      `json:"description" binding:"required"`
	Tags        []primitive.ObjectID        `json:"tags"`
	Config      app_model.FeatureFlagConfig `json:"config" binding:"required"`
}

type FeatureFlagDto struct {
	dto.BaseEntityDto
	AccountID   primitive.ObjectID          `json:"accountId"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Tags        []primitive.ObjectID        `json:"tags"`
	Archived    bool                        `json:"archived"`
	Config      app_model.FeatureFlagConfig `json:"config"`
}

func CreateFromEntity(entity *ff_dal.FeatureFlagEntity) *FeatureFlagDto {
	return &FeatureFlagDto{
		BaseEntityDto: *dto.GetResourceFromEntity(&entity.BaseEntity),
		AccountID:     entity.AccountID,
		Name:          entity.Name,
		Description:   entity.Description,
		Config:        entity.Config,
		Tags:          entity.Tags,
		Archived:      entity.Archived,
	}
}
