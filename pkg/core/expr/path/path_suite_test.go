package path_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core Expression
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core Table
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core Column
func TestPath(t *testing.T) {
	runner.SetupAndRunUnit(t, "path", "unit")
}
