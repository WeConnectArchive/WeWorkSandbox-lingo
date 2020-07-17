package expr

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/sql"
)

func ToBinary(expr lingo.Expression) Binary {
	return expr.ToSQL
}

type Binary lingo.Expr

func (p Binary) Eq(value []byte) operator.Binary {
	return operator.Eq(p, NewValue(value))
}

func (p Binary) EqPath(exp lingo.Expression) operator.Binary {
	return operator.Eq(p, exp)
}

func (p Binary) NotEq(value []byte) operator.Binary {
	return operator.NotEq(p, NewValue(value))
}

func (p Binary) NotEqPath(exp lingo.Expression) operator.Binary {
	return operator.NotEq(p, exp)
}

func (p Binary) LT(value []byte) operator.Binary {
	return operator.LessThan(p, NewValue(value))
}

func (p Binary) LTPath(exp lingo.Expression) operator.Binary {
	return operator.LessThan(p, exp)
}

func (p Binary) LTOrEq(value []byte) operator.Binary {
	return operator.LessThanOrEqual(p, NewValue(value))
}

func (p Binary) LTOrEqPath(exp lingo.Expression) operator.Binary {
	return operator.LessThanOrEqual(p, exp)
}

func (p Binary) GT(value []byte) operator.Binary {
	return operator.GreaterThan(p, NewValue(value))
}

func (p Binary) GTPath(exp lingo.Expression) operator.Binary {
	return operator.GreaterThan(p, exp)
}

func (p Binary) GTOrEq(value []byte) operator.Binary {
	return operator.GreaterThanOrEqual(p, NewValue(value))
}

func (p Binary) GTOrEqPath(exp lingo.Expression) operator.Binary {
	return operator.GreaterThanOrEqual(p, exp)
}

func (p Binary) IsNull() operator.Unary {
	return operator.IsNull(p)
}

func (p Binary) IsNotNull() operator.Unary {
	return operator.IsNotNull(p)
}

func (p Binary) In(values ...[]byte) operator.Binary {
	return operator.In(p, NewParens(NewValue(values)))
}

func (p Binary) InPaths(exps ...lingo.Expression) operator.Binary {
	return operator.In(p, NewParens(ToList(exps)))
}

func (p Binary) NotIn(values ...[]byte) operator.Binary {
	return operator.NotIn(p, NewParens(NewValue(values)))
}

func (p Binary) NotInPaths(exps ...lingo.Expression) operator.Binary {
	return operator.NotIn(p, NewParens(ToList(exps)))
}

func (p Binary) Between(first, second []byte) operator.Binary {
	return operator.Between(p, NewValue(first), NewValue(second))
}

func (p Binary) BetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.Between(p, first, second)
}

func (p Binary) NotBetween(first, second []byte) operator.Binary {
	return operator.NotBetween(p, NewValue(first), NewValue(second))
}

func (p Binary) NotBetweenPaths(first, second lingo.Expression) operator.Binary {
	return operator.NotBetween(p, first, second)
}

