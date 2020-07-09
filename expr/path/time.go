package path

import (
	"time"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/sql"
)

func NewTimeWithAlias(e lingo.Table, name, alias string) Time {
	return Time{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewTime(e lingo.Table, name string) Time {
	return NewTimeWithAlias(e, name, "")
}

type Time struct {
	entity lingo.Table
	name   string
	alias  string
}

func (t Time) GetParent() lingo.Table {
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

func (t Time) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, t)
}

func (t Time) To(value time.Time) lingo.Set {
	return expr.NewSet(t, expr.NewValue(value))
}

func (t Time) ToExpr(setExp lingo.Expression) lingo.Set {
	return expr.NewSet(t, setExp)
}

func (t Time) Eq(equalTo time.Time) lingo.ComboExpression {
	return operator.NewOperator(t, operator.Eq, expr.NewValue(equalTo))
}

func (t Time) EqPath(equalTo lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(t, operator.Eq, equalTo)
}

func (t Time) NotEq(notEqualTo time.Time) lingo.ComboExpression {
	return operator.NewOperator(t, operator.NotEq, expr.NewValue(notEqualTo))
}

func (t Time) NotEqPath(notEqualTo lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(t, operator.NotEq, notEqualTo)
}

func (t Time) LT(lessThan time.Time) lingo.ComboExpression {
	return operator.NewOperator(t, operator.LessThan, expr.NewValue(lessThan))
}

func (t Time) LTPath(lessThan lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(t, operator.LessThan, lessThan)
}

func (t Time) LTOrEq(lessThanOrEqual time.Time) lingo.ComboExpression {
	return operator.NewOperator(t, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (t Time) LTOrEqPath(lessThanOrEqual lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(t, operator.LessThanOrEqual, lessThanOrEqual)
}

func (t Time) GT(greaterThan time.Time) lingo.ComboExpression {
	return operator.NewOperator(t, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (t Time) GTPath(greaterThan lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(t, operator.GreaterThan, greaterThan)
}

func (t Time) GTOrEq(greaterThanOrEqual time.Time) lingo.ComboExpression {
	return operator.NewOperator(t, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (t Time) GTOrEqPath(greaterThanOrEqual lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(t, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (t Time) IsNull() lingo.ComboExpression {
	return operator.NewOperator(t, operator.Null)
}

func (t Time) IsNotNull() lingo.ComboExpression {
	return operator.NewOperator(t, operator.NotNull)
}

func (t Time) In(values ...time.Time) lingo.ComboExpression {
	return operator.NewOperator(t, operator.In, expr.NewValue(values))
}

func (t Time) InPaths(values ...lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(t, operator.In, values...)
}

func (t Time) NotIn(values ...time.Time) lingo.ComboExpression {
	return operator.NewOperator(t, operator.NotIn, expr.NewValue(values))
}

func (t Time) NotInPaths(values ...lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(t, operator.NotIn, values...)
}

func (t Time) Between(first, second time.Time) lingo.ComboExpression {
	return operator.NewOperator(t, operator.Between, expr.NewValue(first).And(expr.NewValue(second)))
}

func (t Time) BetweenPaths(first, second lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(t, operator.Between, operator.NewOperator(first, operator.And, second))
}

func (t Time) NotBetween(first, second time.Time) lingo.ComboExpression {
	return operator.NewOperator(t, operator.NotBetween, expr.NewValue(first).And(expr.NewValue(second)))
}

func (t Time) NotBetweenPaths(first, second lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(t, operator.NotBetween, operator.NewOperator(first, operator.And, second))
}
