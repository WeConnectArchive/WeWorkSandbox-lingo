package expression

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
)

type Set interface {
	Set(left, value core.SQL) (core.SQL, error)
}

func NewSet(left core.Expression, value core.Expression) core.ComboExpression {
	e := &set{
		left:  left,
		value: value,
	}
	e.exp = e
	return e
}

type set struct {
	ComboExpression
	left  core.Expression
	value core.Expression
}

func (c set) GetSQL(d core.Dialect) (core.SQL, error) {
	set, ok := d.(Set)
	if !ok {
		return nil, DialectFunctionNotSupported("Set")
	}

	if check.IsValueNilOrEmpty(c.left) {
		return nil, ExpressionIsNil("left")
	}
	left, lerr := c.left.GetSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	if check.IsValueNilOrEmpty(c.value) {
		return nil, ErrorAroundSQL(ExpressionIsNil("value"), left.String())
	}
	value, verr := c.value.GetSQL(d)
	if verr != nil {
		return nil, ErrorAroundSQL(verr, left.String())
	}

	return set.Set(left, value)
}
