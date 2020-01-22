package path

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/expression"
)

type ExpandTable interface {
	ExpandTable(entity core.Table, sql core.SQL) error
}

func ExpandTableWithDialect(d core.Dialect, entity core.Table, sql core.SQL) error {
	expand, ok := d.(ExpandTable)
	if !ok {
		return expression.DialectFunctionNotSupported("ExpandTable")
	}
	return expand.ExpandTable(entity, sql)
}
