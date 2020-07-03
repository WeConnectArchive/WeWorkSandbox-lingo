package expression

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type Set interface {
	Set(left, value sql.Data) (sql.Data, error)
}

func NewSet(left core.Expression, value core.Expression) core.ComboExpression {
	return set{
		left:  left,
		value: value,
	}
}

type set struct {
	left  core.Expression
	value core.Expression
}

func (s set) And(exp core.Expression) core.ComboExpression {
	return NewOperator(s, operator.And, exp)
}

func (s set) Or(exp core.Expression) core.ComboExpression {
	return NewOperator(s, operator.Or, exp)
}

func (s set) ToSQL(d core.Dialect) (sql.Data, error) {
	set, ok := d.(Set)
	if !ok {
		return nil, DialectFunctionNotSupported("Set")
	}

	if check.IsValueNilOrEmpty(s.left) {
		return nil, ExpressionIsNil("left")
	}
	left, lerr := s.left.ToSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	if check.IsValueNilOrEmpty(s.value) {
		return nil, ErrorAroundSQL(ExpressionIsNil("value"), left.String())
	}
	value, verr := s.value.ToSQL(d)
	if verr != nil {
		return nil, ErrorAroundSQL(verr, left.String())
	}

	return set.Set(left, value)
}
