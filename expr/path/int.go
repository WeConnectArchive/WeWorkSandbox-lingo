package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/set"
	"github.com/weworksandbox/lingo/sql"
)

func NewIntWithAlias(e lingo.Table, name, alias string) Int {
	return Int{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewInt(e lingo.Table, name string) Int {
	return NewIntWithAlias(e, name, "")
}

type Int struct {
	entity lingo.Table
	name   string
	alias  string
}

func (i Int) GetParent() lingo.Table {
	return i.entity
}

func (i Int) GetName() string {
	return i.name
}

func (i Int) GetAlias() string {
	return i.alias
}

func (i Int) As(alias string) Int {
	i.alias = alias
	return i
}

func (i Int) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, i)
}

func (i Int) To(value int) set.Set {
	return set.NewSet(i, expr.NewValue(value))
}

func (i Int) ToExpr(setExp lingo.Expression) set.Set {
	return set.NewSet(i, setExp)
}

func (i Int) Eq(equalTo int) operator.Binary {
	return operator.NewBinary(i, operator.Eq, expr.NewValue(equalTo))
}

func (i Int) EqPath(equalTo lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.Eq, equalTo)
}

func (i Int) NotEq(notEqualTo int) operator.Binary {
	return operator.NewBinary(i, operator.NotEq, expr.NewValue(notEqualTo))
}

func (i Int) NotEqPath(notEqualTo lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.NotEq, notEqualTo)
}

func (i Int) LT(lessThan int) operator.Binary {
	return operator.NewBinary(i, operator.LessThan, expr.NewValue(lessThan))
}

func (i Int) LTPath(lessThan lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.LessThan, lessThan)
}

func (i Int) LTOrEq(lessThanOrEqual int) operator.Binary {
	return operator.NewBinary(i, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (i Int) LTOrEqPath(lessThanOrEqual lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.LessThanOrEqual, lessThanOrEqual)
}

func (i Int) GT(greaterThan int) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (i Int) GTPath(greaterThan lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThan, greaterThan)
}

func (i Int) GTOrEq(greaterThanOrEqual int) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (i Int) GTOrEqPath(greaterThanOrEqual lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (i Int) IsNull() operator.Unary {
	return operator.NewUnary(i, operator.Null)
}

func (i Int) IsNotNull() operator.Unary {
	return operator.NewUnary(i, operator.NotNull)
}

func (i Int) In(values ...int) operator.Binary {
	return operator.NewBinary(i, operator.In, expr.NewParens(expr.NewValue(values)))
}

func (i Int) InPaths(values ...lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.In, expr.NewParens(expr.ToList(values)))
}

func (i Int) NotIn(values ...int) operator.Binary {
	return operator.NewBinary(i, operator.NotIn, expr.NewParens(expr.NewValue(values)))
}

func (i Int) NotInPaths(values ...lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.NotIn, expr.NewParens(expr.ToList(values)))
}

func (i Int) Between(first, second int) operator.Binary {
	return operator.NewBinary(i, operator.Between, expr.NewParens(expr.NewValue(first).And(expr.NewValue(second))))
}

func (i Int) BetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.Between, expr.NewParens(operator.NewBinary(first, operator.And, second)))
}

func (i Int) NotBetween(first, second int) operator.Binary {
	return operator.NewBinary(i, operator.NotBetween, expr.NewParens(expr.NewValue(first).And(expr.NewValue(second))))
}

func (i Int) NotBetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NewBinary(i, operator.NotBetween, expr.NewParens(operator.NewBinary(first, operator.And, second)))
}
