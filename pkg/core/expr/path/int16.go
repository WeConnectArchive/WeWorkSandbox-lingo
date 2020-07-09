package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expr"
	"github.com/weworksandbox/lingo/pkg/core/expr/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func NewInt16WithAlias(e core.Table, name, alias string) Int16 {
	return Int16{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt16(e core.Table, name string) Int16 {
	return NewInt16WithAlias(e, name, "")
}

type Int16 struct {
	entity core.Table
	name   string
	alias  string
}

func (i Int16) GetParent() core.Table {
	return i.entity
}

func (i Int16) GetName() string {
	return i.name
}

func (i Int16) GetAlias() string {
	return i.alias
}

func (i Int16) As(alias string) Int16 {
	i.alias = alias
	return i
}

func (i Int16) ToSQL(d core.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int16) To(value int16) core.Set {
	return expr.NewSet(i, expr.NewValue(value))
}

func (i Int16) ToExpr(setExp core.Expression) core.Set {
	return expr.NewSet(i, setExp)
}

func (i Int16) Eq(equalTo int16) core.ComboExpression {
	return operator.NewOperator(i, operator.Eq, expr.NewValue(equalTo))
}

func (i Int16) EqPath(equalTo core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.Eq, equalTo)
}

func (i Int16) NotEq(notEqualTo int16) core.ComboExpression {
	return operator.NewOperator(i, operator.NotEq, expr.NewValue(notEqualTo))
}

func (i Int16) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.NotEq, notEqualTo)
}

func (i Int16) LT(lessThan int16) core.ComboExpression {
	return operator.NewOperator(i, operator.LessThan, expr.NewValue(lessThan))
}

func (i Int16) LTPath(lessThan core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.LessThan, lessThan)
}

func (i Int16) LTOrEq(lessThanOrEqual int16) core.ComboExpression {
	return operator.NewOperator(i, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (i Int16) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int16) GT(greaterThan int16) core.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (i Int16) GTPath(greaterThan core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThan, greaterThan)
}

func (i Int16) GTOrEq(greaterThanOrEqual int16) core.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (i Int16) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int16) IsNull() core.ComboExpression {
	return operator.NewOperator(i, operator.Null)
}

func (i Int16) IsNotNull() core.ComboExpression {
	return operator.NewOperator(i, operator.NotNull)
}

func (i Int16) In(values ...int16) core.ComboExpression {
	return operator.NewOperator(i, operator.In, expr.NewValue(values))
}

func (i Int16) InPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.In, values...)
}

func (i Int16) NotIn(values ...int16) core.ComboExpression {
	return operator.NewOperator(i, operator.NotIn, expr.NewValue(values))
}

func (i Int16) NotInPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.NotIn, values...)
}

func (i Int16) Between(first, second int16) core.ComboExpression {
	return operator.NewOperator(i, operator.Between, expr.NewValue(first).And(expr.NewValue(second)))
}

func (i Int16) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.Between, operator.NewOperator(first, operator.And, second))
}

func (i Int16) NotBetween(first, second int16) core.ComboExpression {
	return operator.NewOperator(i, operator.NotBetween, expr.NewValue(first).And(expr.NewValue(second)))
}

func (i Int16) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.NotBetween, operator.NewOperator(first, operator.And, second))
}
