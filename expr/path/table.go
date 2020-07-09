package path

import (
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/sql"
)

type ExpandTableDialect interface {
	ExpandTable(entity lingo.Table) (sql.Data, error)
}

func ExpandTableWithDialect(d lingo.Dialect, entity lingo.Table) (sql.Data, error) {
	expand, ok := d.(ExpandTableDialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'path.ExpandTableDialect'", d.GetName())
	}
	return expand.ExpandTable(entity)
}
