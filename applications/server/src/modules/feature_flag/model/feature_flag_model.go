package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type FeatureFlagConfig struct {
	Identifier string  `bson:"indentifier" json:"identifier" binding:"required"`
	Percentage float32 `bson:"percentage" json:"percentage" binding:"required"`
}

type CreateFeatureFlagDto struct {
	AccountID   primitive.ObjectID
	Name        string
	Description string
	ServiceID   primitive.ObjectID
	Archived    bool              `bson:"archived"`
	Config      FeatureFlagConfig `bson:"config"`
}
