package app_model

import "lightup/src/common/matching_rule"

type MatchingSegment struct {
	MatchingRules []matching_rule.MatchingRule `bson:"matchingRules" json:"matchingRules" binding:"required"`
	Percentage    float32                      `bson:"percentage" json:"percentage" binding:"required,min=0,max=100"`
}

type FeatureFlagConfig struct {
	MatchingSegments []MatchingSegment `bson:"machingSegment" json:"machingSegment" binding:"required"`
}
