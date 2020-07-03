package expression_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/expression/matchers"
	"github.com/weworksandbox/lingo/pkg/core/expression/operator"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var _ = Describe("Operator", func() {

	Context("Calling `NewOperator`", func() {

		var (
			left   core.Expression
			op     operator.Operand
			values []core.Expression

			newOperator core.ComboExpression
		)

		BeforeEach(func() {
			left = NewMockExpression()
			op = operator.Between
			values = []core.Expression{NewMockExpression(), NewMockExpression()}
		})

		JustBeforeEach(func() {
			newOperator = expression.NewOperator(left, op, values...)
		})

		It("Is not nil", func() {
			Expect(newOperator).ToNot(BeNil())
		})

		Context("`ToSQL`", func() {

			var (
				d core.Dialect

				s   sql.Data
				err error
			)

			BeforeEach(func() {
				d = operatorDialectSuccess{}
				pegomock.When(left.ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("left sql"), nil)
				pegomock.When(values[0].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("values[0] sql"), nil)
				pegomock.When(values[1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("values[1] sql"), nil)
			})

			JustBeforeEach(func() {
				s, err = newOperator.ToSQL(d)
			})

			It("Returns Operator SQL string", func() {
				Expect(s).ToNot(BeNil())
				Expect(s).To(MatchSQLString("operator sql"))
			})

			It("Returns no errors", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			Context("Dialect does not support `Operator`", func() {

				BeforeEach(func() {
					d = NewMockDialect()
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(EqString("dialect function '%s' not supported", "Operator")))
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
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "left")))
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
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "expressions[0]")))
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
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "expressions[1]")))
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

			Context("dialect `Operator` returns an error", func() {

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

type operatorDialectSuccess struct{}

func (operatorDialectSuccess) GetName() string { return "operator success" }
func (operatorDialectSuccess) Operator(sql.Data, operator.Operand, []sql.Data) (sql.Data, error) {
	return sql.String("operator sql"), nil
}

type operatorDialectFailure struct{ operatorDialectSuccess }

func (operatorDialectFailure) Operator(sql.Data, operator.Operand, []sql.Data) (sql.Data, error) {
	return nil, errors.New("operator failure")
}
