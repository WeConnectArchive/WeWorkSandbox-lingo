package path

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expr"
	"github.com/weworksandbox/lingo/pkg/core/expr/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

func NewStringWithAlias(e core.Table, name, alias string) String {
	return String{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewString(e core.Table, name string) String {
	return NewStringWithAlias(e, name, "")
}

type String struct {
	entity core.Table
	name   string
	alias  string
}

func (s String) GetParent() core.Table {
	return s.entity
}

func (s String) GetName() string {
	return s.name
}

func (s String) GetAlias() string {
	return s.alias
}

func (s String) As(alias string) String {
	s.alias = alias
	return s
}

func (s String) ToSQL(d core.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, s)
}

func (s String) To(value string) core.Set {
	return expr.NewSet(s, expr.NewValue(value))
}

func (s String) ToExpression(setExp core.Expression) core.Set {
	return expr.NewSet(s, setExp)
}

func (s String) Eq(equalTo string) core.ComboExpression {
	return operator.NewOperator(s, operator.Eq, expr.NewValue(equalTo))
}

func (s String) EqPath(equalTo core.Expression) core.ComboExpression {
	return operator.NewOperator(s, operator.Eq, equalTo)
}

func (s String) NotEq(notEqualTo string) core.ComboExpression {
	return operator.NewOperator(s, operator.NotEq, expr.NewValue(notEqualTo))
}

func (s String) NotEqPath(notEqualTo core.Expression) core.ComboExpression {
	return operator.NewOperator(s, operator.NotEq, notEqualTo)
}

func (s String) Like(like string) core.ComboExpression {
	return operator.NewOperator(s, operator.Like, expr.NewValue(like))
}

func (s String) LikePath(likePath core.Expression) core.ComboExpression {
	return operator.NewOperator(s, operator.Like, likePath)
}

func (s String) NotLike(like string) core.ComboExpression {
	return operator.NewOperator(s, operator.NotLike, expr.NewValue(like))
}

func (s String) NotLikePath(notLikePath core.Expression) core.ComboExpression {
	return operator.NewOperator(s, operator.NotLike, notLikePath)
}

func (s String) IsNull() core.ComboExpression {
	return operator.NewOperator(s, operator.Null)
}

func (s String) IsNotNull() core.ComboExpression {
	return operator.NewOperator(s, operator.NotNull)
}

func (s String) In(values ...string) core.ComboExpression {
	return operator.NewOperator(s, operator.In, expr.NewValue(values))
}

func (s String) InPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(s, operator.In, values...)
}

func (s String) NotIn(values ...string) core.ComboExpression {
	return operator.NewOperator(s, operator.NotIn, expr.NewValue(values))
}

func (s String) NotInPaths(values ...core.Expression) core.ComboExpression {
	return operator.NewOperator(s, operator.NotIn, values...)
}
