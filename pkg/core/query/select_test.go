package query_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/dialect"
	"github.com/weworksandbox/lingo/pkg/core/expression/join"
	"github.com/weworksandbox/lingo/pkg/core/expression/sort"
	"github.com/weworksandbox/lingo/pkg/core/query"
	"github.com/weworksandbox/lingo/pkg/core/query/matchers"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var _ = Describe("select", func() {

	Context("Select", func() {

		var (
			paths     []core.Expression
			from      core.Table
			where     []core.Expression
			orderBy   []core.Expression
			direction []sort.Direction
			joins     [][]core.Expression
			joinType  []join.Type
			modifier  query.Modifier

			q *query.SelectQuery
		)

		BeforeEach(func() {
			paths = []core.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(paths[0].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("path[0].sql"), nil)
			pegomock.When(paths[1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("path[1].sql"), nil)

			from = NewMockTable()
			pegomock.When(from.ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("from.sql"), nil)

			where = []core.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(where[0].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("where[0].sql"), nil)
			pegomock.When(where[1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("where[1].sql"), nil)

			orderBy = []core.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			direction = []sort.Direction{
				sort.Ascending,
				sort.Descending,
			}
			pegomock.When(orderBy[0].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("orderBy[0].sql"), nil)
			pegomock.When(orderBy[1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("orderBy[1].sql"), nil)

			joins = [][]core.Expression{
				{
					NewMockExpression(),
					NewMockExpression(),
				},
				{
					NewMockExpression(),
					NewMockExpression(),
				},
			}
			joinType = []join.Type{
				join.Left,
				join.Right,
			}
			pegomock.When(joins[0][0].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("joins[0][0].sql"), nil)
			pegomock.When(joins[0][1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("joins[0][1].sql"), nil)
			pegomock.When(joins[1][0].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("joins[1][0].sql"), nil)
			pegomock.When(joins[1][1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("joins[1][1].sql"), nil)

			modifier = NewMockModifier()
			pegomock.When(modifier.IsZero()).ThenReturn(false)
			pegomock.When(modifier.Limit()).ThenReturn(uint64(10), true)
			pegomock.When(modifier.Offset()).ThenReturn(uint64(3), true)
		})

		JustBeforeEach(func() {
			q = query.Select(paths...).From(from).Where(where...).Restrict(modifier)
			for i, order := range orderBy {
				q = q.OrderBy(order, direction[i])
			}
			for i, join := range joins {
				q = q.Join(join[0], joinType[i], join[1])
			}
		})

		Context("#ToSQL", func() {

			var (
				d core.Dialect

				sql sql.Data
				err error
			)

			BeforeEach(func() {
				d = dialect.Default{}
			})

			JustBeforeEach(func() {
				sql, err = q.ToSQL(d)
			})

			It("Returns a valid SQL string", func() {
				Expect(sql).To(MatchSQLString("SELECT path[0].sql, path[1].sql FROM from.sql LEFT JOIN joins[0][0].sql ON joins[0][1].sql RIGHT JOIN joins[1][0].sql ON joins[1][1].sql WHERE (where[0].sql AND where[1].sql) ORDER BY orderBy[0].sql ASC, orderBy[1].sql DESC LIMIT ? OFFSET ?"))
			})

			It("Returns no error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			Context("Without any columns", func() {

				BeforeEach(func() {
					paths = nil
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns an columns cannot be empty error", func() {
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be empty", "columns")))
				})
			})

			Context("Error build path SQL", func() {

				BeforeEach(func() {
					pegomock.When(paths[len(paths)-1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("path error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns an columns cannot be empty error", func() {
					Expect(err).To(MatchError(ContainSubstring("path error")))
				})
			})

			Context("From is nil", func() {

				BeforeEach(func() {
					from = nil
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a from cannot be nil error", func() {
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "from")))
				})
			})

			Context("Error building from SQL", func() {

				BeforeEach(func() {
					pegomock.When(from.ToSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("from error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a from error", func() {
					Expect(err).To(MatchError(ContainSubstring("from error")))
				})
			})

			Context("No joins", func() {

				BeforeEach(func() {
					joins = nil
				})

				It("Returns a valid SQL string", func() {
					Expect(sql).To(MatchSQLString("SELECT path[0].sql, path[1].sql FROM from.sql WHERE (where[0].sql AND where[1].sql) ORDER BY orderBy[0].sql ASC, orderBy[1].sql DESC LIMIT ? OFFSET ?"))
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("Error on left side of joins", func() {

				BeforeEach(func() {
					pegomock.When(joins[len(joins)-1][0].ToSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("left joins error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a left joins error", func() {
					Expect(err).To(MatchError(ContainSubstring("left joins error")))
				})
			})

			Context("Error on on of joins", func() {

				BeforeEach(func() {
					pegomock.When(joins[len(joins)-1][1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("on joins error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns an on joins error", func() {
					Expect(err).To(MatchError(ContainSubstring("on joins error")))
				})
			})

			Context("Where is nil", func() {

				BeforeEach(func() {
					where = nil
				})

				It("Returns a valid SQL string", func() {
					Expect(sql).To(MatchSQLString("SELECT path[0].sql, path[1].sql FROM from.sql LEFT JOIN joins[0][0].sql ON joins[0][1].sql RIGHT JOIN joins[1][0].sql ON joins[1][1].sql ORDER BY orderBy[0].sql ASC, orderBy[1].sql DESC LIMIT ? OFFSET ?"))
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("Where has error", func() {

				BeforeEach(func() {
					pegomock.When(where[len(where)-1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("where error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a where error", func() {
					Expect(err).To(MatchError(ContainSubstring("where error")))
				})
			})

			Context("Order By is nil", func() {

				BeforeEach(func() {
					orderBy = nil
				})

				It("Returns a valid SQL string", func() {
					Expect(sql).To(MatchSQLString("SELECT path[0].sql, path[1].sql FROM from.sql LEFT JOIN joins[0][0].sql ON joins[0][1].sql RIGHT JOIN joins[1][0].sql ON joins[1][1].sql WHERE (where[0].sql AND where[1].sql) LIMIT ? OFFSET ?"))
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("Order By has error", func() {

				BeforeEach(func() {
					pegomock.When(orderBy[len(where)-1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("order by error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a order by error", func() {
					Expect(err).To(MatchError(ContainSubstring("order by error")))
				})
			})

			Context("modifier IsZero = true", func() {

				BeforeEach(func() {
					modifier = NewMockModifier()
					pegomock.When(modifier.IsZero()).ThenReturn(true)
				})

				It("Returns a valid SQL", func() {
					Expect(sql).To(MatchSQLString("SELECT path[0].sql, path[1].sql FROM from.sql LEFT JOIN joins[0][0].sql ON joins[0][1].sql RIGHT JOIN joins[1][0].sql ON joins[1][1].sql WHERE (where[0].sql AND where[1].sql) ORDER BY orderBy[0].sql ASC, orderBy[1].sql DESC"))
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})
		})
	})

	Context("SelectFrom", func() {

		var (
			from core.Table
			cols []core.Column

			q *query.SelectQuery
		)

		BeforeEach(func() {
			cols = []core.Column{
				NewMockColumn(),
				NewMockColumn(),
			}
			pegomock.When(cols[0].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("cols[0].sql"), nil)
			pegomock.When(cols[1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("cols[1].sql"), nil)

			from = NewMockTable()
			pegomock.When(from.ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("from.sql"), nil)
			pegomock.When(from.GetColumns()).ThenReturn(cols)
		})

		JustBeforeEach(func() {
			q = query.SelectFrom(from)
		})

		Context("#ToSQL", func() {

			var (
				d core.Dialect

				sql sql.Data
				err error
			)

			BeforeEach(func() {
				d = dialect.Default{}
			})

			JustBeforeEach(func() {
				sql, err = q.ToSQL(d)
			})

			It("Returns a valid SQL string", func() {
				Expect(sql).To(MatchSQLString("SELECT cols[0].sql, cols[1].sql FROM from.sql"))
			})

			It("Returns no error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
