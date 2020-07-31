package dialect

import (
	"fmt"

	"github.com/weworksandbox/lingo/expr"
)

type Syntax string

// opSyntax is a thin wrapper to help with merge opMap
type opSyntax map[expr.Operator]Syntax

// Merge overwrites any values in this opSyntax
func (o opSyntax) Merge(other opSyntax) opSyntax {
	for k, v := range other {
		o[k] = v
	}
	return o
}

var defaultSyntax = requireSyntaxForEveryOperator(opSyntax{
	// CatArithmetic
	expr.OpAddition:       "{0} + {1}",
	expr.OpSubtraction:    "{0} - {1}",
	expr.OpMultiplication: "{0} * {1}",
	expr.OpDivision:       "{0} / {1}",
	expr.OpModulo:         "{0} % {1}",

	// CatAssignment
	expr.OpAssign:      "{0} = {1}",
	expr.OpTable:       "{1}", // {0} & {1} only
	expr.OpTableAlias:  "{1} AS {2}",
	expr.OpSchema:      "", // {0} only
	expr.OpPath:        "{0}.{1}",
	expr.OpColumnAlias: "{0}.{1} AS {2}",

	// CatBitwise
	expr.OpBitwiseAND: "{0} & {1}",
	expr.OpBitwiseOR:  "{0} | {1}",
	expr.OpBitwiseXOR: "{0} ^ {1}",
	expr.OpBitwiseNOT: "~{0}",

	// CatComparison
	expr.OpEq:                 "{0} = {1}",
	expr.OpNotEq:              "{0} <> {1}",
	expr.OpLessThan:           "{0} < {1}",
	expr.OpGreaterThan:        "{0} > {1}",
	expr.OpLessThanOrEqual:    "{0} <= {1}",
	expr.OpGreaterThanOrEqual: "{0} >= {1}",
	expr.OpIsNull:             "{0} IS NULL",
	expr.OpIsNotNull:          "{0} IS NOT NULL",

	// CatLogical
	expr.OpAnd:        "{0} AND {1}",
	expr.OpOr:         "{0} OR {1}",
	expr.OpNot:        "NOT {0}",
	expr.OpIn:         "{0} IN {1}",
	expr.OpNotIn:      "{0} NOT IN {1}",
	expr.OpBetween:    "{0} BETWEEN {1} AND {2}",
	expr.OpNotBetween: "{0} NOT BETWEEN {1} AND {2}",
	expr.OpAny:        "ANY ({0})",
	expr.OpAll:        "ALL ({0})",
	expr.OpSome:       "SOME ({0})",
	expr.OpExists:     "EXISTS ({0})",

	// CatSet
	expr.OpUnion:     "{0} UNION {1}",
	expr.OpExcept:    "{0} EXCEPT {1}",
	expr.OpIntersect: "{0} INTERSECT {1}",

	// CatString
	expr.OpStringConcat: "CONCAT({0}, {1})",

	// CatUnary
	expr.OpSingleton: "{0}",
	expr.OpNegative:  "-{0}",

	// ???
	expr.OpList:             "{0}, {1}",
	expr.OpCount:            "COUNT({0})",
	expr.OpLike:             "{0} LIKE {1}",
	expr.OpNotLike:          "{0} NOT LIKE {1}",
	expr.OpCurrentTimestamp: "CURRENT_TIMESTAMP",
})

func requireSyntaxForEveryOperator(ops opSyntax) opSyntax {
	for op := expr.OpUnknown + 1; op < expr.OpLastOperation; op++ {
		_, found := ops[op]
		if !found {
			panic(fmt.Sprintf("Operator %d does not have a default Syntax", op))
		}
	}
	return ops
}
