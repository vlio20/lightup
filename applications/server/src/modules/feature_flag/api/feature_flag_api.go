package api

import (
	app_dto "lightup/src/common/dto"
	"lightup/src/common/http"
	"lightup/src/modules/feature_flag/bl"
	"lightup/src/modules/feature_flag/dto"
	"lightup/src/modules/feature_flag/model"
)

func GetFeatureFlagById(id string) (*dto.FeatureFlagDto, error) {
	entity, err := bl.GetFeatureFlagById(id)

	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, &http.HttpError{StatusCode: 404, Message: "Feature flag not found"}
	}

	return dto.CreateFromEntity(entity), nil
}

func CreateFeatureFlag(createDto *dto.CreateFeatureFlagDto) (*app_dto.CreatedEntityDto, error) {
	input := model.CreateFeatureFlagDto{
		Name:        createDto.Name,
		Description: createDto.Description,
	}
	entity, err := bl.CreateFeatureFlag(&input)

	if err != nil {
		return nil, http.GetHttpServerError(err)
	}

	return &app_dto.CreatedEntityDto{
		ID: string(entity.ID),
	}, nil
}
