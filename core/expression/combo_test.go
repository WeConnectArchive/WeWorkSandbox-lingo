package expression_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/expression"
	. "github.com/weworksandbox/lingo/core/expression/matchers"
)

var _ = Describe("ComboExpression", func() {

	Context("New `ComboExpression`", func() {

		var (
			exp core.Expression

			combo *expression.ComboExpression
		)

		BeforeEach(func() {
			exp = NewMockExpression()
			pegomock.When(exp.GetSQL(AnyCoreDialect())).ThenReturn(core.NewSQLf("expression sql"), nil)
		})

		JustBeforeEach(func() {
			combo = expression.NewComboExpression(exp)
		})

		It("Returns a `ComboExpression`", func() {
			Expect(combo).ToNot(BeNil())
		})

		Context("`And`", func() {

			var (
				withExp core.Expression

				combined core.ComboExpression
			)

			BeforeEach(func() {
				withExp = NewMockExpression()
				pegomock.When(withExp.GetSQL(AnyCoreDialect())).ThenReturn(core.NewSQLf("with expression sql"), nil)
			})

			JustBeforeEach(func() {
				combined = combo.And(withExp)
			})

			Context("`GetSQL`", func() {

				var (
					d core.Dialect

					sql core.SQL
					err error
				)

				BeforeEach(func() {
					d = operatorDialectSuccess{}
				})

				JustBeforeEach(func() {
					sql, err = combined.GetSQL(d)
				})

				It("returns a valid `SQL`", func() {
					Expect(sql).To(matchers.MatchSQLString("operator sql"))
				})

				It("returns a no errors", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})
		})

		Context("`Or`", func() {

			var (
				withExp core.Expression

				combined core.ComboExpression
			)

			BeforeEach(func() {
				withExp = NewMockExpression()
				pegomock.When(withExp.GetSQL(AnyCoreDialect())).ThenReturn(core.NewSQLf("with expression sql"), nil)
			})

			JustBeforeEach(func() {
				combined = combo.Or(withExp)
			})

			Context("`GetSQL`", func() {

				var (
					d core.Dialect

					sql core.SQL
					err error
				)

				BeforeEach(func() {
					d = operatorDialectSuccess{}
				})

				JustBeforeEach(func() {
					sql, err = combined.GetSQL(d)
				})

				It("returns a valid `SQL`", func() {
					Expect(sql).To(matchers.MatchSQLString("operator sql"))
				})

				It("returns a no errors", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})
		})
	})
})
