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
	"github.com/weworksandbox/lingo/pkg/core/expression/sort"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var _ = Describe("Order", func() {

	Context("Calling `NewOrderBy`", func() {

		var (
			left      core.Expression
			direction sort.Direction

			orderBy core.OrderBy
		)

		BeforeEach(func() {
			left = NewMockExpression()
			direction = sort.Ascending
		})

		JustBeforeEach(func() {
			orderBy = expression.NewOrderBy(left, direction)
		})

		It("Returns a `core.OrderBy`", func() {
			Expect(orderBy).ToNot(BeNil())
		})

		Context("Calling `ToSQL`", func() {

			var (
				d core.Dialect

				s   sql.Data
				err error
			)

			BeforeEach(func() {
				d = orderDialectSuccess{}
				pegomock.When(left.ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("left sql"), nil)
			})

			JustBeforeEach(func() {
				s, err = orderBy.ToSQL(d)
			})

			It("Returns Set SQL string", func() {
				Expect(s).ToNot(BeNil())
				Expect(s).To(MatchSQLString("order by sql"))
			})

			It("Returns no errors", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			Context("Dialect does not support `Order`", func() {

				BeforeEach(func() {
					d = NewMockDialect()
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(EqString("dialect function '%s' not supported", "Order")))
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

			Context("left return an error", func() {

				BeforeEach(func() {
					pegomock.When(left.ToSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("left error"))
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("left error")))
				})
			})

			Context("direction is `Unknown`", func() {

				BeforeEach(func() {
					direction = sort.Unknown
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "direction")))
				})
			})

			Context("`OrderBy` returns an error", func() {

				BeforeEach(func() {
					d = orderDialectFailure{}
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("order by failure")))
				})
			})
		})
	})
})

type orderDialectSuccess struct{}

func (orderDialectSuccess) GetName() string { return "order by dialect" }
func (orderDialectSuccess) OrderBy(left sql.Data, direction sort.Direction) (sql.Data, error) {
	return sql.String("order by sql"), nil
}

type orderDialectFailure struct{ orderDialectSuccess }

func (orderDialectFailure) OrderBy(left sql.Data, direction sort.Direction) (sql.Data, error) {
	return nil, errors.New("order by failure")
}
