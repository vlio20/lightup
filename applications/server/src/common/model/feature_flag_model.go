package app_model

import (
	"lightup/src/common/http"
	"lightup/src/common/matching_rule"
)

type FeatureFlagConfig struct {
	Segment    MatchingSegment `json:"segment" bson:"segment" binding:"required"`
	Identifier string          `json:"identifier" bson:"identifier" binding:"required"`
	Percentage float32         `bson:"percentage" json:"percentage" binding:"required,min=0,max=1"`
}

type SegmentOperator string

const (
	Or  SegmentOperator = "or"
	And SegmentOperator = "and"
)

type SegmentNesting struct {
	Operator SegmentOperator  `json:"operator" bson:"operator" binding:"required"`
	Segment  *MatchingSegment `json:"segment" bson:"segment" binding:"required"`
}

type MatchingSegment struct {
	Name    string              `json:"name" bson:"name"`
	Rule    *matching_rule.Rule `json:"rule" bson:"rule"`
	Nesting *SegmentNesting     `json:"nesting" bson:"nesting"`
}

func (s *MatchingSegment) IsMatch(valuesMap map[string][]string) (bool, error) {
	var match = false
	var err error

	if s.Rule != nil {
		vals := valuesMap[s.Rule.Key]

		if len(vals) == 0 {
			return false, http.Error{
				StatusCode: 400,
				Message:    "Missing key: " + s.Rule.Key,
			}
		}

		val := vals[0]

		match, err = s.Rule.IsMatch(val)
	} else {
		match = true
	}

	if s.Nesting != nil {
		if s.Nesting.Segment != nil {
			var nestedMatch bool

			nestedMatch, err = s.Nesting.Segment.IsMatch(valuesMap)

			if err != nil {
				return false, err
			}

			switch s.Nesting.Operator {
			case Or:
				match = match || nestedMatch
			case And:
				match = match && nestedMatch
			}
		}
	}

	return match, err
}
