package expr

import (
	"fmt"
)

type Category int

const (
	// CatUnknown is the default Category - used to denote an unset Category.
	CatUnknown Category = -iota

	// CatArithmetic operators can perform arithmetical operations on numeric operands.
	CatArithmetic

	// CatAssignment operators assigns a value to a variable or of a column or field of a table.
	CatAssignment

	// CatBitwise operators perform bit manipulations between two numeric expressions.
	CatBitwise

	// CatComparison (or relational) operators are mathematical symbols which are used to compare two values.
	// The result of a comparison can be TRUE, FALSE, or UNKNOWN.
	CatComparison

	// CatLogical operators are those that are true or false. They return a true or false values to combine one or
	// more true or false values.
	CatLogical

	// CatSet operators merges two queries using different algorithms
	CatSet

	// CatString operators can combine two or more string data types into one expression
	CatString

	// CatUnary operators perform such an operation which contain only one expression.
	CatUnary

	// TODO - Update this name!!!
	// CatOthers just dont know what the heck to call this...
	CatOthers

	// CatLastCategory is the last entry in Category. Allows for iterating over Category.
	CatLastCategory
)

// Operators that belong to this Category
func (c Category) Operators() []Operator {
	ops := make([]Operator, len(categoryToOperators[c]))
	copy(ops, categoryToOperators[c])
	return ops
}

// Operator is the type of an Operation
type Operator int

const (
	// OpUnknown is the default Operator - used to denote an unset Operator.
	OpUnknown Operator = -iota

	// CatArithmetic Operators

	OpAddition
	OpSubtraction
	OpMultiplication
	OpDivision
	OpModulo

	// CatAssignment Operators

	OpAssign
	OpTableAlias
	OpColumnAlias

	// CatBitwise Operators

	OpBitwiseAND
	OpBitwiseNOT
	OpBitwiseOR
	OpBitwiseXOR

	// CatComparison Operators

	OpNull
	OpNotNull
	OpEq
	OpNotEq
	OpLessThan
	OpLessThanOrEqual
	OpGreaterThan
	OpGreaterThanOrEqual

	// CatLogical Operators

	OpAnd
	OpOr
	OpNot
	OpIn
	OpNotIn
	OpBetween
	OpNotBetween
	OpAny
	OpAll
	OpSome
	OpExists

	// CatSet Operators

	OpUnion
	OpExcept
	OpIntersect

	// CatString Operators

	OpStringConcat

	// CatUnary Operators

	OpSingleton
	OpNegative
	OpCurrentTimestamp

	// CatOthers Operators

	OpList

	OpLike
	OpNotLike

	// OpLastOperator is the last entry in Operator list. Allows for iterating over Operator.
	OpLastOperation

	// Put aliases for Operator here to not affect OpLastOperator
	// Ex: OpNewName = OpOldName
)

// String serves as the name of the operator. It is a single SnakeCased ASCII word that is used in generating paths.
func (o Operator) String() string {
	return operatorToString[o]
}

// Category of operations this Operator belongs to.
func (o Operator) Category() Category {
	return operatorToCategory[o]
}

// This is used to ensure:
// - all categories have an operator
// - all operators have method strings
// - each operator only exists in one category
// Each Operator string must be a single SnakeCased ASCII word that is used in generating paths.
var _ = checkAllOpsInACategoryWithStrings(map[Category]map[Operator]string{
	CatUnknown: {
		OpUnknown: "Unknown",
	},
	CatArithmetic: {
		OpAddition:       "Addition",
		OpSubtraction:    "Subtraction",
		OpMultiplication: "Multiplication",
		OpDivision:       "Division",
		OpModulo:         "Modulo",
	},
	CatAssignment: {
		OpAssign:      "Assign",
		OpTableAlias:  "TableAlias",
		OpColumnAlias: "ColumnAlias",
	},
	CatBitwise: {
		OpBitwiseAND: "BitwiseAND",
		OpBitwiseNOT: "BitwiseNOT",
		OpBitwiseOR:  "BitwiseOR",
		OpBitwiseXOR: "BitwiseXOR",
	},
	CatComparison: {
		OpNull:               "Null",
		OpNotNull:            "NotNull",
		OpEq:                 "Eq",
		OpNotEq:              "NotEq",
		OpLessThan:           "LessThan",
		OpLessThanOrEqual:    "LessThanOrEqual",
		OpGreaterThan:        "GreaterThan",
		OpGreaterThanOrEqual: "GreaterThanOrEqual",
	},
	CatLogical: {
		OpAnd:        "And",
		OpOr:         "Or",
		OpNot:        "Not",
		OpIn:         "In",
		OpNotIn:      "NotIn",
		OpBetween:    "Between",
		OpNotBetween: "NotBetween",
		OpAny:        "Any",
		OpAll:        "All",
		OpSome:       "Some",
		OpExists:     "Exists",
	},
	CatSet: {
		OpUnion:     "Union",
		OpExcept:    "Except",
		OpIntersect: "Intersect",
	},
	CatString: {
		OpStringConcat: "StringConcat",
	},
	CatUnary: {
		OpSingleton: "Singleton",
		OpNegative:  "Negative",
	},
	CatOthers: {
		OpList:             "List",
		OpLike:             "Like",
		OpNotLike:          "NotLike",
		OpCurrentTimestamp: "CurrentTimestamp",
	},
})

var operatorToString = make(map[Operator]string)
var operatorToCategory = make(map[Operator]Category)
var categoryToOperators = make(map[Category][]Operator)

func checkAllOpsInACategoryWithStrings(catToOps map[Category]map[Operator]string) map[Category]map[Operator]string {
	foundOps := make(map[Operator]Category)
	foundOpStrs := make(map[string]Category)

	for cat := CatUnknown; cat < CatLastCategory; cat++ {
		opsToStr, found := catToOps[cat]
		if !found {
			panic(fmt.Sprintf("Category %d is not in catToOps", cat))
		}

		for op, opStr := range opsToStr {
			// Check if this operation already exists / is in another category
			otherCat, alreadyHaveOp := foundOps[op]
			if alreadyHaveOp {
				panic(fmt.Sprintf("Operator %s is already in Category %d", opStr, otherCat))
			}
			foundOps[op] = cat

			// Check if this operation's string value already exists / is in another category
			otherCat, alreadyHaveOpStr := foundOpStrs[opStr]
			if alreadyHaveOpStr {
				panic(fmt.Sprintf("Operator string %s is already in Category %d", opStr, otherCat))
			}
			foundOpStrs[opStr] = cat

			categoryToOperators[cat] = append(categoryToOperators[cat], op)
			operatorToCategory[op] = cat
			operatorToString[op] = opStr
		}
	}

	for op := OpUnknown; op < OpLastOperation; op++ {
		if _, found := foundOps[op]; !found {
			panic(fmt.Sprintf("Operator value %d was not found in any category", op))
		}
	}
	return catToOps
}
