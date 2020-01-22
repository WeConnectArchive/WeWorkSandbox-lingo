package expression

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/helpers"
)

type Set interface {
	Set(left, value core.SQL) error
}

func NewSet(left core.Expression, value core.Expression) *set {
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

func (c set) GetSQL(d core.Dialect, sql core.SQL) error {
	set, ok := d.(Set)
	if !ok {
		return DialectFunctionNotSupported("Set")
	}

	if helpers.IsValueNilOrEmpty(c.left) {
		return ExpressionIsNil("left")
	}
	if helpers.IsValueNilOrEmpty(c.value) {
		return ErrorAroundSql(ExpressionIsNil("value"), sql.String())
	}

	if lerr := c.left.GetSQL(d, sql); lerr != nil {
		return lerr
	}

	valueSQL := sql.New()
	if verr := c.value.GetSQL(d, valueSQL); verr != nil {
		return ErrorAroundSql(verr, sql.String())
	}
	return set.Set(sql, valueSQL)
}
