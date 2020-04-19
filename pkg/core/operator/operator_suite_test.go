package operator_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core Dialect
func TestOperator(t *testing.T) {
	runner.SetupAndRunUnit(t, "operator", "unit")
}
