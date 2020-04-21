package expression

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/helpers"
	"github.com/weworksandbox/lingo/pkg/core/join"
)

type Joiner interface {
	Join(left core.SQL, joinType join.Type, on core.SQL) (core.SQL, error)
}

func NewJoinOn(left core.Expression, joinType join.Type, on core.Expression) core.ComboExpression {
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
	joinType join.Type
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
