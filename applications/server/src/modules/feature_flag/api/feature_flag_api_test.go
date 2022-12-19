package api

import (
	"lightup/src/common/db"
	"lightup/src/common/http"
	app_model "lightup/src/common/model"
	ff_dto "lightup/src/modules/feature_flag/dto"
	"lightup/src/modules/feature_flag/ff_dal"
	"lightup/src/modules/feature_flag/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ffMock = genFeatureFlagEntity()

type FeatureFlagBlMock struct {
}

func (bl *FeatureFlagBlMock) GetFeatureFlagById(id primitive.ObjectID) (*ff_dal.FeatureFlagEntity, error) {
	if id == ffMock.ID {
		return ffMock, nil
	}

	return nil, nil
}

func (bl *FeatureFlagBlMock) CreateFeatureFlag(input *model.CreateFeatureFlagInput) (*ff_dal.FeatureFlagEntity, error) {
	return ffMock, nil
}

func (bl *FeatureFlagBlMock) GetFeatureFlag(accountId primitive.ObjectID, name string) (*ff_dal.FeatureFlagEntity, error) {
	return ffMock, nil
}

func genFeatureFlagEntity() *ff_dal.FeatureFlagEntity {
	return &ff_dal.FeatureFlagEntity{
		BaseEntity: db.BaseEntity{
			ID: primitive.NewObjectID(),
		},
		Tags:        []primitive.ObjectID{primitive.NewObjectID()},
		Name:        "name",
		Description: "description",
		Archived:    false,
		Config:      app_model.FeatureFlagConfig{},
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

	assert.Equal(t, err.(*http.Error).StatusCode, 404)
	assert.Nil(t, res)
}

func comperEntityAndDto(t *testing.T, entity *ff_dal.FeatureFlagEntity, dto *ff_dto.FeatureFlagDto) {
	assert.Equal(t, entity.ID, dto.ID)
	assert.Equal(t, entity.Tags, dto.Tags)
	assert.Equal(t, entity.AccountID, dto.AccountID)
	assert.Equal(t, entity.Name, dto.Name)
	assert.Equal(t, entity.Description, dto.Description)
	assert.Equal(t, entity.Archived, dto.Archived)
	assert.Equal(t, entity.Config, dto.Config)
}
