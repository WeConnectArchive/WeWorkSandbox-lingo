package path_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/expression/operator"
	"github.com/weworksandbox/lingo/pkg/core/expression/path"
)

var _ = Describe("String", func() {

	Context("NewStringWithAlias", func() {

		var (
			e     core.Table
			name  string
			alias string

			p path.String
		)

		BeforeEach(func() {
			e = NewMockTable()
			name = "name"
			alias = "alias"
		})

		JustBeforeEach(func() {
			p = path.NewStringWithAlias(e, name, alias)
		})

		It("Returns a `String`", func() {
			Expect(p).To(BeAssignableToTypeOf(path.String{}))
		})

		It("Has the given parent table", func() {
			Expect(p.GetParent()).To(Equal(e))
		})

		It("Has the given name", func() {
			Expect(p.GetName()).To(Equal(name))
		})

		It("Has the given alias", func() {
			Expect(p.GetAlias()).To(Equal(alias))
		})
	})

	Context("NewString", func() {

		var (
			e    core.Table
			name string

			p path.String
		)

		BeforeEach(func() {
			e = NewMockTable()
			name = "name"
		})

		JustBeforeEach(func() {
			p = path.NewString(e, name)
		})

		It("Returns a `String`", func() {
			Expect(p).To(BeAssignableToTypeOf(path.String{}))
		})

		It("Has the given parent table", func() {
			Expect(p.GetParent()).To(Equal(e))
		})

		It("Has the given name", func() {
			Expect(p.GetName()).To(Equal(name))
		})

		It("Has a blank alias", func() {
			Expect(p.GetAlias()).To(BeEmpty())
		})

		Context("As", func() {

			var (
				alias string
			)

			BeforeEach(func() {
				alias = "new_name"
			})

			JustBeforeEach(func() {
				p = p.As(alias)
			})

			It("Adds the alias", func() {
				Expect(p.GetAlias()).To(Equal(alias))
			})
		})

		Context("To", func() {

			var (
				value string
				set   core.Set
			)

			BeforeEach(func() {
				value = "example text"
			})

			JustBeforeEach(func() {
				set = p.To(value)
			})

			It("Returns a valid `core.SetDialect`", func() {
				Expect(set).ToNot(BeNil())
				Expect(set).To(Equal(expression.NewSet(p, expression.NewValue(value))))
			})
		})

		Context("ToExpression", func() {

			var (
				value core.Expression
				set   core.Set
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				set = p.ToExpression(value)
			})

			It("Returns a valid `core.SetDialect`", func() {
				Expect(set).ToNot(BeNil())
				Expect(set).To(Equal(expression.NewSet(p, value)))
			})
		})

		Context("Eq", func() {

			var (
				value  string
				result core.Expression
			)

			BeforeEach(func() {
				value = "example text"
			})

			JustBeforeEach(func() {
				result = p.Eq(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Eq, expression.NewValue(value))))
			})
		})

		Context("EqPath", func() {

			var (
				value  core.Expression
				result core.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.EqPath(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Eq, value)))
			})
		})

		Context("NotEq", func() {

			var (
				value  string
				result core.Expression
			)

			BeforeEach(func() {
				value = "example text"
			})

			JustBeforeEach(func() {
				result = p.NotEq(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotEq, expression.NewValue(value))))
			})
		})

		Context("NotEqPath", func() {

			var (
				value  core.Expression
				result core.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.NotEqPath(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotEq, value)))
			})
		})

		Context("IsNull", func() {

			var (
				result core.Expression
			)

			JustBeforeEach(func() {
				result = p.IsNull()
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Null)))
			})
		})

		Context("IsNotNull", func() {

			var (
				result core.Expression
			)

			JustBeforeEach(func() {
				result = p.IsNotNull()
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotNull)))
			})
		})

		Context("Like", func() {

			var (
				value  string
				result core.Expression
			)

			BeforeEach(func() {
				value = "example text%"
			})

			JustBeforeEach(func() {
				result = p.Like(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Like, expression.NewValue(value))))
			})
		})

		Context("LikePath", func() {

			var (
				value  core.Expression
				result core.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.LikePath(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Like, value)))
			})
		})

		Context("NotLike", func() {

			var (
				value  string
				result core.Expression
			)

			BeforeEach(func() {
				value = "example text%"
			})

			JustBeforeEach(func() {
				result = p.NotLike(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotLike, expression.NewValue(value))))
			})
		})

		Context("NotLikePath", func() {

			var (
				value  core.Expression
				result core.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.NotLikePath(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotLike, value)))
			})
		})

		Context("In", func() {

			var (
				value  []string
				result core.Expression
			)

			BeforeEach(func() {
				value = []string{
					"example text",
					"second example text",
				}
			})

			JustBeforeEach(func() {
				// Doing it this way to make it more apparent we are passing in multiple params
				result = p.In(value[0], value[1])
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.In, expression.NewValue(value))))
			})
		})

		Context("InPath", func() {

			var (
				value  []core.Expression
				result core.Expression
			)

			BeforeEach(func() {
				value = []core.Expression{
					NewMockExpression(),
					NewMockExpression(),
				}
			})

			JustBeforeEach(func() {
				result = p.InPaths(value[0], value[1])
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.In, value[0], value[1])))
			})
		})

		Context("NotIn", func() {

			var (
				value  []string
				result core.Expression
			)

			BeforeEach(func() {
				value = []string{
					"example text",
					"second example text",
				}
			})

			JustBeforeEach(func() {
				// Doing it this way to make it more apparent we are passing in multiple params
				result = p.NotIn(value[0], value[1])
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotIn, expression.NewValue(value))))
			})
		})

		Context("NotInPath", func() {

			var (
				value  []core.Expression
				result core.Expression
			)

			BeforeEach(func() {
				value = []core.Expression{
					NewMockExpression(),
					NewMockExpression(),
				}
			})

			JustBeforeEach(func() {
				result = p.NotInPaths(value[0], value[1])
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotIn, value[0], value[1])))
			})
		})
	})
})
