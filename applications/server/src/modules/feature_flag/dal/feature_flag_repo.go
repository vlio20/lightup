package dal

import (
	"lightup/src/common/db"
)

var getDb = db.GetDB

type FeatureFlagRepo struct {
	db.Repository[FeatureFlagEntity]
}

func NewFeatureFlagRepository() *FeatureFlagRepo {
	return &FeatureFlagRepo{
		Repository: db.Repository[FeatureFlagEntity]{
			DB:         getDb(),
			Collection: getDb().Collection("featureFlag"),
		},
	}
}

func (repo *FeatureFlagRepo) GetFeatureFlag(accountId string, serviceId string, name string) (*FeatureFlagEntity, error) {

	var entity = &FeatureFlagEntity{
		AccountID: repo.StrIdToObjectID(accountId),
		ServiceID: repo.StrIdToObjectID(serviceId),
		Name:      name,
	}

	return repo.FindOne(entity)
}
