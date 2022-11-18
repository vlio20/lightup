package dto

import "lightup/src/common/db"

type CreatedEntityDto struct {
	ID string `json:"id"`
}

type BaseEntityDto struct {
	ID        string `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func GetResourceFromEntity(entity *db.BaseEntity) *BaseEntityDto {
	return &BaseEntityDto{
		ID:        entity.ID.ToString(),
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
