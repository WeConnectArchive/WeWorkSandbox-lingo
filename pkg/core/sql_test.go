package core_test

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/sql"
	"testing"
)

type q interface {
	String() string
	Values() []interface{}
}

type assert struct {
	s string
	v []interface{}
}
func (a assert) String() string {
	return a.s
}
func (a assert) Values() []interface{} {
	return a.v
}

func Benchmark(b *testing.B) {

	benchmarks := []struct {
		focus bool
		name string
		createSQL func() q
		result assert
	}{
		// String
		{
			name: "Core_String",
			createSQL: func() q {
				return core.NewSQLString("sqlString")
			},
			result: assert{
				s: "sqlString",
				v: []interface{}{},
			},
		},
		{
			name: "SQL_String",
			createSQL: func() q {
				return sql.String("sqlString")
			},
			result: assert{
				s: "sqlString",
				v: []interface{}{},
			},
		},
		// Format
		{
			name: "Core_Format",
			createSQL: func() q {
				return core.NewSQLf("%d sqlString %s", 10, "my string")
			},
			result: assert{
				s: "10 sqlString my string",
				v: []interface{}{},
			},
		},
		{
			name: "SQL_Format",
			createSQL: func() q {
				return sql.Format("%d sqlString %s", 10, "my string")
			},
			result: assert{
				s: "10 sqlString my string",
				v: []interface{}{},
			},
		},
		// String_Append
		{
			name: "Core_String_AppendString",
			createSQL: func() q {
				return core.NewSQLString("sqlString").AppendString(" my string")
			},
			result: assert{
				s: "sqlString my string",
				v: []interface{}{},
			},
		},
		{
			name: "SQL_String_AppendString",
			createSQL: func() q {
				return sql.String("sqlString").Append(sql.String(" my string"))
			},
			result: assert{
				s: "sqlString my string",
				v: []interface{}{},
			},
		},
		// String_AppendStringFormat
		{
			name: "Core_String_AppendStringFormat",
			createSQL: func() q {
				return core.NewSQLString("sqlString").AppendFormat(" my string %d", 10)
			},
			result: assert{
				s: "sqlString my string 10",
				v: []interface{}{},
			},
		},
		{
			name: "SQL_String_AppendStringFormat",
			createSQL: func() q {
				return sql.String("sqlString").Append(sql.Format(" my string %d", 10))
			},
			result: assert{
				s: "sqlString my string 10",
				v: []interface{}{},
			},
		},
		// String_String_Append
		{
			name: "Core_String_String_Append",
			createSQL: func() q {
				a := core.NewSQLString("sqlString1")
				b := core.NewSQLString(" sqlString2")
				return a.AppendSQL(b)
			},
			result: assert{
				s: "sqlString1 sqlString2",
				v: []interface{}{},
			},
		},
		{
			name: "SQL_String_String_Append",
			createSQL: func() q {
				a := sql.String("sqlString1")
				b := sql.String(" sqlString2")
				return a.Append(b)
			},
			result: assert{
				s: "sqlString1 sqlString2",
				v: []interface{}{},
			},
		},
		// String_String_Append_ManipulateA
		{
			name: "Core_String_String_Append_ManipulateA",
			createSQL: func() q {
				a := core.NewSQLString("sqlString1")
				b := core.NewSQLString(" sqlString2")
				c := a.AppendSQL(b)
				a.AppendString("otherInvalid")
				return c
			},
			result: assert{
				s: "sqlString1 sqlString2",
				v: []interface{}{},
			},
		},
		{
			name: "SQL_String_String_Append_ManipulateA",
			createSQL: func() q {
				a := sql.String("sqlString1")
				b := sql.String(" sqlString2")
				c := a.Append(b)
				a.Append(sql.String("otherInvalid"))
				return c
			},
			result: assert{
				s: "sqlString1 sqlString2",
				v: []interface{}{},
			},
		},
		// String_String_Append_ManipulateA
		{
			name: "Core_String_String_Append_ManipulateA",
			createSQL: func() q {
				a := core.NewSQLString("sqlString1")
				b := core.NewSQLString(" sqlString2")
				c := a.AppendSQL(b)
				_ = a.AppendString("otherInvalid")
				d := c.AppendSQL(core.NewSQLString("sqlString3").SurroundWithParens())
				return d
			},
			result: assert{
				s: "sqlString1 sqlString2(sqlString3)",
				v: []interface{}{},
			},
		},
		{
			name: "SQL_String_String_Append_ManipulateA",
			createSQL: func() q {
				a := sql.String("sqlString1")
				b := sql.String(" sqlString2")
				c := a.Append(b)
				_ = a.Append(sql.String("otherInvalid"))
				d := c.SurroundAppend("(", ")", sql.String("sqlString3"))
				return d
			},
			result: assert{
				s: "sqlString1 sqlString2(sqlString3)",
				v: []interface{}{},
			},
		},
	}

	focused := false
	for idx := range benchmarks {
		if benchmarks[idx].focus {
			focused = true
			break
		}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			if focused && !bm.focus {
				b.SkipNow()
			}

			for i := 0; i < b.N; i++ {
				sqlQuery := bm.createSQL()
				s := sqlQuery.String()
				v := sqlQuery.Values()

				if s != bm.result.s {
					b.Errorf("Expected: %s\nGot: %s", bm.result.s, s)
				}
				if len(bm.result.v) != len(v) {
					b.Errorf("Expected: %d\nGot: %d", len(bm.result.v), len(v))
				}
				for idx := range bm.result.v {
					if bm.result.v[idx] != v[idx] {
						b.Errorf("Expected: %s\nGot: %s", bm.result.v[idx], v[idx])
					}
				}
			}
		})
	}
}
