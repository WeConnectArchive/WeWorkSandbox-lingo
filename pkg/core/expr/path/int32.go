package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expr"
	"github.com/weworksandbox/lingo/pkg/core/expr/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func NewInt32WithAlias(e core.Table, name, alias string) Int32 {
	return Int32{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt32(e core.Table, name string) Int32 {
	return NewInt32WithAlias(e, name, "")
}

type Int32 struct {
	entity core.Table
	name   string
	alias  string
}

func (i Int32) GetParent() core.Table {
	return i.entity
}

func (i Int32) GetName() string {
	return i.name
}

func (i Int32) GetAlias() string {
	return i.alias
}

func (i Int32) As(alias string) Int32 {
	i.alias = alias
	return i
}

func (i Int32) ToSQL(d core.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int32) To(value int32) core.Set {
	return expr.NewSet(i, expr.NewValue(value))
}

func (i Int32) ToExpr(setExp core.Expression) core.Set {
	return expr.NewSet(i, setExp)
}

func (i Int32) Eq(equalTo int32) core.ComboExpression {
	return operator.NewOperator(i, operator.Eq, expr.NewValue(equalTo))
}

func (i Int32) EqPath(equalTo core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.Eq, equalTo)
}

func (i Int32) NotEq(notEqualTo int32) core.ComboExpression {
	return operator.NewOperator(i, operator.NotEq, expr.NewValue(notEqualTo))
}

func (i Int32) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.NotEq, notEqualTo)
}

func (i Int32) LT(lessThan int32) core.ComboExpression {
	return operator.NewOperator(i, operator.LessThan, expr.NewValue(lessThan))
}

func (i Int32) LTPath(lessThan core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.LessThan, lessThan)
}

func (i Int32) LTOrEq(lessThanOrEqual int32) core.ComboExpression {
	return operator.NewOperator(i, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (i Int32) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int32) GT(greaterThan int32) core.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (i Int32) GTPath(greaterThan core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThan, greaterThan)
}

func (i Int32) GTOrEq(greaterThanOrEqual int32) core.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (i Int32) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int32) IsNull() core.ComboExpression {
	return operator.NewOperator(i, operator.Null)
}

func (i Int32) IsNotNull() core.ComboExpression {
	return operator.NewOperator(i, operator.NotNull)
}

func (i Int32) In(values ...int32) core.ComboExpression {
	return operator.NewOperator(i, operator.In, expr.NewValue(values))
}

func (i Int32) InPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.In, values...)
}

func (i Int32) NotIn(values ...int32) core.ComboExpression {
	return operator.NewOperator(i, operator.NotIn, expr.NewValue(values))
}

func (i Int32) NotInPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.NotIn, values...)
}

func (i Int32) Between(first, second int32) core.ComboExpression {
	return operator.NewOperator(i, operator.Between, expr.NewValue(first).And(expr.NewValue(second)))
}

func (i Int32) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.Between, operator.NewOperator(first, operator.And, second))
}

func (i Int32) NotBetween(first, second int32) core.ComboExpression {
	return operator.NewOperator(i, operator.NotBetween, expr.NewValue(first).And(expr.NewValue(second)))
}

func (i Int32) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.NotBetween, operator.NewOperator(first, operator.And, second))
}
