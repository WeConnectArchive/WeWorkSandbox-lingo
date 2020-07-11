package expr_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/sql"
)

var _ = Describe("list.go", func() {

	Context("#ToList", func() {

		var (
			values []lingo.Expression

			list lingo.Expression
		)

		BeforeEach(func() {
			values = []lingo.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
		})

		JustBeforeEach(func() {
			list = expr.ToList(values)
		})

		It("Returns non nil", func() {
			Expect(list).ShouldNot(BeNil())
		})

		Context("Calling `ToSQL`", func() {

			var (
				d lingo.Dialect

				s   sql.Data
				err error
			)

			BeforeEach(func() {
				d = NewMockDialect()
				pegomock.When(values[0].ToSQL(d)).ThenReturn(sql.New("values[0] ?", []interface{}{10}), nil)
				pegomock.When(values[1].ToSQL(d)).ThenReturn(sql.New("values[1] ?", []interface{}{55}), nil)
			})

			JustBeforeEach(func() {
				s, err = list.ToSQL(d)
			})

			It("SQL should create a comma+space separated list`", func() {
				Expect(s).Should(matchers.MatchSQLString(`values[0] ?, values[1] ?`))
			})

			It("SQL should have 2 values", func() {
				Expect(s).Should(matchers.MatchSQLValues(matchers.AllInSlice(10, 55)))
			})

			It("Returns nil error", func() {
				Expect(err).Should(BeNil())
			})

			Context("Expression is nil", func() {

				BeforeEach(func() {
					values = nil
				})

				It("Returns empty SQL", func() {
					Expect(s).Should(matchers.MatchSQLString(""))
				})

				It("Returns nil error", func() {
					Expect(err).Should(BeNil())
				})
			})

			Context("expr returns an error", func() {

				BeforeEach(func() {
					pegomock.When(values[1].ToSQL(d)).ThenReturn(nil, errors.New("values[1] error"))
				})

				It("Returns a nil SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns the expr error", func() {
					Expect(err).To(MatchError("values[1] error"))
				})
			})
		})
	})
})
