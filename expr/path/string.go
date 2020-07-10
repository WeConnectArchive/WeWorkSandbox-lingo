package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/set"
	"github.com/weworksandbox/lingo/sql"
)

func NewStringWithAlias(e lingo.Table, name, alias string) String {
	return String{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewString(e lingo.Table, name string) String {
	return NewStringWithAlias(e, name, "")
}

type String struct {
	entity lingo.Table
	name   string
	alias  string
}

func (s String) GetParent() lingo.Table {
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

func (s String) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, s)
}

func (s String) To(value string) set.Set {
	return set.NewSet(s, expr.NewValue(value))
}

func (s String) ToExpr(setExp lingo.Expression) set.Set {
	return set.NewSet(s, setExp)
}

func (s String) Eq(equalTo string) operator.Operator {
	return operator.NewOperator(s, operator.Eq, expr.NewValue(equalTo))
}

func (s String) EqPath(equalTo lingo.Expression) operator.Operator {
	return operator.NewOperator(s, operator.Eq, equalTo)
}

func (s String) NotEq(notEqualTo string) operator.Operator {
	return operator.NewOperator(s, operator.NotEq, expr.NewValue(notEqualTo))
}

func (s String) NotEqPath(notEqualTo lingo.Expression) operator.Operator {
	return operator.NewOperator(s, operator.NotEq, notEqualTo)
}

func (s String) Like(like string) operator.Operator {
	return operator.NewOperator(s, operator.Like, expr.NewValue(like))
}

func (s String) LikePath(likePath lingo.Expression) operator.Operator {
	return operator.NewOperator(s, operator.Like, likePath)
}

func (s String) NotLike(like string) operator.Operator {
	return operator.NewOperator(s, operator.NotLike, expr.NewValue(like))
}

func (s String) NotLikePath(notLikePath lingo.Expression) operator.Operator {
	return operator.NewOperator(s, operator.NotLike, notLikePath)
}

func (s String) IsNull() operator.Operator {
	return operator.NewOperator(s, operator.Null)
}

func (s String) IsNotNull() operator.Operator {
	return operator.NewOperator(s, operator.NotNull)
}

func (s String) In(values ...string) operator.Operator {
	return operator.NewOperator(s, operator.In, expr.NewValue(values))
}

func (s String) InPaths(values ...lingo.Expression) operator.Operator {
	return operator.NewOperator(s, operator.In, values...)
}

func (s String) NotIn(values ...string) operator.Operator {
	return operator.NewOperator(s, operator.NotIn, expr.NewValue(values))
}

func (s String) NotInPaths(values ...lingo.Expression) operator.Operator {
	return operator.NewOperator(s, operator.NotIn, values...)
}
