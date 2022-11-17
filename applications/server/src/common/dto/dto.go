package dto

import "lightup/src/common/db"

type ResourceDto struct {
	ID        string `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func GetResourceFromEntity(entity *db.BaseEntity) *ResourceDto {
	return &ResourceDto{
		ID:        entity.ID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
