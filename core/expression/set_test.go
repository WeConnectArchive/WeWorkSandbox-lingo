package expression_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/core"
	"github.com/weworksandbox/lingo/core/expression"
	"github.com/weworksandbox/lingo/core/expression/matchers"
	"errors"
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

		Context("Calling `GetSQL`", func() {

			var (
				d core.Dialect

				sql core.SQL
				err error
			)

			BeforeEach(func() {
				d = setDialectSuccess{}
				pegomock.When(left.GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("left sql"), nil)
				pegomock.When(value.GetSQL(matchers.AnyCoreDialect())).ThenReturn(core.NewSQLf("value sql"), nil)
			})

			JustBeforeEach(func() {
				sql, err = set.GetSQL(d)
			})

			It("Returns Set SQL string", func() {
				Expect(sql).ToNot(BeNil())
				Expect(sql).To(MatchSQLString("set sql"))
			})

			It("Returns no errors", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			Context("Dialect does not support Set", func() {

				BeforeEach(func() {
					d = NewMockDialect()
				})

				It("Returns no SQL", func() {
					Expect(sql).To(BeNil())
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
					Expect(sql).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(EqString("expression '%s' cannot be nil", "left")))
				})
			})

			Context("left returns an error", func() {

				BeforeEach(func() {
					pegomock.When(left.GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("left error"))
				})

				It("Returns no SQL", func() {
					Expect(sql).To(BeNil())
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
					Expect(sql).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "value")))
				})
			})

			Context("value returns an error", func() {

				BeforeEach(func() {
					pegomock.When(value.GetSQL(matchers.AnyCoreDialect())).ThenReturn(nil, errors.New("value error"))
				})

				It("Returns no SQL", func() {
					Expect(sql).To(BeNil())
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
					Expect(sql).To(BeNil())
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
func (setDialectSuccess) Set(left, value core.SQL) (core.SQL, error) {
	return core.NewSQLf("set sql"), nil
}

type setDialectFailure struct{ setDialectSuccess }

func (setDialectFailure) Set(left, value core.SQL) (core.SQL, error) {
	return nil, errors.New("set error")
}
