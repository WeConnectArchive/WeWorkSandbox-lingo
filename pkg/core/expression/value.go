package expression

import (
	"reflect"
	"time"

	"github.com/weworksandbox/lingo/pkg/core"
)

var timeType = reflect.TypeOf(time.Time{})

type Value interface {
	Value(value []interface{}) (core.SQL, error)
}

func NewValue(v interface{}) core.ComboExpression {
	e := &value{
		value: v,
	}
	e.exp = e
	return e
}

type value struct {
	ComboExpression
	value interface{}
}

func (v value) GetSQL(d core.Dialect) (core.SQL, error) {
	constant, ok := d.(Value)
	if !ok {
		return nil, DialectFunctionNotSupported("Value")
	}

	if v.value == nil {
		return nil, ConstantIsNil()
	}

	reflectOfV := reflect.ValueOf(v.value)
	if err := v.validateOverallKind(reflectOfV); err != nil {
		return nil, err
	}

	splitValues := convertToSlice(v.value)
	return constant.Value(splitValues)
}

func (v value) validateOverallKind(reflectOfV reflect.Value) error {
	// If this is a slice or array, check the inner type.
	var reflectType = reflectOfV.Type()
	switch reflectOfV.Kind() {
	case reflect.Slice, reflect.Array:
		// Pull out the inner type to validate
		reflectType = reflectOfV.Type().Elem()
	}
	if !validateSimpleKind(reflectType) {
		return ValueIsComplexType(reflectType)
	}
	return nil
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

func convertToSlice(actual interface{}) []interface{} {
	refVal := reflect.ValueOf(actual)
	switch refVal.Kind() {
	case reflect.Slice, reflect.Array:
		// If we have byte slice or array, they must mean they want Binary, so include
		// the entire slice as a single value to do binary lookups.
		if refVal.Type().Elem().Kind() == reflect.Uint8 {
			break
		}

		var values = make([]interface{}, 0, refVal.Len())
		for index := 0; index < refVal.Len(); index++ {
			indexed := refVal.Index(index)
			values = append(values, indexed.Interface())
		}
		return values
	}
	return []interface{}{actual}
}
