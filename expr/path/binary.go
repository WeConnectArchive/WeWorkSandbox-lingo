package path

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/set"
	"github.com/weworksandbox/lingo/sql"
)

func NewBinaryWithAlias(e lingo.Table, name, alias string) Binary {
	return Binary{
		entity: e,
		name:   name,
		alias:  alias,
	}
}

func NewBinary(e lingo.Table, name string) Binary {
	return NewBinaryWithAlias(e, name, "")
}

type Binary struct {
	entity lingo.Table
	name   string
	alias  string
}

func (b Binary) GetParent() lingo.Table {
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

func (b Binary) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, b)
}

func (b Binary) To(value []byte) set.Set {
	return set.NewSet(b, expr.NewValue(value))
}

func (b Binary) ToExpr(setExp lingo.Expression) set.Set {
	return set.NewSet(b, setExp)
}

func (b Binary) Eq(equalTo []byte) operator.Binary {
	return operator.NewBinary(b, operator.Eq, expr.NewValue(equalTo))
}

func (b Binary) EqPath(equalTo lingo.Expression) operator.Binary {
	return operator.NewBinary(b, operator.Eq, equalTo)
}

func (b Binary) NotEq(notEqualTo []byte) operator.Binary {
	return operator.NewBinary(b, operator.NotEq, expr.NewValue(notEqualTo))
}

func (b Binary) NotEqPath(notEqualTo lingo.Expression) operator.Binary {
	return operator.NewBinary(b, operator.NotEq, notEqualTo)
}

func (b Binary) LT(lessThan []byte) operator.Binary {
	return operator.NewBinary(b, operator.LessThan, expr.NewValue(lessThan))
}

func (b Binary) LTPath(lessThan lingo.Expression) operator.Binary {
	return operator.NewBinary(b, operator.LessThan, lessThan)
}

func (b Binary) LTOrEq(lessThanOrEqual []byte) operator.Binary {
	return operator.NewBinary(b, operator.LessThanOrEqual, expr.NewValue(lessThanOrEqual))
}

func (b Binary) LTOrEqPath(lessThanOrEqual lingo.Expression) operator.Binary {
	return operator.NewBinary(b, operator.LessThanOrEqual, lessThanOrEqual)
}

func (b Binary) GT(greaterThan []byte) operator.Binary {
	return operator.NewBinary(b, operator.GreaterThan, expr.NewValue(greaterThan))
}

func (b Binary) GTPath(greaterThan lingo.Expression) operator.Binary {
	return operator.NewBinary(b, operator.GreaterThan, greaterThan)
}

func (b Binary) GTOrEq(greaterThanOrEqual []byte) operator.Binary {
	return operator.NewBinary(b, operator.GreaterThanOrEqual, expr.NewValue(greaterThanOrEqual))
}

func (b Binary) GTOrEqPath(greaterThanOrEqual lingo.Expression) operator.Binary {
	return operator.NewBinary(b, operator.GreaterThanOrEqual, greaterThanOrEqual)
}

func (b Binary) IsNull() operator.Unary {
	return operator.NewUnary(b, operator.Null)
}

func (b Binary) IsNotNull() operator.Unary {
	return operator.NewUnary(b, operator.NotNull)
}

func (b Binary) In(values ...[]byte) operator.Binary {
	return operator.NewBinary(b, operator.In, expr.NewParens(expr.NewValue(values)))
}

func (b Binary) InPaths(values ...lingo.Expression) operator.Binary {
	return operator.NewBinary(b, operator.In, expr.NewParens(expr.ToList(values)))
}

func (b Binary) NotIn(values ...[]byte) operator.Binary {
	return operator.NewBinary(b, operator.NotIn, expr.NewParens(expr.NewValue(values)))
}

func (b Binary) NotInPaths(values ...lingo.Expression) operator.Binary {
	return operator.NewBinary(b, operator.NotIn, expr.NewParens(expr.ToList(values)))
}

func (b Binary) Between(first, second []byte) operator.Binary {
	return operator.NewBinary(b, operator.Between, expr.NewParens(expr.NewValue(first).And(expr.NewValue(second))))
}

func (b Binary) BetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NewBinary(b, operator.Between, expr.NewParens(operator.NewBinary(first, operator.And, second)))
}

func (b Binary) NotBetween(first, second []byte) operator.Binary {
	return operator.NewBinary(b, operator.NotBetween, expr.NewParens(expr.NewValue(first).And(expr.NewValue(second))))
}

func (b Binary) NotBetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NewBinary(b, operator.NotBetween, expr.NewParens(operator.NewBinary(first, operator.And, second)))
}
