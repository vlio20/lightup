package dto

import (
	"lightup/src/common/dto"
	"lightup/src/modules/user/dal"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserDto struct {
	AccountID primitive.ObjectID `json:"accountId" binding:"required"`
	Name      string             `json:"name" binding:"required"`
	Email     string             `json:"email" binding:"required,email"`
	Password  string             `json:"password" binding:"required,min=8"`
}

type UserDto struct {
	dto.BaseEntityDto
	AccountID primitive.ObjectID `json:"accountId"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	Archived  bool               `json:"archived"`
}

func CreateFromEntity(entity *dal.UserEntity) *UserDto {
	return &UserDto{
		BaseEntityDto: *dto.GetResourceFromEntity(&entity.BaseEntity),
		Name:          entity.Name,
		AccountID:     entity.AccountID,
		Archived:      entity.Archived,
	}
}

type CreateTokenDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreatedTokenDto struct {
	Token string `json:"token"`
}
