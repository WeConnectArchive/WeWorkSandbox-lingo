package execute_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate database/sql Result
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core -m Expression
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core -m --output mock_core_sql_test.go --mock-name=MockCoreSQL SQL
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core/execute -m SQL
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core/execute -m RowScanner
func TestExecute(t *testing.T) {
	runner.SetupAndRunUnit(t, "expression", "unit")
}
