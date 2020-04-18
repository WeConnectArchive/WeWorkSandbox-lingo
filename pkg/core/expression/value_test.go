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
)

var _ = Describe("Value", func() {

	Context("Calling `NewValue`", func() {
		type i interface{}
		var iFace i = &struct{}{}
		DescribeTable("`GetSQL`",

			func(d core.Dialect, v interface{}, matchSql, matchErr types.GomegaMatcher) {
				value := expression.NewValue(v)
				sql, err := value.GetSQL(d)

				Expect(err).To(matchErr)
				Expect(sql).To(matchSql)
			},

			// Edge Cases / Odd balls - Failures
			Entry("dialect does not support `Value`", NewMockDialect(), 1, BeNil(), MatchError(EqString("dialect function '%s' not supported", "Value"))),
			Entry("`Value` fails", valueDialectFailure{}, 1, BeNil(), MatchError("value failure")),

			// Basic Types
			// - Successful
			Entry("1", valueDialectSuccess{}, 1, MatchSQLString("value sql"), Not(HaveOccurred())),
			Entry("-1", valueDialectSuccess{}, -1, MatchSQLString("value sql"), Not(HaveOccurred())),
			Entry("String", valueDialectSuccess{}, "String", MatchSQLString("value sql"), Not(HaveOccurred())),
			Entry("1.55", valueDialectSuccess{}, 1.55, MatchSQLString("value sql"), Not(HaveOccurred())),
			Entry("1.55", valueDialectSuccess{}, 1.55, MatchSQLString("value sql"), Not(HaveOccurred())),
			// - Failures
			Entry("nil", valueDialectSuccess{}, nil, BeNil(), MatchError("constant is nil, use IsNull instead")),

			// Complex Types
			// - Successful
			Entry("string array", valueDialectSuccess{}, [1]string{"aaa"}, MatchSQLString("value sql"), Not(HaveOccurred())),
			Entry("string slice", valueDialectSuccess{}, []string{"aaa"}, MatchSQLString("value sql"), Not(HaveOccurred())),
			Entry("int slice", valueDialectSuccess{}, []int{5}, MatchSQLString("value sql"), Not(HaveOccurred())),
			Entry("time.Time slice", valueDialectSuccess{}, []time.Time{{}, time.Now()}, MatchSQLString("value sql"), Not(HaveOccurred())),
			Entry("byte slice", valueDialectSuccess{}, []byte{0x01, 0x02, 0x03, 0x04}, MatchSQLString("value sql"), Not(HaveOccurred())),
			// - Failures
			Entry("chan", valueDialectSuccess{}, make(chan string), BeNil(), MatchError(fmt.Sprintf("value is complex type '%s' when it should be a simple type", "chan string"))),
			Entry("func", valueDialectSuccess{}, func() {}, BeNil(), MatchError(fmt.Sprintf("value is complex type '%s' when it should be a simple type", "func()"))),
			Entry("struct{}", valueDialectSuccess{}, struct{}{}, BeNil(), MatchError(fmt.Sprintf("value is complex type '%s' when it should be a simple type", "struct {}"))),
			Entry("interface", valueDialectSuccess{}, iFace, BeNil(), MatchError(fmt.Sprintf("value is complex type '%s' when it should be a simple type", "*struct {}"))),
		)
	})
})

type valueDialectSuccess struct{}

func (valueDialectSuccess) GetName() string { return "value by dialect" }
func (valueDialectSuccess) Value(_ []interface{}) (core.SQL, error) {
	return core.NewSQLf("value sql"), nil
}

type valueDialectFailure struct{ valueDialectSuccess }

func (valueDialectFailure) Value(_ []interface{}) (core.SQL, error) {
	return nil, errors.New("value failure")
}
