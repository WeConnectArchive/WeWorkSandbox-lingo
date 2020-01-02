package query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/dialect"
	"github.com/weworksandbox/lingo/core/expression"
	"github.com/weworksandbox/lingo/core/query"
	"github.com/weworksandbox/lingo/core/query/matchers"
	"errors"
)

var _ = Describe("Delete", func() {

	Context("Delete", func() {

		var (
			from     core.Table
			join     [][]core.Expression
			joinType []expression.JoinType
			where    []core.Expression

			q *query.DeleteQuery
		)

		BeforeEach(func() {
			from = NewMockTable()
			pegomock.When(from.GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("from.sql"), nil)

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

			where = []core.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(where[0].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("where[0].sql"), nil)
			pegomock.When(where[1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("where[1].sql"), nil)
		})

		JustBeforeEach(func() {
			q = query.Delete(from).Where(where...)
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
				Expect(sql).To(MatchSQLString("DELETE FROM from.sql LEFT JOIN join[0][0].sql ON join[0][1].sql RIGHT JOIN join[1][0].sql ON join[1][1].sql WHERE (where[0].sql AND where[1].sql)"))
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
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "path entry")))
				})
			})

			Context("From GetSQL returns an error", func() {

				BeforeEach(func() {
					pegomock.When(from.GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("from error"))
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
					join = nil
				})

				It("Returns a valid SQL string", func() {
					Expect(sql).To(MatchSQLString("DELETE FROM from.sql WHERE (where[0].sql AND where[1].sql)"))
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("Left part of Join returns an error", func() {

				BeforeEach(func() {
					pegomock.When(join[len(join)-1][0].GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("left join error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the left join error", func() {
					Expect(err).To(MatchError(ContainSubstring("left join error")))
				})
			})

			Context("On part of Join returns an error", func() {

				BeforeEach(func() {
					pegomock.When(join[len(join)-1][1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("on join error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the on join error", func() {
					Expect(err).To(MatchError(ContainSubstring("on join error")))
				})
			})

			Context("Where is empty", func() {

				BeforeEach(func() {
					where = nil
				})

				It("Returns a valid SQL string", func() {
					Expect(sql).To(MatchSQLString("DELETE FROM from.sql LEFT JOIN join[0][0].sql ON join[0][1].sql RIGHT JOIN join[1][0].sql ON join[1][1].sql"))
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("Where GetSQL returns an error", func() {

				BeforeEach(func() {
					pegomock.When(where[len(where)-1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("where error"))
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
