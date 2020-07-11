package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/set"
	"github.com/weworksandbox/lingo/sql"
)

func NewFloat64WithAlias(e lingo.Table, name, alias string) Float64 {
	return Float64{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewFloat64(e lingo.Table, name string) Float64 {
	return NewFloat64WithAlias(e, name, "")
}

type Float64 struct {
	entity lingo.Table
	name   string
	alias  string
}

func (f Float64) GetParent() lingo.Table {
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

func (f Float64) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, f)
}

func (f Float64) To(value float64) set.Set {
	return set.NewSet(f, expr.NewValue(value))
}

func (f Float64) ToExpr(setExp lingo.Expression) set.Set {
	return set.NewSet(f, setExp)
}

func (f Float64) Eq(equalTo float64) operator.Binary {
	return operator.NewBinary(f, operator.Eq, expr.NewValue(equalTo))
}

func (f Float64) EqPath(equalTo lingo.Expression) operator.Binary {
	return operator.NewBinary(f, operator.Eq, equalTo)
}

func (f Float64) NotEq(notEqualTo float64) operator.Binary {
	return operator.NewBinary(f, operator.NotEq, expr.NewValue(notEqualTo))
}

func (f Float64) NotEqPath(notEqualTo lingo.Expression) operator.Binary {
	return operator.NewBinary(f, operator.NotEq, notEqualTo)
}

func (f Float64) LT(lessThan float64) operator.Binary {
	return operator.NewBinary(f, operator.LessThan, expr.NewValue(lessThan))
}

func (f Float64) LTPath(lessThan lingo.Expression) operator.Binary {
	return operator.NewBinary(f, operator.LessThan, lessThan)
}

func (f Float64) LTOrEq(lessThanOrEqual float64) operator.Binary {
	return operator.NewBinary(f, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (f Float64) LTOrEqPath(lessThanOrEqual lingo.Expression) operator.Binary {
	return operator.NewBinary(f, operator.LessThanOrEqual, lessThanOrEqual)
}

func (f Float64) GT(greaterThan float64) operator.Binary {
	return operator.NewBinary(f, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (f Float64) GTPath(greaterThan lingo.Expression) operator.Binary {
	return operator.NewBinary(f, operator.GreaterThan, greaterThan)
}

func (f Float64) GTOrEq(greaterThanOrEqual float64) operator.Binary {
	return operator.NewBinary(f, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (f Float64) GTOrEqPath(greaterThanOrEqual lingo.Expression) operator.Binary {
	return operator.NewBinary(f, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (f Float64) IsNull() operator.Unary {
	return operator.NewUnary(f, operator.Null)
}

func (f Float64) IsNotNull() operator.Unary {
	return operator.NewUnary(f, operator.NotNull)
}

func (f Float64) In(values ...float64) operator.Binary {
	return operator.NewBinary(f, operator.In, expr.NewParens(expr.NewValue(values)))
}

func (f Float64) InPaths(values ...lingo.Expression) operator.Binary {
	return operator.NewBinary(f, operator.In, expr.NewParens(expr.ToList(values)))
}

func (f Float64) NotIn(values ...float64) operator.Binary {
	return operator.NewBinary(f, operator.NotIn, expr.NewParens(expr.NewValue(values)))
}

func (f Float64) NotInPaths(values ...lingo.Expression) operator.Binary {
	return operator.NewBinary(f, operator.NotIn, expr.NewParens(expr.ToList(values)))
}

func (f Float64) Between(first, second float64) operator.Binary {
	return operator.NewBinary(f, operator.Between, expr.NewParens(expr.NewValue(first).And(expr.NewValue(second))))
}

func (f Float64) BetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NewBinary(f, operator.Between, expr.NewParens(operator.NewBinary(first, operator.And, second)))
}

func (f Float64) NotBetween(first, second float64) operator.Binary {
	return operator.NewBinary(f, operator.NotBetween, expr.NewParens(expr.NewValue(first).And(expr.NewValue(second))))
}

func (f Float64) NotBetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NewBinary(f, operator.NotBetween, expr.NewParens(operator.NewBinary(first, operator.And, second)))
}
