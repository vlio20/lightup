package model

import (
	"lightup/src/common/matching_rule"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureFlagConfig struct {
	MatchingRules []matching_rule.MatchingRule `bson:"matchingRules" json:"matchingRules" binding:"required"`
}

type CreateFeatureFlagInput struct {
	AccountID   primitive.ObjectID
	Name        string
	Description string
	ServiceID   primitive.ObjectID
	Archived    bool
	Config      FeatureFlagConfig `bson:"config"`
}
