package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expr"
	"github.com/weworksandbox/lingo/pkg/core/expr/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func NewFloat64WithAlias(e core.Table, name, alias string) Float64 {
	return Float64{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewFloat64(e core.Table, name string) Float64 {
	return NewFloat64WithAlias(e, name, "")
}

type Float64 struct {
	entity core.Table
	name   string
	alias  string
}

func (f Float64) GetParent() core.Table {
	return f.entity
}

func (f Float64) GetName() string {
	return f.name
}

func (f Float64) GetAlias() string {
	return f.alias
}

func (f Float64) As(alias string) Float64 {
	f.alias = alias
	return f
}

func (f Float64) ToSQL(d core.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, f)
}

func (f Float64) To(value float64) core.Set {
	return expr.NewSet(f, expr.NewValue(value))
}

func (f Float64) ToExpr(setExp core.Expression) core.Set {
	return expr.NewSet(f, setExp)
}

func (f Float64) Eq(equalTo float64) core.ComboExpression {
	return operator.NewOperator(f, operator.Eq, expr.NewValue(equalTo))
}

func (f Float64) EqPath(equalTo core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.Eq, equalTo)
}

func (f Float64) NotEq(notEqualTo float64) core.ComboExpression {
	return operator.NewOperator(f, operator.NotEq, expr.NewValue(notEqualTo))
}

func (f Float64) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.NotEq, notEqualTo)
}

func (f Float64) LT(lessThan float64) core.ComboExpression {
	return operator.NewOperator(f, operator.LessThan, expr.NewValue(lessThan))
}

func (f Float64) LTPath(lessThan core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.LessThan, lessThan)
}

func (f Float64) LTOrEq(lessThanOrEqual float64) core.ComboExpression {
	return operator.NewOperator(f, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (f Float64) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.LessThanOrEqual, lessThanOrEqual)
}

func (f Float64) GT(greaterThan float64) core.ComboExpression {
	return operator.NewOperator(f, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (f Float64) GTPath(greaterThan core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.GreaterThan, greaterThan)
}

func (f Float64) GTOrEq(greaterThanOrEqual float64) core.ComboExpression {
	return operator.NewOperator(f, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (f Float64) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (f Float64) IsNull() core.ComboExpression {
	return operator.NewOperator(f, operator.Null)
}

func (f Float64) IsNotNull() core.ComboExpression {
	return operator.NewOperator(f, operator.NotNull)
}

func (f Float64) In(values ...float64) core.ComboExpression {
	return operator.NewOperator(f, operator.In, expr.NewValue(values))
}

func (f Float64) InPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.In, values...)
}

func (f Float64) NotIn(values ...float64) core.ComboExpression {
	return operator.NewOperator(f, operator.NotIn, expr.NewValue(values))
}

func (f Float64) NotInPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.NotIn, values...)
}

func (f Float64) Between(first, second float64) core.ComboExpression {
	return operator.NewOperator(f, operator.Between, expr.NewValue(first).And(expr.NewValue(second)))
}

func (f Float64) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.Between, operator.NewOperator(first, operator.And, second))
}

func (f Float64) NotBetween(first, second float64) core.ComboExpression {
	return operator.NewOperator(f, operator.NotBetween, expr.NewValue(first).And(expr.NewValue(second)))
}

func (f Float64) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.NotBetween, operator.NewOperator(first, operator.And, second))
}
