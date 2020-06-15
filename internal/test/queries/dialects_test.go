package queries_test

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/dialect"
)

func DefaultDialect() (core.Dialect, error) {
	return dialect.NewDefault()
}

func DefaultDialectWithSchema() (core.Dialect, error) {
	return dialect.NewDefault(
		dialect.WithSchemaNameIncluded(true),
	)
}

func MySQLDialect() (core.Dialect, error) {
	return dialect.NewMySQL()
}
