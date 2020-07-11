package operator

import (
	"errors"
	"fmt"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/check"
	"github.com/weworksandbox/lingo/sql"
)

func NewBinary(left lingo.Expression, op Operator, right lingo.Expression) Binary {
	return Binary{
		left:  left,
		op:    op,
		right: right,
	}
}

type Binary struct {
	left  lingo.Expression
	op    Operator
	right lingo.Expression
}

func (b Binary) And(exp lingo.Expression) lingo.ComboExpression {
	return NewBinary(b, And, exp)
}

func (b Binary) Or(exp lingo.Expression) lingo.ComboExpression {
	return NewBinary(b, Or, exp)
}

func (b Binary) ToSQL(d lingo.Dialect) (sql.Data, error) {
	operand, ok := d.(Dialect)
	if !ok {
		return nil, fmt.Errorf("dialect '%s' does not support 'operator.Dialect'", d.GetName())
	}

	if check.IsValueNilOrEmpty(b.left) {
		return nil, errors.New("left of operator.Binary cannot be empty")
	}
	leftSQL, err := b.left.ToSQL(d)
	if err != nil {
		return nil, err
	}

	if check.IsValueNilOrEmpty(b.right) {
		return nil, errors.New("right of operator.Binary cannot be empty")
	}
	rightSQL, err := b.right.ToSQL(d)
	if err != nil {
		return nil, err
	}
	return operand.BinaryOperator(leftSQL, b.op, rightSQL)
}
