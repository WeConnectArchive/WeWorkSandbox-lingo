package execute_test

import (
	"reflect"
	"testing"

	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo/internal/test/runner"
)

//go:generate pegomock generate database/sql Result
//go:generate pegomock generate go.opentelemetry.io/otel/api/trace -m Span
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core Dialect
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core -m Expression
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core/sql Data
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core/execute -m SQL
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core/execute -m RowScanner
//go:generate pegomock generate github.com/weworksandbox/lingo/pkg/core/execute -m TxSQL
func TestExecute(t *testing.T) {
	runner.SetupAndRunUnit(t, "expression", "unit")
}

func AnyError() error {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(error))(nil)).Elem()))
	var nullValue error
	return nullValue
}
