package app_model

import "lightup/src/common/matching_rule"

type FeatureFlagConfig struct {
	Segment    MatchingSegment `json:"segment" bson:"segment" binding:"required"`
	Percentage float32         `bson:"percentage" json:"percentage" binding:"required,min=0,max=100"`
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
	Name    string                     `json:"name" bson:"name"`
	Rule    matching_rule.MatchingRule `json:"rule" bson:"rule" binding:"required"`
	Nesting SegmentNesting             `json:"nesting" bson:"nesting"`
}

func (s *MatchingSegment) IsMatch(val map[string][]string) (bool, error) {
	var match = false
	var err error

	if s.Rule == nil {
		vals := val[s.Rule.GetKey()]

		if len(vals) == 0 {
			return false, nil
		}

		val := vals[0]

		match, err = s.Rule.IsMatch(val)
	}

	if s.Nesting.Segment != nil {
		var nestedMatch bool

		nestedMatch, err = s.Nesting.Segment.IsMatch(val)

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

	return match, err
}
