package api

import (
	"lightup/src/common/db"
	"lightup/src/modules/feature_flag/dal"
	"lightup/src/modules/feature_flag/model"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureFlagBlMock struct {
}

func (bl *FeatureFlagBlMock) GetFeatureFlagById(id string) (*dal.FeatureFlagEntity, error) {
	return genFeatureFlagEntity(), nil
}

func (bl *FeatureFlagBlMock) CreateFeatureFlag(input *model.CreateFeatureFlagDto) (*dal.FeatureFlagEntity, error) {
	return genFeatureFlagEntity(), nil
}

func genFeatureFlagEntity() *dal.FeatureFlagEntity {
	return &dal.FeatureFlagEntity{
		BaseEntity: db.BaseEntity{
			ID: primitive.NewObjectID(),
		},
		ServiceID:   primitive.NewObjectID(),
		Name:        "asdasd",
		Description: "asdasd",
		Archived:    false,
		Config: model.FeatureFlagConfig{
			Identifier: "id",
			Percentage: 0,
		},
	}
}

func TestCreateFeatureFlag(t *testing.T) {
	tested := &FeatureFlagApi{
		featureFlagBl: &FeatureFlagBlMock{},
	}

	res, err := tested.featureFlagBl.GetFeatureFlagById("234")

	if err != nil {
		t.Error(err)
	}

	if res.ID.Hex() == "" {
		t.Error("Invalid feature flag id")
	}
}
