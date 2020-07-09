package sql_test

import (
	"testing"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

func TestSQL(t *testing.T) {
	runner.SetupAndRunUnit(t, "sql", "unit")
}
