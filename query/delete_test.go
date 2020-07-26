package query_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/dialect"
	"github.com/weworksandbox/lingo/expr/join"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/query"
	"github.com/weworksandbox/lingo/query/matchers"
	"github.com/weworksandbox/lingo/sql"
)

var _ = Describe("Delete", func() {

	Context("Delete", func() {

		var (
			from     lingo.Table
			joinOn   [][]lingo.Expression
			joinType []join.Type
			where    []lingo.Expression

			q *query.DeleteQuery
		)

		BeforeEach(func() {
			from = NewMockTable()
			pegomock.When(from.ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("from.sqlStr"), nil)

			joinOn = [][]lingo.Expression{
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
			pegomock.When(joinOn[0][0].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("joinOn[0][0].sqlStr"), nil)
			pegomock.When(joinOn[0][1].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("joinOn[0][1].sqlStr"), nil)
			pegomock.When(joinOn[1][0].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("joinOn[1][0].sqlStr"), nil)
			pegomock.When(joinOn[1][1].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("joinOn[1][1].sqlStr"), nil)

			where = []lingo.Expression{
				NewMockExpression(),
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(where[0].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("whereSQL[0].sqlStr"), nil)
			pegomock.When(where[1].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("whereSQL[1].sqlStr"), nil)
			pegomock.When(where[2].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("whereSQL[2].sqlStr"), nil)
		})

		JustBeforeEach(func() {
			q = query.Delete(from).Where(where...)
			for i, join := range joinOn {
				q = q.Join(join[0], joinType[i], join[1])
			}
		})

		Context("#ToSQL", func() {

			var (
				d lingo.Dialect

				sql sql.Data
				err error
			)

			BeforeEach(func() {
				d, err = dialect.NewDialect()
				Expect(err).ToNot(HaveOccurred())
			})

			JustBeforeEach(func() {
				sql, err = q.ToSQL(d)
			})

			It("Returns a valid SQL string", func() {
				Expect(sql).To(MatchSQLString("DELETE FROM from.sqlStr LEFT JOIN joinOn[0][0].sqlStr ON joinOn[0][1].sqlStr RIGHT JOIN joinOn[1][0].sqlStr ON joinOn[1][1].sqlStr WHERE whereSQL[0].sqlStr AND whereSQL[1].sqlStr AND whereSQL[2].sqlStr"))
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
					Expect(err).To(MatchError(ContainSubstring("from cannot be empty")))
				})
			})

			Context("From ToSQL returns an error", func() {

				BeforeEach(func() {
					pegomock.When(from.ToSQL(matchers.AnyLingoDialect())).ThenReturn(nil, errors.New("from error"))
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
					Expect(sql).To(MatchSQLString("DELETE FROM from.sqlStr WHERE whereSQL[0].sqlStr AND whereSQL[1].sqlStr AND whereSQL[2].sqlStr"))
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("Left part of Join returns an error", func() {

				BeforeEach(func() {
					pegomock.When(joinOn[len(joinOn)-1][0].ToSQL(matchers.AnyLingoDialect())).ThenReturn(nil, errors.New("left joinOn error"))
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
					pegomock.When(joinOn[len(joinOn)-1][1].ToSQL(matchers.AnyLingoDialect())).ThenReturn(nil, errors.New("on joinOn error"))
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
					pegomock.When(where[len(where)-1].ToSQL(matchers.AnyLingoDialect())).ThenReturn(nil, errors.New("whereSQL error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the whereSQL error", func() {
					Expect(err).To(MatchError(ContainSubstring("whereSQL error")))
				})
			})
		})
	})
})
