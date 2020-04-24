package path

import (
	"time"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/operator"
)

func NewTimePathWithAlias(e core.Table, name, alias string) Time {
	return Time{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewTimePath(e core.Table, name string) Time {
	return NewTimePathWithAlias(e, name, "")
}

type Time struct {
	entity core.Table
	name   string
	alias  string
}

func (t Time) GetParent() core.Table {
	return t.entity
}

func (t Time) GetName() string {
	return t.name
}

func (t Time) GetAlias() string {
	return t.alias
}

func (t Time) As(alias string) Time {
	t.alias = alias
	return t
}

func (t Time) GetSQL(d core.Dialect) (core.SQL, error) {
	return ExpandColumnWithDialect(d, t)
}

func (t Time) To(value time.Time) core.Set {
	return expression.NewSet(t, expression.NewValue(value))
}

func (t Time) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(t, setExp)
}

func (t Time) Eq(equalTo time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.Eq, expression.NewValue(equalTo))
}

func (t Time) EqPath(equalTo core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.Eq, equalTo)
}

func (t Time) NotEq(notEqualTo time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.NotEq, expression.NewValue(notEqualTo))
}

func (t Time) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.NotEq, notEqualTo)
}

func (t Time) LT(lessThan time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.LessThan, expression.NewValue(lessThan))
}

func (t Time) LTPath(lessThan core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.LessThan, lessThan)
}

func (t Time) LTOrEq(lessThanOrEqual time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.LessThanOrEqual, expression.NewValue(lessThanOrEqual))
}

func (t Time) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.LessThanOrEqual, lessThanOrEqual)
}

func (t Time) GT(greaterThan time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.GreaterThan, expression.NewValue(greaterThan))
}

func (t Time) GTPath(greaterThan core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.GreaterThan, greaterThan)
}

func (t Time) GTOrEq(greaterThanOrEqual time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.GreaterThanOrEqual, expression.NewValue(greaterThanOrEqual))
}

func (t Time) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (t Time) IsNull() core.ComboExpression {
	return expression.NewOperator(t, operator.Null)
}

func (t Time) IsNotNull() core.ComboExpression {
	return expression.NewOperator(t, operator.NotNull)
}

func (t Time) In(values ...time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.In, expression.NewValue(values))
}

func (t Time) InPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.In, values...)
}

func (t Time) NotIn(values ...time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.NotIn, expression.NewValue(values))
}

func (t Time) NotInPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.NotIn, values...)
}

func (t Time) Between(first, second time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.Between, expression.NewValue(first).And(expression.NewValue(second)))
}

func (t Time) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.Between, expression.NewOperator(first, operator.And, second))
}

func (t Time) NotBetween(first, second time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.NotBetween, expression.NewValue(first).And(expression.NewValue(second)))
}

func (t Time) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.NotBetween, expression.NewOperator(first, operator.And, second))
}
