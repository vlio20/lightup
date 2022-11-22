package api

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/http"
	"lightup/src/modules/feature_flag/bl"
	"lightup/src/modules/feature_flag/dto"
	"lightup/src/modules/feature_flag/model"
)

type FeatureFlagApi struct {
	featureFlagBl bl.FeatureFlagBl
}

func New() *FeatureFlagApi {
	return &FeatureFlagApi{
		featureFlagBl: bl.New(),
	}
}

func (api *FeatureFlagApi) GetFeatureFlagById(id string) (*dto.FeatureFlagDto, error) {
	entity, err := api.featureFlagBl.GetFeatureFlagById(id)

	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, &http.HttpError{StatusCode: 404, Message: "Feature flag not found"}
	}

	return dto.CreateFromEntity(entity), nil
}

func (api *FeatureFlagApi) CreateFeatureFlag(createDto *dto.CreateFeatureFlagDto) (*app_dto.CreatedEntityDto, error) {
	input := model.CreateFeatureFlagDto{
		Name:        createDto.Name,
		Description: createDto.Description,
		Archived:    false,
		Config:      createDto.Config,
	}
	entity, err := api.featureFlagBl.CreateFeatureFlag(&input)

	if err != nil {
		return nil, err
	}

	return &app_dto.CreatedEntityDto{
		ID: entity.ID.Hex(),
	}, nil
}
