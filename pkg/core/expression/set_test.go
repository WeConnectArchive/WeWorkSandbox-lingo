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

var _ = Describe("SetDialect", func() {

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

		It("Returns a `core.SetDialect`", func() {
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

			It("Returns SetDialect SQL string", func() {
				Expect(s).ToNot(BeNil())
				Expect(s).To(MatchSQLString("Set sql"))
			})

			It("Returns no errors", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			Context("Dialect does not support SetDialect", func() {

				BeforeEach(func() {
					d = NewMockDialect()
					pegomock.When(d.GetName()).ThenReturn("mock")
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(EqString("dialect '%s' does not support '%s'", "mock", "expression.SetDialect")))
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
					Expect(err).To(MatchError(EqString("left of 'set' cannot be empty")))
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
					Expect(err).To(MatchError("set 'value' cannot be empty"))
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
					Expect(err).To(MatchError("value error"))
				})
			})

			Context("`SetDialect` fails", func() {

				BeforeEach(func() {
					d = setDialectFailure{}
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError("Set error"))
				})
			})
		})
	})
})

type setDialectSuccess struct{}

func (setDialectSuccess) GetName() string { return "dialect name" }
func (setDialectSuccess) Set(_, _ sql.Data) (sql.Data, error) {
	return sql.String("Set sql"), nil
}

type setDialectFailure struct{ setDialectSuccess }

func (setDialectFailure) Set(_, _ sql.Data) (sql.Data, error) {
	return nil, errors.New("Set error")
}
