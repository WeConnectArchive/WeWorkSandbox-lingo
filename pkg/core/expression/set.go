package expression

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/check"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

type Set interface {
	Set(left, value sql.Data) (sql.Data, error)
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

func (c set) ToSQL(d core.Dialect) (sql.Data, error) {
	set, ok := d.(Set)
	if !ok {
		return nil, DialectFunctionNotSupported("Set")
	}

	if check.IsValueNilOrEmpty(c.left) {
		return nil, ExpressionIsNil("left")
	}
	left, lerr := c.left.ToSQL(d)
	if lerr != nil {
		return nil, lerr
	}

	if check.IsValueNilOrEmpty(c.value) {
		return nil, ErrorAroundSQL(ExpressionIsNil("value"), left.String())
	}
	value, verr := c.value.ToSQL(d)
	if verr != nil {
		return nil, ErrorAroundSQL(verr, left.String())
	}

	return set.Set(left, value)
}
