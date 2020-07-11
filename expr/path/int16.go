package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/set"
	"github.com/weworksandbox/lingo/sql"
)

func NewInt16WithAlias(e lingo.Table, name, alias string) Int16 {
	return Int16{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt16(e lingo.Table, name string) Int16 {
	return NewInt16WithAlias(e, name, "")
}

type Int16 struct {
	entity lingo.Table
	name   string
	alias  string
}

func (i Int16) GetParent() lingo.Table {
	return i.entity
}

func (i Int16) GetName() string {
	return i.name
}

func (i Int16) GetAlias() string {
	return i.alias
}

func (i Int16) As(alias string) Int16 {
	i.alias = alias
	return i
}

func (i Int16) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int16) To(value int16) set.Set {
	return set.NewSet(i, expr.NewValue(value))
}

func (i Int16) ToExpr(setExp lingo.Expression) set.Set {
	return set.NewSet(i, setExp)
}

func (i Int16) Eq(equalTo int16) operator.Binary {
	return operator.NewBinary(i, operator.Eq, expr.NewValue(equalTo))
}

func (i Int16) EqPath(equalTo lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.Eq, equalTo)
}

func (i Int16) NotEq(notEqualTo int16) operator.Binary {
	return operator.NewBinary(i, operator.NotEq, expr.NewValue(notEqualTo))
}

func (i Int16) NotEqPath(notEqualTo lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.NotEq, notEqualTo)
}

func (i Int16) LT(lessThan int16) operator.Binary {
	return operator.NewBinary(i, operator.LessThan, expr.NewValue(lessThan))
}

func (i Int16) LTPath(lessThan lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.LessThan, lessThan)
}

func (i Int16) LTOrEq(lessThanOrEqual int16) operator.Binary {
	return operator.NewBinary(i, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (i Int16) LTOrEqPath(lessThanOrEqual lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int16) GT(greaterThan int16) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (i Int16) GTPath(greaterThan lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThan, greaterThan)
}

func (i Int16) GTOrEq(greaterThanOrEqual int16) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (i Int16) GTOrEqPath(greaterThanOrEqual lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int16) IsNull() operator.Unary {
	return operator.NewUnary(i, operator.Null)
}

func (i Int16) IsNotNull() operator.Unary {
	return operator.NewUnary(i, operator.NotNull)
}

func (i Int16) In(values ...int16) operator.Binary {
	return operator.NewBinary(i, operator.In, expr.NewParens(expr.NewValue(values)))
}

func (i Int16) InPaths(values ...lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.In, expr.NewParens(expr.ToList(values)))
}

func (i Int16) NotIn(values ...int16) operator.Binary {
	return operator.NewBinary(i, operator.NotIn, expr.NewParens(expr.NewValue(values)))
}

func (i Int16) NotInPaths(values ...lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.NotIn, expr.NewParens(expr.ToList(values)))
}

func (i Int16) Between(first, second int16) operator.Binary {
	return operator.NewBinary(i, operator.Between, expr.NewParens(expr.NewValue(first).And(expr.NewValue(second))))
}

func (i Int16) BetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.Between, expr.NewParens(operator.NewBinary(first, operator.And, second)))
}

func (i Int16) NotBetween(first, second int16) operator.Binary {
	return operator.NewBinary(i, operator.NotBetween, expr.NewParens(expr.NewValue(first).And(expr.NewValue(second))))
}

func (i Int16) NotBetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.NotBetween, expr.NewParens(operator.NewBinary(first, operator.And, second)))
}
