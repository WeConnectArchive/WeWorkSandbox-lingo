package matchers

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// NOTE: This entire file was copied and pasted from the Gomega library for use in our custom matchers
// From revision 7615b9433f86a8bdf29709bf288bc4fd0636a369

func isBool(a interface{}) bool {
	return reflect.TypeOf(a).Kind() == reflect.Bool
}

func isNumber(a interface{}) bool {
	if a == nil {
		return false
	}
	kind := reflect.TypeOf(a).Kind()
	return reflect.Int <= kind && kind <= reflect.Float64
}

func isInteger(a interface{}) bool {
	kind := reflect.TypeOf(a).Kind()
	return reflect.Int <= kind && kind <= reflect.Int64
}

func isUnsignedInteger(a interface{}) bool {
	kind := reflect.TypeOf(a).Kind()
	return reflect.Uint <= kind && kind <= reflect.Uint64
}

func isFloat(a interface{}) bool {
	kind := reflect.TypeOf(a).Kind()
	return reflect.Float32 <= kind && kind <= reflect.Float64
}

func toInteger(a interface{}) int64 {
	if isInteger(a) {
		return reflect.ValueOf(a).Int()
	} else if isUnsignedInteger(a) {
		return int64(reflect.ValueOf(a).Uint())
	} else if isFloat(a) {
		return int64(reflect.ValueOf(a).Float())
	}
	panic(fmt.Sprintf("Expected a number!  Got <%T> %#v", a, a))
}

func toUnsignedInteger(a interface{}) uint64 {
	if isInteger(a) {
		return uint64(reflect.ValueOf(a).Int())
	} else if isUnsignedInteger(a) {
		return reflect.ValueOf(a).Uint()
	} else if isFloat(a) {
		return uint64(reflect.ValueOf(a).Float())
	}
	panic(fmt.Sprintf("Expected a number!  Got <%T> %#v", a, a))
}

func toFloat(a interface{}) float64 {
	if isInteger(a) {
		return float64(reflect.ValueOf(a).Int())
	} else if isUnsignedInteger(a) {
		return float64(reflect.ValueOf(a).Uint())
	} else if isFloat(a) {
		return reflect.ValueOf(a).Float()
	}
	panic(fmt.Sprintf("Expected a number!  Got <%T> %#v", a, a))
}

func isError(a interface{}) bool {
	_, ok := a.(error)
	return ok
}

func isChan(a interface{}) bool {
	if isNil(a) {
		return false
	}
	return reflect.TypeOf(a).Kind() == reflect.Chan
}

func isMap(a interface{}) bool {
	if a == nil {
		return false
	}
	return reflect.TypeOf(a).Kind() == reflect.Map
}

func isArrayOrSlice(a interface{}) bool {
	if a == nil {
		return false
	}
	switch reflect.TypeOf(a).Kind() {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}
}

func isString(a interface{}) bool {
	if a == nil {
		return false
	}
	return reflect.TypeOf(a).Kind() == reflect.String
}

func toString(a interface{}) (string, bool) {
	aString, isString := a.(string)
	if isString {
		return aString, true
	}

	aBytes, isBytes := a.([]byte)
	if isBytes {
		return string(aBytes), true
	}

	aStringer, isStringer := a.(fmt.Stringer)
	if isStringer {
		return aStringer.String(), true
	}

	aJSONRawMessage, isJSONRawMessage := a.(json.RawMessage)
	if isJSONRawMessage {
		return string(aJSONRawMessage), true
	}

	return "", false
}

func lengthOf(a interface{}) (int, bool) {
	if a == nil {
		return 0, false
	}
	switch reflect.TypeOf(a).Kind() {
	case reflect.Map, reflect.Array, reflect.String, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(a).Len(), true
	default:
		return 0, false
	}
}

func capOf(a interface{}) (int, bool) {
	if a == nil {
		return 0, false
	}
	switch reflect.TypeOf(a).Kind() {
	case reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(a).Cap(), true
	default:
		return 0, false
	}
}

func isNil(a interface{}) bool {
	if a == nil {
		return true
	}

	switch reflect.TypeOf(a).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(a).IsNil()
	}

	return false
}

func isPtr(a interface{}) bool {
	if a == nil {
		return false
	}
	switch reflect.TypeOf(a).Kind() {
	case reflect.Ptr:
		return true
	}
	return false
}
