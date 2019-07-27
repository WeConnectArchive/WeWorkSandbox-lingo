package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
)

type ExpandColumn interface {
	ExpandColumn(entity core.Column) (core.SQL, error)
}

func ExpandColumnWithDialect(d core.Dialect, path core.Column) (core.SQL, error) {
	expand, ok := d.(ExpandColumn)
	if !ok {
		return nil, expression.DialectFunctionNotSupported("ExpandColumn")
	}
	return expand.ExpandColumn(path)
}
