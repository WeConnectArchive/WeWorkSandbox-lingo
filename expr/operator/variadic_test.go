package operator_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/matchers"
	"github.com/weworksandbox/lingo/expr/operator"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/sql"
)

var _ = Describe("variadic.go", func() {

	Context("#NewVariadic", func() {

		var (
			left   lingo.Expression
			op     operator.Operator
			values []lingo.Expression

			newOperator lingo.ComboExpression
		)

		BeforeEach(func() {
			left = NewMockExpression()
			op = operator.Between
			values = []lingo.Expression{NewMockExpression(), NewMockExpression()}
		})

		JustBeforeEach(func() {
			newOperator = operator.NewVariadic(left, op, values)
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
				pegomock.When(values[0].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("values[0] sql"), nil)
				pegomock.When(values[1].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("values[1] sql"), nil)
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
					Expect(err).To(MatchError("left of operator.Variadic cannot be empty"))
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

			Context("first value is nil", func() {

				BeforeEach(func() {
					values[0] = nil
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError("expressions[0] of operator.Variadic cannot be empty"))
				})
			})

			Context("first value returns an error", func() {

				BeforeEach(func() {
					pegomock.When(values[0].ToSQL(d)).ThenReturn(nil, errors.New("values[0] error"))
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("values[0] error")))
				})
			})

			Context("second value is nil", func() {

				BeforeEach(func() {
					values[1] = nil
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError("expressions[1] of operator.Variadic cannot be empty"))
				})
			})

			Context("second value returns an error", func() {

				BeforeEach(func() {
					pegomock.When(values[1].ToSQL(d)).ThenReturn(nil, errors.New("values[1] error"))
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("values[1] error")))
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
