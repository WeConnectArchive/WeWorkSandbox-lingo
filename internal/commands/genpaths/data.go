package main

import (
	"fmt"

	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/internal/generate"
)

func toOperatorInfo(ops []expr.Operator) map[expr.Operator]generate.OperatorInfo {
	result := make(map[expr.Operator]generate.OperatorInfo, len(ops))
	for _, op := range ops {
		found, ok := opsToInfo[op]
		if !ok {
			panic(fmt.Errorf("missing mapping from expr.Operator to generate.OperatorInfo for %d-%s", op, op))
		}
		result[op] = found
	}
	return result
}

const (
	strValue  = "value"
	strValues = "values"
	strFirst  = "first"
	strSecond = "second"
)

var opsToInfo = map[expr.Operator]generate.OperatorInfo{
	expr.OpIsNull: {
		ArgNames: []string{},
	},
	expr.OpIsNotNull: {
		ArgNames: []string{},
	},
	expr.OpEq: {
		ArgNames: []string{strValue},
	},
	expr.OpNotEq: {
		ArgNames: []string{strValue},
	},
	expr.OpLessThan: {
		ArgNames: []string{strValue},
	},
	expr.OpLessThanOrEqual: {
		ArgNames: []string{strValue},
	},
	expr.OpGreaterThan: {
		ArgNames: []string{strValue},
	},
	expr.OpGreaterThanOrEqual: {
		ArgNames: []string{strValue},
	},
	expr.OpIn: {
		ArgNames: []string{strValues},
	},
	expr.OpNotIn: {
		ArgNames: []string{strValues},
	},
	expr.OpBetween: {
		ArgNames: []string{strFirst, strSecond},
	},
	expr.OpNotBetween: {
		ArgNames: []string{strFirst, strSecond},
	},
	expr.OpStringConcat: {
		ArgNames: []string{strValue},
	},
	expr.OpCurrentTimestamp: {
		ArgNames: []string{},
	},
	expr.OpLike: {
		ArgNames: []string{strValue},
	},
	expr.OpNotLike: {
		ArgNames: []string{strValue},
	},
}

var pathData = generate.Paths{
	Package: generate.PkgShortExpr,
	Imports: []string{
		"time",
		"",
		generate.PkgLingo,
		generate.PkgSQL,
	},
	Paths: []generate.Path{
		{
			Name:   "Binary",
			GoType: "[]byte",
			Operators: toOperatorInfo([]expr.Operator{
				expr.OpIsNull,
				expr.OpIsNotNull,
				expr.OpEq,
				expr.OpNotEq,
				expr.OpLessThan,
				expr.OpLessThanOrEqual,
				expr.OpGreaterThan,
				expr.OpGreaterThanOrEqual,
				//expr.OpIn,
				//expr.OpNotIn,
				expr.OpBetween,
				expr.OpNotBetween,
			}),
		},
		{
			Name:   "Bool",
			GoType: "bool",
			Operators: toOperatorInfo([]expr.Operator{
				expr.OpIsNull,
				expr.OpIsNotNull,
				expr.OpEq,
				expr.OpNotEq,
				//expr.OpIn,
				//expr.OpNotIn,
			}),
		},
		{
			Name:   "Float32",
			GoType: "float32",
			Operators: toOperatorInfo([]expr.Operator{
				expr.OpIsNull,
				expr.OpIsNotNull,
				expr.OpEq,
				expr.OpNotEq,
				expr.OpLessThan,
				expr.OpLessThanOrEqual,
				expr.OpGreaterThan,
				expr.OpGreaterThanOrEqual,
				//expr.OpIn,
				//expr.OpNotIn,
				expr.OpBetween,
				expr.OpNotBetween,
			}),
		},
		{
			Name:   "Float64",
			GoType: "float64",
			Operators: toOperatorInfo([]expr.Operator{
				expr.OpIsNull,
				expr.OpIsNotNull,
				expr.OpEq,
				expr.OpNotEq,
				expr.OpLessThan,
				expr.OpLessThanOrEqual,
				expr.OpGreaterThan,
				expr.OpGreaterThanOrEqual,
				//expr.OpIn,
				//expr.OpNotIn,
				expr.OpBetween,
				expr.OpNotBetween,
			}),
		},
		{
			Name:   "Int",
			GoType: "int",
			Operators: toOperatorInfo([]expr.Operator{
				expr.OpIsNull,
				expr.OpIsNotNull,
				expr.OpEq,
				expr.OpNotEq,
				expr.OpLessThan,
				expr.OpLessThanOrEqual,
				expr.OpGreaterThan,
				expr.OpGreaterThanOrEqual,
				//expr.OpIn,
				//expr.OpNotIn,
				expr.OpBetween,
				expr.OpNotBetween,
			}),
		},
		{
			Name:   "Int8",
			GoType: "int8",
			Operators: toOperatorInfo([]expr.Operator{
				expr.OpIsNull,
				expr.OpIsNotNull,
				expr.OpEq,
				expr.OpNotEq,
				expr.OpLessThan,
				expr.OpLessThanOrEqual,
				expr.OpGreaterThan,
				expr.OpGreaterThanOrEqual,
				//expr.OpIn,
				//expr.OpNotIn,
				expr.OpBetween,
				expr.OpNotBetween,
			}),
		},
		{
			Name:   "Int16",
			GoType: "int16",
			Operators: toOperatorInfo([]expr.Operator{
				expr.OpIsNull,
				expr.OpIsNotNull,
				expr.OpEq,
				expr.OpNotEq,
				expr.OpLessThan,
				expr.OpLessThanOrEqual,
				expr.OpGreaterThan,
				expr.OpGreaterThanOrEqual,
				//expr.OpIn,
				//expr.OpNotIn,
				expr.OpBetween,
				expr.OpNotBetween,
			}),
		},
		{
			Name:   "Int32",
			GoType: "int32",
			Operators: toOperatorInfo([]expr.Operator{
				expr.OpIsNull,
				expr.OpIsNotNull,
				expr.OpEq,
				expr.OpNotEq,
				expr.OpLessThan,
				expr.OpLessThanOrEqual,
				expr.OpGreaterThan,
				expr.OpGreaterThanOrEqual,
				//expr.OpIn,
				//expr.OpNotIn,
				expr.OpBetween,
				expr.OpNotBetween,
			}),
		},
		{
			Name:   "Int64",
			GoType: "int64",
			Operators: toOperatorInfo([]expr.Operator{
				expr.OpIsNull,
				expr.OpIsNotNull,
				expr.OpEq,
				expr.OpNotEq,
				expr.OpLessThan,
				expr.OpLessThanOrEqual,
				expr.OpGreaterThan,
				expr.OpGreaterThanOrEqual,
				//expr.OpIn,
				//expr.OpNotIn,
				expr.OpBetween,
				expr.OpNotBetween,
			}),
		},
		{
			Name:   "String",
			GoType: "string",
			Operators: toOperatorInfo([]expr.Operator{
				expr.OpIsNull,
				expr.OpIsNotNull,
				expr.OpEq,
				expr.OpNotEq,
				expr.OpLessThan,
				expr.OpLessThanOrEqual,
				expr.OpGreaterThan,
				expr.OpGreaterThanOrEqual,
				//expr.OpIn,
				//expr.OpNotIn,
				expr.OpBetween,
				expr.OpNotBetween,
				expr.OpStringConcat,
			}),
		},
		{
			Name:   "Time",
			GoType: "time.Time",
			Operators: toOperatorInfo([]expr.Operator{
				expr.OpIsNull,
				expr.OpIsNotNull,
				expr.OpEq,
				expr.OpNotEq,
				expr.OpLessThan,
				expr.OpLessThanOrEqual,
				expr.OpGreaterThan,
				expr.OpGreaterThanOrEqual,
				//expr.OpIn,
				//expr.OpNotIn,
				expr.OpBetween,
				expr.OpNotBetween,
			}),
		},
	},
}
