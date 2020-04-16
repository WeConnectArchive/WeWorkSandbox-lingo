package expression

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/helpers"
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
	Join(left core.SQL, joinType JoinType, on core.SQL) (core.SQL, error)
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

func (j joinOn) GetSQL(d core.Dialect) (core.SQL, error) {
	joiner, ok := d.(Joiner)
	if !ok {
		return nil, DialectFunctionNotSupported("Joiner")
	}

	if helpers.IsValueNilOrEmpty(j.on) {
		return nil, ExpressionIsNil("on")
	}
	on, oerr := j.on.GetSQL(d)
	if oerr != nil {
		return nil, oerr
	}

	if helpers.IsValueNilOrEmpty(j.left) {
		return nil, ExpressionIsNil("left")
	}
	left, lerr := j.left.GetSQL(d)
	if lerr != nil {
		return nil, ErrorAroundSQL(lerr, on.String())
	}

	return joiner.Join(left, j.joinType, on)
}
