// Code generated by Lingo for table information_schema.GLOBAL_VARIABLES - DO NOT EDIT

package qglobalvariables

import "github.com/weworksandbox/lingo/core/path"

var instance = New()

func Q() QGlobalVariables {
	return instance
}

func VariableName() path.StringPath {
	return instance.variableName
}

func VariableValue() path.StringPath {
	return instance.variableValue
}
