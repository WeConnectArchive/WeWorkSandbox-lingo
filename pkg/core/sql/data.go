package sql

import (
	"fmt"
	"strings"
)

type Data interface {
	Append(other Data) Data
	AppendWithSpace(other Data) Data
	SurroundAppend(l, r string, other Data) Data

	String() string
	Values() []interface{}
}

// Empty can be used before a for loop to initialize your Data
func Empty() Data { return String("") }

func New(s string, values []interface{}) Data {
	return data{
		str:    s,
		values: values,
	}
}

func Newf(values []interface{}, format string, a ...interface{}) Data {
	return New(fmt.Sprintf(format, a...), values)
}

func String(s string) Data {
	return data{
		str: s,
	}
}

func Format(format string, a ...interface{}) Data {
	return String(fmt.Sprintf(format, a...))
}

func Values(values []interface{}) Data {
	return data{
		values: values,
	}
}

func Surround(l, r string, d Data) Data {
	return data{
		str:    l + d.String() + r,
		values: d.Values(), // Point to the same array. Data should be immutable - minus random slice shenanigans.
	}
}

// Join will combine the Data with each other with sep in between each String value.
// It will not append empty String values, but will append the Values, if any.
func Join(sep string, other []Data) Data {
	var s strings.Builder
	var v []interface{}
	for idx := range other {
		// Copy and load the values
		otherVals := other[idx].Values()

		// Ensure to expand both of these! Without, it will put all the Data
		// values (for one Data) into a single top level entry.
		// Example if you dont expand:
		// []interface{}{ []interface{}{data1for1, data2for1}, []interface{}{data1for2} }
		copiedVals := append(otherVals[:0:0], otherVals...)
		v = append(v, copiedVals...)

		// Load the strings
		otherStr := other[idx].String()
		if len(otherStr) == 0 {
			continue
		}

		// Only add the separator if there is data already in the builder.
		if s.Len() > 0 {
			_, _ = s.WriteString(sep)
		}
		_, _ = s.WriteString(otherStr)
	}
	return data{
		values: v,
		str:    s.String(),
	}
}

const (
	space = " "
)

// TODO - Data might be able to be replaced with some sort of linked
// list type construct. Could be more useful than copying strings
// and backed value slices.

type data struct {
	str    string
	values []interface{}
}

func (d data) String() string {
	return d.str
}
func (d data) Values() []interface{} {
	return d.values
}

func (d data) Append(other Data) Data {
	return data{
		values: combineValues(d.values, other.Values()),
		str:    d.str + other.String(),
	}
}

func (d data) AppendWithSpace(other Data) Data {
	return data{
		values: combineValues(d.values, other.Values()),
		str:    d.str + space + other.String(),
	}
}

func (d data) SurroundAppend(l, r string, other Data) Data {
	return data{
		values: combineValues(d.values, other.Values()),
		str:    d.str + l + other.String() + r,
	}
}

// combineValues creates a final array and copies into it
func combineValues(v1, v2 []interface{}) []interface{} {
	lv1 := len(v1)
	lv2 := len(v2)

	v := make([]interface{}, lv1+lv2)
	copy(v, v1)
	copy(v[lv1:], v2)
	return v
}
