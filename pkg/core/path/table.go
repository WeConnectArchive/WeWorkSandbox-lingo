package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type ExpandTable interface {
	ExpandTable(entity core.Table) (sql.Data, error)
}

func ExpandTableWithDialect(d core.Dialect, entity core.Table) (sql.Data, error) {
	expand, ok := d.(ExpandTable)
	if !ok {
		return nil, expression.DialectFunctionNotSupported("ExpandTable")
	}
	return expand.ExpandTable(entity)
}
