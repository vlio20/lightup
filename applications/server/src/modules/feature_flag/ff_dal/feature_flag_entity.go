package ff_dal

import (
	"lightup/src/common/db"
	app_model "lightup/src/common/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureFlagEntity struct {
	db.BaseEntity `bson:",inline"`
	AccountID     primitive.ObjectID          `bson:"accountId, omitempty" `
	Tags          []primitive.ObjectID        `bson:"tags, omitempty"`
	Name          string                      `bson:"name, omitempty"`
	Description   string                      `bson:"description, omitempty"`
	Archived      bool                        `bson:"archived, omitempty"`
	Config        app_model.FeatureFlagConfig `bson:"config, omitempty"`
}
