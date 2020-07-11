package path

import (
	"time"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/set"
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

func (t Time) To(value time.Time) set.Set {
	return set.NewSet(t, expr.NewValue(value))
}

func (t Time) ToExpr(setExp lingo.Expression) set.Set {
	return set.NewSet(t, setExp)
}

func (t Time) Eq(equalTo time.Time) operator.Binary {
	return operator.NewBinary(t, operator.Eq, expr.NewValue(equalTo))
}

func (t Time) EqPath(equalTo lingo.Expression) operator.Binary {
	return operator.NewBinary(t, operator.Eq, equalTo)
}

func (t Time) NotEq(notEqualTo time.Time) operator.Binary {
	return operator.NewBinary(t, operator.NotEq, expr.NewValue(notEqualTo))
}

func (t Time) NotEqPath(notEqualTo lingo.Expression) operator.Binary {
	return operator.NewBinary(t, operator.NotEq, notEqualTo)
}

func (t Time) LT(lessThan time.Time) operator.Binary {
	return operator.NewBinary(t, operator.LessThan, expr.NewValue(lessThan))
}

func (t Time) LTPath(lessThan lingo.Expression) operator.Binary {
	return operator.NewBinary(t, operator.LessThan, lessThan)
}

func (t Time) LTOrEq(lessThanOrEqual time.Time) operator.Binary {
	return operator.NewBinary(t, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (t Time) LTOrEqPath(lessThanOrEqual lingo.Expression) operator.Binary {
	return operator.NewBinary(t, operator.LessThanOrEqual, lessThanOrEqual)
}

func (t Time) GT(greaterThan time.Time) operator.Binary {
	return operator.NewBinary(t, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (t Time) GTPath(greaterThan lingo.Expression) operator.Binary {
	return operator.NewBinary(t, operator.GreaterThan, greaterThan)
}

func (t Time) GTOrEq(greaterThanOrEqual time.Time) operator.Binary {
	return operator.NewBinary(t, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (t Time) GTOrEqPath(greaterThanOrEqual lingo.Expression) operator.Binary {
	return operator.NewBinary(t, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (t Time) IsNull() operator.Unary {
	return operator.NewUnary(t, operator.Null)
}

func (t Time) IsNotNull() operator.Unary {
	return operator.NewUnary(t, operator.NotNull)
}

func (t Time) In(values ...time.Time) operator.Binary {
	return operator.NewBinary(t, operator.In, expr.NewParens(expr.NewValue(values)))
}

func (t Time) InPaths(values ...lingo.Expression) operator.Binary {
	return operator.NewBinary(t, operator.In, expr.NewParens(expr.ToList(values)))
}

func (t Time) NotIn(values ...time.Time) operator.Binary {
	return operator.NewBinary(t, operator.NotIn, expr.NewParens(expr.NewValue(values)))
}

func (t Time) NotInPaths(values ...lingo.Expression) operator.Binary {
	return operator.NewBinary(t, operator.NotIn, expr.NewParens(expr.ToList(values)))
}

func (t Time) Between(first, second time.Time) operator.Binary {
	return operator.NewBinary(t, operator.Between, expr.NewParens(expr.NewValue(first).And(expr.NewValue(second))))
}

func (t Time) BetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NewBinary(t, operator.Between, expr.NewParens(operator.NewBinary(first, operator.And, second)))
}

func (t Time) NotBetween(first, second time.Time) operator.Binary {
	return operator.NewBinary(t, operator.NotBetween, expr.NewParens(expr.NewValue(first).And(expr.NewValue(second))))
}

func (t Time) NotBetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NewBinary(t, operator.NotBetween, expr.NewParens(operator.NewBinary(first, operator.And, second)))
}
