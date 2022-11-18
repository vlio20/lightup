package dal

import (
	"lightup/src/common/db"
	"lightup/src/modules/feature_flag/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureFlagEntity struct {
	db.BaseEntity `bson:",inline"`
	ServiceID     primitive.ObjectID      `bson:"serviceId"`
	Name          string                  `bson:"name"`
	Description   string                  `bson:"description"`
	Archived      bool                    `bson:"archived"`
	Config        model.FeatureFlagConfig `bson:"config"`
}
