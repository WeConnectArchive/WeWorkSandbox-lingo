package query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/dialect"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/query"
	"github.com/weworksandbox/lingo/pkg/core/query/matchers"
	"github.com/weworksandbox/lingo/pkg/core/sort"
	"golang.org/x/xerrors"
)

var _ = Describe("select", func() {

	Context("Select", func() {

		var (
			paths     []core.Expression
			from      core.Table
			where     []core.Expression
			orderBy   []core.Expression
			direction []sort.Direction
			join      [][]core.Expression
			joinType  []expression.JoinType

			q *query.SelectQuery
		)

		BeforeEach(func() {
			paths = []core.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(paths[0].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("path[0].sql"), nil)
			pegomock.When(paths[1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("path[1].sql"), nil)

			from = NewMockTable()
			pegomock.When(from.GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("from.sql"), nil)

			where = []core.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(where[0].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("where[0].sql"), nil)
			pegomock.When(where[1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("where[1].sql"), nil)

			orderBy = []core.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			direction = []sort.Direction{
				sort.Ascending,
				sort.Descending,
			}
			pegomock.When(orderBy[0].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("orderBy[0].sql"), nil)
			pegomock.When(orderBy[1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("orderBy[1].sql"), nil)

			join = [][]core.Expression{
				{
					NewMockExpression(),
					NewMockExpression(),
				},
				{
					NewMockExpression(),
					NewMockExpression(),
				},
			}
			joinType = []expression.JoinType{
				expression.LeftJoin,
				expression.RightJoin,
			}
			pegomock.When(join[0][0].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("join[0][0].sql"), nil)
			pegomock.When(join[0][1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("join[0][1].sql"), nil)
			pegomock.When(join[1][0].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("join[1][0].sql"), nil)
			pegomock.When(join[1][1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("join[1][1].sql"), nil)
		})

		JustBeforeEach(func() {
			q = query.Select(paths...).From(from).Where(where...)
			for i, order := range orderBy {
				q = q.OrderBy(order, direction[i])
			}
			for i, join := range join {
				q = q.Join(join[0], joinType[i], join[1])
			}
		})

		Context("#GetSQL", func() {

			var (
				d core.Dialect

				sql core.SQL
				err error
			)

			BeforeEach(func() {
				d = dialect.Default{}
			})

			JustBeforeEach(func() {
				sql, err = q.GetSQL(d)
			})

			It("Returns a valid SQL string", func() {
				Expect(sql).To(MatchSQLString("SELECT path[0].sql, path[1].sql FROM from.sql LEFT JOIN join[0][0].sql ON join[0][1].sql RIGHT JOIN join[1][0].sql ON join[1][1].sql WHERE (where[0].sql AND where[1].sql) ORDER BY orderBy[0].sql ASC, orderBy[1].sql DESC"))
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
					pegomock.When(paths[len(paths)-1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, xerrors.New("path error"))
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
					pegomock.When(from.GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, xerrors.New("from error"))
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
					join = nil
				})

				It("Returns a valid SQL string", func() {
					Expect(sql).To(MatchSQLString("SELECT path[0].sql, path[1].sql FROM from.sql WHERE (where[0].sql AND where[1].sql) ORDER BY orderBy[0].sql ASC, orderBy[1].sql DESC"))
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("Error on left side of join", func() {

				BeforeEach(func() {
					pegomock.When(join[len(join)-1][0].GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, xerrors.New("left join error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a left join error", func() {
					Expect(err).To(MatchError(ContainSubstring("left join error")))
				})
			})

			Context("Error on on of join", func() {

				BeforeEach(func() {
					pegomock.When(join[len(join)-1][1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, xerrors.New("on join error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns an on join error", func() {
					Expect(err).To(MatchError(ContainSubstring("on join error")))
				})
			})

			Context("Where is nil", func() {

				BeforeEach(func() {
					where = nil
				})

				It("Returns a valid SQL string", func() {
					Expect(sql).To(MatchSQLString("SELECT path[0].sql, path[1].sql FROM from.sql LEFT JOIN join[0][0].sql ON join[0][1].sql RIGHT JOIN join[1][0].sql ON join[1][1].sql ORDER BY orderBy[0].sql ASC, orderBy[1].sql DESC"))
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("Where has error", func() {

				BeforeEach(func() {
					pegomock.When(where[len(where)-1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, xerrors.New("where error"))
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
					Expect(sql).To(MatchSQLString("SELECT path[0].sql, path[1].sql FROM from.sql LEFT JOIN join[0][0].sql ON join[0][1].sql RIGHT JOIN join[1][0].sql ON join[1][1].sql WHERE (where[0].sql AND where[1].sql)"))
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("Order By has error", func() {

				BeforeEach(func() {
					pegomock.When(orderBy[len(where)-1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, xerrors.New("order by error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a order by error", func() {
					Expect(err).To(MatchError(ContainSubstring("order by error")))
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
			pegomock.When(cols[0].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("cols[0].sql"), nil)
			pegomock.When(cols[1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("cols[1].sql"), nil)

			from = NewMockTable()
			pegomock.When(from.GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("from.sql"), nil)
			pegomock.When(from.GetColumns()).ThenReturn(cols)
		})

		JustBeforeEach(func() {
			q = query.SelectFrom(from)
		})

		Context("#GetSQL", func() {

			var (
				d core.Dialect

				sql core.SQL
				err error
			)

			BeforeEach(func() {
				d = dialect.Default{}
			})

			JustBeforeEach(func() {
				sql, err = q.GetSQL(d)
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
