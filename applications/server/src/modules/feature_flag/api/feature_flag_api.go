package api

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/http"
	"lightup/src/modules/feature_flag/bl"
	"lightup/src/modules/feature_flag/dto"
	"lightup/src/modules/feature_flag/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureFlagApi struct {
	featureFlagBl bl.FeatureFlagBl
}

func New() *FeatureFlagApi {
	return &FeatureFlagApi{
		featureFlagBl: bl.New(),
	}
}

func (api *FeatureFlagApi) GetFeatureFlagById(id primitive.ObjectID) (*dto.FeatureFlagDto, error) {
	entity, err := api.featureFlagBl.GetFeatureFlagById(id)

	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, &http.HttpError{StatusCode: 404, Message: "Feature flag not found"}
	}

	return dto.CreateFromEntity(entity), nil
}

func (api *FeatureFlagApi) CreateFeatureFlag(accountID primitive.ObjectID, createDto *dto.CreateFeatureFlagDto) (*app_dto.CreatedEntityDto, error) {
	exisistingFeatureFlag, err := api.featureFlagBl.GetFeatureFlag(accountID, createDto.ServiceID, createDto.Name)

	if err != nil {
		return nil, err
	}

	if exisistingFeatureFlag != nil {
		return nil, &http.HttpError{StatusCode: 409, Message: "Feature flag already exists"}
	}

	input := model.CreateFeatureFlagInput{
		AccountID:   accountID,
		Name:        createDto.Name,
		Description: createDto.Description,
		ServiceID:   createDto.ServiceID,
		Archived:    false,
		Config:      createDto.Config,
	}

	entity, err := api.featureFlagBl.CreateFeatureFlag(&input)

	if err != nil {
		return nil, err
	}

	return &app_dto.CreatedEntityDto{
		ID: entity.ID,
	}, nil
}
