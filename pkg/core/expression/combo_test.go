package expression_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	. "github.com/weworksandbox/lingo/pkg/core/expression/matchers"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var _ = Describe("ComboExpression", func() {

	Context("New `ComboExpression`", func() {

		var (
			exp core.Expression

			combo *expression.ComboExpression
		)

		BeforeEach(func() {
			exp = NewMockExpression()
			pegomock.When(exp.ToSQL(AnyCoreDialect())).ThenReturn(sql.String("expression sql"), nil)
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
				pegomock.When(withExp.ToSQL(AnyCoreDialect())).ThenReturn(sql.String("with expression sql"), nil)
			})

			JustBeforeEach(func() {
				combined = combo.And(withExp)
			})

			Context("`ToSQL`", func() {

				var (
					d core.Dialect

					sql sql.Data
					err error
				)

				BeforeEach(func() {
					d = operatorDialectSuccess{}
				})

				JustBeforeEach(func() {
					sql, err = combined.ToSQL(d)
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
				pegomock.When(withExp.ToSQL(AnyCoreDialect())).ThenReturn(sql.String("with expression sql"), nil)
			})

			JustBeforeEach(func() {
				combined = combo.Or(withExp)
			})

			Context("`ToSQL`", func() {

				var (
					d core.Dialect

					sql sql.Data
					err error
				)

				BeforeEach(func() {
					d = operatorDialectSuccess{}
				})

				JustBeforeEach(func() {
					sql, err = combined.ToSQL(d)
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
