package sort

import (
	"github.com/weworksandbox/lingo/pkg/core"
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

func (d Direction) GetSQL(dialect core.Dialect) (core.SQL, error) {
	return core.NewSQL(d.String(), nil), nil
}
