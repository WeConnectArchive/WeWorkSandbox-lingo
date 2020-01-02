package query_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate github.com/weworksandbox/lingo/core Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo/core -m Column
//go:generate pegomock generate github.com/weworksandbox/lingo/core -m Expression
//go:generate pegomock generate github.com/weworksandbox/lingo/core -m OrderBy
//go:generate pegomock generate github.com/weworksandbox/lingo/core -m Set
//go:generate pegomock generate github.com/weworksandbox/lingo/core -m Table
func TestQuery(t *testing.T) {
	runner.SetupAndRun(t, "Query")
}
