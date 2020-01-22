package expression

import (
	"reflect"
	"time"

	"github.com/weworksandbox/lingo/core"
)

type Value interface {
	Value(value interface{}, sql core.SQL) error
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

func (v value) GetSQL(d core.Dialect, sql core.SQL) error {
	constant, ok := d.(Value)
	if !ok {
		return DialectFunctionNotSupported("Value")
	}

	reflectOfV := reflect.ValueOf(v.value)
	if v.value == nil {
		return ConstantIsNil()
	}
	switch reflectOfV.Kind() {
	case reflect.Slice, reflect.Array:
		if reflectOfV.Type().Elem().Kind() != reflect.Uint8 {
			return ValueIsComplexType(reflectOfV)
		}
	case reflect.Struct:
		if !reflectOfV.Type().AssignableTo(timeType) {
			return ValueIsComplexType(reflectOfV)
		}
	case reflect.Chan, reflect.Func, reflect.Interface,
		reflect.Invalid, reflect.Map, reflect.Ptr, reflect.UnsafePointer:
		return ValueIsComplexType(reflectOfV)
	}

	return constant.Value(v.value, sql)
}

var timeType = reflect.TypeOf(time.Time{})
