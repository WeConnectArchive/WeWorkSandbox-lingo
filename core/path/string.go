package path

import (
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/expression"
	"github.com/weworksandbox/lingo/core/operator"
)

func NewStringPathWithAlias(e core.Table, name, alias string) StringPath {
	return StringPath{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewStringPath(e core.Table, name string) StringPath {
	return NewStringPathWithAlias(e, name, "")
}

type StringPath struct {
	entity core.Table
	name   string
	alias  string
}

func (s StringPath) GetParent() core.Table {
	return s.entity
}

func (s StringPath) GetName() string {
	return s.name
}

func (s StringPath) GetAlias() string {
	return s.alias
}

func (s StringPath) As(alias string) StringPath {
	s.alias = alias
	return s
}

func (s StringPath) GetSQL(d core.Dialect, sql core.SQL) error {
	return ExpandColumnWithDialect(d, s, sql)
}

func (s StringPath) To(value string) core.Set {
	return expression.NewSet(s, expression.NewValue(value))
}

func (s StringPath) ToExpression(setExp core.Expression) core.Set {
	return expression.NewSet(s, setExp)
}

func (s StringPath) Eq(equalTo string) core.ComboExpression {
	return expression.NewOperator(s, operator.Eq, expression.NewValue(equalTo))
}

func (s StringPath) EqPath(equalTo core.Expression) core.ComboExpression {
	return expression.NewOperator(s, operator.Eq, equalTo)
}

func (s StringPath) NotEq(notEqualTo string) core.ComboExpression {
	return expression.NewOperator(s, operator.NotEq, expression.NewValue(notEqualTo))
}

func (s StringPath) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return expression.NewOperator(s, operator.NotEq, notEqualTo)
}

func (s StringPath) Like(like string) core.ComboExpression {
	return expression.NewOperator(s, operator.Like, expression.NewValue(like))
}

func (s StringPath) LikePath(likePath core.Expression) core.ComboExpression {
	return expression.NewOperator(s, operator.Like, likePath)
}

func (s StringPath) NotLike(like string) core.ComboExpression {
	return expression.NewOperator(s, operator.NotLike, expression.NewValue(like))
}

func (s StringPath) NotLikePath(notLikePath core.Expression) core.ComboExpression {
	return expression.NewOperator(s, operator.NotLike, notLikePath)
}

func (s StringPath) IsNull() core.ComboExpression {
	return expression.NewOperator(s, operator.Null)
}

func (s StringPath) IsNotNull() core.ComboExpression {
	return expression.NewOperator(s, operator.NotNull)
}

func (s StringPath) In(values ...string) core.ComboExpression {
	return expression.NewOperator(s, operator.In, expression.NewValue(values))
}

func (s StringPath) InPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(s, operator.In, values...)
}

func (s StringPath) NotIn(values ...string) core.ComboExpression {
	return expression.NewOperator(s, operator.NotIn, expression.NewValue(values))
}

func (s StringPath) NotInPaths(values ...core.Expression) core.ComboExpression {
	return expression.NewOperator(s, operator.NotIn, values...)
}
