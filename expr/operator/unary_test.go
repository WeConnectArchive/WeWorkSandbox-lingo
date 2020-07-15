package operator_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/operator/matchers"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/sql"
)

var _ = Describe("unary.go", func() {

	Context("#NewUnary", func() {

		var (
			left lingo.Expression
			op   operator.Operator

			newOperator lingo.Expression
		)

		BeforeEach(func() {
			left = NewMockExpression()
			op = operator.OpBetween
		})

		JustBeforeEach(func() {
			newOperator = operator.NewUnary(left, op)
		})

		It("Is not nil", func() {
			Expect(newOperator).ToNot(BeNil())
		})

		Context("#ToSQL", func() {

			var (
				d lingo.Dialect

				s   sql.Data
				err error
			)

			BeforeEach(func() {
				d = operatorDialectSuccess{}
				pegomock.When(left.ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("left sql"), nil)
			})

			JustBeforeEach(func() {
				s, err = newOperator.ToSQL(d)
			})

			It("Returns Dialect SQL string", func() {
				Expect(s).ToNot(BeNil())
				Expect(s).To(MatchSQLString("operator sql"))
			})

			It("Returns no errors", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			Context("Dialect does not support `Dialect`", func() {

				BeforeEach(func() {
					d = NewMockDialect()
					pegomock.When(d.GetName()).ThenReturn("mock")
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(EqString("dialect '%s' does not support '%s'", "mock", "operator.Dialect")))
				})
			})

			Context("left is nil", func() {

				BeforeEach(func() {
					left = nil
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError("left of operator.Unary cannot be empty"))
				})
			})

			Context("left returns an error", func() {

				BeforeEach(func() {
					pegomock.When(left.ToSQL(d)).ThenReturn(nil, errors.New("left error"))
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError("left error"))
				})
			})

			Context("dialect `Dialect` returns an error", func() {

				BeforeEach(func() {
					d = operatorDialectFailure{}
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("operator failure")))
				})
			})
		})
	})
})
