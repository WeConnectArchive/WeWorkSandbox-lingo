package expression

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/join"
	"github.com/weworksandbox/lingo/pkg/core/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type Joiner interface {
	Join(left sql.Data, joinType join.Type, on sql.Data) (sql.Data, error)
}

func NewJoinOn(left core.Expression, joinType join.Type, on core.Expression) core.ComboExpression {
	return joinOn{
		left:     left,
		on:       on,
		joinType: joinType,
	}
}

type joinOn struct {
	left     core.Expression
	on       core.Expression
	joinType join.Type
}

func (j joinOn) And(exp core.Expression) core.ComboExpression {
	return NewOperator(j, operator.And, exp)
}

func (j joinOn) Or(exp core.Expression) core.ComboExpression {
	return NewOperator(j, operator.Or, exp)
}

func (j joinOn) ToSQL(d core.Dialect) (sql.Data, error) {
	joiner, ok := d.(Joiner)
	if !ok {
		return nil, DialectFunctionNotSupported("Joiner")
	}

	if check.IsValueNilOrEmpty(j.on) {
		return nil, ExpressionIsNil("on")
	}
	on, oerr := j.on.ToSQL(d)
	if oerr != nil {
		return nil, oerr
	}

	if check.IsValueNilOrEmpty(j.left) {
		return nil, ExpressionIsNil("left")
	}
	left, lerr := j.left.ToSQL(d)
	if lerr != nil {
		return nil, ErrorAroundSQL(lerr, on.String())
	}

	return joiner.Join(left, j.joinType, on)
}
