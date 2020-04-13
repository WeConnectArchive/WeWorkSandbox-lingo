package expressions_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expressions"
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
			count = expressions.Count(countOn)
		})

		It("Returns non nil", func() {
			Expect(count).ShouldNot(BeNil())
		})

		Context("Calling `GetSQL`", func() {

			var (
				d core.Dialect

				sql core.SQL
				err error
			)

			BeforeEach(func() {
				d = NewMockDialect()
				pegomock.When(countOn.GetSQL(d)).ThenReturn(core.NewSQL("countOn sql ?", []interface{}{10}), nil)
			})

			JustBeforeEach(func() {
				sql, err = count.GetSQL(d)
			})

			It("SQL should match `COUNT()`", func() {
				Expect(sql).Should(matchers.MatchSQLString(MatchRegexp(`COUNT\(.+\)`)))
			})

			It("SQL should have no values", func() {
				Expect(sql).Should(matchers.MatchSQLValues(ContainElement(10)))
			})

			It("Returns nil error", func() {
				Expect(err).Should(BeNil())
			})

			Context("Expression is nil", func() {

				BeforeEach(func() {
					countOn = nil
				})

				It("Returns no SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "countOn")))
				})
			})

			Context("expression returns an error", func() {

				BeforeEach(func() {
					pegomock.When(countOn.GetSQL(d)).ThenReturn(nil, errors.New("countOn error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the expression error", func() {
					Expect(err).To(MatchError("countOn error"))
				})
			})
		})
	})
})
