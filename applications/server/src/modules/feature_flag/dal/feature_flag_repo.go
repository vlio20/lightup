package dal

import (
	"lightup/src/common/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repo *FeatureFlagRepo) GetFeatureFlag(accountId primitive.ObjectID, serviceId primitive.ObjectID, name string) (*FeatureFlagEntity, error) {

	var entity = &FeatureFlagEntity{
		AccountID: accountId,
		ServiceID: serviceId,
		Name:      name,
	}

	return repo.FindOne(entity)
}
