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
	"github.com/weworksandbox/lingo/pkg/core/join"
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

var _ = Describe("JoinOn", func() {

	Context("Calling `NewJoinOn`", func() {

		var (
			left     core.Expression
			joinType join.Type
			on       core.Expression

			joinOn core.Expression
		)

		BeforeEach(func() {
			left = NewMockExpression()
			joinType = join.Outer
			on = NewMockExpression()
		})

		JustBeforeEach(func() {
			joinOn = expression.NewJoinOn(left, joinType, on)
		})

		It("Returns a `core.JoinOn`", func() {
			Expect(joinOn).ToNot(BeNil())
		})

		Context("`ToSQL`", func() {

			var (
				d core.Dialect

				s   sql.Data
				err error
			)

			BeforeEach(func() {
				d = joinerDialectSuccess{}
				pegomock.When(left.ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("left sql"), nil)
				pegomock.When(on.ToSQL(matchers.AnyCoreDialect())).ThenReturn(sql.String("on sql"), nil)
			})

			JustBeforeEach(func() {
				s, err = joinOn.ToSQL(d)
			})

			It("Returns Join SQL string", func() {
				Expect(s).ToNot(BeNil())
				Expect(s).To(MatchSQLString("joiner sql"))
			})

			It("Returns no errors", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			Context("Dialect does not support `Joiner`", func() {

				BeforeEach(func() {
					d = NewMockDialect()
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(EqString("dialect function '%s' not supported", "Joiner")))
				})
			})

			Context("on is nil", func() {

				BeforeEach(func() {
					on = nil
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "on")))
				})
			})

			Context("on returns an error", func() {

				BeforeEach(func() {
					pegomock.When(on.ToSQL(d)).ThenReturn(nil, errors.New("on error"))
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError("on error"))
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
					Expect(err).To(MatchError(ContainSubstring("expression '%s' cannot be nil", "left")))
				})
			})

			Context("left returns an error", func() {

				BeforeEach(func() {
					pegomock.When(left.ToSQL(d)).ThenReturn(nil, errors.New("left error"))
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError(ContainSubstring("left error")))
				})
			})

			Context("`Join` returns an error", func() {

				BeforeEach(func() {
					d = joinerDialectFailure{}
				})

				It("Returns no SQL", func() {
					Expect(s).To(BeNil())
				})

				It("Returns an error", func() {
					Expect(err).To(MatchError("joiner failure"))
				})
			})
		})
	})
})

type joinerDialectSuccess struct{}

func (joinerDialectSuccess) GetName() string { return "joiner success" }
func (joinerDialectSuccess) Join(sql.Data, join.Type, sql.Data) (sql.Data, error) {
	return sql.String("joiner sql"), nil
}

type joinerDialectFailure struct{ joinerDialectSuccess }

func (joinerDialectFailure) Join(sql.Data, join.Type, sql.Data) (sql.Data, error) {
	return nil, errors.New("joiner failure")
}
