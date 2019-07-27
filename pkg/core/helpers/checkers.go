package helpers

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
