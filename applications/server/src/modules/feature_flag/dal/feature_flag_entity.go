package dal

import "lightup/src/common/db"

type FeatureFlagEntity struct {
	db.BaseEntity `bson:",inline"`
	Name          string `bson:"name"`
	Description   string `bson:"description"`
}
