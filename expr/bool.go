package expr

import (
	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/sql"
)

func ToBool(exp lingo.Expression) Bool {
	return exp.ToSQL
}

type Bool lingo.Expr

func (p Bool) ToSQL(d lingo.Dialect) (sql.Data, error) {
	return p(d)
}

func (p Bool) Eq(value bool) operator.Binary {
	return operator.Eq(p, NewValue(value))
}

func (p Bool) EqPath(exp lingo.Expression) operator.Binary {
	return operator.Eq(p, exp)
}

func (p Bool) NotEq(value bool) operator.Binary {
	return operator.NotEq(p, NewValue(value))
}

func (p Bool) NotEqPath(exp lingo.Expression) operator.Binary {
	return operator.NotEq(p, exp)
}

func (p Bool) IsNull() operator.Unary {
	return operator.IsNull(p)
}

func (p Bool) IsNotNull() operator.Unary {
	return operator.IsNotNull(p)
}
