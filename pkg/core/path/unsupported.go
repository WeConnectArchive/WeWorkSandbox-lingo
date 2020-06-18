package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func NewUnsupportedWithAlias(e core.Table, name, alias string) Unsupported {
	return Unsupported{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewUnsupported(e core.Table, name string) Unsupported {
	return NewUnsupportedWithAlias(e, name, "")
}

type Unsupported struct {
	entity core.Table
	name   string
	alias  string
}

func (i Unsupported) GetParent() core.Table {
	return i.entity
}

func (i Unsupported) GetName() string {
	return i.name
}

func (i Unsupported) GetAlias() string {
	return i.alias
}

func (Unsupported) ToSQL(_ core.Dialect) (sql.Data, error) {
	// TODO - Revisit how we want to deal with unsupported columns. Right now we just ignore them.
	return sql.Empty(), nil
}
