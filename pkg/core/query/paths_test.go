package query_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/query"
)

var _ = Describe("Paths", func() {

	Context("ExpandTables", func() {

		var (
			paths []core.Expression

			result []core.Expression
		)

		BeforeEach(func() {
			col1 := NewMockColumn()
			col2 := NewMockColumn()
			table := NewMockTable()
			paths = []core.Expression{
				table,
			}
			pegomock.When(table.GetColumns()).ThenReturn([]core.Column{col1, col2})
		})

		JustBeforeEach(func() {
			result = query.ExpandTables(paths)
		})

		It("Returns 2 column expressions", func() {
			Expect(result).To(HaveLen(2))
		})

		Context("With `core.Expression`", func() {

			BeforeEach(func() {
				paths = []core.Expression{
					NewMockColumn(),
					NewMockExpression(),
				}
			})

			It("Returns two expressions", func() {
				Expect(result).To(HaveLen(2))
				ExpectWithOffset(0, result).To(ContainElement(paths[0]))
				ExpectWithOffset(1, result).To(ContainElement(paths[1]))
			})
		})
	})

	Context("CombinePathSQL", func() {

		var (
			d     core.Dialect
			paths []core.Expression

			sql core.SQL
			err error
		)

		BeforeEach(func() {
			d = NewMockDialect()
			paths = []core.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(paths[0].GetSQL(d)).ThenReturn(core.NewSQLf("exp 1 sql"), nil)
			pegomock.When(paths[1].GetSQL(d)).ThenReturn(core.NewSQLf("exp 2 sql"), nil)
		})

		JustBeforeEach(func() {
			sql, err = query.CombinePathSQL(d, paths)
		})

		It("Returns a combined SQL", func() {
			Expect(sql).To(MatchSQLString("exp 1 sql, exp 2 sql"))
		})

		It("Returns no error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		Context("With nil columns", func() {

			BeforeEach(func() {
				paths = nil
			})

			It("Returns an empty SQL", func() {
				Expect(sql).To(MatchSQLString(""))
			})

			It("Returns a nil error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("With an error on the second expression", func() {

			BeforeEach(func() {
				pegomock.When(paths[1].GetSQL(d)).ThenReturn(nil, errors.New("second exp error"))
			})

			It("Returns a nil SQL", func() {
				Expect(sql).To(BeNil())
			})

			It("Returns the second error", func() {
				Expect(err).To(MatchError(ContainSubstring("second exp error")))
			})
		})
	})

	Context("CombineSQL", func() {

		var (
			d     core.Dialect
			paths []core.Expression

			sql core.SQL
			err error
		)

		BeforeEach(func() {
			d = NewMockDialect()
			paths = []core.Expression{
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(paths[0].GetSQL(d)).ThenReturn(core.NewSQLf("exp 1 sql"), nil)
			pegomock.When(paths[1].GetSQL(d)).ThenReturn(core.NewSQLf("exp 2 sql"), nil)
		})

		JustBeforeEach(func() {
			sql, err = query.CombineSQL(d, paths)
		})

		It("Returns a combined SQL", func() {
			Expect(sql).To(MatchSQLString("exp 1 sql exp 2 sql"))
		})

		It("Returns no error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		Context("With nil columns", func() {

			BeforeEach(func() {
				paths = nil
			})

			It("Returns an empty SQL", func() {
				Expect(sql).To(MatchSQLString(""))
			})

			It("Returns a nil error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("With an embedded empty path", func() {

			BeforeEach(func() {
				paths[len(paths)-1] = nil
			})

			It("Returns a nil SQL", func() {
				Expect(sql).To(BeNil())
			})

			It("Returns a nil expression error", func() {
				Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "path entry[1]")))
			})
		})

		Context("With an error on the second expression", func() {

			BeforeEach(func() {
				pegomock.When(paths[1].GetSQL(d)).ThenReturn(nil, errors.New("second exp error"))
			})

			It("Returns a nil SQL", func() {
				Expect(sql).To(BeNil())
			})

			It("Returns the second SQL error", func() {
				Expect(err).To(MatchError(ContainSubstring("second exp error")))
			})
		})
	})
})
