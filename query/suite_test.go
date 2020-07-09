package query_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate github.com/weworksandbox/lingo Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo -m Column
//go:generate pegomock generate github.com/weworksandbox/lingo -m Expression
//go:generate pegomock generate github.com/weworksandbox/lingo -m OrderBy
//go:generate pegomock generate github.com/weworksandbox/lingo -m Set
//go:generate pegomock generate github.com/weworksandbox/lingo -m Table
//go:generate pegomock generate github.com/weworksandbox/lingo/query Modifier
func TestQuery(t *testing.T) {
	runner.SetupAndRunUnit(t, "query", "unit")
}
