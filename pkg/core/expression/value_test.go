package expression_test

import (
	"errors"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"

	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var _ = Describe("Value", func() {

	Context("Calling `NewValue`", func() {
		complexTypeErrFmt := "value is complex type '%s' when it should be a simple type " +
			"or a pointer to a simple type"

		type i interface{}
		var iFace i = &struct{}{}
		var myStr = "my string is hereeeeeeeeeee"
		var myTime = time.Now()

		DescribeTable("`ToSQL`",

			func(d core.Dialect, v interface{}, matchSql, matchErr types.GomegaMatcher) {
				value := expression.NewValue(v)
				sql, err := value.ToSQL(d)

				Expect(err).To(matchErr)
				Expect(sql).To(matchSql)
			},

			// Edge Cases / Odd balls - Failures
			Entry("dialect does not support `Value`", NewMockDialect(), 1, BeNil(), MatchError(EqString("dialect function '%s' not supported", "Value"))),
			Entry("`Value` fails", valueDialectFailure{}, 1, BeNil(), MatchError("value failure")),

			// Basic Types
			// - Successful
			Entry("1", valueDialectSuccess{}, 1, MatchSQLString("[1]"), Not(HaveOccurred())),
			Entry("-1", valueDialectSuccess{}, -1, MatchSQLString("[-1]"), Not(HaveOccurred())),
			Entry("String", valueDialectSuccess{}, "String", MatchSQLString("[String]"), Not(HaveOccurred())),
			Entry("1.55", valueDialectSuccess{}, float32(1.55), MatchSQLString("[1.55]"), Not(HaveOccurred())),
			Entry("1.55", valueDialectSuccess{}, float64(1.55e19), MatchSQLString("[1.55e+19]"), Not(HaveOccurred())),
			// - Failures
			Entry("nil", valueDialectSuccess{}, nil, BeNil(), MatchError("constant is nil, use IsNull instead")),

			// Complex Types
			// - Successful
			Entry("string array", valueDialectSuccess{}, [1]string{"aaa"}, MatchSQLString("[aaa]"), Not(HaveOccurred())),
			Entry("string slice", valueDialectSuccess{}, []string{"aaa"}, MatchSQLString("[aaa]"), Not(HaveOccurred())),
			Entry("int slice", valueDialectSuccess{}, []int{5}, MatchSQLString("[5]"), Not(HaveOccurred())),
			Entry("time.Time slice", valueDialectSuccess{}, []time.Time{{}, myTime}, MatchSQLString(fmt.Sprintf("[%+v %+v]", time.Time{}, myTime)), Not(HaveOccurred())),
			Entry("byte slice", valueDialectSuccess{}, []byte{0x20, 0x02, 0x03, 0x04}, MatchSQLString("[[32 2 3 4]]"), Not(HaveOccurred())),
			Entry("string ptr", valueDialectSuccess{}, &myStr, MatchSQLString("[my string is hereeeeeeeeeee]"), Not(HaveOccurred())),
			// - Failures
			Entry("chan", valueDialectSuccess{}, make(chan string), BeNil(), MatchError(fmt.Sprintf(complexTypeErrFmt, "chan string"))),
			Entry("func", valueDialectSuccess{}, func() {}, BeNil(), MatchError(fmt.Sprintf(complexTypeErrFmt, "func()"))),
			Entry("struct{}", valueDialectSuccess{}, struct{}{}, BeNil(), MatchError(fmt.Sprintf(complexTypeErrFmt, "struct {}"))),
			Entry("interface", valueDialectSuccess{}, iFace, BeNil(), MatchError(fmt.Sprintf(complexTypeErrFmt, "struct {}"))),
		)
	})
})

type valueDialectSuccess struct{}

func (valueDialectSuccess) GetName() string { return "value by dialect" }
func (valueDialectSuccess) Value(i []interface{}) (sql.Data, error) {
	return sql.Format("%+v", i), nil
}

type valueDialectFailure struct{ valueDialectSuccess }

func (valueDialectFailure) Value(_ []interface{}) (sql.Data, error) {
	return nil, errors.New("value failure")
}
