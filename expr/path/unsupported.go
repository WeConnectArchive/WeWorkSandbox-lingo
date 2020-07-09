package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/sql"
)

func NewUnsupportedWithAlias(e lingo.Table, name, alias string) Unsupported {
	return Unsupported{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewUnsupported(e lingo.Table, name string) Unsupported {
	return NewUnsupportedWithAlias(e, name, "")
}

type Unsupported struct {
	entity lingo.Table
	name   string
	alias  string
}

func (i Unsupported) GetParent() lingo.Table {
	return i.entity
}

func (i Unsupported) GetName() string {
	return i.name
}

func (i Unsupported) GetAlias() string {
	return i.alias
}

func (Unsupported) ToSQL(_ lingo.Dialect) (sql.Data, error) {
	// TODO - Revisit how we want to deal with unsupported columns. Right now we just ignore them.
	return sql.Empty(), nil
}
