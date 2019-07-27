package expression_test

import (
	"github.com/weworksandbox/lingo/internal/test/runner"
	"testing"
)

//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core -m Expression
func TestExpression(t *testing.T) {
	runner.SetupAndRun(t, "expression")
}
