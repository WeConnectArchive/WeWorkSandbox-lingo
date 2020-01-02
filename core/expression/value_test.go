package expression_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/expression"
	"errors"
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

			// Odd balls
			Entry("dialect does not support `Value`", NewMockDialect(), 1, BeNil(), MatchError(EqString("dialect function '%s' not supported", "Value"))),
			Entry("`Value` fails", valueDialectFailure{}, 1, BeNil(), MatchError("value failure")),

			// Basic Types
			Entry("nil", valueDialectSuccess{}, nil, BeNil(), MatchError("constant is nil, use IsNull instead")),
			Entry("1", valueDialectSuccess{}, 1, MatchSQLString("value sql"), Not(HaveOccurred())),
			Entry("-1", valueDialectSuccess{}, -1, MatchSQLString("value sql"), Not(HaveOccurred())),
			Entry("String", valueDialectSuccess{}, "String", MatchSQLString("value sql"), Not(HaveOccurred())),
			Entry("1.55", valueDialectSuccess{}, 1.55, MatchSQLString("value sql"), Not(HaveOccurred())),
			Entry("1.55", valueDialectSuccess{}, 1.55, MatchSQLString("value sql"), Not(HaveOccurred())),

			// Complex Types
			Entry("string array", valueDialectSuccess{}, []string{"aaa"}, BeNil(), MatchError(fmt.Sprintf("value is complex type '%s' when it should be a simple type", "<[]string Value>"))),
			Entry("int array", valueDialectSuccess{}, []int{5}, BeNil(), MatchError(fmt.Sprintf("value is complex type '%s' when it should be a simple type", "<[]int Value>"))),
			// Yes, I know the channel is a memory leak, but its a unit test, I do not care.
			Entry("chan", valueDialectSuccess{}, make(chan string, 0), BeNil(), MatchError(fmt.Sprintf("value is complex type '%s' when it should be a simple type", "<chan string Value>"))),
			Entry("func", valueDialectSuccess{}, func() {}, BeNil(), MatchError(fmt.Sprintf("value is complex type '%s' when it should be a simple type", "<func() Value>"))),
			Entry("struct{}", valueDialectSuccess{}, struct{}{}, BeNil(), MatchError(fmt.Sprintf("value is complex type '%s' when it should be a simple type", "<struct {} Value>"))),
			Entry("interface", valueDialectSuccess{}, iFace, BeNil(), MatchError(fmt.Sprintf("value is complex type '%s' when it should be a simple type", "<*struct {} Value>"))),
		)
	})

	Context("Calling `NewValues`", func() {

		var (
			v interface{}

			values []core.Expression
		)

		BeforeEach(func() {
			v = 1234
		})

		JustBeforeEach(func() {
			values = expression.NewValues(v)
		})

		It("Returns an array of one `Value`", func() {
			Expect(values).To(HaveLen(1))
			Expect(values).To(ContainElement(expression.NewValue(v)))
		})

		Context("With a slice of bytes", func() {

			BeforeEach(func() {
				v = []byte{0x01, 0x02, 0x03}
			})

			It("Returns an array with a single `Value`", func() {
				Expect(values).To(HaveLen(1))
				Expect(values).To(ContainElement(expression.NewValue(v)))
			})
		})

		Context("With a slice of strings", func() {

			BeforeEach(func() {
				v = []string{"a", "b"}
			})

			It("Returns an array with a two `Value`s", func() {
				Expect(values).To(HaveLen(2))
				Expect(values).To(ContainElement(expression.NewValue("a")))
				Expect(values).To(ContainElement(expression.NewValue("b")))
			})
		})

		Context("With an array of bytes", func() {

			BeforeEach(func() {
				v = [2]byte{0x01, 0x02}
			})

			It("Returns an array with one `Value`", func() {
				Expect(values).To(HaveLen(1))
				Expect(values).To(ContainElement(expression.NewValue([2]byte{0x01, 0x02})))
			})
		})
	})
})

type valueDialectSuccess struct{}

func (valueDialectSuccess) GetName() string { return "value by dialect" }
func (valueDialectSuccess) Value(value interface{}) (core.SQL, error) {
	return core.NewSQLf("value sql"), nil
}

type valueDialectFailure struct{ valueDialectSuccess }

func (valueDialectFailure) Value(value interface{}) (core.SQL, error) {
	return nil, errors.New("value failure")
}
