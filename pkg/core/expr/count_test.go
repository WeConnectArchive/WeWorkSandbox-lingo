package expr_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expr"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var _ = Describe("Count", func() {

	Context("Calling `Count`", func() {

		var (
			countOn core.Expression

			count core.Expression
		)

		BeforeEach(func() {
			countOn = NewMockExpression()
		})

		JustBeforeEach(func() {
			count = expr.Count(countOn)
		})

		It("Returns non nil", func() {
			Expect(count).ShouldNot(BeNil())
		})

		Context("Calling `ToSQL`", func() {

			var (
				d core.Dialect

				s   sql.Data
				err error
			)

			BeforeEach(func() {
				d = NewMockDialect()
				pegomock.When(countOn.ToSQL(d)).ThenReturn(sql.New("countOn s ?", []interface{}{10}), nil)
			})

			JustBeforeEach(func() {
				s, err = count.ToSQL(d)
			})

			It("SQL should match `COUNT()`", func() {
				Expect(s).Should(matchers.MatchSQLString(MatchRegexp(`COUNT\(.+\)`)))
			})

			It("SQL should have no values", func() {
				Expect(s).Should(matchers.MatchSQLValues(matchers.AllInSlice(10)))
			})

			It("Returns nil error", func() {
				Expect(err).Should(BeNil())
			})

			Context("Expression is nil", func() {

				BeforeEach(func() {
					countOn = nil
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError("countOn value cannot be empty"))
				})
			})

			Context("expr returns an error", func() {

				BeforeEach(func() {
					pegomock.When(countOn.ToSQL(d)).ThenReturn(nil, errors.New("countOn error"))
				})

				It("Returns a nil SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns the expr error", func() {
					Expect(err).To(MatchError("countOn error"))
				})
			})
		})
	})
})
