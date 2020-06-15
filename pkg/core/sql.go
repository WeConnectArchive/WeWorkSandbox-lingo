package core

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

type sql struct {
	sql    string
	values []interface{}
}

func NewEmptySQL() SQL {
	return &sql{
		sql:    "",
		values: make([]interface{}, 0),
	}
}

func NewSQL(sqlStr string, values []interface{}) SQL {
	if values == nil {
		values = make([]interface{}, 0)
	}
	return &sql{
		sql:    sqlStr,
		values: values,
	}
}

func NewSQLf(format string, values ...interface{}) SQL {
	if format == "" {
		return NewEmptySQL()
	}
	return NewSQL(fmt.Sprintf(format, values...), nil)
}

func NewSQLString(s string) SQL {
	return NewSQL(s, nil)
}

func (s *sql) String() string {
	if s == nil {
		return ""
	}
	return s.sql
}

func (s *sql) Values() []interface{} {
	if s == nil {
		return nil
	}
	return s.values
}

func (s *sql) AppendSQL(right SQL) SQL {
	if s == nil {
		return right
	}
	return s.AppendString(right.String()).AppendValues(right.Values())
}

func (s *sql) AppendSQLWithSpace(right SQL) SQL {
	if s == nil {
		return right
	}
	s.ensureSingleSpace()
	return s.AppendSQL(right)
}

func (s *sql) AppendStringWithSpace(str string) SQL {
	if s == nil {
		return NewSQLf(" %s", str)
	}
	s.ensureSingleSpace()
	return s.AppendString(str)
}

func (s *sql) AppendFormat(format string, values ...interface{}) SQL {
	if s == nil {
		return NewSQL(fmt.Sprintf(format, values...), nil)
	}
	return s.AppendString(fmt.Sprintf(format, values...))
}

func (s *sql) AppendValuesWithFormat(appendValues []interface{}, format string, values ...interface{}) SQL {
	if s == nil {
		return NewSQL(fmt.Sprintf(format, values...), appendValues)
	}
	return s.AppendFormat(format, values...).AppendValues(appendValues)
}

func (s *sql) AppendString(str string) SQL {
	if s == nil {
		return NewSQL(str, nil)
	}
	return NewSQL(s.String()+str, s.Values())
}

func (s *sql) AppendSQLValues(sql SQL) SQL {
	if s == nil {
		return NewSQL("", sql.Values())
	}
	return s.AppendValues(sql.Values())
}

func (s *sql) AppendValues(values []interface{}) SQL {
	if s == nil {
		return NewEmptySQL()
	}
	return NewSQL(s.String(), append(s.Values(), values...))
}

func (s *sql) CombineWithSeparator(sqls []SQL, separator string) SQL {
	var previousSQL SQL = s
	for _, rangeSQL := range sqls {
		previousSQL = previousSQL.AppendFormat("%s%s", separator, rangeSQL.String())
		previousSQL = previousSQL.AppendValues(rangeSQL.Values())
	}
	return previousSQL
}

func (s *sql) CombinePaths(sqls []SQL) SQL {
	var previousSQL SQL = s
	for i, rangeSQL := range sqls {
		if i == 0 {
			previousSQL = previousSQL.AppendSQL(rangeSQL)
		} else {
			previousSQL = previousSQL.AppendSQLValues(rangeSQL).AppendFormat("%s%s", ", ", rangeSQL.String())
		}
	}
	return previousSQL
}

func (s *sql) SurroundWithParens() SQL {
	return s.SurroundWith("(", ")")
}

func (s *sql) SurroundWith(left string, right string) SQL {
	sql := fmt.Sprintf("%s%s%s", left, s.String(), right)
	return NewSQL(sql, s.Values())
}

func (s *sql) InjectValues() string {
	questionMarkSQL := s.String()
	values := s.Values()
	valuesLen := len(values)

	// Find the locations of all the '?' values and inject proper type
	var currentValueIndex = valuesLen
	for i := len(questionMarkSQL); ; {
		trimmedSQL := questionMarkSQL[:i]
		if i = strings.LastIndex(trimmedSQL, "?"); i != -1 {
			// Decrement since we found a question mark
			currentValueIndex--
			valueToInsert := values[currentValueIndex]

			pre := questionMarkSQL[:i]
			post := questionMarkSQL[i+1:]

			valueStr := stringify(valueToInsert)
			questionMarkSQL = pre + valueStr + post
		} else {
			break
		}
	}

	// Check what we found with what we expected
	if currentValueIndex != 0 {
		panic("somehow got a different number of '?' compared to the number of Values")
	}
	return questionMarkSQL
}

func stringify(v interface{}) string {
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.String:
		return fmt.Sprintf("'%v'", value)
	case reflect.Bool:
		return fmt.Sprintf("'%v'", value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%v", value)

	case reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("'%s'", v)

	case reflect.Array, reflect.Slice:
		if value.Type() == reflect.TypeOf([]byte(nil)) {
			return fmt.Sprintf("BINARY(%X)", value.Bytes())
		}

		var str string
		for index := 0; index < value.Len(); index++ {
			indexedValue := value.Index(index)
			if str != "" {
				str += ", "
			}
			str += stringify(indexedValue)
		}
		return str

	case reflect.Chan, reflect.Func, reflect.Struct, reflect.Uintptr, reflect.Ptr, reflect.UnsafePointer:
		panic(fmt.Sprintf("invalid type for String: %s - %s - %+v", value.Kind(), value.String(), value.Interface()))
	}
	panic("some random unknown `reflect.Kind`!")
}

func (s *sql) ensureSingleSpace() {
	r, _ := utf8.DecodeLastRuneInString(s.sql)
	if r == utf8.RuneError || !unicode.IsSpace(r) {
		s.sql += " "
	}
}
