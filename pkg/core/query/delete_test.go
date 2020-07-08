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
	"github.com/weworksandbox/lingo/pkg/core/query"
	"github.com/weworksandbox/lingo/pkg/core/query/matchers"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var _ = Describe("Delete", func() {

	Context("Delete", func() {

		var (
			from     core.Table
			joinOn   [][]core.Expression
			joinType []join.Type
			where    []core.Expression

			q *query.DeleteQuery
		)

		BeforeEach(func() {
			from = NewMockTable()
			pegomock.When(from.ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("from.sqlStr"), nil)

			joinOn = [][]core.Expression{
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
			pegomock.When(joinOn[0][0].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("joinOn[0][0].sqlStr"), nil)
			pegomock.When(joinOn[0][1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("joinOn[0][1].sqlStr"), nil)
			pegomock.When(joinOn[1][0].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("joinOn[1][0].sqlStr"), nil)
			pegomock.When(joinOn[1][1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("joinOn[1][1].sqlStr"), nil)

			where = []core.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(where[0].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("where[0].sqlStr"), nil)
			pegomock.When(where[1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("where[1].sqlStr"), nil)
		})

		JustBeforeEach(func() {
			q = query.Delete(from).Where(where...)
			for i, join := range joinOn {
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
				Expect(sql).To(MatchSQLString("DELETE FROM from.sqlStr LEFT JOIN joinOn[0][0].sqlStr ON joinOn[0][1].sqlStr RIGHT JOIN joinOn[1][0].sqlStr ON joinOn[1][1].sqlStr WHERE (where[0].sqlStr AND where[1].sqlStr)"))
			})

			It("Returns no error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			Context("From is nil", func() {

				BeforeEach(func() {
					from = nil
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a path entry is nil error", func() {
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "from")))
				})
			})

			Context("From ToSQL returns an error", func() {

				BeforeEach(func() {
					pegomock.When(from.ToSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("from error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the from error", func() {
					Expect(err).To(MatchError(ContainSubstring("from error")))
				})
			})

			Context("Join is nil/empty", func() {

				BeforeEach(func() {
					joinOn = nil
				})

				It("Returns a valid SQL string", func() {
					Expect(sql).To(MatchSQLString("DELETE FROM from.sqlStr WHERE (where[0].sqlStr AND where[1].sqlStr)"))
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("Left part of Join returns an error", func() {

				BeforeEach(func() {
					pegomock.When(joinOn[len(joinOn)-1][0].ToSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("left joinOn error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the left joinOn error", func() {
					Expect(err).To(MatchError(ContainSubstring("left joinOn error")))
				})
			})

			Context("On part of Join returns an error", func() {

				BeforeEach(func() {
					pegomock.When(joinOn[len(joinOn)-1][1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("on joinOn error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the on joinOn error", func() {
					Expect(err).To(MatchError(ContainSubstring("on joinOn error")))
				})
			})

			Context("Where is empty", func() {

				BeforeEach(func() {
					where = nil
				})

				It("Returns a valid SQL string", func() {
					Expect(sql).To(MatchSQLString("DELETE FROM from.sqlStr LEFT JOIN joinOn[0][0].sqlStr ON joinOn[0][1].sqlStr RIGHT JOIN joinOn[1][0].sqlStr ON joinOn[1][1].sqlStr"))
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("Where ToSQL returns an error", func() {

				BeforeEach(func() {
					pegomock.When(where[len(where)-1].ToSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("where error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the where error", func() {
					Expect(err).To(MatchError(ContainSubstring("where error")))
				})
			})
		})
	})
})
