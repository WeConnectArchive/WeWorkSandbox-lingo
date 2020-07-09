package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/sql"
)

func NewIntWithAlias(e lingo.Table, name, alias string) Int {
	return Int{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt(e lingo.Table, name string) Int {
	return NewIntWithAlias(e, name, "")
}

type Int struct {
	entity lingo.Table
	name   string
	alias  string
}

func (i Int) GetParent() lingo.Table {
	return i.entity
}

func (i Int) GetName() string {
	return i.name
}

func (i Int) GetAlias() string {
	return i.alias
}

func (i Int) As(alias string) Int {
	i.alias = alias
	return i
}

func (i Int) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int) To(value int) lingo.Set {
	return expr.NewSet(i, expr.NewValue(value))
}

func (i Int) ToExpr(setExp lingo.Expression) lingo.Set {
	return expr.NewSet(i, setExp)
}

func (i Int) Eq(equalTo int) lingo.ComboExpression {
	return operator.NewOperator(i, operator.Eq, expr.NewValue(equalTo))
}

func (i Int) EqPath(equalTo lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.Eq, equalTo)
}

func (i Int) NotEq(notEqualTo int) lingo.ComboExpression {
	return operator.NewOperator(i, operator.NotEq, expr.NewValue(notEqualTo))
}

func (i Int) NotEqPath(notEqualTo lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.NotEq, notEqualTo)
}

func (i Int) LT(lessThan int) lingo.ComboExpression {
	return operator.NewOperator(i, operator.LessThan, expr.NewValue(lessThan))
}

func (i Int) LTPath(lessThan lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.LessThan, lessThan)
}

func (i Int) LTOrEq(lessThanOrEqual int) lingo.ComboExpression {
	return operator.NewOperator(i, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (i Int) LTOrEqPath(lessThanOrEqual lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int) GT(greaterThan int) lingo.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (i Int) GTPath(greaterThan lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThan, greaterThan)
}

func (i Int) GTOrEq(greaterThanOrEqual int) lingo.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (i Int) GTOrEqPath(greaterThanOrEqual lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int) IsNull() lingo.ComboExpression {
	return operator.NewOperator(i, operator.Null)
}

func (i Int) IsNotNull() lingo.ComboExpression {
	return operator.NewOperator(i, operator.NotNull)
}

func (i Int) In(values ...int) lingo.ComboExpression {
	return operator.NewOperator(i, operator.In, expr.NewValue(values))
}

func (i Int) InPaths(values ...lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.In, values...)
}

func (i Int) NotIn(values ...int) lingo.ComboExpression {
	return operator.NewOperator(i, operator.NotIn, expr.NewValue(values))
}

func (i Int) NotInPaths(values ...lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.NotIn, values...)
}

func (i Int) Between(first, second int) lingo.ComboExpression {
	return operator.NewOperator(i, operator.Between, expr.NewValue(first).And(expr.NewValue(second)))
}

func (i Int) BetweenPaths(first, second lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.Between, operator.NewOperator(first, operator.And, second))
}

func (i Int) NotBetween(first, second int) lingo.ComboExpression {
	return operator.NewOperator(i, operator.NotBetween, expr.NewValue(first).And(expr.NewValue(second)))
}

func (i Int) NotBetweenPaths(first, second lingo.Expression) lingo.ComboExpression {
	return operator.NewOperator(i, operator.NotBetween, operator.NewOperator(first, operator.And, second))
}
