package dto

import (
	"lightup/src/common/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreatedEntityDto struct {
	ID primitive.ObjectID `json:"id"`
}

type BaseEntityDto struct {
	ID        primitive.ObjectID `json:"id"`
	CreatedAt int64              `json:"createdAt"`
	UpdatedAt int64              `json:"updatedAt"`
}

func GetResourceFromEntity(entity *db.BaseEntity) *BaseEntityDto {
	return &BaseEntityDto{
		ID:        entity.ID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
