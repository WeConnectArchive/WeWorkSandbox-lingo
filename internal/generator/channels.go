package generator

import (
	"reflect"
)

func pipeErrors(errorsOut chan<- error, errorsIn <-chan error) {
	for err := range errorsIn {
		errorsOut <- err
	}
}

func forEachPipeErrors(inputChannel interface{}, errorsIn <-chan error, errorsOut chan<- error, doWork func(interface{})) bool {
	var cases = []reflect.SelectCase{
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(inputChannel),
		},
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(errorsIn),
		},
	}

	var hadErrors = false
	for len(cases) > 0 {
		i, v, ok := reflect.Select(cases)
		if !ok {
			cases = append(cases[:i], cases[i+1:]...)
			continue
		}

		if cases[i].Chan == reflect.ValueOf(inputChannel) {
			doWork(v.Interface())
		} else {
			hadErrors = true
			errorsOut <- v.Interface().(error)
		}
	}
	return hadErrors
}
