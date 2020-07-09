package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expr"
	"github.com/weworksandbox/lingo/pkg/core/expr/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func NewFloat32WithAlias(e core.Table, name, alias string) Float32 {
	return Float32{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewFloat32(e core.Table, name string) Float32 {
	return NewFloat32WithAlias(e, name, "")
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

func (f Float32) ToSQL(d core.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, f)
}

func (f Float32) To(value float32) core.Set {
	return expr.NewSet(f, expr.NewValue(value))
}

func (f Float32) ToExpr(setExp core.Expression) core.Set {
	return expr.NewSet(f, setExp)
}

func (f Float32) Eq(equalTo float32) core.ComboExpression {
	return operator.NewOperator(f, operator.Eq, expr.NewValue(equalTo))
}

func (f Float32) EqPath(equalTo core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.Eq, equalTo)
}

func (f Float32) NotEq(notEqualTo float32) core.ComboExpression {
	return operator.NewOperator(f, operator.NotEq, expr.NewValue(notEqualTo))
}

func (f Float32) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.NotEq, notEqualTo)
}

func (f Float32) LT(lessThan float32) core.ComboExpression {
	return operator.NewOperator(f, operator.LessThan, expr.NewValue(lessThan))
}

func (f Float32) LTPath(lessThan core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.LessThan, lessThan)
}

func (f Float32) LTOrEq(lessThanOrEqual float32) core.ComboExpression {
	return operator.NewOperator(f, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (f Float32) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.LessThanOrEqual, lessThanOrEqual)
}

func (f Float32) GT(greaterThan float32) core.ComboExpression {
	return operator.NewOperator(f, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (f Float32) GTPath(greaterThan core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.GreaterThan, greaterThan)
}

func (f Float32) GTOrEq(greaterThanOrEqual float32) core.ComboExpression {
	return operator.NewOperator(f, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (f Float32) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (f Float32) IsNull() core.ComboExpression {
	return operator.NewOperator(f, operator.Null)
}

func (f Float32) IsNotNull() core.ComboExpression {
	return operator.NewOperator(f, operator.NotNull)
}

func (f Float32) In(values ...float32) core.ComboExpression {
	return operator.NewOperator(f, operator.In, expr.NewValue(values))
}

func (f Float32) InPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.In, values...)
}

func (f Float32) NotIn(values ...float32) core.ComboExpression {
	return operator.NewOperator(f, operator.NotIn, expr.NewValue(values))
}

func (f Float32) NotInPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.NotIn, values...)
}

func (f Float32) Between(first, second float32) core.ComboExpression {
	return operator.NewOperator(f, operator.Between, expr.NewValue(first).And(expr.NewValue(second)))
}

func (f Float32) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.Between, operator.NewOperator(first, operator.And, second))
}

func (f Float32) NotBetween(first, second float32) core.ComboExpression {
	return operator.NewOperator(f, operator.NotBetween, expr.NewValue(first).And(expr.NewValue(second)))
}

func (f Float32) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return operator.NewOperator(f, operator.NotBetween, operator.NewOperator(first, operator.And, second))
}
