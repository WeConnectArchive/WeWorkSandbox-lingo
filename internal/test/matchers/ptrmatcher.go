package matchers

import (
	"fmt"
	"reflect"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

// PtrMatcher expects the value to be a pointer.
type PtrMatcher struct {
	Expected types.GomegaMatcher
}

func (matcher *PtrMatcher) Match(actual interface{}) (success bool, err error) {
	if isNil(actual) {
		return false, fmt.Errorf("expected a pointer, got nil")
	}

	if !isPtr(actual) {
		return false, fmt.Errorf("expected a pointer.  Got:\n%s", format.Object(actual, 1))
	}

	reflectValueNoPtr := reflect.ValueOf(actual).Elem()
	valueItself := reflectValueNoPtr.Interface()

	return matcher.Expected.Match(valueItself)
}

func (matcher *PtrMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "be a pointer", matcher.Expected)
}

func (matcher *PtrMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not be a pointer", matcher.Expected)
}
