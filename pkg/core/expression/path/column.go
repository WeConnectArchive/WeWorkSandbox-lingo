package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type ExpandColumn interface {
	ExpandColumn(entity core.Column) (sql.Data, error)
}

func ExpandColumnWithDialect(d core.Dialect, path core.Column) (sql.Data, error) {
	expand, ok := d.(ExpandColumn)
	if !ok {
		return nil, expression.DialectFunctionNotSupported("ExpandColumn")
	}
	return expand.ExpandColumn(path)
}
