package operator

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/sql"
)

func NewUnary(exp lingo.Expression, op Operator) Unary {
	return Unary{
		op:  op,
		exp: exp,
	}
}

type Unary struct {
	op  Operator
	exp lingo.Expression
}

func (u Unary) ToSQL(d lingo.Dialect) (sql.Data, error) {
	operand, ok := d.(Dialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'operator.Dialect'", d.GetName())
	}

	if check.IsValueNilOrEmpty(u.exp) {
		return nil, errors.New("left of operator.Unary cannot be empty")
	}
	expSQL, err := u.exp.ToSQL(d)
	if err != nil {
		return nil, err
	}
	return operand.UnaryOperator(expSQL, u.op)
}
