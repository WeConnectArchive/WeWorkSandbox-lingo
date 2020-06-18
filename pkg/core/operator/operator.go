package operator

import (
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type Operand int

const (
	Unknown Operand = iota

	And
	Or

	Eq
	NotEq
	Like
	NotLike

	LessThan
	LessThanOrEqual
	GreaterThan
	GreaterThanOrEqual

	Null
	NotNull

	In
	NotIn

	Between
	NotBetween
)

var _names = map[Operand]sql.Data{
	And:                sql.String("AND"),
	Or:                 sql.String("OR"),
	Eq:                 sql.String("="),
	NotEq:              sql.String("<>"),
	LessThan:           sql.String("<"),
	LessThanOrEqual:    sql.String("<="),
	GreaterThan:        sql.String(">"),
	GreaterThanOrEqual: sql.String(">="),
	Like:               sql.String("LIKE"),
	NotLike:            sql.String("NOT LIKE"),
	Null:               sql.String("IS NULL"),
	NotNull:            sql.String("IS NOT NULL"),
	In:                 sql.String("IN"),
	NotIn:              sql.String("NOT IN"),
	Between:            sql.String("BETWEEN"),
	NotBetween:         sql.String("NOT BETWEEN"),
}

func (o Operand) String() string {
	return _names[o].String()
}
