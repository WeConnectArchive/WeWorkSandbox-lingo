package sql

import (
	"fmt"
	"strings"
)

func Query(s string, values []interface{}) Data {
	return Data{
		str:    s,
		values: values,
	}
}

func String(s string) Data {
	return Data{
		str: s,
		values: nil,
	}
}

func Format(format string, values ...interface{}) Data {
	return String(fmt.Sprintf(format, values...))
}

func Values(values []interface{}) Data {
	return Data{
		values: values,
	}
}

func Join(sep string, other ...Data) Data {
	// Do the work of strings.Join while loading the values, more efficient.
	var s strings.Builder
	var v = make([]interface{}, 0, len(other))
	for idx := range other {
		v = append(v, other[idx].Values()...) // Ensure to expand!

		if idx > 0 {
			s.WriteString(sep)
		}
		s.WriteString(other[idx].String())
	}
	return Data{
		values: v,
		str: s.String(),
	}
}

const (
	space = " "
)

type Data struct {
	str string
	values []interface{}
}

func (d Data) String() string {
	return d.str
}
func (d Data) Values() []interface{} {
	return d.values
}

func (d Data) Append(other Data) Data {
	return Data{
		values: combineValues(d.values, other.values),
		str:    d.str + other.String(),
	}
}

func (d Data) AppendWithSpace(other Data) Data {
	return Data{
		values: combineValues(d.values, other.values),
		str:    d.str + space + other.String(),
	}
}

func (d Data) SurroundAppend(l, r string, other Data) Data {
	return Data{
		values: combineValues(d.values, other.Values()),
		str:    d.str + l + other.String() + r,
	}
}

// combineValues creates a final array and copies into it
func combineValues(v1, v2 []interface{}) []interface{} {
	v := make([]interface{}, len(v1) + len(v2))
	copy(v, v1)
	copy(v, v2)
	return v
}


