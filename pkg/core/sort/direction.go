package sort

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type Direction int

const (
	Unknown Direction = iota

	Ascending
	Descending
)

var _names = map[Direction]string{
	Ascending:  "ASC",
	Descending: "DESC",
}

func (d Direction) String() string {
	return _names[d]
}

func (d Direction) ToSQL(_ core.Dialect) (sql.Data, error) {
	return sql.String(d.String()), nil
}
