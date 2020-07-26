package main

import (
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/internal/generate"
	"github.com/weworksandbox/lingo/internal/generate/paths"
)

var pathData = []paths.Path{
	{
		Name:     "Binary",
		Filename: "binary.go",
		GoType:   "[]byte",
		Imports: []string{
			generate.PkgLingo,
			generate.PkgExp,
			generate.PkgSet,
			generate.PkgSQL,
		},
		Operators: []expr.Operator{

		},
	},
	{
		Name:     "Bool",
		Filename: "bool.go",
		GoType:   "bool",
		Imports: []string{
			generate.PkgLingo,
			generate.PkgExp,
			generate.PkgSet,
			generate.PkgSQL,
		},
		Operators: []expr.Operator{

		},
	},
	{
		Name:     "Float32",
		Filename: "float32.go",
		GoType:   "float32",
		Imports: []string{
			generate.PkgLingo,
			generate.PkgExp,
			generate.PkgSet,
			generate.PkgSQL,
		},
		Operators: []expr.Operator{

		},
	},
	{
		Name:     "Float64",
		Filename: "float64.go",
		GoType:   "float64",
		Imports: []string{
			generate.PkgLingo,
			generate.PkgExp,
			generate.PkgSet,
			generate.PkgSQL,
		},
		Operators: []expr.Operator{

		},
	},
	{
		Name:     "Int",
		Filename: "int.go",
		GoType:   "int",
		Imports: []string{
			generate.PkgLingo,
			generate.PkgExp,
			generate.PkgSet,
			generate.PkgSQL,
		},
		Operators: []expr.Operator{

		},
	},
	{
		Name:     "Int8",
		Filename: "int8.go",
		GoType:   "int8",
		Imports: []string{
			generate.PkgLingo,
			generate.PkgExp,
			generate.PkgSet,
			generate.PkgSQL,
		},
		Operators: []expr.Operator{

		},
	},
	{
		Name:     "Int16",
		Filename: "int16.go",
		GoType:   "int16",
		Imports: []string{
			generate.PkgLingo,
			generate.PkgExp,
			generate.PkgSet,
			generate.PkgSQL,
		},
		Operators: []expr.Operator{

		},
	},
	{
		Name:     "Int32",
		Filename: "int32.go",
		GoType:   "int32",
		Imports: []string{
			generate.PkgLingo,
			generate.PkgExp,
			generate.PkgSet,
			generate.PkgSQL,
		},
		Operators: []expr.Operator{

		},
	},
	{
		Name:     "Int64",
		Filename: "int64.go",
		GoType:   "int64",
		Imports: []string{
			generate.PkgLingo,
			generate.PkgExp,
			generate.PkgSet,
			generate.PkgSQL,
		},
		Operators: []expr.Operator{

		},
	},
	{
		Name:     "String",
		Filename: "string.go",
		GoType:   "string",
		Imports: []string{
			generate.PkgLingo,
			generate.PkgExp,
			generate.PkgSet,
			generate.PkgSQL,
		},
		Operators: []expr.Operator{

		},
	},
	{
		Name:     "Time",
		Filename: "time.go",
		GoType:   "time.Time",
		Imports: []string{
			"time",
			"",
			generate.PkgLingo,
			generate.PkgExp,
			generate.PkgSet,
			generate.PkgSQL,
		},
		Operators: []expr.Operator{

		},
	},
}
