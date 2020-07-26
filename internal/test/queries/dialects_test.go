package queries_test

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/dialect"
)

func DefaultDialect() (lingo.Dialect, error) {
	return dialect.NewDialect()
}

func DefaultDialectWithSchema() (lingo.Dialect, error) {
	return dialect.NewDialect(
		dialect.WithSchemaNameIncluded(true),
	)
}
