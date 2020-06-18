package expression_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/expression/matchers"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var _ = Describe("Set", func() {

	Context("Calling `NewSet`", func() {

		var (
			left  core.Expression
			value core.Expression

			set core.Set
		)

		BeforeEach(func() {
			left = NewMockExpression()
			value = NewMockExpression()
		})

		JustBeforeEach(func() {
			set = expression.NewSet(left, value)
		})

		It("Returns a `core.Set`", func() {
			Expect(set).ToNot(BeNil())
		})

		Context("Calling `ToSQL`", func() {

			var (
				d core.Dialect

				s   sql.Data
				err error
			)

			BeforeEach(func() {
				d = setDialectSuccess{}
				pegomock.When(left.ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("left sql"), nil)
				pegomock.When(value.ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("value sql"), nil)
			})

			JustBeforeEach(func() {
				s, err = set.ToSQL(d)
			})

			It("Returns Set SQL string", func() {
				Expect(s).ToNot(BeNil())
				Expect(s).To(MatchSQLString("set sql"))
			})

			It("Returns no errors", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			Context("Dialect does not support Set", func() {

				BeforeEach(func() {
					d = NewMockDialect()
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(EqString("dialect function '%s' not supported", "Set")))
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
					Expect(err).To(MatchError(EqString("expression '%s' cannot be nil", "left")))
				})
			})

			Context("left returns an error", func() {

				BeforeEach(func() {
					pegomock.When(left.ToSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("left error"))
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError("left error"))
				})
			})

			Context("value is nil", func() {

				BeforeEach(func() {
					value = nil
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "value")))
				})
			})

			Context("value returns an error", func() {

				BeforeEach(func() {
					pegomock.When(value.ToSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("value error"))
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("value error")))
				})
			})

			Context("`Set` fails", func() {

				BeforeEach(func() {
					d = setDialectFailure{}
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError("set error"))
				})
			})
		})
	})
})

type setDialectSuccess struct{}

func (setDialectSuccess) GetName() string { return "dialect name" }
func (setDialectSuccess) Set(left, value sql.Data) (sql.Data, error) {
	return sql.String("set sql"), nil
}

type setDialectFailure struct{ setDialectSuccess }

func (setDialectFailure) Set(left, value sql.Data) (sql.Data, error) {
	return nil, errors.New("set error")
}
