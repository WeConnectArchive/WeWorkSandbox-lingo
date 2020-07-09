package path_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate github.com/weworksandbox/lingo Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo Expression
//go:generate pegomock generate github.com/weworksandbox/lingo Table
//go:generate pegomock generate github.com/weworksandbox/lingo Column
func TestPath(t *testing.T) {
	runner.SetupAndRunUnit(t, "path", "unit")
}
