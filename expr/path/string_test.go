package path_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/operator"
	"github.com/weworksandbox/lingo/expr/path"
)

var _ = Describe("String", func() {

	Context("NewStringWithAlias", func() {

		var (
			e     lingo.Table
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
			e    lingo.Table
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
				set   lingo.Set
			)

			BeforeEach(func() {
				value = "example text"
			})

			JustBeforeEach(func() {
				set = p.To(value)
			})

			It("Returns a valid `lingo.SetDialect`", func() {
				Expect(set).ToNot(BeNil())
				Expect(set).To(Equal(expr.NewSet(p, expr.NewValue(value))))
			})
		})

		Context("ToExpression", func() {

			var (
				value lingo.Expression
				set   lingo.Set
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				set = p.ToExpr(value)
			})

			It("Returns a valid `lingo.SetDialect`", func() {
				Expect(set).ToNot(BeNil())
				Expect(set).To(Equal(expr.NewSet(p, value)))
			})
		})

		Context("Eq", func() {

			var (
				value  string
				result lingo.Expression
			)

			BeforeEach(func() {
				value = "example text"
			})

			JustBeforeEach(func() {
				result = p.Eq(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Eq, expr.NewValue(value))))
			})
		})

		Context("EqPath", func() {

			var (
				value  lingo.Expression
				result lingo.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.EqPath(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Eq, value)))
			})
		})

		Context("NotEq", func() {

			var (
				value  string
				result lingo.Expression
			)

			BeforeEach(func() {
				value = "example text"
			})

			JustBeforeEach(func() {
				result = p.NotEq(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotEq, expr.NewValue(value))))
			})
		})

		Context("NotEqPath", func() {

			var (
				value  lingo.Expression
				result lingo.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.NotEqPath(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotEq, value)))
			})
		})

		Context("IsNull", func() {

			var (
				result lingo.Expression
			)

			JustBeforeEach(func() {
				result = p.IsNull()
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Null)))
			})
		})

		Context("IsNotNull", func() {

			var (
				result lingo.Expression
			)

			JustBeforeEach(func() {
				result = p.IsNotNull()
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotNull)))
			})
		})

		Context("Like", func() {

			var (
				value  string
				result lingo.Expression
			)

			BeforeEach(func() {
				value = "example text%"
			})

			JustBeforeEach(func() {
				result = p.Like(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Like, expr.NewValue(value))))
			})
		})

		Context("LikePath", func() {

			var (
				value  lingo.Expression
				result lingo.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.LikePath(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.Like, value)))
			})
		})

		Context("NotLike", func() {

			var (
				value  string
				result lingo.Expression
			)

			BeforeEach(func() {
				value = "example text%"
			})

			JustBeforeEach(func() {
				result = p.NotLike(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotLike, expr.NewValue(value))))
			})
		})

		Context("NotLikePath", func() {

			var (
				value  lingo.Expression
				result lingo.Expression
			)

			BeforeEach(func() {
				value = NewMockExpression()
			})

			JustBeforeEach(func() {
				result = p.NotLikePath(value)
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotLike, value)))
			})
		})

		Context("In", func() {

			var (
				value  []string
				result lingo.Expression
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

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.In, expr.NewValue(value))))
			})
		})

		Context("InPath", func() {

			var (
				value  []lingo.Expression
				result lingo.Expression
			)

			BeforeEach(func() {
				value = []lingo.Expression{
					NewMockExpression(),
					NewMockExpression(),
				}
			})

			JustBeforeEach(func() {
				result = p.InPaths(value[0], value[1])
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.In, value[0], value[1])))
			})
		})

		Context("NotIn", func() {

			var (
				value  []string
				result lingo.Expression
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

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotIn, expr.NewValue(value))))
			})
		})

		Context("NotInPath", func() {

			var (
				value  []lingo.Expression
				result lingo.Expression
			)

			BeforeEach(func() {
				value = []lingo.Expression{
					NewMockExpression(),
					NewMockExpression(),
				}
			})

			JustBeforeEach(func() {
				result = p.NotInPaths(value[0], value[1])
			})

			It("Returns a valid `lingo.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(operator.NewOperator(p, operator.NotIn, value[0], value[1])))
			})
		})
	})
})
