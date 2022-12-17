package ff_dal

import (
	"go.mongodb.org/mongo-driver/bson"
	"lightup/src/common/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlagRepo struct {
	db.Repository[FeatureFlagEntity]
}

func NewFlagRepository() *FlagRepo {
	return &FlagRepo{
		Repository: db.Repository[FeatureFlagEntity]{
			DB:         db.GetDB(),
			Collection: db.GetDB().Collection("featureFlag"),
		},
	}
}

func (repo *FlagRepo) GetFeatureFlag(accountId primitive.ObjectID, name string) (*FeatureFlagEntity, error) {
	var entity = &FeatureFlagEntity{
		AccountID: accountId,
		Name:      name,
	}

	return repo.FindOne(bson.M{"accountId": entity.AccountID, "name": entity.Name})
}

func (repo *FlagRepo) GetFlagByName(accountId primitive.ObjectID, name string) (*FeatureFlagEntity, error) {
	return repo.FindOne(bson.M{"accountId": accountId, "name": name})
}
