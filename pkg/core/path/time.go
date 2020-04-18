package path

import (
	"time"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/operator"
)

func NewTimePathWithAlias(e core.Table, name, alias string) TimePath {
	return TimePath{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewTimePath(e core.Table, name string) TimePath {
	return NewTimePathWithAlias(e, name, "")
}

type TimePath struct {
	entity core.Table
	name   string
	alias  string
}

func (t TimePath) GetParent() core.Table {
	return t.entity
}

func (t TimePath) GetName() string {
	return t.name
}

func (t TimePath) GetAlias() string {
	return t.alias
}

func (t TimePath) As(alias string) TimePath {
	t.alias = alias
	return t
}

func (t TimePath) GetSQL(d core.Dialect) (core.SQL, error) {
	return ExpandColumnWithDialect(d, t)
}

func (t TimePath) To(value time.Time) core.Set {
	return expression.NewSet(t, expression.NewValue(value))
}

func (t TimePath) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(t, setExp)
}

func (t TimePath) Eq(equalTo time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.Eq, expression.NewValue(equalTo))
}

func (t TimePath) EqPath(equalTo core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.Eq, equalTo)
}

func (t TimePath) NotEq(notEqualTo time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.NotEq, expression.NewValue(notEqualTo))
}

func (t TimePath) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.NotEq, notEqualTo)
}

func (t TimePath) LT(lessThan time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.LessThan, expression.NewValue(lessThan))
}

func (t TimePath) LTPath(lessThan core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.LessThan, lessThan)
}

func (t TimePath) LTOrEq(lessThanOrEqual time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.LessThanOrEqual, expression.NewValue(lessThanOrEqual))
}

func (t TimePath) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.LessThanOrEqual, lessThanOrEqual)
}

func (t TimePath) GT(greaterThan time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.GreaterThan, expression.NewValue(greaterThan))
}

func (t TimePath) GTPath(greaterThan core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.GreaterThan, greaterThan)
}

func (t TimePath) GTOrEq(greaterThanOrEqual time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.GreaterThanOrEqual, expression.NewValue(greaterThanOrEqual))
}

func (t TimePath) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (t TimePath) IsNull() core.ComboExpression {
	return expression.NewOperator(t, operator.Null)
}

func (t TimePath) IsNotNull() core.ComboExpression {
	return expression.NewOperator(t, operator.NotNull)
}

func (t TimePath) In(values ...time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.In, expression.NewValue(values))
}

func (t TimePath) InPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.In, values...)
}

func (t TimePath) NotIn(values ...time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.NotIn, expression.NewValue(values))
}

func (t TimePath) NotInPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.NotIn, values...)
}

func (t TimePath) Between(first, second time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.Between, expression.NewValue(first).And(expression.NewValue(second)))
}

func (t TimePath) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.Between, expression.NewOperator(first, operator.And, second))
}

func (t TimePath) NotBetween(first, second time.Time) core.ComboExpression {
	return expression.NewOperator(t, operator.NotBetween, expression.NewValue(first).And(expression.NewValue(second)))
}

func (t TimePath) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(t, operator.NotBetween, expression.NewOperator(first, operator.And, second))
}
