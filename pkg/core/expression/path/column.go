package path

import (
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type ExpandColumnDialect interface {
	ExpandColumn(entity core.Column) (sql.Data, error)
}

func ExpandColumnWithDialect(d core.Dialect, path core.Column) (sql.Data, error) {
	expand, ok := d.(ExpandColumnDialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'ExpandColumnDialect'", d.GetName())
	}
	return expand.ExpandColumn(path)
}
