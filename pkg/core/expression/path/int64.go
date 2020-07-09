package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/expression/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func NewInt64PathWithAlias(e core.Table, name, alias string) Int64 {
	return Int64{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt64Path(e core.Table, name string) Int64 {
	return NewInt64PathWithAlias(e, name, "")
}

type Int64 struct {
	entity core.Table
	name   string
	alias  string
}

func (i Int64) GetParent() core.Table {
	return i.entity
}

func (i Int64) GetName() string {
	return i.name
}

func (i Int64) GetAlias() string {
	return i.alias
}

func (i Int64) As(alias string) Int64 {
	i.alias = alias
	return i
}

func (i Int64) ToSQL(d core.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int64) To(value int64) core.Set {
	return expression.NewSet(i, expression.NewValue(value))
}

func (i Int64) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(i, setExp)
}

func (i Int64) Eq(equalTo int64) core.ComboExpression {
	return operator.NewOperator(i, operator.Eq, expression.NewValue(equalTo))
}

func (i Int64) EqPath(equalTo core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.Eq, equalTo)
}

func (i Int64) NotEq(notEqualTo int64) core.ComboExpression {
	return operator.NewOperator(i, operator.NotEq, expression.NewValue(notEqualTo))
}

func (i Int64) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.NotEq, notEqualTo)
}

func (i Int64) LT(lessThan int64) core.ComboExpression {
	return operator.NewOperator(i, operator.LessThan, expression.NewValue(lessThan))
}

func (i Int64) LTPath(lessThan core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.LessThan, lessThan)
}

func (i Int64) LTOrEq(lessThanOrEqual int64) core.ComboExpression {
	return operator.NewOperator(i, operator.LessThanOrEqual, expression.NewValue(lessThanOrEqual))
}

func (i Int64) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int64) GT(greaterThan int64) core.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThan, expression.NewValue(greaterThan))
}

func (i Int64) GTPath(greaterThan core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThan, greaterThan)
}

func (i Int64) GTOrEq(greaterThanOrEqual int64) core.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThanOrEqual, expression.NewValue(greaterThanOrEqual))
}

func (i Int64) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int64) IsNull() core.ComboExpression {
	return operator.NewOperator(i, operator.Null)
}

func (i Int64) IsNotNull() core.ComboExpression {
	return operator.NewOperator(i, operator.NotNull)
}

func (i Int64) In(values ...int64) core.ComboExpression {
	return operator.NewOperator(i, operator.In, expression.NewValue(values))
}

func (i Int64) InPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.In, values...)
}

func (i Int64) NotIn(values ...int64) core.ComboExpression {
	return operator.NewOperator(i, operator.NotIn, expression.NewValue(values))
}

func (i Int64) NotInPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.NotIn, values...)
}

func (i Int64) Between(first, second int64) core.ComboExpression {
	return operator.NewOperator(i, operator.Between, expression.NewValue(first).And(expression.NewValue(second)))
}

func (i Int64) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.Between, operator.NewOperator(first, operator.And, second))
}

func (i Int64) NotBetween(first, second int64) core.ComboExpression {
	return operator.NewOperator(i, operator.NotBetween, expression.NewValue(first).And(expression.NewValue(second)))
}

func (i Int64) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return operator.NewOperator(i, operator.NotBetween, operator.NewOperator(first, operator.And, second))
}
