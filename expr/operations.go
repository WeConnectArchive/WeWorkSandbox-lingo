package expr

import (
	"github.com/weworksandbox/lingo"
)

func Schema() lingo.Expression {
	return Operation(OpSchema)
}

func Table(table lingo.Table) lingo.Expression {
	name := TableName(table)
	if alias := table.GetAlias(); alias != nil {
		return Operation(OpTableAlias, Schema(), name, alias)
	}
	return Operation(OpTable, Schema(), name)
}

func TableName(table lingo.Table) lingo.Expression {
	return Operation(OpSingleton, Lit(table.GetTableName()))
}

func Column(table lingo.Table, name lingo.Expression) lingo.Expression {
	var parent = table.GetName()
	if alias := table.GetAlias(); alias != nil {
		parent = alias
	}
	return Path(parent, name)
}

func Path(parent, name lingo.Expression) lingo.Expression {
	return Operation(OpPath, parent, name)
}

// List creates a comma separated left, right
func List(left, right lingo.Expression) lingo.Expression {
	return Operation(OpList, left, right)
}

// Assign toValue to exp for a setter operation (updates)
func Assign(exp lingo.Expression, toValue lingo.Expression) lingo.Expression {
	return Operation(OpAssign, exp, toValue)
}

// ====================================================================
// TODO - other types above

// And creates an AND Operation expression
func And(left, right lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpAnd, left, right))
}

// Or creates an OR Operation expression
func Or(left, right lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpOr, left, right))
}

// IsNull creates an Operation "null" expression (not literally a nil lingo.Expression!)
func IsNull(exp lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpIsNull, exp))
}

// IsNotNull creates a not null Operation expression
func IsNotNull(exp lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpIsNotNull, exp))
}

// Eq creates an equals Operation expression
func Eq(left, right lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpEq, left, right))
}

// NotEq creates an not equal Operation expression
func NotEq(left, right lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpNotEq, left, right))
}

// Like creates a like Operation expression
func Like(left, right lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpLike, left, right))
}

// NotLike creates a not like Operation expression
func NotLike(left, right lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpNotLike, left, right))
}

// LessThan creates a less than Operation expression
func LessThan(left, right lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpLessThan, left, right))
}

// LessThanOrEqual creates a less than or equal to Operation expression
func LessThanOrEqual(left, right lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpLessThanOrEqual, left, right))
}

// GreaterThan creates a greater than Operation expression
func GreaterThan(left, right lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpGreaterThan, left, right))
}

// GreaterThanOrEqual creates a greater than or equal to Operation expression
func GreaterThanOrEqual(left, right lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpGreaterThanOrEqual, left, right))
}

// Between creates a between Operation expression
func Between(left, first, second lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpBetween, left, first, second))
}

// NotBetween creates a not between Operation expression
func NotBetween(left, first, second lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpNotBetween, left, first, second))
}

func StringConcat(left, right lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpStringConcat, left, right))
}

func Count(exp lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpCount, exp))
}

// In creates an in Operation expression
func In(left lingo.Expression, values []lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpIn, append([]lingo.Expression{left}, values...)...))
}

// NotIn creates a not in Operation expression
func NotIn(left lingo.Expression, values []lingo.Expression) lingo.ComboExpression {
	return ComboOperation(Operation(OpNotIn, append([]lingo.Expression{left}, values...)...))
}

// CurrentTimestamp represents a CurrentTimestamp function
func CurrentTimestamp() lingo.ComboExpression {
	return ComboOperation(Operation(OpCurrentTimestamp))
}
