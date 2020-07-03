package expression

import (
	"reflect"
	"time"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var timeType = reflect.TypeOf(time.Time{})

type Value interface {
	Value(value []interface{}) (sql.Data, error)
}

func NewValue(v interface{}) core.ComboExpression {
	return value{
		value: v,
	}
}

type value struct {
	value interface{}
}

func (v value) And(exp core.Expression) core.ComboExpression {
	return NewOperator(v, operator.And, exp)
}

func (v value) Or(exp core.Expression) core.ComboExpression {
	return NewOperator(v, operator.Or, exp)
}

func (v value) ToSQL(d core.Dialect) (sql.Data, error) {
	constant, ok := d.(Value)
	if !ok {
		return nil, DialectFunctionNotSupported("Value")
	}

	if v.value == nil {
		return nil, ConstantIsNil()
	}

	reflectOfV := reflect.ValueOf(v.value)
	if err := validateOverallKind(reflectOfV); err != nil {
		return nil, err
	}

	splitValues := convertToSlice(reflectOfV)
	return constant.Value(splitValues)
}

func validateOverallKind(reflectOfV reflect.Value) error {
	reflectOfV = removePtrIfExists(reflectOfV)

	var underlyingType = reflectOfV.Type()

	// If this is a slice or array, check the inner type.
	switch reflectOfV.Kind() {
	case reflect.Slice, reflect.Array:
		// Pull out the inner type to validate
		underlyingType = reflectOfV.Type().Elem()
	}
	if !validateSimpleKind(underlyingType) {
		return ValueIsComplexType(reflectOfV.Type())
	}
	return nil
}

// removePtrIfExists so the underlying type can be exposed. This is helpful
// when they want to change the value in a loop for example.
func removePtrIfExists(reflectOfV reflect.Value) reflect.Value {
	switch reflectOfV.Kind() {
	case reflect.Ptr:
		reflectOfV = reflectOfV.Elem()
	}
	return reflectOfV
}

func validateSimpleKind(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Struct:
		// Check to see if it is of type time.Time
		if !t.AssignableTo(timeType) {
			return false
		}
	case reflect.Chan, reflect.Func, reflect.Interface,
		reflect.Invalid, reflect.Map, reflect.Ptr, reflect.UnsafePointer:
		return false
	}
	return true
}

func convertToSlice(value reflect.Value) []interface{} {
	value = removePtrIfExists(value)

	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		// If we have byte slice or array, they must mean they want Binary, so include
		// the entire slice as a single value to do binary lookups.
		if value.Type().Elem().Kind() == reflect.Uint8 {
			break
		}

		var values = make([]interface{}, 0, value.Len())
		for index := 0; index < value.Len(); index++ {
			indexed := value.Index(index)
			values = append(values, indexed.Interface())
		}
		return values
	}
	return []interface{}{value.Interface()}
}
