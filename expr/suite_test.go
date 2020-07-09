package expr_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate github.com/weworksandbox/lingo Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo -m Expression
func TestExpr(t *testing.T) {
	runner.SetupAndRunUnit(t, "expr", "unit")
}
