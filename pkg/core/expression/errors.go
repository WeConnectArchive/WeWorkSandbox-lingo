package expression

import (
	"fmt"
	"reflect"
)

type dialectFunctionNotSupported struct {
	function string
}

func DialectFunctionNotSupported(function string) error {
	return dialectFunctionNotSupported{function: function}
}
func (e dialectFunctionNotSupported) Error() string {
	return fmt.Sprintf("dialect function '%s' not supported", e.FunctionName())
}
func (e dialectFunctionNotSupported) FunctionName() string { return e.function }

type expresionIsNil struct {
	name string
}

func ExpressionIsNil(name string) error { return expresionIsNil{name: name} }
func (e expresionIsNil) Error() string  { return fmt.Sprintf("expression '%s' cannot be nil", e.Name()) }
func (e expresionIsNil) Name() string   { return e.name }

type constantIsNil struct{}

func ConstantIsNil() error          { return constantIsNil{} }
func (constantIsNil) Error() string { return "constant is nil, use IsNull instead" }

type valueIsComplexType struct {
	value reflect.Type
}

func ValueIsComplexType(value reflect.Type) error { return valueIsComplexType{value: value} }
func (e valueIsComplexType) Error() string {
	return fmt.Sprintf("value is complex type '%s' when it should be a simple type "+
		"or a pointer to a simple type", e.TypeName())
}
func (e valueIsComplexType) TypeName() string { return e.value.String() }

type errorAtSQL struct {
	err error
	sql string
}

func ErrorAroundSQL(err error, sql string) error {
	return errorAtSQL{err: err, sql: sql}
}
func (e errorAtSQL) Error() string {
	return fmt.Sprintf("an error occurred around sql '%s': %s", e.lastChars(), e.Unwrap().Error())
}
func (e errorAtSQL) SQL() string   { return e.sql }
func (e errorAtSQL) Unwrap() error { return e.err }
func (e errorAtSQL) lastChars() string {
	const length = 20
	var sqlLen = len(e.SQL())
	if sqlLen <= length {
		return e.SQL()
	}
	return "..." + e.SQL()[sqlLen-length:]
}

type expressionCannotBeEmpty struct {
	name string
}

func ExpressionCannotBeEmpty(name string) error { return expressionCannotBeEmpty{name: name} }
func (e expressionCannotBeEmpty) Error() string {
	return fmt.Sprintf("expression '%s' cannot be empty", e.Name())
}
func (e expressionCannotBeEmpty) Name() string { return e.name }

type enumIsInvalid struct {
	value interface{}
	name  string
}

func EnumIsInvalid(name string, value interface{}) error {
	return enumIsInvalid{name: name, value: value}
}
func (e enumIsInvalid) Error() string {
	return fmt.Sprintf("value '%s' for enum '%s' is invalid", e.Name(), e.Value())
}
func (e enumIsInvalid) Name() string  { return e.name }
func (e enumIsInvalid) Value() string { return fmt.Sprintf("%d:%s", e.value, e.value) }
