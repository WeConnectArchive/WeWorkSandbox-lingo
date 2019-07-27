package query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/operator"
	"github.com/weworksandbox/lingo/pkg/core/query"
	"golang.org/x/xerrors"
)

var _ = Describe("where", func() {

	Context("#BuildWhereSQL", func() {

		var (
			d      core.Dialect
			values []core.Expression

			sql core.SQL
			err error
		)

		BeforeEach(func() {
			d = whereDialectSuccess{}
			values = []core.Expression{
				NewMockExpression(),
				NewMockExpression(),
				NewMockExpression(),
			}
			pegomock.When(values[0].GetSQL(d)).ThenReturn(core.NewSQLf("where 0 sql"), nil)
			pegomock.When(values[1].GetSQL(d)).ThenReturn(core.NewSQLf("where 1 sql"), nil)
			pegomock.When(values[2].GetSQL(d)).ThenReturn(core.NewSQLf("where 2 sql"), nil)
		})

		JustBeforeEach(func() {
			sql, err = query.BuildWhereSQL(d, values)
		})

		It("Combines all SQL with commas and `WHERE`", func() {
			Expect(sql).To(matchers.MatchSQLString("WHERE where 0 sql AND where 1 sql AND where 2 sql"))
		})

		It("Returns no error", func() {
			Expect(err).ShouldNot(HaveOccurred())
		})

		Context("With an error returning", func() {

			BeforeEach(func() {
				pegomock.When(values[2].GetSQL(d)).ThenReturn(nil, xerrors.New("last error"))
			})

			It("Returns a nil SQL", func() {
				Expect(sql).To(BeNil())
			})

			It("Returns the error", func() {
				Expect(err).To(MatchError(ContainSubstring("last error")))
			})
		})

		Context("With 2 values", func() {

			BeforeEach(func() {
				values = values[:len(values)-1]
			})

			It("Combines all SQL with commas and `WHERE`", func() {
				Expect(sql).To(matchers.MatchSQLString("WHERE where 0 sql AND where 1 sql"))
			})

			It("Returns no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("With 1 value", func() {

			BeforeEach(func() {
				values = values[:1]
			})

			It("Combines all SQL with commas and `WHERE`", func() {
				Expect(sql).To(matchers.MatchSQLString("WHERE where 0 sql"))
			})

			It("Returns no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})

			Context("With an error returning", func() {

				BeforeEach(func() {
					pegomock.When(values[0].GetSQL(d)).ThenReturn(nil, xerrors.New("last error"))
				})

				It("Returns a nil SQL", func() {
					Expect(sql).To(BeNil())
				})

				It("Returns the error", func() {
					Expect(err).To(MatchError(ContainSubstring("last error")))
				})
			})
		})

		Context("With 0 values", func() {

			BeforeEach(func() {
				values = []core.Expression{}
			})

			It("Combines all SQL with commas and `WHERE`", func() {
				Expect(sql).To(matchers.MatchSQLString(""))
			})

			It("Returns no error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})

type whereDialectSuccess struct{}

func (whereDialectSuccess) GetName() string { return "where dialect success" }
func (whereDialectSuccess) Operator(left core.SQL, op operator.Operand, values []core.SQL) (core.SQL, error) {
	var sql = left
	for _, value := range values {
		sql = sql.AppendStringWithSpace(op.String()).AppendSqlWithSpace(value)
	}
	return sql, nil
}
