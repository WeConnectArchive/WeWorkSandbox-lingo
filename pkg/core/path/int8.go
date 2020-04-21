package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/operator"
)

func NewInt8PathWithAlias(e core.Table, name, alias string) Int8Path {
	return Int8Path{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt8Path(e core.Table, name string) Int8Path {
	return NewInt8PathWithAlias(e, name, "")
}

type Int8Path struct {
	entity core.Table
	name   string
	alias  string
}

func (i Int8Path) GetParent() core.Table {
	return i.entity
}

func (i Int8Path) GetName() string {
	return i.name
}

func (i Int8Path) GetAlias() string {
	return i.alias
}

func (i Int8Path) As(alias string) Int8Path {
	i.alias = alias
	return i
}

func (i Int8Path) GetSQL(d core.Dialect) (core.SQL, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int8Path) To(value int8) core.Set {
	return expression.NewSet(i, expression.NewValue(value))
}

func (i Int8Path) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(i, setExp)
}

func (i Int8Path) Eq(equalTo int8) core.ComboExpression {
	return expression.NewOperator(i, operator.Eq, expression.NewValue(equalTo))
}

func (i Int8Path) EqPath(equalTo core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.Eq, equalTo)
}

func (i Int8Path) NotEq(notEqualTo int8) core.ComboExpression {
	return expression.NewOperator(i, operator.NotEq, expression.NewValue(notEqualTo))
}

func (i Int8Path) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotEq, notEqualTo)
}

func (i Int8Path) LT(lessThan int8) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThan, expression.NewValue(lessThan))
}

func (i Int8Path) LTPath(lessThan core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThan, lessThan)
}

func (i Int8Path) LTOrEq(lessThanOrEqual int8) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThanOrEqual, expression.NewValue(lessThanOrEqual))
}

func (i Int8Path) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int8Path) GT(greaterThan int8) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThan, expression.NewValue(greaterThan))
}

func (i Int8Path) GTPath(greaterThan core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThan, greaterThan)
}

func (i Int8Path) GTOrEq(greaterThanOrEqual int8) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThanOrEqual, expression.NewValue(greaterThanOrEqual))
}

func (i Int8Path) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int8Path) IsNull() core.ComboExpression {
	return expression.NewOperator(i, operator.Null)
}

func (i Int8Path) IsNotNull() core.ComboExpression {
	return expression.NewOperator(i, operator.NotNull)
}

func (i Int8Path) In(values ...int8) core.ComboExpression {
	return expression.NewOperator(i, operator.In, expression.NewValue(values))
}

func (i Int8Path) InPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.In, values...)
}

func (i Int8Path) NotIn(values ...int8) core.ComboExpression {
	return expression.NewOperator(i, operator.NotIn, expression.NewValue(values))
}

func (i Int8Path) NotInPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotIn, values...)
}

func (i Int8Path) Between(first, second int8) core.ComboExpression {
	return expression.NewOperator(i, operator.Between, expression.NewValue(first).And(expression.NewValue(second)))
}

func (i Int8Path) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.Between, expression.NewOperator(first, operator.And, second))
}

func (i Int8Path) NotBetween(first, second int8) core.ComboExpression {
	return expression.NewOperator(i, operator.NotBetween, expression.NewValue(first).And(expression.NewValue(second)))
}

func (i Int8Path) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotBetween, expression.NewOperator(first, operator.And, second))
}
