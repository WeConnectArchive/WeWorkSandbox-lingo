package join

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type Dialect interface {
	Join(left sql.Data, joinType Type, on sql.Data) (sql.Data, error)
}

func NewJoinOn(left core.Expression, joinType Type, on core.Expression) core.ComboExpression {
	return joinOn{
		left:     left,
		on:       on,
		joinType: joinType,
	}
}

type joinOn struct {
	left     core.Expression
	on       core.Expression
	joinType Type
}

func (j joinOn) And(exp core.Expression) core.ComboExpression {
	return operator.NewOperator(j, operator.And, exp)
}

func (j joinOn) Or(exp core.Expression) core.ComboExpression {
	return operator.NewOperator(j, operator.Or, exp)
}

func (j joinOn) ToSQL(d core.Dialect) (sql.Data, error) {
	joiner, ok := d.(Dialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'join.Dialect'", d.GetName())
	}

	if check.IsValueNilOrEmpty(j.on) {
		return nil, errors.New("join 'on' cannot be empty")
	}
	on, oerr := j.on.ToSQL(d)
	if oerr != nil {
		return nil, oerr
	}

	if check.IsValueNilOrEmpty(j.left) {
		return nil, errors.New("left of join cannot be empty")
	}
	left, lerr := j.left.ToSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	return joiner.Join(left, j.joinType, on)
}
