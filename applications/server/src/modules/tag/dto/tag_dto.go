package dto

import (
	"lightup/src/common/dto"
	"lightup/src/modules/tag/dal"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTagDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type TagDto struct {
	dto.BaseEntityDto
	AccountID   primitive.ObjectID `json:"accountId"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Archived    bool               `json:"archived"`
}

func CreateFromEntity(entity *dal.TagEntity) *TagDto {
	return &TagDto{
		BaseEntityDto: *dto.GetResourceFromEntity(&entity.BaseEntity),
		AccountID:     entity.AccountID,
		Name:          entity.Name,
		Description:   entity.Description,
		Archived:      entity.Archived,
	}
}
