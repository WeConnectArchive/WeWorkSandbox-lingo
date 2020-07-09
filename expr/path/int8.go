package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/sql"
)

func NewInt8WithAlias(e lingo.Table, name, alias string) Int8 {
	return Int8{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt8(e lingo.Table, name string) Int8 {
	return NewInt8WithAlias(e, name, "")
}

type Int8 struct {
	entity lingo.Table
	name   string
	alias  string
}

func (i Int8) GetParent() lingo.Table {
	return i.entity
}

func (i Int8) GetName() string {
	return i.name
}

func (i Int8) GetAlias() string {
	return i.alias
}

func (i Int8) As(alias string) Int8 {
	i.alias = alias
	return i
}

func (i Int8) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int8) To(value int8) lingo.Set {
	return expr.NewSet(i, expr.NewValue(value))
}

func (i Int8) ToExpr(setExp lingo.Expression) lingo.Set {
	return expr.NewSet(i, setExp)
}

func (i Int8) Eq(equalTo int8) lingo.ComboExpression {
	return operator.NewOperator(i, operator.Eq, expr.NewValue(equalTo))
}

func (i Int8) EqPath(equalTo lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.Eq, equalTo)
}

func (i Int8) NotEq(notEqualTo int8) lingo.ComboExpression {
	return operator.NewOperator(i, operator.NotEq, expr.NewValue(notEqualTo))
}

func (i Int8) NotEqPath(notEqualTo lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.NotEq, notEqualTo)
}

func (i Int8) LT(lessThan int8) lingo.ComboExpression {
	return operator.NewOperator(i, operator.LessThan, expr.NewValue(lessThan))
}

func (i Int8) LTPath(lessThan lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.LessThan, lessThan)
}

func (i Int8) LTOrEq(lessThanOrEqual int8) lingo.ComboExpression {
	return operator.NewOperator(i, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (i Int8) LTOrEqPath(lessThanOrEqual lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int8) GT(greaterThan int8) lingo.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (i Int8) GTPath(greaterThan lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThan, greaterThan)
}

func (i Int8) GTOrEq(greaterThanOrEqual int8) lingo.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (i Int8) GTOrEqPath(greaterThanOrEqual lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int8) IsNull() lingo.ComboExpression {
	return operator.NewOperator(i, operator.Null)
}

func (i Int8) IsNotNull() lingo.ComboExpression {
	return operator.NewOperator(i, operator.NotNull)
}

func (i Int8) In(values ...int8) lingo.ComboExpression {
	return operator.NewOperator(i, operator.In, expr.NewValue(values))
}

func (i Int8) InPaths(values ...lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.In, values...)
}

func (i Int8) NotIn(values ...int8) lingo.ComboExpression {
	return operator.NewOperator(i, operator.NotIn, expr.NewValue(values))
}

func (i Int8) NotInPaths(values ...lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.NotIn, values...)
}

func (i Int8) Between(first, second int8) lingo.ComboExpression {
	return operator.NewOperator(i, operator.Between, expr.NewValue(first).And(expr.NewValue(second)))
}

func (i Int8) BetweenPaths(first, second lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.Between, operator.NewOperator(first, operator.And, second))
}

func (i Int8) NotBetween(first, second int8) lingo.ComboExpression {
	return operator.NewOperator(i, operator.NotBetween, expr.NewValue(first).And(expr.NewValue(second)))
}

func (i Int8) NotBetweenPaths(first, second lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.NotBetween, operator.NewOperator(first, operator.And, second))
}
