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
	return And(b, exp)
}

func (b Binary) Or(exp lingo.Expression) lingo.ComboExpression {
	return And(b, exp)
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

// And creates an AND operator.Binary expression
func And(left, right lingo.Expression) Binary {
	return NewBinary(left, OpAnd, right)
}

// Or creates an OR operator.Binary expression
func Or(left, right lingo.Expression) Binary {
	return NewBinary(left, OpOr, right)
}

// Eq creates an equals operator.Binary expression
func Eq(left, right lingo.Expression) Binary {
	return NewBinary(left, OpEq, right)
}

// NotEq creates an not equal operator.Binary expression
func NotEq(left, right lingo.Expression) Binary {
	return NewBinary(left, OpNotEq, right)
}

// Like creates a like operator.Binary expression
func Like(left, right lingo.Expression) Binary {
	return NewBinary(left, OpLike, right)
}

// NotLike creates a not like operator.Binary expression
func NotLike(left, right lingo.Expression) Binary {
	return NewBinary(left, OpNotLike, right)
}

// LessThan creates a less than operator.Binary expression
func LessThan(left, right lingo.Expression) Binary {
	return NewBinary(left, OpLessThan, right)
}

// LessThanOrEqual creates a less than or equal to operator.Binary expression
func LessThanOrEqual(left, right lingo.Expression) Binary {
	return NewBinary(left, OpLessThanOrEqual, right)
}

// GreaterThan creates a greater than operator.Binary expression
func GreaterThan(left, right lingo.Expression) Binary {
	return NewBinary(left, OpGreaterThan, right)
}

// GreaterThanOrEqual creates a greater than or equal to operator.Binary expression
func GreaterThanOrEqual(left, right lingo.Expression) Binary {
	return NewBinary(left, OpGreaterThanOrEqual, right)
}

// Between creates a between operator.Binary expression, adding the And expression for the first and second values
func Between(left, first, second lingo.Expression) Binary {
	return NewBinary(left, OpBetween, And(first, second))
}

// NotBetween creates a not between operator.Binary expression, adding the And expression for the
// first and second values
func NotBetween(left, first, second lingo.Expression) Binary {
	return NewBinary(left, OpNotBetween, And(first, second))
}

// In creates an in operator.Binary expression
func In(left lingo.Expression, values lingo.Expression) Binary {
	return NewBinary(left, OpIn, values)
}

// NotIn creates a not in operator.Binary expression
func NotIn(left lingo.Expression, values lingo.Expression) Binary {
	return NewBinary(left, OpNotIn, values)
}
