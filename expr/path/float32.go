package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/set"
	"github.com/weworksandbox/lingo/sql"
)

func NewFloat32WithAlias(e lingo.Table, name, alias string) Float32 {
	return Float32{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewFloat32(e lingo.Table, name string) Float32 {
	return NewFloat32WithAlias(e, name, "")
}

type Float32 struct {
	entity lingo.Table
	name   string
	alias  string
}

func (f Float32) GetParent() lingo.Table {
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

func (f Float32) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, f)
}

func (f Float32) To(value float32) set.Set {
	return set.NewSet(f, expr.NewValue(value))
}

func (f Float32) ToExpr(setExp lingo.Expression) set.Set {
	return set.NewSet(f, setExp)
}

func (f Float32) Eq(equalTo float32) operator.Operator {
	return operator.NewOperator(f, operator.Eq, expr.NewValue(equalTo))
}

func (f Float32) EqPath(equalTo lingo.Expression) operator.Operator {
	return operator.NewOperator(f, operator.Eq, equalTo)
}

func (f Float32) NotEq(notEqualTo float32) operator.Operator {
	return operator.NewOperator(f, operator.NotEq, expr.NewValue(notEqualTo))
}

func (f Float32) NotEqPath(notEqualTo lingo.Expression) operator.Operator {
	return operator.NewOperator(f, operator.NotEq, notEqualTo)
}

func (f Float32) LT(lessThan float32) operator.Operator {
	return operator.NewOperator(f, operator.LessThan, expr.NewValue(lessThan))
}

func (f Float32) LTPath(lessThan lingo.Expression) operator.Operator {
	return operator.NewOperator(f, operator.LessThan, lessThan)
}

func (f Float32) LTOrEq(lessThanOrEqual float32) operator.Operator {
	return operator.NewOperator(f, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (f Float32) LTOrEqPath(lessThanOrEqual lingo.Expression) operator.Operator {
	return operator.NewOperator(f, operator.LessThanOrEqual, lessThanOrEqual)
}

func (f Float32) GT(greaterThan float32) operator.Operator {
	return operator.NewOperator(f, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (f Float32) GTPath(greaterThan lingo.Expression) operator.Operator {
	return operator.NewOperator(f, operator.GreaterThan, greaterThan)
}

func (f Float32) GTOrEq(greaterThanOrEqual float32) operator.Operator {
	return operator.NewOperator(f, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (f Float32) GTOrEqPath(greaterThanOrEqual lingo.Expression) operator.Operator {
	return operator.NewOperator(f, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (f Float32) IsNull() operator.Operator {
	return operator.NewOperator(f, operator.Null)
}

func (f Float32) IsNotNull() operator.Operator {
	return operator.NewOperator(f, operator.NotNull)
}

func (f Float32) In(values ...float32) operator.Operator {
	return operator.NewOperator(f, operator.In, expr.NewValue(values))
}

func (f Float32) InPaths(values ...lingo.Expression) operator.Operator {
	return operator.NewOperator(f, operator.In, values...)
}

func (f Float32) NotIn(values ...float32) operator.Operator {
	return operator.NewOperator(f, operator.NotIn, expr.NewValue(values))
}

func (f Float32) NotInPaths(values ...lingo.Expression) operator.Operator {
	return operator.NewOperator(f, operator.NotIn, values...)
}

func (f Float32) Between(first, second float32) operator.Operator {
	return operator.NewOperator(f, operator.Between, expr.NewValue(first).And(expr.NewValue(second)))
}

func (f Float32) BetweenPaths(first, second lingo.Expression) operator.Operator {
	return operator.NewOperator(f, operator.Between, operator.NewOperator(first, operator.And, second))
}

func (f Float32) NotBetween(first, second float32) operator.Operator {
	return operator.NewOperator(f, operator.NotBetween, expr.NewValue(first).And(expr.NewValue(second)))
}

func (f Float32) NotBetweenPaths(first, second lingo.Expression) operator.Operator {
	return operator.NewOperator(f, operator.NotBetween, operator.NewOperator(first, operator.And, second))
}
