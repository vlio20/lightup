package bl

import (
	"lightup/src/common/db"
	"lightup/src/modules/feature_flag/dal"
	"lightup/src/modules/feature_flag/model"
)

var featureFlagRepo *dal.FeatureFlagRepo

var getFeatureFlegRepo = func() *dal.FeatureFlagRepo {
	if featureFlagRepo == nil {
		featureFlagRepo = dal.NewFeatureFlagRepository()
	}

	return featureFlagRepo
}

func GetFeatureFlagById(id string) (*dal.FeatureFlagEntity, error) {
	return getFeatureFlegRepo().GetById(id)
}

func CreateFeatureFlag(input *model.CreateFeatureFlagDto) (*dal.FeatureFlagEntity, error) {
	entity := &dal.FeatureFlagEntity{
		Name:        input.Name,
		Description: input.Description,
		BaseEntity:  *db.GetBaseEntity(),
	}

	return getFeatureFlegRepo().Add(entity)
}
