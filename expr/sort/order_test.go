package sort_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr/matchers"
	"github.com/weworksandbox/lingo/expr/sort"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/sql"
)

var _ = Describe("Dialect", func() {

	Context("Calling `NewOrderBy`", func() {

		var (
			left      lingo.Expression
			direction sort.Direction

			orderBy lingo.OrderBy
		)

		BeforeEach(func() {
			left = NewMockExpression()
			direction = sort.OpAscending
		})

		JustBeforeEach(func() {
			orderBy = sort.NewOrderBy(left, direction)
		})

		It("Returns a `lingo.By`", func() {
			Expect(orderBy).ToNot(BeNil())
		})

		Context("Calling `ToSQL`", func() {

			var (
				d lingo.Dialect

				s   sql.Data
				err error
			)

			BeforeEach(func() {
				d = orderDialectSuccess{}
				pegomock.When(left.ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("left sql"), nil)
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

			Context("Dialect does not support `Dialect`", func() {

				BeforeEach(func() {
					d = NewMockDialect()
					pegomock.When(d.GetName()).ThenReturn("mock")
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(EqString("dialect '%s' does not support '%s'", "mock", "sort.Dialect")))
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
					Expect(err).To(MatchError("left of 'order by' cannot be empty"))
				})
			})

			Context("left return an error", func() {

				BeforeEach(func() {
					pegomock.When(left.ToSQL(matchers.AnyLingoDialect())).ThenReturn(nil, errors.New("left error"))
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("left error")))
				})
			})

			Context("`By` returns an error", func() {

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
func (orderDialectSuccess) OrderBy(_ sql.Data, _ sort.Direction) (sql.Data, error) {
	return sql.String("order by sql"), nil
}

type orderDialectFailure struct{ orderDialectSuccess }

func (orderDialectFailure) OrderBy(_ sql.Data, _ sort.Direction) (sql.Data, error) {
	return nil, errors.New("order by failure")
}
