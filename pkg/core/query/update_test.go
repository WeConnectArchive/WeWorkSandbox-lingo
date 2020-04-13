package query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/dialect"
	"github.com/weworksandbox/lingo/pkg/core/query"
	"github.com/weworksandbox/lingo/pkg/core/query/matchers"
	"errors"
)

var _ = Describe("Update", func() {

	Context("Update", func() {

		var (
			table core.Table
			where []core.Expression
			set   []core.Set

			q *query.UpdateQuery
		)

		BeforeEach(func() {
			table = NewMockTable()
			pegomock.When(table.GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("table.sql"), nil)

			where = []core.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(where[0].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("where[0].sql"), nil)
			pegomock.When(where[1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("where[1].sql"), nil)

			set = []core.Set{
				NewMockSet(),
				NewMockSet(),
			}
			pegomock.When(set[0].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("set[0].sql"), nil)
			pegomock.When(set[1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("set[1].sql"), nil)
		})

		JustBeforeEach(func() {
			q = query.Update(table).Where(where...).Set(set...)
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

			It("returns SQL", func() {
				Expect(sql).To(MatchSQLString("UPDATE table.sql SET set[0].sql, set[1].sql WHERE (where[0].sql AND where[1].sql)"))
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
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "table")))
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

			Context("Table GetSQL returns an error", func() {

				BeforeEach(func() {
					pegomock.When(table.GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("table error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the GetSQL error", func() {
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
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "set")))
				})
			})

			Context("Set returns an error", func() {

				BeforeEach(func() {
					pegomock.When(set[len(set)-1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("set error"))
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
					Expect(sql).To(MatchSQLString("UPDATE table.sql SET set[0].sql, set[1].sql"))
				})

				It("should not error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("Where returns an error", func() {

				BeforeEach(func() {
					pegomock.When(where[len(where)-1].GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("where error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns a where is empty error", func() {
					Expect(err).To(MatchError(ContainSubstring("where error")))
				})
			})
		})
	})
})
