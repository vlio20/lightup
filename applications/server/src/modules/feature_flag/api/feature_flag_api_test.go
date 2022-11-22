package api

import (
	"errors"
	"lightup/src/common/db"
	"lightup/src/modules/feature_flag/dal"
	ff_dto "lightup/src/modules/feature_flag/dto"
	"lightup/src/modules/feature_flag/model"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ffMock = genFeatureFlagEntity()
var errMock = errors.New("not found")

type FeatureFlagBlMock struct {
}

func (bl *FeatureFlagBlMock) GetFeatureFlagById(id string) (*dal.FeatureFlagEntity, error) {
	if id == ffMock.ID.Hex() {
		return ffMock, nil
	}

	return nil, errMock
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
	comperEntityAndDto(t, ffMock, res)
}

func TestGetFeatureFlag_whenNotFound_returnAnError(t *testing.T) {
	ffApi := &FeatureFlagApi{
		featureFlagBl: &FeatureFlagBlMock{},
	}

	res, err := ffApi.GetFeatureFlagById("123")

	assert.Equal(t, err, errMock)
	assert.Nil(t, res)
}

func comperEntityAndDto(t *testing.T, entity *dal.FeatureFlagEntity, dto *ff_dto.FeatureFlagDto) {
	assert.Equal(t, entity.ID.Hex(), dto.ID)
	assert.Equal(t, entity.ServiceID.Hex(), dto.ServiceID)
	assert.Equal(t, entity.Name, dto.Name)
	assert.Equal(t, entity.Description, dto.Description)
	assert.Equal(t, entity.Archived, dto.Archived)
	assert.Equal(t, entity.Config, dto.Config)
}
