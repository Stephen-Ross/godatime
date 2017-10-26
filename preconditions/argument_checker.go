package preconditions

import (
	"errors"
	"fmt"
)

func CheckArgumentRange(parameterName string, value, minInclusive, maxInclusive int) error {
	if value < minInclusive || value > maxInclusive {
		return errors.New(fmt.Sprintf("%s' value should be in range [%v-%v]", parameterName, minInclusive, maxInclusive))
	}

	return nil
}

func CheckArgumentRangeInt64(parameterName string, value, minInclusive, maxInclusive int64) error {
	if value < minInclusive || value > maxInclusive {
		return errors.New(fmt.Sprintf("%s' value should be in range [%v-%v]", parameterName, minInclusive, maxInclusive))
	}

	return nil
}

func CheckArgumentRangeFloat64(parameterName string, value, minInclusive, maxInclusive int64) error {
	if value < minInclusive || value > maxInclusive {
		return errors.New(fmt.Sprintf("%s' value should be in range [%v-%v]", parameterName, minInclusive, maxInclusive))
	}

	return nil
}