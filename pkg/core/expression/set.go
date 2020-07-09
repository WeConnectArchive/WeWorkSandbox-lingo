package expression

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/expression/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type SetDialect interface {
	Set(left, value sql.Data) (sql.Data, error)
}

func NewSet(left core.Expression, value core.Expression) Set {
	return Set{
		left:  left,
		value: value,
	}
}

type Set struct {
	left  core.Expression
	value core.Expression
}

func (s Set) And(exp core.Expression) core.ComboExpression {
	return operator.NewOperator(s, operator.And, exp)
}

func (s Set) Or(exp core.Expression) core.ComboExpression {
	return operator.NewOperator(s, operator.Or, exp)
}

func (s Set) ToSQL(d core.Dialect) (sql.Data, error) {
	setFunc, ok := d.(SetDialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'expression.SetDialect'", d.GetName())
	}

	if check.IsValueNilOrEmpty(s.left) {
		return nil, errors.New("left of 'set' cannot be empty")
	}
	left, lerr := s.left.ToSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	if check.IsValueNilOrEmpty(s.value) {
		return nil, errors.New("set 'value' cannot be empty")
	}
	v, verr := s.value.ToSQL(d)
	if verr != nil {
		return nil, verr
	}

	return setFunc.Set(left, v)
}
