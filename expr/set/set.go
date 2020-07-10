package set

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/sql"
)

type Dialect interface {
	Set(left, value sql.Data) (sql.Data, error)
}

func NewSet(left lingo.Expression, value lingo.Expression) Set {
	return Set{
		left:  left,
		value: value,
	}
}

type Set struct {
	left  lingo.Expression
	value lingo.Expression
}

func (s Set) And(exp lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(s, operator.And, exp)
}

func (s Set) Or(exp lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(s, operator.Or, exp)
}

func (s Set) ToSQL(d lingo.Dialect) (sql.Data, error) {
	setFunc, ok := d.(Dialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'set.Dialect'", d.GetName())
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
