package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/set"
	"github.com/weworksandbox/lingo/sql"
)

func NewInt64PathWithAlias(e lingo.Table, name, alias string) Int64 {
	return Int64{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt64Path(e lingo.Table, name string) Int64 {
	return NewInt64PathWithAlias(e, name, "")
}

type Int64 struct {
	entity lingo.Table
	name   string
	alias  string
}

func (i Int64) GetParent() lingo.Table {
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

func (i Int64) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int64) To(value int64) set.Set {
	return set.NewSet(i, expr.NewValue(value))
}

func (i Int64) ToExpr(setExp lingo.Expression) set.Set {
	return set.NewSet(i, setExp)
}

func (i Int64) Eq(equalTo int64) operator.Operator {
	return operator.NewOperator(i, operator.Eq, expr.NewValue(equalTo))
}

func (i Int64) EqPath(equalTo lingo.Expression) operator.Operator {
	return operator.NewOperator(i, operator.Eq, equalTo)
}

func (i Int64) NotEq(notEqualTo int64) operator.Operator {
	return operator.NewOperator(i, operator.NotEq, expr.NewValue(notEqualTo))
}

func (i Int64) NotEqPath(notEqualTo lingo.Expression) operator.Operator {
	return operator.NewOperator(i, operator.NotEq, notEqualTo)
}

func (i Int64) LT(lessThan int64) operator.Operator {
	return operator.NewOperator(i, operator.LessThan, expr.NewValue(lessThan))
}

func (i Int64) LTPath(lessThan lingo.Expression) operator.Operator {
	return operator.NewOperator(i, operator.LessThan, lessThan)
}

func (i Int64) LTOrEq(lessThanOrEqual int64) operator.Operator {
	return operator.NewOperator(i, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (i Int64) LTOrEqPath(lessThanOrEqual lingo.Expression) operator.Operator {
	return operator.NewOperator(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int64) GT(greaterThan int64) operator.Operator {
	return operator.NewOperator(i, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (i Int64) GTPath(greaterThan lingo.Expression) operator.Operator {
	return operator.NewOperator(i, operator.GreaterThan, greaterThan)
}

func (i Int64) GTOrEq(greaterThanOrEqual int64) operator.Operator {
	return operator.NewOperator(i, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (i Int64) GTOrEqPath(greaterThanOrEqual lingo.Expression) operator.Operator {
	return operator.NewOperator(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int64) IsNull() operator.Operator {
	return operator.NewOperator(i, operator.Null)
}

func (i Int64) IsNotNull() operator.Operator {
	return operator.NewOperator(i, operator.NotNull)
}

func (i Int64) In(values ...int64) operator.Operator {
	return operator.NewOperator(i, operator.In, expr.NewValue(values))
}

func (i Int64) InPaths(values ...lingo.Expression) operator.Operator {
	return operator.NewOperator(i, operator.In, values...)
}

func (i Int64) NotIn(values ...int64) operator.Operator {
	return operator.NewOperator(i, operator.NotIn, expr.NewValue(values))
}

func (i Int64) NotInPaths(values ...lingo.Expression) operator.Operator {
	return operator.NewOperator(i, operator.NotIn, values...)
}

func (i Int64) Between(first, second int64) operator.Operator {
	return operator.NewOperator(i, operator.Between, expr.NewValue(first).And(expr.NewValue(second)))
}

func (i Int64) BetweenPaths(first, second lingo.Expression) operator.Operator {
	return operator.NewOperator(i, operator.Between, operator.NewOperator(first, operator.And, second))
}

func (i Int64) NotBetween(first, second int64) operator.Operator {
	return operator.NewOperator(i, operator.NotBetween, expr.NewValue(first).And(expr.NewValue(second)))
}

func (i Int64) NotBetweenPaths(first, second lingo.Expression) operator.Operator {
	return operator.NewOperator(i, operator.NotBetween, operator.NewOperator(first, operator.And, second))
}
