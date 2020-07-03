package join

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type Joiner interface {
	Join(left sql.Data, joinType Type, on sql.Data) (sql.Data, error)
}

func NewJoinOn(left core.Expression, joinType Type, on core.Expression) core.ComboExpression {
	j := &joinOn{
		left:     left,
		on:       on,
		joinType: joinType,
	}
	j.exp = j
	return j
}

type joinOn struct {
	expression.ComboExpression
	left     core.Expression
	on       core.Expression
	joinType Type
}

func (j joinOn) ToSQL(d core.Dialect) (sql.Data, error) {
	joiner, ok := d.(Joiner)
	if !ok {
		return nil, expression.DialectFunctionNotSupported("Joiner")
	}

	if check.IsValueNilOrEmpty(j.on) {
		return nil, expression.ExpressionIsNil("on")
	}
	on, oerr := j.on.ToSQL(d)
	if oerr != nil {
		return nil, oerr
	}

	if check.IsValueNilOrEmpty(j.left) {
		return nil, expression.ExpressionIsNil("left")
	}
	left, lerr := j.left.ToSQL(d)
	if lerr != nil {
		return nil, expression.ErrorAroundSQL(lerr, on.String())
	}

	return joiner.Join(left, j.joinType, on)
}
