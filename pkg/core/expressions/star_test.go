package expressions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expressions"
)

var _ = Describe("Star", func() {

	Context("Calling `Star`", func() {

		var (
			star core.Expression
		)

		JustBeforeEach(func() {
			star = expressions.Star()
		})

		It("Returns non nil", func() {
			Expect(star).ShouldNot(BeNil())
		})

		Context("Calling `GetSQL`", func() {

			var (
				d core.Dialect

				sql core.SQL
				err error
			)

			BeforeEach(func() {
				d = NewMockDialect()
			})

			JustBeforeEach(func() {
				sql, err = star.GetSQL(d)
			})

			It("SQL should match `*`", func() {
				Expect(sql).Should(matchers.MatchSQLString("*"))
			})

			It("SQL should have no values", func() {
				Expect(sql).Should(matchers.MatchSQLValues(BeEmpty()))
			})

			It("Returns nil error", func() {
				Expect(err).Should(BeNil())
			})
		})
	})
})
