package queries_test

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/dialect"
)

func DefaultDialect() (lingo.Dialect, error) {
	return dialect.NewDefault()
}

func DefaultDialectWithSchema() (lingo.Dialect, error) {
	return dialect.NewDefault(
		dialect.WithSchemaNameIncluded(true),
	)
}

func MySQLDialect() (lingo.Dialect, error) {
	return dialect.NewMySQL()
}
