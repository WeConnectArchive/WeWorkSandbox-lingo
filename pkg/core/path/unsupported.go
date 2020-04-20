package path

import "github.com/weworksandbox/lingo/pkg/core"

func NewUnsupportedPathWithAlias(e core.Table, name, alias string) UnsupportedPath {
	return UnsupportedPath{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewUnsupportedPath(e core.Table, name string) UnsupportedPath {
	return NewUnsupportedPathWithAlias(e, name, "")
}

type UnsupportedPath struct {
	entity core.Table
	name   string
	alias  string
}

func (i UnsupportedPath) GetParent() core.Table {
	return i.entity
}

func (i UnsupportedPath) GetName() string {
	return i.name
}

func (i UnsupportedPath) GetAlias() string {
	return i.alias
}

func (UnsupportedPath) GetSQL(_ core.Dialect) (core.SQL, error) {
	// TODO - Revisit how we want to deal with unsupported columns
	return core.NewEmptySQL(), nil
}
