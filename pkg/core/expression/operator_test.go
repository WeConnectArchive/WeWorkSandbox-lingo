package expression_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/expression/matchers"
	"github.com/weworksandbox/lingo/pkg/core/operator"
	"errors"
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

		Context("`GetSQL`", func() {

			var (
				d core.Dialect

				sql core.SQL
				err error
			)

			BeforeEach(func() {
				d = operatorDialectSuccess{}
				pegomock.When(left.GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("left sql"), nil)
				pegomock.When(values[0].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("values[0] sql"), nil)
				pegomock.When(values[1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("values[1] sql"), nil)
			})

			JustBeforeEach(func() {
				sql, err = newOperator.GetSQL(d)
			})

			It("Returns Operator SQL string", func() {
				Expect(sql).ToNot(BeNil())
				Expect(sql).To(MatchSQLString("operator sql"))
			})

			It("Returns no errors", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			Context("Dialect does not support `Operator`", func() {

				BeforeEach(func() {
					d = NewMockDialect()
				})

				It("Returns no SQL", func() {
					Expect(sql).To(BeNil())
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
					Expect(sql).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "left")))
				})
			})

			Context("left returns an error", func() {

				BeforeEach(func() {
					pegomock.When(left.GetSQL(d)).ThenReturn(nil, errors.New("left error"))
				})

				It("Returns no SQL", func() {
					Expect(sql).To(BeNil())
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
					Expect(sql).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "expressions[0]")))
				})
			})

			Context("first value returns an error", func() {

				BeforeEach(func() {
					pegomock.When(values[0].GetSQL(d)).ThenReturn(nil, errors.New("values[0] error"))
				})

				It("Returns no SQL", func() {
					Expect(sql).To(BeNil())
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
					Expect(sql).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "expressions[1]")))
				})
			})

			Context("second value returns an error", func() {

				BeforeEach(func() {
					pegomock.When(values[1].GetSQL(d)).ThenReturn(nil, errors.New("values[1] error"))
				})

				It("Returns no SQL", func() {
					Expect(sql).To(BeNil())
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
					Expect(sql).To(BeNil())
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
func (operatorDialectSuccess) Operator(core.SQL, operator.Operand, []core.SQL) (core.SQL, error) {
	return core.NewSQLf("operator sql"), nil
}

type operatorDialectFailure struct{ operatorDialectSuccess }

func (operatorDialectFailure) Operator(core.SQL, operator.Operand, []core.SQL) (core.SQL, error) {
	return nil, errors.New("operator failure")
}
