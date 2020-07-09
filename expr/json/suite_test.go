package json_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate github.com/weworksandbox/lingo Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo -m Expression
func TestJSON(t *testing.T) {
	runner.SetupAndRunUnit(t, "json", "unit")
}
