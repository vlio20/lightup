package api

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/http"
	"lightup/src/modules/account/bl"
	"lightup/src/modules/account/dto"
	"lightup/src/modules/account/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountApi struct {
	accountBl bl.AccountBl
}

func New() *AccountApi {
	return &AccountApi{
		accountBl: bl.New(),
	}
}

func (api *AccountApi) GetAccountById(id primitive.ObjectID) (*dto.AccountDto, error) {
	entity, err := api.accountBl.GetAccountById(id)

	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, &http.HttpError{StatusCode: 404, Message: "Feature flag not found"}
	}

	return dto.CreateFromEntity(entity), nil
}

func (api *AccountApi) CreateAccount(accountID primitive.ObjectID, createDto *dto.CreateAccountDto) (*app_dto.CreatedEntityDto, error) {
	input := model.CreateAccountModel{
		Name:        createDto.Name,
		Description: createDto.Description,
	}

	entity, err := api.accountBl.CreateAccount(&input)

	if err != nil {
		return nil, err
	}

	return &app_dto.CreatedEntityDto{
		ID: entity.ID,
	}, nil
}
