package dto

import (
	"lightup/src/common/dto"
	"lightup/src/modules/feature_flag/dal"
	"lightup/src/modules/feature_flag/model"
)

type CreateFeatureFlagDto struct {
	Name        string                  `json:"name" binding:"required"`
	Description string                  `json:"description" binding:"required"`
	ServiceID   string                  `json:"serviceId"`
	Config      model.FeatureFlagConfig `json:"config" binding:"required"`
}

type FeatureFlagDto struct {
	dto.BaseEntityDto
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	ServiceID   string                  `json:"serviceId"`
	Archived    bool                    `json:"archived"`
	Config      model.FeatureFlagConfig `json:"config"`
}

func CreateFromEntity(entity *dal.FeatureFlagEntity) *FeatureFlagDto {
	return &FeatureFlagDto{
		BaseEntityDto: *dto.GetResourceFromEntity(&entity.BaseEntity),
		Name:          entity.Name,
		Description:   entity.Description,
		Config:        entity.Config,
		ServiceID:     entity.ServiceID.Hex(),
		Archived:      entity.Archived,
	}
}
