package path

import (
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/sql"
)

type ExpandColumnDialect interface {
	ExpandColumn(entity lingo.Column) (sql.Data, error)
}

func ExpandColumnWithDialect(d lingo.Dialect, path lingo.Column) (sql.Data, error) {
	expand, ok := d.(ExpandColumnDialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'path.ExpandColumnDialect'", d.GetName())
	}
	return expand.ExpandColumn(path)
}
