package execute_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core -m Expression
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core/execute -m SQLExp
func TestExecute(t *testing.T) {
	runner.SetupAndRunUnit(t, "expression", "unit")
}
