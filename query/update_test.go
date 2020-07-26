package query_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/dialect"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/query"
	"github.com/weworksandbox/lingo/query/matchers"
	"github.com/weworksandbox/lingo/sql"
)

var _ = Describe("Update", func() {

	Context("Update", func() {

		var (
			table lingo.Table
			where []lingo.Expression
			set   []lingo.Set

			q *query.UpdateQuery
		)

		BeforeEach(func() {
			table = NewMockTable()
			pegomock.When(table.ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("table.sqlStr"), nil)

			where = []lingo.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(where[0].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("whereSQL[0].sqlStr"), nil)
			pegomock.When(where[1].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("whereSQL[1].sqlStr"), nil)

			set = []lingo.Set{
				NewMockSet(),
				NewMockSet(),
			}
			pegomock.When(set[0].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("set[0].sqlStr"), nil)
			pegomock.When(set[1].ToSQL(matchers.AnyLingoDialect())).ThenReturn(sql.String("set[1].sqlStr"), nil)
		})

		JustBeforeEach(func() {
			q = query.Update(table).Where(where...).Set(set...)
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

			It("returns SQL", func() {
				Expect(sql).To(MatchSQLString("UPDATE table.sqlStr SET set[0].sqlStr, set[1].sqlStr WHERE whereSQL[0].sqlStr AND whereSQL[1].sqlStr"))
			})

			It("should not error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})

			Context("With a nil Table", func() {

				BeforeEach(func() {
					table = nil
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a table is nil error", func() {
					Expect(err).To(MatchError(ContainSubstring("table cannot be empty")))
				})
			})

			Context("Table Alias is not nil", func() {

				BeforeEach(func() {
					pegomock.When(table.GetAlias()).ThenReturn("alias is here")
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns an alias must be unset error", func() {
					Expect(err).To(MatchError(ContainSubstring("table alias must be unset")))
				})
			})

			Context("Table ToSQL returns an error", func() {

				BeforeEach(func() {
					pegomock.When(table.ToSQL(matchers.AnyLingoDialect())).ThenReturn(nil, errors.New("table error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the ToSQL error", func() {
					Expect(err).To(MatchError(ContainSubstring("table error")))
				})
			})

			Context("Set columns are nil", func() {

				BeforeEach(func() {
					set = nil
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a set is empty error", func() {
					Expect(err).To(MatchError(ContainSubstring("set cannot be empty")))
				})
			})

			Context("Set returns an error", func() {

				BeforeEach(func() {
					pegomock.When(set[len(set)-1].ToSQL(matchers.AnyLingoDialect())).ThenReturn(nil, errors.New("set error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a set is empty error", func() {
					Expect(err).To(MatchError(ContainSubstring("set error")))
				})
			})

			Context("Where is empty", func() {

				BeforeEach(func() {
					where = nil
				})

				It("Returns a valid SQL", func() {
					Expect(sql).To(MatchSQLString("UPDATE table.sqlStr SET set[0].sqlStr, set[1].sqlStr"))
				})

				It("should not error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("Where returns an error", func() {

				BeforeEach(func() {
					pegomock.When(where[len(where)-1].ToSQL(matchers.AnyLingoDialect())).ThenReturn(nil, errors.New("whereSQL error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a whereSQL is empty error", func() {
					Expect(err).To(MatchError(ContainSubstring("whereSQL error")))
				})
			})
		})
	})
})
