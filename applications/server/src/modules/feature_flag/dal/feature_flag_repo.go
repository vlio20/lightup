package dal

import (
	"lightup/src/common/db"
)

var getDb = db.GetDB

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
