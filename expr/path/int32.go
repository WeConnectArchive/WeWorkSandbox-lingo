package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/set"
	"github.com/weworksandbox/lingo/sql"
)

func NewInt32WithAlias(e lingo.Table, name, alias string) Int32 {
	return Int32{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt32(e lingo.Table, name string) Int32 {
	return NewInt32WithAlias(e, name, "")
}

type Int32 struct {
	entity lingo.Table
	name   string
	alias  string
}

func (i Int32) GetParent() lingo.Table {
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

func (i Int32) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int32) To(value int32) set.Set {
	return set.NewSet(i, expr.NewValue(value))
}

func (i Int32) ToExpr(setExp lingo.Expression) set.Set {
	return set.NewSet(i, setExp)
}

func (i Int32) Eq(equalTo int32) operator.Binary {
	return operator.NewBinary(i, operator.Eq, expr.NewValue(equalTo))
}

func (i Int32) EqPath(equalTo lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.Eq, equalTo)
}

func (i Int32) NotEq(notEqualTo int32) operator.Binary {
	return operator.NewBinary(i, operator.NotEq, expr.NewValue(notEqualTo))
}

func (i Int32) NotEqPath(notEqualTo lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.NotEq, notEqualTo)
}

func (i Int32) LT(lessThan int32) operator.Binary {
	return operator.NewBinary(i, operator.LessThan, expr.NewValue(lessThan))
}

func (i Int32) LTPath(lessThan lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.LessThan, lessThan)
}

func (i Int32) LTOrEq(lessThanOrEqual int32) operator.Binary {
	return operator.NewBinary(i, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (i Int32) LTOrEqPath(lessThanOrEqual lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int32) GT(greaterThan int32) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (i Int32) GTPath(greaterThan lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThan, greaterThan)
}

func (i Int32) GTOrEq(greaterThanOrEqual int32) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (i Int32) GTOrEqPath(greaterThanOrEqual lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int32) IsNull() operator.Unary {
	return operator.NewUnary(i, operator.Null)
}

func (i Int32) IsNotNull() operator.Unary {
	return operator.NewUnary(i, operator.NotNull)
}

func (i Int32) In(values ...int32) operator.Binary {
	return operator.NewBinary(i, operator.In, expr.NewParens(expr.NewValue(values)))
}

func (i Int32) InPaths(values ...lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.In, expr.NewParens(expr.ToList(values)))
}

func (i Int32) NotIn(values ...int32) operator.Binary {
	return operator.NewBinary(i, operator.NotIn, expr.NewParens(expr.NewValue(values)))
}

func (i Int32) NotInPaths(values ...lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.NotIn, expr.NewParens(expr.ToList(values)))
}

func (i Int32) Between(first, second int32) operator.Binary {
	return operator.NewBinary(i, operator.Between, expr.NewParens(expr.NewValue(first).And(expr.NewValue(second))))
}

func (i Int32) BetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.Between, expr.NewParens(operator.NewBinary(first, operator.And, second)))
}

func (i Int32) NotBetween(first, second int32) operator.Binary {
	return operator.NewBinary(i, operator.NotBetween, expr.NewParens(expr.NewValue(first).And(expr.NewValue(second))))
}

func (i Int32) NotBetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.NotBetween, expr.NewParens(operator.NewBinary(first, operator.And, second)))
}
