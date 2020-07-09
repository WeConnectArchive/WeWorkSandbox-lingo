package execute_test

import (
	"reflect"
	"testing"

	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate database/sql Result
//go:generate pegomock generate go.opentelemetry.io/otel/api/trace -m Span
//go:generate pegomock generate github.com/weworksandbox/lingo Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo -m Expression
//go:generate pegomock generate github.com/weworksandbox/lingo/sql Data
//go:generate pegomock generate github.com/weworksandbox/lingo/execute -m SQL
//go:generate pegomock generate github.com/weworksandbox/lingo/execute -m RowScanner
//go:generate pegomock generate github.com/weworksandbox/lingo/execute -m TxSQL
func TestExecute(t *testing.T) {
	runner.SetupAndRunUnit(t, "execute", "unit")
}

func AnyError() error {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(error))(nil)).Elem()))
	var nullValue error
	return nullValue
}
