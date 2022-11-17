package bl

import "lightup/src/modules/feature_flag/dal"

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
