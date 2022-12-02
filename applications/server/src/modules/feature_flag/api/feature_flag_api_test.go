package api

import (
	"lightup/src/common/db"
	"lightup/src/common/http"
	"lightup/src/common/matching_rule"
	"lightup/src/modules/feature_flag/dal"
	ff_dto "lightup/src/modules/feature_flag/dto"
	"lightup/src/modules/feature_flag/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ffMock = genFeatureFlagEntity()

type FeatureFlagBlMock struct {
}

func (bl *FeatureFlagBlMock) GetFeatureFlagById(id primitive.ObjectID) (*dal.FeatureFlagEntity, error) {
	if id == ffMock.ID {
		return ffMock, nil
	}

	return nil, nil
}

func (bl *FeatureFlagBlMock) CreateFeatureFlag(input *model.CreateFeatureFlagInput) (*dal.FeatureFlagEntity, error) {
	return ffMock, nil
}

func (bl *FeatureFlagBlMock) GetFeatureFlag(accountId primitive.ObjectID, serviceId primitive.ObjectID, name string) (*dal.FeatureFlagEntity, error) {
	return ffMock, nil
}

func genFeatureFlagEntity() *dal.FeatureFlagEntity {
	return &dal.FeatureFlagEntity{
		BaseEntity: db.BaseEntity{
			ID: primitive.NewObjectID(),
		},
		ServiceID:   primitive.NewObjectID(),
		Name:        "name",
		Description: "description",
		Archived:    false,
		Config: model.FeatureFlagConfig{
			MatchingRules: []matching_rule.MatchingRule{},
		},
	}
}

func TestGetFeatureFlag_whenFound_returnDto(t *testing.T) {
	ffApi := &FeatureFlagApi{
		featureFlagBl: &FeatureFlagBlMock{},
	}

	res, err := ffApi.GetFeatureFlagById(ffMock.ID)

	assert.Nil(t, err)
	comperEntityAndDto(t, ffMock, res)
}

func TestGetFeatureFlag_whenNotFound_returnAnError(t *testing.T) {
	ffApi := &FeatureFlagApi{
		featureFlagBl: &FeatureFlagBlMock{},
	}

	res, err := ffApi.GetFeatureFlagById(primitive.NewObjectID())

	assert.Equal(t, err.(*http.HttpError).StatusCode, 404)
	assert.Nil(t, res)
}

func comperEntityAndDto(t *testing.T, entity *dal.FeatureFlagEntity, dto *ff_dto.FeatureFlagDto) {
	assert.Equal(t, entity.ID, dto.ID)
	assert.Equal(t, entity.ServiceID, dto.ServiceID)
	assert.Equal(t, entity.AccountID, dto.AccountID)
	assert.Equal(t, entity.Name, dto.Name)
	assert.Equal(t, entity.Description, dto.Description)
	assert.Equal(t, entity.Archived, dto.Archived)
	assert.Equal(t, entity.Config, dto.Config)
}
