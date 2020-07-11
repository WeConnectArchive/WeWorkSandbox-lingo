// Code generated by an internal Lingo tool, genpaths.go - DO NOT EDIT

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

func (p Int) GetParent() lingo.Table {
	return p.entity
}

func (p Int) GetName() string {
	return p.name
}

func (p Int) GetAlias() string {
	return p.alias
}

func (p Int) As(alias string) Int {
	p.alias = alias
	return p
}

func (p Int) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return ExpandColumnWithDialect(d, p)
}

func (p Int) To(value int) set.Set {
	return set.NewSet(p, expr.NewValue(value))
}

func (p Int) ToExpr(exp lingo.Expression) set.Set {
	return set.NewSet(p, exp)
}

func (p Int) Eq(value int) operator.Binary {
	return operator.NewBinary(p, operator.Eq, expr.NewValue(value))
}

func (p Int) EqPath(exp lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.Eq, exp)
}

func (p Int) NotEq(value int) operator.Binary {
	return operator.NewBinary(p, operator.NotEq, expr.NewValue(value))
}

func (p Int) NotEqPath(exp lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.NotEq, exp)
}

func (p Int) LT(value int) operator.Binary {
	return operator.NewBinary(p, operator.LessThan, expr.NewValue(value))
}

func (p Int) LTPath(exp lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.LessThan, exp)
}

func (p Int) LTOrEq(value int) operator.Binary {
	return operator.NewBinary(p, operator.LessThanOrEqual, expr.NewValue(value))
}

func (p Int) LTOrEqPath(exp lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.LessThanOrEqual, exp)
}

func (p Int) GT(value int) operator.Binary {
	return operator.NewBinary(p, operator.GreaterThan, expr.NewValue(value))
}

func (p Int) GTPath(exp lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.GreaterThan, exp)
}

func (p Int) GTOrEq(value int) operator.Binary {
	return operator.NewBinary(p, operator.GreaterThanOrEqual, expr.NewValue(value))
}

func (p Int) GTOrEqPath(exp lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.GreaterThanOrEqual, exp)
}

func (p Int) IsNull() operator.Unary {
	return operator.NewUnary(p, operator.Null)
}

func (p Int) IsNotNull() operator.Unary {
	return operator.NewUnary(p, operator.NotNull)
}

func (p Int) In(values ...int) operator.Binary {
	return operator.NewBinary(p, operator.In, expr.NewParens(expr.NewValue(values)))
}

func (p Int) InPaths(exps ...lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.In, expr.NewParens(expr.ToList(exps)))
}

func (p Int) NotIn(values ...int) operator.Binary {
	return operator.NewBinary(p, operator.NotIn, expr.NewParens(expr.NewValue(values)))
}

func (p Int) NotInPaths(exps ...lingo.Expression) operator.Binary {
	return operator.NewBinary(p, operator.NotIn, expr.NewParens(expr.ToList(exps)))
}

func (p Int) Between(first, second int) operator.Binary {
	and := expr.NewParens(expr.NewValue(first).And(expr.NewValue(second)))
	return operator.NewBinary(p, operator.Between, and)
}

func (p Int) BetweenPaths(firstExp, secondExp lingo.Expression) operator.Binary {
	and := expr.NewParens(operator.NewBinary(firstExp, operator.And, secondExp))
	return operator.NewBinary(p, operator.Between, and)
}

func (p Int) NotBetween(first, second int) operator.Binary {
	and := expr.NewParens(expr.NewValue(first).And(expr.NewValue(second)))
	return operator.NewBinary(p, operator.NotBetween, and)
}

func (p Int) NotBetweenPaths(firstExp, secondExp lingo.Expression) operator.Binary {
	and := expr.NewParens(operator.NewBinary(firstExp, operator.And, secondExp))
	return operator.NewBinary(p, operator.NotBetween, and)
}
