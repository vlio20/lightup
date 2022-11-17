package dal

import "lightup/src/common/db"

type FeatureFlagEntity struct {
	db.BaseEntity
	Name        string `bson:"name"`
	Description string `bson:"description"`
}
