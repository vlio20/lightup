package model

import (
	ff_model "lightup/src/common/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateFeatureFlagInput struct {
	AccountID   primitive.ObjectID
	Name        string
	Description string
	ServiceID   primitive.ObjectID
	Archived    bool
	Config      ff_model.FeatureFlagConfig `bson:"config"`
}
