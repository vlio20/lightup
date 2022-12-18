package num_rule

import (
	"errors"
	"lightup/src/common/http"
	"strconv"
)

type NumericOperator string

const (
	LessThan       = "lt"
	LessOrEqual    = "le"
	Equal          = "eq"
	GreaterOrEqual = "ge"
	GreaterThan    = "gt"
)

func IsMatch(operator string, confStrVal string, servingVal string) (bool, error) {
	confVal, err := strconv.ParseFloat(confStrVal, 64)

	if err != nil {
		return false, errors.New("Invalid config value" + servingVal)
	}

	val, err := strconv.ParseFloat(servingVal, 64)

	if err != nil {
		return false, http.Error{
			StatusCode:    400,
			Message:       "Invalid serving value" + servingVal,
			OriginalError: err,
		}
	}

	switch operator {
	case LessThan:
		return val < confVal, nil
	case LessOrEqual:
		return val <= confVal, nil
	case Equal:
		return val == confVal, nil
	case GreaterOrEqual:
		return val >= confVal, nil
	case GreaterThan:
		return val > confVal, nil
	default:
		return false, errors.New("Invalid operator " + operator)
	}
}
