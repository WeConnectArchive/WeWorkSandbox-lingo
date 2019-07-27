package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/operator"
)

func NewBinaryPathWithAlias(e core.Table, name, alias string) BinaryPath {
	return BinaryPath{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewBinaryPath(e core.Table, name string) BinaryPath {
	return NewBinaryPathWithAlias(e, name, "")
}

type BinaryPath struct {
	entity core.Table
	name   string
	alias  string
}

func (b BinaryPath) GetParent() core.Table {
	return b.entity
}

func (b BinaryPath) GetName() string {
	return b.name
}

func (b BinaryPath) GetAlias() string {
	return b.alias
}

func (s BinaryPath) As(alias string) BinaryPath {
	s.alias = alias
	return s
}

func (b BinaryPath) GetSQL(d core.Dialect) (core.SQL, error) {
	return ExpandColumnWithDialect(d, b)
}

func (b BinaryPath) To(value []byte) core.Set {
	return expression.NewSet(b, expression.NewValue(value))
}

func (b BinaryPath) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(b, setExp)
}

func (b BinaryPath) Eq(equalTo []byte) core.ComboExpression {
	return expression.NewOperator(b, operator.Eq, expression.NewValue(equalTo))
}

func (b BinaryPath) EqPath(equalTo core.Expression) core.ComboExpression {
	return expression.NewOperator(b, operator.Eq, equalTo)
}

func (b BinaryPath) NotEq(notEqualTo []byte) core.ComboExpression {
	return expression.NewOperator(b, operator.NotEq, expression.NewValue(notEqualTo))
}

func (b BinaryPath) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return expression.NewOperator(b, operator.NotEq, notEqualTo)
}

func (b BinaryPath) LT(lessThan []byte) core.ComboExpression {
	return expression.NewOperator(b, operator.LessThan, expression.NewValue(lessThan))
}

func (b BinaryPath) LTPath(lessThan core.Expression) core.ComboExpression {
	return expression.NewOperator(b, operator.LessThan, lessThan)
}

func (b BinaryPath) LTOrEq(lessThanOrEqual []byte) core.ComboExpression {
	return expression.NewOperator(b, operator.LessThanOrEqual, expression.NewValue(lessThanOrEqual))
}

func (b BinaryPath) LTOrEqPath(lessThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(b, operator.LessThanOrEqual, lessThanOrEqual)
}

func (b BinaryPath) GT(greaterThan []byte) core.ComboExpression {
	return expression.NewOperator(b, operator.GreaterThan, expression.NewValue(greaterThan))
}

func (b BinaryPath) GTPath(greaterThan core.Expression) core.ComboExpression {
	return expression.NewOperator(b, operator.GreaterThan, greaterThan)
}

func (b BinaryPath) GTOrEq(greaterThanOrEqual []byte) core.ComboExpression {
	return expression.NewOperator(b, operator.GreaterThanOrEqual, expression.NewValue(greaterThanOrEqual))
}

func (b BinaryPath) GTOrEqPath(greaterThanOrEqual core.Expression) core.ComboExpression {
	return expression.NewOperator(b, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (b BinaryPath) IsNull() core.ComboExpression {
	return expression.NewOperator(b, operator.Null)
}

func (b BinaryPath) IsNotNull() core.ComboExpression {
	return expression.NewOperator(b, operator.NotNull)
}

func (b BinaryPath) In(values ...[]byte) core.ComboExpression {
	return expression.NewOperator(b, operator.In, expression.NewValues(values)...)
}

func (b BinaryPath) InPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(b, operator.In, values...)
}

func (b BinaryPath) NotIn(values ...[]byte) core.ComboExpression {
	return expression.NewOperator(b, operator.NotIn, expression.NewValues(values)...)
}

func (b BinaryPath) NotInPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(b, operator.NotIn, values...)
}

func (b BinaryPath) Between(first, second []byte) core.ComboExpression {
	return expression.NewOperator(b, operator.Between, expression.NewValue(first).And(expression.NewValue(second)))
}

func (b BinaryPath) BetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(b, operator.Between, expression.NewOperator(first, operator.And, second))
}

func (b BinaryPath) NotBetween(first, second []byte) core.ComboExpression {
	return expression.NewOperator(b, operator.NotBetween, expression.NewValue(first).And(expression.NewValue(second)))
}

func (b BinaryPath) NotBetweenPaths(first, second core.Expression) core.ComboExpression {
	return expression.NewOperator(b, operator.NotBetween, expression.NewOperator(first, operator.And, second))
}
