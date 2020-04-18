package query_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core -m Column
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core -m Expression
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core -m OrderBy
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core -m Set
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core -m Table
func TestQuery(t *testing.T) {
	runner.SetupAndRunUnit(t, "SQL")
}
