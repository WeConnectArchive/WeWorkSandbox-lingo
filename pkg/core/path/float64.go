package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/operator"
)

func NewFloat64PathWithAlias(e core.Table, name, alias string) Float64Path {
	return Float64Path{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewFloat64Path(e core.Table, name string) Float64Path {
	return NewFloat64PathWithAlias(e, name, "")
}

type Float64Path struct {
	entity core.Table
	name   string
	alias  string
}

func (f Float64Path) GetParent() core.Table {
	return f.entity
}

func (f Float64Path) GetName() string {
	return f.name
}

func (f Float64Path) GetAlias() string {
	return f.alias
}

func (f Float64Path) As(alias string) Float64Path {
	f.alias = alias
	return f
}

func (f Float64Path) GetSQL(d core.Dialect) (core.SQL, error) {
	return ExpandColumnWithDialect(d, f)
}

func (f Float64Path) To(value float64) core.Set {
	return expression.NewSet(f, expression.NewValue(value))
}

func (f Float64Path) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(f, setExp)
}

func (f Float64Path) Eq(equalTo float64) core.ComboExpression {
	return expression.NewOperator(f, operator.Eq, expression.NewValue(equalTo))
}

func (f Float64Path) EqPath(equalTo core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.Eq, equalTo)
}

func (f Float64Path) NotEq(notEqualTo float64) core.ComboExpression {
	return expression.NewOperator(f, operator.NotEq, expression.NewValue(notEqualTo))
}

func (f Float64Path) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.NotEq, notEqualTo)
}

func (f Float64Path) LT(lessThan float64) core.ComboExpression {
	return expression.NewOperator(f, operator.LessThan, expression.NewValue(lessThan))
}

func (f Float64Path) LTPath(lessThan core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.LessThan, lessThan)
}

func (f Float64Path) LTOrEq(lessThanOrEqual float64) core.ComboExpression {
	return expression.NewOperator(f, operator.LessThanOrEqual, expression.NewValue(lessThanOrEqual))
}

func (f Float64Path) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.LessThanOrEqual, lessThanOrEqual)
}

func (f Float64Path) GT(greaterThan float64) core.ComboExpression {
	return expression.NewOperator(f, operator.GreaterThan, expression.NewValue(greaterThan))
}

func (f Float64Path) GTPath(greaterThan core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.GreaterThan, greaterThan)
}

func (f Float64Path) GTOrEq(greaterThanOrEqual float64) core.ComboExpression {
	return expression.NewOperator(f, operator.GreaterThanOrEqual, expression.NewValue(greaterThanOrEqual))
}

func (f Float64Path) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (f Float64Path) IsNull() core.ComboExpression {
	return expression.NewOperator(f, operator.Null)
}

func (f Float64Path) IsNotNull() core.ComboExpression {
	return expression.NewOperator(f, operator.NotNull)
}

func (f Float64Path) In(values ...float64) core.ComboExpression {
	return expression.NewOperator(f, operator.In, expression.NewValue(values))
}

func (f Float64Path) InPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.In, values...)
}

func (f Float64Path) NotIn(values ...float64) core.ComboExpression {
	return expression.NewOperator(f, operator.NotIn, expression.NewValue(values))
}

func (f Float64Path) NotInPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.NotIn, values...)
}

func (f Float64Path) Between(first, second float64) core.ComboExpression {
	return expression.NewOperator(f, operator.Between, expression.NewValue(first).And(expression.NewValue(second)))
}

func (f Float64Path) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.Between, expression.NewOperator(first, operator.And, second))
}

func (f Float64Path) NotBetween(first, second float64) core.ComboExpression {
	return expression.NewOperator(f, operator.NotBetween, expression.NewValue(first).And(expression.NewValue(second)))
}

func (f Float64Path) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.NotBetween, expression.NewOperator(first, operator.And, second))
}
