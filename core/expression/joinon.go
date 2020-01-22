package expression

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/helpers"
)

type JoinType int

// Feel free to add any additional `JoinType`'s in a given dialect.
// Just ensure the `int` value for `JoinType` is positive as to not conflict
// with these `JoinType`s
const (
	InnerJoin JoinType = -iota // The `-` in front ensures all values are negative, yay C++ macros!
	OuterJoin
	LeftJoin
	RightJoin
)

type Joiner interface {
	Join(left core.SQL, joinType JoinType, on core.SQL) error
}

func NewJoinOn(left core.Expression, joinType JoinType, on core.Expression) core.ComboExpression {
	j := &joinOn{
		left:     left,
		on:       on,
		joinType: joinType,
	}
	j.exp = j
	return j
}

type joinOn struct {
	ComboExpression
	left     core.Expression
	on       core.Expression
	joinType JoinType
}

func (j joinOn) GetSQL(d core.Dialect, sql core.SQL) error {
	joiner, ok := d.(Joiner)
	if !ok {
		return DialectFunctionNotSupported("Joiner")
	}

	if helpers.IsValueNilOrEmpty(j.left) {
		return ExpressionIsNil("left")
	}
	if helpers.IsValueNilOrEmpty(j.on) {
		return ExpressionIsNil("on")
	}

	if lerr := j.left.GetSQL(d, sql); lerr != nil {
		return ErrorAroundSql(lerr, sql.String())
	}
	onSQL := sql.New()
	if oerr := j.on.GetSQL(d, onSQL); oerr != nil {
		return oerr
	}

	return joiner.Join(sql, j.joinType, onSQL)
}
