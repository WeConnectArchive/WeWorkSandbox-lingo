package path

import (
	"time"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expr"
	"github.com/weworksandbox/lingo/pkg/core/expr/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func NewTimeWithAlias(e core.Table, name, alias string) Time {
	return Time{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewTime(e core.Table, name string) Time {
	return NewTimeWithAlias(e, name, "")
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

func (t Time) ToSQL(d core.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, t)
}

func (t Time) To(value time.Time) core.Set {
	return expr.NewSet(t, expr.NewValue(value))
}

func (t Time) ToExpression(setExp core.Expression) core.Set {
	return expr.NewSet(t, setExp)
}

func (t Time) Eq(equalTo time.Time) core.ComboExpression {
	return operator.NewOperator(t, operator.Eq, expr.NewValue(equalTo))
}

func (t Time) EqPath(equalTo core.Expression) core.ComboExpression {
	return operator.NewOperator(t, operator.Eq, equalTo)
}

func (t Time) NotEq(notEqualTo time.Time) core.ComboExpression {
	return operator.NewOperator(t, operator.NotEq, expr.NewValue(notEqualTo))
}

func (t Time) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return operator.NewOperator(t, operator.NotEq, notEqualTo)
}

func (t Time) LT(lessThan time.Time) core.ComboExpression {
	return operator.NewOperator(t, operator.LessThan, expr.NewValue(lessThan))
}

func (t Time) LTPath(lessThan core.Expression) core.ComboExpression {
	return operator.NewOperator(t, operator.LessThan, lessThan)
}

func (t Time) LTOrEq(lessThanOrEqual time.Time) core.ComboExpression {
	return operator.NewOperator(t, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (t Time) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return operator.NewOperator(t, operator.LessThanOrEqual, lessThanOrEqual)
}

func (t Time) GT(greaterThan time.Time) core.ComboExpression {
	return operator.NewOperator(t, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (t Time) GTPath(greaterThan core.Expression) core.ComboExpression {
	return operator.NewOperator(t, operator.GreaterThan, greaterThan)
}

func (t Time) GTOrEq(greaterThanOrEqual time.Time) core.ComboExpression {
	return operator.NewOperator(t, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (t Time) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return operator.NewOperator(t, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (t Time) IsNull() core.ComboExpression {
	return operator.NewOperator(t, operator.Null)
}

func (t Time) IsNotNull() core.ComboExpression {
	return operator.NewOperator(t, operator.NotNull)
}

func (t Time) In(values ...time.Time) core.ComboExpression {
	return operator.NewOperator(t, operator.In, expr.NewValue(values))
}

func (t Time) InPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(t, operator.In, values...)
}

func (t Time) NotIn(values ...time.Time) core.ComboExpression {
	return operator.NewOperator(t, operator.NotIn, expr.NewValue(values))
}

func (t Time) NotInPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(t, operator.NotIn, values...)
}

func (t Time) Between(first, second time.Time) core.ComboExpression {
	return operator.NewOperator(t, operator.Between, expr.NewValue(first).And(expr.NewValue(second)))
}

func (t Time) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return operator.NewOperator(t, operator.Between, operator.NewOperator(first, operator.And, second))
}

func (t Time) NotBetween(first, second time.Time) core.ComboExpression {
	return operator.NewOperator(t, operator.NotBetween, expr.NewValue(first).And(expr.NewValue(second)))
}

func (t Time) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return operator.NewOperator(t, operator.NotBetween, operator.NewOperator(first, operator.And, second))
}
