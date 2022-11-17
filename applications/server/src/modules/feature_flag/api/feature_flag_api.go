package api

import (
	"lightup/src/common/http"
	"lightup/src/modules/feature_flag/bl"
	"lightup/src/modules/feature_flag/dto"
)

func GetFeatureFlagById(id string) (*dto.FeatureFlagDto, *http.HttpError) {
	entity, err := bl.GetFeatureFlagById(id)

	if err != nil {
		return nil, http.GetHttpServerError(err)
	}

	if entity == nil {
		return nil, &http.HttpError{StatusCode: 404, Message: "Feature flag not found"}
	}

	return dto.CreateFromEntity(entity), nil
}
