package matchers

import (
	"fmt"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

// MatchSQLString creates a matcher to test if a given `core.SQL` string output
// has a specific string or a matcher of a string.
func MatchSQLString(expected interface{}) types.GomegaMatcher {
	return &SQLStringMatcher{
		Expected: expected,
	}
}

func MatchSQLValues(expected interface{}) types.GomegaMatcher {
	return &SQLValuesMatcher{
		Expected: expected,
	}
}

func EqString(format string, args ...interface{}) types.GomegaMatcher {
	return gomega.Equal(fmt.Sprintf(format, args...))
}
