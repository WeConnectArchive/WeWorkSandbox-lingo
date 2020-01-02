package path_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate github.com/weworksandbox/lingo/core Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo/core Expression
//go:generate pegomock generate github.com/weworksandbox/lingo/core Table
//go:generate pegomock generate github.com/weworksandbox/lingo/core Column
func TestPath(t *testing.T) {
	runner.SetupAndRun(t, "path")
}
