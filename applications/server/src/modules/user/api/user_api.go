package api

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/http"
	"lightup/src/modules/user/bl"
	"lightup/src/modules/user/dto"
	"lightup/src/modules/user/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserApi struct {
	userBl bl.UserBl
	authBl bl.AuthBl
}

func New() *UserApi {
	return &UserApi{
		userBl: bl.New(),
		authBl: bl.NewAuth(),
	}
}

func (api *UserApi) GetUserById(id primitive.ObjectID) (*dto.UserDto, error) {
	entity, err := api.userBl.GetUserById(id)

	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, &http.Error{StatusCode: 404, Message: "Feature flag not found"}
	}

	return dto.CreateFromEntity(entity), nil
}

func (api *UserApi) CreateUser(createDto *dto.CreateUserDto) (*app_dto.CreatedEntityDto, error) {
	input := model.CreateUserModel{
		AccountID: createDto.AccountID,
		Name:      createDto.Name,
		Email:     createDto.Email,
		Password:  createDto.Password,
	}

	entity, err := api.userBl.CreateUser(&input)

	if err != nil {
		return nil, err
	}

	return &app_dto.CreatedEntityDto{
		ID: entity.ID,
	}, nil
}

func (api *UserApi) CreateToken(tokenDto *dto.CreateTokenDto) (*dto.CreatedTokenDto, error) {
	token, err := api.authBl.CreateToken(tokenDto.Email, tokenDto.Password)

	if err != nil {
		return nil, err
	}

	return &dto.CreatedTokenDto{
		Token: token,
	}, nil
}
