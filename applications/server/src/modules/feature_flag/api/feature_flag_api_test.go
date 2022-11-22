package api

import (
	"errors"
	"lightup/src/common/db"
	"lightup/src/modules/feature_flag/dal"
	"lightup/src/modules/feature_flag/model"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ffMock = genFeatureFlagEntity()
var notFoundError = errors.New("not found")

type FeatureFlagBlMock struct {
}

func (bl *FeatureFlagBlMock) GetFeatureFlagById(id string) (*dal.FeatureFlagEntity, error) {
	if id == ffMock.ID.Hex() {
		return ffMock, nil
	}

	return nil, notFoundError
}

func (bl *FeatureFlagBlMock) CreateFeatureFlag(input *model.CreateFeatureFlagDto) (*dal.FeatureFlagEntity, error) {
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
			Identifier: "id",
			Percentage: 0,
		},
	}
}

func TestGetFeatureFlag_whenFound_returnDto(t *testing.T) {
	ffApi := &FeatureFlagApi{
		featureFlagBl: &FeatureFlagBlMock{},
	}

	res, err := ffApi.GetFeatureFlagById(ffMock.ID.Hex())

	assert.Nil(t, err)
	assert.Equal(t, res.ID, ffMock.ID.Hex())
}

func TestGetFeatureFlag_whenNotFound_returnAnError(t *testing.T) {
	ffApi := &FeatureFlagApi{
		featureFlagBl: &FeatureFlagBlMock{},
	}

	res, err := ffApi.GetFeatureFlagById("123")

	assert.Equal(t, err, notFoundError)
	assert.Nil(t, res)
}
