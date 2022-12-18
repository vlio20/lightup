package str_rule

import (
	"errors"
	"strings"
)

const (
	StartsWith = "startsWith"
	EndsWith   = "endsWith"
	Contains   = "contains"
	Equals     = "equals"
)

func IsMatch(operator string, confVal string, servingVal string) (bool, error) {
	switch operator {
	case StartsWith:
		return strings.HasPrefix(servingVal, confVal), nil
	case EndsWith:
		return strings.HasSuffix(servingVal, confVal), nil
	case Contains:
		return strings.Contains(servingVal, confVal), nil
	case Equals:
		return servingVal == confVal, nil
	default:
		return false, errors.New("Invalid operator " + string(operator))
	}
}
