package sql_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"

	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var _ = Describe("data.go", func() {

	table.DescribeTable("Data",
		func(d sql.Data, strMatcher, valMatcher types.GomegaMatcher) {
			Expect(d.String()).To(strMatcher)
			Expect(d.Values()).To(valMatcher)
		},
		table.Entry(
			"sql.New - nil",
			sql.New("myString", nil),
			Equal("myString"),
			BeEmpty(),
		),
		table.Entry(
			"sql.New - {}",
			sql.New("myString", []interface{}{}),
			Equal("myString"),
			Equal([]interface{}{}),
		),
		table.Entry(
			"sql.New - d,s,f",
			sql.New("myString", []interface{}{0x02, "str", 10.0}),
			Equal("myString"),
			Equal([]interface{}{0x02, "str", 10.0}),
		),
		table.Entry(
			"sql.Newf - s+d,s,f",
			sql.Newf([]interface{}{0x01, "1string", 14.0}, "myString %d, %s, %0.2f", 0x12, "str", 10.0),
			Equal("myString 18, str, 10.00"),
			Equal([]interface{}{0x01, "1string", 14.0}),
		),
		table.Entry(
			"sql.String - myString",
			sql.String("myString"),
			Equal("myString"),
			BeEmpty(),
		),
		table.Entry(
			"sql.Format - s+d,s,f",
			sql.Format("myString %d, %s, %.2f", 0x12, "str", 10.0),
			Equal("myString 18, str, 10.00"),
			BeEmpty(),
		),
		table.Entry(
			"sql.Values - d,s,f",
			sql.Values([]interface{}{0x02, "str", 10.0}),
			BeEmpty(),
			Equal([]interface{}{0x02, "str", 10.0}),
		),
		table.Entry(
			"sql.Join - empty",
			sql.Join("!@", nil),
			BeEmpty(),
			BeEmpty(),
		),
		table.Entry(
			"sql.Join - s - d,s,f",
			sql.Join(
				"!@",
				[]sql.Data{
					sql.String("myString"),
					sql.Values([]interface{}{0x02, "str", 10.0}),
				},
			),
			Equal("myString"),
			Equal([]interface{}{0x02, "str", 10.0}),
		),
		table.Entry(
			"sql.Join - s - d,s,f - s+d,s,f",
			sql.Join(
				"!@",
				[]sql.Data{
					sql.String(""),
					sql.String("myString"),
					sql.Values([]interface{}{0x02, "str", 10.0}),
					sql.Format("myString %d, %s, %.2f", 0x12, "str", 10.0),
					sql.String(""),
					sql.Newf([]interface{}{}, "%s", "myThirdString"),
				},
			),
			Equal("myString!@myString 18, str, 10.00!@myThirdString"),
			Equal([]interface{}{0x02, "str", 10.0}),
		),
		table.Entry(
			"sql.Append - s+d - s+s",
			sql.New("abc1", []interface{}{1}).Append(sql.New("2cba", []interface{}{"a"})),
			Equal("abc12cba"),
			Equal([]interface{}{1, "a"}),
		),
		table.Entry(
			"sql.AppendWithSpace - s+d - s+s",
			sql.New("abc1", []interface{}{1}).AppendWithSpace(sql.New("2cba", []interface{}{"a"})),
			Equal("abc1 2cba"),
			Equal([]interface{}{1, "a"}),
		),
		table.Entry(
			"sql.AppendWithSpace - s+d - s+s",
			sql.New("abc1", []interface{}{1}).SurroundAppend("firstBang!", "!secondBang", sql.New("2cba", []interface{}{"a"})),
			Equal("abc1firstBang!2cba!secondBang"),
			Equal([]interface{}{1, "a"}),
		),
	)
})
