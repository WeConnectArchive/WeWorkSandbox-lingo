package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/operator"
)

func NewFloat32PathWithAlias(e core.Table, name, alias string) Float32 {
	return Float32{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewFloat32Path(e core.Table, name string) Float32 {
	return NewFloat32PathWithAlias(e, name, "")
}

type Float32 struct {
	entity core.Table
	name   string
	alias  string
}

func (f Float32) GetParent() core.Table {
	return f.entity
}

func (f Float32) GetName() string {
	return f.name
}

func (f Float32) GetAlias() string {
	return f.alias
}

func (f Float32) As(alias string) Float32 {
	f.alias = alias
	return f
}

func (f Float32) GetSQL(d core.Dialect) (core.SQL, error) {
	return ExpandColumnWithDialect(d, f)
}

func (f Float32) To(value float32) core.Set {
	return expression.NewSet(f, expression.NewValue(value))
}

func (f Float32) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(f, setExp)
}

func (f Float32) Eq(equalTo float32) core.ComboExpression {
	return expression.NewOperator(f, operator.Eq, expression.NewValue(equalTo))
}

func (f Float32) EqPath(equalTo core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.Eq, equalTo)
}

func (f Float32) NotEq(notEqualTo float32) core.ComboExpression {
	return expression.NewOperator(f, operator.NotEq, expression.NewValue(notEqualTo))
}

func (f Float32) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.NotEq, notEqualTo)
}

func (f Float32) LT(lessThan float32) core.ComboExpression {
	return expression.NewOperator(f, operator.LessThan, expression.NewValue(lessThan))
}

func (f Float32) LTPath(lessThan core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.LessThan, lessThan)
}

func (f Float32) LTOrEq(lessThanOrEqual float32) core.ComboExpression {
	return expression.NewOperator(f, operator.LessThanOrEqual, expression.NewValue(lessThanOrEqual))
}

func (f Float32) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.LessThanOrEqual, lessThanOrEqual)
}

func (f Float32) GT(greaterThan float32) core.ComboExpression {
	return expression.NewOperator(f, operator.GreaterThan, expression.NewValue(greaterThan))
}

func (f Float32) GTPath(greaterThan core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.GreaterThan, greaterThan)
}

func (f Float32) GTOrEq(greaterThanOrEqual float32) core.ComboExpression {
	return expression.NewOperator(f, operator.GreaterThanOrEqual, expression.NewValue(greaterThanOrEqual))
}

func (f Float32) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (f Float32) IsNull() core.ComboExpression {
	return expression.NewOperator(f, operator.Null)
}

func (f Float32) IsNotNull() core.ComboExpression {
	return expression.NewOperator(f, operator.NotNull)
}

func (f Float32) In(values ...float32) core.ComboExpression {
	return expression.NewOperator(f, operator.In, expression.NewValue(values))
}

func (f Float32) InPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.In, values...)
}

func (f Float32) NotIn(values ...float32) core.ComboExpression {
	return expression.NewOperator(f, operator.NotIn, expression.NewValue(values))
}

func (f Float32) NotInPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.NotIn, values...)
}

func (f Float32) Between(first, second float32) core.ComboExpression {
	return expression.NewOperator(f, operator.Between, expression.NewValue(first).And(expression.NewValue(second)))
}

func (f Float32) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.Between, expression.NewOperator(first, operator.And, second))
}

func (f Float32) NotBetween(first, second float32) core.ComboExpression {
	return expression.NewOperator(f, operator.NotBetween, expression.NewValue(first).And(expression.NewValue(second)))
}

func (f Float32) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(f, operator.NotBetween, expression.NewOperator(first, operator.And, second))
}
