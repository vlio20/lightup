package dto

import (
	"lightup/src/common/dto"
	"lightup/src/modules/account/dal"
)

type CreateAccountDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type AccountDto struct {
	dto.BaseEntityDto
	Name        string `json:"name"`
	Description string `json:"description"`
	Archived    bool   `json:"archived"`
}

func CreateFromEntity(entity *dal.AccountEntity) *AccountDto {
	return &AccountDto{
		BaseEntityDto: *dto.GetResourceFromEntity(&entity.BaseEntity),
		Name:          entity.Name,
		Description:   entity.Description,
		Archived:      entity.Archived,
	}
}
