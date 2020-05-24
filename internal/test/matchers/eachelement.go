package matchers

import (
	"fmt"
	"reflect"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

// MatchEachElement matches each element in the slice, array, or map Value (not map Keys)
type MatchEachElement struct {
	Element interface{}
}

func (matcher *MatchEachElement) Match(actual interface{}) (success bool, err error) {
	if isNil(actual) {
		return false, fmt.Errorf("expected an array, slice, or map, but got nil")
	}

	if !isArrayOrSlice(actual) && !isMap(actual) {
		return false, fmt.Errorf("expected an array, slice, or map.  Got:\n%s",
			format.Object(actual, 1))
	}

	elemMatcher, elementIsMatcher := matcher.Element.(gomega.OmegaMatcher)
	if !elementIsMatcher {
		return false, fmt.Errorf("expected Element matcher to be a matcher. Got:\n%s",
			format.Object(matcher.Element, 1))
	}

	value := reflect.ValueOf(actual)
	var valueAt func(int) interface{}
	if isMap(actual) {
		keys := value.MapKeys()
		valueAt = func(i int) interface{} {
			return value.MapIndex(keys[i]).Interface()
		}
	} else {
		valueAt = func(i int) interface{} {
			return value.Index(i).Interface()
		}
	}

	for i := 0; i < value.Len(); i++ {
		success, err := elemMatcher.Match(valueAt(i))
		if err != nil {
			return false, fmt.Errorf("failed match at index %d: %w", i, err)
		}
		if !success {
			return false, fmt.Errorf("failed match at index %d", i)
		}
	}
	return true, nil
}

func (matcher *MatchEachElement) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to have each element matching", matcher.Element)
}

func (matcher *MatchEachElement) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to have each element matching", matcher.Element)
}
