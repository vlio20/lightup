package app_model

import "lightup/src/common/matching_rule"

type MatchingSegment struct {
	Name          string                       `bson:"name" json:"name" binding:"required"`
	MatchingRules []matching_rule.MatchingRule `bson:"matchingRules" json:"matchingRules" binding:"required"`
}

type FeatureFlagConfig struct {
	MatchingSegments []MatchingSegment `bson:"machingSegment" json:"machingSegment" binding:"required"`
	Percentage       float32           `bson:"percentage" json:"percentage" binding:"required,min=0,max=100"`
}
