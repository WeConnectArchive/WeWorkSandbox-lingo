package expr

import (
	"github.com/weworksandbox/lingo"
)

// List creates a comma separated left, right
func List(left, right lingo.Expression) Operation {
	return NewOperation(OpList, left, right)
}

// Assign toValue to exp for a setter operation (updates)
func Assign(exp lingo.Expression, toValue lingo.Expression) Operation {
	return NewOperation(OpAssign, exp, toValue)
}

// ====================================================================
// TODO - other types above

// CurrentTimestamp represents a CurrentTimestamp function
func CurrentTimestamp() Operation {
	return NewOperation(OpCurrentTimestamp)
}

// IsNull creates an Operation "null" expression (not literally a nil lingo.Expression!)
func IsNull(exp lingo.Expression) Operation {
	return NewOperation(OpNull, exp)
}

// IsNotNull creates a not null Operation expression
func IsNotNull(exp lingo.Expression) Operation {
	return NewOperation(OpNotNull, exp)
}

// And creates an AND Operation expression
func And(left, right lingo.Expression) Operation {
	return NewOperation(OpAnd, left, right)
}

// Or creates an OR Operation expression
func Or(left, right lingo.Expression) Operation {
	return NewOperation(OpOr, left, right)
}

// Eq creates an equals Operation expression
func Eq(left, right lingo.Expression) Operation {
	return NewOperation(OpEq, left, right)
}

// NotEq creates an not equal Operation expression
func NotEq(left, right lingo.Expression) Operation {
	return NewOperation(OpNotEq, left, right)
}

// Like creates a like Operation expression
func Like(left, right lingo.Expression) Operation {
	return NewOperation(OpLike, left, right)
}

// NotLike creates a not like Operation expression
func NotLike(left, right lingo.Expression) Operation {
	return NewOperation(OpNotLike, left, right)
}

// LessThan creates a less than Operation expression
func LessThan(left, right lingo.Expression) Operation {
	return NewOperation(OpLessThan, left, right)
}

// LessThanOrEqual creates a less than or equal to Operation expression
func LessThanOrEqual(left, right lingo.Expression) Operation {
	return NewOperation(OpLessThanOrEqual, left, right)
}

// GreaterThan creates a greater than Operation expression
func GreaterThan(left, right lingo.Expression) Operation {
	return NewOperation(OpGreaterThan, left, right)
}

// GreaterThanOrEqual creates a greater than or equal to Operation expression
func GreaterThanOrEqual(left, right lingo.Expression) Operation {
	return NewOperation(OpGreaterThanOrEqual, left, right)
}

// Between creates a between Operation expression
func Between(left, first, second lingo.Expression) Operation {
	return NewOperation(OpBetween, left, first, second)
}

// NotBetween creates a not between Operation expression
func NotBetween(left, first, second lingo.Expression) Operation {
	return NewOperation(OpNotBetween, left, first, second)
}

// In creates an in Operation expression
func In(left lingo.Expression, values []lingo.Expression) Operation {
	return NewOperation(OpIn, append([]lingo.Expression{left}, values...)...)
}

// NotIn creates a not in Operation expression
func NotIn(left lingo.Expression, values []lingo.Expression) Operation {
	return NewOperation(OpNotIn, append([]lingo.Expression{left}, values...)...)
}
