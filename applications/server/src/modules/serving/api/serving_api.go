package api

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lightup/src/modules/serving/bl"
	"lightup/src/modules/serving/dto"
	"lightup/src/modules/serving/model"
	"net/url"
)

type ServingApi struct {
	servingBl bl.ServingBl
}

func New() *ServingApi {
	return &ServingApi{
		servingBl: bl.New(),
	}
}

func (api *ServingApi) GetFeatureFlagState(
	accountID primitive.ObjectID, flagParams *model.FlagStateParams,
	query url.Values,
) (*dto.FeatureFlagStateDto, error) {
	state, err := api.servingBl.GetFeatureFlagState(accountID, flagParams, query)

	if err != nil {
		return nil, err
	}

	return &dto.FeatureFlagStateDto{
		Active: state,
	}, nil
}
