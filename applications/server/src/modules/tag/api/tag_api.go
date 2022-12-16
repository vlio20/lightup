package api

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/http"
	"lightup/src/modules/tag/bl"
	"lightup/src/modules/tag/dto"
	"lightup/src/modules/tag/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TagApi struct {
	tagBl bl.TagBl
}

func New() *TagApi {
	return &TagApi{
		tagBl: bl.New(),
	}
}

func (api *TagApi) GetTagById(id primitive.ObjectID) (*dto.TagDto, error) {
	entity, err := api.tagBl.GetTagById(id)

	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, &http.Error{StatusCode: 404, Message: "Feature flag not found"}
	}

	return dto.CreateFromEntity(entity), nil
}

func (api *TagApi) CreateTag(accountID primitive.ObjectID, createDto *dto.CreateTagDto) (*app_dto.CreatedEntityDto, error) {
	exisistingTag, err := api.tagBl.GetTag(accountID, createDto.Name)

	if err != nil {
		return nil, err
	}

	if exisistingTag != nil {
		return nil, &http.Error{StatusCode: 409, Message: "Feature flag already exists"}
	}

	input := model.CreateTagModel{
		AccountID:   accountID,
		Name:        createDto.Name,
		Description: createDto.Description,
	}

	entity, err := api.tagBl.CreateTag(&input)

	if err != nil {
		return nil, err
	}

	return &app_dto.CreatedEntityDto{
		ID: entity.ID,
	}, nil
}
