package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/expression/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func NewBinaryWithAlias(e core.Table, name, alias string) Binary {
	return Binary{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewBinary(e core.Table, name string) Binary {
	return NewBinaryWithAlias(e, name, "")
}

type Binary struct {
	entity core.Table
	name   string
	alias  string
}

func (b Binary) GetParent() core.Table {
	return b.entity
}

func (b Binary) GetName() string {
	return b.name
}

func (b Binary) GetAlias() string {
	return b.alias
}

func (b Binary) As(alias string) Binary {
	b.alias = alias
	return b
}

func (b Binary) ToSQL(d core.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, b)
}

func (b Binary) To(value []byte) core.Set {
	return expression.NewSet(b, expression.NewValue(value))
}

func (b Binary) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(b, setExp)
}

func (b Binary) Eq(equalTo []byte) core.ComboExpression {
	return operator.NewOperator(b, operator.Eq, expression.NewValue(equalTo))
}

func (b Binary) EqPath(equalTo core.Expression) core.ComboExpression {
	return operator.NewOperator(b, operator.Eq, equalTo)
}

func (b Binary) NotEq(notEqualTo []byte) core.ComboExpression {
	return operator.NewOperator(b, operator.NotEq, expression.NewValue(notEqualTo))
}

func (b Binary) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return operator.NewOperator(b, operator.NotEq, notEqualTo)
}

func (b Binary) LT(lessThan []byte) core.ComboExpression {
	return operator.NewOperator(b, operator.LessThan, expression.NewValue(lessThan))
}

func (b Binary) LTPath(lessThan core.Expression) core.ComboExpression {
	return operator.NewOperator(b, operator.LessThan, lessThan)
}

func (b Binary) LTOrEq(lessThanOrEqual []byte) core.ComboExpression {
	return operator.NewOperator(b, operator.LessThanOrEqual, expression.NewValue(lessThanOrEqual))
}

func (b Binary) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return operator.NewOperator(b, operator.LessThanOrEqual, lessThanOrEqual)
}

func (b Binary) GT(greaterThan []byte) core.ComboExpression {
	return operator.NewOperator(b, operator.GreaterThan, expression.NewValue(greaterThan))
}

func (b Binary) GTPath(greaterThan core.Expression) core.ComboExpression {
	return operator.NewOperator(b, operator.GreaterThan, greaterThan)
}

func (b Binary) GTOrEq(greaterThanOrEqual []byte) core.ComboExpression {
	return operator.NewOperator(b, operator.GreaterThanOrEqual, expression.NewValue(greaterThanOrEqual))
}

func (b Binary) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return operator.NewOperator(b, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (b Binary) IsNull() core.ComboExpression {
	return operator.NewOperator(b, operator.Null)
}

func (b Binary) IsNotNull() core.ComboExpression {
	return operator.NewOperator(b, operator.NotNull)
}

func (b Binary) In(values ...[]byte) core.ComboExpression {
	return operator.NewOperator(b, operator.In, expression.NewValue(values))
}

func (b Binary) InPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(b, operator.In, values...)
}

func (b Binary) NotIn(values ...[]byte) core.ComboExpression {
	return operator.NewOperator(b, operator.NotIn, expression.NewValue(values))
}

func (b Binary) NotInPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(b, operator.NotIn, values...)
}

func (b Binary) Between(first, second []byte) core.ComboExpression {
	return operator.NewOperator(b, operator.Between, expression.NewValue(first).And(expression.NewValue(second)))
}

func (b Binary) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return operator.NewOperator(b, operator.Between, operator.NewOperator(first, operator.And, second))
}

func (b Binary) NotBetween(first, second []byte) core.ComboExpression {
	return operator.NewOperator(b, operator.NotBetween, expression.NewValue(first).And(expression.NewValue(second)))
}

func (b Binary) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return operator.NewOperator(b, operator.NotBetween, operator.NewOperator(first, operator.And, second))
}
