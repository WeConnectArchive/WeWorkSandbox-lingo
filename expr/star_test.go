package expr_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/sql"
)

var _ = Describe("Star", func() {

	Context("Calling `Star`", func() {

		var (
			star lingo.Expression
		)

		JustBeforeEach(func() {
			star = expr.Star()
		})

		It("Returns non nil", func() {
			Expect(star).ShouldNot(BeNil())
		})

		Context("Calling `ToSQL`", func() {

			var (
				d lingo.Dialect

				s   sql.Data
				err error
			)

			BeforeEach(func() {
				d = NewMockDialect()
			})

			JustBeforeEach(func() {
				s, err = star.ToSQL(d)
			})

			It("SQL should match `*`", func() {
				Expect(s).Should(matchers.MatchSQLString("*"))
			})

			It("SQL should have no values", func() {
				Expect(s).Should(matchers.MatchSQLValues(BeEmpty()))
			})

			It("Returns nil error", func() {
				Expect(err).Should(BeNil())
			})
		})
	})
})
