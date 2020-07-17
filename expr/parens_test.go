package expr_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/sql"
)

var _ = Describe("parens.go", func() {

	Context("#NewParens", func() {

		var (
			exp lingo.Expression

			list lingo.Expression
		)

		BeforeEach(func() {
			exp = NewMockExpression()
		})

		JustBeforeEach(func() {
			list = expr.NewParens(exp)
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
				pegomock.When(exp.ToSQL(d)).ThenReturn(sql.New("value?", []interface{}{10}), nil)
			})

			JustBeforeEach(func() {
				s, err = list.ToSQL(d)
			})

			It("SQL should value with parens and no spaces`", func() {
				Expect(s).Should(matchers.MatchSQLString(`(value?)`))
			})

			It("SQL should have 1 value", func() {
				Expect(s).Should(matchers.MatchSQLValues(matchers.AllInSlice(10)))
			})

			It("Returns nil error", func() {
				Expect(err).Should(BeNil())
			})

			Context("Expr is nil", func() {

				BeforeEach(func() {
					exp = nil
				})

				It("Returns a nil SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns the expr error", func() {
					Expect(err).To(MatchError("paren Expr cannot be empty"))
				})
			})
		})
	})
})
