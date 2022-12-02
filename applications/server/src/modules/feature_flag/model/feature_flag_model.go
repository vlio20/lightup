package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RuleType string

const (
	Number RuleType = "numer"
	String RuleType = "string"
	Ip     RuleType = "ip"
)

type Rule struct {
	Type RuleType `json:"type" bson:"type"`
}

type FeatureFlagConfig struct {
	Identifier string  `bson:"indentifier" json:"identifier" binding:"required"`
	Percentage float32 `bson:"percentage" json:"percentage" binding:"required"`
}

type CreateFeatureFlagDto struct {
	AccountID     primitive.ObjectID
	Name          string
	Description   string
	ServiceID     primitive.ObjectID
	Archived      bool              `bson:"archived"`
	Config        FeatureFlagConfig `bson:"config"`
	MatchingRules []Rule            `bson:"mathchingrules"`
}
