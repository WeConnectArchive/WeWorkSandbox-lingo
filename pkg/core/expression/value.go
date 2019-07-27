package expression

import (
	"reflect"
	"time"

	"github.com/weworksandbox/lingo/pkg/core"
)

type Value interface {
	Value(value interface{}) (core.SQL, error)
}

func NewValues(v interface{}) []core.Expression {
	value := reflect.ValueOf(v)
	// For slices and arrays, extract each element out
	if value.Kind() == reflect.Array || value.Kind() == reflect.Slice {
		// If we have byte slice or array, they must mean they want Binary, so include
		// the entire slice as a single value to do binary lookups.
		if value.Type().Elem().Kind() == reflect.Uint8 {
			return []core.Expression{NewValue(v)}
		}

		var values = make([]core.Expression, 0, value.Len())
		for index := 0; index < value.Len(); index++ {
			indexed := value.Index(index)
			values = append(values, NewValue(indexed.Interface()))
		}
		return values
	}
	// TODO - We have a single value in our new values, this is incorrect!
	// Hope for the best and just return one value?
	return []core.Expression{NewValue(v)}
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

	reflectOfV := reflect.ValueOf(v.value)
	if v.value == nil {
		return nil, ConstantIsNil()
	}
	switch reflectOfV.Kind() {
	case reflect.Slice, reflect.Array:
		if reflectOfV.Type().Elem().Kind() != reflect.Uint8 {
			return nil, ValueIsComplexType(reflectOfV)
		}
	case reflect.Struct:
		if !reflectOfV.Type().AssignableTo(timeType) {
			return nil, ValueIsComplexType(reflectOfV)
		}
	case reflect.Chan, reflect.Func, reflect.Interface,
		reflect.Invalid, reflect.Map, reflect.Ptr, reflect.UnsafePointer:
		return nil, ValueIsComplexType(reflectOfV)
	}

	return constant.Value(v.value)
}

var timeType = reflect.TypeOf(time.Time{})
