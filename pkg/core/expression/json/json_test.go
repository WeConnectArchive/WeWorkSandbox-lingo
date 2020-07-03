package json_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/expression/json"
	. "github.com/weworksandbox/lingo/pkg/core/expression/matchers"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var _ = Describe("JSON", func() {

	Context("#NewJSONOperation", func() {

		var (
			left        core.Expression
			op          Operand
			expressions []core.Expression

			operation core.ComboExpression
		)

		BeforeEach(func() {
			left = expression.NewMockExpression()
			pegomock.When(left.ToSQL(AnyCoreDialect())).ThenReturn(sql.String("left sql"), nil)

			op = Extract
			expressions = []core.Expression{
				expression.NewMockExpression(),
				expression.NewMockExpression(),
			}
			pegomock.When(expressions[0].ToSQL(AnyCoreDialect())).ThenReturn(sql.String("expressions[0]"), nil)
			pegomock.When(expressions[1].ToSQL(AnyCoreDialect())).ThenReturn(sql.String("expressions[1]"), nil)
		})

		JustBeforeEach(func() {
			operation = json.NewJSONOperation(left, op, expressions...)
		})

		It("Creates a valid operation", func() {
			Expect(operation).ToNot(BeNil())
		})

		Context("#ToSQL", func() {

			var (
				d core.Dialect

				sql sql.Data
				err error
			)

			BeforeEach(func() {
				d = jsonDialectSuccess{}
			})

			JustBeforeEach(func() {
				sql, err = operation.ToSQL(d)
			})

			It("Returns a valid SQL", func() {
				Expect(sql).To(matchers.MatchSQLString("json sql"))
			})

			It("Does not error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			Context("Dialect does not support `JSONOperator`", func() {

				BeforeEach(func() {
					d = expression.NewMockDialect()
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns an error that dialect does not support `JSONOperator`", func() {
					Expect(err).To(MatchError(ContainSubstring("dialect function '%s' not supported", "JSONOperation")))
				})
			})

			Context("left is nil", func() {

				BeforeEach(func() {
					left = nil
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns an error that left is nil", func() {
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "left")))
				})
			})

			Context("left returns an error", func() {

				BeforeEach(func() {
					leftErr := errors.New("left error")
					pegomock.When(left.ToSQL(AnyCoreDialect())).ThenReturn(nil, leftErr)
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the left error", func() {
					Expect(err).To(MatchError(ContainSubstring("left error")))
				})
			})

			Context("an expression is nil", func() {

				BeforeEach(func() {
					expressions[1] = nil
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns an error that an expression is nil", func() {
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "expressions[1]")))
				})
			})

			Context("an expression returns an error", func() {

				BeforeEach(func() {
					expErr := errors.New("exp error")
					pegomock.When(expressions[1].ToSQL(AnyCoreDialect())).ThenReturn(nil, expErr)
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the expression error", func() {
					Expect(err).To(MatchError(ContainSubstring("exp error")))
				})
			})

			Context("JSONOperator returns an error", func() {

				BeforeEach(func() {
					d = jsonDialectFailure{}
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns JSONOperator error", func() {
					Expect(err).To(MatchError(ContainSubstring("json failure")))
				})
			})
		})
	})
})

type jsonDialectSuccess struct{}

func (jsonDialectSuccess) GetName() string { return "json success" }
func (jsonDialectSuccess) JSONOperator(sql.Data, Operand, []sql.Data) (sql.Data, error) {
	return sql.String("json sql"), nil
}

type jsonDialectFailure struct{ jsonDialectSuccess }

func (jsonDialectFailure) JSONOperator(sql.Data, Operand, []sql.Data) (sql.Data, error) {
	return nil, errors.New("json failure")
}
