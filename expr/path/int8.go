package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/set"
	"github.com/weworksandbox/lingo/sql"
)

func NewInt8WithAlias(e lingo.Table, name, alias string) Int8 {
	return Int8{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt8(e lingo.Table, name string) Int8 {
	return NewInt8WithAlias(e, name, "")
}

type Int8 struct {
	entity lingo.Table
	name   string
	alias  string
}

func (i Int8) GetParent() lingo.Table {
	return i.entity
}

func (i Int8) GetName() string {
	return i.name
}

func (i Int8) GetAlias() string {
	return i.alias
}

func (i Int8) As(alias string) Int8 {
	i.alias = alias
	return i
}

func (i Int8) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int8) To(value int8) set.Set {
	return set.NewSet(i, expr.NewValue(value))
}

func (i Int8) ToExpr(setExp lingo.Expression) set.Set {
	return set.NewSet(i, setExp)
}

func (i Int8) Eq(equalTo int8) operator.Binary {
	return operator.NewBinary(i, operator.Eq, expr.NewValue(equalTo))
}

func (i Int8) EqPath(equalTo lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.Eq, equalTo)
}

func (i Int8) NotEq(notEqualTo int8) operator.Binary {
	return operator.NewBinary(i, operator.NotEq, expr.NewValue(notEqualTo))
}

func (i Int8) NotEqPath(notEqualTo lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.NotEq, notEqualTo)
}

func (i Int8) LT(lessThan int8) operator.Binary {
	return operator.NewBinary(i, operator.LessThan, expr.NewValue(lessThan))
}

func (i Int8) LTPath(lessThan lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.LessThan, lessThan)
}

func (i Int8) LTOrEq(lessThanOrEqual int8) operator.Binary {
	return operator.NewBinary(i, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (i Int8) LTOrEqPath(lessThanOrEqual lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int8) GT(greaterThan int8) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (i Int8) GTPath(greaterThan lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThan, greaterThan)
}

func (i Int8) GTOrEq(greaterThanOrEqual int8) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (i Int8) GTOrEqPath(greaterThanOrEqual lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int8) IsNull() operator.Unary {
	return operator.NewUnary(i, operator.Null)
}

func (i Int8) IsNotNull() operator.Unary {
	return operator.NewUnary(i, operator.NotNull)
}

func (i Int8) In(values ...int8) operator.Binary {
	return operator.NewBinary(i, operator.In, expr.NewParens(expr.NewValue(values)))
}

func (i Int8) InPaths(values ...lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.In, expr.NewParens(expr.ToList(values)))
}

func (i Int8) NotIn(values ...int8) operator.Binary {
	return operator.NewBinary(i, operator.NotIn, expr.NewParens(expr.NewValue(values)))
}

func (i Int8) NotInPaths(values ...lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.NotIn, expr.NewParens(expr.ToList(values)))
}

func (i Int8) Between(first, second int8) operator.Binary {
	return operator.NewBinary(i, operator.Between, expr.NewParens(expr.NewValue(first).And(expr.NewValue(second))))
}

func (i Int8) BetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.Between, expr.NewParens(operator.NewBinary(first, operator.And, second)))
}

func (i Int8) NotBetween(first, second int8) operator.Binary {
	return operator.NewBinary(i, operator.NotBetween, expr.NewParens(expr.NewValue(first).And(expr.NewValue(second))))
}

func (i Int8) NotBetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.NotBetween, expr.NewParens(operator.NewBinary(first, operator.And, second)))
}
