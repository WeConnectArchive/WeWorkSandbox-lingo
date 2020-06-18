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
func Empty() Data { return empty }

var empty = stringData("")

func New(s string, values []interface{}) Data {
	if len(values) == 0 && len(s) == 0 {
		return Empty()
	}
	if len(s) == 0 {
		return Values(values)
	}
	if len(values) == 0 {
		return String(s)
	}
	return data{
		str:    s,
		values: values,
	}
}

func Newf(values []interface{}, format string, a ...interface{}) Data {
	return New(fmt.Sprintf(format, a...), values)
}

func String(s string) Data {
	if s == "" {
		return Empty()
	}
	return stringData(s)
}

func Format(format string, a ...interface{}) Data {
	return String(fmt.Sprintf(format, a...))
}

func Values(values []interface{}) Data {
	if len(values) == 0 {
		return Empty()
	}
	return valuesData(values)
}

func Surround(l, r string, d Data) Data {
	// Point to the same array. Data should be immutable - minus random slice shenanigans.
	return New(l+d.String()+r, d.Values())
}

// Join will combine the Data with each other with sep in between each String value.
// It will not append empty String values, but will append the Values, if any.
func Join(sep string, d []Data) Data {
	// Determine final lengths first
	var sLen int
	var vLen int
	for idx := range d {
		sLen += len(d[idx].String())
		vLen += len(d[idx].Values())
	}

	// Setup result buffers
	var v = make([]interface{}, 0, vLen)
	var s strings.Builder

	// Grow to the combined text length, plus the required number of separators.
	// Note that this can be more than what needs to be allocated due to possible empty strings
	sepCount := sLen + len(sep)*(len(d)-1)
	if sepCount > 0 {
		s.Grow(sepCount)
	}

	// Build result buffers
	for idx := range d {
		str := d[idx].String()

		// Only add if we have data, if not, its just empty, might have values still though.
		if len(str) > 0 {
			// Only add the separator if the str is non-empty & the result buffer already has other data.
			// We do not want to put the separate at the beginning of it.
			if s.Len() > 0 {
				_, _ = s.WriteString(sep)
			}
			_, _ = s.WriteString(str)
		}

		// Ensure we do not append nils
		if values := d[idx].Values(); len(values) > 0 {
			// Ensure to expand both of these! Without, it will put all the Data
			// values (for one Data) into a single top level entry.
			// Example if you dont expand:
			// []interface{}{ []interface{}{data1for1, data2for1}, []interface{}{data1for2} }
			v = append(v, values...)
		}
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
	return New(
		d.str+other.String(),
		copyValues(d.values, other.Values()),
	)
}

func (d data) AppendWithSpace(other Data) Data {
	return New(
		d.str+space+other.String(),
		copyValues(d.values, other.Values()),
	)
}

func (d data) SurroundAppend(l, r string, other Data) Data {
	return New(
		d.str+l+other.String()+r,
		copyValues(d.values, other.Values()),
	)
}

type stringData string

func (d stringData) Append(other Data) Data {
	s := string(d) + other.String()
	return New(s, other.Values())
}

func (d stringData) AppendWithSpace(other Data) Data {
	s := string(d) + space + other.String()
	return New(s, other.Values())
}

func (d stringData) SurroundAppend(l, r string, other Data) Data {
	s := string(d) + l + other.String() + r
	return New(s, other.Values())
}

func (d stringData) String() string      { return string(d) }
func (stringData) Values() []interface{} { return nil }

type valuesData []interface{}

func (v valuesData) Append(other Data) Data {
	return New(other.String(), copyValues(v, other.Values()))
}

func (v valuesData) AppendWithSpace(other Data) Data {
	return New(space+other.String(), copyValues(v, other.Values()))
}

func (v valuesData) SurroundAppend(l, r string, other Data) Data {
	return New(l+other.String()+r, copyValues(v, other.Values()))
}

func (valuesData) String() string          { return "" }
func (v valuesData) Values() []interface{} { return v }

// copyValues creates a new backing slice and copies the values of both slices into it.
func copyValues(v1, v2 []interface{}) []interface{} {
	lv1 := len(v1)

	v := make([]interface{}, lv1+len(v2))
	copy(v, v1)
	copy(v[lv1:], v2)
	return v
}
