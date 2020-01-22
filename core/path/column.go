package path

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/expression"
)

type ExpandColumn interface {
	ExpandColumn(entity core.Column, sql core.SQL) error
}

func ExpandColumnWithDialect(d core.Dialect, path core.Column, sql core.SQL) error {
	expand, ok := d.(ExpandColumn)
	if !ok {
		return expression.DialectFunctionNotSupported("ExpandColumn")
	}
	return expand.ExpandColumn(path, sql)
}
