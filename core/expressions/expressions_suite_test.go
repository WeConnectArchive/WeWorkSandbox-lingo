package expressions_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate github.com/weworksandbox/lingo/core Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo/core Expression
func TestExpressions(t *testing.T) {
	runner.SetupAndRun(t, "expressions")
}
