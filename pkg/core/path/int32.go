package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/operator"
)

func NewInt32PathWithAlias(e core.Table, name, alias string) Int32Path {
	return Int32Path{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt32Path(e core.Table, name string) Int32Path {
	return NewInt32PathWithAlias(e, name, "")
}

type Int32Path struct {
	entity core.Table
	name   string
	alias  string
}

func (i Int32Path) GetParent() core.Table {
	return i.entity
}

func (i Int32Path) GetName() string {
	return i.name
}

func (i Int32Path) GetAlias() string {
	return i.alias
}

func (i Int32Path) As(alias string) Int32Path {
	i.alias = alias
	return i
}

func (i Int32Path) GetSQL(d core.Dialect) (core.SQL, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int32Path) To(value int32) core.Set {
	return expression.NewSet(i, expression.NewValue(value))
}

func (i Int32Path) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(i, setExp)
}

func (i Int32Path) Eq(equalTo int32) core.ComboExpression {
	return expression.NewOperator(i, operator.Eq, expression.NewValue(equalTo))
}

func (i Int32Path) EqPath(equalTo core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.Eq, equalTo)
}

func (i Int32Path) NotEq(notEqualTo int32) core.ComboExpression {
	return expression.NewOperator(i, operator.NotEq, expression.NewValue(notEqualTo))
}

func (i Int32Path) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotEq, notEqualTo)
}

func (i Int32Path) LT(lessThan int32) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThan, expression.NewValue(lessThan))
}

func (i Int32Path) LTPath(lessThan core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThan, lessThan)
}

func (i Int32Path) LTOrEq(lessThanOrEqual int32) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThanOrEqual, expression.NewValue(lessThanOrEqual))
}

func (i Int32Path) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int32Path) GT(greaterThan int32) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThan, expression.NewValue(greaterThan))
}

func (i Int32Path) GTPath(greaterThan core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThan, greaterThan)
}

func (i Int32Path) GTOrEq(greaterThanOrEqual int32) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThanOrEqual, expression.NewValue(greaterThanOrEqual))
}

func (i Int32Path) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int32Path) IsNull() core.ComboExpression {
	return expression.NewOperator(i, operator.Null)
}

func (i Int32Path) IsNotNull() core.ComboExpression {
	return expression.NewOperator(i, operator.NotNull)
}

func (i Int32Path) In(values ...int32) core.ComboExpression {
	return expression.NewOperator(i, operator.In, expression.NewValue(values))
}

func (i Int32Path) InPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.In, values...)
}

func (i Int32Path) NotIn(values ...int32) core.ComboExpression {
	return expression.NewOperator(i, operator.NotIn, expression.NewValue(values))
}

func (i Int32Path) NotInPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotIn, values...)
}

func (i Int32Path) Between(first, second int32) core.ComboExpression {
	return expression.NewOperator(i, operator.Between, expression.NewValue(first).And(expression.NewValue(second)))
}

func (i Int32Path) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.Between, expression.NewOperator(first, operator.And, second))
}

func (i Int32Path) NotBetween(first, second int32) core.ComboExpression {
	return expression.NewOperator(i, operator.NotBetween, expression.NewValue(first).And(expression.NewValue(second)))
}

func (i Int32Path) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotBetween, expression.NewOperator(first, operator.And, second))
}
