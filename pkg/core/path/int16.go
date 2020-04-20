package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/operator"
)

func NewInt16PathWithAlias(e core.Table, name, alias string) Int16Path {
	return Int16Path{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt16Path(e core.Table, name string) Int16Path {
	return NewInt16PathWithAlias(e, name, "")
}

type Int16Path struct {
	entity core.Table
	name   string
	alias  string
}

func (i Int16Path) GetParent() core.Table {
	return i.entity
}

func (i Int16Path) GetName() string {
	return i.name
}

func (i Int16Path) GetAlias() string {
	return i.alias
}

func (i Int16Path) As(alias string) Int16Path {
	i.alias = alias
	return i
}

func (i Int16Path) GetSQL(d core.Dialect) (core.SQL, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int16Path) To(value int16) core.Set {
	return expression.NewSet(i, expression.NewValue(value))
}

func (i Int16Path) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(i, setExp)
}

func (i Int16Path) Eq(equalTo int16) core.ComboExpression {
	return expression.NewOperator(i, operator.Eq, expression.NewValue(equalTo))
}

func (i Int16Path) EqPath(equalTo core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.Eq, equalTo)
}

func (i Int16Path) NotEq(notEqualTo int16) core.ComboExpression {
	return expression.NewOperator(i, operator.NotEq, expression.NewValue(notEqualTo))
}

func (i Int16Path) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotEq, notEqualTo)
}

func (i Int16Path) LT(lessThan int16) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThan, expression.NewValue(lessThan))
}

func (i Int16Path) LTPath(lessThan core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThan, lessThan)
}

func (i Int16Path) LTOrEq(lessThanOrEqual int16) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThanOrEqual, expression.NewValue(lessThanOrEqual))
}

func (i Int16Path) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int16Path) GT(greaterThan int16) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThan, expression.NewValue(greaterThan))
}

func (i Int16Path) GTPath(greaterThan core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThan, greaterThan)
}

func (i Int16Path) GTOrEq(greaterThanOrEqual int16) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThanOrEqual, expression.NewValue(greaterThanOrEqual))
}

func (i Int16Path) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int16Path) IsNull() core.ComboExpression {
	return expression.NewOperator(i, operator.Null)
}

func (i Int16Path) IsNotNull() core.ComboExpression {
	return expression.NewOperator(i, operator.NotNull)
}

func (i Int16Path) In(values ...int16) core.ComboExpression {
	return expression.NewOperator(i, operator.In, expression.NewValue(values))
}

func (i Int16Path) InPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.In, values...)
}

func (i Int16Path) NotIn(values ...int16) core.ComboExpression {
	return expression.NewOperator(i, operator.NotIn, expression.NewValue(values))
}

func (i Int16Path) NotInPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotIn, values...)
}

func (i Int16Path) Between(first, second int16) core.ComboExpression {
	return expression.NewOperator(i, operator.Between, expression.NewValue(first).And(expression.NewValue(second)))
}

func (i Int16Path) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.Between, expression.NewOperator(first, operator.And, second))
}

func (i Int16Path) NotBetween(first, second int16) core.ComboExpression {
	return expression.NewOperator(i, operator.NotBetween, expression.NewValue(first).And(expression.NewValue(second)))
}

func (i Int16Path) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotBetween, expression.NewOperator(first, operator.And, second))
}
