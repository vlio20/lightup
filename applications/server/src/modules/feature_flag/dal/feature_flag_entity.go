package dal

import (
	"lightup/src/common/db"
	"lightup/src/modules/feature_flag/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureFlagEntity struct {
	db.BaseEntity `bson:",inline"`
	AccountID     primitive.ObjectID      `bson:"accountId, omitempty" `
	ServiceID     primitive.ObjectID      `bson:"serviceId, omitempty"`
	Name          string                  `bson:"name, omitempty"`
	Description   string                  `bson:"description, omitempty"`
	Archived      bool                    `bson:"archived, omitempty"`
	Config        model.FeatureFlagConfig `bson:"config, omitempty"`
}
