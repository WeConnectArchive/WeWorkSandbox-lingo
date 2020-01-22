package path

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/expression"
	"github.com/weworksandbox/lingo/core/operator"
)

func NewIntPathWithAlias(e core.Table, name, alias string) IntPath {
	return IntPath{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewIntPath(e core.Table, name string) IntPath {
	return NewIntPathWithAlias(e, name, "")
}

type IntPath struct {
	entity core.Table
	name   string
	alias  string
}

func (i IntPath) GetParent() core.Table {
	return i.entity
}

func (i IntPath) GetName() string {
	return i.name
}

func (i IntPath) GetAlias() string {
	return i.alias
}

func (i IntPath) As(alias string) IntPath {
	i.alias = alias
	return i
}

func (i IntPath) GetSQL(d core.Dialect, sql core.SQL) error {
	return ExpandColumnWithDialect(d, i, sql)
}

func (i IntPath) To(value int) core.Set {
	return expression.NewSet(i, expression.NewValue(value))
}

func (i IntPath) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(i, setExp)
}

func (i IntPath) Eq(equalTo int) core.ComboExpression {
	return expression.NewOperator(i, operator.Eq, expression.NewValue(equalTo))
}

func (i IntPath) EqPath(equalTo core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.Eq, equalTo)
}

func (i IntPath) NotEq(notEqualTo int) core.ComboExpression {
	return expression.NewOperator(i, operator.NotEq, expression.NewValue(notEqualTo))
}

func (i IntPath) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotEq, notEqualTo)
}

func (i IntPath) LT(lessThan int) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThan, expression.NewValue(lessThan))
}

func (i IntPath) LTPath(lessThan core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThan, lessThan)
}

func (i IntPath) LTOrEq(lessThanOrEqual int) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThanOrEqual, expression.NewValue(lessThanOrEqual))
}

func (i IntPath) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i IntPath) GT(greaterThan int) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThan, expression.NewValue(greaterThan))
}

func (i IntPath) GTPath(greaterThan core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThan, greaterThan)
}

func (i IntPath) GTOrEq(greaterThanOrEqual int) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThanOrEqual, expression.NewValue(greaterThanOrEqual))
}

func (i IntPath) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i IntPath) IsNull() core.ComboExpression {
	return expression.NewOperator(i, operator.Null)
}

func (i IntPath) IsNotNull() core.ComboExpression {
	return expression.NewOperator(i, operator.NotNull)
}

func (i IntPath) In(values ...int) core.ComboExpression {
	return expression.NewOperator(i, operator.In, expression.NewValue(values))
}

func (i IntPath) InPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.In, values...)
}

func (i IntPath) NotIn(values ...int) core.ComboExpression {
	return expression.NewOperator(i, operator.NotIn, expression.NewValue(values))
}

func (i IntPath) NotInPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotIn, values...)
}

func (i IntPath) Between(first, second int) core.ComboExpression {
	return expression.NewOperator(i, operator.Between, expression.NewValue(first).And(expression.NewValue(second)))
}

func (i IntPath) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.Between, expression.NewOperator(first, operator.And, second))
}

func (i IntPath) NotBetween(first, second int) core.ComboExpression {
	return expression.NewOperator(i, operator.NotBetween, expression.NewValue(first).And(expression.NewValue(second)))
}

func (i IntPath) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotBetween, expression.NewOperator(first, operator.And, second))
}
