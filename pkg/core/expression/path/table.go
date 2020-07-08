package path

import (
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type ExpandTableDialect interface {
	ExpandTable(entity core.Table) (sql.Data, error)
}

func ExpandTableWithDialect(d core.Dialect, entity core.Table) (sql.Data, error) {
	expand, ok := d.(ExpandTableDialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'ExpandTableDialect'", d.GetName())
	}
	return expand.ExpandTable(entity)
}
