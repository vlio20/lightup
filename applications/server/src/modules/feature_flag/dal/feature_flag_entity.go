package dal

import (
	"lightup/src/common/db"
	"lightup/src/modules/feature_flag/model"
)

type FeatureFlagEntity struct {
	db.BaseEntity `bson:",inline"`
	Name          string                  `bson:"name"`
	Description   string                  `bson:"description"`
	Archived      bool                    `bson:"archived"`
	Config        model.FeatureFlagConfig `bson:"config"`
}
