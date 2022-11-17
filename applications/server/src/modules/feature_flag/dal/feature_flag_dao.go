package dal

import (
	"lightup/src/common/db"

	"go.mongodb.org/mongo-driver/mongo"
)

var getDb = func() *mongo.Database {
	return db.GetDB("lightup")
}

type A struct {
	A int
}

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
