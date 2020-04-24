package path_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/operator"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

var _ = Describe("Bool", func() {

	Context("NewBoolPathWithAlias", func() {

		var (
			e     core.Table
			name  string
			alias string

			p path.Bool
		)

		BeforeEach(func() {
			e = NewMockTable()
			name = "name"
			alias = "alias"
		})

		JustBeforeEach(func() {
			p = path.NewBoolPathWithAlias(e, name, alias)
		})

		It("Returns a `Bool`", func() {
			Expect(p).To(BeAssignableToTypeOf(path.Bool{}))
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

	Context("NewBoolPath", func() {

		var (
			e    core.Table
			name string

			p path.Bool
		)

		BeforeEach(func() {
			e = NewMockTable()
			name = "name"
		})

		JustBeforeEach(func() {
			p = path.NewBoolPath(e, name)
		})

		It("Returns a `Bool`", func() {
			Expect(p).To(BeAssignableToTypeOf(path.Bool{}))
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
				value bool
				set   core.Set
			)

			BeforeEach(func() {
				value = true
			})

			JustBeforeEach(func() {
				set = p.To(value)
			})

			It("Returns a valid `core.Set`", func() {
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

			It("Returns a valid `core.Set`", func() {
				Expect(set).ToNot(BeNil())
				Expect(set).To(Equal(expression.NewSet(p, value)))
			})
		})

		Context("Eq", func() {

			var (
				value  bool
				result core.Expression
			)

			BeforeEach(func() {
				value = true
			})

			JustBeforeEach(func() {
				result = p.Eq(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(expression.NewOperator(p, operator.Eq, expression.NewValue(value))))
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
				Expect(result).To(Equal(expression.NewOperator(p, operator.Eq, value)))
			})
		})

		Context("NotEq", func() {

			var (
				value  bool
				result core.Expression
			)

			BeforeEach(func() {
				value = true
			})

			JustBeforeEach(func() {
				result = p.NotEq(value)
			})

			It("Returns a valid `core.ComboExpression`", func() {
				Expect(result).ToNot(BeNil())
				Expect(result).To(Equal(expression.NewOperator(p, operator.NotEq, expression.NewValue(value))))
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
				Expect(result).To(Equal(expression.NewOperator(p, operator.NotEq, value)))
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
				Expect(result).To(Equal(expression.NewOperator(p, operator.Null)))
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
				Expect(result).To(Equal(expression.NewOperator(p, operator.NotNull)))
			})
		})
	})
})
