package dto

import (
	"lightup/src/common/dto"
	"lightup/src/modules/feature_flag/dal"
)

type CreateFeatureFlagDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type FeatureFlagDto struct {
	dto.BaseEntityDto
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateFromEntity(entity *dal.FeatureFlagEntity) *FeatureFlagDto {
	return &FeatureFlagDto{
		Name:          entity.Name,
		Description:   entity.Description,
		BaseEntityDto: *dto.GetResourceFromEntity(&entity.BaseEntity),
	}
}
