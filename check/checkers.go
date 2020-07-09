package check

import (
	"reflect"
)

func IsValueNilOrEmpty(c interface{}) bool {
	if c == nil {
		return true
	}
	value := reflect.ValueOf(c)
	switch value.Kind() {
	case reflect.Ptr:
		return value.IsNil()
	case reflect.Slice, reflect.Array, reflect.String:
		return value.Len() == 0
	}
	return false
}

func IsValueNilOrBlank(c interface{}) bool {
	if c == nil {
		return true
	}
	value := reflect.ValueOf(c)
	switch value.Kind() {
	case reflect.Ptr:
		return value.IsNil()
	case reflect.String:
		return value.Len() == 0
	}
	return false
}

func AreValuesNilOrBlank(c interface{}) bool {
	if c == nil {
		return true
	}

	value := reflect.ValueOf(c)
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		valuesLen := value.Len()
		for idx := 0; idx < valuesLen; idx++ {
			if IsValueNilOrBlank(value.Index(idx).Interface()) {
				return true
			}
		}
	}
	return IsValueNilOrBlank(c)
}
