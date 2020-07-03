package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/expression/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func NewInt8WithAlias(e core.Table, name, alias string) Int8 {
	return Int8{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt8(e core.Table, name string) Int8 {
	return NewInt8WithAlias(e, name, "")
}

type Int8 struct {
	entity core.Table
	name   string
	alias  string
}

func (i Int8) GetParent() core.Table {
	return i.entity
}

func (i Int8) GetName() string {
	return i.name
}

func (i Int8) GetAlias() string {
	return i.alias
}

func (i Int8) As(alias string) Int8 {
	i.alias = alias
	return i
}

func (i Int8) ToSQL(d core.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int8) To(value int8) core.Set {
	return expression.NewSet(i, expression.NewValue(value))
}

func (i Int8) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(i, setExp)
}

func (i Int8) Eq(equalTo int8) core.ComboExpression {
	return expression.NewOperator(i, operator.Eq, expression.NewValue(equalTo))
}

func (i Int8) EqPath(equalTo core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.Eq, equalTo)
}

func (i Int8) NotEq(notEqualTo int8) core.ComboExpression {
	return expression.NewOperator(i, operator.NotEq, expression.NewValue(notEqualTo))
}

func (i Int8) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotEq, notEqualTo)
}

func (i Int8) LT(lessThan int8) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThan, expression.NewValue(lessThan))
}

func (i Int8) LTPath(lessThan core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThan, lessThan)
}

func (i Int8) LTOrEq(lessThanOrEqual int8) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThanOrEqual, expression.NewValue(lessThanOrEqual))
}

func (i Int8) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int8) GT(greaterThan int8) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThan, expression.NewValue(greaterThan))
}

func (i Int8) GTPath(greaterThan core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThan, greaterThan)
}

func (i Int8) GTOrEq(greaterThanOrEqual int8) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThanOrEqual, expression.NewValue(greaterThanOrEqual))
}

func (i Int8) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int8) IsNull() core.ComboExpression {
	return expression.NewOperator(i, operator.Null)
}

func (i Int8) IsNotNull() core.ComboExpression {
	return expression.NewOperator(i, operator.NotNull)
}

func (i Int8) In(values ...int8) core.ComboExpression {
	return expression.NewOperator(i, operator.In, expression.NewValue(values))
}

func (i Int8) InPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.In, values...)
}

func (i Int8) NotIn(values ...int8) core.ComboExpression {
	return expression.NewOperator(i, operator.NotIn, expression.NewValue(values))
}

func (i Int8) NotInPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotIn, values...)
}

func (i Int8) Between(first, second int8) core.ComboExpression {
	return expression.NewOperator(i, operator.Between, expression.NewValue(first).And(expression.NewValue(second)))
}

func (i Int8) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.Between, expression.NewOperator(first, operator.And, second))
}

func (i Int8) NotBetween(first, second int8) core.ComboExpression {
	return expression.NewOperator(i, operator.NotBetween, expression.NewValue(first).And(expression.NewValue(second)))
}

func (i Int8) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(i, operator.NotBetween, expression.NewOperator(first, operator.And, second))
}
